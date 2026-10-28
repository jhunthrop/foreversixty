// web/tests/captures/premium-examples/playwright.capture.config.ts
// Run from web/ with the site served on port 4481 (CAPTURE_PORT overrides):
//   cp -R dist <dir> && npx astro preview --outDir <dir> --port 4481   # in the background
//   SIM_ID=<a finished server Top Gear run> npx playwright test -c tests/captures/premium-examples/playwright.capture.config.ts
// Writes web/public/premium/example-server-run.{avif,webp} and web/src/data/premium-examples.json
// (see capture.spec.ts).
import { defineConfig } from '@playwright/test';

const PORT = process.env.CAPTURE_PORT ?? '4481';
const TEST_TIMEOUT_MS = 90_000;

export default defineConfig({
  testDir: '.',
  testMatch: 'capture*.spec.ts',
  timeout: TEST_TIMEOUT_MS,
  workers: 1,
  reporter: 'list',
  use: {
    baseURL: `http://localhost:${PORT}`,
    browserName: 'chromium',
    deviceScaleFactor: 2,
    timezoneId: 'America/Chicago',
    reducedMotion: 'reduce',
  },
});
