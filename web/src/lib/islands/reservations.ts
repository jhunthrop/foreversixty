// web/src/lib/islands/reservations.ts
// A shell's pre-hydration height reservation: the min-h classes an Astro page puts on an island's
// mount element so the footer stays put until the island draws. Once the island mounts it governs
// its own height, and a reservation left in place floors the page there for good, a hole above
// the footer on a short page. sim-island.ts and sim-tools-island.ts drop every one on mount.

/** The mount element's min-height classes, at every breakpoint. */
export function shellReservations(element: HTMLElement): string[] {
  return [...element.classList].filter((name) => /(^|:)min-h-\[/.test(name));
}
