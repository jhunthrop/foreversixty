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
import { COMPANION_DOWNLOADS, downloadFor } from '../../src/lib/setup/downloads';
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

/** Screenshots of a state at the desktop width and the owner's ~2000 px, for the persona reviews. */
async function shoot(page: import('@playwright/test').Page, name: string): Promise<void> {
  for (const width of [1440, 2000]) {
    await page.setViewportSize({ width, height: 1100 });
    await page.screenshot({ path: `test-results/import-states/${name}-${width}.png`, fullPage: true });
  }
}

test('an import that left characters out: row 1 done, the missing count under the signed-in line', async ({
  page,
  isMobile,
}) => {
  await openSetupState(page, 'partial');
  await expect(plate(page).getByRole('img', { name: 'Step 1: Done' })).toBeVisible();
  await expect(page.getByTestId('setup-steps-step-1')).toContainText(addonCopy.setupImportMissing(2));
  await expect(plate(page).getByRole('alert')).toHaveCount(0);
  await expect(page.getByTestId('setup-action-refresh')).toHaveCount(0);
  if (!isMobile) await shoot(page, 'partial');
});

test('an import that keyed nothing: the failed line leads the no-character lines, Refresh primary', async ({
  page,
  isMobile,
}) => {
  await openSetupState(page, 'import-failed');
  await expect(page.getByTestId('setup-steps-step-1')).toContainText(addonCopy.setupImportFailed);
  await expect(page.getByTestId('setup-steps-step-1')).toContainText(characterListCopy.emptyLead);
  await expect(page.getByTestId('setup-action-refresh')).toBeVisible();
  if (!isMobile) await shoot(page, 'import-failed');
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

const LATEST = 'companion-v0.2.37';
const releasePath = (asset: string) => `/dl/companion/${LATEST}/${asset}`;

for (const os of ['windows', 'mac'] as const) {
  test(`a ${os} visitor: one button for the own build, straight to the release file, its version and size, the rest under Other platforms`, async ({
    page,
  }) => {
    await openSetupState(page, os === 'mac' ? 'mac' : 'out');
    const own = downloadFor(os);
    if (own === null) throw new Error(`companion.json lists no ${os} build`);
    const button = page.getByTestId('setup-action-download');
    await expect(button).toHaveText(addonCopy.setupCompanionDownload(own.platform));
    await expect(button).toHaveAttribute('href', releasePath(own.asset));
    await expect(button).toHaveAttribute('download', own.saveAs);
    await expect(page.getByTestId('setup-download-meta')).toHaveText(/^Version 0\.2\.37 · \d+\.\d MB$/);
    await expect(page.getByTestId('setup-companion-first-run')).toHaveText(own.firstRun);
    const other = COMPANION_DOWNLOADS.find((build) => build.os !== os);
    if (other === undefined) return;
    const others = page.getByTestId('setup-other-platforms');
    await others.getByText(addonCopy.setupOtherPlatforms).click();
    await expect(others.getByTestId(`setup-download-${other.os}`)).toHaveAttribute(
      'href',
      releasePath(other.asset),
    );
    await expect(others.getByTestId(`setup-download-${other.os}`)).toHaveAttribute('download', other.saveAs);
  });
}

test('a system the page cannot name: every build side by side, none gold, as the server drew it', async ({
  page,
}) => {
  await openSetupState(page, 'unknown');
  await expect(page.getByTestId('setup-action-download')).toHaveCount(0);
  for (const build of COMPANION_DOWNLOADS) {
    await expect(page.getByTestId(`setup-download-${build.os}`)).toHaveText(
      addonCopy.setupCompanionDownload(build.platform),
    );
  }
  await expect(page.getByTestId('setup-download-meta')).toHaveText('Version 0.2.37');
  await expect(page.getByTestId('setup-companion-first-run')).toHaveCount(COMPANION_DOWNLOADS.length);
});

test('a phone: the companion runs on a PC or Mac, Copy link copies this page, and no download', async ({
  page,
  context,
}) => {
  await context.grantPermissions(['clipboard-read', 'clipboard-write']);
  await openSetupState(page, 'mobile');
  await expect(page.getByTestId('setup-steps-step-2')).toContainText(addonCopy.setupCompanionMobile);
  await expect(page.getByTestId('setup-download')).toHaveCount(0);
  await expect(page.getByTestId('setup-companion-first-run')).toHaveCount(0);
  const copy = page.getByTestId('setup-action-copylink');
  await copy.click();
  await expect(copy).toHaveText(addonCopy.setupLinkCopied);
  expect(await page.evaluate(() => navigator.clipboard.readText())).toMatch(/\/setup$/);
  // The head and The companion say the same and copy the same link.
  for (const place of ['setup-head', 'setup-companion']) {
    await expect(page.getByTestId(`${place}-copy`)).toContainText(addonCopy.setupCompanionMobile);
    await expect(page.getByTestId(`${place}-action-copylink`)).toBeVisible();
  }
});

test('the download sits at the top and right under The companion heading, the same build, never gold', async ({
  page,
}) => {
  await openSetupState(page, 'out');
  const own = downloadFor('windows');
  if (own === null) throw new Error('companion.json lists no windows build');
  // The download is the head's actions, beside its answer (ui/PageHeader).
  const head = page.locator('[data-region="head"]');
  const section = page.getByTestId('setup-companion');
  for (const [scope, place] of [
    [head, 'setup-head'],
    [section, 'setup-companion'],
  ] as const) {
    const button = scope.getByTestId(`${place}-action-download`);
    await expect(button).toHaveText(addonCopy.setupCompanionDownload(own.platform));
    await expect(button).toHaveAttribute('href', releasePath(own.asset));
    await expect(scope.getByTestId(`${place}-download-meta`)).toHaveText(/^Version 0\.2\.37 · \d+\.\d MB$/);
  }
  // The plate keeps the page's one gold step (sign in); the copies are secondary.
  const lead = await section.getByTestId('setup-companion-lead').boundingBox();
  const download = await section.getByTestId('setup-companion-download').boundingBox();
  expect(download?.y ?? Infinity).toBeLessThan(lead?.y ?? 0);
});

test('a pairing code out: the companion step waits with the code, no gold step', async ({ page }) => {
  await openSetupState(page, 'pairing');
  await expect(plate(page).getByRole('img', { name: 'Step 2: In progress' })).toBeVisible();
  await expect(page.getByTestId('setup-steps-step-2')).toContainText(pairingCopy.waiting);
  await expect(page.getByTestId('setup-pairing-code')).toBeVisible();
  await expect(plate(page).getByRole('img', { name: /: Next$/ })).toHaveCount(0);
});

test('everything done: the quiet install links stay in the Sections (U4)', async ({ page }) => {
  await openSetupState(page, 'done');
  await expect(page.getByTestId('setup-curseforge-again')).toHaveAttribute('rel', 'noopener');
  // Done or not, the download stays at the top and in The companion (owner 2026-11-08).
  for (const place of ['setup-head', 'setup-companion']) {
    await expect(page.getByTestId(`${place}-action-download`)).toHaveAttribute('rel', 'noopener');
  }
  await expect(page.getByTestId('setup-device-status')).toContainText('JHUNT-PC · Windows');
});

test('the companion comes second and the addon third, the addon a neutral to-do with both stores', async ({
  page,
}) => {
  await openSetupState(page, 'bnet');
  await expect(page.getByTestId('setup-steps-step-2')).toContainText(addonCopy.setupCompanionTitle);
  const addon = page.getByTestId('setup-steps-step-3');
  await expect(addon).toContainText(addonCopy.setupAddonTitle);
  await expect(plate(page).getByRole('img', { name: 'Step 3: Not done' })).toBeVisible();
  await expect(plate(page).getByRole('alert')).toHaveCount(0);
  await expect(page.getByTestId('setup-action-curseforge')).toHaveAttribute(
    'href',
    /curseforge\.com\/projects\//,
  );
  await expect(page.getByTestId('setup-action-wago')).toHaveAttribute(
    'href',
    /addons\.wago\.io\/addons\/foreversixty/,
  );
});

test('a companion that found no addon: the addon is the gold to-do, naming the PC, never failed', async ({
  page,
}) => {
  await openSetupState(page, 'companion');
  await expect(plate(page).getByRole('img', { name: 'Step 2: Done' })).toBeVisible();
  await expect(plate(page).getByRole('img', { name: 'Step 3: Next' })).toBeVisible();
  await expect(page.getByTestId('setup-steps-step-3')).toContainText(addonCopy.setupAddonNotOn('JHUNT-PC'));
  await expect(plate(page).getByRole('alert')).toHaveCount(0);
});

test('a companion that found the addon ticks the step with the PC and version', async ({ page }) => {
  await openSetupState(page, 'installed');
  await expect(plate(page).getByRole('img', { name: /^Step \d: Done$/ })).toHaveCount(3);
  await expect(page.getByTestId('setup-steps-step-3')).toContainText(
    addonCopy.setupAddonInstalled('JHUNT-PC', '2.0.37'),
  );
  await expect(page.getByTestId('setup-steps-step-3')).not.toContainText('is out');
  await expect(page.getByTestId('setup-curseforge-again')).toBeVisible();
});

test('an addon behind the feed stays ticked and says what is out and where to update', async ({ page }) => {
  await openSetupState(page, 'outdated');
  await expect(plate(page).getByRole('img', { name: 'Step 3: Done' })).toBeVisible();
  const addon = page.getByTestId('setup-steps-step-3');
  await expect(addon).toContainText(addonCopy.setupAddonInstalled('JHUNT-PC', '2.0.30'));
  await expect(addon).toContainText(addonCopy.setupAddonBehind('2.0.37'));
  await expect(page.getByTestId('setup-action-curseforge')).toContainText(addonCopy.setupUpdateCurseForge);
  await expect(page.getByTestId('setup-action-wago')).toContainText(addonCopy.setupUpdateWago);
  // Ticked, so it carries no gold: the plate's one primary is never an update.
  await expect(plate(page).getByRole('img', { name: /: Next$/ })).toHaveCount(0);
});

test('opened from a claim, the way back is a Notice under the head (spec 4.6)', async ({ page }) => {
  await openSetupState(page, 'claim');
  const notice = page.getByTestId('setup-back-to-claim-notice');
  await expect(
    notice.getByRole('link', { name: `${guildClaimCopy.backToClaim('Employed Millennials')} ›` }),
  ).toHaveAttribute('href', CLAIM_NEXT);
  // The plate keeps its own gold step: the companion download.
  await expect(plate(page).getByRole('img', { name: 'Step 2: Next' })).toBeVisible();
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
