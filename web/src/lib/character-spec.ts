// web/src/lib/character-spec.ts
// The one rule for "which spec is this character": the spec its talent points lean to when the
// record carries talent ranks, else the spec the addon reported. Every surface that prints a
// character's spec (the header chip and card, the character page, the account, the planner card,
// the simulator card) asks here, so a record whose reported spec and ranks disagree is read the
// same everywhere, and the name is the display name ("Beast Mastery Hunter"), never the export's
// run-together "BeastMastery".
import { specName } from './guild/spec-name';
import { SPECS } from './sim/specs';

/** The tree names of a class in tab order ("Beast Mastery", "Marksmanship", "Survival"). */
function treeNames(classSlug: string): string[] {
  return SPECS.filter((row) => row.class_slug === classSlug)
    .sort((a, b) => a.tree_index - b.tree_index)
    .map((row) => row.name);
}

/**
 * How far the leading tree must lead every other before the points name it, and the fewest spent
 * points that can name a tree at all: the addon's Talents.SURE_LEAD and SURE_MIN_POINTS (owner
 * 2026-11-11, "a wrong silent guess is worse than one question"). api/internal/character holds
 * the Go copy; all three read the same cases in their tests.
 */
export const SURE_LEAD = 5;
export const SURE_MIN_POINTS = 2;

/**
 * The tab index the points name beyond doubt, as the addon's Talents.decideTree reads them: at
 * least SURE_MIN_POINTS spent and all in one tree, or the leading tree holding more than half of
 * the points and leading every other by SURE_LEAD. Null otherwise: a tie, a close split, nothing
 * spent, where the addon asks the player.
 */
export function decideTree(points: readonly number[]): number | null {
  const spent = points.reduce((sum, count) => sum + Math.max(count, 0), 0);
  if (spent < SURE_MIN_POINTS) return null;
  const ranked = points
    .map((count, index) => ({ count, index }))
    .sort((a, b) => b.count - a.count || a.index - b.index);
  const [lead, second] = ranked;
  if (lead === undefined) return null;
  const runnerUp = Math.max(second?.count ?? 0, 0);
  if (runnerUp === 0) return lead.index;
  return lead.count * 2 > spent && lead.count - runnerUp >= SURE_LEAD ? lead.index : null;
}

/**
 * The spec named by a per-tree point split (decideTree). Null when the points do not name a tree,
 * or the class has no tree at that position; callers fall back to the spec the addon reported.
 */
export function specFromPoints(classSlug: string, points: readonly number[]): string | null {
  const tree = decideTree(points);
  if (tree === null) return null;
  return treeNames(classSlug.toLowerCase())[tree] ?? null;
}

/** The pieces of a character record the rule reads. */
export interface SpecSource {
  class?: string;
  spec?: string;
  build?: { talents?: { points: readonly number[] } };
}

/**
 * The character's spec as a player says it: from its talent ranks when it has them, else from the
 * spec the addon reported. Undefined when neither names one.
 */
export function characterSpecName(character: SpecSource): string | undefined {
  const points = character.build?.talents?.points;
  if (points !== undefined && character.class !== undefined) {
    const derived = specFromPoints(character.class, points);
    if (derived !== null) return derived;
  }
  const reported = character.spec === undefined ? '' : specName(character.spec);
  return reported === '' ? undefined : reported;
}

/** "Beast Mastery Hunter": the spec then the class, or the class alone when no spec is known. */
export function characterSpecClass(character: SpecSource): string {
  const spec = characterSpecName(character);
  const className = character.class === undefined ? '' : specName(character.class);
  return [spec, className].filter((part) => part !== undefined && part !== '').join(' ');
}

/** The record with `spec` replaced by the rule's answer, for the readers that take a `spec` field. */
export function withRuleSpec<T extends SpecSource>(character: T): T {
  const spec = characterSpecName(character);
  return spec === character.spec ? character : { ...character, spec };
}

/** The canonical spec key ("hunter-beast-mastery") of a record by the one rule; undefined when unknown. */
export function characterSpecKey(character: SpecSource): string | undefined {
  const name = characterSpecName(character);
  if (name === undefined || character.class === undefined) return undefined;
  const classSlug = character.class.toLowerCase();
  return SPECS.find((row) => row.class_slug === classSlug && row.name === name)?.spec;
}

/** False for a healing or tanking spec: the simulator does not model them. Unknown specs stay runnable. */
export function characterIsSimulated(character: SpecSource): boolean {
  const key = characterSpecKey(character);
  return key === undefined || SPECS.find((row) => row.spec === key)?.role === 'dps';
}
