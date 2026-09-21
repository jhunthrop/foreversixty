// web/src/lib/planner/fs1.test.ts
import { describe, expect, it } from 'vitest';
import talents from '../../fixtures/planner/talents/warrior.json';
import { indexTalents } from './rules';
import type { TalentFile } from './types';
import { decodeFS1, encodeFS1, encodeFS1V2, MAX_CODE_LENGTH, orderFromRanks } from './fs1';

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
    const code = encodeFS1({
      dataBuild: '1',
      classSlug: 'mage',
      raceSlug: 'gnome',
      treeRanks: [[12], [], []],
      gear: {},
    });
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

  it('refuses a gear value that is not entirely digits, rather than truncating it', () => {
    // Number.parseInt stops at the first non-digit and returns what came before it, so these
    // would otherwise silently become the item id 12640 instead of being refused.
    expect(decodeFS1('FS1:1:mage:gnome:1/0/0:head=12640abc')).toEqual({
      ok: false,
      message: 'That code has an unreadable gear entry: head=12640abc.',
    });
    expect(decodeFS1('FS1:1:mage:gnome:1/0/0:head=12640.5')).toEqual({
      ok: false,
      message: 'That code has an unreadable gear entry: head=12640.5.',
    });
  });

  it('refuses a code past the sane length bound before parsing any of it', () => {
    const hostile = `FS1:1:mage:gnome:${'1'.repeat(MAX_CODE_LENGTH)}/0/0:`;
    expect(decodeFS1(hostile)).toEqual({ ok: false, message: 'That code is too long to read.' });
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

describe('version 2 sections', () => {
  const V1 = 'FS1:1.15.9:warrior:orc:0/5530515/0:head=12640,main_hand=11726';

  it('reads a version 1 code unchanged, with the new fields empty', () => {
    const decoded = decodeFS1(V1);
    expect(decoded.ok).toBe(true);
    if (!decoded.ok) return;
    expect(decoded.build.gear).toEqual({ head: 12640, main_hand: 11726 });
    expect(decoded.build.gearSlots).toEqual([
      { slot: 'head', itemId: 12640 },
      { slot: 'main_hand', itemId: 11726 },
    ]);
    expect(decoded.build.bags).toEqual([]);
    expect(decoded.build.bank).toEqual([]);
    expect(decoded.build.sets).toEqual([]);
    expect(decoded.build.loadouts).toEqual([]);
    expect(decoded.build.professions).toEqual([]);
    expect(decoded.build.ignored).toEqual([]);
  });

  it('reads an enchant and a suffix on a gear entry (contract 10.5), keeping the id map lossy', () => {
    const decoded = decodeFS1('FS1:1.15.9:warrior:orc:0/5530515/0:head=12640:2504,main_hand=11726:2505:1820');
    expect(decoded.ok).toBe(true);
    if (!decoded.ok) return;
    // `gear` stays the planner's map of ids -- the strip, the planner link and the build
    // draft all read it and none of them models an enchant.
    expect(decoded.build.gear).toEqual({ head: 12640, main_hand: 11726 });
    // `gearSlots` is the whole truth, and is what SimCharacter and the request carry.
    expect(decoded.build.gearSlots).toEqual([
      { slot: 'head', itemId: 12640, enchant: 2504 },
      { slot: 'main_hand', itemId: 11726, enchant: 2505, suffix: 1820 },
    ]);
  });

  it('reads the professions section', () => {
    const decoded = decodeFS1(`${V1}|professions=engineering,blacksmithing`);
    expect(decoded.ok).toBe(true);
    if (!decoded.ok) return;
    expect(decoded.build.professions).toEqual(['engineering', 'blacksmithing']);
  });

  it('reads the guild section', () => {
    const decoded = decodeFS1(`${V1}|guild=Iron%20Vanguard:2`);
    expect(decoded.ok).toBe(true);
    if (!decoded.ok) return;
    expect(decoded.build.guild).toEqual({ name: 'Iron Vanguard', rankIndex: 2 });
  });

  it('leaves guild undefined, never an empty object, when the code carries no guild section', () => {
    const decoded = decodeFS1(V1);
    expect(decoded.ok).toBe(true);
    if (!decoded.ok) return;
    expect(decoded.build.guild).toBeUndefined();
  });

  it('refuses a non-numeric guild rank rather than silently dropping the section', () => {
    const decoded = decodeFS1(`${V1}|guild=Iron%20Vanguard:officer`);
    expect(decoded.ok).toBe(false);
    if (decoded.ok) return;
    expect(decoded.message).toBe('That code has an unreadable guild rank: officer.');
  });

  it('splits the guild payload on the first colon, so a name containing one still reads (pre-escaped)', () => {
    // A literal colon in a guild name is always percent-encoded by encodeURIComponent
    // before it reaches this format (RFC 3986's unreserved set excludes ':'), so this
    // is exercising the decoder's own split rule, not a real un-escaped name.
    const decoded = decodeFS1(`${V1}|guild=A%3AB:3`);
    expect(decoded.ok).toBe(true);
    if (!decoded.ok) return;
    expect(decoded.build.guild).toEqual({ name: 'A:B', rankIndex: 3 });
  });

  it('reads bags and bank, with the optional enchant and suffix', () => {
    const decoded = decodeFS1(`${V1}|bags=16963,17076:2504,19360:2505:1820|bank=12640`);
    expect(decoded.ok).toBe(true);
    if (!decoded.ok) return;
    expect(decoded.build.bags).toEqual([
      { itemId: 16963 },
      { itemId: 17076, enchant: 2504 },
      { itemId: 19360, enchant: 2505, suffix: 1820 },
    ]);
    expect(decoded.build.bank).toEqual([{ itemId: 12640 }]);
  });

  it('reads named sets whose gear is a gear list, enchants and suffixes included', () => {
    const decoded = decodeFS1(`${V1}|sets=AQ%20set=head=21329:2504,chest=21330;PvP=head=16963`);
    expect(decoded.ok).toBe(true);
    if (!decoded.ok) return;
    expect(decoded.build.sets).toEqual([
      {
        name: 'AQ set',
        gear: [
          { slot: 'head', itemId: 21329, enchant: 2504 },
          { slot: 'chest', itemId: 21330 },
        ],
      },
      { name: 'PvP', gear: [{ slot: 'head', itemId: 16963 }] },
    ]);
  });

  it('reads named loadouts as three trees each', () => {
    const decoded = decodeFS1(`${V1}|loadouts=Deep%20Fury=0/5530515/0;Arms=5530515/0/0`);
    expect(decoded.ok).toBe(true);
    if (!decoded.ok) return;
    expect(decoded.build.loadouts.map((row) => row.name)).toEqual(['Deep Fury', 'Arms']);
    expect(decoded.build.loadouts[0].treeRanks[1]).toEqual([5, 5, 3, 0, 5, 1, 5]);
  });

  it('ignores a section it does not know and reports its name', () => {
    const decoded = decodeFS1(`${V1}|bags=16963|quiver=1234|bank=12640`);
    expect(decoded.ok).toBe(true);
    if (!decoded.ok) return;
    expect(decoded.build.ignored).toEqual(['quiver']);
    expect(decoded.build.bags).toEqual([{ itemId: 16963 }]);
    expect(decoded.build.bank).toEqual([{ itemId: 12640 }]);
  });

  it('refuses an unreadable item entry rather than silently dropping it', () => {
    const decoded = decodeFS1(`${V1}|bags=16963,notanid`);
    expect(decoded.ok).toBe(false);
    if (decoded.ok) return;
    expect(decoded.message).toContain('notanid');
  });

  it('refuses a gear entry carrying a second "=" rather than silently discarding it', () => {
    const decoded = decodeFS1(`${V1.replace(':head=12640', ':head=12640=99')}`);
    expect(decoded.ok).toBe(false);
    if (decoded.ok) return;
    expect(decoded.message).toContain('head=12640=99');
  });

  it('drops an empty profession slug from a stray comma rather than forwarding it', () => {
    const decoded = decodeFS1(`${V1}|professions=engineering,,alchemy,`);
    expect(decoded.ok).toBe(true);
    if (!decoded.ok) return;
    expect(decoded.build.professions).toEqual(['engineering', 'alchemy']);
  });

  // A malformed percent-encoding must never throw: `Planner.svelte` calls `decodeFS1`
  // synchronously with no try, and `sources.ts`'s `fromAddonExport` calls
  // `characterFromFs1` outside its own -- a third-party addon exporting a set or loadout
  // name it did not URL-encode (a bare "%", or "%" not followed by two hex digits) is an
  // honest string the player typed, not an attack to refuse.
  it('reads a bare "%" in a set name as itself rather than throwing', () => {
    expect(() => decodeFS1(`${V1}|sets=%`)).not.toThrow();
    const decoded = decodeFS1(`${V1}|sets=%`);
    expect(decoded.ok).toBe(true);
    if (!decoded.ok) return;
    expect(decoded.build.sets).toEqual([{ name: '%', gear: [] }]);
  });

  it('reads an unescaped "%" in a loadout name as itself rather than throwing', () => {
    expect(() => decodeFS1(`${V1}|loadouts=Hit% Set=1/0/0`)).not.toThrow();
    const decoded = decodeFS1(`${V1}|loadouts=Hit% Set=1/0/0`);
    expect(decoded.ok).toBe(true);
    if (!decoded.ok) return;
    expect(decoded.build.loadouts[0].name).toBe('Hit% Set');
  });

  it('accepts a code up to the new bound and refuses one past it', () => {
    expect(MAX_CODE_LENGTH).toBe(16_384);
    const long = `${V1}|bags=${Array.from({ length: 900 }, () => '16963').join(',')}`;
    expect(long.length).toBeLessThan(MAX_CODE_LENGTH);
    expect(decodeFS1(long).ok).toBe(true);
    expect(decodeFS1('x'.repeat(MAX_CODE_LENGTH + 1)).ok).toBe(false);
  });
});

describe('encodeFS1V2', () => {
  it('is encodeFS1 exactly when there is nothing after the gear and nothing to enchant', () => {
    const build = {
      dataBuild: '1.15.9',
      classSlug: 'warrior',
      raceSlug: 'orc',
      treeRanks: [[], [5, 5, 3, 0, 5, 1, 5], []],
      gear: { head: 12640 },
      bags: [],
      bank: [],
      sets: [],
      loadouts: [],
      professions: [],
      ignored: [],
    };
    expect(encodeFS1V2(build)).toBe(encodeFS1(build));
  });

  it('writes an enchant and a suffix onto a gear entry, and round-trips them', () => {
    const code = encodeFS1V2({
      dataBuild: '1.15.9',
      classSlug: 'warrior',
      raceSlug: 'orc',
      treeRanks: [[], [], []],
      gear: { head: 12640 },
      gearSlots: [{ slot: 'head', itemId: 12640, enchant: 2504, suffix: 1820 }],
      bags: [],
      bank: [],
      sets: [],
      loadouts: [],
      professions: [],
      ignored: [],
    });
    expect(code).toContain('head=12640:2504:1820');
    const decoded = decodeFS1(code);
    expect(decoded.ok).toBe(true);
    if (decoded.ok) {
      expect(decoded.build.gearSlots).toEqual([{ slot: 'head', itemId: 12640, enchant: 2504, suffix: 1820 }]);
    }
  });

  it('accepts a build that omits every version 2 field, per the contract’s "all optional"', () => {
    // Spec section 7: the decoder returns FS1Build's bags/bank/sets/loadouts "all
    // optional". A version-1-only caller (Planner.svelte's own "Sim this build" link,
    // for one) should not have to write six empty arrays to say it has nothing to add.
    expect(
      encodeFS1V2({
        dataBuild: '1.15.9',
        classSlug: 'warrior',
        raceSlug: 'orc',
        treeRanks: [[], [], []],
        gear: { head: 12640 },
      }),
    ).toBe('FS1:1.15.9:warrior:orc:0/0/0:head=12640');
  });

  it('writes gear from gearSlots alone, rather than silently dropping it when gear is empty', () => {
    // A caller can hold gearSlots with nothing mirrored into the lossy `gear` map -- the
    // doc comment on FS1Build.gearSlots says as much. Falling through to encodeFS1 (which
    // only ever reads `gear`) for a build with no enchant would have written no gear field
    // at all here.
    const code = encodeFS1V2({
      dataBuild: '1.15.9',
      classSlug: 'warrior',
      raceSlug: 'orc',
      treeRanks: [[], [], []],
      gear: {},
      gearSlots: [{ slot: 'head', itemId: 12640 }],
      bags: [],
      bank: [],
      sets: [],
      loadouts: [],
      professions: [],
      ignored: [],
    });
    expect(code).toContain('head=12640');
    const decoded = decodeFS1(code);
    expect(decoded.ok).toBe(true);
    if (decoded.ok) expect(decoded.build.gearSlots).toEqual([{ slot: 'head', itemId: 12640 }]);
  });

  it('writes the sections in the contract’s order and round-trips them', () => {
    const build = {
      dataBuild: '1.15.9',
      classSlug: 'warrior',
      raceSlug: 'orc',
      treeRanks: [[], [5, 5, 3, 0, 5, 1, 5], []],
      gear: { head: 12640 },
      bags: [{ itemId: 16963, enchant: 2504 }],
      bank: [{ itemId: 19360, enchant: 2505, suffix: 1820 }],
      sets: [{ name: 'AQ set', gear: [{ slot: 'head' as const, itemId: 21329 }] }],
      loadouts: [{ name: 'Deep Fury', treeRanks: [[], [5, 5, 3, 0, 5, 1, 5], []] }],
      professions: ['engineering', 'blacksmithing'],
      ignored: [],
    };
    const code = encodeFS1V2(build);
    expect(code.indexOf('|bags=')).toBeLessThan(code.indexOf('|bank='));
    expect(code.indexOf('|bank=')).toBeLessThan(code.indexOf('|sets='));
    expect(code.indexOf('|sets=')).toBeLessThan(code.indexOf('|loadouts='));
    expect(code.indexOf('|loadouts=')).toBeLessThan(code.indexOf('|professions='));

    const decoded = decodeFS1(code);
    expect(decoded.ok).toBe(true);
    if (!decoded.ok) return;
    expect(decoded.build.bags).toEqual(build.bags);
    expect(decoded.build.bank).toEqual(build.bank);
    expect(decoded.build.sets).toEqual(build.sets);
    expect(decoded.build.loadouts[0].name).toBe('Deep Fury');
    expect(decoded.build.professions).toEqual(['engineering', 'blacksmithing']);
  });

  it('writes the guild section after professions and round-trips it', () => {
    const build = {
      dataBuild: '1.15.9',
      classSlug: 'warrior',
      raceSlug: 'orc',
      treeRanks: [[], [], []],
      gear: {},
      bags: [],
      bank: [],
      sets: [],
      loadouts: [],
      professions: ['engineering'],
      guild: { name: 'Iron Vanguard', rankIndex: 2 },
      ignored: [],
    };
    const code = encodeFS1V2(build);
    expect(code.indexOf('|professions=')).toBeLessThan(code.indexOf('|guild='));
    expect(code).toContain('|guild=Iron%20Vanguard:2');

    const decoded = decodeFS1(code);
    expect(decoded.ok).toBe(true);
    if (decoded.ok) expect(decoded.build.guild).toEqual({ name: 'Iron Vanguard', rankIndex: 2 });
  });

  it('omits the guild section entirely when build.guild is absent', () => {
    const code = encodeFS1V2({
      dataBuild: '1.15.9',
      classSlug: 'warrior',
      raceSlug: 'orc',
      treeRanks: [[], [], []],
      gear: {},
      bags: [],
      bank: [],
      sets: [],
      loadouts: [],
      professions: [],
      ignored: [],
    });
    expect(code).not.toContain('guild=');
  });

  it('escapes a name carrying the separators it would otherwise break on', () => {
    const code = encodeFS1V2({
      dataBuild: '1.15.9',
      classSlug: 'warrior',
      raceSlug: 'orc',
      treeRanks: [[], [], []],
      gear: {},
      bags: [],
      bank: [],
      sets: [{ name: 'a;b=c|d', gear: [{ slot: 'head' as const, itemId: 1 }] }],
      loadouts: [],
      professions: [],
      ignored: [],
    });
    const decoded = decodeFS1(code);
    expect(decoded.ok).toBe(true);
    if (decoded.ok) expect(decoded.build.sets[0].name).toBe('a;b=c|d');
  });
});
