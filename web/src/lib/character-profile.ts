// web/src/lib/character-profile.ts
// Small reads off the character page's GET /v1/characters/... answer (rankings/api.ts's
// CharacterPage) that the profile header and the Forever Score card share.
import type { CharacterFight, CharacterPage } from './rankings/api';

/** The spec of the character's newest fight: `history[0]`, else the first best row that has
 *  one. Null when no fight names a spec (the header then shows the class alone). */
export function newestSpec(page: Pick<CharacterPage, 'history' | 'best'>): string | null {
  const spec = page.history[0]?.spec ?? page.best.find((row) => row.spec !== undefined)?.spec;
  return spec === undefined || spec.trim() === '' ? null : spec;
}

/** "Holy Priest" when the spec is known, else "Priest"; empty with no class. */
export function specClassLabel(className: string | undefined, spec: string | null): string {
  if (className === undefined || className === '') return '';
  return spec === null ? className : `${spec} ${className}`;
}

function slug(text: string): string {
  return text.trim().toLowerCase().replace(/\s+/g, '-');
}

/** The best-in-slot page for a class and spec (`/bis/priest/holy`), or null when either is
 *  unknown. */
export function bisHref(className: string | undefined, spec: string | null): string | null {
  if (className === undefined || className === '' || spec === null) return null;
  return `/bis/${slug(className)}/${slug(spec)}`;
}

/** The two fight tables' column headers (fix round 2 item 6). */
export const characterTableCopy = {
  boss: 'Boss',
  spec: 'Spec',
  parse: 'Parse %',
  date: 'Date',
  execution: 'Execution',
} as const;

/** The amount column's header: "HPS" or "DPS" when every row is one metric, else both. */
export function amountHeader(rows: readonly Pick<CharacterFight, 'metric'>[]): string {
  const metrics = new Set(rows.map((row) => row.metric));
  if (metrics.size === 1 && metrics.has('hps')) return 'HPS';
  if (metrics.size === 1 && metrics.has('dps')) return 'DPS';
  return 'DPS or HPS';
}

/** Whether "Best per encounter" would only repeat "Every ranked fight": the same fights, by
 *  report and fight index, in both. */
export function bestRepeatsHistory(
  best: readonly Pick<CharacterFight, 'report_id' | 'fight_index'>[],
  history: readonly Pick<CharacterFight, 'report_id' | 'fight_index'>[],
): boolean {
  if (best.length !== history.length) return false;
  const key = (row: Pick<CharacterFight, 'report_id' | 'fight_index'>): string =>
    `${row.report_id}:${row.fight_index}`;
  const seen = new Set(history.map(key));
  return best.every((row) => seen.has(key(row)));
}
