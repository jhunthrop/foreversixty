// web/src/lib/sim/talent-loadouts.test.ts
import { describe, expect, it } from 'vitest';
import talentsFixture from '../../fixtures/planner/talents/warrior.json';
import { indexTalents } from '../planner/rules';
import type { TalentFile } from '../planner/types';
import { bulkCopy } from './copy';
import { poolQualityCopy } from './pool-quality-copy';
import type { BisFile } from '../bis/types';
import {
  customLoadouts,
  guideLoadout,
  hasUnaddedBuild,
  loadoutTrees,
  savedBuildsMessage,
} from './talent-loadouts';
import type { TalentLoadout } from './types';

describe('guideLoadout', () => {
  const band = (level: number, faction: string, talents: string) => ({
    band: level,
    faction,
    preset: 'raid',
    talents,
  });
  const file = {
    spec: 'warrior-fury',
    bands: [band(50, 'horde', '1-2-3'), band(60, 'horde', '343-051-2'), band(60, 'alliance', 'a-b-c')],
  } as unknown as BisFile;

  it('is the band’s own talents at the character’s level band and faction', () => {
    expect(guideLoadout(file, 60, 'orc')).toEqual({
      name: bulkCopy.talentsGuide(60, '10/6/2'),
      talents: '343-051-2',
    });
    expect(guideLoadout(file, 57, 'troll')).toEqual({
      name: bulkCopy.talentsGuide(50, '1/2/3'),
      talents: '1-2-3',
    });
  });

  it('is null for a race with no faction or a band the file lacks', () => {
    expect(guideLoadout(file, 60, 'murloc')).toBeNull();
    expect(guideLoadout(file, 40, 'orc')).toBeNull();
  });
});

const index = indexTalents(talentsFixture as unknown as TalentFile);

function loadout(name: string): TalentLoadout {
  return { name, talents: `talents-for-${name}` };
}

describe('customLoadouts', () => {
  const own = loadout('Your current build');
  const saved = [loadout('Raid build')];
  const exported = [loadout('Addon loadout 1')];

  it('is empty when every picked loadout is already named by one of the known lists', () => {
    const picked = [own, saved[0], exported[0]];
    expect(customLoadouts(picked, own, [saved, exported])).toEqual([]);
  });

  /**
   * dps D31/E2: ADD A BUILD ticks a pasted or hand-built loadout into `picked`
   * (`store.addLoadout`) immediately, but nothing rendered a row for it once it was neither
   * the character's own build, a signed-in player's saved build, nor an addon-exported
   * loadout -- so a signed-out player pasting a second build to compare against their own
   * had no way to see it was added, or to run it.
   */
  it('names a picked loadout none of the known lists claims', () => {
    const custom = loadout('Build 1');
    const picked = [own, custom];
    expect(customLoadouts(picked, own, [saved, exported])).toEqual([custom]);
  });

  it('treats a null own build the same as one that claims nothing', () => {
    const custom = loadout('Build 1');
    expect(customLoadouts([custom], null, [saved, exported])).toEqual([custom]);
  });

  it('preserves picked order across more than one custom loadout', () => {
    const first = loadout('Build 1');
    const second = loadout('Build 2');
    expect(customLoadouts([own, first, second], own, [[], []])).toEqual([first, second]);
  });
});

describe('hasUnaddedBuild', () => {
  const VALID_CODE = 'FS1:1:warrior:orc:1/0/0:';

  /**
   * newcomer round 4, review.md:83-110: two buttons both read "Add a build" -- the outer
   * toggle (sim-loadout-add) and the inner accept (sim-loadout-accept) -- and a player who
   * pastes into IMPORT FROM ADDON, then clicks RUN without ever finding the second button,
   * gets a silent one-row table with no explanation. `hasUnaddedBuild` is the trap's own
   * precondition: a decodable build sitting in the editor's `customCode` that `addCustom`
   * has not yet turned into a picked loadout.
   */
  it('is true once the editor holds a build addCustom could actually accept', () => {
    expect(hasUnaddedBuild(VALID_CODE, index)).toBe(true);
  });

  it('is false with nothing pasted', () => {
    expect(hasUnaddedBuild('', index)).toBe(false);
  });

  it('is false for a string that does not decode', () => {
    expect(hasUnaddedBuild('not a build code', index)).toBe(false);
  });

  it('is false while talents have not loaded yet -- addCustom itself would no-op', () => {
    expect(hasUnaddedBuild(VALID_CODE, null)).toBe(false);
  });
});

describe('savedBuildsMessage', () => {
  it('reads as guidance, not an error, when the failure was signing out', () => {
    expect(savedBuildsMessage(true)).toBe(poolQualityCopy.talentsSavedSignedOut);
    expect(savedBuildsMessage(true)).not.toBe(bulkCopy.talentsSavedUnavailable);
  });

  it('keeps the original message for a real read failure', () => {
    expect(savedBuildsMessage(false)).toBe(bulkCopy.talentsSavedUnavailable);
  });
});

describe('loadoutTrees', () => {
  const index = indexTalents(talentsFixture as unknown as TalentFile);

  it('reads one tree per talent tab, each talent at the rank its position in the string gives it', () => {
    const trees = loadoutTrees(index, '23');
    expect(trees).toHaveLength(index.trees.length);
    const [first, second] = index.trees[0].talents;
    expect(trees[0].talents[0]).toMatchObject({ name: first.name, rank: Math.min(2, first.max_rank) });
    expect(trees[0].points).toBe(Math.min(2, first.max_rank) + Math.min(3, second.max_rank));
    expect(trees[1].points).toBe(0);
  });

  it('never draws a rank past the talent’s own maximum', () => {
    const [first] = loadoutTrees(index, '9');
    expect(first.talents[0].rank).toBe(index.trees[0].talents[0].max_rank);
  });
});
