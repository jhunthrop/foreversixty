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
  // Sim-accuracy round 18: the pre-pull Frostbolt runs only without Heating Up, so the raid
  // build (which takes it) never casts it; the level 60 build does, with Improved Frostbolt
  // 5/5, a 2.5 s cast (3 s base, client 10181, less 0.5 s), which the APL starts at -2.5 s and
  // the prose names. The card read the raid build's 3 s.
  it('times Arcane’s pre-pull Frostbolt by the build that casts it', () => {
    const { leveling, raid } = guideTrees('mage/arcane.md');
    expect(rankOf(raid!, 'Heating Up')).toBeGreaterThan(0);
    expect(rankOf(leveling, 'Heating Up')).toBe(0);
    expect(rankOf(leveling, 'Improved Frostbolt')).toBe(5);
    const trees = describedBuildTrees(leveling, raid);
    const frostbolt = guideRotation(BUILD, 'mage', 'mage-arcane', 'dps', trees, leveling)?.opener.find(
      (item) => item.name === 'Frostbolt',
    );
    expect(frostbolt?.clause).toBe('2.5 s before the pull, without Heating Up');
  });

  it('times a pre-pull cast both builds run by the described build', () => {
    const { leveling, raid } = guideTrees('mage/arcane.md');
    const pyroblast = guideRotation(BUILD, 'mage', 'mage-arcane', 'dps', raid!, leveling)?.opener.find(
      (item) => item.name === 'Pyroblast',
    );
    expect(pyroblast?.clause).toBe('6 s before the pull, with Heating Up');
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
