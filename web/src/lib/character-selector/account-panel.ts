// web/src/lib/character-selector/account-panel.ts
// What the account control's panel adds to the character list (spec 2026-10-28-polish 1.3): the
// head's identity and Premium tag, the Plan row's words and the Your profile link. Pure: the
// caller passes `/v1/me`, the current account character and the clock, and everything here is
// read from what the selector already holds (no request).
import type { Me, MeCharacter } from '../account/api';
import { accountPlanLine, premiumSource } from '../account/plan-line';
import { characterHref } from '../characters';
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
  href: string;
  founder: boolean;
}

export interface ProfileRow {
  href: string;
  name: string;
}

export interface AccountPanel {
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
export function accountPanelOf(me: Me, current: MeCharacter | null, nowMs: number): AccountPanel {
  const now = new Date(nowMs);
  const source = premiumSource(me, now);
  const line = accountPlanLine(me, now);
  const free = source === 'none';
  const { account } = selectorCopy;
  return {
    head: headOf(me, source !== 'none' && source !== 'fallback'),
    plan: {
      sub: free ? account.planSubFree : line.line,
      href: free ? account.freePlanHref : account.planHref,
      founder: line.founder,
    },
    profile:
      current === null
        ? null
        : { href: characterHref(current.region, current.ruleset, current.name), name: current.name },
  };
}
