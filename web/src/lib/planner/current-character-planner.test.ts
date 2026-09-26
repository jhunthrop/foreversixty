// web/src/lib/planner/current-character-planner.test.ts
import { describe, expect, it } from 'vitest';
import fixtureClasses from '../../fixtures/planner/classes.json';
import fixtureCombos from '../../fixtures/planner/combos.json';
import fixtureRaces from '../../fixtures/planner/races.json';
import fixtureTalents from '../../fixtures/planner/talents/warrior.json';
import { readCurrent, type CurrentCharacter } from '../current-character';
import {
  decidePlannerLoad,
  isBarePlannerUrl,
  labelForPlannerLoad,
  plannerAddonCode,
  recordPlannerCharacter,
  unsavedPlannerHref,
  writePlannerPointer,
  type PlannerLoadDecision,
} from './current-character-planner';
import { decodeFS1, orderFromRanks } from './fs1';
import { indexTalents } from './rules';
import { createPlannerStore, type PlannerStore } from './store.svelte';
import type { ClassRow, Combo, RaceRow, TalentFile } from './types';

const BUILD = '1.15.9.69722';

/** Same loading sequence as store.test.ts's own `loaded()`: reference data plus this
 *  class's talent file, the two things `characterFromPlanner` needs to convert the store's
 *  live build into a `SimCharacter`. Named apart from `writePlannerPointer`'s own local
 *  `loadedStore` fixture below, which is a plain object literal rather than a real store. */
function loadedPlannerStore(
  overrides: Partial<Parameters<typeof createPlannerStore>[0]> | null = null,
): PlannerStore {
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
  return store;
}

function fakeStorage(): Storage {
  const map = new Map<string, string>();
  return {
    getItem: (key) => map.get(key) ?? null,
    setItem: (key, value) => void map.set(key, value),
    removeItem: (key) => void map.delete(key),
    clear: () => map.clear(),
    key: (index) => [...map.keys()][index] ?? null,
    get length() {
      return map.size;
    },
  };
}

function stored(source: CurrentCharacter['source'], ref: string): CurrentCharacter {
  return { source, ref, label: 'Warrior', classSlug: 'warrior', savedAt: '2026-09-21T00:00:00.000Z' };
}

const GOOD_CODE = 'FS1:1:warrior:orc:0/5530515/0:';
const GOOD_DECODED = decodeFS1(GOOD_CODE);
const OTHER_DECODED = decodeFS1('FS1:1:warrior:orc:1:');

describe('isBarePlannerUrl', () => {
  it('is true for an empty query', () => {
    expect(isBarePlannerUrl('')).toBe(true);
  });
  it('is false when ?code= is present', () => {
    expect(isBarePlannerUrl('?code=FS1:1:warrior:orc:0/0/0:')).toBe(false);
  });
  it('is false when ?class= is present', () => {
    expect(isBarePlannerUrl('?class=warrior')).toBe(false);
  });
  it('is false when ?race= is present', () => {
    expect(isBarePlannerUrl('?race=orc')).toBe(false);
  });
});

describe('decidePlannerLoad', () => {
  it('prefers the URL code over a stored pointer, but still surfaces the pointer for the chip', () => {
    const pointer = stored('addon', 'FS1:1:warrior:orc:1:');
    const decision = decidePlannerLoad(GOOD_CODE, true, false, true, pointer);
    expect(decision).toEqual<PlannerLoadDecision>({
      codeParam: GOOD_CODE,
      decoded: GOOD_DECODED,
      restored: false,
      deadPointer: false,
      pointer,
      initialClassSlug: null,
    });
  });

  it('restores a code-sourced pointer on a bare, standalone, record-less mount', () => {
    const pointer = stored('code', GOOD_CODE);
    const decision = decidePlannerLoad(null, true, false, true, pointer);
    expect(decision).toEqual<PlannerLoadDecision>({
      codeParam: GOOD_CODE,
      decoded: GOOD_DECODED,
      restored: true,
      deadPointer: false,
      pointer,
      initialClassSlug: 'warrior',
    });
  });

  it('restores an addon-sourced pointer the same way', () => {
    const pointer = stored('addon', GOOD_CODE);
    const decision = decidePlannerLoad(null, true, false, true, pointer);
    expect(decision).toEqual<PlannerLoadDecision>({
      codeParam: GOOD_CODE,
      decoded: GOOD_DECODED,
      restored: true,
      deadPointer: false,
      pointer,
      initialClassSlug: 'warrior',
    });
  });

  it('does not restore a build-sourced pointer -- it has its own permalink -- but the chip still shows it', () => {
    const pointer = stored('build', 'b1');
    const decision = decidePlannerLoad(null, true, false, true, pointer);
    expect(decision).toEqual<PlannerLoadDecision>({
      codeParam: null,
      decoded: null,
      restored: false,
      deadPointer: false,
      pointer,
      initialClassSlug: 'warrior',
    });
  });

  it('does not restore a fight- or armory-sourced pointer', () => {
    const fight = stored('fight', 'abc:1');
    expect(decidePlannerLoad(null, true, false, true, fight)).toEqual<PlannerLoadDecision>({
      codeParam: null,
      decoded: null,
      restored: false,
      deadPointer: false,
      pointer: fight,
      initialClassSlug: 'warrior',
    });
    const armory = stored('armory', 'us/normal/simfury');
    expect(decidePlannerLoad(null, true, false, true, armory)).toEqual<PlannerLoadDecision>({
      codeParam: null,
      decoded: null,
      restored: false,
      deadPointer: false,
      pointer: armory,
      initialClassSlug: 'warrior',
    });
  });

  it('does not restore when the URL is not bare, but the chip still shows the pointer', () => {
    const pointer = stored('code', GOOD_CODE);
    const decision = decidePlannerLoad(null, false, false, true, pointer);
    expect(decision).toEqual<PlannerLoadDecision>({
      codeParam: null,
      decoded: null,
      restored: false,
      deadPointer: false,
      pointer,
      initialClassSlug: 'warrior',
    });
  });

  it('does not restore when a record is present (/b/:id already has its own build)', () => {
    const pointer = stored('code', GOOD_CODE);
    const decision = decidePlannerLoad(null, true, true, true, pointer);
    expect(decision).toEqual<PlannerLoadDecision>({
      codeParam: null,
      decoded: null,
      restored: false,
      deadPointer: false,
      pointer,
      initialClassSlug: 'warrior',
    });
  });

  it('shows no pointer at all when the mount is not standalone (Top Gear\'s inline "add a build")', () => {
    const decision = decidePlannerLoad(null, true, false, false, stored('code', GOOD_CODE));
    expect(decision).toEqual<PlannerLoadDecision>({
      codeParam: null,
      decoded: null,
      restored: false,
      deadPointer: false,
      pointer: null,
      initialClassSlug: null,
    });
  });

  it('does not restore when nothing is stored', () => {
    const decision = decidePlannerLoad(null, true, false, true, null);
    expect(decision).toEqual<PlannerLoadDecision>({
      codeParam: null,
      decoded: null,
      restored: false,
      deadPointer: false,
      pointer: null,
      initialClassSlug: null,
    });
  });

  it('forgets a stored code that no longer decodes, with no error, no restore and no pointer', () => {
    const decision = decidePlannerLoad(null, true, false, true, stored('code', 'not an fs1 code'));
    expect(decision).toEqual<PlannerLoadDecision>({
      codeParam: null,
      decoded: null,
      restored: false,
      deadPointer: true,
      pointer: null,
      initialClassSlug: null,
    });
  });

  it('decodes the URL code exactly once, handing the result back rather than making the caller redo it', () => {
    const decision = decidePlannerLoad('FS1:1:warrior:orc:1:', true, false, true, null);
    expect(decision.decoded).toEqual(OTHER_DECODED);
  });

  it('uses the pointer classSlug as the initial class even for a non-restorable (armory) pointer', () => {
    const pointer = stored('armory', 'us/normal/simfury');
    const decision = decidePlannerLoad(null, true, false, true, pointer);
    expect(decision.initialClassSlug).toBe('warrior');
  });

  it('uses the URL code classSlug is irrelevant here -- initialClassSlug is null when the URL already names a code', () => {
    const decision = decidePlannerLoad(GOOD_CODE, true, false, true, null);
    expect(decision.initialClassSlug).toBeNull();
  });

  it('is null with no pointer at all, so the caller resolves the main async', () => {
    const decision = decidePlannerLoad(null, true, false, true, null);
    expect(decision.initialClassSlug).toBeNull();
  });

  it('is null when the mount is not standalone, matching the rest of the decision', () => {
    const decision = decidePlannerLoad(null, true, false, false, stored('armory', 'us/normal/simfury'));
    expect(decision.initialClassSlug).toBeNull();
  });
});

describe('labelForPlannerLoad', () => {
  const index = indexTalents(fixtureTalents as TalentFile);
  const order = GOOD_DECODED.ok ? orderFromRanks(index, GOOD_DECODED.build.treeRanks).order : [];

  // Fix round 1, Important: specLabel already names the class ("Fury Warrior"), so a title
  // and a derivable spec join as "<title> · <specLabel>", never "<title> · <className> ·
  // <specLabel>", and an untitled, spec-derivable load is `specLabel` alone, not
  // "<className> · <specLabel>".
  it('joins the title with the spec label when both are present', () => {
    expect(labelForPlannerLoad('My Fury Build', 'Warrior', index, order)).toBe(
      'My Fury Build · Fury Warrior',
    );
  });

  it('ignores a blank title, same as no title at all', () => {
    expect(labelForPlannerLoad('   ', 'Warrior', index, order)).toBe('Fury Warrior');
  });

  it('is the title alone when the spec is not derivable (no talent data yet)', () => {
    expect(labelForPlannerLoad('My Fury Build', 'Warrior', null, [])).toBe('My Fury Build');
  });

  it('is the spec label alone, with no title and a derivable spec', () => {
    expect(labelForPlannerLoad(undefined, 'Warrior', index, order)).toBe('Fury Warrior');
  });

  it('is the class name alone, with no title and no derivable spec', () => {
    expect(labelForPlannerLoad(undefined, 'Warrior', null, [])).toBe('Warrior');
  });
});

describe('recordPlannerCharacter', () => {
  it('writes a code-sourced pointer with the given label and class slug', () => {
    const storage = fakeStorage();
    recordPlannerCharacter('code', GOOD_CODE, 'Warrior · Fury Warrior', 'warrior', storage);
    expect(readCurrent(storage)).toMatchObject({
      source: 'code',
      ref: GOOD_CODE,
      label: 'Warrior · Fury Warrior',
      classSlug: 'warrior',
    });
  });

  it('writes a build-sourced pointer', () => {
    const storage = fakeStorage();
    recordPlannerCharacter('build', 'b1', 'My Build', 'mage', storage);
    expect(readCurrent(storage)).toMatchObject({ source: 'build', ref: 'b1', classSlug: 'mage' });
  });

  it('writes an addon-sourced pointer', () => {
    const storage = fakeStorage();
    recordPlannerCharacter('addon', GOOD_CODE, 'Warrior', 'warrior', storage);
    expect(readCurrent(storage)).toMatchObject({ source: 'addon', ref: GOOD_CODE });
  });
});

describe('plannerAddonCode', () => {
  it('is empty with no talent index', () => {
    const store = { talentIndex: null } as unknown as Parameters<typeof plannerAddonCode>[0];
    expect(plannerAddonCode(store)).toBe('');
  });

  it('builds the addon export string once talent data has loaded', () => {
    const index = indexTalents(fixtureTalents as TalentFile);
    const order = GOOD_DECODED.ok ? orderFromRanks(index, GOOD_DECODED.build.treeRanks).order : [];
    const store = {
      talentIndex: index,
      treeVersion: '1',
      classSlug: 'warrior',
      order,
      gear: {},
      itemIndex: new Map(),
    } as unknown as Parameters<typeof plannerAddonCode>[0];
    expect(plannerAddonCode(store)).toMatch(/^FSB1:1:warrior:/);
  });
});

describe('writePlannerPointer', () => {
  const index = indexTalents(fixtureTalents as TalentFile);
  const order = GOOD_DECODED.ok ? orderFromRanks(index, GOOD_DECODED.build.treeRanks).order : [];
  const loadedStore = {
    talentIndex: index,
    classRow: { id: 1, slug: 'warrior', name: 'Warrior' },
    order,
  } as unknown as PlannerStore;
  const emptyStore = { talentIndex: null } as unknown as PlannerStore;

  it('writes and returns the fresh pointer once talent data has loaded, standalone', () => {
    const storage = fakeStorage();
    const result = writePlannerPointer(loadedStore, true, 'code', GOOD_CODE, 'warrior', undefined, storage);
    expect(result).toMatchObject({ source: 'code', ref: GOOD_CODE, label: 'Fury Warrior' });
    expect(readCurrent(storage)).toEqual(result);
  });

  it('is a no-op inline (standalone false), leaving any existing pointer untouched', () => {
    const storage = fakeStorage();
    recordPlannerCharacter('build', 'b1', 'Old', 'mage', storage);
    const before = readCurrent(storage);
    const result = writePlannerPointer(loadedStore, false, 'code', GOOD_CODE, 'warrior', undefined, storage);
    expect(result).toEqual(before);
    expect(readCurrent(storage)).toEqual(before);
  });

  it('is a no-op before talent data has loaded', () => {
    const storage = fakeStorage();
    const result = writePlannerPointer(emptyStore, true, 'code', GOOD_CODE, 'warrior', undefined, storage);
    expect(result).toBeNull();
    expect(readCurrent(storage)).toBeNull();
  });
});

describe('unsavedPlannerHref', () => {
  it('is empty before talent data has loaded', () => {
    const store = { talentIndex: null } as unknown as PlannerStore;
    expect(unsavedPlannerHref(store)).toBe('');
  });

  it('is empty when the store cannot resolve a class against the loaded reference data', () => {
    const store = loadedPlannerStore({ classSlug: 'not-a-class' });
    expect(unsavedPlannerHref(store)).toBe('');
  });

  it('is empty when the store cannot resolve a race against the loaded reference data', () => {
    // `setReference` auto-repairs an illegal `raceSlug` to a legal one for the class
    // (`repairRaceForClass`), so an unresolvable race never survives a real load -- this
    // exercises `characterFromPlanner`'s own `raceRow === null` guard directly instead.
    const store = {
      talentIndex: indexTalents(fixtureTalents as TalentFile),
      talents: fixtureTalents as TalentFile,
      classes: fixtureClasses as ClassRow[],
      classRow: (fixtureClasses as ClassRow[])[0],
      raceRow: null,
    } as unknown as PlannerStore;
    expect(unsavedPlannerHref(store)).toBe('');
  });

  it('builds a /planner?code= link -- path and query only, no origin', () => {
    const store = loadedPlannerStore();
    store.addPoint(1001);
    const href = unsavedPlannerHref(store);
    expect(href.startsWith('/planner?code=')).toBe(true);
    expect(href).not.toContain('://');
  });

  it('round-trips through decodeFS1 back to the same talent order', () => {
    const store = loadedPlannerStore();
    store.addPoint(1001);
    store.addPoint(1001);
    const href = unsavedPlannerHref(store);
    const code = new URLSearchParams(href.slice(href.indexOf('?'))).get('code') ?? '';
    const decoded = decodeFS1(code);
    expect(decoded.ok).toBe(true);
    if (!decoded.ok) return;
    const index = indexTalents(fixtureTalents as TalentFile);
    expect(orderFromRanks(index, decoded.build.treeRanks).order).toEqual(store.order);
  });
});
