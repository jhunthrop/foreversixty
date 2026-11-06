// web/tests/e2e/support/character-not-found.ts
// The character page's not-found state (design/specs/2026-11-06-character.md C10), for the fixture
// character (Elyra Duskvale, hardcore US): a 404, or a known character with no score, no gear and
// no fights seen by a viewer who does not hold it. One check, shared by every spec that lands there.
import type { Page } from '@playwright/test';
import { expect } from './test';
import { ctaCopy } from '../../../src/lib/addon/cta-copy';

/** C10's one line for the fixture's address (characterPageCopy.notFound, spelled out because
 *  page-view.ts's import chain reaches JSON Playwright's loader cannot take). */
export const FIXTURE_NOT_FOUND_LINE = 'No character named Elyra Duskvale on US Hardcore.';
/** characterPageCopy.signIn: the row a signed-out viewer gets. */
export const SIGN_IN_ROW = 'Sign in to see your characters';

/** The not-found page: the name as the title, the line, the lookup filled with the name, the add
 *  rows (sign in only when signed out, then the addon), and none of the ready page's Sections. */
export async function expectFixtureNotFound(page: Page, { signedIn }: { signedIn: boolean }): Promise<void> {
  await expect(page.getByTestId('character-not-found')).toBeVisible();
  await expect(page.getByTestId('character-head').getByTestId('page-title')).toHaveText('Elyra Duskvale');
  await expect(page.getByTestId('character-not-found-line')).toHaveText(FIXTURE_NOT_FOUND_LINE);
  const lookup = page.getByTestId('character-lookup');
  await expect(lookup.getByRole('textbox', { name: 'Name' })).toHaveValue('Elyra Duskvale');
  await expect(lookup.getByRole('button', { name: 'Look up' })).toBeVisible();
  const add = page.getByTestId('character-add').getByRole('link');
  const rows = signedIn ? [ctaCopy.getAddonAndCompanion] : [SIGN_IN_ROW, ctaCopy.getAddonAndCompanion];
  await expect(add).toHaveCount(rows.length);
  await expect(add).toContainText(rows);
  await expect(add.last()).toHaveAttribute('href', '/setup');
  await expect(page.getByTestId('character')).toHaveCount(0);
  await expect(page.getByTestId('character-score')).toHaveCount(0);
  await expect(page.getByTestId('character-gear-section')).toHaveCount(0);
}
