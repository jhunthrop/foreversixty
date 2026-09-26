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
    contestClaim,
    fetchGuildHome,
    removeCharacter,
    type GuildHome,
    type GuildHomeReport,
    type GuildRosterRow,
  } from '../lib/guild/api';
  import { guildHomeCopy } from '../lib/guild/copy';
  import { SECONDARY_BUTTON_FIXED } from '../lib/planner/styles';
  import { classColorVar, formatAmount, rowLink } from '../lib/report/format';
  import { encounterSlug, fetchGuild, type GuildPage } from '../lib/rankings/api';
  import { RANKING_METRICS } from '../lib/rankings/url';
  import { executionLabel, executionTitle } from '../lib/sim/execution';
  import GuildRosterHandoff from './GuildRosterHandoff.svelte';
  import GuildStatus from './GuildStatus.svelte';
  import EmptyState from './ui/EmptyState.svelte';
  import { GUILD_LOADING } from '../lib/guild/layout';

  let { path = null }: { path?: CharacterPath | null } = $props();

  const resolved = $derived(
    path ?? (typeof window === 'undefined' ? null : parseGuildPath(window.location.pathname)),
  );

  let data = $state<GuildPage | null>(null);
  let status = $state<'loading' | 'ready' | 'failed' | 'missing'>('loading');
  let error = $state('');
  let attempt = $state(0);

  /**
   * Same reasoning as Character.svelte's effect: this page has no filter or tab state, so
   * `resolved` never changes after mount and this effect only ever fires once. The guard
   * is kept for the same reason -- cheap, and it is what keeps this correct if that stops
   * being true later, rather than leaning on today's absence of a second trigger.
   */
  $effect(() => {
    void attempt;
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
  let homeStatus = $state<'idle' | 'loading' | 'ready'>('idle');
  let rosterBusy = $state<string | null>(null); // character_key currently being approved/removed
  let rosterActionError = $state('');
  let contestBusy = $state(false);
  let contestError = $state('');
  let showContestConfirm = $state(false);

  let myGuildMembership = $state<{ rank: string; verified: boolean } | null>(null);
  let myCharacterKeys = $state<Set<string>>(new Set());
  let isModerator = $state(false);

  /**
   * A second, independent effect from the public page's own: it never blocks or delays the
   * public sections, and a failure here (network error, signed out, not a member) simply
   * leaves `home` null -- the same "fail toward the public-only view" rule fetchMe's own
   * callers (SessionNav, MyReports) already follow, so this component behaves identically
   * whether /v1/me is stubbed, refuses, or is not reachable at all (guild-phone.spec.ts
   * exercises exactly this last case today and must keep passing unmodified). There is no
   * visible error state for this section: a stranger to the guild is not owed an
   * explanation for why they see no member content, and a signed-in member's own error is
   * already the whole page's concern via `status`/`error` above, not this section's to
   * duplicate. So every rejection here -- network failure, signed out, not a member, or
   * the home fetch itself failing -- lands on the same `'idle'`, which renders nothing.
   */
  $effect(() => {
    const requested = resolved;
    if (requested === null) return;
    homeStatus = 'loading';
    void fetchMeOnce()
      .then((result) => {
        if (resolved !== requested) return;
        // A guild's real identity is (region, ruleset, name), not (region, ruleset) alone
        // -- many guilds share a region and ruleset. Matching on the slugified name too is
        // what keeps a signed-in member of one guild from being matched onto a DIFFERENT
        // guild's page that merely shares their region and ruleset, which would render
        // their own guild's private home panel (reports, full roster, officer controls)
        // underneath the other guild's public header.
        const membership = result?.guilds.find(
          (g) =>
            g.region === requested.region &&
            g.ruleset === requested.ruleset &&
            characterSlug(g.name) === requested.slug,
        );
        myCharacterKeys = new Set((result?.characters ?? []).map((c) => c.key));
        isModerator = result?.user.role === 'moderator' || result?.user.role === 'admin';
        if (membership === undefined) {
          homeStatus = 'idle';
          return null;
        }
        myGuildMembership = { rank: membership.rank ?? 'member', verified: membership.verified };
        return fetchGuildHome(membership.id).then((page) => {
          if (resolved !== requested) return;
          home = page;
          homeStatus = 'ready';
        });
      })
      .catch(() => {
        if (resolved !== requested) return;
        homeStatus = 'idle';
      });
  });

  const canManage = $derived(
    myGuildMembership !== null &&
      myGuildMembership.verified &&
      (myGuildMembership.rank === 'officer' || myGuildMembership.rank === 'leader'),
  );

  /** A raw membership row (any rank, verified or not) is enough to offer contesting --
   *  the API enforces the real officer/leader-or-rank-0 rule server side (spec section
   *  3.3's amendment); a plain member who tries gets a 403 with a sentence. */
  const canContest = $derived(myGuildMembership !== null);

  async function onApprove(row: GuildRosterRow): Promise<void> {
    if (home === null) return;
    rosterBusy = row.character_key;
    rosterActionError = '';
    try {
      await approveCharacter(home.guild.id, row.region, row.ruleset, characterSlug(row.name));
      home = {
        ...home,
        roster: home.roster.map((r) =>
          r.character_key === row.character_key ? { ...r, verified: true } : r,
        ),
      };
    } catch (thrown) {
      rosterActionError = thrown instanceof Error ? thrown.message : 'That did not work; try again';
    } finally {
      rosterBusy = null;
    }
  }

  async function onRemove(row: GuildRosterRow): Promise<void> {
    if (home === null) return;
    rosterBusy = row.character_key;
    rosterActionError = '';
    try {
      await removeCharacter(home.guild.id, row.region, row.ruleset, characterSlug(row.name));
      home = { ...home, roster: home.roster.filter((r) => r.character_key !== row.character_key) };
    } catch (thrown) {
      rosterActionError = thrown instanceof Error ? thrown.message : 'That did not work; try again';
    } finally {
      rosterBusy = null;
    }
  }

  /**
   * A contest always freezes officer tools now (a later security-review response
   * simplified the freeze rule to exactly `state === 'contested'`), so `home.claim.frozen`
   * is always true right after a successful contest -- but the contest response itself
   * carries no `claim` object, so the page still re-fetches home to pick up the fresh
   * claim/roster state (may_remove per row, etc.) rather than guessing it client-side.
   *
   * The re-fetch is wrapped separately from `contestClaim` itself: once `contestClaim`
   * resolves, the contest is recorded server-side no matter what happens next, so a
   * failure of the follow-up `fetchGuildHome` must never be reported with the generic
   * "that did not work" message -- that would tell the viewer their contest failed when
   * it actually succeeded, inviting a retry that only hits a 409.
   */
  async function onContest(): Promise<void> {
    if (home === null) return;
    contestBusy = true;
    contestError = '';
    try {
      await contestClaim(home.guild.id);
      showContestConfirm = false;
    } catch (thrown) {
      contestError = thrown instanceof Error ? thrown.message : 'That did not work; try again';
      contestBusy = false;
      return;
    }
    try {
      home = await fetchGuildHome(home.guild.id);
    } catch {
      contestError = guildHomeCopy.contestRecordedRefreshFailed;
    } finally {
      contestBusy = false;
    }
  }

  function rowPath(row: GuildRosterRow): CharacterPath {
    return { region: row.region as Region, ruleset: row.ruleset as Ruleset, slug: characterSlug(row.name) };
  }

  /** "Am I the guild's only member" is counted by distinct account, not row count: a
   *  solo officer with two characters is still one account. `Me.characters[].key`
   *  exact-matches `character_key`, so any roster row not in that set belongs to
   *  someone else. */
  const soloRoster = $derived(
    home === null ? true : home.roster.every((row) => myCharacterKeys.has(row.character_key)),
  );

  const reportWipeCount = (report: GuildHomeReport): number =>
    Math.max(0, report.fight_count - report.kill_count);

  /** An empty title falls back to the zone (matching ReportView.svelte's own
   *  `title === '' ? zone : title` pattern), and only when both are empty does the row
   *  fall back to the honest "untitled" placeholder. */
  const reportLabel = (report: GuildHomeReport): string => {
    if (report.title !== '') return report.title;
    if (report.zone !== '') return report.zone;
    return guildHomeCopy.untitledReport;
  };

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
{:else if status === 'loading' || status === 'failed'}
  <GuildStatus
    status={status === 'loading' ? 'loading' : 'failed'}
    error={status === 'failed' ? error : ''}
    onRetry={() => (attempt += 1)}
    lines={GUILD_LOADING.home.lines}
    minHeight={GUILD_LOADING.home.minHeight}
    testid="guild"
  />
{:else if data !== null && resolved !== null}
  <div class="reveal flex flex-col gap-[22px] md:gap-8" data-testid="guild" id="guild">
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
          {#if canManage}
            <a
              class="text-[13px] font-semibold"
              href={guildSettingsHref(resolved.region, resolved.ruleset, home.guild.name)}
              data-testid="guild-settings-link"
            >
              {guildHomeCopy.settingsLink}
            </a>
          {:else if myGuildMembership !== null && !myGuildMembership.verified}
            <a
              class="text-[13px] font-semibold"
              href={guildClaimHref(resolved.region, resolved.ruleset, home.guild.name)}
              data-testid="guild-claim-link"
            >
              {guildHomeCopy.claimLink}
            </a>
          {/if}
        </div>

        {#if home.claim.frozen}
          <p class="text-[13px]" role="alert" data-testid="guild-home-frozen">{guildHomeCopy.frozenNotice}</p>
        {/if}

        {#if canContest && home.claim.state !== 'unclaimed' && home.claim.state !== 'contested'}
          <div class="flex flex-wrap items-center gap-3">
            <button
              class="{SECONDARY_BUTTON_FIXED} border-line-warm text-text px-3"
              onclick={() => {
                showContestConfirm = true;
                contestError = '';
              }}
              disabled={contestBusy}
              data-testid="guild-home-contest-button"
            >
              {guildHomeCopy.contestButton}
            </button>
          </div>
        {/if}

        {#if !showContestConfirm && contestError !== ''}
          <p class="text-[13px]" role="alert" data-testid="guild-home-contest-error">{contestError}</p>
        {/if}

        {#if showContestConfirm}
          <div
            class="border-line-soft flex flex-col gap-3 border p-4"
            data-testid="guild-home-contest-confirm"
          >
            <ul class="flex flex-col gap-1 text-[13px]">
              {#each guildHomeCopy.contestRules as rule (rule)}
                <li>{rule}</li>
              {/each}
            </ul>
            <div class="flex gap-3">
              <button
                class="{SECONDARY_BUTTON_FIXED} border-line-warm-strong text-strong px-3"
                onclick={() => void onContest()}
                disabled={contestBusy}
                data-testid="guild-home-contest-confirm-button"
              >
                {guildHomeCopy.contestConfirmButton}
              </button>
              <button
                class="{SECONDARY_BUTTON_FIXED} border-line-warm text-text px-3"
                onclick={() => {
                  showContestConfirm = false;
                  contestError = '';
                }}
                disabled={contestBusy}
              >
                {guildHomeCopy.cancel}
              </button>
            </div>
            {#if contestError !== ''}<p class="text-[13px]" role="alert">{contestError}</p>{/if}
          </div>
        {/if}

        {#if myGuildMembership !== null && !myGuildMembership.verified}
          <p class="text-muted text-[13px]" data-testid="guild-home-unverified-note">
            {guildHomeCopy.reportsUnverifiedNote}
          </p>
        {/if}
        {#if home.reports.length === 0}
          <EmptyState message={guildHomeCopy.noReports} testid="guild-home-empty-reports" />
        {:else}
          <ul class="flex flex-col" data-testid="guild-home-reports">
            {#each home.reports as report (report.id)}
              <li
                class="border-line-soft flex min-h-11 flex-wrap items-center gap-3 border-b px-2 py-2 text-[14px]"
              >
                <a class={rowLink} href={`/reports/${report.id}`}>
                  {reportLabel(report)}
                </a>
                <span class="text-muted tabular font-mono text-[13px]">
                  {guildHomeCopy.reportSummary(report.kill_count, reportWipeCount(report))}
                </span>
              </li>
            {/each}
          </ul>
        {/if}

        <h2 class="section-title text-[18px]">{guildHomeCopy.rosterHeading}</h2>
        {#if soloRoster}
          <EmptyState
            message={canManage ? guildHomeCopy.emptyRosterOfficer : guildHomeCopy.emptyRosterMember}
            action={canManage
              ? {
                  label: guildHomeCopy.manageInvite,
                  href: guildSettingsHref(resolved.region, resolved.ruleset, home.guild.name),
                }
              : undefined}
            testid="guild-home-empty-roster"
          />
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
                  <span class="text-muted tabular font-mono text-[13px]"
                    >{guildHomeCopy.itemLevelLabel} {row.item_level}</span
                  >
                {/if}
                {#if row.consent !== 'roster'}
                  <GuildRosterHandoff path={rowPath(row)} />
                {/if}
                {#if canManage && !row.verified}
                  <button
                    class="{SECONDARY_BUTTON_FIXED} border-line-warm text-text px-3"
                    onclick={() => void onApprove(row)}
                    disabled={rosterBusy === row.character_key || home.claim.frozen}
                    data-testid="guild-roster-approve"
                  >
                    {guildHomeCopy.approve}
                  </button>
                {/if}
                {#if row.may_remove}
                  <button
                    class="{SECONDARY_BUTTON_FIXED} border-line-warm text-text px-3"
                    onclick={() => void onRemove(row)}
                    disabled={rosterBusy === row.character_key ||
                      (home.claim.frozen && !myCharacterKeys.has(row.character_key) && !isModerator)}
                    data-testid="guild-roster-remove"
                  >
                    {guildHomeCopy.remove}
                  </button>
                {/if}
              </li>
            {/each}
          </ul>
        {/if}
        {#if rosterActionError !== ''}
          <p class="text-[13px]" role="alert" data-testid="guild-roster-action-error">{rosterActionError}</p>
        {/if}
      </section>
    {/if}

    <section class="flex flex-col gap-2">
      <h2 class="section-title text-[18px]">{guildHomeCopy.progressionHeading}</h2>
      {#if data.progression.length === 0}
        <EmptyState message={guildHomeCopy.noProgression} testid="guild-empty" />
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
