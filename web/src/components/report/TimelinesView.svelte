<!-- web/src/components/report/TimelinesView.svelte -->
<!-- One lane per player: their casts as ticks, the auras they held as bars, and their
     death as a full-height mark. The same window the tables use, drawn horizontally, so
     "who was doing nothing for eight seconds" is a glance rather than a calculation.

     Every source drawn here -- cast sequence, aura segments, deaths -- is exact under a
     window (window.ts's scopeSummary), so nothing on this lane needs an approximate mark. -->
<script lang="ts">
  import { splitUnitName } from '../../lib/characters';
  import { classColorVar, formatDuration } from '../../lib/report/format';
  import type { Summary } from '../../lib/report/types';
  import type { TimeWindow } from '../../lib/report/window';

  let {
    summary,
    window: current,
    classOf,
  }: { summary: Summary; window: TimeWindow; classOf: Map<string, string> } = $props();

  const span = $derived(Math.max(current.endMs - current.startMs, 1));
  const pct = (ms: number): number => ((ms - current.startMs) / span) * 100;

  const lanes = $derived(
    [...summary.roster]
      .sort((a, b) => b.damage_done + b.healing_done - (a.damage_done + a.healing_done))
      .map((row) => ({
        guid: row.guid,
        name: splitUnitName(row.name).name,
        color: classColorVar(row.class ?? classOf.get(row.guid)),
        casts: summary.casts.filter((cast) => cast.guid === row.guid).flatMap((cast) => cast.sequence),
        auras: summary.auras.filter((track) => track.target_guid === row.guid).flatMap((track) => track.segments),
        deaths: summary.deaths.filter((death) => death.guid === row.guid).map((death) => death.at_ms),
      })),
  );
</script>

{#if lanes.length === 0}
  <p class="text-muted text-[14px]" data-testid="table-empty">Nobody was in this window.</p>
{:else}
  <div class="flex flex-col gap-1" data-testid="timelines">
    <p class="text-muted label">
      {formatDuration(current.startMs)} to {formatDuration(current.endMs)} · casts as ticks, auras as bars
    </p>
    <ul class="flex flex-col">
      {#each lanes as lane (lane.guid)}
        <li
          class="border-line-soft grid min-h-11 grid-cols-[minmax(96px,140px)_minmax(0,1fr)] items-center gap-3 border-b py-2"
          data-testid={`lane-${lane.guid}`}
        >
          <span class="truncate text-[13px] font-semibold" style={`color: ${lane.color}`}>{lane.name}</span>
          <span class="bg-line-soft relative block h-[18px] w-full">
            {#each lane.auras as segment (segment.start_ms)}
              <span
                class="absolute top-0 h-[6px]"
                style={`left: ${pct(segment.start_ms)}%; width: ${Math.max(pct(segment.end_ms) - pct(segment.start_ms), 0.4)}%; background: ${lane.color}; opacity: .45`}
              ></span>
            {/each}
            {#each lane.casts as at (at)}
              <span
                class="absolute bottom-0 h-[10px] w-[2px]"
                style={`left: ${pct(at)}%; background: ${lane.color}`}
                title={formatDuration(at)}
              ></span>
            {/each}
            {#each lane.deaths as at (at)}
              <span
                class="bg-ember absolute top-0 h-full w-[2px]"
                style={`left: ${pct(at)}%`}
                title={`died at ${formatDuration(at)}`}
              ></span>
            {/each}
          </span>
        </li>
      {/each}
    </ul>
  </div>
{/if}
