// web/src/lib/planner/rules.test.ts
import { describe, expect, it } from 'vitest';
import fixtureItems from '../../fixtures/planner/items/warrior.json';
import fixtureTalents from '../../fixtures/planner/talents/warrior.json';
import {
  canAddPoint,
  canEquip,
  comboIsLegal,
  indexItems,
  indexTalents,
  messages,
  canRemovePoint,
  slotsForItem,
  validateGear,
  validateOrder,
  withPoint,
  withoutLastPoint,
} from './rules';
import type { Combo, Gear, ItemFile, TalentFile } from './types';
import { MAX_POINTS } from './types';

const index = indexTalents(fixtureTalents as TalentFile);
const items = indexItems(fixtureItems as ItemFile);

// Arms: 1001 Improved Heroic Strike (3, tier 0), 1002 Deflection (5, tier 0),
//       1003 Improved Rend (3, tier 1), 1004 Tactical Mastery (5, tier 1, needs 1002 x2),
//       1005 Sweeping Strikes (1, tier 2, needs 1003 x3), 1006 Impale (2, tier 2),
//       1007 Axe Specialization (5, tier 3).
// Fury: 2001..2007, same shape.
const repeat = (id: number, n: number): number[] => Array.from({ length: n }, () => id);

describe('withPoint and withoutLastPoint', () => {
  it('return new arrays and never mutate the input', () => {
    const order = [1001, 1002];
    expect(withPoint(order, 1003)).toEqual([1001, 1002, 1003]);
    expect(withoutLastPoint([1001, 1002, 1001], 1001)).toEqual([1001, 1002]);
    expect(order).toEqual([1001, 1002]);
  });

  it('returns the order unchanged when the talent has no points', () => {
    expect(withoutLastPoint([1001], 1006)).toEqual([1001]);
  });
});

describe('rule 3: tier locks at 5 and 10 points in the tree', () => {
  it('refuses tier 1 until five points sit in the tree', () => {
    expect(canAddPoint(index, repeat(1001, 3).concat(repeat(1002, 1)), 1003)).toEqual({
      ok: false,
      reason: messages.tierLocked(1, 'Arms'),
    });
    expect(messages.tierLocked(1, 'Arms')).toBe('Tier 1 of Arms needs 5 points in Arms first');
  });

  it('allows tier 1 on the fifth point in the tree', () => {
    expect(canAddPoint(index, repeat(1001, 3).concat(repeat(1002, 2)), 1003)).toEqual({ ok: true });
  });

  it('refuses tier 2 until ten points sit in the tree', () => {
    const nine = repeat(1001, 3).concat(repeat(1002, 5), repeat(1003, 1));
    expect(canAddPoint(index, nine, 1006)).toEqual({ ok: false, reason: messages.tierLocked(2, 'Arms') });
    expect(messages.tierLocked(2, 'Arms')).toBe('Tier 2 of Arms needs 10 points in Arms first');
    expect(canAddPoint(index, nine.concat(1003), 1006)).toEqual({ ok: true });
  });

  it('counts each tree separately', () => {
    const tenInArms = repeat(1001, 3).concat(repeat(1002, 5), repeat(1003, 2));
    expect(canAddPoint(index, tenInArms, 2003)).toEqual({
      ok: false,
      reason: messages.tierLocked(1, 'Fury'),
    });
  });
});

describe('rule 4: prerequisite rank', () => {
  it('refuses a talent whose prerequisite has too few points', () => {
    const order = repeat(1001, 3).concat(repeat(1002, 1), repeat(1003, 1));
    expect(canAddPoint(index, order, 1004)).toEqual({
      ok: false,
      reason: messages.prereqMissing('Tactical Mastery', 2, 'Deflection'),
    });
    expect(messages.prereqMissing('Tactical Mastery', 2, 'Deflection')).toBe(
      'Tactical Mastery needs 2 points in Deflection first',
    );
  });

  it('uses the singular when one point is enough', () => {
    expect(messages.prereqMissing('Improved Battle Shout', 1, 'Booming Voice')).toBe(
      'Improved Battle Shout needs 1 point in Booming Voice first',
    );
  });

  it('allows the talent once the prerequisite rank is met', () => {
    const order = repeat(1001, 3).concat(repeat(1002, 2));
    expect(canAddPoint(index, order, 1004)).toEqual({ ok: true });
  });
});

describe('rule 5: max rank and the 51-point cap', () => {
  it('refuses a point past the talent maximum', () => {
    expect(canAddPoint(index, repeat(1001, 3), 1001)).toEqual({
      ok: false,
      reason: messages.maxRank('Improved Heroic Strike', 3),
    });
    expect(messages.maxRank('Improved Heroic Strike', 3)).toBe(
      'Improved Heroic Strike is already at 3 of 3 points',
    );
  });

  it('refuses the fifty-second point', () => {
    // All 24 Arms points plus 27 of Fury's 31: exactly 51, breaking no other rule.
    const order = [
      ...repeat(1001, 3),
      ...repeat(1002, 5),
      ...repeat(1003, 3),
      ...repeat(1004, 5),
      ...repeat(1005, 1),
      ...repeat(1006, 2),
      ...repeat(1007, 5),
      ...repeat(2001, 5),
      ...repeat(2002, 5),
      ...repeat(2003, 5),
      ...repeat(2004, 5),
      ...repeat(2005, 5),
      ...repeat(2006, 1),
      ...repeat(2007, 1),
    ];
    expect(order).toHaveLength(MAX_POINTS);
    expect(validateOrder(index, order)).toEqual([]);
    expect(canAddPoint(index, order, 2007)).toEqual({ ok: false, reason: messages.capReached() });
    expect(messages.capReached()).toBe('A build spends at most 51 points');
  });
});

describe('rule 2: talents must belong to the class', () => {
  it('refuses an unknown talent id and reports the offending index', () => {
    expect(canAddPoint(index, [], 9999)).toEqual({ ok: false, reason: messages.unknownTalent(9999) });
    expect(validateOrder(index, [1001, 9999])).toEqual([
      { field: 'point_order[1]', message: messages.unknownTalent(9999) },
    ]);
  });
});

describe('removal', () => {
  it('refuses to remove a point that is not there', () => {
    expect(canRemovePoint(index, [1001], 1006)).toEqual({
      ok: false,
      reason: messages.noPoints('Impale'),
    });
  });

  it('removes the last point in a talent when nothing depends on it', () => {
    const order = repeat(1001, 3).concat(repeat(1002, 2));
    expect(canRemovePoint(index, order, 1002)).toEqual({ ok: true });
    expect(withoutLastPoint(order, 1002)).toEqual(repeat(1001, 3).concat(1002));
  });

  it('refuses when a later point would fall below its tier unlock', () => {
    const order = repeat(1001, 3).concat(repeat(1002, 2), 1003);
    expect(canRemovePoint(index, order, 1001)).toEqual({
      ok: false,
      reason: messages.tierLocked(1, 'Arms'),
    });
  });

  it('refuses when a later point would lose its prerequisite rank', () => {
    // Fury keeps 5 points without the removed one, so the tier stays unlocked and the
    // prerequisite is the only rule that breaks.
    const order = [2001, ...repeat(2002, 5), 2004];
    expect(validateOrder(index, order)).toEqual([]);
    expect(canRemovePoint(index, order, 2001)).toEqual({
      ok: false,
      reason: messages.prereqMissing('Improved Battle Shout', 1, 'Booming Voice'),
    });
  });

  it('refuses dropping a prerequisite below the rank a dependent still needs', () => {
    // Deflection to rank 2 (what Tactical Mastery needs), then one point in it.
    const order = [1002, 1002, 1004];
    expect(canRemovePoint(index, order, 1002)).toEqual({
      ok: false,
      reason: messages.prereqMissing('Tactical Mastery', 2, 'Deflection'),
    });
    // With Deflection at 3 there is a point to spare, so it comes back out.
    expect(canRemovePoint(index, [1002, 1002, 1002, 1004], 1002)).toEqual({ ok: true });
  });
});

describe('rule 1: legal race and class combinations', () => {
  const combos: Combo[] = [
    { race_id: 5, class_id: 2, new_in_forever: true },
    { race_id: 1, class_id: 1, new_in_forever: false },
  ];

  it('accepts a listed pair and refuses an unlisted one', () => {
    expect(comboIsLegal(combos, 5, 2)).toBe(true);
    expect(comboIsLegal(combos, 5, 1)).toBe(false);
  });
});

describe('rule 6: gear', () => {
  it('maps finger and trinket items to both numbered slots', () => {
    expect(slotsForItem(items.get(19325)!)).toEqual(['finger1', 'finger2']);
    expect(slotsForItem(items.get(13968)!)).toEqual(['trinket1', 'trinket2']);
    expect(slotsForItem(items.get(12640)!)).toEqual(['head']);
  });

  it('refuses an item that does not fit the slot', () => {
    expect(canEquip(items, {}, 'chest', 12640)).toEqual({
      ok: false,
      reason: messages.wrongSlot('Lionheart Helm', 'Chest'),
    });
  });

  it('refuses an unknown slot and an unknown item', () => {
    expect(canEquip(items, {}, 'backpack' as never, 12640)).toEqual({
      ok: false,
      reason: messages.unknownSlot('backpack'),
    });
    expect(canEquip(items, {}, 'head', 1)).toEqual({ ok: false, reason: messages.unknownItem(1) });
  });

  it('refuses the same unique ring in both finger slots but allows a shared non-unique one', () => {
    expect(canEquip(items, { finger1: 19325 }, 'finger2', 19325)).toEqual({
      ok: false,
      reason: messages.duplicateUnique('Band of Accuria'),
    });
    expect(canEquip(items, { trinket1: 13968 }, 'trinket2', 13968)).toEqual({ ok: true });
  });

  it('validates a whole gear map and names the offending slot', () => {
    const gear: Gear = { head: 12640, chest: 12784 };
    expect(validateGear(items, gear)).toEqual([
      { field: 'gear.chest', message: messages.wrongSlot('Arcanite Reaper', 'Chest') },
    ]);
    expect(validateGear(items, { head: 12640, main_hand: 12784 })).toEqual([]);
  });
});

describe('validateOrder with more than one offending index', () => {
  it('counts a refused point toward the tree total, so a later tier can unlock on it', () => {
    // 1001 x3 + 1002 x1 = 4 real, legal points in Arms (tier 0 needs none of them).
    // A 5th 1001 is refused (already at its 3-point max) at index 4, but the tree
    // counter still advances to 5 for it, as if it were a legal point.
    // 1004 (tier 1, needs 5 points in Arms) then reads inTree === 5 and treats tier 1
    // as unlocked, even though only 4 points were ever legally spent there. Tier 1 being
    // "unlocked" only means the tier check does not fire; validateOrder falls through to
    // the next rule, and 1004 still needs rank 2 in its prerequisite (1002), which has
    // only rank 1. So index 5 is refused for the prerequisite, not the tier lock it would
    // have hit had the phantom count not been added.
    const order = [1001, 1001, 1001, 1002, 1001, 1004];
    expect(validateOrder(index, order)).toEqual([
      { field: 'point_order[4]', message: messages.maxRank('Improved Heroic Strike', 3) },
      {
        field: 'point_order[5]',
        message: messages.prereqMissing('Tactical Mastery', 2, 'Deflection'),
      },
    ]);
  });
});

describe('rule 6: two-handed main hand', () => {
  it('refuses an off-hand item beside a two-handed main hand', () => {
    const items = new Map([
      [1, { id: 1, slot: 'main_hand', two_hand: true, name: 'Big Axe' }],
      [2, { id: 2, slot: 'off_hand', two_hand: false, name: 'Shield' }],
    ]) as never;
    const errors = validateGear(items, { main_hand: 1, off_hand: 2 });
    expect(errors.map((error) => error.message)).toContain(messages.twoHandOffHand('Big Axe'));
  });

  it('allows an off-hand item beside a one-handed main hand', () => {
    const items = new Map([
      [1, { id: 1, slot: 'main_hand', two_hand: false, name: 'Sword' }],
      [2, { id: 2, slot: 'off_hand', two_hand: false, name: 'Shield' }],
    ]) as never;
    expect(validateGear(items, { main_hand: 1, off_hand: 2 })).toEqual([]);
  });
});
