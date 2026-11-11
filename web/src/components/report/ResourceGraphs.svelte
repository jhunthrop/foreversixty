<!-- web/src/components/report/ResourceGraphs.svelte -->
<!-- Resources (logs spec 2026-11-07 section 6.2): the tab's one Plate, a StatTable row per
     actor and power type: the crest, the name over the power and its peak, low and cap
     figures, a Spark of the bar over the fight in the class colour (the time at the cap
     shaded, the cap drawn, deaths marked), then Gained, Spent, Wasted and Time empty. Time
     at zero is the number a mana-starved healer is looking for, so it is a column.

     The line and the number beside it are answering two different questions. The series
     is exact under any window: window.ts's scopeResource only slices it to the window's
     buckets. `zero_ms` (and `gained`/`spent`, not shown here) are not sliced at all --
     they come through scopeResource untouched, so they are silently the whole fight's
     total under every window, the same lie Interrupts and Dispels tell. It carries the
     `†` mark unconditionally, exactly as ExchangeTable's Count does. -->
<script lang="ts">
  import { splitUnitName } from '../../lib/characters';
  import { clearChartCursor, cursorMsFor, setChartCursor } from '../../lib/chart/cursor.svelte';
  import { clockLabel } from '../../lib/chart/time-axis';
  import SyncedCrosshair from '../chart/SyncedCrosshair.svelte';
  import { classColorVar, wholeFightMark } from '../../lib/report/format';
  import type { ResourceTrack } from '../../lib/report/types';
  import { classSlugOf } from '../../lib/report/planner-link';
  import type { StatColumn, StatFigure } from '../../lib/ui/stat-table';
  import { clockDuration } from '../../lib/format/duration';
  import { compact, percent } from '../../lib/format/number';
  import Crest from '../ui/Crest.svelte';
  import Section from '../ui/Section.svelte';
  import Spark from '../ui/Spark.svelte';
  import StatRow from '../ui/StatRow.svelte';
  import StatTable from '../ui/StatTable.svelte';
  import CopyCsv from './CopyCsv.svelte';

  let {
    tracks,
    durationMs,
    deaths = [],
    windowed = false,
    syncKey = '',
    windowStartMs = 0,
    classOf = new Map<string, string>(),
  }: {
    tracks: ResourceTrack[];
    durationMs: number;
    deaths?: { guid: string; at_ms: number }[];
    /** True while a window is brushed, so the at-cap figure can say what it measured. */
    windowed?: boolean;
    /** The page's chart cursor key (the fight); empty draws no shared crosshair (a sim). */
    syncKey?: string;
    /** Where the sliced series starts on the fight's clock, so the cursor's ms finds its second. */
    windowStartMs?: number;
    /** GUID to class, for the crest and the line's colour. */
    classOf?: ReadonlyMap<string, string>;
  } = $props();
  const owner = 'resource-graphs';
  /** The page cursor's second inside this window, from any chart or lane on the page. */
  const cursorMs = $derived(syncKey === '' ? null : cursorMsFor(syncKey));
  const shownWindow = $derived({ startMs: windowStartMs, endMs: windowStartMs + durationMs });
  /** The value at the shared cursor: the figure slot beside the line while the cursor is set. */
  function atCursor(series: number[]): string | null {
    if (cursorMs === null) return null;
    const index = Math.floor((cursorMs - windowStartMs) / 1000);
    if (index < 0 || index >= series.length) return null;
    return `at ${clockLabel(cursorMs)} · ${compact(series[index] ?? 0)}`;
  }
  /**
   * What the at-cap share is a share of. The figure has always been the window's own --
   * the series is sliced and `durationMs` is the window's length -- but the words beside
   * it said "of the fight" whatever the brush was doing, so a four-minute slice capped
   * throughout read "at cap 100.0% of the fight" on a twelve-minute pull.
   */
  const capSpan = $derived(windowed ? 'window' : 'fight');

  /** The reading under the pointer, per line: "at 40.6s · 1,188", so the picture has a number. */
  let readouts = $state<Record<string, string>>({});
  /** A mouse leaving the line takes its readout with it; a tap's readout stays for reading. */
  function clearReadout(event: PointerEvent, key: string): void {
    if (event.pointerType !== 'mouse') return;
    const { [key]: _gone, ...rest } = readouts;
    readouts = rest;
    clearChartCursor(owner);
  }
  function readAt(event: PointerEvent, key: string, series: number[]): void {
    const box = (event.currentTarget as HTMLElement).getBoundingClientRect();
    if (box.width === 0 || series.length === 0) return;
    const index = Math.min(
      series.length - 1,
      Math.max(0, Math.round(((event.clientX - box.left) / box.width) * (series.length - 1))),
    );
    readouts = {
      ...readouts,
      [key]: `at ${clockLabel(windowStartMs + index * 1000)} · ${compact(series[index] ?? 0)}`,
    };
    if (syncKey !== '') setChartCursor({ syncKey, ms: windowStartMs + index * 1000, owner });
  }

  /** The lowest point of a series and the second it happened, for the figure beside the line. */
  function low(series: number[]): { value: number; atMs: number } {
    let value = Number.POSITIVE_INFINITY;
    let atMs = 0;
    // The seconds before the first reading are not zero power, they are no reading.
    const first = series.findIndex((point) => point > 0);
    series.forEach((point, index) => {
      if (index >= first && first >= 0 && point < value) {
        value = point;
        atMs = index * 1000;
      }
    });
    return { value: Number.isFinite(value) ? value : 0, atMs };
  }

  const peakOf = (series: number[]): number => series.reduce((highest, value) => Math.max(highest, value), 0);

  /** The top of the sparkline's scale: the peak reading, or the cap when the line never reached it. */
  const scaleTop = (peak: number, max: number | undefined): number => Math.max(peak, max ?? 0);
  /** A reading's y in the Spark's 24-unit box (lib/ui/spark.ts: baseline 22, rise 20), so the cap
   *  line sits on the line's own scale. */
  const yOf = (value: number, top: number): number => (top === 0 ? 22 : 22 - (value / top) * 20);
  /** The seconds whose reading is at the cap, as [from, to] percentages of the line's width. */
  function atMaxSpans(series: number[], max: number): { from: number; to: number }[] {
    if (max <= 0 || series.length === 0) return [];
    const width = 100 / series.length;
    const out: { from: number; to: number }[] = [];
    series.forEach((value, index) => {
      if (value < max) return;
      const from = index * width;
      const last = out[out.length - 1];
      if (last !== undefined && Math.abs(last.to - from) < 0.0001) last.to = from + width;
      else out.push({ from, to: from + width });
    });
    return out;
  }
  /** The share of the window the bar spent full. */
  function atCapPct(atMaxMs: number): number {
    return durationMs === 0 ? 0 : (atMaxMs / durationMs) * 100;
  }
  /** Every series as lines: the player, the power, the second, the reading and whether it was at the cap. */
  function seriesCsv(): string[][] {
    return [
      ['Player', 'Power', 'Second', 'Reading', 'At cap'],
      ...rows.flatMap((track) =>
        track.series.map((value, index) => [
          splitUnitName(track.name).name,
          powerName(track.power_type),
          String(index),
          String(value),
          track.max !== undefined && track.max > 0 && value >= track.max ? 'yes' : 'no',
        ]),
      ),
    ];
  }

  /** The game's power indices. Anything else is shown by its number, not guessed at. */
  const POWER_NAMES = new Map<number, string>([
    [0, 'Mana'],
    [1, 'Rage'],
    [2, 'Focus'],
    [3, 'Energy'],
    [4, 'Combo points'],
    [6, 'Runic power'],
    [7, 'Soul shards'],
    [8, 'Astral power'],
    [9, 'Holy power'],
    [11, 'Maelstrom'],
    [12, 'Chi'],
    [13, 'Insanity'],
    [16, 'Arcane charges'],
    [17, 'Fury'],
    [18, 'Pain'],
    [19, 'Essence'],
  ]);

  /**
   * A sim reports no per-second reading at all (sim/adapter/adapter.go's `resources`: the
   * engine gives it a total gained/spent/wasted, and nothing else) -- `track.series` is
   * always `[]` for one, never merely mostly-zero. Filtering on the series alone, as this
   * used to, excluded EVERY sim resource row unconditionally: a fight where Bloodrage
   * generated real rage (gained/spent both well above zero) still read "No resource
   * changes in this window", contradicting the ONE ITERATION tab's own cast log right next
   * to it (2026-09-21 result-page review, Defect 4). `wasted` is optional and only ever
   * carries a real figure when it is worth reporting (adapter.go's own `if w > 0`), so a
   * present `wasted` is itself already evidence of activity.
   */
  const hadActivity = (track: ResourceTrack): boolean =>
    track.series.some((value) => value > 0) || track.gained > 0 || track.spent > 0 || (track.wasted ?? 0) > 0;

  const powerName = (type: number): string => POWER_NAMES.get(type) ?? `Power ${type}`;
  const rows = $derived(
    tracks.filter(hadActivity).sort((a, b) => a.name.localeCompare(b.name) || a.power_type - b.power_type),
  );

  const mark = $derived(wholeFightMark(windowed));
  const hasCap = (track: ResourceTrack): boolean => track.max !== undefined && track.max > 0;

  const columns: readonly StatColumn[] = [
    { id: 'gained', label: 'Gained' },
    { id: 'spent', label: 'Spent' },
    { id: 'wasted', label: 'Wasted' },
    { id: 'empty', label: 'Time empty' },
  ];
  function figuresOf(track: ResourceTrack): StatFigure[] {
    return [
      { value: `${mark}${compact(track.gained)}` },
      { value: `${mark}${compact(track.spent)}` },
      track.wasted === undefined ? { value: '–', muted: true } : { value: `${mark}${compact(track.wasted)}` },
      {
        value: `${mark}${track.zero_ms === 0 ? '0:00' : clockDuration(track.zero_ms)}`,
        testid: 'resource-zero',
      },
    ];
  }
  const SEPARATOR = ' · ';
</script>

<Section
  title="Resources"
  meta={`Mana, rage, energy and focus, ${capSpan === 'window' ? 'this window' : 'whole fight'}`}
  panel
  flush
>
  {#snippet head()}{#if rows.some((track) => track.series.length > 0)}<CopyCsv
        lines={seriesCsv}
        label="Copy every line as CSV"
        inHead
      />{/if}{/snippet}
  <div>
    {#if rows.length === 0}
      <p class="text-muted px-4 py-4 text-[14px]" data-testid="table-empty">
        No resource changes in this window.
      </p>
    {:else}
      <StatTable
        {columns}
        nameHead="Player"
        barHead={`Over the ${capSpan}`}
        label="Resources"
        testid="resource-graphs"
      >
        {#each rows as track (`${track.guid}-${track.power_type}`)}
          {@const lowest = low(track.series)}
          {@const peak = peakOf(track.series)}
          {@const top = scaleTop(peak, track.max)}
          {@const colour = classColorVar(classOf.get(track.guid))}
          {@const key = `${track.guid}-${track.power_type}`}
          {@const reading = readouts[key] ?? atCursor(track.series)}
          <StatRow
            figures={figuresOf(track)}
            phone={`${compact(track.spent)} spent`}
            testid={`resource-${track.guid}-${track.power_type}`}
          >
            {#snippet icon()}<Crest
                cls={classSlugOf(classOf.get(track.guid)) ?? ''}
                size={28}
              />{/snippet}
            {#snippet name()}<span style:color={colour}>{splitUnitName(track.name).name}</span>{/snippet}
            {#snippet sub()}
              <!-- A block, so the figures wrap between pairs on a phone (each pair stays whole). -->
              <div>
                <span class="whitespace-nowrap">{powerName(track.power_type)}</span
                >{#if track.series.length === 0}{SEPARATOR}<span
                    >No per-second reading or cap for a simulated fight; these are the fight's own totals.</span
                  ><span class="ui-st-phone" data-testid="resource-totals-only"
                    >{SEPARATOR}gained {compact(track.gained)}{SEPARATOR}spent {compact(
                      track.spent,
                    )}{#if (track.wasted ?? 0) > 0}{`${SEPARATOR}wasted ${compact(track.wasted ?? 0)}`}{/if}</span
                  >{:else}{SEPARATOR}<span data-testid={hasCap(track) ? 'resource-cap-figures' : undefined}
                    ><span class="whitespace-nowrap" title="The top of the line">peak {compact(peak)}</span
                    >{SEPARATOR}<span
                      class="whitespace-nowrap"
                      title="The lowest point and when it was reached"
                      data-testid="resource-low"
                      >low {compact(lowest.value)} at {clockDuration(lowest.atMs)}</span
                    >{#if hasCap(track)}{SEPARATOR}<span
                        class="whitespace-nowrap"
                        title="The share of this window the bar spent full, measured from the window's own seconds"
                        data-testid="resource-at-cap"
                        >at cap {percent(atCapPct(track.at_max_ms ?? 0))} of the {capSpan}</span
                      >{SEPARATOR}<span class="whitespace-nowrap"
                        >{mark}wasted {compact(track.wasted ?? 0)}</span
                      >{/if}</span
                  >{/if}
              </div>
            {/snippet}
            {#snippet meter()}
              {#if track.series.length > 0}
                <!-- A pointer over the line reads the second under it; a tap on a phone does the same. -->
                <span
                  class="relative block touch-none"
                  role="presentation"
                  onpointermove={(event) => readAt(event, key, track.series)}
                  onpointerdown={(event) => readAt(event, key, track.series)}
                  onpointerleave={(event) => clearReadout(event, key)}
                >
                  <Spark values={track.series} {top} color={colour} />
                  <svg
                    class="pointer-events-none absolute inset-0 h-6 w-full"
                    viewBox="0 0 1000 24"
                    preserveAspectRatio="none"
                    aria-hidden="true"
                  >
                    {#each deaths as death, i (`${death.at_ms}-${i}`)}
                      <rect
                        x={durationMs === 0 ? 0 : (death.at_ms / durationMs) * 1000}
                        y="0"
                        width={death.guid === track.guid ? 4 : 2}
                        height="24"
                        fill={death.guid === track.guid ? 'var(--color-death)' : 'var(--color-death-soft)'}
                      />
                    {/each}
                    {#if hasCap(track)}
                      {#each atMaxSpans(track.series, track.max ?? 0) as span, i (`${span.from}-${i}`)}
                        <rect
                          x={span.from * 10}
                          y="0"
                          width={(span.to - span.from) * 10}
                          height="24"
                          fill="var(--color-accent)"
                          opacity="0.18"
                          data-testid="resource-at-max"
                        />
                      {/each}
                      <!-- A filled rect, not a <line>: a horizontal line has no height for a visibility check. -->
                      <rect
                        x="0"
                        y={yOf(track.max ?? 0, top) - 0.6}
                        width="1000"
                        height="1.2"
                        fill="var(--color-accent)"
                        opacity="0.6"
                        data-testid="resource-cap-line"
                      />
                    {/if}
                  </svg>
                  {#if syncKey !== ''}<SyncedCrosshair {syncKey} window={shownWindow} />{/if}
                </span>
                {#if reading}<span
                    class="text-strong block text-[12px] leading-4 tabular-nums"
                    data-testid="resource-readout">{reading}</span
                  >{/if}
              {/if}
            {/snippet}
          </StatRow>
        {/each}
      </StatTable>
    {/if}
  </div>
  {#if rows.length > 0 && windowed}
    <div class="text-muted flex flex-col gap-1 px-5 pt-3 pb-4 text-[13px] leading-[18px]">
      <p>Fight length in this window: <span class="tabular-nums">{clockDuration(durationMs)}</span>.</p>
      <p data-testid="resource-wholefight-note">
        Gained, spent, wasted and time empty are marked {mark} because the summary tracks them for the whole fight
        only; the line, the cap and the time at the cap are this window's own.
      </p>
    </div>
  {/if}
</Section>
