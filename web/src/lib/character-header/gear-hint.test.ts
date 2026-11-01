import { describe, expect, it } from 'vitest';
import { HEADER_GEARED, HEADER_GEAR_HINT_KEY, rememberHeaderGear } from './gear-hint';
import { PLACEHOLDER_CHARACTER, PLACEHOLDER_GEARLESS_CHARACTER } from './placeholder';

function memoryStorage(): Pick<Storage, 'setItem'> & { values: Map<string, string> } {
  const values = new Map<string, string>();
  return { values, setItem: (key, value) => void values.set(key, value) };
}

describe('rememberHeaderGear', () => {
  it('stores the geared value for a character with gear on file', () => {
    const storage = memoryStorage();
    rememberHeaderGear(PLACEHOLDER_CHARACTER, storage);
    expect(storage.values.get(HEADER_GEAR_HINT_KEY)).toBe(HEADER_GEARED);
  });

  it('stores another value for a character with none', () => {
    const storage = memoryStorage();
    rememberHeaderGear(PLACEHOLDER_GEARLESS_CHARACTER, storage);
    expect(storage.values.get(HEADER_GEAR_HINT_KEY)).not.toBe(HEADER_GEARED);
  });

  it('swallows a storage that refuses writes', () => {
    const refusing = {
      setItem: () => {
        throw new Error('quota');
      },
    };
    expect(() => rememberHeaderGear(PLACEHOLDER_CHARACTER, refusing)).not.toThrow();
  });
});
