<!-- web/src/components/Guild.svelte -->
<!-- Progression per boss with pull counts and kill dates, the roster's bests, and the
     guild's reports. Progression is the page's reason to exist, so it leads. -->
<script lang="ts">
  import { characterHref, parseGuildPath, rulesetLabel, splitUnitName, type CharacterPath } from '../lib/characters';
  import { classColorVar, formatAmount } from '../lib/report/format';
  import { fetchGuild, type GuildPage } from '../lib/rankings/api';

  let { path = null }: { path?: CharacterPath | null } = $props();

  const resolved = $derived(
    path ?? (typeof window === 'undefined' ? null : parseGuildPath(window.location.pathname)),
  );

  /**
   * An `<a>` is inline: its own box is only as tall as its text, not the `min-h-11` row it
   * sits in. Rankings.svelte's row links already carry this fix; the phone audit here
   * caught the same shape of miss on this page's encounter, roster and report links, so
   * every anchor that is its own tap target gets it too, not just the row around it.
   */
  const rowLink = 'inline-flex min-h-11 items-center';

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

  const killed = $derived((data?.progression ?? []).filter((row) => row.kills > 0).length);
  const killedAt = (row: { first_kill_at?: string }): string =>
    row.first_kill_at === undefined ? 'not killed' : row.first_kill_at.slice(0, 10);
  const pulls = $derived((data?.progression ?? []).reduce((total, row) => total + row.pull_count, 0));
</script>

{#if status === 'missing'}
  <p class="text-[14px]" data-testid="guild-missing">
    That is not a guild address. They look like <code class="font-mono">/guild/eu/normal/the-last-watch</code>.
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
        {rulesetLabel(resolved.ruleset)} {resolved.region.toUpperCase()} · {killed} bosses down · {pulls} pulls
      </p>
    </header>

    <section class="flex flex-col gap-2">
      <h2 class="section-title text-[18px]">Progression</h2>
      {#if data.progression.length === 0}
        <p class="text-muted text-[14px]" data-testid="guild-empty">No pulls recorded yet.</p>
      {:else}
        <ul class="flex flex-col" data-testid="guild-progression">
          {#each data.progression as row (row.encounter)}
            <li class="border-line-soft grid min-h-11 grid-cols-[minmax(0,1fr)_auto_auto] items-center gap-3 border-b px-2 py-2 text-[14px]">
              <a class="{rowLink} truncate" href={`/rankings/${row.encounter.toLowerCase().replace(/[^a-z0-9]+/g, '-')}`}>
                {row.encounter}
              </a>
              <span class="text-muted font-mono tabular text-right text-[13px]">{row.pull_count} pulls</span>
              <span class="font-mono tabular w-[104px] text-right" data-testid="guild-kill">
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
            <li class="border-line-soft grid min-h-11 grid-cols-[minmax(0,1fr)_auto] items-center gap-3 border-b px-2 py-2 text-[14px] md:grid-cols-[minmax(120px,1fr)_minmax(0,1fr)_96px]">
              <a
                class="{rowLink} truncate font-semibold"
                style={`color: ${classColorVar(row.player.class)}`}
                href={characterHref(resolved.region, resolved.ruleset, row.player.name)}
              >
                {splitUnitName(row.player.name).name}
              </a>
              <span class="text-muted truncate hidden text-[13px] md:inline">{row.encounter} · {row.player.spec}</span>
              <span class="font-mono tabular text-right">{formatAmount(Math.round(row.value))}</span>
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
            <li class="border-line-soft flex min-h-11 flex-wrap items-center gap-3 border-b px-2 py-2 text-[14px]">
              <a class={rowLink} href={`/reports/${report.id}`}>{report.title === '' ? report.zone : report.title}</a>
              <span class="text-muted font-mono tabular text-[13px]">{report.created_at.slice(0, 10)}</span>
            </li>
          {/each}
        </ul>
      </section>
    {/if}
  </div>
{/if}
