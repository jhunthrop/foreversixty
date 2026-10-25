// @vitest-environment jsdom
import { beforeEach, describe, expect, it, vi } from 'vitest';
import { onCurrentCharacterChange } from '../current-character-events';
import type { CurrentCharacter } from '../current-character';
import type { Me, MeCharacter } from './api';
import { pointerToHeal, settlePointerForSession } from './session-pointer';

const JACKZON: MeCharacter = {
  key: 'us/pvp/bow-jackzon',
  region: 'us',
  ruleset: 'pvp',
  name: 'Bow Jackzon',
  class: 'hunter',
  spec: 'Survival',
};
const PIPSQUEAK: MeCharacter = {
  key: 'us/pvp/pipsqueak',
  region: 'us',
  ruleset: 'pvp',
  name: 'Pipsqueak',
  class: 'priest',
};
const ME = { characters: [JACKZON, PIPSQUEAK], main_character_key: JACKZON.key } as Me;

const pointer = (over: Partial<CurrentCharacter>): CurrentCharacter => ({
  source: 'armory',
  ref: PIPSQUEAK.key,
  label: 'Pipsqueak · Priest',
  classSlug: 'priest',
  savedAt: '2026-10-21T00:00:00Z',
  ...over,
});

describe('pointerToHeal', () => {
  it('writes the main when there is no pointer', () => {
    expect(pointerToHeal(null, ME)).toMatchObject({
      source: 'armory',
      ref: JACKZON.key,
      label: 'Bow Jackzon · Survival Hunter',
    });
  });

  it('replaces a nameless pasted code, which names nobody', () => {
    expect(
      pointerToHeal(pointer({ source: 'code', ref: 'FS1:x', label: 'Hunter', classSlug: 'hunter' }), ME)?.ref,
    ).toBe(JACKZON.key);
  });

  it('replaces an armory key the account does not hold', () => {
    expect(pointerToHeal(pointer({ ref: 'us/pvp/someone-else' }), ME)?.ref).toBe(JACKZON.key);
  });

  it('leaves a choice alone: an account character, a named export, a fight or a build', () => {
    expect(pointerToHeal(pointer({}), ME)).toBeNull();
    expect(
      pointerToHeal(
        pointer({ source: 'addon', ref: 'FS1:x', label: 'Stranger · Mage', classSlug: 'mage' }),
        ME,
      ),
    ).toBeNull();
    expect(pointerToHeal(pointer({ source: 'fight', ref: 'abc:1' }), ME)).toBeNull();
    expect(pointerToHeal(pointer({ source: 'build', ref: 'b1' }), ME)).toBeNull();
  });

  it('writes nothing for an account with no characters', () => {
    expect(pointerToHeal(null, { characters: [] } as unknown as Me)).toBeNull();
  });
});

describe('settlePointerForSession', () => {
  beforeEach(() => {
    window.localStorage.clear();
    window.sessionStorage.clear();
  });

  it('announces the heal as a page load, never as a player choosing', () => {
    const listener = vi.fn();
    const stop = onCurrentCharacterChange(listener);
    settlePointerForSession(ME);
    stop();
    expect(listener.mock.calls).toEqual([[{ fromPageLoad: true }]]);
  });

  it('runs once per browser session', () => {
    settlePointerForSession(ME);
    const listener = vi.fn();
    const stop = onCurrentCharacterChange(listener);
    settlePointerForSession(ME);
    stop();
    expect(listener).not.toHaveBeenCalled();
  });
});
