import { describe, expect, it } from 'vitest';
import { codeTrees, guideBuildRows, ranksTrees, treesSplit, type GuideBuild } from './guide-builds';

const SURVIVAL = 'FS1:1.60.1.70790:hunter:night-elf:5/005005201/500230230050222151:';
const BM = 'FS1:1.60.1.70790:hunter:night-elf:5420001505001251/00524510012/0:';
const BM_RAID = 'FS1:1.60.1.70790:hunter:night-elf:5420001505001251/00523500014/0:';
const BUILDS: GuideBuild[] = [
  { classSlug: 'hunter', specSlug: 'survival', kind: 'leveling', code: SURVIVAL },
  { classSlug: 'hunter', specSlug: 'beast-mastery', kind: 'raid', code: BM_RAID },
  { classSlug: 'hunter', specSlug: 'beast-mastery', kind: 'leveling', code: BM },
  { classSlug: 'mage', specSlug: 'fire', kind: 'leveling', code: 'FS1:1:mage:human:0/5/0:' },
  { classSlug: 'hunter', specSlug: 'marksmanship', kind: 'leveling', code: 'not a code' },
];

describe('guide builds', () => {
  it('reads a code tree segment with each tree trailing zeros trimmed', () => {
    expect(codeTrees(SURVIVAL)).toBe('5/005005201/500230230050222151');
    expect(codeTrees('FS1:1:hunter:human:50/0/0100:')).toBe('5//01');
    expect(codeTrees('nope')).toBeNull();
  });

  it('puts a build ranks in the same form, so a loaded guide build compares equal', () => {
    expect(ranksTrees([[5, 0], [0], [0, 1, 0, 0]])).toBe('5//01');
  });

  it('sums each tree into the split', () => {
    expect(treesSplit('5/005005201/500230230050222151')).toEqual([5, 13, 33]);
  });

  it('lists the class builds by spec in tree order, the levelling build first, with a guide load link', () => {
    const rows = guideBuildRows(BUILDS, 'hunter');
    expect(rows.map((row) => row.key)).toEqual([
      'hunter-beast-mastery-leveling',
      'hunter-beast-mastery-raid',
      'hunter-survival-leveling',
    ]);
    expect(rows[0]).toMatchObject({ name: 'Beast Mastery', label: 'Level 60 build', split: [31, 20, 0] });
    expect(rows[1].label).toBe('Raid build');
    expect(rows[2].href).toBe(`/planner?code=${encodeURIComponent(SURVIVAL)}&spec=hunter-survival`);
    expect(rows[2].icon).toBeTruthy();
  });
});
