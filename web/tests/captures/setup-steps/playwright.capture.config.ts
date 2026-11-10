// web/tests/captures/setup-steps/playwright.capture.config.ts
// /setup's checklist with the companion second and the addon third (owner 2026-11-10), in every
// state the review covers, at 1440 and 390. Run from web/ against a fixture build on a private port:
//   FOREVER_DATA=fixture npx astro build
//   npx astro preview --port 4794
//   CAPTURE_PORT=4794 npx playwright test -c tests/captures/setup-steps/playwright.capture.config.ts
// PNGs go to design/reviews/2026-11-10-setup-steps/.
import { defineConfig } from '@playwright/test';
import { HERMETIC_LAUNCH_ARGS } from '../../hermetic';

const PORT = process.env.CAPTURE_PORT ?? '4794';
const TEST_TIMEOUT_MS = 90_000;

export default defineConfig({
  testDir: '.',
  testMatch: 'capture*.spec.ts',
  timeout: TEST_TIMEOUT_MS,
  workers: 1,
  reporter: 'list',
  use: {
    launchOptions: { args: [...HERMETIC_LAUNCH_ARGS, '--font-render-hinting=none'] },
    baseURL: `http://localhost:${PORT}`,
    browserName: 'chromium',
    deviceScaleFactor: 1,
    timezoneId: 'America/Chicago',
    reducedMotion: 'reduce',
  },
});
