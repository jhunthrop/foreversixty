// web/src/lib/guild/fact-line.ts
// A line of facts on a guild page (P.1, player review S2-1, UX P2-1): the header's "PvP US · 1 boss
// down · 2 pulls", the Overview's "5 raiders with gear on the site · 0 waiting for approval", a
// night's "1 boss (Onyxia) · 1 kill · 2 trash packs · 1 wipe" and a roster row's stats. Body face,
// " · " with a space each side, and each fact one unbreakable piece so a narrow screen wraps whole
// facts ("1 wipe" never splits, never stands alone mid-phrase). Pure, so the words are pinned
// without a page; FactLine.svelte draws them.
import { counted } from '../plural';

export const FACT_SEPARATOR = ' · ';

/** One fact: its words, and optionally the hover, a class for the one that needs colour, and a testid. */
export interface Fact {
  text: string;
  title?: string;
  class?: string;
  testid?: string;
}

/** The dot that ends every fact but the last: a no-break space then the dot, so a wrap never starts a
 *  line with one. */
export const FACT_TAIL = '\u00a0·';

/** The line as plain words: "PvP US · 1 boss down · 2 pulls". */
export function factsText(facts: readonly (Fact | string)[]): string {
  return facts.map((fact) => (typeof fact === 'string' ? fact : fact.text)).join(FACT_SEPARATOR);
}

/** The facts as objects, whatever the caller held. */
export function asFacts(facts: readonly (Fact | string)[]): Fact[] {
  return facts.map((fact) => (typeof fact === 'string' ? { text: fact } : fact));
}

/** The header's line: the ruleset and region, the tier's boss and pull counts as two facts, and when the
 *  guild was last updated ("Updated today", when the page knows), all one sentence. */
export function bandFacts(
  ruleset: string,
  region: string,
  counts: { killed: number; pulls: number } | null,
  updated = '',
): string[] {
  const facts = [`${ruleset} ${region.toUpperCase()}`];
  if (counts !== null) {
    facts.push(`${counted(counts.killed, 'boss', 'bosses')} down`, counted(counts.pulls, 'pull'));
  }
  if (updated !== '') facts.push(updated);
  return facts;
}
