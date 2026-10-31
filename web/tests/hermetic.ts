// web/tests/hermetic.ts
// No test, local or CI, ever reaches the production API (2026-10-31: unstubbed calls from
// local Playwright runs spent the owner's own per-IP rate limit and a real raid save got a
// 429). Three layers, all defined here so the e2e suite, the capture configs and the capture
// scripts share one copy:
//   1. the e2e site is built against DEAD_API_BASE, a port nothing listens on, so an
//      unstubbed call is refused on the machine instead of answered by a real server;
//   2. every test browser launches with HERMETIC_LAUNCH_ARGS, which makes the production
//      API's host unresolvable, so even a dist built with the production default (a local
//      `npm run build` without PUBLIC_API_BASE_URL) cannot reach it;
//   3. tests/e2e/support/test.ts fails any test whose page asks an API origin for something
//      the test did not stub (isApiRequest below).

/** The API origin the e2e and capture builds carry: the discard port on loopback, refused at once. */
export const DEAD_API_BASE = 'http://127.0.0.1:9';

/** The production API's host: the default `lib/planner/config.ts` falls back to. */
export const PRODUCTION_API_HOST = 'api.foreversixty.gg';

/** Chromium's host-resolver rule that makes the production API's host fail DNS. lighthouserc.json
 *  carries the same rule in its chromeFlags (src/e2e-hermetic.test.ts pins the two equal). */
export const HOST_BLOCK_RULE = `MAP ${PRODUCTION_API_HOST} ~NOTFOUND`;

/** Chromium flags that make the production API's host fail DNS inside the test browser. */
export const HERMETIC_LAUNCH_ARGS: readonly string[] = [`--host-resolver-rules=${HOST_BLOCK_RULE}`];

const API_HOSTS = new Set([new URL(DEAD_API_BASE).host, PRODUCTION_API_HOST]);

/** True for a request to any origin the site's API calls can go to under test. */
export function isApiRequest(url: URL | string): boolean {
  const parsed = typeof url === 'string' ? new URL(url) : url;
  return API_HOSTS.has(parsed.host);
}
