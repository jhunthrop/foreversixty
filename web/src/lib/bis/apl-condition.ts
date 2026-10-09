// web/src/lib/bis/apl-condition.ts
// A curated APL step's condition and target, in plain words (tenet 3: an ability is never
// just a name). The healer lists carry one step per target and per threshold -- "Heal the
// tank below 80%", "Heal a party member below 65%" -- and the step's own note is on only
// one of them, so the words have to come from the condition tree itself.
//
// Only the shapes the lists use today are read (health and mana comparisons, the mana
// pace, "not already on the target"); a step with any other clause reads as `null` and
// keeps its note alone rather than a half-translated sentence.

/** The fake raid the healer lists are written against: player 5 is the tank, 1-4 the party. */
const TANK_PLAYER_INDEX = 5;
const PERCENT = 100;

const TANK = 'the tank';
const PARTY_MEMBER = 'a party member';

const COMPARISON_WORDS: Record<string, string> = {
  OpLt: 'below',
  OpLe: 'at or below',
  OpGt: 'above',
  OpGe: 'at or above',
};

type Node = Record<string, unknown>;

/** What one step says in words: who it is cast on, and what has to hold. */
export interface StepWords {
  /** "the tank below 35%", "a party member", or null for a self-cast with no threshold. */
  subject: string | null;
  /** Clauses that do not name the target: "your mana is below 50%". */
  qualifiers: string[];
}

export interface AplStepInput {
  target?: unknown;
  condition?: unknown;
}

function isNode(value: unknown): value is Node {
  return typeof value === 'object' && value !== null && !Array.isArray(value);
}

function playerWho(unit: unknown): string | null {
  if (!isNode(unit) || unit.type !== 'Player' || typeof unit.index !== 'number') return null;
  return unit.index === TANK_PLAYER_INDEX ? TANK : PARTY_MEMBER;
}

function constantOf(node: unknown): number | null {
  if (!isNode(node) || !isNode(node.const)) return null;
  const value = Number(node.const.val);
  return Number.isFinite(value) ? value : null;
}

function percentText(fraction: number): string {
  return `${Math.round(fraction * PERCENT)}%`;
}

interface Atom {
  /** Who the clause is about: a target phrase, "you", or null for a clause about no one. */
  who: string | null;
  /** The clause without its subject: "below 35%". */
  predicate: string;
}

function healthAtom(lhs: Node, word: string, bound: number): Atom {
  const source = isNode(lhs.currentHealthPercent) ? lhs.currentHealthPercent.sourceUnit : undefined;
  if (source === undefined) return { who: null, predicate: `your health is ${word} ${percentText(bound)}` };
  return { who: playerWho(source), predicate: `${word} ${percentText(bound)}` };
}

function manaPaceAtom(node: Node): Atom | null {
  const rhs = node.rhs;
  if (!isNode(rhs) || !isNode(rhs.math) || rhs.math.op !== 'OpMul') return null;
  const factor = constantOf(rhs.math.rhs);
  if (factor === null || !isNode(rhs.math.lhs) || !isNode(rhs.math.lhs.remainingTimePercent)) return null;
  return {
    who: null,
    predicate: `your mana is at least ${percentText(factor)} of the share of the fight left`,
  };
}

function comparisonAtom(node: Node): Atom | null {
  const word = COMPARISON_WORDS[String(node.op)];
  if (word === undefined || !isNode(node.lhs)) return null;
  const bound = constantOf(node.rhs);
  if (isNode(node.lhs.currentHealthPercent) && bound !== null) return healthAtom(node.lhs, word, bound);
  if (isNode(node.lhs.currentManaPercent)) {
    if (bound !== null) return { who: null, predicate: `your mana is ${word} ${percentText(bound)}` };
    return node.op === 'OpGe' ? manaPaceAtom(node) : null;
  }
  return null;
}

function atomsOf(condition: unknown): Atom[] | null {
  if (!isNode(condition)) return null;
  if (isNode(condition.cmp)) {
    const atom = comparisonAtom(condition.cmp);
    return atom === null ? null : [atom];
  }
  if (isNode(condition.not) && isNode(condition.not.val) && isNode(condition.not.val.dotIsActive)) {
    return [{ who: null, predicate: 'it is not already running on the target' }];
  }
  if (isNode(condition.and) && Array.isArray(condition.and.vals)) {
    const parts = condition.and.vals.map(atomsOf);
    return parts.some((part) => part === null) ? null : parts.flat().filter((atom) => atom !== null);
  }
  return null;
}

/** The step in words, or `null` when its condition has a clause this module does not read. */
export function stepWords(step: AplStepInput): StepWords | null {
  const target = playerWho(step.target);
  const atoms = step.condition === undefined ? [] : atomsOf(step.condition);
  if (atoms === null) return null;
  const subjectParts: string[] = [];
  const qualifiers: string[] = [];
  for (const atom of atoms) {
    if (atom.who !== null && atom.who === target) subjectParts.push(atom.predicate);
    else if (atom.who === null) qualifiers.push(atom.predicate);
    else qualifiers.push(`${atom.who} ${atom.predicate}`);
  }
  if (target === null) return { subject: null, qualifiers };
  return { subject: [target, ...subjectParts].join(' '), qualifiers };
}

/** Groups steps that share qualifiers and lists each group's distinct subjects in order. */
export function describeSteps(steps: readonly StepWords[]): string {
  const groups = new Map<string, { qualifiers: string[]; subjects: string[] }>();
  for (const step of steps) {
    const key = step.qualifiers.join(' and ');
    const group = groups.get(key) ?? { qualifiers: step.qualifiers, subjects: [] };
    if (step.subject !== null && !group.subjects.includes(step.subject)) group.subjects.push(step.subject);
    groups.set(key, group);
  }
  const clauses = [...groups.values()]
    .map(({ qualifiers, subjects }) => {
      const onWhom = subjects.length > 0 ? `on ${subjects.join(', then ')}` : '';
      const when = qualifiers.length > 0 ? `when ${qualifiers.join(' and ')}` : '';
      return [onWhom, when].filter((part) => part !== '').join(', ');
    })
    .filter((clause) => clause !== '');
  return clauses.length === 0 ? '' : `Cast ${clauses.join('; or ')}.`;
}
