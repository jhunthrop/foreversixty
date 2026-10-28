// web/tests/e2e/premium.spec.ts
// The launch plans page (design/specs/2026-10-27-launch-pricing.md section 9 cases 1-7 and 18,
// with the R-P8 ruling: founder prices until 6 January 2027, no "before" state). Prices come
// from lib/billing/plans, never typed here.
import { createRequire } from 'node:module';
import { expect, test, type Page } from '@playwright/test';
import { priceFor } from '../../src/lib/billing/plan-costs';
import { FOUNDER_WINDOW_CLOSES, PLAN_ORDER, formatPrice } from '../../src/lib/billing/plans';
import { PREMIUM_SOURCE_KEY } from '../../src/lib/account/premium-hint';
import { plansPageCopy } from '../../src/lib/billing/plans-page-copy';
import {
  captionFields,
  exampleHref,
  isPlaceholder,
  serverRunExample,
} from '../../src/lib/billing/premium-examples';
import { serverRunCaption, serverRunCaptionTemplates } from '../../src/lib/billing/premium-example-copy';
import { EXAMPLE_IMAGE_HEIGHT, EXAMPLE_IMAGE_WIDTH } from '../../src/lib/billing/example-image';
import {
  COST_BOX_PX,
  cardHeights,
  founderLineBoxes,
  guildBlockPx,
  guildPriceBoxes,
  type GuildPriceBoxes,
} from './support/premium-layout';
import { API } from './support/local-stack';
import { premiumSourceCopy } from '../../src/lib/billing/premium-source-copy';
import {
  PREMIUM_GRANT_OPEN_ENDED,
  PREMIUM_OWN_FOUNDER,
  PREMIUM_SEAT_WITH_PORTAL,
} from '../../src/test-support/subscribe-fixtures';
import { plantSessionHint, routeMeCell } from './support/subscribe-stubs';

const LAUNCH_DAY = '2026-10-27T18:00:00Z';
const DECEMBER = '2026-12-08T18:00:00Z';
const JUST_BEFORE_CLOSE = '2027-01-07T07:59:59Z';
const AT_CLOSE = FOUNDER_WINDOW_CLOSES;
const PHONE = { width: 360, height: 780 };
const MIN_TAP_PX = 44;
const FIRST_SCREEN = { width: 1280, height: 800 };
const BUTTONS_END_ABOVE_PX = 520;
const BAND_LINK_ABOVE_PX = 300;
const HEADING_TOP_TOLERANCE_PX = 40;
const SUBPIXEL_PX = 1;
const FOUNDER_LINE_WIDTHS = [360, 1024] as const;
const GUILD_WIDTHS = [360, 768, 1024, 1440] as const;
/** The three stored Premium hints: the page must hold its layout with and without one. */
const PREMIUM_STATES = [
  { name: 'free visitor', stored: null },
  { name: 'premium (data-premium set)', stored: 'own' },
] as const;

const premiumPrice = (interval: 'yearly' | 'monthly', founder: boolean): string =>
  formatPrice(priceFor('premium', interval, founder), interval);

interface OpenOptions {
  path?: string;
  /** The stored Premium source the head script reads into `data-premium`. */
  stored?: string | null;
}

async function openAt(page: Page, time: string, options: OpenOptions = {}): Promise<void> {
  const { path = '/premium', stored = null } = options;
  if (stored !== null) {
    await page.addInitScript(
      ([key, value]) => localStorage.setItem(key as string, value as string),
      [PREMIUM_SOURCE_KEY, stored],
    );
  }
  await page.clock.install({ time });
  await page.goto(path);
}

const premiumButton = (page: Page) => page.getByRole('link', { name: /^Subscribe · / });

// The page itself ships no island: the one astro-island is the layout's CharacterSelector.
for (const { name, stored } of PREMIUM_STATES) {
  test(`ships only the layout character selector island (${name})`, async ({ page }) => {
    await openAt(page, DECEMBER, { stored });
    const islands = page.locator('astro-island');
    await expect(islands).toHaveCount(1);
    await expect(islands.first()).toHaveAttribute('component-url', /CharacterSelector/);
  });
}

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
  await page.getByLabel(plansPageCopy.interval.labels.monthly).check();
  await expect(page.getByTestId('price-raid_guild').locator('.list-line:visible')).toHaveText(
    '$25 a month after 6 Jan 2027',
  );
});

test('in December Yearly is checked and the Premium button links to the yearly checkout', async ({
  page,
}) => {
  await openAt(page, DECEMBER);
  await expect(page.getByTestId('founder-row')).toBeVisible();
  await expect(page.getByLabel(plansPageCopy.interval.labels.yearly)).toBeChecked();
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

  await page.getByLabel(plansPageCopy.interval.labels.monthly).check();
  await expect(premiumButton(page)).toHaveText(`Subscribe · ${premiumPrice('monthly', true)}`);
  await expect(premiumButton(page)).toHaveAttribute('href', /plan=premium&interval=monthly/);
  for (const plan of PLAN_ORDER) {
    await expect(page.getByTestId(`price-${plan}`)).toContainText(`$${priceFor(plan, 'monthly', true)}`);
  }
  expect(await heights()).toEqual(before);

  await page.getByLabel(plansPageCopy.interval.labels.yearly).focus();
  await page.keyboard.press('ArrowRight');
  await expect(page.getByLabel(plansPageCopy.interval.labels.monthly)).toBeChecked();
  await page.keyboard.press('ArrowLeft');
  await expect(page.getByLabel(plansPageCopy.interval.labels.yearly)).toBeChecked();
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
  const order = await page.locator('[data-plan]').evaluateAll((nodes) =>
    nodes
      .filter((node) => node.getBoundingClientRect().height > 0)
      .sort((a, b) => a.getBoundingClientRect().top - b.getBoundingClientRect().top)
      .map((node) => node.getAttribute('data-plan')),
  );
  expect(order).toEqual(['free', 'premium', 'example', 'guild', 'raid_guild', 'history']);
});

// Spec B3.2 (premium-value amendment): from 1024 Premium sits in region A and the two guild
// plans in region B below it, so the buttons share a height and the guild pair shares a top.
test('at 1024 the subscribe buttons share a height and the guild pair shares a top', async ({ page }) => {
  await page.setViewportSize({ width: 1024, height: 900 });
  await openAt(page, DECEMBER);
  const boxes = await page.locator('[data-plan] .action-slot a:visible').evaluateAll((links) =>
    links.map((link) => {
      const { top, height } = link.getBoundingClientRect();
      return { plan: link.closest('[data-plan]')?.getAttribute('data-plan'), top, height };
    }),
  );
  expect(boxes.map((box) => box.plan)).toEqual([...PLAN_ORDER]);
  for (const box of boxes) expect(box.height).toBe(boxes[0]!.height);
  const guild = boxes.find((box) => box.plan === 'guild');
  const raid = boxes.find((box) => box.plan === 'raid_guild');
  expect(raid?.top).toBe(guild?.top);
});

test('at 1024 the guild row puts the history promise in column 3', async ({ page }) => {
  await page.setViewportSize({ width: 1024, height: 900 });
  await openAt(page, DECEMBER);
  const [guild, raid, history] = await Promise.all(
    ['plan-guild', 'plan-raid_guild', 'history-promise'].map((id) =>
      page.getByTestId(id).evaluate((node) => {
        const { left, top } = node.getBoundingClientRect();
        return { left, top };
      }),
    ),
  );
  expect(history.left).toBeGreaterThan(raid.left);
  expect(raid.left).toBeGreaterThan(guild.left);
  expect(history.top).toBeCloseTo(guild.top, 0);
});

async function expectNoCardMovement(page: Page, width: number): Promise<void> {
  const check = (boxes: GuildPriceBoxes[]): void => {
    for (const box of boxes) {
      expect(box.block, `${box.plan} block at ${width}`).toBe(guildBlockPx(width));
      expect(box.costBox, `${box.plan} cost box at ${width}`).toBe(COST_BOX_PX);
      expect(box.costOverflows, `${box.plan} cost line overflows at ${width}`).toBe(false);
    }
  };
  const yearly = { cards: await cardHeights(page), boxes: await guildPriceBoxes(page) };
  check(yearly.boxes);
  await page.getByLabel(plansPageCopy.interval.labels.monthly).check();
  const monthly = { cards: await cardHeights(page), boxes: await guildPriceBoxes(page) };
  check(monthly.boxes);
  expect(monthly.cards, `card heights at ${width}`).toEqual(yearly.cards);
}

for (const width of GUILD_WIDTHS) {
  for (const [label, time] of [
    ['founder open', DECEMBER],
    ['founder closed', AT_CLOSE],
  ] as const) {
    test(`guild price blocks hold their height across interval at ${width}, ${label}`, async ({ page }) => {
      await page.setViewportSize({ width, height: 900 });
      await openAt(page, time);
      await expectNoCardMovement(page, width);
    });
  }
}

for (const { name, stored } of PREMIUM_STATES) {
  test.describe(`first screen at 1280 x 800 (${name})`, () => {
    test.use({ viewport: FIRST_SCREEN });

    test('buttons end above y 520 and the band link above y 300', async ({ page }) => {
      await openAt(page, DECEMBER, { stored });
      for (const plan of ['plan-free', 'plan-premium']) {
        const bottom = await page
          .getByTestId(plan)
          .locator('.action-slot')
          .evaluate((node) => node.getBoundingClientRect().bottom);
        expect(bottom, plan).toBeLessThan(BUTTONS_END_ABOVE_PX);
      }
      const link = page.getByTestId('band-guild-link');
      const linkBottom = await link.evaluate((node) => node.getBoundingClientRect().bottom);
      expect(linkBottom).toBeLessThan(BAND_LINK_ABOVE_PX);
    });

    test('the band link brings the guild heading to the top with both guild buttons in view', async ({
      page,
    }) => {
      await openAt(page, DECEMBER, { stored });
      await page.getByTestId('band-guild-link').click();
      const heading = page.getByRole('heading', { level: 2, name: plansPageCopy.guildRegion.heading });
      await expect
        .poll(() => heading.evaluate((node) => node.getBoundingClientRect().top))
        .toBeLessThan(HEADING_TOP_TOLERANCE_PX);
      expect(await heading.evaluate((node) => node.getBoundingClientRect().top)).toBeGreaterThanOrEqual(0);
      for (const plan of ['plan-guild', 'plan-raid_guild']) {
        const card = page.getByTestId(plan);
        for (const part of ['.action-slot', '.slot']) {
          const bottom = await card.locator(part).evaluate((node) => node.getBoundingClientRect().bottom);
          expect(bottom, `${plan} ${part}`).toBeLessThanOrEqual(FIRST_SCREEN.height);
        }
      }
    });
  });
}

test('the guild section holds both guild plans, the history promise and the rules in order', async ({
  page,
}) => {
  await openAt(page, DECEMBER);
  const section = page.locator('#guild-leaders');
  await expect(section.getByTestId('plan-guild')).toHaveCount(1);
  await expect(section.getByTestId('plan-raid_guild')).toHaveCount(1);
  await expect(section.getByTestId('history-promise')).toHaveText(plansPageCopy.historyPromise);
  await expect(section.getByRole('heading', { name: plansPageCopy.howItWorks })).toBeVisible();
  const terms = await section.locator('[data-rule-term]').allTextContents();
  const [first] = plansPageCopy.rules;
  expect(terms[0]?.trim()).toBe(first?.term);
  // The copy order is Verified members, Size, Paying together.
  const sizeTerm = plansPageCopy.rules[1]?.term ?? '';
  const paying = plansPageCopy.rules[2]?.term ?? '';
  const trimmed = terms.map((term) => term.trim());
  expect(trimmed.indexOf(paying)).toBe(trimmed.indexOf(sizeTerm) + 1);
  await expect(section.locator('.row').first().locator('dd')).toHaveText(first?.sentence ?? '');
  expect(await section.textContent()).not.toContain('Battle.net');
});

test('the Free card renders once at every width', async ({ page }) => {
  for (const width of GUILD_WIDTHS) {
    await page.setViewportSize({ width, height: 900 });
    await openAt(page, DECEMBER);
    await expect(page.locator('#free')).toHaveCount(1);
    await expect(page.getByTestId('plan-free')).toHaveCount(1);
    await expect(page.getByRole('heading', { name: plansPageCopy.free.name, exact: true })).toHaveCount(1);
  }
});

test('forbidden words stay off the page, and the coach line is the first Premium benefit', async ({
  page,
}) => {
  await openAt(page, DECEMBER);
  const text = (await page.locator('main').textContent()) ?? '';
  expect(text).not.toMatch(/chip in|retention|compare more than/i);
  expect(text).not.toContain('!');
  expect(text).toContain(plansPageCopy.premium.features[0]);
  expect(text).not.toContain(plansPageCopy.trialEmail);
  const trialText = await page.locator('.slot, .row:has([data-rule-term]) dd').allTextContents();
  expect(trialText.join(' ')).not.toMatch(/email you/i);
});

test('no loot-council example is on the page', async ({ page }) => {
  await openAt(page, DECEMBER);
  const text = (await page.locator('main').textContent()) ?? '';
  expect(text).not.toMatch(/EXAMPLE · A REAL LOOT COUNCIL/i);
  await expect(page.getByTestId('premium-example')).toHaveCount(1);
});

test('the example card reads the data file: label, caption template, link, image box', async ({ page }) => {
  await openAt(page, DECEMBER);
  const template = serverRunExample.preset
    ? serverRunCaptionTemplates.preset
    : serverRunCaptionTemplates.ownGear;
  const caption = template(captionFields(serverRunExample));
  expect(serverRunCaption(serverRunExample)).toBe(caption);
  const card = page.getByTestId('premium-example');
  await expect(card.getByTestId('premium-example-caption')).toHaveText(caption);
  await expect(card.locator('[data-part="label"]')).toHaveText(plansPageCopy.examples.serverRun.label);
  const link = card.getByTestId('premium-example-link');
  await expect(link).toHaveText(plansPageCopy.examples.serverRun.link);
  await expect(link).toHaveAttribute('href', exampleHref(serverRunExample));
  const image = card.locator('img');
  await expect(image).toHaveAttribute('width', String(EXAMPLE_IMAGE_WIDTH));
  await expect(image).toHaveAttribute('height', String(EXAMPLE_IMAGE_HEIGHT));
  await expect(image).toHaveAttribute('alt', caption);
});

test('the example image box is reserved at 4:3 whether or not the picture loads', async ({ page }) => {
  await page.route('**/premium/example-server-run.*', (route) => route.abort());
  await openAt(page, DECEMBER);
  const box = page.getByTestId('premium-example').locator('[data-part="image"]');
  const { width, height } = await box.evaluate((node) => node.getBoundingClientRect());
  expect(Math.abs(height - (width * 3) / 4)).toBeLessThanOrEqual(SUBPIXEL_PX);
});

const SHIP = process.env.SHIP === '1';

test('the example ships real data, not the placeholder', () => {
  if (!SHIP && isPlaceholder(serverRunExample)) {
    test
      .info()
      .annotations.push({ type: 'placeholder', description: 'premium-examples.json is a placeholder' });
    test.skip(true, 'placeholder data; set SHIP=1 to make this a failure');
  }
  expect(
    isPlaceholder(serverRunExample),
    'premium-examples.json is still the placeholder — run the PV19 capture before shipping',
  ).toBe(false);
});

test('Open this result answers 200 signed out', async ({ request }) => {
  test.skip(!SHIP && isPlaceholder(serverRunExample), 'placeholder data; set SHIP=1 to run');
  const response = await request.get(`${API}/v1/sims/${encodeURIComponent(serverRunExample.simId)}`);
  expect(response.status()).toBe(200);
});

test('the founder line sets on one line at 360 and 1024', async ({ page }) => {
  for (const width of FOUNDER_LINE_WIDTHS) {
    await page.setViewportSize({ width, height: 900 });
    await openAt(page, DECEMBER);
    const lines = await founderLineBoxes(page);
    expect(lines.length, `founder lines at ${width}`).toBeGreaterThan(0);
    for (const line of lines) expect(line.lines, `${line.text} at ${width}`).toBe(1);
  }
});

interface AxeResults {
  violations: { id: string }[];
}
const AXE_SOURCE = createRequire(import.meta.url).resolve('axe-core/axe.min.js');

for (const { name, stored } of PREMIUM_STATES) {
  test(`axe finds no violations and every control is at least 44px tall at 360 (${name})`, async ({
    page,
  }) => {
    await page.setViewportSize(PHONE);
    await openAt(page, DECEMBER, { stored });
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
    expect(heights.length).toBeGreaterThan(0);
    for (const height of heights) expect(height).toBeGreaterThanOrEqual(MIN_TAP_PX);
  });
}

test('the band link and Open this result are 44px tall from 768 up', async ({ page }) => {
  await page.setViewportSize({ width: 1024, height: 900 });
  await openAt(page, DECEMBER);
  for (const link of [page.getByTestId('band-guild-link'), page.getByTestId('premium-example-link')]) {
    expect((await link.boundingBox())?.height).toBeGreaterThanOrEqual(MIN_TAP_PX);
  }
});

// --- the Premium card for an account that has Premium (subscribe-and-join spec §4.4, case 16) ---

const SOURCE_STATES = [
  { source: 'own', me: PREMIUM_OWN_FOUNDER },
  { source: 'seat', me: PREMIUM_SEAT_WITH_PORTAL },
  { source: 'grant', me: PREMIUM_GRANT_OPEN_ENDED },
] as const;

const premiumCard = (page: Page) => page.getByTestId('plan-premium');

async function signedOutPremiumHeight(page: Page): Promise<number> {
  await routeMeCell(page, () => null);
  await openAt(page, DECEMBER);
  return (await cardHeights(page)).premium ?? 0;
}

for (const { source, me } of SOURCE_STATES) {
  test(`stored source ${source}: the Premium card reads its line and See your plan on first paint, at the signed-out height`, async ({
    page,
  }) => {
    const signedOutHeight = await signedOutPremiumHeight(page);
    expect(signedOutHeight).toBeGreaterThan(0);

    await page.evaluate(() => localStorage.clear());
    await plantSessionHint(page);
    await routeMeCell(page, () => me);
    await page.addInitScript(
      ([key, value]) => localStorage.setItem(key as string, value as string),
      [PREMIUM_SOURCE_KEY, source],
    );
    await page.reload();

    const card = premiumCard(page);
    await expect(card.locator(`.owned-line[data-source="${source}"]`)).toBeVisible();
    await expect(card.locator(`.owned-line[data-source="${source}"]`)).toHaveText(
      premiumSourceCopy.premiumCard[source],
    );
    await expect(card.getByRole('link', { name: premiumSourceCopy.buttons.seePlan })).toHaveAttribute(
      'href',
      '/account#plan',
    );
    await expect(card.getByRole('link', { name: /^Subscribe · / })).toBeHidden();
    await expect(card.locator('[data-default]')).toBeHidden();
    expect((await cardHeights(page)).premium).toBe(signedOutHeight);
    expect(await page.evaluate((key) => localStorage.getItem(key), PREMIUM_SOURCE_KEY)).toBe(source);
  });
}

test('the swap holds at 360 with the See your plan link 44px tall', async ({ page }) => {
  await page.setViewportSize(PHONE);
  await plantSessionHint(page);
  await routeMeCell(page, () => PREMIUM_GRANT_OPEN_ENDED);
  await openAt(page, DECEMBER, { stored: 'grant' });
  const link = premiumCard(page).getByRole('link', { name: premiumSourceCopy.buttons.seePlan });
  await expect(link).toBeVisible();
  expect((await link.boundingBox())?.height ?? 0).toBeGreaterThanOrEqual(MIN_TAP_PX);
});

test('a malformed stored source renders the default card and the key is gone', async ({ page }) => {
  await plantSessionHint(page);
  await routeMeCell(page, () => null);
  await openAt(page, DECEMBER, { stored: 'premium' });
  const card = premiumCard(page);
  await expect(card.locator('[data-default]')).toBeVisible();
  await expect(card.locator('.owned-line:visible')).toHaveCount(0);
  await expect(premiumButton(page)).toBeVisible();
  await expect(page.locator('html')).not.toHaveAttribute('data-premium', /.*/);
  await expect.poll(() => page.evaluate((key) => localStorage.getItem(key), PREMIUM_SOURCE_KEY)).toBeNull();
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
