// @vitest-environment jsdom
import { afterEach, describe, expect, it, vi } from 'vitest';
import { API_BASE_URL } from '../planner/config';
import { resetQueryCache, setQueryData, subscribe } from '../data/query';
import { meKey, type Me } from './api';
import { discordLinkStartUrl, unlinkDiscord } from './discord-link-api';

const me = (extra: Partial<Me> = {}): Me => ({
  user: { id: 1, battletag: null, email: null, role: 'user', anonymize: false },
  characters: [],
  guilds: [],
  discord_link_available: true,
  discord_link: { display_name: 'zulmara' },
  ...extra,
});

function cachedMe(): Me | null | undefined {
  let cached: Me | null | undefined;
  subscribe<Me | null>(meKey(API_BASE_URL), (state) => {
    cached = state.data;
  })();
  return cached;
}

afterEach(() => {
  vi.restoreAllMocks();
  resetQueryCache();
});

describe('discordLinkStartUrl', () => {
  it('is the API start route', () => {
    expect(discordLinkStartUrl()).toBe(`${API_BASE_URL}/v1/discord/link/start`);
  });
});

describe('unlinkDiscord', () => {
  it('patches the current cache entry, not the stale prop', async () => {
    vi.spyOn(globalThis, 'fetch').mockResolvedValue(new Response(null, { status: 204 }));
    setQueryData(meKey(API_BASE_URL), me({ main_character_key: 'newer' }));
    await unlinkDiscord(me());
    expect(cachedMe()?.discord_link).toBeNull();
    expect(cachedMe()?.main_character_key).toBe('newer');
  });

  it('leaves the cache alone when the call fails', async () => {
    vi.spyOn(globalThis, 'fetch').mockResolvedValue(new Response('{}', { status: 500 }));
    setQueryData(meKey(API_BASE_URL), me());
    await expect(unlinkDiscord(me())).rejects.toBeDefined();
    expect(cachedMe()?.discord_link).toEqual({ display_name: 'zulmara' });
  });
});
