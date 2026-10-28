// web/tests/e2e/support/checkout-guild.ts
// The /premium/checkout guild step, API stubbed (spec 2026-10-27-subscribe-and-join section 5): a
// mutable /v1/me cell the claim and refresh paths can change or hold, recorded POST routes, the
// Stripe hosts, and the fixture builders the step specs share.
import { expect, type Page } from '@playwright/test';
import type { Me, MeGuild } from '../../../src/lib/account/me-types';
import type { GuildCoveredView } from '../../../src/lib/billing/contract';
import { OLYMPUS, STEP_FIXTURES, buildMe, coveredView } from '../../../src/test-support/subscribe-fixtures';
import { envelope } from './guild-billing';
import { routeMeCell, type MeCell } from './subscribe-stubs';

export { routeMeCell };

export const GUILD_URL = '/premium/checkout?plan=guild&interval=yearly';
export const RAID_GUILD_URL = '/premium/checkout?plan=raid_guild&interval=yearly';
export const GUILD_ID = OLYMPUS.id;
export const CHECKOUT_PATH = '**/v1/billing/checkout';
export const PORTAL_PATH = '**/v1/billing/portal';
export const GUILD_COUNT_HINT_KEY = 'fs.meGuildCount';

const STRIPE_CHECKOUT_HOST = 'https://checkout.stripe.com/**';
const STRIPE_PORTAL_HOST = 'https://billing.stripe.com/**';
export const STRIPE_CHECKOUT_URL = 'https://checkout.stripe.com/x';
export const STRIPE_PORTAL_URL = 'https://billing.stripe.com/x';

type Reply = ReturnType<typeof envelope>;

/** Answers a POST route and records each request body (null when the request carried none). */
export async function routePost(
  page: Page,
  glob: string,
  reply: (body: unknown) => Reply | Promise<Reply>,
): Promise<unknown[]> {
  const bodies: unknown[] = [];
  await page.route(glob, async (route) => {
    const body = route.request().postDataJSON() as unknown;
    bodies.push(body);
    return route.fulfill(await reply(body));
  });
  return bodies;
}

export const checkoutReply = (url: string = STRIPE_CHECKOUT_URL): Reply => envelope({ checkout_url: url });
export const portalReply = (url: string = STRIPE_PORTAL_URL): Reply => envelope({ portal_url: url });
export const refusal = (status: number, code: string, fields?: Record<string, string>): Reply =>
  envelope({ code, message: 'the api sentence, never shown', ...(fields ? { fields } : {}) }, status);

/** The hosts a redirect leaves for: fulfilled, so the navigation can be waited on. */
export async function routeStripe(page: Page): Promise<void> {
  for (const host of [STRIPE_CHECKOUT_HOST, STRIPE_PORTAL_HOST]) {
    await page.route(host, (route) => route.fulfill({ status: 200, body: 'ok' }));
  }
}

/** Seeds the session's guild-count hint, as a returning player's browser holds it. */
export async function hintGuildCount(page: Page, count: number): Promise<void> {
  await page.addInitScript(
    ([key, value]) => {
      localStorage.setItem(key, value);
    },
    [GUILD_COUNT_HINT_KEY, String(count)] as const,
  );
}

const STEP_READY = 'checkout-guild-lead';
const EMPTY_STATES = 'checkout-guild-empty-a|checkout-guild-empty-b';

/** Stubs /v1/me and opens the step, waiting for the cards (or the empty panel) to replace the skeleton. */
export async function openStep(page: Page, me: Me, url: string = GUILD_URL): Promise<MeCell> {
  const cell = await routeMeCell(page, me);
  await page.goto(url);
  await expect(page.getByTestId(new RegExp(`^(${STEP_READY}|${EMPTY_STATES})$`)).first()).toBeVisible();
  return cell;
}

// --- fixture builders ---------------------------------------------------------------------------

/** The fixtures' `ready` guild (trial unused, no plan), with overrides. */
export const readyGuild = (overrides: Partial<MeGuild> = {}): MeGuild => ({
  ...STEP_FIXTURES.ready.guilds[0]!,
  ...overrides,
});

/** The fixtures' `on_plan` guild the viewer pays for, with the `covered` variant given. */
export function onPlanGuild(covered: Partial<GuildCoveredView>): MeGuild {
  const base = STEP_FIXTURES.on_plan.guilds[0]!;
  return { ...base, covered: coveredView({ ...base.covered, ...covered }) };
}

export const meWithGuilds = (...guilds: MeGuild[]): Me => buildMe({ guilds });
