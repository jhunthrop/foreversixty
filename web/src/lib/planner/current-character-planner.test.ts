// web/src/lib/planner/current-character-planner.test.ts
import { describe, expect, it } from 'vitest';
import fixtureClasses from '../../fixtures/planner/classes.json';
import fixtureCombos from '../../fixtures/planner/combos.json';
import fixtureRaces from '../../fixtures/planner/races.json';
import fixtureTalents from '../../fixtures/planner/talents/warrior.json';
import type { CurrentCharacter } from '../current-character';
import {
  decidePlannerLoad,
  followPointerInPlanner,
  isBarePlannerUrl,
  plannerAddonCode,
  unsavedPlannerHref,
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
  it('is true when only ?spec= is present, so the home page’s own entry can still restore', () => {
    expect(isBarePlannerUrl('?spec=hunter-beast-mastery')).toBe(true);
  });
  it('is false when ?talents= is present, so a stored pointer never wins the race against it', () => {
    expect(isBarePlannerUrl('?talents=503200000-0-0')).toBe(false);
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
      characterKey: null,
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
      characterKey: null,
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
      characterKey: null,
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
      characterKey: null,
    });
  });

  it('does not restore a fight- or armory-sourced pointer as a code; an armory pointer names its character', () => {
    const fight = stored('fight', 'abc:1');
    expect(decidePlannerLoad(null, true, false, true, fight)).toEqual<PlannerLoadDecision>({
      codeParam: null,
      decoded: null,
      restored: false,
      deadPointer: false,
      pointer: fight,
      initialClassSlug: 'warrior',
      characterKey: null,
    });
    const armory = stored('armory', 'us/normal/simfury');
    expect(decidePlannerLoad(null, true, false, true, armory)).toEqual<PlannerLoadDecision>({
      codeParam: null,
      decoded: null,
      restored: false,
      deadPointer: false,
      pointer: armory,
      initialClassSlug: 'warrior',
      characterKey: 'us/normal/simfury',
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
      characterKey: null,
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
      characterKey: null,
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
      characterKey: null,
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
      characterKey: null,
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
      characterKey: null,
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

describe('decidePlannerLoad with an account character', () => {
  const armory = stored('armory', 'us/pvp/bow-jackzon');

  it('opens a bare planner on the armory pointer’s character', () => {
    expect(decidePlannerLoad(null, true, false, true, armory).characterKey).toBe('us/pvp/bow-jackzon');
  });

  it('opens ?character= on that character, over any pointer', () => {
    expect(decidePlannerLoad(null, false, false, true, armory, 'us/pvp/pipsqueak').characterKey).toBe(
      'us/pvp/pipsqueak',
    );
  });

  it('never loads a character into a saved build, an inline planner or a ?class= tile', () => {
    expect(decidePlannerLoad(null, true, true, true, armory).characterKey).toBeNull();
    expect(decidePlannerLoad(null, true, false, false, armory).characterKey).toBeNull();
    expect(decidePlannerLoad(null, false, false, true, armory).characterKey).toBeNull();
  });

  it('treats ?character= as a named address, not a bare one', () => {
    expect(isBarePlannerUrl('?character=us%2Fpvp%2Fbow-jackzon')).toBe(false);
  });
});

describe('followPointerInPlanner', () => {
  it('imports a pasted export, naming its class so the planner can switch first', () => {
    expect(followPointerInPlanner(stored('addon', GOOD_CODE), 'mage')).toEqual({
      kind: 'import',
      classSlug: 'warrior',
      code: GOOD_CODE,
    });
    expect(followPointerInPlanner(stored('code', GOOD_CODE), 'warrior').kind).toBe('import');
  });

  it('does nothing for an export that no longer decodes, or for no pointer', () => {
    expect(followPointerInPlanner(stored('addon', 'FS1:broken'), 'mage')).toEqual({ kind: 'none' });
    expect(followPointerInPlanner(null, 'mage')).toEqual({ kind: 'none' });
  });

  it('loads an account character with its own build, whatever class the planner is on', () => {
    const hunter = { ...stored('armory', 'us/normal/bow'), classSlug: 'hunter' };
    expect(followPointerInPlanner(hunter, 'warrior')).toEqual({ kind: 'character', key: 'us/normal/bow' });
    expect(followPointerInPlanner(hunter, 'hunter')).toEqual({ kind: 'character', key: 'us/normal/bow' });
  });

  it('moves a fight pointer to its class and keeps the build when the class is the same', () => {
    const fight = { ...stored('fight', 'abc:1'), classSlug: 'hunter' };
    expect(followPointerInPlanner(fight, 'warrior')).toEqual({ kind: 'class', classSlug: 'hunter' });
    expect(followPointerInPlanner(fight, 'hunter')).toEqual({ kind: 'none' });
  });
});
