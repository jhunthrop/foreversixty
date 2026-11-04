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
  guideLevelingBuilds,
  guideTargets,
  ladderFirstFor,
  ladderTalentString,
  loadLadderTrees,
  type LadderTalentTree,
  guideLadderOrder,
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

describe('the ladder order every leveling surface reads', () => {
  it.each(guideBuilds().map((row) => [row.spec, row] as const))(
    '%s: its first N points are the ladder’s build at level 9 + N, every level 10 to 60',
    (spec, row) => {
      const trees = loadLadderTrees(BUILD, row.classSlug)!;
      const digits = buildCodeDigits(row.build)!;
      const targets = guideTargets(trees, digits);
      const first = ladderFirstFor(spec);
      const order = guideLadderOrder(BUILD, row.classSlug, spec, row.treeIndex, row.build);
      const cellOf = new Map(
        trees.flatMap((tree, ti) => tree.talents.map((talent, cell) => [talent.id, [ti, cell]] as const)),
      );
      const ladderAt = (level: number) => ladderTalentString(trees, targets, row.treeIndex, level, first);
      // The levels at which the ladder moves a point from one talent to another: before one, the
      // order already shows the talent the ladder settles on (one point differs, never more).
      const moves = Array.from({ length: 51 }, (_, i) => i + 10).filter((level) =>
        [...ladderAt(level).replace(/-/g, '')].some(
          (digit, i) => Number(digit) < Number(ladderAt(level - 1).replace(/-/g, '')[i]),
        ),
      );
      const ranks = trees.map((tree) => new Array<number>(tree.talents.length).fill(0));
      for (let level = 10; level <= 60; level += 1) {
        const id = order[level - 10];
        if (id !== undefined) {
          const [ti, cell] = cellOf.get(id)!;
          ranks[ti]![cell]! += 1;
        }
        const got = ranks.map((r) => r.join('')).join('-');
        const want = ladderAt(level);
        if (moves.some((move) => move > level)) {
          const flat = want.replace(/-/g, '');
          const off = [...got.replace(/-/g, '')].reduce(
            (sum, digit, i) => sum + Math.abs(Number(digit) - Number(flat[i] ?? 0)),
            0,
          );
          expect(off, `level ${level}: ${got} vs ${want}`).toBeLessThanOrEqual(2);
        } else {
          expect(got, `level ${level}`).toBe(want);
        }
      }
      // And the order is one the planner accepts point by point.
      const talentFile = JSON.parse(
        readFileSync(path.join(REPO, 'data/builds', BUILD, 'talents', `${row.classSlug}.json`), 'utf8'),
      );
      expect(validateOrder(indexTalents(talentFile), order)).toEqual([]);
    },
  );
});
