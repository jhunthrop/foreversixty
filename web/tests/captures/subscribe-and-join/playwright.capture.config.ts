// web/tests/captures/subscribe-and-join/playwright.capture.config.ts
// Run from web/ with the fixture site (a COPY of dist) served on a private port (Astro allows one
// preview per project, so kill it by PID afterwards):
//   cp -R dist /tmp/sj-dist && npx astro preview --outDir /tmp/sj-dist --port 4483   # in the background
//   CAPTURE_PORT=4483 npx playwright test -c tests/captures/subscribe-and-join/playwright.capture.config.ts
// Row 17 needs the baseline first: build `main`, serve it the same way and run with
//   CAPTURE_VARIANT=before   # writes the "before" PNGs; the default run writes "after" and diffs them
// PNGs go to design/reviews/2026-10-27-subscribe-and-join/ (see capture.spec.ts).
import { defineConfig } from '@playwright/test';
import { HERMETIC_LAUNCH_ARGS } from '../../hermetic';

const PORT = process.env.CAPTURE_PORT ?? '4483';
const TEST_TIMEOUT_MS = 90_000;

export default defineConfig({
  testDir: '.',
  testMatch: 'capture*.spec.ts',
  timeout: TEST_TIMEOUT_MS,
  workers: 4,
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
