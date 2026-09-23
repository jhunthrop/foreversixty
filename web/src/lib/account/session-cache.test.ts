// web/src/lib/account/session-cache.test.ts
import { describe, expect, it } from 'vitest';
import type { Me } from './api';
import {
  SNAPSHOT_TTL_MS,
  clearSnapshot,
  readSnapshot,
  sameMe,
  sessionHinted,
  writeSnapshot,
} from './session-cache';

function memoryStorage(): Storage {
  const map = new Map<string, string>();
  return {
    get length() {
      return map.size;
    },
    clear: () => map.clear(),
    getItem: (k) => map.get(k) ?? null,
    key: (i) => [...map.keys()][i] ?? null,
    removeItem: (k) => void map.delete(k),
    setItem: (k, v) => void map.set(k, v),
  };
}

const ME: Me = {
  user: { id: 1, battletag: 'Fixture#1', email: null, role: 'user', anonymize: false },
  characters: [],
  guilds: [],
  entitlements: { server_sims: false, guild_tools: false, personal: { plan: null, status: null } },
} as unknown as Me;

describe('the session snapshot', () => {
  it('is only read while the session cookie is present', () => {
    const store = memoryStorage();
    writeSnapshot(ME, 1000, store);
    expect(readSnapshot(2000, store, 'fs_csrf=abc; other=1')).toEqual(ME);
    expect(readSnapshot(2000, store, 'other=1')).toBeNull();
  });

  it('expires after its TTL and clears on demand', () => {
    const store = memoryStorage();
    writeSnapshot(ME, 1000, store);
    expect(readSnapshot(1000 + SNAPSHOT_TTL_MS + 1, store, 'fs_csrf=abc')).toBeNull();
    writeSnapshot(ME, 1000, store);
    clearSnapshot(store);
    expect(readSnapshot(1000, store, 'fs_csrf=abc')).toBeNull();
  });

  it('survives garbage in the store', () => {
    const store = memoryStorage();
    store.setItem('fs.me', '{not json');
    expect(readSnapshot(1000, store, 'fs_csrf=abc')).toBeNull();
  });

  it('compares answers structurally', () => {
    expect(sameMe(ME, { ...ME })).toBe(true);
    expect(sameMe(ME, null)).toBe(false);
    expect(sessionHinted('fs_csrf=x')).toBe(true);
    expect(sessionHinted('fs_csrf=')).toBe(false);
  });
});
