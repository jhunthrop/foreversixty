// web/tests/e2e/first-run-signed-in-empty.spec.ts
// A signed-in account with no characters reads as signed in (header, at desktop and phone
// widths), and the first-run copy leads with the addon and the companion (/account, /setup,
// /logs). Runs in both the desktop and the mobile project.
import { apiUnreachable, expect, test } from './support/test';
import { signInWith } from './support/selector';

// Reads this spec does not answer go unanswered, as the API out of reach (support/test.ts).
test.beforeEach(({ page }) => apiUnreachable(page, '**/v1/devices'));

const envelope = (data: unknown) => ({
  status: 200,
  contentType: 'application/json',
  body: JSON.stringify({ ok: true, data, error: null, request_id: 'r' }),
});

test.describe('the header of a signed-in account with no characters', () => {
  test('shows the battletag and "No characters yet", never "Sign in"', async ({ page }) => {
    await signInWith(page, []);
    await page.goto('/guides');
    const control = page.getByTestId('character-selector');
    await expect(control).toBeVisible();
    await expect(control).toContainText('Fixture#1');
    await expect(control).toContainText('No characters yet');
    await expect(control).not.toContainText('Sign in');
    await expect(control).not.toContainText('or paste an export');
    expect(
      await page.evaluate(() => document.documentElement.scrollWidth <= document.documentElement.clientWidth),
    ).toBe(true);
  });

  test('opens the signed-in empty panel with Get set up and Your account', async ({ page }) => {
    await signInWith(page, []);
    await page.goto('/guides');
    await page.getByTestId('character-selector').click();
    const panel = page.getByTestId('selector-panel');
    await expect(panel).toContainText('Your characters appear here once the addon has seen them.');
    await expect(panel.getByRole('link', { name: 'Get set up' })).toHaveAttribute('href', '/setup');
    await expect(panel.getByRole('link', { name: 'Your account' })).toHaveAttribute('href', '/account');
  });
});

test.describe('/account with no characters', () => {
  test.beforeEach(async ({ page }) => {
    await signInWith(page, []);
    await page.route('**/v1/devices', (route) => route.fulfill(envelope([])));
  });

  test('leads the Characters empty state with Get set up, then Paste an export', async ({ page }) => {
    await page.goto('/account');
    const empty = page.getByTestId('account-characters-empty');
    await expect(empty).toContainText('Install the addon and the companion, then log out once in game');
    await expect(empty.getByRole('link', { name: 'Get set up' })).toHaveAttribute('href', '/setup');
    await expect(empty.getByRole('link', { name: 'Paste an export' })).toBeVisible();
  });

  test('says the companion signs itself in, with the code as a text link', async ({ page }) => {
    await page.goto('/account');
    const devices = page.getByTestId('account-devices');
    await expect(devices).toContainText('The companion signs in with Battle.net from its Device page');
    // The download is the button; the code is a text link under it (W.9).
    await expect(devices.getByRole('link', { name: 'Download the companion' })).toHaveAttribute(
      'href',
      '/setup#companion',
    );
    await expect(devices.getByRole('button', { name: 'Use a pairing code instead' })).toBeVisible();
  });

  test('greets the arrival by battletag instead of naming a character', async ({ page }) => {
    await page.goto('/account?signed_in=1');
    await expect(page.getByTestId('account-signed-in-banner')).toHaveText(
      'Signed in as Fixture#1. Install the addon and the companion to bring your characters here.',
    );
  });
});

test.describe('/setup and /logs first-run copy', () => {
  test('/setup says the companion connects itself', async ({ page }) => {
    await signInWith(page, []);
    await page.goto('/setup');
    await expect(page.getByTestId('setup-line-signin')).toContainText('Signed in as Fixture#1');
    await expect(page.getByTestId('setup-companion-pair')).toHaveText(
      'Sign in with Battle.net inside the companion; it connects itself.',
    );
  });

  test('/logs puts the Battle.net line above a text-link code', async ({ page }) => {
    await signInWith(page, []);
    await page.route('**/v1/devices', (route) => route.fulfill(envelope([])));
    await page.goto('/logs');
    const line = page.getByTestId('pairing-companion-first');
    await expect(line).toHaveText(
      'Using the companion? Press Sign in with Battle.net in it. The code below is only for the fallback.',
    );
    await expect(page.getByRole('button', { name: 'Show pairing code' })).toBeVisible();
  });
});
