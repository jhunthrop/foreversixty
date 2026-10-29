// web/src/lib/billing/account-plan-copy.ts
// The account page Plan panel strings of design/specs/2026-10-27-launch-pricing.md section 6.3:
// Offer, Plan line, Founder, Seat ended and Guild row. Dates arrive as stored: date-only
// values read as written, instants on the viewer calendar; prices are built from ./plans.
import { zonedDayMonth, zonedDayMonthYear } from '../dates';
import { PLANS, formatCents, type Interval } from './plans';

export const accountPlanCopy = {
  offer: {
    title: (guild: string): string => `${guild} offers you a Premium seat.`,
    from: (name: string, date: string, zone?: string): string =>
      `From ${name}, ${zonedDayMonth(date, zone)}.`,
    consequence: (ownEnds: string, amountCents: number, interval: Interval, zone?: string): string =>
      `If you accept, your own Premium stops renewing and ends ${zonedDayMonthYear(ownEnds, zone)}, and the seat starts then. You are not charged twice. If the seat is ever taken away, we email you a link to come back at ${formatCents(amountCents, interval)}, the price you pay now.`,
    accept: 'Accept the seat',
    accepting: 'Accepting…',
    decline: 'Decline',
    accepted: (ownEnds: string, zone?: string): string =>
      `Accepted. Your Premium ends ${zonedDayMonthYear(ownEnds, zone)}; the seat starts then.`,
    acceptedNoDate: 'Accepted. The seat starts when your Premium ends.',
    declined: 'Declined. Your Premium carries on as before.',
    failed: 'That did not save. Try again.',
  },
  planLine: {
    renews: (date: string, amountCents: number, interval: Interval, zone?: string): string =>
      `Premium · renews ${zonedDayMonthYear(date, zone)} · ${formatCents(amountCents, interval)}`,
    endsThenSeat: (date: string, guild: string, zone?: string): string =>
      `Premium · ends ${zonedDayMonthYear(date, zone)} · then a seat from ${guild}`,
    ends: (date: string, zone?: string): string => `Premium · ends ${zonedDayMonthYear(date, zone)}`,
    renewsNoPrice: (date: string, zone?: string): string =>
      `Premium · renews ${zonedDayMonthYear(date, zone)}`,
    seat: (guild: string): string => `Premium · a seat from ${guild}`,
    givenBy: (name: string): string =>
      `Given by ${name}. It lasts while the guild's plan does; nothing is billed to you.`,
  },
  /** The plan card's head (polish spec 4.3): the plan name and the date it renews or ends. */
  head: {
    names: { premium: 'Premium', free: 'Free', seat: 'Guild seat' },
    renews: 'Renews',
    ends: 'Ends',
    until: 'Until',
    price: (amountCents: number, interval: Interval): string => formatCents(amountCents, interval),
    thenSeat: (guild: string): string => `Then a seat from ${guild}`,
    /** Under a renewing plan with its own subscription (player ruling, polish task 4). */
    cancelAnytime: 'Cancel anytime in Manage billing.',
    seatFrom: (guild: string): string => `A seat from ${guild}`,
  },
  founder: {
    pill: 'Founder',
    title: 'Founder price, kept while your plan stays active.',
  },
  seatEnded: {
    text: (guild: string, date: string, zone?: string): string =>
      `Your seat from ${guild} ended ${zonedDayMonth(date, zone)}. `,
    action: (amountCents: number, interval: Interval): string =>
      `Come back at ${formatCents(amountCents, interval)}`,
    actionNoPrice: 'Come back',
  },
  guildRow: {
    youHoldSeat: ' · you hold a seat',
    /** The plan row's value when the plan has no date to show. */
    active: 'Active',
    notVerified: 'Not verified',
    /** The plan tag in a guild block's name line. */
    planTag: { guild: PLANS.guild.name, raid_guild: PLANS.raid_guild.name, none: 'No plan' },
    noPlan: (guild: string): string => `${guild} has no plan.`,
    nextStep: 'See the next step',
    nextStepLabel: (action: string, guild: string): string => `${action} for ${guild}`,
  },
} as const;
