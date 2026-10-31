// web/tests/captures/log-coach/playwright.capture.config.ts
// Run from web/ with the fixture site (a copy of dist) served on port 4482:
//   cp -R dist /tmp/coach-dist && npx astro preview --outDir /tmp/coach-dist --port 4482   # background
//   npx playwright test -c tests/captures/log-coach/playwright.capture.config.ts
// PNGs, the logged y values and the side-by-side boards go to design/reviews/2026-10-27-log-coach/.
import { defineConfig } from '@playwright/test';
import { HERMETIC_LAUNCH_ARGS } from '../../hermetic';

const PORT = process.env.CAPTURE_PORT ?? '4482';
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
