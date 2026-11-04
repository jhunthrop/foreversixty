// web/src/lib/bis/apl-gates.ts
// What an APL step's condition gates it on, read from the condition tree: a talent or spell the
// character must (or must not) know, and an aura that must (or must not) be up, on you or on
// the target. The guide's rotation cards word these ("1.5 s before the pull, with Unsealed
// Wrath, once Judgement of the Crusader is on the target") so no card shows a step as
// unconditional when the APL gates it (2026-11-04: Retribution's opener printed three Seals at
// -1.5 s after sim round 15 combined its gates under `and`).
//
// Every condition node the curated APLs use is listed in CONDITION_NODES, so a new shape fails
// apl-gates.test.ts instead of rendering ungated. A gate counts when every way through the
// condition passes it: `and` adds its children's gates, `not` flips a gate, `or` keeps only the
// gates all its branches share. Comparisons and every other value node gate on nothing a card
// names: they are what the step's own note explains.

type Node = Record<string, unknown>;

function isNode(value: unknown): value is Node {
  return typeof value === 'object' && value !== null && !Array.isArray(value);
}

/** A gate on knowing a talent or spell. */
export interface KnownGate {
  kind: 'known';
  spellId: number;
  /** False for `not`: the step runs without it. */
  positive: boolean;
}

/** A gate on an aura being up. */
export interface ActiveGate {
  kind: 'active';
  spellId: number;
  positive: boolean;
  /** True for an aura on the current target, false for one on you. */
  onTarget: boolean;
}

export type AplGate = KnownGate | ActiveGate;

/** Every node a curated condition is built from today. A gate leaf, a combinator, or a value
 *  that gates on nothing a card names. */
export const CONDITION_NODES: ReadonlySet<string> = new Set([
  // gate leaves
  'auraIsKnown',
  'spellIsKnown',
  'auraIsActive',
  // combinators
  'and',
  'or',
  'not',
  // values and comparisons: no card-named gate
  'auraNumStacks',
  'auraRemainingTime',
  'autoTimeToNext',
  'cmp',
  'const',
  'currentAttackPower',
  'currentComboPoints',
  'currentEnergy',
  'currentHealthPercent',
  'currentLevel',
  'currentMana',
  'currentManaPercent',
  'currentRage',
  'currentSealRemainingTime',
  'currentTime',
  'dotIsActive',
  'dotRemainingTime',
  'isExecutePhase',
  'math',
  'numberTargets',
  'remainingTime',
  'remainingTimePercent',
  'spellCanCast',
  'spellCastTime',
  'spellCurrentCost',
  'spellIsReady',
  'spellTimeToReady',
  'totemRemainingTime',
]);

function spellIdOf(node: unknown): number | undefined {
  return isNode(node) && typeof node.spellId === 'number' ? node.spellId : undefined;
}

function leafGate(node: Node): AplGate | undefined {
  if (isNode(node.auraIsKnown)) {
    const spellId = spellIdOf(node.auraIsKnown.auraId);
    return spellId === undefined ? undefined : { kind: 'known', spellId, positive: true };
  }
  if (isNode(node.spellIsKnown)) {
    const spellId = spellIdOf(node.spellIsKnown.spellId);
    return spellId === undefined ? undefined : { kind: 'known', spellId, positive: true };
  }
  if (isNode(node.auraIsActive)) {
    const spellId = spellIdOf(node.auraIsActive.auraId);
    const unit = node.auraIsActive.sourceUnit;
    const onTarget = isNode(unit) && unit.type === 'CurrentTarget';
    return spellId === undefined ? undefined : { kind: 'active', spellId, positive: true, onTarget };
  }
  return undefined;
}

function sameGate(a: AplGate, b: AplGate): boolean {
  return (
    a.kind === b.kind &&
    a.spellId === b.spellId &&
    a.positive === b.positive &&
    (a.kind !== 'active' || b.kind !== 'active' || a.onTarget === b.onTarget)
  );
}

function unique(gates: readonly AplGate[]): AplGate[] {
  return gates.filter((gate, index) => gates.findIndex((other) => sameGate(gate, other)) === index);
}

/** The gates every way through `condition` passes; empty for none or no condition. */
export function aplGates(condition: unknown): AplGate[] {
  if (!isNode(condition)) return [];
  const leaf = leafGate(condition);
  if (leaf !== undefined) return [leaf];
  if (isNode(condition.not)) {
    const inner = condition.not.val;
    const gate = isNode(inner) ? leafGate(inner) : undefined;
    return gate === undefined ? [] : [{ ...gate, positive: !gate.positive }];
  }
  if (isNode(condition.and) && Array.isArray(condition.and.vals)) {
    return unique(condition.and.vals.flatMap(aplGates));
  }
  if (isNode(condition.or) && Array.isArray(condition.or.vals)) {
    const [first, ...rest] = condition.or.vals.map(aplGates);
    return (first ?? []).filter((gate) =>
      rest.every((branch) => branch.some((other) => sameGate(gate, other))),
    );
  }
  return [];
}

/** The gates two steps share, for a card row that stands for both. */
export function sharedGates(a: readonly AplGate[], b: readonly AplGate[]): AplGate[] {
  return a.filter((gate) => b.some((other) => sameGate(gate, other)));
}

/** Every node name `condition` uses, for the test that holds CONDITION_NODES to the APLs. */
export function conditionNodeNames(condition: unknown, into: Set<string> = new Set()): Set<string> {
  if (Array.isArray(condition)) {
    for (const item of condition) conditionNodeNames(item, into);
  } else if (isNode(condition)) {
    for (const [key, value] of Object.entries(condition)) {
      if (key === 'vals' || key === 'val' || key === 'lhs' || key === 'rhs') {
        conditionNodeNames(value, into);
      } else if (isNode(value) && !['auraId', 'spellId', 'sourceUnit', 'targetUnit'].includes(key)) {
        into.add(key);
        conditionNodeNames(value, into);
      }
    }
  }
  return into;
}

/** "with Unsealed Wrath", "without Unsealed Wrath", "once Judgement of the Crusader is on the
 *  target", "while Judgement of the Crusader is not on the target", "while X is up". */
export function gateWords(gate: AplGate, name: string): string {
  if (gate.kind === 'known') return `${gate.positive ? 'with' : 'without'} ${name}`;
  if (gate.onTarget)
    return gate.positive ? `once ${name} is on the target` : `while ${name} is not on the target`;
  return gate.positive ? `while ${name} is up` : `while ${name} is not up`;
}
