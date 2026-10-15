// web/src/lib/character-search-api.test.ts
// The typeahead's one read: GET /v1/characters/search, the URL it builds and the signal it passes.
import { afterEach, describe, expect, it, vi } from 'vitest';
import { searchCharacters } from './character-search-api';

const API = 'https://api.foreversixty.test';
type GlobalFetch = (...args: Parameters<typeof fetch>) => Promise<Response>;

const ROW = {
  key: 'us/normal/bolt',
  region: 'us',
  ruleset: 'normal',
  name: 'Bolt',
  class: 'Mage',
  level: 60,
  spec: 'Fire',
};

function stubFetch(rows: unknown[] = [ROW]) {
  const upstream = vi.fn<GlobalFetch>(
    async () =>
      new Response(JSON.stringify({ ok: true, data: { rows }, error: null, request_id: 'r' }), {
        status: 200,
        headers: { 'content-type': 'application/json' },
      }),
  );
  vi.stubGlobal('fetch', upstream);
  return upstream;
}

const sent = (upstream: ReturnType<typeof stubFetch>): Request => upstream.mock.calls[0][0] as Request;

afterEach(() => vi.unstubAllGlobals());

describe('searchCharacters', () => {
  it('encodes the text and leaves out the filters that are not set', async () => {
    const upstream = stubFetch();
    const rows = await searchCharacters('Bo Lt&', { signal: new AbortController().signal }, API);
    expect(rows).toEqual([ROW]);
    expect(sent(upstream).url).toBe(`${API}/v1/characters/search?q=Bo+Lt%26`);
  });

  it('sends region, ruleset and limit when given', async () => {
    const upstream = stubFetch([]);
    const signal = new AbortController().signal;
    await searchCharacters('bo', { region: 'eu', ruleset: 'hardcore', limit: 5, signal }, API);
    expect(sent(upstream).url).toBe(`${API}/v1/characters/search?q=bo&region=eu&ruleset=hardcore&limit=5`);
  });

  it('passes the abort signal to the request', async () => {
    const upstream = stubFetch();
    const controller = new AbortController();
    await searchCharacters('bo', { signal: controller.signal }, API);
    controller.abort();
    expect(sent(upstream).signal.aborted).toBe(true);
  });
});
