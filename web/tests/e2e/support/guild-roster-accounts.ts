// web/tests/e2e/support/guild-roster-accounts.ts
// The Roster tab leads with the dense raid roster (components/guild/dense/GuildRaidRoster.svelte);
// the account table (ranks, Remove, approvals, members with no account) sits under it in the
// collapsed "Ranks, accounts and removals" disclosure. Specs that read the account table open it
// first. A page with no disclosure (a visitor, a gated roster) is left as it is.
import type { Page } from '@playwright/test';
import { expect } from './test';

/** How long to wait for the disclosure before deciding the page has none. */
const ACCOUNTS_WAIT_MS = 5_000;

export async function openRosterAccounts(page: Page): Promise<void> {
  const accounts = page.getByTestId('guild-roster-accounts');
  try {
    await accounts.waitFor({ state: 'visible', timeout: ACCOUNTS_WAIT_MS });
  } catch {
    return;
  }
  const open = await accounts.evaluate((el) => (el as HTMLDetailsElement).open);
  if (!open) await accounts.locator('summary').first().click();
  await expect(page.getByTestId('guild-roster-tab')).toBeVisible();
}
