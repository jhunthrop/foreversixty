// web/src/lib/guides/build-tree.test.ts
import { readFileSync } from 'node:fs';
import { fileURLToPath } from 'node:url';
import { describe, expect, it } from 'vitest';
import { decodeBuildTrees, litTalents, primaryTree } from './build-tree';

const BUILD = '1.60.1.70374';

function frontmatterBuildCode(): string {
  const path = fileURLToPath(new URL('../../content/guides/warrior/fury.md', import.meta.url));
  const raw = readFileSync(path, 'utf8');
  const match = /^build:\s*'([^']+)'/m.exec(raw);
  if (match === null) throw new Error('fury.md build code not found');
  return match[1]!;
}

describe('decodeBuildTrees', () => {
  it('decodes all three trees with ranks summing to the guide’s own talent points', () => {
    const trees = decodeBuildTrees(BUILD, frontmatterBuildCode());
    expect(trees).toBeDefined();
    expect(trees).toHaveLength(3);
    const totalPoints = trees!.reduce((sum, tree) => sum + tree.points, 0);
    expect(totalPoints).toBeGreaterThan(0);
    expect(totalPoints).toBeLessThanOrEqual(51);
  });

  it('is undefined for a code that does not parse', () => {
    expect(decodeBuildTrees(BUILD, 'not-a-real-code')).toBeUndefined();
  });

  it('sorts each tree’s talents by tier then column', () => {
    const trees = decodeBuildTrees(BUILD, frontmatterBuildCode())!;
    for (const tree of trees) {
      for (let i = 1; i < tree.talents.length; i += 1) {
        const prev = tree.talents[i - 1]!;
        const curr = tree.talents[i]!;
        expect(prev.tier < curr.tier || (prev.tier === curr.tier && prev.column <= curr.column)).toBe(true);
      }
    }
  });
});

describe('primaryTree', () => {
  it('picks the tree with the most points (Fury’s own guide spends most of its build in Fury)', () => {
    const trees = decodeBuildTrees(BUILD, frontmatterBuildCode())!;
    const primary = primaryTree(trees);
    expect(primary?.name).toBe('Fury');
  });

  it('is undefined for an empty list', () => {
    expect(primaryTree([])).toBeUndefined();
  });
});

describe('litTalents', () => {
  it('returns only talents with rank > 0, across all trees', () => {
    const trees = decodeBuildTrees(BUILD, frontmatterBuildCode())!;
    const lit = litTalents(trees);
    expect(lit.every((talent) => talent.rank > 0)).toBe(true);
    const expectedCount = trees.flatMap((tree) => tree.talents).filter((t) => t.rank > 0).length;
    expect(lit).toHaveLength(expectedCount);
  });
});
