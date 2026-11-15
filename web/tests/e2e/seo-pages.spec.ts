// web/tests/e2e/seo-pages.spec.ts
// What a crawler gets (SEO lane 2026-11-08, design/marketing/2026-11-08-marketing-plan.md): every
// BiS page is a server-rendered URL with its own title, canonical and h1, its list in the HTML
// before any script (the spec page with its folded mid-level bands, level 60 its own page); the band controls are real links; the bare /guild
// page is the Find a guild landing, not a loading shell; personal pages carry noindex.
import { test, expect } from './support/test';
import { chooseControl, openBis } from './support/bis-page';

test.describe('without JavaScript, as a crawler fetches', () => {
  test.use({ javaScriptEnabled: false });

  test('the level 60 page carries its own head, its h1 and its list', async ({ page }) => {
    await page.goto('/bis/warrior/fury/60');
    await expect(page).toHaveTitle(/^Fury Warrior (Pre-Raid )?BiS, Level 60, WoW Forever/);
    await expect(page.locator('link[rel="canonical"]')).toHaveAttribute(
      'href',
      'https://foreversixty.gg/bis/warrior/fury/60',
    );
    await expect(page.getByRole('heading', { level: 1 })).toHaveText('Fury Warrior BiS');
    await expect(page.getByTestId('bis-band-alliance-60')).toBeVisible();
    await expect(page.getByTestId('bis-doll')).toContainText(/\S/);
    // The band controls are real links: the first band the spec's own path, a folded band its hash.
    await expect(page.getByTestId('control-band-20')).toHaveAttribute('href', '/bis/warrior/fury');
    await expect(page.getByTestId('control-band-40')).toHaveAttribute(
      'href',
      '/bis/warrior/fury#levels-40-49',
    );
    await expect(page.getByTestId('control-band-60')).toHaveAttribute('href', '/bis/warrior/fury/60');
    // The folded bands live on the spec page alone.
    await expect(page.getByTestId('bis-leveling')).toHaveCount(0);
    const crumbs = await page.locator('script[type="application/ld+json"]').first().textContent();
    expect(JSON.parse(crumbs ?? '{}')['@type']).toBe('BreadcrumbList');
  });

  test('the spec page stays canonical for its first band', async ({ page }) => {
    await page.goto('/bis/warrior/fury');
    await expect(page.locator('link[rel="canonical"]')).toHaveAttribute(
      'href',
      'https://foreversixty.gg/bis/warrior/fury',
    );
    await expect(page.getByRole('heading', { level: 1 })).toHaveText('Fury Warrior BiS');
    await expect(page.locator('.pdf-frame .pdf-meta')).toHaveText('20 to 29');
    // The folded mid-level bands are in its HTML, Alliance shown.
    const leveling = page.getByTestId('bis-leveling');
    await expect(leveling.getByRole('heading', { level: 2 })).toHaveText('Leveling 30 to 59');
    for (const band of [30, 40, 50]) {
      await expect(leveling.locator(`[data-band="${band}"][data-faction="alliance"]`).first()).toBeVisible();
    }
  });

  test('/guild is the Find a guild landing, not a loading shell', async ({ page }) => {
    await page.goto('/guild');
    await expect(page.getByRole('heading', { level: 1 })).toHaveText('Guilds');
    await expect(page.getByRole('textbox', { name: 'Guild name' })).toBeVisible();
    await expect(page.locator('meta[name="robots"]')).toHaveCount(0);
  });

  for (const path of ['/account', '/connect', '/login', '/reports', '/premium/checkout']) {
    test(`${path} is personal or a step in a flow: noindex`, async ({ page }) => {
      await page.goto(path);
      await expect(page.locator('meta[name="robots"]')).toHaveAttribute('content', 'noindex');
    });
  }
});

test('choosing a band moves the address to its page or hash without a reload', async ({ page }) => {
  await openBis(page, '/bis/warrior/fury#levels-40-49');
  await expect(page.getByTestId('bis-band-alliance-40')).toBeVisible();
  await page.evaluate(() => {
    (window as unknown as { __marker?: boolean }).__marker = true;
  });
  await chooseControl(page, 'band', '60');
  await expect(page).toHaveURL(/\/bis\/warrior\/fury\/60$/);
  await expect(page.getByTestId('bis-band-alliance-60')).toBeVisible();
  await chooseControl(page, 'band', '50');
  await expect(page).toHaveURL(/\/bis\/warrior\/fury#levels-50-59$/);
  await expect(page.getByTestId('bis-band-alliance-50')).toBeVisible();
  await expect(page.getByRole('heading', { level: 1 })).toContainText('Fury Warrior');
  expect(await page.evaluate(() => (window as unknown as { __marker?: boolean }).__marker)).toBe(true);
});
