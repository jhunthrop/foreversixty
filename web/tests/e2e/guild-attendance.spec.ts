// web/tests/e2e/guild-attendance.spec.ts
// Attendance and the raid-night report acceptance (design/specs/2026-10-27-guild-attendance.md §9,
// the E2E paragraph, plus the browser behaviours Tasks 7 and 8 left to this file): the ended band and
// the poll that turns it done; the match picker's keyboard, races and closing; the clear confirm;
// Excused and the summary; Show all; Your night and R-A1 for members; R-A3 (no "no-show"); Copy and
// Post to Discord; officer-read-only; the roster chip; the readiness next-raid filter; axe; the 44 px
// floor at 360. The API is stubbed by support/guild-attendance.ts.
import type { Page } from '@playwright/test';
import { expect, test } from './support/test';
import { attendanceCopy } from '../../src/lib/guild/attendance-copy';
import {
  AFTER_END,
  AFTER_WINDOW,
  DONE,
  MATCHED_REPORT,
  NO_EXPORT,
  SECOND_REPORT,
  STALE_EXPORT,
  TAKEN_REPORT,
  bareReport,
  expectedDiscordText,
  fixtureCandidates,
  memberNoShow,
  rosterAttendance,
  signedKeys,
  storedNight,
  stubGuildAttendance,
  type AttendanceStub,
  type AttendanceStubOptions,
} from './support/guild-attendance';
import { CHICAGO, GUILD_URL, NOW, ONYXIA_ID, keyOf, raidUrl, writesTo } from './support/guild-raids';
import { gotoHydrated } from './support/hydrated';
import { expectNoAxeViolations } from './support/officer-tools';
import { shortControls } from './support/short-controls';

const POLL_MS = 10_000;
const PHONE = { width: 360, height: 800 };
const EXCUSE_DELAY_MS = 300;

test.use({ timezoneId: CHICAGO });

async function openRaid(page: Page, options: AttendanceStubOptions, at = AFTER_END): Promise<AttendanceStub> {
  const stub = await stubGuildAttendance(page, options);
  await page.clock.install({ time: at });
  await gotoHydrated(page, raidUrl(ONYXIA_ID), 'raid-page');
  return stub;
}

const officerDone = (extra: Partial<AttendanceStubOptions> = {}): AttendanceStubOptions => ({
  viewer: 'officer',
  nights: DONE,
  ...extra,
});

async function openPicker(page: Page): Promise<void> {
  await page.getByTestId('raid-match-open').click();
  await expect(page.getByTestId('raid-match-rows')).toBeVisible();
}

const row = (page: Page, id: string) => page.getByTestId(`raid-match-row-${id}`);

/** Every text node on the page, lower-cased, that holds the word R-A3 forbids. */
function noShowTexts(page: Page): Promise<string[]> {
  return page.evaluate(() => {
    const walker = document.createTreeWalker(document.body, NodeFilter.SHOW_TEXT);
    const found: string[] = [];
    for (let node = walker.nextNode(); node !== null; node = walker.nextNode()) {
      if (/no-show/i.test(node.textContent ?? '')) found.push(node.textContent ?? '');
    }
    return found;
  });
}

test.describe('ended with no log, and the poll (§3.1, §4 "Poll")', () => {
  test('the band says Ended · no log yet, inside the window and after it', async ({ page }) => {
    await openRaid(page, { viewer: 'member' });
    await expect(page.getByTestId('raid-ended-line')).toContainText(
      `${attendanceCopy.ended.noLog}${attendanceCopy.ended.matchSoon}`,
    );
    await expect(page.getByTestId('raid-match-open')).toHaveCount(0);

    await openRaid(page, { viewer: 'officer' }, AFTER_WINDOW);
    const ended = page.getByTestId('raid-ended-line');
    await expect(ended).toContainText(attendanceCopy.ended.uploadLink);
    await expect(page.getByTestId('raid-match-open')).toHaveText(attendanceCopy.ended.matchALog);
  });

  test('the poll matches the log in place and the live region speaks', async ({ page }) => {
    const stub = await openRaid(page, { viewer: 'officer' });
    await expect(page.getByTestId('raid-live-region')).toHaveText('');
    stub.matchFromElsewhere(ONYXIA_ID);
    await page.clock.runFor(POLL_MS);
    await expect(page.getByTestId('raid-night-line')).toContainText('13 of 14 came');
    await expect(page.getByTestId('raid-live-region')).toHaveText(attendanceCopy.liveRegion);
    await expect(page.getByTestId('raid-report')).toBeVisible();
    await expect(page.getByTestId('raid-signup-strip')).toHaveCount(0);
  });
});

test.describe('the match picker (§3.2)', () => {
  test('Closest is chosen and focused, the taken row is disabled, arrows and Space choose, Match matches', async ({
    page,
  }) => {
    const stub = await openRaid(page, { viewer: 'officer' });
    await openPicker(page);
    await expect(row(page, MATCHED_REPORT)).toHaveAttribute('aria-checked', 'true');
    await expect(row(page, MATCHED_REPORT)).toBeFocused();
    await expect(row(page, TAKEN_REPORT)).toHaveAttribute('aria-disabled', 'true');
    await expect(page.getByTestId('raid-match-elsewhere')).toContainText('Clear it there first.');

    await page.keyboard.press('ArrowDown');
    await expect(row(page, SECOND_REPORT)).toBeFocused();
    await page.keyboard.press('ArrowDown');
    await expect(row(page, MATCHED_REPORT)).toBeFocused();
    await page.keyboard.press('ArrowUp');
    await expect(row(page, SECOND_REPORT)).toBeFocused();
    await page.keyboard.press('Space');
    await expect(row(page, SECOND_REPORT)).toHaveAttribute('aria-checked', 'true');
    await expect(row(page, MATCHED_REPORT)).toHaveAttribute('aria-checked', 'false');
    await row(page, TAKEN_REPORT).click({ force: true });
    await expect(row(page, TAKEN_REPORT)).toHaveAttribute('aria-checked', 'false');

    await page.getByTestId('raid-match-submit').click();
    await expect(page.getByTestId('raid-match-picker')).toHaveCount(0);
    await expect(page.getByTestId('raid-open-night')).toBeFocused();
    await expect(page.getByTestId('raid-match-uploaded')).toContainText(attendanceCopy.match.byOfficer);
    expect(writesTo(stub.raids, 'PUT', /\/raids\/7\/match$/).map((w) => w.body)).toEqual([
      { report_id: SECOND_REPORT },
    ]);
  });

  test('Cancel closes it and returns focus to Match a log', async ({ page }) => {
    await openRaid(page, { viewer: 'officer' });
    await openPicker(page);
    await page.getByTestId('raid-match-cancel').click();
    await expect(page.getByTestId('raid-match-picker')).toHaveCount(0);
    await expect(page.getByTestId('raid-match-open')).toBeFocused();
  });

  test('409 report_matched turns the row matched-elsewhere with the focused line', async ({ page }) => {
    const stub = await openRaid(page, { viewer: 'officer' });
    await openPicker(page);
    stub.refuseNext = { match: { code: 'report_matched', status: 409 } };
    stub.candidates = fixtureCandidates();
    await page.getByTestId('raid-match-submit').click();
    const line = page.getByTestId('raid-match-status');
    await expect(line).toHaveText(attendanceCopy.picker.matchedElsewhere);
    await expect(line).toBeFocused();
    await expect(row(page, MATCHED_REPORT)).toHaveAttribute('aria-disabled', 'true');
    await expect(page.getByTestId('raid-match-submit')).toBeDisabled();
  });

  test('403 guild_plan_required closes the picker and says the tools ended', async ({ page }) => {
    const stub = await openRaid(page, { viewer: 'officer', officerTools: true });
    await openPicker(page);
    stub.refuseNext = { match: { code: 'guild_plan_required', status: 403 } };
    await page.getByTestId('raid-match-submit').click();
    await expect(page.getByTestId('raid-match-picker')).toHaveCount(0);
    const refused = page.getByTestId('raid-match-refused');
    await expect(refused).toHaveText(attendanceCopy.errors.officerToolsEnded);
    await expect(refused).toBeFocused();
    await expect(page.getByTestId('raid-match-open')).toBeDisabled();
  });

  test('a poll that brings a match closes the picker and hands focus to Open the night', async ({ page }) => {
    const stub = await openRaid(page, { viewer: 'officer' });
    await openPicker(page);
    await expect(row(page, MATCHED_REPORT)).toBeFocused();
    stub.matchFromElsewhere(ONYXIA_ID);
    await page.clock.runFor(POLL_MS);
    await expect(page.getByTestId('raid-match-picker')).toHaveCount(0);
    await expect(page.getByTestId('raid-open-night')).toBeFocused();
  });

  test('the empty list and the unknown log zone read their lines', async ({ page }) => {
    const stub = await openRaid(page, { viewer: 'officer' });
    stub.candidates = { zone_known: true, rows: [] };
    await page.getByTestId('raid-match-open').click();
    await expect(page.getByTestId('raid-match-empty')).toContainText("No Onyxia's Lair logs uploaded");
    await page.getByTestId('raid-match-cancel').click();
    stub.candidates = fixtureCandidates({ zone_known: false });
    await openPicker(page);
    await expect(page.getByTestId('raid-match-helper')).toContainText('so every zone shows');
  });
});

test.describe('Not this log (§3.1 clear confirm)', () => {
  test('focus lands on Keep it; Escape keeps it; Clear the match clears', async ({ page }) => {
    const stub = await openRaid(page, officerDone());
    await page.getByTestId('raid-not-this-log').click();
    await expect(page.getByTestId('raid-clear-confirm')).toContainText(attendanceCopy.clearConfirm.prompt);
    await expect(page.getByTestId('raid-clear-keep')).toBeFocused();
    await page.keyboard.press('Escape');
    await expect(page.getByTestId('raid-clear-confirm')).toHaveCount(0);
    await expect(page.getByTestId('raid-not-this-log')).toBeFocused();
    expect(writesTo(stub.raids, 'DELETE', /\/match$/)).toEqual([]);

    await page.getByTestId('raid-not-this-log').click();
    await page.getByTestId('raid-clear-confirm-button').click();
    await expect(page.getByTestId('raid-ended-line')).toBeVisible();
    await expect(page.getByTestId('raid-match-open')).toBeFocused();
    await expect(page.getByTestId('raid-report')).toHaveCount(0);
    expect(writesTo(stub.raids, 'DELETE', /\/raids\/7\/match$/)).toHaveLength(1);
  });
});

test.describe('the report and attendance, officer (§3.4, §3.5)', () => {
  test('Excused ticks overlap and the summary counts both', async ({ page }) => {
    const stub = await openRaid(page, officerDone({ excuseDelayMs: EXCUSE_DELAY_MS }));
    const summary = page.getByTestId('raid-attendance-summary');
    await expect(summary).toHaveText('13 of 14 came · 1 late · 1 not in the log · 2 more joined');
    await page.getByLabel('Excused, Hexbramble').check();
    await page.getByLabel('Excused, Mendwhisper').check();
    await expect(summary).toContainText(' · 2 excused');
    await expect(page.getByLabel('Excused, Hexbramble')).toBeChecked();
    await expect(page.getByLabel('Excused, Mendwhisper')).toBeChecked();
    await expect.poll(() => writesTo(stub.raids, 'PUT', /\/attendance\//).length).toBe(2);
    await page.waitForTimeout(EXCUSE_DELAY_MS * 2);
    await expect(summary).toContainText(' · 2 excused');
    await expect(page.getByLabel('Excused, Hexbramble')).toBeChecked();
    expect(writesTo(stub.raids, 'PUT', /attendance\/us\/pvp\/hexbramble$/).map((w) => w.body)).toEqual([
      { excused: true },
    ]);
  });

  test('a refused tick reverts with the row line', async ({ page }) => {
    // The refusal is held, so the optimistic tick is seen before it rolls back; `click`, not
    // `check`, since check() itself fails when the box ends unticked.
    const stub = await openRaid(page, officerDone({ excuseDelayMs: EXCUSE_DELAY_MS }));
    stub.refuseNext = { excuse: { code: 'internal', status: 500 } };
    const box = page.getByLabel('Excused, Hexbramble');
    await box.click();
    await expect(box).toBeChecked();
    await expect(page.getByTestId('raid-attendance-line')).toHaveText(attendanceCopy.errors.rows);
    await expect(box).not.toBeChecked();
  });

  test('every group shows, with the Came as row', async ({ page }) => {
    await openRaid(page, officerDone());
    for (const group of ['noShow', 'late', 'joined', 'present', 'benchOut']) {
      await expect(page.getByTestId(`raid-attendance-${group}`)).toBeVisible();
    }
    await expect(page.getByTestId('raid-attendance-noShow')).toContainText('Hexbramble');
    // Present rows are compact: the alt reads `Felsnap as Mirestalk` there.
    await expect(page.getByTestId('raid-attendance-present')).toContainText('Felsnap as Mirestalk');
  });

  test('Show all moves focus to the first revealed death row', async ({ page }) => {
    await openRaid(page, officerDone());
    const rows = page.getByTestId('raid-death-row');
    await expect(rows).toHaveCount(10);
    await page.getByTestId('raid-deaths-show-all').click();
    await expect(rows).toHaveCount(11);
    await expect(rows.nth(10)).toBeFocused();
    await expect(page.getByTestId('raid-deaths-show-all')).toHaveCount(0);
  });

  test('Copy for Discord writes the fixture text', async ({ page, context }) => {
    await context.grantPermissions(['clipboard-read', 'clipboard-write']);
    await openRaid(page, officerDone());
    const copy = page.getByTestId('raid-report-copy');
    await copy.click();
    await expect(copy).toHaveText(attendanceCopy.report.copied);
    const copied = await page.evaluate(() => navigator.clipboard.readText());
    const url = new URL(page.url());
    expect(copied).toBe(expectedDiscordText(`${url.origin}${url.pathname}#report`));
    expect(copied.split('\n').length).toBeLessThanOrEqual(12);
    expect(copied).not.toContain('Hexbramble');
  });

  test('Post to Discord posts once and the stub sends one report_ready', async ({ page }) => {
    const stub = await openRaid(page, officerDone({ discordConnected: true }));
    await page.getByTestId('raid-report-post').click();
    await expect(page.getByTestId('raid-report-posted')).toContainText('Posted to Discord at');
    await expect(page.getByTestId('raid-report-post')).toHaveCount(0);
    expect(stub.reportReady).toBe(1);
    expect(writesTo(stub.raids, 'POST', /\/report\/post$/)).toHaveLength(1);
  });

  test('409 already_posted re-reads and shows the posted line', async ({ page }) => {
    const stub = await openRaid(page, officerDone({ discordConnected: true }));
    const night = storedNight();
    stub.nights.set(ONYXIA_ID, { ...night, report: { ...night.report, posted_at: '2026-12-11T06:01:00Z' } });
    await page.getByTestId('raid-report-post').click();
    await expect(page.getByTestId('raid-report-posted')).toBeVisible();
    await expect(page.getByTestId('raid-report-post-error')).toHaveCount(0);
    expect(stub.reportReady).toBe(0);
  });

  test('no buff data, no previous raid and no loot each read their line', async ({ page }) => {
    await openRaid(page, officerDone({ nights: { [ONYXIA_ID]: storedNight({ report: bareReport() }) } }));
    await expect(page.getByTestId('raid-consumables-no-buffs')).toHaveText(
      attendanceCopy.consumables.noBuffData,
    );
    await expect(page.getByTestId('raid-moves-first')).toContainText("First Onyxia's Lair raid logged.");
    await expect(page.getByTestId('raid-loot-none')).toContainText('No loot awarded for this raid.');
  });
});

test.describe('a done raid band and its closed sections (§3.1, R-A4)', () => {
  test("a done raid draws no officer row; Not this log is the band's one officer control", async ({
    page,
  }) => {
    await openRaid(page, officerDone());
    await expect(page.getByTestId('raid-officer-row')).toHaveCount(0);
    await expect(page.getByTestId('raid-not-this-log')).toBeVisible();
  });

  test('an ended raid with no log keeps Edit alone beside Match a log', async ({ page }) => {
    await openRaid(page, { viewer: 'officer' });
    await expect(page.getByTestId('raid-officer-edit')).toBeVisible();
    for (const id of ['raid-officer-lock', 'raid-officer-copy', 'raid-officer-cancel']) {
      await expect(page.getByTestId(id)).toHaveCount(0);
    }
    await expect(page.getByTestId('raid-match-open')).toBeVisible();
  });

  test('Notes and Sign-ups sit closed below Attendance; opened, the selects are disabled', async ({
    page,
  }) => {
    await openRaid(page, officerDone());
    for (const id of ['raid-notes-disclosure', 'raid-roster-disclosure']) {
      await expect(page.getByTestId(id)).not.toHaveAttribute('open', '');
    }
    await expect(page.getByTestId('raid-roster')).toBeHidden();
    await page.getByTestId('raid-roster-disclosure-summary').click();
    await expect(page.getByTestId('raid-roster')).toBeVisible();
    const selects = page.locator('select[data-testid^="raid-roster-select-"]');
    expect(await selects.count()).toBeGreaterThan(0);
    expect(await selects.evaluateAll((n) => n.every((s) => (s as HTMLSelectElement).disabled))).toBe(true);
  });

  test('loot draws the client icons, a quest item too, and opens a tooltip for a quest item', async ({
    page,
  }) => {
    await openRaid(page, officerDone());
    const row = (name: string) => page.getByTestId('raid-loot-row').filter({ hasText: name });
    await expect(row('Netherwind Crown').getByTestId('raid-loot-icon').locator('img')).toHaveAttribute(
      'src',
      /icons\/inv_helmet_70\.webp$/,
    );
    await expect(row('Head of Onyxia').getByTestId('raid-loot-icon').locator('img')).toHaveAttribute(
      'src',
      /icons\/inv_misc_head_dragon_01\.webp$/,
    );
    await page.getByTestId('item-hover-18423').first().focus();
    await expect(page.getByTestId('item-tooltip')).toContainText('Head of Onyxia');
  });
});

test.describe('a member (R-A1, R-A3)', () => {
  test('sees Your night and none of the officer rows but their own', async ({ page }) => {
    await openRaid(page, { viewer: 'member', nights: DONE });
    await expect(page.getByTestId('raid-your-night-sentence')).toHaveText(
      'Signed In as DPS. Present from the first pull, 4 of 4 pulls.',
    );
    await expect(page.getByTestId('raid-your-night-facts')).toContainText('0 deaths');
    await expect(page.getByTestId('raid-your-night-facts')).toContainText('up 16');
    await expect(page.getByTestId('raid-attendance')).toHaveCount(0);
    for (const block of ['raid-deaths', 'raid-consumables', 'raid-moves']) {
      await expect(page.getByTestId(block)).toHaveCount(0);
    }
    await expect(page.getByTestId('raid-kill-row')).toHaveCount(1);
    await expect(page.getByTestId('raid-loot-row')).toHaveCount(4);
    await expect(page.getByTestId('raid-report-copy')).toHaveCount(0);
  });

  test('the no-show sentence renders and no text node says no-show', async ({ page }) => {
    await openRaid(page, { viewer: 'member', nights: memberNoShow() });
    await expect(page.getByTestId('raid-your-night-sentence')).toHaveText(
      'Signed In as DPS. You are not in this log. If you were there under another name, tell an officer.',
    );
    expect(await noShowTexts(page)).toEqual([]);
    // Not in the log: no deaths and no DPS fact (UX 5, WP 2).
    await expect(page.getByTestId('raid-your-night-facts')).toHaveCount(0);
  });

  test('the officer page carries no no-show text either', async ({ page }) => {
    await openRaid(page, officerDone());
    expect(await noShowTexts(page)).toEqual([]);
  });
});

test('officer-read-only: match, clear, Excused and Post disabled; Copy live', async ({ page, context }) => {
  await context.grantPermissions(['clipboard-read', 'clipboard-write']);
  const stub = await openRaid(page, officerDone({ officerTools: false, discordConnected: true }));
  await expect(page.getByTestId('raid-not-this-log')).toBeDisabled();
  await expect(page.getByTestId('raid-report-post')).toBeDisabled();
  const boxes = page.getByTestId('raid-attendance-excused');
  await expect(boxes).toHaveCount(2);
  // A done raid draws no officer row; the gate sentence shows once, under the Report header.
  await expect(page.getByTestId('raid-officer-row')).toHaveCount(0);
  await expect(page.locator('#officer-tools-gate-raids')).toHaveCount(1);
  await expect(page.getByTestId('guild-raids-plan-gate-report')).toBeVisible();
  for (const box of await boxes.all()) await expect(box).toBeDisabled();
  await page.getByTestId('raid-report-copy').click();
  await expect(page.getByTestId('raid-report-copy')).toHaveText(attendanceCopy.report.copied);
  expect(stub.writes).toEqual([]);

  const ended = await context.newPage();
  await openRaid(ended, { viewer: 'officer', officerTools: false });
  await expect(ended.getByTestId('raid-match-open')).toBeDisabled();
});

test.describe('the roster chip (§3.6)', () => {
  test('raid mode: two windows, No raids yet and the footnote', async ({ page }) => {
    await stubGuildAttendance(page, { viewer: 'member', attendanceSource: 'raids' });
    await page.clock.install({ time: NOW });
    await gotoHydrated(page, `${GUILD_URL}#roster`, 'guild-roster-tab');
    const rows = page.getByTestId('guild-roster-rows');
    await expect(rows).toContainText(/Came \d of 8 · \d+ of 30 raids/);
    await expect(rows).toContainText(attendanceCopy.roster.noRaids);
    // One read for the roster and the guild line (WP 4): Kraggor, who led the raid, has raids.
    const { last8, last30 } = rosterAttendance('Kraggor');
    await expect(rows).toContainText(`Came ${last8.present} of 8 · ${last30.present} of 30 raids`);
    await expect(page.getByTestId('guild-roster-attendance-footnote')).toHaveText(
      attendanceCopy.roster.footnote,
    );
  });

  test('report mode: the nights chip and no footnote', async ({ page }) => {
    await stubGuildAttendance(page, { viewer: 'member', attendanceSource: 'reports' });
    await page.clock.install({ time: NOW });
    await gotoHydrated(page, `${GUILD_URL}#roster`, 'guild-roster-tab');
    await expect(page.getByTestId('guild-roster-rows')).not.toContainText('Came ');
    await expect(page.getByTestId('guild-roster-attendance-footnote')).toHaveCount(0);
  });
});

test('readiness: Next raid is on by default, its rows are the sign-ups, and off shows the guild', async ({
  page,
}) => {
  const stub = await stubGuildAttendance(page, { viewer: 'officer' });
  await page.clock.install({ time: NOW });
  await gotoHydrated(page, `${GUILD_URL}#readiness`, 'guild-readiness-tab');
  const toggle = page.getByTestId('guild-readiness-next-raid');
  await expect(toggle).toHaveText('Next raid · Thu 10 Dec');
  await expect(toggle).toHaveAttribute('aria-pressed', 'true');
  const rowKeys = () =>
    page
      .locator('[data-testid^="guild-readiness-row-"]:visible')
      .evaluateAll((nodes) =>
        nodes.map((n) => (n.getAttribute('data-testid') ?? '').replace('guild-readiness-row-', '')),
      );
  const signed = signedKeys(stub);
  await expect.poll(async () => (await rowKeys()).sort()).toEqual([...signed].sort());
  await expect(page.getByTestId('guild-readiness-showing')).toContainText(`Showing the ${signed.length}`);
  await expect(page.getByTestId('guild-readiness-stale').first()).toContainText('Synced');
  await expect(
    page
      .getByTestId(`guild-readiness-row-${keyOf(STALE_EXPORT)}`)
      .locator('[data-testid="guild-readiness-stale"]:visible'),
  ).toBeVisible();
  await expect(
    page
      .getByTestId(`guild-readiness-row-${keyOf(NO_EXPORT)}`)
      .locator('[data-testid="guild-readiness-no-export"]:visible'),
  ).toBeVisible();
  expect(stub.readinessReads).toContain(String(ONYXIA_ID));
  // No export: Gear gap through Item level read a dash (UX 16, WP 14).
  await expect(
    page
      .getByTestId(`guild-readiness-row-${keyOf(NO_EXPORT)}`)
      .locator('[data-testid="guild-readiness-dash"]:visible'),
  ).toHaveCount(5);

  await toggle.click();
  await expect(toggle).toHaveAttribute('aria-pressed', 'false');
  await expect.poll(async () => (await rowKeys()).length).toBeGreaterThan(signed.length);
});

test.describe('axe (§9)', () => {
  const states: [string, AttendanceStubOptions, Date][] = [
    ['ended, officer', { viewer: 'officer' }, AFTER_END],
    ['ended, member after the window', { viewer: 'member' }, AFTER_WINDOW],
    ['done, officer', officerDone({ discordConnected: true }), AFTER_END],
    ['done, member', { viewer: 'member', nights: DONE }, AFTER_END],
    ['done, member no-show', { viewer: 'member', nights: memberNoShow() }, AFTER_END],
    ['done, officer-read-only', officerDone({ officerTools: false, discordConnected: true }), AFTER_END],
  ];
  for (const [name, options, at] of states) {
    test(`the raid page is clean: ${name}`, async ({ page }) => {
      await openRaid(page, options, at);
      await expectNoAxeViolations(page, 'main');
    });
  }

  test('the picker and the clear confirm are clean', async ({ page }) => {
    await openRaid(page, { viewer: 'officer' });
    await openPicker(page);
    await expectNoAxeViolations(page, 'main');
    await openRaid(page, officerDone());
    await page.getByTestId('raid-not-this-log').click();
    await expectNoAxeViolations(page, 'main');
  });

  test('the roster in raid mode and readiness on the next raid are clean', async ({ page }) => {
    await stubGuildAttendance(page, { viewer: 'officer', attendanceSource: 'raids' });
    await page.clock.install({ time: NOW });
    await gotoHydrated(page, `${GUILD_URL}#roster`, 'guild-roster-tab');
    // The rows and the footnote: the sort chips' aria-sort predates this lane (see the report).
    await expectNoAxeViolations(page, '[data-testid="guild-roster-rows"]');
    await expectNoAxeViolations(page, '[data-testid="guild-roster-attendance-footnote"]');
    await gotoHydrated(page, `${GUILD_URL}#readiness`, 'guild-readiness-next-raid');
    await expectNoAxeViolations(page, '[data-testid="guild-readiness-tab"]');
  });
});

test.describe('every control is at least 44 px tall at 360', () => {
  test.beforeEach(({ page }, testInfo) => {
    test.skip(testInfo.project.name !== 'mobile', 'the 360 floor is a phone rule');
    return page.setViewportSize(PHONE);
  });

  test('the done page for the officer, deaths and consumables expanded', async ({ page }) => {
    await openRaid(page, officerDone({ discordConnected: true }));
    await page.getByTestId('raid-deaths-show-all').click();
    expect(await shortControls(page.getByTestId('raid-page'))).toEqual([]);
  });

  test('the done page for the member', async ({ page }) => {
    await openRaid(page, { viewer: 'member', nights: DONE });
    expect(await shortControls(page.getByTestId('raid-page'))).toEqual([]);
  });

  test('the ended band, the picker and the clear confirm', async ({ page }) => {
    await openRaid(page, { viewer: 'officer' }, AFTER_WINDOW);
    expect(await shortControls(page.getByTestId('raid-page'))).toEqual([]);
    await openPicker(page);
    expect(await shortControls(page.getByTestId('raid-match-picker'))).toEqual([]);
    await openRaid(page, officerDone());
    await page.getByTestId('raid-not-this-log').click();
    expect(await shortControls(page.getByTestId('raid-match-line'))).toEqual([]);
  });

  test('the readiness filter', async ({ page }) => {
    await stubGuildAttendance(page, { viewer: 'officer' });
    await page.clock.install({ time: NOW });
    await gotoHydrated(page, `${GUILD_URL}#readiness`, 'guild-readiness-next-raid');
    expect(await shortControls(page.getByTestId('guild-readiness-header'))).toEqual([]);
    expect(await shortControls(page.getByTestId('guild-readiness-showing'))).toEqual([]);
  });
});
