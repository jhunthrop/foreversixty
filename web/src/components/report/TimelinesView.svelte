<!-- web/src/components/report/TimelinesView.svelte -->
<!-- One lane per player: their casts as ticks, the auras they held as bars, and their
     death as a full-height mark. The same window the tables use, drawn horizontally, so
     "who was doing nothing for eight seconds" is a glance rather than a calculation.

     Every source drawn here -- cast sequence, aura segments, deaths -- is exact under a
     window (window.ts's scopeSummary), so nothing on this lane needs an approximate mark. -->
<script lang="ts">
  import { splitUnitName } from '../../lib/characters';
  import { classColorVar, formatDuration, tauntKey } from '../../lib/report/format';
  import type { AuraTrack, Summary, Taunt } from '../../lib/report/types';
  import type { TimeWindow } from '../../lib/report/window';

  let {
    summary,
    window: current,
    classOf,
    bossName = '',
    auraOrder = [],
    players = new Set<string>(),
    allCasts = [],
    taunts = [],
  }: {
    summary: Summary;
    window: TimeWindow;
    classOf: Map<string, string>;
    /** Every aura name in the whole fight, in order, so a band keeps its row when the window moves. */
    auraOrder?: string[];
    /** The encounter's name, which is the boss unit's name; '' on trash. */
    bossName?: string;
    players?: ReadonlySet<string>;
    /** The whole fight's cast rows, enemies included, for the boss lane. */
    allCasts?: Summary['casts'];
    /** Who taunted what, when: a mark on the taunter's lane and on the boss lane. */
    taunts?: Taunt[];
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
  /** A mouse leaving the lane takes its readout with it; a tap's readout stays for reading. */
  function clearPick(event: PointerEvent): void {
    if (event.pointerType === 'mouse') picked = null;
  }

  function pickNearest(
    event: PointerEvent,
    casts: { at: number; name: string }[],
    auras: AuraBand[] = [],
    laneTaunts: Taunt[] = [],
  ): void {
    // The lane by its mark, not currentTarget: a delegated pointer event can hand over the
    // island's root, whose width made every readout land a few seconds off the pointer.
    const lane = (event.target as HTMLElement).closest<HTMLElement>('[data-lane]');
    if (lane === null) return;
    const bounds = lane.getBoundingClientRect();
    const at = current.startMs + ((event.clientX - bounds.left) / Math.max(bounds.width, 1)) * span;
    // The aura bands run from the lane's top, one per aura: the band under the pointer
    // names its aura and the segment's span; a gap in it says what else was up.
    const y = event.clientY - bounds.top;
    const bandCount = auras.reduce((top, aura) => Math.max(top, aura.band), -1) + 1;
    if (y < bandCount * BAND_PX) {
      const band = Math.floor(y / BAND_PX);
      const under = auras.find((aura) => aura.band === band && aura.start_ms <= at && at <= aura.end_ms);
      if (under !== undefined) {
        picked = {
          at,
          name: `${under.name} · ${formatDuration(under.start_ms)} to ${formatDuration(under.end_ms)}`,
        };
        return;
      }
      const held = auras.filter((aura) => aura.start_ms <= at && at <= aura.end_ms);
      if (held.length > 0) {
        picked = { at, name: [...new Set(held.map((aura) => aura.name))].join(', ') };
        return;
      }
    }
    // A taunt within reach beats a cast: it is the rarer, more legible event, and a mark
    // drawn right beside a cast tick would otherwise always lose to the tick underneath it.
    if (laneTaunts.length > 0) {
      let nearestTaunt = laneTaunts[0];
      for (const taunt of laneTaunts)
        if (Math.abs(taunt.at_ms - at) < Math.abs(nearestTaunt.at_ms - at)) nearestTaunt = taunt;
      if (Math.abs(nearestTaunt.at_ms - at) <= span / 40) {
        picked = { at: nearestTaunt.at_ms, name: tauntReadout(nearestTaunt) };
        return;
      }
    }
    if (casts.length === 0) return;
    let nearest = casts[0];
    for (const cast of casts) if (Math.abs(cast.at - at) < Math.abs(nearest.at - at)) nearest = cast;
    // Within a fortieth of the window: past that the pointer is in a gap, not on a tick.
    if (Math.abs(nearest.at - at) <= span / 40) picked = nearest;
  }

  /** "Taunt · <spell> on <target> by <taunter>", the pull's label appended on the night. */
  function tauntReadout(taunt: Taunt): string {
    const target = splitUnitName(taunt.target_name).name;
    const source = splitUnitName(taunt.source_name).name;
    const pull = taunt.label === undefined ? '' : ` · ${taunt.label}`;
    return `Taunt · ${taunt.spell_name} on ${target} by ${source}${pull}`;
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
  /** Taunts clipped to the window, the same way a cast or a death is. */
  const windowedTaunts = $derived(
    taunts.filter((taunt) => taunt.at_ms >= current.startMs && taunt.at_ms <= current.endMs),
  );
  /**
   * The boss's own mark: every taunt that landed on it, wherever the taunter's lane is --
   * matched by name rather than guid, the same way bossCasts is, since an add's guid is
   * new on every pull.
   */
  const bossTaunts = $derived(
    bossName === ''
      ? []
      : windowedTaunts.filter((taunt) => splitUnitName(taunt.target_name).name === bossName),
  );
  /**
   * Every segment of a player's auras, each aura kept on one of three bands for the whole
   * lane, so an aura reads as one line across the fight rather than hopping bands between
   * occurrences.
   */
  interface AuraBand {
    start_ms: number;
    end_ms: number;
    name: string;
    band: number;
    color: string;
  }
  const BAND_PX = 3;
  const CASTS_PX = 12;
  /** A hue of its own per aura name, stable across lanes and fights. */
  function hueOf(name: string): number {
    let hash = 0;
    for (const char of name) hash = (hash * 31 + char.charCodeAt(0)) % 360;
    return hash;
  }
  /** Every aura on its own band, in name order, so each reads as one line and can be pointed at. */
  function bandAuras(tracks: AuraTrack[]): AuraBand[] {
    // By name, not by uptime: an aura keeps its row when the window moves.
    const ordered = [...tracks].sort((a, b) => a.name.localeCompare(b.name));
    // eslint-disable-next-line svelte/prefer-svelte-reactivity
    const bands = new Map<string, number>(auraOrder.map((name, index) => [name, index]));
    return ordered.flatMap((track) => {
      const band = bands.get(track.name) ?? bands.size;
      bands.set(track.name, band);
      const color = `hsl(${hueOf(track.name)} 65% 62%)`;
      return track.segments.map((segment) => ({ ...segment, name: track.name, band, color }));
    });
  }
  const laneHeight = (auras: AuraBand[]): number =>
    Math.max(18, (auras.reduce((top, aura) => Math.max(top, aura.band), -1) + 1) * BAND_PX + CASTS_PX);

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
        auras: bandAuras(summary.auras.filter((track) => track.target_guid === row.guid)),
        deaths: summary.deaths.filter((death) => death.guid === row.guid).map((death) => death.at_ms),
        taunts: windowedTaunts.filter((taunt) => taunt.source_guid === row.guid),
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
      {#if bossCasts.length > 0}
        <span
          ><span class="bg-wipe mr-1 inline-block h-[10px] w-[2px] align-middle" aria-hidden="true"
          ></span>boss cast</span
        >
      {/if}
      {#if windowedTaunts.length > 0}
        <span
          ><span class="bg-gold mr-1 inline-block h-[14px] w-[3px] align-middle" aria-hidden="true"
          ></span>taunt</span
        >
      {/if}
    </p>
    <!-- The readout on one line of fixed height: eight aura names must never wrap and push
         the lanes out from under the pointer that is reading them. -->
    <p
      class="text-muted label h-5 truncate"
      title={picked === null ? undefined : `${picked.name} · at ${formatDuration(picked.at)}`}
    >
      {#if picked}
        <span class="text-strong normal-case" data-testid="timeline-picked"
          >{picked.name} · at {formatDuration(picked.at)}</span
        >
      {:else}
        <span>tap or hover a tick for the spell, or an aura band for what was up</span>
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
      {#if bossCasts.length > 0 || bossTaunts.length > 0}
        <li
          class="border-line-soft grid min-h-11 grid-cols-[minmax(96px,140px)_minmax(0,1fr)] items-center gap-3 border-b py-2"
          data-testid="lane-boss"
        >
          <span class="text-wipe truncate text-[13px] font-semibold">{bossName}</span>
          <span
            class="bg-line-soft relative block h-[18px] w-full touch-none"
            data-lane
            onpointerdown={(event) => pickNearest(event, bossCasts, [], bossTaunts)}
            onpointermove={(event) => pickNearest(event, bossCasts, [], bossTaunts)}
            onpointerleave={clearPick}
          >
            {#each bossCasts as cast, i (`${cast.at}-${i}`)}
              <span
                class="bg-wipe absolute bottom-0 h-[14px] w-[2px]"
                style={`left: ${pct(cast.at)}%`}
                title={`${cast.name} · ${formatDuration(cast.at)}`}
              ></span>
            {/each}
            {#each bossTaunts as taunt (tauntKey(taunt))}
              <span
                class="bg-gold absolute top-0 h-full w-[3px]"
                style={`left: ${pct(taunt.at_ms)}%`}
                data-testid="taunt-mark"
                title={`Taunt · ${formatDuration(taunt.at_ms)} · by ${splitUnitName(taunt.source_name).name}`}
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
            class="bg-line-soft relative block w-full touch-none"
            style={`height: ${laneHeight(lane.auras)}px`}
            data-lane
            onpointerdown={(event) => pickNearest(event, lane.casts, lane.auras, lane.taunts)}
            onpointermove={(event) => pickNearest(event, lane.casts, lane.auras, lane.taunts)}
            onpointerleave={clearPick}
          >
            {#each lane.auras as segment, i (`${segment.start_ms}-${i}`)}
              <span
                class="absolute h-[3px]"
                style={`top: ${segment.band * BAND_PX}px; left: ${pct(segment.start_ms)}%; width: ${Math.max(pct(segment.end_ms) - pct(segment.start_ms), 0.4)}%; background: ${segment.color}`}
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
            {#each lane.taunts as taunt (tauntKey(taunt))}
              <span
                class="bg-gold absolute top-0 h-full w-[3px]"
                style={`left: ${pct(taunt.at_ms)}%`}
                data-testid="taunt-mark"
                title={`Taunt · ${formatDuration(taunt.at_ms)} · on ${splitUnitName(taunt.target_name).name}`}
              ></span>
            {/each}
          </span>
        </li>
      {/each}
    </ul>
    {#if auraOrder.length > 0}
      <!-- The bands are 3px each, too thin to label in place: this names each row's aura in
           its colour, in the order the rows run from the top of every lane. -->
      <details class="text-[12px]" data-testid="timeline-aura-legend">
        <summary class="label text-muted min-h-11 cursor-pointer md:min-h-0"
          >Aura rows, top to bottom ({auraOrder.length})</summary
        >
        <ol class="mt-1 flex flex-wrap gap-x-4 gap-y-1">
          {#each auraOrder as name, index (name)}
            <li class="flex items-center gap-1">
              <span class="text-muted tabular font-mono text-[11px]">{index + 1}</span>
              <span class="inline-block h-[3px] w-[14px]" style={`background: hsl(${hueOf(name)} 65% 62%)`}
              ></span>
              {name}
            </li>
          {/each}
        </ol>
      </details>
    {/if}
  </div>
{/if}
