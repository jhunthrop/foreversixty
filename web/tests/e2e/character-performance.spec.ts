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

test.describe('the /ratings sample at 1440 (logs spec 2026-11-07 section 8)', () => {
  test.use({ viewport: { width: 1440, height: 900 } });
  test.skip(() => test.info().project.name !== 'desktop', 'a desktop width');

  // The trend reads as the last fights' numbers in a line since the report tail's round 5 (no sparkline).
  test('reads as three figures on one line', async ({ page }) => {
    await page.goto('/ratings');
    const sample = page.getByTestId('ratings-rating-sample');
    await expect(sample).toContainText('65');
    const tops = await sample.evaluate((element) =>
      [...element.children].map((child) => Math.round(child.getBoundingClientRect().top)),
    );
    expect(new Set(tops).size).toBe(1);
  });
});
