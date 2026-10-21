// web/src/lib/report/columns.ts
// Which columns of a report's tables say anything. A column that knows nothing is hidden, not
// filled with dashes or question marks: a visitor reads a column of "–" as a broken page.
import type { FightEntry } from './types';

/** The Parse column exists only on a report that has a boss kill to rank; trash and wipes are never ranked. */
export function hasRankedKill(fights: readonly Pick<FightEntry, 'kind' | 'kill'>[]): boolean {
  return fights.some((fight) => fight.kind === 'encounter' && fight.kill);
}

/** The Spec column exists only when at least one row has a spec; a class alone is the name's colour. */
export function hasAnySpec(rows: readonly { spec?: string }[]): boolean {
  return rows.some((row) => row.spec !== undefined && row.spec !== '');
}
