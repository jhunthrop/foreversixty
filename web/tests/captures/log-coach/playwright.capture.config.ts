// web/tests/captures/log-coach/playwright.capture.config.ts
// Run from web/ with the fixture site (a copy of dist) served on port 4482:
//   cp -R dist /tmp/coach-dist && npx astro preview --outDir /tmp/coach-dist --port 4482   # background
//   npx playwright test -c tests/captures/log-coach/playwright.capture.config.ts
// PNGs, the logged y values and the side-by-side boards go to design/reviews/2026-10-27-log-coach/.
import { defineConfig } from '@playwright/test';

const PORT = process.env.CAPTURE_PORT ?? '4482';
const TEST_TIMEOUT_MS = 90_000;

export default defineConfig({
  testDir: '.',
  testMatch: 'capture*.spec.ts',
  timeout: TEST_TIMEOUT_MS,
  workers: 4,
  reporter: 'list',
  use: {
    baseURL: `http://localhost:${PORT}`,
    browserName: 'chromium',
    deviceScaleFactor: 1,
    timezoneId: 'America/Chicago',
    reducedMotion: 'reduce',
  },
});
