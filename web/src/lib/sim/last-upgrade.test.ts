// web/src/lib/sim/last-upgrade.test.ts
import { describe, expect, it } from 'vitest';
import { readLastUpgrade, upgradeKeyOf, writeLastUpgrade } from './last-upgrade';

function fakeStorage(): Storage {
  const map = new Map<string, string>();
  return {
    getItem: (key) => map.get(key) ?? null,
    setItem: (key, value) => void map.set(key, value),
    removeItem: (key) => void map.delete(key),
    clear: () => map.clear(),
    key: () => null,
    get length() {
      return map.size;
    },
  } as Storage;
}

describe('last-upgrade', () => {
  it('round-trips a written upgrade', () => {
    const storage = fakeStorage();
    writeLastUpgrade(
      {
        itemName: 'Bracers of X',
        sourceName: 'Blackfathom Deeps',
        gain: '+14 ± 3',
        savedAt: '2026-01-01T00:00:00Z',
        characterKey: 'armory:us/forever/thrall:warrior-fury',
      },
      storage,
    );
    expect(readLastUpgrade('armory:us/forever/thrall:warrior-fury', storage)).toEqual({
      itemName: 'Bracers of X',
      sourceName: 'Blackfathom Deeps',
      gain: '+14 ± 3',
      savedAt: '2026-01-01T00:00:00Z',
      characterKey: 'armory:us/forever/thrall:warrior-fury',
    });
  });

  it('reads null when nothing is stored', () => {
    expect(readLastUpgrade('k', fakeStorage())).toBeNull();
  });

  it('reads null for malformed JSON rather than throwing', () => {
    const storage = fakeStorage();
    storage.setItem('fs.lastDroptimizerUpgrade', '{not json');
    expect(readLastUpgrade('k', storage)).toBeNull();
  });

  it('reads null for a value missing a required field', () => {
    const storage = fakeStorage();
    storage.setItem('fs.lastDroptimizerUpgrade', JSON.stringify({ itemName: 'X' }));
    expect(readLastUpgrade('k', storage)).toBeNull();
  });

  it('never hands one character’s upgrade to another, or to a different spec of the same one', () => {
    const storage = fakeStorage();
    const upgrade = {
      itemName: 'Bracers of X',
      sourceName: 'BFD',
      gain: '+14 ± 3',
      savedAt: '2026-01-01T00:00:00Z',
    };
    writeLastUpgrade({ ...upgrade, characterKey: 'armory:us/forever/thrall:warrior-fury' }, storage);
    expect(readLastUpgrade('armory:us/forever/thrall:warrior-arms', storage)).toBeNull();
    expect(readLastUpgrade('manual:FS1-rogue:rogue-combat', storage)).toBeNull();
  });

  it('reads null for an upgrade stored before it named its character', () => {
    const storage = fakeStorage();
    storage.setItem(
      'fs.lastDroptimizerUpgrade',
      JSON.stringify({ itemName: 'X', sourceName: 'Y', gain: '+1', savedAt: '2026-01-01T00:00:00Z' }),
    );
    expect(readLastUpgrade('k', storage)).toBeNull();
  });

  it('keys an upgrade by the character’s source and spec', () => {
    expect(
      upgradeKeyOf({
        spec: 'rogue-combat',
        source: { kind: 'armory', ref: 'us/forever/x' },
      }),
    ).toBe('armory:us/forever/x:rogue-combat');
  });
});
