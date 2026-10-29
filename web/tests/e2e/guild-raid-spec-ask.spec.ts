// web/tests/e2e/guild-raid-spec-ask.spec.ts
// The raid page asks the spec (design/specs/2026-10-28-discord-guided-cards.md P.10.7): In, Late
// and Bench wait for the pick while no raid spec is saved, the pick commits the sign-up with the
// spec's role, Out never asks, a saved spec reads `Change spec`, and the select holds the 44 px
// floor at 360. The API is the RaidStub with `askSpec`.
import { expect, test, type Locator, type Page } from '@playwright/test';
import { VIEWER_MEMBER } from '../../src/fixtures/guild/mock-guild';
import { raidCopy } from '../../src/lib/guild/raid-copy';
import {
  CHICAGO,
  NOW,
  ONYXIA_ID,
  raidUrl,
  stripId,
  stubGuildRaids,
  writesTo,
  type RaidStub,
} from './support/guild-raids';
import { gotoHydrated } from './support/hydrated';
import { shortControls } from './support/short-controls';

test.use({ timezoneId: CHICAGO });

const SIGNUP_PUTS = /signups\/us\/pvp\/zulmara$/;
const PHONE = { width: 360, height: 800 };

async function openAsking(page: Page): Promise<RaidStub> {
  const stub = await stubGuildRaids(page, { viewer: 'member', askSpec: true });
  await page.clock.install({ time: NOW });
  await gotoHydrated(page, raidUrl(ONYXIA_ID), 'raid-page');
  return stub;
}

const specSelect = (page: Page): Locator => page.getByTestId(stripId('spec', VIEWER_MEMBER));
const toggle = (page: Page): Locator => page.getByTestId(`${stripId('spec', VIEWER_MEMBER)}-toggle`);
const specLabel = (page: Page): Locator => page.getByTestId(`${stripId('spec', VIEWER_MEMBER)}-label`);
const status = (page: Page): Locator => page.getByTestId(stripId('status', VIEWER_MEMBER));
const segment = (page: Page, state: string): Locator =>
  page.getByTestId(`${stripId('segments', VIEWER_MEMBER)}-${state}`);
const bodies = (stub: RaidStub): unknown[] => writesTo(stub, 'PUT', SIGNUP_PUTS).map((w) => w.body);

test('In waits for the pick; the pick signs up with that spec and says it is kept', async ({ page }) => {
  const stub = await openAsking(page);
  await expect(specSelect(page)).toHaveCount(0);

  await segment(page, 'in').click();
  await expect(specSelect(page)).toBeVisible();
  await expect(specLabel(page)).toHaveText(raidCopy.spec.title);
  await expect(segment(page, 'in')).toHaveAttribute('aria-checked', 'false');
  await expect(segment(page, 'in')).toHaveAttribute('aria-busy', 'true');
  expect(bodies(stub)).toEqual([]);
  await page.reload();
  await expect(page.getByTestId('raid-page')).toBeVisible();
  await expect(segment(page, 'in')).toHaveAttribute('aria-checked', 'false');
  await expect(specSelect(page)).toHaveCount(0);
  expect(bodies(stub)).toEqual([]);
  await segment(page, 'in').click();

  await specSelect(page).selectOption('hunter-survival');

  await expect(status(page)).toContainText(raidCopy.spec.kept('Survival', VIEWER_MEMBER));
  expect(bodies(stub)).toEqual([{ state: 'in', role: 'dps', note: '', spec: 'hunter-survival' }]);
  await expect(toggle(page)).toHaveText(`Survival · ${raidCopy.spec.change}`);
  await expect(toggle(page)).toHaveAttribute('aria-expanded', 'false');
  await expect(specSelect(page)).toBeHidden();
  await toggle(page).click();
  await expect(specSelect(page)).toHaveValue('hunter-survival');
});

test('Escape drops a press that is waiting for its pick', async ({ page }) => {
  const stub = await openAsking(page);
  await segment(page, 'in').click();
  await expect(segment(page, 'in')).toHaveAttribute('aria-busy', 'true');
  await page.keyboard.press('Escape');
  await expect(segment(page, 'in')).not.toHaveAttribute('aria-busy', 'true');
  await expect(specSelect(page)).toHaveCount(0);
  expect(bodies(stub)).toEqual([]);
});

test('an idle press drops after 60 s out of focus', async ({ page }) => {
  await openAsking(page);
  await segment(page, 'in').click();
  await page.evaluate(() => (document.activeElement as HTMLElement | null)?.blur());
  await page.clock.fastForward(61_000);
  await expect(segment(page, 'in')).not.toHaveAttribute('aria-busy', 'true');
});

test('pressing a state shifts nothing below the row', async ({ page }) => {
  for (const size of [PHONE, { width: 1440, height: 900 }]) {
    await page.setViewportSize(size);
    await openAsking(page);
    const row = page.getByTestId(stripId('row', VIEWER_MEMBER));
    const before = (await row.boundingBox())?.height ?? 0;
    await segment(page, 'in').click();
    await expect(specSelect(page)).toBeVisible();
    const after = (await row.boundingBox())?.height ?? 0;
    expect(Math.abs(after - before)).toBeLessThanOrEqual(2);
  }
});

test('the next press asks nothing: the saved spec stands', async ({ page }) => {
  const stub = await openAsking(page);
  await segment(page, 'in').click();
  await specSelect(page).selectOption('hunter-marksmanship');
  await expect(status(page)).toContainText(raidCopy.stripStatus.saved);

  await segment(page, 'late').click();

  await expect(segment(page, 'late')).toHaveAttribute('aria-checked', 'true');
  await expect.poll(() => bodies(stub).length).toBe(2);
  expect(bodies(stub)[1]).toMatchObject({ state: 'late' });
  expect(bodies(stub)[1]).not.toHaveProperty('spec');
});

test('Change spec saves another spec and keeps the state', async ({ page }) => {
  const stub = await openAsking(page);
  await segment(page, 'bench').click();
  await specSelect(page).selectOption('hunter-beast-mastery');
  await expect(toggle(page)).toBeVisible();
  await toggle(page).click();

  await specSelect(page).focus();
  await specSelect(page).selectOption('hunter-survival');

  await expect(toggle(page)).toHaveText(`Survival · ${raidCopy.spec.change}`);
  await expect(toggle(page)).toBeFocused();
  await expect(segment(page, 'bench')).toHaveAttribute('aria-checked', 'true');
  await expect.poll(() => bodies(stub).length).toBe(2);
  expect(bodies(stub)[1]).toMatchObject({ state: 'bench', spec: 'hunter-survival' });
});

test('Out never asks and saves at once', async ({ page }) => {
  const stub = await openAsking(page);
  await segment(page, 'out').click();
  await expect(status(page)).toHaveText(raidCopy.stripStatus.saved);
  await expect(specSelect(page)).toHaveCount(0);
  expect(bodies(stub)).toEqual([{ state: 'out', role: 'dps', note: '' }]);
});

test('a refused pick says so and makes no sign-up', async ({ page }) => {
  const stub = await openAsking(page);
  stub.refuseNext = { signup: { code: 'server_error', status: 500 } };
  await segment(page, 'in').click();
  await specSelect(page).selectOption('hunter-survival');
  await expect(status(page)).toHaveText(raidCopy.stripStatus.failed);
  await expect(segment(page, 'in')).toHaveAttribute('aria-checked', 'false');
});

test('the select holds the 44 px floor at 360 and the page does not scroll sideways', async ({ page }) => {
  await page.setViewportSize(PHONE);
  await openAsking(page);
  await segment(page, 'in').click();
  await expect(specSelect(page)).toBeVisible();
  expect(await shortControls(page.getByTestId(stripId('row', VIEWER_MEMBER)))).toEqual([]);
  expect(await page.evaluate(() => document.documentElement.scrollWidth <= window.innerWidth)).toBe(true);
});
