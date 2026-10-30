// web/tests/e2e/report-brush.spec.ts
// The time window is the report's primary interaction, so the three ways into it -- the
// drag, the two range inputs a keyboard uses, and the presets -- all have to write the
// same URL and move the same label.
import { expect, test } from '@playwright/test';
import { openChartIfCollapsed } from './support/report-chart';

const FIGHT = '/reports/fixture2abcd?fight=3';
const FIG = `${FIGHT}&tab=damage-done`;

test('dragging across the chart sets the window in the URL and rescopes the page', async ({
  page,
}, testInfo) => {
  // Drag-to-zoom is a fine pointer's (spec §B6.8: a phone zooms with the brush and the presets).
  test.skip(testInfo.project.name === 'mobile', 'the plot takes no drag on touch');
  await page.goto(FIG);
  await openChartIfCollapsed(page);
  const chart = page.getByTestId('time-chart');
  await expect(chart.getByTestId('window-select')).toHaveValue('0');
  await expect(chart.locator('[data-testid="chart-frame"][data-state="ready"]')).toBeVisible();

  const plot = chart.getByTestId('chart-plot');
  await plot.scrollIntoViewIfNeeded();
  const box = (await plot.boundingBox())!;
  await page.mouse.move(box.x + box.width * 0.2, box.y + box.height / 2);
  await page.mouse.down();
  await page.mouse.move(box.x + box.width * 0.6, box.y + box.height / 2, { steps: 8 });
  // The drag names its span while it is in progress (spec §B6.2).
  await expect(chart.getByTestId('drag-label')).toContainText(/\d:\d\d to \d:\d\d · \d+ s/);
  await page.mouse.up();

  await expect(page).toHaveURL(/start=\d+&end=\d+/);
  // The Window select names the dragged window, and the brush shows its edges.
  await expect(chart.getByTestId('window-select')).toHaveValue('custom');
  await expect(chart.getByTestId('window-select')).toContainText(/\d:\d\d to \d:\d\d/);
  await expect(chart.getByTestId('brush-start-label')).toBeVisible();
});

test('zoom keeps the plot, the brush and the URL on one window, and double-click returns', async ({
  page,
}, testInfo) => {
  test.skip(testInfo.project.name === 'mobile', 'the plot takes no drag on touch');
  await page.goto(FIG);
  const chart = page.getByTestId('time-chart');
  await expect(chart.locator('[data-testid="chart-frame"][data-state="ready"]')).toBeVisible();
  const plot = chart.getByTestId('chart-plot');
  const box = (await plot.boundingBox())!;
  await page.mouse.move(box.x + box.width * 0.3, box.y + box.height / 2);
  await page.mouse.down();
  await page.mouse.move(box.x + box.width * 0.7, box.y + box.height / 2, { steps: 6 });
  await page.mouse.up();
  await expect(page).toHaveURL(/start=(\d+)&end=(\d+)/);
  const url = new URL(page.url());
  const start = Number(url.searchParams.get('start'));
  const end = Number(url.searchParams.get('end'));
  // The brush's handles sit on the same window the URL carries.
  await expect(chart.getByTestId('window-start')).toHaveValue(String(start));
  await expect(chart.getByTestId('window-end')).toHaveValue(String(end));
  // The plot's accessible name reads the window it is zoomed to.
  await expect(plot).toHaveAttribute('aria-label', /chart, 0:\d\d to 0:\d\d\. Arrow keys/);
  // Moving the brush's handle moves the URL too.
  await chart.getByTestId('window-start').focus();
  await page.keyboard.press('ArrowLeft');
  await expect(page).toHaveURL(new RegExp(`start=${start - 1000}&`));
  // A double-click on the plot is the whole fight again.
  await plot.dblclick({ position: { x: box.width / 2, y: box.height / 2 } });
  await expect(page).not.toHaveURL(/start=/);
  await expect(chart.getByTestId('window-select')).toHaveValue('0');
});

test('the range inputs move the same window, so a keyboard can brush', async ({ page }) => {
  await page.goto(FIG);
  await openChartIfCollapsed(page);
  await page.getByTestId('window-start').focus();
  for (let press = 0; press < 5; press += 1) await page.keyboard.press('ArrowRight');
  await expect(page).toHaveURL(/start=5000/);
  await expect(page.getByTestId('brush-start-label')).toHaveText('0:05');
  // Shift moves ten seconds, End goes to the fight's end.
  await page.keyboard.press('Shift+ArrowRight');
  await expect(page).toHaveURL(/start=15000/);
});

test('a death preset sets the window to the seconds before it, and Whole fight clears it', async ({
  page,
}) => {
  await page.goto(FIG);
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
  await expect(page.getByTestId('brush-start-label')).toHaveCount(0);
});

test('switching fight drops the window, because milliseconds do not carry across', async ({ page }) => {
  await page.goto(`${FIGHT}&start=1000&end=5000`);
  await page.getByTestId('toggle-trash').click();
  await page.getByTestId('fight-group-1').click();
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
