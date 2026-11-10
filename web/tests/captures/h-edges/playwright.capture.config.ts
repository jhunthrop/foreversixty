// web/tests/captures/h-edges/playwright.capture.config.ts
// Batch H lane h-edges (design/reviews/2026-11-09-h-edges): the guild landing, join, member
// profile and night report, the admin pages, the no-script character page and the S11 system
// captures, at 390, 1024, 1440 and 2000 (2560 for the system). Run from web/ against a fixture build:
//   FOREVER_DATA=fixture PUBLIC_API_BASE_URL=http://127.0.0.1:9 npx astro build
//   npx astro preview --port 4377
//   CAPTURE_PORT=4377 CAPTURE_PHASE=after npx playwright test -c tests/captures/h-edges/playwright.capture.config.ts --workers=1
// PNGs go to design/reviews/2026-11-09-h-edges/<phase>/ and design/reviews/2026-11-09-redesign/system/.
import { defineConfig } from '@playwright/test';
import { HERMETIC_LAUNCH_ARGS } from '../../hermetic';

const PORT = process.env.CAPTURE_PORT ?? '4377';
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
