// web/src/lib/billing/trial-copy.ts
// Every word of the free trials (design/specs/2026-11-08-premium-trial.md): the plan card's line,
// the checkout chooser and its creator code, the confirmation, the account page and guild settings
// line, the cancel confirm, and the reminder and receipt emails and notification. Every date is
// computed from the trial's real start (never a typed date) and read in UTC as the site's house
// date ("15 Nov 2026"); every price comes from plans.ts. The API's billingmail mirrors the email
// sentences and a test holds the two together.
import { priceFor } from './plan-costs';
import { FOUNDER_LAST_DAY, GUILD_TRIAL_DAYS, PLANS, formatPrice, type Interval, type PlanKey } from './plans';
import { formatDayMonthYear } from '../dates';

/** Premium's free trial without a code; a guild plan's is GUILD_TRIAL_DAYS. */
export const PREMIUM_TRIAL_DAYS = 7;
/** The free month a creator's code gives, on Premium or a guild plan. */
export const CODE_TRIAL_DAYS = 30;
/** The reminder goes out this many days before the first charge... */
export const REMINDER_LEAD_DAYS = 3;
/** ...or this many for a trial too short to leave three clear days after it starts. */
export const SHORT_REMINDER_LEAD_DAYS = 2;
/** The shortest trial that still gets the 3-day lead: day 3 of 6 is past the first half. */
const SHORT_TRIAL_DAYS = 5;

const DAY_MS = 24 * 60 * 60 * 1000;

/** The trial a checkout gives: 30 days with a creator code, else the plan's own. */
export function trialDays(plan: PlanKey, withCode: boolean): number {
  if (withCode) return CODE_TRIAL_DAYS;
  return plan === 'premium' ? PREMIUM_TRIAL_DAYS : GUILD_TRIAL_DAYS;
}

/** When a trial started at `start` ends and the first charge is taken. */
export function trialEndsAt(start: Date, days: number): Date {
  return new Date(start.getTime() + days * DAY_MS);
}

/** Days before the first charge that the reminder goes out (spec section 4.2). */
export function reminderLeadDays(days: number): number {
  return days <= SHORT_TRIAL_DAYS ? SHORT_REMINDER_LEAD_DAYS : REMINDER_LEAD_DAYS;
}

/** The house date of an instant: "15 Nov 2026", UTC. */
export function trialDate(at: Date): string {
  return formatDayMonthYear(at);
}

function otherInterval(interval: Interval): Interval {
  return interval === 'monthly' ? 'yearly' : 'monthly';
}

/** "$5 a month", the price the trial turns into. */
export function chargePrice(plan: PlanKey, interval: Interval, founder: boolean): string {
  return formatPrice(priceFor(plan, interval, founder), interval);
}

/** What a trial on a plan costs after it, and when: every surface's facts, from one start instant. */
export interface TrialTerms {
  plan: PlanKey;
  days: number;
  endsAt: Date;
  /** "15 Nov 2026": the day the first charge is taken; "before" it means before that day starts. */
  endDate: string;
  /** "$5 a month" */
  price: string;
  /** "$40 a year": the other interval, said in brackets. */
  otherPrice: string;
  founder: boolean;
}

export function trialTerms(
  plan: PlanKey,
  interval: Interval,
  founder: boolean,
  withCode: boolean,
  start: Date,
): TrialTerms {
  const days = trialDays(plan, withCode);
  const endsAt = trialEndsAt(start, days);
  return {
    plan,
    days,
    endsAt,
    endDate: trialDate(endsAt),
    price: chargePrice(plan, interval, founder),
    otherPrice: chargePrice(plan, otherInterval(interval), founder),
    founder,
  };
}

const founderNote = `Founder price for anyone who subscribes by ${FOUNDER_LAST_DAY}.`;

/** The one sentence a plan card, the chooser and the confirmation all carry (spec section 3.1). */
export function trialLine(t: TrialTerms): string {
  const scope = PLANS[t.plan].guild ? ' for the whole guild' : '';
  const line =
    `Free for ${t.days} days, then ${t.price} (or ${t.otherPrice})${scope}. ` +
    `Cancel any time before ${t.endDate} and you pay nothing.`;
  return t.founder ? `${line} ${founderNote}` : line;
}

/** "Trial ends 15 Nov 2026, then $5 a month": the account page and guild settings, always. */
export function trialStatusLine(endDate: string, price: string): string {
  return `Trial ends ${endDate}, then ${price}.`;
}

export const trialCopy = {
  /** The Premium plan card's tag beside the price. */
  cardTag: 'Free trial',
  chooser: {
    title: 'Start your free trial',
    intervalLabel: 'Billing',
    codeLabel: 'Creator code',
    codeHint: 'Optional. A creator’s code makes the trial a free month.',
    codePlaceholder: 'e.g. GRIMBREW',
    apply: 'Apply',
    checking: 'Checking',
    codeApplied: (code: string): string => `${code} applied: your first month is free.`,
    remove: 'Remove',
    codeUnknown: 'That isn’t a creator’s code. Check the spelling, or start without one.',
    trialUsed:
      'This account already had its free trial, so a code can’t give another. You can still subscribe.',
    guildTrialUsed:
      'This guild already had its free trial, so a code can’t give another. You can still subscribe.',
    cardNote: 'Stripe asks for a card now and shows the same dates before you confirm.',
    start: (days: number): string => `Start ${days} free days`,
    subscribe: (price: string): string => `Subscribe · ${price}`,
  },
  confirmation: {
    title: 'Your free Premium is on',
    guildTitle: (guild: string): string => `${guild}’s free trial is on`,
    cancel: 'Cancel the trial',
    keep: 'Nothing else to do. We remind you 3 days before the first charge.',
  },
  account: {
    cancel: 'Cancel trial',
    canceled: (endDate: string): string => `Trial ends ${endDate}. You won’t be charged.`,
    resume: 'Keep Premium after the trial',
    resumeGuild: 'Keep the plan after the trial',
    canceling: 'Cancelling',
    cancelFailed: 'That didn’t go through; nothing changed. Try again.',
  },
  confirm: {
    title: 'Cancel your free trial?',
    body: (endDate: string): string => `Premium stays on until ${endDate} and you won’t be charged.`,
    guildBody: (guild: string, endDate: string): string =>
      `${guild} keeps the plan until ${endDate} and you won’t be charged.`,
    yes: 'Cancel trial',
    no: 'Keep my trial',
  },
  /** The page the email's cancel link opens (a POST does the cancelling, never the link itself). */
  cancelPage: {
    title: 'Cancel your free trial',
    expired: 'This link has expired or was already used. You can cancel from your account page.',
    done: (endDate: string): string => `Done. Your trial ends ${endDate} and you won’t be charged.`,
    accountLink: 'Your account',
  },
} as const;

/** The reminder, by email and in the app (spec section 4). `link` is the one-click cancel page. */
export function reminderEmail(
  endDate: string,
  price: string,
  link: string,
): { subject: string; text: string } {
  return {
    subject: `Your free Premium ends on ${endDate}`,
    text:
      `Your free Premium ends on ${endDate}. On that day you'll be charged ${price}. Cancel in one click: ${link}\n\n` +
      `To keep Premium, do nothing.`,
  };
}

export function guildReminderEmail(
  guild: string,
  plan: string,
  endDate: string,
  price: string,
  link: string,
): { subject: string; text: string } {
  return {
    subject: `${guild}'s free ${plan} ends on ${endDate}`,
    text:
      `${guild}'s free ${plan} ends on ${endDate}. On that day you'll be charged ${price}. Cancel in one click: ${link}\n\n` +
      `To keep the plan, do nothing.`,
  };
}

/** The in-app notification: the same sentence, short. */
export function reminderNotification(endDate: string, price: string): { title: string; body: string } {
  return {
    title: `Free Premium ends ${endDate}`,
    body: `On that day you'll be charged ${price}. Tap to cancel.`,
  };
}

/** The receipt after the first charge. `receipt` is Stripe's hosted invoice page. */
export function receiptEmail(
  price: string,
  chargedOn: string,
  nextOn: string,
  receipt: string,
  account: string,
): { subject: string; text: string } {
  return {
    subject: `Receipt: Premium, ${price}`,
    text:
      `Your free trial ended and we charged ${price} on ${chargedOn}. Your receipt: ${receipt}\n\n` +
      `The next charge is on ${nextOn}. Change or cancel any time: ${account}`,
  };
}
