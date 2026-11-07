// web/tests/e2e/bnet-hub.spec.ts
// Task 5 (spec 2026-09-22 §3.1): a fresh Battle.net sign-in lands on /account?signed_in=1,
// picks the main character, points the site at it, and shows the hero band with a Logs
// link and a one-line banner. This depends on meBnetFixture's guilded character (Thoradin)
// carrying a `build` field so mainCharacter picks it deterministically over the unguilded,
// level-less Elyra -- that field lands in a later, fixture-touching task, so this spec is
// created now but left failing until then (that task runs it as its own verification step).
import { expect, test } from './support/test';
import { meBnetFixture } from '../../src/fixtures/me-bnet';

const fulfil = (body: unknown, status = 200) => ({
  status,
  contentType: 'application/json',
  body: JSON.stringify(body),
});

test("signing in lands on the account with the main character as the head's answer", async ({ page }) => {
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
  // The head's answer is the main character (premium spec 2026-11-08 7.1): Open goes to its page.
  const answer = page.getByTestId('account-head');
  await expect(answer).toContainText('Thoradin');
  const open = answer.getByTestId('setup-answer-open');
  await expect(open).toHaveAttribute('href', '/character/us/pvp/thoradin');
});
