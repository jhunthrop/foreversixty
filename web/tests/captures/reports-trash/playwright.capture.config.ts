// web/tests/captures/reports-trash/playwright.capture.config.ts
// Run from web/ with the fixture site served on CAPTURE_PORT (4707 by default):
//   FOREVER_DATA=fixture npm run sync && FOREVER_DATA=fixture npm run build
//   npx astro preview --port 4707   # in the background
//   npx playwright test -c tests/captures/reports-trash/playwright.capture.config.ts
// The PNGs go to design/reviews/2026-10-31-reports-trash/ (see capture.spec.ts).
import { defineConfig } from '@playwright/test';

const PORT = process.env.CAPTURE_PORT ?? '4707';
const TEST_TIMEOUT_MS = 60_000;

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
