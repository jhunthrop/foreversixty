// web/tests/e2e/bis-signed-in.spec.ts
// Signed in on the paperdoll (design/specs/2026-11-09-directions.md, Paperdoll; owner 2026-11-08):
// the medallion holds the character's Battle.net head render, the centre stands the full render
// behind the tooltip with the name and "As last logged out" above it, and the slot's lines and
// Open in the sim go in the bar under the frame. A render that fails falls back: the medallion
// to the seal, the centre to no model, never a broken image. Stub images only, never Blizzard's.
import { readFileSync } from 'node:fs';
import path from 'node:path';
import type { Page } from '@playwright/test';
import { test, expect } from './support/test';
import { openBis } from './support/bis-page';
import { signInWith, type FixtureCharacter } from './support/selector';
import { ACTIVE_BUILD } from './support/active-build';
import { REPO_ROOT } from './support/parity';

const AVATAR = '/fixtures/avatar-placeholder.jpg';
const RENDER = '/fixtures/render-placeholder.png';
const HUNTER: FixtureCharacter = {
  name: 'Frickin Laserbeams',
  class: 'Hunter',
  spec: 'Beast Mastery',
  race: 'Night Elf',
  level: 24,
  buildDaysAgo: 0.01,
  gear: {},
  avatarUrl: AVATAR,
  renderUrl: RENDER,
};

/** The spec's published list the page compares against once signed in (data/builds, every band). */
async function routeCompareFiles(page: Page): Promise<void> {
  const file = path.join(REPO_ROOT, 'data/builds', ACTIVE_BUILD, 'bis/hunter-beast-mastery.json');
  await page.route('**/data/*/bis/hunter-beast-mastery.json', (route) =>
    route.fulfill({ status: 200, contentType: 'application/json', body: readFileSync(file, 'utf8') }),
  );
}

async function signIn(page: Page): Promise<void> {
  await routeCompareFiles(page);
  await signInWith(page, [HUNTER]);
}

const medallion = (page: Page) => page.locator('.pdf-frame [data-region="paperdoll-medallion"] img');
const centre = (page: Page) => page.locator('[data-region="paperdoll-centre"]');

test.describe('signed in, on the frame', () => {
  test.skip(({ isMobile }) => isMobile, 'the frame shows from 423 px; a phone has the list form');

  test('the medallion is the head render and the centre stands the model under the name', async ({
    page,
  }) => {
    await signIn(page);
    await openBis(page, '/bis/hunter/beast-mastery');
    await expect(medallion(page)).toHaveAttribute('src', AVATAR);
    await expect(centre(page).locator('img.pdf-model')).toHaveAttribute('src', RENDER);
    await expect(centre(page).getByTestId('bis-who')).toContainText('Frickin Laserbeams');
    await expect(centre(page).getByTestId('bis-who')).toContainText('Level 24 Night Elf Hunter');
    await expect(centre(page).getByTestId('bis-who')).toContainText('As last logged out');
    // Only the tooltip stands over the model: the lines and the sim link sit under the frame.
    await expect(centre(page).getByTestId('bis-sim-link')).toHaveCount(0);
    await expect(page.getByTestId('bis-after').getByTestId('bis-sim-link')).toBeVisible();
  });

  test('a render that fails leaves the seal and no model, never a broken image', async ({ page }) => {
    await page.route(`**${AVATAR}`, (route) => route.fulfill({ status: 403, body: '' }));
    await page.route(`**${RENDER}`, (route) => route.fulfill({ status: 404, body: '' }));
    await signIn(page);
    await openBis(page, '/bis/hunter/beast-mastery');
    await expect(medallion(page)).toHaveAttribute('src', /seal/);
    await expect(centre(page).locator('img.pdf-model')).toHaveCount(0);
  });
});

test('signed out, the medallion is the seal and the centre has no model', async ({ page, isMobile }) => {
  test.skip(isMobile, 'the frame shows from 423 px');
  await openBis(page, '/bis/hunter/beast-mastery');
  await expect(medallion(page)).toHaveAttribute('src', /seal/);
  await expect(centre(page).locator('img.pdf-model')).toHaveCount(0);
  await expect(centre(page).getByTestId('bis-sim-link')).toBeVisible();
});
