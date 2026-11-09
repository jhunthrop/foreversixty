<!-- web/src/components/report/TimeChart.svelte -->
<!-- The report's main chart (controls and charts spec §B5, D5; logs spec 2026-11-07 6.1, the
     mock library's LineChart): an open Section under the tab's Plate, "{Metric} over the fight"
     with the scope's figure as its meta, the Players control and the log scale at the head's
     right, then the legend and the shared TimeSeriesChart with its brush. The table's metric per
     player, averaged over the fight's bucket (lib/chart/buckets.ts), the top 4 and the bottom 4 of
     N plus the Source player and you, and the raid median. The Window choice lives in the report's
     control bar. Below 768 the chart folds into a Disclosure (spec 9, "Chart (phone Disclosure)"). -->
<script lang="ts">
  import { splitUnitName } from '../../lib/characters';
  import { togglePlayer } from '../../lib/chart/choice';
  import { chartCopy, type RateMetric } from '../../lib/chart/copy';
  import type { DeadSpan, DrawPreset, PlayersChoice } from '../../lib/chart/drawn';
  import { LegendState } from '../../lib/chart/legend-state.svelte';
  import { entityStyle } from '../../lib/chart/palette';
  import { buildRateModel, classMedianInfo, RAID_MEDIAN_ID, type RatePlayer } from '../../lib/chart/rate';
  import type { ChartEvent, ChartPhase, ChartSeries, PlayerEntry } from '../../lib/chart/types';
  import { formatRate } from '../../lib/report/frame';
  import type { ScopeNoun } from '../../lib/report/group';
  import type { Actor, RosterRow } from '../../lib/report/types';
  import { isFullWindow, type TimeWindow } from '../../lib/report/window';
  import ChartLegend from '../chart/ChartLegend.svelte';
  import LegendSheet from '../chart/LegendSheet.svelte';
  import PlayersControl from '../chart/PlayersControl.svelte';
  import TimeSeriesChart from '../chart/TimeSeriesChart.svelte';
  import { MediaQuery } from 'svelte/reactivity';
  import Disclosure from '../ui/Disclosure.svelte';
  import Section from '../ui/Section.svelte';
  import Switch from '../ui/Switch.svelte';

  let {
    metric,
    actors,
    roster,
    durationMs,
    window: current,
    phases = [],
    events = [],
    deadSpans = [],
    brushShape = [],
    onWindow,
    sourcePlayer,
    enemies = false,
    youGuid,
    chosen,
    preset = 'split',
    onChoose,
    onPreset,
    logScale,
    onLogScale,
    targetName,
    extra = null,
    syncKey,
    onOpenDeath,
    noun = 'raid',
    testid = 'time-chart',
  }: {
    metric: RateMetric;
    /** The table's actors over the whole fight, under the source scope. */
    actors: readonly Actor[];
    roster: readonly RosterRow[];
    durationMs: number;
    window: TimeWindow;
    phases?: readonly ChartPhase[];
    events?: readonly ChartEvent[];
    deadSpans?: readonly (DeadSpan & { guid: string })[];
    brushShape?: readonly number[];
    onWindow: (window: TimeWindow | null) => void;
    /** One player as the Source. */
    sourcePlayer?: string;
    /** The Source is the enemies: lines take the palette, not class colours. */
    enemies?: boolean;
    youGuid?: string;
    /** The Players control's own set, from the url; null keeps the preset. */
    chosen: readonly string[] | null;
    /** The default draw (fix round 2 ruling): top 4 and bottom 4, or Top 8, or Bottom 8. */
    preset?: DrawPreset;
    onChoose: (chosen: string[] | null) => void;
    /** A preset picked in the Players control: it replaces any own set. */
    onPreset: (preset: DrawPreset) => void;
    logScale: boolean;
    onLogScale: (on: boolean) => void;
    /** The filter bar's target, by name, for the head's meta. */
    targetName?: string;
    /** A charted ability: one more line in the palette's second colour. */
    extra?: { label: string; series: readonly number[] } | null;
    syncKey: string;
    onOpenDeath?: (deaths: readonly ChartEvent[]) => void;
    /** What the drawn players are (lib/report/group.ts): it names the head's figure and the median. */
    noun?: ScopeNoun;
    testid?: string;
  } = $props();

  /** The Section's title per metric: the parity region is its slug (section-damage-over-the-fight). */
  const SECTION_TITLE: Readonly<Record<RateMetric, string>> = {
    'damage-done': 'Damage over the fight',
    healing: 'Healing over the fight',
    'damage-taken': 'Damage taken over the fight',
  };
  const wide = new MediaQuery('min-width: 768px', true);

  /** The phone's open chart is remembered (it was before the report frame, 2026-09-26): a reader
   *  who opened it once finds it open on the next report. Read after mount, so the server's closed
   *  Disclosure hydrates as drawn; written only once read. */
  const CHART_OPEN_KEY = 'fs.report.chartOpen';
  let chartOpen = $state(false);
  let chartOpenRead = $state(false);
  $effect(() => {
    try {
      chartOpen = globalThis.localStorage?.getItem(CHART_OPEN_KEY) === 'true';
    } catch {
      // No storage (a private window, blocked site data): the chart starts closed.
    }
    chartOpenRead = true;
  });
  $effect(() => {
    const open = chartOpen;
    if (!chartOpenRead) return;
    try {
      globalThis.localStorage?.setItem(CHART_OPEN_KEY, String(open));
    } catch {
      // Nothing to remember it in: the choice lasts this page.
    }
  });

  const legend = new LegendState();
  let classMedian = $state<string | null>(null);
  let raidMedianOn = $state(true);

  const roles = $derived(new Map(roster.map((row) => [row.guid, row])));
  const players = $derived<RatePlayer[]>(
    actors
      .filter((actor) => enemies || roles.has(actor.guid))
      .map((actor, index) => {
        const row = roles.get(actor.guid);
        const role = row?.role;
        return {
          guid: actor.guid,
          name: splitUnitName(actor.name).name,
          className: row?.class ?? actor.class,
          role: role === 'tank' || role === 'healer' || role === 'dps' ? role : 'other',
          series: actor.series,
          colour: enemies ? entityStyle(index).colour : undefined,
        };
      }),
  );
  const youClass = $derived(
    players.find((player) => player.guid === (sourcePlayer ?? youGuid))?.className ?? null,
  );
  const spans = $derived.by(() => {
    // eslint-disable-next-line svelte/prefer-svelte-reactivity
    const out = new Map<string, DeadSpan[]>();
    for (const span of deadSpans) out.set(span.guid, [...(out.get(span.guid) ?? []), span]);
    return out;
  });
  const model = $derived(
    buildRateModel({
      metric,
      players,
      window: current,
      durationMs,
      pinned: new Set([sourcePlayer, youGuid].filter((guid): guid is string => guid !== undefined)),
      youGuid,
      chosen:
        sourcePlayer !== undefined && chosen === null ? new Set() : chosen === null ? null : new Set(chosen),
      preset,
      compare: { raid: raidMedianOn && !enemies, className: enemies ? null : classMedian },
      deadSpans: spans,
      single: sourcePlayer,
      noun,
    }),
  );
  const series = $derived<ChartSeries[]>([
    ...model.series,
    ...(extra === null
      ? []
      : [
          {
            id: 'extra',
            label: extra.label,
            colour: entityStyle(1).colour,
            kind: 'extra' as const,
            ...model.pointsOf(extra.series),
            figure: 0,
          },
        ]),
  ]);
  const medians = $derived(
    enemies
      ? []
      : [
          { ...model.medians[0], on: raidMedianOn },
          ...(youClass === null ? [] : [model.medians[1] ?? classMedianInfo(players, youClass, false)]),
        ],
  );
  const whole = $derived(isFullWindow(current, durationMs));
  const single = $derived(
    sourcePlayer === undefined ? undefined : model.entries.find((entry) => entry.id === sourcePlayer),
  );
  const phaseOfWindow = $derived(
    phases.find((phase) => current.startMs >= phase.startMs && current.endMs <= phase.endMs)?.name,
  );
  const meta = $derived.by(() => {
    if (model.raidFigure === 0 && players.length > 0)
      return chartCopy.emptyMeta(chartCopy.rateNoun(metric), phaseOfWindow, current.startMs, current.endMs);
    const span = chartCopy.span(current.startMs, current.endMs, whole);
    if (single !== undefined)
      return chartCopy.playerMeta(metric, single.label, formatRate(single.figure), span, targetName);
    return chartCopy.rateMeta(metric, formatRate(model.raidFigure), span, targetName, noun);
  });

  function choose(id: string): void {
    onChoose([...togglePlayer(chosen === null ? null : new Set(chosen), model.entries, id)]);
  }
  function toggleMedian(id: string): void {
    if (id === RAID_MEDIAN_ID) raidMedianOn = !raidMedianOn;
    else classMedian = classMedian === null ? youClass : null;
  }
  let coarse = $state(false);
  $effect(() => {
    const query = window.matchMedia('(pointer: coarse), (hover: none)');
    const read = (): void => {
      coarse = query.matches;
    };
    read();
    query.addEventListener('change', read);
    return () => query.removeEventListener('change', read);
  });
  const title = $derived(chartCopy.title[metric]);
  const choice = $derived<PlayersChoice>(chosen !== null ? 'custom' : preset);
  const PRESETS: readonly DrawPreset[] = ['split', 'top', 'bottom'];
  const notDrawn = $derived(model.entries.filter((entry: PlayerEntry) => !entry.drawn).length);
</script>

{#snippet controls()}
  <!-- On a phone the head wraps (fix round 2 ruling): the title on its row, then each field on
       its own full-width row, so every field shows its value whole. -->
  <div
    class="flex flex-wrap items-center justify-end gap-x-2 gap-y-2 max-md:grid max-md:w-full max-md:grid-cols-[64px_minmax(0,1fr)] max-md:justify-items-stretch max-md:text-left"
    data-testid="chart-head-controls"
  >
    {#if !enemies}
      <span class="text-small text-muted">{chartCopy.players}</span>
      <PlayersControl
        id="chart-players"
        entries={model.entries}
        {medians}
        {choice}
        presets={PRESETS}
        seen={noun === 'everyone'}
        formatValue={formatRate}
        onToggle={choose}
        onMedian={toggleMedian}
        {onPreset}
        onCoarse={() => (legend.sheetOpen = true)}
      />
    {/if}
    <span class="inline-flex items-center gap-2 max-md:col-span-2">
      <Switch checked={logScale} label={chartCopy.logScale} onchange={onLogScale} testid="log-scale" />
      <span class="text-text text-[13px]" aria-hidden="true">{chartCopy.logScale}</span>
    </span>
  </div>
{/snippet}

{#snippet plot()}
  <ChartLegend
    {series}
    hidden={legend.hidden}
    highlightId={legend.highlightId}
    {notDrawn}
    {coarse}
    scrubMs={legend.scrubMs}
    formatValue={formatRate}
    unit={chartCopy.unit(metric)}
    onToggle={(id) => legend.toggle(id)}
    onOnly={(id) => legend.only(id, series)}
    onHighlight={(id) => legend.highlight(id)}
    onMore={() =>
      coarse ? (legend.sheetOpen = true) : document.getElementById('chart-players-trigger')?.click()}
  />
  <TimeSeriesChart
    {series}
    {events}
    {phases}
    domainMs={durationMs}
    window={current}
    {onWindow}
    plotHeights={{ base: 140, md: 140, lg: 140, xl: 140 }}
    scale={logScale ? 'log' : 'linear'}
    labelFloorS={model.bucket.labelFloorS}
    {syncKey}
    label={title}
    emptyCopy={chartCopy.emptyRate(metric)}
    hidden={legend.hidden}
    highlightId={legend.highlightId}
    {brushShape}
    formatValue={formatRate}
    rowsLabel={chartCopy.tooltipRate(metric, model.bucket.seconds)}
    onScrub={(ms) => (legend.scrubMs = ms)}
    {onOpenDeath}
  />
{/snippet}

<Section
  title={SECTION_TITLE[metric]}
  meta={wide.current ? meta : undefined}
  metaTestid="chart-head-meta"
  {testid}
  panel
>
  {#snippet head()}{#if wide.current}{@render controls()}{/if}{/snippet}
  {#if wide.current}
    {@render plot()}
  {:else}
    <Disclosure summary={meta} bind:open={chartOpen} testid="chart-collapsed">
      {@render controls()}
      {@render plot()}
    </Disclosure>
  {/if}
</Section>
<LegendSheet
  open={legend.sheetOpen}
  entries={model.entries}
  {medians}
  {choice}
  presets={PRESETS}
  formatValue={formatRate}
  onToggle={choose}
  onMedian={toggleMedian}
  {onPreset}
  onClose={() => (legend.sheetOpen = false)}
/>
