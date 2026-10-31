// web/tests/e2e/logs.spec.ts
// The logs landing (design/specs/2026-10-11-logs.md): signed in, the shared character header
// leads with the character's latest kill and links the newest report when it is newer; the
// companion sits above Upload only when a device is paired; the outline Upload button appears
// only after a file is chosen; the header selector moves the band in place. Signed out, the
// shared band stands where the header would and the page has no sign-in button. On a phone the
// reading column comes first and the header's action opens the newest report.
import type { Page } from '@playwright/test';
import { expect, test } from './support/test';
import { routeSampleReportMeta } from './support/logs-sample-fixture';
import {
  OBNOXIOUS,
  armoryPointer,
  keyOf,
  signInWith,
  signedOut,
  storePointer,
  type FixtureCharacter,
} from './support/selector';

const DESKTOP = { width: 1440, height: 900 };
const PHONE = { width: 390, height: 844 };
const DAY_MS = 86_400_000;
const NEWEST_REPORT_ID = 'newestreport1';
const HUNTER: FixtureCharacter = {
  name: 'Quickshot',
  class: 'Hunter',
  spec: 'Marksmanship',
  level: 33,
  faction: 'horde',
  buildDaysAgo: 1,
};

const envelope = (data: unknown, status = 200) => ({
  status,
  contentType: 'application/json',
  body: JSON.stringify({ ok: status < 400, data, error: null, request_id: 'e2e' }),
});

const daysAgo = (days: number): string => new Date(Date.now() - days * DAY_MS - 60_000).toISOString();

function fightRow(overrides: Record<string, unknown>): Record<string, unknown> {
  return {
    encounter: 'Onyxia',
    encounter_id: 1084,
    difficulty: 1,
    metric: 'dps',
    value: 612.4,
    fought_at: daysAgo(3),
    report_id: 'killreport01',
    fight_index: 4,
    execution_score: null,
    ...overrides,
  };
}

async function routeCharacter(page: Page, character: FixtureCharacter, history: unknown[]): Promise<void> {
  const slug = keyOf(character).split('/').pop() as string;
  await page.route(`**/v1/characters/us/normal/${slug}`, (route) =>
    route.fulfill(
      envelope({
        player: { key: keyOf(character), name: character.name, class: character.class },
        best: [],
        history,
        builds_seen: [],
      }),
    ),
  );
}

async function routeReports(page: Page, createdDaysAgo: number | null): Promise<void> {
  const rows =
    createdDaysAgo === null
      ? []
      : [
          {
            id: NEWEST_REPORT_ID,
            title: 'Onyxia, wipes',
            zone: "Onyxia's Lair",
            status: 'complete',
            visibility: 'public',
            created_at: daysAgo(createdDaysAgo),
            fight_count: 6,
            kill_count: 0,
          },
        ];
  await page.route('**/v1/reports?mine=1**', (route) =>
    route.fulfill(envelope({ rows, total: rows.length, page: 1, per_page: 100 })),
  );
  await page.route('**/v1/reports/recent**', (route) => route.fulfill(envelope({ rows: [] })));
  await routeSampleReportMeta(page);
}

async function routeDevices(page: Page, paired: boolean): Promise<void> {
  const devices = paired
    ? [
        {
          id: 'dev1',
          name: 'MacBook Pro',
          platform: 'macOS',
          created_at: daysAgo(30),
          last_seen_at: new Date(Date.now() - 4 * 60_000).toISOString(),
        },
      ]
    : [];
  await page.route('**/v1/devices', (route) => route.fulfill(envelope(devices)));
}

async function openSignedIn(
  page: Page,
  options: { paired?: boolean; reportDaysAgo?: number | null; characters?: FixtureCharacter[] } = {},
): Promise<void> {
  const characters = options.characters ?? [OBNOXIOUS];
  await signInWith(page, characters);
  await storePointer(page, armoryPointer(characters[0]!));
  await routeDevices(page, options.paired ?? false);
  await routeReports(page, options.reportDaysAgo === undefined ? 1 : options.reportDaysAgo);
  await routeCharacter(page, OBNOXIOUS, [fightRow({})]);
  await routeCharacter(page, HUNTER, [
    fightRow({
      encounter: 'Lucifron',
      metric: 'dps',
      value: 1250.4,
      fought_at: daysAgo(20),
      report_id: 'hunterrep01',
    }),
  ]);
  await page.goto('/logs');
  await expect(page.getByTestId('character-header-name')).toHaveText(characters[0]!.name);
}

const header = (page: Page) => page.getByTestId('character-header');
const top = async (page: Page, testId: string): Promise<number> =>
  ((await page.locator(testId).first().boundingBox())?.y ?? Number.NaN) as number;

test.describe('signed in, desktop', () => {
  test.use({ viewport: DESKTOP });

  test('the band shows the latest kill and links the report that is newer than it', async ({ page }) => {
    await openSignedIn(page);
    await expect(page.getByTestId('character-header-answer-label')).toHaveText('Latest kill');
    await expect(page.getByTestId('character-header-figure-1')).toHaveText('612.4 DPS');
    const note = page.getByTestId('character-header-note');
    await expect(note).toHaveText('Onyxia kill · 3 days ago · Open the fight →');
    await expect(note.getByRole('link')).toHaveAttribute('href', '/reports/killreport01?fight=4');
    const second = page.getByTestId('character-header-second-note');
    await expect(second).toHaveText('Last report 1 day ago');
    await expect(second).toHaveAttribute('href', `/reports/${NEWEST_REPORT_ID}`);
  });

  test('no last-report line when the newest report is older than the kill', async ({ page }) => {
    await openSignedIn(page, { reportDaysAgo: 6 });
    await expect(page.getByTestId('character-header-note')).toBeVisible();
    await expect(page.getByTestId('character-header-second-note')).toHaveCount(0);
  });

  test('a character with no logs reads "No logs yet" and the action stays Upload a log', async ({ page }) => {
    await signInWith(page, [OBNOXIOUS]);
    await routeDevices(page, false);
    await routeReports(page, null);
    await routeCharacter(page, OBNOXIOUS, []);
    await page.goto('/logs');
    await expect(page.getByTestId('character-header-sentence')).toHaveText(
      'No logs yet. Upload one, or log live with the companion.',
    );
    await expect(page.getByTestId('character-header-action')).toHaveText('Upload a log');
  });

  test('the header action opens the file picker, and the page spine is gone', async ({ page }) => {
    await openSignedIn(page);
    const action = page.getByTestId('character-header-action');
    await expect(action).toHaveText('Upload a log');
    const chooser = page.waitForEvent('filechooser');
    await action.click();
    await chooser;
    await expect(page.getByTestId('current-character-bar')).toHaveCount(0);
  });

  test('report rows read zone and bosses with a relative age, and no Sample pill remains', async ({
    page,
  }) => {
    await openSignedIn(page);
    await expect(page.getByTestId('report-row-summary').first()).toHaveText("Onyxia's Lair · 6 fights");
    await expect(page.getByTestId('report-row-when').first()).toHaveText('1 day ago');
    await expect(page.locator('.pill-sample')).toHaveCount(0);
    await expect(page.getByTestId('upload-start')).toHaveCount(0);
  });

  test('Upload comes first, then the companion, when no device is paired', async ({ page }) => {
    await openSignedIn(page, { paired: false });
    await expect(page.getByTestId('companion-status')).toHaveCount(0);
    expect(await top(page, '#upload')).toBeLessThan(await top(page, '#companion'));
  });

  test('the companion sits above Upload when a device is paired', async ({ page }) => {
    await openSignedIn(page, { paired: true });
    await expect(page.getByTestId('companion-status-line')).toContainText('MacBook Pro');
    expect(await top(page, '#companion')).toBeLessThan(await top(page, '#upload'));
  });

  test('the outline Upload button appears only after a file is chosen', async ({ page }) => {
    await openSignedIn(page);
    await expect(page.getByTestId('upload-start')).toHaveCount(0);
    await page.getByTestId('upload-file').setInputFiles({
      name: 'WoWCombatLog.txt',
      mimeType: 'text/plain',
      buffer: Buffer.from('9/26 20:10:00.000  COMBAT_LOG_VERSION,16\n'),
    });
    await expect(page.getByTestId('upload-start')).toBeVisible();
    await expect(page.getByTestId('upload-start')).toHaveText('Upload');
  });

  test('a header switch updates the band in place', async ({ page }) => {
    await openSignedIn(page, { characters: [OBNOXIOUS, HUNTER] });
    await expect(page.getByTestId('character-header-figure-1')).toHaveText('612.4 DPS');
    await page.evaluate(() => {
      (window as unknown as { __kept: boolean }).__kept = true;
    });
    await page.getByTestId('character-selector').click();
    await page
      .getByTestId('selector-panel')
      .getByTestId(`selector-row-${keyOf(HUNTER)}`)
      .getByRole('button')
      .first()
      .click();
    await expect(page.getByTestId('character-header-name')).toHaveText(HUNTER.name);
    await expect(page.getByTestId('character-header-figure-1')).toHaveText('1,250 DPS');
    await expect(page.getByTestId('character-header-note')).toHaveText(
      `Lucifron kill · ${await hunterDate()} · Open the fight →`,
    );
    expect(await page.evaluate(() => (window as unknown as { __kept?: boolean }).__kept)).toBe(true);
  });

  test("the 1024 note never wraps its arrow away from the link's words", async ({ page }) => {
    await page.setViewportSize({ width: 1024, height: 900 });
    await openSignedIn(page);
    const link = page.getByTestId('character-header-note').getByRole('link');
    const box = await link.boundingBox();
    expect(box?.height ?? Infinity).toBeLessThan(24);
  });
});

/** The hunter's kill date, as the page prints it (UTC "Mon D") for a kill 20 days ago. */
async function hunterDate(): Promise<string> {
  const date = new Date(Date.now() - 20 * DAY_MS - 60_000);
  const month = ['Jan', 'Feb', 'Mar', 'Apr', 'May', 'Jun', 'Jul', 'Aug', 'Sep', 'Oct', 'Nov', 'Dec'][
    date.getUTCMonth()
  ];
  return `${month} ${date.getUTCDate()}`;
}

test.describe('signed out, desktop', () => {
  test.use({ viewport: DESKTOP });

  test.beforeEach(async ({ page }) => {
    await signedOut(page);
    await routeReports(page, null);
    await page.goto('/logs');
  });

  test('signed out the header slot is empty, and the one sign-in button is in Your reports', async ({
    page,
  }) => {
    await expect(page.getByTestId('character-header-signed-out')).toHaveCount(0);
    await expect(header(page)).toHaveCount(0);
    const main = page.locator('#main');
    await expect(main.getByRole('link', { name: /sign in/i })).toHaveCount(1);
    await expect(page.getByTestId('reports-sign-in')).toBeVisible();
    await expect(main.getByRole('button', { name: /sign in/i })).toHaveCount(0);
    await expect(page.getByTestId('upload-signed-out')).toBeVisible();
    // The companion card is a sentence and a button, with no pairing block (W.9).
    await expect(page.getByTestId('pairing-signed-out')).toHaveCount(0);
    await expect(page.getByTestId('logs-companion-link')).toBeVisible();
  });

  test('Upload comes first, then the companion, and the sample is one row', async ({ page }) => {
    expect(await top(page, '#upload')).toBeLessThan(await top(page, '#companion'));
    await expect(page.getByTestId('recent-reports-sample').getByTestId('report-row')).toHaveCount(1);
  });
});

test.describe('signed in, phone', () => {
  test.use({ viewport: PHONE });

  test('the action opens the latest report, and reports come before Upload', async ({ page }) => {
    await openSignedIn(page);
    const action = page.getByTestId('character-header-action');
    await expect(action).toHaveText('Open your latest report →');
    await expect(action).toHaveAttribute('href', `/reports/${NEWEST_REPORT_ID}`);
    await expect(page.getByTestId('upload-start')).toHaveCount(0);
    const reports = await top(page, '#reports');
    const upload = await top(page, '#upload');
    expect(reports).toBeLessThan(upload);
  });
});
