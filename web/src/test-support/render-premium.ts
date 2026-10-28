// web/src/test-support/render-premium.ts
// The plans page's shared container render. It lives outside any *.test.ts so a test file that
// imports it does not also re-run another file's suites.
import { JSDOM } from 'jsdom';
import Premium from '../pages/premium.astro';
import { createPageContainer } from './page-container';

/** Renders /premium through Base.astro and returns its parsed document. */
export async function renderPremium(): Promise<Document> {
  const container = await createPageContainer();
  const html = await container.renderToString(Premium);
  return new JSDOM(html).window.document;
}
