// web/src/lib/sim/head-meta.ts
// The simulator head's meta line (spec 2026-11-06-sim §4.3, §4.5): the character's race, spec,
// level, talent split and average item level, and where the gear came from and when; the saved
// head's fight facts. Pure, so each line is a unit test.
import { simPageCopy as copy } from './page-copy';
import { PLACEHOLDER_CHARACTER_NAMES } from './character';
import { talentSplit } from './result-band';
import { specLabel } from './spec-label';
import type { CharacterSource } from './types';

/** "high-order-skyborne" -> "High Order Skyborne". */
export function raceWords(slug: string): string {
  return slug
    .split('-')
    .filter((word) => word !== '')
    .map((word) => word.charAt(0).toUpperCase() + word.slice(1))
    .join(' ');
}

/** The worn items' average item level to one decimal, or null when any worn item is unknown. */
export function averageWornItemLevel(
  itemIds: readonly number[],
  itemLevel: (id: number) => number | undefined,
): string | null {
  if (itemIds.length === 0) return null;
  const levels = itemIds.map(itemLevel);
  if (levels.some((level) => level === undefined)) return null;
  return ((levels as number[]).reduce((sum, level) => sum + level, 0) / levels.length).toFixed(1);
}

const DAY_MONTH = new Intl.DateTimeFormat('en-GB', { day: 'numeric', month: 'short', timeZone: 'UTC' });

/** "2026-11-04T…" -> "4 Nov"; null for a missing or unreadable date. */
export function dayMonth(iso: string): string | null {
  const date = new Date(iso);
  return iso === '' || Number.isNaN(date.getTime()) ? null : DAY_MONTH.format(date).replace('Sept', 'Sep');
}

export interface HeadCharacter {
  race: string;
  spec: string;
  level: number;
  talents: string;
  source: CharacterSource;
}

/** Line 1 `{Race} {Spec} {Class} · Level {n} · {a/b/c} talents · item level {x}`, line 2 the source. */
export function characterMeta(character: HeadCharacter, itemLevel: string | null): [string, string] {
  const race = raceWords(character.race);
  const split = talentSplit(character.talents);
  const spec = specLabel(character.spec);
  const line1 =
    itemLevel === null
      ? copy.head.metaNoGear(race, spec, character.level, split)
      : copy.head.meta(race, spec, character.level, split, itemLevel);
  const from = copy.head.sourceKind[character.source.kind] ?? copy.head.sourceKind.manual;
  return [line1.trim(), copy.head.source(from, dayMonth(character.source.captured_at))];
}

/** The title (ruling S1): the character's name, or the spec when the source carried none (a
 *  published set's code, a placeholder name), the saved page's own rule (`bandHeading`). */
export function characterTitle(name: string, spec: string): string {
  return name !== '' && !PLACEHOLDER_CHARACTER_NAMES.has(name) ? name : specLabel(spec);
}
