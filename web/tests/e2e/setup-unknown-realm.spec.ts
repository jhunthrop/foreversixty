// web/tests/e2e/setup-unknown-realm.spec.ts
// The unknown-realm question on /setup (api/internal/bnetimport/unknown_realm.go): a Battle.net
// character on a realm Blizzard has not typed is imported under the ruleset `unknown`, and the
// plate asks its realm type with the four rulesets as buttons. A tap posts the answer, the account
// is read again and the question is gone. Captures for the persona reviews land in
// test-results/unknown-realm (design/reviews/2026-11-13-unknown-realm).
import type { Page } from '@playwright/test';
import { REALM_ANSWER_PATH, unknownRealmCopy } from '../../src/lib/account/unknown-realm';
import { meBnet, openSetupState } from './support/setup-page';
import { expect, test } from './support/test';

const WAITING = {
  key: 'us/unknown/newly-rolled',
  region: 'us',
  ruleset: 'unknown',
  name: 'Newly Rolled',
  class: 'rogue',
  level: 12,
  faction: 'alliance',
  race: 'Gnome',
  source: 'bnet',
};

const envelope = (data: unknown) => ({
  status: 200,
  contentType: 'application/json',
  body: JSON.stringify({ ok: true, data, error: null, request_id: 'unknown-realm' }),
});

/** /setup for an account whose newest character waits for its realm type; the answer moves it. */
async function openWaiting(page: Page): Promise<{ answers: unknown[] }> {
  const answers: unknown[] = [];
  const fixture = meBnet();
  // A neutral BattleTag in the review captures.
  const base = { ...fixture, user: { ...(fixture.user as Record<string, unknown>), battletag: 'Player#1001' } };
  let account: Record<string, unknown> = { ...base, characters: [WAITING, ...base.characters] };
  await page.route('**/v1/me', (route) => route.fulfill(envelope(account)));
  await page.route(`**${REALM_ANSWER_PATH}`, async (route) => {
    answers.push(route.request().postDataJSON());
    account = { ...base, characters: [{ ...WAITING, key: 'us/pvp/newly-rolled', ruleset: 'pvp' }, ...base.characters] };
    await route.fulfill(envelope({ key: 'us/pvp/newly-rolled', ruleset: 'pvp' }));
  });
  await openSetupState(page, 'bnet');
  // openSetupState routes its own /v1/me after ours; ours must answer, so route it again.
  await page.unroute('**/v1/me');
  await page.route('**/v1/me', (route) => route.fulfill(envelope(account)));
  await page.reload();
  await expect(page.getByTestId('setup-steps')).toBeVisible();
  return { answers };
}

async function shoot(page: Page, name: string): Promise<void> {
  for (const width of [1440, 2000, 390]) {
    await page.setViewportSize({ width, height: width === 390 ? 844 : 1100 });
    await page.screenshot({ path: `test-results/unknown-realm/${name}-${width}.png`, fullPage: true });
  }
  await page.setViewportSize({ width: 1440, height: 1100 });
  const question = page.getByTestId('unknown-realm');
  if ((await question.count()) > 0) await question.screenshot({ path: `test-results/unknown-realm/${name}-close.png` });
}

test('asks the realm type of a character on an unknown realm, and an answer clears the question', async ({
  page,
  isMobile,
}) => {
  test.setTimeout(120_000);
  const { answers } = await openWaiting(page);
  const question = page.getByTestId('unknown-realm');
  await expect(question.getByRole('heading')).toHaveText(
    `${unknownRealmCopy.questionLead} Newly Rolled ${unknownRealmCopy.questionTail}`,
  );
  for (const label of ['Normal', 'PvP', 'Roleplay', 'Hardcore']) {
    await expect(question.getByRole('button', { name: label })).toBeVisible();
  }
  if (!isMobile) await shoot(page, 'question');

  await question.getByRole('button', { name: 'PvP' }).click();
  await expect(page.getByTestId('unknown-realm')).toHaveCount(0);
  expect(answers).toEqual([{ key: 'us/unknown/newly-rolled', ruleset: 'pvp' }]);
  if (!isMobile) await shoot(page, 'answered');
});

test('a refused answer says why in the wipe colour and keeps the question', async ({ page }) => {
  await openWaiting(page);
  await page.unroute(`**${REALM_ANSWER_PATH}`);
  await page.route(`**${REALM_ANSWER_PATH}`, (route) =>
    route.fulfill({
      status: 409,
      contentType: 'application/json',
      body: JSON.stringify({
        ok: false,
        data: null,
        error: { code: 'taken', message: 'another account already has a character of that name on that realm type' },
        request_id: 'unknown-realm',
      }),
    }),
  );
  const question = page.getByTestId('unknown-realm');
  await question.getByRole('button', { name: 'Roleplay' }).click();
  await expect(question.getByRole('alert')).toContainText('another account already has a character');
  await expect(question.getByRole('button', { name: 'Normal' })).toBeEnabled();
  await page.getByTestId('unknown-realm').screenshot({ path: 'test-results/unknown-realm/refused-close.png' });
});
