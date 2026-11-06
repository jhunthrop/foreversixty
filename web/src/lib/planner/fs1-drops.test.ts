// web/src/lib/planner/fs1-drops.test.ts
import { describe, expect, it } from 'vitest';
import { decodeFS1 } from './fs1';

const HEAD = 'FS1:1.60.1.70790:warrior:human:0/0/0:head=16866|level=60';

describe('decodeFS1 drops section', () => {
  it('reads the corpses, items and winners, and does not report the section as ignored', () => {
    const decoded = decodeFS1(
      `${HEAD}|drops=1762480000:249:1084:master:Creature-0-1-249-7-10184-0A:16900x1@Thr%C3%A2ll,18205x2@A+B`,
    );
    expect(decoded.ok).toBe(true);
    if (!decoded.ok) return;
    expect(decoded.build.drops).toEqual([
      {
        at: 1762480000,
        instanceId: 249,
        encounterId: 1084,
        method: 'master',
        source: 'Creature-0-1-249-7-10184-0A',
        items: [
          { itemId: 16900, count: 1, winners: ['Thrâll'] },
          { itemId: 18205, count: 2, winners: ['A', 'B'] },
        ],
      },
    ]);
    expect(decoded.build.ignored).toEqual([]);
  });

  it('drops a malformed corpse or item instead of refusing the code', () => {
    const decoded = decodeFS1(`${HEAD}|drops=bad;5:1:1:::zz,7x1`);
    expect(decoded.ok).toBe(true);
    if (!decoded.ok) return;
    expect(decoded.build.drops).toEqual([
      { at: 5, instanceId: 1, encounterId: 1, method: '', source: '', items: [{ itemId: 7, count: 1, winners: [] }] },
    ]);
  });
});
