// web/tests/e2e/report-charts.spec.ts
// The report charts on the shared uPlot base (controls and charts spec Part B, §C, §D captures 9
// to 17). The 40-player tests serve a synthetic Ragnaros pull over fight 3 by route mocking
// (support/raid40.ts); nothing here changes the fixture the site serves.
import type { Locator, Page } from '@playwright/test';
import { expect, test } from './support/test';
import { BROKK, beLightningbolltz, guidOf, raid40DpsSum, RAID40_URL, serveRaid40 } from './support/raid40';
import { fieldBox, listOf } from './support/select';

/** The Players field's closed text (`Top 8 of 40`). */
const playersField = (chart: Locator): Locator => fieldBox(chart, 'players').locator('.sel-label');

const FIXTURE = '/reports/fixture2abcd?fight=3';
const isPhone = (page: Page): boolean => (page.viewportSize()?.width ?? 1280) < 768;

async function ready(chart: Locator): Promise<void> {
  await expect(chart.locator('[data-state="ready"]').first()).toBeVisible();
}

/** Hovers the plot at a fight time, with the plot zoomed to the whole fight. */
async function hoverAt(page: Page, chart: Locator, ms: number, durationMs: number): Promise<void> {
  const plot = chart.getByTestId('chart-plot');
  await plot.scrollIntoViewIfNeeded();
  const box = (await plot.boundingBox())!;
  const gutter = isPhone(page) ? 40 : 52;
  const width = box.width - gutter - 12;
  await page.mouse.move(box.x + gutter + (ms / durationMs) * width, box.y + box.height / 3);
}

test.describe('overlays never sit under a chart (spec §C)', () => {
  test('the open Players list is the topmost element over the plot', async ({ page }) => {
    test.skip(isPhone(page), 'a coarse pointer opens the legend sheet instead');
    await page.goto(`${FIXTURE}&tab=damage-done`);
    const chart = page.getByTestId('time-chart');
    await ready(chart);
    await chart.getByTestId('players-trigger').click();
    const list = listOf(chart, 'players');
    await expect(list).toBeVisible();
    for (const row of await list.locator('[role="option"]').all()) {
      const box = (await row.boundingBox())!;
      if (box.y + box.height / 2 > (page.viewportSize()?.height ?? 0)) continue;
      const onTop = await page.evaluate(
        ({ x, y }) => document.elementFromPoint(x, y)?.closest('.sel-pop') !== null,
        { x: box.x + box.width / 2, y: box.y + box.height / 2 },
      );
      expect(onTop, 'a Players row is covered').toBe(true);
    }
    // While a menu is open the chart stands down: a hover shows no tooltip.
    await hoverAt(page, chart, 20_000, 60_000);
    await expect(page.locator('[data-testid="chart-tooltip"].shown')).toHaveCount(0);
    await page.keyboard.press('Escape');
    await expect(list).toBeHidden();
  });

  test('the tooltip renders in the top layer, over the chart and inside the viewport', async ({ page }) => {
    test.skip(isPhone(page), 'a phone reads the event line instead of a tooltip');
    await page.goto(`${FIXTURE}&tab=damage-done`);
    const chart = page.getByTestId('time-chart');
    await ready(chart);
    await hoverAt(page, chart, 20_000, 60_000);
    const tooltip = page.getByTestId('chart-tooltip');
    await expect(tooltip).toHaveClass(/shown/);
    expect(await tooltip.evaluate((el) => el.matches(':popover-open'))).toBe(true);
    const box = (await tooltip.boundingBox())!;
    const viewport = page.viewportSize()!;
    expect(box.x).toBeGreaterThanOrEqual(8);
    expect(box.x + box.width).toBeLessThanOrEqual(viewport.width - 8);
    await expect(tooltip.getByTestId('chart-tooltip-time')).toHaveText('0:20');
  });
});

test('a death marker opens that death on the Deaths tab, open and in view', async ({ page }) => {
  await page.goto(`${FIXTURE}&tab=damage-done`);
  const toggle = page.getByTestId('chart-toggle');
  if (isPhone(page)) await toggle.click();
  const chart = page.getByTestId('time-chart');
  await ready(chart);
  const marker = chart.getByTestId('death-marker');
  await expect(marker).toHaveCount(1);
  // The marker names the killing blow and says where it goes.
  await expect(marker).toHaveAttribute(
    'aria-label',
    /^Thalgrit died · Anima Lash from Warden Kelthas · 2,400 \(100 overkill\)\. Open on the Deaths tab\.$/,
  );
  await expect(marker).toHaveAttribute('href', /tab=deaths&death=Player-4184-000000A4-10100/);
  await marker.click();
  await expect(page).toHaveURL(/tab=deaths/);
  await expect(page).toHaveURL(/death=Player-4184-000000A4-10100/);
  const card = page.locator('#death-Player-4184-000000A4-10100');
  await expect(card).toBeInViewport();
});

test.describe('the 40-player damage chart (spec D5, capture 10)', () => {
  test.beforeEach(async ({ page }) => {
    await serveRaid40(page);
    await beLightningbolltz(page);
  });

  test('draws the top 8 of 40, you pinned, the raid median, and a head that sums every player', async ({
    page,
  }) => {
    await page.goto(`${RAID40_URL}&tab=damage-done`);
    if (isPhone(page)) await page.getByTestId('chart-toggle').click();
    const chart = page.getByTestId('time-chart');
    await ready(chart);
    const chips = chart.getByTestId('legend-chip');
    // The raid median, the top 4, the bottom 4 and you (fix round 2 ruling).
    if (!isPhone(page)) {
      await expect(chips).toHaveCount(10);
      await expect(chips.first()).toContainText('Raid median');
      await expect(playersField(chart)).toHaveText('Top 4, bottom 4 of 40');
      await expect(chart.getByTestId('legend-more')).toHaveText('+31 more');
    }
    const you = chart.locator(`[data-testid="legend-chip"][data-id="${guidOf(13)}"]`);
    await expect(you).toContainText('Lightningbolltz');
    await expect(you).toContainText('YOU');
    // The head's raid figure is the sum of all 40 players' rounded DPS.
    await expect(chart.getByTestId('chart-head-meta')).toHaveText(
      `Raid ${raid40DpsSum().toLocaleString('en-US')} DPS over the whole fight`,
    );
  });

  test('the submerge draws flat at zero for every line', async ({ page }) => {
    test.skip(isPhone(page), 'reads the desktop tooltip');
    await page.goto(`${RAID40_URL}&tab=damage-done`);
    const chart = page.getByTestId('time-chart');
    await ready(chart);
    await hoverAt(page, chart, 225_000, 340_000);
    const rows = page.getByTestId('chart-tooltip-row');
    await expect(rows.first()).toBeVisible();
    for (const value of await rows.locator('span.numeric').allTextContents()) expect(value).toBe('0');
    await expect(page.getByTestId('chart-tooltip')).toContainText('Submerge phase · DPS, 5 s average');
  });

  test('the Players control draws a player at once, and Top 8 restores the default', async ({ page }) => {
    test.skip(isPhone(page), 'a coarse pointer uses the legend sheet');
    await page.goto(`${RAID40_URL}&tab=damage-done`);
    const chart = page.getByTestId('time-chart');
    await ready(chart);
    await chart.getByTestId('players-trigger').click();
    const list = listOf(chart, 'players');
    await list.locator('input[role="combobox"]').fill('osric');
    await list.locator(`[role="option"][data-value="${guidOf(39)}"]`).click();
    await expect(chart.locator(`[data-testid="legend-chip"][data-id="${guidOf(39)}"]`)).toBeVisible();
    await expect(page).toHaveURL(/players=/);
    await expect(playersField(chart)).toHaveText('10 of 40');
    await expect(list.locator('[data-foot-id="custom"]')).toHaveAttribute('data-lit', 'true');
    await list.locator('button[data-foot-id="split"]').click();
    await expect(page).not.toHaveURL(/players=/);
    await expect(chart.locator(`[data-testid="legend-chip"][data-id="${guidOf(39)}"]`)).toHaveCount(0);
  });

  test('the presets draw the top 8 or the bottom 8, and ride in the URL', async ({ page }) => {
    test.skip(isPhone(page), 'a coarse pointer uses the legend sheet');
    await page.goto(`${RAID40_URL}&tab=damage-done`);
    const chart = page.getByTestId('time-chart');
    await ready(chart);
    await chart.getByTestId('players-trigger').click();
    const list = listOf(chart, 'players');
    await list.locator('button[data-foot-id="bottom"]').click();
    await expect(page).toHaveURL(/players=bottom/);
    await expect(playersField(chart)).toHaveText('Bottom 8 of 40');
    // The bottom 8 are damage dealers: no healer's small damage is named a low performer.
    await expect(chart.locator(`[data-testid="legend-chip"][data-id="${guidOf(35)}"]`)).toHaveCount(0);
    await expect(chart.locator(`[data-testid="legend-chip"][data-id="${guidOf(29)}"]`)).toBeVisible();
    await list.locator('button[data-foot-id="top"]').click();
    await expect(page).toHaveURL(/players=top/);
    await expect(playersField(chart)).toHaveText('Top 8 of 40');
  });

  test('the 40 players’ figures add up to the head’s raid figure, in the same unit', async ({ page }) => {
    test.skip(isPhone(page), 'reads the desktop Players list');
    await page.goto(`${RAID40_URL}&tab=damage-done`);
    const chart = page.getByTestId('time-chart');
    await ready(chart);
    await chart.getByTestId('players-trigger').click();
    const figures = await listOf(chart, 'players').locator('.sel-row-meta').allTextContents();
    expect(figures).toHaveLength(40);
    const sum = figures.reduce((total, text) => total + Number(text.replace(/,/g, '')), 0);
    await expect(chart.getByTestId('chart-head-meta')).toHaveText(
      `Raid ${sum.toLocaleString('en-US')} DPS over the whole fight`,
    );
    // Every chip names its unit.
    for (const value of await chart.getByTestId('legend-value').allTextContents())
      expect(value).toMatch(/DPS$/);
  });

  test('the legend isolates a line with Only, and the log scale rides in the URL', async ({ page }) => {
    test.skip(isPhone(page), 'hover and Only are a fine pointer’s');
    await page.goto(`${RAID40_URL}&tab=damage-done`);
    const chart = page.getByTestId('time-chart');
    await ready(chart);
    // The median leads the legend, then you (pinned); the third chip is the top player.
    const first = chart.locator('[data-testid="legend-chip"]').nth(2);
    await first.hover();
    await chart.getByTestId('legend-only').click();
    // The line alone, plus the pinned (you), and the raid median stays as the reference.
    await expect(chart.locator('[data-testid="legend-chip"][aria-pressed="true"]')).toHaveCount(3);
    await chart.getByTestId('log-scale').click();
    await expect(page).toHaveURL(/scale=log/);
    await expect(chart.locator('[data-testid="chart-frame"]')).toHaveAttribute('data-scale', 'log');
  });

  test('the cast lane marks Power Infusion and Rebirth by spell id', async ({ page }) => {
    await page.goto(`${RAID40_URL}&tab=damage-done`);
    if (isPhone(page)) await page.getByTestId('chart-toggle').click();
    const chart = page.getByTestId('time-chart');
    await ready(chart);
    const casts = chart.locator('[data-testid="lane-marker"][data-kind="cast"]');
    await expect(casts).toHaveCount(2);
    await expect(casts.first()).toHaveAttribute('aria-label', /^Power Infusion · /);
    await expect(casts.nth(1)).toHaveAttribute('aria-label', /^Rebirth · /);
  });
});

test.describe('the 40-player threat chart (spec §B12, capture 9)', () => {
  test.beforeEach(async ({ page }) => serveRaid40(page));

  test('always draws the tanks, marks the taunt and the target change, and lists events within ±2 s', async ({
    page,
  }) => {
    await page.goto(`${RAID40_URL}&tab=threat`);
    const chart = page.getByTestId('threat-chart');
    await ready(chart);
    // Both tanks are drawn though their threat is a third of the top: pinned.
    for (const tank of [guidOf(8), guidOf(BROKK)]) {
      if (!isPhone(page))
        await expect(chart.locator(`[data-testid="legend-chip"][data-id="${tank}"]`)).toBeVisible();
    }
    await expect(chart.locator('[data-testid="lane-marker"][data-kind="taunt"]')).toHaveCount(1);
    await expect(chart.locator('[data-testid="lane-marker"][data-kind="target-change"]')).toHaveCount(2);
    if (isPhone(page)) return;
    await hoverAt(page, chart, 127_000, 340_000);
    const events = page.getByTestId('chart-tooltip-events');
    await expect(events).toContainText('Brokk died · Melee from Ragnaros · 3,862 (1,104 overkill)');
    await expect(events).toContainText('Ragnaros turned to Thandril');
    await expect(events).toContainText('Taunt · Thandril on Ragnaros');
    // One second's events read death, target change, taunt.
    await expect(events.locator('li').nth(1)).toContainText('turned to');
    await expect(page.getByTestId('chart-tooltip')).toContainText('estimated threat, every enemy');
    // The share column is of every player's estimated threat at that second.
    await expect(page.getByTestId('chart-tooltip')).toContainText('Share');
    // The table's leader is the legend's first chip.
    const leader = (await page
      .locator('li[data-testid^="threat-Player"]')
      .first()
      .locator('span')
      .first()
      .textContent())!.trim();
    await expect(chart.getByTestId('legend-chip').filter({ hasText: leader })).toBeVisible();
  });

  test('on a 360 × 640 phone the whole chart fits the first screen after a tab change', async ({ page }) => {
    test.skip(!isPhone(page), 'the phone fold stack');
    await page.setViewportSize({ width: 360, height: 640 });
    await page.goto(`${RAID40_URL}&tab=damage-done`);
    await page.locator('#report-tab-select').selectOption('threat', { force: true });
    const chart = page.getByTestId('threat-chart');
    await ready(chart);
    await page.waitForTimeout(800);
    const box = (await chart.boundingBox())!;
    expect(box.y + box.height).toBeLessThanOrEqual(640);
    // The legend is one row ending in +n, and +n opens the legend sheet.
    await chart.getByTestId('legend-more').click();
    const sheet = page.getByTestId('legend-sheet');
    await expect(sheet).toBeVisible();
    await expect(sheet).toContainText('Players');
    await sheet.getByTestId('legend-sheet-done').click();
    await expect(sheet).toBeHidden();
  });
});

test('on a phone every head field shows its value whole, one field to a row', async ({ page }) => {
  test.skip(!isPhone(page), 'the phone head');
  await page.setViewportSize({ width: 390, height: 844 });
  await serveRaid40(page);
  await page.goto(`${RAID40_URL}&tab=damage-done`);
  await page.getByTestId('chart-toggle').click();
  const chart = page.getByTestId('time-chart');
  await ready(chart);
  const windowField = chart.getByTestId('window-select');
  await expect(windowField).toHaveValue('0');
  const field = (await fieldBox(chart, 'window-select').boundingBox())!;
  expect(field.width).toBeGreaterThanOrEqual(240);
  const players = (await fieldBox(chart, 'players').boundingBox())!;
  expect(players.width).toBeGreaterThanOrEqual(240);
  expect(Math.abs(players.x - field.x)).toBeLessThanOrEqual(1);
  expect(await page.evaluate(() => document.documentElement.scrollWidth <= window.innerWidth)).toBe(true);
});

test('the Resources tab draws no damage chart and heads its rows', async ({ page }) => {
  await page.goto(`${FIXTURE}&tab=resources`);
  await expect(page.getByTestId('resource-graphs')).toBeVisible();
  await expect(page.getByTestId('resource-head')).toContainText('Player');
  await expect(page.getByTestId('resource-head')).toContainText('Time empty');
  await expect(page.getByTestId('time-chart')).toHaveCount(0);
  await expect(page.getByTestId('chart-toggle')).toHaveCount(0);
});

test('the main chart’s cursor draws the crosshair on the Raid cooldowns lanes too', async ({ page }) => {
  test.skip(isPhone(page), 'a fine pointer’s hover');
  await page.goto(`${FIXTURE}&tab=buffs`);
  const chart = page.getByTestId('time-chart');
  await ready(chart);
  const lanes = page.getByTestId('raid-cooldowns');
  test.skip((await lanes.count()) === 0, 'the fixture fight has no raid cooldown');
  await hoverAt(page, chart, 20_000, 60_000);
  await expect(lanes.getByTestId('synced-crosshair').first()).toBeVisible();
});
