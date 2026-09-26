// web/tests/e2e/report-brush.spec.ts
// The time window is the report's primary interaction, so the three ways into it -- the
// drag, the two range inputs a keyboard uses, and the presets -- all have to write the
// same URL and move the same label.
import { expect, test } from '@playwright/test';
import { openChartIfCollapsed } from './support/report-chart';

const FIGHT = '/reports/fixture2abcd?fight=3';

test('dragging across the chart sets the window in the URL and rescopes the page', async ({ page }) => {
  await page.goto(FIGHT);
  await openChartIfCollapsed(page);
  await expect(page.getByTestId('window-label')).toContainText('Whole fight');

  const canvas = page.getByTestId('time-chart-canvas');
  // On a phone the chart sits below the fight list; a drag at coordinates off the screen
  // lands nowhere.
  await canvas.scrollIntoViewIfNeeded();
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
  await openChartIfCollapsed(page);
  await page.getByTestId('window-start').focus();
  for (let press = 0; press < 5; press += 1) await page.keyboard.press('ArrowRight');
  await expect(page).toHaveURL(/start=5000/);
  await expect(page.getByTestId('window-label')).toContainText('5.0s to');
});

test('a death preset sets the window to the seconds before it, and Whole fight clears it', async ({
  page,
}) => {
  await page.goto(FIGHT);
  await openChartIfCollapsed(page);
  const select = page.getByTestId('window-select');
  // The label carries the death's own timestamp (a battle-rez can die twice), so this
  // reads the option's value off its text rather than pinning the exact suffix.
  const value = await select
    .locator('option')
    .filter({ hasText: 'before Thalgrit died' })
    .getAttribute('value');
  await select.selectOption(value!);
  await expect(page).toHaveURL(/start=0&end=11000/);

  // The select's first option is the whole fight -- the old separate reset button's job.
  await page.getByTestId('window-select').selectOption({ index: 0 });
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

  test('the chart is collapsed by default, and opens to a legible, reachable 360px chart', async ({
    page,
  }, testInfo) => {
    test.skip(testInfo.project.name !== 'mobile', 'phone layout');
    await page.goto(FIGHT);
    // Collapsed by default (design review 2026-09-26 finding 1): a one-line strip with the
    // window summary and a 44px toggle, not the chart, its sliders and six preset buttons.
    await expect(page.getByTestId('time-chart')).toHaveCount(0);
    const strip = page.getByTestId('chart-collapsed');
    await expect(strip).toContainText('Whole fight');
    const toggle = page.getByTestId('chart-toggle');
    const toggleBox = (await toggle.boundingBox())!;
    expect(toggleBox.height).toBeGreaterThanOrEqual(44);

    await toggle.click();
    const box = (await page.getByTestId('time-chart').boundingBox())!;
    expect(box.width).toBeLessThanOrEqual(360);
    const select = (await page.getByTestId('window-select').boundingBox())!;
    expect(select.height).toBeGreaterThanOrEqual(44);
    // The brush handles are named by type in the 44px rule. A range input drags from a
    // press anywhere inside its own box, so the box is the hit target -- the wrapping
    // label's min-h-11 does nothing for it.
    for (const handle of ['window-start', 'window-end']) {
      const box = (await page.getByTestId(handle).boundingBox())!;
      expect(box.height, `${handle} hit target`).toBeGreaterThanOrEqual(44);
    }
    expect(await page.evaluate(() => document.documentElement.scrollWidth <= window.innerWidth)).toBe(true);
  });

  test('the choice to open the chart survives a reload', async ({ page }, testInfo) => {
    test.skip(testInfo.project.name !== 'mobile', 'phone layout');
    await page.goto(FIGHT);
    await page.getByTestId('chart-toggle').click();
    await expect(page.getByTestId('time-chart')).toBeVisible();

    await page.reload();
    await expect(page.getByTestId('time-chart')).toBeVisible();
    await expect(page.getByTestId('chart-collapsed')).toHaveCount(0);
  });
});
