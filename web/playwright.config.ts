// web/playwright.config.ts
import { defineConfig, devices } from '@playwright/test';

// E2E_PORT lets a second checkout run its own preview beside another one (the persona
// review harness keeps the main build on 4321), so the suite is not pinned to one port.
const port = process.env.E2E_PORT ?? '4321';

export default defineConfig({
  testDir: './tests/e2e',
  // 30 s is right for everything but tests/e2e/report-queries.spec.ts, where the first
  // query downloads and instantiates DuckDB's WebAssembly build. Those tests call
  // test.slow(), which triples whatever this is, so 45 s gives them 135 s on a cold CI
  // runner without slowing anything else down.
  timeout: 45_000,
  use: { baseURL: `http://localhost:${port}`, trace: 'retain-on-failure' },
  // ASTRO_PREVIEW_BACKGROUND disables Astro's auto-detected-AI-agent daemon mode for `astro preview`
  // (see astro/dist/cli/preview/index.js `isRunByAgent()`), which otherwise forks preview into the
  // background and exits the foreground process immediately -- Playwright then reports
  // "Process from config.webServer exited early" even though the server is actually up.
  // FOREVER_DATA picks which planner data the built site carries (see scripts/sync-data.mjs).
  // The suite defaults to the fixture: every spec but real-data.spec.ts asserts on the
  // fixture's talent names, item ids and counts, which the real pipeline output moves. Run
  // `npm run test:e2e:real` to build with real data and run the smoke suite against it.
  webServer: {
    command: `npm run build && npm run preview -- --port ${port}`,
    port: Number(port),
    reuseExistingServer: !process.env.CI,
    timeout: 180_000,
    env: { ASTRO_PREVIEW_BACKGROUND: '1', FOREVER_DATA: process.env.FOREVER_DATA ?? 'fixture' },
  },
  projects: [
    { name: 'desktop', use: { ...devices['Desktop Chrome'] } },
    { name: 'mobile', use: { ...devices['Pixel 7'] } },
  ],
});
