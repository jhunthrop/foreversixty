<!-- web/src/components/report/TimeChart.svelte -->
<!-- The report's main chart (controls and charts spec §B5, D5): the table's metric per player,
     averaged over the fight's bucket (lib/chart/buckets.ts: every second on a short pull, 5 s on a long fight), the top 4 and the bottom 4 of N (fix round 2) plus the Source player and you, the raid median, the Players
     control, the Window presets and the log scale, all on the shared TimeSeriesChart. The head's
     raid figure is the sum of every player's figure as the legend shows it. The sample report's
     `compact` strip, on a phone only, is the same chart at 96 px with no legend, lane or brush. -->
<script lang="ts">
  import { splitUnitName } from '../../lib/characters';
  import { togglePlayer } from '../../lib/chart/choice';
  import { chartCopy, type RateMetric } from '../../lib/chart/copy';
  import type { DeadSpan, DrawPreset, PlayersChoice } from '../../lib/chart/drawn';
  import { LegendState } from '../../lib/chart/legend-state.svelte';
  import { entityStyle } from '../../lib/chart/palette';
  import { buildRateModel, classMedianInfo, RAID_MEDIAN_ID, type RatePlayer } from '../../lib/chart/rate';
  import { clockLabel } from '../../lib/chart/time-axis';
  import type { ChartEvent, ChartPhase, ChartSeries, PlayerEntry } from '../../lib/chart/types';
  import { formatAmount } from '../../lib/report/format';
  import type { ScopeNoun } from '../../lib/report/group';
  import type { Actor, RosterRow } from '../../lib/report/types';
  import { isFullWindow, type TimeWindow, type WindowPreset } from '../../lib/report/window';
  import ChartLegend from '../chart/ChartLegend.svelte';
  import LegendSheet from '../chart/LegendSheet.svelte';
  import PlayersControl from '../chart/PlayersControl.svelte';
  import TimeSeriesChart from '../chart/TimeSeriesChart.svelte';
  import Card from '../ui/Card.svelte';
  import Select from '../ui/Select.svelte';
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
    presets,
    onSelectPreset,
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
    compact = false,
    noun = 'raid',
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
    presets: readonly WindowPreset[];
    onSelectPreset: (index: number) => void;
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
    compact?: boolean;
    /** What the drawn players are (lib/report/group.ts): it names the head's figure and the median. */
    noun?: ScopeNoun;
  } = $props();

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
  const selectedPreset = $derived(
    presets.findIndex((preset) =>
      preset.window === null
        ? whole
        : preset.window.startMs === current.startMs && preset.window.endMs === current.endMs,
    ),
  );
  const windowOptions = $derived([
    ...presets.map((preset, index) => ({ value: String(index), label: preset.label })),
    ...(selectedPreset < 0
      ? [{ value: 'custom', label: `${clockLabel(current.startMs)} to ${clockLabel(current.endMs)}` }]
      : []),
  ]);
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
      return chartCopy.playerMeta(metric, single.label, formatAmount(single.figure), span, targetName);
    return chartCopy.rateMeta(metric, formatAmount(model.raidFigure), span, targetName, noun);
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

{#snippet headMeta()}
  <span class="text-muted numeric text-[13px]" data-testid="chart-head-meta">{meta}</span>
{/snippet}

{#snippet controls()}
  <!-- On a phone the head wraps (fix round 2 ruling): the title on its row, then each field on
       its own full-width row, so every field shows its value whole. -->
  <div
    class="flex flex-wrap items-center justify-end gap-x-2 gap-y-2 max-md:grid max-md:w-full max-md:grid-cols-[64px_minmax(0,1fr)] max-md:justify-items-stretch max-md:text-left"
    data-testid="chart-head-controls"
  >
    {#if !enemies}
      <span class="quiet-label">{chartCopy.players}</span>
      <PlayersControl
        id="chart-players"
        entries={model.entries}
        {medians}
        {choice}
        presets={PRESETS}
        seen={noun === 'everyone'}
        formatValue={formatAmount}
        onToggle={choose}
        onMedian={toggleMedian}
        {onPreset}
        onCoarse={() => (legend.sheetOpen = true)}
      />
    {/if}
    <label class="quiet-label" for="window-select">{chartCopy.window}</label>
    <div class="w-[180px] max-md:w-full">
      <Select
        id="window-select"
        class="w-full"
        size="dense"
        align="end"
        testid="window-select"
        options={windowOptions}
        value={selectedPreset < 0 ? 'custom' : String(selectedPreset)}
        onchange={(event) => {
          const value = (event.currentTarget as HTMLSelectElement).value;
          if (value !== 'custom') onSelectPreset(Number(value));
        }}
      />
    </div>
    <span class="inline-flex items-center gap-2 max-md:col-span-2">
      <Switch checked={logScale} label={chartCopy.logScale} onchange={onLogScale} testid="log-scale" />
      <span class="text-text text-[13px]" aria-hidden="true">{chartCopy.logScale}</span>
    </span>
  </div>
{/snippet}

{#if compact}
  <div class="card-surface-quiet relative overflow-hidden" data-testid="time-chart" data-compact="true">
    {#if !whole}
      <button
        type="button"
        class="border-line-warm bg-raised text-text rounded-control absolute top-1 right-1 z-10 inline-flex h-7 items-center border px-2 text-[11px] font-bold tracking-[0.06em] uppercase"
        onclick={() => onWindow(null)}
        data-testid="window-reset-compact">{chartCopy.wholeFight}</button
      >
    {/if}
    <TimeSeriesChart
      series={model.series}
      domainMs={durationMs}
      window={current}
      {onWindow}
      plotHeights={{ base: 74, md: 74, lg: 74, xl: 74 }}
      labelFloorS={model.bucket.labelFloorS}
      {syncKey}
      label={title}
      emptyCopy={chartCopy.emptyRate(metric)}
      formatValue={formatAmount}
      compact
    />
  </div>
{:else}
  <Card
    quiet
    compact
    body="none"
    {title}
    titleSnippet={headMeta}
    aside={controls}
    headClass="max-md:[&>div:last-child]:ml-0 max-md:[&>div:last-child]:basis-full"
    testid="time-chart"
  >
    <ChartLegend
      {series}
      hidden={legend.hidden}
      highlightId={legend.highlightId}
      {notDrawn}
      {coarse}
      scrubMs={legend.scrubMs}
      formatValue={formatAmount}
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
      plotHeights={{ base: 140, md: 160, lg: 200, xl: 220 }}
      scale={logScale ? 'log' : 'linear'}
      labelFloorS={model.bucket.labelFloorS}
      {syncKey}
      label={title}
      emptyCopy={chartCopy.emptyRate(metric)}
      hidden={legend.hidden}
      highlightId={legend.highlightId}
      {brushShape}
      formatValue={formatAmount}
      rowsLabel={chartCopy.tooltipRate(metric, model.bucket.seconds)}
      onScrub={(ms) => (legend.scrubMs = ms)}
      {onOpenDeath}
    />
  </Card>
  <LegendSheet
    open={legend.sheetOpen}
    entries={model.entries}
    {medians}
    {choice}
    presets={PRESETS}
    formatValue={formatAmount}
    onToggle={choose}
    onMedian={toggleMedian}
    {onPreset}
    onClose={() => (legend.sheetOpen = false)}
  />
{/if}
