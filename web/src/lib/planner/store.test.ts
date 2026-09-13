// web/src/lib/planner/store.test.ts
import { describe, expect, it } from 'vitest';
import fixtureClasses from '../../fixtures/planner/classes.json';
import fixtureCombos from '../../fixtures/planner/combos.json';
import fixtureItems from '../../fixtures/planner/items/warrior.json';
import fixtureRaces from '../../fixtures/planner/races.json';
import fixtureSets from '../../fixtures/planner/sets.json';
import fixtureTalents from '../../fixtures/planner/talents/warrior.json';
import { messages } from './rules';
import { createPlannerStore, READ_ONLY_REASON } from './store.svelte';
import type { ClassRow, Combo, ItemFile, ItemSet, RaceRow, TalentFile } from './types';

const BUILD = '1.15.9.69722';

function loaded(overrides: Parameters<typeof createPlannerStore>[0] | null = null) {
  const store = createPlannerStore({
    treeVersion: BUILD,
    classSlug: 'warrior',
    raceSlug: 'human',
    ...(overrides ?? {}),
  });
  store.setReference({
    classes: fixtureClasses as ClassRow[],
    races: fixtureRaces as RaceRow[],
    combos: fixtureCombos as Combo[],
  });
  store.setTalents(fixtureTalents as TalentFile);
  store.setItems(fixtureItems as ItemFile);
  store.setSets(fixtureSets as ItemSet[]);
  return store;
}

describe('spending points', () => {
  it('adds a legal point, tracks the level and the split, and clears the refusal', () => {
    const store = loaded();
    store.addPoint(1001);
    store.addPoint(2001);
    expect(store.order).toEqual([1001, 2001]);
    expect(store.spent).toBe(2);
    expect(store.level).toBe(11);
    expect(store.splitLabel).toBe('1/1');
    expect(store.refusal).toBeNull();
  });

  it('refuses an illegal point, records the reason and leaves the order alone', () => {
    const store = loaded();
    store.addPoint(1003);
    expect(store.order).toEqual([]);
    expect(store.refusal).toBe(messages.tierLocked(1, 'Arms'));
    store.addPoint(1001);
    expect(store.order).toEqual([1001]);
    expect(store.refusal).toBeNull();
  });

  it('removes the last point in a talent and refuses when something depends on it', () => {
    const store = loaded();
    for (const id of [1001, 1001, 1001, 1002, 1002]) store.addPoint(id);
    store.addPoint(1003);
    store.removePoint(1001);
    expect(store.refusal).toBe(messages.tierLocked(1, 'Arms'));
    expect(store.order).toHaveLength(6);
    store.removePoint(1003);
    expect(store.refusal).toBeNull();
    expect(store.order).toEqual([1001, 1001, 1001, 1002, 1002]);
  });

  it('exposes per-talent ranks for the cells', () => {
    const store = loaded();
    store.addPoint(1001);
    store.addPoint(1001);
    expect(store.ranks.get(1001)).toBe(2);
    expect(store.ranks.get(1002)).toBeUndefined();
  });
});

describe('class and race selection', () => {
  it('clears the build and the loaded trees when the class changes', () => {
    const store = loaded();
    store.addPoint(1001);
    store.selectClass('paladin');
    expect(store.classSlug).toBe('paladin');
    expect(store.order).toEqual([]);
    expect(store.talents).toBeNull();
    expect(store.talentIndex).toBeNull();
  });

  it('keeps the build when only the race changes', () => {
    const store = loaded();
    store.addPoint(1001);
    store.selectRace('dwarf');
    expect(store.raceSlug).toBe('dwarf');
    expect(store.order).toEqual([1001]);
  });

  it('moves to the first legal race when the current one cannot be the new class', () => {
    const store = loaded({ treeVersion: BUILD, classSlug: 'warrior', raceSlug: 'gnome' });
    store.selectClass('shaman');
    expect(store.legalRaces.map((r) => r.slug)).toEqual(['orc', 'tauren', 'troll']);
    expect(store.raceSlug).toBe('orc');
  });
});

describe('read-only builds, fork and reset', () => {
  it('refuses every edit while read-only', () => {
    const store = loaded({
      treeVersion: BUILD,
      classSlug: 'warrior',
      raceSlug: 'human',
      order: [1001],
      readOnly: true,
      sourceId: 'k7x2qm4a',
    });
    store.addPoint(1002);
    expect(store.order).toEqual([1001]);
    expect(store.refusal).toBe(READ_ONLY_REASON);
    store.removePoint(1001);
    expect(store.order).toEqual([1001]);
  });

  it('fork makes the same build editable and drops the source id and title', () => {
    const store = loaded({
      treeVersion: BUILD,
      classSlug: 'warrior',
      raceSlug: 'human',
      order: [1001],
      title: 'Original',
      readOnly: true,
      sourceId: 'k7x2qm4a',
    });
    store.fork();
    expect(store.readOnly).toBe(false);
    expect(store.sourceId).toBeNull();
    expect(store.title).toBe('');
    expect(store.order).toEqual([1001]);
    store.addPoint(1002);
    expect(store.order).toEqual([1001, 1002]);
  });

  it('reset clears points, gear and the refusal', () => {
    const store = loaded();
    store.addPoint(1001);
    store.equip('head', 12640);
    store.addPoint(1003);
    store.reset();
    expect(store.order).toEqual([]);
    expect(store.gear).toEqual({});
    expect(store.refusal).toBeNull();
  });
});

describe('gear', () => {
  it('equips a fitting item and refuses one that does not fit', () => {
    const store = loaded();
    store.equip('head', 12640);
    expect(store.gear).toEqual({ head: 12640 });
    store.equip('chest', 12640);
    expect(store.gear).toEqual({ head: 12640 });
    expect(store.refusal).toBe(messages.wrongSlot('Lionheart Helm', 'Chest'));
  });

  it('unequips a slot and keeps the rest', () => {
    const store = loaded();
    store.equip('head', 12640);
    store.equip('main_hand', 12784);
    store.unequip('head');
    expect(store.gear).toEqual({ main_hand: 12784 });
  });

  it('totals stats and activates set bonuses from the equipped items', () => {
    const store = loaded();
    store.equip('head', 16963);
    store.equip('shoulder', 16966);
    expect(store.statTotals).toEqual({ strength: 45, stamina: 38, armor: 1110 });
    expect(store.activeSets).toHaveLength(1);
    expect(store.activeSets[0].pieces).toBe(2);
    expect(store.activeSets[0].active).toHaveLength(1);
  });
});

describe('toDraft', () => {
  it('produces exactly the POST /v1/builds body', () => {
    const store = loaded();
    store.addPoint(1001);
    store.equip('head', 12640);
    store.setTitle('  Arms leveling  ');
    expect(store.toDraft()).toEqual({
      class_id: 1,
      race_id: 1,
      tree_version: BUILD,
      point_order: [1001],
      gear: { head: 12640 },
      title: 'Arms leveling',
    });
  });

  it('omits an empty title and caps a long one at 60 characters', () => {
    const store = loaded();
    expect(store.toDraft().title).toBeUndefined();
    store.setTitle('x'.repeat(80));
    expect(store.toDraft().title).toHaveLength(60);
  });
});
