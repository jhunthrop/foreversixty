// web/src/lib/current-character-bar.test.ts
import { describe, expect, it } from 'vitest';
import { resolveSpineClassSlug, spineDoorsFor } from './current-character-bar';
import type { CurrentCharacter } from './current-character';
import type { MeCharacter } from './account/api';

const ARMORY: CurrentCharacter = {
  source: 'armory',
  ref: 'us/normal/simfury',
  label: 'Simfury · Fury Warrior',
  classSlug: 'warrior',
  savedAt: '2026-09-25T00:00:00.000Z',
};

const CODE: CurrentCharacter = {
  source: 'code',
  ref: 'FS1:1:warrior:orc:0/0/0:',
  label: 'Fury Warrior',
  classSlug: 'warrior',
  savedAt: '2026-09-25T00:00:00.000Z',
};

const MAIN: MeCharacter = {
  key: 'us/normal/roland',
  region: 'us',
  ruleset: 'normal',
  name: 'Roland',
  class: 'Mage',
};
const POINTER_CHAR: MeCharacter = {
  key: 'us/normal/simfury',
  region: 'us',
  ruleset: 'normal',
  name: 'Simfury',
  class: 'Warrior',
};

describe('resolveSpineClassSlug', () => {
  it('prefers the resolved pointer character', () => {
    expect(resolveSpineClassSlug(POINTER_CHAR, ARMORY, MAIN)).toBe('warrior');
  });

  it('falls back to the raw pointer classSlug when it does not resolve to a MeCharacter', () => {
    expect(resolveSpineClassSlug(null, CODE, MAIN)).toBe('warrior');
  });

  it('falls back to the main character when there is no pointer at all', () => {
    expect(resolveSpineClassSlug(null, null, MAIN)).toBe('mage');
  });

  it('is null when nothing is known', () => {
    expect(resolveSpineClassSlug(null, null, null)).toBeNull();
  });

  it('is null when the main character has no class on file yet', () => {
    const noClass: MeCharacter = { key: 'us/normal/x', region: 'us', ruleset: 'normal', name: 'X' };
    expect(resolveSpineClassSlug(null, null, noClass)).toBeNull();
  });
});

describe('spineDoorsFor', () => {
  it('builds all four doors, Plan/Sim carrying an armory pointer, Rankings pre-filtered', () => {
    const doors = spineDoorsFor('warrior', ARMORY, 'us', 'normal');
    expect(doors.map((d) => d.id)).toEqual(['plan', 'sim', 'logs', 'rankings']);
    expect(doors[0]).toMatchObject({ label: 'Plan', testid: 'current-character-bar-plan' });
    expect(doors[0].href).toBe('/planner?class=warrior');
    expect(doors[1].href).toContain('/sim?source=armory&ref=us%2Fnormal%2Fsimfury');
    expect(doors[2]).toMatchObject({ label: 'Logs', href: '/logs' });
    expect(doors[3].label).toBe('Rankings for Warrior');
    expect(doors[3].href).toBe('/rankings?class=warrior&ruleset=normal');
  });

  it('carries a code/addon pointer through Plan and Sim via their own FS1 hrefs', () => {
    const doors = spineDoorsFor('warrior', CODE, null, null);
    expect(doors[0].href).toBe(`/planner?code=${encodeURIComponent(CODE.ref)}`);
    expect(doors[1].href).toContain(`code=${encodeURIComponent(CODE.ref)}`);
    // No region/ruleset known for a code pointer: Rankings filters by class alone.
    expect(doors[3].href).toBe('/rankings?class=warrior');
  });

  it('falls back to bare Plan/Sim links with no pointer at all (main-fallback case)', () => {
    const doors = spineDoorsFor('mage', null, 'us', 'normal');
    expect(doors[0].href).toBe('/planner?class=mage');
    expect(doors[1].href).toBe('/sim');
    expect(doors[3].href).toBe('/rankings?class=mage&ruleset=normal');
  });

  it('omits the Rankings door entirely when no class is known', () => {
    const doors = spineDoorsFor(null, null, null, null);
    expect(doors.map((d) => d.id)).toEqual(['plan', 'sim', 'logs']);
  });
});
