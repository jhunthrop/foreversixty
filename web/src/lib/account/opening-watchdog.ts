// web/src/lib/account/opening-watchdog.ts
// The connect link's escape hatch: a click that never navigates (a blocked or swallowed
// navigation) must not leave `Opening Discord…` on the button for good. If the page is still here
// this long after the click, `onStuck` runs; leaving the page or `pageshow` cancels it.

export const OPENING_TIMEOUT_MS = 10_000;

export interface OpeningWatchdog {
  start: () => void;
  cancel: () => void;
}

export function createOpeningWatchdog(onStuck: () => void): OpeningWatchdog {
  let timer: ReturnType<typeof setTimeout> | undefined;
  const cancel = (): void => clearTimeout(timer);
  return {
    start: () => {
      cancel();
      timer = setTimeout(onStuck, OPENING_TIMEOUT_MS);
    },
    cancel,
  };
}
