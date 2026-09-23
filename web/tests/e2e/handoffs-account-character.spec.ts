// web/tests/e2e/handoffs-account-character.spec.ts
// The account Characters list's per-row action is built from `MeCharacter.build` alone, with
// no per-row `sim-input` fetch (spec 2026-09-22 §3.4, CharacterRowLink.svelte's own header
// comment) -- unlike the hero band and the character page, which still use
// CharacterHandoffLinks and its own real fetch. A row therefore offers "Open in simulator"
// (via the fetch-free armory-ref href) exactly when the character carries a `build`, and
// "No export yet" otherwise; there is no per-row "Open in planner" link.
import { expect, test } from '@playwright/test';
import { NEEDS_EXPORT_TEXT } from '../../src/lib/handoff-copy';

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

test('a signed-in member sees hand-off links only for a character with a build', async ({ page }) => {
  await page.route('**/v1/me', (route) => route.fulfill(envelope(ME)));
  await page.route('**/v1/devices', (route) => route.fulfill(envelope([])));

  await page.goto('/account');

  const thrallgarRow = page.getByRole('listitem').filter({ hasText: 'Thrallgar' });
  await expect(thrallgarRow.getByTestId('character-open-sim')).toHaveAttribute(
    'href',
    '/sim?source=armory&ref=us%2Fnormal%2Fthrallgar',
  );

  const rolandRow = page.getByRole('listitem').filter({ hasText: 'Roland' });
  await expect(rolandRow.getByTestId('character-needs-addon')).toHaveText(NEEDS_EXPORT_TEXT);
  await expect(rolandRow.getByTestId('character-open-sim')).toHaveCount(0);
});
