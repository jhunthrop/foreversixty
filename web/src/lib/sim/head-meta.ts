// web/src/lib/sim/head-meta.ts
// The simulator head's meta line (spec 2026-11-06-sim §4.3, §4.5): the character's race, spec,
// level, talent split and average item level, and where the gear came from and when; the saved
// head's fight facts. Pure, so each line is a unit test.
import { simPageCopy as copy } from './page-copy';
import { PLACEHOLDER_CHARACTER_NAMES } from './character';
import { talentSplit } from './result-band';
import { specLabel } from './spec-label';
import type { CharacterSource } from './types';
import { utcDayMonth } from '../dates';

/** "high-order-skyborne" -> "High Order Skyborne". */
export function raceWords(slug: string): string {
  return slug
    .split('-')
    .filter((word) => word !== '')
    .map((word) => word.charAt(0).toUpperCase() + word.slice(1))
    .join(' ');
}

/** "2026-11-04T…" -> "4 Nov" in UTC; null for a missing or unreadable date. */
export function runDay(iso: string): string | null {
  const day = iso === '' ? '' : utcDayMonth(iso);
  return day === '' ? null : day;
}

export interface HeadCharacter {
  race: string;
  spec: string;
  level: number;
  talents: string;
  source: CharacterSource;
}

/** Line 1 `{Race} {Spec} {Class} · Level {n} · {a/b/c} talents · item level {x}`, line 2 the source.
 *  The default says nothing: an addon export with no capture date has no source line, since the
 *  line would only restate where a sim's gear comes from; a date, or any other source, keeps it. */
export function characterMeta(character: HeadCharacter, itemLevel: string | null): readonly string[] {
  const race = raceWords(character.race);
  const split = talentSplit(character.talents);
  const spec = specLabel(character.spec);
  const line1 =
    itemLevel === null
      ? copy.head.metaNoGear(race, spec, character.level, split)
      : copy.head.meta(race, spec, character.level, split, itemLevel);
  const date = runDay(character.source.captured_at);
  if (character.source.kind === 'addon' && date === null) return [line1.trim()];
  const from = copy.head.sourceKind[character.source.kind] ?? copy.head.sourceKind.manual;
  return [line1.trim(), copy.head.source(from, date)];
}

/** The title (ruling S1): the character's name, or the spec when the source carried none (a
 *  published set's code, a placeholder name), the saved page's own rule (`bandHeading`). */
export function characterTitle(name: string, spec: string): string {
  return name !== '' && !PLACEHOLDER_CHARACTER_NAMES.has(name) ? name : specLabel(spec);
}
