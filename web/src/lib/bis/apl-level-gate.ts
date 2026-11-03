// web/src/lib/bis/apl-level-gate.ts
// Whether a curated APL step can fire at a character level, read off its condition tree.
// The lists gate leveling steps on what the character has trained: "Below level 20, before
// Cat Form" is `not spellIsKnown(Claw)`, "from level 24" is `spellIsKnown(Ferocious Bite)`.
// The engine resolves the named id to the family's highest learned rank, so a spell is
// known at a level when any rank of its family is learned by then. A talent spell is known
// only with the talent, which the level alone does not decide: below its learn level it is
// not known, from there on the answer is open.
//
// The condition is read three-valued (true, false, open): every clause that is not a
// `spellIsKnown` is open, so a step is ruled out only when its level gate alone makes the
// whole condition false.

import { levelledRanks } from './apl-steps';

type Node = Record<string, unknown>;

/** true, false, or `undefined` when the level does not decide it. */
export type Truth = boolean | undefined;

/** Whether the spell `spellId` names is known at `level`. */
export type KnownAt = (spellId: number, level: number) => Truth;

function isNode(value: unknown): value is Node {
  return typeof value === 'object' && value !== null && !Array.isArray(value);
}

function spellIdOf(node: unknown): number | undefined {
  if (!isNode(node) || !isNode(node.spellId)) return undefined;
  return typeof node.spellId.spellId === 'number' ? node.spellId.spellId : undefined;
}

function allOf(values: readonly Truth[]): Truth {
  if (values.includes(false)) return false;
  return values.every((value) => value === true) ? true : undefined;
}

function anyOf(values: readonly Truth[]): Truth {
  if (values.includes(true)) return true;
  return values.every((value) => value === false) ? false : undefined;
}

/** The condition at `level`, as far as what the character has trained decides it. */
export function conditionAt(condition: unknown, level: number, knownAt: KnownAt): Truth {
  if (condition === undefined) return true;
  if (!isNode(condition)) return undefined;
  if ('spellIsKnown' in condition) {
    const spellId = spellIdOf(condition.spellIsKnown);
    return spellId === undefined ? undefined : knownAt(spellId, level);
  }
  if (isNode(condition.not)) {
    const inner = conditionAt(condition.not.val, level, knownAt);
    return inner === undefined ? undefined : !inner;
  }
  for (const [key, combine] of [
    ['and', allOf],
    ['or', anyOf],
  ] as const) {
    const group = condition[key];
    if (isNode(group) && Array.isArray(group.vals)) {
      return combine(group.vals.map((value) => conditionAt(value, level, knownAt)));
    }
  }
  return undefined;
}

/** Whether the condition can hold at some level from `fromLevel` through `toLevel`: false
 *  only when the level gate rules it out across the whole range. */
export function canApplyBetween(
  condition: unknown,
  fromLevel: number,
  toLevel: number,
  knownAt: KnownAt,
): boolean {
  for (let level = fromLevel; level <= Math.max(fromLevel, toLevel); level += 1) {
    if (conditionAt(condition, level, knownAt) !== false) return true;
  }
  return false;
}

/** A `KnownAt` over a class's rank families (`spellranks.json` `classes.<class>`), with
 *  `talentNames` the families only a talent teaches. An id in no family is open. */
export function knownAtFor(
  families: Readonly<Record<string, readonly { id: number; level: number }[]>>,
  talentNames: ReadonlySet<string>,
): KnownAt {
  const familyOf = new Map<number, { name: string; firstLevel: number }>();
  for (const [name, ranks] of Object.entries(families)) {
    const levelled = levelledRanks(ranks);
    const firstLevel = Math.min(...levelled.map((rank) => rank.level));
    for (const rank of levelled) familyOf.set(rank.id, { name, firstLevel });
  }
  return (spellId, level) => {
    const family = familyOf.get(spellId);
    if (family === undefined) return undefined;
    if (level < family.firstLevel) return false;
    return talentNames.has(family.name) ? undefined : true;
  };
}
