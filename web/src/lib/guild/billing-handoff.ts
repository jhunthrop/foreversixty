// web/src/lib/guild/billing-handoff.ts
// Spec 2026-10-27 §6.4: an officer who does not hold the claim, in a guild with no live plan,
// gets a link to send the claim holder; and §6.4's `?plan&interval&upgrade` the settings page
// reads. Pure, so both are pinned without a page.
import type { GuildBillingSettings, GuildTier } from '../billing/contract';
import { parseInterval, parsePlan } from '../billing/plan-params';
import type { Interval } from '../billing/plans';
import { handoffUrl } from '../billing/subscribe-step';
import type { CharacterPath } from '../characters';
import {
  canBill,
  fitsTier,
  movesFromPlanLine,
  showSizeLine,
  showsChooser,
  showsReadOnlyLine,
  type BillingViewState,
} from './billing-state';

const HANDOFF_INTERVAL: Interval = 'yearly';

/** The line and link draw for a read-only reader of a guild with no live plan (none, lapsed, scoped grant). */
export function showsBillingHandoff(view: GuildBillingSettings, state: BillingViewState): boolean {
  return showsReadOnlyLine(view, canBill(view)) && showsChooser(state);
}

/** The settings `#billing` link for the claim holder, naming the plan that fits the verified count. */
export function billingHandoffUrl(path: CharacterPath, verifiedCount: number): string {
  return handoffUrl(path, 'claimed_ask', fitsTier(verifiedCount), HANDOFF_INTERVAL);
}

export interface BillingPreselect {
  tier: GuildTier | null;
  interval: Interval | null;
  /** `upgrade=1`: open the Move to Raid guild confirm. */
  upgrade: boolean;
}

/** `plan`, `interval` and `upgrade` off the settings URL; anything outside the closed sets is dropped. */
export function parseBillingPreselect(search: string): BillingPreselect {
  const query = new URLSearchParams(search);
  const plan = parsePlan(query.get('plan'));
  return {
    tier: plan === 'guild' || plan === 'raid_guild' ? plan : null,
    interval: parseInterval(query.get('interval')) ?? null,
    upgrade: query.get('upgrade') === '1',
  };
}

/** The confirm opens once, when a Move to Raid guild button is drawn for this reader (plan line or size line). */
export function shouldOpenUpgrade(
  wanted: boolean,
  view: GuildBillingSettings,
  state: BillingViewState,
  alreadyOpened: boolean,
): boolean {
  const moveDrawn = movesFromPlanLine(view, state) || (canBill(view) && showSizeLine(view));
  return wanted && !alreadyOpened && moveDrawn;
}

/** The chooser's starting props off the settings link: `?plan` picks the card, `?interval` the toggle. */
export function chooserPreselect(preselect: BillingPreselect): {
  initialTier: GuildTier | null;
  initialInterval: Interval;
} {
  return { initialTier: preselect.tier, initialInterval: preselect.interval ?? HANDOFF_INTERVAL };
}
