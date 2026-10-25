<!-- web/src/components/character/ParseChart.svelte -->
<!-- The History section's parse column (experience spec 2026-10-24-character-history 3.4):
     sub-heading, the boss picker, the readout with its step buttons, the chart of the selected
     boss's kills, the net change and that boss's kills list. A boss with one kill is a sentence
     and a one-row list; a character with no kill at all is one sentence. -->
<script lang="ts">
  import type { CharacterHistory } from '../../lib/character/history-api';
  import {
    bossPills,
    buildParseChart,
    defaultBossPill,
    parseReadout,
    parseSrRows,
  } from '../../lib/character/history-chart';
  import { historyCopy } from '../../lib/character/history-copy';
  import { PLACEHOLDER_SIZE } from '../../lib/character/history-layout';
  import type { ChartSize } from '../../lib/character/history-scales';
  import {
    canStep,
    killListFights,
    killMarkKeys,
    markedKeys,
    select,
    shownIndex,
    step,
    withHover,
    type Selection,
    type Step,
  } from '../../lib/character/history-select';
  import BossPicker from './BossPicker.svelte';
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
  let pickedBoss = $state<string | null>(null);
  let picked = $state<number | null>(null);
  let hover = $state<number | null>(null);

  const pills = $derived(bossPills(history.encounters));
  const pill = $derived(pills.find((entry) => entry.id === pickedBoss) ?? defaultBossPill(pills));
  const size = $derived(sizeOverride ?? measured);
  const chart = $derived(buildParseChart(pill?.encounter, size ?? PLACEHOLDER_SIZE));
  const count = $derived(chart.kills.length);
  /** The committed kill: marks and announcements follow it, never a hover preview. */
  const committed = $derived(picked ?? count - 1);
  const selection = $derived<Selection>({ selected: committed, hover });
  const marked = $derived(
    pill === undefined
      ? new Set<string>()
      : markedKeys(killMarkKeys(pill.encounter), { selected: committed, hover: null }),
  );

  function pickBoss(id: string): void {
    pickedBoss = id;
    picked = null;
    hover = null;
  }
  const onstep = (move: Step): void => {
    picked = step(selection, move, count).selected;
  };
  const onselect = (index: number): void => {
    picked = select(selection, index, count).selected;
  };
  const onhover = (index: number | null): void => {
    hover = withHover(selection, index, count).hover;
  };
</script>

<div
  class="flex flex-col gap-3 xl:row-span-5 xl:grid xl:grid-rows-subgrid xl:gap-y-3"
  data-testid="history-parse"
>
  <h3 class="text-[15px] font-semibold">{historyCopy.parseHeading}</h3>

  {#if pill === undefined}
    <p class="text-muted min-h-11 text-[14px] xl:row-span-4" data-testid="history-parse-empty">
      {historyCopy.parseNoKills}
    </p>
  {:else}
    <BossPicker {pills} activeId={pill.id} onpick={pickBoss} />

    {#if chart.mode === 'line'}
      <ChartReadout
        readout={parseReadout(chart.kills, shownIndex(selection))}
        liveText={parseReadout(chart.kills, committed).text}
        earlierLabel={historyCopy.step.parse.earlier}
        laterLabel={historyCopy.step.parse.later}
        canEarlier={canStep(selection, 'earlier', count)}
        canLater={canStep(selection, 'later', count)}
        {onstep}
        testid="history-parse-readout"
      />
      <ChartPlot
        geometry={chart.geometry}
        {size}
        label={historyCopy.parsePlotLabel(pill.label)}
        selected={selection.selected}
        hover={selection.hover}
        selectedFill={chart.kills[selection.selected].color}
        onmeasure={(next) => (measured = next)}
        {onselect}
        {onhover}
        {onstep}
      />
    {:else}
      <p class="text-muted min-h-11 text-[14px] xl:row-span-2" data-testid="history-parse-one-kill">
        {historyCopy.parseOneKill}
      </p>
    {/if}

    <div class="flex flex-col gap-3">
      {#if chart.change !== null}
        <p class="text-text text-[12px]" data-testid="history-parse-change">{chart.change}</p>
      {/if}
      {#key pill.id}
        <HistoryFights
          fights={killListFights(history, pill.encounter)}
          label={historyCopy.parseListLabel(pill.label)}
          kind="kills"
          lead="date"
          {marked}
          testid="history-parse-fights"
        />
      {/key}
      <ChartSrTable
        caption={historyCopy.srParseCaption}
        valueHeading={historyCopy.srColumns.parse}
        rows={parseSrRows(chart)}
        testid="history-parse-table"
      />
    </div>
  {/if}
</div>
