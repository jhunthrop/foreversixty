<!-- web/src/components/report/TimelinesView.svelte -->
<!-- One lane per player: their casts as ticks, the auras they held as bars, and their
     death as a full-height mark. The same window the tables use, drawn horizontally, so
     "who was doing nothing for eight seconds" is a glance rather than a calculation.

     Every source drawn here -- cast sequence, aura segments, deaths -- is exact under a
     window (window.ts's scopeSummary), so nothing on this lane needs an approximate mark.
     Drawn to round 5 (design/mocks/directions/report.py, b1-report-timeline): the panel
     `Timeline` with the window as its meta, the Show chips (each its mark's swatch, so the row is
     the legend), a 180 px label column (crest and name in the class colour, the boss first), one
     lane per actor on one time axis (casts gold, buffs green, debuffs red, a death red, a taunt
     gold full height), the axis under. Each lane stacks only its own auras, one 3 px band each.
     Below 768 the labels are 124 px and the plot scrolls sideways while the labels stay put. The
     summary carries no per-player hit stream, so the lanes draw no hit marks. -->
<script lang="ts">
  import EmptyState from '../ui/EmptyState.svelte';
  import { splitUnitName } from '../../lib/characters';
  import { classColorVar, tauntKey } from '../../lib/report/format';
  import { classSlugOf } from '../../lib/report/planner-link';
  import type { AuraTrack, Summary, Taunt } from '../../lib/report/types';
  import type { TimeWindow } from '../../lib/report/window';
  import { clearChartCursor, setChartCursor } from '../../lib/chart/cursor.svelte';
  import { clockLabel, timeTicks } from '../../lib/chart/time-axis';
  import { CHART_INK, inkCss } from '../../lib/chart/theme';
  import { clockDuration } from '../../lib/format/duration';
  import SyncedCrosshair from '../chart/SyncedCrosshair.svelte';
  import Crest from '../ui/Crest.svelte';
  import Chips from '../ui/Chips.svelte';
  import Section from '../ui/Section.svelte';

  let {
    summary,
    window: current,
    classOf,
    bossName = '',
    auraOrder = [],
    players = new Set<string>(),
    allCasts = [],
    taunts = [],
    syncKey = '',
  }: {
    summary: Summary;
    window: TimeWindow;
    classOf: Map<string, string>;
    /** Every aura name in the whole fight, in order, so a lane's bands keep their order when the window moves. */
    auraOrder?: string[];
    /** The encounter's name, which is the boss unit's name; '' on trash. */
    bossName?: string;
    players?: ReadonlySet<string>;
    /** The whole fight's cast rows, enemies included, for the boss lane. */
    allCasts?: Summary['casts'];
    /** Who taunted what, when: a mark on the taunter's lane and on the boss lane. */
    taunts?: Taunt[];
    /** The page's chart cursor key (the fight): the lanes share the charts' crosshair. Empty over a night. */
    syncKey?: string;
  } = $props();
  const owner = 'timelines';

  const span = $derived(Math.max(current.endMs - current.startMs, 1));
  /** The boss's phase changes inside the window: a dashed line down every lane, as the charts draw them. */
  const PHASE_INK = inkCss(CHART_INK.phase);
  const phaseStarts = $derived(
    (summary.phases ?? []).filter(
      (phase) => phase.start_ms > current.startMs && phase.start_ms < current.endMs,
    ),
  );
  const pct = (ms: number): number => ((ms - current.startMs) / span) * 100;

  /** The time axis's ticks (lib/chart/time-axis: the step that keeps labels 64 px apart). */
  let axisWidth = $state(600);
  const axis = $derived(
    timeTicks(current.startMs / 1000, current.endMs / 1000, axisWidth).map((seconds) => seconds * 1000),
  );

  /** Which marks the lanes draw (the Show chips); every kind on by default. */
  type MarkKind = 'cast' | 'buff' | 'debuff' | 'death' | 'taunt';
  let shown = $state<Record<MarkKind, boolean>>({
    cast: true,
    buff: true,
    debuff: true,
    death: true,
    taunt: true,
  });
  function toggleMark(id: string): void {
    if (id in shown) shown = { ...shown, [id]: !shown[id as MarkKind] };
  }

  /** The cast last tapped or hovered, read out in the legend for screens with no hover. */
  let picked = $state<{ at: number; name: string } | null>(null);

  /**
   * A tap or a pass of the pointer anywhere on a lane picks the nearest tick: a tick is
   * two pixels wide, and a finger needs the whole lane to be the target.
   */
  /** A mouse leaving the lane takes its readout with it; a tap's readout stays for reading. */
  function clearPick(event: PointerEvent): void {
    if (event.pointerType !== 'mouse') return;
    picked = null;
    clearChartCursor(owner);
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
    if (syncKey !== '') setChartCursor({ syncKey, ms: Math.round(at / 1000) * 1000, owner });
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
          name: `${under.name} · ${clockDuration(under.start_ms)} to ${clockDuration(under.end_ms)}`,
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
    const taunt = nearest(laneTaunts, (one) => one.at_ms, at);
    if (taunt !== null) {
      picked = { at: taunt.at_ms, name: tauntReadout(taunt) };
      return;
    }
    const cast = nearest(casts, (one) => one.at, at);
    if (cast !== null) picked = cast;
  }

  /**
   * The item under the pointer, or null when the pointer is in a gap: within a fortieth
   * of the window is on a mark, past that is between them. Both lanes ask the same
   * question of two differently shaped lists, so `atOf` reads the time out of each.
   */
  function nearest<T>(list: T[], atOf: (item: T) => number, at: number): T | null {
    if (list.length === 0) return null;
    let closest = list[0];
    for (const item of list) if (Math.abs(atOf(item) - at) < Math.abs(atOf(closest) - at)) closest = item;
    return Math.abs(atOf(closest) - at) <= span / 40 ? closest : null;
  }

  /** "Taunt · <spell> on <target> by <taunter>", the pull's label appended on the night. */
  function tauntReadout(taunt: Taunt): string {
    const target = splitUnitName(taunt.target_name).name;
    const source = splitUnitName(taunt.source_name).name;
    const pull = taunt.label === undefined ? '' : ` · ${taunt.label}`;
    const before = taunt.pre_pull ? ' · cast before the pull' : '';
    return `Taunt · ${taunt.spell_name} on ${target} by ${source}${pull}${before}`;
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
  /** One enemy's lane: its casts as ticks and the taunts that landed on it. */
  interface EnemyLane {
    name: string;
    casts: { at: number; name: string }[];
    taunts: Taunt[];
  }
  /**
   * The enemies with a lane: the boss, and every other enemy someone taunted. An add that
   * took nine of a pull's ten taunts and killed four people is what the taunt marks point
   * at, and a mark pointing at a creature with no lane is a mark pointing at nothing.
   */
  const enemyLanes = $derived.by<EnemyLane[]>(() => {
    const lanes: EnemyLane[] = [];
    if (bossCasts.length > 0 || bossTaunts.length > 0) {
      lanes.push({ name: bossName, casts: bossCasts, taunts: bossTaunts });
    }
    const others = [
      ...new Set(
        windowedTaunts
          .map((taunt) => splitUnitName(taunt.target_name).name)
          .filter((name) => name !== '' && name !== bossName),
      ),
    ].sort((a, b) => a.localeCompare(b));
    for (const name of others) {
      lanes.push({
        name,
        casts: allCasts
          .filter((cast) => !players.has(cast.guid) && splitUnitName(cast.name).name === name)
          .flatMap((cast) => cast.sequence.map((at) => ({ at, name: cast.spell_name })))
          .filter((cast) => cast.at >= current.startMs && cast.at <= current.endMs),
        taunts: windowedTaunts.filter((taunt) => splitUnitName(taunt.target_name).name === name),
      });
    }
    return lanes;
  });
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
    /** A buff is green, a debuff red (the legend's Buff up and Debuff up). */
    debuff: boolean;
  }
  const BAND_PX = 3;
  const CASTS_PX = 12;
  /** A lane's row: its plot's height plus ROW_PAD_PX, never under 30 px. */
  const ROW_MIN_PX = 30;
  /** 8 px under the cast line, the rest above the bands: a clear gap between one lane and the next. */
  const ROW_PAD_PX = 14;
  const ENEMY_LANE_PX = 18;
  /** A tick this close to the right edge ends there instead of centring past it. */
  const END_LABEL_PCT = 97;
  /**
   * Every aura the lane holds on its own band, so each reads as one line and can be pointed at.
   * Only the lane's own auras take a band (a lane never reserves rows for another's), ordered by
   * the fight's aura order, so an aura keeps its place when the window moves.
   */
  function bandAuras(tracks: AuraTrack[], order: readonly string[] = auraOrder): AuraBand[] {
    const visible = tracks.filter((track) => (track.type === 'DEBUFF' ? shown.debuff : shown.buff));
    const rank = (name: string): number => {
      const index = order.indexOf(name);
      return index === -1 ? order.length : index;
    };
    const names = [...new Set(visible.map((track) => track.name))].sort(
      (a, b) => rank(a) - rank(b) || a.localeCompare(b),
    );
    return visible.flatMap((track) => {
      const band = names.indexOf(track.name);
      const debuff = track.type === 'DEBUFF';
      return track.segments.map((segment) => ({ ...segment, name: track.name, band, debuff }));
    });
  }
  const laneHeight = (auras: AuraBand[]): number =>
    Math.max(18, (auras.reduce((top, aura) => Math.max(top, aura.band), -1) + 1) * BAND_PX + CASTS_PX);
  const rowHeight = (plot: number): number => Math.max(ROW_MIN_PX, plot + ROW_PAD_PX);
  /**
   * The enemy lanes with the group's debuffs on them (Sunder Armor, Thunder Clap on the boss),
   * each lane's own bands in name order, every enemy lane at the tallest one's height.
   */
  const enemyRows = $derived(
    enemyLanes.map((lane) => {
      const auras = bandAuras(
        summary.auras.filter(
          (track) =>
            track.type === 'DEBUFF' &&
            !players.has(track.target_guid) &&
            splitUnitName(track.target_name).name === lane.name,
        ),
        [],
      );
      return { ...lane, auras, height: Math.max(ENEMY_LANE_PX, laneHeight(auras)) };
    }),
  );
  const enemyRow = $derived(Math.max(ENEMY_LANE_PX, ...enemyRows.map((lane) => lane.height)));

  const lanes = $derived(
    [...summary.roster]
      .sort((a, b) => b.damage_done + b.healing_done - (a.damage_done + a.healing_done))
      .map((row) => ({
        guid: row.guid,
        name: splitUnitName(row.name).name,
        color: classColorVar(row.class ?? classOf.get(row.guid)),
        crest: classSlugOf(row.class ?? classOf.get(row.guid)) ?? '',
        casts: summary.casts
          .filter((cast) => cast.guid === row.guid)
          .flatMap((cast) => cast.sequence.map((at) => ({ at, name: cast.spell_name }))),
        auras: bandAuras(summary.auras.filter((track) => track.target_guid === row.guid)),
        deaths: summary.deaths.filter((death) => death.guid === row.guid).map((death) => death.at_ms),
        taunts: windowedTaunts.filter((taunt) => taunt.source_guid === row.guid),
      }))
      .map((lane) => ({ ...lane, height: laneHeight(lane.auras) })),
  );
  /** Every player lane at the tallest one's height, so the rows run on an even pitch. */
  const laneRow = $derived(Math.max(18, ...lanes.map((lane) => lane.height)));
  /** One pitch for every lane, enemy and player, so the rows run evenly down the plot. */
  const pitch = $derived(Math.max(laneRow, enemyRow));
  /** The Show chips, each with its mark's colour: the row is the legend. Taunts only when someone taunted. */
  const MARKS: readonly { id: MarkKind; label: string; swatch: string }[] = [
    { id: 'cast', label: 'Casts', swatch: 'var(--color-gold)' },
    { id: 'buff', label: 'Buffs', swatch: 'color-mix(in srgb, var(--color-kill) 60%, transparent)' },
    { id: 'debuff', label: 'Debuffs', swatch: 'color-mix(in srgb, var(--color-wipe) 60%, transparent)' },
    { id: 'death', label: 'Deaths', swatch: 'var(--color-death)' },
    { id: 'taunt', label: 'Taunts', swatch: 'var(--color-accent)' },
  ];
  const chips = $derived(
    MARKS.filter((mark) => mark.id !== 'taunt' || windowedTaunts.length > 0).map((mark) => ({
      ...mark,
      on: shown[mark.id],
    })),
  );
</script>

{#snippet crossing()}{#if syncKey !== ''}<SyncedCrosshair {syncKey} window={current} />{/if}{/snippet}

<Section
  title="Timeline"
  meta={`${clockLabel(current.startMs)} to ${clockLabel(current.endMs)}`}
  testid="timelines-section"
  panel
  flush
>
  <div>
    {#if lanes.length === 0}
      <EmptyState inset message="Nothing in this window." testid="table-empty" />
    {:else}
      <div class="flex flex-col gap-2 px-4 pb-3 md:px-5" data-testid="timelines">
        <Chips {chips} label="Show" onToggle={toggleMark} testid="timeline-show" />
        <!-- The readout on one line of fixed height: eight aura names must never wrap and push
             the lanes out from under the pointer that is reading them. -->
        <p
          class="text-muted h-5 truncate text-[13px] leading-5"
          title={picked === null ? undefined : `${picked.name} · at ${clockDuration(picked.at)}`}
        >
          {#if picked}
            <span class="text-strong font-semibold" data-testid="timeline-picked"
              >{picked.name} · at {clockDuration(picked.at)}</span
            >
          {:else}
            <span aria-hidden="true">&nbsp;</span>
          {/if}
        </p>
        <div class="grid grid-cols-[124px_minmax(0,1fr)] md:grid-cols-[180px_minmax(0,1fr)]">
          <!-- The labels stay put; only the plot scrolls sideways on a phone. -->
          <ul class="flex flex-col pr-3">
            {#each enemyRows as lane (lane.name)}
              <!-- Every label sits on its lane's cast line, the bands rising above it. -->
              <li
                class="flex items-end gap-2 overflow-hidden pb-[3px]"
                style:height="{rowHeight(pitch)}px"
                title={lane.name}
              >
                <Crest cls="" size={20} />
                <span
                  class="text-strong min-w-0 text-[13px] leading-4 font-semibold md:truncate md:text-[14px] md:leading-5"
                  >{lane.name}</span
                >
              </li>
            {/each}
            {#each lanes as lane (lane.guid)}
              <li
                class="flex items-end gap-2 overflow-hidden pb-[3px]"
                style:height="{rowHeight(pitch)}px"
                title={lane.name}
              >
                <Crest cls={lane.crest} size={20} />
                <span
                  class="min-w-0 text-[13px] leading-4 font-semibold md:truncate md:text-[14px] md:leading-5"
                  style:color={lane.color}>{lane.name}</span
                >
              </li>
            {/each}
          </ul>
          <div class="min-w-0 overflow-x-auto">
            <div class="relative" aria-hidden="true">
              {#each phaseStarts as phase (phase.name + phase.start_ms)}
                <span
                  class="pointer-events-none absolute inset-y-0 z-[1] w-0 border-l border-dashed"
                  style:left="{pct(phase.start_ms)}%"
                  style:border-color={PHASE_INK}
                  data-testid="timeline-phase"
                  title={phase.name}
                ></span>
              {/each}
              <ul class="flex flex-col">
                {#each enemyRows as lane (lane.name)}
                  <li class="flex items-end pb-2" style:height="{rowHeight(pitch)}px">
                    <!-- svelte-ignore a11y_no_static_element_interactions (the plot is decorative, aria-hidden: the pointer readout is a sighted aid) -->
                    <span
                      class="relative block w-full touch-none"
                      style:height="{lane.height}px"
                      data-lane
                      data-testid={lane.name === bossName ? 'lane-boss' : `lane-enemy-${lane.name}`}
                      onpointerdown={(event) => pickNearest(event, lane.casts, lane.auras, lane.taunts)}
                      onpointermove={(event) => pickNearest(event, lane.casts, lane.auras, lane.taunts)}
                      onpointerleave={clearPick}
                    >
                      <span class="bg-line-soft absolute inset-x-0 bottom-[4px] block h-[2px]"></span>
                      {#each lane.auras as segment, i (`${segment.start_ms}-${i}`)}
                        <span
                          class="bg-wipe/60 absolute h-[3px] rounded-[1px]"
                          style:top="{segment.band * BAND_PX}px"
                          style:left="{pct(segment.start_ms)}%"
                          style:width="{Math.max(pct(segment.end_ms) - pct(segment.start_ms), 0.4)}%"
                          title={`${segment.name} · ${clockDuration(segment.start_ms)} to ${clockDuration(segment.end_ms)}`}
                        ></span>
                      {/each}
                      {#each shown.cast ? lane.casts : [] as cast, i (`${cast.at}-${i}`)}
                        <span
                          class="bg-gold absolute bottom-0 h-[10px] w-[2px]"
                          style:left="{pct(cast.at)}%"
                          title={`${cast.name} · ${clockDuration(cast.at)}`}
                        ></span>
                      {/each}
                      {#each shown.taunt ? lane.taunts : [] as taunt (tauntKey(taunt))}
                        <span
                          class="bg-accent absolute -top-[4px] h-[calc(100%+8px)] w-[3px]"
                          style:left="{pct(taunt.at_ms)}%"
                          data-testid="taunt-mark"
                          title={`Taunt · ${clockDuration(taunt.at_ms)} · by ${splitUnitName(taunt.source_name).name}`}
                        ></span>
                      {/each}
                      {@render crossing()}
                    </span>
                  </li>
                {/each}
                <!--
                  Keys carry the index as well as the timestamp: two casts in one millisecond, two
                  auras starting on the same tick and two deaths at the same instant are all real
                  (multi-target spells, a trinket and its proc, an AoE wipe), and a timestamp alone
                  made Svelte throw on the duplicate and leave the whole view blank.
                -->
                {#each lanes as lane (lane.guid)}
                  <li class="flex items-end pb-2" style:height="{rowHeight(pitch)}px">
                    <!-- svelte-ignore a11y_no_static_element_interactions (the plot is decorative, aria-hidden: the pointer readout is a sighted aid) -->
                    <span
                      class="relative block w-full touch-none"
                      style:height="{lane.height}px"
                      data-lane
                      data-testid={`lane-${lane.guid}`}
                      onpointerdown={(event) => pickNearest(event, lane.casts, lane.auras, lane.taunts)}
                      onpointermove={(event) => pickNearest(event, lane.casts, lane.auras, lane.taunts)}
                      onpointerleave={clearPick}
                    >
                      <span class="bg-line-soft absolute inset-x-0 bottom-[4px] block h-[2px]"></span>
                      {#each lane.auras as segment, i (`${segment.start_ms}-${i}`)}
                        <span
                          class={[
                            'absolute h-[3px] rounded-[1px]',
                            segment.debuff ? 'bg-wipe/60' : 'bg-kill/60',
                          ]}
                          style:top="{segment.band * BAND_PX}px"
                          style:left="{pct(segment.start_ms)}%"
                          style:width="{Math.max(pct(segment.end_ms) - pct(segment.start_ms), 0.4)}%"
                          title={`${segment.name} · ${clockDuration(segment.start_ms)} to ${clockDuration(segment.end_ms)}`}
                        ></span>
                      {/each}
                      {#each shown.cast ? lane.casts : [] as cast, i (`${cast.at}-${i}`)}
                        <span
                          class="bg-gold absolute bottom-0 h-[10px] w-[2px]"
                          style:left="{pct(cast.at)}%"
                          title={`${cast.name} · ${clockDuration(cast.at)}`}
                        ></span>
                      {/each}
                      {#each shown.death ? lane.deaths : [] as at, i (`${at}-${i}`)}
                        <span
                          class="bg-death absolute top-0 h-full w-[3px]"
                          style:left="{pct(at)}%"
                          title={`died at ${clockDuration(at)}`}
                        ></span>
                      {/each}
                      {#each shown.taunt ? lane.taunts : [] as taunt (tauntKey(taunt))}
                        <span
                          class="bg-accent absolute -top-[4px] h-[calc(100%+8px)] w-[3px]"
                          style:left="{pct(taunt.at_ms)}%"
                          data-testid="taunt-mark"
                          title={`Taunt · ${clockDuration(taunt.at_ms)} · on ${splitUnitName(taunt.target_name).name}`}
                        ></span>
                      {/each}
                      {@render crossing()}
                    </span>
                  </li>
                {/each}
              </ul>
              <span
                class="text-muted relative mt-1 block h-4 text-[12px] leading-4 font-semibold tabular-nums"
                bind:clientWidth={axisWidth}
                data-testid="timeline-axis"
              >
                {#each axis as at (at)}
                  <span
                    class={[
                      'absolute top-0 whitespace-nowrap',
                      // The first label starts at the edge, the last ends at it; the rest are centred.
                      at === axis[0]
                        ? ''
                        : pct(at) >= END_LABEL_PCT
                          ? '-translate-x-full'
                          : '-translate-x-1/2',
                    ]}
                    style:left="{pct(at)}%">{clockLabel(at)}</span
                  >
                {/each}
              </span>
            </div>
          </div>
        </div>
      </div>
    {/if}
  </div>
</Section>
