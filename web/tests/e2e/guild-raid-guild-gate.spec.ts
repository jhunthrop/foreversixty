// web/tests/e2e/guild-raid-guild-gate.spec.ts
// Launch pricing spec §9 case 16: on a Guild-tier guild (officer tools on, no Raid guild) an
// officer's Loot and Readiness tabs say the Raid guild sentence, its link goes to `#billing`, and
// Award, Undo and Nudge are in place and disabled. A guild with Raid guild shows no sentence.
import type { Page } from '@playwright/test';
import { expect, test, apiUnreachable, GUILD_OVERVIEW_READS } from './support/test';
import { guildGateCopy } from '../../src/lib/guild/copy';
import { gotoHydrated } from './support/hydrated';
import {
  GUILD_SETTINGS_PATH,
  GUILD_URL,
  awardButtons,
  nudgeButtons,
  stubGuildCentre,
  undoButtons,
  unawardedItem,
  awardedItemId,
  type GuildCentreStub,
} from './support/officer-tools';

test.beforeEach(({ page }) => apiUnreachable(page, ...GUILD_OVERVIEW_READS));

const GUILD_TIER: GuildCentreStub = {
  viewer: 'officer',
  officerTools: true,
  raidGuildTools: false,
  claimHolder: true,
};

type Tab = 'loot' | 'readiness';

async function openTab(page: Page, tab: Tab, stub: GuildCentreStub): Promise<void> {
  await stubGuildCentre(page, stub);
  await gotoHydrated(page, `${GUILD_URL}#${tab}`, `guild-${tab}-tab`);
}

for (const tab of ['loot', 'readiness'] as const) {
  test(`Guild tier: ${tab} says the Raid guild sentence and opens the Move confirm`, async ({ page }) => {
    await openTab(page, tab, GUILD_TIER);
    const gate = page.getByTestId(`guild-${tab}-plan-gate`);
    await expect(gate).toBeVisible();
    await expect(gate).toHaveText(`${guildGateCopy.lead[tab]} ${guildGateCopy.moveToRaidGuild}.`);
    const link = gate.getByRole('link', { name: guildGateCopy.moveToRaidGuild });
    await expect(link).toHaveAttribute('href', new RegExp(`${GUILD_SETTINGS_PATH}\\?upgrade=1#billing$`));
  });
}

test('Guild tier: every Award, Undo and Nudge is in place and disabled', async ({ page }) => {
  await openTab(page, 'loot', GUILD_TIER);
  // The council's controls live on the picked drop: an open one carries Award, the awarded one Undo.
  for (const [itemId, controls] of [
    [unawardedItem().itemId, awardButtons(page)],
    [awardedItemId(), undoButtons(page)],
  ] as const) {
    await page.getByTestId(`guild-loot-drop-${itemId}`).click();
    await expect.poll(() => controls.count()).toBeGreaterThan(0);
    await expect(controls.first()).toBeDisabled();
    expect(await controls.evaluateAll((nodes) => nodes.every((n) => (n as HTMLButtonElement).disabled))).toBe(
      true,
    );
  }
  await page.getByTestId('guild-tab-readiness').click();
  const nudges = nudgeButtons(page);
  // The Readiness tab reads its board after the click; count() does not wait for it.
  await expect.poll(() => nudges.count()).toBeGreaterThan(0);
  expect(await nudges.evaluateAll((nodes) => nodes.every((n) => (n as HTMLButtonElement).disabled))).toBe(
    true,
  );
});

test('Raid guild: no sentence and every control live', async ({ page }) => {
  await openTab(page, 'loot', { viewer: 'officer', officerTools: true, raidGuildTools: true });
  await expect(page.getByTestId('guild-loot-plan-gate')).toHaveCount(0);
  await page.getByTestId(`guild-loot-drop-${unawardedItem().itemId}`).click();
  await expect(awardButtons(page).first()).toBeEnabled();
});

test('an officer without the claim is told who turns it on', async ({ page }) => {
  await openTab(page, 'loot', { viewer: 'officer', officerTools: false, raidGuildTools: false });
  await expect(page.getByTestId('guild-loot-plan-gate')).toHaveText(
    `${guildGateCopy.lead.loot} ${guildGateCopy.ask}Kraggor${guildGateCopy.askTail}`,
  );
  await expect(page.getByTestId('guild-loot-plan-gate').getByRole('link')).toHaveCount(0);
});
