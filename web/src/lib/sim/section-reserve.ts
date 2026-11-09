// web/src/lib/sim/section-reserve.ts
// How tall each lazy section of a sim result will be, held only while its chunk loads (spec
// 2026-11-04-sim-results 4.2 and 8; sim spec 2026-11-06 §9): computed from this result's own data
// where the height follows it (the timeline's lanes, the rage sources, the aura rows, the kit, the
// buffs), from the section's measured frame elsewhere. LazySection drops the reserve the moment its
// section renders or fails, so an estimate that is off moves only what is still below the fold.
//
// The frame figures were measured on the built page (2026-11-06, re-measured 2026-11-09 on the
// round-5 panels, `simfullfixt1`, at 390 for phone and 1440 for wide); the timeline's rows come from the very
// lane rules the chart draws with (lib/sim/timeline-lanes.ts). Talents is an estimate: the fixture
// build's talent layout does not fit its runs, so it draws the one-line fallback there.
import { plotHeightOf, ROW_HEIGHT, timelineShape } from './timeline-lanes';
import type { SimResult } from './types';

/** A section's reserved height under 768 px and from 768 px. */
export interface SectionReserve {
  phone: number;
  wide: number;
}

export type LazySectionId = 'timeline' | 'gear' | 'talents' | 'resource' | 'buffs' | 'rotation' | 'settings';

type Widths<T> = Readonly<{ phone: T; wide: T }>;

/** One open StatTable row in the Rage and Buffs sections; a phone puts the bar under the name. */
const LIST_ROW: Widths<number> = { phone: 85, wide: 56 };

/** The Section head, the toolbar and legend over the plot, and the cast count line under it. */
const TIMELINE_FRAME: Widths<number> = { phone: 321, wide: 198 }; // re-measured 2026-11-09 (round 5 panels)

function timelineHeight(result: SimResult, phone: boolean): number | null {
  const shape = timelineShape(result);
  if (shape === null) return null;
  return (
    (phone ? TIMELINE_FRAME.phone : TIMELINE_FRAME.wide) + plotHeightOf(shape, phone) + ROW_HEIGHT.castFoot
  );
}

/** Rage: the head and three Figures, then the two source lists side by side (stacked on a phone), or one line. */
const RESOURCE: Widths<{ lists: number; needsNewer: number }> = {
  phone: { lists: 173, needsNewer: 137 },
  wide: { lists: 155, needsNewer: 143 },
};

function resourceHeight(result: SimResult, phone: boolean): number {
  const at = phone ? RESOURCE.phone : RESOURCE.wide;
  const sources = result.resource_sources?.[0];
  if (sources === undefined) return at.needsNewer;
  const gained = sources.gained.length;
  const spent = sources.spent.length;
  const rows = phone ? gained + spent : Math.max(gained, spent);
  return at.lists + (phone ? LIST_ROW.phone : LIST_ROW.wide) * rows;
}

/** Buffs: one uptime row per aura that came and went; the frame holds the heads and Up all fight. */
const BUFFS_FRAME: Widths<number> = { phone: 732, wide: 426 };
/** The Overpower window is the rotation's, not a buff: the Buffs section leaves it out (aura-card.ts). */
const OVERPOWER_WINDOWS: ReadonlySet<string> = new Set([
  'spell:7384',
  'spell:7887',
  'spell:11584',
  'spell:11585',
]);
/** The Buffs section files an aura at this share or more under "Up all fight", not as a row. */
const BUFF_UP_ALL_FIGHT = 0.999;

function buffsHeight(result: SimResult, phone: boolean): number {
  const duration = result.summary.duration_ms;
  const bars = (type: string): number =>
    result.summary.auras.filter(
      (track) =>
        track.type === type &&
        track.uptime_ms > 0 &&
        // A swing's window and a queue's own window are never rows.
        !track.name.startsWith('other:') &&
        !track.name.includes('/') &&
        !OVERPOWER_WINDOWS.has(track.name) &&
        track.uptime_ms < duration * BUFF_UP_ALL_FIGHT,
    ).length;
  const player = bars('BUFF');
  const target = bars('DEBUFF');
  return phone
    ? BUFFS_FRAME.phone + LIST_ROW.phone * (player + target)
    : BUFFS_FRAME.wide + LIST_ROW.wide * Math.max(player, target);
}

/** Gear: the open doll is the same frame for any kit; line 2 and the notes grow with it. */
const GEAR: Widths<{ base: number; perItem: number }> = {
  phone: { base: 1012.5, perItem: 35.7 },
  wide: { base: 557, perItem: 9.5 },
};

function gearHeight(result: SimResult, phone: boolean): number | null {
  const items = result.request.character.gear.length;
  if (items === 0) return null;
  const at = phone ? GEAR.phone : GEAR.wide;
  return at.base + at.perItem * items;
}

/** This run: the Facts, then a wrapping IconList entry per buff and consumable, then the actions. */
const SETTINGS: Widths<{ base: number; perEntry: number }> = {
  phone: { base: 741, perEntry: 22.5 },
  wide: { base: 704, perEntry: 25.2 },
};

function settingsHeight(result: SimResult, phone: boolean): number {
  const at = phone ? SETTINGS.phone : SETTINGS.wide;
  const { buffs, consumes } = result.request.character;
  return at.base + at.perEntry * (buffs.length + consumes.length);
}

/** Sections whose height comes from data the main chunk does not have (the trees, the rotation file). */
const TALENTS: SectionReserve = { phone: 1128, wide: 430 };
const ROTATION: SectionReserve = { phone: 1781, wide: 1913 };

function both(
  height: (result: SimResult, phone: boolean) => number | null,
  result: SimResult,
): SectionReserve | null {
  const phone = height(result, true);
  const wide = height(result, false);
  return phone === null || wide === null ? null : { phone: Math.round(phone), wide: Math.round(wide) };
}

/** Every lazy section's reserve for this result; null where the section is one line (no fight, no kit). */
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
