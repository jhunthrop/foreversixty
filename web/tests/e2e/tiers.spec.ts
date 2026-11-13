// web/tests/e2e/tiers.spec.ts
// The tier list (directions mock g-tiers, round 5, over the rules of 2026-10-09-tier-list.md): the
// three role lists against the published BiS files, the faction choice, the S to D letter column,
// the spec panel signed out and in, keyboard on the control bar and a row, and the phone.
import { readFileSync } from 'node:fs';
import type { Page } from '@playwright/test';
import { expect, test } from './support/test';
import {
  expectedRows,
  metricText,
  oneDecimal,
  wholeNumber,
  type TierFaction,
  type TierRole,
} from './support/tier-list';
import { bisBand } from './support/bis-file';

/** The catalogue's full spec name (web/src/data/specs.json), which every tier row links by. */
const SPEC_NAMES = new Map(
  (
    JSON.parse(readFileSync(new URL('../../src/data/specs.json', import.meta.url), 'utf8')) as {
      class_slug: string;
      spec_slug: string;
      full_name: string;
    }[]
  ).map((spec) => [`${spec.class_slug}/${spec.spec_slug}`, spec.full_name]),
);
function fullName(classSlug: string, specSlug: string): string {
  const name = SPEC_NAMES.get(`${classSlug}/${specSlug}`);
  if (!name) throw new Error(`no catalogue spec ${classSlug}/${specSlug}`);
  return name;
}

const ROLE_PATH: Record<TierRole, string> = { dps: '/tiers', tank: '/tiers/tank', healer: '/tiers/healer' };
const PHONE = { width: 390, height: 844 };
const ACCENT = 'rgb(0, 147, 178)';

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
  test(`${role}: rows follow the published files in order, named in full, with their figures`, async ({
    page,
  }) => {
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
      // The spec's full name links to its best in slot. A DPS row's figure links to the
      // ranking sim behind it (1b3853d60); a tank or healer figure stays plain, since the
      // browser sim reproduces neither damage taken nor healing.
      const bisLink = rendered.locator(`a[href^="/bis/${row.entry.class_slug}/${row.entry.spec_slug}/60"]`);
      await expect(bisLink).toHaveCount(1);
      await expect(bisLink).toContainText(fullName(row.entry.class_slug, row.entry.spec_slug));
      await expect(rendered.getByRole('link')).toHaveCount(role === 'dps' ? 2 : 1);
      if (role === 'dps') {
        await expect(rendered.getByTestId('tier-number')).toHaveAttribute('href', /^\/sim\?/);
      }
    }
  });
}

test('the title band names the fight, and the bars are the accent from 90% of the lowest spec', async ({
  page,
}) => {
  await page.goto('/tiers');
  await expect(page.getByRole('heading', { level: 1 })).toHaveText('Tier list at 60');
  await expect(page.getByTestId('tier-page')).toContainText(
    'Best in slot, raid buffs · vs a level 63 boss · 3 min',
  );
  await expect(activePanel(page).getByTestId('rank-group-letter').first()).toHaveAttribute(
    'title',
    'Within 5% of #1',
  );
  const rows = activePanel(page).locator('[data-testid^="tier-row-"]');
  const fill = (index: number) =>
    rows
      .nth(index)
      .locator('.ui-meter > i')
      .evaluate((node) => ({
        width: parseFloat((node as HTMLElement).style.width),
        colour: getComputedStyle(node).backgroundColor,
      }));
  const top = await fill(0);
  const last = await fill((await rows.count()) - 1);
  expect(top).toEqual({ width: 100, colour: ACCENT });
  // The lowest spec sits a tenth above the floor: (x - 0.9x) / (top - 0.9x) of the bar, well under the share.
  expect(last.width).toBeGreaterThan(0);
  expect(last.width).toBeLessThan(50);
});

test('only the healer list says under its title what HPS does not count', async ({ page }) => {
  await page.goto('/tiers/healer');
  // The sentence itself is tier-copy.test.ts's; here it is the one lead on the page, on this list only.
  await expect(page.locator('.ui-page-header-lead')).toContainText('Shields count for what they soak');
  // It names the fight the list ranks on: the logged raid kill, or the stand-in until then.
  await expect(page.locator('.ui-page-header-lead')).toContainText(
    /Ranked on (one healer’s share|a stand-in fight)/,
  );
  await page.goto('/tiers');
  await expect(page.locator('.ui-page-header-lead')).toHaveCount(0);
});

test('tank list sorts on damage taken, shows maximum health, and the best of each column is green', async ({
  page,
}) => {
  await page.goto('/tiers/tank');
  const expected = expectedRows('tank', 'alliance');
  const dtps = expected.map((row) => row.band.metrics!.dtps);
  expect(dtps).toEqual([...dtps].sort((a, b) => a - b));
  const rows = activePanel(page).locator('[data-testid^="tier-row-"]');
  const leader = rows.first();
  // The least-hit tank reads "lowest"; every other row its extra damage taken, whole percent.
  await expect(leader.getByTestId('rank-gap')).toHaveText('lowest');
  await expect(leader.getByTestId('tier-number')).toHaveCSS('color', 'rgb(127, 212, 138)');
  await expect(rows.nth(1).getByTestId('rank-gap')).toHaveText(/^\+\d+% taken$/);
  const healths = expected.map((row) => row.band.metrics!.health);
  if (healths.some((health) => health === undefined)) {
    // A file ranked before health was published: no half-filled column, never effective health.
    await expect(activePanel(page).getByTestId('tier-health')).toHaveCount(0);
    await expect(activePanel(page).locator('.ui-rank-heads')).not.toContainText('Health');
    return;
  }
  const bestHealth = Math.max(...(healths as number[]));
  await expect(rows.nth(healths.indexOf(bestHealth)).getByTestId('tier-health')).toHaveCSS(
    'color',
    'rgb(127, 212, 138)',
  );
});

test("tank bars are the least damage taken over the row's own, from zero", async ({ page }) => {
  await page.goto('/tiers/tank');
  const expected = expectedRows('tank', 'alliance');
  const rows = activePanel(page).locator('[data-testid^="tier-row-"]');
  const width = (index: number): Promise<number> =>
    rows
      .nth(index)
      .locator('.ui-meter > i')
      .evaluate((node) => parseFloat((node as HTMLElement).style.width));
  const least = expected[0]!.band.metrics!.dtps;
  const last = expected.length - 1;
  expect(await width(0)).toBeCloseTo(100, 0);
  expect(await width(last)).toBeCloseTo((100 * least) / expected[last]!.band.metrics!.dtps, 0);
});

test('tank figures are whole numbers and every column head is one short line beside the rail', async ({
  page,
}) => {
  await page.goto('/tiers/tank');
  const leader = expectedRows('tank', 'alliance')[0]!;
  const row = activePanel(page).locator('[data-testid^="tier-row-"]').first();
  const health = leader.band.metrics!.health;
  if (health !== undefined) await expect(row.getByTestId('tier-health')).toHaveText(wholeNumber(health));
  await expect(row.getByTestId('tier-threat')).toHaveText(wholeNumber(leader.band.metrics!.tps));
  for (const width of [1280, 1440, 1760, 2000]) {
    await page.setViewportSize({ width, height: 900 });
    const heads = activePanel(page).locator('.ui-rank-heads > span');
    await expect(heads.nth(4)).toHaveText('Taken / s');
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
    // Every name holds one line, and the bars start at one x down the list.
    for (const link of await activePanel(page).getByTestId('tier-spec-link').all()) {
      const { height, lineHeight } = await link.evaluate((node) => ({
        height: node.getBoundingClientRect().height,
        lineHeight: parseFloat(getComputedStyle(node).lineHeight),
      }));
      expect(height).toBeLessThanOrEqual(lineHeight + 1);
    }
    const bars = await activePanel(page)
      .locator('.ui-rr-bar')
      .evaluateAll((nodes) => nodes.map((node) => Math.round(node.getBoundingClientRect().left)));
    expect(new Set(bars).size).toBe(1);
    // The list fits its column: nothing spills past the plate.
    const plate = activePanel(page).getByTestId('tier-list');
    const { scrollWidth, clientWidth } = await plate.evaluate((node) => ({
      scrollWidth: node.scrollWidth,
      clientWidth: node.clientWidth,
    }));
    expect(scrollWidth).toBeLessThanOrEqual(clientWidth);
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

test('the DPS list groups its rows under a letter column, S to D', async ({ page }) => {
  await page.goto('/tiers');
  const expected = expectedRows('dps', 'alliance');
  const letters = expectedLetters(expected.map((row) => row.metric));
  const present = TIER_LINES.map((line) => line.letter).filter((letter) => letters.includes(letter));
  const heads = activePanel(page).getByTestId('rank-group-head');
  await expect(heads).toHaveCount(present.length);
  const groups = activePanel(page).getByTestId('tier-group');
  for (const [index, letter] of present.entries()) {
    await expect(heads.nth(index).getByTestId('rank-group-letter')).toContainText(letter);
    await expect(groups.nth(index).getByRole('list')).toHaveAccessibleName(`Tier ${letter}`);
    await expect(groups.nth(index).locator('[data-testid^="tier-row-"]')).toHaveCount(
      letters.filter((l) => l === letter).length,
    );
  }
});

for (const viewport of [{ width: 1440, height: 900 }, PHONE]) {
  test(`at ${viewport.width} the letter column sits beside its rows`, async ({ page }) => {
    await page.setViewportSize(viewport);
    await page.goto('/tiers');
    const group = activePanel(page).getByTestId('tier-group').first();
    const head = await group.getByTestId('rank-group-head').boundingBox();
    const row = await group.locator('[data-testid^="tier-row-"]').first().boundingBox();
    expect(head!.x + head!.width).toBeLessThanOrEqual(row!.x + 1);
    expect(Math.abs(head!.y - row!.y)).toBeLessThanOrEqual(1);
  });
}

for (const role of ['tank', 'healer'] as const) {
  test(`the ${role} list carries no tier letters`, async ({ page }) => {
    await page.goto(ROLE_PATH[role]);
    await expect(activePanel(page).getByTestId('tier-list')).toBeVisible();
    await expect(activePanel(page).getByTestId('rank-group-head')).toHaveCount(0);
  });
}

test('a shared healer rank names the healer it ties with; no other row carries a tie', async ({ page }) => {
  await page.goto('/tiers/healer');
  const rows = activePanel(page).locator('[data-testid^="tier-row-"]');
  const count = await rows.count();
  for (let index = 0; index < count; index += 1) {
    const rank = (await rows.nth(index).getByTestId('rank-rank').textContent()) ?? '';
    const note = rows.nth(index).getByTestId('rank-note');
    if (rank.startsWith('=')) await expect(note).toHaveText(/^≈ tie with /);
    else await expect(note.filter({ hasText: '≈ tie' })).toHaveCount(0);
  }
  await page.goto('/tiers');
  await expect(activePanel(page).getByTestId('rank-note').filter({ hasText: '≈ tie' })).toHaveCount(0);
});

test('the DPS list keeps the less-certain note', async ({ page }) => {
  await page.goto('/tiers');
  const expected = expectedRows('dps', 'alliance');
  const uncertain = expected.filter((row) => row.band.weights_low_confidence === true).length;
  await expect(activePanel(page).getByTestId('rank-note')).toHaveCount(uncertain);
});

test('the faction tab shows the Horde numbers, and the roles keep the faction', async ({ page }) => {
  await page.goto('/tiers');
  await openBar(page);
  await page.getByTestId('control-faction-horde').click();
  await expect(page).toHaveURL(/\?faction=horde$/);
  await expect(activePanel(page, 'horde')).toBeVisible();
  await expect(activePanel(page)).toBeHidden();
  const expected = expectedRows('dps', 'horde');
  const first = activePanel(page, 'horde').locator('[data-testid^="tier-row-"]').first();
  await expect(first.getByTestId('tier-number')).toHaveText(oneDecimal(expected[0]!.metric));
  await expect(page.getByTestId('control-role-tank')).toHaveAttribute('href', '/tiers/tank?faction=horde');
  await expect(page.getByTestId('control-role-tank')).toHaveText('Tanks');
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

test('signed out: the spec panel offers sign-in and no row is marked', async ({ page }) => {
  await page.goto('/tiers');
  const card = page.getByTestId('tier-card');
  await expect(page.getByTestId('tier-card-signed-out')).toBeVisible();
  await expect(card.getByRole('link', { name: 'Sign in' })).toHaveAttribute('href', '/login?next=%2Ftiers');
  await expect(page.getByTestId('tier-answer-body')).toBeHidden();
  await expect(activePanel(page).locator('[aria-current="true"]')).toHaveCount(0);
  await expect(activePanel(page).getByTestId('tier-list')).toBeVisible();
});

test('signed in on a DPS spec: the panel names it and marks its row, with no line repeating it', async ({
  page,
}) => {
  await signIn(page, { class: 'Warrior', spec: 'Fury' });
  await page.goto('/tiers');
  await expect(page.getByTestId('tier-answer-name')).toHaveText('Fury Warrior');
  await expect(page.getByTestId('tier-answer-standing')).toBeHidden();
  await expect(page.getByTestId('tier-answer-bis')).toHaveAttribute('href', '/bis/warrior/fury/60');
  await expect(page.getByTestId('tier-answer-guide')).toHaveAttribute('href', '/guides/warrior/fury');
  await expect(page.getByTestId('tier-card-signed-out')).toBeHidden();
  const mine = page.getByTestId('tier-row-warrior-fury').first();
  await expect(mine).toHaveAttribute('aria-current', 'true');
  const edge = await mine.evaluate((row) => getComputedStyle(row, '::before').backgroundColor);
  expect(edge).toBe(ACCENT);
});

test('signed in on a tank spec: the DPS page points at the Tanks list', async ({ page }) => {
  await signIn(page, { class: 'Paladin', spec: 'Protection' });
  await page.goto('/tiers');
  await expect(page.getByTestId('tier-answer-other')).toHaveText('On the Tanks list ›');
  await expect(page.getByTestId('tier-answer-other')).toHaveAttribute('href', '/tiers/tank');
  await expect(page.getByTestId('tier-answer-guide')).toHaveAttribute('href', '/guides/paladin/protection');
  await expect(activePanel(page).locator('[aria-current="true"]')).toHaveCount(0);
  await page.goto('/tiers/tank');
  await expect(page.getByTestId('tier-answer-name')).toHaveText('Protection Paladin');
  await expect(page.getByTestId('tier-answer-standing')).toBeHidden();
});

test('the spec panel holds its size from loading to answered: nothing under it moves', async ({ page }) => {
  await page.setViewportSize(PHONE);
  let release: () => void = () => undefined;
  const held = new Promise<void>((resolve) => (release = resolve));
  await page.context().addCookies([{ name: 'fs_csrf', value: 'token', domain: 'localhost', path: '/' }]);
  await page.route('**/v1/me', async (route) => {
    await held;
    await route.fulfill(
      envelope({
        user: { id: 1, battletag: 'Fixture#1', email: null, role: 'user', anonymize: false },
        characters: [
          {
            key: 'us/normal/t',
            region: 'us',
            ruleset: 'normal',
            name: 'Tester',
            level: 60,
            class: 'Warrior',
            spec: 'Fury',
          },
        ],
        guilds: [],
      }),
    );
  });
  await page.goto('/tiers');
  const list = activePanel(page).getByTestId('tier-list');
  const before = (await list.boundingBox())!.y;
  release();
  await expect(page.getByTestId('tier-answer-name')).toHaveText('Fury Warrior');
  expect((await list.boundingBox())!.y).toBe(before);
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
  await expect(link).toHaveCSS('outline-color', 'rgb(255, 236, 113)');
});

test('Tier list is a door of its own and current, not Logs', async ({ page }) => {
  await page.goto('/tiers');
  const nav = page.getByTestId('primary-nav');
  const tiers = nav.locator('a.site-door[data-door="tier-list"]');
  await expect(tiers).toHaveAttribute('aria-current', 'page');
  await expect(tiers).toHaveAttribute('href', '/tiers');
  await expect(nav.locator('a.site-door[data-door="logs"]')).not.toHaveAttribute('aria-current', 'page');
  const logs = await page.locator('#site-panel-logs a.site-panel-row').allTextContents();
  expect(logs.some((row) => row.trim().toLowerCase().startsWith('tier list'))).toBe(false);
});

test('the BiS index links to the tier list', async ({ page }) => {
  await page.goto('/bis');
  await expect(page.getByTestId('bis-index-tier-link')).toHaveAttribute('href', '/tiers');
});

test.describe('phone', () => {
  test.use({ viewport: PHONE });

  for (const role of ['dps', 'tank', 'healer'] as const) {
    test(`${role}: no sideways scroll, the spec panel first, a 44 px name link on every row`, async ({
      page,
    }) => {
      await page.goto(ROLE_PATH[role]);
      const { scrollWidth, clientWidth } = await page.evaluate(() => ({
        scrollWidth: document.documentElement.scrollWidth,
        clientWidth: document.documentElement.clientWidth,
      }));
      expect(scrollWidth).toBeLessThanOrEqual(clientWidth);
      const card = (await page.getByTestId('tier-card').boundingBox())!;
      const list = (await activePanel(page).getByTestId('tier-list').boundingBox())!;
      expect(card.y + card.height).toBeLessThanOrEqual(list.y);
      const links = activePanel(page).getByTestId('tier-spec-link');
      const count = await links.count();
      for (let index = 0; index < count; index += 1) {
        expect((await links.nth(index).boundingBox())!.height).toBeGreaterThanOrEqual(44);
      }
    });
  }

  test('tank rows keep health, threat and vs #1 on their own line, labelled', async ({ page }) => {
    await page.goto('/tiers/tank');
    const rows = activePanel(page).locator('[data-testid^="tier-row-"]');
    await expect(rows.first()).not.toContainText('Effective health');
    await expect(rows.first()).toContainText('Threat');
    await expect(rows.nth(1).getByTestId('tier-gap-phone')).toBeVisible();
    await expect(rows.nth(1).getByTestId('tier-gap-phone')).toContainText(/\+\d+% taken/);
  });
});
