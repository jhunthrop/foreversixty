// web/src/lib/sim/saved-sim-state.svelte.ts
// The one-shot saved-sim fetch (/sim/<id>) SimView.svelte reads: pulled out of the
// component so its own state and retry logic don't grow SimView.svelte past the 800-line
// budget task 7's fix round hit. `load` resets `error` before each attempt so a retry
// never leaves a stale message showing while the new request is in flight.
import { fetchSim } from './api';
import { simCopy } from './copy';
import type { SimResult } from './types';

export function createSavedSimState(initial: SimResult | null) {
  let result = $state<SimResult | null>(initial);
  let error = $state<string | null>(null);

  function load(simId: string): void {
    error = null;
    void fetchSim(simId)
      .then((fetched) => (result = fetched))
      .catch((thrown: unknown) => {
        error = thrown instanceof Error ? thrown.message : simCopy.loadFailed;
      });
  }

  return {
    get result() {
      return result;
    },
    get error() {
      return error;
    },
    load,
  };
}
