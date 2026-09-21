<!-- web/src/components/Guild.svelte -->
<!-- Progression per boss with pull counts and kill dates, the roster's bests, and the
     guild's reports. Progression is the page's reason to exist, so it leads. -->
<script lang="ts">
  import { fetchMeOnce } from '../lib/account/api';
  import {
    characterHref,
    characterSlug,
    guildClaimHref,
    guildSettingsHref,
    parseGuildPath,
    rulesetLabel,
    splitUnitName,
    type CharacterPath,
    type Region,
    type Ruleset,
  } from '../lib/characters';
  import {
    approveCharacter,
    fetchGuildHome,
    removeCharacter,
    type GuildHome,
    type GuildRosterRow,
  } from '../lib/guild/api';
  import { guildHomeCopy } from '../lib/guild/copy';
  import { SECONDARY_BUTTON_FIXED } from '../lib/planner/styles';
  import { classColorVar, formatAmount, rowLink } from '../lib/report/format';
  import { encounterSlug, fetchGuild, type GuildPage } from '../lib/rankings/api';
  import { RANKING_METRICS } from '../lib/rankings/url';
  import { executionLabel, executionTitle } from '../lib/sim/execution';
  import GuildRosterHandoff from './GuildRosterHandoff.svelte';

  let { path = null }: { path?: CharacterPath | null } = $props();

  const resolved = $derived(
    path ?? (typeof window === 'undefined' ? null : parseGuildPath(window.location.pathname)),
  );

  let data = $state<GuildPage | null>(null);
  let status = $state<'loading' | 'ready' | 'failed' | 'missing'>('loading');
  let error = $state('');

  /**
   * Same reasoning as Character.svelte's effect: this page has no filter or tab state, so
   * `resolved` never changes after mount and this effect only ever fires once. The guard
   * is kept for the same reason -- cheap, and it is what keeps this correct if that stops
   * being true later, rather than leaning on today's absence of a second trigger.
   */
  $effect(() => {
    const requested = resolved;
    if (requested === null) {
      status = 'missing';
      return;
    }
    status = 'loading';
    void fetchGuild(requested)
      .then((result) => {
        if (resolved !== requested) return;
        data = result;
        status = 'ready';
      })
      .catch((thrown: unknown) => {
        if (resolved !== requested) return;
        status = 'failed';
        error = thrown instanceof Error ? thrown.message : 'That guild did not load.';
      });
  });

  let home = $state<GuildHome | null>(null);
  let homeStatus = $state<'idle' | 'loading' | 'ready' | 'failed'>('idle');
  let rosterBusy = $state<string | null>(null); // character_key currently being approved/removed

  /**
   * A second, independent effect from the public page's own: it never blocks or delays the
   * public sections, and a failure here (network error, signed out, not a member) simply
   * leaves `home` null -- the same "fail toward the public-only view" rule fetchMe's own
   * callers (SessionNav, MyReports) already follow, so this component behaves identically
   * whether /v1/me is stubbed, refuses, or is not reachable at all (guild-phone.spec.ts
   * exercises exactly this last case today and must keep passing unmodified).
   */
  $effect(() => {
    const requested = resolved;
    if (requested === null) return;
    homeStatus = 'loading';
    void fetchMeOnce()
      .then((result) => {
        if (resolved !== requested) return;
        const membership = result?.guilds.find(
          (g) => g.region === requested.region && g.ruleset === requested.ruleset,
        );
        if (membership === undefined) {
          homeStatus = 'idle';
          return null;
        }
        return fetchGuildHome(membership.id).then((page) => {
          if (resolved !== requested) return;
          home = page;
          homeStatus = 'ready';
        });
      })
      .catch(() => {
        if (resolved !== requested) return;
        homeStatus = homeStatus === 'loading' ? 'failed' : homeStatus;
      });
  });

  async function onApprove(characterKey: string): Promise<void> {
    if (home === null) return;
    rosterBusy = characterKey;
    try {
      const updated = await approveCharacter(home.guild.id, characterKey);
      home = {
        ...home,
        roster: home.roster.map((row) => (row.character_key === characterKey ? updated : row)),
      };
    } finally {
      rosterBusy = null;
    }
  }

  async function onRemove(characterKey: string): Promise<void> {
    if (home === null) return;
    rosterBusy = characterKey;
    try {
      await removeCharacter(home.guild.id, characterKey);
      home = { ...home, roster: home.roster.filter((row) => row.character_key !== characterKey) };
    } finally {
      rosterBusy = null;
    }
  }

  function rowPath(row: GuildRosterRow): CharacterPath {
    return { region: row.region as Region, ruleset: row.ruleset as Ruleset, slug: characterSlug(row.name) };
  }

  const killed = $derived((data?.progression ?? []).filter((row) => row.kills > 0).length);
  const killedAt = (row: { first_kill_at?: string }): string =>
    row.first_kill_at === undefined ? 'not killed' : row.first_kill_at.slice(0, 10);
  /**
   * "not killed" already says what it is; a bare date does not -- read on its own, out of
   * a screen reader's per-row traversal, "2026-12-09" is not obviously the boss's first
   * kill date rather than a pull's date or the report's. This is the whole reason the date
   * carries its own label rather than only the visible column position.
   */
  const killedAtAriaLabel = (row: { first_kill_at?: string }): string =>
    row.first_kill_at === undefined ? 'not killed' : `first killed ${killedAt(row)}`;
  const pulls = $derived((data?.progression ?? []).reduce((total, row) => total + row.pull_count, 0));

  /**
   * A roster-best value has no column heading at any breakpoint to say which metric it
   * is. `RANKING_METRICS` already holds the one label table for `dps`/`hps`/
   * `damage_taken`; echoed as-is if the API returns an id this list has not heard of, the
   * same fallback shape `rulesetLabel` uses.
   */
  function metricLabel(id: string): string {
    return RANKING_METRICS.find((metric) => metric.id === id)?.label ?? id;
  }
</script>

{#if status === 'missing'}
  <p class="text-[14px]" data-testid="guild-missing">
    That is not a guild address. They look like <code class="font-mono">/guild/eu/normal/the-last-watch</code
    >.
  </p>
{:else if status === 'loading'}
  <p class="text-muted text-[14px]">Loading.</p>
{:else if status === 'failed'}
  <p class="text-[14px]" role="alert" data-testid="guild-error">{error}</p>
{:else if data !== null && resolved !== null}
  <div class="flex flex-col gap-[22px] md:gap-8" data-testid="guild" id="guild">
    <header class="flex flex-col gap-1">
      <h1 class="section-title text-[18px]">{data.guild.name}</h1>
      <p class="text-muted text-[13px]">
        {rulesetLabel(resolved.ruleset)}
        {resolved.region.toUpperCase()} ·
        <span class="tabular font-mono">{killed}</span> bosses down ·
        <span class="tabular font-mono">{pulls}</span> pulls
      </p>
    </header>

    {#if homeStatus === 'ready' && home !== null}
      <section class="border-line-soft flex flex-col gap-4 border-b pb-6" data-testid="guild-home">
        <div class="flex flex-wrap items-center justify-between gap-3">
          <h2 class="section-title text-[18px]">{guildHomeCopy.reportsHeading}</h2>
          {#if home.viewer.can_manage}
            <a
              class="text-[13px] font-semibold"
              href={guildSettingsHref(resolved.region, resolved.ruleset, home.guild.name)}
              data-testid="guild-settings-link"
            >
              {guildHomeCopy.settingsLink}
            </a>
          {:else if !home.viewer.verified}
            <a
              class="text-[13px] font-semibold"
              href={guildClaimHref(resolved.region, resolved.ruleset, home.guild.name)}
              data-testid="guild-claim-link"
            >
              {guildHomeCopy.claimLink}
            </a>
          {/if}
        </div>
        {#if home.reports.rows.length === 0}
          <p class="text-muted text-[14px]" data-testid="guild-home-empty-reports">
            {guildHomeCopy.noReports}
          </p>
        {:else}
          <ul class="flex flex-col" data-testid="guild-home-reports">
            {#each home.reports.rows as report (report.id)}
              <li
                class="border-line-soft flex min-h-11 flex-wrap items-center gap-3 border-b px-2 py-2 text-[14px]"
              >
                <a class={rowLink} href={`/reports/${report.id}`}
                  >{report.title === '' ? report.zone : report.title}</a
                >
                <span class="text-muted tabular font-mono text-[13px]">
                  {report.kill_count} kills · {report.wipe_count} wipes
                </span>
              </li>
            {/each}
          </ul>
        {/if}

        <h2 class="section-title text-[18px]">{guildHomeCopy.rosterHeading}</h2>
        {#if home.roster.length <= 1}
          <p class="text-muted text-[14px]" data-testid="guild-home-empty-roster">
            {home.viewer.can_manage ? guildHomeCopy.emptyRosterOfficer : guildHomeCopy.emptyRosterMember}
          </p>
        {:else}
          <ul class="flex flex-col" data-testid="guild-home-roster">
            {#each home.roster as row (row.character_key)}
              <li
                class="border-line-soft flex min-h-11 flex-wrap items-center gap-3 border-b px-2 py-2 text-[14px]"
              >
                <a
                  class="{rowLink} truncate font-semibold"
                  style={`color: ${classColorVar(row.class ?? '')}`}
                  href={characterHref(row.region, row.ruleset, row.name)}
                >
                  {row.name}
                </a>
                <span class="text-muted text-[13px]">{row.spec ?? ''}</span>
                {#if row.logged_recently}
                  <span class="pill pill-site text-[11px]">{guildHomeCopy.loggedRecently}</span>
                {/if}
                {#if !row.verified}
                  <span class="pill pill-site text-[11px]" data-testid="guild-roster-unverified">
                    {guildHomeCopy.unverified}
                  </span>
                {/if}
                {#if row.item_level !== undefined}
                  <span class="text-muted tabular font-mono text-[13px]">ilvl {row.item_level}</span>
                {/if}
                {#if row.consent !== 'roster'}
                  <GuildRosterHandoff path={rowPath(row)} />
                {/if}
                {#if home.viewer.can_manage}
                  {#if !row.verified}
                    <button
                      class="{SECONDARY_BUTTON_FIXED} border-line-warm text-text px-3"
                      onclick={() => void onApprove(row.character_key)}
                      disabled={rosterBusy === row.character_key}
                      data-testid="guild-roster-approve"
                    >
                      {guildHomeCopy.approve}
                    </button>
                  {/if}
                  <button
                    class="{SECONDARY_BUTTON_FIXED} border-line-warm text-text px-3"
                    onclick={() => void onRemove(row.character_key)}
                    disabled={rosterBusy === row.character_key}
                    data-testid="guild-roster-remove"
                  >
                    {guildHomeCopy.remove}
                  </button>
                {/if}
              </li>
            {/each}
          </ul>
        {/if}
      </section>
    {:else if homeStatus === 'failed'}
      <p class="text-[14px]" role="alert" data-testid="guild-home-error">{guildHomeCopy.failed}</p>
    {/if}

    <section class="flex flex-col gap-2">
      <h2 class="section-title text-[18px]">Progression</h2>
      {#if data.progression.length === 0}
        <p class="text-muted text-[14px]" data-testid="guild-empty">No pulls recorded yet.</p>
      {:else}
        <ul class="flex flex-col" data-testid="guild-progression">
          {#each data.progression as row (row.encounter)}
            <li
              class="border-line-soft grid min-h-11 grid-cols-[minmax(0,1fr)_auto_auto] items-center gap-3 border-b px-2 py-2 text-[14px]"
            >
              <a class="{rowLink} truncate" href={`/rankings/${encounterSlug(row.encounter)}`}>
                {row.encounter}
              </a>
              <span class="text-muted tabular text-right font-mono text-[13px]">{row.pull_count} pulls</span>
              <span
                class="tabular w-[104px] text-right font-mono"
                data-testid="guild-kill"
                aria-label={killedAtAriaLabel(row)}
              >
                {killedAt(row)}
              </span>
            </li>
          {/each}
        </ul>
      {/if}
    </section>

    {#if data.roster_best.length > 0}
      <section class="flex flex-col gap-2">
        <h2 class="section-title text-[18px]">Roster bests</h2>
        <ul class="flex flex-col" data-testid="guild-roster">
          {#each data.roster_best as row (`${row.player.key}-${row.encounter_id}-${row.metric}`)}
            <li
              class="border-line-soft grid min-h-11 grid-cols-[minmax(0,1fr)_auto] items-center gap-3 border-b px-2 py-2 text-[14px] md:grid-cols-[minmax(120px,1fr)_minmax(0,1fr)_96px_72px]"
            >
              <a
                class="{rowLink} truncate font-semibold"
                style={`color: ${classColorVar(row.player.class)}`}
                href={characterHref(resolved.region, resolved.ruleset, row.player.name)}
              >
                {splitUnitName(row.player.name).name}
              </a>
              <span class="text-muted hidden truncate text-[13px] md:inline"
                >{row.encounter} · {row.player.spec}</span
              >
              <span
                class="tabular text-right font-mono"
                aria-label={`${formatAmount(Math.round(row.value))} ${metricLabel(row.metric)}`}
              >
                {formatAmount(Math.round(row.value))}
              </span>
              <!-- `roster_best` rows have no report_id/fight_index -- they are per-encounter
                   aggregates, not one fight -- so this score is text, never a compare-mode link,
                   unlike the same score on the rankings and character rows. -->
              <span
                class="text-muted tabular hidden text-right font-mono text-[13px] md:inline"
                title={executionTitle(row.execution_score)}
                aria-label={executionTitle(row.execution_score)}
                data-testid="guild-execution">{executionLabel(row.execution_score)}</span
              >
            </li>
          {/each}
        </ul>
      </section>
    {/if}

    {#if data.reports.length > 0}
      <section class="flex flex-col gap-2">
        <h2 class="section-title text-[18px]">Reports</h2>
        <ul class="flex flex-col" data-testid="guild-reports">
          {#each data.reports as report (report.id)}
            <li
              class="border-line-soft flex min-h-11 flex-wrap items-center gap-3 border-b px-2 py-2 text-[14px]"
            >
              <a class={rowLink} href={`/reports/${report.id}`}
                >{report.title === '' ? report.zone : report.title}</a
              >
              <span class="text-muted tabular font-mono text-[13px]">{report.created_at.slice(0, 10)}</span>
            </li>
          {/each}
        </ul>
      </section>
    {/if}
  </div>
{/if}
