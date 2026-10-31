// web/src/lib/character-selector/restored.test.ts
// The selector's "Restored your last character" line and its Forget: shown only while the
// pointer names the character this tab's session started with, never for the account's main,
// and Forget clears the pointer and tells every island.
import { describe, expect, it, vi } from 'vitest';
import {
  CURRENT_CHARACTER_CHANGED,
  readCurrent,
  writeCurrent,
  type CurrentCharacter,
} from '../current-character';
import {
  forgetPointer,
  isRestoredPointer,
  pointerIdentity,
  readSessionStartPointer,
  SESSION_START_POINTER_KEY,
} from './restored';

const THRALL = 'armory:us/normal/thrall';

function memoryStorage(initial: Record<string, string> = {}): Storage {
  const data = new Map(Object.entries(initial));
  return {
    get length() {
      return data.size;
    },
    clear: () => data.clear(),
    getItem: (key) => data.get(key) ?? null,
    key: (index) => [...data.keys()][index] ?? null,
    removeItem: (key) => void data.delete(key),
    setItem: (key, value) => void data.set(key, value),
  };
}

function pointer(overrides: Partial<CurrentCharacter> = {}): CurrentCharacter {
  return {
    source: 'armory',
    ref: 'us/normal/thrall',
    label: 'Thrall · Fury Warrior',
    classSlug: 'warrior',
    savedAt: '2026-10-30T12:00:00Z',
    ...overrides,
  };
}

describe('readSessionStartPointer', () => {
  it('reads the identity the pre-paint script recorded, empty for no pointer', () => {
    expect(readSessionStartPointer(memoryStorage({ [SESSION_START_POINTER_KEY]: THRALL }))).toBe(THRALL);
    expect(readSessionStartPointer(memoryStorage({ [SESSION_START_POINTER_KEY]: '' }))).toBe('');
  });

  it('is null when nothing was recorded', () => {
    expect(readSessionStartPointer(memoryStorage())).toBeNull();
  });
});

describe('isRestoredPointer', () => {
  it('is true while the pointer names the character the session started with', () => {
    expect(isRestoredPointer(pointer(), THRALL, null)).toBe(true);
  });

  it('survives a rewrite of the same character (the simulator re-saves what it restored)', () => {
    expect(
      isRestoredPointer(
        pointer({ savedAt: '2026-10-31T13:00:00Z', label: 'Thrall · Arms Warrior' }),
        THRALL,
        null,
      ),
    ).toBe(true);
  });

  it('is false once another character is chosen in this session', () => {
    expect(isRestoredPointer(pointer({ ref: 'us/normal/jaina' }), THRALL, null)).toBe(false);
  });

  it('is false with no pointer, a session that started with none, or nothing recorded', () => {
    expect(isRestoredPointer(null, THRALL, null)).toBe(false);
    expect(isRestoredPointer(pointer(), '', null)).toBe(false);
    expect(isRestoredPointer(pointer(), null, null)).toBe(false);
  });

  it('is false for the account main: forgetting it would change nothing', () => {
    expect(isRestoredPointer(pointer(), THRALL, 'us/normal/thrall')).toBe(false);
    expect(isRestoredPointer(pointer(), THRALL, 'us/normal/jaina')).toBe(true);
  });

  it('holds for a pasted export, whatever the main', () => {
    const pasted = pointer({ source: 'addon', ref: 'FS1:x' });
    expect(isRestoredPointer(pasted, pointerIdentity(pasted), 'us/normal/thrall')).toBe(true);
  });
});

describe('forgetPointer', () => {
  it('clears the stored pointer and announces the change', () => {
    const storage = memoryStorage();
    writeCurrent(pointer(), storage);
    const dispatch = vi.fn();
    forgetPointer({ storage, dispatch });
    expect(readCurrent(storage)).toBeNull();
    expect(dispatch).toHaveBeenCalledTimes(1);
    expect((dispatch.mock.calls[0][0] as Event).type).toBe(CURRENT_CHARACTER_CHANGED);
  });
});
