// web/src/lib/coach/list-gate.ts
// Spec 7: the combatant list waits for `/v1/me` (so the viewer's own row paints with its strip slot
// in place) for at most this long. After that the list renders anyway; a late answer then adds the
// strip to the viewer's own row only.
export const LIST_GATE_MS = 150;

export type ListGateReason = 'waiting' | 'resolved' | 'timed_out';

/** Why the list may or may not render: the answer is in (signed in or out), or the wait is spent. */
export function listGate(resolved: boolean, timedOut: boolean): ListGateReason {
  if (resolved) return 'resolved';
  return timedOut ? 'timed_out' : 'waiting';
}

export const listGateOpen = (reason: ListGateReason): boolean => reason !== 'waiting';
