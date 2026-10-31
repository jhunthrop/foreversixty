// web/tests/e2e/support/test.ts
// The one `test` every e2e spec imports (src/e2e-hermetic.test.ts enforces it). It extends
// Playwright's own with the API guard (tests/hermetic.ts): the browser context answers any API
// request a spec did not stub by aborting it, records it, and fails the test with the list once
// the test body is done. Routes a spec adds (page or context) win over these, since Playwright
// tries the newest handler first and page routes before context routes.
//
// AMBIENT_READS below are the only exceptions: background reads that pages across the site make
// whatever the spec is about, each with the one answer it gets when no spec stubs it. Every
// other unstubbed request, and any unstubbed write, fails the test. A spec that cares about one
// of these reads stubs it itself.
import { test as base, type BrowserContext, type Page, type Route } from '@playwright/test';
import { DEAD_API_BASE, PRODUCTION_API_HOST, isApiRequest } from '../../hermetic';

export { expect } from '@playwright/test';

const SIGNED_OUT = { data: null, error: { code: 'unauthorized', message: 'Sign in first.' } };

type AmbientAnswer = (route: Route) => Promise<void>;

/** The signed-out visitor: what the header renders for anyone without a session. */
const signedOut: AmbientAnswer = (route) => route.fulfill({ status: 401, json: SIGNED_OUT });

/** The API out of reach: these reads are enrichment the page renders without. */
const unreachable: AmbientAnswer = (route) => route.abort('connectionrefused');

interface AmbientRead {
  path: RegExp;
  answer: AmbientAnswer;
}

const CHARACTER = '/v1/characters/[a-z]+/[a-z]+/[a-z0-9-]+';

/** GET reads with a default answer, by path. Keep this short: a new read fails loudly until a
 *  spec stubs it or it is shown to be site-wide background like these. */
const AMBIENT_READS: readonly AmbientRead[] = [
  // Every page's header asks who is signed in.
  { path: /^\/v1\/me$/, answer: signedOut },
  // The character selector and the sim pages enrich a known character with its sim input and rating.
  { path: new RegExp(`^${CHARACTER}/(sim-input|rating)$`), answer: unreachable },
  // The planner, sim and rankings pages refresh the build's spec and phase lists.
  { path: /^\/v1\/(specs|phases)$/, answer: unreachable },
  // Report pages: percentiles and per-fight ratings.
  { path: /^\/v1\/rankings\/percentile$/, answer: unreachable },
  { path: /^\/v1\/reports\/[a-z0-9]+\/fights\/\d+\/ratings$/, answer: unreachable },
  // The home and logs feeds, the header's finished-sim line and the coach's recent list.
  { path: /^\/v1\/(reports|reports\/recent|rankings\/guilds|sims|me\/coach\/recent)$/, answer: unreachable },
];

function ambientAnswer(route: Route): AmbientAnswer | null {
  const request = route.request();
  if (request.method() !== 'GET') return null;
  const { pathname } = new URL(request.url());
  return AMBIENT_READS.find((read) => read.path.test(pathname))?.answer ?? null;
}

async function guardApi(context: BrowserContext, unstubbed: string[]): Promise<void> {
  await context.route(isApiRequest, async (route) => {
    const answer = ambientAnswer(route);
    if (answer !== null) {
      await answer(route);
      return;
    }
    unstubbed.push(`${route.request().method()} ${route.request().url()}`);
    await route.abort('blockedbyclient');
  });
}

function unstubbedMessage(unstubbed: readonly string[]): string {
  const lines = [
    'Unstubbed API requests (stub them with page.route; tests never reach a real API):',
    ...unstubbed.map((request) => `  ${request}`),
  ];
  if (unstubbed.some((request) => request.includes(`//${PRODUCTION_API_HOST}/`))) {
    lines.push(`The served dist points at production: rebuild it with PUBLIC_API_BASE_URL=${DEAD_API_BASE}.`);
  }
  return lines.join('\n');
}

/** A spec's explicit "this read goes unanswered here": the page is tested without it. Routed on
 *  the context, so any stub the test itself adds (page or context) still wins. */
export async function apiUnreachable(page: Page, ...globs: readonly string[]): Promise<void> {
  for (const glob of globs) await page.context().route(glob, unreachable);
}

export const test = base.extend({
  context: async ({ context }, use) => {
    const unstubbed: string[] = [];
    await guardApi(context, unstubbed);
    await use(context);
    if (unstubbed.length > 0) throw new Error(unstubbedMessage(unstubbed));
  },
});
