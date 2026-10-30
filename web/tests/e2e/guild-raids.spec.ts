// web/tests/e2e/guild-raids.spec.ts
// Guild raids and sign-ups acceptance (design/specs/2026-10-26-guild-raids.md §9, the E2E
// paragraph, plus the browser behaviours Task 7 left to this file): a member signs up from the
// raid page and from the band line and R3 moves the other character to Bench; an officer posts a
// raid; the 10 s poll; the race lines; officer-read-only; R1 for visitors; the guild clock (R4);
// axe; and the 44 px floor at 360. The API is stubbed by support/guild-raids.ts.
import { expect, test, type Locator, type Page } from '@playwright/test';
import { copyList, raidCopy } from '../../src/lib/guild/raid-copy';
import { gotoHydrated } from './support/hydrated';
import { expectNoAxeViolations } from './support/officer-tools';
import { shortControls } from './support/short-controls';
import {
  CHICAGO,
  FIRST_NEW_ID,
  GUILD_URL,
  HYJAL_ID,
  MEMBER_ALT,
  NOW,
  ONYXIA_ID,
  keyOf,
  raidUrl,
  repeatStarts,
  stripId,
  storedRaid,
  stubGuildRaids,
  threePlanned,
  writesTo,
  type RaidStub,
  type RaidStubOptions,
} from './support/guild-raids';
import { VIEWER_MEMBER } from '../../src/fixtures/guild/mock-guild';

const NEW_YORK = 'America/New_York';
const POLL_MS = 10_000;
const COPIED_MS = 4_000;
const PHONE = { width: 360, height: 800 };

test.use({ timezoneId: CHICAGO });

async function openRaid(page: Page, options: RaidStubOptions, id = ONYXIA_ID, at = NOW): Promise<RaidStub> {
  const stub = await stubGuildRaids(page, options);
  await page.clock.install({ time: at });
  await gotoHydrated(page, raidUrl(id), 'raid-page');
  return stub;
}

async function openRaidsTab(page: Page, options: RaidStubOptions, at = NOW): Promise<RaidStub> {
  const stub = await stubGuildRaids(page, options);
  await page.clock.install({ time: at });
  await gotoHydrated(page, `${GUILD_URL}#raids`, 'guild-raids-tab');
  return stub;
}

const segment = (page: Page, name: string, state: string): Locator =>
  page.getByTestId(`${stripId('segments', name)}-${state}`);
const rosterSelects = (page: Page): Locator => page.locator('select[data-testid^="raid-roster-select-"]');

async function openForm(page: Page): Promise<Locator> {
  await page.getByTestId('guild-raids-plan').click();
  const form = page.getByTestId('raid-form-create');
  await expect(form).toBeVisible();
  return form;
}

test.describe('a member signs up (§9, R3)', () => {
  test('from the raid page: the alt moves to Bench and says so', async ({ page }) => {
    const stub = await openRaid(page, { viewer: 'member' });
    await expect(segment(page, VIEWER_MEMBER, 'in')).toHaveAttribute('aria-checked', 'false');
    await expect(segment(page, MEMBER_ALT, 'in')).toHaveAttribute('aria-checked', 'true');

    await segment(page, VIEWER_MEMBER, 'in').click();

    await expect(segment(page, VIEWER_MEMBER, 'in')).toHaveAttribute('aria-checked', 'true');
    await expect(segment(page, MEMBER_ALT, 'bench')).toHaveAttribute('aria-checked', 'true');
    await expect(page.getByTestId(stripId('status', MEMBER_ALT))).toHaveText(
      raidCopy.stripStatus.movedToBench(MEMBER_ALT),
    );
    await expect(page.getByTestId(stripId('status', VIEWER_MEMBER))).toHaveText(raidCopy.stripStatus.saved);
    const puts = writesTo(stub, 'PUT', /\/raids\/7\/signups\/us\/pvp\/zulmara$/);
    expect(puts.map((w) => w.body)).toEqual([{ state: 'in', role: 'dps', note: '' }]);
    await expect(page.getByTestId('raid-roster-bench')).toContainText(MEMBER_ALT);
    await expect(page.getByTestId('raid-roster-dps')).toContainText(VIEWER_MEMBER);
  });

  test('a note typed and then a segment pressed: one save carries both', async ({ page }) => {
    const stub = await openRaid(page, { viewer: 'member' });
    await page.getByTestId(stripId('note', VIEWER_MEMBER)).fill('late by 10');
    await segment(page, VIEWER_MEMBER, 'late').click();
    await expect(page.getByTestId(stripId('status', VIEWER_MEMBER))).toHaveText(raidCopy.stripStatus.saved);
    const puts = writesTo(stub, 'PUT', /signups\/us\/pvp\/zulmara$/);
    expect(puts.map((w) => w.body)).toEqual([{ state: 'late', role: 'dps', note: 'late by 10' }]);
  });

  test('a segment pressed while the note saves is sent when that save settles, never dropped', async ({
    page,
  }) => {
    const stub = await openRaid(page, { viewer: 'member' });
    await page.getByTestId(stripId('note', MEMBER_ALT)).fill('late by 10');
    await segment(page, MEMBER_ALT, 'late').click();
    await expect(segment(page, MEMBER_ALT, 'late')).toHaveAttribute('aria-checked', 'true');
    await expect(page.getByTestId(stripId('status', MEMBER_ALT))).toHaveText(raidCopy.stripStatus.saved);
    const puts = writesTo(stub, 'PUT', /signups\/us\/pvp\/duskstrider$/).map((w) => w.body);
    expect(puts.at(-1)).toEqual({ state: 'late', role: 'dps', note: 'late by 10' });
  });

  test('from the band line: the segments act on the character the sentence names', async ({ page }) => {
    const stub = await stubGuildRaids(page, { viewer: 'member' });
    await page.clock.install({ time: NOW });
    await gotoHydrated(page, GUILD_URL, 'guild-raid-line');
    const sentence = page.getByTestId('guild-raid-line-sentence');
    const segments = page.getByTestId('guild-raid-line-segments');
    // The alt is In, so the sentence names it and the segments show and act on it (spec 3.4).
    await expect(sentence).toHaveText(
      `Thursday 8 pm, you are In on ${MEMBER_ALT} as DPS, 14 of 40 signed up.`,
    );
    await expect(segments).toHaveAttribute('aria-label', `${MEMBER_ALT}'s sign-up`);
    await expect(page.getByTestId('guild-raid-line-segments-in')).toHaveAttribute('aria-checked', 'true');
    await page.getByTestId('guild-raid-line-segments-late').click();
    await expect(sentence).toHaveText(
      `Thursday 8 pm, you are Late on ${MEMBER_ALT} as DPS, 14 of 40 signed up.`,
    );
    await expect(page.getByTestId('guild-raid-line-segments-late')).toHaveAttribute('aria-checked', 'true');
    expect(writesTo(stub, 'PUT', /signups\/us\/pvp\/duskstrider$/).map((w) => w.body)).toMatchObject([
      { state: 'late', role: 'dps' },
    ]);
    expect(writesTo(stub, 'PUT', /signups\/us\/pvp\/zulmara$/)).toEqual([]);
    await expect(page.getByTestId('guild-raid-line-link')).toHaveAttribute('href', raidUrl(ONYXIA_ID));
  });
});

test('Plan a raid hides while the form is open, then returns with focus', async ({ page }) => {
  await openRaidsTab(page, { viewer: 'officer' });
  await openForm(page);
  await expect(page.getByTestId('guild-raids-plan')).toHaveCount(0);
  await page.getByTestId('raid-form-cancel').click();
  await expect(page.getByTestId('raid-form-create')).toHaveCount(0);
  await expect(page.getByTestId('guild-raids-plan')).toBeFocused();
});

test('an officer posts a raid with Post to Discord, and the planned row opens its page', async ({ page }) => {
  const stub = await openRaidsTab(page, { viewer: 'officer', discordConnected: true });
  await openForm(page);
  await page.getByTestId('raid-form-discord').check();
  await page.getByTestId('raid-form-submit').click();
  await expect(page.getByTestId('raid-form-create')).toHaveCount(0);

  const posts = writesTo(stub, 'POST', /\/raids$/);
  expect(posts).toHaveLength(1);
  expect(posts[0].body).toMatchObject({
    zone: 'raid:onyxias-lair',
    starts_at: '2026-12-18T02:00:00.000Z',
    timezone: CHICAGO,
    repeat_weeks: 0,
    post_to_discord: true,
  });

  const row = page.getByTestId(`planned-raid-${FIRST_NEW_ID}`);
  await expect(row).toBeVisible();
  await row.click();
  await page.waitForURL(`**${raidUrl(FIRST_NEW_ID)}`);
  await expect(page.getByTestId('raid-title')).toContainText("Onyxia's Lair");
});

test('counts change on the 10 s poll without a reload', async ({ page }) => {
  const stub = await openRaid(page, { viewer: 'member' });
  const count = page.getByTestId('raid-count');
  await expect(count).toContainText('14 of 40 signed up');
  await page.evaluate(() => ((window as unknown as { pollMarker: number }).pollMarker = 1));
  stub.applySignup(ONYXIA_ID, keyOf('Thornhide'), { state: 'in', role: 'dps', note: '' });
  stub.applySignup(ONYXIA_ID, keyOf('Grimvow'), { state: 'late', role: 'dps', note: '' });

  await page.clock.runFor(POLL_MS);

  await expect(count).toContainText('16 of 40 signed up');
  await expect(page.getByTestId('raid-roster-late')).toContainText('Grimvow');
  expect(await page.evaluate(() => (window as unknown as { pollMarker?: number }).pollMarker)).toBe(1);
});

test.describe('race lines take focus (§4 "Races")', () => {
  test('a lock while signing up reverts the press and locks the page', async ({ page }) => {
    const stub = await openRaid(page, { viewer: 'member' });
    stub.refuseNext = { signup: { code: 'raid_locked', status: 409 } };
    await segment(page, VIEWER_MEMBER, 'late').click();
    const line = page.getByTestId(stripId('status', VIEWER_MEMBER)).getByRole('alert');
    await expect(line).toHaveText(raidCopy.races.locked);
    await expect(line).toBeFocused();
    await expect(segment(page, VIEWER_MEMBER, 'late')).toHaveAttribute('aria-checked', 'false');
    await expect(page.getByTestId('raid-strip-state')).toHaveText(raidCopy.raidStates.locked);
    await expect(segment(page, VIEWER_MEMBER, 'in')).toBeDisabled();
    await expect(page.getByTestId('raid-status-chip')).toHaveText('Locked');
  });

  test('a cancel while signing up shows the page-level line, focused', async ({ page }) => {
    const stub = await openRaid(page, { viewer: 'member' });
    stub.refuseNext = { signup: { code: 'raid_cancelled', status: 409 } };
    await segment(page, VIEWER_MEMBER, 'in').click();
    const line = page.getByTestId('raid-race-line');
    await expect(line).toHaveText(raidCopy.races.cancelled);
    await expect(line).toBeFocused();
    await expect(page.getByTestId('raid-strip-state')).toHaveText(raidCopy.raidStates.cancelled);
  });
});

test.describe('officer-read-only (§3.2, gate spec §3.4)', () => {
  test('every officer control is present and disabled, and Copy still copies', async ({ browser }) => {
    const live = await browser.newPage();
    let selects = 0;
    try {
      await openRaid(live, { viewer: 'officer', officerTools: true });
      selects = await rosterSelects(live).count();
      expect(selects).toBeGreaterThan(0);
      expect(
        await rosterSelects(live).evaluateAll((n) => n.every((s) => !(s as HTMLSelectElement).disabled)),
      ).toBe(true);
    } finally {
      await live.close();
    }

    const context = await browser.newContext({ timezoneId: CHICAGO });
    try {
      await context.grantPermissions(['clipboard-read', 'clipboard-write']);
      const page = await context.newPage();
      const stub = await openRaid(page, { viewer: 'officer', officerTools: false });
      for (const id of ['raid-officer-edit', 'raid-officer-lock', 'raid-officer-cancel']) {
        await expect(page.getByTestId(id)).toBeDisabled();
        await expect(page.getByTestId(id)).toHaveAttribute('aria-describedby', 'officer-tools-gate-raids');
      }
      await expect(page.getByTestId('guild-raids-plan-gate')).toBeVisible();
      await expect(rosterSelects(page)).toHaveCount(selects);
      expect(
        await rosterSelects(page).evaluateAll((n) => n.every((s) => (s as HTMLSelectElement).disabled)),
      ).toBe(true);

      const copy = page.getByTestId('raid-officer-copy');
      await expect(copy).toBeEnabled();
      await copy.click();
      await expect(copy).toHaveText(raidCopy.officerRow.copied);
      const copied = await page.evaluate(() => navigator.clipboard.readText());
      expect(copied).toBe(copyList(stub.view(ONYXIA_ID), CHICAGO));
      expect(copied.split('\n')[0]).toBe("Onyxia's Lair · Thu 10 Dec, 8:00 pm CST · <t:1796954400:F>");
      expect(stub.writes).toEqual([]);
    } finally {
      await context.close();
    }
  });

  test('the Raids tab keeps Plan a raid, disabled under the gate sentence', async ({ page }) => {
    await openRaidsTab(page, { viewer: 'officer', officerTools: false });
    await expect(page.getByTestId('guild-raids-plan')).toBeDisabled();
    await expect(page.getByTestId('guild-raids-plan')).toHaveAttribute(
      'aria-describedby',
      'officer-tools-gate-raids',
    );
    await expect(page.getByTestId('guild-raids-plan-gate')).toBeVisible();
  });
});

test.describe('R1: visitors see nothing of planned raids', () => {
  for (const viewer of ['visitor', 'outsider', 'unverified'] as const) {
    test(`${viewer}: no Planned on the Raids tab, and the raid route shows the not-visible line`, async ({
      page,
    }) => {
      const stub = await stubGuildRaids(page, { viewer });
      await page.clock.install({ time: NOW });
      await gotoHydrated(page, `${GUILD_URL}#raids`, 'guild-raids-tab');
      await expect(page.getByTestId('guild-raids-rows')).toBeVisible();
      await expect(page.getByTestId('guild-planned')).toHaveCount(0);
      await expect(page.getByTestId('guild-raids-plan')).toHaveCount(0);
      await expect(page.locator('[data-testid^="guild-raid-line"]')).toHaveCount(0);

      await gotoHydrated(page, raidUrl(ONYXIA_ID), 'raid-not-visible');
      const expected =
        viewer === 'unverified'
          ? raidCopy.notVisible.unverified
          : raidCopy.notVisible.members('Olympus XXVII');
      await expect(page.getByTestId('raid-not-visible')).toContainText(expected);
      if (viewer === 'visitor') await expect(page.getByTestId('raid-sign-in-prompt')).toBeVisible();
      await expect(page.getByTestId('raid-page')).toHaveCount(0);
      expect(stub.raidReads).toEqual([]);
    });
  }
});

test.describe('the guild clock (R4)', () => {
  test.use({ timezoneId: NEW_YORK });

  test('the For you line reads the viewer clock with the guild on Chicago', async ({ page }) => {
    await openRaidsTab(page, { viewer: 'officer' });
    await openForm(page);
    await expect(page.getByTestId('raid-form-timezone')).toHaveValue(CHICAGO);
    await expect(page.getByTestId('raid-form-time')).toHaveValue('20:00');
    await expect(page.getByTestId('raid-form-for-you')).toHaveText('For you: Thu 17 Dec, 9:00 pm.');
    await page.getByTestId('raid-form-timezone').selectOption(NEW_YORK);
    await expect(page.getByTestId('raid-form-for-you')).toHaveText('For you: same time.');
  });

  test('a repeat across 14 March 2027 keeps 8:00 pm guild time', async ({ page }) => {
    const stub = await openRaidsTab(page, { viewer: 'officer', raids: [] }, new Date('2027-03-01T18:00:00Z'));
    await openForm(page);
    await page.getByTestId('raid-form-zone').selectOption('raid:onyxias-lair');
    await page.getByTestId('raid-form-date').fill('2027-03-11');
    await page.getByTestId('raid-form-time').fill('20:00');
    await page.getByTestId('raid-form-repeat').check();
    await page.getByTestId('raid-form-submit').click();
    await expect(page.getByTestId('raid-form-create')).toHaveCount(0);

    const [post] = writesTo(stub, 'POST', /\/raids$/);
    expect(post.body).toMatchObject({
      starts_at: '2027-03-12T02:00:00.000Z',
      timezone: CHICAGO,
      repeat_weeks: 3,
    });
    // The four raids the API makes from that body: one wall time, two offsets.
    const starts = repeatStarts('2027-03-12T02:00:00.000Z', CHICAGO, 3);
    expect(starts).toEqual([
      '2027-03-12T02:00:00.000Z',
      '2027-03-19T01:00:00.000Z',
      '2027-03-26T01:00:00.000Z',
      '2027-04-02T01:00:00.000Z',
    ]);
    const rows = page.locator('[data-testid^="planned-raid-"][href]');
    await expect(rows).toHaveCount(4);
    const labels = [
      '8:00 pm CST guild time',
      '8:00 pm CDT guild time',
      '8:00 pm CDT guild time',
      '8:00 pm CDT guild time',
    ];
    for (const [index, label] of labels.entries()) {
      await expect(rows.nth(index)).toContainText(label);
      await expect(rows.nth(index)).toContainText('9:00 pm');
    }
  });
});

test.describe('axe (§9)', () => {
  for (const viewer of ['member', 'officer', 'read-only'] as const) {
    test(`the raid page is clean for the ${viewer}`, async ({ page }) => {
      await openRaid(page, {
        viewer: viewer === 'member' ? 'member' : 'officer',
        officerTools: viewer !== 'read-only',
      });
      await expectNoAxeViolations(page, 'main');
    });
  }

  test('the Raids tab and the form are clean', async ({ page }) => {
    await openRaidsTab(page, { viewer: 'officer', raids: threePlanned() });
    await expectNoAxeViolations(page, '[data-testid="guild-raids-tab"]');
    await openForm(page);
    await expectNoAxeViolations(page, '[data-testid="raid-form-create"]');
  });
});

test.describe('every control is at least 44 px tall at 360', () => {
  test.beforeEach(({ page }, testInfo) => {
    test.skip(testInfo.project.name !== 'mobile', 'the 360 floor is a phone rule');
    return page.setViewportSize(PHONE);
  });

  for (const viewer of ['member', 'officer'] as const) {
    test(`the raid page for the ${viewer}`, async ({ page }) => {
      await openRaid(page, { viewer });
      expect(await shortControls(page.getByTestId('raid-page'))).toEqual([]);
    });
  }

  test('the Raids tab and the band line', async ({ page }) => {
    await openRaidsTab(page, { viewer: 'officer', raids: threePlanned() });
    expect(await shortControls(page.getByTestId('guild-raids-tab'))).toEqual([]);
    expect(await shortControls(page.getByTestId('guild-raid-line'))).toEqual([]);
  });

  test('the form with a Discord channel connected', async ({ page }) => {
    await openRaidsTab(page, { viewer: 'officer', discordConnected: true });
    expect(await shortControls(await openForm(page))).toEqual([]);
  });

  test('the form with no webhook: its inline Settings link', async ({ page }) => {
    // The `Connect a Discord channel in Settings…` line's inline link takes the gate link's
    // 44 px ::after hit area below md.
    await openRaidsTab(page, { viewer: 'officer' });
    expect(await shortControls(await openForm(page))).toEqual([]);
  });
});

test.describe('the officer row in a browser (Task 7)', () => {
  test('Cancel raid focuses Keep it, and Escape keeps the raid', async ({ page }) => {
    const stub = await openRaid(page, { viewer: 'officer' });
    await page.getByTestId('raid-officer-cancel').click();
    await expect(page.getByTestId('raid-cancel-keep')).toBeFocused();
    await page.keyboard.press('Escape');
    await expect(page.getByTestId('raid-cancel-confirm')).toHaveCount(0);
    await expect(page.getByTestId('raid-officer-cancel')).toBeFocused();
    expect(writesTo(stub, 'POST', /\/cancel$/)).toEqual([]);

    await page.getByTestId('raid-officer-cancel').click();
    await expect(page.getByTestId('raid-cancel-discord')).not.toBeChecked();
    await page.getByTestId('raid-cancel-confirm-button').click();
    await expect(page.getByTestId('raid-status-chip')).toHaveText('Cancelled');
    expect(writesTo(stub, 'POST', /\/raids\/7\/cancel$/).map((w) => w.body)).toEqual([
      { post_to_discord: false },
    ]);
    await expect(page.getByTestId('raid-back-link')).toBeFocused();
  });

  test('Copied shows for 4 s', async ({ page, context }) => {
    await context.grantPermissions(['clipboard-read', 'clipboard-write']);
    await openRaid(page, { viewer: 'officer' });
    const copy = page.getByTestId('raid-officer-copy');
    await copy.click();
    await expect(copy).toHaveText(raidCopy.officerRow.copied);
    await page.clock.runFor(COPIED_MS - 200);
    await expect(copy).toHaveText(raidCopy.officerRow.copied);
    await page.clock.runFor(400);
    await expect(copy).toHaveText(raidCopy.officerRow.copy);
  });

  test('a refused clipboard shows the line and the list, selected', async ({ page }) => {
    await page.addInitScript(() => {
      Object.defineProperty(navigator, 'clipboard', {
        configurable: true,
        value: { writeText: () => Promise.reject(new Error('denied')) },
      });
    });
    const stub = await openRaid(page, { viewer: 'officer' });
    await page.getByTestId('raid-officer-copy').click();
    await expect(page.getByTestId('raid-copy-failed')).toHaveText(raidCopy.copyFailed);
    const text = page.getByTestId('raid-copy-text');
    await expect(text).toHaveValue(copyList(stub.view(ONYXIA_ID), CHICAGO));
    const selection = await text.evaluate((node) => {
      const area = node as HTMLTextAreaElement;
      return [area.selectionStart, area.selectionEnd, area.value.length];
    });
    expect(selection).toEqual([0, selection[2], selection[2]]);
  });
});

test('a cancelled raid keeps its row on the Raids tab with the Cancelled chip', async ({ page }) => {
  await openRaidsTab(page, { viewer: 'member', raids: threePlanned() });
  await expect(page.getByTestId('guild-planned').getByTestId('card-head-meta')).toHaveText('3');
  const chip = page.locator(
    `[data-testid^="planned-raid-${HYJAL_ID}-"][data-testid$="-status-chip"]:visible`,
  );
  await expect(chip).toHaveText('Cancelled');
  // Spec 3.1 (review round): a cancelled row shows Cancelled alone, no viewer chip.
  await expect(
    page.locator(`[data-testid^="planned-raid-${HYJAL_ID}-"][data-testid$="-viewer-chip"]:visible`),
  ).toHaveCount(0);
  await expect(
    page.locator(`[data-testid^="planned-raid-${ONYXIA_ID}-"][data-testid$="-viewer-chip"]:visible`),
  ).toHaveText('In · DPS');
});

const BOT_CHANNEL = { discordConnected: true, postDefault: true, channelName: 'raid-signups' } as const;
const LATER_HELPER = raidCopy.form.postLater('#raid-signups', 'Thursday 10 December');
/** A raid posted through the bot to #raid-signups, before its reminder. */
const POSTED = {
  posted_at: '2026-12-01T18:00:00Z',
  post_at: '2026-12-01T18:00:00Z',
  reminded_at: null,
  remind: true,
};

test.describe('Post to Discord and the reminder (guided cards P.3)', () => {
  test('Post to Discord is on by default and the reminder under it', async ({ page }) => {
    await openRaidsTab(page, { viewer: 'officer', ...BOT_CHANNEL });
    const form = await openForm(page);
    const post = page.getByTestId('raid-form-discord');
    const remind = page.getByTestId('raid-form-remind');
    await expect(post).toHaveAttribute('role', 'switch');
    await expect(post).toBeChecked();
    await expect(post).toHaveAttribute('aria-checked', 'true');
    await expect(remind).toHaveAttribute('role', 'switch');
    await expect(remind).toBeChecked();
    await expect(page.getByTestId('raid-form-discord-helper')).toHaveText(LATER_HELPER);
    await expect(page.getByTestId('raid-form-remind-helper')).toHaveText(raidCopy.form.remind);
    await expectNoAxeViolations(page, '[data-testid="raid-form-create"]');
    await post.focus();
    await page.keyboard.press('Space');
    await expect(post).toHaveAttribute('aria-checked', 'false');
    await expect(form.getByTestId('raid-form-discord-helper')).toHaveText(raidCopy.form.postOff);
  });

  test('the first raid with no date posts as soon as it is saved', async ({ page }) => {
    await openRaidsTab(page, { viewer: 'officer', raids: [], ...BOT_CHANNEL });
    await openForm(page);
    await expect(page.getByTestId('raid-form-discord-helper')).toHaveText(
      raidCopy.form.postNow('#raid-signups'),
    );
  });

  test('turning Post to Discord off disables the reminder', async ({ page }) => {
    await openRaidsTab(page, { viewer: 'officer', ...BOT_CHANNEL });
    await openForm(page);
    await page.getByTestId('raid-form-discord').uncheck();
    const remind = page.getByTestId('raid-form-remind');
    await expect(remind).toBeDisabled();
    await expect(remind).toBeChecked();
    await expect(page.getByTestId('raid-form-remind-helper')).toHaveText(raidCopy.form.remindNeedsPost);
    await page.getByTestId('raid-form-discord').check();
    await expect(remind).toBeEnabled();
    await remind.focus();
    await page.keyboard.press('Space');
    await expect(remind).toHaveAttribute('aria-checked', 'false');
  });

  test('the saved values return as the next raid’s defaults', async ({ page }) => {
    const stub = await openRaidsTab(page, { viewer: 'officer', ...BOT_CHANNEL });
    await openForm(page);
    await page.getByTestId('raid-form-remind').uncheck();
    await page.getByTestId('raid-form-discord').uncheck();
    await page.getByTestId('raid-form-submit').click();
    await expect(page.getByTestId('raid-form-create')).toHaveCount(0);
    const posts = writesTo(stub, 'POST', /\/raids$/);
    expect(posts[0].body).toMatchObject({ post_to_discord: false, remind_day_before: false });
    await openForm(page);
    await expect(page.getByTestId('raid-form-discord')).not.toBeChecked();
    await expect(page.getByTestId('raid-form-remind')).not.toBeChecked();
    await expect(page.getByTestId('raid-form-remind')).toBeDisabled();
  });

  test('edit of a posted raid shows the posted line', async ({ page }) => {
    await openRaid(page, { viewer: 'officer', ...BOT_CHANNEL, discordPost: POSTED });
    await page.getByTestId('raid-officer-edit').click();
    await expect(page.getByTestId('raid-form-discord-posted')).toHaveText(
      raidCopy.form.posted('#raid-signups'),
    );
    await expect(page.getByTestId('raid-form-discord')).toHaveCount(0);
    await expect(page.getByTestId('raid-form-remind')).toBeEnabled();
  });

  test('a sent reminder reads Reminder sent', async ({ page }) => {
    const discordPost = { ...POSTED, reminded_at: '2026-12-10T02:00:00Z' };
    await openRaid(page, { viewer: 'officer', ...BOT_CHANNEL, discordPost });
    await page.getByTestId('raid-officer-edit').click();
    await expect(page.getByTestId('raid-form-remind-helper')).toHaveText(raidCopy.form.reminded);
    await expect(page.getByTestId('raid-form-remind')).toHaveCount(0);
  });

  for (const viewport of [PHONE, { width: 1440, height: 900 }]) {
    test(`the form and its Discord block shift nothing inside the form at ${viewport.width}`, async ({
      page,
    }) => {
      await page.setViewportSize(viewport);
      await page.addInitScript(() => {
        const shifts: number[] = [];
        Object.assign(window, { __shifts: shifts });
        const inForm = (node: Node | null): boolean =>
          (node instanceof Element ? node : (node?.parentElement ?? null))?.closest(
            '[data-testid="raid-form-create"]',
          ) != null;
        type Shift = { value: number; hadRecentInput: boolean; sources: { node: Node | null }[] };
        new PerformanceObserver((list) => {
          for (const entry of list.getEntries() as unknown as Shift[])
            if (!entry.hadRecentInput && entry.sources.some((source) => inForm(source.node)))
              shifts.push(entry.value);
        }).observe({ type: 'layout-shift', buffered: true });
      });
      await openRaidsTab(page, { viewer: 'officer', ...BOT_CHANNEL });
      await openForm(page);
      await expect(page.getByTestId('raid-form-discord-helper')).toBeVisible();
      await page.waitForTimeout(300);
      const height = await page
        .getByTestId('raid-form-discord')
        .evaluate((node) => node.closest('label')?.parentElement?.getBoundingClientRect().height ?? 0);
      expect(height).toBeGreaterThanOrEqual(viewport.width < 640 ? 164 : 144);
      const shifts = await page.evaluate(() => (window as unknown as { __shifts: number[] }).__shifts);
      expect(shifts.reduce((sum, value) => sum + value, 0)).toBe(0);
    });
  }

  test('both rows are 44 px at 360', async ({ page }) => {
    await page.setViewportSize(PHONE);
    await openRaidsTab(page, { viewer: 'officer', ...BOT_CHANNEL });
    await openForm(page);
    for (const id of ['raid-form-discord', 'raid-form-remind']) {
      const box = await page.getByTestId(id).locator('xpath=ancestor::label').boundingBox();
      expect(box?.height).toBeGreaterThanOrEqual(44);
    }
  });

  test('a posted raid link opens the raid page', async ({ page }) => {
    await stubGuildRaids(page, { viewer: 'member', raids: [storedRaid({ id: 1 })] });
    await page.clock.install({ time: NOW });
    await gotoHydrated(page, raidUrl(1), 'raid-page');
    await expect(page.getByTestId('raid-title')).toContainText("Onyxia's Lair");
  });
});
