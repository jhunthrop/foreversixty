// web/src/lib/billing/account-plan-copy.ts
// The account page Plan panel strings of design/specs/2026-10-27-launch-pricing.md section 6.3:
// Offer, Plan line, Founder, Seat ended and Guild row. Dates arrive as stored and are read in
// UTC; prices are built from ./plans.
import { utcDayMonth, utcDayMonthYear } from '../dates';
import { PLANS, formatCents, type Interval } from './plans';

export const accountPlanCopy = {
  offer: {
    title: (guild: string): string => `${guild} offers you a Premium seat.`,
    from: (name: string, date: string): string => `From ${name}, ${utcDayMonth(date)}.`,
    consequence: (ownEnds: string, amountCents: number, interval: Interval): string =>
      `If you accept, your own Premium stops renewing and ends ${utcDayMonthYear(ownEnds)}, and the seat starts then. You are not charged twice. If the seat is ever taken away, we email you a link to come back at ${formatCents(amountCents, interval)}, the price you pay now.`,
    accept: 'Accept the seat',
    accepting: 'Accepting…',
    decline: 'Decline',
    accepted: (ownEnds: string): string =>
      `Accepted. Your Premium ends ${utcDayMonthYear(ownEnds)}; the seat starts then.`,
    declined: 'Declined. Your Premium carries on as before.',
    failed: 'That did not save. Try again.',
  },
  planLine: {
    renews: (date: string, amountCents: number, interval: Interval): string =>
      `Premium · renews ${utcDayMonthYear(date)} · ${formatCents(amountCents, interval)}`,
    endsThenSeat: (date: string, guild: string): string =>
      `Premium · ends ${utcDayMonthYear(date)} · then a seat from ${guild}`,
    seat: (guild: string): string => `Premium · a seat from ${guild}`,
    givenBy: (name: string): string =>
      `Given by ${name}. It lasts while the guild's plan does; nothing is billed to you.`,
  },
  founder: {
    pill: 'Founder',
    title: 'Founder price, kept while your plan stays active.',
  },
  seatEnded: {
    text: (guild: string, date: string): string => `Your seat from ${guild} ended ${utcDayMonth(date)}. `,
    action: (amountCents: number, interval: Interval): string =>
      `Come back at ${formatCents(amountCents, interval)}`,
  },
  guildRow: {
    plan: `${PLANS.guild.name} plan`,
    raidPlan: `${PLANS.raid_guild.name} plan`,
    youHoldSeat: ' · you hold a seat',
  },
} as const;
