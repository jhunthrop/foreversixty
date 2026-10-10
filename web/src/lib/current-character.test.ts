// web/src/lib/current-character.test.ts
// @vitest-environment jsdom
import { describe, expect, it, vi } from 'vitest';
import {
  CURRENT_CHARACTER_CHANGED,
  announcePageLoadedCharacter,
  clearCurrent,
  onCurrentCharacterChange,
  plannerHrefFor,
  readCurrent,
  simHrefFor,
  writeCurrent,
  type CurrentCharacter,
} from './current-character';

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

const sample: CurrentCharacter = {
  source: 'addon',
  ref: 'FS1:1:warrior:orc:0/0/0:',
  label: 'Simfury · Fury Warrior',
  classSlug: 'warrior',
  savedAt: '2026-09-21T00:00:00.000Z',
};

describe('readCurrent / writeCurrent / clearCurrent', () => {
  it('round-trips a written value', () => {
    const storage = fakeStorage();
    writeCurrent(sample, storage);
    expect(readCurrent(storage)).toEqual(sample);
  });

  it('is null when nothing has been written', () => {
    expect(readCurrent(fakeStorage())).toBeNull();
  });

  it('is null for malformed JSON rather than throwing', () => {
    const storage = fakeStorage();
    storage.setItem('fs.currentCharacter', '{not json');
    expect(readCurrent(storage)).toBeNull();
  });

  it('refuses a write past the 16 KB cap', () => {
    const storage = fakeStorage();
    const huge: CurrentCharacter = { ...sample, ref: 'x'.repeat(20_000) };
    writeCurrent(huge, storage);
    expect(readCurrent(storage)).toBeNull();
  });

  it('clear removes the stored value', () => {
    const storage = fakeStorage();
    writeCurrent(sample, storage);
    clearCurrent(storage);
    expect(readCurrent(storage)).toBeNull();
  });

  it('never throws when storage.getItem throws (private browsing)', () => {
    const storage = fakeStorage();
    storage.getItem = () => {
      throw new Error('blocked');
    };
    expect(readCurrent(storage)).toBeNull();
  });

  it('never throws when storage.setItem throws (private browsing / quota)', () => {
    const storage = fakeStorage();
    storage.setItem = () => {
      throw new Error('blocked');
    };
    expect(() => writeCurrent(sample, storage)).not.toThrow();
  });

  it('is null for a source outside the known enum', () => {
    const storage = fakeStorage();
    storage.setItem('fs.currentCharacter', JSON.stringify({ ...sample, source: 'manual' }));
    expect(readCurrent(storage)).toBeNull();
  });

  it('is null when a field has the wrong type', () => {
    const storage = fakeStorage();
    storage.setItem('fs.currentCharacter', JSON.stringify({ ...sample, ref: 5 }));
    expect(readCurrent(storage)).toBeNull();
  });

  it('is null when a required field is missing', () => {
    const storage = fakeStorage();
    const { label, ...withoutLabel } = sample;
    storage.setItem('fs.currentCharacter', JSON.stringify(withoutLabel));
    expect(readCurrent(storage)).toBeNull();
  });
});

describe('plannerHrefFor', () => {
  it('is a ?code= link for an addon-sourced pointer', () => {
    expect(plannerHrefFor(sample)).toBe(`/planner?code=${encodeURIComponent(sample.ref)}`);
  });

  it('is a ?code= link for a code-sourced pointer', () => {
    expect(plannerHrefFor({ ...sample, source: 'code' })).toBe(
      `/planner?code=${encodeURIComponent(sample.ref)}`,
    );
  });

  it('is the saved-build permalink for a build-sourced pointer', () => {
    expect(plannerHrefFor({ ...sample, source: 'build', ref: 'b1' })).toBe('/b/b1');
  });

  it('is the honest class-only fallback for a fight-sourced pointer', () => {
    expect(plannerHrefFor({ ...sample, source: 'fight', ref: 'abc:1' })).toBe('/planner?class=warrior');
  });

  it('is the honest class-only fallback for an armory-sourced pointer', () => {
    expect(plannerHrefFor({ ...sample, source: 'armory', ref: 'us/normal/simfury' })).toBe(
      '/planner?class=warrior',
    );
  });
});

describe('simHrefFor', () => {
  it('is a ?code= link on the bare /sim path for an addon-sourced pointer', () => {
    expect(simHrefFor(sample)).toBe(`/sim?code=${encodeURIComponent(sample.ref)}`);
  });

  it('carries ?source=&ref= for a build-sourced pointer', () => {
    expect(simHrefFor({ ...sample, source: 'build', ref: 'b1' })).toBe('/sim?source=build&ref=b1');
  });

  it('targets the named tab’s own path', () => {
    expect(simHrefFor({ ...sample, source: 'build', ref: 'b1' }, 'drops')).toBe(
      '/sim/drops?source=build&ref=b1',
    );
  });
});

describe('the pre-paint pointer attribute', () => {
  it('follows writeCurrent and clearCurrent on the document element', () => {
    delete document.documentElement.dataset.pointer;
    writeCurrent({
      source: 'code',
      ref: 'FS1:x',
      label: 'Simfury · Fury Warrior',
      classSlug: 'warrior',
      savedAt: '2026-09-25T00:00:00Z',
    });
    expect(document.documentElement.dataset.pointer).toBe('1');
    clearCurrent();
    expect(document.documentElement.dataset.pointer).toBeUndefined();
  });
});

describe('onCurrentCharacterChange', () => {
  it('calls the listener on every pointer event until unsubscribed', () => {
    const listener = vi.fn();
    const stop = onCurrentCharacterChange(listener);
    window.dispatchEvent(new Event(CURRENT_CHARACTER_CHANGED));
    window.dispatchEvent(new Event(CURRENT_CHARACTER_CHANGED));
    expect(listener).toHaveBeenCalledTimes(2);
    stop();
    window.dispatchEvent(new Event(CURRENT_CHARACTER_CHANGED));
    expect(listener).toHaveBeenCalledTimes(2);
  });

  it('tells a page-load write apart from a selector write', () => {
    const listener = vi.fn();
    const stop = onCurrentCharacterChange(listener);
    window.dispatchEvent(new Event(CURRENT_CHARACTER_CHANGED));
    announcePageLoadedCharacter();
    expect(listener.mock.calls).toEqual([[{ fromPageLoad: false }], [{ fromPageLoad: true }]]);
    stop();
  });

  it('also follows a pointer change made in another tab', () => {
    const listener = vi.fn();
    const stop = onCurrentCharacterChange(listener);
    window.dispatchEvent(new StorageEvent('storage', { key: 'fs.currentCharacter' }));
    window.dispatchEvent(new StorageEvent('storage', { key: 'something-else' }));
    expect(listener).toHaveBeenCalledTimes(1);
    stop();
  });

  it('is a no-op without a window', () => {
    vi.stubGlobal('window', undefined);
    try {
      expect(onCurrentCharacterChange(() => undefined)()).toBeUndefined();
    } finally {
      vi.unstubAllGlobals();
    }
  });
});
