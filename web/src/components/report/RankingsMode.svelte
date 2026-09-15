<!-- web/src/components/report/RankingsMode.svelte -->
<!-- Where this kill sits, from the same endpoint the standalone rankings pages read. This
     report's own row is marked, because the question is always "where am I on this
     list" -- and a moderation state (contract Amendments: ok | at_risk | removed) is
     printed inline rather than silently filtered, the same way a live fight says "Live"
     instead of hiding until it closes. -->
<script lang="ts">
  import {
    characterHref,
    guildHref,
    parseCharacterKey,
    rulesetLabel,
    splitUnitName,
  } from '../../lib/characters';
  import { classColorVar, formatAmount, formatDuration } from '../../lib/report/format';
  import {
    fetchRankings,
    type RankingMetric,
    type RankingRow,
    type RankingsPage,
  } from '../../lib/rankings/api';
  import type { FightEntry } from '../../lib/report/types';

  let {
    fight,
    reportId,
    encounterSlug,
    spec = '',
    metric: metricParam = '',
    onPatch = () => {},
  }: {
    fight: FightEntry;
    reportId: string;
    encounterSlug: string;
    /** '' ranks every spec together; a spec name narrows the board to it. From the url, so a link keeps it. */
    spec?: string;
    /** The metric id from the url; '' means dps. */
    metric?: string;
    onPatch?: (patch: { rankingsSpec?: string; rankingsMetric?: string }) => void;
  } = $props();

  /** The picker's options and the word the value column is filed under: one list, so the
      label on a phone card cannot drift from the metric the visitor chose. */
  const METRICS: { id: RankingMetric; label: string }[] = [
    { id: 'dps', label: 'DPS' },
    { id: 'hps', label: 'HPS' },
    { id: 'damage_taken', label: 'Damage taken per second' },
  ];

  const metric = $derived<RankingMetric>(
    METRICS.some((entry) => entry.id === metricParam) ? (metricParam as RankingMetric) : 'dps',
  );
  let page = $state<RankingsPage | null>(null);
  let status = $state<'idle' | 'loading' | 'ready' | 'failed'>('idle');
  let error = $state('');

  const metricLabel = $derived(METRICS.find((option) => option.id === metric)?.label ?? 'Value');
  /** The specs on the board plus the chosen one, so a narrowed board still offers the rest. */
  let seenSpecs = $state<string[]>([]);
  const specs = $derived(
    [...new Set([...seenSpecs, ...(page?.rows ?? []).map((row) => row.player.spec)])].sort(),
  );
  $effect(() => {
    if (page === null || spec !== '') return;
    seenSpecs = [...new Set(page.rows.map((row) => row.player.spec))];
  });

  /**
   * A row's three figures with the words their columns never carried. Built here rather
   * than written out three times in the markup so the phone strip cannot drift from the
   * columns it stands in for -- the same shape SummaryTab.svelte's `figuresFor` uses.
   */
  function figuresFor(row: RankingRow): { label: string; value: string }[] {
    return [
      { label: 'Split', value: row.talent_split },
      { label: metricLabel, value: formatAmount(Math.round(row.value)) },
      { label: 'Duration', value: formatDuration(row.duration_ms) },
    ];
  }

  /**
   * A row's own region and ruleset, read out of `player.key` rather than the row's guild --
   * an unguilded row still fought in a real region and ruleset, and defaulting them (the
   * way `us`/`normal` would) links a real player at someone else's server. Null when
   * `player.key` does not parse, rather than a guess: `Rankings.svelte` carries the same
   * helper, and both call sites take the fix together since both carried the bug.
   */
  function characterRowHref(row: RankingRow): string | null {
    const key = parseCharacterKey(row.player.key);
    return key === null ? null : characterHref(key.region, key.ruleset, row.player.name);
  }

  /**
   * The metric picker is not the only thing that can change while a request is in
   * flight: the visitor can pick a different fight in the selector, which stays visible
   * in every mode. Two rankings requests can therefore be in flight at once and settle
   * in either order, the same hazard ReportView.svelte's own Effect 3 guards against on
   * both its success and its rejection path. `wantedFight` and `wantedMetric`, captured
   * before the request and re-checked against the live props/state after it, are that
   * guard here: an answer for a fight or a metric nobody has selected any more is
   * dropped rather than painted under the current selection's name, on both paths.
   */
  $effect(() => {
    const wantedMetric = metric;
    const wantedSpec = spec;
    const wantedFight = fight.index;
    const encounterId = fight.encounter_id;
    const difficulty = fight.difficulty;
    if (encounterId === undefined) {
      status = 'idle';
      return;
    }
    status = 'loading';
    void fetchRankings({
      encounter: encounterSlug,
      difficulty,
      metric: wantedMetric,
      spec: wantedSpec === '' ? undefined : wantedSpec,
      page: 1,
    })
      .then((result) => {
        if (fight.index !== wantedFight || metric !== wantedMetric || spec !== wantedSpec) return;
        page = result;
        status = 'ready';
      })
      .catch((thrown: unknown) => {
        if (fight.index !== wantedFight || metric !== wantedMetric || spec !== wantedSpec) return;
        status = 'failed';
        error = thrown instanceof Error ? thrown.message : 'Rankings did not load.';
      });
  });
</script>

<div class="flex flex-col gap-3" data-testid="rankings-mode">
  {#if fight.encounter_id === undefined}
    <p class="text-muted text-[14px]">Trash is not ranked. Pick an encounter to see its rankings.</p>
  {:else}
    <div class="flex flex-wrap items-center gap-x-4 gap-y-2">
      <label class="label text-muted flex items-center gap-2" for="rankings-metric">
        Metric
        <select
          id="rankings-metric"
          class="border-line-warm bg-raised rounded-control text-text h-11 px-2 text-[13px] md:h-9"
          value={metric}
          onchange={(event) => onPatch({ rankingsMetric: (event.currentTarget as HTMLSelectElement).value })}
          data-testid="rankings-metric"
        >
          {#each METRICS as option (option.id)}
            <option value={option.id}>{option.label}</option>
          {/each}
        </select>
      </label>
      <label class="label text-muted flex items-center gap-2" for="rankings-spec">
        Spec
        <select
          id="rankings-spec"
          class="border-line-warm bg-raised rounded-control text-text h-11 px-2 text-[13px] md:h-9"
          value={spec}
          onchange={(event) => onPatch({ rankingsSpec: (event.currentTarget as HTMLSelectElement).value })}
          data-testid="rankings-spec"
        >
          <option value="">Every spec</option>
          {#each specs as option (option)}
            <option value={option}>{option}</option>
          {/each}
        </select>
      </label>
      <!-- A target beside the metric picker, not a word in a sentence, so it carries the
           44px the design system asks of a phone. SummaryTab.svelte's build link is the
           same shape. -->
      <a class="inline-flex min-h-11 items-center text-[13px] md:min-h-0" href={`/rankings/${encounterSlug}`}>
        Full rankings for {fight.name}
      </a>
    </div>

    {#if status === 'loading'}
      <p class="text-muted text-[14px]">Loading rankings.</p>
    {:else if status === 'failed'}
      <p class="text-[14px]" role="alert">{error}</p>
    {:else if page !== null}
      <p class="text-muted text-[13px]">
        <span class="tabular font-mono">{page.total}</span> ranked {page.total === 1 ? 'kill' : 'kills'} of
        {fight.name} on this ruleset, every report counted · updated
        <span class="tabular font-mono">{page.updated_at.slice(0, 10)}</span>
        {#if !fight.kill}
          · <span class="text-wipe">this pull was a wipe</span>, and only kills are ranked
        {/if}
      </p>
      <div
        class="text-muted label hidden grid-cols-[40px_minmax(120px,1.4fr)_minmax(100px,1fr)_88px_96px_72px] gap-x-3 px-2 pb-1 md:grid"
      >
        <span>#</span>
        <span>Player</span>
        <span>Guild · spec</span>
        <span class="text-right" title="Talent points per tree">Split</span>
        <span class="text-right">{metricLabel}</span>
        <span class="text-right">Length</span>
      </div>
      <ul class="flex flex-col" data-testid="rankings-rows">
        {#each page.rows as row (`${row.report_id}-${row.fight_index}-${row.player.key}`)}
          {@const characterLinkHref = characterRowHref(row)}
          <li
            class="border-line-soft grid min-h-11 grid-cols-[40px_minmax(0,1fr)_auto] items-center gap-x-3 gap-y-1 border-b px-2 py-2 text-[14px] md:grid-cols-[40px_minmax(120px,1.4fr)_minmax(100px,1fr)_88px_96px_72px]"
            class:bg-card-top={row.report_id === reportId}
            data-testid={row.report_id === reportId ? 'rankings-mine' : 'rankings-row'}
          >
            <span class="text-muted tabular font-mono text-[12px]">{row.rank}</span>
            <span class="flex min-w-0 items-center gap-2">
              {#if characterLinkHref !== null}
                <a
                  class="truncate font-semibold"
                  style={`color: ${classColorVar(row.player.class)}`}
                  href={characterLinkHref}
                >
                  {splitUnitName(row.player.name).name}
                </a>
              {:else}
                <!-- player.key did not parse into a region and ruleset: an unlinked name
                     is the legible failure, not a guessed link to the wrong character. -->
                <span class="truncate font-semibold" style={`color: ${classColorVar(row.player.class)}`}>
                  {splitUnitName(row.player.name).name}
                </span>
              {/if}
              {#if row.report_id === reportId}
                <!-- The row's own background also marks this, but a colour alone is not a
                     safe way to say "this is the report you are reading": this text does
                     the same job in words. -->
                <span class="pill pill-site shrink-0">This report</span>
              {/if}
            </span>
            <!-- Its own line on a phone card. Sharing line one with the player leaves the
                 name a `minmax(0,1fr)` track against the guild's `auto` one, and a row
                 carrying the "This report" pill truncated "Baelgrim" to "B". -->
            <span class="text-muted col-span-full truncate text-[13px] md:col-auto">
              {#if row.guild}
                <a href={guildHref(row.guild.region, row.guild.ruleset, row.guild.name)}>{row.guild.name}</a>
                · {rulesetLabel(row.guild.ruleset)} ·
              {/if}
              {row.player.spec}
            </span>
            <span class="text-muted tabular hidden text-right font-mono text-[13px] md:inline"
              >{row.talent_split}</span
            >
            <span class="tabular hidden text-right font-mono md:inline"
              >{formatAmount(Math.round(row.value))}</span
            >
            <span class="text-muted tabular hidden text-right font-mono text-[13px] md:inline">
              {#if row.report_id === reportId}
                {formatDuration(row.duration_ms)}
              {:else}
                <a
                  class="underline-offset-2 hover:underline"
                  href={`/reports/${row.report_id}?fight=${row.fight_index}`}
                  title="Open this kill's log"
                  data-testid="ranking-open">{formatDuration(row.duration_ms)}</a
                >
              {/if}
            </span>

            <!-- The three figures above sit in unlabelled columns, and a phone card has
                 neither the columns nor the width to keep them side by side, so below `md`
                 they are replaced by a strip where each says what it is. The value takes
                 the word from the metric picker above, not a word of its own. -->
            <span
              class="text-muted label col-span-full flex flex-wrap gap-x-3 gap-y-1 md:hidden"
              data-testid="ranking-figures"
            >
              {#each figuresFor(row) as figure (figure.label)}
                <span>{figure.label} <span class="tabular font-mono">{figure.value}</span></span>
              {/each}
            </span>
            {#if row.state !== 'ok'}
              <span class="pill pill-sample col-span-full md:col-auto">{row.state.replace('_', ' ')}</span>
            {/if}
          </li>
        {/each}
      </ul>
    {/if}
  {/if}
</div>
