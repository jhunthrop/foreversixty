// web/src/lib/bis/apl-steps.ts
// The curated APL's cast steps for one spec at one level, in the order
// `data/pipeline/addonrotation.py` publishes them as `addon-data.json` rotation lines (the
// same flattening and the same highest-learned-rank rule), each with the target and
// condition words the published line drops. The published line stays the source of the
// spell, rank and icon; these steps only add the words, and are used only when they line up
// with it one to one.
import { existsSync, readFileSync } from 'node:fs';
import path from 'node:path';
import { fileURLToPath } from 'node:url';
import { stepWords, type StepWords } from './apl-condition';

const REPO_ROOT = path.resolve(path.dirname(fileURLToPath(import.meta.url)), '../../../..');

type Node = Record<string, unknown>;

interface RankRow {
  id: number;
  level: number;
}

type SpellRanksByName = Record<string, RankRow[]>;

interface CastStep {
  spellId: number;
  target?: unknown;
  condition?: unknown;
}

export interface AplStep {
  /** The id the published line casts at this level. */
  resolvedSpellId: number;
  words: StepWords | null;
}

function isNode(value: unknown): value is Node {
  return typeof value === 'object' && value !== null && !Array.isArray(value);
}

/** Depth first, like addonrotation.py's `_collect_cast_lines`: a direct castSpell is one
 *  step, a sequence unfolds into one per child and shares the sequence's condition. */
function collectCasts(action: unknown, inherited: unknown, out: CastStep[]): void {
  if (!isNode(action)) return;
  const condition = action.condition ?? inherited;
  if (isNode(action.castSpell)) {
    const spellId = isNode(action.castSpell.spellId) ? action.castSpell.spellId.spellId : undefined;
    if (typeof spellId === 'number') out.push({ spellId, target: action.castSpell.target, condition });
    return;
  }
  for (const key of ['strictSequence', 'sequence']) {
    const sequence = action[key];
    if (isNode(sequence) && Array.isArray(sequence.actions)) {
      for (const child of sequence.actions) collectCasts(child, condition, out);
    }
  }
}

function castSteps(apl: unknown): CastStep[] {
  const list = isNode(apl) && isNode(apl.rotation) ? apl.rotation.priorityList : undefined;
  const out: CastStep[] = [];
  if (!Array.isArray(list)) return out;
  for (const entry of list) collectCasts(isNode(entry) ? entry.action : undefined, undefined, out);
  return out;
}

/** A family's ranks by learn level; a level-0 alternate id is kept only when the family
 *  names no real level (addonrotation.py `_family_ranks`). */
function levelledRanks(ranks: readonly RankRow[]): RankRow[] {
  const sorted = [...ranks].sort((a, b) => a.level - b.level || a.id - b.id);
  const levelled = sorted.filter((rank) => rank.level > 0);
  return levelled.length > 0 ? levelled : sorted;
}

function resolveAt(families: SpellRanksByName, spellId: number, level: number): number | null {
  for (const ranks of Object.values(families)) {
    const levelled = levelledRanks(ranks);
    if (!levelled.some((rank) => rank.id === spellId)) continue;
    const learned = levelled.filter((rank) => rank.level <= level);
    return learned.length === 0 ? null : learned[learned.length - 1].id;
  }
  return null;
}

/** The spec's cast steps resolved at `level`, or `undefined` when its curated APL is not
 *  on disk. Steps whose ability is not learned yet are left out, as in the published
 *  line list. */
export function loadAplSteps(spec: string, families: SpellRanksByName, level: number): AplStep[] | undefined {
  const file = path.join(REPO_ROOT, 'data/curated/apl', `${spec}.json`);
  if (!existsSync(file)) return undefined;
  const steps: AplStep[] = [];
  for (const cast of castSteps(JSON.parse(readFileSync(file, 'utf8')))) {
    const resolvedSpellId = resolveAt(families, cast.spellId, level);
    if (resolvedSpellId !== null) steps.push({ resolvedSpellId, words: stepWords(cast) });
  }
  return steps;
}
