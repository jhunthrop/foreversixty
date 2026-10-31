// web/tests/captures/raid-spec-ask/playwright.capture.config.ts
// Run from web/ with the fixture site served (any static server of a FOREVER_DATA=fixture build):
//   CAPTURE_PORT=4392 npx playwright test -c tests/captures/raid-spec-ask/playwright.capture.config.ts
import { defineConfig } from '@playwright/test';
import { HERMETIC_LAUNCH_ARGS } from '../../hermetic';

const PORT = process.env.CAPTURE_PORT ?? '4441';

export default defineConfig({
  testDir: '.',
  testMatch: 'capture*.spec.ts',
  timeout: 60_000,
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
