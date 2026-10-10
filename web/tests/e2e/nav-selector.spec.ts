// web/tests/e2e/nav-selector.spec.ts
// The site navigation bar and its character selector (design/specs/2026-10-09-nav-character-
// selector.md): the doors in order with BiS, the selector's closed states and widths, the list
// (order, stale, failed, filter), choosing a character, the paste flow, the signed-out and
// expired states, the keyboard path, and the phone Menu and sheet.
import { expect, test, type Page } from '@playwright/test';
import { ACTIVE_BUILD } from './support/active-build';
import {
  OBNOXIOUS,
  armoryPointer,
  keyOf,
  signInWith,
  signedOut,
  storePointer,
  type FixtureCharacter,
} from './support/selector';

const DOORS = ['Planner', 'BiS', 'Simulator', 'Logs', 'Rankings', 'Tier List', 'Guides', 'Get set up'];
const FURY = `FS1:${ACTIVE_BUILD}:warrior:orc:0/5530515/0:head=12640,main_hand=11726`;

const FROSTBYTE: FixtureCharacter = {
  name: 'Frostbyte',
  class: 'Mage',
  spec: 'Frost',
  level: 42,
  faction: 'horde',
  ruleset: 'hardcore',
  buildDaysAgo: 3,
  buildSource: 'blizzard',
};
const HORDE_FURY: FixtureCharacter = {
  name: 'Rageclaw',
  class: 'Warrior',
  spec: 'Fury',
  level: 34,
  faction: 'horde',
  buildDaysAgo: 1,
};
const SHADOWMEND: FixtureCharacter = {
  name: 'Shadowmend',
  class: 'Priest',
  spec: 'Shadow',
  level: 35,
  faction: 'horde',
  buildDaysAgo: 2,
};
const OAKHEART: FixtureCharacter = {
  name: 'Oakheart',
  class: 'Paladin',
  spec: 'Protection',
  buildDaysAgo: 3,
  buildSource: 'blizzard',
  syncError: 'refresh failed',
};
const TREEWALKER: FixtureCharacter = {
  name: 'Treewalker',
  class: 'Druid',
  spec: 'Balance',
  buildDaysAgo: 19,
};

const selector = (page: Page) => page.getByTestId('character-selector');
const panel = (page: Page) => page.getByTestId('selector-panel');
const rowButton = (page: Page, character: FixtureCharacter) =>
  panel(page)
    .getByTestId(`selector-row-${keyOf(character)}`)
    .getByRole('button')
    .first();

async function ready(page: Page): Promise<void> {
  await expect(selector(page)).toBeVisible();
}

async function barBox(page: Page) {
  return page.getByRole('banner').boundingBox();
}

async function noSideScroll(page: Page) {
  return page.evaluate(() => document.documentElement.scrollWidth <= document.documentElement.clientWidth);
}

test.describe('the doors', () => {
  test.skip(() => test.info().project.name !== 'desktop', 'desktop layout only');
  test.use({ viewport: { width: 1440, height: 900 } });

  test('are in order, with BiS where Leveling BiS was, and the door goes to /bis', async ({ page }) => {
    await signedOut(page);
    await page.goto('/');
    const links = page.getByTestId('primary-nav').getByRole('link');
    await expect(links).toHaveText(DOORS);
    await expect(page.getByRole('banner')).not.toContainText('Leveling BiS');
    await page.getByTestId('primary-nav').getByRole('link', { name: 'BiS', exact: true }).click();
    await expect(page).toHaveURL(/\/bis$/);
    await expect(page.getByRole('heading', { name: 'BiS', level: 1 })).toBeVisible();
  });

  test('Discord shows its label at 1440 and the mark with a tooltip below', async ({ page }) => {
    await signedOut(page);
    await page.goto('/');
    const discord = page.getByTestId('discord-link');
    await expect(discord).toHaveAttribute('aria-label', 'Discord');
    await expect(discord.locator('.site-discord-label')).toBeVisible();
    await page.setViewportSize({ width: 1439, height: 900 });
    await expect(discord.locator('.site-discord-label')).toBeHidden();
    await discord.hover();
    await expect(page.getByRole('tooltip', { name: 'Discord' })).toBeVisible();
  });
});

test.describe('the bar at every width', () => {
  test.skip(() => test.info().project.name !== 'desktop', 'desktop layout only');

  const widths = [
    { width: 2000, height: 81 },
    { width: 1440, height: 81 },
    { width: 1439, height: 81 },
    { width: 1280, height: 81 },
    { width: 1279, height: 81 },
    { width: 1024, height: 81 },
    { width: 1023, height: 56 },
    { width: 768, height: 56 },
  ];
  for (const { width, height } of widths) {
    test(`${width}px: one row ${height}px tall, no sideways scroll, the full name unclipped`, async ({
      page,
    }) => {
      await signInWith(page, [OBNOXIOUS]);
      await storePointer(page, armoryPointer(OBNOXIOUS));
      await page.setViewportSize({ width, height: 900 });
      await page.goto('/guides');
      await ready(page);
      expect((await barBox(page))!.height).toBeGreaterThanOrEqual(height);
      expect((await barBox(page))!.height).toBeLessThanOrEqual(height + 1);
      expect(await noSideScroll(page)).toBe(true);
      const name = width >= 1024 ? page.getByTestId('selector-name') : page.locator('.csel-name-bar');
      await expect(name).toHaveText('Obnoxious Yell');
      const clipped = await name.evaluate((el) => el.scrollWidth > el.clientWidth);
      expect(clipped).toBe(false);
      if (width >= 1024) {
        const doors = await page
          .getByTestId('primary-nav')
          .getByRole('link')
          .evaluateAll((links) => links.map((link) => Math.round(link.getBoundingClientRect().top)));
        expect(new Set(doors).size).toBe(1);
      }
    });
  }

  test('1280 and up show spec, class and level; 1024 to 1279 show spec and level', async ({ page }) => {
    await signInWith(page, [OBNOXIOUS]);
    await storePointer(page, armoryPointer(OBNOXIOUS));
    await page.setViewportSize({ width: 1280, height: 900 });
    await page.goto('/guides');
    await expect(page.getByTestId('selector-sub')).toHaveText('Fury Warrior · 60');
    await page.setViewportSize({ width: 1279, height: 900 });
    await expect(page.locator('.csel-sub-short')).toHaveText('Fury · 60');
    await expect(page.getByTestId('selector-sub')).toBeHidden();
  });

  test('the wordmark and Discord sit on the 1344px content edges at 2000', async ({ page }) => {
    await signInWith(page, [OBNOXIOUS]);
    await storePointer(page, armoryPointer(OBNOXIOUS));
    await page.setViewportSize({ width: 2000, height: 900 });
    await page.goto('/guides');
    await ready(page);
    const mark = (await page.getByRole('banner').getByRole('link', { name: 'Forever Sixty' }).boundingBox())!;
    const discord = (await page.getByTestId('discord-link').boundingBox())!;
    expect(mark.x).toBeGreaterThanOrEqual(328 - 1);
    expect(discord.x + discord.width).toBeLessThanOrEqual(1672 + 1);
  });

  test('a 22-character name keeps the bar one row and ends in an ellipsis, with the whole name in title', async ({
    page,
  }) => {
    const long: FixtureCharacter = { ...OBNOXIOUS, name: 'Bartholomewtheunbearable' };
    await signInWith(page, [long]);
    await storePointer(page, armoryPointer(long));
    await page.setViewportSize({ width: 1440, height: 900 });
    await page.goto('/guides');
    await ready(page);
    expect((await barBox(page))!.height).toBeLessThanOrEqual(82);
    await expect(selector(page)).toHaveAttribute('title', long.name);
    const clipped = await page.getByTestId('selector-name').evaluate((el) => el.scrollWidth > el.clientWidth);
    expect(clipped).toBe(true);
  });
});

test.describe('the list', () => {
  test.skip(() => test.info().project.name !== 'desktop', 'desktop layout only');
  test.use({ viewport: { width: 1440, height: 900 } });

  test('opens under the selector with the current character first, then last chosen, then newest sync', async ({
    page,
  }) => {
    await signInWith(page, [TREEWALKER, OAKHEART, SHADOWMEND, FROSTBYTE, OBNOXIOUS]);
    await storePointer(page, armoryPointer(OBNOXIOUS));
    await page.addInitScript(
      (keys) => window.localStorage.setItem('fs.characterOrder', JSON.stringify(keys)),
      [keyOf(FROSTBYTE), keyOf(SHADOWMEND)],
    );
    await page.goto('/guides');
    await ready(page);
    await selector(page).click();
    await expect(panel(page)).toBeVisible();
    await expect(panel(page)).toHaveAttribute('aria-label', 'Choose a character');
    const names = panel(page).locator('.csel-r1 > span:first-child');
    await expect(names).toHaveText(['Obnoxious Yell', 'Frostbyte', 'Shadowmend', 'Oakheart', 'Treewalker']);
    await expect(panel(page).locator('[aria-current="true"]')).toContainText('Obnoxious Yell');
    await expect(panel(page).getByTestId('selector-fade')).toHaveCount(0);
    await expect(panel(page).getByTestId('selector-filter')).toHaveCount(0);
    await expect(panel(page).getByText('Your characters')).toBeVisible();
    await expect(panel(page).getByTestId('selector-manage')).toHaveAttribute('href', '/account#characters');
    const box = (await panel(page).boundingBox())!;
    const trigger = (await selector(page).boundingBox())!;
    expect(box.width).toBe(392);
    expect(Math.round(box.x + box.width)).toBe(Math.round(trigger.x + trigger.width));
  });

  test('marks a stale character, a failed one with Retry, and Retry only when the API reports an error', async ({
    page,
  }) => {
    await signInWith(page, [TREEWALKER, OAKHEART, SHADOWMEND, OBNOXIOUS]);
    await storePointer(page, armoryPointer(OBNOXIOUS));
    await page.goto('/guides');
    await ready(page);
    await selector(page).click();
    const stale = panel(page).getByTestId(`selector-row-${keyOf(TREEWALKER)}`);
    await expect(stale).toContainText('Addon · 19 days ago');
    await expect(stale).toContainText('Log in to the game to update');
    const failed = panel(page).getByTestId(`selector-row-${keyOf(OAKHEART)}`);
    await expect(failed).toContainText('Battle.net refresh failed · last good 3 days ago');
    await expect(panel(page).getByTestId(`selector-retry-${keyOf(OAKHEART)}`)).toBeVisible();
    await expect(panel(page).getByTestId(`selector-retry-${keyOf(TREEWALKER)}`)).toHaveCount(0);
    await expect(panel(page).getByTestId(`selector-retry-${keyOf(SHADOWMEND)}`)).toHaveCount(0);
  });

  test('puts a ring mark and a tooltip on the closed crest when the current character is stale', async ({
    page,
  }) => {
    await signInWith(page, [TREEWALKER, OBNOXIOUS]);
    await storePointer(page, armoryPointer(TREEWALKER));
    await page.goto('/guides');
    await ready(page);
    await expect(page.getByTestId('selector-mark')).toHaveAttribute('data-mark', 'stale');
    await expect(selector(page)).toHaveAttribute('aria-label', /Sync is stale\. Change character\.$/);
    await selector(page).hover();
    await expect(page.getByRole('tooltip', { name: /Last synced 19 days ago/ })).toBeVisible();
  });

  test('shows the Get the addon row only without an addon character', async ({ page }) => {
    await signInWith(page, [FROSTBYTE]);
    await storePointer(page, armoryPointer(FROSTBYTE));
    await page.goto('/guides');
    await ready(page);
    await selector(page).click();
    await expect(panel(page).getByTestId('selector-addon')).toBeVisible();
    await expect(panel(page).getByTestId('selector-addon')).toHaveAttribute('href', '/setup');
    await page.keyboard.press('Escape');

    await page.unroute('**/v1/me');
    await signInWith(page, [FROSTBYTE, OBNOXIOUS]);
    await page.reload();
    await ready(page);
    await selector(page).click();
    await expect(panel(page).getByTestId('selector-addon')).toHaveCount(0);
  });

  test('scrolls under a fade from the seventh character and filters from the ninth', async ({ page }) => {
    const many: FixtureCharacter[] = Array.from({ length: 9 }, (_, index) => ({
      ...OBNOXIOUS,
      name: index === 0 ? 'Obnoxious Yell' : `Altchar${String.fromCharCode(97 + index)}`,
      buildDaysAgo: 0.01 + index,
    }));
    await signInWith(page, many);
    await storePointer(page, armoryPointer(many[0]));
    await page.goto('/guides');
    await ready(page);
    await selector(page).click();
    await expect(panel(page).getByTestId('selector-fade')).toBeVisible();
    const filter = panel(page).getByTestId('selector-filter');
    await expect(filter).toBeVisible();
    await filter.fill('altchard');
    await expect(panel(page).locator('.csel-row')).toHaveCount(1);
  });

  test('closes on Escape, on an outside click, and when focus leaves', async ({ page }) => {
    await signInWith(page, [OBNOXIOUS, FROSTBYTE]);
    await storePointer(page, armoryPointer(OBNOXIOUS));
    await page.goto('/guides');
    await ready(page);
    await selector(page).click();
    await expect(panel(page)).toBeVisible();
    await page.keyboard.press('Escape');
    await expect(panel(page)).toBeHidden();
    await expect(selector(page)).toBeFocused();
    await selector(page).click();
    await page.mouse.click(700, 600);
    await expect(panel(page)).toBeHidden();
  });
});

test.describe('choosing a character', () => {
  test.skip(() => test.info().project.name !== 'desktop', 'desktop layout only');
  test.use({ viewport: { width: 1440, height: 900 } });

  test('writes the pointer, moves the page along without a reload, and remembers the choice', async ({
    page,
  }) => {
    await signInWith(page, [OBNOXIOUS, FROSTBYTE, SHADOWMEND]);
    await storePointer(page, armoryPointer(OBNOXIOUS));
    await page.goto('/rankings');
    await ready(page);
    await page.evaluate(() => {
      (window as unknown as { __kept: boolean }).__kept = true;
    });
    await selector(page).click();
    await panel(page)
      .getByTestId(`selector-row-${keyOf(FROSTBYTE)}`)
      .getByRole('button')
      .first()
      .click();
    await expect(panel(page)).toBeHidden();
    await expect(selector(page)).toBeFocused();
    await expect(page.getByTestId('selector-name')).toHaveText('Frostbyte');
    await expect(page.getByTestId('selector-live')).toHaveText('Now using Frostbyte.');
    expect(await page.evaluate(() => (window as unknown as { __kept?: boolean }).__kept)).toBe(true);
    const stored = await page.evaluate(() => ({
      pointer: JSON.parse(window.localStorage.getItem('fs.currentCharacter') ?? 'null'),
      order: JSON.parse(window.localStorage.getItem('fs.characterOrder') ?? 'null'),
      cls: document.documentElement.dataset.pointerClass,
    }));
    expect(stored.pointer).toMatchObject({ source: 'armory', ref: keyOf(FROSTBYTE), classSlug: 'mage' });
    expect(stored.pointer.label).toBe('Frostbyte · Frost Mage');
    expect(stored.order).toEqual([keyOf(FROSTBYTE)]);
    expect(stored.cls).toBe('mage');
    await expect(page.getByTestId('current-character-bar-identity')).toContainText('Frostbyte');
    await page.goto('/tiers');
    await ready(page);
    await expect(page.getByTestId('selector-name')).toHaveText('Frostbyte');
  });

  test('on /planner, points the URL at the chosen character and opens its class', async ({ page }) => {
    await signInWith(page, [OBNOXIOUS, FROSTBYTE]);
    await storePointer(page, armoryPointer(OBNOXIOUS));
    await page.goto('/planner');
    await ready(page);
    await selector(page).click();
    await panel(page)
      .getByTestId(`selector-row-${keyOf(FROSTBYTE)}`)
      .getByRole('button')
      .first()
      .click();
    await expect(page).toHaveURL(/\/planner\?class=mage/);
    await expect(page.getByTestId('selector-name')).toHaveText('Frostbyte');
  });

  async function choose(page: Page, character: FixtureCharacter): Promise<void> {
    await page.evaluate(() => {
      (window as unknown as { __kept: boolean }).__kept = true;
    });
    await selector(page).click();
    await rowButton(page, character).click();
    await expect(page.getByTestId('selector-name')).toHaveText(character.name);
  }

  async function pageWasNotReloaded(page: Page): Promise<boolean> {
    return page.evaluate(() => (window as unknown as { __kept?: boolean }).__kept === true);
  }

  test('on /tiers, the callout and the YOUR SPEC badge move to the new character without a reload', async ({
    page,
  }) => {
    await signInWith(page, [OBNOXIOUS, FROSTBYTE]);
    await storePointer(page, armoryPointer(OBNOXIOUS));
    await page.goto('/tiers');
    await ready(page);
    await expect(page.getByTestId('tier-callout')).toContainText('Fury Warrior is');
    await expect(page.getByTestId('tier-row-warrior-fury').first()).toHaveClass(/is-you/);
    await choose(page, FROSTBYTE);
    await expect(page.getByTestId('tier-callout')).toContainText('Frost Mage is');
    await expect(page.getByTestId('tier-callout')).not.toContainText('Fury Warrior is');
    await expect(page.getByTestId('tier-row-mage-frost').first()).toHaveClass(/is-you/);
    await expect(page.getByTestId('tier-row-warrior-fury').first()).not.toHaveClass(/is-you/);
    await expect(page.locator('[data-testid="tier-you-pill"]:visible')).toHaveCount(1);
    expect(await pageWasNotReloaded(page)).toBe(true);
  });

  test('on /bis, a character of another class moves the page to that character’s own spec page, band and faction', async ({
    page,
  }) => {
    await signInWith(page, [OBNOXIOUS, FROSTBYTE]);
    await storePointer(page, armoryPointer(OBNOXIOUS));
    await page.goto('/bis/warrior/fury#band-alliance-20');
    await ready(page);
    await expect(page.getByTestId('bis-character-card')).toContainText('Obnoxious Yell');
    await choose(page, FROSTBYTE);
    await expect(page).toHaveURL(/\/bis\/mage\/frost\?faction=horde#band-horde-40$/);
    await expect(page.getByRole('heading', { level: 1 })).toContainText('Frost');
    await expect(page.getByTestId('bis-character-card')).toContainText('Frostbyte');
    await expect(page.getByTestId('bis-character-card')).not.toContainText('Obnoxious Yell');
    await expect(page.getByTestId('bis-character-card')).not.toContainText('Switch');
  });

  test('on /bis, a same-spec character updates band and faction in place, with no reload', async ({
    page,
  }) => {
    await signInWith(page, [OBNOXIOUS, HORDE_FURY]);
    await storePointer(page, armoryPointer(OBNOXIOUS));
    await page.goto('/bis/warrior/fury#band-alliance-60');
    await ready(page);
    await choose(page, HORDE_FURY);
    await expect(page).toHaveURL(/\/bis\/warrior\/fury\?faction=horde#band-horde-30$/);
    await expect(page.getByTestId('bis-faction-panel-horde')).toBeVisible();
    await expect(page.getByTestId('bis-band-horde-30')).toBeVisible();
    await expect(page.getByTestId('bis-character-card')).toContainText('Rageclaw');
    expect(await pageWasNotReloaded(page)).toBe(true);
  });

  test('on /planner, the open planner switches to the new character in place', async ({ page }) => {
    await signInWith(page, [OBNOXIOUS, FROSTBYTE]);
    await storePointer(page, armoryPointer(OBNOXIOUS));
    await page.goto('/planner');
    await ready(page);
    await expect(page.getByTestId('planner-header-h1')).toContainText('Warrior');
    await choose(page, FROSTBYTE);
    await expect(page.getByTestId('planner-header-h1')).toContainText('Mage');
    await expect(page).toHaveURL(/\/planner\?class=mage/);
    await expect(page.getByTestId('planner-character-card')).toContainText('Frostbyte');
    expect(await pageWasNotReloaded(page)).toBe(true);
  });

  test('a faction shows as its real emblem with its name, never an abstract mark', async ({ page }) => {
    await signInWith(page, [OBNOXIOUS, FROSTBYTE]);
    await storePointer(page, armoryPointer(OBNOXIOUS));
    await page.goto('/guides');
    await ready(page);
    await selector(page).click();
    const alliance = panel(page)
      .getByTestId(`selector-row-${keyOf(OBNOXIOUS)}`)
      .getByRole('img', { name: 'Alliance' });
    const horde = panel(page)
      .getByTestId(`selector-row-${keyOf(FROSTBYTE)}`)
      .getByRole('img', { name: 'Horde' });
    await expect(alliance).toHaveAttribute('src', /\/icons\/hd\/faction\/alliance\.webp$/);
    await expect(horde).toHaveAttribute('src', /\/icons\/hd\/faction\/horde\.webp$/);
    await expect(horde).toHaveAttribute('title', 'Horde');
    expect(await horde.evaluate((img) => (img as HTMLImageElement).width)).toBe(16);
  });

  test('walks the list with the keyboard: open, arrow, choose, focus returns', async ({ page }) => {
    await signInWith(page, [OBNOXIOUS, FROSTBYTE, SHADOWMEND]);
    await storePointer(page, armoryPointer(OBNOXIOUS));
    await page.goto('/guides');
    await ready(page);
    await selector(page).focus();
    await page.keyboard.press('ArrowDown');
    await expect(panel(page)).toBeVisible();
    await expect(panel(page).locator('[aria-current="true"]')).toBeFocused();
    // The never-chosen tail is newest sync first: Shadowmend (2 days) before Frostbyte (3).
    await page.keyboard.press('ArrowDown');
    await expect(rowButton(page, SHADOWMEND)).toBeFocused();
    await page.keyboard.press('f');
    await expect(rowButton(page, FROSTBYTE)).toBeFocused();
    await page.keyboard.press('Home');
    await expect(panel(page).locator('[aria-current="true"]')).toBeFocused();
    await page.keyboard.press('End');
    await expect(rowButton(page, FROSTBYTE)).toBeFocused();
    await page.keyboard.press('Enter');
    await expect(panel(page)).toBeHidden();
    await expect(selector(page)).toBeFocused();
    await expect(page.getByTestId('selector-name')).toHaveText('Frostbyte');
  });
});

test.describe('paste an export', () => {
  test.skip(() => test.info().project.name !== 'desktop', 'desktop layout only');
  test.use({ viewport: { width: 1440, height: 900 } });

  test('rejects a string that is not an export and keeps the box', async ({ page }) => {
    await signInWith(page, [OBNOXIOUS]);
    await storePointer(page, armoryPointer(OBNOXIOUS));
    await page.goto('/guides');
    await ready(page);
    await selector(page).click();
    await panel(page).getByTestId('selector-paste').click();
    await expect(panel(page).getByTestId('selector-paste-input')).toBeFocused();
    await panel(page).getByTestId('selector-paste-input').fill('not an export');
    await panel(page).getByTestId('selector-paste-use').click();
    await expect(panel(page).getByTestId('selector-paste-error')).toHaveText(
      'That is not a Forever Sixty export.',
    );
    await expect(panel(page).getByTestId('selector-paste-input')).toHaveValue('not an export');
    await panel(page).getByRole('button', { name: 'Cancel' }).click();
    await expect(panel(page).getByTestId('selector-paste')).toBeVisible();
  });

  test('a valid export becomes the current character as a Pasted export row', async ({ page }) => {
    await signInWith(page, [OBNOXIOUS]);
    await storePointer(page, armoryPointer(OBNOXIOUS));
    await page.goto('/guides');
    await ready(page);
    await selector(page).click();
    await panel(page).getByTestId('selector-paste').click();
    await panel(page).getByTestId('selector-paste-input').fill(FURY);
    await panel(page).getByTestId('selector-paste-use').click();
    await expect(panel(page)).toBeHidden();
    await expect(selector(page)).toBeFocused();
    const pointer = await page.evaluate(() =>
      JSON.parse(window.localStorage.getItem('fs.currentCharacter') ?? 'null'),
    );
    expect(pointer).toMatchObject({ source: 'code', ref: FURY, classSlug: 'warrior' });
    await selector(page).click();
    const current = panel(page).locator('[aria-current="true"]');
    await expect(current).toContainText('Pasted export');
  });
});

test.describe('signed out and expired', () => {
  test.skip(() => test.info().project.name !== 'desktop', 'desktop layout only');
  test.use({ viewport: { width: 1440, height: 900 } });

  test('shows Sign in, and opens Add your character with paste, the addon and a quiet sign-in', async ({
    page,
  }) => {
    await signedOut(page);
    await page.goto('/guides');
    await expect(selector(page)).toHaveAttribute('aria-label', 'Sign in or choose a character');
    await expect(selector(page)).toContainText('Sign in');
    await expect(selector(page)).toContainText('or paste an export');
    await selector(page).click();
    await expect(panel(page).getByText('Add your character')).toBeVisible();
    await expect(panel(page).getByTestId('selector-manage')).toHaveCount(0);
    await expect(panel(page).getByTestId('selector-paste')).toContainText('Paste an export');
    await expect(panel(page).getByTestId('selector-addon')).toContainText('Get the addon');
    await expect(panel(page).getByTestId('selector-sign-in')).toHaveAttribute('href', '/login');
  });

  test('a pasted export with no session is the current row, with the not-saved note and a sign-in', async ({
    page,
  }) => {
    await signedOut(page);
    await storePointer(page, {
      source: 'code',
      ref: FURY,
      label: 'Obnoxious Yell · Fury Warrior',
      classSlug: 'warrior',
    });
    await page.goto('/guides');
    await ready(page);
    await expect(page.getByTestId('selector-name')).toHaveText('Obnoxious Yell');
    await selector(page).click();
    await expect(panel(page)).toContainText('Pasted export');
    await expect(panel(page)).toContainText('Not saved to an account. Sign in to keep it.');
    await expect(panel(page).getByTestId('selector-sign-in')).toBeVisible();
  });

  test('an account pointer with no session shows the ember mark and the session-ended line', async ({
    page,
  }) => {
    await signedOut(page);
    await storePointer(page, armoryPointer(OBNOXIOUS));
    await page.goto('/guides');
    await ready(page);
    await expect(page.getByTestId('selector-mark')).toHaveAttribute('data-mark', 'session');
    await expect(selector(page)).toHaveAttribute('aria-label', /Signed out\. Change character\.$/);
    await selector(page).click();
    await expect(panel(page)).toContainText('Your session ended. Sign in to see your other characters.');
    await expect(panel(page).getByTestId('selector-sign-in')).toBeVisible();
  });

  test('a list that did not load says so and offers Try again and Paste', async ({ page }) => {
    await page.context().addCookies([{ name: 'fs_csrf', value: 'token', domain: 'localhost', path: '/' }]);
    await page.route('**/v1/me', (route) => route.abort());
    await page.goto('/guides');
    await expect(selector(page)).toBeVisible();
    await selector(page).click();
    await expect(panel(page).getByTestId('selector-load-error')).toContainText(
      'Your characters did not load.',
    );
    await expect(panel(page).getByRole('button', { name: 'Try again' })).toBeVisible();
    await expect(panel(page).getByTestId('selector-paste')).toBeVisible();
  });

  test('a signed-in account with no characters shows the none state with both setup rows', async ({
    page,
  }) => {
    await signInWith(page, []);
    await page.goto('/guides');
    await expect(selector(page)).toHaveAttribute('aria-label', 'Sign in or choose a character');
    await selector(page).click();
    await expect(panel(page).getByTestId('selector-none')).toContainText('No characters yet.');
    await expect(panel(page).getByTestId('selector-paste')).toBeVisible();
    await expect(panel(page).getByTestId('selector-addon')).toBeVisible();
  });
});

test.describe('before the island hydrates', () => {
  test.skip(() => test.info().project.name !== 'desktop', 'desktop layout only');
  test.use({ viewport: { width: 1440, height: 900 } });

  test('the crest, the name and the spec line are painted from the pointer with the bar at full height', async ({
    page,
  }) => {
    await signInWith(page, [OBNOXIOUS]);
    await storePointer(page, armoryPointer(OBNOXIOUS));
    await page.route('**/_astro/**/*.js', (route) => route.abort());
    await page.route('**/_astro/*.js', (route) => route.abort());
    await page.goto('/guides');
    const link = page.getByTestId('character-selector-pre');
    await expect(link).toBeVisible();
    await expect(link.locator('.csel-pre-char .csel-name')).toHaveText('Obnoxious Yell');
    await expect(link.locator('.csel-pre-char .csel-sub')).toHaveText('Fury Warrior');
    await expect(link.locator('.csel-pre-crest')).toHaveCSS('background-image', /crests\/warrior\.webp/);
    expect((await barBox(page))!.height).toBeGreaterThanOrEqual(81);
    expect((await link.boundingBox())!.width).toBe(232);
  });
});

test.describe('the phone bar', () => {
  test.skip(() => test.info().project.name !== 'mobile', 'phone layout only');

  test('is one 56px bar with a 44px crest and a 44px Menu button, no sideways scroll', async ({ page }) => {
    await signInWith(page, [OBNOXIOUS]);
    await storePointer(page, armoryPointer(OBNOXIOUS));
    await page.goto('/guides');
    await ready(page);
    expect((await barBox(page))!.height).toBeLessThanOrEqual(57);
    const crest = (await selector(page).boundingBox())!;
    const menu = (await page.getByTestId('menu-button').boundingBox())!;
    expect(crest.width).toBeGreaterThanOrEqual(44);
    expect(crest.height).toBeGreaterThanOrEqual(44);
    expect(menu.width).toBe(44);
    expect(menu.height).toBe(44);
    expect(crest.x + crest.width).toBeLessThanOrEqual(menu.x);
    expect(await noSideScroll(page)).toBe(true);
    await expect(page.getByTestId('selector-name')).toBeHidden();
  });

  test('the Menu leads with the full name, then the doors with BiS, then Discord', async ({ page }) => {
    await signInWith(page, [OBNOXIOUS]);
    await storePointer(page, armoryPointer(OBNOXIOUS));
    await page.goto('/guides');
    await ready(page);
    await page.getByTestId('menu-button').click();
    const row = page.getByTestId('menu-character');
    await expect(row).toContainText('Obnoxious Yell');
    await expect(row).toContainText('Fury Warrior · 60');
    await expect(page.getByTestId('primary-nav').getByRole('link')).toHaveText(DOORS);
    const rowBox = (await row.boundingBox())!;
    const navBox = (await page.getByTestId('primary-nav').boundingBox())!;
    expect(rowBox.y + rowBox.height).toBeLessThanOrEqual(navBox.y);
    expect(rowBox.height).toBeGreaterThanOrEqual(56);
    expect(await noSideScroll(page)).toBe(true);
  });

  test('the sheet opens from the crest and from the Menu row, holds every row and closes four ways', async ({
    page,
  }) => {
    await signInWith(page, [OBNOXIOUS, FROSTBYTE, SHADOWMEND, OAKHEART, TREEWALKER]);
    await storePointer(page, armoryPointer(OBNOXIOUS));
    await page.goto('/guides');
    await ready(page);
    await selector(page).click();
    await expect(panel(page)).toHaveAttribute('aria-modal', 'true');
    await expect(panel(page).locator('.csel-row')).toHaveCount(5);
    const sheet = (await panel(page).boundingBox())!;
    const viewport = page.viewportSize()!;
    expect(Math.round(sheet.y + sheet.height)).toBe(viewport.height);
    expect(sheet.width).toBe(viewport.width);
    for (const target of [
      panel(page).getByRole('button', { name: 'Close' }),
      panel(page).getByTestId('selector-paste'),
      panel(page).locator('.csel-row-main').first(),
    ]) {
      const box = (await target.boundingBox())!;
      expect(box.height).toBeGreaterThanOrEqual(44);
    }
    await panel(page).getByRole('button', { name: 'Close' }).click();
    await expect(panel(page)).toBeHidden();

    await selector(page).click();
    await page.keyboard.press('Escape');
    await expect(panel(page)).toBeHidden();

    await selector(page).click();
    await page.mouse.click(200, 60);
    await expect(panel(page)).toBeHidden();

    await page.getByTestId('menu-button').click();
    await page.getByTestId('menu-character').click();
    await expect(panel(page)).toBeVisible();
    await expect(page.getByTestId('primary-nav')).toBeHidden();
    await panel(page)
      .getByTestId(`selector-row-${keyOf(FROSTBYTE)}`)
      .getByRole('button')
      .first()
      .click();
    await expect(panel(page)).toBeHidden();
    await expect(page.getByTestId('selector-live')).toHaveText('Now using Frostbyte.');
  });

  test('signed out, the Menu leads with Sign in and the sheet says Add your character', async ({ page }) => {
    await signedOut(page);
    await page.goto('/guides');
    await ready(page);
    await page.getByTestId('menu-button').click();
    await expect(page.getByTestId('menu-character')).toContainText('Sign in');
    await expect(page.getByTestId('menu-character')).toContainText('or paste an export');
    await page.getByTestId('menu-character').click();
    await expect(panel(page).getByText('Add your character')).toBeVisible();
    await expect(panel(page).getByTestId('selector-sign-in')).toBeVisible();
  });
});

test.describe('the tablet bar', () => {
  test.skip(() => test.info().project.name !== 'desktop', 'desktop layout only');
  test.use({ viewport: { width: 768, height: 900 } });

  test('shows the full name beside the crest and keeps the doors in the Menu', async ({ page }) => {
    await signInWith(page, [OBNOXIOUS]);
    await storePointer(page, armoryPointer(OBNOXIOUS));
    await page.goto('/guides');
    await ready(page);
    await expect(page.locator('.csel-name-bar')).toBeVisible();
    await expect(page.getByTestId('menu-button')).toBeVisible();
    await expect(page.getByTestId('primary-nav')).toBeHidden();
  });
});
