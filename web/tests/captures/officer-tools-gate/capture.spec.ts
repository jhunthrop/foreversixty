// web/tests/captures/officer-tools-gate/capture.spec.ts
// Acceptance screenshots for the officer tools gate (experience spec
// design/specs/2026-10-25-officer-tools-gate.md, section 8). Not a test of behaviour: it drives
// the guild control centre, guild settings and the account page into the 13 states, writes
// PNGs to design/reviews/2026-10-25-officer-tools-gate/, and measures what the sentence adds to
// each tab panel against §3.5's table (annotations plus heights-<tab>-<width>.json there).
import { writeFileSync } from 'node:fs';
import path from 'node:path';
import { fileURLToPath } from 'node:url';
import { test, type Locator, type Page } from '@playwright/test';
import {
  OFFICER_CAN,
  SETTINGS_GUILD_PATH,
  guildBilling,
  meWithGuildPlan,
  routeGuildSettingsPage,
  scopedGrant,
} from '../../../src/test-support/billing-fixtures';
import { gotoHydrated } from '../../e2e/support/hydrated';
import {
  GUILD_URL,
  awardedItemId,
  envelope,
  refuseAwards,
  stubAccountWithTwoGuilds,
  stubGuildCentre,
  unawardedItem,
  type GuildCentreStub,
} from '../../e2e/support/officer-tools';

const OUT_DIR = path.resolve(
  path.dirname(fileURLToPath(import.meta.url)),
  '../../../../design/reviews/2026-10-25-officer-tools-gate',
);
const PAD = 24;
const TOP_CROP_PX = 520;
const VIEWPORT_HEIGHT = 1000;
const SETTLE_MS = 400;
const PAIR = [390, 1440];
const WITH_360 = [360, ...PAIR];
const ALL = [360, 390, 1440, 2000];

/** §3.5: what the sentence may add to each panel, by width. */
const CAP_PX: Record<'loot' | 'readiness', Record<number, number>> = {
  loot: { 360: 56, 390: 56, 1440: 36, 2000: 36 },
  readiness: { 360: 52, 390: 52, 1440: 32, 2000: 32 },
};

type Tab = 'loot' | 'readiness';

const target = (name: string): string => path.join(OUT_DIR, name);
const settle = (page: Page): Promise<void> => page.waitForTimeout(SETTLE_MS);

async function openTab(page: Page, width: number, tab: Tab, stub: GuildCentreStub): Promise<void> {
  await page.setViewportSize({ width, height: VIEWPORT_HEIGHT });
  await stubGuildCentre(page, stub);
  await gotoHydrated(page, `${GUILD_URL}#${tab}`, `guild-${tab}-tab`);
  await settle(page);
}

/** The element's page rectangle grown by PAD, at most `maxHeight` tall, kept inside the page. */
async function cropOf(page: Page, element: Locator, file: string, maxHeight = Infinity): Promise<void> {
  const rect = await element.evaluate((node) => {
    const box = node.getBoundingClientRect();
    return {
      x: box.x + window.scrollX,
      y: box.y + window.scrollY,
      width: box.width,
      height: box.height,
      pageWidth: document.documentElement.clientWidth,
    };
  });
  const x = Math.max(0, rect.x - PAD);
  const y = Math.max(0, rect.y - PAD);
  await page.screenshot({
    path: target(file),
    fullPage: true,
    clip: {
      x,
      y,
      width: Math.min(rect.width + PAD * 2, rect.pageWidth - x),
      height: Math.min(rect.height + PAD * 2, maxHeight),
    },
  });
}

/** A full-page capture and a crop of `subject` (its top TOP_CROP_PX for a long tab). */
async function shoot(page: Page, state: string, width: number, subject: Locator, maxHeight?: number) {
  await page.screenshot({ path: target(`officer-tools-${state}-${width}.png`), fullPage: true });
  await cropOf(page, subject, `officer-tools-${state}-panel-${width}.png`, maxHeight);
}

function atWidths(title: string, widths: number[], body: (page: Page, width: number) => Promise<void>): void {
  for (const width of widths) {
    test(`${title} @${width}`, ({ page }) => body(page, width));
  }
}

const tabPanel = (page: Page, tab: Tab): Locator => page.getByTestId(`guild-${tab}-tab`);

// 1. Loot, officer with tools on: the controls as they were.
atWidths('1 loot tools on', PAIR, async (page, width) => {
  await openTab(page, width, 'loot', { viewer: 'officer', officerTools: true });
  await shoot(page, '01-loot-tools-on', width, tabPanel(page, 'loot'), TOP_CROP_PX);
});

// 2. Loot, officer-read-only, drops listed.
atWidths('2 loot read-only', ALL, async (page, width) => {
  await openTab(page, width, 'loot', { viewer: 'officer', officerTools: false });
  await shoot(page, '02-loot-read-only', width, tabPanel(page, 'loot'), TOP_CROP_PX);
});

// 3. Loot, officer-read-only, no kills yet.
atWidths('3 loot read-only empty', PAIR, async (page, width) => {
  await openTab(page, width, 'loot', { viewer: 'officer', officerTools: false, emptyLoot: true });
  await shoot(page, '03-loot-read-only-empty', width, tabPanel(page, 'loot'));
});

// 4. Loot, officer-read-only, the awarded item: Undo disabled.
atWidths('4 loot read-only awarded item', PAIR, async (page, width) => {
  await openTab(page, width, 'loot', { viewer: 'officer', officerTools: false });
  const card = page.getByTestId(`guild-loot-item-${awardedItemId()}`);
  await card.scrollIntoViewIfNeeded();
  await shoot(page, '04-loot-read-only-awarded', width, card);
});

// 5. Loot, member: no sentence, no controls.
atWidths('5 loot member', PAIR, async (page, width) => {
  await openTab(page, width, 'loot', { viewer: 'member', officerTools: false });
  await shoot(page, '05-loot-member', width, tabPanel(page, 'loot'), TOP_CROP_PX);
});

// 6. Loot after the 403 on Award: the refused line focused, the sentence shown.
atWidths('6 loot after the 403', PAIR, async (page, width) => {
  await openTab(page, width, 'loot', { viewer: 'officer', officerTools: true });
  await refuseAwards(page);
  const { itemId, characterKey } = unawardedItem();
  await page.getByTestId(`guild-loot-award-${itemId}-${characterKey}`).click();
  await page.getByTestId(`guild-loot-award-yes-${itemId}`).click();
  await page.getByTestId(`guild-loot-plan-refused-${itemId}`).waitFor();
  await settle(page);
  await page.screenshot({ path: target(`officer-tools-06-loot-refused-${width}.png`), fullPage: true });
  await cropOf(page, tabPanel(page, 'loot'), `officer-tools-06-loot-refused-panel-${width}.png`, TOP_CROP_PX);
  await cropOf(
    page,
    page.getByTestId(`guild-loot-item-${itemId}`),
    `officer-tools-06-loot-refused-card-${width}.png`,
  );
});

// 7. Readiness, officer-read-only (grid on desktop, stacked cards on a phone).
atWidths('7 readiness read-only', ALL, async (page, width) => {
  await openTab(page, width, 'readiness', { viewer: 'officer', officerTools: false });
  await shoot(page, '07-readiness-read-only', width, tabPanel(page, 'readiness'), TOP_CROP_PX);
});

// 8. Readiness, member.
atWidths('8 readiness member', PAIR, async (page, width) => {
  await openTab(page, width, 'readiness', { viewer: 'member', officerTools: false });
  await shoot(page, '08-readiness-member', width, tabPanel(page, 'readiness'), TOP_CROP_PX);
});

// 9. Focus-visible on Subscribe the guild, reached with the keyboard.
atWidths('9 link focus-visible', PAIR, async (page, width) => {
  await openTab(page, width, 'loot', { viewer: 'officer', officerTools: false });
  await page.keyboard.press('Shift');
  await page.getByTestId('guild-loot-plan-gate').getByRole('link').focus();
  await settle(page);
  await shoot(page, '09-link-focus', width, page.getByTestId('guild-loot-plan-gate'));
});

async function openSettings(page: Page, width: number, me: ReturnType<typeof meWithGuildPlan>) {
  await page.setViewportSize({ width, height: VIEWPORT_HEIGHT });
  await routeGuildSettingsPage(page);
  await page.route('**/v1/me', (route) => route.fulfill(envelope(me.data)));
  await gotoHydrated(page, SETTINGS_GUILD_PATH, 'guild-settings-billing');
  await settle(page);
}

const billing = (page: Page): Locator => page.getByTestId('guild-settings-billing');

// 10. Settings, scoped grant, claim holder.
atWidths('10 settings scoped claim holder', WITH_360, async (page, width) => {
  await openSettings(page, width, meWithGuildPlan(scopedGrant()));
  await shoot(page, '10-settings-scoped-claim-holder', width, billing(page));
});

// 11. Settings, scoped grant, an officer without the claim.
atWidths('11 settings scoped officer', PAIR, async (page, width) => {
  await openSettings(page, width, meWithGuildPlan(scopedGrant(), OFFICER_CAN));
  await shoot(page, '11-settings-scoped-officer', width, billing(page));
});

// 12. Settings, granted full plan (regression).
atWidths('12 settings granted plan', PAIR, async (page, width) => {
  const granted = guildBilling({
    current_period_end: '2026-12-31T00:00:00Z',
    billed_by: '',
    you_are_billing_contact: false,
  });
  await openSettings(page, width, meWithGuildPlan(granted));
  await shoot(page, '12-settings-granted', width, billing(page));
});

// 13. Account, Guild panel: one scoped row and one paid row.
atWidths('13 account guild rows', WITH_360, async (page, width) => {
  await page.setViewportSize({ width, height: VIEWPORT_HEIGHT });
  await stubAccountWithTwoGuilds(page);
  await gotoHydrated(page, '/account', 'account-guilds');
  await settle(page);
  await shoot(page, '13-account-guild-rows', width, page.getByTestId('account-guilds-panel'));
});

// §3.5: the panel with the sentence minus the same panel without it, at the same width.
for (const tab of ['loot', 'readiness'] as const) {
  atWidths(`height ${tab}`, ALL, async (page, width) => {
    const heightOf = async (on: Page, officerTools: boolean): Promise<number> => {
      await openTab(on, width, tab, { viewer: 'officer', officerTools });
      return (await tabPanel(on, tab).boundingBox())?.height ?? 0;
    };
    // A second page: a goto to the same URL and hash would not reload the first.
    const live = await page.context().newPage();
    const without = await heightOf(live, true);
    await live.close();
    const withGate = await heightOf(page, false);
    const sentence = (await page.getByTestId(`guild-${tab}-plan-gate`).boundingBox())?.height ?? 0;
    const added = Math.round((withGate - without) * 10) / 10;
    const result = { tab, width, added, sentence, cap: CAP_PX[tab][width] };
    test.info().annotations.push({
      type: 'height',
      description: `${tab} @${width}: adds ${added} px (sentence ${sentence} px), cap ${result.cap}`,
    });
    writeFileSync(target(`heights-${tab}-${width}.json`), `${JSON.stringify(result)}\n`);
  });
}
