import { describe, expect, it } from 'vitest';
import { fixtureResult } from '../../test-support/sim-api';
import { parseActionKey, resolveActionName, type ActionNames } from './action-names';

const names: ActionNames = {
  spell: { '25286': 'Heroic Strike', '20662': 'Execute' },
  item: { '14554': 'Cloudkeeper Legplates' },
};

describe('parseActionKey', () => {
  it('reads every shape sim/adapter/identity.go emits', () => {
    expect(parseActionKey('spell:25286')).toEqual({
      kind: 'spell',
      id: 25286,
      label: '25286',
      tag: 0,
      rank: 0,
    });
    expect(parseActionKey('spell:25286/1')).toEqual({
      kind: 'spell',
      id: 25286,
      label: '25286',
      tag: 1,
      rank: 0,
    });
    expect(parseActionKey('spell:25286+r3')).toEqual({
      kind: 'spell',
      id: 25286,
      label: '25286',
      tag: 0,
      rank: 3,
    });
    expect(parseActionKey('spell:25286/1+r3')).toEqual({
      kind: 'spell',
      id: 25286,
      label: '25286',
      tag: 1,
      rank: 3,
    });
    expect(parseActionKey('item:14554')).toEqual({
      kind: 'item',
      id: 14554,
      label: '14554',
      tag: 0,
      rank: 0,
    });
    expect(parseActionKey('other:attack/2')).toEqual({
      kind: 'other',
      id: 0,
      label: 'attack',
      tag: 2,
      rank: 0,
    });
    expect(parseActionKey('other:rage_gain')).toEqual({
      kind: 'other',
      id: 0,
      label: 'rage_gain',
      tag: 0,
      rank: 0,
    });
    expect(parseActionKey('unknown')).toEqual({ kind: 'unknown', id: 0, label: '', tag: 0, rank: 0 });
  });

  it('reads every key the checked-in fixture actually carries', () => {
    const summary = fixtureResult.summary;
    const keys = [
      ...summary.damage_done.flatMap((actor) => actor.abilities.map((ability) => ability.name)),
      ...summary.casts.map((row) => row.spell_name),
      ...summary.auras.map((track) => track.name),
    ];
    expect(keys.length).toBeGreaterThan(0);
    for (const key of keys) expect(parseActionKey(key)).not.toBeNull();
  });

  it('is null for a display name, which is what a logged fight carries', () => {
    expect(parseActionKey('Bloodthirst')).toBeNull();
    expect(parseActionKey('spell:')).toBeNull();
    expect(parseActionKey('spell:abc')).toBeNull();
    expect(parseActionKey('')).toBeNull();
  });
});

describe('resolveActionName', () => {
  it('names a spell, an item and an other action', () => {
    expect(resolveActionName('spell:25286', names)).toBe('Heroic Strike');
    expect(resolveActionName('item:14554', names)).toBe('Cloudkeeper Legplates');
    // The engine writes its own OtherAction name into the key; this only tidies the case.
    expect(resolveActionName('other:attack', names)).toBe('Attack');
    expect(resolveActionName('other:rage_gain', names)).toBe('Rage gain');
  });

  it('marks a tagged or ranked action as the variant it is, not as a duplicate row', () => {
    // The engine splits one spell into several metric rows -- the three tags of a white
    // swing, a ranked cast -- and rows all reading "Heroic Strike" would look like a bug.
    // "other:attack" is excluded here (its own test below): its tags name a hand, not a
    // row number, so it does not go through this generic variant numbering.
    expect(resolveActionName('spell:25286/1', names)).toBe('Heroic Strike (2)');
    expect(resolveActionName('other:rage_gain/2', names)).toBe('Rage gain (3)');
    expect(resolveActionName('spell:25286+r3', names)).toBe('Heroic Strike (Rank 3)');
    expect(resolveActionName('spell:25286/1+r3', names)).toBe('Heroic Strike (2, Rank 3)');
  });

  it('names the auto-attack tag by the hand it swung from, not by row number', () => {
    // sim/core/attack.go: tagMainhand = 1, tagOffhand = 2, tagExtraAttack = 3. The old
    // behaviour numbered tag 1 as "Attack (2)" -- the *second* row -- and copy.ts's
    // sentence prose then read that as off-hand, so a two-handed weapon's only attack
    // row was described as off-hand damage (dps-minmaxer review round 1, D2).
    expect(resolveActionName('other:attack/1', names)).toBe('Main-hand attacks');
    expect(resolveActionName('other:attack/2', names)).toBe('Off-hand attacks');
    expect(resolveActionName('other:attack/3', names)).toBe('Extra attacks');
  });

  it('falls back to a humanised label for an attack tag this build has never seen', () => {
    // No wowsims/classic tag reaches 9; this pins that an unrecognised tag still reads as
    // English and never leaks the raw key.
    expect(resolveActionName('other:attack/9', names)).toBe('Attack (10)');
  });

  it('falls back to prose, never the raw key or the id itself, when the build does not know the id', () => {
    // dps-minmaxer review round 2, D48: a saved run rendered "spell:20662" as an ability
    // name. 2026-09-21 result-page review round 2: "Spell 20662" -- action-names.ts's own
    // fix for D48 -- turned out to still be an id reaching a player, just spelled out. An
    // id the table never carries -- rather than one merely still loading, the next test --
    // reads as prose with no number in it at all.
    expect(resolveActionName('spell:999999', names)).toBe('An unnamed spell');
    expect(resolveActionName('item:1', names)).toBe('An unnamed item');
  });

  it('falls back to prose, never the raw key or the id itself, when the names have not loaded yet', () => {
    expect(resolveActionName('spell:25286', null)).toBe('An unnamed spell');
    // An other action needs no table, so it reads properly even before one loads.
    expect(resolveActionName('other:attack', null)).toBe('Attack');
  });

  it('keeps a tagged or ranked action’s variant suffix on the unresolved fallback too', () => {
    expect(resolveActionName('spell:999999/1', names)).toBe('An unnamed spell (2)');
    expect(resolveActionName('spell:999999+r3', names)).toBe('An unnamed spell (Rank 3)');
  });

  it('passes a real display name straight through, so a logged fight renders unchanged', () => {
    // compare.ts puts a sim row beside a fight row, and the fight's rows carry real names.
    expect(resolveActionName('Bloodthirst', names)).toBe('Bloodthirst');
  });
});
