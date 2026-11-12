// web/src/lib/creators/model.ts
// The creator pages' types (the API's shapes, api/openapi.yaml) and the little arithmetic they
// show: what one player is worth to a creator in a first year (the public page's worked examples,
// from PROGRAM_TERMS and the plan table), money and counts through the shared formatters, a month's
// label, the Stripe status Tag, and which action a statement offers. Every rule that decides money
// lives in the API (api/internal/creators); this module only formats what the API answered.
import type { TagTone } from '../../components/ui/Tag.svelte';
import {
  CENTS_PER_DOLLAR,
  FOUNDER_LAST_DAY,
  PLANS,
  entryPrice,
  formatPrice,
  formatUsd,
  type Interval,
  type PlanKey,
} from '../billing/plans';
import { DAY_MS, MONTHS, utcDayMonth, utcDayMonthYear } from '../dates';
import { grouped, usdCents } from '../format/number';
import {
  PROGRAM_TERMS,
  adminCreatorsCopy,
  applyCopy,
  creatorStatsCopy,
  payoutsCopy,
  requestsCopy,
  stripeStatusCopy,
  type Platform,
} from './copy';

const MONTHS_PER_YEAR = 12;
const PERCENT = 100;

/** The plans a worked example shows: a player's Premium both ways, and a guild's officer on each guild plan. */
export const EXAMPLE_PLANS: readonly { plan: PlanKey; interval: Interval }[] = [
  { plan: 'premium', interval: 'yearly' },
  { plan: 'premium', interval: 'monthly' },
  { plan: 'guild', interval: 'yearly' },
  { plan: 'raid_guild', interval: 'yearly' },
];

/** Dollars a creator earns in a player's first year on `plan`, paid per `interval`. */
export function firstYearEarnings(plan: PlanKey, interval: Interval, founder: boolean): number {
  const price = entryPrice(PLANS[plan], interval, founder);
  const paymentsInShare =
    interval === 'monthly' ? PROGRAM_TERMS.shareMonths : PROGRAM_TERMS.shareMonths / MONTHS_PER_YEAR;
  return PROGRAM_TERMS.bountyUsd + (price * paymentsInShare * PROGRAM_TERMS.sharePercent) / PERCENT;
}

export interface EarningsRow {
  readonly id: string;
  /** `Premium, $40 a year`. */
  readonly label: string;
  /** `$22`. */
  readonly value: string;
  /** While founder prices last, the same player at list price: `$27 at list price after 6 Jan 2027`. */
  readonly sub?: string;
}

/** The worked examples at today's prices (founder prices while the window is open). */
export function earningsRows(founder: boolean): readonly EarningsRow[] {
  return EXAMPLE_PLANS.map(({ plan, interval }) => ({
    id: `${plan}-${interval}`,
    label: `${PLANS[plan].name}, ${formatPrice(entryPrice(PLANS[plan], interval, founder), interval)}`,
    value: formatUsd(firstYearEarnings(plan, interval, founder)),
    ...(founder
      ? {
          sub: `${formatUsd(firstYearEarnings(plan, interval, false))} at list price after ${FOUNDER_LAST_DAY}`,
        }
      : {}),
  }));
}

export type StripeStatus = keyof typeof stripeStatusCopy;

export interface CreatorFigures {
  clicks: number;
  sign_ups: number;
  returning: number;
  revenue_cents: number;
  /** The open month owed_cents covers, "2026-12"; the four counts above are all time. */
  month: string;
  owed_cents: number;
  capped: boolean;
}

export interface Creator {
  id: number;
  name: string;
  channel_url: string;
  slug: string;
  code: string;
  starts_on: string;
  cap_cents: number;
  stripe_status: StripeStatus;
  figures: CreatorFigures;
}

export interface CreatorInput {
  name: string;
  channel_url: string;
  slug: string;
  code: string;
  starts_on: string;
  cap_cents: number;
}

export type StatementStatus = 'open' | 'ready' | 'nothing' | 'carry' | 'blocked' | 'paid';

export interface StatementLine {
  kind: 'bounty' | 'share' | 'refund';
  player: number;
  on: string;
  amount_cents: number;
}

export interface Statement {
  creator_id: number;
  name: string;
  month: string;
  status: StatementStatus;
  bounty_count: number;
  bounty_cents: number;
  share_cents: number;
  refund_cents: number;
  capped_cents: number;
  total_cents: number;
  paid_at: string | null;
  transfer_id: string | null;
  lines: StatementLine[];
}

/** The month a statements read covers, and whether the server lets a payout be made at all. */
export interface StatementsPage {
  month: string;
  payouts_enabled: boolean;
  /** The server's Stripe key cannot reach Connect (its startup check). */
  connect_denied: boolean;
  statements: Statement[];
}

/** One week (a Monday, UTC) of a creator's sign-ups as the creator may see it: returning and paying
 *  are null while the week is still running (section 1, Privacy). */
export interface CreatorWeek {
  week_of: string;
  signed_up: number;
  returning: number | null;
  paying: number | null;
}

export interface CreatorPayout {
  month: string;
  amount_cents: number;
  paid_at: string;
}

export interface CreatorMe {
  name: string;
  slug: string;
  code: string;
  cap_cents: number;
  /** The open month the estimate covers, "2026-12". */
  month: string;
  month_cents: number;
  players: number;
  returning: number;
  paying: number;
  weeks: CreatorWeek[];
  payouts: CreatorPayout[];
  stripe_status: StripeStatus;
}

/** The returning-player bounty, in cents: the API's creators.BountyCents. */
export const BOUNTY_CENTS = 200;

/** 1234 -> "$12.34", always to the cent so a column of amounts lines up; a refund total stays
 *  positive and is labelled by its column. */
export function money(cents: number): string {
  return usdCents(cents);
}

/** A cap as the owner types it: whole dollars (`$500`), cents only when it has them. */
export function capDollars(cents: number): string {
  return formatUsd(cents / CENTS_PER_DOLLAR);
}

export function count(value: number): string {
  return grouped(value);
}

/** "2026-11" -> "Nov 2026": months are three letters everywhere (owner ruling 2026-11-10). */
export function monthLabel(month: string): string {
  const [year, mon] = month.split('-').map(Number);
  return `${MONTHS[(mon ?? 1) - 1] ?? ''} ${year ?? ''}`;
}

/** The month before "2026-12": "2026-11". */
export function previousMonth(month: string): string {
  const [year, mon] = month.split('-').map(Number);
  const d = new Date(Date.UTC(year ?? 0, (mon ?? 1) - 2, 1));
  return `${d.getUTCFullYear()}-${String(d.getUTCMonth() + 1).padStart(2, '0')}`;
}

/** The day a closed month is approved by: the 7th of the open month, "7 Dec". */
export function approvedByDay(openMonth: string): string {
  return utcDayMonth(`${openMonth}-07`);
}

/** "2026-11-08" or an ISO instant -> "8 Nov 2026", read in UTC as the API wrote it. */
export function dayLabel(value: string): string {
  return utcDayMonthYear(value);
}

/** The month the clock is in, "2026-11", and the twelve closed months before it, newest first. */
export function recentMonths(now: Date): string[] {
  return Array.from({ length: 13 }, (_, i) => {
    const d = new Date(Date.UTC(now.getUTCFullYear(), now.getUTCMonth() - i, 1));
    return `${d.getUTCFullYear()}-${String(d.getUTCMonth() + 1).padStart(2, '0')}`;
  });
}

export const STATUS_TONE: Readonly<Record<StripeStatus, TagTone>> = {
  ready: 'success',
  onboarding: 'muted',
  needs_info: 'warn',
  none: 'muted',
};

/** A statement offers one action, Approve, and only when it is ready and owes money. A month that
 *  owes nothing has no button: its lines (a refund that outweighed the earnings) roll into the next. */
export function canApprove(s: Pick<Statement, 'status' | 'total_cents'>): boolean {
  return s.status === 'ready' && s.total_cents > 0;
}

/** The month's total still to pay across creators (the ready ones). */
export function toPay(statements: readonly Statement[]): number {
  return statements.reduce((sum, s) => sum + (s.status === 'ready' ? s.total_cents : 0), 0);
}

/** One figure a row shows: the label, the value, the line under it, and whether it is the row's answer. */
export interface ShownFigure {
  key: string;
  label: string;
  value: string;
  note: string;
  /** The one emphasised figure per row (gold): what the owner or the creator came for. */
  lead: boolean;
}

function shown(key: string, label: string, value: string, note = '', lead = false): ShownFigure {
  return { key, label, value, note, lead };
}

/** A creator Row's foot on /admin/creators: four all-time totals, then what the open month owes so far. */
export function creatorFigures(c: Pick<Creator, 'figures' | 'cap_cents'>): ShownFigure[] {
  const k = adminCreatorsCopy.columns;
  const f = c.figures;
  const owedNote = adminCreatorsCopy.owedNote(monthLabel(f.month), capDollars(c.cap_cents), f.capped);
  return [
    shown('clicks', k.clicks, count(f.clicks)),
    shown('sign-ups', k.signUps, count(f.sign_ups)),
    shown('returning', k.returning, count(f.returning)),
    shown('revenue', k.revenue, money(f.revenue_cents)),
    shown('owed', k.owed, money(f.owed_cents), owedNote, true),
  ];
}

/** A statement Row's foot: the total first (what Approve pays), then how it was made. */
export function statementFigures(s: Statement): ShownFigure[] {
  const k = payoutsCopy.columns;
  return [
    shown('total', k.total, money(s.total_cents), '', true),
    shown(
      'bounties',
      k.bounties,
      money(s.bounty_cents),
      payoutsCopy.bounties(s.bounty_count, money(BOUNTY_CENTS)),
    ),
    shown('share', k.share, money(s.share_cents)),
    shown('refunds', k.refunds, money(s.refund_cents)),
    shown('cap', k.cap, money(s.capped_cents)),
  ];
}

/** A share as a whole percent, "62%"; nothing of nothing is 0%. */
function share(part: number, whole: number): string {
  return `${whole > 0 ? Math.round((part / whole) * 100) : 0}%`;
}

/** The creator's own four figures on /creators/stats: the open month's estimate, then the all-time
 *  counts with each one's rate of the one before it. */
export function statsFigures(me: CreatorMe): ShownFigure[] {
  const c = creatorStatsCopy;
  const k = c.figures;
  return [
    shown(
      'month',
      k.month,
      money(me.month_cents),
      c.monthNote(monthLabel(me.month), capDollars(me.cap_cents)),
      true,
    ),
    shown('players', k.players, count(me.players)),
    shown('returning', k.returning, count(me.returning), c.returningNote(share(me.returning, me.players))),
    shown('paying', k.paying, count(me.paying), c.payingNote(share(me.paying, me.returning))),
  ];
}

/** One week of sign-ups as a row: the title, its three counts (the running week shows only its
 *  sign-ups, spec section 1, Privacy) and whether it is still filling (under 14 days old). */
export interface WeekRow {
  key: string;
  title: string;
  filling: boolean;
  figures: ShownFigure[];
}

const FILL_DAYS = 14;

export function weekRow(w: CreatorWeek, now: Date): WeekRow {
  const k = creatorStatsCopy.columns;
  const start = new Date(`${w.week_of}T00:00:00Z`).getTime();
  const running = w.returning === null || w.paying === null;
  return {
    key: w.week_of,
    title: creatorStatsCopy.weekTitle(utcDayMonth(w.week_of)),
    filling: running || now.getTime() - start < FILL_DAYS * DAY_MS,
    figures: [
      shown('signed-up', k.signedUp, count(w.signed_up)),
      ...(running
        ? []
        : [
            shown('returning', k.returning, count(w.returning ?? 0)),
            shown('paying', k.paying, count(w.paying ?? 0)),
          ]),
    ],
  };
}

/** A statement Row's sub line. */
export function statementStatusLine(s: Statement): string {
  if (s.status !== 'paid') return payoutsCopy.status[s.status];
  const on = dayLabel(s.paid_at ?? '');
  return s.transfer_id ? payoutsCopy.status.paid(on, s.transfer_id) : payoutsCopy.status.settled(on);
}

/** A statement line as the Lines rail shows it, a refund signed. */
export function lineText(l: StatementLine): { text: string; amount: string; refund: boolean } {
  return {
    text: payoutsCopy.line[l.kind](l.player, utcDayMonth(l.on)),
    amount: (l.kind === 'refund' ? '−' : '') + money(l.amount_cents),
    refund: l.kind === 'refund',
  };
}

export type ApplicationStatus = 'pending' | 'approved' | 'declined';

/** The applicant's own application (GET /v1/creators/application). */
export interface MyApplication {
  status: ApplicationStatus;
  created_at: string;
}

/** A request as the owner sees it. The contact email is the owner's alone. */
export interface CreatorRequest {
  id: number;
  channel_url: string;
  platform: Platform;
  audience: number | null;
  about: string;
  contact_email: string | null;
  created_at: string;
  /** Suggestions for the Approve form: the channel's own name, a free slug and code. */
  suggested: { name: string; slug: string; code: string };
}

export interface ApplicationInput {
  channel_url: string;
  platform: Platform;
  audience: number | null;
  about: string;
  contact_email: string;
}

/** The API's own limit on "What do you make about Forever?". */
export const ABOUT_MAX = 500;

/** A request Row's sub line: platform, the size they gave, the day it came. */
export function requestSub(r: CreatorRequest): string {
  const size = r.audience === null ? requestsCopy.sizeUnknown : requestsCopy.size(count(r.audience));
  return requestsCopy.sub(applyCopy.platforms[r.platform], size, dayLabel(r.created_at));
}

/** The channel as a short title: the address without its scheme. */
export function channelTitle(url: string): string {
  return url.replace(/^https:\/\/(www\.)?/, '').replace(/\/$/, '');
}
