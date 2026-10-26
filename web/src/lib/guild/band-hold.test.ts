import { describe, expect, it } from 'vitest';
import { bandHoldKey, forgetBandHold, readBandHold, rememberBandHold } from './band-hold';

function memoryStorage(): Pick<Storage, 'getItem' | 'setItem' | 'removeItem'> {
  const map = new Map<string, string>();
  return {
    getItem: (key) => map.get(key) ?? null,
    setItem: (key, value) => void map.set(key, value),
    removeItem: (key) => void map.delete(key),
  };
}

const KEY = bandHoldKey('us', 'pvp', 'olympus-xxvii');

describe('the guild band hold', () => {
  it('keys the hold by guild', () => {
    expect(KEY).toBe('fs.guildBand.us.pvp.olympus-xxvii');
  });

  it('holds nothing for a viewer never seen verified here', () => {
    expect(readBandHold(memoryStorage(), KEY, 'desktop')).toBeNull();
    expect(readBandHold(null, KEY, 'desktop')).toBeNull();
  });

  it('reads back the height each layout last drew', () => {
    const storage = memoryStorage();
    rememberBandHold(storage, KEY, 'desktop', 131.6);
    rememberBandHold(storage, KEY, 'phone', 208);
    expect(readBandHold(storage, KEY, 'desktop')).toBe(132);
    expect(readBandHold(storage, KEY, 'phone')).toBe(208);
  });

  it('keeps the smaller of the stored and the measured height, so a transient tall band never sticks', () => {
    const storage = memoryStorage();
    rememberBandHold(storage, KEY, 'desktop', 132);
    rememberBandHold(storage, KEY, 'desktop', 160);
    expect(readBandHold(storage, KEY, 'desktop')).toBe(132);
    rememberBandHold(storage, KEY, 'desktop', 120);
    expect(readBandHold(storage, KEY, 'desktop')).toBe(120);
  });

  it('replaces a stored value it does not trust', () => {
    const storage = memoryStorage();
    storage.setItem(KEY, '{"desktop":-4}');
    rememberBandHold(storage, KEY, 'desktop', 132);
    expect(readBandHold(storage, KEY, 'desktop')).toBe(132);
  });

  it('forgets the hold when the viewer is no longer a verified member', () => {
    const storage = memoryStorage();
    rememberBandHold(storage, KEY, 'desktop', 132);
    forgetBandHold(storage, KEY);
    expect(readBandHold(storage, KEY, 'desktop')).toBeNull();
  });

  it('trusts no nonsense and survives a blocked store', () => {
    const storage = memoryStorage();
    storage.setItem(KEY, '{"desktop":-4,"phone":99999}');
    expect(readBandHold(storage, KEY, 'desktop')).toBeNull();
    expect(readBandHold(storage, KEY, 'phone')).toBeNull();
    storage.setItem(KEY, 'not json');
    expect(readBandHold(storage, KEY, 'desktop')).toBeNull();
    const blocked = {
      getItem: (): string | null => {
        throw new Error('blocked');
      },
      setItem: (): void => {
        throw new Error('blocked');
      },
      removeItem: (): void => {
        throw new Error('blocked');
      },
    };
    expect(readBandHold(blocked, KEY, 'desktop')).toBeNull();
    expect(() => rememberBandHold(blocked, KEY, 'desktop', 132)).not.toThrow();
    expect(() => forgetBandHold(blocked, KEY)).not.toThrow();
  });
});
