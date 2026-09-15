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
    bossName = '',
    players = new Set<string>(),
    allCasts = [],
  }: {
    summary: Summary;
    window: TimeWindow;
    classOf: Map<string, string>;
    /** The encounter's name, which is the boss unit's name; '' on trash. */
    bossName?: string;
    players?: ReadonlySet<string>;
    /** The whole fight's cast rows, enemies included, for the boss lane. */
    allCasts?: Summary['casts'];
  } = $props();

  const span = $derived(Math.max(current.endMs - current.startMs, 1));
  const pct = (ms: number): number => ((ms - current.startMs) / span) * 100;

  /** Axis ticks every 10s, 30s or minute, whichever keeps them under about twelve. */
  const axis = $derived.by(() => {
    const step = span <= 120_000 ? 10_000 : span <= 360_000 ? 30_000 : 60_000;
    const first = Math.ceil(current.startMs / step) * step;
    const ticks: number[] = [];
    for (let at = first; at <= current.endMs; at += step) ticks.push(at);
    return ticks;
  });

  /** The cast last tapped or hovered, read out in the legend for screens with no hover. */
  let picked = $state<{ at: number; name: string } | null>(null);

  /**
   * A tap or a pass of the pointer anywhere on a lane picks the nearest tick: a tick is
   * two pixels wide, and a finger needs the whole lane to be the target.
   */
  function pickNearest(event: PointerEvent, casts: { at: number; name: string }[]): void {
    if (casts.length === 0) return;
    const bounds = (event.currentTarget as HTMLElement).getBoundingClientRect();
    const at = current.startMs + ((event.clientX - bounds.left) / Math.max(bounds.width, 1)) * span;
    let nearest = casts[0];
    for (const cast of casts) if (Math.abs(cast.at - at) < Math.abs(nearest.at - at)) nearest = cast;
    // Within a fortieth of the window: past that the pointer is in a gap, not on a tick.
    if (Math.abs(nearest.at - at) <= span / 40) picked = nearest;
  }

  /** The boss's casts, one tick each, named on hover: the thing to line a death up against. */
  const bossCasts = $derived(
    bossName === ''
      ? []
      : allCasts
          .filter((cast) => !players.has(cast.guid) && cast.name === bossName)
          .flatMap((cast) => cast.sequence.map((at) => ({ at, name: cast.spell_name })))
          // Clipped to the window: the whole fight's casts are read so the lane exists
          // under any scope, but a tick past the window's edge stretches the page.
          .filter((cast) => cast.at >= current.startMs && cast.at <= current.endMs),
  );
  const lanes = $derived(
    [...summary.roster]
      .sort((a, b) => b.damage_done + b.healing_done - (a.damage_done + a.healing_done))
      .map((row) => ({
        guid: row.guid,
        name: splitUnitName(row.name).name,
        color: classColorVar(row.class ?? classOf.get(row.guid)),
        casts: summary.casts
          .filter((cast) => cast.guid === row.guid)
          .flatMap((cast) => cast.sequence.map((at) => ({ at, name: cast.spell_name }))),
        auras: summary.auras
          .filter((track) => track.target_guid === row.guid)
          .flatMap((track) => track.segments.map((segment) => ({ ...segment, name: track.name }))),
        deaths: summary.deaths.filter((death) => death.guid === row.guid).map((death) => death.at_ms),
      })),
  );
</script>

{#if lanes.length === 0}
  <p class="text-muted text-[14px]" data-testid="table-empty">Nobody was in this window.</p>
{:else}
  <div class="flex flex-col gap-1" data-testid="timelines">
    <p class="text-muted label flex flex-wrap items-center gap-x-3">
      <span class="tabular font-mono"
        >{formatDuration(current.startMs)} to {formatDuration(current.endMs)}</span
      >
      <span
        ><span class="bg-text mr-1 inline-block h-[10px] w-[2px] align-middle" aria-hidden="true"
        ></span>cast</span
      >
      <span
        ><span class="bg-text mr-1 inline-block h-[6px] w-[14px] align-middle opacity-45" aria-hidden="true"
        ></span>aura held</span
      >
      <span
        ><span class="bg-death mr-1 inline-block h-[14px] w-[3px] align-middle" aria-hidden="true"
        ></span>death</span
      >
      {#if picked}
        <span class="text-strong normal-case" data-testid="timeline-picked"
          >{picked.name} · {formatDuration(picked.at)}</span
        >
      {:else}
        <span>tap or hover a tick for the spell</span>
      {/if}
      {#if bossCasts.length > 0}
        <span
          ><span class="bg-wipe mr-1 inline-block h-[10px] w-[2px] align-middle" aria-hidden="true"
          ></span>boss cast</span
        >
      {/if}
    </p>
    <div
      class="grid grid-cols-[minmax(96px,140px)_minmax(0,1fr)] gap-3"
      aria-hidden="true"
      data-testid="timeline-axis"
    >
      <span></span>
      <span
        class="text-muted tabular relative block h-[14px] font-mono text-[10px] [&>span:nth-child(even)]:hidden md:[&>span:nth-child(even)]:inline"
      >
        {#each axis as at (at)}
          <span class="absolute top-0 -translate-x-1/2" style={`left: ${pct(at)}%`}>{formatDuration(at)}</span
          >
        {/each}
      </span>
    </div>
    <ul class="flex flex-col">
      {#if bossCasts.length > 0}
        <li
          class="border-line-soft grid min-h-11 grid-cols-[minmax(96px,140px)_minmax(0,1fr)] items-center gap-3 border-b py-2"
          data-testid="lane-boss"
        >
          <span class="text-wipe truncate text-[13px] font-semibold">{bossName}</span>
          <span
            class="bg-line-soft relative block h-[18px] w-full touch-none"
            onpointerdown={(event) => pickNearest(event, bossCasts)}
            onpointermove={(event) => pickNearest(event, bossCasts)}
          >
            {#each bossCasts as cast, i (`${cast.at}-${i}`)}
              <span
                class="bg-wipe absolute bottom-0 h-[14px] w-[2px]"
                style={`left: ${pct(cast.at)}%`}
                title={`${cast.name} · ${formatDuration(cast.at)}`}
              ></span>
            {/each}
          </span>
        </li>
      {/if}
      <!--
        Keys carry the index as well as the timestamp: two casts in one millisecond, two
        auras starting on the same tick and two deaths at the same instant are all real
        (multi-target spells, a trinket and its proc, an AoE wipe), and a timestamp alone
        made Svelte throw on the duplicate and leave the whole view blank.
      -->
      {#each lanes as lane (lane.guid)}
        <li
          class="border-line-soft grid min-h-11 grid-cols-[minmax(96px,140px)_minmax(0,1fr)] items-center gap-3 border-b py-2"
          data-testid={`lane-${lane.guid}`}
        >
          <span class="truncate text-[13px] font-semibold" style={`color: ${lane.color}`}>{lane.name}</span>
          <span
            class="bg-line-soft relative block h-[18px] w-full touch-none"
            onpointerdown={(event) => pickNearest(event, lane.casts)}
            onpointermove={(event) => pickNearest(event, lane.casts)}
          >
            {#each lane.auras as segment, i (`${segment.start_ms}-${i}`)}
              <span
                class="absolute top-0 h-[6px]"
                style={`left: ${pct(segment.start_ms)}%; width: ${Math.max(pct(segment.end_ms) - pct(segment.start_ms), 0.4)}%; background: ${lane.color}; opacity: .45`}
                title={`${segment.name} · ${formatDuration(segment.start_ms)} to ${formatDuration(segment.end_ms)}`}
              ></span>
            {/each}
            {#each lane.casts as cast, i (`${cast.at}-${i}`)}
              <span
                class="absolute bottom-0 h-[10px] w-[2px]"
                style={`left: ${pct(cast.at)}%; background-color: ${lane.color}`}
                title={`${cast.name} · ${formatDuration(cast.at)}`}
              ></span>
            {/each}
            {#each lane.deaths as at, i (`${at}-${i}`)}
              <span
                class="bg-death absolute top-0 h-full w-[3px]"
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
