// web/tests/captures/paperdoll/playwright.capture.config.ts
// Acceptance captures for the paperdoll (design/specs/2026-10-31-paperdoll.md §9), against a
// site built with the real active build's data (FOREVER_DATA=real) and a dead API base
// (PUBLIC_API_BASE_URL=http://127.0.0.1:9), served on 4482:
//   FOREVER_DATA=real PUBLIC_API_BASE_URL=http://127.0.0.1:9 npm run build
//   cp -R dist <dir> && npx astro preview --outDir <dir> --port 4482   # in the background
//   npx playwright test -c tests/captures/paperdoll/playwright.capture.config.ts
// PNGs go to design/reviews/2026-10-31-paperdoll/build/ (see capture.spec.ts).
import { defineConfig } from '@playwright/test';

const PORT = process.env.CAPTURE_PORT ?? '4482';

export default defineConfig({
  testDir: '.',
  testMatch: 'capture*.spec.ts',
  timeout: 120_000,
  workers: 4,
  reporter: 'list',
  use: {
    baseURL: `http://localhost:${PORT}`,
    browserName: 'chromium',
    deviceScaleFactor: 2,
    timezoneId: 'America/Chicago',
    reducedMotion: 'reduce',
  },
});
