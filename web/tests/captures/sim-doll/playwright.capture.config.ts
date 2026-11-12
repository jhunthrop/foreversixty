// web/tests/captures/sim-doll/playwright.capture.config.ts
// Run from web/ with a COPY of dist served on a private port (one preview at a time; kill it by
// PID afterwards):
//   cp -R dist <copy> && npx astro preview --outDir <copy> --port 4493   # in the background
//   CAPTURE_PORT=4493 npx playwright test -c tests/captures/sim-doll/playwright.capture.config.ts
import { defineConfig } from '@playwright/test';
import { HERMETIC_LAUNCH_ARGS } from '../../hermetic';

const PORT = process.env.CAPTURE_PORT ?? '4493';
const TEST_TIMEOUT_MS = 90_000;

export default defineConfig({
  testDir: '.',
  testMatch: 'capture*.spec.ts',
  timeout: TEST_TIMEOUT_MS,
  workers: 1,
  reporter: 'list',
  use: {
    launchOptions: { args: [...HERMETIC_LAUNCH_ARGS] },
    baseURL: `http://localhost:${PORT}`,
    browserName: 'chromium',
    timezoneId: 'America/Chicago',
  },
});
