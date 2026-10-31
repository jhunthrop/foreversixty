// web/tests/e2e/export-not-read.spec.ts
// W.3: an export that carried no gear and no talents is "not read yet" on the character page,
// not seventeen "Empty" rows; the band's best in slot stays below it. And the upload form's two
// advisory warnings, which never block. Captures at 1280 go to docs/superpowers/reviews/
// lifecycle/r1 when E2E_CAPTURE is set.
import type { Page } from '@playwright/test';
import { expect, test } from './support/test';
import {
  CHARACTER,
  CHARACTER_PATH,
  CHARACTER_URL,
  GEAR,
  GEAR_ITEMS,
  routeGearIcons,
} from './support/character-fixture';
import { routeSampleReportMeta } from './support/logs-sample-fixture';

const CAPTURE_DIR = process.env.E2E_CAPTURE;

const envelope = (data: unknown, status = 200) => ({
  status,
  contentType: 'application/json',
  body: JSON.stringify({ ok: status < 400, data: status < 400 ? data : null, error: null, request_id: 'r' }),
});

const UNREAD_GEAR = {
  ...GEAR,
  gear: GEAR.gear.map((slot) => ({ ...slot, worn: null, gap: slot.pick === null ? null : 3 })),
  gear_gap_total: null,
};

async function openCharacter(page: Page, gear: unknown): Promise<void> {
  await page.route(CHARACTER_URL, (route) =>
    route.fulfill(envelope({ ...CHARACTER.data, ...(gear as object) })),
  );
  await page.route(`${CHARACTER_URL}/rating`, (route) => route.fulfill(envelope(null, 404)));
  await page.route('**/data/*/items/priest.json', (route) =>
    route.fulfill({ status: 200, contentType: 'application/json', body: JSON.stringify(GEAR_ITEMS) }),
  );
  await routeGearIcons(page);
  await page.goto(CHARACTER_PATH);
}

test.describe('character page with a gearless export', () => {
  test.use({ viewport: { width: 1280, height: 900 } });

  test('says the gear has not been read, with the band’s best in slot below', async ({ page }) => {
    await openCharacter(page, UNREAD_GEAR);
    const panel = page.getByTestId('export-not-read');
    await expect(panel).toBeVisible();
    await expect(panel).toContainText('This character’s gear has not been read yet.');
    await expect(page.getByTestId('character-gear-title')).toContainText('Best in slot for level');
    await expect(page.getByTestId('character-gear-row-head')).toContainText('Halo of Transcendence');
    // Nothing reads as a verdict on the character.
    await expect(page.getByText('Empty', { exact: true })).toHaveCount(0);
    await expect(page.getByText('No pick for your level range yet')).toHaveCount(0);
    await expect(page.getByTestId('character-gear-totals')).toHaveCount(0);
  });

  test('a character with gear still draws the paperdoll', async ({ page }) => {
    await openCharacter(page, GEAR);
    await expect(page.getByTestId('character-gear-paperdoll')).toBeVisible();
    await expect(page.getByTestId('export-not-read')).toHaveCount(0);
  });

  test('the owner sees the full wording, and the closing line says the page updates by itself', async ({
    page,
  }) => {
    await page.route('**/v1/me', (route) =>
      route.fulfill(
        envelope({
          user: { id: 1, battletag: 'F#1', email: null, role: 'user', anonymize: false },
          characters: [
            {
              key: 'us/hardcore/elyra-duskvale',
              region: 'us',
              ruleset: 'hardcore',
              name: 'Elyra Duskvale',
              class: 'Priest',
              level: 60,
              build: { source: 'addon', captured_at: '2026-10-19T21:40:00Z' },
            },
          ],
          guilds: [],
        }),
      ),
    );
    await openCharacter(page, UNREAD_GEAR);
    const line = page.getByTestId('character-in-game-planner');
    await expect(line).toContainText('This page updates by itself each time you log out');
    await expect(line.getByRole('link', { name: 'paste it on the setup page' })).toHaveAttribute(
      'href',
      '/setup',
    );
    await expect(page.getByTestId('export-not-read')).toContainText('Your gear has not been read yet.');
    await expect(page.getByTestId('export-not-read')).toContainText('The last sync from the game (');
    if (CAPTURE_DIR !== undefined) {
      await page.getByTestId('character-gear-title').scrollIntoViewIfNeeded();
      await page.screenshot({ path: `${CAPTURE_DIR}/w3-character-not-read-1280.png`, fullPage: false });
    }
  });
});

test.describe('upload warnings', () => {
  test.use({ viewport: { width: 1280, height: 900 } });

  const HEADER_ON = '10/19 21:00:00.000  COMBAT_LOG_VERSION,20,ADVANCED_LOG_ENABLED,1,BUILD_VERSION,1.15.9\n';
  const HEADER_OFF =
    '10/19 21:00:00.000  COMBAT_LOG_VERSION,20,ADVANCED_LOG_ENABLED,0,BUILD_VERSION,1.15.9\n';
  const TRASH = '10/19 21:00:01.000  SWING_DAMAGE,Player-1,"Bow",0x511\n'.repeat(40);
  const BOSS = '10/19 21:05:00.000  ENCOUNTER_START,1084,"Onyxia",9,20,249\n';

  async function openLogs(page: Page): Promise<void> {
    await page.route('**/v1/reports/recent**', (route) => route.fulfill(envelope({ rows: [] })));
    await page.route('**/v1/me', (route) =>
      route.fulfill(
        envelope({
          user: { id: 1, battletag: 'F#1', email: null, role: 'user', anonymize: false },
          characters: [],
          guilds: [],
        }),
      ),
    );
    await page.route('**/v1/devices', (route) => route.fulfill(envelope([])));
    await page.route('**/v1/reports?mine=1**', (route) =>
      route.fulfill(envelope({ rows: [], total: 0, page: 1, per_page: 100 })),
    );
    await routeSampleReportMeta(page);
    await page.goto('/logs');
    await expect(page.locator('astro-island[ssr]:has([data-testid="logs-grid"])')).toHaveCount(0);
  }

  const choose = (page: Page, text: string) =>
    page
      .getByTestId('upload-file')
      .setInputFiles({ name: 'WoWCombatLog.txt', mimeType: 'text/plain', buffer: Buffer.from(text) });

  test('a trash-only log with advanced logging off gets both warnings, and Upload stays live', async ({
    page,
  }) => {
    await openLogs(page);
    await choose(page, HEADER_OFF + TRASH);
    await expect(page.getByTestId('upload-warning-advanced')).toHaveText(
      'Advanced combat logging was off for this file: no gear, resources or positions in the report.',
    );
    await expect(page.getByTestId('upload-warning-encounters')).toHaveText(
      'No boss fights found: this report will have no ratings or rankings.',
    );
    await expect(page.getByTestId('upload-start')).toBeEnabled();
    if (CAPTURE_DIR !== undefined) {
      await page.getByTestId('upload-start').scrollIntoViewIfNeeded();
      await page.screenshot({ path: `${CAPTURE_DIR}/w3-upload-warnings-1280.png`, fullPage: false });
    }
  });

  test('a full log gets no warning', async ({ page }) => {
    await openLogs(page);
    await choose(page, HEADER_ON + TRASH + BOSS);
    await expect(page.getByTestId('upload-start')).toBeVisible();
    await expect(page.getByTestId('upload-warnings')).toHaveCount(0);
  });

  test('a log with boss fights but advanced logging off gets only the first warning', async ({ page }) => {
    await openLogs(page);
    await choose(page, HEADER_OFF + TRASH + BOSS);
    await expect(page.getByTestId('upload-warning-advanced')).toBeVisible();
    await expect(page.getByTestId('upload-warning-encounters')).toHaveCount(0);
  });
});
