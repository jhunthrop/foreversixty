// web/tests/e2e/report-queries.spec.ts
// The one place the real DuckDB runs: vitest cannot start it (the Node build wants a
// web-worker shim), so src/lib/report/query.test.ts drives a fake engine and everything
// that needs an actual WebAssembly instance is asserted here, against the checked-in
// fixture Parquet (src/fixtures/report/fights/3/events.parquet).
import { expect, test, type Page } from '@playwright/test';
import { DUCKDB_WASM_VERSION } from '../../src/lib/report/duckdb-runtime';
import { serveDuckdbRuntime } from './support/duckdb-runtime';
import { readFileSync } from 'node:fs';

/** The lines of fight 3 (Warden Kelthas, encounter 9001) in the fixture log, boundaries included. */
function fightThreeLines(): number {
  const lines = readFileSync(new URL('../../src/fixtures/report/fixture.log', import.meta.url), 'utf8').split(
    '\n',
  );
  const start = lines.findIndex((line) => line.includes('ENCOUNTER_START,9001,'));
  const end = lines.findIndex((line) => line.includes('ENCOUNTER_END,9001,'));
  if (start < 0 || end < start) throw new Error('fixture.log has no complete encounter 9001');
  return end - start + 1;
}

const REPORT = '/reports/fixture2abcd';
const QUERIES = `${REPORT}?fight=3&view=queries`;
// The preview's own address, which E2E_PORT moves (see playwright.config.ts): a run on a
// second port must still recognise its own bytes as first-party.
const ORIGIN = `http://localhost:${process.env.E2E_PORT ?? '4321'}`;
// The report island calls the rankings API for its percentiles (Task 12), which is a
// first-party origin the contract names. Everything else -- a CDN, a font host, DuckDB's
// own extension repository -- is what "no third-party bytes, on any page, ever" forbids.
const OWN_ORIGINS = new Set([ORIGIN, 'https://api.foreversixty.gg']);

/** Every request the page made, newest last, so a test can assert on what was not asked for. */
function recordRequests(page: Page): string[] {
  const seen: string[] = [];
  page.on('request', (request) => seen.push(request.url()));
  return seen;
}

// Both halves of the runtime: the worker script and the extension under /duckdb/, and the
// engine module under /duckdb-runtime/, which is served from R2 rather than from dist/.
const duckdbAssets = (urls: string[]): string[] =>
  urls.filter((url) => /\/duckdb(-runtime)?\//.test(new URL(url).pathname));
const parquetFiles = (urls: string[]): string[] => urls.filter((url) => /events\.parquet(\?|$)/.test(url));

// Spec section 4: "DuckDB-WASM loads lazily on the first deep interaction, never on page
// load." Opening the Queries view is not the deep interaction -- running a query is -- so
// browsing the whole report, tabs included, must still cost nothing.
test('nothing DuckDB-related loads on page load, while browsing, or on opening Queries', async ({ page }) => {
  const fetched = recordRequests(page);

  await page.goto(REPORT);
  await expect(page.getByTestId('summary-tab')).toBeVisible();
  const phone = (page.viewportSize()?.width ?? 1280) < 1024;
  if (phone) await page.getByTestId('tab-select').selectOption('damage-done');
  else await page.getByTestId('tab-damage-done').click();
  if (phone) {
    await page.getByTestId('tab-select').selectOption('casts');
  } else {
    // Casts lives under the category row's "More" menu now (design review 2026-09-26
    // finding 3): a native `<details>`, opened by clicking its own summary.
    await page.getByTestId('tab-more').click();
    await page.getByTestId('tab-casts').click();
  }
  await page.getByTestId('view-events').click();
  await expect(page.getByTestId('events-view')).toBeVisible();

  expect(duckdbAssets(fetched)).toEqual([]);
  expect(parquetFiles(fetched)).toEqual([]);

  await page.getByTestId('view-queries').click();
  await expect(page.getByTestId('queries-view')).toBeVisible();
  await expect(page.getByTestId('query-run')).toBeEnabled();

  // The engine and the two-to-ten megabyte event file are still untouched: the view is
  // rendered, the SQL box is filled in, and not a byte of either has been asked for.
  expect(duckdbAssets(fetched)).toEqual([]);
  expect(parquetFiles(fetched)).toEqual([]);
});

test('a statement that is not a read is refused before anything loads', async ({ page }) => {
  const fetched = recordRequests(page);
  await page.goto(QUERIES);

  await page.getByTestId('query-sql').fill('DROP TABLE events');
  await page.getByTestId('query-run').click();

  await expect(page.getByTestId('query-error')).toHaveText(
    'Queries here read the fight; they cannot change it.',
  );
  expect(duckdbAssets(fetched)).toEqual([]);
  expect(parquetFiles(fetched)).toEqual([]);
});

test('a query runs against the fight’s own Parquet, and every byte comes from this origin', async ({
  page,
}) => {
  test.slow(); // The first run downloads and instantiates the WebAssembly build.
  const external: string[] = [];
  const fetched = recordRequests(page);
  page.on('request', (request) => {
    if (!OWN_ORIGINS.has(new URL(request.url()).origin)) external.push(request.url());
  });
  await serveDuckdbRuntime(page);

  await page.goto(QUERIES);
  await page.getByTestId('template-event-counts').click();
  await page.getByTestId('query-run').click();

  const table = page.getByTestId('query-result');
  await expect(table).toBeVisible({ timeout: 90_000 });
  await expect(table).toContainText('SPELL_DAMAGE');
  // The cold run paid for the engine and the download, and says so, so nobody reads the
  // number as the cost of the query.
  await expect(page.getByTestId('query-status')).toContainText('That run included the one-time load');

  // A warm run answers from memory: no second engine, no second download.
  await page.getByTestId('template-damage-by-ability').click();
  await page.getByTestId('query-run').click();
  await expect(page.getByTestId('query-result')).toContainText('Anima Lash');
  await expect(page.getByTestId('query-status')).not.toContainText('one-time load');

  // A query with nothing to say says nothing rather than rendering an empty grid: the
  // fixture fight has no misses at all.
  await page.getByTestId('template-misses').click();
  await page.getByTestId('query-run').click();
  await expect(page.getByTestId('query-empty')).toHaveText('No rows matched.');

  // A malformed query is DuckDB's own sentence, not an unhandled rejection, and the page
  // is still usable afterwards.
  await page.getByTestId('query-sql').fill("SELECT * FROM read_parquet('nope.parquet')");
  await page.getByTestId('query-run').click();
  await expect(page.getByTestId('query-error')).toBeVisible();

  // A query naming an address is refused before it reaches the engine: it would autoload
  // httpfs, which is not among the vendored extensions, and then fetch from that address.
  await page.getByTestId('query-sql').fill("SELECT * FROM read_parquet('https://example.invalid/x.parquet')");
  await page.getByTestId('query-run').click();
  await expect(page.getByTestId('query-error')).toHaveText(
    'Queries here read this fight\u2019s own file; they cannot fetch from another address.',
  );

  // And the other half of the same worry, which the guard above deliberately does not
  // cover: an extension that is simply not vendored, reached without naming an address.
  // custom_extension_repository has to replace DuckDB's default rather than extend it.
  // The error naming spatial is what says the query reached the engine rather than
  // leaving the previous alert on screen; the trailing same-origin assertion is what says
  // there was no fallback to extensions.duckdb.org.
  await page.getByTestId('query-sql').fill("SELECT * FROM st_read('nothing.shp')");
  await page.getByTestId('query-run').click();
  await expect(page.getByTestId('query-error')).toContainText('it exists in the spatial extension');
  await page.getByTestId('query-sql').fill("SELECT count(*) AS n FROM read_parquet('events.parquet')");
  await page.getByTestId('query-run').click();
  await expect(page.getByTestId('query-error')).toHaveCount(0);
  // Every line of fight 3 in the fixture log, ENCOUNTER_START and ENCOUNTER_END included,
  // counted from the log itself so a line added to the fixture moves this with it.
  await expect(page.getByTestId('query-result')).toContainText(String(fightThreeLines()));

  expect(external).toEqual([]);

  // The engine module comes from the version-pinned Worker route, not from a static asset
  // under /duckdb/: at 32.7 MiB it is over Cloudflare's 25 MiB per-file limit, so shipping
  // it in dist/ fails the whole site's deploy (scripts/duckdb-runtime.mjs). The worker
  // script beside it is small enough to stay a static asset, and is version-pinned too so
  // public/_headers can cache both for a year.
  expect(duckdbAssets(fetched)).toContain(`${ORIGIN}/duckdb-runtime/${DUCKDB_WASM_VERSION}/duckdb-eh.wasm`);
  expect(duckdbAssets(fetched)).toContain(
    `${ORIGIN}/duckdb/${DUCKDB_WASM_VERSION}/duckdb-browser-eh.worker.js`,
  );
  const staticModules = duckdbAssets(fetched).filter(
    (url) => url.includes('/duckdb/') && url.endsWith('.wasm') && !url.includes('/extensions/'),
  );
  expect(staticModules).toEqual([]);
});

// The engine build, the download and the query all resolve long after the click, so a
// fight switch can land in the middle of any of them. This holds fight 3's event file
// open until fight 1 is already on screen, then releases it: an answer computed from
// fight 3's events must not be painted under fight 1's name.
test('an answer for the fight that was left behind is dropped, not painted', async ({ page }) => {
  test.slow();
  await serveDuckdbRuntime(page);
  await page.goto(QUERIES);

  let markStarted = (): void => {};
  let release = (): void => {};
  const started = new Promise<void>((resolve) => (markStarted = resolve));
  const gate = new Promise<void>((resolve) => (release = resolve));
  await page.route('**/fights/3/events.parquet*', async (route) => {
    markStarted();
    await gate;
    await route.continue();
  });

  await page.getByTestId('template-event-counts').click();
  await page.getByTestId('query-run').click();
  await started;

  await page.getByTestId('toggle-trash').click();
  await page.getByTestId('fight-1').click();
  await expect(page.getByTestId('fight-1')).toHaveAttribute('aria-current', 'true');

  const answered = page.waitForResponse((response) =>
    /\/fights\/3\/events\.parquet(\?|$)/.test(response.url()),
  );
  release();
  await answered;
  await page.waitForTimeout(1000);

  await expect(page.getByTestId('query-result')).toHaveCount(0);
  await expect(page.getByTestId('query-empty')).toHaveCount(0);
  await expect(page.getByTestId('query-error')).toHaveCount(0);
  await expect(page.getByTestId('query-run')).toBeEnabled();
});

// The island's budget is 140 KB gzipped (scripts/check-island-size.mjs) and DuckDB is an
// order of magnitude larger than that, so the library has to be in the chunk the dynamic
// import splits out, not in the entry every report page loads.
test('the report island bundle does not carry DuckDB', async ({ request }) => {
  const bundle = await (await request.get('/report-island.js')).text();

  // A protocol string from the worker binding, and the CDN helper the site's "no
  // third-party bytes" rule rules out: neither may appear in the eagerly loaded entry.
  expect(bundle).not.toContain('REGISTER_FILE_BUFFER');
  expect(bundle).not.toContain('cdn.jsdelivr.net');
  expect(bundle.length).toBeLessThan(700_000);
});
