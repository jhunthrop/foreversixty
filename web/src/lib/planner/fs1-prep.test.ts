// web/src/lib/planner/fs1-prep.test.ts
import { describe, expect, it } from 'vitest';
import { decodeFS1 } from './fs1';

const HEAD = 'FS1:1.60.1.70790:warrior:human:0/0/0:head=16866|level=60';

describe('decodeFS1 raid-prep sections', () => {
  it('reads school, supplies, locks and quests, and reports none as ignored', () => {
    const decoded = decodeFS1(
      `${HEAD}|school=52:10:0:15:5|supplies=13510:1,20452:20|locks=13704|quests=6502:1,6602:0|owned=16309:1`,
    );
    expect(decoded.ok).toBe(true);
    if (!decoded.ok) return;
    expect(decoded.build.resistances).toEqual({ fire: 52, nature: 10, frost: 0, shadow: 15, arcane: 5 });
    expect(decoded.build.supplies).toEqual([
      { itemId: 13510, count: 1 },
      { itemId: 20452, count: 20 },
    ]);
    expect(decoded.build.keyring).toEqual([13704]);
    expect(decoded.build.quests).toEqual([
      { questId: 6502, done: true },
      { questId: 6602, done: false },
    ]);
    expect(decoded.build.owned).toEqual([{ itemId: 16309, held: true }]);
    expect(decoded.build.ignored).toEqual([]);
  });

  it('drops a malformed entry instead of refusing the code', () => {
    const decoded = decodeFS1(`${HEAD}|school=1:2:3|supplies=x:1,13510:2|locks=a,13704|quests=6502:2,6602:1`);
    expect(decoded.ok).toBe(true);
    if (!decoded.ok) return;
    expect(decoded.build.resistances).toBeUndefined();
    expect(decoded.build.supplies).toEqual([{ itemId: 13510, count: 2 }]);
    expect(decoded.build.keyring).toEqual([13704]);
    expect(decoded.build.quests).toEqual([{ questId: 6602, done: true }]);
  });

  it('leaves the fields absent on an export that predates them, and empty on an empty read', () => {
    const old = decodeFS1(HEAD);
    expect(old.ok && old.build.supplies).toBeUndefined();
    const empty = decodeFS1(`${HEAD}|supplies=|locks=|quests=`);
    expect(empty.ok && [empty.build.supplies, empty.build.keyring, empty.build.quests]).toEqual([[], [], []]);
  });
});
