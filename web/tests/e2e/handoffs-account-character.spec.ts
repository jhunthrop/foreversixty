// web/tests/e2e/handoffs-account-character.spec.ts
import { expect, test } from '@playwright/test';
import { NEEDS_EXPORT_TEXT } from '../../src/lib/handoff-copy';
import { ACTIVE_BUILD } from './support/active-build';

function envelope(data: unknown, status = 200) {
  return {
    status,
    contentType: 'application/json',
    body: JSON.stringify({ ok: status < 400, data, error: null, request_id: 'req-test' }),
  };
}

const ADDON_CODE = `FS1:${ACTIVE_BUILD}:warrior:human:0/0/0:`;

const ME = {
  user: { id: 7, battletag: 'Fixture#1234', email: null, role: 'user', anonymize: false, premium: false },
  characters: [
    { key: 'us/normal/thrallgar', region: 'us', ruleset: 'normal', name: 'Thrallgar', class: 'Warrior' },
    { key: 'us/normal/roland', region: 'us', ruleset: 'normal', name: 'Roland', class: 'Mage' },
  ],
  guilds: [],
};

test('a signed-in member sees hand-off links only for a character with an addon export', async ({ page }) => {
  await page.route('**/v1/me', (route) => route.fulfill(envelope(ME)));
  await page.route('**/v1/devices', (route) => route.fulfill(envelope({ devices: [] })));
  await page.route('**/v1/characters/us/normal/thrallgar/sim-input', (route) =>
    route.fulfill(
      envelope({
        spec: 'warrior-fury',
        gear: ADDON_CODE,
        talents: '',
        buffs: [],
        captured_at: new Date().toISOString(),
        source: 'addon',
      }),
    ),
  );
  await page.route('**/v1/characters/us/normal/roland/sim-input', (route) =>
    route.fulfill(
      envelope({
        spec: 'mage-fire',
        gear: { trinkets: [] },
        talents: '31/0/20',
        buffs: [],
        captured_at: new Date().toISOString(),
        source: 'fight',
      }),
    ),
  );

  await page.goto('/account');

  const thrallgarRow = page.getByRole('listitem').filter({ hasText: 'Thrallgar' });
  await expect(thrallgarRow.getByTestId('character-open-sim')).toHaveAttribute(
    'href',
    `/sim?code=${encodeURIComponent(ADDON_CODE)}`,
  );
  await expect(thrallgarRow.getByTestId('character-open-planner')).toHaveAttribute(
    'href',
    `/planner?code=${encodeURIComponent(ADDON_CODE)}`,
  );

  const rolandRow = page.getByRole('listitem').filter({ hasText: 'Roland' });
  await expect(rolandRow.getByTestId('character-needs-addon')).toHaveText(NEEDS_EXPORT_TEXT);
});
