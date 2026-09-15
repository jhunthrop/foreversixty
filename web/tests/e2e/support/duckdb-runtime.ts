// web/tests/e2e/support/duckdb-runtime.ts
// The two DuckDB engine modules are over Cloudflare's 25 MiB static-asset cap, so they are
// not published into dist/: in production src/worker.ts serves them out of the LOGS bucket
// at /duckdb-runtime/<version>/<file>, and `astro dev` serves them off the staging
// directory (astro.config.mjs's duckdbRuntime integration).
//
// `astro preview`, which is what this suite runs against, can do neither: it is Vite's
// preview server over dist/, and Astro strips user Vite plugins before starting it. So the
// one spec that instantiates the real engine fulfils that route itself, from the same
// staged bytes the upload would have put in the bucket. The URL the page asks for is still
// the production one, on this origin, which is what report-queries.spec.ts asserts.
import { readFile } from 'node:fs/promises';
import path from 'node:path';
import type { Page } from '@playwright/test';
import {
  DUCKDB_RUNTIME_MODULES,
  DUCKDB_WASM_VERSION,
  duckdbRuntimeUrl,
} from '../../../src/lib/report/duckdb-runtime';

const STAGED = path.resolve(process.cwd(), 'build/duckdb-runtime', DUCKDB_WASM_VERSION);

/** Answers /duckdb-runtime/<version>/<module> the way the Worker does in production. */
export async function serveDuckdbRuntime(page: Page): Promise<void> {
  for (const file of DUCKDB_RUNTIME_MODULES) {
    // Read inside the handler, not here: only one of the two builds is ever asked for
    // (Chromium takes the exception-handling one) and they are 33 and 38 MB.
    await page.route(`**${duckdbRuntimeUrl(file)}`, async (route) =>
      route.fulfill({
        status: 200,
        contentType: 'application/wasm',
        headers: { 'cache-control': 'public, max-age=31536000, immutable' },
        body: await readFile(path.join(STAGED, file)),
      }),
    );
  }
}
