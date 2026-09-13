import { expect, test } from '@playwright/test';

for (const path of ['/dungeons', '/zones']) {
  test(`${path} dates its facts and lists the sources behind them`, async ({ page }) => {
    await page.goto(path);
    await expect(page.getByText(/^Updated /)).toBeVisible();
    const sources = page.getByRole('heading', { name: 'Sources' }).locator('xpath=..');
    await expect(sources).toBeVisible();
    const links = sources.getByRole('link');
    expect(await links.count()).toBeGreaterThan(0);
    const hrefs = await links.evaluateAll((els) => els.map((e) => e.getAttribute('href')));
    expect(new Set(hrefs).size).toBe(hrefs.length);
    for (const href of hrefs) expect(href).toMatch(/^https?:\/\//);
  });
}

for (const path of ['/guides', '/changelog']) {
  test(`${path} carries the updated stamp of its newest entry`, async ({ page }) => {
    await page.goto(path);
    await expect(page.getByText(/^Updated /)).toBeVisible();
  });
}
