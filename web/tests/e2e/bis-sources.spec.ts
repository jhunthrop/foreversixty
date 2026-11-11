// web/tests/e2e/bis-sources.spec.ts
// The Sources filter (web layout spec 2026-11-05 §5.11): turning a kind off re-picks the slots
// it supplied from what remains (nothing is hidden), the choice is in the URL and remembered per
// browser, and the head's figure is labelled the unfiltered one when no filtered set was simmed.
import type { Page } from '@playwright/test';
import { test, expect } from './support/test';
import { openBis } from './support/bis-page';

const PAGE = '/bis/hunter/marksmanship?band=60&view=slot';

async function picks(page: Page): Promise<Map<string, { name: string; kind: string | null }>> {
  const rows = await page
    .locator('[data-testid="bis-doll"] [data-doll-slot][data-state="item"]')
    .evaluateAll((cells) =>
      cells.map((cell) => [
        (cell as HTMLElement).dataset.dollSlot ?? '',
        cell.querySelector('[data-testid^="item-hover-"]')?.textContent?.trim() ?? '',
        cell.querySelector('[data-source-kind]')?.getAttribute('data-source-kind') ?? null,
      ]),
    );
  return new Map(rows.map(([slot, name, kind]) => [slot!, { name: name!, kind }]));
}

test('Dungeon only re-picks every slot from dungeons, hides none, and labels the figure', async ({
  page,
}) => {
  await openBis(page, PAGE);
  const all = await picks(page);
  const outside = [...all].filter(([, pick]) => pick.kind !== 'dungeon');
  test.skip(outside.length === 0, 'every pick at 60 is a dungeon drop with the current data');

  await openBis(page, `${PAGE}&sources=dungeon`);
  await expect(page.locator('[data-testid="bis-doll"] [data-doll-slot]')).toHaveCount(17);
  const dungeon = await picks(page);
  for (const [, pick] of dungeon) expect(pick.kind).toBe('dungeon');
  const changed = outside.filter(([slot]) => dungeon.get(slot)?.name !== all.get(slot)?.name);
  expect(changed.length).toBeGreaterThan(0);
  for (const [slot] of changed) {
    const cell = page.getByTestId(`bis-doll-slot-${slot}`);
    // Option C: a re-picked slot names the best overall pick under its gain; the why is in its detail.
    if ((await cell.getAttribute('data-state')) === 'item')
      await expect(cell).toContainText(/vs best: \S|as \S/);
  }
  await expect(page.getByTestId('bis-answer-figure')).toContainText(
    /Set: dungeon only|Best set, all sources/i,
  );
});

test('from 768 the Sources button opens a checklist popover; Escape closes it and focus returns', async ({
  page,
  isMobile,
}) => {
  test.skip(isMobile, 'the phone keeps the chips in the expanded control stack');
  await openBis(page, PAGE);
  const button = page.getByTestId('bis-sources-button');
  await expect(button).toHaveText('All sources');
  await button.click();
  await expect(button).toHaveAttribute('aria-expanded', 'true');
  const popover = page.getByTestId('bis-sources-popover');
  await expect(popover).toBeVisible();
  await expect(popover.getByRole('checkbox')).toHaveCount(8);
  await expect(popover).toContainText('Counts are gear slots with an option');

  const quest = page.getByTestId('bis-source-check-quest').getByRole('checkbox');
  await quest.uncheck();
  await expect(page).toHaveURL(/[?&]sources=/);
  await expect(button).toHaveText('Sources (7)');
  await expect(popover).toBeVisible();

  await page.keyboard.press('Escape');
  await expect(popover).toBeHidden();
  await expect(button).toHaveAttribute('aria-expanded', 'false');
  await expect(button).toBeFocused();
});

test.describe('on a phone', () => {
  test.use({ viewport: { width: 390, height: 844 } });

  test('a chip turns a kind off, writes ?sources= and is remembered on the next visit', async ({ page }) => {
    await openBis(page, PAGE);
    const before = await picks(page);
    const kind = [...before.values()].find((pick) => pick.kind !== null && pick.kind !== 'dungeon')?.kind;
    test.skip(kind == null, 'no non-dungeon pick at 60 with the current data');
    await page.getByTestId('bis-controls').locator('[data-control-bar-toggle]').click();
    const chip = page.getByTestId(`bis-source-chip-${kind}`);
    await expect(chip).toHaveAttribute('aria-pressed', 'true');
    await chip.click();
    await expect(chip).toHaveAttribute('aria-pressed', 'false');
    await expect(page).toHaveURL(/[?&]sources=/);
    const after = await picks(page);
    for (const [, pick] of after) expect(pick.kind).not.toBe(kind);

    await openBis(page, '/bis/hunter/marksmanship?band=60&view=slot');
    const remembered = await picks(page);
    for (const [, pick] of remembered) expect(pick.kind).not.toBe(kind);
  });
});
