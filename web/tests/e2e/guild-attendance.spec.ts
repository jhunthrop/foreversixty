// web/tests/e2e/guild-attendance.spec.ts
// Attendance and the raid-night report acceptance (design/specs/2026-10-27-guild-attendance.md §9,
// the E2E paragraph, plus the browser behaviours Tasks 7 and 8 left to this file): the ended band and
// the poll that turns it done; the match picker's keyboard, races and closing; the clear confirm;
// the raid report (web layout spec 2026-11-05 §7.2): the answer figures, a pull's deaths, DPS and
// HPS, consumables by pull, the attendance fix and the summary; Your night and a member's own rows;
// R-A3 (no "no-show"); Copy, Post and Post again; officer-read-only; the roster chip; the readiness
// next-raid filter; axe; the 44 px floor at 360. The API is stubbed by support/guild-attendance.ts.
import type { Page } from '@playwright/test';
import { expect, test, apiUnreachable, GUILD_OVERVIEW_READS } from './support/test';
import { attendanceCopy, rosterChip } from '../../src/lib/guild/attendance-copy';
import { reportNightCopy as copyOf } from '../../src/lib/guild/report-night-copy';
import {
  AFTER_END,
  AFTER_WINDOW,
  DONE,
  MATCHED_REPORT,
  NO_EXPORT,
  SECOND_REPORT,
  STALE_EXPORT,
  TAKEN_REPORT,
  allAttendance,
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
import { choose } from './support/select';
import { shortControls } from './support/short-controls';

test.beforeEach(({ page }) => apiUnreachable(page, ...GUILD_OVERVIEW_READS));

const POLL_MS = 10_000;
const PHONE = { width: 360, height: 800 };
const FIX_DELAY_MS = 300;

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
    await expect(page.getByTestId('raid-attendance-summary')).toContainText('13 of 14 came');
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

/** The officer's fix rows sit behind Fix attendance (mock web-guild-report). */
async function openFixRows(page: Page): Promise<void> {
  await page.getByTestId('raid-attendance-fix-open').locator('summary').click();
}

test.describe('the report and attendance, officer (web layout spec 2026-11-05 §7.2)', () => {
  const fix = (page: Page, name: string) => page.getByTestId(`raid-attendance-fix-${keyOf(name)}`);
  const attendanceRowOf = (page: Page, name: string) =>
    page.getByTestId('raid-attendance-row').filter({ hasText: name });

  test('the head answers kills, wipes, attempts and deaths', async ({ page }) => {
    await openRaid(page, officerDone());
    await expect(page.getByTestId('raid-report-figures')).toHaveText(
      /Kills\s*1\s*Wipes\s*3\s*Attempts\s*4\s*Deaths\s*17/,
    );
  });

  test('a fix to Bench reads Bench with the log muted, the summary moves, Back to the log restores', async ({
    page,
  }) => {
    const stub = await openRaid(page, officerDone({ fixDelayMs: FIX_DELAY_MS }));
    await openFixRows(page);
    const summary = page.getByTestId('raid-attendance-summary');
    await expect(summary).toHaveText('13 of 14 came · 1 late · 1 not in the log · 2 more joined');
    await choose(fix(page, 'Hexbramble'), 'bench');
    await expect(attendanceRowOf(page, 'Hexbramble').getByTestId('raid-attendance-outcome')).toHaveText(
      'Bench, set by Kraggor',
    );
    await expect(attendanceRowOf(page, 'Hexbramble').getByTestId('raid-attendance-auto')).toHaveText(
      'The log said Not in the log',
    );
    await expect(summary).toHaveText('14 of 14 came · 1 late · 2 more joined');
    expect(
      writesTo(stub.raids, 'PUT', /attendance\/us\/pvp\/hexbramble\/override$/).map((w) => w.body),
    ).toEqual([{ status: 'bench' }]);
    await attendanceRowOf(page, 'Hexbramble').getByTestId('raid-attendance-back').click();
    await expect(attendanceRowOf(page, 'Hexbramble').getByTestId('raid-attendance-outcome')).toHaveText(
      'Not in the log',
    );
    await expect(summary).toHaveText('13 of 14 came · 1 late · 1 not in the log · 2 more joined');
    expect(writesTo(stub.raids, 'DELETE', /hexbramble\/override$/)).toHaveLength(1);
  });

  test('a refused fix keeps the row and says so under it', async ({ page }) => {
    const stub = await openRaid(page, officerDone());
    stub.refuseNext = { fix: { code: 'internal', status: 500 } };
    await openFixRows(page);
    await choose(fix(page, 'Hexbramble'), 'present');
    await expect(attendanceRowOf(page, 'Hexbramble').getByTestId('raid-attendance-line')).toHaveText(
      attendanceCopy.errors.rows,
    );
    await expect(attendanceRowOf(page, 'Hexbramble').getByTestId('raid-attendance-outcome')).toHaveText(
      'Not in the log',
    );
  });

  test('every raider has a row, the likely fixes first, the alt read as Came as', async ({ page }) => {
    await openRaid(page, officerDone());
    await openFixRows(page);
    const rows = page.getByTestId('raid-attendance-row');
    await expect(rows).toHaveCount(allAttendance().length);
    await expect(rows.first()).toContainText('Hexbramble');
    await expect(attendanceRowOf(page, 'Felsnap')).toContainText('Came as Mirestalk');
  });

  test('a wipe opens its deaths: the killing blow and the last hits', async ({ page }) => {
    await openRaid(page, officerDone());
    const pulls = page.getByTestId('raid-pull-row');
    await expect(pulls).toHaveCount(4);
    const inspector = page.getByTestId('raid-pull-inspector');
    const wide = (page.viewportSize()?.width ?? 0) >= 1024;
    if (wide) {
      // From 1024 the Inspector sits beside the plate on the first wipe with deaths.
      await expect(inspector).toHaveAttribute('data-placement', 'side');
    } else {
      await expect(inspector).toHaveCount(0);
      await pulls.first().getByTestId('raid-pull-open').click();
      await expect(inspector).toHaveAttribute('data-placement', 'inline');
    }
    await expect(pulls.first().getByTestId('raid-pull-open')).toHaveAttribute('aria-expanded', 'true');
    await expect(inspector.getByTestId('raid-pull-death')).toHaveCount(11);
    await expect(inspector.getByTestId('raid-pull-killing-blow').first()).toContainText(
      'Flame Breath from Onyxia, 4,210 (1,050 overkill)',
    );
    await expect(inspector.getByTestId('raid-pull-last-hit').first()).toContainText(
      'Melee from Onyxian Whelp',
    );
    await pulls.nth(1).getByTestId('raid-pull-open').click();
    await expect(inspector.getByTestId('raid-pull-death')).toHaveCount(3);
  });

  test('DPS and HPS read the guild unit, and the consumables grid says why a pull is unknown', async ({
    page,
  }) => {
    await openRaid(page, officerDone());
    await expect(page.getByTestId('raid-output-row')).toHaveCount(5);
    await expect(page.getByTestId('raid-output')).toContainText(/291\s*DPS/);
    await expect(page.getByTestId('raid-output')).toContainText(/443\s*HPS/);
    await expect(page.getByTestId('raid-consumables-grid')).toBeVisible();
    await expect(page.getByTestId('raid-consumables-unknown')).toHaveText(
      'Onyxia, pull 1: the log carried no character sheet, so nothing can be said.',
    );
  });

  test('Copy for Discord writes the fixture text, naming nobody for dying', async ({ page, context }) => {
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
    expect(copied).not.toContain('Most deaths');
  });

  test('Post to Discord posts once, then Post again reposts', async ({ page }) => {
    const stub = await openRaid(page, officerDone({ discordConnected: true }));
    const post = page.getByTestId('raid-report-post');
    await expect(post).toHaveText(copyOf.discord.post);
    await post.click();
    await expect(page.getByTestId('raid-report-posted')).toContainText('Posted to Discord at');
    await expect(post).toHaveText(copyOf.discord.postAgain);
    expect(stub.reportReady).toBe(1);
    await post.click();
    await expect.poll(() => stub.reportReady).toBe(2);
    expect(writesTo(stub.raids, 'POST', /\/report\/post$/).map((w) => w.body)).toEqual([
      null,
      { repost: true },
    ]);
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

  test('no character sheets and no loot each read their line, never red', async ({ page }) => {
    await openRaid(page, officerDone({ nights: { [ONYXIA_ID]: storedNight({ report: bareReport() }) } }));
    await expect(page.getByTestId('raid-consumables-grid')).toHaveCount(0);
    await expect(page.getByTestId('raid-consumables-unknown')).toHaveCount(4);
    await expect(page.getByTestId('raid-consumables')).not.toContainText(
      attendanceCopy.consumables.noBuffData,
    );
    await expect(page.getByTestId('raid-loot-empty')).toContainText('No loot awarded for this raid.');
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

  test('a done raid draws no sign-up strip and no sign-up roster', async ({ page }) => {
    await openRaid(page, officerDone());
    await expect(page.getByTestId('raid-signup-strip')).toHaveCount(0);
    await expect(page.getByTestId('raid-roster')).toHaveCount(0);
  });

  test('loot draws the client icons, a quest item too, and opens a tooltip for a quest item', async ({
    page,
  }) => {
    await openRaid(page, officerDone());
    const row = (name: string) => page.getByTestId('raid-loot-row').filter({ hasText: name });
    await expect(row('Netherwind Crown').locator('img')).toHaveAttribute(
      'src',
      /icons\/inv_helmet_70\.webp$/,
    );
    await expect(row('Head of Onyxia').locator('img')).toHaveAttribute(
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
    // Every pull, but only their own deaths, output and consumables; no one else's rows.
    await expect(page.getByTestId('raid-report-figures')).toContainText('17');
    await expect(page.getByTestId('raid-pull-row')).toHaveCount(4);
    await expect(page.getByTestId('raid-pull-open')).toHaveCount(0);
    await expect(page.getByTestId('raid-output-row')).toHaveCount(1);
    await expect(page.getByTestId('raid-consumables-grid-row')).toHaveCount(1);
    await expect(page.getByTestId('raid-attendance-row')).toHaveCount(0);
    await expect(page.getByTestId('raid-loot-row')).toHaveCount(4);
    for (const officerOnly of ['raid-discord', 'raid-report-copy']) {
      await expect(page.getByTestId(officerOnly)).toHaveCount(0);
    }
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

test('officer-read-only: match, clear, the fix and Post disabled; Copy live', async ({ page, context }) => {
  await context.grantPermissions(['clipboard-read', 'clipboard-write']);
  const stub = await openRaid(page, officerDone({ officerTools: false, discordConnected: true }));
  await expect(page.getByTestId('raid-not-this-log')).toBeDisabled();
  await expect(page.getByTestId('raid-report-post')).toBeDisabled();
  await openFixRows(page);
  const fixes = page.locator('select[data-testid^="raid-attendance-fix-"]');
  await expect(fixes).toHaveCount(allAttendance().length);
  // A done raid draws no officer row; the gate sentence shows once, above the report.
  await expect(page.getByTestId('raid-officer-row')).toHaveCount(0);
  await expect(page.locator('#officer-tools-gate-raids')).toHaveCount(1);
  await expect(page.getByTestId('guild-raids-plan-gate-report')).toBeVisible();
  expect(await fixes.evaluateAll((nodes) => nodes.every((n) => (n as HTMLSelectElement).disabled))).toBe(
    true,
  );
  await page.getByTestId('raid-report-copy').click();
  await expect(page.getByTestId('raid-report-copy')).toHaveText(attendanceCopy.report.copied);
  expect(stub.writes).toEqual([]);

  const ended = await context.newPage();
  await openRaid(ended, { viewer: 'officer', officerTools: false });
  await expect(ended.getByTestId('raid-match-open')).toBeDisabled();
});

test.describe('the roster chip (§3.6)', () => {
  test('raid mode: two windows and No raid nights yet, no footnote', async ({ page }) => {
    await stubGuildAttendance(page, { viewer: 'member', attendanceSource: 'raids' });
    await page.clock.install({ time: NOW });
    await gotoHydrated(page, `${GUILD_URL}#roster`, 'guild-roster-tab');
    const rows = page.getByTestId('guild-roster-rows');
    // The chip reads the bot's words (attendanceChip): `6 of 8 raid nights · 25 of the last 30`.
    await expect(rows).toContainText(/\d of 8 raid nights · \d+ of the last 30/);
    await expect(rows).toContainText(attendanceCopy.roster.noNights);
    // One read for the roster and the guild line (WP 4): Kraggor, who led the raid, has raids.
    const { last8, last30 } = rosterAttendance('Kraggor');
    await expect(rows).toContainText(rosterChip(last8, last30));
    await expect(page.getByTestId('guild-roster-attendance-footnote')).toHaveCount(0);
  });

  test('report mode: the nights chip and no footnote', async ({ page }) => {
    await stubGuildAttendance(page, { viewer: 'member', attendanceSource: 'reports' });
    await page.clock.install({ time: NOW });
    await gotoHydrated(page, `${GUILD_URL}#roster`, 'guild-roster-tab');
    await expect(page.getByTestId('guild-roster-rows')).toContainText(/\d of 8 raid nights/);
    await expect(page.getByTestId('guild-roster-rows')).not.toContainText('of the last 30');
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
    // The rows: the sort chips' aria-sort predates this lane (see the report).
    await expectNoAxeViolations(page, '[data-testid="guild-roster-rows"]');
    await gotoHydrated(page, `${GUILD_URL}#readiness`, 'guild-readiness-next-raid');
    await expectNoAxeViolations(page, '[data-testid="guild-readiness-tab"]');
  });
});

test.describe('every control is at least 44 px tall at 360', () => {
  test.beforeEach(({ page }, testInfo) => {
    test.skip(testInfo.project.name !== 'mobile', 'the 360 floor is a phone rule');
    return page.setViewportSize(PHONE);
  });

  test('the done page for the officer, a pull opened', async ({ page }) => {
    await openRaid(page, officerDone({ discordConnected: true }));
    await page.getByTestId('raid-pull-open').first().click();
    await expect(page.getByTestId('raid-pull-inspector')).toBeVisible();
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
