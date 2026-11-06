// web/tests/e2e/character-performance.spec.ts
// Task L8.5's rulings that outlive the character page's rebuild (design/specs/2026-11-06-character.md
// cut the Performance Rating card, the score card's part bars and the "Every ranked fight" table;
// the Forever Score Section is covered by character-score.spec.ts and character-rating.spec.ts):
//   Ruling 5. /setup says "type /fs copy, press Ctrl+C".
//   Fix round 1 item 5. The /ratings sample rating card stays one column in its 560px container.
import { expect, test } from './support/test';

test('/setup says "type /fs copy, press Ctrl+C" (L8.5 ruling 5)', async ({ page }) => {
  await page.goto('/setup');
  await expect(
    page.getByText(
      'In game, type /fs copy, press Ctrl+C, then paste it into the planner’s Import from addon box.',
    ),
  ).toHaveCount(1);
  await expect(page.getByText('/fs export')).toHaveCount(0);
});

test.describe('the /ratings sample at 1440 (fix round 1 item 5)', () => {
  test.use({ viewport: { width: 1440, height: 900 } });
  test.skip(() => test.info().project.name !== 'desktop', 'a desktop width');

  test('stays one column in its 560px container, the sparkline under the line', async ({ page }) => {
    await page.goto('/ratings');
    const sample = page.getByTestId('ratings-rating-sample');
    await expect(sample.getByTestId('rating-card')).toBeVisible();
    const boxes = await sample.evaluate((element) => {
      const box = (testid: string) =>
        element.querySelector(`[data-testid="${testid}"]`)!.getBoundingClientRect();
      const main = box('rating-card-main');
      const side = box('rating-card-side');
      return { mainBottom: main.bottom, mainLeft: main.left, sideTop: side.top, sideLeft: side.left };
    });
    expect(boxes.sideTop).toBeGreaterThanOrEqual(boxes.mainBottom - 0.5);
    expect(Math.abs(boxes.sideLeft - boxes.mainLeft)).toBeLessThan(1);
  });
});
