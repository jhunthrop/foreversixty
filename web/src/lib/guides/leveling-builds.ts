// web/src/lib/guides/leveling-builds.ts
// The guide's leveling builds: its own level 60 build cut down to the points a character of
// each level range has spent, the same walk the sim's rotation ladder and the BiS ranker use
// (sim/leveling LadderTalentString and ForSpec, ported line for line; a test holds this port
// to the committed ladder goldens). Read from the guide's own build code, never from the
// published BiS file, so the Leveling strip follows the guide the moment its build changes.
import { existsSync, readFileSync } from 'node:fs';
import path from 'node:path';
import { repoRoot } from '../data/build-files';
import { FIRST_POINT_LEVEL } from '../planner/types';

export interface LadderTalentNode {
  id: number;
  name: string;
  tier: number;
  column: number;
  max_rank: number;
  prereq_talent_id: number | null;
  prereq_rank: number | null;
}

export interface LadderTalentTree {
  name: string;
  position: number;
  talents: LadderTalentNode[];
}

export interface LadderFirstRank {
  talent: string;
  rank: number;
}

type Targets = ReadonlyMap<number, number>;
type Ranks = Map<number, number>;

/** The planner's tier gate (web/src/lib/planner/types.ts POINTS_PER_TIER). */
const POINTS_PER_TIER = 5;

/** The level ranges the Leveling strip shows: 20 to 29 up to 60. */
export const LEVELING_BANDS: readonly number[] = [20, 30, 40, 50, 60];
/** The level the ladder ends on: the guide's own build. */
const MAX_LADDER_LEVEL = 60;
/** Levels per leveling range ("Level 30 to 39"). */
const LEVELS_PER_BAND = 10;

/** The class's trees, by position, each sorted top row down, left to right (LoadTalentTrees). */
export function loadLadderTrees(build: string, classSlug: string): LadderTalentTree[] | undefined {
  const file = path.join(repoRoot(), 'data/builds', build, 'talents', `${classSlug}.json`);
  if (!existsSync(file)) return undefined;
  const trees = (JSON.parse(readFileSync(file, 'utf8')) as { trees: LadderTalentTree[] }).trees;
  return [...trees]
    .sort((a, b) => a.position - b.position)
    .map((tree) => ({
      ...tree,
      talents: [...tree.talents].sort((a, b) => a.tier - b.tier || a.column - b.column),
    }));
}

/** The per-tree digit strings of an FS1 build code, or undefined for a code that is not one. */
export function buildCodeDigits(code: string): [string, string, string] | undefined {
  const match = /^FS1:[^:]+:[^:]+:[^:]+:([^/]+)\/([^/]+)\/([^:]+):$/.exec(code);
  return match === null ? undefined : [match[1]!, match[2]!, match[3]!];
}

/** Each talent's guide rank by id, read positionally against `trees` (GuideTalentTargets). */
export function guideTargets(
  trees: readonly LadderTalentTree[],
  digits: readonly string[],
): Map<number, number> {
  const targets = new Map<number, number>();
  trees.forEach((tree, i) => {
    const row = digits[i] ?? '';
    tree.talents.forEach((node, j) => {
      const rank = Number.parseInt(row[j] ?? '', 10);
      targets.set(node.id, Number.isNaN(rank) ? 0 : rank);
    });
  });
  return targets;
}

function target(targets: Targets, id: number): number {
  return targets.get(id) ?? 0;
}

function spent(ranks: ReadonlyMap<number, number>): number {
  let total = 0;
  for (const rank of ranks.values()) total += rank;
  return total;
}

function spendTree(
  digits: number[],
  tree: LadderTalentTree,
  targets: Targets,
  reserved: ReadonlyMap<number, number> | undefined,
  budget: number,
): number {
  let left = budget;
  tree.talents.forEach((node, j) => {
    digits[j] = reserved?.get(node.id) ?? 0;
    left -= digits[j]!;
  });
  for (let j = 0; j < tree.talents.length; j++) {
    if (left <= 0) break;
    const node = tree.talents[j]!;
    const extra = Math.min(target(targets, node.id), node.max_rank) - digits[j]!;
    const give = Math.max(Math.min(extra, left), 0);
    digits[j]! += give;
    left -= give;
  }
  return left;
}

function keepsEveryTalent(before: readonly number[], after: readonly number[]): boolean {
  return before.every((rank, j) => rank === 0 || after[j] !== 0);
}

function pathToDeepestTarget(tree: LadderTalentTree, targets: Targets): LadderTalentNode[] {
  const byId = new Map(tree.talents.map((node) => [node.id, node]));
  let deepest: LadderTalentNode | undefined;
  for (const node of tree.talents) {
    if (target(targets, node.id) > 0 && (deepest === undefined || node.tier > deepest.tier)) deepest = node;
  }
  if (deepest === undefined) return [];
  const chain = [deepest];
  while (chain.length <= tree.talents.length) {
    const prereq = chain[0]!.prereq_talent_id ?? 0;
    const parent = byId.get(prereq);
    if (prereq === 0 || parent === undefined) break;
    chain.unshift(parent);
  }
  return chain;
}

function spineRanks(chain: readonly LadderTalentNode[]): Ranks {
  const ranks: Ranks = new Map();
  chain.forEach((step, j) => {
    const rank = j + 1 < chain.length ? Math.max(chain[j + 1]!.prereq_rank ?? 0, 1) : 1;
    ranks.set(step.id, Math.min(rank, step.max_rank));
  });
  return ranks;
}

function pointsInRowsAbove(tree: LadderTalentTree, targets: Targets, tier: number): number {
  return tree.talents
    .filter((node) => node.tier < tier)
    .reduce((total, node) => total + Math.min(target(targets, node.id), node.max_rank), 0);
}

function plainWalkReaches(tree: LadderTalentTree, targets: Targets, budget: number, id: number): boolean {
  const digits = new Array<number>(tree.talents.length).fill(0);
  spendTree(digits, tree, targets, undefined, budget);
  const j = tree.talents.findIndex((node) => node.id === id);
  return j >= 0 && digits[j]! > 0;
}

function ownTreeSpine(tree: LadderTalentTree, targets: Targets, budget: number): Ranks | undefined {
  const chain = pathToDeepestTarget(tree, targets);
  if (chain.length === 0) return undefined;
  const capstone = chain[chain.length - 1]!;
  const gate = POINTS_PER_TIER * capstone.tier;
  if (capstone.tier === 0 || gate + 1 > budget || pointsInRowsAbove(tree, targets, capstone.tier) < gate) {
    return undefined;
  }
  if (plainWalkReaches(tree, targets, budget, capstone.id)) return undefined;
  const spine = spineRanks(chain);
  return spent(spine) > budget ? undefined : spine;
}

function spendOwnTree(digits: number[], tree: LadderTalentTree, targets: Targets, budget: number): number {
  const plain = new Array<number>(digits.length).fill(0);
  const plainLeft = spendTree(plain, tree, targets, undefined, budget);
  const spine = ownTreeSpine(tree, targets, budget);
  if (spine !== undefined) {
    const trial = new Array<number>(digits.length).fill(0);
    const left = spendTree(trial, tree, targets, spine, budget);
    if (keepsEveryTalent(plain, trial)) {
      trial.forEach((rank, j) => (digits[j] = rank));
      return left;
    }
  }
  plain.forEach((rank, j) => (digits[j] = rank));
  return plainLeft;
}

function reservedRanks(tree: LadderTalentTree, targets: Targets, first: readonly LadderFirstRank[]): Ranks[] {
  const byName = new Map(tree.talents.map((node) => [node.name, node]));
  const steps: Ranks[] = [];
  let acc: Ranks = new Map();
  for (const row of first) {
    const node = byName.get(row.talent);
    if (node === undefined) continue;
    const top = Math.min(row.rank, target(targets, node.id), node.max_rank);
    for (let rank = (acc.get(node.id) ?? 0) + 1; rank <= top; rank++) {
      acc = new Map(acc).set(node.id, rank);
      steps.push(acc);
    }
  }
  return steps;
}

function treeIsLegal(tree: LadderTalentTree, digits: readonly number[]): boolean {
  const rankOf = new Map(tree.talents.map((node, j) => [node.id, digits[j]!]));
  return tree.talents.every((node, j) => {
    if (digits[j] === 0) return true;
    const above = tree.talents.reduce(
      (total, other, k) => total + (other.tier < node.tier ? digits[k]! : 0),
      0,
    );
    if (above < POINTS_PER_TIER * node.tier) return false;
    const prereq = node.prereq_talent_id ?? 0;
    return prereq === 0 || (rankOf.get(prereq) ?? 0) >= Math.max(node.prereq_rank ?? 0, 1);
  });
}

function spendReserved(
  digits: number[],
  tree: LadderTalentTree,
  targets: Targets,
  steps: readonly Ranks[],
  budget: number,
): number {
  for (let k = steps.length; k > 0; k--) {
    const reserve = steps[k - 1]!;
    if (spent(reserve) > budget) continue;
    const trial = new Array<number>(digits.length).fill(0);
    const left = spendTree(trial, tree, targets, reserve, budget);
    if (treeIsLegal(tree, trial)) {
      trial.forEach((rank, j) => (digits[j] = rank));
      return left;
    }
  }
  return spendTree(digits, tree, targets, undefined, budget);
}

/** The talent string ("a-b-c", one digit per talent) for `level` (LadderTalentString): the
 *  ladder climbed one point a level from 10, so a level never takes back a point the level
 *  below holds (climbLadder). */
export function ladderTalentString(
  trees: readonly LadderTalentTree[],
  targets: Targets,
  ownTreeIndex: number,
  level: number,
  first: readonly LadderFirstRank[] = [],
): string {
  return climbLadder(trees, targets, ownTreeIndex, level, first)
    .climb.map((row) => row.join(''))
    .join('-');
}

/** The talent each point goes into, one a level from 10 up to `level`, in the order the
 *  ladder spends them (sim/leveling LadderPointOrder): the order the Leveling strip and the
 *  low-level home show, so they agree with the sim's ladder and the addon. */
export function ladderPointOrder(
  trees: readonly LadderTalentTree[],
  targets: Targets,
  ownTreeIndex: number,
  level: number,
  first: readonly LadderFirstRank[] = [],
): number[] {
  return climbLadder(trees, targets, ownTreeIndex, level, first).order;
}

/** The planner's rule for adding one point (sim/leveling climbLadder's `legal`). */
function canAddPoint(tree: LadderTalentTree, digits: readonly number[], j: number): boolean {
  const node = tree.talents[j]!;
  const inTree = digits.reduce((sum, d) => sum + d, 0);
  if (inTree < POINTS_PER_TIER * node.tier || digits[j]! >= node.max_rank) return false;
  if (node.prereq_talent_id === null) return true;
  const prereq = tree.talents.findIndex((other) => other.id === node.prereq_talent_id);
  const held = prereq >= 0 ? digits[prereq]! : 0;
  return held >= (node.prereq_rank ?? 0);
}

/** One point a level from 10 up to `level`, never removed (sim/leveling climbLadder): the
 *  first legal cell, own tree first, then tier and column, where that level's walk holds
 *  more ranks than the climb so far, else where the guide's targets do. */
function climbLadder(
  trees: readonly LadderTalentTree[],
  targets: Targets,
  ownTreeIndex: number,
  level: number,
  first: readonly LadderFirstRank[],
): { climb: number[][]; order: number[] } {
  const climb = trees.map((tree) => new Array<number>(tree.talents.length).fill(0));
  const order: number[] = [];
  const treeOrder = [ownTreeIndex, ...trees.map((_, i) => i).filter((i) => i !== ownTreeIndex)].filter(
    (ti) => trees[ti] !== undefined,
  );
  const firstCell = (wants: (ti: number, j: number) => number): [number, number] | undefined => {
    for (const ti of treeOrder) {
      const tree = trees[ti]!;
      for (let j = 0; j < tree.talents.length; j++) {
        if (wants(ti, j) > climb[ti]![j]! && canAddPoint(tree, climb[ti]!, j)) return [ti, j];
      }
    }
    return undefined;
  };
  for (let lvl = 10; lvl <= level; lvl++) {
    const walk = walkLevel(trees, targets, ownTreeIndex, lvl, first);
    const pick =
      firstCell((ti, j) => walk[ti]![j]!) ??
      firstCell((ti, j) => {
        const node = trees[ti]!.talents[j]!;
        return Math.min(target(targets, node.id), node.max_rank);
      });
    if (pick === undefined) continue;
    const row = climb[pick[0]]!;
    row[pick[1]] = row[pick[1]]! + 1;
    order.push(trees[pick[0]]!.talents[pick[1]]!.id);
  }
  return { climb, order };
}

/** The walk for one level on its own, as digits per tree (sim/leveling walkLevel). */
function walkLevel(
  trees: readonly LadderTalentTree[],
  targets: Targets,
  ownTreeIndex: number,
  level: number,
  first: readonly LadderFirstRank[],
): number[][] {
  let remaining = Math.max(level - 9, 0);
  const digits = trees.map((tree) => new Array<number>(tree.talents.length).fill(0));
  const order = [ownTreeIndex, ...trees.map((_, i) => i).filter((i) => i !== ownTreeIndex)];
  const offTree = reserveOffTree(trees, targets, ownTreeIndex, first, remaining);
  remaining = offTree.budget;
  for (const ti of order) {
    const tree = trees[ti];
    if (tree === undefined) continue;
    if (ti === ownTreeIndex && first.length > 0) {
      remaining = spendReserved(digits[ti]!, tree, targets, reservedRanks(tree, targets, first), remaining);
    } else if (ti === ownTreeIndex) {
      remaining = spendOwnTree(digits[ti]!, tree, targets, remaining);
    } else {
      const reserved = offTree.reserved.get(ti);
      remaining = spendTree(
        digits[ti]!,
        tree,
        targets,
        reserved,
        remaining + (reserved ? spent(reserved) : 0),
      );
    }
  }
  return digits;
}

/** A reservation's ranks in another tree than the spec's own (Shadow's Wand
 *  Specialization, in Discipline), spent before the own tree one rank at a time in list
 *  order, each capped at the guide's target and its max rank, and only for a first-row
 *  talent with no prerequisite (sim/leveling reserveOffTree). */
function reserveOffTree(
  trees: readonly LadderTalentTree[],
  targets: Targets,
  ownTreeIndex: number,
  first: readonly LadderFirstRank[],
  budget: number,
): { reserved: Map<number, Ranks>; budget: number } {
  const reserved = new Map<number, Ranks>();
  let left = budget;
  for (const row of first) {
    trees.forEach((tree, ti) => {
      if (ti === ownTreeIndex) return;
      for (const node of tree.talents) {
        if (node.name !== row.talent || node.tier !== 0 || node.prereq_talent_id !== null) continue;
        const ranks = reserved.get(ti) ?? new Map<number, number>();
        reserved.set(ti, ranks);
        const cap = Math.min(row.rank, target(targets, node.id), node.max_rank);
        while ((ranks.get(node.id) ?? 0) < cap && left > 0) {
          ranks.set(node.id, (ranks.get(node.id) ?? 0) + 1);
          left--;
        }
      }
    });
  }
  return { reserved, budget: left };
}

interface LadderFirstFile {
  specs: Record<string, { first: LadderFirstRank[] }>;
}

function curatedReservation(file: string, spec: string): LadderFirstRank[] {
  const full = path.join(repoRoot(), 'data/curated', file);
  if (!existsSync(full)) return [];
  return (JSON.parse(readFileSync(full, 'utf8')) as LadderFirstFile).specs[spec]?.first ?? [];
}

/** The spec's curated leveling reservation, as sim/leveling ForSpec spends it: its
 *  ladder-first.json row, then its hand-set order from ladder-order.json (an order that
 *  moves no checkpoint build, sim/leveling/ladder_order.go). */
export function ladderFirstFor(spec: string): LadderFirstRank[] {
  return [...curatedReservation('ladder-first.json', spec), ...curatedReservation('ladder-order.json', spec)];
}

export interface LevelingBuild {
  band: number;
  /** Points at the range's first level, and the build there. */
  points: number;
  talents: string;
  /** The range's last level (59 for 50 to 59; 60 for 60), and the build a player holds there:
   *  the Leveling strip lists this one, so a reader anywhere in the range sees every point. */
  lastLevel: number;
  lastTalents: string;
  /** The first level whose point the range's list names: the range's own first level, or the
   *  first talent point (10) for the first range, which also carries the points before it. */
  pointsFrom: number;
}

/** The guide's ladder point order through level 60, from its own level 60 code: one talent
 *  id per point (`ladderPointOrder`), the one order the guide's Leveling strip, the planner's
 *  band loads and the low-level home read; empty when the build or the talent file does not read. */
export function guideLadderOrder(
  build: string,
  classSlug: string,
  spec: string,
  ownTreeIndex: number,
  code: string,
): number[] {
  const trees = loadLadderTrees(build, classSlug);
  const digits = buildCodeDigits(code);
  if (trees === undefined || digits === undefined) return [];
  return ladderPointOrder(
    trees,
    guideTargets(trees, digits),
    ownTreeIndex,
    MAX_LADDER_LEVEL,
    ladderFirstFor(spec),
  );
}

/** The guide's build at each leveling range's first and last level, from its own level 60 code. */
export function guideLevelingBuilds(
  build: string,
  classSlug: string,
  spec: string,
  ownTreeIndex: number,
  code: string,
): LevelingBuild[] {
  const trees = loadLadderTrees(build, classSlug);
  const digits = buildCodeDigits(code);
  if (trees === undefined || digits === undefined) return [];
  const targets = guideTargets(trees, digits);
  const first = ladderFirstFor(spec);
  return LEVELING_BANDS.map((band, index) => {
    const lastLevel = Math.min(band + LEVELS_PER_BAND - 1, MAX_LADDER_LEVEL);
    return {
      band,
      points: Math.max(band - 9, 0),
      talents: ladderTalentString(trees, targets, ownTreeIndex, band, first),
      lastLevel,
      lastTalents: ladderTalentString(trees, targets, ownTreeIndex, lastLevel, first),
      pointsFrom: index === 0 ? FIRST_POINT_LEVEL : band,
    };
  });
}
