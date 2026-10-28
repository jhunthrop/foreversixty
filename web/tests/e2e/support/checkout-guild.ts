// web/tests/e2e/support/checkout-guild.ts
// The /premium/checkout guild step, API stubbed (spec 2026-10-27-subscribe-and-join section 5): a
// mutable /v1/me cell the claim and refresh paths can change or hold, recorded POST routes, the
// Stripe hosts, and the fixture builders the step specs share that subscribe-fixtures does not export.
import { expect, type Page } from '@playwright/test';
import type { Me, MeGuild } from '../../../src/lib/account/me-types';
import type { GuildCoveredView, SubscribeStep } from '../../../src/lib/billing/contract';
import {
  OLYMPUS,
  buildMe,
  coveredView,
  stepGuild,
  subscribeView,
} from '../../../src/test-support/subscribe-fixtures';
import { envelope } from './guild-billing';

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

/** The `/v1/me` answers a spec can steer: swap the account, hold the read, or fail it. */
export interface MeCell {
  set: (me: Me) => void;
  /** Holds every read until the returned release runs; the re-read after a claim is the one to hold. */
  hold: () => () => void;
  failReads: (failing: boolean) => void;
  reads: () => number;
}

export async function routeMeCell(page: Page, initial: Me, delayMs = 0): Promise<MeCell> {
  let me = initial;
  let gate: Promise<void> | null = null;
  let failing = false;
  let reads = 0;
  await page.route('**/v1/me', async (route) => {
    reads += 1;
    if (delayMs > 0) await new Promise((resolve) => setTimeout(resolve, delayMs));
    if (gate !== null) await gate;
    if (failing) return route.fulfill(envelope({ code: 'internal', message: 'api sentence' }, 500));
    return route.fulfill(envelope(me));
  });
  return {
    set: (next) => {
      me = next;
    },
    hold: () => {
      let release = (): void => undefined;
      gate = new Promise<void>((resolve) => {
        release = resolve;
      });
      return () => {
        release();
        gate = null;
      };
    },
    failReads: (next) => {
      failing = next;
    },
    reads: () => reads,
  };
}

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
      sessionStorage.setItem(key, value);
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

// --- fixture builders the shared fixtures do not export -----------------------------------------

/** A `ready` guild the viewer holds the claim of; `verified` members, trial unused unless overridden. */
export function readyGuild(overrides: Partial<MeGuild> = {}, trialUsed = false): MeGuild {
  return stepGuild({
    rank: 'officer',
    verified: true,
    verified_count: 12,
    claim: { state: 'claimed', yours: true, pending_yours: false, holder_name: 'Obnoxious Yell' },
    covered: coveredView({ tier: 'guild', state: 'lapsed', ends_at: null, billed_by: '' }),
    subscribe: subscribeView('ready', { trial_used: trialUsed }),
    ...overrides,
  });
}

/** An `on_plan` guild the viewer is an officer of; `covered` picks the variant. */
export function onPlanGuild(covered: Partial<GuildCoveredView>, overrides: Partial<MeGuild> = {}): MeGuild {
  return stepGuild({
    rank: 'officer',
    verified: true,
    verified_count: 12,
    claim: { state: 'claimed', yours: false, pending_yours: false, holder_name: 'Thornwick' },
    covered: coveredView(covered),
    subscribe: subscribeView('on_plan'),
    ...overrides,
  });
}

/** One guild on a step (the owner's OLYMPUS XXVII by default) as the whole account. */
export function meWithGuilds(...guilds: MeGuild[]): Me {
  return buildMe({ guilds });
}

export const stepOnly = (step: SubscribeStep, overrides: Partial<MeGuild> = {}): Me =>
  meWithGuilds(stepGuild({ step, ...overrides }));
