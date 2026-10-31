// web/tests/captures/launch-pricing/playwright.capture.config.ts
// Run from web/ with the fixture site (a copy of dist) served on port 4481:
//   cp -R dist <dir> && npx astro preview --outDir <dir> --port 4481   # in the background
//   npx playwright test -c tests/captures/launch-pricing/playwright.capture.config.ts
// PNGs go to design/reviews/2026-10-27-launch-pricing/ (see capture.spec.ts).
import { defineConfig } from '@playwright/test';
import { HERMETIC_LAUNCH_ARGS } from '../../hermetic';

const PORT = process.env.CAPTURE_PORT ?? '4481';
const TEST_TIMEOUT_MS = 90_000;

export default defineConfig({
  testDir: '.',
  testMatch: 'capture*.spec.ts',
  timeout: TEST_TIMEOUT_MS,
  workers: 4,
  reporter: 'list',
  use: {
    // The production API's host does not resolve in this browser (tests/hermetic.ts).
    launchOptions: { args: [...HERMETIC_LAUNCH_ARGS] },
    baseURL: `http://localhost:${PORT}`,
    browserName: 'chromium',
    deviceScaleFactor: 1,
    timezoneId: 'America/Chicago',
    reducedMotion: 'reduce',
  },
});
