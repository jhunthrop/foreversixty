// design/mocks/2026-10-29-site-alignment/playwright.capture.config.ts
// The before captures for the site-alignment spec (design/specs/2026-10-29-site-alignment.md).
// Run from web/ with the fixture site served on CAPTURE_PORT (default 4391):
//   ln -s ../../../web/node_modules ../design/mocks/2026-10-29-site-alignment/node_modules (once, uncommitted), then
//   npx playwright test -c ../design/mocks/2026-10-29-site-alignment/playwright.capture.config.ts
// No import here, so the config resolves from design/; the node_modules link lets the spec find @playwright/test.
const PORT = process.env.CAPTURE_PORT ?? '4391';

export default {
  testDir: '.',
  testMatch: 'capture*.spec.ts',
  timeout: 90_000,
  workers: 4,
  reporter: 'list',
  use: {
    baseURL: `http://localhost:${PORT}`,
    browserName: 'chromium',
    deviceScaleFactor: 1,
    timezoneId: 'America/Chicago',
    reducedMotion: 'reduce',
  },
};
