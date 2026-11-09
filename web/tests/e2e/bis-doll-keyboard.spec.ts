// web/tests/e2e/bis-doll-keyboard.spec.ts
// The paperdoll on the keyboard: every filled slot is a button in the client's order (head down
// the left, hands down the right, the weapons along the bottom); focus selects it and the centre
// shows its tooltip, as a hover does.
import { test, expect } from './support/test';
import { openBis } from './support/bis-page';

const URL = '/bis/hunter/marksmanship?band=60';

test('Tab walks the slots in the client order and focus swaps the centre', async ({ page, isMobile }) => {
  test.skip(isMobile, 'the frame shows from 423 px; a phone has the list form');
  await page.setViewportSize({ width: 1440, height: 1000 });
  await openBis(page, URL);
  const head = page.getByTestId('paperdoll-frame-slot-head');
  await head.focus();
  await expect(head).toHaveAttribute('aria-pressed', 'true');
  await page.keyboard.press('Tab');
  const next = page.locator('[data-testid^="paperdoll-frame-slot-"]:focus');
  await expect(next).toHaveAttribute('data-testid', 'paperdoll-frame-slot-neck');
  await expect(next).toHaveAttribute('aria-pressed', 'true');
  const item = (await next.getAttribute('aria-label'))!.split(': ')[1]!;
  await expect(page.locator('[data-region="paperdoll-centre"]').getByTestId('item-tooltip')).toContainText(
    item,
  );
});
