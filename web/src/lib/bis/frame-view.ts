// web/src/lib/bis/frame-view.ts
// The paperdoll-led BiS page's own views (design/specs/2026-11-09-directions.md, Paperdoll and
// Best in slot): beside the frame, What to stack (the top weights and what 1.00 is worth) and
// Where to go (the stops and how many pieces each gives); below it, the itinerary (each stop's
// pieces, each row saying only what its heading does not). Everything reads the band's model and
// the doll's cells, so the frame, the side and the itinerary never disagree.
import type { DollCell } from '../ui/paperdoll';
import type { GainFigure } from '../ui/primitives';
import { gearCopy } from '../gear/copy';
import type { BisBandModel } from './page-model';
import type { ScaleRow } from './scale-view';
import { whereGroups } from './doll-view';

/** What to stack lists this many stats; the sim's weights page has them all. */
export const STACK_ROWS = 4;

/** The source line's separator (doll-view's lines read "Skum · Wailing Caverns (15-24) · 50%"). */
const SEP = ' · ';
/** An instance with its level range, as a source line names it: "Wailing Caverns (15-24)". */
const PLACE_RANGE = /^(.*) \((\d+)-(\d+)\)$/;

export interface StackView {
  rows: readonly ScaleRow[];
  /** "1.00 = 0.316 DPS per point of Agility". */
  unitLine?: string;
}

/** The band's significant weights, loudest first, and what 1.00 is worth. */
export function stackView(model: BisBandModel): StackView {
  const rows = model.weights.rows.filter((row) => row.significant).slice(0, STACK_ROWS);
  const lead = rows[0];
  if (lead?.dpsPerPoint === undefined) return { rows };
  return {
    rows,
    unitLine: `${lead.scaleFactor.toFixed(2)} = ${lead.dpsPerPoint.toFixed(3)} ${lead.perPointUnit} per point of ${lead.label}`,
  };
}

export interface ItineraryRow {
  cell: DollCell;
  /** The slot, then what the source adds to its stop's heading: "Back · Skum · 50%". */
  line: string;
}

export interface ItineraryStop {
  key: string;
  /** "Wailing Caverns", "Crafted · Leatherworking", "World / BoE". */
  title: string;
  /** An instance's level range: "15 to 24". */
  range?: string;
  /** The id Where to go links to. */
  anchor: string;
  rows: ItineraryRow[];
}

/** The stops in route order with their pieces (doll-view's Where to go groups). */
export function itinerary(model: BisBandModel, cells: readonly DollCell[]): ItineraryStop[] {
  const bySlot = new Map(cells.map((cell) => [cell.slot, cell]));
  return whereGroups(model, cells).map((group) => {
    const rows = group.slots.flatMap((slot) => {
      const cell = bySlot.get(slot);
      return cell === undefined ? [] : [cell];
    });
    const range = rows
      .map((cell) => placeRange(cell.source, group.title))
      .find((found) => found !== undefined);
    return {
      key: group.key,
      title: group.title,
      ...(range === undefined ? {} : { range }),
      anchor: `stop-${group.key.toLowerCase().replace(/[^a-z0-9]+/g, '-')}`,
      rows: rows.map((cell) => ({ cell, line: rowLine(cell, group.title) })),
    };
  });
}

/** The level range of the instance `place` a source line names: "15 to 24". A world drop's own
 *  range ("World drop (20-26)") is the item's, not the stop's, and stays on its row. */
export function placeRange(source: string | undefined, place: string): string | undefined {
  for (const part of (source ?? '').split(SEP)) {
    const match = PLACE_RANGE.exec(part);
    if (match && match[1] === place) return `${match[2]} to ${match[3]}`;
  }
  return undefined;
}

/** A row's line under its stop: the slot, then the source's parts the heading does not say. */
export function rowLine(cell: DollCell, title: string): string {
  const said = new Set(title.split(SEP));
  const parts = (cell.source ?? '').split(SEP).filter((part) => {
    if (part === '' || said.has(part)) return false;
    const match = PLACE_RANGE.exec(part);
    return match === null || !said.has(match[1] ?? '');
  });
  return [cell.label, ...parts].join(SEP);
}

/** The band's set on the sim's weights page: the same code Open in the sim carries. */
export function simWeightsHref(simHref: string | undefined): string | undefined {
  return simHref?.replace(/^\/sim\?/, '/sim/weights?');
}

/** A gain with a figure ("+0.3", "−1.2"), not a tie, an only option or a worn item's word. */
export function isMeasuredGain(value: string | undefined): boolean {
  return value !== undefined && /^[+−-]\d/.test(value);
}

/** A row's figure in the itinerary, so no row is a silent blank: a measured gain in the gain
 *  colour (a re-pick's loss in red); a tie with the next pick as "Same DPS"; a worn pick as Worn;
 *  an only option or an unscored one by its own words ("Only option", "Not sim-checked"). */
export function rowFigure(gain: GainFigure | undefined, unit: string): GainFigure | undefined {
  if (gain === undefined) return undefined;
  if (isMeasuredGain(gain.value)) {
    return { value: gain.value, unit: gain.unit, tone: gain.tone === 'wipe' ? 'wipe' : 'gain' };
  }
  if (gain.value === gearCopy.paperdoll.same) return { value: gain.value, unit, tone: 'muted' };
  if (gain.value === gearCopy.paperdoll.worn) return { value: gain.value, tone: gain.tone ?? 'kill' };
  return { value: gain.basis?.prefix ?? gain.value, tone: 'muted' };
}
