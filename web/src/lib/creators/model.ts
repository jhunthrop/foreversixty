// web/src/lib/creators/model.ts
// The creator pages' types (the API's shapes, api/openapi.yaml) and the little arithmetic they
// show: money in dollars, a month's label, the Stripe status Tag, and which action a statement
// offers. Every rule that decides money lives in the API (api/internal/creators); this module
// only formats what the API answered.
import type { TagTone } from '../../components/ui/Tag.svelte';
import { adminCreatorsCopy, creatorStatsCopy, payoutsCopy, stripeStatusCopy } from './copy';

export type StripeStatus = keyof typeof stripeStatusCopy;

export interface CreatorFigures {
  clicks: number;
  sign_ups: number;
  returning: number;
  revenue_cents: number;
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

export type StatementStatus = 'open' | 'ready' | 'nothing' | 'blocked' | 'paid';

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

/** One of the creator's players as the creator may see them: the week they signed up (a Monday), and
 *  whether they returned and pay, null while the player is under 7 days old (section 1, Privacy). */
export interface CreatorPlayer {
  week_of: string;
  returning: boolean | null;
  paying: boolean | null;
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
  month_cents: number;
  players: number;
  returning: number;
  paying: number;
  recent: CreatorPlayer[];
  payouts: CreatorPayout[];
  stripe_status: StripeStatus;
}

/** The returning-player bounty, in cents: the API's creators.BountyCents. */
export const BOUNTY_CENTS = 200;

const DOLLARS = new Intl.NumberFormat('en-US', { style: 'currency', currency: 'USD' });
const COUNT = new Intl.NumberFormat('en-US');

/** 1234 -> "$12.34"; a refund total stays positive and is labelled by its column. */
export function money(cents: number): string {
  return DOLLARS.format(cents / 100);
}

/** A cap as the owner types it: whole dollars. */
export function capDollars(cents: number): string {
  return DOLLARS.format(cents / 100).replace(/\.00$/, '');
}

export function count(value: number): string {
  return COUNT.format(value);
}

/** "2026-11" -> "November 2026". */
export function monthLabel(month: string): string {
  const [year, mon] = month.split('-').map(Number);
  return new Date(Date.UTC(year, mon - 1, 1)).toLocaleDateString('en-US', {
    month: 'long',
    year: 'numeric',
    timeZone: 'UTC',
  });
}

/** "2026-11-08" or an ISO instant -> "8 Nov 2026". */
export function dayLabel(value: string): string {
  const d = new Date(value.length === 10 ? `${value}T00:00:00Z` : value);
  return d.toLocaleDateString('en-GB', { day: 'numeric', month: 'short', year: 'numeric', timeZone: 'UTC' });
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

/** The one action a statement offers: pay it, settle a $0 month, or nothing. */
export type StatementAction = 'approve' | 'settle' | 'none';

export function statementAction(s: Pick<Statement, 'status' | 'total_cents'>): StatementAction {
  if (s.status !== 'ready' && s.status !== 'nothing') return 'none';
  return s.total_cents > 0 ? 'approve' : 'settle';
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

/** A creator Row's foot on /admin/creators: four totals since the start, then what is owed now. */
export function creatorFigures(c: Pick<Creator, 'figures' | 'cap_cents'>): ShownFigure[] {
  const k = adminCreatorsCopy.columns;
  const f = c.figures;
  const owedNote = f.capped ? adminCreatorsCopy.capReached : adminCreatorsCopy.ofCap(capDollars(c.cap_cents));
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
    shown('bounties', k.bounties, money(s.bounty_cents), payoutsCopy.bounties(s.bounty_count, money(BOUNTY_CENTS))),
    shown('share', k.share, money(s.share_cents)),
    shown('refunds', k.refunds, money(s.refund_cents)),
    shown('cap', k.cap, money(s.capped_cents)),
  ];
}

/** The creator's own four figures on /creators/stats. */
export function statsFigures(me: CreatorMe): ShownFigure[] {
  const k = creatorStatsCopy.figures;
  return [
    shown('month', k.month, money(me.month_cents), creatorStatsCopy.monthNote(capDollars(me.cap_cents)), true),
    shown('players', k.players, count(me.players)),
    shown('returning', k.returning, count(me.returning)),
    shown('paying', k.paying, count(me.paying)),
  ];
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
    text: payoutsCopy.line[l.kind](l.player, dayLabel(l.on)),
    amount: (l.kind === 'refund' ? '−' : '') + money(l.amount_cents),
    refund: l.kind === 'refund',
  };
}
