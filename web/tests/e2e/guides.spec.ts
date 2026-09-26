// web/tests/e2e/guides.spec.ts
// Spec section 9 (lane C): "a spec that Load this build lands in the planner with the
// points lit."
//
// The suite defaults to FOREVER_DATA=fixture (see playwright.config.ts's own comment and
// real-data.spec.ts), because every other spec asserts on the fixture's fixed talent names,
// item ids and counts. Guide pages break that assumption: their `build:` frontmatter carries
// a real-build FS1 code (see src/content/guides/warrior/fury.md), because a guide states real
// talent facts, not the fixture's small two-tree warrior. Decoding that code against the
// fixture's different talent ids would either fail to parse or light the wrong cell -- the
// fixture's warrior tree does not share Real Fury's talent ids or tier layout -- so, like
// real-data.spec.ts, this spec only runs meaningfully under FOREVER_DATA=real and skips
// otherwise (`FOREVER_DATA=real E2E_PORT=4387 npx playwright test tests/e2e/guides.spec.ts`).
// active-build.json already names the same build the guides were written against
// (1.60.1.69893), so no extra sync step beyond the normal `prebuild` hook is needed -- just
// the env var that picks which source `sync-data.mjs` publishes.
import { test, expect } from '@playwright/test';

test.skip(
  process.env.FOREVER_DATA !== 'real',
  'guide build-load smoke; the fixture talent tree cannot legally hold a real-build FS1 code. Run with FOREVER_DATA=real, e.g. FOREVER_DATA=real E2E_PORT=4387 npx playwright test tests/e2e/guides.spec.ts',
);

test('Load this build on the Fury Warrior guide lands in the planner with points already spent', async ({
  page,
}) => {
  await page.goto('/guides/warrior/fury');
  const loadLink = page.getByTestId('guide-load-build');
  await expect(loadLink).toBeVisible();
  const href = await loadLink.getAttribute('href');
  expect(href).toMatch(/^\/planner\?code=/);

  await loadLink.click();
  await expect(page).toHaveURL(/\/planner\?code=/);

  // Bloodthirst is Fury's own named capstone (tier 6, single rank, tab-order's last talent
  // in the tree) -- fury.md's `build:` FS1 code spends its Fury segment's final, untrimmed
  // digit on it, so it comes back lit at rank 1 once the planner decodes the loaded code
  // against the real talent data this same build id ships.
  const bloodthirstCell = page.locator('[data-testid^="talent-"][aria-label^="Bloodthirst"]');
  await expect(bloodthirstCell).toHaveAttribute('data-rank', '1');
});
