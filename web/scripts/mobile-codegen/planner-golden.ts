// web/scripts/mobile-codegen/planner-golden.ts
// The planner's golden vectors (spec 2026-10-31-mobile-wave3 F13 "Rules live in one place"):
// every add, remove, validate, points-available, split and FS1-to-order decision both apps
// must reproduce, computed here by the site planner's own functions (web/src/lib/planner)
// over the active build's real talent files. web/scripts/gen-planner-rules.mjs writes the
// result to apps/shared/fixtures/planner-rules.json; each app's tests replay every vector.
//
// The walks are seeded, so the file only changes when the rules, the talent files or the
// copy change. The generator fails when a refusal kind is missing from the cases, so the
// file never silently stops covering one.
import { readFileSync } from 'node:fs';
import { join } from 'node:path';
import { decodeFS1, orderFromRanks } from '../../src/lib/planner/fs1';
import { pointsAvailableAt, pointsPerTree } from '../../src/lib/planner/derive';
import {
  canAddPoint,
  canRemovePoint,
  indexTalents,
  validateOrder,
  withPoint,
  withoutLastPoint,
  type TalentIndex,
} from '../../src/lib/planner/rules';
import { FIRST_POINT_LEVEL, MAX_POINTS, POINTS_PER_TIER, type TalentFile } from '../../src/lib/planner/types';

/** The classes the vectors walk: one with deep prerequisite chains, one with many 5-rank rows. */
export const GOLDEN_CLASSES = ['shaman', 'warrior', 'druid'] as const;
const SEED = 0x5eed_f13;
const WALK_STEPS = 260;
const SAMPLE_EVERY = 3;
const LEVELS = [1, 9, 10, 11, 25, 59, 60, 61];
/** The guide fixtures whose build codes feed the FS1 vectors. */
const GUIDE_FIXTURES = ['shaman-elemental', 'mage-fire', 'warrior-protection'];

export interface GoldenTalent {
  id: number;
  name: string;
  tier: number;
  column: number;
  max_rank: number;
  prereq_talent_id: number | null;
  prereq_rank: number | null;
}

export interface GoldenClass {
  class_slug: string;
  trees: { id: number; name: string; position: number; talents: GoldenTalent[] }[];
}

export interface Decision {
  class_slug: string;
  level: number;
  order: number[];
  talent: number;
  ok: boolean;
  reason: string | null;
}

export interface GoldenFile {
  note: string;
  build: string;
  constants: { max_points: number; points_per_tier: number; first_point_level: number };
  /** The app-only budget refusal: a level holds `level - 9` points (spec F13.1). */
  none_left: string;
  classes: GoldenClass[];
  points_available: { level: number; points: number }[];
  add: Decision[];
  remove: Decision[];
  validate: { class_slug: string; order: number[]; errors: { field: string; message: string }[] }[];
  split: { class_slug: string; order: number[]; totals: number[] }[];
  fs1: {
    code: string;
    class_slug: string;
    race_slug: string;
    tree_ranks: number[][];
    order: number[];
    dropped: number[];
  }[];
}

/** mulberry32: a tiny seeded generator, so the file is the same on every run. */
function random(seed: number): () => number {
  let a = seed >>> 0;
  return () => {
    a = (a + 0x6d2b79f5) >>> 0;
    let t = a;
    t = Math.imul(t ^ (t >>> 15), t | 1);
    t ^= t + Math.imul(t ^ (t >>> 7), t | 61);
    return ((t ^ (t >>> 14)) >>> 0) / 4294967296;
  };
}

function slim(file: TalentFile): GoldenClass {
  return {
    class_slug: file.class_slug,
    trees: [...file.trees]
      .sort((a, b) => a.position - b.position)
      .map((tree) => ({
        id: tree.id,
        name: tree.name,
        position: tree.position,
        talents: tree.talents.map((t) => ({
          id: t.id,
          name: t.name,
          tier: t.tier,
          column: t.column,
          max_rank: t.max_rank,
          prereq_talent_id: t.prereq_talent_id,
          prereq_rank: t.prereq_rank,
        })),
      })),
  };
}

/** The app's add: the level budget first (spec F13.2), then the site planner's own rules. */
export function appAdd(index: TalentIndex, order: number[], talent: number, level: number, noneLeft: string) {
  if (order.length >= pointsAvailableAt(level)) return { ok: false, reason: noneLeft };
  const decision = canAddPoint(index, order, talent);
  return decision.ok ? { ok: true, reason: null } : { ok: false, reason: decision.reason };
}

interface Walk {
  add: Decision[];
  remove: Decision[];
  orders: number[][];
}

/**
 * A seeded walk: mostly adds (biased toward the tree the walk has favoured, so deep tiers
 * are reached), some removes, a level drawn now and then. Every attempt is a vector; the
 * order is recorded before the attempt.
 */
function walk(index: TalentIndex, classSlug: string, noneLeft: string, seed: number): Walk {
  const next = random(seed);
  const all = index.trees.flatMap((tree) => tree.talents.map((t) => t.id));
  const out: Walk = { add: [], remove: [], orders: [] };
  for (let round = 0; round < 3; round += 1) {
    let order: number[] = [];
    const favourite = index.trees[round % index.trees.length];
    const favoured = favourite.talents.map((t) => t.id);
    for (let step = 0; step < WALK_STEPS; step += 1) {
      const level = next() < 0.15 ? LEVELS[Math.floor(next() * LEVELS.length)] : 60;
      const removing = order.length > 0 && next() < 0.22;
      if (removing) {
        const talent =
          next() < 0.85 ? order[Math.floor(next() * order.length)] : all[Math.floor(next() * all.length)];
        const decision = canRemovePoint(index, order, talent);
        if (step % SAMPLE_EVERY === 0 || !decision.ok) {
          out.remove.push({
            class_slug: classSlug,
            level,
            order,
            talent,
            ok: decision.ok,
            reason: decision.ok ? null : decision.reason,
          });
        }
        if (decision.ok) order = withoutLastPoint(order, talent);
      } else {
        const pool = next() < 0.7 ? favoured : all;
        const talent = pool[Math.floor(next() * pool.length)];
        const decision = appAdd(index, order, talent, level, noneLeft);
        if (step % SAMPLE_EVERY === 0 || !decision.ok) {
          out.add.push({ class_slug: classSlug, level, order, talent, ...decision });
        }
        if (decision.ok) order = withPoint(order, talent);
      }
      if (step % 20 === 0) out.orders.push(order);
    }
    out.orders.push(order);
  }
  return out;
}

/** Spends points in `tree`, lowest tier first, skipping `avoid`, until it holds `points`. */
function fillTree(
  index: TalentIndex,
  order: number[],
  tree: TalentIndex['trees'][number],
  points: number,
  avoid: Set<number>,
): number[] {
  let next = order;
  const byTier = [...tree.talents].sort((a, b) => a.tier - b.tier || a.column - b.column);
  const inTree = () => next.filter((id) => index.treeOf.get(id)?.id === tree.id).length;
  let progress = true;
  while (inTree() < points && progress) {
    progress = false;
    for (const talent of byTier) {
      if (inTree() >= points) break;
      if (avoid.has(talent.id)) continue;
      if (canAddPoint(index, next, talent.id).ok) {
        next = withPoint(next, talent.id);
        progress = true;
      }
    }
  }
  return next;
}

/**
 * The prerequisite cases a random walk rarely reaches: for every talent with a prerequisite,
 * an add with the tier open and the prerequisite short, and a remove of the prerequisite's
 * last point while the talent holds one.
 */
function prerequisiteCases(index: TalentIndex, classSlug: string, noneLeft: string): Walk {
  const out: Walk = { add: [], remove: [], orders: [] };
  for (const tree of index.trees) {
    for (const talent of tree.talents.filter((t) => t.prereq_talent_id !== null)) {
      const prereq = talent.prereq_talent_id as number;
      const open = fillTree(index, [], tree, POINTS_PER_TIER * talent.tier, new Set([prereq, talent.id]));
      out.add.push({
        class_slug: classSlug,
        level: 60,
        order: open,
        talent: talent.id,
        ...appAdd(index, open, talent.id, 60, noneLeft),
      });
      let held = fillTree(index, [], tree, POINTS_PER_TIER * talent.tier, new Set([talent.id]));
      while (
        held.filter((id) => id === prereq).length < (talent.prereq_rank ?? 0) &&
        canAddPoint(index, held, prereq).ok
      ) {
        held = withPoint(held, prereq);
      }
      if (!canAddPoint(index, held, talent.id).ok) continue;
      held = withPoint(held, talent.id);
      const decision = canRemovePoint(index, held, prereq);
      out.remove.push({
        class_slug: classSlug,
        level: 60,
        order: held,
        talent: prereq,
        ok: decision.ok,
        reason: decision.ok ? null : decision.reason,
      });
      out.orders.push(held);
    }
  }
  return out;
}

/** Keeps at most `limit` vectors per distinct reason (and per ok), so the file stays small. */
function thin(decisions: Decision[], limit: number): Decision[] {
  const seen = new Map<string, number>();
  return decisions.filter((d) => {
    const kind = d.ok ? 'ok' : (d.reason ?? '').replace(/[A-Z][a-zA-Z' ]+/g, '').replace(/\d+/g, 'n');
    const count = seen.get(`${d.class_slug}|${kind}`) ?? 0;
    seen.set(`${d.class_slug}|${kind}`, count + 1);
    return count < limit;
  });
}

const KINDS: [string, RegExp][] = [
  ['tier', /^Tier \d+ of .+ needs \d+ points in .+ first$/],
  ['prerequisite', /^(?!Tier \d+ of ).+ needs \d+ points? in .+ first$/],
  ['max rank', /^.+ is already at \d+ of \d+ points$/],
  ['no points to remove', /^.+ has no points to remove$/],
];

function requireKinds(name: string, decisions: Decision[], kinds: [string, RegExp][]): void {
  for (const [kind, pattern] of kinds) {
    if (!decisions.some((d) => !d.ok && pattern.test(d.reason ?? ''))) {
      throw new Error(`planner golden vectors: no ${name} case refused for ${kind}`);
    }
  }
}

function fs1Vectors(repoRoot: string, indexes: Map<string, TalentIndex>): GoldenFile['fs1'] {
  const codes = GUIDE_FIXTURES.flatMap((name) => {
    const guide = JSON.parse(
      readFileSync(join(repoRoot, `docs/contracts/fixtures/app/cdn.guides.${name}.json`), 'utf8'),
    );
    return [guide.talents?.leveling_build?.code, guide.talents?.raid_build?.code].filter(
      (code): code is string => typeof code === 'string',
    );
  });
  return codes.flatMap((code) => {
    const decoded = decodeFS1(code);
    if (!decoded.ok) throw new Error(`planner golden vectors: ${code}: ${decoded.message}`);
    const index = indexes.get(decoded.build.classSlug);
    if (!index) return [];
    const { order, dropped } = orderFromRanks(index, decoded.build.treeRanks);
    return [
      {
        code,
        class_slug: decoded.build.classSlug,
        race_slug: decoded.build.raceSlug,
        tree_ranks: decoded.build.treeRanks,
        order,
        dropped,
      },
    ];
  });
}

export function plannerGolden(repoRoot: string): GoldenFile {
  const build = JSON.parse(readFileSync(join(repoRoot, 'web/src/data/active-build.json'), 'utf8'))
    .build as string;
  const copy = JSON.parse(readFileSync(join(repoRoot, 'apps/shared/copy/en.json'), 'utf8')).strings;
  const noneLeft = copy['plan.refuse.noneLeft'] as string;
  const files = GOLDEN_CLASSES.map(
    (slug) =>
      JSON.parse(
        readFileSync(join(repoRoot, `data/builds/${build}/talents/${slug}.json`), 'utf8'),
      ) as TalentFile,
  );
  const indexes = new Map(files.map((file) => [file.class_slug, indexTalents(file)]));

  const add: Decision[] = [];
  const remove: Decision[] = [];
  const validate: GoldenFile['validate'] = [];
  const split: GoldenFile['split'] = [];
  files.forEach((file, i) => {
    const index = indexes.get(file.class_slug)!;
    const w = walk(index, file.class_slug, noneLeft, SEED + i);
    const p = prerequisiteCases(index, file.class_slug, noneLeft);
    add.push(...thin(w.add, 14), ...p.add.slice(0, 6));
    remove.push(...thin(w.remove, 14), ...p.remove.slice(0, 6));
    w.orders.push(...p.orders.slice(0, 3));
    for (const order of w.orders)
      split.push({ class_slug: file.class_slug, order, totals: pointsPerTree(index, order) });
    // validate: each walked order as is, and reversed (which breaks tiers and prerequisites).
    for (const order of w.orders.filter((o) => o.length > 0).slice(0, 8)) {
      validate.push({ class_slug: file.class_slug, order, errors: validateOrder(index, order) });
      const reversed = [...order].reverse();
      validate.push({ class_slug: file.class_slug, order: reversed, errors: validateOrder(index, reversed) });
    }
    const unknown = [...(w.orders.at(-1) ?? []), 999_999_999];
    validate.push({ class_slug: file.class_slug, order: unknown, errors: validateOrder(index, unknown) });
  });

  requireKinds('add', add, [
    ...KINDS.filter(([k]) => k !== 'no points to remove'),
    ['no points left', /^No points left\./],
  ]);
  requireKinds('remove', remove, [
    ['no points to remove', KINDS[3][1]],
    ['a talent that needs the point', KINDS[1][1]],
  ]);

  return {
    note: "Generated by web/scripts/gen-planner-rules.mjs from web/src/lib/planner over the active build's talent files. Generated, do not edit. Both apps replay every vector (spec 2026-10-31-mobile-wave3 F13): add is the level budget then canAddPoint, remove is canRemovePoint, validate is validateOrder, split is pointsPerTree, fs1 is decodeFS1 then orderFromRanks.",
    build,
    constants: {
      max_points: MAX_POINTS,
      points_per_tier: POINTS_PER_TIER,
      first_point_level: FIRST_POINT_LEVEL,
    },
    none_left: noneLeft,
    classes: files.map(slim),
    points_available: [...LEVELS, 0, 30, 70].map((level) => ({ level, points: pointsAvailableAt(level) })),
    add,
    remove,
    validate,
    split,
    fs1: fs1Vectors(repoRoot, indexes),
  };
}

/** The file's bytes: one vector per line, so a diff names the vector that moved. */
export function renderGolden(golden: GoldenFile): string {
  const lines = (rows: unknown[]) => rows.map((row) => `    ${JSON.stringify(row)}`).join(',\n');
  const head = {
    note: golden.note,
    build: golden.build,
    constants: golden.constants,
    none_left: golden.none_left,
  };
  const body = [
    JSON.stringify(head, null, 2).slice(0, -2),
    `  "classes": ${JSON.stringify(golden.classes)}`,
    ...(['points_available', 'add', 'remove', 'validate', 'split', 'fs1'] as const).map(
      (key) => `  "${key}": [\n${lines(golden[key])}\n  ]`,
    ),
  ];
  return `${body.join(',\n')}\n}\n`;
}
