// web/src/lib/home/outcome-rows.ts
// The landing's three outcome rows (P.1, replacing the five feature cards): what a player does
// here, in the player's words, one short line each. No product names as headings, no counts.
// When the visitor has a current character (lib/current-character.ts) the planner and simulator
// rows open that character; without one they open the tool.
import { plannerHrefFor, simHrefFor, type CurrentCharacter } from '../current-character';

export type OutcomeRowKey = 'planner' | 'sim' | 'logs';

export interface OutcomeRow {
  key: OutcomeRowKey;
  /** The row's title, in the player's words. */
  title: string;
  /** One short line under it. */
  line: string;
  /** Where it goes with no current character. */
  href: string;
}

export const outcomeRowsHeading = 'Get started';

export const outcomeRows: readonly OutcomeRow[] = [
  {
    key: 'planner',
    title: 'Plan your talents',
    line: 'Pick a build and share it by link.',
    href: '/planner',
  },
  {
    key: 'sim',
    title: 'Sim your gear',
    line: 'See what each upgrade is worth to your damage.',
    href: '/sim',
  },
  { key: 'logs', title: 'Upload tonight’s raid', line: 'Read every pull, and who did what.', href: '/logs' },
] as const;

/** The rows' links for a visitor: the planner and simulator rows open the current character when
 *  there is one; the logs row is the same page for everyone. */
export function outcomeHrefs(current: CurrentCharacter | null): Record<OutcomeRowKey, string> {
  return {
    planner: current === null ? '/planner' : plannerHrefFor(current),
    sim: current === null ? '/sim' : simHrefFor(current),
    logs: '/logs',
  };
}
