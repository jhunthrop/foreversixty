// web/src/lib/account/plan-line.ts
// The account Plan panel's headline: which of five states the account is in (own Premium, own
// Premium ending into a seat, a seat, free, a seat that ended) and the words for it, from
// accountPlanCopy. Pure: the caller passes the clock and the viewer's zone.
import { accountPlanCopy } from '../billing/account-plan-copy';
import { billingBlockCopy } from '../billing/copy';
import { guildPlanLine } from '../billing/guild-billing-copy';
import { settingsWords } from '../guild/settings-view';
import type { GuildBillingView, Me } from './api';

export interface EndedSeatLine {
  /** The sentence before the link, with its trailing space. */
  text: string;
  /** The link's words. */
  action: string;
  href: string;
}

export interface AccountPlanLine {
  line: string;
  line2: string | null;
  founder: boolean;
  manageBilling: boolean;
  ended: EndedSeatLine | null;
}

function startsInFuture(startsAt: string | null, now: Date): boolean {
  return startsAt !== null && new Date(startsAt).getTime() > now.getTime();
}

function ownPremiumLine(me: Me, zone?: string): string {
  const billing = me.entitlements?.billing;
  const end = billing?.current_period_end ?? null;
  if (billing === null || billing === undefined || end === null) return billingBlockCopy.notSubscribed;
  if (billing.cancel_at_period_end) return accountPlanCopy.planLine.ends(end, zone);
  if (billing.amount_cents === null || billing.interval === null) {
    return accountPlanCopy.planLine.renewsNoPrice(end, zone);
  }
  return accountPlanCopy.planLine.renews(end, billing.amount_cents, billing.interval, zone);
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

export function accountPlanLine(me: Me, now: Date, zone?: string): AccountPlanLine {
  const billing = me.entitlements?.billing ?? null;
  const seat = me.seat ?? null;
  const founder = billing?.founder === true;

  if (seat !== null && billing !== null && startsInFuture(seat.starts_at, now)) {
    const end = billing.current_period_end;
    return {
      line:
        end === null
          ? accountPlanCopy.planLine.seat(seat.guild_name)
          : accountPlanCopy.planLine.endsThenSeat(end, seat.guild_name, zone),
      line2: null,
      founder,
      manageBilling: true,
      ended: null,
    };
  }
  if (seat !== null) {
    return {
      line: accountPlanCopy.planLine.seat(seat.guild_name),
      line2: accountPlanCopy.planLine.givenBy(seat.given_by),
      founder: false,
      manageBilling: false,
      ended: null,
    };
  }
  return {
    line: ownPremiumLine(me, zone),
    line2: null,
    founder,
    manageBilling: billing !== null,
    ended: billing === null ? endedSeatLine(me, zone) : null,
  };
}

/** A guild row's line: the plan words under the tier's own name, plus the seat note when this account holds one. */
export function guildRowPlanLine(plan: GuildBillingView, zone?: string): string {
  const base = guildPlanLine(plan, zone);
  const named =
    plan.tier === 'raid_guild'
      ? base.replace(settingsWords.guildPlan, accountPlanCopy.guildRow.raidPlan)
      : base;
  return plan.you_hold_seat ? `${named}${accountPlanCopy.guildRow.youHoldSeat}` : named;
}
