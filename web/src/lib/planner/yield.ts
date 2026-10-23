// web/src/lib/planner/yield.ts
// A break in a long run of main-thread work. The planner's boot (the data landing in the store,
// the three talent trees drawing) is several milliseconds of script each; under Lighthouse's 4x
// CPU throttling their sum is one task past 100 ms, and the page's total blocking time counts
// every millisecond of a task beyond 50. Awaiting this between the pieces ends the task, so the
// browser can paint and answer input before the next piece starts.

type SchedulerWithYield = { yield?: () => Promise<void> };

/** Resolves in a new task: `scheduler.yield()` where the browser has it, a zero timeout where not. */
export function yieldToMain(): Promise<void> {
  const scheduler = (globalThis as { scheduler?: SchedulerWithYield }).scheduler;
  if (scheduler?.yield !== undefined) return scheduler.yield();
  return new Promise((resolve) => setTimeout(resolve, 0));
}
