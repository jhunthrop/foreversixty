// web/tests/captures/premium-entry/playwright.capture.config.ts
// Acceptance screenshots for design/specs/2026-10-29-premium-entry.md. Run from web/ against a
// fixture build served on a private port (a separate outDir, never the shared dist):
//   npx astro build --outDir dist-premium && npx astro preview --outDir dist-premium --port 4791
//   CAPTURE_PORT=4791 npx playwright test -c tests/captures/premium-entry/playwright.capture.config.ts
// PNGs go to design/reviews/2026-10-29-premium-entry/.
import { defineConfig } from '@playwright/test';
import { HERMETIC_LAUNCH_ARGS } from '../../hermetic';

const PORT = process.env.CAPTURE_PORT ?? '4791';
const TEST_TIMEOUT_MS = 90_000;

export default defineConfig({
  testDir: '.',
  testMatch: 'capture*.spec.ts',
  timeout: TEST_TIMEOUT_MS,
  workers: 2,
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
