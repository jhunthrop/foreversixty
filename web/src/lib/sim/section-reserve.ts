// web/src/lib/sim/section-reserve.ts
// How tall each lazy card of the saved result will be, held only while its chunk loads (spec
// 2026-11-04-sim-results 4.2 and 8): computed from this result's own data where the card's height
// follows it (the timeline's lanes, the rage sources, the aura rows, the kit, the chips), from the
// card's measured frame elsewhere. LazySection drops the reserve the moment its card renders or
// fails, so an estimate that is off moves only what is still below the fold, never leaves a hole.
//
// The frame figures were measured on the built page (2026-11-05, `simfullfixt1` and
// `simfixtureab`, at 390 for phone and 1440 for wide); the timeline's rows come from the very lane
// rules the card draws with (lib/sim/timeline-lanes.ts), so its reserve is its ready height.
import { COLLAPSED_CAST_ROWS, plotHeightOf, timelineShape } from './timeline-lanes';
import type { SimResult } from './types';

/** A card's reserved height under 768 px and from 768 px. */
export interface SectionReserve {
  phone: number;
  wide: number;
}

export type LazySectionId = 'timeline' | 'gear' | 'talents' | 'resource' | 'buffs' | 'rotation' | 'settings';

type Widths<T> = Readonly<{ phone: T; wide: T }>;

/** A card's own frame: its 1 px border top and bottom. */
const CARD_BORDERS = 2;
/** A CardHead: 68 px plus its rule; on a phone a meta or a long label wraps under the title. */
const HEAD: Widths<number> = { phone: 71, wide: 71 };
const HEAD_WRAPPED: Widths<number> = { phone: 98.5, wide: 71 };
/** One list row in the Rage and Buffs cards (`min-h-10`). */
const LIST_ROW = 40;

const TIMELINE: Widths<{ toolbar: number; castFoot: number }> = {
  // The toolbar wraps to three lines on a phone; the cast list's foot wraps to two.
  phone: { toolbar: 158, castFoot: 108 },
  wide: { toolbar: 59, castFoot: 53 },
};
/** The plot's top rule, the brush row and its rule, the live region. */
const TIMELINE_FIXED = 1 + 45 + 1;

function timelineHeight(result: SimResult, phone: boolean): number | null {
  const shape = timelineShape(result);
  if (shape === null) return null;
  const at = phone ? TIMELINE.phone : TIMELINE.wide;
  return (
    CARD_BORDERS +
    (phone ? HEAD_WRAPPED.phone : HEAD_WRAPPED.wide) +
    at.toolbar +
    plotHeightOf(shape, phone) +
    TIMELINE_FIXED +
    at.castFoot +
    LIST_ROW * Math.min(COLLAPSED_CAST_ROWS, shape.presses)
  );
}

/** Rage: three tiles, then the two source lists side by side (stacked on a phone), or one line. */
const RESOURCE: Widths<{ tiles: number; lists: number; needsNewer: number }> = {
  phone: { tiles: 121, lists: 110, needsNewer: 75 },
  wide: { tiles: 118.5, lists: 49.5, needsNewer: 54 },
};

function resourceHeight(result: SimResult, phone: boolean): number {
  const at = phone ? RESOURCE.phone : RESOURCE.wide;
  const sources = result.resource_sources?.[0];
  const frame = CARD_BORDERS + HEAD.wide + at.tiles;
  if (sources === undefined) return frame + at.needsNewer;
  const gained = sources.gained.length;
  const spent = sources.spent.length;
  return frame + at.lists + LIST_ROW * (phone ? gained + spent : Math.max(gained, spent));
}

/** Buffs: one bar row per aura the run tracked that came and went; the rest is the frame. */
const BUFFS_FRAME: Widths<number> = { phone: 790.5, wide: 342.5 };
/** The Overpower window is the rotation's, not a buff: the Buffs card leaves it out (aura-card.ts). */
const OVERPOWER_WINDOWS: ReadonlySet<string> = new Set([
  'spell:7384',
  'spell:7887',
  'spell:11584',
  'spell:11585',
]);
/** The Buffs card files an aura at this share or more under "Up all fight", not as a bar. */
const BUFF_UP_ALL_FIGHT = 0.999;

function buffsHeight(result: SimResult, phone: boolean): number {
  const duration = result.summary.duration_ms;
  const bars = (type: string): number =>
    result.summary.auras.filter(
      (track) =>
        track.type === type &&
        track.uptime_ms > 0 &&
        // A swing's window and a queue's own window are never bars.
        !track.name.startsWith('other:') &&
        !track.name.includes('/') &&
        !OVERPOWER_WINDOWS.has(track.name) &&
        track.uptime_ms < duration * BUFF_UP_ALL_FIGHT,
    ).length;
  const player = bars('BUFF');
  const target = bars('DEBUFF');
  return (
    (phone ? BUFFS_FRAME.phone : BUFFS_FRAME.wide) +
    LIST_ROW * (phone ? player + target : Math.max(player, target))
  );
}

/** Gear: the paperdoll is the same frame for any kit; enchant lines and set feet grow with it. */
const GEAR: Widths<{ base: number; perItem: number }> = {
  phone: { base: 901.7, perItem: 17.1 },
  wide: { base: 912.6, perItem: 10.2 },
};

function gearHeight(result: SimResult, phone: boolean): number | null {
  const items = result.request.character.gear.length;
  if (items === 0) return null;
  const at = phone ? GEAR.phone : GEAR.wide;
  return at.base + at.perItem * items;
}

/** Settings: the rows, then a chip per buff and consumable (a phone shows six a group, then Show all). */
const SETTINGS: Widths<{ base: number; perChip: number; chipsPerGroup: number }> = {
  phone: { base: 687.5, perChip: 63, chipsPerGroup: 6 },
  wide: { base: 423.2, perChip: 11.9, chipsPerGroup: Number.POSITIVE_INFINITY },
};

function settingsHeight(result: SimResult, phone: boolean): number {
  const at = phone ? SETTINGS.phone : SETTINGS.wide;
  const { buffs, consumes } = result.request.character;
  const chips = [buffs.length, consumes.length].reduce(
    (sum, count) => sum + Math.min(count, at.chipsPerGroup),
    0,
  );
  return at.base + at.perChip * chips;
}

/** Cards whose height comes from data the main chunk does not have (the trees, the rotation file). */
const TALENTS: SectionReserve = { phone: 687, wide: 582 };
const ROTATION: SectionReserve = { phone: 2005, wide: 1353 };

function both(
  height: (result: SimResult, phone: boolean) => number | null,
  result: SimResult,
): SectionReserve | null {
  const phone = height(result, true);
  const wide = height(result, false);
  return phone === null || wide === null ? null : { phone: Math.round(phone), wide: Math.round(wide) };
}

/** Every lazy card's reserve for this result; null where the card is one line (no fight, no kit). */
export function sectionReserves(result: SimResult): Record<LazySectionId, SectionReserve | null> {
  return {
    timeline: both(timelineHeight, result),
    gear: both(gearHeight, result),
    talents: TALENTS,
    resource: both(resourceHeight, result),
    buffs: both(buffsHeight, result),
    rotation: ROTATION,
    settings: both(settingsHeight, result),
  };
}
