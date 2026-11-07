// web/tests/e2e/setup-states.spec.ts
// /setup's states beyond the parity four (design/specs/2026-11-07-setup.md section 5), on the
// fixtures the mocks read (support/setup-page.ts): the failed Battle.net return, an account read
// that fails, the import still running, a refresh Battle.net did not answer, a Mac visitor, the
// pairing code out, the guild claim's way back, and no layout shift as the reads answer.
import { addonCopy } from '../../src/lib/addon/copy';
import { characterListCopy } from '../../src/lib/account/character-list-copy';
import { guildClaimCopy } from '../../src/lib/guild/copy';
import { pairingCopy } from '../../src/lib/reports/copy';
import { yourReportsPageCopy } from '../../src/lib/reports/my-reports-copy';
import { storePointer } from './support/selector';
import { companionDownloadPath } from '../../src/lib/downloads/paths';
import { downloadFor } from '../../src/lib/setup/downloads';
import { CLAIM_NEXT, openSetupState } from './support/setup-page';
import { expect, test } from './support/test';

const plate = (page: import('@playwright/test').Page) => page.getByRole('region', { name: 'Your setup' });

test('back from a failed Battle.net sign-in: row 1 failed, said once as an alert, the query dropped', async ({
  page,
}) => {
  await openSetupState(page, 'signin-failed');
  await expect(plate(page).getByRole('img', { name: 'Step 1: Failed' })).toBeVisible();
  await expect(plate(page).getByRole('alert')).toHaveText(addonCopy.setupSignInFailed);
  await expect(page.getByTestId('setup-action-signin')).toBeVisible();
  expect(new URL(page.url()).search).toBe('');
});

test('an account read that fails is an error with Try again, never "signed out" (B1)', async ({ page }) => {
  await openSetupState(page, 'error');
  await expect(plate(page).getByRole('alert')).toHaveText(yourReportsPageCopy.sessionFailed);
  await expect(page.getByTestId('setup-action-retry')).toHaveText(characterListCopy.tryAgain);
  await expect(page.getByTestId('setup-action-signin')).toHaveCount(0);
});

test('while the import runs, row 1 is busy, no step is gold, and /v1/me is read again', async ({ page }) => {
  let reads = 0;
  page.on('request', (request) => {
    if (new URL(request.url()).pathname === '/v1/me') reads += 1;
  });
  await openSetupState(page, 'importing');
  await expect(plate(page).getByRole('img', { name: 'Step 1: In progress' })).toBeVisible();
  await expect(page.getByTestId('setup-steps-step-1')).toContainText(addonCopy.setupImportWait);
  await expect(plate(page).getByRole('img', { name: /: Next$/ })).toHaveCount(0);
  const before = reads;
  await page.clock.runFor(5_500);
  await expect.poll(() => reads).toBeGreaterThan(before);
});

test('back from a refresh Battle.net did not answer: the refresh error and Try again', async ({ page }) => {
  await openSetupState(page, 'failed');
  await expect(plate(page).getByRole('alert')).toHaveText(characterListCopy.refreshError);
  await expect(page.getByTestId('setup-action-refresh')).toHaveText(characterListCopy.tryAgain);
  expect(new URL(page.url()).search).toBe('');
});

test('a Mac visitor: the Mac download with its own signing line, or while there is no Mac build Try an export and no SmartScreen line (U7, U10)', async ({
  page,
}) => {
  await openSetupState(page, 'mac');
  const mac = downloadFor('mac');
  await expect(page.getByTestId('setup-companion-platforms')).toHaveCount(0);
  if (mac === null) {
    await expect(page.getByTestId('setup-action-paste')).toHaveAttribute('href', '#paste');
    await expect(page.getByTestId('setup-action-download')).toHaveCount(0);
    await expect(page.getByTestId('setup-companion-first-run')).toHaveCount(0);
    return;
  }
  await expect(page.getByTestId('setup-action-download')).toHaveAttribute(
    'href',
    companionDownloadPath(mac.asset),
  );
  await expect(page.getByTestId('setup-action-download')).toHaveText(
    addonCopy.setupCompanionDownload(mac.platform),
  );
  await expect(page.getByTestId('setup-companion-first-run')).toHaveText(mac.firstRun);
});

test('a pairing code out: row 3 waits with the code, no gold step', async ({ page }) => {
  await openSetupState(page, 'pairing');
  await expect(plate(page).getByRole('img', { name: 'Step 3: In progress' })).toBeVisible();
  await expect(page.getByTestId('setup-steps-step-3')).toContainText(pairingCopy.waiting);
  await expect(page.getByTestId('setup-pairing-code')).toBeVisible();
  await expect(plate(page).getByRole('img', { name: /: Next$/ })).toHaveCount(0);
});

test('everything done: the quiet install links stay in the Sections (U4)', async ({ page }) => {
  await openSetupState(page, 'done');
  await expect(page.getByTestId('setup-curseforge-again')).toHaveAttribute('rel', 'noopener');
  await expect(page.getByTestId('setup-download-another')).toHaveAttribute('rel', 'noopener');
  await expect(page.getByTestId('setup-device-status')).toContainText('JHUNT-PC · Windows');
  await expect(page.getByTestId('setup-answer-open')).toHaveText('Open Felsnap ›');
});

test('opened from a claim, the way back is a Notice under the head (spec 4.6)', async ({ page }) => {
  await openSetupState(page, 'claim');
  const notice = page.getByTestId('setup-back-to-claim-notice');
  await expect(
    notice.getByRole('link', { name: `${guildClaimCopy.backToClaim('Employed Millennials')} ›` }),
  ).toHaveAttribute('href', CLAIM_NEXT);
  // The plate keeps its own gold step: the companion download.
  await expect(plate(page).getByRole('img', { name: 'Step 3: Next' })).toBeVisible();
});

test('a returning signed-in visitor sees no shift as the reads answer', async ({ page }) => {
  // A returning visitor has the header's character stored (the pre-paint data-pointer hint).
  await storePointer(page, {
    source: 'armory',
    ref: 'us/normal/felsnap',
    label: 'Felsnap · Beast Mastery Hunter',
    classSlug: 'hunter',
  });
  await page.addInitScript(() => {
    (window as unknown as { __shift: number }).__shift = 0;
    new PerformanceObserver((list) => {
      for (const entry of list.getEntries() as (PerformanceEntry & {
        value: number;
        hadRecentInput: boolean;
      })[]) {
        if (!entry.hadRecentInput) (window as unknown as { __shift: number }).__shift += entry.value;
      }
    }).observe({ type: 'layout-shift', buffered: true });
  });
  await openSetupState(page, 'done');
  const shift = await page.evaluate(() => (window as unknown as { __shift: number }).__shift);
  expect(shift).toBeLessThanOrEqual(0.05);
});
