// web/tests/e2e/account-delete.spec.ts
// Self-service account deletion (DELETE /v1/me, operation deleteAccount): /account's last rail
// Section opens the confirm dialog; its danger button stays off until the BattleTag is typed; a
// refusal the server explains (409 guild_obligation) is shown verbatim and nothing navigates; a
// refused confirm (400 confirm_mismatch) is the line under the field; the happy path sends what
// was typed, lands on home and says the account is gone once. The account is Hunthrop#1894
// (fixtures/home/me-34.json), every answer routed here.
import { readFileSync } from 'node:fs';
import path from 'node:path';
import type { Page, Route } from '@playwright/test';
import { accountPageCopy } from '../../src/lib/account/account-page-copy';
import { expect, test } from './support/test';

const FIXTURES = path.join(import.meta.dirname, 'fixtures');
const ME = JSON.parse(readFileSync(path.join(FIXTURES, 'home/me-34.json'), 'utf8')) as {
  user: { battletag: string };
};
const BATTLETAG = ME.user.battletag;
const CONTRACT = path.join(import.meta.dirname, '../../../docs/contracts/fixtures/app');
const contract = (name: string): { error: { message: string } } =>
  JSON.parse(readFileSync(path.join(CONTRACT, `${name}.json`), 'utf8')) as { error: { message: string } };

const envelope = (data: unknown) => ({
  status: 200,
  contentType: 'application/json',
  body: JSON.stringify({ ok: true, data, error: null, request_id: 'delete' }),
});

/** One delete attempt's answer: a status and the envelope the API sends with it. */
type DeleteAnswer = { status: number; body: unknown };

/**
 * Routes /account for the account and DELETE /v1/me with `answers`, one per attempt. Returns the
 * bodies the page sent. A 200 clears the cookies, as the server does, and /v1/me answers 401 after.
 */
async function routeAccount(page: Page, answers: DeleteAnswer[]): Promise<unknown[]> {
  const sent: unknown[] = [];
  let deleted = false;
  await page.context().addCookies([{ name: 'fs_csrf', value: 'token', domain: 'localhost', path: '/' }]);
  await page.route('**/v1/me', async (route: Route) => {
    const request = route.request();
    if (request.method() === 'DELETE') {
      sent.push(request.postDataJSON());
      const answer = answers.shift() ?? { status: 500, body: {} };
      if (answer.status === 200) {
        deleted = true;
        await page.context().clearCookies();
      }
      await route.fulfill({
        status: answer.status,
        contentType: 'application/json',
        body: JSON.stringify(answer.body),
      });
      return;
    }
    if (deleted) {
      await route.fulfill({
        status: 401,
        json: { ok: false, data: null, error: { code: 'unauthorized', message: 'Sign in first.' } },
      });
      return;
    }
    await route.fulfill(envelope(ME));
  });
  await page.route('**/v1/devices', (route) => route.fulfill(envelope([])));
  return sent;
}

async function openDialog(page: Page): Promise<void> {
  await page.goto('/account');
  await expect(page.getByTestId('account-grid')).toBeVisible();
  await page.getByTestId('account-delete-open').click();
  await expect(page.getByRole('dialog', { name: accountPageCopy.deleteDialog.title })).toBeVisible();
}

const dialog = (page: Page) => page.getByTestId('account-delete-dialog');
const field = (page: Page) => page.getByLabel(accountPageCopy.deleteDialog.confirmLabel(BATTLETAG));
const submit = (page: Page) => page.getByTestId('account-delete-submit');

test.describe('delete account', () => {
  test.skip(({ isMobile }) => isMobile, 'one flow; the phone sheet is the parity spec');

  test('the dialog lists what goes and what stays, and opens and closes with focus kept', async ({
    page,
  }) => {
    await routeAccount(page, []);
    await openDialog(page);
    const copy = accountPageCopy.deleteDialog;

    await expect(dialog(page).getByRole('heading', { level: 2 })).toBeFocused();
    await expect(dialog(page)).toContainText(copy.intro);
    await expect(page.getByTestId('account-delete-deleted').getByRole('listitem')).toHaveText(
      copy.deleted.map((line) => line.text),
    );
    await expect(page.getByTestId('account-delete-kept').getByRole('listitem')).toHaveText(
      copy.kept.map((line) => line.text),
    );
    await expect(dialog(page).getByRole('link', { name: 'refund policy' })).toHaveAttribute(
      'href',
      '/refunds',
    );
    await expect(field(page)).toHaveAttribute('autocomplete', 'off');
    await expect(field(page)).toHaveAttribute('spellcheck', 'false');
    await expect(field(page)).toHaveAttribute('autocapitalize', 'off');

    await page.keyboard.press('Escape');
    await expect(dialog(page)).toBeHidden();
    await expect(page.getByTestId('account-delete-open')).toBeFocused();

    await page.getByTestId('account-delete-open').click();
    await page.getByTestId('account-delete-cancel').click();
    await expect(dialog(page)).toBeHidden();
    await expect(page.getByTestId('account-delete-open')).toBeFocused();
  });

  test('the button stays off until the BattleTag is typed, in any case', async ({ page }) => {
    await routeAccount(page, []);
    await openDialog(page);

    await expect(submit(page)).toBeDisabled();
    await field(page).fill(BATTLETAG.slice(0, -1));
    await expect(submit(page)).toBeDisabled();
    await field(page).fill(`  ${BATTLETAG.toUpperCase()} `);
    await expect(submit(page)).toBeEnabled();
    await field(page).fill('Delete');
    await expect(submit(page)).toBeEnabled();
    await field(page).fill('');
    await expect(submit(page)).toBeDisabled();
  });

  test('a guild obligation is shown verbatim and nothing navigates', async ({ page }) => {
    const refusal = contract('me.delete.error-guild-obligation');
    const sent = await routeAccount(page, [{ status: 409, body: refusal }]);
    await openDialog(page);

    await field(page).fill(BATTLETAG);
    await submit(page).click();

    await expect(page.getByTestId('account-delete-refusal')).toHaveText(refusal.error.message);
    await expect(page.getByTestId('account-delete-refusal')).toHaveAttribute('role', 'alert');
    expect(sent).toEqual([{ confirm: BATTLETAG }]);
    await expect(page).toHaveURL(/\/account$/);
    await expect(dialog(page)).toBeVisible();
    await expect(submit(page)).toBeEnabled();
  });

  test('a refused confirm is the line under the field', async ({ page }) => {
    await routeAccount(page, [{ status: 400, body: contract('me.delete.error-confirm-mismatch') }]);
    await openDialog(page);

    await field(page).fill('delete');
    await submit(page).click();

    const line = page.getByTestId('account-delete-mismatch');
    await expect(line).toHaveText(accountPageCopy.deleteDialog.mismatch(BATTLETAG));
    await expect(field(page)).toHaveAttribute('aria-invalid', 'true');
    await expect(field(page)).toBeFocused();
    await expect(page).toHaveURL(/\/account$/);
  });

  test('a deletion sends what was typed, lands on home and says so once', async ({ page }) => {
    const sent = await routeAccount(page, [
      { status: 200, body: JSON.parse(readFileSync(path.join(CONTRACT, 'me.delete.json'), 'utf8')) },
    ]);
    await openDialog(page);

    await field(page).fill(` ${BATTLETAG} `);
    await submit(page).click();

    await expect(page).toHaveURL(/\/$/);
    expect(sent).toEqual([{ confirm: ` ${BATTLETAG} ` }]);
    const notice = page.getByTestId('home-deleted-notice');
    await expect(notice).toHaveText(accountPageCopy.deletedNotice);
    // Read once: the flag left the address bar, so a reload says nothing.
    expect(new URL(page.url()).search).toBe('');
    await page.reload();
    await expect(page.getByTestId('home-signed-out')).toBeVisible();
    // The notice's island has mounted (Astro drops `ssr` once it has) and drew nothing.
    const island = page.locator('astro-island[component-url*="HomeDeletedNotice"]');
    await expect(island).toHaveCount(1);
    await expect(island).not.toHaveAttribute('ssr');
    await expect(notice).toHaveCount(0);
  });
});
