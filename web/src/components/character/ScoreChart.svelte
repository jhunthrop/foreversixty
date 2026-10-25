<!-- web/src/components/character/ScoreChart.svelte -->
<!-- The History section's score column (experience spec 2026-10-24-character-history 3.3):
     sub-heading, the collapse note, the readout with its step buttons, the chart, the net
     change and the gap notes, and the fights behind the score. In a two-column section its
     children take the four shared rows (subgrid) so the chart tops line up with the parse
     column's. -->
<script lang="ts">
  import type { CharacterHistory } from '../../lib/character/history-api';
  import { buildScoreChart, scoreReadout, scoreSrRows } from '../../lib/character/history-chart';
  import { historyCopy } from '../../lib/character/history-copy';
  import { PLACEHOLDER_SIZE } from '../../lib/character/history-layout';
  import type { ChartSize } from '../../lib/character/history-scales';
  import {
    canStep,
    markedKeys,
    scoreListFights,
    scoreMarkKeys,
    shownIndex,
    step,
    select,
    withHover,
    type Selection,
    type Step,
  } from '../../lib/character/history-select';
  import { shortDate, tierColor } from '../../lib/score/score';
  import ChartPlot from './ChartPlot.svelte';
  import ChartReadout from './ChartReadout.svelte';
  import ChartSrTable from './ChartSrTable.svelte';
  import HistoryFights from './HistoryFights.svelte';

  let {
    history,
    sizeOverride = null,
  }: {
    history: CharacterHistory;
    /** A fixed drawing size, for tests and captures; the plot's own measure otherwise. */
    sizeOverride?: ChartSize | null;
  } = $props();

  let measured = $state<ChartSize | null>(null);
  let picked = $state<number | null>(null);
  let hover = $state<number | null>(null);

  const size = $derived(sizeOverride ?? measured);
  const chart = $derived(buildScoreChart(history, size ?? PLACEHOLDER_SIZE));
  const count = $derived(chart.days.length);
  /** The committed point: marks and announcements follow it, never a hover preview. */
  const committed = $derived(picked ?? count - 1);
  const selection = $derived<Selection>({ selected: committed, hover });
  const day = $derived(chart.days[shownIndex(selection)]);
  const rings = $derived(
    new Set(chart.days.flatMap((entry, index) => (entry.visibility === 'public' ? [] : [index]))),
  );
  const marked = $derived(markedKeys(scoreMarkKeys(chart.days), { selected: committed, hover: null }));
  const selectedDay = $derived(chart.days[selection.selected]);

  function commit(next: Selection): void {
    picked = next.selected;
  }
  const onstep = (move: Step): void => commit(step(selection, move, count));
  const onselect = (index: number): void => commit(select(selection, index, count));
  const onhover = (index: number | null): void => {
    hover = withHover(selection, index, count).hover;
  };
</script>

<div
  class="flex flex-col gap-3 xl:row-span-5 xl:grid xl:grid-rows-subgrid xl:gap-y-3"
  data-testid="history-score"
>
  <h3 class="text-[15px] font-semibold">{historyCopy.scoreHeading}</h3>

  <p class="text-muted flex min-h-11 items-center text-[12px]" data-testid="history-score-note">
    {chart.mode === 'line' ? historyCopy.scoreNote(chart.hasRings) : ''}
  </p>

  {#if day !== undefined}
    <ChartReadout
      readout={scoreReadout(day)}
      liveText={scoreReadout(chart.days[committed]).text}
      earlierLabel={historyCopy.step.score.earlier}
      laterLabel={historyCopy.step.score.later}
      canEarlier={canStep(selection, 'earlier', count)}
      canLater={canStep(selection, 'later', count)}
      {onstep}
      testid="history-score-readout"
    />
  {/if}

  {#if chart.mode === 'line'}
    <ChartPlot
      geometry={chart.geometry}
      {size}
      label={historyCopy.scorePlotLabel}
      selected={selection.selected}
      hover={selection.hover}
      selectedFill={tierColor(selectedDay.tier)}
      {rings}
      onmeasure={(next) => (measured = next)}
      {onselect}
      {onhover}
      {onstep}
    />
  {:else}
    <p class="text-text text-[13px]" data-testid="history-score-single">
      {historyCopy.scoreOnePoint(shortDate(chart.days[0].date) ?? chart.days[0].date)}
    </p>
  {/if}

  <div class="flex flex-col gap-3">
    {#if chart.change !== null}
      <p class="text-text text-[12px]" data-testid="history-score-change">{chart.change}</p>
    {/if}
    {#each chart.gapNotes as note (note)}
      <p class="text-muted text-[12px]" data-testid="history-gap-note">{note}</p>
    {/each}
    <HistoryFights
      fights={scoreListFights(history)}
      label={historyCopy.scoreListLabel}
      kind="fights"
      {marked}
      testid="history-score-fights"
    />
    <ChartSrTable
      caption={historyCopy.srScoreCaption}
      valueHeading={historyCopy.srColumns.score}
      rows={scoreSrRows(chart, history.gaps)}
      testid="history-score-table"
    />
  </div>
</div>
