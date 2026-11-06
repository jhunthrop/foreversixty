// web/src/lib/guild/cover-view.ts
// The buff and debuff cover grid's view (owner decision D15): the API's rows (raid order: party
// buffs, raid buffs, debuffs) split into the two sections, each cell's Grid tone from the need
// against the players signed In or Late, and the spell's icon through the same build icon path
// SpellIcon draws everywhere. The counts are the API's, never recomputed here.
import { isRealIcon } from '../icons/placeholder';
import { dataUrl } from '../planner/load';
import type { MarkTone } from '../ui/cells';
import type { GuildRaidCoverRow } from './api';

/** A Grid cell's tone: covered, covered on a talent we cannot see or too few, nobody. */
export type CoverTone = MarkTone;

export interface CoverSection {
  kind: GuildRaidCoverRow['kind'];
  rows: GuildRaidCoverRow[];
}

/** The sections in the order a raid leader reads them, each in the API's order; an empty
 *  section is left out. */
export function coverSections(rows: readonly GuildRaidCoverRow[] | null | undefined): CoverSection[] {
  const kinds: GuildRaidCoverRow['kind'][] = ['buff', 'debuff'];
  return kinds
    .map((kind) => ({ kind, rows: (rows ?? []).filter((row) => row.kind === kind) }))
    .filter((section) => section.rows.length > 0);
}

/** The players signed In or Late who can provide the row. */
export function coverSigned(row: GuildRaidCoverRow): number {
  return row.providers_in + row.providers_late;
}

/** ok when the In and Late providers meet the need without counting a talent their sign-ups
 *  cannot show; warn when the need is met only that way or some are signed but too few; miss
 *  when none are (red at 0). */
export function coverTone(row: GuildRaidCoverRow): CoverTone {
  const signed = coverSigned(row);
  if (signed === 0) return 'miss';
  return signed - row.likely >= row.need ? 'ok' : 'warn';
}

/** The cell's count, `2 of 3`. */
export function coverCount(row: GuildRaidCoverRow): string {
  return `${coverSigned(row)} of ${row.need}`;
}

/** The row's spell icon in `build`, or null when it has none to draw. */
export function coverIconUrl(build: string, row: GuildRaidCoverRow): string | null {
  return isRealIcon(row.icon) ? dataUrl(build, `icons/${row.icon}.webp`) : null;
}
