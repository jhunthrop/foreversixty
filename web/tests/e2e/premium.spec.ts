// web/tests/e2e/premium.spec.ts
// The launch plans page (design/specs/2026-10-27-launch-pricing.md section 9 cases 1-7 and 18,
// with the R-P8 ruling: founder prices until 6 January 2027, no "before" state). Prices come
// from lib/billing/plans, never typed here.
import { createRequire } from 'node:module';
import { expect, test, type Page } from '@playwright/test';
import { priceFor } from '../../src/lib/billing/plan-costs';
import { FOUNDER_WINDOW_CLOSES, PLAN_ORDER, formatPrice } from '../../src/lib/billing/plans';
import { plansPageCopy } from '../../src/lib/billing/plans-page-copy';

const LAUNCH_DAY = '2026-10-27T18:00:00Z';
const DECEMBER = '2026-12-08T18:00:00Z';
const JUST_BEFORE_CLOSE = '2027-01-07T07:59:59Z';
const AT_CLOSE = FOUNDER_WINDOW_CLOSES;
const PHONE = { width: 360, height: 780 };
const MIN_TAP_PX = 44;

const premiumPrice = (interval: 'yearly' | 'monthly', founder: boolean): string =>
  formatPrice(priceFor('premium', interval, founder), interval);

async function openAt(page: Page, time: string, path = '/premium'): Promise<void> {
  await page.clock.install({ time });
  await page.goto(path);
}

const premiumButton = (page: Page) => page.getByRole('link', { name: /^Subscribe · / });

// The page itself ships no island: the one astro-island is the layout's CharacterSelector.
test('ships only the layout character selector island', async ({ page }) => {
  await page.goto('/premium');
  const islands = page.locator('astro-island');
  await expect(islands).toHaveCount(1);
  await expect(islands.first()).toHaveAttribute('component-url', /CharacterSelector/);
});

test('on launch day the founder row shows, prices are founder prices and the list price is plain text', async ({
  page,
}) => {
  await openAt(page, LAUNCH_DAY);
  await expect(page.getByTestId('founder-row')).toContainText(plansPageCopy.founder.open);
  await expect(premiumButton(page)).toHaveText(`Subscribe · ${premiumPrice('yearly', true)}`);
  await expect(page.getByTestId('price-premium').locator('.list-line:visible')).toHaveText(
    '$50 after 6 Jan 2027',
  );
  await expect(page.locator('s:visible')).toHaveCount(0);
  await page.getByLabel('Monthly').check();
  await expect(page.getByTestId('price-raid_guild').locator('.list-line:visible')).toHaveText(
    '$25 a month after 6 Jan 2027',
  );
});

test('in December Yearly is checked and the Premium button links to the yearly checkout', async ({
  page,
}) => {
  await openAt(page, DECEMBER);
  await expect(page.getByTestId('founder-row')).toBeVisible();
  await expect(page.getByLabel('Yearly')).toBeChecked();
  const button = premiumButton(page);
  await expect(button).toHaveText(`Subscribe · ${premiumPrice('yearly', true)}`);
  await expect(button).toHaveAttribute('href', /plan=premium&interval=yearly/);
});

test('Monthly changes every price and button, no card changes height, arrow keys switch', async ({
  page,
}) => {
  await openAt(page, DECEMBER);
  const heights = async (): Promise<number[]> =>
    page
      .locator('[data-plan]')
      .evaluateAll((cards) => cards.map((card) => card.getBoundingClientRect().height));
  const before = await heights();

  await page.getByLabel('Monthly').check();
  await expect(premiumButton(page)).toHaveText(`Subscribe · ${premiumPrice('monthly', true)}`);
  await expect(premiumButton(page)).toHaveAttribute('href', /plan=premium&interval=monthly/);
  for (const plan of PLAN_ORDER) {
    await expect(page.getByTestId(`price-${plan}`)).toContainText(`$${priceFor(plan, 'monthly', true)}`);
  }
  expect(await heights()).toEqual(before);

  await page.getByLabel('Yearly').focus();
  await page.keyboard.press('ArrowRight');
  await expect(page.getByLabel('Monthly')).toBeChecked();
  await page.keyboard.press('ArrowLeft');
  await expect(page.getByLabel('Yearly')).toBeChecked();
});

test('the founder window ends at the instant, to the second', async ({ page }) => {
  await openAt(page, JUST_BEFORE_CLOSE);
  await expect(premiumButton(page)).toHaveText(`Subscribe · ${premiumPrice('yearly', true)}`);
  await expect(page.getByTestId('founder-row')).toBeVisible();

  await openAt(page, AT_CLOSE);
  await expect(page.getByTestId('founder-row')).toBeHidden();
  await expect(premiumButton(page)).toHaveText(`Subscribe · ${premiumPrice('yearly', false)}`);
  await expect(page.getByTestId('price-premium').locator('.list-line:visible')).toHaveCount(0);
});

test('at 360 the Free card comes first, once, with the free list and the promise, before Premium', async ({
  page,
}) => {
  await page.setViewportSize(PHONE);
  await openAt(page, DECEMBER);
  await expect(page.getByTestId('free-summary')).toHaveCount(0);
  await expect(page.locator('#free')).toContainText(plansPageCopy.free.promise);
  const order = await page.locator('[data-plan], [data-testid="history-promise"]').evaluateAll((nodes) =>
    nodes
      .filter((node) => node.getBoundingClientRect().height > 0)
      .sort((a, b) => a.getBoundingClientRect().top - b.getBoundingClientRect().top)
      .map((node) => node.getAttribute('data-plan') ?? node.getAttribute('data-testid')),
  );
  expect(order).toEqual(['free', 'premium', 'guild', 'raid_guild', 'history-promise']);
});

test('at 1024 the three subscribe buttons share top and height', async ({ page }) => {
  await page.setViewportSize({ width: 1024, height: 900 });
  await openAt(page, DECEMBER);
  const boxes = await page
    .locator('[data-plan] a:visible')
    .evaluateAll((links) =>
      links.map((link) => link.getBoundingClientRect()).map(({ top, height }) => ({ top, height })),
    );
  expect(boxes).toHaveLength(PLAN_ORDER.length);
  for (const box of boxes) expect(box).toEqual(boxes[0]);
});

test('no text says retention, compare more than, or shouts', async ({ page }) => {
  await openAt(page, DECEMBER);
  const text = await page.locator('main').innerText();
  expect(text).not.toMatch(/retention|compare more than/i);
  expect(text).not.toContain('!');
});

interface AxeResults {
  violations: { id: string }[];
}
const AXE_SOURCE = createRequire(import.meta.url).resolve('axe-core/axe.min.js');

test('axe finds no violations and every control is at least 44px tall at 360', async ({ page }) => {
  await page.setViewportSize(PHONE);
  await openAt(page, DECEMBER);
  await page.addScriptTag({ path: AXE_SOURCE });
  const results = await page.evaluate(async (): Promise<AxeResults> => {
    const axe = (window as unknown as { axe: { run: () => Promise<AxeResults> } }).axe;
    return axe.run();
  });
  expect(results.violations.map((violation) => violation.id)).toEqual([]);

  const heights = await page
    .locator('main a:visible, main input:visible, main summary:visible')
    .evaluateAll((nodes) =>
      nodes.filter((node) => !node.closest('p, li')).map((node) => node.getBoundingClientRect().height),
    );
  for (const height of heights) expect(height).toBeGreaterThanOrEqual(MIN_TAP_PX);
});

test('?canceled=1 acknowledges a canceled Stripe Checkout and loads no third-party script', async ({
  page,
}) => {
  const scripts: string[] = [];
  page.on('request', (r) => {
    if (r.resourceType() === 'script') scripts.push(r.url());
  });
  await page.goto('/premium?canceled=1');
  await expect(page.getByTestId('premium-canceled-notice')).toContainText(
    'Checkout was canceled. Nothing was charged.',
  );
  const origin = new URL(page.url()).origin;
  expect(scripts.filter((url) => !url.startsWith(`${origin}/_astro/`))).toEqual([]);

  await page.goto('/premium');
  await expect(page.getByTestId('premium-canceled-notice')).toBeHidden();
});
