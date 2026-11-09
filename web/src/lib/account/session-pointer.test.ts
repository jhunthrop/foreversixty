// @vitest-environment jsdom
import { beforeEach, describe, expect, it, vi } from 'vitest';
import { onCurrentCharacterChange } from '../current-character-events';
import { readCurrent, writeCurrent, type CurrentCharacter } from '../current-character';
import {
  pointerIdentity,
  SESSION_FRESH_KEY,
  SESSION_START_POINTER_KEY,
} from '../character-selector/restored';
import type { Me, MeCharacter } from './api';
import { pointerForFreshSession, pointerToHeal, settlePointerForSession } from './session-pointer';

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

  it('announces the heal tagged as a session heal, not as a player choosing', () => {
    const listener = vi.fn();
    const stop = onCurrentCharacterChange(listener);
    settlePointerForSession(ME);
    stop();
    expect(listener.mock.calls).toEqual([[{ fromPageLoad: false, fromSessionHeal: true }]]);
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

const AT = (hoursAgo: number): string =>
  new Date(Date.parse('2026-11-09T12:00:00Z') - hoursAgo * 3_600_000).toISOString();
const played = (character: MeCharacter, hoursAgo: number): MeCharacter => ({
  ...character,
  build: { source: 'addon', captured_at: AT(hoursAgo) },
});
const JACKZON_POINTER = pointer({
  ref: JACKZON.key,
  label: 'Bow Jackzon · Survival Hunter',
  classSlug: 'hunter',
});
const NAMELESS_PASTE = pointer({ source: 'code', ref: 'FS1:x', label: 'Hunter', classSlug: 'hunter' });

describe('pointerForFreshSession', () => {
  const jackzon = played(JACKZON, 48);
  const pipsqueak = played(PIPSQUEAK, 0.3);
  const me = (characters: MeCharacter[]): Me => ({ characters, main_character_key: JACKZON.key }) as Me;

  it('gives the restored pointer way to the character played since', () => {
    expect(pointerForFreshSession(JACKZON_POINTER, me([jackzon, pipsqueak]))).toMatchObject({
      source: 'armory',
      ref: PIPSQUEAK.key,
    });
  });

  it('lets the restored pointer stand when it names the character played last', () => {
    expect(pointerForFreshSession(pointer({}), me([jackzon, pipsqueak]))).toBeNull();
  });

  it('lets the restored pointer stand on a tie with the character it names', () => {
    expect(pointerForFreshSession(pointer({}), me([played(JACKZON, 0.3), pipsqueak]))).toBeNull();
    expect(pointerForFreshSession(JACKZON_POINTER, me([played(JACKZON, 0.3), pipsqueak]))).toBeNull();
  });

  it('gives way when the character it names has no data and another does', () => {
    expect(pointerForFreshSession(pointer({}), me([jackzon, PIPSQUEAK]))?.ref).toBe(JACKZON.key);
  });

  it('stands when no character has data to judge by', () => {
    expect(pointerForFreshSession(pointer({}), me([JACKZON, PIPSQUEAK]))).toBeNull();
  });

  it('replaces a nameless paste with the character played last, not the main', () => {
    expect(pointerForFreshSession(NAMELESS_PASTE, me([jackzon, pipsqueak]))?.ref).toBe(PIPSQUEAK.key);
  });

  it('still replaces a nameless paste with the main when nobody has data', () => {
    expect(pointerForFreshSession(NAMELESS_PASTE, me([JACKZON, PIPSQUEAK]))?.ref).toBe(JACKZON.key);
  });
});

describe('settlePointerForSession in a fresh browser session', () => {
  const me = {
    characters: [played(JACKZON, 48), played(PIPSQUEAK, 0.3)],
    main_character_key: JACKZON.key,
  } as Me;

  beforeEach(() => {
    window.localStorage.clear();
    window.sessionStorage.clear();
  });

  function startSession(stored: CurrentCharacter, fresh: boolean): void {
    writeCurrent(stored);
    window.sessionStorage.setItem(SESSION_START_POINTER_KEY, pointerIdentity(stored));
    if (fresh) window.sessionStorage.setItem(SESSION_FRESH_KEY, '1');
  }

  it('moves a restored pointer to the character played since', () => {
    startSession(JACKZON_POINTER, true);
    settlePointerForSession(me);
    expect(readCurrent()?.ref).toBe(PIPSQUEAK.key);
  });

  it('leaves the pointer alone in a session that is not fresh', () => {
    startSession(JACKZON_POINTER, false);
    settlePointerForSession(me);
    expect(readCurrent()?.ref).toBe(JACKZON.key);
  });

  it('never overrides a choice made in this session', () => {
    startSession(pointer({ ref: 'us/pvp/earlier', label: 'Earlier · Mage', classSlug: 'mage' }), true);
    writeCurrent(JACKZON_POINTER);
    settlePointerForSession(me);
    expect(readCurrent()?.ref).toBe(JACKZON.key);
  });
});
