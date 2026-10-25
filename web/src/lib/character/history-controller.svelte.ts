// web/src/lib/character/history-controller.svelte.ts
// The reactive read of the owner's character history for the History section: one fetch, the
// state it is in, and a retry that re-reads in place. The fetcher is injected so the section
// can be given any state without a network.
import type { CharacterPath } from '../characters';
import { fetchCharacterHistory, type CharacterHistory } from './history-api';

export type HistoryStatus = 'loading' | 'ready' | 'error';

export type HistoryFetcher = (path: CharacterPath) => Promise<CharacterHistory>;

export interface HistoryController {
  readonly status: HistoryStatus;
  readonly history: CharacterHistory | null;
  /** Reads the history; calling it again after an error retries, and a stale answer is dropped. */
  load: () => void;
}

export function createHistoryController(
  path: CharacterPath,
  fetcher: HistoryFetcher = fetchCharacterHistory,
): HistoryController {
  let status = $state<HistoryStatus>('loading');
  let history = $state<CharacterHistory | null>(null);
  let attempt = 0;
  return {
    get status() {
      return status;
    },
    get history() {
      return history;
    },
    load(): void {
      attempt += 1;
      const mine = attempt;
      status = 'loading';
      void fetcher(path)
        .then((result) => {
          if (mine !== attempt) return;
          history = result;
          status = 'ready';
        })
        .catch(() => {
          if (mine !== attempt) return;
          history = null;
          status = 'error';
        });
    },
  };
}
