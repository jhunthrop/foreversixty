// web/src/lib/guides/talent-cells.test.ts
import { readFileSync } from 'node:fs';
import { fileURLToPath } from 'node:url';
import { parse as parseYaml } from 'yaml';
import { describe, expect, it } from 'vitest';
import activeBuild from '../../data/active-build.json';
import type { TalentTree } from '../../components/ui/TalentTrees.svelte';
import { decodeBuildTrees, type DecodedTalentView, type DecodedTreeView } from './build-tree';
import { buildCodeDigits } from './leveling-builds';
import { talentTreesOf, treesAtStep } from './talent-cells';

function talent(id: number, tier: number, column: number, rank: number, maxRank: number): DecodedTalentView {
  return {
    id,
    name: `Talent ${id}`,
    icon: `icon_${id}`,
    max_rank: maxRank,
    tier,
    column,
    prereq_talent_id: null,
    prereq_rank: null,
    ranks: [],
    spell_id: id * 10,
    rank,
  };
}

const TREES: DecodedTreeView[] = [
  { name: 'Arcane', position: 0, points: 3, talents: [talent(1, 0, 0, 2, 2), talent(2, 0, 1, 1, 5)] },
  { name: 'Fire', position: 1, points: 0, talents: [talent(3, 0, 0, 0, 3)] },
];

describe('talentTreesOf', () => {
  it('maps each decoded talent onto its cell: place, icon, rank and max rank', () => {
    const [arcane] = talentTreesOf(TREES);
    expect(arcane).toEqual({
      name: 'Arcane',
      points: 3,
      talents: [
        { tier: 0, column: 0, name: 'Talent 1', icon: 'icon_1', rank: 2, maxRank: 2 },
        { tier: 0, column: 1, name: 'Talent 2', icon: 'icon_2', rank: 1, maxRank: 5 },
      ],
    });
  });

  it('sets the tooltip key only for a talent the bundle carries, at rank 1 when unlearned', () => {
    const bundle = { talents: { '1': { ranks: [] }, '3': { ranks: [] } } } as never;
    const [arcane, fire] = talentTreesOf(TREES, bundle);
    expect(arcane!.talents.map((cell) => cell.tipKey)).toEqual(['talent:1:2', undefined]);
    expect(fire!.talents[0]!.tipKey).toBe('talent:3:1');
  });
});

describe('treesAtStep', () => {
  const layout: TalentTree[] = talentTreesOf(TREES);

  it('reads one digit per talent in each tree’s order and recomputes the points', () => {
    const [arcane, fire] = treesAtStep(layout, '13-2');
    expect(arcane!.talents.map((cell) => cell.rank)).toEqual([1, 3]);
    expect(arcane!.points).toBe(4);
    expect(fire!.talents[0]!.rank).toBe(2);
    expect(fire!.points).toBe(2);
  });

  it('reads a missing tree or digit as nothing spent', () => {
    const [arcane, fire] = treesAtStep(layout, '1');
    expect(arcane!.talents.map((cell) => cell.rank)).toEqual([1, 0]);
    expect(fire!.points).toBe(0);
  });

  it('keeps the layout it was handed (a new tree, never a mutation)', () => {
    treesAtStep(layout, '00-0');
    expect(layout[0]!.points).toBe(3);
    expect(layout[0]!.talents[0]!.rank).toBe(2);
  });

  it('gives back the build itself from the build’s own digits (a real guide)', () => {
    const path = fileURLToPath(new URL('../../content/guides/mage/frost.md', import.meta.url));
    const code = (parseYaml(readFileSync(path, 'utf8').split('---')[1]!) as { build: string }).build;
    const decoded = decodeBuildTrees(activeBuild.build, code);
    const digits = buildCodeDigits(code);
    if (decoded === undefined || digits === undefined) throw new Error('the Frost build does not decode');
    const trees = talentTreesOf(decoded);
    expect(treesAtStep(trees, digits.join('-'))).toEqual(trees);
  });
});
