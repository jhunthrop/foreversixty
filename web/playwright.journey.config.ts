// web/playwright.journey.config.ts
// The one-character journey against the real local API (tools/dev-stack) through installed Chrome.
// It serves nothing: start `npm run build` (PUBLIC_API_BASE_URL=http://localhost:8080) and
// `astro preview --port 4331` yourself, then `W11_JOURNEY=1 npx playwright test -c playwright.journey.config.ts`.
import { defineConfig } from '@playwright/test';
import { HERMETIC_LAUNCH_ARGS } from './tests/hermetic';

export default defineConfig({
  testDir: './tests/e2e',
  testMatch: /one-character-journey.spec.ts/,
  timeout: 120_000,
  workers: 1,
  use: {
    // The production API's host does not resolve in this browser (tests/hermetic.ts).
    launchOptions: { args: [...HERMETIC_LAUNCH_ARGS] },
    baseURL: process.env.W11_WEB ?? 'http://localhost:4331',
    channel: 'chrome',
    viewport: { width: 1280, height: 900 },
  },
});
