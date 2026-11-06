// web/tests/e2e/tiers.spec.ts
// The tier list (design/specs/2026-11-06-tiers.md over the rules of 2026-10-09-tier-list.md): the
// three role lists against the published BiS files, the faction choice, ties named on both rows of
// a pair, the signed-in answer and pointer, keyboard on the control bar and a row, and the phone.
import type { Page } from '@playwright/test';
import { expect, test } from './support/test';
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
const ordinalOf = (n: number): string =>
  `${n}${n % 100 >= 11 && n % 100 <= 13 ? 'th' : (({ 1: 'st', 2: 'nd', 3: 'rd' } as Record<number, string>)[n % 10] ?? 'th')}`;

const envelope = (data: unknown) => ({
  status: 200,
  contentType: 'application/json',
  body: JSON.stringify({ ok: true, data, error: null, request_id: 'r' }),
});

function activePanel(page: Page, faction: TierFaction = 'alliance') {
  return page.getByTestId(`tier-faction-panel-${faction}`);
}

/** The phone bar folds its controls behind "Change"; open it, then wait for the script. */
async function openBar(page: Page): Promise<void> {
  const bar = page.locator('[data-control-bar]');
  await expect(bar).toHaveAttribute('data-booted', 'true');
  const toggle = bar.locator('[data-control-bar-toggle]');
  if (await toggle.isVisible()) await toggle.click();
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
  await expect(leader.getByTestId('rank-gap')).toHaveText('Least');
  await expect(leader.getByTestId('tier-number')).toHaveCSS('color', 'rgb(127, 212, 138)');
  const bestHealth = Math.max(...expected.map((row) => row.band.metrics!.effective_health));
  const bestRow = expected.findIndex((row) => row.band.metrics!.effective_health === bestHealth);
  await expect(
    activePanel(page).locator('[data-testid^="tier-row-"]').nth(bestRow).getByTestId('tier-effective-health'),
  ).toHaveCSS('color', 'rgb(127, 212, 138)');
});

test('tank figures are whole numbers and every column head is one short line', async ({ page }) => {
  await page.goto('/tiers/tank');
  const leader = expectedRows('tank', 'alliance')[0]!;
  const row = activePanel(page).locator('[data-testid^="tier-row-"]').first();
  await expect(row.getByTestId('tier-effective-health')).toHaveText(
    wholeNumber(leader.band.metrics!.effective_health),
  );
  await expect(row.getByTestId('tier-threat')).toHaveText(wholeNumber(leader.band.metrics!.tps));
  for (const width of [1280, 1440]) {
    await page.setViewportSize({ width, height: 900 });
    const heads = activePanel(page).locator('.ui-rank-heads > span');
    await expect(heads.nth(3)).toHaveText('Survival vs the best');
    for (const head of await heads.all()) {
      const { height, lineHeight, scroll, client } = await head.evaluate((node) => ({
        height: node.getBoundingClientRect().height,
        lineHeight: parseFloat(getComputedStyle(node).lineHeight),
        scroll: node.scrollWidth,
        client: node.clientWidth,
      }));
      expect(height).toBeLessThanOrEqual(lineHeight + 1);
      expect(scroll).toBeLessThanOrEqual(client);
    }
  }
});

for (const faction of ['alliance', 'horde'] as const) {
  test(`${faction}: both rows of a tied pair carry the mark, naming the partner`, async ({ page }) => {
    await page.goto(`/tiers?faction=${faction}`);
    const expected = expectedRows('dps', faction);
    const panel = activePanel(page, faction);
    await expect(panel.getByTestId('rank-tie')).toHaveCount(expectedTieCount(expected, 'dps'));
    const rows = panel.locator('[data-testid^="tier-row-"]');
    const within = (a: number, b: number) => Math.abs(1 - a / b) * 100 <= 1;
    for (const [index, row] of expected.entries()) {
      const partners = expected.filter(
        (other, j) =>
          j !== index && within(Math.min(row.metric, other.metric), Math.max(row.metric, other.metric)),
      );
      const tie = rows.nth(index).getByTestId('rank-tie');
      if (partners.length === 0) {
        await expect(tie).toHaveCount(0);
        continue;
      }
      for (const partner of partners) await expect(tie).toContainText(partner.entry.name);
    }
  });
}

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
  const chips = activePanel(page).getByTestId('rank-group-head');
  await expect(chips).toHaveCount(present.length);
  for (const [index, letter] of present.entries()) {
    const chip = chips.nth(index);
    await expect(chip.getByTestId('rank-group-letter')).toContainText(letter);
    await expect(chip).toContainText(`${letters.filter((l) => l === letter).length} spec`);
  }
  await expect(chips.first()).toContainText('Within 5% of the top');
  const groups = activePanel(page).getByTestId('tier-group');
  for (const [index, letter] of present.entries()) {
    await expect(groups.nth(index).locator('[data-testid^="tier-row-"]')).toHaveCount(
      letters.filter((l) => l === letter).length,
    );
  }
  await expect(activePanel(page).getByTestId('tier-howto')).toContainText(
    'S is within 5% of the top, A within 10%, B within 20%, C within 30%, D beyond.',
  );
});

for (const role of ['tank', 'healer'] as const) {
  test(`the ${role} list carries no tier letters`, async ({ page }) => {
    await page.goto(ROLE_PATH[role]);
    await expect(activePanel(page).getByTestId('tier-list')).toBeVisible();
    await expect(activePanel(page).getByTestId('rank-group-head')).toHaveCount(0);
    await expect(activePanel(page).getByTestId('rank-group-letter')).toHaveCount(0);
  });
}

test('the DPS list keeps the less-certain note', async ({ page }) => {
  await page.goto('/tiers');
  const expected = expectedRows('dps', 'alliance');
  const uncertain = expected.filter((row) => row.band.weights_low_confidence === true).length;
  await expect(activePanel(page).getByTestId('rank-note')).toHaveCount(uncertain);
});

test('on a phone the tier chip is a header row above its rows', async ({ page }) => {
  await page.setViewportSize(PHONE);
  await page.goto('/tiers');
  const group = activePanel(page).getByTestId('tier-group').first();
  const chip = await group.getByTestId('rank-group-head').boundingBox();
  const row = await group.locator('[data-testid^="tier-row-"]').first().boundingBox();
  expect(chip!.y + chip!.height).toBeLessThanOrEqual(row!.y + 1);
  expect(chip!.height).toBeLessThan(60);
});

for (const viewport of [{ width: 1440, height: 900 }, PHONE]) {
  test(`at ${viewport.width} the Your spec edge is its slice of the tier rail`, async ({ page }) => {
    await page.setViewportSize(viewport);
    await signIn(page, { class: 'Warrior', spec: 'Fury' });
    await page.goto('/tiers');
    const you = activePanel(page).locator('[data-rank-key][aria-current="true"]');
    await expect(you).toHaveCount(1);
    const chip = activePanel(page)
      .getByTestId('tier-group')
      .filter({ has: page.locator('[data-rank-key][aria-current="true"]') })
      .getByTestId('rank-group-head');
    const slice = await you.evaluate((row) => {
      const edge = getComputedStyle(row, '::before');
      return { x: row.getBoundingClientRect().left + parseFloat(edge.left), width: edge.width };
    });
    const rail = await chip.evaluate((el) => ({
      x: el.getBoundingClientRect().left,
      width: getComputedStyle(el).borderLeftWidth,
    }));
    expect(slice).toEqual(rail);
  });
}

test('the faction pill shows the Horde numbers and races, and the roles keep the faction', async ({
  page,
}) => {
  await page.goto('/tiers');
  await openBar(page);
  await page.getByTestId('control-faction-horde').click();
  await expect(page).toHaveURL(/\?faction=horde$/);
  await expect(activePanel(page, 'horde')).toBeVisible();
  await expect(activePanel(page)).toBeHidden();
  const expected = expectedRows('dps', 'horde');
  const first = activePanel(page, 'horde').locator('[data-testid^="tier-row-"]').first();
  await expect(first.getByTestId('tier-number')).toHaveText(oneDecimal(expected[0]!.metric));
  await expect(first).toContainText(raceName(expected[0]!.band.race!));
  await expect(page.getByTestId('control-role-tank')).toHaveAttribute('href', '/tiers/tank?faction=horde');
  await expect(page.getByTestId('control-faction-horde')).toHaveAttribute('aria-checked', 'true');
  await openBar(page);
  await page.getByTestId('control-faction-alliance').click();
  await expect(page).toHaveURL(/\/tiers$/);
  await expect(activePanel(page)).toBeVisible();
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
  await expect(page.getByTestId('control-faction-horde')).toHaveAttribute('aria-checked', 'true');
});

test('a Horde current character opens on Horde before paint; the URL still wins', async ({ page }) => {
  await page.addInitScript(() => {
    window.localStorage.setItem(
      'fs.currentCharacter',
      JSON.stringify({
        source: 'armory',
        ref: 'us/normal/tester',
        label: 'Tester · Fury Warrior',
        classSlug: 'warrior',
        faction: 'horde',
        savedAt: '2026-11-06T00:00:00.000Z',
      }),
    );
  });
  await page.goto('/tiers');
  expect(await page.evaluate(() => document.documentElement.getAttribute('data-tier-faction'))).toBe('horde');
  await expect(activePanel(page, 'horde')).toBeVisible();
  await expect(page.getByTestId('control-faction-horde')).toHaveAttribute('aria-checked', 'true');
  await page.goto('/tiers?faction=alliance');
  await expect(activePanel(page)).toBeVisible();
  await expect(activePanel(page, 'horde')).toBeHidden();
});

test('signed out: no answer and no marked row, the list is there and the caption offers sign-in', async ({
  page,
}) => {
  await page.goto('/tiers');
  await expect(page.getByTestId('tier-answer')).toBeHidden();
  await expect(activePanel(page).locator('[aria-current="true"]')).toHaveCount(0);
  await expect(activePanel(page).locator('[data-testid="tier-you-tag"]:visible')).toHaveCount(0);
  await expect(activePanel(page).getByTestId('tier-list')).toBeVisible();
  const signIn = activePanel(page).getByTestId('tier-caption').getByRole('link', { name: 'Sign in' });
  await expect(signIn).toHaveAttribute('href', '/login?next=%2Ftiers');
  await expect(activePanel(page).getByTestId('tier-caption')).toContainText('to see your spec marked.');
});

test('signed in on a DPS spec: the callout ranks it and its row is marked', async ({ page }) => {
  await signIn(page, { class: 'Warrior', spec: 'Fury' });
  await page.goto('/tiers');
  const expected = expectedRows('dps', 'alliance');
  const rank = expected.findIndex((row) => row.entry.spec === 'warrior-fury') + 1;
  const answer = page.getByTestId('tier-answer-body');
  await expect(answer).toContainText('Your spec · DPS');
  await expect(answer).toContainText('Fury Warrior');
  await expect(page.getByTestId('tier-answer-figure')).toHaveText(`${ordinalOf(rank)}of ${expected.length}`);
  await expect(page.getByTestId('tier-answer-detail')).toContainText(
    rank === 1 ? 'Level with the top spec.' : '% behind',
  );
  await expect(page.getByTestId('tier-answer-detail')).toContainText(rank === 1 ? '' : 'Tier ');
  await expect(page.getByTestId('tier-answer-bis')).toHaveAttribute('href', '/bis/warrior/fury');
  await expect(page.getByTestId('tier-answer-sim')).toHaveAttribute('href', '/sim');
  await expect(page.getByTestId('tier-row-warrior-fury').first()).toHaveAttribute('aria-current', 'true');
  await expect(page.getByTestId('tier-row-warrior-fury').first().getByTestId('tier-you-tag')).toBeVisible();
});

test('signed in on a tank spec: the DPS page points at the Tank list', async ({ page }) => {
  await signIn(page, { class: 'Paladin', spec: 'Protection' });
  await page.goto('/tiers');
  await expect(page.getByTestId('tier-answer-detail')).toHaveText('Protection Paladin is on the Tank list.');
  await expect(page.getByTestId('tier-answer-other')).toHaveAttribute('href', '/tiers/tank');
  await expect(page.getByTestId('tier-answer-figure')).toBeHidden();
  await expect(activePanel(page).locator('[aria-current="true"]')).toHaveCount(0);
  await page.goto('/tiers/tank');
  await expect(page.getByTestId('tier-answer-body')).toContainText('Protection Paladin');
  await expect(page.getByTestId('tier-answer-figure')).toContainText('of 3');
});

test('keyboard: arrows choose the faction and the role, and a row link shows a focus ring', async ({
  page,
}) => {
  await page.goto('/tiers');
  await openBar(page);
  await page.getByTestId('control-faction-alliance').focus();
  await page.keyboard.press('ArrowRight');
  await expect(page.getByTestId('control-faction-horde')).toHaveAttribute('aria-checked', 'true');
  await expect(page).toHaveURL(/\?faction=horde$/);
  await expect(activePanel(page, 'horde')).toBeVisible();
  await openBar(page);
  await page.getByTestId('control-role-dps').focus();
  await page.keyboard.press('End');
  await expect(page).toHaveURL(/\/tiers\/healer\?faction=horde$/);
  await page.goto('/tiers');
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

test('a DPS row links the ranking sim behind its number; a tank row, which /sim cannot run, has none', async ({
  page,
}) => {
  await page.goto('/tiers');
  const dpsRow = activePanel(page).locator('[data-testid^="tier-row-"]').first();
  const sim = dpsRow.getByTestId('tier-sim-link');
  await expect(sim).toHaveAccessibleName(/^View ranking sim: /);
  const href = await sim.getAttribute('href');
  expect(href).toMatch(/^\/sim\?code=/);
  expect(href).toContain('run=1');
  expect((await sim.boundingBox())!.height).toBeGreaterThanOrEqual(44);
  await page.goto('/tiers/tank');
  const tankRow = activePanel(page).locator('[data-testid^="tier-row-"]').first();
  await expect(tankRow).toBeVisible();
  await expect(tankRow.getByTestId('tier-sim-link')).toHaveCount(0);
});

test('Tier list is a Tools panel row and the Tools door is current, not Logs', async ({ page }) => {
  await page.goto('/tiers');
  const nav = page.getByTestId('primary-nav');
  await expect(nav.locator('a.site-door[data-door="tools"]')).toHaveAttribute('aria-current', 'page');
  await expect(nav.locator('a.site-door[data-door="logs"]')).not.toHaveAttribute('aria-current', 'page');
  const rows = await page.locator('#site-panel-tools a.site-panel-row').allTextContents();
  const titles = rows.map((row) => row.trim().toLowerCase());
  expect(titles.findIndex((row) => row.startsWith('tier list'))).toBe(3);
  const logs = await page.locator('#site-panel-logs a.site-panel-row').allTextContents();
  expect(logs.some((row) => row.trim().toLowerCase().startsWith('tier list'))).toBe(false);
  await expect(page.getByTestId('site-panel-row-tier-list')).toHaveAttribute('href', '/tiers');
});

test('the BiS index links to the tier list', async ({ page }) => {
  await page.goto('/bis');
  await expect(page.getByTestId('bis-index-tier-link')).toHaveAttribute('href', '/tiers');
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
        const links = rows.nth(index).locator('.ui-rr-links a');
        // A DPS row adds the ranking sim; /sim cannot rerun a tank's or a healer's figure.
        const linkCount = role === 'dps' ? 4 : 3;
        await expect(links).toHaveCount(linkCount);
        for (let l = 0; l < linkCount; l += 1) {
          expect((await links.nth(l).boundingBox())!.height).toBeGreaterThanOrEqual(44);
        }
      }
    });
  }

  test('tank rows keep effective health and threat on their own line', async ({ page }) => {
    await page.goto('/tiers/tank');
    const first = activePanel(page).locator('[data-testid^="tier-row-"]').first();
    await expect(first).toContainText('Effective health');
    await expect(first).toContainText('Threat');
    await expect(first).toContainText('threat / s');
  });
});
