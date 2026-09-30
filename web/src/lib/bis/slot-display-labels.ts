// web/src/lib/bis/slot-display-labels.ts
// The best-in-slot list's own slot column (bis rebuild spec §4.D): the client's own
// paperdoll labels both ring sockets and both trinket sockets identically ("Finger" /
// "Trinket"), unlike `planner/types.ts`'s own `SLOT_LABELS`, which stays the disambiguating
// internal label every other caller (the planner, `data-testid`s) still needs. A second map,
// not an overload of `SLOT_LABELS` itself, because two different callers need two different
// strings for the same key.
//
// Lives here, not alongside `SLOT_LABELS` in `lib/planner/types.ts`, even though it is built
// from that same map: `planner/types.ts` is the entry point of the planner's own Vite island
// bundle (`vite.island.config.ts`), and adding a second export there -- even one the planner
// itself never imports -- changed that bundle's own chunking enough to push it over its
// Lighthouse budget (harness review, 2026-09-30: planner-island.js grew from 240,139 to
// 242,869 raw bytes measuring this exact change). A leveling-BiS-only module the planner
// never reaches keeps that island byte-for-byte what main already ships.
import { SLOT_LABELS } from '../planner/types';
import type { Slot } from '../planner/types';

export const SLOT_DISPLAY_LABELS: Record<Slot, string> = {
  ...SLOT_LABELS,
  finger1: 'Ring',
  finger2: 'Ring',
  trinket1: 'Trinket',
  trinket2: 'Trinket',
};
