<!-- web/src/components/report/ThreatChart.svelte -->
<!-- Estimated threat on one enemy (controls and charts spec §B5, §B12): one cumulative line per
     player, the top 8 by standing at the window's end plus the tanks, the taunters and everyone
     the enemy attacked; taunt markers, target-change markers where the summary carries them, and
     deaths with their killing blow. The tooltip's share is of every player's estimated threat on
     the enemy at that second. The caveat line says what the estimate leaves out. -->
<script lang="ts">
  import type { Snippet } from 'svelte';
  import { togglePlayer } from '../../lib/chart/choice';
  import { chartCopy } from '../../lib/chart/copy';
  import { LegendState } from '../../lib/chart/legend-state.svelte';
  import { buildThreatModel, shareAt, type ThreatPlayer } from '../../lib/chart/threat';
  import type { ChartEvent, ChartPhase } from '../../lib/chart/types';
  import { formatAmount, formatPercent } from '../../lib/report/format';
  import type { Actor, Taunt, TargetChange, ThreatPair } from '../../lib/report/types';
  import { isFullWindow, type TimeWindow } from '../../lib/report/window';
  import ChartLegend from '../chart/ChartLegend.svelte';
  import LegendSheet from '../chart/LegendSheet.svelte';
  import PlayersControl from '../chart/PlayersControl.svelte';
  import TimeSeriesChart from '../chart/TimeSeriesChart.svelte';
  import Card from '../ui/Card.svelte';
  import Glyph from '../ui/Glyph.svelte';

  let {
    enemyName,
    enemyGuids,
    pairs,
    players,
    taunts,
    targetChanges,
    damageTaken,
    events,
    phases = [],
    durationMs,
    window: current,
    onWindow,
    sourcePlayer,
    youGuid,
    chosen,
    onChoose,
    brushShape = [],
    syncKey,
    onOpenDeath,
    enemyPicker,
  }: {
    /** The picked enemy's name; null with "Every enemy" picked, where every enemy's threat is summed. */
    enemyName: string | null;
    enemyGuids: readonly string[];
    /** Whole-fight pairs, every player. */
    pairs: readonly ThreatPair[];
    players: ReadonlyMap<string, ThreatPlayer>;
    taunts: readonly Taunt[];
    targetChanges?: readonly TargetChange[];
    damageTaken?: readonly Actor[];
    events: readonly ChartEvent[];
    phases?: readonly ChartPhase[];
    durationMs: number;
    window: TimeWindow;
    onWindow: (window: TimeWindow | null) => void;
    sourcePlayer?: string;
    youGuid?: string;
    chosen: readonly string[] | null;
    onChoose: (chosen: string[] | null) => void;
    brushShape?: readonly number[];
    syncKey: string;
    onOpenDeath?: (deaths: readonly ChartEvent[]) => void;
    /** The Enemy select, owned by the table that also reads it. */
    enemyPicker: Snippet;
  } = $props();

  const legend = new LegendState();
  const model = $derived(
    buildThreatModel({
      pairs,
      enemyGuids,
      players,
      window: current,
      taunts,
      targetChanges,
      damageTaken,
      sourcePlayer,
      youGuid,
      chosen: chosen === null ? null : new Set(chosen),
    }),
  );
  const whole = $derived(isFullWindow(current, durationMs));
  let viewport = $state(1440);
  let coarse = $state(false);
  $effect(() => {
    const query = window.matchMedia('(pointer: coarse), (hover: none)');
    const read = (): void => {
      viewport = window.innerWidth;
      coarse = query.matches;
    };
    read();
    window.addEventListener('resize', read);
    query.addEventListener('change', read);
    return () => {
      window.removeEventListener('resize', read);
      query.removeEventListener('change', read);
    };
  });
  const phone = $derived(viewport < 768);
  const title = $derived(chartCopy.threatTitle(enemyName));
  const notDrawn = $derived(model.entries.filter((entry) => !entry.drawn).length);
  function choose(id: string): void {
    onChoose([...togglePlayer(chosen === null ? null : new Set(chosen), model.entries, id)]);
  }
</script>

{#snippet headMeta()}
  <!-- On a phone the title takes the head's one row (it names the enemy, fix round 2); the brush
       under the plot carries the window's edges there. -->
  <span class="text-muted numeric text-[13px] max-md:hidden" data-testid="chart-head-meta"
    >{chartCopy.threatMeta(current.startMs, current.endMs, whole)}</span
  >
{/snippet}

{#snippet controls()}
  <div class="flex flex-wrap items-center justify-end gap-x-2 gap-y-2">
    {#if !phone}
      <span class="quiet-label">{chartCopy.players}</span>
      <PlayersControl
        id="threat-players"
        entries={model.entries}
        choice={model.isDefault ? 'top' : 'custom'}
        presets={['top']}
        formatValue={formatAmount}
        {coarse}
        onToggle={choose}
        onPreset={() => onChoose(null)}
        onCoarse={() => (legend.sheetOpen = true)}
      />
      {@render enemyPicker()}
    {/if}
    {#if !whole}
      <button
        type="button"
        class="reset inline-flex items-center gap-1.5"
        onclick={() => onWindow(null)}
        data-testid="window-reset"
      >
        <Glyph name="close" size={14} />{chartCopy.wholeFight}
      </button>
    {/if}
  </div>
{/snippet}

{#snippet caveat()}
  <p class="text-muted text-[12px] leading-4" data-testid="threat-caveat">
    {phone ? chartCopy.threatCaveatPhone : chartCopy.threatCaveat}
    <a class="caveat-link whitespace-nowrap" href="#glossary-threat">{chartCopy.threatCaveatLink}</a>
  </p>
{/snippet}

<Card quiet compact body="none" {title} titleSnippet={headMeta} aside={controls} testid="threat-chart">
  {#if phone}
    <div class="px-4 pb-2">{@render enemyPicker()}</div>
  {/if}
  <ChartLegend
    series={model.series}
    hidden={legend.hidden}
    highlightId={legend.highlightId}
    {notDrawn}
    {coarse}
    scrubMs={legend.scrubMs}
    formatValue={formatAmount}
    onToggle={(id) => legend.toggle(id)}
    onOnly={(id) => legend.only(id, model.series)}
    onHighlight={(id) => legend.highlight(id)}
    onMore={() => (coarse ? (legend.sheetOpen = true) : document.getElementById('threat-players')?.click())}
  />
  <TimeSeriesChart
    series={model.series}
    {events}
    {phases}
    domainMs={durationMs}
    window={current}
    {onWindow}
    plotHeights={{ base: 160, md: 220, lg: 260, xl: 300 }}
    mode="cumulative"
    {syncKey}
    label={chartCopy.threatTitle(enemyName)}
    emptyCopy={chartCopy.emptyThreat(chartCopy.threatSubject(enemyName))}
    rowsLabel={chartCopy.tooltipThreat(chartCopy.threatSubject(enemyName))}
    hidden={legend.hidden}
    highlightId={legend.highlightId}
    {brushShape}
    tooltipWidth={300}
    formatValue={formatAmount}
    shareOf={(value, ms) => formatPercent(shareAt(model, value, ms) * 100)}
    onScrub={(ms) => (legend.scrubMs = ms)}
    {onOpenDeath}
    testid="threat-time-chart"
  />
  <div class="border-line-soft flex min-h-7 items-center border-t px-4 py-1">{@render caveat()}</div>
</Card>
<LegendSheet
  open={legend.sheetOpen}
  entries={model.entries}
  choice={model.isDefault ? 'top' : 'custom'}
  presets={['top']}
  formatValue={formatAmount}
  onToggle={choose}
  onPreset={() => onChoose(null)}
  onClose={() => (legend.sheetOpen = false)}
/>

<style>
  .reset {
    min-height: 44px;
    padding: 0 8px;
    color: var(--color-gold);
    font-size: 12px;
    font-weight: 700;
    letter-spacing: 0.06em;
    text-transform: uppercase;
  }
  @media (min-width: 768px) {
    .reset {
      min-height: 32px;
    }
  }
  .caveat-link {
    color: var(--color-gold);
    text-decoration: underline;
    text-underline-offset: 3px;
  }
</style>
