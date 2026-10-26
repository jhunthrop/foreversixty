// web/src/lib/guild/minute-clock.svelte.ts
// The one minute tick the raid surfaces read `now` from (raids spec 3.1, 3.4): countdowns step and
// `Tonight` turns over at midnight without a reload. A fixed `now` (tests, captures) stops the tick.
// Call it while a component initialises: it owns an effect.

export const MINUTE_MS = 60_000;

export interface MinuteClock {
  readonly now: Date;
}

/** `fixed` wins while it returns a date; `active` gates the timer (a tab with nothing to count). */
export function minuteClock(fixed: () => Date | undefined, active: () => boolean = () => true): MinuteClock {
  /** The tick as epoch ms: a primitive, so each tick is a new value and no Date is shared. */
  let tickedMs = $state(Date.now());
  $effect(() => {
    if (fixed() !== undefined || !active()) return;
    const timer = setInterval(() => (tickedMs = Date.now()), MINUTE_MS);
    return () => clearInterval(timer);
  });
  return {
    get now(): Date {
      // A snapshot read from the reactive tick, never mutated: SvelteDate would add nothing.
      // eslint-disable-next-line svelte/prefer-svelte-reactivity
      return fixed() ?? new Date(tickedMs);
    },
  };
}
