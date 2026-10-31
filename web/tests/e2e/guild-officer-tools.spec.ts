// web/tests/e2e/guild-officer-tools.spec.ts
// Officer tools gate acceptance (design/specs/2026-10-25-officer-tools-gate.md §8). An officer
// whose guild has no officer tools (`home.viewer.officer_tools: false`) sees every Award, Undo
// and Nudge in place and disabled, with one sentence per tab that links to guild settings; a
// 403 `guild_plan_required` mid-session flips both tabs without a reload; guild settings and
// the account page's guild rows say what the guild has. Members see nothing new.
import type { Locator, Page } from '@playwright/test';
import { expect, test } from './support/test';
import { guildGateCopy } from '../../src/lib/guild/copy';
import type { GuildBillingSettings } from '../../src/lib/billing/contract';
import {
  SETTINGS_GUILD,
  SETTINGS_GUILD_PATH,
  guildBillingRead,
  planView,
} from '../../src/test-support/billing-fixtures';
import { openWith } from './support/guild-billing';
import { gotoHydrated } from './support/hydrated';
import {
  GUILD_SETTINGS_PATH,
  GUILD_URL,
  awardButtons,
  awardedItemId,
  envelope,
  expectNoAxeViolations,
  nudgeButtons,
  refuseAwards,
  stubAccountGuildRows,
  stubGuildCentre,
  undoButtons,
  unawardedItem,
  type GuildCentreStub,
} from './support/officer-tools';

const SENTENCE = 'Award and Nudge are part of the guild plan. Subscribe the guild.';
const MASTER_ONLY = 'The subscriber or the guild leader manages billing and seats.';
const SCOPED_PLAN_LINE = 'Officer tools · until 12 Dec 2026';
const CLIPBOARD_SENTINEL = 'clipboard untouched';
/** The scoped CLI grant as GET /v1/guilds/{id}/billing sends it: Award, Undo and Nudge until 12 Dec 2026. */
const SCOPED_GRANT_PLAN = planView({
  scope: 'officer_tools',
  current_period_end: '2026-12-12T00:00:00Z',
  interval: null,
  amount_cents: null,
  billed_by: '',
  you_are_billing_contact: false,
});

type Tab = 'loot' | 'readiness';

async function openTab(page: Page, tab: Tab, stub: GuildCentreStub): Promise<void> {
  await stubGuildCentre(page, stub);
  await gotoHydrated(page, `${GUILD_URL}#${tab}`, `guild-${tab}-tab`);
}

const gate = (page: Page, tab: Tab): Locator => page.getByTestId(`guild-${tab}-plan-gate`);

/** [total, enabled] for a control family. */
async function counts(controls: Locator): Promise<[number, number]> {
  const states = await controls.evaluateAll((nodes) => nodes.map((n) => (n as HTMLButtonElement).disabled));
  return [states.length, states.filter((disabled) => !disabled).length];
}

async function expectAllDisabled(controls: Locator, expected: number): Promise<void> {
  expect(await counts(controls)).toEqual([expected, 0]);
  const described = await controls.evaluateAll((nodes) =>
    nodes.map((n) => n.getAttribute('aria-describedby')),
  );
  expect(new Set(described).size).toBe(1);
}

/** What the officer with tools on sees enabled (state 1): the baseline the disabled counts match. */
async function enabledCounts(page: Page): Promise<{ award: number; undo: number; nudge: number }> {
  await openTab(page, 'loot', { viewer: 'officer', officerTools: true });
  const [award, awardEnabled] = await counts(awardButtons(page));
  const [undo, undoEnabled] = await counts(undoButtons(page));
  expect(award).toBe(awardEnabled);
  expect(undo).toBe(undoEnabled);
  await page.getByTestId('guild-tab-readiness').click();
  await expect(page.getByTestId('guild-readiness-tab')).toBeVisible();
  const [nudge, nudgeEnabled] = await counts(nudgeButtons(page));
  expect(nudge).toBe(nudgeEnabled);
  expect(award).toBeGreaterThan(0);
  expect(undo).toBe(1);
  expect(nudge).toBeGreaterThan(0);
  return { award, undo, nudge };
}

/** The Guild settings link row that holds an account guild plan value. */
const guildSettingsLinkOf = (value: Locator): Locator =>
  value.locator('xpath=ancestor::a').filter({ hasText: 'Guild settings' });
test.describe('officer-read-only: every control in place and disabled (§3.4)', () => {
  test('state 1: tools on shows no sentence and every control live', async ({ page, context }) => {
    await context.grantPermissions(['clipboard-read', 'clipboard-write']);
    await enabledCounts(page);
    await expect(gate(page, 'readiness')).toHaveCount(0);
    // The positive control for the disabled-Nudge test below: a live Nudge does write.
    await page.evaluate((text) => navigator.clipboard.writeText(text), CLIPBOARD_SENTINEL);
    await page.locator('button[data-testid^="guild-readiness-nudge-"]:visible').first().click();
    await expect.poll(() => page.evaluate(() => navigator.clipboard.readText())).not.toBe(CLIPBOARD_SENTINEL);
    await page.getByTestId('guild-tab-loot').click();
    await expect(gate(page, 'loot')).toHaveCount(0);
  });

  test('an older API without the field keeps the tools on', async ({ page }) => {
    await openTab(page, 'loot', { viewer: 'officer' });
    await expect(gate(page, 'loot')).toHaveCount(0);
    expect((await counts(awardButtons(page)))[1]).toBeGreaterThan(0);
  });

  test('states 2 and 7: the counts equal state 1, each disabled and described by the sentence', async ({
    browser,
  }) => {
    const live = await browser.newPage();
    const expected = await enabledCounts(live);
    await live.close();

    const page = await browser.newPage();
    await openTab(page, 'loot', { viewer: 'officer', officerTools: false });
    await expect(gate(page, 'loot')).toHaveText(SENTENCE);
    await expect(page.getByTestId('guild-loot-tab').getByText(guildGateCopy.readOnly)).toBeVisible();
    await expectAllDisabled(awardButtons(page), expected.award);
    await expectAllDisabled(undoButtons(page), expected.undo);
    await expect(awardButtons(page).first()).toHaveAttribute('aria-describedby', 'officer-tools-gate-loot');

    await page.getByTestId('guild-tab-readiness').click();
    await expect(gate(page, 'readiness')).toHaveText(SENTENCE);
    await expect(page.getByTestId('guild-readiness-tab').getByText(guildGateCopy.readOnly)).toBeVisible();
    await expectAllDisabled(nudgeButtons(page), expected.nudge);
    await expect(nudgeButtons(page).first()).toHaveAttribute(
      'aria-describedby',
      'officer-tools-gate-readiness',
    );
    await expect(page.getByText('Nudge copies a message you can paste to them.')).toBeVisible();
    await page.close();
  });

  test('state 3: the sentence sits above the empty state', async ({ page }) => {
    await openTab(page, 'loot', { viewer: 'officer', officerTools: false, emptyLoot: true });
    const sentence = await gate(page, 'loot').boundingBox();
    const empty = await page.getByTestId('guild-loot-empty').boundingBox();
    expect(sentence && empty && sentence.y + sentence.height <= empty.y).toBe(true);
  });

  test('a click on a disabled Award opens no confirm row and posts nothing', async ({ page }) => {
    const posts: string[] = [];
    await page.route('**/v1/guilds/*/loot/awards**', (route) => {
      posts.push(route.request().method());
      return route.fulfill(envelope({ award_id: 1, status: 'awarded' }));
    });
    await openTab(page, 'loot', { viewer: 'officer', officerTools: false });
    const { itemId, characterKey } = unawardedItem();
    await page.getByTestId(`guild-loot-award-${itemId}-${characterKey}`).click({ force: true });
    await page.getByTestId(`guild-loot-row-undo-${awardedItemId()}`).click({ force: true });
    await expect(page.locator('[data-testid^="guild-loot-award-confirm-"]')).toHaveCount(0);
    expect(posts).toEqual([]);
  });

  test('a click on a disabled Nudge writes nothing to the clipboard', async ({ page, context }) => {
    await context.grantPermissions(['clipboard-read', 'clipboard-write']);
    await openTab(page, 'readiness', { viewer: 'officer', officerTools: false });
    await page.evaluate((text) => navigator.clipboard.writeText(text), CLIPBOARD_SENTINEL);
    const nudge = page.locator('button[data-testid^="guild-readiness-nudge-"]:visible').first();
    await nudge.click({ force: true });
    expect(await page.evaluate(() => navigator.clipboard.readText())).toBe(CLIPBOARD_SENTINEL);
  });

  for (const tab of ['loot', 'readiness'] as const) {
    test(`the ${tab} sentence links to the guild's settings`, async ({ page }) => {
      await openTab(page, tab, { viewer: 'officer', officerTools: false });
      const link = gate(page, tab).getByRole('link', { name: guildGateCopy.link });
      await expect(link).toHaveAttribute('href', GUILD_SETTINGS_PATH);
      await page.route(`**/v1/guilds/*/settings`, (route) => route.fulfill(envelope({})));
      await link.click();
      await page.waitForURL(`**${GUILD_SETTINGS_PATH}`);
    });
  }

  for (const width of [360, 1440]) {
    for (const tab of ['loot', 'readiness'] as const) {
      test(`the ${tab} sentence is inside the viewport on load at ${width}`, async ({ page }) => {
        await page.setViewportSize({ width, height: width === 360 ? 740 : 900 });
        await openTab(page, tab, { viewer: 'officer', officerTools: false });
        const box = await gate(page, tab).boundingBox();
        const viewport = page.viewportSize();
        expect(box && viewport).toBeTruthy();
        if (box && viewport) {
          expect(box.x).toBeGreaterThanOrEqual(0);
          expect(box.y).toBeGreaterThanOrEqual(0);
          expect(box.x + box.width).toBeLessThanOrEqual(viewport.width);
          expect(box.y + box.height).toBeLessThanOrEqual(viewport.height);
        }
      });
    }
  }

  for (const [width, hitArea] of [
    [360, 44],
    [1440, null],
  ] as const) {
    test(`the link's phone hit area at ${width}: ${hitArea ?? 'none'}`, async ({ page }) => {
      await page.setViewportSize({ width, height: 900 });
      await openTab(page, 'loot', { viewer: 'officer', officerTools: false });
      const link = gate(page, 'loot').getByRole('link');
      const after = await link.evaluate((node) => {
        const style = getComputedStyle(node, '::after');
        return { display: style.display, height: parseFloat(style.height) };
      });
      if (hitArea === null) expect(after.display).toBe('none');
      else expect(after.height).toBeGreaterThanOrEqual(hitArea);
    });
  }

  test('keyboard: the link follows the tab strip and every disabled control is out of the order', async ({
    page,
  }) => {
    await openTab(page, 'readiness', { viewer: 'officer', officerTools: false });
    const link = gate(page, 'readiness').getByRole('link');
    await link.focus();
    await page.keyboard.press('Tab');
    const next = await page.evaluate(() => document.activeElement?.getAttribute('data-testid') ?? '');
    expect(next.startsWith('guild-readiness-nudge-')).toBe(false);
  });
});

test.describe('axe on the gate states (§8: 2, 7, 10, 13)', () => {
  test.use({ reducedMotion: 'reduce' });

  test('state 2: Loot, officer-read-only', async ({ page }) => {
    await openTab(page, 'loot', { viewer: 'officer', officerTools: false });
    await expectNoAxeViolations(page, '[data-testid="guild-loot-tab"]');
  });

  test('state 7: Readiness, officer-read-only', async ({ page }) => {
    await openTab(page, 'readiness', { viewer: 'officer', officerTools: false });
    await expectNoAxeViolations(page, '[data-testid="guild-readiness-tab"]');
  });

  test('state 10: settings, scoped grant, claim holder', async ({ page }) => {
    await openWith(page, guildBillingRead({ billing: SCOPED_GRANT_PLAN }));
    await expectNoAxeViolations(page, '[data-testid="guild-settings-billing"]');
  });

  test('state 13: account, one scoped row and one paid row', async ({ page }) => {
    await stubAccountGuildRows(page);
    await gotoHydrated(page, '/account', 'account-guilds');
    await expectNoAxeViolations(page, '[data-testid="account-guilds-panel"]');
  });
});

test.describe('the 403 guild_plan_required path (§4)', () => {
  test('state 6: no reload, the refused line takes focus, both tabs turn read-only', async ({ page }) => {
    await stubGuildCentre(page, { viewer: 'officer', officerTools: true });
    const seen = await refuseAwards(page);
    await gotoHydrated(page, `${GUILD_URL}#loot`, 'guild-loot-tab');
    await page.evaluate(() => {
      (window as unknown as { __noReload: boolean }).__noReload = true;
    });
    const { itemId, characterKey } = unawardedItem();
    await page.getByTestId(`guild-loot-award-${itemId}-${characterKey}`).click();
    await page.getByTestId(`guild-loot-award-yes-${itemId}`).click();

    const refused = page.getByTestId(`guild-loot-plan-refused-${itemId}`);
    await expect(refused).toHaveText(guildGateCopy.awardRefused);
    await expect(refused).toBeFocused();
    await expect(refused).toHaveAttribute('role', 'alert');
    expect(seen).toEqual(['POST']);
    await expect(
      page.getByTestId(`guild-loot-item-${itemId}`).getByTestId(`guild-loot-plan-refused-${itemId}`),
    ).toHaveCount(1);
    await expect(gate(page, 'loot')).toHaveText(SENTENCE);
    await expect(page.getByTestId('guild-loot-error')).toHaveCount(0);
    expect((await counts(awardButtons(page)))[1]).toBe(0);

    await page.getByTestId('guild-tab-readiness').click();
    await expect(gate(page, 'readiness')).toHaveText(SENTENCE);
    expect((await counts(nudgeButtons(page)))[1]).toBe(0);
    expect(await page.evaluate(() => (window as unknown as { __noReload?: boolean }).__noReload)).toBe(true);
  });
});

test.describe('members see nothing new (states 5 and 8)', () => {
  test('state 5: Loot for a member has no sentence and no control', async ({ page }) => {
    await openTab(page, 'loot', { viewer: 'member', officerTools: false });
    await expect(gate(page, 'loot')).toHaveCount(0);
    await expect(awardButtons(page)).toHaveCount(0);
    await expect(undoButtons(page)).toHaveCount(0);
  });

  test('state 8: Readiness for a member has no sentence, label or Nudge', async ({ page }) => {
    await openTab(page, 'readiness', { viewer: 'member', officerTools: false });
    await expect(gate(page, 'readiness')).toHaveCount(0);
    await expect(page.getByTestId('guild-readiness-tab').getByText(guildGateCopy.readOnly)).toHaveCount(0);
    await expect(nudgeButtons(page)).toHaveCount(0);
  });
});

// Guild settings #billing reads GET /v1/guilds/{id}/billing (launch pricing spec 3.2): a grant keeps
// its line and explanation, then the chooser for the claim holder (3.2.1); an officer without the
// claim reads the plan line and the master-only line, with no buttons (3.2, 6.2).
test.describe('guild settings billing on a scoped grant (§5, states 10–12)', () => {
  const GRANT_ID = SETTINGS_GUILD.id;
  const READ_ONLY: GuildBillingSettings['can'] = [];
  const FULL_GRANT_PLAN = planView({
    current_period_end: '2026-12-31T00:00:00Z',
    interval: null,
    amount_cents: null,
    billed_by: '',
    you_are_billing_contact: false,
  });

  test('state 10: the claim holder reads the grant and can subscribe the guild', async ({ page }) => {
    await openWith(page, guildBillingRead({ billing: SCOPED_GRANT_PLAN }));
    await expect(page.getByTestId('guild-billing-plan')).toHaveText(SCOPED_PLAN_LINE);
    await expect(page.getByTestId('guild-billing-contact')).toHaveText(
      'Granted by Forever Sixty for the first raid nights: Award, Undo and Nudge only, until 12 Dec 2026. After that they need the guild plan.',
    );
    await expect(page.getByTestId('guild-billing-chooser')).toBeVisible();
    await expect(page.getByTestId('guild-billing-subscribe')).toHaveAttribute(
      'href',
      `/premium/checkout?plan=guild&interval=yearly&guild_id=${GRANT_ID}`,
    );
    await expect(page.getByTestId('guild-manage-billing')).toHaveCount(0);
    await expect(page.getByTestId('guild-settings-master-only')).toHaveCount(0);
  });

  test('state 11: an officer without the claim reads the grant, then the master-only line', async ({
    page,
  }) => {
    await openWith(page, guildBillingRead({ billing: SCOPED_GRANT_PLAN, can: READ_ONLY }));
    await expect(page.getByTestId('guild-billing-plan')).toHaveText(SCOPED_PLAN_LINE);
    await expect(page.getByTestId('guild-billing-contact')).toContainText('Award, Undo and Nudge only');
    await expect(page.getByTestId('guild-settings-master-only')).toHaveText(MASTER_ONLY);
    await expect(page.getByTestId('guild-billing-chooser')).toHaveCount(0);
    await expect(page.getByTestId('guild-billing-subscribe')).toHaveCount(0);
    await expect(page.getByTestId('guild-manage-billing')).toHaveCount(0);
  });

  test('an officer without the claim and no plan reads only the master-only line', async ({ page }) => {
    await openWith(page, guildBillingRead({ can: READ_ONLY }));
    await expect(page.getByTestId('guild-settings-master-only')).toHaveText(MASTER_ONLY);
    await expect(page.getByTestId('guild-billing-plan')).toHaveCount(0);
    await expect(page.getByTestId('guild-billing-chooser')).toHaveCount(0);
    await expect(page.getByTestId('guild-billing-subscribe')).toHaveCount(0);
    await expect(page.getByRole('link', { name: guildGateCopy.link })).toHaveCount(0);
  });

  test('the claim holder with no plan gets the chooser and nothing else above Records', async ({ page }) => {
    const billing = page.getByTestId('guild-settings-billing');
    await openWith(page, guildBillingRead());
    await expect(page.getByTestId('guild-billing-chooser')).toBeVisible();
    await expect(page.getByTestId('guild-billing-subscribe')).toHaveAttribute(
      'href',
      `/premium/checkout?plan=guild&interval=yearly&guild_id=${GRANT_ID}`,
    );
    await expect(page.getByTestId('guild-billing-plan')).toHaveCount(0);
    await expect(billing.getByTestId('guild-settings-master-only')).toHaveCount(0);
  });

  test('state 12: a granted full plan is unchanged', async ({ page }) => {
    await openWith(page, guildBillingRead({ billing: FULL_GRANT_PLAN }));
    await expect(page.getByTestId('guild-billing-plan')).toHaveText('Guild plan · until 31 Dec 2026');
    await expect(page.getByTestId('guild-billing-contact')).toHaveText('Granted by Forever Sixty.');
    await expect(page.getByTestId('guild-billing-chooser')).toHaveCount(0);
    await expect(page.getByTestId('guild-billing-subscribe')).toHaveCount(0);
    await expect(page.getByTestId('guild-manage-billing')).toHaveCount(0);
  });
});

// A Stripe renewal reads on the viewer's calendar (launch pricing spec §6); the fixture's instants are
// midnight UTC, so the viewer is in UTC for the dates to read as written.
test.describe('the account page guild rows', () => {
  test.use({ timezoneId: 'UTC' });

  test('state 13: the account page shows each guild plan on its own row, never in Premium', async ({
    page,
  }) => {
    await stubAccountGuildRows(page);
    await gotoHydrated(page, '/account', 'account-guilds');
    const lines = page.getByTestId('account-guild-plan');
    await expect(page.getByTestId('account-guild-row')).toHaveCount(3);
    await expect(lines).toHaveCount(2);
    await expect(
      page
        .getByTestId('account-guild-row')
        .filter({ hasText: 'Ashen Vigil' })
        .getByTestId('account-guild-plan'),
    ).toHaveCount(0);
    await expect(lines.nth(0)).toHaveText('Officer tools · until 12 Dec 2026');
    await expect(lines.nth(1)).toHaveText('Renews 31 Dec 2026');
    await expect(guildSettingsLinkOf(lines.nth(0))).toHaveAttribute('href', GUILD_SETTINGS_PATH);
    await expect(guildSettingsLinkOf(lines.nth(1))).toHaveAttribute('href', SETTINGS_GUILD_PATH);
    await expect(page.getByTestId('account-premium-block').getByText('Officer tools')).toHaveCount(0);
  });
});
