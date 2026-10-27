// web/src/lib/billing/seats-copy.ts
// The Seats, Seat picker, Who pays and Receiver rows of design/specs/2026-10-27-launch-pricing.md
// section 6.2, with a function for every {slot}. Dates arrive as stored and are read in UTC;
// prices are built from ./plans.
import { utcDayMonth, utcDayMonthYear } from '../dates';
import { SEAT_MOVE_LOCK_DAYS, formatCents, type Interval } from './plans';

export const seatsCopy = {
  heading: 'Premium seats',
  given: (given: number, total: number): string => `${given} of ${total} given`,
  explainer: `A seat gives one verified member Premium while the plan lasts. A seat can move once every ${SEAT_MOVE_LOCK_DAYS} days.`,
  givenOn: (date: string): string => `Given ${utcDayMonth(date)}`,
  canMoveFrom: (date: string): string => `Can move from ${utcDayMonth(date)}`,
  canMoveNow: 'Can move now',
  move: 'Move seat',
  moving: 'Moving…',
  waiting: 'Waiting',
  offeredOn: (date: string): string =>
    `Offered ${utcDayMonth(date)}. They pay for Premium, so they accept it first.`,
  withdraw: 'Withdraw',
  withdrawing: 'Withdrawing…',
  accepted: (starts: string): string =>
    `Accepted · starts ${utcDayMonth(starts)}, when their own Premium ends`,
  notGiven: 'Not given',
  give: 'Give seat',
  giving: 'Giving…',
  freedSince: (date: string, name: string): string =>
    `Free since ${utcDayMonth(date)}: ${name} left the guild.`,
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
    lockedError: (date: string): string => `That seat can move from ${utcDayMonth(date)}.`,
    notMember: (name: string): string => `${name} is no longer a verified member.`,
  },
  whoPays: {
    heading: 'Who pays',
    help: 'Hand the plan to another officer, for example when the guild leader changes. The founder price goes with it.',
    transfer: 'Transfer the plan',
    transferTo: 'Transfer to',
    offerNote: (name: string, from: string, amountCents: number, interval: Interval): string =>
      `${name} gets an offer to take over. When they add a card, they are billed from ${utcDayMonthYear(from)} at ${formatCents(amountCents, interval)}, and your plan stops renewing. Nothing changes until they accept.`,
    send: 'Send the offer',
    sending: 'Sending…',
    cancel: 'Cancel',
    offered: (name: string, date: string): string =>
      `Offered to ${name} on ${utcDayMonth(date)}. Waiting for them.`,
    withdraw: 'Withdraw',
    declined: (name: string): string => `${name} declined.`,
    noCandidates: 'No other officer has a Forever Sixty account yet.',
    failed: 'The offer did not send. Try again.',
  },
  receiver: {
    offer: (name: string, from: string, amountCents: number, interval: Interval, founder: boolean): string =>
      `${name} offers you the guild plan. You would be billed from ${utcDayMonthYear(from)} at ${formatCents(amountCents, interval)}${founder ? ', the founder price' : ''}. Nothing is charged before then.`,
    take: 'Take over the plan',
    opening: 'Opening checkout…',
    notNow: 'Not now',
  },
} as const;
