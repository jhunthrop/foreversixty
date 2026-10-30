// web/tests/e2e/one-character-journey.spec.ts
// W.11 acceptance: the chosen character follows the player through every page, against the real
// local API (docs/superpowers/reviews/lifecycle/r4/local-stack.md), signed in as hunthrop#1894,
// who holds Bow Jackzon (a hunter with gear and talents from the seed; its race, spec and gear are read from the
// account's own record, never assumed) and Pipsqueak.
//
// Not part of the CI suite: it needs the dev stack, so it runs only with W11_JOURNEY=1 through
// playwright.journey.config.ts, against a built site served by `astro preview` (the simulator's
// island and the character and guild shells exist only in a build). W11_CAPTURES=<dir> writes a
// 1280px capture of each stop.
import { expect, test, type Page } from '@playwright/test';
import { API, devLogin, pointerOf, serveShells, signInThroughProxy } from './support/local-stack';
import { nativeSelect } from './support/select';

const WEB = process.env.W11_WEB ?? 'http://localhost:4331';
const CAPTURES = process.env.W11_CAPTURES;
const BOW = 'us/pvp/bow-jackzon';
const PIPSQUEAK = 'us/pvp/pipsqueak';

/** What the account's record says Bow Jackzon is: the journey asserts the page agrees with it. */
const bow = { race: '', raceSlug: '', spec: '', headId: 0, points: 0, split: '' };

test.skip(process.env.W11_JOURNEY !== '1', 'needs the local stack (W11_JOURNEY=1)');
test.use({ viewport: { width: 1280, height: 900 } });

test.beforeEach(async ({ page, context }) => {
  const session = await devLogin('hunthrop#1894');
  const me = (await (await fetch(`${API}/v1/me`, { headers: { cookie: session.cookie } })).json()) as {
    data: {
      characters: {
        key: string;
        race?: string;
        spec?: string;
        build?: { gear?: Record<string, number>; talents?: { points: number[] } };
      }[];
    };
  };
  const record = me.data.characters.find((character) => character.key === BOW);
  bow.race = record?.race ?? '';
  bow.raceSlug = bow.race.toLowerCase().replace(/[ ]+/g, '-');
  bow.spec = (record?.spec ?? '').replace(/([a-z])([A-Z])/g, '$1 $2');
  bow.headId = record?.build?.gear?.head ?? 0;
  const points = record?.build?.talents?.points ?? [];
  bow.points = points.reduce((sum, count) => sum + count, 0);
  bow.split = points.join('/');
  await signInThroughProxy(context, WEB, session);
  await serveShells(context, `${process.cwd()}/dist`);
  await page.goto('/');
  await page.evaluate(() => {
    window.localStorage.clear();
    window.sessionStorage.clear();
  });
});

async function shot(page: Page, name: string): Promise<void> {
  if (CAPTURES === undefined) return;
  await page.waitForTimeout(400);
  await page.screenshot({ path: `${CAPTURES}/w12-${name}.png` });
}

const selector = (page: Page) => page.getByTestId('character-selector');

async function headerIs(page: Page, name: string, line: string): Promise<void> {
  await expect(page.getByTestId('selector-name')).toHaveText(name);
  await expect(selector(page)).toContainText(line);
}

async function choose(page: Page, key: string): Promise<void> {
  await selector(page).click();
  await page.getByTestId(`selector-row-${key}`).getByRole('button').first().click();
}

test('the chosen character follows the player through every page', async ({ page }) => {
  // 1. Choose Bow Jackzon in the header selector.
  await page.goto('/tiers');
  await expect(page.getByTestId('selector-name')).toBeVisible();
  // The selector open, the choice highlighted.
  await selector(page).click();
  const bowRow = page.getByTestId(`selector-row-${BOW}`);
  await bowRow.getByRole('button').first().hover();
  await shot(page, '01a-selector-open');
  await bowRow.getByRole('button').first().click();
  await headerIs(page, 'Bow Jackzon', `${bow.spec} Hunter · 60`);
  expect(await pointerOf(page)).toMatchObject({
    source: 'armory',
    ref: BOW,
    label: `Bow Jackzon · ${bow.spec} Hunter`,
    classSlug: 'hunter',
  });
  await shot(page, '01-selector-chosen');

  // 2. The planner opens on that character: its race, level, talents and gear, not the first legal race.
  await page.goto('/planner');
  await expect(nativeSelect(page, 'Class')).toHaveValue('hunter');
  await expect(nativeSelect(page, 'Race')).toHaveValue(bow.raceSlug);
  await expect(page.getByTestId('planner-code-note')).toContainText('Talents loaded from a character');
  const main = page.getByRole('main');
  await expect(main).toContainText(`Hunter talents · ${bow.race} · Level 60 · ${bow.spec} Hunter`);
  await expect(main).toContainText('Bow Jackzon');
  const spent = Number((await main.innerText()).match(/SPENT\s+(\d+)\/51/)?.[1] ?? '0');
  expect(spent).toBeGreaterThan(0);
  // The planner shows the points the record carries, never "25/51" beside "has spent 51".
  expect(spent).toBe(bow.points);
  await expect(page.getByTestId('planner-build-origin')).toHaveCount(0);
  await expect(page.getByTestId('character-header-action')).toHaveText('Send this build to Bow Jackzon');
  const simLink = page.getByTestId('planner-sim-link');
  // "Sim this build" is in the card under Send, inside the first screen.
  const sendBox = (await page.getByTestId('character-header-action').boundingBox())!;
  const simBox = (await simLink.boundingBox())!;
  expect(simBox.y).toBeGreaterThan(sendBox.y);
  expect(simBox.y + simBox.height).toBeLessThan(900);
  const href = (await simLink.getAttribute('href')) ?? '';
  const code = new URL(href, WEB).searchParams.get('code') ?? '';
  expect(code).toContain(`hunter:${bow.raceSlug}`);
  expect(code).toContain(`head=${bow.headId}`);
  expect(code).toContain('who=Bow%20Jackzon');
  expect(code).toContain('level=60');
  await expect(selector(page)).toContainText('Bow Jackzon');
  expect(await pointerOf(page)).toMatchObject({ source: 'armory', ref: BOW });
  await shot(page, '02-planner');

  // 3. "Sim this build": the simulator shows the character, not a bare code.
  await simLink.click();
  await expect(page.getByTestId('sim-character-name')).toHaveText('Bow Jackzon');
  await expect(page.getByTestId('sim-character-descriptor')).toContainText(
    `${bow.race} ${bow.spec} Hunter · 60`,
  );
  await expect(page.getByTestId('sim-source-pill')).toHaveText('Build from planner · Bow Jackzon’s gear');
  await expect(page.getByTestId('sim-gear')).toBeVisible();
  await expect(page.getByTestId('sim-no-gear')).toHaveCount(0);
  await headerIs(page, 'Bow Jackzon', `${bow.spec} Hunter · 60`);
  expect(await pointerOf(page)).toMatchObject({
    source: 'armory',
    ref: BOW,
    label: `Bow Jackzon · ${bow.spec} Hunter`,
  });
  // The spine bar's Plan door opens the planner on the account character, not a class.
  await expect(page.getByTestId('current-character-bar-plan')).toHaveAttribute(
    'href',
    `/planner?character=${encodeURIComponent(BOW)}`,
  );
  await expect(page.getByTestId('current-character-bar-rankings')).toContainText('Rankings for Hunter');
  await shot(page, '03-sim');
  // A bare /sim still restores the same character, and still leaves the pointer alone.
  await page.goto('/sim');
  await expect(page.getByTestId('sim-character-name')).toHaveText('Bow Jackzon');
  await expect(page.getByText('Sign in with Battle.net to find your characters')).toHaveCount(0);
  expect(await pointerOf(page)).toMatchObject({ source: 'armory', ref: BOW });

  // 4. The character's own page reads as owned.
  await page.goto('/character/us/pvp/bow-jackzon');
  // Owned: the spine bar speaks for this character, and the page reads its gear against the player's levels.
  // The page's card carries the identity, so the strip shows only its doors.
  await expect(page.getByTestId('current-character-bar-identity')).toHaveCount(0);
  await expect(page.getByTestId('current-character-bar-rankings')).toContainText('Rankings for Hunter');
  await expect(page.getByRole('main')).toContainText('from your talents');
  // One spec: the owner's talents decide it, spelt as a player says it, with the split beside it.
  await expect(page.getByRole('main')).toContainText(`${bow.spec} Hunter · ${bow.split}`);
  await expect(page.getByRole('main')).not.toContainText('BeastMastery');
  await expect(page.getByRole('main')).toContainText(/[0-9]+ of [0-9]+ slots best in slot/);
  await headerIs(page, 'Bow Jackzon', `${bow.spec} Hunter · 60`);
  await shot(page, '04-character');

  // 5. Rankings, BiS and the account page speak for the same character.
  await page.goto('/rankings');
  await expect(page.getByRole('main')).toContainText(`${bow.spec} Hunters`);
  await expect(page.getByRole('main')).toContainText('Bow Jackzon');
  await shot(page, '05-rankings');
  await page.goto('/bis');
  await expect(page.getByRole('main')).toContainText('Bow Jackzon');
  await expect(page.getByRole('main')).toContainText(bow.spec);
  await shot(page, '06-bis');
  await page.goto('/account');
  await expect(page.getByRole('main')).toContainText(/CURRENT CHARACTER\s+Bow Jackzon/i);
  await expect(page.getByTestId('account-hero')).toContainText(`${bow.race} ${bow.spec} Hunter · Level 60`);
  await expect(page.getByTestId('character-current-pill')).toHaveCount(1);
  await shot(page, '07-account');

  // 6. A guild member's gear opens in the planner without changing the header.
  await page.goto('/guild/us/pvp/lifecycle-test#roster');
  const ashenveil = page.getByText('Ashenveil').first();
  await expect(ashenveil).toBeVisible();
  // Ashenveil's row, by her class: the link carries her own export. A roster that holds her back
  // for approval has no such link (the shared local guild moves between walks), so the same
  // export, named as her addon names it, opens the planner directly.
  const rosterLink = page.locator('a[href*="%3Apaladin%3A"]', { hasText: 'Open in planner' }).first();
  if ((await rosterLink.count()) > 0) {
    await rosterLink.click();
  } else {
    const input = await fetch(`${API}/v1/characters/us/pvp/ashenveil/sim-input`);
    const exported = ((await input.json()) as { data: { gear: string } }).data.gear;
    const named = exported.includes('who=') ? exported : `${exported}|who=Ashenveil:Realm`;
    await page.goto(`/planner?code=${encodeURIComponent(named)}`);
  }
  await expect(nativeSelect(page, 'Class')).toHaveValue('paladin');
  await headerIs(page, 'Bow Jackzon', `${bow.spec} Hunter · 60`);
  expect(await pointerOf(page)).toMatchObject({ source: 'armory', ref: BOW });
  // Whose build it is, and never a Paladin build sent to a Hunter.
  await expect(page.getByTestId('planner-build-origin')).toContainText("Ashenveil's build ·");
  await expect(page.getByTestId('planner-build-origin')).toContainText('Paladin');
  await expect(page.getByTestId('character-header-action')).toHaveText(
    'This is a Paladin build; Bow Jackzon is a Hunter',
  );
  await expect(page.getByTestId('character-header-action')).toBeDisabled();
  await shot(page, '08-guild-gear-in-planner');

  // 7. Forget clears the pointer and the header falls back to the account's main.
  await page.goto('/sim');
  await expect(page.getByTestId('sim-character-name')).toHaveText('Bow Jackzon');
  await page.getByTestId('current-character-forget').click();
  expect(await pointerOf(page)).toBeNull();
  await page.goto('/tiers');
  await expect(page.getByTestId('selector-name')).toHaveText('Pipsqueak');
  await shot(page, '09-forgotten');
  // The simulator's landing lists the signed-in account's characters, never a sign-in prompt.
  await page.goto('/sim');
  await expect(page.getByText('Sign in with Battle.net to find your characters')).toHaveCount(0);
  await expect(page.getByRole('main')).toContainText('Bow Jackzon');
  await expect(page.getByRole('main')).toContainText('Pipsqueak');
  // The current character (the main, after Forget) is listed first, and a Holy Priest has no gold
  // "Run sim": a quiet "Not simulated yet" stands where the button would be, in the card and the row.
  const listed = await page
    .locator('[data-testid^="sim-character-link-"]')
    .evaluateAll((els) => els.map((el) => el.getAttribute('data-testid')));
  expect(listed).toEqual([`sim-character-link-${PIPSQUEAK}`, `sim-character-link-${BOW}`]);
  await expect(page.getByTestId('sim-run-block-action')).toHaveCount(0);
  await expect(page.getByTestId('sim-run-block-not-simulated')).toHaveText('Not simulated yet');
  await expect(page.getByTestId(`sim-pick-${PIPSQUEAK}`)).toHaveCount(0);
  await expect(page.getByTestId(`sim-pick-${BOW}`)).toBeVisible();
  // The card carries the identity once: the strip above it shows only its doors.
  await expect(page.getByTestId('current-character-bar-identity')).toHaveCount(0);
  await shot(page, '13-sim-landing');
});

test('a class tile opens a class and leaves the character alone', async ({ page }) => {
  await page.goto('/tiers');
  await choose(page, BOW);
  await headerIs(page, 'Bow Jackzon', `${bow.spec} Hunter · 60`);
  await page.goto('/planner?class=mage');
  await expect(nativeSelect(page, 'Class')).toHaveValue('mage');
  await expect(page.getByTestId('planner-sim-link')).toHaveCount(0);
  await headerIs(page, 'Bow Jackzon', `${bow.spec} Hunter · 60`);
  expect(await pointerOf(page)).toMatchObject({ source: 'armory', ref: BOW });
  await expect(page.getByTestId('character-header-action')).toHaveText(
    'This is a Mage build; Bow Jackzon is a Hunter',
  );
  await expect(page.getByTestId('character-header-action')).toBeDisabled();
  await shot(page, '10-class-tile');
});

test('a linked code never replaces the character, and a nameless one is never the pointer', async ({
  page,
}) => {
  await page.goto('/tiers');
  await choose(page, BOW);
  const nameless = 'FS1:1.60.1.70450:hunter:human:0/0/0:';
  await page.goto(`/sim?code=${encodeURIComponent(nameless)}`);
  await expect(page.getByTestId('sim-character')).toBeVisible();
  await expect(page.getByTestId('sim-character-name')).toHaveCount(0);
  await expect(page.getByTestId('sim-source-pill')).toHaveText('Entered by hand');
  expect(await pointerOf(page)).toMatchObject({ source: 'armory', ref: BOW });
  await page.goto(`/planner?code=${encodeURIComponent(nameless)}`);
  await expect(nativeSelect(page, 'Class')).toHaveValue('hunter');
  expect(await pointerOf(page)).toMatchObject({ source: 'armory', ref: BOW });
  await headerIs(page, 'Bow Jackzon', `${bow.spec} Hunter · 60`);
  await shot(page, '11-linked-code');
});

test('pasting Pipsqueak’s export makes Pipsqueak the current character', async ({ page }) => {
  const session = await devLogin('hunthrop#1894');
  const input = await fetch(`${API}/v1/characters/us/pvp/pipsqueak/sim-input`, {
    headers: { cookie: session.cookie },
  });
  const exported = ((await input.json()) as { data: { gear: string } }).data.gear;
  // The seed's export carries no name section; the addon's carries one.
  const named = exported.includes('who=') ? exported : `${exported}|who=Pipsqueak:Realm`;
  await page.goto('/tiers');
  await choose(page, BOW);
  await selector(page).click();
  await page.getByTestId('selector-paste').click();
  await page.getByTestId('selector-paste-input').fill(named);
  // Bow Jackzon is current until the paste is used: the capture shows the box, then the change.
  await headerIs(page, 'Bow Jackzon', `${bow.spec} Hunter · 60`);
  await shot(page, '12a-paste-box');
  await page.getByTestId('selector-paste-use').click();
  await headerIs(page, 'Pipsqueak', 'Holy Priest · 60');
  expect(await pointerOf(page)).toMatchObject({ source: 'armory', ref: PIPSQUEAK });
  await shot(page, '12-pasted-export');
});

test('signing in settles the pointer on the main, and signing out clears it', async ({ page }) => {
  // A previous visit left a pointer for a character this account does not hold: sign-in settles it.
  await page.evaluate(() =>
    window.localStorage.setItem(
      'fs.currentCharacter',
      JSON.stringify({
        source: 'armory',
        ref: 'us/pvp/someone-else',
        label: 'Someone Else · Fury Warrior',
        classSlug: 'warrior',
        savedAt: new Date().toISOString(),
      }),
    ),
  );
  await page.goto('/tiers');
  await expect(page.getByTestId('selector-name')).toHaveText('Pipsqueak');
  expect(await pointerOf(page)).toMatchObject({ source: 'armory', ref: PIPSQUEAK });
  await shot(page, '14-sign-in-settled');
  // Sign out from the account page: the pointer goes with the session.
  await page.goto('/account');
  await expect(page.getByTestId('account-hero')).toBeVisible();
  await page.getByRole('button', { name: /sign out/i }).click();
  await expect.poll(async () => await pointerOf(page)).toBeNull();
  await page.goto('/tiers');
  await expect(page.getByTestId('selector-name')).toHaveCount(0);
  await shot(page, '15-signed-out');
});

test.describe('at 360 px', () => {
  test.use({ viewport: { width: 360, height: 800 } });

  test('the planner card, the simulator card and the account card hold at phone width', async ({ page }) => {
    // The header chooser is a desktop control: choose there, then narrow the window.
    await page.setViewportSize({ width: 1280, height: 900 });
    await page.goto('/tiers');
    await choose(page, BOW);
    await page.setViewportSize({ width: 360, height: 800 });
    await page.goto('/planner');
    const card = page.getByTestId('character-header');
    await expect(card).toContainText('Bow Jackzon');
    await expect(page.getByTestId('character-header-action')).toHaveText('Send this build to Bow Jackzon');
    await expect(page.getByTestId('planner-sim-link')).toBeVisible();
    // The card itself holds at phone width: nothing in it is wider than it is, and it sits inside 360.
    const fits = await card.evaluate((el) => el.scrollWidth <= el.clientWidth);
    expect(fits).toBe(true);
    expect((await card.boundingBox())!.x + (await card.boundingBox())!.width).toBeLessThanOrEqual(360);
    await card.scrollIntoViewIfNeeded();
    if (CAPTURES !== undefined) await card.screenshot({ path: `${CAPTURES}/w12-16-planner-card-360.png` });

    await page.getByTestId('planner-sim-link').click();
    await expect(page.getByTestId('sim-character-name')).toHaveText('Bow Jackzon');
    expect(await page.evaluate(() => document.documentElement.scrollWidth)).toBeLessThanOrEqual(360);
    if (CAPTURES !== undefined)
      await page.getByTestId('sim-character').screenshot({ path: `${CAPTURES}/w12-17-sim-card-360.png` });

    await page.goto('/account');
    await expect(page.getByTestId('account-hero')).toBeVisible();
    expect(await page.evaluate(() => document.documentElement.scrollWidth)).toBeLessThanOrEqual(360);
    await page.waitForTimeout(800); // the page's reveal fade
    if (CAPTURES !== undefined)
      await page.getByTestId('account-hero').screenshot({ path: `${CAPTURES}/w12-18-account-card-360.png` });
  });
});
