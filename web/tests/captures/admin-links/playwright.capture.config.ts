// web/tests/captures/admin-links/playwright.capture.config.ts
// Run from web/ with the fixture site served (npx astro preview --port <port>, in the background):
//   CAPTURE_PORT=<port> npx playwright test -c tests/captures/admin-links/playwright.capture.config.ts
// PNGs go to design/reviews/2026-11-02-admin-links/ (see capture.spec.ts).
import { defineConfig } from '@playwright/test';
import { HERMETIC_LAUNCH_ARGS } from '../../hermetic';

const PORT = process.env.CAPTURE_PORT ?? '4481';

export default defineConfig({
  testDir: '.',
  testMatch: 'capture*.spec.ts',
  timeout: 90_000,
  workers: 3,
  reporter: 'list',
  use: {
    launchOptions: { args: [...HERMETIC_LAUNCH_ARGS] },
    baseURL: `http://localhost:${PORT}`,
    browserName: 'chromium',
    timezoneId: 'America/Chicago',
    reducedMotion: 'reduce',
  },
});
