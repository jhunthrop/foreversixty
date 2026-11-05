// web/tests/e2e/guild-raid-spec-ask.spec.ts
// The strip's one spec picker (design/specs/2026-11-05-guild-raid-comp.md 3.2): with no spec on
// file the picker reads `Choose a spec`; In, Late and Bench then wait (`Pick a spec first.`, the
// picker red and focused) and the pick signs up with state and spec in one request; a pick on a
// signed-up row saves at once; Out never needs a spec; the row holds the 44 px floor at 360. The
// API is the RaidStub with `askSpec` (no raid spec and no default spec on file).
import type { Locator, Page } from '@playwright/test';
import { expect, test } from './support/test';
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
import { choose, fieldControl } from './support/select';

test.use({ timezoneId: CHICAGO });

const SIGNUP_PUTS = /signups\/us\/pvp\/zulmara$/;
const PHONE = { width: 360, height: 800 };
/** Raid comp spec 3.2: a phone strip row is at most 180 px. */
const PHONE_ROW_MAX = 180;

async function openAsking(page: Page): Promise<RaidStub> {
  const stub = await stubGuildRaids(page, { viewer: 'member', askSpec: true });
  await page.clock.install({ time: NOW });
  await gotoHydrated(page, raidUrl(ONYXIA_ID), 'raid-page');
  return stub;
}

const specSelect = (page: Page): Locator => page.getByTestId(stripId('spec', VIEWER_MEMBER));
// What the member sees of the field: the native select, or the listbox trigger that covers it
// once select-boot.ts enhances the field after idle on a fine pointer (either can be showing).
const specControl = (page: Page): Locator => fieldControl(page, stripId('spec', VIEWER_MEMBER));
const status = (page: Page): Locator => page.getByTestId(stripId('status', VIEWER_MEMBER));
const who = (page: Page): Locator => page.getByTestId(stripId('who', VIEWER_MEMBER));
const segment = (page: Page, state: string): Locator =>
  page.getByTestId(`${stripId('segments', VIEWER_MEMBER)}-${state}`);
const bodies = (stub: RaidStub): unknown[] => writesTo(stub, 'PUT', SIGNUP_PUTS).map((w) => w.body);

test('In with no spec waits: Pick a spec first., then the pick signs up in one request', async ({ page }) => {
  const stub = await openAsking(page);
  await expect(specControl(page)).toBeVisible();
  await expect(specControl(page)).toContainText(raidCopy.spec.placeholder);

  await segment(page, 'in').click();
  await expect(status(page)).toHaveText(raidCopy.spec.pickFirst);
  await expect(status(page).getByRole('alert')).toHaveText(raidCopy.spec.pickFirst);
  await expect(specSelect(page)).toHaveAttribute('aria-invalid', 'true');
  await expect(segment(page, 'in')).toHaveAttribute('aria-checked', 'false');
  expect(bodies(stub)).toEqual([]);

  await choose(specSelect(page), 'hunter-survival');

  await expect(status(page)).toHaveText(raidCopy.stripStatus.saved);
  await expect(segment(page, 'in')).toHaveAttribute('aria-checked', 'true');
  expect(bodies(stub)).toEqual([{ state: 'in', role: 'dps', note: '', spec: 'hunter-survival' }]);
  await expect(who(page)).toHaveText('Survival Hunter');
  await expect(page.getByTestId('raid-roster-melee')).toContainText(VIEWER_MEMBER);
});

test('Escape drops a press that is waiting for its pick', async ({ page }) => {
  const stub = await openAsking(page);
  await segment(page, 'in').click();
  await expect(segment(page, 'in')).toHaveAttribute('aria-busy', 'true');
  await page.keyboard.press('Escape');
  await expect(segment(page, 'in')).not.toHaveAttribute('aria-busy', 'true');
  await expect(status(page)).toHaveText('');
  expect(bodies(stub)).toEqual([]);
});

test('an idle press drops after 60 s out of focus', async ({ page }) => {
  await openAsking(page);
  await segment(page, 'in').click();
  await page.evaluate(() => (document.activeElement as HTMLElement | null)?.blur());
  await page.clock.fastForward(61_000);
  await expect(segment(page, 'in')).not.toHaveAttribute('aria-busy', 'true');
});

test('a pick on an unsigned row saves nothing; the next press sends it', async ({ page }) => {
  const stub = await openAsking(page);
  await choose(specSelect(page), 'hunter-marksmanship');
  expect(bodies(stub)).toEqual([]);
  await segment(page, 'late').click();
  await expect(segment(page, 'late')).toHaveAttribute('aria-checked', 'true');
  await expect.poll(() => bodies(stub).length).toBe(1);
  expect(bodies(stub)[0]).toEqual({ state: 'late', role: 'dps', note: '', spec: 'hunter-marksmanship' });
});

test('a pick on a signed-up row saves at once and keeps the state', async ({ page }) => {
  const stub = await openAsking(page);
  await choose(specSelect(page), 'hunter-beast-mastery');
  await segment(page, 'bench').click();
  await expect.poll(() => bodies(stub).length).toBe(1);

  await choose(specSelect(page), 'hunter-survival');

  await expect(segment(page, 'bench')).toHaveAttribute('aria-checked', 'true');
  await expect.poll(() => bodies(stub).length).toBe(2);
  expect(bodies(stub)[1]).toMatchObject({ state: 'bench', spec: 'hunter-survival' });
  await expect(status(page)).toHaveText(raidCopy.stripStatus.saved);
});

test('pressing a state shifts nothing below the row', async ({ page }) => {
  for (const size of [PHONE, { width: 1440, height: 900 }]) {
    await page.setViewportSize(size);
    await openAsking(page);
    const row = page.getByTestId(stripId('row', VIEWER_MEMBER));
    const before = (await row.boundingBox())?.height ?? 0;
    await segment(page, 'in').click();
    await expect(status(page)).toHaveText(raidCopy.spec.pickFirst);
    const after = (await row.boundingBox())?.height ?? 0;
    expect(Math.abs(after - before)).toBeLessThanOrEqual(2);
  }
});

test('Out never needs a spec and saves at once', async ({ page }) => {
  const stub = await openAsking(page);
  await segment(page, 'out').click();
  await expect(status(page)).toHaveText(raidCopy.stripStatus.saved);
  expect(bodies(stub)).toEqual([{ state: 'out', role: 'dps', note: '' }]);
});

test('a refused pick says so and makes no sign-up', async ({ page }) => {
  const stub = await openAsking(page);
  stub.refuseNext = { signup: { code: 'server_error', status: 500 } };
  await segment(page, 'in').click();
  await choose(specSelect(page), 'hunter-survival');
  await expect(status(page)).toHaveText(raidCopy.stripStatus.failed);
  await expect(segment(page, 'in')).toHaveAttribute('aria-checked', 'false');
});

test('the row holds the 44 px floor at 360, at most 180 px, and the page does not scroll sideways', async ({
  page,
}) => {
  await page.setViewportSize(PHONE);
  await openAsking(page);
  const row = page.getByTestId(stripId('row', VIEWER_MEMBER));
  expect(await shortControls(row)).toEqual([]);
  expect((await row.boundingBox())?.height ?? 0).toBeLessThanOrEqual(PHONE_ROW_MAX);
  expect(await page.evaluate(() => document.documentElement.scrollWidth <= window.innerWidth)).toBe(true);
});
