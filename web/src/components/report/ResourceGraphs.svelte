<!-- web/src/components/report/ResourceGraphs.svelte -->
<!-- One sparkline per actor and power type, drawn as an inline SVG polyline rather than a
     canvas: these are small, numerous, and never interactive, so markup beats a second
     drawing surface. Time at zero is the number a mana-starved healer is looking for, so
     it sits beside the line rather than in a tooltip.

     The line and the number beside it are answering two different questions. The series
     is exact under any window: window.ts's scopeResource only slices it to the window's
     buckets. `zero_ms` (and `gained`/`spent`, not shown here) are not sliced at all --
     they come through scopeResource untouched, so they are silently the whole fight's
     total under every window, the same lie Interrupts and Dispels tell. It carries the
     `†` mark unconditionally, exactly as ExchangeTable's Count does. -->
<script lang="ts">
  import { splitUnitName } from '../../lib/characters';
  import {
    formatAmount,
    formatDuration,
    formatPercent,
    wholeFightAriaLabel,
    wholeFightMark,
    wholeFightTitle,
  } from '../../lib/report/format';
  import type { ResourceTrack } from '../../lib/report/types';
  import CopyCsv from './CopyCsv.svelte';

  let {
    tracks,
    durationMs,
    deaths = [],
  }: { tracks: ResourceTrack[]; durationMs: number; deaths?: { guid: string; at_ms: number }[] } = $props();

  /** The reading under the pointer, per line: "at 40.6s · 1,188", so the picture has a number. */
  let readouts = $state<Record<string, string>>({});
  /** A mouse leaving the line takes its readout with it; a tap's readout stays for reading. */
  function clearReadout(event: PointerEvent, key: string): void {
    if (event.pointerType !== 'mouse') return;
    const { [key]: _gone, ...rest } = readouts;
    readouts = rest;
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
      [key]: `at ${formatDuration(index * 1000)} · ${formatAmount(series[index] ?? 0)}`,
    };
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

  /** Where the cap sits on the 26-unit viewBox the sparkline is drawn in. */
  function capY(series: number[], max: number): number {
    const peak = peakOf(series);
    const top = Math.max(peak, max);
    return top === 0 ? 24 : 24 - (max / top) * 22;
  }
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
  /** One track's series as lines: the second, the reading, and whether it was at the cap. */
  function seriesCsv(track: ResourceTrack): string[][] {
    return [
      ['Second', 'Reading', 'At cap'],
      ...track.series.map((value, index) => [
        String(index),
        String(value),
        track.max !== undefined && track.max > 0 && value >= track.max ? 'yes' : 'no',
      ]),
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

  const rows = $derived(
    tracks
      .filter((track) => track.series.some((value) => value > 0))
      .sort((a, b) => a.name.localeCompare(b.name) || a.power_type - b.power_type),
  );

  function points(series: number[], max = 0): string {
    const top = Math.max(peakOf(series), max);
    if (top === 0 || series.length < 2) return '';
    return series
      .map((value, index) => `${(index / (series.length - 1)) * 100},${24 - (value / top) * 22}`)
      .join(' ');
  }

  const mark = wholeFightMark(true);
  const title = wholeFightTitle(true);

  function zeroText(zeroMs: number): string {
    return zeroMs === 0 ? 'never empty' : `${formatDuration(zeroMs)} empty`;
  }
</script>

{#if rows.length === 0}
  <p class="text-muted text-[14px]" data-testid="table-empty">No resource changes in this window.</p>
{:else}
  <ul class="flex flex-col" data-testid="resource-graphs">
    {#each rows as track (`${track.guid}-${track.power_type}`)}
      {@const lowest = low(track.series)}
      {@const peak = peakOf(track.series)}
      <li
        class="border-line-soft grid min-h-11 grid-cols-[minmax(0,1fr)_auto] items-center gap-x-3 gap-y-1 border-b px-2 py-2 text-[14px] md:grid-cols-[minmax(120px,1.2fr)_96px_minmax(0,3fr)_96px]"
        data-testid={`resource-${track.guid}-${track.power_type}`}
      >
        <span class="truncate font-semibold">{splitUnitName(track.name).name}</span>
        <span class="text-muted text-[13px]"
          >{POWER_NAMES.get(track.power_type) ?? `Power ${track.power_type}`}</span
        >
        <span class="col-span-2 flex flex-col gap-0.5 md:col-span-1">
          <!-- A pointer over the line reads the second under it; a tap on a phone does the same. -->
          <span
            class="block touch-none"
            onpointermove={(event) => readAt(event, `${track.guid}-${track.power_type}`, track.series)}
            onpointerdown={(event) => readAt(event, `${track.guid}-${track.power_type}`, track.series)}
            onpointerleave={(event) => clearReadout(event, `${track.guid}-${track.power_type}`)}
          >
            <svg class="h-[40px] w-full" viewBox="0 0 100 26" preserveAspectRatio="none" aria-hidden="true">
              <line
                x1="0"
                y1="2"
                x2="100"
                y2="2"
                stroke="var(--color-line-soft)"
                stroke-width="1"
                vector-effect="non-scaling-stroke"
              />
              <line
                x1="0"
                y1="24"
                x2="100"
                y2="24"
                stroke="var(--color-line-soft)"
                stroke-width="1"
                vector-effect="non-scaling-stroke"
              />
              {#each deaths as death, i (`${death.at_ms}-${i}`)}
                <line
                  x1={durationMs === 0 ? 0 : (death.at_ms / durationMs) * 100}
                  y1="0"
                  x2={durationMs === 0 ? 0 : (death.at_ms / durationMs) * 100}
                  y2="26"
                  stroke={death.guid === track.guid ? 'var(--color-death)' : 'var(--color-death-soft)'}
                  stroke-width={death.guid === track.guid ? 2 : 1}
                  vector-effect="non-scaling-stroke"
                />
              {/each}
              {#if track.max !== undefined && track.max > 0}
                {#each atMaxSpans(track.series, track.max) as span, i (`${span.from}-${i}`)}
                  <rect
                    x={span.from}
                    y="0"
                    width={span.to - span.from}
                    height="26"
                    fill="var(--color-gold)"
                    opacity="0.18"
                    data-testid="resource-at-max"
                  />
                {/each}
                <!-- A perfectly horizontal <line> has zero geometric height under
                     getBoundingClientRect, so a visibility check on it always reads
                     hidden regardless of where it is drawn; a thin filled <rect> at the
                     same y (from capY, the same scale the sparkline itself uses) carries
                     genuine area and reads the same as a cap line on screen. -->
                <rect
                  x="0"
                  y={capY(track.series, track.max) - 0.4}
                  width="100"
                  height="0.8"
                  fill="var(--color-gold)"
                  data-testid="resource-cap-line"
                />
              {/if}
              <polyline
                points={points(track.series, track.max ?? 0)}
                fill="none"
                stroke="var(--color-gold)"
                stroke-width="1.5"
                vector-effect="non-scaling-stroke"
              />
            </svg>
          </span>
          {#if readouts[`${track.guid}-${track.power_type}`]}
            <span class="text-strong tabular font-mono text-[11px]" data-testid="resource-readout"
              >{readouts[`${track.guid}-${track.power_type}`]}</span
            >
          {/if}
          <!-- The peak/low reading always renders; when the track carries a reported cap the
               same line also carries the cap figures, so a test (or a reader) that asks for
               "the figures" sees the peak beside the cap, the time at it and the waste. -->
          <span
            class={track.max !== undefined && track.max > 0
              ? 'text-muted tabular flex flex-wrap items-baseline gap-x-3 font-mono text-[11px]'
              : 'text-muted tabular flex justify-between font-mono text-[11px]'}
            data-testid={track.max !== undefined && track.max > 0 ? 'resource-cap-figures' : undefined}
          >
            <span title="The top of the line">peak {formatAmount(peak)}</span>
            <span title="The lowest point and when it was reached" data-testid="resource-low"
              >low {formatAmount(lowest.value)} at {formatDuration(lowest.atMs)}</span
            >
            {#if track.max !== undefined && track.max > 0}
              <span title="The cap the log reported for this power">cap {formatAmount(track.max)}</span>
              <span
                title="The share of this window the bar spent full, measured from the window's own seconds"
                >at cap {formatPercent(atCapPct(track.at_max_ms ?? 0))} of the fight</span
              >
              <!-- One title per element: the whole-fight one from wholeFightTitle(true),
                   already bound to this file's `title` const, and the words that explain
                   what the figure is in the aria-label, which is where a title on an
                   element with visible text would not reliably be read out anyway. -->
              <span
                {title}
                aria-label={wholeFightAriaLabel(
                  true,
                  `${formatAmount(track.wasted ?? 0)} of power gained past the cap and thrown away`,
                )}>{mark}wasted {formatAmount(track.wasted ?? 0)}</span
              >
            {/if}
          </span>
          <CopyCsv lines={() => seriesCsv(track)} label="Copy this line as CSV" />
        </span>
        <span
          class="text-muted tabular text-right font-mono text-[13px]"
          {title}
          aria-label={wholeFightAriaLabel(true, zeroText(track.zero_ms))}
          data-testid="resource-zero"
        >
          {mark}{zeroText(track.zero_ms)}
        </span>
      </li>
    {/each}
  </ul>
  <p class="text-muted text-[12px]">
    Fight length in this window: <span class="tabular font-mono">{formatDuration(durationMs)}</span>.
  </p>
  <p class="text-muted text-[12px]" data-testid="resource-wholefight-note">
    Time empty and power wasted are marked {mark} because the summary tracks them for the whole fight only; the
    line, the cap and the time at the cap are this window's own.
  </p>
{/if}
