// web/tests/e2e/support/guild-parity.ts
// The guild states of the parity gate (design/parity.md, "The guild states"): each state id maps to
// the opener of the page group that owns it (support/guild-parity/*.ts). Every opener routes the
// committed fixtures (tests/e2e/fixtures/guild-parity, written by design/mocks/guild_fixtures.py,
// the same files ws_guild2.py draws the mocks from) and waits until the page shows its data.
import type { Page } from '@playwright/test';
import { OVERVIEW_OPENERS } from './guild-parity/overview';
import { RAID_OPENERS } from './guild-parity/raid';

export type GuildStateOpener = (page: Page) => Promise<void>;

const OPENERS: Readonly<Record<string, GuildStateOpener>> = {
  ...OVERVIEW_OPENERS,
  ...RAID_OPENERS,
};

export async function openGuildState(page: Page, id: string): Promise<void> {
  const open = OPENERS[id];
  if (open === undefined)
    throw new Error(`guild parity: no opener for the state "${id}" (support/guild-parity.ts)`);
  await open(page);
}
