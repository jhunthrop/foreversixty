// web/tests/e2e/report-brush.spec.ts
// The time window is the report's primary interaction, so the three ways into it -- the
// drag, the two range inputs a keyboard uses, and the presets -- all have to write the
// same URL and move the same label.
import { expect, test } from '@playwright/test';

const FIGHT = '/reports/fixture2abcd?fight=3';

test('dragging across the chart sets the window in the URL and rescopes the page', async ({ page }) => {
  await page.goto(FIGHT);
  await expect(page.getByTestId('window-label')).toContainText('Whole fight');

  const canvas = page.getByTestId('time-chart-canvas');
  const box = (await canvas.boundingBox())!;
  await page.mouse.move(box.x + box.width * 0.1, box.y + box.height / 2);
  await page.mouse.down();
  await page.mouse.move(box.x + box.width * 0.35, box.y + box.height / 2, { steps: 8 });
  await page.mouse.up();

  await expect(page).toHaveURL(/start=\d+&end=\d+/);
  await expect(page.getByTestId('window-label')).not.toContainText('Whole fight');
});

test('the range inputs move the same window, so a keyboard can brush', async ({ page }) => {
  await page.goto(FIGHT);
  await page.getByTestId('window-start').focus();
  for (let press = 0; press < 5; press += 1) await page.keyboard.press('ArrowRight');
  await expect(page).toHaveURL(/start=5000/);
  await expect(page.getByTestId('window-label')).toContainText('5.0s to');
});

test('a death preset sets the window to the seconds before it, and Whole fight clears it', async ({
  page,
}) => {
  await page.goto(FIGHT);
  await page.getByRole('button', { name: 'Before Thalgrit died' }).click();
  await expect(page).toHaveURL(/start=0&end=10100/);

  await page.getByTestId('window-reset').click();
  await expect(page).not.toHaveURL(/start=/);
  await expect(page.getByTestId('window-label')).toContainText('Whole fight');
});

test('switching fight drops the window, because milliseconds do not carry across', async ({ page }) => {
  await page.goto(`${FIGHT}&start=1000&end=5000`);
  await page.getByTestId('toggle-trash').click();
  await page.getByTestId('fight-2').click();
  await expect(page).not.toHaveURL(/start=/);
});

// 360x800 is the narrowest phone the design system covers, and the project skip keeps this
// to one run rather than one per device.
test.describe('the chart on a phone', () => {
  test.use({ viewport: { width: 360, height: 800 } });

  test('the chart is legible and reachable at 360px', async ({ page }, testInfo) => {
    test.skip(testInfo.project.name !== 'mobile', 'phone layout');
    await page.goto(FIGHT);
    const box = (await page.getByTestId('time-chart').boundingBox())!;
    expect(box.width).toBeLessThanOrEqual(360);
    const reset = (await page.getByTestId('window-reset').boundingBox())!;
    expect(reset.height).toBeGreaterThanOrEqual(44);
    expect(await page.evaluate(() => document.documentElement.scrollWidth <= window.innerWidth)).toBe(true);
  });
});
