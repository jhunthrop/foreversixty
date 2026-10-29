// web/src/lib/character-selector/account-panel.ts
// What the account control's panel adds to the character list (spec 2026-10-28-polish 1.3): the
// head's identity and Premium tag, the Plan row's words and the Your profile link. Pure: the
// caller passes `/v1/me`, the current account character and the clock, and everything here is
// read from what the selector already holds (no request).
import type { Me, MeCharacter } from '../account/api';
import { accountPlanLine, premiumSource, type AccountPlanLine } from '../account/plan-line';
import type { PremiumSource } from '../billing/contract';
import { CENTS_PER_DOLLAR, formatAmount } from '../billing/plans';
import { characterHref } from '../characters';
import { zonedDayMonthYear } from '../dates';
import { selectorCopy } from './copy';

export interface AccountHead {
  /** The BattleTag, or the whole email address of an email account. */
  name: string;
  via: 'bnet' | 'email';
  /** A live plan, a grant or a held seat; never the fallback entitlement. */
  premium: boolean;
}

export interface PlanRow {
  sub: string;
  /** The second line (`Renews 26 Oct 2027`), only for the account's own plan. */
  sub2: string | null;
  href: string;
  founder: boolean;
  /** The gold `See plans` word before the chevron: a free account only (`source === 'none'`). */
  seePlans: boolean;
}

export interface ProfileRow {
  href: string;
}

export interface AccountPanel {
  /** Who pays for the account's Premium; the header's Premium tag follows it. */
  source: PremiumSource;
  head: AccountHead;
  plan: PlanRow;
  /** Null when the current character is not saved to the account. */
  profile: ProfileRow | null;
}

function headOf(me: Me, premium: boolean): AccountHead {
  const battletag = me.user.battletag?.trim() ?? '';
  if (battletag !== '') return { name: battletag, via: 'bnet', premium };
  const email = me.user.email?.trim() ?? '';
  return { name: email === '' ? 'Signed in' : email, via: 'email', premium };
}

/** `nowMs` is the clock in milliseconds, a number so reactive state can hold it. */
function ownPlanWords(me: Me): { sub: string; sub2: string } | null {
  const billing = me.entitlements?.billing ?? null;
  const end = billing?.current_period_end ?? null;
  if (billing === null || end === null || billing.amount_cents === null || billing.interval === null) {
    return null;
  }
  const { account } = selectorCopy;
  const unit = billing.interval === 'monthly' ? account.perMonth : account.perYear;
  const price = `${formatAmount(billing.amount_cents / CENTS_PER_DOLLAR)}${unit}`;
  const lead = billing.cancel_at_period_end ? account.ends : account.renews;
  return { sub: account.planOwn(price), sub2: `${lead} ${zonedDayMonthYear(end)}` };
}

function planOf(me: Me, source: PremiumSource, line: AccountPlanLine): PlanRow {
  const { account } = selectorCopy;
  if (source === 'none')
    return {
      sub: account.planSubFree,
      sub2: null,
      href: account.freePlanHref,
      founder: false,
      seePlans: true,
    };
  const own = source === 'own' ? ownPlanWords(me) : null;
  return {
    sub: own?.sub ?? line.line,
    sub2: own?.sub2 ?? null,
    href: account.planHref,
    founder: line.founder,
    seePlans: false,
  };
}

export function accountPanelOf(me: Me, current: MeCharacter | null, nowMs: number): AccountPanel {
  const now = new Date(nowMs);
  const source = premiumSource(me, now);
  const line = accountPlanLine(me, now);
  return {
    source,
    head: headOf(me, source !== 'none' && source !== 'fallback'),
    plan: planOf(me, source, line),
    profile: current === null ? null : { href: characterHref(current.region, current.ruleset, current.name) },
  };
}
