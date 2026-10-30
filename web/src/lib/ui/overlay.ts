// web/src/lib/ui/overlay.ts
// One open overlay at a time (design/specs/2026-10-30-controls-charts.md §C): every menu and
// list registers here while open, opening one closes the others, and charts subscribe to
// the count so they stand down (no tooltip, no crosshair, no hover) while a menu is up.
// Framework-free so the static pages' fields use it without a Svelte runtime.

type Close = () => void;
type Listener = (openCount: number) => void;

const open = new Set<Close>();
const listeners = new Set<Listener>();

function notify(): void {
  for (const listener of listeners) listener(open.size);
}

/** Registers an opened overlay, closing every other one first. Returns its release. */
export function claimOverlay(close: Close): () => void {
  for (const other of [...open]) {
    if (other !== close) {
      open.delete(other);
      other();
    }
  }
  open.add(close);
  notify();
  return () => {
    if (open.delete(close)) notify();
  };
}

/** The number of overlays open now. */
export function openOverlayCount(): number {
  return open.size;
}

/** Calls `listener` with the open count now and on every change. Returns the unsubscribe. */
export function subscribeOverlays(listener: Listener): () => void {
  listeners.add(listener);
  listener(open.size);
  return () => {
    listeners.delete(listener);
  };
}
