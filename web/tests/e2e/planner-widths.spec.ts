// web/tests/e2e/planner-widths.spec.ts
// Fix round 6 (owner: signed-in-1024.png rendered 1210px wide for a 1024 viewport -- the
// page scrolled horizontally at laptop widths): the header band's own full-bleed breakout
// (`width: 100vw; margin-left: calc(50% - 50vw)`, PlannerHeaderBand.svelte) combined with
// content that did not yet give way at 1024 (tree panels squeezed into 8 of 12 columns, a
// stretched facts rail, the site nav's own `flex-nowrap`) pushed the document wider than
// the viewport. Asserts the one invariant that actually matters for all of that at once:
// the page never scrolls sideways, at every width this lane was asked to check, signed out
// and signed in (the state the designer's own screenshot caught; signed in, the rail's
// In the game section holds the send, spec 2026-11-08 4.7).
import { expect, test } from './support/test';
import { meAddonFixture } from '../../src/fixtures/me-addon';

const WIDTHS = [1024, 1280, 1440, 1920];

async function noHorizontalScroll(page: import('@playwright/test').Page): Promise<void> {
  const { scrollWidth, clientWidth } = await page.evaluate(() => ({
    scrollWidth: document.documentElement.scrollWidth,
    clientWidth: document.documentElement.clientWidth,
  }));
  expect(scrollWidth, `scrollWidth (${scrollWidth}) should equal clientWidth (${clientWidth})`).toBe(
    clientWidth,
  );
}

test.describe('the planner never scrolls sideways', () => {
  for (const width of WIDTHS) {
    test(`signed out at ${width}px`, async ({ page }) => {
      await page.setViewportSize({ width, height: 1000 });
      await page.route('**/v1/me', (route) =>
        route.fulfill({
          status: 401,
          contentType: 'application/json',
          body: JSON.stringify({ ok: false, data: null, error: null, request_id: 'e2e' }),
        }),
      );
      await page.goto('/planner');
      await page.getByTestId('band-compare').waitFor();
      await noHorizontalScroll(page);
    });

    test(`signed in at ${width}px`, async ({ page }) => {
      await page.setViewportSize({ width, height: 1000 });
      // meAddonFixture's own `main_character_key` names Zulmara -- the `mainCharacter()`
      // fallback finds her with no stored pointer at all, the same way a signed-in visitor with
      // no current-character pointer yet still has a character to send to.
      await page.route('**/v1/me', (route) =>
        route.fulfill({
          status: 200,
          contentType: 'application/json',
          body: JSON.stringify({ ok: true, data: meAddonFixture, error: null, request_id: 'e2e' }),
        }),
      );
      await page.goto('/planner');
      await page.getByTestId('band-compare').waitFor();
      // The send is the whole point of this width/mode combination -- confirms the signed-in
      // shape (not the signed-out one) is actually what got measured below, not a fallback
      // that happened to also pass.
      await page.getByTestId('planner-send').waitFor();
      await noHorizontalScroll(page);
    });
  }
});
