// @vitest-environment jsdom
import { describe, expect, it, vi } from 'vitest';
import { recordCurrentCharacter } from './current-character-bridge';
import { readCurrent } from '../current-character';
import type { SimCharacter } from './character';

const FURY = { name: 'Simfury', spec: 'warrior-fury', class_slug: 'warrior' } as unknown as SimCharacter;

function memoryStorage(): Storage & { writes: number } {
  const map = new Map<string, string>();
  return {
    writes: 0,
    get length() {
      return map.size;
    },
    clear: () => map.clear(),
    getItem: (key: string) => map.get(key) ?? null,
    key: (index: number) => [...map.keys()][index] ?? null,
    removeItem: (key: string) => void map.delete(key),
    setItem(key: string, value: string) {
      this.writes += 1;
      map.set(key, value);
    },
  };
}

describe('recordCurrentCharacter', () => {
  it('writes the pointer for a character it does not hold yet', () => {
    const storage = memoryStorage();
    recordCurrentCharacter(FURY, 'addon', 'FS1:a', storage);
    expect(readCurrent(storage)).toMatchObject({ source: 'addon', ref: 'FS1:a', classSlug: 'warrior' });
  });

  it('leaves the pointer alone when it already names this character, so a restore writes nothing', () => {
    // A restore that wrote the pointer again reached every other open tab as a storage change,
    // which loaded the same character there and wrote it back: two tabs passed it to and fro, and
    // a Forget in one landed between their writes and was undone (current-character.spec.ts).
    const storage = memoryStorage();
    const announced = vi.fn();
    window.addEventListener('fs:current-character', announced);
    recordCurrentCharacter(FURY, 'addon', 'FS1:a', storage);
    recordCurrentCharacter(FURY, 'addon', 'FS1:a', storage);
    window.removeEventListener('fs:current-character', announced);
    expect(storage.writes).toBe(1);
    expect(announced).toHaveBeenCalledTimes(1);
  });

  it('writes again for another character', () => {
    const storage = memoryStorage();
    recordCurrentCharacter(FURY, 'addon', 'FS1:a', storage);
    recordCurrentCharacter(FURY, 'addon', 'FS1:b', storage);
    expect(storage.writes).toBe(2);
    expect(readCurrent(storage)?.ref).toBe('FS1:b');
  });
});
