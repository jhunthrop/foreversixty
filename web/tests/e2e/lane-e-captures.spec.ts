// web/tests/e2e/lane-e-captures.spec.ts
// Not a test: lane E's review captures (batch L). Skipped unless LANE_E_CAPTURE=1:
//
//   LANE_E_CAPTURE=1 npx playwright test tests/e2e/lane-e-captures.spec.ts --project=desktop --workers=1
//
// An empty guild page, an error state, a filter popover, the planner share panel and the boss
// picker, full page at 1440 and 390, into design/reviews/2026-11-11-lane-e/.
import { mkdirSync } from 'node:fs';
import path from 'node:path';
import type { Page } from '@playwright/test';
import { expect, test } from './support/test';
import { gotoHydrated } from './support/hydrated';
import { addSessionHint } from './support/account-reserve';
import { openGuildState } from './support/guild-parity';
import { GUILD_PARITY_NOW, GUILD_PARITY_ZONE } from './support/guild-parity/clock';
import { REPO_ROOT, prepareParityPage, settleParityPage } from './support/parity';
import { stubRankingsApi } from './support/rankings-fixture';
import { signedOut } from './support/selector';
import { shareBuild } from './support/planner';
import { openBis } from './support/bis-page';
import {
  CHARACTER,
  CHARACTER_PATH,
  CHARACTER_URL,
  openHistoryDisclosure,
  ownerMeAnswer,
  routeCharacterHistory,
} from './support/character-fixture';

const OUT = path.join(REPO_ROOT, 'design/reviews/2026-11-11-lane-e');
const WIDTHS = [1440, 390];

async function shoot(page: Page, name: string, width: number): Promise<void> {
  mkdirSync(OUT, { recursive: true });
  await page.waitForTimeout(300);
  await page.screenshot({ path: path.join(OUT, `${name}-${width}.png`), fullPage: true });
}

const fulfil = (body: unknown, status = 200) => ({
  status,
  contentType: 'application/json',
  body: JSON.stringify(body),
});

test.describe('Lane E captures', () => {
  test.skip(() => process.env.LANE_E_CAPTURE !== '1', 'a capture tool, run on demand');
  test.skip(({ isMobile }) => isMobile, 'the widths are set per capture');

  for (const width of WIDTHS) {
    test.describe(`${width}`, () => {
      test.use({ viewport: { width, height: 1000 }, timezoneId: GUILD_PARITY_ZONE });

      test('guild with no raid planned', async ({ page }) => {
        await page.clock.setFixedTime(GUILD_PARITY_NOW);
        await prepareParityPage(page);
        await openGuildState(page, 'no-raid');
        await settleParityPage(page);
        await shoot(page, 'guild-empty', width);
      });

      test('account load error', async ({ page, baseURL }) => {
        await addSessionHint(page, baseURL);
        await page.route('**/v1/me', (route) =>
          route.fulfill(fulfil({ ok: false, data: null, error: { message: 'oops' }, request_id: 'r' }, 500)),
        );
        await page.goto('/account');
        await expect(page.getByTestId('account-load-error')).toBeVisible();
        await shoot(page, 'error-account', width);
      });

      test('rankings filters', async ({ page }) => {
        await signedOut(page);
        await stubRankingsApi(page);
        await gotoHydrated(page, '/rankings/warden-kelthas', 'rankings-heading');
        if (width >= 768) {
          await page.getByTestId('rankings-filters-button').click();
          await expect(page.getByTestId('rankings-filters-popover')).toBeVisible();
        }
        await shoot(page, 'filter-popover-rankings', width);
      });

      test('bis sources', async ({ page }) => {
        await openBis(page, '/bis/hunter/marksmanship?band=60&view=slot');
        const button = page.getByTestId('bis-sources-button');
        if (width >= 768 && (await button.count()) > 0) {
          await button.click();
          await expect(page.getByTestId('bis-sources-popover')).toBeVisible();
        }
        await shoot(page, 'filter-popover-bis', width);
      });

      test('planner share panel', async ({ page }) => {
        await page.route('**/v1/builds', (route) =>
          route.fulfill(
            fulfil(
              {
                ok: true,
                data: { id: 'k7x2qm4a', url: 'https://foreversixty.gg/b/k7x2qm4a' },
                error: null,
                request_id: 'r',
              },
              201,
            ),
          ),
        );
        await page.goto('/planner');
        for (let i = 0; i < 3; i += 1) await page.getByTestId('talent-1001').click();
        await page.getByLabel('Title').fill('Arms leveling');
        await shareBuild(page);
        await expect(page.getByTestId('share-link')).toBeVisible();
        await shoot(page, 'planner-share', width);
      });

      test('boss picker', async ({ page }) => {
        await page.route(CHARACTER_URL, (route) =>
          route.fulfill(
            fulfil({
              ...CHARACTER,
              data: {
                ...CHARACTER.data,
                score: 1830,
                tier: 'epic',
                as_of: '2026-11-04',
                components: { gear: 412, experience: 640, parses: 510, attendance: 180, discipline: 88 },
                experience: { dungeons: 4, raid_kills: 12, alts: 1 },
              },
            }),
          ),
        );
        await page.route(`${CHARACTER_URL}/rating`, (route) => route.fulfill({ status: 404, body: '{}' }));
        const me = ownerMeAnswer() as { data: { entitlements?: Record<string, unknown> } };
        me.data.entitlements = { ...(me.data.entitlements ?? {}), history: true };
        await page.route('**/v1/me', (route) => route.fulfill(fulfil(me)));
        await routeCharacterHistory(page);
        await page.goto(CHARACTER_PATH);
        await openHistoryDisclosure(page);
        await expect(page.getByTestId('history-picker')).toBeVisible();
        await page.getByTestId('history-picker').scrollIntoViewIfNeeded();
        await shoot(page, 'boss-picker', width);
      });
    });
  }
});
