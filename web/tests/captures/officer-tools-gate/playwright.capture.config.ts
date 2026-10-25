// web/tests/captures/officer-tools-gate/playwright.capture.config.ts
// Run from web/ with the fixture site served on port 4431:
//   FOREVER_DATA=fixture npx astro build --outDir <dir> && npx astro preview --outDir <dir> --port 4431
//   npx playwright test -c tests/captures/officer-tools-gate/playwright.capture.config.ts
// It lives under web/ so that `@playwright/test` resolves; the PNGs go to
// design/reviews/2026-10-25-officer-tools-gate/ (see capture.spec.ts).
import { defineConfig } from '@playwright/test';

const PORT = process.env.CAPTURE_PORT ?? '4431';
const TEST_TIMEOUT_MS = 60_000;

export default defineConfig({
  testDir: '.',
  testMatch: 'capture.spec.ts',
  timeout: TEST_TIMEOUT_MS,
  workers: 2,
  reporter: 'list',
  use: {
    baseURL: `http://localhost:${PORT}`,
    browserName: 'chromium',
    deviceScaleFactor: 1,
    // The 160 ms reveal fade and hover transitions settle at once, so every capture is the end state.
    reducedMotion: 'reduce',
  },
});
