// web/tests/e2e/handoffs-account-character.spec.ts
// The account Characters Plate's rows (premium spec 2026-11-08 7.2) are drawn from `/v1/me` alone,
// with no per-row `sim-input` fetch: each row says where its build came from (`Addon`,
// `Battle.net`, or `No gear yet` when it has none) and opens the character's page, where the
// simulator, planner and paste hand-offs live.
import { expect, test } from './support/test';

function envelope(data: unknown, status = 200) {
  return {
    status,
    contentType: 'application/json',
    body: JSON.stringify({ ok: status < 400, data, error: null, request_id: 'req-test' }),
  };
}

const ME = {
  user: { id: 7, battletag: 'Fixture#1234', email: null, role: 'user', anonymize: false, premium: false },
  characters: [
    {
      key: 'us/normal/thrallgar',
      region: 'us',
      ruleset: 'normal',
      name: 'Thrallgar',
      class: 'Warrior',
      build: { source: 'addon', captured_at: '2026-09-21T03:14:00Z' },
    },
    { key: 'us/normal/roland', region: 'us', ruleset: 'normal', name: 'Roland', class: 'Mage' },
  ],
  guilds: [],
};

test('each row names its build source and opens the character page', async ({ page }) => {
  await page.route('**/v1/me', (route) => route.fulfill(envelope(ME)));
  await page.route('**/v1/devices', (route) => route.fulfill(envelope([])));
  const simInputs: string[] = [];
  await page.route('**/v1/characters/**', (route) => {
    simInputs.push(route.request().url());
    return route.fulfill(envelope(null, 404));
  });

  await page.goto('/account');

  const rows = page.getByTestId('account-character-row');
  const thrallgarRow = rows.filter({ hasText: 'Thrallgar' });
  await expect(thrallgarRow.getByTestId('character-source')).toHaveText('Addon');
  await expect(thrallgarRow.getByTestId('character-open')).toHaveAttribute(
    'href',
    '/character/us/normal/thrallgar',
  );

  const rolandRow = rows.filter({ hasText: 'Roland' });
  await expect(rolandRow.getByTestId('character-source')).toHaveText('No gear yet');
  await expect(rolandRow.getByTestId('character-open')).toHaveAttribute(
    'href',
    '/character/us/normal/roland',
  );
  expect(simInputs.filter((url) => url.includes('sim-input'))).toEqual([]);
});
