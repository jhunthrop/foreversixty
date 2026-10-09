// web/tests/e2e/tiers.spec.ts
// The tier list (design/specs/2026-10-09-tier-list.md): the three role lists against the
// published BiS files, the faction toggle, ties on both rows of a pair, the signed-in callout
// and pointer, keyboard focus on a row, and the phone layout.
import { expect, test, type Page } from '@playwright/test';
import {
  expectedRows,
  expectedTieCount,
  metricText,
  oneDecimal,
  wholeNumber,
  raceName,
  type TierFaction,
  type TierRole,
} from './support/tier-list';
import { bisBand } from './support/bis-file';

const ROLE_PATH: Record<TierRole, string> = { dps: '/tiers', tank: '/tiers/tank', healer: '/tiers/healer' };
const PHONE = { width: 390, height: 844 };

const envelope = (data: unknown) => ({
  status: 200,
  contentType: 'application/json',
  body: JSON.stringify({ ok: true, data, error: null, request_id: 'r' }),
});

function activePanel(page: Page, faction: TierFaction = 'alliance') {
  return page.getByTestId(`tier-faction-panel-${faction}`);
}

async function signIn(page: Page, character: { class: string; spec: string; faction?: string }) {
  await page.context().addCookies([{ name: 'fs_csrf', value: 'token', domain: 'localhost', path: '/' }]);
  await page.route('**/v1/me', (route) =>
    route.fulfill(
      envelope({
        user: { id: 1, battletag: 'Fixture#1', email: null, role: 'user', anonymize: false },
        characters: [
          {
            key: 'us/normal/tester',
            region: 'us',
            ruleset: 'normal',
            name: 'Tester',
            level: 60,
            ...character,
          },
        ],
        guilds: [],
      }),
    ),
  );
}

for (const role of ['dps', 'tank', 'healer'] as const) {
  test(`${role}: rows follow the published files in order, with their figures`, async ({ page }) => {
    await page.goto(ROLE_PATH[role]);
    const expected = expectedRows(role, 'alliance');
    const rows = activePanel(page).locator('[data-testid^="tier-row-"]');
    await expect(rows).toHaveCount(expected.length);
    for (const [index, row] of expected.entries()) {
      const rendered = rows.nth(index);
      await expect(rendered).toHaveAttribute(
        'data-testid',
        `tier-row-${row.entry.class_slug}-${row.entry.spec_slug}`,
      );
      await expect(rendered.getByTestId('tier-number')).toHaveText(metricText(role, row.metric));
      await expect(rendered).toContainText(raceName(row.band.race!));
    }
  });
}

test('tank list sorts on damage taken, not effective health, and the best of each column is green', async ({
  page,
}) => {
  await page.goto('/tiers/tank');
  const expected = expectedRows('tank', 'alliance');
  const dtps = expected.map((row) => row.band.metrics!.dtps);
  expect(dtps).toEqual([...dtps].sort((a, b) => a - b));
  const leader = activePanel(page).locator('[data-testid^="tier-row-"]').first();
  await expect(leader.getByTestId('tier-gap')).toHaveText('Least');
  await expect(leader.getByTestId('tier-number')).toHaveCSS('color', 'rgb(127, 212, 138)');
  const bestHealth = Math.max(...expected.map((row) => row.band.metrics!.effective_health));
  const bestRow = expected.findIndex((row) => row.band.metrics!.effective_health === bestHealth);
  await expect(
    activePanel(page).locator('[data-testid^="tier-row-"]').nth(bestRow).getByTestId('tier-effective-health'),
  ).toHaveCSS('color', 'rgb(127, 212, 138)');
});

test('tank figures are whole numbers and the damage-taken head wraps to at most two lines', async ({
  page,
}) => {
  await page.goto('/tiers/tank');
  const leader = expectedRows('tank', 'alliance')[0]!;
  const row = activePanel(page).locator('[data-testid^="tier-row-"]').first();
  await expect(row.getByTestId('tier-effective-health')).toHaveText(
    wholeNumber(leader.band.metrics!.effective_health),
  );
  await expect(row.getByTestId('tier-threat')).toHaveText(wholeNumber(leader.band.metrics!.tps));
  const head = activePanel(page).locator('.tier-head-num').first();
  const { height, fontSize } = await head.evaluate((node) => ({
    height: node.getBoundingClientRect().height,
    fontSize: parseFloat(getComputedStyle(node).fontSize),
  }));
  // Two lines of a label at its normal line height (at most about 1.6 em each).
  expect(height).toBeLessThanOrEqual(fontSize * 1.6 * 2);
});

test('both rows of a tied pair carry the mark', async ({ page }) => {
  await page.goto('/tiers');
  const expected = expectedRows('dps', 'alliance');
  await expect(activePanel(page).getByTestId('tier-tie')).toHaveCount(expectedTieCount(expected, 'dps'));
  const rows = activePanel(page).locator('[data-testid^="tier-row-"]');
  for (const [index, row] of expected.entries()) {
    const next = expected[index + 1];
    if (next === undefined || Math.abs(1 - next.metric / row.metric) * 100 > 1) continue;
    await expect(rows.nth(index).getByTestId('tier-tie')).toBeVisible();
    await expect(rows.nth(index + 1).getByTestId('tier-tie')).toBeVisible();
  }
});

const TIER_LINES = [
  { letter: 'S', from: 0 },
  { letter: 'A', from: 5 },
  { letter: 'B', from: 10 },
  { letter: 'C', from: 20 },
  { letter: 'D', from: 30 },
] as const;

function expectedLetters(metrics: number[]): string[] {
  const top = metrics[0]!;
  return metrics.map((metric) => {
    const gap = Math.round((1 - metric / top) * 100 * 1e6) / 1e6;
    return [...TIER_LINES].reverse().find((line) => gap >= line.from)!.letter;
  });
}

test('the DPS list carries a tier chip per tier with its letter, band and count', async ({ page }) => {
  await page.goto('/tiers');
  const expected = expectedRows('dps', 'alliance');
  const letters = expectedLetters(expected.map((row) => row.metric));
  const present = TIER_LINES.map((line) => line.letter).filter((letter) => letters.includes(letter));
  const chips = activePanel(page).getByTestId('tier-chip');
  await expect(chips).toHaveCount(present.length);
  for (const [index, letter] of present.entries()) {
    const chip = chips.nth(index);
    await expect(chip.getByTestId('tier-letter')).toContainText(letter);
    await expect(chip).toContainText(`${letters.filter((l) => l === letter).length} spec`);
  }
  await expect(chips.first()).toContainText('Within 5% of the top');
  const groups = activePanel(page).getByTestId('tier-group');
  for (const [index, letter] of present.entries()) {
    await expect(groups.nth(index).locator('[data-testid^="tier-row-"]')).toHaveCount(
      letters.filter((l) => l === letter).length,
    );
  }
  await expect(activePanel(page).getByTestId('tier-ruler')).toHaveCount(0);
  await expect(activePanel(page).getByTestId('tier-notes')).toContainText(
    'S is within 5% of the top, A within 10%, B within 20%, C within 30%, D beyond; a ≈ tie across a line is a tie.',
  );
});

for (const role of ['tank', 'healer'] as const) {
  test(`the ${role} list carries no tier letters`, async ({ page }) => {
    await page.goto(ROLE_PATH[role]);
    await expect(activePanel(page).getByTestId('tier-list')).toBeVisible();
    await expect(activePanel(page).getByTestId('tier-chip')).toHaveCount(0);
    await expect(activePanel(page).getByTestId('tier-letter')).toHaveCount(0);
  });
}

test('the DPS list keeps the less-certain note', async ({ page }) => {
  await page.goto('/tiers');
  const expected = expectedRows('dps', 'alliance');
  const uncertain = expected.filter((row) => row.band.weights_low_confidence === true).length;
  await expect(activePanel(page).getByTestId('tier-low-confidence')).toHaveCount(uncertain);
});

test('on a phone the tier chip is a header row above its rows', async ({ page }) => {
  await page.setViewportSize(PHONE);
  await page.goto('/tiers');
  const group = activePanel(page).getByTestId('tier-group').first();
  const chip = await group.getByTestId('tier-chip').boundingBox();
  const row = await group.locator('[data-testid^="tier-row-"]').first().boundingBox();
  expect(chip!.y + chip!.height).toBeLessThanOrEqual(row!.y + 1);
  expect(chip!.height).toBeLessThan(60);
});

test('the faction pill shows the Horde numbers and races, and the roles keep the faction', async ({
  page,
}) => {
  await page.goto('/tiers');
  await page.getByTestId('tier-faction-horde').click();
  await expect(page).toHaveURL(/\?faction=horde$/);
  await expect(activePanel(page, 'horde')).toBeVisible();
  await expect(activePanel(page)).toBeHidden();
  const expected = expectedRows('dps', 'horde');
  const first = activePanel(page, 'horde').locator('[data-testid^="tier-row-"]').first();
  await expect(first.getByTestId('tier-number')).toHaveText(oneDecimal(expected[0]!.metric));
  await expect(first).toContainText(raceName(expected[0]!.band.race!));
  await expect(page.getByTestId('tier-tab-tank')).toHaveAttribute('href', '/tiers/tank?faction=horde');
  await expect(page.getByTestId('tier-faction-horde')).toHaveAttribute('aria-pressed', 'true');
});

test('?faction=horde loads the Horde list with no flash of the Alliance one', async ({ page }) => {
  await page.goto('/tiers/healer?faction=horde');
  const attribute = await page.evaluate(() => document.documentElement.getAttribute('data-tier-faction'));
  expect(attribute).toBe('horde');
  await expect(activePanel(page, 'horde')).toBeVisible();
  const expected = expectedRows('healer', 'horde');
  await expect(activePanel(page, 'horde').getByTestId('tier-number').first()).toHaveText(
    oneDecimal(expected[0]!.metric),
  );
  expect(bisBand(expected[0]!.entry.spec, 'horde', 60).race).toBeTruthy();
});

test('signed out: no callout and no pointer, the list is still there', async ({ page }) => {
  await page.goto('/tiers');
  await expect(page.getByTestId('tier-callout')).toHaveCount(0);
  await expect(page.getByTestId('tier-pointer')).toHaveCount(0);
  await expect(page.getByTestId('tier-you-slot')).toBeHidden();
  await expect(activePanel(page).getByTestId('tier-you-pill')).toHaveCount(0);
});

test('signed in on a DPS spec: the callout ranks it and its row is marked', async ({ page }) => {
  await signIn(page, { class: 'Warrior', spec: 'Fury' });
  await page.goto('/tiers');
  const expected = expectedRows('dps', 'alliance');
  const rank = expected.findIndex((row) => row.entry.spec === 'warrior-fury') + 1;
  const callout = page.getByTestId('tier-callout');
  await expect(callout).toContainText(`Fury Warrior is`);
  await expect(callout).toContainText(`of ${expected.length} DPS specs.`);
  await expect(callout).toContainText(rank === 1 ? 'Level with the top spec.' : '% behind');
  await expect(page.getByTestId('tier-callout-button')).toHaveAttribute('href', '/bis/warrior/fury');
  await expect(page.getByTestId('tier-row-warrior-fury').first()).toHaveClass(/is-you/);
  await expect(page.getByTestId('tier-row-warrior-fury').first().getByTestId('tier-you-pill')).toBeVisible();
});

test('signed in on a tank spec: the DPS page points at the Tank list', async ({ page }) => {
  await signIn(page, { class: 'Paladin', spec: 'Protection' });
  await page.goto('/tiers');
  const pointer = page.getByTestId('tier-pointer');
  await expect(pointer).toContainText('Protection Paladin is on the Tank list.');
  await expect(pointer).toHaveAttribute('href', '/tiers/tank');
  await expect(page.getByTestId('tier-callout')).toHaveCount(0);
  await page.goto('/tiers/tank');
  await expect(page.getByTestId('tier-callout')).toContainText('Protection Paladin is');
  await expect(page.getByTestId('tier-callout')).toContainText('of 3 Tank specs.');
});

test('keyboard: arrow keys move across the role tabs and a row link shows a focus ring', async ({ page }) => {
  await page.goto('/tiers');
  await page.getByTestId('tier-tab-dps').focus();
  await page.keyboard.press('ArrowRight');
  await expect(page.getByTestId('tier-tab-tank')).toBeFocused();
  await page.keyboard.press('End');
  await expect(page.getByTestId('tier-tab-healer')).toBeFocused();
  const link = activePanel(page).locator('[data-testid="tier-spec-link"]').first();
  await link.focus();
  await expect(link).toBeFocused();
  await expect(link).toHaveCSS('outline-style', 'solid');
  await expect(link).toHaveCSS('outline-color', 'rgb(229, 185, 85)');
});

test('every row offers BiS, Guide and Planner at 44px', async ({ page }) => {
  await page.goto('/tiers/healer');
  const row = activePanel(page).locator('[data-testid^="tier-row-"]').first();
  for (const name of ['BiS', 'Guide', 'Planner']) {
    const link = row.getByRole('link', { name: new RegExp(`^${name}:`) });
    await expect(link).toBeVisible();
    const box = await link.boundingBox();
    expect(box!.height).toBeGreaterThanOrEqual(44);
  }
});

test('Tier List sits between Rankings and Guides in the nav', async ({ page }) => {
  await page.goto('/tiers');
  const labels = await page.getByTestId('primary-nav').locator('a').allTextContents();
  const clean = labels.map((label) => label.trim().toLowerCase());
  expect(clean.indexOf('tier list')).toBe(clean.indexOf('rankings') + 1);
  expect(clean.indexOf('guides')).toBe(clean.indexOf('tier list') + 1);
  await expect(page.getByTestId('primary-nav').locator('a', { hasText: 'Tier List' })).toHaveAttribute(
    'aria-current',
    'page',
  );
});

test('the BiS index and the home page link to the tier list', async ({ page }) => {
  await page.goto('/bis');
  await expect(page.getByTestId('bis-index-tier-link')).toHaveAttribute('href', '/tiers');
  await page.goto('/');
  await expect(page.getByTestId('home-tier-list-link').first()).toHaveAttribute('href', '/tiers');
});

test('the BiS index shows damage taken per second beside effective health for a tank', async ({ page }) => {
  await page.goto('/bis');
  const tank = expectedRows('tank', 'alliance')[0]!;
  const summary = page.getByTestId(`bis-index-dps-${tank.entry.class_slug}-${tank.entry.spec_slug}`);
  await expect(summary).toContainText('effective health');
  await expect(summary).toContainText(`${wholeNumber(tank.band.metrics!.dtps)} damage taken per second`);
});

test.describe('phone', () => {
  test.use({ viewport: PHONE });

  for (const role of ['dps', 'tank', 'healer'] as const) {
    test(`${role}: no sideways scroll, a link strip on every row`, async ({ page }) => {
      await page.goto(ROLE_PATH[role]);
      const { scrollWidth, clientWidth } = await page.evaluate(() => ({
        scrollWidth: document.documentElement.scrollWidth,
        clientWidth: document.documentElement.clientWidth,
      }));
      expect(scrollWidth).toBeLessThanOrEqual(clientWidth);
      const rows = activePanel(page).locator('[data-testid^="tier-row-"]');
      const count = await rows.count();
      for (let index = 0; index < count; index += 1) {
        const links = rows.nth(index).locator('.tier-link');
        await expect(links).toHaveCount(3);
        for (let l = 0; l < 3; l += 1) {
          expect((await links.nth(l).boundingBox())!.height).toBeGreaterThanOrEqual(44);
        }
      }
    });
  }

  test('tank rows keep effective health and threat on their own line', async ({ page }) => {
    await page.goto('/tiers/tank');
    const first = activePanel(page).locator('[data-testid^="tier-row-"]').first();
    await expect(first).toContainText('Effective health');
    await expect(first).toContainText('Threat per second');
  });
});
