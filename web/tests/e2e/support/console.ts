// web/tests/e2e/support/console.ts
// "No console errors" means no errors from the page's own code. Every page now asks the API
// who is signed in (the session link is in every header), and under test there is no API:
// the browser reports that as a console error of its own -- a CORS refusal from localhost
// and a "Failed to load resource" line -- which says nothing about the page. A spec that
// cares about the session stubs `/v1/me` itself; the rest collect errors through here.
import type { Page } from '@playwright/test';

/** The browser's own network reports, as opposed to anything a script threw or logged. */
const NETWORK_NOISE = [/Failed to load resource/, /blocked by CORS policy/, /net::ERR_/];

/** Starts collecting; read the returned array after the page has done its work. */
export function collectPageErrors(page: Page): string[] {
  const errors: string[] = [];
  page.on('console', (message) => {
    if (message.type() !== 'error') return;
    const text = message.text();
    if (!NETWORK_NOISE.some((pattern) => pattern.test(text))) errors.push(text);
  });
  // A thrown exception is never noise, whatever it says.
  page.on('pageerror', (error) => errors.push(String(error)));
  return errors;
}
