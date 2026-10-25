import { afterEach, beforeEach, describe, expect, it } from 'vitest';
import { characterHistoryFixture } from '../../fixtures/character-history';
import { createHistoryApi } from '../../test-support/history-api';
import { HistoryApiError, fetchCharacterHistory } from './history-api';

const PATH = { region: 'us', ruleset: 'hardcore', slug: 'elyra-duskvale' } as const;
const stub = createHistoryApi();

beforeEach(() => stub.install());
afterEach(() => stub.reset());

describe('fetchCharacterHistory', () => {
  it('reads the owner route and returns the payload', async () => {
    const history = await fetchCharacterHistory(PATH, 'https://api.test');
    expect(stub.lastUrl()).toBe('https://api.test/v1/characters/us/hardcore/elyra-duskvale/history');
    expect(history).toEqual(characterHistoryFixture());
  });

  it('throws a HistoryApiError carrying 402 for a lapsed plan', async () => {
    stub.respondWith(402);
    await expect(fetchCharacterHistory(PATH, 'https://api.test')).rejects.toMatchObject({
      name: 'HistoryApiError',
      status: 402,
    });
  });

  it('throws a HistoryApiError for a server failure', async () => {
    stub.respondWith(500);
    await expect(fetchCharacterHistory(PATH, 'https://api.test')).rejects.toBeInstanceOf(HistoryApiError);
  });
});
