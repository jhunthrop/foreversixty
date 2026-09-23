// web/tests/e2e/bnet-hub.spec.ts
// Task 5 (spec 2026-09-22 §3.1): a fresh Battle.net sign-in lands on /account?signed_in=1,
// picks the main character, points the site at it, and shows the hero band with a Logs
// link and a one-line banner. This depends on meBnetFixture's guilded character (Thoradin)
// carrying a `build` field so mainCharacter picks it deterministically over the unguilded,
// level-less Elyra -- that field lands in a later, fixture-touching task, so this spec is
// created now but left failing until then (that task runs it as its own verification step).
import { expect, test } from '@playwright/test';
import { meBnetFixture } from '../../src/fixtures/me-bnet';

const fulfil = (body: unknown, status = 200) => ({
  status,
  contentType: 'application/json',
  body: JSON.stringify(body),
});

test('signing in lands on the hub with the main character in the hero band', async ({ page }) => {
  await page.route('**/v1/me', (route) =>
    route.fulfill(fulfil({ ok: true, data: meBnetFixture, error: null, request_id: 'r' })),
  );
  await page.route('**/v1/devices', (route) =>
    route.fulfill(fulfil({ ok: true, data: [], error: null, request_id: 'r' })),
  );
  await page.route('**/v1/reports**', (route) =>
    route.fulfill(
      fulfil({ ok: true, data: { rows: [], total: 0, page: 1, per_page: 20 }, error: null, request_id: 'r' }),
    ),
  );
  await page.route('**/v1/characters/**/rating', (route) =>
    route.fulfill(fulfil({ ok: false, data: null, error: null, request_id: 'r' }, 404)),
  );
  await page.route('**/v1/characters/us/pvp/thoradin/sim-input', (route) =>
    route.fulfill(
      fulfil({
        ok: true,
        data: {
          spec: 'fury',
          gear: 'FS1:1.60.1.69893:warrior:orc:1a/0/0:',
          talents: '',
          buffs: [],
          captured_at: '2026-09-21T03:14:00Z',
          source: 'blizzard',
        },
        error: null,
        request_id: 'r',
      }),
    ),
  );

  await page.goto('/account?signed_in=1');

  await expect(page.getByTestId('account-signed-in-banner')).toContainText(
    'Thoradin is your current character',
  );
  await expect.poll(() => new URL(page.url()).search).toBe('');
  await expect(page.getByTestId('account-hero')).toContainText('Thoradin');
  await expect(page.getByTestId('character-open-sim')).toHaveAttribute('href', /^\/sim\?code=/);

  await page.getByTestId('character-open-sim').click();
  await expect(page).toHaveURL(/\/sim\?code=/);
});
