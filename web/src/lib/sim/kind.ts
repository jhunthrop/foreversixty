// web/src/lib/sim/kind.ts
// The contract's derived kind (2026-09-19-simulator-parity-interfaces.md, 1.1): "Kind is
// derived, never sent". `api.SimRequest.Kind()` is the Go original; this is its mirror, so
// the page can render a saved sim by what its stored request actually is rather than by a
// field a client could set to anything.
import type { SimRequest } from './types';

export type SimKind = 'run' | 'gear' | 'talents' | 'drops' | 'weights';

export const SIM_KINDS: readonly SimKind[] = ['run', 'gear', 'talents', 'drops', 'weights'];

const BULK_MODES: readonly SimKind[] = ['gear', 'talents', 'drops'];

/**
 * Bulk wins over weights: a request carrying both is malformed, and reporting one kind
 * rather than two is what lets `/sim/<id>` pick a single renderer without a tie-break of
 * its own. A `bulk.mode` outside the vocabulary reads as a plain run -- the engine will
 * refuse the request anyway, and the page must not render a bulk table for it meanwhile.
 */
export function requestKind(request: Pick<SimRequest, 'bulk' | 'weights'>): SimKind {
  const mode = request.bulk?.mode;
  if (mode !== undefined) {
    return BULK_MODES.find((known) => known === mode) ?? 'run';
  }
  return request.weights === undefined ? 'run' : 'weights';
}
