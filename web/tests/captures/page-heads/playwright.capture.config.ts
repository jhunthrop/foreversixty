// web/tests/captures/page-heads/playwright.capture.config.ts
// The unified page head (design/reviews/2026-11-09-sim-head): every page's head, signed in with a
// realistic character and signed out, at 1440, 390 and 2000 for a sample. Run from web/ against a
// fixture build served on a private port:
//   FOREVER_DATA=fixture PUBLIC_API_BASE_URL=http://127.0.0.1:9 npx astro build
//   npx astro preview --port 4793
//   CAPTURE_PORT=4793 CAPTURE_PHASE=after npx playwright test -c tests/captures/page-heads/playwright.capture.config.ts --workers=1
// PNGs go to design/reviews/2026-11-09-sim-head/<phase>/.
import { defineConfig } from '@playwright/test';
import { HERMETIC_LAUNCH_ARGS } from '../../hermetic';

const PORT = process.env.CAPTURE_PORT ?? '4793';
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
