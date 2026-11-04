// web/src/lib/guides/leveling-talents.test.ts
import { readdirSync, readFileSync } from 'node:fs';
import { join } from 'node:path';
import activeBuild from '../../data/active-build.json';
import { repoRoot } from '../data/build-files';
import { describe, expect, it } from 'vitest';
import { loadBisFile } from '../bis/load';
import { decodeBandTalents, type DecodedTalentView, type DecodedTreeView } from './build-tree';
import { guideLadderOrder, guideLevelingBuilds } from './leveling-builds';
import { levelingChipRows } from './leveling-talents';

function talent(id: number, name: string, tier: number, column: number, rank: number): DecodedTalentView {
  return {
    id,
    name,
    icon: `icon_${id}`,
    max_rank: 5,
    tier,
    column,
    prereq_talent_id: null,
    prereq_rank: null,
    ranks: [],
    spell_id: id,
    rank,
  };
}

function tree(position: number, ...talents: DecodedTalentView[]): DecodedTreeView {
  return { name: `T${position}`, position, talents, points: talents.reduce((sum, t) => sum + t.rank, 0) };
}

const BUILD = activeBuild.build;

describe('levelingChipRows', () => {
  const band20 = [
    tree(0, talent(1, 'Fire One', 0, 0, 1)),
    tree(1, talent(2, 'Frost Two', 1, 0, 3), talent(3, 'Frost One', 0, 1, 2)),
  ];
  const band30 = [
    tree(0, talent(1, 'Fire One', 0, 0, 1)),
    tree(
      1,
      talent(2, 'Frost Two', 1, 0, 5),
      talent(3, 'Frost One', 0, 1, 2),
      talent(4, 'Frost Three', 2, 0, 1),
    ),
  ];

  // The order the points are spent in: Frost Three's point lands before Frost Two's last.
  const spend = [3, 3, 2, 2, 2, 4, 2, 2, 1];

  it('lists a band’s talents in the order their first point is spent', () => {
    const [first] = levelingChipRows([band20], spend);
    expect(first!.map((chip) => chip.name)).toEqual(['Frost One', 'Frost Two', 'Fire One']);
  });

  it('keeps every band in that one order, so a talent never moves between bands', () => {
    const [, second] = levelingChipRows([band20, band30], spend);
    expect(second!.map((chip) => chip.name)).toEqual(['Frost One', 'Frost Two', 'Frost Three', 'Fire One']);
  });

  it('puts a talent the order does not name after the ones it does', () => {
    const [first] = levelingChipRows([band20], [2]);
    expect(first!.map((chip) => chip.name)).toEqual(['Frost Two', 'Fire One', 'Frost One']);
  });

  it('carries the icon, the rank and the maximum rank', () => {
    expect(levelingChipRows([band20], spend)[0]![1]).toMatchObject({ icon: 'icon_2', rank: 3, maxRank: 5 });
  });

  it('marks nothing new on the first band, and a rise or a new talent after it', () => {
    const [first, second] = levelingChipRows([band20, band30], spend);
    expect(first!.some((chip) => chip.isNew)).toBe(false);
    expect(second!.filter((chip) => chip.isNew).map((chip) => chip.name)).toEqual([
      'Frost Two',
      'Frost Three',
    ]);
  });

  it('gives a band that did not decode an empty list', () => {
    expect(levelingChipRows([band20, undefined, band30], spend)[1]).toEqual([]);
  });
});

describe('a real band talent string', () => {
  it('decodes into chips with a name, an icon and a rank for every lit talent', () => {
    const file = loadBisFile('mage-frost', BUILD);
    const band = file?.bands.find((b) => b.band === 40 && b.faction === 'alliance');
    if (band === undefined) throw new Error('mage-frost band 40 missing');
    const trees = decodeBandTalents(BUILD, 'mage', band.talents);
    if (trees === undefined) throw new Error('mage talent file missing');
    const chips = levelingChipRows([trees], [])[0]!;
    expect(chips.length).toBeGreaterThan(5);
    expect(chips.reduce((sum, chip) => sum + chip.rank, 0)).toBe(band.talent_points);
    for (const chip of chips) {
      expect(chip.name).not.toBe('');
      expect(chip.icon).not.toBe('');
      expect(chip.rank).toBeGreaterThan(0);
      expect(chip.rank).toBeLessThanOrEqual(chip.maxRank);
    }
  });
});

describe('the Leveling strip of every guide', () => {
  const guides = guideBuilds();

  it('reads every guide’s build', () => {
    expect(guides.length).toBeGreaterThan(20);
  });

  // Sim-accuracy round 15 wave 6: the Shadow ladder takes Shadow Focus (17-21) before Mind
  // Flay (28) and Improved Mind Flay (29-30), and band 30 holds all three
  // (data/curated/ladder-first.json priest-shadow).
  it('shows the Shadow priest’s Shadow Focus before Mind Flay and Improved Mind Flay at band 30', () => {
    const shadow = guides.find((guide) => guide.spec === 'priest-shadow')!;
    const names = stripRows(shadow).find((row) => row.band === 30)!.names;
    expect(names.indexOf('Mind Flay')).toBeGreaterThan(-1);
    expect(names.indexOf('Shadow Focus')).toBeLessThan(names.indexOf('Mind Flay'));
    expect(names.indexOf('Mind Flay')).toBeLessThan(names.indexOf('Improved Mind Flay'));
  });

  // Sim-accuracy round 16: the strip follows the ladder's own order, which the sim's ladder
  // and the addon spend (Wand Specialization at 10 and 11, Spirit Tap from 12), not the
  // planner's reconstruction of the level 60 build (Blackout first).
  it('starts the Shadow priest’s first range on Wand Specialization, then Spirit Tap', () => {
    const shadow = guides.find((guide) => guide.spec === 'priest-shadow')!;
    const names = stripRows(shadow).find((row) => row.band === 20)!.names;
    expect(names.slice(0, 2)).toEqual(['Wand Specialization', 'Spirit Tap']);
  });

  it('lists every band in the order the ladder spends the guide’s points', () => {
    for (const guide of guides) {
      const rows = stripRows(guide);
      const order = guideLadderOrder(BUILD, guide.classSlug, guide.spec, guide.treeIndex, guide.code);
      const firstPoint = (name: string): number =>
        order.findIndex((id) =>
          rows.at(-1)!.trees!.some((tree) => tree.talents.some((t) => t.id === id && t.name === name)),
        );
      for (const row of rows) {
        const points = row.names.map(firstPoint);
        expect(points, `${guide.spec} band ${row.band}`).toEqual([...points].sort((a, b) => a - b));
      }
    }
  });
});

interface GuideBuild {
  spec: string;
  classSlug: string;
  code: string;
  treeIndex: number;
}

function guideBuilds(): GuideBuild[] {
  const root = join(repoRoot(), 'web/src/content/guides');
  const specs = JSON.parse(readFileSync(join(repoRoot(), 'data/curated/specs.json'), 'utf8')) as unknown;
  const rows = (Array.isArray(specs) ? specs : (specs as { specs: unknown[] }).specs) as {
    class_slug: string;
    spec_slug: string;
    tree_index: number;
  }[];
  return readdirSync(root, { withFileTypes: true })
    .filter((entry) => entry.isDirectory())
    .flatMap((dir) =>
      readdirSync(join(root, dir.name))
        .filter((file) => file.endsWith('.md') && !file.startsWith('_'))
        .map((file): GuideBuild | undefined => {
          const text = readFileSync(join(root, dir.name, file), 'utf8');
          const code = /^build: '([^']+)'/m.exec(text)?.[1];
          const specSlug = file.replace(/\.md$/, '');
          const row = rows.find((r) => r.class_slug === dir.name && r.spec_slug === specSlug);
          if (code === undefined || row === undefined) return undefined;
          return { spec: `${dir.name}-${specSlug}`, classSlug: dir.name, code, treeIndex: row.tree_index };
        })
        .filter((guide): guide is GuideBuild => guide !== undefined),
    );
}

function stripRows(guide: GuideBuild) {
  const bands = guideLevelingBuilds(BUILD, guide.classSlug, guide.spec, guide.treeIndex, guide.code);
  const decoded = bands.map((band) => decodeBandTalents(BUILD, guide.classSlug, band.talents));
  const order = guideLadderOrder(BUILD, guide.classSlug, guide.spec, guide.treeIndex, guide.code);
  const chips = levelingChipRows(decoded, order);
  return bands.map((band, i) => ({
    band: band.band,
    talents: band.talents,
    trees: decoded[i],
    names: chips[i]!.map((chip) => chip.name),
  }));
}
