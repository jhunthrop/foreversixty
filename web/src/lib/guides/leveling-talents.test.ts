// web/src/lib/guides/leveling-talents.test.ts
import activeBuild from '../../data/active-build.json';
import { describe, expect, it } from 'vitest';
import { loadBisFile } from '../bis/load';
import { decodeBandTalents, type DecodedTalentView, type DecodedTreeView } from './build-tree';
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

  it('lists a band’s first talents with the main tree first, each from its first row', () => {
    const [first] = levelingChipRows([band20]);
    expect(first!.map((chip) => chip.name)).toEqual(['Frost One', 'Frost Two', 'Fire One']);
  });

  it('keeps the earlier band’s order and puts what a band adds after it', () => {
    const [, second] = levelingChipRows([band20, band30]);
    expect(second!.map((chip) => chip.name)).toEqual(['Frost One', 'Frost Two', 'Fire One', 'Frost Three']);
  });

  it('carries the icon, the rank and the maximum rank', () => {
    expect(levelingChipRows([band20])[0]![1]).toMatchObject({ icon: 'icon_2', rank: 3, maxRank: 5 });
  });

  it('marks nothing new on the first band, and a rise or a new talent after it', () => {
    const [first, second] = levelingChipRows([band20, band30]);
    expect(first!.some((chip) => chip.isNew)).toBe(false);
    expect(second!.filter((chip) => chip.isNew).map((chip) => chip.name)).toEqual([
      'Frost Two',
      'Frost Three',
    ]);
  });

  it('gives a band that did not decode an empty list', () => {
    expect(levelingChipRows([band20, undefined, band30])[1]).toEqual([]);
  });
});

describe('a real band talent string', () => {
  it('decodes into chips with a name, an icon and a rank for every lit talent', () => {
    const file = loadBisFile('mage-frost', BUILD);
    const band = file?.bands.find((b) => b.band === 40 && b.faction === 'alliance');
    if (band === undefined) throw new Error('mage-frost band 40 missing');
    const trees = decodeBandTalents(BUILD, 'mage', band.talents);
    if (trees === undefined) throw new Error('mage talent file missing');
    const chips = levelingChipRows([trees])[0]!;
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
