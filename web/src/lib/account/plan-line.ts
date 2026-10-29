// web/src/lib/account/plan-line.ts
// The account Plan panel's headline (design/specs/2026-10-27-subscribe-and-join.md 4.1 and 4.3):
// which one source pays for the account's Premium and the words for it. Pure: the caller passes
// the clock and the viewer's zone.
import { accountPlanCopy } from '../billing/account-plan-copy';
import type { PremiumSource } from '../billing/contract';
import { billingBlockCopy } from '../billing/copy';
import { guildPlanLine } from '../billing/guild-billing-copy';
import { premiumSourceCopy } from '../billing/premium-source-copy';
import { zonedDayMonthYear } from '../dates';
import { settingsWords } from '../guild/settings-view';
import type { BillingView, GuildBillingView, Me } from './api';

export interface EndedSeatLine {
  /** The sentence before the link, with its trailing space. */
  text: string;
  /** The link's words. */
  action: string;
  href: string;
}

export type BillingButton = 'manage' | 'history';

/** The plan card's head: the name, the date it renews or ends, and one muted note. */
export interface PlanHead {
  name: string;
  when: { label: string; date: string } | null;
  /** The price, or what follows the plan (`Then a seat from X`); null when there is nothing to add. */
  extra: string | null;
}

export interface AccountPlanLine {
  head: PlanHead;
  line: string;
  line2: string | null;
  founder: boolean;
  billingButton: BillingButton | null;
  ended: EndedSeatLine | null;
}

type LineWords = Pick<AccountPlanLine, 'line' | 'line2'>;

const LIVE_STATUSES: readonly string[] = ['active', 'trialing', 'past_due'];

function startsInFuture(startsAt: string | null, now: Date): boolean {
  return startsAt !== null && new Date(startsAt).getTime() > now.getTime();
}

function liveBilling(me: Me): BillingView | null {
  const billing = me.entitlements?.billing ?? null;
  return billing !== null && LIVE_STATUSES.includes(billing.status) ? billing : null;
}

/**
 * Whether the row is an operator grant. An API that predates `billing.source` sends grants
 * without it, but a grant never carries a price, so a source-less row is Stripe only when it
 * does.
 */
function isGrant(billing: BillingView): boolean {
  if (billing.source !== undefined) return billing.source === 'grant';
  return billing.amount_cents === null || billing.interval === null;
}

/** Section 4.1: exactly one source pays for the account's Premium; the first match wins. */
export function premiumSource(me: Me, now: Date): PremiumSource {
  const billing = liveBilling(me);
  const seat = me.seat ?? null;
  if (billing !== null && !isGrant(billing)) {
    return seat !== null && startsInFuture(seat.starts_at, now) ? 'own_then_seat' : 'own';
  }
  if (billing !== null) return 'grant';
  if (seat !== null) return 'seat';
  return me.entitlements?.server_sims === true ? 'fallback' : 'none';
}

function ownWords(billing: BillingView, zone?: string): LineWords {
  const end = billing.current_period_end;
  // A live Stripe subscription always has a period end; this guards bad data from reading as a plan.
  if (end === null) return { line: premiumSourceCopy.fallback, line2: null };
  if (billing.cancel_at_period_end) return { line: accountPlanCopy.planLine.ends(end, zone), line2: null };
  const line =
    billing.amount_cents === null || billing.interval === null
      ? accountPlanCopy.planLine.renewsNoPrice(end, zone)
      : accountPlanCopy.planLine.renews(end, billing.amount_cents, billing.interval, zone);
  return { line, line2: null };
}

function ownThenSeatWords(me: Me, billing: BillingView, zone?: string): LineWords {
  const end = billing.current_period_end;
  const seat = me.seat;
  if (end === null || seat === null || seat === undefined) return ownWords(billing, zone);
  return { line: accountPlanCopy.planLine.endsThenSeat(end, seat.guild_name, zone), line2: null };
}

function grantWords(billing: BillingView, zone?: string): LineWords {
  const end = billing.current_period_end;
  return end === null
    ? { line: premiumSourceCopy.grant.noEnd, line2: premiumSourceCopy.grant.givenNoEnd }
    : { line: premiumSourceCopy.grant.until(end, zone), line2: premiumSourceCopy.grant.givenUntil };
}

function seatWords(me: Me): LineWords {
  const seat = me.seat;
  if (seat === null || seat === undefined) return { line: premiumSourceCopy.fallback, line2: null };
  return {
    line: accountPlanCopy.planLine.seat(seat.guild_name),
    line2: accountPlanCopy.planLine.givenBy(seat.given_by),
  };
}

function sourceWords(source: PremiumSource, me: Me, zone?: string): LineWords {
  const billing = liveBilling(me);
  if (billing !== null && source === 'own') return ownWords(billing, zone);
  if (billing !== null && source === 'own_then_seat') return ownThenSeatWords(me, billing, zone);
  if (billing !== null && source === 'grant') return grantWords(billing, zone);
  if (source === 'seat') return seatWords(me);
  if (source === 'fallback') return { line: premiumSourceCopy.fallback, line2: null };
  return { line: billingBlockCopy.notSubscribed, line2: null };
}

/** `Manage billing` only for the account's own subscription; `Billing history` for the other sources that have had one. */
function billingButton(source: PremiumSource, me: Me): BillingButton | null {
  if (source === 'own' || source === 'own_then_seat') return 'manage';
  return me.entitlements?.personal_portal === true && source !== 'fallback' ? 'history' : null;
}

function endedSeatLine(me: Me, zone?: string): EndedSeatLine | null {
  const ended = me.ended_seat;
  if (ended === null || ended === undefined) return null;
  const action =
    ended.amount_cents !== null && ended.interval !== null
      ? accountPlanCopy.seatEnded.action(ended.amount_cents, ended.interval)
      : accountPlanCopy.seatEnded.actionNoPrice;
  return {
    text: accountPlanCopy.seatEnded.text(ended.guild_name, ended.ended_at, zone),
    action,
    href: ended.resubscribe_url,
  };
}

function headWords(source: PremiumSource, me: Me, zone?: string): PlanHead {
  const words = accountPlanCopy.head;
  const billing = liveBilling(me);
  const end = billing?.current_period_end ?? null;
  const date = end === null ? null : zonedDayMonthYear(end, zone);
  if ((source === 'own' || source === 'own_then_seat') && billing !== null && date !== null) {
    const price =
      billing.amount_cents === null || billing.interval === null
        ? null
        : words.price(billing.amount_cents, billing.interval);
    const label = billing.cancel_at_period_end || source === 'own_then_seat' ? words.ends : words.renews;
    const seat = me.seat ?? null;
    const extra = source === 'own_then_seat' && seat !== null ? words.thenSeat(seat.guild_name) : price;
    return { name: words.names.premium, when: { label, date }, extra };
  }
  if (source === 'grant') {
    return {
      name: words.names.premium,
      when: date === null ? null : { label: words.until, date },
      extra: null,
    };
  }
  if (source === 'seat') {
    const guild = me.seat?.guild_name;
    return { name: words.names.seat, when: null, extra: guild === undefined ? null : words.seatFrom(guild) };
  }
  if (source === 'fallback') return { name: words.names.premium, when: null, extra: null };
  return { name: words.names.free, when: null, extra: null };
}

export function accountPlanLine(me: Me, now: Date, zone?: string): AccountPlanLine {
  const source = premiumSource(me, now);
  const owns = source === 'own' || source === 'own_then_seat';
  return {
    ...sourceWords(source, me, zone),
    head: headWords(source, me, zone),
    founder: owns && me.entitlements?.billing?.founder === true,
    billingButton: billingButton(source, me),
    ended: source === 'none' ? endedSeatLine(me, zone) : null,
  };
}

/**
 * A guild block's plan value: the verb and date (`Renews 12 Dec 2026`), `Active` with no date, plus
 * the seat note when this account holds one. The tier is the plan tag beside the guild name, so
 * the words never repeat it.
 */
export function guildRowPlanLine(plan: GuildBillingView, zone?: string): string {
  const base = guildPlanLine(plan, zone);
  const prefix = `${settingsWords.guildPlan} · `;
  const value = base === settingsWords.guildPlan ? accountPlanCopy.guildRow.active : base;
  const bare = value.startsWith(prefix) ? value.slice(prefix.length) : value;
  const sentence = `${bare.charAt(0).toUpperCase()}${bare.slice(1)}`;
  return plan.you_hold_seat ? `${sentence}${accountPlanCopy.guildRow.youHoldSeat}` : sentence;
}

/**
 * Whether the account is on Premium by any road: its own plan, a held seat, a grant, or the
 * resolved Premium features. Such an account never sees the free plan line's See plans link.
 */
export function onPremium(me: Me | null, now: Date = new Date()): boolean {
  return me !== null && premiumSource(me, now) !== 'none';
}
