// web/tests/captures/discord-notifications/playwright.capture.config.ts
// Run from web/ with the fixture site served on port 4411:
//   FOREVER_DATA=fixture npx astro build --outDir <dir> && npx astro preview --outDir <dir> --port 4411
//   npx playwright test -c tests/captures/discord-notifications/playwright.capture.config.ts
// It lives under web/ so that `@playwright/test` resolves; the PNGs go to
// design/reviews/2026-10-25-discord-notifications/ (see capture.spec.ts).
import { defineConfig } from '@playwright/test';
import { HERMETIC_LAUNCH_ARGS } from '../../hermetic';

const PORT = process.env.CAPTURE_PORT ?? '4411';
const TEST_TIMEOUT_MS = 60_000;

export default defineConfig({
  testDir: '.',
  testMatch: 'capture.spec.ts',
  timeout: TEST_TIMEOUT_MS,
  workers: 2,
  reporter: 'list',
  use: {
    // The production API's host does not resolve in this browser (tests/hermetic.ts).
    launchOptions: { args: [...HERMETIC_LAUNCH_ARGS] },
    baseURL: `http://localhost:${PORT}`,
    browserName: 'chromium',
    deviceScaleFactor: 1,
    // The 160 ms reveal fade and hover transitions settle at once, so every capture is the end state.
    reducedMotion: 'reduce',
  },
});
