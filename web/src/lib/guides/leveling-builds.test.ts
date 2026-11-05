// web/src/lib/guides/leveling-builds.test.ts
import activeBuild from '../../data/active-build.json';
import { existsSync, readdirSync, readFileSync } from 'node:fs';
import path from 'node:path';
import { fileURLToPath } from 'node:url';
import { parse as parseYaml } from 'yaml';
import { describe, expect, it } from 'vitest';
import { readSpecCatalog } from '../bis/load';
import { indexTalents, validateOrder } from '../planner/rules';
import {
  buildCodeDigits,
  guideLadderSteps,
  guideLevelingBuilds,
  guideLadderOrder,
  guideTargets,
  ladderFirstFor,
  ladderPointOrder,
  ladderTalentString,
  loadLadderTrees,
  type LadderTalentTree,
} from './leveling-builds';

const BUILD = activeBuild.build;
const REPO = path.resolve(path.dirname(fileURLToPath(import.meta.url)), '../../../..');
const GUIDES = path.join(REPO, 'web/src/content/guides');
const GOLDENS = path.join(REPO, 'sim/request/testdata/ladder');

interface GuideBuildRow {
  spec: string;
  classSlug: string;
  build: string;
  raidBuild?: string;
  treeIndex: number;
}

function guideBuilds(): GuideBuildRow[] {
  const catalog = readSpecCatalog();
  return readdirSync(GUIDES, { withFileTypes: true })
    .filter((entry) => entry.isDirectory())
    .flatMap((entry) =>
      readdirSync(path.join(GUIDES, entry.name))
        .filter((file) => file.endsWith('.md') && file !== 'index.md')
        .flatMap((file): GuideBuildRow[] => {
          const [, front] = readFileSync(path.join(GUIDES, entry.name, file), 'utf8').split('---');
          const data = parseYaml(front ?? '') as { build?: string; raidBuild?: string };
          const specSlug = file.slice(0, -'.md'.length);
          const row = catalog.find((s) => s.class_slug === entry.name && s.spec_slug === specSlug);
          if (data.build === undefined || row === undefined) return [];
          return [
            {
              spec: `${entry.name}-${specSlug}`,
              classSlug: entry.name,
              build: data.build,
              raidBuild: data.raidBuild,
              treeIndex: row.tree_index,
            },
          ];
        }),
    );
}

/** The golden's ladder table: talent string by level. */
function goldenTalents(spec: string): Map<number, string> {
  const file = path.join(GOLDENS, `${spec}.golden.md`);
  const rows = new Map<number, string>();
  if (!existsSync(file)) return rows;
  const ladder = readFileSync(file, 'utf8').split('## Ladder')[1]?.split('\n## ')[0] ?? '';
  for (const line of ladder.split('\n')) {
    const match = /^\| (\d+) \| ([0-9-]+) \|/.exec(line);
    if (match !== null) rows.set(Number(match[1]), match[2]!);
  }
  return rows;
}

function node(id: number, tier: number, column: number, maxRank: number, prereq: number | null = null) {
  return {
    id,
    name: `T${id}`,
    tier,
    column,
    max_rank: maxRank,
    prereq_talent_id: prereq,
    prereq_rank: prereq === null ? null : maxRank,
  };
}

describe('ladderTalentString', () => {
  const trees: LadderTalentTree[] = [
    { name: 'A', position: 0, talents: [node(1, 0, 0, 5), node(2, 1, 0, 5), node(3, 2, 0, 1, 2)] },
    { name: 'B', position: 1, talents: [node(4, 0, 0, 5)] },
  ];
  const targets = new Map([
    [1, 5],
    [2, 5],
    [3, 1],
    [4, 3],
  ]);

  it('spends level minus nine points, own tree first, top row down', () => {
    expect(ladderTalentString(trees, targets, 0, 10)).toBe('100-0');
    expect(ladderTalentString(trees, targets, 0, 20)).toBe('551-0');
    expect(ladderTalentString(trees, targets, 0, 30)).toBe('551-3');
    expect(ladderTalentString(trees, targets, 1, 15)).toBe('300-3');
  });

  it('never spends past the guide build', () => {
    expect(ladderTalentString(trees, targets, 0, 60)).toBe('551-3');
  });

  it('reads a build code by tree', () => {
    expect(buildCodeDigits('FS1:1.60.1.70753:warrior:human:0332/0505/005:')).toEqual(['0332', '0505', '005']);
    expect(buildCodeDigits('not a code')).toBeUndefined();
    expect([...guideTargets(trees, ['55', '3'])]).toEqual([
      [1, 5],
      [2, 5],
      [3, 0],
      [4, 3],
    ]);
  });
});

describe('the guide leveling builds match the sim’s ladder', () => {
  it.each(guideBuilds().map((row) => [row.spec, row] as const))('%s', (spec, row) => {
    const golden = goldenTalents(spec);
    const trees = loadLadderTrees(BUILD, row.classSlug)!;
    const first = ladderFirstFor(spec);
    const checked: number[] = [];
    for (const [level, talents] of golden) {
      // The ladder cuts down the guide’s level 60 build at every level.
      const code = row.build;
      const digits = buildCodeDigits(code);
      if (digits === undefined) continue;
      expect(
        ladderTalentString(trees, guideTargets(trees, digits), row.treeIndex, level, first),
        `level ${level}`,
      ).toBe(talents);
      checked.push(level);
    }
    expect(checked.length).toBeGreaterThan(0);
  });

  it('gives the strip one build per range, 60 being the guide’s own', () => {
    const frost = guideBuilds().find((row) => row.spec === 'mage-frost')!;
    const builds = guideLevelingBuilds(BUILD, 'mage', 'mage-frost', frost.treeIndex, frost.build);
    expect(builds.map((b) => b.band)).toEqual([20, 30, 40, 50, 60]);
    expect(builds.map((b) => b.points)).toEqual([11, 21, 31, 41, 51]);
    const total = (talents: string): number =>
      [...talents.replace(/-/g, '')].reduce((sum, d) => sum + Number(d), 0);
    expect(builds.every((b) => total(b.talents) <= b.points)).toBe(true);
  });
});

describe('every guide ladder only adds points', () => {
  // No character takes a point back without a respec (sim/leveling
  // TestEveryLadderOnlyAddsPoints); Protection Warrior's walk dropped Shield Slam at 41.
  it.each(guideBuilds().map((row) => [row.spec, row] as const))('%s', (spec, row) => {
    const trees = loadLadderTrees(BUILD, row.classSlug)!;
    const digits = buildCodeDigits(row.build);
    if (digits === undefined) return;
    const targets = guideTargets(trees, digits);
    const first = ladderFirstFor(spec);
    let prev = ladderTalentString(trees, targets, row.treeIndex, 10, first);
    for (let level = 11; level <= 60; level++) {
      const at = ladderTalentString(trees, targets, row.treeIndex, level, first);
      const dropped = [...at].some((d, i) => d !== '-' && Number(d) < Number(prev[i]));
      expect(dropped, `level ${level}: ${at} after ${prev}`).toBe(false);
      prev = at;
    }
  });
});

describe('the ladder’s point order', () => {
  // Sim-accuracy round 16: the Leveling strip and the low-level home follow the order the
  // ladder spends its points in, one a level from 10, not the planner's reconstruction.
  it.each(guideBuilds().map((row) => [row.spec, row] as const))(
    '%s replays the ladder level by level',
    (spec, row) => {
      const trees = loadLadderTrees(BUILD, row.classSlug)!;
      const digits = buildCodeDigits(row.build);
      if (digits === undefined) return;
      const targets = guideTargets(trees, digits);
      const first = ladderFirstFor(spec);
      const order = ladderPointOrder(trees, targets, row.treeIndex, 60, first);
      const cellOf = new Map(
        trees.flatMap((tree, ti) => tree.talents.map((t, j) => [t.id, [ti, j]] as const)),
      );
      const climb = trees.map((tree) => new Array<number>(tree.talents.length).fill(0));
      order.forEach((id, point) => {
        const [ti, j] = cellOf.get(id)!;
        climb[ti]![j]! += 1;
        const level = 10 + point;
        expect(climb.map((r) => r.join('')).join('-'), `level ${level}`).toBe(
          ladderTalentString(trees, targets, row.treeIndex, level, first),
        );
      });
      // And the order is one the planner accepts point by point (its band loads read it).
      const talentFile = JSON.parse(
        readFileSync(path.join(REPO, 'data/builds', BUILD, 'talents', `${row.classSlug}.json`), 'utf8'),
      );
      expect(validateOrder(indexTalents(talentFile), order)).toEqual([]);
    },
  );

  it('starts Shadow on Wand Specialization at 10 and 11, then Spirit Tap from 12', () => {
    const shadow = guideBuilds().find((row) => row.spec === 'priest-shadow')!;
    const order = guideLadderOrder(BUILD, 'priest', 'priest-shadow', shadow.treeIndex, shadow.build);
    const trees = loadLadderTrees(BUILD, 'priest')!;
    const nameOf = new Map(trees.flatMap((tree) => tree.talents.map((t) => [t.id, t.name] as const)));
    const names = order.map((id) => nameOf.get(id));
    expect(names.slice(0, 7)).toEqual([
      'Wand Specialization',
      'Wand Specialization',
      'Spirit Tap',
      'Spirit Tap',
      'Spirit Tap',
      'Spirit Tap',
      'Spirit Tap',
    ]);
  });

  // Sim-accuracy round 24: a reservation is spent in its own order across trees, so Fire's
  // Wand Specialization (Arcane) comes right after Combustion, not after the whole Fire tree.
  it('takes Fire’s Wand Specialization at 41 and 42, right after Combustion', () => {
    const fire = guideBuilds().find((row) => row.spec === 'mage-fire')!;
    const order = guideLadderOrder(BUILD, 'mage', 'mage-fire', fire.treeIndex, fire.build);
    const trees = loadLadderTrees(BUILD, 'mage')!;
    const nameOf = new Map(trees.flatMap((tree) => tree.talents.map((t) => [t.id, t.name] as const)));
    const at = (level: number) => nameOf.get(order[level - 10]!);
    expect([at(40), at(41), at(42)]).toEqual(['Combustion', 'Wand Specialization', 'Wand Specialization']);
  });
});

describe('guideLadderSteps', () => {
  it.each(['rogue-combat', 'priest-shadow', 'hunter-beast-mastery'])(
    '%s: one step a level from 10 to 60, each the ladder’s own build at that level',
    (spec) => {
      const guide = guideBuilds().find((row) => row.spec === spec);
      if (guide === undefined) throw new Error(`${spec} has no guide build`);
      const steps = guideLadderSteps(BUILD, guide.classSlug, spec, guide.treeIndex, guide.build);
      expect(steps.map((step) => step.level)).toEqual(Array.from({ length: 51 }, (_, i) => i + 10));
      const trees = loadLadderTrees(BUILD, guide.classSlug)!;
      const targets = guideTargets(trees, buildCodeDigits(guide.build)!);
      const first = ladderFirstFor(spec);
      for (const step of steps) {
        expect(step.talents, `${spec} level ${step.level}`).toBe(
          ladderTalentString(trees, targets, guide.treeIndex, step.level, first),
        );
      }
      // The points, in order, are the ladder's own order, each talent's rank counted up.
      const order = guideLadderOrder(BUILD, guide.classSlug, spec, guide.treeIndex, guide.build);
      expect(steps.flatMap((step) => (step.talentId === null ? [] : [step.talentId]))).toEqual(order);
      const held = new Map<number, number>();
      for (const step of steps) {
        if (step.talentId === null) continue;
        held.set(step.talentId, (held.get(step.talentId) ?? 0) + 1);
        expect(step.rank).toBe(held.get(step.talentId));
        expect(step.name).not.toBeNull();
      }
    },
  );

  it('is empty for a code that does not read', () => {
    expect(guideLadderSteps(BUILD, 'rogue', 'rogue-combat', 1, 'not a code')).toEqual([]);
  });
});
