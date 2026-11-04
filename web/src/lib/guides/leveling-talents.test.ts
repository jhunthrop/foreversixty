// web/src/lib/guides/leveling-talents.test.ts
import { readdirSync, readFileSync } from 'node:fs';
import { join } from 'node:path';
import activeBuild from '../../data/active-build.json';
import { repoRoot } from '../data/build-files';
import { describe, expect, it } from 'vitest';
import { loadBisFile } from '../bis/load';
import { decodeBandTalents, type DecodedTalentView, type DecodedTreeView } from './build-tree';
import { guideLadderOrder, guideLevelingBuilds } from './leveling-builds';
import { levelingChipRows, levelRuns, levelRunsText } from './leveling-talents';

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

  // The order the points are spent in, one a level from 10: Frost Three's point (15) lands
  // before Frost Two's last (16, 17).
  const spend = [3, 3, 2, 2, 2, 4, 2, 2, 1];
  // Two ranges in miniature: levels 10 to 13 (band20's build), then 14 to 18 (band30's).
  const first = { trees: band20, from: 10, to: 13 };
  const second = { trees: band30, from: 14, to: 18 };

  it('lists a band’s talents in the order their first point is spent', () => {
    const [chips] = levelingChipRows([first], spend);
    expect(chips!.map((chip) => chip.name)).toEqual(['Frost One', 'Frost Two', 'Fire One']);
  });

  it('lists what a range carries first, then what it adds, in the order the points come', () => {
    const [, chips] = levelingChipRows([first, second], spend);
    expect(chips!.map((chip) => chip.name)).toEqual(['Frost One', 'Frost Two', 'Frost Three', 'Fire One']);
    // Frost Three's point at 14 comes before Frost Two's next at 15, so it leads the added ones.
    const later = [3, 3, 2, 2, 4, 2, 2, 2, 1];
    const [, reordered] = levelingChipRows([first, second], later);
    expect(reordered!.map((chip) => chip.name)).toEqual([
      'Frost One',
      'Frost Three',
      'Frost Two',
      'Fire One',
    ]);
  });

  it('puts a talent the order does not name after the ones it does', () => {
    const [chips] = levelingChipRows([first], [2]);
    expect(chips!.map((chip) => chip.name)).toEqual(['Frost Two', 'Fire One', 'Frost One']);
  });

  it('carries the icon, the rank and the maximum rank', () => {
    expect(levelingChipRows([first], spend)[0]![1]).toMatchObject({ icon: 'icon_2', rank: 3, maxRank: 5 });
  });

  it('marks nothing new on the first band, and a rise or a new talent after it', () => {
    const [one, two] = levelingChipRows([first, second], spend);
    expect(one!.some((chip) => chip.isNew)).toBe(false);
    expect(two!.filter((chip) => chip.isNew).map((chip) => chip.name)).toEqual(['Frost Two', 'Frost Three']);
  });

  // Sim-accuracy round 18: a range's list is what a player holds by its last level, and each
  // talent says the levels its points come at inside the range, so a level 16 reader of the
  // second range sees Frost Three (15) and the Frost Two point at 16.
  it('names the levels each talent’s points come at inside the range', () => {
    const [one, two] = levelingChipRows([first, second], spend);
    const levels = (chips: typeof one) => Object.fromEntries(chips!.map((chip) => [chip.name, chip.levels]));
    expect(levels(one)).toEqual({ 'Frost One': [10, 11], 'Frost Two': [12, 13], 'Fire One': [] });
    expect(levels(two)).toEqual({
      'Frost One': [],
      'Frost Two': [14, 16, 17],
      'Frost Three': [15],
      'Fire One': [18],
    });
  });

  it('gives a band that did not decode an empty list', () => {
    expect(levelingChipRows([first, { trees: undefined, from: 14, to: 18 }, second], spend)[1]).toEqual([]);
  });
});

describe('levelRuns', () => {
  it('folds levels into runs', () => {
    expect(levelRuns([30, 31, 32, 35, 37, 38])).toEqual([
      [30, 32],
      [35, 35],
      [37, 38],
    ]);
    expect(levelRuns([])).toEqual([]);
  });

  it('writes the runs a reader scans', () => {
    expect(levelRunsText([30, 31, 32, 35])).toBe('30-32, 35');
    expect(levelRunsText([44])).toBe('44');
  });
});

describe('a real band talent string', () => {
  it('decodes into chips with a name, an icon and a rank for every lit talent', () => {
    const file = loadBisFile('mage-frost', BUILD);
    const band = file?.bands.find((b) => b.band === 40 && b.faction === 'alliance');
    if (band === undefined) throw new Error('mage-frost band 40 missing');
    const trees = decodeBandTalents(BUILD, 'mage', band.talents);
    if (trees === undefined) throw new Error('mage talent file missing');
    const chips = levelingChipRows([{ trees, from: 10, to: 49 }], [])[0]!;
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

  // Sim-accuracy round 18: "Level 30 to 39" showed the level 30 build, so a level 36 reader
  // never saw the points taken at 31 to 36. Each range now lists what a player holds by its
  // last level, and every point of the ladder is named once, at its own level.
  it('lists each range as held at its last level, every point named once at its level', () => {
    for (const guide of guides) {
      const rows = stripRows(guide);
      const order = guideLadderOrder(BUILD, guide.classSlug, guide.spec, guide.treeIndex, guide.code);
      for (const row of rows) {
        const held = order.slice(0, row.lastLevel - 9);
        for (const chip of row.chips) {
          expect(chip.rank, `${guide.spec} ${row.band} ${chip.name}`).toBe(
            held.filter((id) => id === chip.id).length,
          );
        }
      }
      const named = rows.flatMap((row) => row.chips.flatMap((chip) => chip.levels)).sort((a, b) => a - b);
      expect(named, guide.spec).toEqual(order.map((_, point) => 10 + point));
    }
  });

  // A row reads left to right in the order a player takes the points: the talents it carries
  // from earlier ranges first, by their first point, then the ones it adds, by the first level
  // it adds them at (round 18 wow-player: Weaponmaster rising at 41 sat ahead of Mortal Strike
  // at 40).
  it('lists every band in the order the ladder spends the guide’s points', () => {
    for (const guide of guides) {
      const rows = stripRows(guide);
      const order = guideLadderOrder(BUILD, guide.classSlug, guide.spec, guide.treeIndex, guide.code);
      for (const row of rows) {
        const carried = row.chips.filter((chip) => chip.levels.length === 0);
        const added = row.chips.filter((chip) => chip.levels.length > 0);
        expect(row.chips, `${guide.spec} band ${row.band}`).toEqual([...carried, ...added]);
        const firsts = carried.map((chip) => order.indexOf(chip.id));
        expect(firsts, `${guide.spec} band ${row.band}`).toEqual([...firsts].sort((a, b) => a - b));
        const levels = added.map((chip) => chip.levels[0]!);
        expect(levels, `${guide.spec} band ${row.band}`).toEqual([...levels].sort((a, b) => a - b));
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
  const decoded = bands.map((band) => decodeBandTalents(BUILD, guide.classSlug, band.lastTalents));
  const order = guideLadderOrder(BUILD, guide.classSlug, guide.spec, guide.treeIndex, guide.code);
  const chips = levelingChipRows(
    bands.map((band, i) => ({ trees: decoded[i], from: band.pointsFrom, to: band.lastLevel })),
    order,
  );
  return bands.map((band, i) => ({
    band: band.band,
    lastLevel: band.lastLevel,
    talents: band.lastTalents,
    trees: decoded[i],
    chips: chips[i]!,
    names: chips[i]!.map((chip) => chip.name),
  }));
}
