// web/src/lib/billing/seats-copy.ts
// The Seats, Seat picker, Who pays and Receiver rows of design/specs/2026-10-27-launch-pricing.md
// section 6.2, with a function for every {slot}. Dates arrive as stored: date-only values read as written, instants on the viewer calendar;
// prices are built from ./plans.
import { zonedDayMonth, zonedDayMonthYear } from '../dates';
import { SEAT_MOVE_LOCK_DAYS, formatCents, type Interval } from './plans';

export const seatsCopy = {
  heading: 'Premium seats',
  given: (given: number, total: number): string => `${given} of ${total} given`,
  explainer: `A seat gives one verified member Premium while the plan lasts. A seat can move once every ${SEAT_MOVE_LOCK_DAYS} days.`,
  givenOn: (date: string, zone?: string): string => `Given ${zonedDayMonth(date, zone)}`,
  canMoveFrom: (date: string, zone?: string): string => `Can move from ${zonedDayMonth(date, zone)}`,
  canMoveNow: 'Can move now',
  move: 'Move seat',
  moving: 'Moving…',
  waiting: 'Waiting',
  offeredOn: (date: string, zone?: string): string =>
    `Offered ${zonedDayMonth(date, zone)}. They pay for Premium, so they accept it first.`,
  withdraw: 'Withdraw',
  withdrawing: 'Withdrawing…',
  accepted: (starts: string, zone?: string): string =>
    `Accepted · starts ${zonedDayMonth(starts, zone)}, when their own Premium ends`,
  notGiven: 'Not given',
  give: 'Give seat',
  giving: 'Giving…',
  freedSince: (date: string, name: string, zone?: string): string =>
    `Free since ${zonedDayMonth(date, zone)}: ${name} left the guild.`,
  picker: {
    giveTitle: 'Give this seat to',
    moveTitle: 'Move this seat to',
    option: (name: string, className: string, paysPremium: boolean): string =>
      `${name} · ${className}${paysPremium ? ' · pays for Premium' : ''}`,
    paysNote:
      'They pay for Premium, so they get an offer to accept. Their own Premium stops renewing and the seat starts when it ends.',
    keepsNote: (name: string): string => `${name} keeps Premium until the seat moves.`,
    give: 'Give seat',
    move: 'Move seat',
    cancel: 'Cancel',
    everyoneHasPremium: 'Every verified member already has Premium from this guild.',
    saveFailed: 'That did not save. Try again.',
    // Additions to spec 6.2 (the API's 409 `seat_count` and `offer_pending`), pending an owner line.
    seatCount: 'This plan has no seat to give. Reload to see the current seats.',
    offerPending: 'That seat already has an offer waiting. Withdraw it first.',
    lockedError: (date: string, zone?: string): string =>
      `That seat can move from ${zonedDayMonth(date, zone)}.`,
    notMember: (name: string): string => `${name} is no longer a verified member.`,
  },
  whoPays: {
    heading: 'Who pays',
    help: 'Hand the plan to another officer, for example when the guild leader changes. The founder price goes with it.',
    transfer: 'Transfer the plan',
    transferTo: 'Transfer to',
    offerNote: (name: string, from: string, amountCents: number, interval: Interval, zone?: string): string =>
      `${name} gets an offer to take over. When they add a card, they are billed from ${zonedDayMonthYear(from, zone)} at ${formatCents(amountCents, interval)}, and your plan stops renewing. Nothing changes until they accept.`,
    send: 'Send the offer',
    sending: 'Sending…',
    cancel: 'Cancel',
    offered: (name: string, date: string, zone?: string): string =>
      `Offered to ${name} on ${zonedDayMonth(date, zone)}. Waiting for them.`,
    withdraw: 'Withdraw',
    declined: (name: string): string => `${name} declined.`,
    noCandidates: 'No other officer has a Forever Sixty account yet.',
    failed: 'The offer did not send. Try again.',
    // Additions: the design spec has no string for these two refusals.
    inProgress: 'An offer or a plan change is already in progress. Reload to see where it stands.',
    billingUnavailable: 'Billing is not available right now. Try again in a few minutes.',
  },
  receiver: {
    offer: (
      name: string,
      from: string,
      amountCents: number,
      interval: Interval,
      founder: boolean,
      zone?: string,
    ): string =>
      `${name} offers you the guild plan. You would be billed from ${zonedDayMonthYear(from, zone)} at ${formatCents(amountCents, interval)}${founder ? ', the founder price' : ''}. Nothing is charged before then.`,
    take: 'Take over the plan',
    opening: 'Opening checkout…',
    notNow: 'Not now',
  },
} as const;
