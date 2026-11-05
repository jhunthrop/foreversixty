// web/tests/e2e/guild-events.spec.ts
// Dungeon runs and guild events (design/specs/2026-11-05-guild-events.md §10): the form's What and
// Group size selects, the title and open-day errors, a keyed dungeon's note, the planned row and
// the event page's words by kind. The API is stubbed by support/guild-raids.ts; FS_EVENT_SHOTS, a
// directory, also saves the §10 acceptance captures there.
import type { Page } from '@playwright/test';
import { expect, test } from './support/test';
import { gotoHydrated } from './support/hydrated';
import { choose } from './support/select';
import {
  CHICAGO,
  FIRST_NEW_ID,
  GUILD_URL,
  NOW,
  raidUrl,
  storedRaid,
  stubGuildRaids,
  writesTo,
} from './support/guild-raids';

test.use({ timezoneId: CHICAGO });

const NOV_20 = '2026-11-20';
const WEEK_ONE = new Date('2026-11-05T18:00:00Z');
const DEADMINES_ID = 21;
const EVENT_ID = 22;

async function shot(page: Page, name: string): Promise<void> {
  const dir = process.env.FS_EVENT_SHOTS;
  if (dir === undefined) return;
  const width = page.viewportSize()?.width ?? 0;
  await page.screenshot({ path: `${dir}/${name}-${width}.png`, fullPage: true });
}

const deadmines = storedRaid({
  id: DEADMINES_ID,
  zone: 'dungeon:the-deadmines',
  zone_name: 'The Deadmines',
  zone_kind: 'dungeon',
  title: 'The Deadmines',
  starts_at: '2026-11-07T02:00:00Z',
  size: 5,
  notes: '',
  notes_html: '',
  seeds: [
    { name: 'Kraggor', state: 'in' },
    { name: 'Duskstrider', state: 'in' },
  ],
});
const premade = storedRaid({
  id: EVENT_ID,
  zone: 'event:guild',
  zone_name: 'Guild event',
  zone_kind: 'event',
  title: 'Darkspear premade',
  starts_at: '2026-11-08T02:00:00Z',
  size: null,
  notes: '',
  notes_html: '',
  seeds: [],
});

async function openTab(page: Page, raids = [deadmines, premade], at = WEEK_ONE) {
  const stub = await stubGuildRaids(page, { viewer: 'officer', raids });
  await page.clock.install({ time: at });
  await gotoHydrated(page, `${GUILD_URL}#raids`, 'guild-raids-tab');
  return stub;
}

async function openForm(page: Page) {
  await page.getByTestId('guild-raids-plan').click();
  await expect(page.getByTestId('raid-form-create')).toBeVisible();
}

test('the tab holds a dungeon run and a guild event, and offers Plan an event', async ({ page }) => {
  await openTab(page);
  await expect(page.getByTestId('guild-raids-plan')).toHaveText('Plan an event');
  await expect(page.getByTestId(`planned-raid-${DEADMINES_ID}`)).toContainText('The Deadmines');
  await expect(page.getByTestId(`planned-raid-${EVENT_ID}`)).toContainText('Darkspear premade');
  await expect(page.getByTestId(`planned-raid-${EVENT_ID}`)).toContainText('Guild event');
  await shot(page, 'tab');
});

test('a guild event asks for a size and a title, and posts its sized key', async ({ page }) => {
  const stub = await openTab(page, [deadmines]);
  await openForm(page);
  await expect(page.getByTestId('raid-form-group-size')).toHaveCount(0);
  await choose(page.getByTestId('raid-form-zone'), 'event:guild');
  await expect(page.getByTestId('raid-form-group-size')).toBeAttached();
  await choose(page.getByTestId('raid-form-group-size'), '10');
  await page.getByTestId('raid-form-date').fill('2026-11-12');
  await page.getByTestId('raid-form-submit').click();
  await expect(page.getByTestId('raid-form-error-title')).toHaveText('Give the event a title.');
  await shot(page, 'form-event-title-error');

  await page.getByTestId('raid-form-title').fill('Quest group');
  await page.getByTestId('raid-form-submit').click();
  await expect(page.getByTestId('raid-form-create')).toHaveCount(0);
  const [post] = writesTo(stub, 'POST', /\/raids$/);
  expect(post?.body).toMatchObject({ zone: 'event:guild-10', title: 'Quest group' });
  await expect(page.getByTestId(`planned-raid-${FIRST_NEW_ID}`)).toContainText('Guild event · 10');
});

test('the form opens on the last planned key, an event with its size', async ({ page }) => {
  const sized = { ...premade, zone: 'event:guild-10', size: 10 };
  await openTab(page, [deadmines, sized]);
  await openForm(page);
  await expect(page.getByTestId('raid-form-zone')).toHaveValue('event:guild');
  await expect(page.getByTestId('raid-form-group-size')).toHaveValue('10');
});

test('a keyed dungeon names its key, and the size select stays away', async ({ page }) => {
  await openTab(page);
  await openForm(page);
  await choose(page.getByTestId('raid-form-zone'), 'dungeon:upper-blackrock-spire');
  await expect(page.getByTestId('raid-form-key')).toContainText('Needs a key: Seal of Ascension');
  await expect(page.getByTestId('raid-form-group-size')).toHaveCount(0);
  await shot(page, 'form-keyed');
});

test('a raid before it opens sets the date floor and refuses the date', async ({ page }) => {
  const stub = await openTab(page);
  await openForm(page);
  await choose(page.getByTestId('raid-form-zone'), 'raid:onyxias-lair');
  await expect(page.getByTestId('raid-form-opens')).toHaveText("Onyxia's Lair opens Wed 9 Dec.");
  await expect(page.getByTestId('raid-form-date')).toHaveAttribute('min', '2026-12-09');
  await page.getByTestId('raid-form-date').fill(NOV_20);
  await page.getByTestId('raid-form-submit').click();
  await expect(page.getByTestId('raid-form-error-date')).toHaveText(
    "Onyxia's Lair opens Wed 9 Dec. Pick that day or later.",
  );
  expect(writesTo(stub, 'POST', /\/raids$/)).toHaveLength(0);
  await shot(page, 'form-raid-before-open');
});

test('the run page reads DUNGEON and Cancel run, its count against five', async ({ page }) => {
  await stubGuildRaids(page, { viewer: 'officer', raids: [deadmines] });
  await page.clock.install({ time: WEEK_ONE });
  await gotoHydrated(page, raidUrl(DEADMINES_ID), 'raid-page');
  await expect(page.getByTestId('raid-page').locator('.label').first()).toHaveText(/dungeon/i);
  await expect(page.getByTestId('raid-officer-cancel')).toHaveText('Cancel run');
  await expect(page.getByTestId('raid-count')).toContainText('2 of 5');
  await shot(page, 'run-page');
  await page.getByTestId('raid-officer-cancel').click();
  await expect(page.getByTestId('raid-cancel-confirm')).toContainText(
    'Cancel this run? Members see it as cancelled.',
  );
  await expect(page.getByTestId('raid-cancel-confirm-button')).toHaveText('Cancel the run');
});

test('with nothing planned a member reads the new empty line', async ({ page }) => {
  await stubGuildRaids(page, { viewer: 'member', raids: [] });
  await page.clock.install({ time: NOW });
  await gotoHydrated(page, `${GUILD_URL}#raids`, 'guild-raids-tab');
  await expect(page.getByTestId('guild-planned-empty')).toHaveText(
    'Nothing planned yet. Officers post dungeon runs, events and raids here.',
  );
  await shot(page, 'member-empty');
});
