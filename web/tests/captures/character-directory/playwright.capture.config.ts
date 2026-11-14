// web/tests/captures/character-directory/playwright.capture.config.ts
// Run from web/ with a fixture build served on CAPTURE_PORT (default 4486):
//   FOREVER_DATA=fixture PUBLIC_API_BASE_URL=http://127.0.0.1:9 npm run build
//   ASTRO_PREVIEW_BACKGROUND=1 npx astro preview --port 4486 &
//   CD_ROUND=r3 npx playwright test -c tests/captures/character-directory/playwright.capture.config.ts
// PNGs go to design/reviews/2026-11-14-character-directory/<CD_ROUND>/. CD_PROD=1 adds the live
// /remove-my-character page (read only, no form is sent, its API unreachable).
import { defineConfig } from '@playwright/test';
import { HERMETIC_LAUNCH_ARGS } from '../../hermetic';

const PORT = process.env.CAPTURE_PORT ?? '4486';

export default defineConfig({
  testDir: '.',
  testMatch: 'capture*.spec.ts',
  timeout: 30_000,
  workers: 4,
  reporter: 'list',
  use: {
    baseURL: `http://localhost:${PORT}`,
    // The production API's host does not resolve (tests/hermetic.ts); the live page reads as signed out.
    launchOptions: { args: [...HERMETIC_LAUNCH_ARGS] },
    browserName: 'chromium',
    deviceScaleFactor: 1,
    timezoneId: 'America/Chicago',
    reducedMotion: 'reduce',
  },
});
