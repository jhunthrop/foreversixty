// web/src/lib/planner/fs1.test.ts
import { describe, expect, it } from 'vitest';
import talents from '../../fixtures/planner/talents/warrior.json';
import { indexTalents } from './rules';
import type { TalentFile } from './types';
import { decodeFS1, encodeFS1, orderFromRanks } from './fs1';

const index = indexTalents(talents as TalentFile);

describe('encodeFS1', () => {
  it('writes the addon spec’s format, base-36 per talent, trailing zeros trimmed', () => {
    expect(
      encodeFS1({
        dataBuild: '1.15.9.69722',
        classSlug: 'paladin',
        raceSlug: 'human',
        treeRanks: [[5, 0, 3, 2, 0, 0, 0, 0, 0], [0], [0]],
        gear: { head: 12640, chest: 11726 },
      }),
    ).toBe('FS1:1.15.9.69722:paladin:human:5032/0/0:head=12640,chest=11726');
  });

  it('writes an empty gear field when nothing is equipped', () => {
    expect(
      encodeFS1({ dataBuild: '1', classSlug: 'mage', raceSlug: 'gnome', treeRanks: [[1], [], []], gear: {} }),
    ).toBe('FS1:1:mage:gnome:1/0/0:');
  });

  it('uses base 36, so a rank above nine is a letter', () => {
    const code = encodeFS1({ dataBuild: '1', classSlug: 'mage', raceSlug: 'gnome', treeRanks: [[12], [], []], gear: {} });
    expect(code).toContain(':c/0/0:');
  });
});

describe('decodeFS1', () => {
  it('reads back what encode wrote', () => {
    const build = {
      dataBuild: '1.15.9.69722',
      classSlug: 'paladin',
      raceSlug: 'human',
      treeRanks: [[5, 0, 3, 2], [0], [0]],
      gear: { head: 12640, chest: 11726 },
    };
    const decoded = decodeFS1(encodeFS1(build));
    expect(decoded.ok).toBe(true);
    if (!decoded.ok) return;
    expect(decoded.build.classSlug).toBe('paladin');
    expect(decoded.build.raceSlug).toBe('human');
    expect(decoded.build.treeRanks[0]).toEqual([5, 0, 3, 2]);
    expect(decoded.build.gear).toEqual({ head: 12640, chest: 11726 });
  });

  it('names the reason for every malformed input rather than failing generically', () => {
    expect(decodeFS1('FS2:1:mage:gnome:1/0/0:')).toEqual({
      ok: false,
      message: 'That code is FS2; this site reads FS1.',
    });
    expect(decodeFS1('FS1:1:mage:gnome')).toEqual({
      ok: false,
      message: 'That code is missing its talent and gear fields.',
    });
    expect(decodeFS1('FS1:1:mage:gnome:1/0:')).toEqual({
      ok: false,
      message: 'That code has 2 talent trees; a build has 3.',
    });
    expect(decodeFS1('FS1:1:mage:gnome:1/0/0:head=nope')).toEqual({
      ok: false,
      message: 'That code has an unreadable gear entry: head=nope.',
    });
    expect(decodeFS1('FS1:1:mage:gnome:1/0/0:elbow=5')).toEqual({
      ok: false,
      message: 'That code names a slot this planner does not have: elbow.',
    });
  });
});

describe('orderFromRanks', () => {
  it('spends the points in a legal order, lowest tier first', () => {
    const first = index.trees[0].talents.find((talent) => talent.tier === 0)!;
    const ranks = index.trees.map((tree) => tree.talents.map((talent) => (talent.id === first.id ? 3 : 0)));
    const result = orderFromRanks(index, ranks);
    expect(result.order).toEqual([first.id, first.id, first.id]);
    expect(result.dropped).toEqual([]);
  });

  it('reports the ranks it could not legally place instead of producing an illegal build', () => {
    const locked = index.trees[0].talents.find((talent) => talent.tier >= 2)!;
    const ranks = index.trees.map((tree) => tree.talents.map((talent) => (talent.id === locked.id ? 1 : 0)));
    const result = orderFromRanks(index, ranks);
    expect(result.order).toEqual([]);
    expect(result.dropped).toEqual([locked.id]);
  });

  it('produces an order the planner’s own rules accept', () => {
    const tier0 = index.trees[0].talents.filter((talent) => talent.tier === 0);
    const ranks = index.trees.map((tree) =>
      tree.talents.map((talent) => (tier0.some((t) => t.id === talent.id) ? talent.max_rank : 0)),
    );
    const result = orderFromRanks(index, ranks);
    expect(result.order.length).toBe(tier0.reduce((total, talent) => total + talent.max_rank, 0));
    expect(result.dropped).toEqual([]);
  });
});
