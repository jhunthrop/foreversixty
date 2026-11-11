// web/tests/e2e/support/shared-build-page.ts
// The shared build /b/:id as the API serves it: fixtures/shared-build/survival.html and missing.html,
// written by api/internal/site/page_test.go (`go test ./internal/site -run
// TestSharedBuildParityFixtures -update`) from the mock's record (api/internal/site/testdata/
// shared-survival.json). The API's chrome.html is a copy of an earlier build's header and footer
// (its scripts load by stable /_chrome/ path, scripts/chrome-aliases.mjs, but its markup can lag):
// the page is served with this build's own chrome in its place (dist/about.html, the file
// chrome.html is copied from), so the header is the site's.
import { readFileSync } from 'node:fs';
import path from 'node:path';
import type { Page } from '@playwright/test';
import { routePlannerState } from './planner-page';

const WEB_ROOT = path.join(import.meta.dirname, '..', '..', '..');
const FIXTURES = path.join(WEB_ROOT, 'tests', 'e2e', 'fixtures', 'shared-build');

/** The first site <header> and the last <footer> of a page (api/internal/site chrome_test.go's rule). */
function chromeOf(html: string): { header: string; footer: string } {
  const header = html.slice(html.indexOf('<header'), html.indexOf('</header>') + '</header>'.length);
  const footer = html.slice(html.lastIndexOf('<footer'), html.lastIndexOf('</footer>') + '</footer>'.length);
  return { header, footer };
}

/** The API page with this build's chrome in place of the copy it was rendered with. */
function apiPage(fixture: string): string {
  const page = readFileSync(path.join(FIXTURES, fixture), 'utf8');
  const served = chromeOf(page);
  const site = chromeOf(readFileSync(path.join(WEB_ROOT, 'dist', 'about.html'), 'utf8'));
  return page.replace(served.header, site.header).replace(served.footer, site.footer);
}

/** The id of the recorded Survival raid build (survival.html). */
export const SHARED_BUILD_ID = (
  JSON.parse(
    readFileSync(path.join(WEB_ROOT, '..', 'api/internal/site/testdata/shared-survival.json'), 'utf8'),
  ) as { id: string }
).id;

/** Opens /b/<id> as the API renders `fixture`, signed out. */
export async function openSharedBuild(page: Page, fixture: string, id: string): Promise<void> {
  await routePlannerState(page, 'out');
  const body = apiPage(fixture);
  await page.route(`**/b/${id}`, (route) => route.fulfill({ status: 200, contentType: 'text/html', body }));
  await page.goto(`/b/${id}`);
}
