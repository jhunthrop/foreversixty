// web/src/lib/guides/talent-cells.test.ts
import { readFileSync } from 'node:fs';
import { fileURLToPath } from 'node:url';
import { parse as parseYaml } from 'yaml';
import { describe, expect, it } from 'vitest';
import activeBuild from '../../data/active-build.json';
import { treesAtTalents } from '../ui/talent-tree';
import { decodeBuildTrees } from './build-tree';
import { buildCodeDigits } from './leveling-builds';
import { guideTalentTrees } from './talent-cells';

function frostCode(): string {
  const path = fileURLToPath(new URL('../../content/guides/mage/frost.md', import.meta.url));
  return (parseYaml(readFileSync(path, 'utf8').split('---')[1]!) as { build: string }).build;
}

describe('guideTalentTrees', () => {
  const decoded = decodeBuildTrees(activeBuild.build, frostCode());
  if (decoded === undefined) throw new Error('the Frost build does not decode');
  const trees = guideTalentTrees(activeBuild.build, 'mage', decoded);

  it('draws every tree of the class at the build’s ranks, with its points', () => {
    expect(trees.map((tree) => tree.name)).toEqual(decoded.map((tree) => tree.name));
    expect(trees.map((tree) => tree.points)).toEqual(decoded.map((tree) => tree.points));
  });

  it('names each tooltip by key only: the guide’s host serves the text', () => {
    const cells = trees.flatMap((tree) => tree.cells);
    expect(cells.every((cell) => cell.ranks === undefined)).toBe(true);
  });

  it('greys an untaken talent and rings a maxed one gold', () => {
    const cells = trees.flatMap((tree) => tree.cells);
    expect(cells.filter((cell) => cell.rank === 0).every((cell) => cell.state === 'lock')).toBe(true);
    expect(cells.filter((cell) => cell.rank === cell.max).every((cell) => cell.state === 'max')).toBe(true);
  });

  it('carries the arrows with what they need, so a level step can recompute them', () => {
    const connectors = trees.flatMap((tree) => tree.connectors);
    expect(connectors.length).toBeGreaterThan(0);
    expect(connectors.every((connector) => connector.needs > 0)).toBe(true);
  });

  it('gives back the build itself from the build’s own digits (the leveling ladder’s strings)', () => {
    const digits = buildCodeDigits(frostCode());
    if (digits === undefined) throw new Error('the Frost build has no digits');
    expect(treesAtTalents(trees, digits.join('-'))).toEqual(trees);
  });

  it('is empty for a build with no trees', () => {
    expect(guideTalentTrees(activeBuild.build, 'mage', [])).toEqual([]);
  });
});
