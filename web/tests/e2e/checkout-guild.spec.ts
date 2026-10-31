// web/tests/e2e/checkout-guild.spec.ts
// Subscribe-and-join spec (design/specs/2026-10-27-subscribe-and-join.md) section 12, cases 1 to 13
// and 18: the checkout guild step, the API stubbed from subscribe-fixtures. Every asserted string
// comes from the copy modules. The premium owned panel (task 12) and the band reserve (task 15) ride along.
import type { Page } from '@playwright/test';
import { expect, test } from './support/test';
import communityLinks from '../../src/data/links.json' with { type: 'json' };
import type { Me } from '../../src/lib/account/me-types';
import { checkoutGuildCopy as copy } from '../../src/lib/billing/checkout-guild-copy';
import { checkoutCopy } from '../../src/lib/billing/copy';
import { priceFor } from '../../src/lib/billing/plan-costs';
import { PLANS, formatPrice, founderOpen } from '../../src/lib/billing/plans';
import { premiumSourceCopy } from '../../src/lib/billing/premium-source-copy';
import { uiCopy } from '../../src/lib/ui/copy';
import {
  CLAIMED_ASK_MEMBER,
  CLAIMED_ASK_OFFICER,
  CLAIM_BLOCKED_HOLDS_OTHER,
  EMPTY_A,
  EMPTY_B,
  MILLENNIALS,
  OLDER_ME,
  ON_PLAN_ENDING,
  ON_PLAN_OTHER_PAYS,
  ON_PLAN_SEAT,
  ON_PLAN_YOU_PAY,
  OWNER_GUILD_MASTER,
  OWNER_GUILD_MASTER_CLAIMED,
  OWNER_MEMBER,
  PREMIUM_FREE,
  PREMIUM_GRANT_OPEN_ENDED,
  PREMIUM_SEAT_NO_PORTAL,
  READY_GUILD_28,
  READY_RAID_GUILD_12,
  STEP_FIXTURES,
  OLYMPUS,
  stepGuild,
  subscribeView,
} from '../../src/test-support/subscribe-fixtures';
import {
  CHECKOUT_PATH,
  GUILD_ID,
  GUILD_URL,
  PORTAL_PATH,
  RAID_GUILD_URL,
  STRIPE_CHECKOUT_URL,
  checkoutReply,
  hintGuildCount,
  meWithGuilds,
  onPlanGuild,
  openStep,
  portalReply,
  readyGuild,
  refusal,
  routeMeCell,
  routePost,
  routeStripe,
} from './support/checkout-guild';
import { envelope } from './support/guild-billing';
import { drainShifts, recordShifts, sumShifts } from './support/layout-shift';
import { expectNoAxeViolations } from './support/officer-tools';
import { shortControls } from './support/short-controls';

// premium-owned.ts's ACCOUNT_PLAN_HREF: that module reaches lib/planner/config.ts, whose
// import.meta.env does not exist under Playwright's Node loader, so the spec cannot import it.
const ACCOUNT_PLAN_HREF = '/account#plan';
const GUILD_NAME = OLYMPUS.name;
const CLAIM_PATH = `**/v1/guilds/${GUILD_ID}/claim`;
const CLAIM_CONFIRM_PATH = `**/v1/guilds/${GUILD_ID}/claim/confirm`;
const FILL = '[class~="bg-gold"]';
const fills = (page: Page) => page.getByTestId('checkout-guild-step').locator(FILL);
const MAX_CLS = 0.02;
const BAND_TOLERANCE_PX = 1;
const id = (suffix: string): string => `guild-${GUILD_ID}-${suffix}`;
const founder = founderOpen(new Date());
const priceLabel = (plan: 'guild' | 'raid_guild'): string =>
  formatPrice(priceFor(plan, 'yearly', founder), 'yearly');

const TWO_CLAIMS: Me = meWithGuilds(
  stepGuild({ step: 'claim', rank: 'leader' }),
  stepGuild({ ...MILLENNIALS, step: 'claim', rank: 'leader' }),
);
const SECOND_ID = MILLENNIALS.id;

async function openClaim(page: Page, testid: string = id('claim-open')): Promise<void> {
  await page.getByTestId(testid).click();
  await expect(page.getByTestId(id('claim-confirm'))).toBeVisible();
}

// --- 1. the owner -------------------------------------------------------------------------------

test.describe('the owner, a member of two unclaimed guilds', () => {
  test('two cards in order, the hand-off, and no Get Premium', async ({ page }) => {
    await openStep(page, OWNER_MEMBER);
    await expect(page.getByTestId('checkout-guild-lead')).toHaveText(copy.lead.allUnclaimed(2));
    const cards = page.getByTestId('checkout-guild-card');
    await expect(cards).toHaveCount(2);
    await expect(cards.nth(0)).toContainText(GUILD_NAME);
    await expect(cards.nth(1)).toContainText(MILLENNIALS.name);

    const first = cards.nth(0);
    await expect(first.getByTestId('guild-you')).toContainText(copy.rank.member);
    await expect(first.getByTestId('guild-you')).toContainText(copy.verifyMark.notVerified);
    await expect(first.getByTestId('guild-facts')).toContainText(copy.facts.noneVerified);
    await expect(first.getByTestId('guild-facts')).toContainText(copy.facts.noPlan);
    await expect(first.getByTestId('guild-step-sentence')).toHaveText(copy.askLeader(GUILD_NAME));
    await expect(first.getByTestId(id('handoff-url'))).toHaveValue(
      new RegExp(`/guild/us/pvp/${OLYMPUS.slug}/claim\\?plan=guild&interval=yearly$`),
    );
    await expect(first.getByTestId('guild-verify-line')).toHaveText(copy.verify.short);
    const other = first.getByTestId('guild-other-character');
    await expect(other).toContainText(copy.otherCharacter(GUILD_NAME));
    await expect(other.getByRole('link', { name: copy.getSetUp })).toBeVisible();
    await expect(first.getByTestId('guild-moderator-fallback').getByRole('link')).toHaveAttribute(
      'href',
      communityLinks.discordInvite,
    );
    // The owner's second character sits in Employed Millennials, so that card names it too.
    const secondYou = cards.nth(1).getByTestId('guild-you');
    await expect(secondYou).toContainText('Bow Jackzon');
    await expect(secondYou).toContainText(copy.rank.member);
    await expect(secondYou).toContainText(copy.verifyMark.notVerified);
    await expect(page.getByRole('link', { name: copy.foot.getPremium })).toHaveCount(0);
    await expect(page.getByText(copy.facts.noPlan).first()).toBeVisible();
  });

  // --- 2. copy link ------------------------------------------------------------------------------

  test('Copy link writes the clipboard and reads Copied for two seconds', async ({
    page,
    context,
    browserName,
  }) => {
    test.skip(browserName !== 'chromium', 'clipboard permissions are Chromium-only here');
    await context.grantPermissions(['clipboard-read', 'clipboard-write']);
    await openStep(page, OWNER_MEMBER);
    const button = page.getByTestId(id('handoff-copy'));
    await expect(button).toHaveText(uiCopy.copyLink);
    await button.click();
    await expect(button).toHaveText(uiCopy.copiedLink);
    const written = await page.evaluate(() => navigator.clipboard.readText());
    expect(written).toBe(await page.getByTestId(id('handoff-url')).inputValue());
    await expect(button).toHaveText(uiCopy.copyLink, { timeout: 4000 });
  });

  test('a denied clipboard shows the failure line', async ({ page }) => {
    await openStep(page, OWNER_MEMBER);
    await page.evaluate(() => {
      Object.defineProperty(navigator, 'clipboard', {
        value: { writeText: () => Promise.reject(new Error('denied')) },
        configurable: true,
      });
    });
    await page.getByTestId(id('handoff-copy')).click();
    await expect(page.getByTestId(id('handoff-copy-error'))).toHaveText(uiCopy.copyLinkFailed);
  });
});

// --- 3, 4. fills --------------------------------------------------------------------------------

test.describe('ready cards', () => {
  test('one ready (trial unused): the only fill, and it posts plan, interval and guild id', async ({
    page,
  }) => {
    await routeStripe(page);
    const posts = await routePost(page, CHECKOUT_PATH, () => checkoutReply());
    await openStep(page, STEP_FIXTURES.ready);
    const button = page.getByTestId(id('subscribe'));
    await expect(button).toContainText(copy.ready.startTrial);
    await expect(fills(page)).toHaveCount(1);
    await expect(button).toHaveClass(/\bbg-gold\b/);
    await button.click();
    await page.waitForURL(STRIPE_CHECKOUT_URL);
    expect(posts).toEqual([{ plan: 'guild', interval: 'yearly', guild_id: GUILD_ID }]);
  });

  test('two ready: both gold outline, no fill on the page', async ({ page }) => {
    await openStep(page, meWithGuilds(readyGuild(), readyGuild({ ...MILLENNIALS })));
    await expect(page.getByTestId('checkout-guild-lead')).toHaveText(copy.lead.ready);
    await expect(page.locator('[data-testid$="-subscribe"]')).toHaveCount(2);
    await expect(fills(page)).toHaveCount(0);
    for (const testid of [id('subscribe'), `guild-${SECOND_ID}-subscribe`]) {
      await expect(page.getByTestId(testid)).toHaveClass(/\bborder-gold\b/);
    }
  });
});

// --- 5. claim in place --------------------------------------------------------------------------

test.describe('claim in place', () => {
  test('confirm opens with focus on Cancel; Escape closes and returns focus to the opener', async ({
    page,
  }) => {
    await openStep(page, OWNER_GUILD_MASTER);
    await openClaim(page);
    await expect(page.getByTestId(id('claim-cancel'))).toBeFocused();
    await expect(page.getByTestId(id('claim-confirm'))).toContainText(copy.confirm.claim(GUILD_NAME));
    await page.keyboard.press('Escape');
    await expect(page.getByTestId(id('claim-confirm'))).toHaveCount(0);
    await expect(page.getByTestId(id('claim-open'))).toBeFocused();
  });

  test('Yes, claim it posts, re-reads, makes the card ready and focuses its button', async ({ page }) => {
    const cell = await openStep(page, OWNER_GUILD_MASTER);
    const posts = await routePost(page, CLAIM_PATH, () => {
      cell.set(OWNER_GUILD_MASTER_CLAIMED);
      return envelope({ status: 'confirmed' });
    });
    await openClaim(page);
    await page.getByTestId(id('claim-commit')).click();
    await expect(page.getByTestId(id('subscribe'))).toBeVisible();
    await expect(page.getByTestId(id('subscribe'))).toBeFocused();
    await expect(page.getByTestId(id('notice'))).toHaveText(copy.done.claim(GUILD_NAME));
    expect(posts).toHaveLength(1);
  });

  test('during the re-read the opener is busy and the commit waits; one confirm at a time', async ({
    page,
  }) => {
    const cell = await openStep(page, TWO_CLAIMS);
    await openClaim(page);
    await expect(page.getByTestId(`guild-${SECOND_ID}-claim-open`)).toBeDisabled();
    await page.keyboard.press('Escape');
    await expect(page.getByTestId(`guild-${SECOND_ID}-claim-open`)).toBeEnabled();

    await routePost(page, CLAIM_PATH, () => envelope({ status: 'confirmed' }));
    await openClaim(page);
    const release = cell.hold();
    await page.getByTestId(id('claim-commit')).click();
    await expect(page.getByTestId(id('claim-open'))).toHaveAttribute('aria-busy', 'true');
    await expect(page.getByTestId(id('claim-commit'))).toBeDisabled();
    await expect(page.getByTestId(id('claim-commit'))).toHaveAttribute('aria-busy', 'true');
    cell.set(meWithGuilds(readyGuild(), stepGuild({ ...MILLENNIALS, step: 'claim', rank: 'leader' })));
    release();
    await expect(page.getByTestId(id('notice'))).toHaveText(copy.done.claim(GUILD_NAME));
  });
});

// --- 6. start claim, pending --------------------------------------------------------------------

test('start_claim then pending: claim_pending_yours with a copy field', async ({ page }) => {
  const cell = await openStep(page, STEP_FIXTURES.start_claim);
  await routePost(page, CLAIM_PATH, () => {
    cell.set(STEP_FIXTURES.claim_pending_yours);
    return envelope({ status: 'pending' });
  });
  await expect(page.getByTestId(id('claim-open'))).toHaveText(copy.startClaim.button);
  await openClaim(page);
  await expect(page.getByTestId(id('claim-confirm'))).toContainText(copy.confirm.startClaim(GUILD_NAME));
  await page.getByTestId(id('claim-commit')).click();
  await expect(page.getByTestId('guild-step-sentence')).toHaveText(copy.claimPendingYours(GUILD_NAME));
  await expect(page.getByTestId(id('handoff-url'))).toBeVisible();
  await expect(page.getByTestId(id('handoff-token'))).toContainText(copy.copyField.labelSecondOfficer);
  await expect(page.getByTestId(id('notice'))).toHaveText(copy.done.startClaim);
});

// --- 7. refusals --------------------------------------------------------------------------------

interface Refusal {
  name: string;
  start: Me;
  status: number;
  code: string;
  rereadsTo: Me | null;
  line: string;
  inBlock: boolean;
}

const REFUSALS: Refusal[] = [
  {
    name: 'already claimed',
    start: OWNER_GUILD_MASTER,
    status: 409,
    code: 'conflict',
    rereadsTo: STEP_FIXTURES.claimed_ask,
    line: copy.claimErrors.alreadyClaimed(GUILD_NAME),
    inBlock: false,
  },
  {
    name: 'not eligible',
    start: OWNER_GUILD_MASTER,
    status: 403,
    code: 'forbidden',
    rereadsTo: null,
    line: copy.claimErrors.notEligible(GUILD_NAME),
    inBlock: true,
  },
  {
    name: 'rate limited',
    start: OWNER_GUILD_MASTER,
    status: 429,
    code: 'rate_limited',
    rereadsTo: null,
    line: copy.claimErrors.rateLimited,
    inBlock: true,
  },
  {
    name: 'holds another',
    start: OWNER_GUILD_MASTER,
    status: 409,
    code: 'conflict',
    rereadsTo: CLAIM_BLOCKED_HOLDS_OTHER,
    line: copy.claimErrors.holdsAnother,
    inBlock: false,
  },
];

for (const refused of REFUSALS) {
  test(`claim refusal, ${refused.name}: the section 9.1 line and no API sentence`, async ({ page }) => {
    const cell = await openStep(page, refused.start);
    await routePost(page, CLAIM_PATH, () => {
      if (refused.rereadsTo !== null) cell.set(refused.rereadsTo);
      return refusal(refused.status, refused.code);
    });
    await openClaim(page);
    await page.getByTestId(id('claim-commit')).click();
    const line = page.getByTestId(id(refused.inBlock ? 'claim-error' : 'notice'));
    await expect(line).toHaveText(refused.line);
    await expect(line).toHaveAttribute('role', 'alert');
    await expect(page.locator('body')).not.toContainText('the api sentence');
    await expect(page.getByTestId(id('claim-confirm'))).toHaveCount(refused.inBlock ? 1 : 0);
  });
}

test('confirm_claim: Yes, confirm it posts the confirm route and says who holds the guild', async ({
  page,
}) => {
  const cell = await openStep(page, STEP_FIXTURES.confirm_claim);
  const posts = await routePost(page, CLAIM_CONFIRM_PATH, () => {
    cell.set(STEP_FIXTURES.claimed_ask);
    return envelope({ status: 'confirmed', claimed_by: { battletag: 'Thornwick#1' } });
  });
  await openClaim(page);
  await page.getByTestId(id('claim-commit')).click();
  await expect(page.getByTestId(id('notice'))).toHaveText(copy.done.confirmClaim(null, GUILD_NAME));
  expect(posts).toHaveLength(1);
});

// --- 8. claimed_ask -----------------------------------------------------------------------------

test('claimed_ask names the holder to an officer, not to a member', async ({ page }) => {
  await openStep(page, CLAIMED_ASK_OFFICER);
  await expect(page.getByTestId('guild-step-sentence')).toHaveText(
    copy.claimedAsk.officer('Thornwick', GUILD_NAME),
  );
  await expect(page.getByTestId(id('handoff-token'))).toContainText(copy.copyField.labelPerson('Thornwick'));

  await page.unrouteAll({ behavior: 'ignoreErrors' });
  await openStep(page, CLAIMED_ASK_MEMBER);
  await expect(page.getByTestId('guild-step-sentence')).toHaveText(copy.claimedAsk.member(GUILD_NAME));
  await expect(page.locator('body')).not.toContainText('Thornwick');
});

// --- 9. on_plan ---------------------------------------------------------------------------------

test.describe('on_plan', () => {
  test('you pay: Manage billing opens the portal for the guild', async ({ page }) => {
    await routeStripe(page);
    const posts = await routePost(page, PORTAL_PATH, () => portalReply());
    await openStep(page, ON_PLAN_YOU_PAY);
    await expect(page.getByTestId('guild-step-sentence')).toHaveText(copy.onPlan.youPay(GUILD_NAME, 'guild'));
    await page.getByTestId(id('manage')).click();
    await page.waitForURL('https://billing.stripe.com/**');
    expect(posts).toEqual([{ guild_id: GUILD_ID }]);
  });

  test('a failed portal shows its line on the card', async ({ page }) => {
    await routePost(page, PORTAL_PATH, () => refusal(500, 'internal'));
    await openStep(page, ON_PLAN_YOU_PAY);
    await page.getByTestId(id('manage')).click();
    await expect(page.getByTestId(id('notice'))).toHaveText(copy.portalError);
    await expect(page.getByTestId(id('manage'))).toBeEnabled();
  });

  test('you pay on Guild with Raid guild chosen: Move to Raid guild links to upgrade=1', async ({ page }) => {
    await openStep(page, ON_PLAN_YOU_PAY, RAID_GUILD_URL);
    const move = page.getByTestId(id('move'));
    await expect(move).toHaveText(new RegExp(copy.onPlan.moveToRaidGuild));
    await expect(move).toHaveAttribute('href', /\/settings\?upgrade=1#billing$/);
    await expect(page.getByTestId(id('manage'))).toHaveCount(0);
  });

  test('seat held, seats free, seats full, someone else pays', async ({ page }) => {
    await openStep(page, ON_PLAN_SEAT);
    await expect(page.getByTestId('checkout-guild-card')).toContainText(copy.onPlan.seatHeld);
    await expect(page.locator('[data-step-action]')).toHaveCount(0);

    await page.unrouteAll({ behavior: 'ignoreErrors' });
    await openStep(page, ON_PLAN_OTHER_PAYS);
    await expect(page.getByTestId('guild-step-sentence')).toHaveText(
      copy.onPlan.billedBy(GUILD_NAME, 'guild', 'Thornwick'),
    );
    await expect(page.getByTestId('checkout-guild-card')).toContainText(
      copy.onPlan.seatsGiven(0, 1, 'Thornwick'),
    );

    await page.unrouteAll({ behavior: 'ignoreErrors' });
    // Seat lines show only on a card the viewer does not pay for (guild-card-model seatLine).
    const otherPays = ON_PLAN_OTHER_PAYS.guilds[0]?.covered;
    await openStep(page, meWithGuilds(onPlanGuild({ ...otherPays, seats_given: 3, seats_total: 3 })));
    await expect(page.getByTestId('checkout-guild-card')).toContainText(copy.onPlan.seatsAllGiven(3));
  });
});

// --- 10. checkout_open and frozen ---------------------------------------------------------------

test.describe('checkout_open and frozen', () => {
  test('checkout_open: its sentence, no buy button, and Check again re-reads in place', async ({ page }) => {
    const cell = await openStep(page, STEP_FIXTURES.checkout_open);
    await expect(page.getByTestId('guild-step-sentence')).toHaveText(copy.checkoutOpen(GUILD_NAME));
    await expect(page.locator('[data-step-action]')).toHaveCount(0);
    cell.set(STEP_FIXTURES.ready);
    await page.getByTestId('guild-check-again').click();
    await expect(page.getByTestId(id('subscribe'))).toBeVisible();
    await expect(page.getByTestId('checkout-guild-loading')).toHaveCount(0);
  });

  test('a failed Check again says so and leaves the cards', async ({ page }) => {
    const cell = await openStep(page, STEP_FIXTURES.checkout_open);
    cell.failReads(true);
    await page.getByTestId('guild-check-again').click();
    await expect(page.getByTestId('checkout-refresh-error')).toHaveText(checkoutCopy.genericFailure);
    await expect(page.getByTestId('guild-step-sentence')).toHaveText(copy.checkoutOpen(GUILD_NAME));
  });

  test('frozen: its sentence and no buy button', async ({ page }) => {
    await openStep(page, STEP_FIXTURES.frozen);
    await expect(page.getByTestId('guild-step-sentence')).toHaveText(copy.frozen(GUILD_NAME));
    await expect(page.locator('[data-step-action]')).toHaveCount(0);
    await expect(page.getByRole('link', { name: copy.openGuild })).toBeVisible();
  });
});

// --- 11. guild_id arrival -----------------------------------------------------------------------

test.describe('guild_id', () => {
  test('a non-ready guild comes first and in view', async ({ page }) => {
    await page.setViewportSize({ width: 390, height: 600 });
    await openStep(page, TWO_CLAIMS, `${GUILD_URL}&guild_id=${SECOND_ID}`);
    const first = page.getByTestId('checkout-guild-card').first();
    await expect(first).toHaveAttribute('id', `guild-${SECOND_ID}`);
    await expect(first).toBeInViewport();
    const top = (await first.boundingBox())?.y ?? Number.POSITIVE_INFINITY;
    expect(top).toBeLessThan(300);
  });

  test('a ready guild launches Stripe on arrival, once, and shows the busy label', async ({ page }) => {
    await routeStripe(page);
    const posts = await routePost(page, CHECKOUT_PATH, () => checkoutReply());
    await routeMeCell(page, STEP_FIXTURES.ready);
    await page.goto(`${GUILD_URL}&guild_id=${GUILD_ID}`);
    await page.waitForURL(STRIPE_CHECKOUT_URL);
    expect(posts).toEqual([{ plan: 'guild', interval: 'yearly', guild_id: GUILD_ID }]);
  });

  test('an unknown guild shows the list and starts nothing', async ({ page }) => {
    const posts = await routePost(page, CHECKOUT_PATH, () => checkoutReply());
    await openStep(page, OWNER_MEMBER, `${GUILD_URL}&guild_id=999`);
    await expect(page.getByTestId('checkout-guild-card')).toHaveCount(2);
    expect(posts).toHaveLength(0);
  });

  test('back from Stripe (bfcache): busy clears and Subscribe works again, with no relaunch', async ({
    page,
  }) => {
    const posts = await routePost(page, CHECKOUT_PATH, () => checkoutReply('#stripe'));
    await routeMeCell(page, STEP_FIXTURES.ready);
    await page.goto(`${GUILD_URL}&guild_id=${GUILD_ID}`);
    const button = page.getByTestId(id('subscribe'));
    await expect(button).toHaveAttribute('aria-busy', 'true');
    await expect(button).toContainText(copy.busy.checkout);
    await page.evaluate(() => {
      window.dispatchEvent(new PageTransitionEvent('pageshow', { persisted: true }));
    });
    await expect(button).not.toHaveAttribute('aria-busy', 'true');
    await expect(button).toBeEnabled();
    await expect(button).toContainText(copy.ready.startTrial);
    expect(posts).toHaveLength(1);
  });
});

// --- checkout refusals on a card (task 18) -------------------------------------------------------

test.describe('checkout refusals', () => {
  test('503: purchases not open, nothing re-read, button usable', async ({ page }) => {
    await routePost(page, CHECKOUT_PATH, () => refusal(503, 'billing_unavailable'));
    await openStep(page, STEP_FIXTURES.ready);
    await page.getByTestId(id('subscribe')).click();
    await expect(page.getByTestId(id('notice'))).toHaveText(copy.checkoutErrors.notOpen);
    await expect(page.getByTestId(id('subscribe'))).toBeEnabled();
  });

  test('409: already on a plan, re-read, focus on the card new action', async ({ page }) => {
    const cell = await openStep(page, STEP_FIXTURES.ready);
    await routePost(page, CHECKOUT_PATH, () => {
      cell.set(ON_PLAN_YOU_PAY);
      return refusal(409, 'conflict');
    });
    await page.getByTestId(id('subscribe')).click();
    await expect(page.getByTestId(id('manage'))).toBeFocused();
    await expect(page.getByTestId(id('notice'))).toHaveText(copy.checkoutErrors.alreadyOnPlan(GUILD_NAME));
  });

  test('403: re-read with no line of its own', async ({ page }) => {
    const cell = await openStep(page, STEP_FIXTURES.ready);
    await routePost(page, CHECKOUT_PATH, () => {
      cell.set(STEP_FIXTURES.frozen);
      return refusal(403, 'forbidden');
    });
    await page.getByTestId(id('subscribe')).click();
    await expect(page.getByTestId('guild-step-sentence')).toHaveText(copy.frozen(GUILD_NAME));
    await expect(page.getByTestId(id('notice'))).toHaveCount(0);
  });

  test('generic failure: the line and Try again, which posts again', async ({ page }) => {
    const posts = await routePost(page, CHECKOUT_PATH, () => refusal(500, 'internal'));
    await openStep(page, STEP_FIXTURES.ready);
    await page.getByTestId(id('subscribe')).click();
    await expect(page.getByTestId(id('notice'))).toHaveText(copy.checkoutErrors.generic);
    await page.getByTestId(id('try-again')).click();
    await expect.poll(() => posts.length).toBe(2);
  });
});

// --- 12. size lines and switching ---------------------------------------------------------------

test.describe('size lines', () => {
  test('Guild and Raid guild each state their own size, and 40 never appears', async ({ page }) => {
    await openStep(page, OWNER_MEMBER);
    await expect(page.getByTestId('checkout-size-line')).toHaveText(copy.sizeLine.guild);
    await page.goto(RAID_GUILD_URL);
    await expect(page.getByTestId('checkout-size-line')).toHaveText(copy.sizeLine.raidGuild);
    await expect(page.locator('body')).not.toContainText(/40-person|40 players/);
  });

  test('12 verified on Raid guild: Switch to Guild changes band, button and URL with no reload', async ({
    page,
  }) => {
    const subscribed = meWithGuilds(readyGuild({ subscribe: subscribeView('ready', { trial_used: true }) }));
    await openStep(page, subscribed, RAID_GUILD_URL);
    await expect(page.getByTestId(id('subscribe'))).toContainText(
      copy.ready.subscribe(priceLabel('raid_guild')),
    );
    await page.evaluate(() => {
      (window as unknown as { __stay: boolean }).__stay = true;
    });
    await page.getByRole('button', { name: copy.switchPlan.toGuild }).click();
    await expect(page).toHaveURL(/plan=guild&/);
    await expect(page.getByTestId('checkout-size-line')).toHaveText(copy.sizeLine.guild);
    await expect(page.getByTestId('checkout-plan-summary')).toContainText(PLANS.guild.name);
    await expect(page.getByTestId(id('subscribe'))).toContainText(copy.ready.subscribe(priceLabel('guild')));
    expect(await page.evaluate(() => (window as unknown as { __stay?: boolean }).__stay)).toBe(true);
    await expect(page.getByTestId(id('switch-note'))).toHaveCount(0);
  });

  test('28 verified on Guild: Switch to Raid guild, and Get set up then carries plan=raid_guild', async ({
    page,
  }) => {
    await openStep(page, READY_GUILD_28);
    await page.getByRole('button', { name: copy.switchPlan.toRaidGuild }).click();
    await expect(page).toHaveURL(/plan=raid_guild/);
    await expect(page.getByTestId('checkout-size-line')).toHaveText(copy.sizeLine.raidGuild);
    const setUp = page.getByTestId('checkout-guild-foot').getByRole('link', { name: copy.foot.getSetUp });
    expect(decodeURIComponent((await setUp.getAttribute('href')) ?? '')).toContain('plan=raid_guild');
  });

  test('the switch note names 12 verified on Raid guild', async ({ page }) => {
    await openStep(page, READY_RAID_GUILD_12, RAID_GUILD_URL);
    await expect(page.getByTestId(id('switch-note'))).toContainText(copy.switchPlan.withinLimit(12));
  });

  test('the md band reserve: the Guild size line fits on one line at 768', async ({ page }) => {
    await page.setViewportSize({ width: 768, height: 900 });
    await openStep(page, OWNER_MEMBER);
    const sizeLine = page.getByTestId('checkout-size-line');
    const { height, lineHeight } = await sizeLine.evaluate((node) => ({
      height: node.getBoundingClientRect().height,
      lineHeight: Number.parseFloat(getComputedStyle(node).lineHeight),
    }));
    expect(height / lineHeight).toBeLessThan(1.5);
  });
});

// --- band reserve (task 15) ---------------------------------------------------------------------

for (const width of [390, 768, 800, 1440]) {
  test(`the band is the same height before and after the summary appears at ${width}`, async ({
    page,
    browser,
    baseURL,
  }) => {
    const viewport = { width, height: 900 };
    const bare = await browser.newContext({ baseURL, viewport, javaScriptEnabled: false });
    const barePage = await bare.newPage();
    await barePage.goto(GUILD_URL);
    const before = (await barePage.getByTestId('checkout-band').boundingBox())?.height ?? 0;
    await bare.close();

    await page.setViewportSize(viewport);
    await openStep(page, OWNER_MEMBER);
    await expect(page.getByTestId('checkout-plan-summary')).toBeVisible();
    const after = (await page.getByTestId('checkout-band').boundingBox())?.height ?? -1;
    expect(Math.abs(after - before)).toBeLessThanOrEqual(BAND_TOLERANCE_PX);
  });
}

// --- 13. empty A and B --------------------------------------------------------------------------

test.describe('empty states', () => {
  for (const [name, me, testid, text] of [
    ['A (no characters)', EMPTY_A, 'checkout-guild-empty-a', copy.emptyA],
    ['B (no guild)', EMPTY_B, 'checkout-guild-empty-b', copy.emptyB],
  ] as const) {
    test(`empty ${name}: Get set up returns here, and Battle.net is never named`, async ({ page }) => {
      await openStep(page, me);
      const panel = page.getByTestId(testid);
      await expect(panel).toContainText(text.title);
      const link = panel.getByRole('link', { name: text.button });
      await expect(link).toHaveAttribute('href', `/setup?next=${encodeURIComponent(GUILD_URL)}`);
      await expect(page.locator('body')).not.toContainText(/Battle\.net/i);
    });
  }

  test('an older API (no subscribe) draws identity rows and Open the guild only', async ({ page }) => {
    await openStep(page, OLDER_ME);
    await expect(page.getByRole('link', { name: copy.openGuild })).toBeVisible();
    await expect(page.locator('[data-step-action]')).toHaveCount(0);
  });
});

// --- premium owned (task 12) --------------------------------------------------------------------

test.describe('premium checkout for an account that has Premium', () => {
  const PREMIUM_URL = '/premium/checkout?plan=premium&interval=monthly';

  test('a grant shows "You already have Premium" and never posts checkout', async ({ page }) => {
    const posts = await routePost(page, CHECKOUT_PATH, () => checkoutReply());
    await routeMeCell(page, PREMIUM_GRANT_OPEN_ENDED);
    await page.goto(PREMIUM_URL);
    const panel = page.getByTestId('premium-owned');
    await expect(panel).toContainText(premiumSourceCopy.panelTitle);
    await expect(panel).toContainText(premiumSourceCopy.grant.noEnd);
    await expect(panel.getByTestId('premium-owned-action')).toHaveAttribute('href', ACCOUNT_PLAN_HREF);
    expect(posts).toHaveLength(0);
  });

  test('a 409 carrying fields.source re-reads and shows the panel, never the API sentence', async ({
    page,
  }) => {
    const cell = await routeMeCell(page, PREMIUM_FREE);
    const posts = await routePost(page, CHECKOUT_PATH, () => {
      cell.set(PREMIUM_SEAT_NO_PORTAL);
      return refusal(409, 'conflict', { source: 'seat' });
    });
    await page.goto(PREMIUM_URL);
    await expect(page.getByTestId('premium-owned')).toContainText(premiumSourceCopy.panelTitle);
    await expect(page.locator('body')).not.toContainText('the api sentence');
    expect(posts).toHaveLength(1);
  });
});

// --- 18. axe, 44 px at 360, CLS -----------------------------------------------------------------

const A11Y_STATES: ReadonlyArray<readonly [string, Me]> = [
  ...Object.entries(STEP_FIXTURES),
  ['owner member', OWNER_MEMBER],
  ['owner guild master', OWNER_GUILD_MASTER],
  ['claimed_ask officer', CLAIMED_ASK_OFFICER],
  ['claimed_ask member', CLAIMED_ASK_MEMBER],
  ['claim_blocked holds_other', CLAIM_BLOCKED_HOLDS_OTHER],
  ['on_plan seat', ON_PLAN_SEAT],
  ['on_plan other pays', ON_PLAN_OTHER_PAYS],
  ['on_plan ending', ON_PLAN_ENDING],
  ['ready on Guild with 28', READY_GUILD_28],
  ['two ready', meWithGuilds(readyGuild(), readyGuild({ ...MILLENNIALS }))],
  ['empty A', EMPTY_A],
  ['empty B', EMPTY_B],
  ['older API', OLDER_ME],
];

test.describe('accessibility and touch targets at 360', () => {
  test.beforeEach(async ({ page }) => {
    await page.setViewportSize({ width: 360, height: 740 });
  });

  for (const [name, me] of A11Y_STATES) {
    test(`${name}: axe clean, every control at least 44 px`, async ({ page }) => {
      await openStep(page, me);
      expect(await shortControls(page.getByTestId('checkout-guild-step'))).toEqual([]);
      await expectNoAxeViolations(page, 'main');
    });
  }

  test('an open claim confirm: axe clean, controls at least 44 px', async ({ page }) => {
    await openStep(page, OWNER_GUILD_MASTER);
    await openClaim(page);
    expect(await shortControls(page.getByTestId('checkout-guild-step'))).toEqual([]);
    await expectNoAxeViolations(page, 'main');
  });

  test('the premium owned panel: axe clean', async ({ page }) => {
    await routeMeCell(page, PREMIUM_GRANT_OPEN_ENDED);
    await page.goto('/premium/checkout?plan=premium&interval=monthly');
    await expect(page.getByTestId('premium-owned')).toBeVisible();
    await expectNoAxeViolations(page, 'main');
  });
});

test.describe('layout shift from loading to ready', () => {
  const SCENES: ReadonlyArray<readonly [string, Me, number]> = [
    ['the owner (two ask_leader cards)', OWNER_MEMBER, 2],
    ['one ready card', STEP_FIXTURES.ready, 1],
  ];
  const READ_DELAY_MS = 400;

  for (const width of [390, 1440]) {
    for (const [name, me, count] of SCENES) {
      test(`${name} at ${width}: CLS at most ${MAX_CLS}`, async ({ page }) => {
        await page.setViewportSize({ width, height: 900 });
        await hintGuildCount(page, count);
        await recordShifts(page);
        await routeMeCell(page, me, READ_DELAY_MS);
        await page.goto(GUILD_URL);
        await expect(page.getByTestId('checkout-guild-lead')).toBeVisible();
        const cls = sumShifts(await drainShifts(page));
        test.info().annotations.push({ type: 'cls', description: cls.toFixed(4) });
        expect(cls).toBeLessThanOrEqual(MAX_CLS);
      });
    }
  }
});
