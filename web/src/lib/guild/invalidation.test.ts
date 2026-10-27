// web/src/lib/guild/invalidation.test.ts
// @vitest-environment jsdom
// The G.4 rule: every guild mutation drops the keys it changes (`/v1/me` and everything
// cached under that guild), so the next read asks the network. Also: the guild reads
// revalidate on every mount, and Settings asks again when the window regains focus.
import { afterEach, beforeEach, describe, expect, it, vi } from 'vitest';
import { fetchMeFresh, meKey } from '../account/api';
import {
  forgetPrivate,
  invalidate,
  onWindowFocus,
  query,
  resetQueryCache,
  setQueryData,
} from '../data/query';
import {
  acceptInvite,
  approveAllRoster,
  approveCharacter,
  awardLoot,
  clearCachedBotRefusal,
  claimGuild,
  confirmClaim,
  contestClaim,
  deleteGuildCrest,
  fetchGuildHome,
  fetchGuildSettings,
  leaveGuild,
  putGuildCrest,
  refetchGuildSettings,
  releaseClaim,
  removeCharacter,
  rotateInvite,
  unawardLoot,
  updateConsent,
  updateGuildSettings,
} from './api';

const API = 'https://api.foreversixty.test';
type GlobalFetch = (...args: Parameters<typeof fetch>) => Promise<Response>;

function envelope(data: unknown, status = 200): Response {
  return new Response(JSON.stringify({ ok: status < 400, data, error: null, request_id: 'r' }), {
    status,
    headers: { 'content-type': 'application/json' },
  });
}

const GUILD = { id: 7, region: 'us', ruleset: 'hardcore', name: 'Lifecycle Test' };
const OTHER = 70;

/** Keys a guild mutation must drop, and the neighbours it must leave alone. */
const droppedKeys = [
  `${API}/v1/me`,
  `${API}/v1/guilds/7/home`,
  `${API}/v1/guilds/7/home?cursor=abc`,
  `${API}/v1/guilds/7/settings`,
  `${API}/v1/guilds/7/roster`,
  `${API}/v1/guilds/7/loot?encounter=3`,
  `${API}/v1/guilds/7/readiness`,
];
const keptKeys = [`${API}/v1/guilds/${OTHER}/home`, `${API}/v1/guilds/${OTHER}/settings`];

async function seed(): Promise<void> {
  for (const key of [...droppedKeys, ...keptKeys]) {
    await query(key, async () => ({ seeded: key }), { scope: 'private', ttlMs: 60_000 });
  }
}

/** Whether a read of `key` is served from the cache (true) or asks the network. */
async function served(key: string): Promise<boolean> {
  let asked = false;
  await query(
    key,
    async () => {
      asked = true;
      return { reloaded: key };
    },
    { scope: 'private', ttlMs: 60_000 },
  );
  return !asked;
}

const MUTATIONS: [string, () => Promise<unknown>, unknown][] = [
  ['claimGuild', () => claimGuild(7, API), { state: 'claimed' }],
  ['confirmClaim', () => confirmClaim(7, API), { state: 'claimed' }],
  ['releaseClaim', () => releaseClaim(7, API), { state: 'unclaimed' }],
  ['contestClaim', () => contestClaim(7, API), { state: 'contested' }],
  ['rotateInvite', () => rotateInvite(7, API), { invite_path: '/guild/invite/x' }],
  ['approveCharacter', () => approveCharacter(7, 'us', 'hardcore', 'ashenveil', API), { status: 'approved' }],
  ['approveAllRoster', () => approveAllRoster(7, API), { approved: [] }],
  ['removeCharacter', () => removeCharacter(7, 'us', 'hardcore', 'ashenveil', API), { status: 'removed' }],
  ['updateConsent', () => updateConsent(7, 'gear', API), { consent: 'gear' }],
  ['leaveGuild', () => leaveGuild(7, API), { status: 'left' }],
  ['updateGuildSettings', () => updateGuildSettings(7, { default_visibility: 'guild' }, API), {}],
  ['acceptInvite', () => acceptInvite('tok', API), { guild: GUILD, rank: 'member' }],
  [
    'awardLoot',
    () => awardLoot(7, { encounter_id: 3, item_id: 9, character_key: 'us/hardcore/a' }, API),
    { id: 1 },
  ],
  ['unawardLoot', () => unawardLoot(7, 'a1', API), { status: 'removed' }],
  ['putGuildCrest', () => putGuildCrest(7, new File(['x'], 'c.png'), API), { crest_url: '/c.png' }],
  ['deleteGuildCrest', () => deleteGuildCrest(7, API), null],
];

beforeEach(() => {
  resetQueryCache();
});
afterEach(() => {
  forgetPrivate();
  invalidate('');
  vi.unstubAllGlobals();
});

describe('guild mutations drop the keys they change', () => {
  it.each(MUTATIONS)(
    '%s drops /v1/me and the guild reads, and only that guild',
    async (_name, run, reply) => {
      vi.stubGlobal(
        'fetch',
        vi.fn<GlobalFetch>(async () =>
          reply === null ? new Response(null, { status: 204 }) : envelope(reply),
        ),
      );
      await seed();
      for (const key of [...droppedKeys, ...keptKeys]) expect(await served(key)).toBe(true);

      await run();

      for (const key of droppedKeys) expect(await served(key), key).toBe(false);
      for (const key of keptKeys) expect(await served(key), key).toBe(true);
    },
  );

  it('a failed mutation leaves the cache alone', async () => {
    vi.stubGlobal(
      'fetch',
      vi.fn<GlobalFetch>(async () => envelope(null, 500)),
    );
    await seed();
    await expect(leaveGuild(7, API)).rejects.toBeTruthy();
    for (const key of droppedKeys) expect(await served(key), key).toBe(true);
  });
});

describe('guild reads revalidate on every read', () => {
  it('paints the cached home at once, asks again, and hands over a changed answer', async () => {
    const stale = { guild: GUILD, claim: { state: 'unclaimed', frozen: false }, reports: [], roster: [] };
    const fresh = { ...stale, claim: { state: 'claimed', frozen: false } };
    setQueryData(`${API}/v1/guilds/7/home`, stale);
    const upstream = vi.fn<GlobalFetch>(async () => envelope(fresh));
    vi.stubGlobal('fetch', upstream);
    const onFresh = vi.fn();

    const first = await fetchGuildHome(7, undefined, API, onFresh);
    expect(first.claim.state).toBe('unclaimed'); // the cached answer, instantly
    await vi.waitFor(() => expect(onFresh).toHaveBeenCalledTimes(1));
    expect(onFresh.mock.calls[0][0].claim.state).toBe('claimed');
    expect(upstream).toHaveBeenCalledTimes(1);
  });

  it('asks again on a second mount in the same page, and stays quiet when nothing changed', async () => {
    const home = { guild: GUILD, claim: { state: 'claimed', frozen: false }, reports: [], roster: [] };
    const upstream = vi.fn<GlobalFetch>(async () => envelope(home));
    vi.stubGlobal('fetch', upstream);
    const onFresh = vi.fn();

    await fetchGuildHome(7, undefined, API, onFresh); // first mount: the network
    await fetchGuildHome(7, undefined, API, onFresh); // second mount: cached, then a background ask
    await vi.waitFor(() => expect(upstream).toHaveBeenCalledTimes(2));
    expect(onFresh).not.toHaveBeenCalled();
  });

  it('settings revalidate on mount too', async () => {
    const stale = { default_visibility: 'guild', plan: null };
    const fresh = { default_visibility: 'guild', plan: { name: 'guild' } };
    setQueryData(`${API}/v1/guilds/7/settings`, stale);
    vi.stubGlobal(
      'fetch',
      vi.fn<GlobalFetch>(async () => envelope(fresh)),
    );
    const onFresh = vi.fn();
    await fetchGuildSettings(7, API, onFresh);
    await vi.waitFor(() => expect(onFresh).toHaveBeenCalledWith(fresh));
  });

  it('an answer older than a minute is not painted', async () => {
    vi.useFakeTimers({ toFake: ['Date'] });
    try {
      setQueryData(`${API}/v1/guilds/7/settings`, { old: true });
      vi.setSystemTime(Date.now() + 61_000);
      const upstream = vi.fn<GlobalFetch>(async () => envelope({ old: false }));
      vi.stubGlobal('fetch', upstream);
      const result = await fetchGuildSettings(7, API);
      expect(result).toEqual({ old: false });
    } finally {
      vi.useRealTimers();
    }
  });

  it('/v1/me read by a guild page revalidates and reports a join', async () => {
    const before = { user: { id: 1 }, characters: [], guilds: [] };
    const after = { ...before, guilds: [{ ...GUILD, rank: 'member', verified: true }] };
    setQueryData(meKey(API), before);
    vi.stubGlobal(
      'fetch',
      vi.fn<GlobalFetch>(async () => envelope(after)),
    );
    const onFresh = vi.fn();
    expect(await fetchMeFresh(onFresh, API)).toEqual(before);
    await vi.waitFor(() => expect(onFresh).toHaveBeenCalledWith(after));
  });
});

describe('the settings read on window focus', () => {
  it('refetches whatever the cache holds and hands over the answer', async () => {
    setQueryData(`${API}/v1/guilds/7/settings`, { plan: null });
    const upstream = vi.fn<GlobalFetch>(async () => envelope({ plan: 'guild' }));
    vi.stubGlobal('fetch', upstream);
    const onData = vi.fn();

    const off = onWindowFocus(() => refetchGuildSettings(7, onData, API));
    window.dispatchEvent(new Event('focus'));
    await vi.waitFor(() => expect(onData).toHaveBeenCalledWith({ plan: 'guild' }));
    expect(upstream).toHaveBeenCalledTimes(1);

    off();
    window.dispatchEvent(new Event('focus'));
    expect(upstream).toHaveBeenCalledTimes(1);
  });
});

describe('the settings read after a bot install return', () => {
  it('tells the caller when the fresh read fails, and keeps the cached answer', async () => {
    setQueryData(`${API}/v1/guilds/7/settings`, { plan: null });
    vi.stubGlobal(
      'fetch',
      vi.fn<GlobalFetch>(async () => envelope({ code: 'boom', message: 'x' }, 500)),
    );
    const onData = vi.fn();
    const onError = vi.fn();

    refetchGuildSettings(7, onData, API, onError);
    await vi.waitFor(() => expect(onError).toHaveBeenCalledTimes(1));
    expect(onData).not.toHaveBeenCalled();
  });

  it('drops bot_refusal from the cached read once shown, so a reload cannot draw it again', async () => {
    const discord = { available: true, connected: false, bot_available: true, bot: null };
    setQueryData(`${API}/v1/guilds/7/settings`, {
      plan: 'guild',
      discord: { ...discord, bot_refusal: { guild_name: 'Olympus' } },
    });
    const upstream = vi.fn<GlobalFetch>(async () => envelope({}));
    vi.stubGlobal('fetch', upstream);

    await clearCachedBotRefusal(7, API);

    const settings = await fetchGuildSettings(7, API);
    expect(settings.discord?.bot_refusal).toBeNull();
  });
});
