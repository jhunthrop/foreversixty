// web/tests/e2e/selector-restored.spec.ts
// 2026-10-31: the "Jump to" spine bar is gone from every page, and its two survivors live in the
// header selector's panel: "Restored your last character" and Forget, shown only while the
// current character came back from an earlier visit (lib/character-selector/restored.ts).
import type { Page } from '@playwright/test';
import { apiUnreachable, expect, test } from './support/test';
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

// Reads this spec does not answer go unanswered, as the API out of reach (support/test.ts).
test.beforeEach(({ page }) => apiUnreachable(page, '**/v1/characters/*/*/*', '**/v1/encounters*'));

const FURY = `FS1:${ACTIVE_BUILD}:warrior:orc:0/5530515/0:head=12640,main_hand=11726`;

/** The account's main: the freshest build wins (main-character.ts), and Obnoxious Yell's is minutes old. */
const MAIN = OBNOXIOUS;
const ALT: FixtureCharacter = {
  name: 'Frostbyte',
  class: 'Mage',
  spec: 'Frost',
  level: 42,
  buildDaysAgo: 3,
};

const selector = (page: Page) => page.getByTestId('character-selector');
const panel = (page: Page) => page.getByTestId('selector-panel');
const restoredLine = (page: Page) => panel(page).getByTestId('selector-restored');

async function openSelector(page: Page): Promise<void> {
  await expect(selector(page)).toBeVisible();
  await selector(page).click();
  await expect(panel(page)).toBeVisible();
}

async function storedPointer(page: Page): Promise<string | null> {
  return page.evaluate(() => window.localStorage.getItem('fs.currentCharacter'));
}

test.describe('the retired Jump to bar', () => {
  for (const path of ['/sim', '/sim/gear', '/planner', '/character/us/hardcore/elyra-duskvale']) {
    test(`${path} carries no Jump to bar, even with a stored character`, async ({ page }) => {
      await signInWith(page, [MAIN, ALT]);
      await storePointer(page, armoryPointer(ALT));
      await page.goto(path);
      await expect(selector(page)).toBeVisible();
      await expect(page.getByTestId('current-character-bar')).toHaveCount(0);
      await expect(page.getByText('Jump to', { exact: true })).toHaveCount(0);
    });
  }
});

test.describe('the selector’s restored line and Forget', () => {
  test('a character restored from an earlier visit shows the line; Forget falls back to the main', async ({
    page,
  }) => {
    await signInWith(page, [MAIN, ALT]);
    await storePointer(page, armoryPointer(ALT));
    await page.goto('/sim');
    await expect(page.getByTestId('selector-name')).toHaveText(ALT.name);

    await openSelector(page);
    await expect(restoredLine(page)).toContainText('Restored your last character');
    const forget = restoredLine(page).getByRole('button', { name: `Forget ${ALT.name}` });
    await expect(forget).toBeVisible();
    await forget.click();

    await expect(panel(page)).toHaveCount(0);
    await expect(page.getByTestId('selector-name')).toHaveText(MAIN.name);
    await expect(page.getByTestId('selector-live')).toHaveText(`Forgot ${ALT.name}.`);
    expect(await storedPointer(page)).toBeNull();
    await expect(selector(page)).toBeFocused();

    await openSelector(page);
    await expect(restoredLine(page)).toHaveCount(0);
  });

  test('Forget is reachable and works from the keyboard', async ({ page }) => {
    await signInWith(page, [MAIN, ALT]);
    await storePointer(page, armoryPointer(ALT));
    await page.goto('/sim/gear');
    await expect(page.getByTestId('selector-name')).toHaveText(ALT.name);
    await selector(page).focus();
    await page.keyboard.press('ArrowDown');
    await expect(panel(page)).toBeVisible();

    const forget = page.getByTestId('selector-forget');
    for (let step = 0; step < 12; step += 1) {
      if (await forget.evaluate((element) => element === document.activeElement)) break;
      await page.keyboard.press('Tab');
    }
    await expect(forget).toBeFocused();
    await page.keyboard.press('Enter');
    await expect(page.getByTestId('selector-name')).toHaveText(MAIN.name);
    expect(await storedPointer(page)).toBeNull();
  });

  test('a character chosen in this session shows no restored line', async ({ page }) => {
    await signInWith(page, [MAIN, ALT]);
    await storePointer(page, armoryPointer(MAIN));
    await page.goto('/rankings');
    await openSelector(page);
    await panel(page)
      .getByTestId(`selector-row-${keyOf(ALT)}`)
      .getByRole('button')
      .first()
      .click();
    await expect(page.getByTestId('selector-name')).toHaveText(ALT.name);

    await openSelector(page);
    await expect(restoredLine(page)).toHaveCount(0);
    // And on the next page of the same visit.
    await page.goto('/sim');
    await openSelector(page);
    await expect(restoredLine(page)).toHaveCount(0);
  });

  test('a restored main shows no line: forgetting it would change nothing', async ({ page }) => {
    await signInWith(page, [MAIN, ALT]);
    await storePointer(page, armoryPointer(MAIN));
    await page.goto('/sim');
    await expect(page.getByTestId('selector-name')).toHaveText(MAIN.name);
    await openSelector(page);
    await expect(panel(page).getByTestId('selector-list')).toBeVisible();
    await expect(restoredLine(page)).toHaveCount(0);
  });

  test('signed out, a restored pasted export can be forgotten, leaving Sign in', async ({ page }) => {
    await signedOut(page);
    await storePointer(page, {
      source: 'code',
      ref: FURY,
      label: 'Obnoxious Yell · Fury Warrior',
      classSlug: 'warrior',
    });
    await page.goto('/guides');
    await expect(page.getByTestId('selector-name')).toHaveText('Obnoxious Yell');
    await openSelector(page);
    await expect(restoredLine(page)).toContainText('Restored your last character');
    await restoredLine(page).getByRole('button', { name: 'Forget Obnoxious Yell' }).click();
    expect(await storedPointer(page)).toBeNull();
    await expect(selector(page)).toHaveAttribute('aria-label', 'Sign in or choose a character');
  });
});
