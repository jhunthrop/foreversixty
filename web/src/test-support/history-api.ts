// web/src/test-support/history-api.ts
// A fake GET /v1/characters/{region}/{ruleset}/{name}/history for vitest, in the idiom of
// sim-api.ts: a fetch stub answering the section-10 envelope, with the request recorded. An
// unmatched request throws rather than reaching the network.
import { vi } from 'vitest';
import { characterHistoryFixture } from '../fixtures/character-history';
import { envelope, failure } from './sim-api';

export const HISTORY_ROUTE = /\/v1\/characters\/[^/]+\/[^/]+\/[^/]+\/history$/;

export interface HistoryApiStub {
  install(): void;
  reset(): void;
  /** Answers with a status instead of the payload (402 lapsed, 404 not the owner, 500). */
  respondWith(status: number): void;
  lastUrl(): string;
}

export function createHistoryApi(): HistoryApiStub {
  let history = characterHistoryFixture();
  let status = 200;
  let lastUrl = '';
  return {
    install: () => {
      vi.stubGlobal(
        'fetch',
        vi.fn((input: Request | string) => {
          lastUrl = typeof input === 'string' ? input : input.url;
          if (!HISTORY_ROUTE.test(lastUrl)) throw new Error(`history stub: no route for ${lastUrl}`);
          return Promise.resolve(status === 200 ? envelope(history) : failure('history failed', status));
        }),
      );
    },
    reset: () => {
      vi.unstubAllGlobals();
      history = characterHistoryFixture();
      status = 200;
      lastUrl = '';
    },
    respondWith: (value) => {
      status = value;
    },
    lastUrl: () => lastUrl,
  };
}
