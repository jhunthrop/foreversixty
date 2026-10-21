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
 * The spec named by a per-tree point split: the tree with the most points, ties to the first
 * tree. Null when nothing is spent, or the class has no tree for the largest position.
 */
export function specFromPoints(classSlug: string, points: readonly number[]): string | null {
  let best = -1;
  let at = -1;
  points.forEach((count, index) => {
    if (count > best) {
      best = count;
      at = index;
    }
  });
  if (best <= 0) return null;
  return treeNames(classSlug.toLowerCase())[at] ?? null;
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
