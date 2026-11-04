// web/src/lib/guides/guide-rotation.test.ts
// The rotation cards read the build they describe (the raid build, else the level 60 build)
// and no other: its cooldown cuts, its cast times, its pre-pull timings.
import activeBuild from '../../data/active-build.json';
import { readFileSync } from 'node:fs';
import path from 'node:path';
import { fileURLToPath } from 'node:url';
import { describe, expect, it } from 'vitest';
import { decodeBuildTrees, type DecodedTreeView } from './build-tree';
import { describedBuildTrees, guideRotation } from './guide-rotation';
import { litTalentsOf, talentCastTimeCuts } from './rotation-lists';

const BUILD = activeBuild.build;
const GUIDES_DIR = path.resolve(path.dirname(fileURLToPath(import.meta.url)), '../../content/guides');

function guideTrees(file: string): { leveling: DecodedTreeView[]; raid: DecodedTreeView[] | undefined } {
  const raw = readFileSync(path.join(GUIDES_DIR, file), 'utf8');
  const code = (key: string): string | undefined => new RegExp(`^${key}: '([^']+)'`, 'm').exec(raw)?.[1];
  const leveling = decodeBuildTrees(BUILD, code('build') ?? '') ?? [];
  const raidCode = code('raidBuild');
  const raid =
    raidCode === undefined || raidCode === code('build') ? undefined : decodeBuildTrees(BUILD, raidCode);
  return { leveling, raid };
}

function rankOf(trees: readonly DecodedTreeView[], name: string): number {
  return trees.flatMap((tree) => tree.talents).find((talent) => talent.name === name)?.rank ?? 0;
}

describe('guideRotation', () => {
  it('times Arcane’s pre-pull Frostbolt by the raid build’s own cast time', () => {
    const { leveling, raid } = guideTrees('mage/arcane.md');
    const trees = describedBuildTrees(leveling, raid);
    const frostbolt = guideRotation(BUILD, 'mage', 'mage-arcane', 'dps', trees)?.opener.find(
      (item) => item.name === 'Frostbolt',
    );
    // 3 s base, less half a second only when the build takes Improved Frostbolt.
    const want = rankOf(trees, 'Improved Frostbolt') > 0 ? '2.5 s before the pull' : '3 s before the pull';
    expect(frostbolt?.clause).toBe(want);
  });

  it('gives Marksmanship’s Rapid Fire the raid build’s cooldown, not the level 60 build’s Rapid Killing', () => {
    const { leveling, raid } = guideTrees('hunter/marksmanship.md');
    const trees = describedBuildTrees(leveling, raid);
    const rapidFire = guideRotation(BUILD, 'hunter', 'hunter-marksmanship', 'dps', trees)?.cooldowns.find(
      (item) => item.name === 'Rapid Fire',
    );
    const cut = rankOf(trees, 'Rapid Killing');
    expect(rapidFire?.clause.startsWith(`${5 - cut} min cooldown`)).toBe(true);
  });
});

describe('describedBuildTrees', () => {
  it('is the raid build when the guide has its own, else the level 60 build', () => {
    const leveling = [{ name: 'L', position: 0, talents: [], points: 0 }];
    const raid = [{ name: 'R', position: 0, talents: [], points: 0 }];
    expect(describedBuildTrees(leveling, raid)).toBe(raid);
    expect(describedBuildTrees(leveling, undefined)).toBe(leveling);
  });
});

describe('talentCastTimeCuts', () => {
  it('reads lasting cuts and not a proc’s', () => {
    const cuts = talentCastTimeCuts(
      BUILD,
      'mage',
      litTalentsOf([
        {
          talents: [
            { name: 'Improved Frostbolt', rank: 5 },
            { name: 'Heating Up', rank: 1 },
          ],
        },
      ]),
    );
    expect(cuts.get('Frostbolt')).toBe(500);
    expect(cuts.has('Pyroblast')).toBe(false);
  });
});
