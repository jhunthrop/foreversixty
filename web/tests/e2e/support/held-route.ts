// web/tests/e2e/support/held-route.ts
// A response a test holds open until it chooses to release it, so two requests can be made
// to interleave deterministically instead of racing: start the slow one, wait for it to
// reach the network, trigger whatever should happen while it is still in flight, then
// release it and see that its answer is handled correctly (or dropped, if it is stale by
// the time it lands).
//
// report-tabs.spec.ts and report-queries.spec.ts each carry this same shape inline, fixed
// to their own URL. This is the third caller wanting it, generalised to take the pattern
// instead of duplicating the implementation a third time.
import type { Page, Route } from '@playwright/test';

export interface HeldRoute {
  /** Resolves once the route handler has been entered -- the request reached the network. */
  started: Promise<void>;
  /** Lets the held request proceed to `settle`. */
  release: () => void;
}

/**
 * `'abort'` and `'continue'` cover the two shapes report-tabs.spec.ts and
 * report-queries.spec.ts each held open -- a real asset request that should fail or
 * complete against the fixture on disk. A held API call that has to answer with mock JSON
 * (nothing is really listening on the API origin in these tests) needs a third shape: a
 * function that receives the route and decides how to settle it, `route.fulfill(...)`
 * included.
 */
export type HeldRouteSettle = 'abort' | 'continue' | ((route: Route) => Promise<void> | void);

export async function heldRoute(
  page: Page,
  urlPattern: string,
  settle: HeldRouteSettle = 'continue',
): Promise<HeldRoute> {
  let markStarted = (): void => {};
  let release = (): void => {};
  const started = new Promise<void>((resolve) => (markStarted = resolve));
  const gate = new Promise<void>((resolve) => (release = resolve));

  await page.route(urlPattern, async (route) => {
    markStarted();
    await gate;
    if (settle === 'abort') await route.abort();
    else if (settle === 'continue') await route.continue();
    else await settle(route);
  });

  return { started, release: () => release() };
}
