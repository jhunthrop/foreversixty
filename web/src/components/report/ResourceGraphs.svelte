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
    formatDuration,
    wholeFightAriaLabel,
    wholeFightMark,
    wholeFightTitle,
  } from '../../lib/report/format';
  import type { ResourceTrack } from '../../lib/report/types';

  let { tracks, durationMs }: { tracks: ResourceTrack[]; durationMs: number } = $props();

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

  function points(series: number[]): string {
    const peak = series.reduce((highest, value) => Math.max(highest, value), 0);
    if (peak === 0 || series.length < 2) return '';
    return series
      .map((value, index) => `${(index / (series.length - 1)) * 100},${24 - (value / peak) * 22}`)
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
      <li
        class="border-line-soft grid min-h-11 grid-cols-[minmax(0,1fr)_auto] items-center gap-x-3 gap-y-1 border-b px-2 py-2 text-[14px] md:grid-cols-[minmax(120px,1.2fr)_96px_minmax(0,3fr)_96px]"
        data-testid={`resource-${track.guid}-${track.power_type}`}
      >
        <span class="truncate font-semibold">{splitUnitName(track.name).name}</span>
        <span class="text-muted text-[13px]"
          >{POWER_NAMES.get(track.power_type) ?? `Power ${track.power_type}`}</span
        >
        <svg
          class="col-span-2 h-[26px] w-full md:col-span-1"
          viewBox="0 0 100 26"
          preserveAspectRatio="none"
          aria-hidden="true"
        >
          <polyline
            points={points(track.series)}
            fill="none"
            stroke="var(--color-gold)"
            stroke-width="1.5"
            vector-effect="non-scaling-stroke"
          />
        </svg>
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
    Time empty is marked {mark} because the summary tracks it for the whole fight only; the line above is this window's
    own series.
  </p>
{/if}
