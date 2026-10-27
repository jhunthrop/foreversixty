// web/src/lib/billing/contract.ts
// The wire shapes of the guild billing read, the seat routes and the billing half of
// `GET /v1/me`. Mirrors the API's guild billing read (billing/guild_view.go) and the seats
// package (seats/handler.go, auth/me_billing.go) field for field; every price the API
// states arrives as `amount_cents` and is rendered through ./plans copy, never typed here.
import type { Interval } from './plans';

export type GuildTier = 'guild' | 'raid_guild';
export type SeatState = 'empty' | 'held' | 'offered' | 'accepted';
export type GuildBillingCan = 'billing' | 'seats' | 'transfer';
export type TransferState = 'offered' | 'declined' | 'accepted';

/** The price facts stamped on a subscription row; the founder flag is the Stripe price it was bought at. */
export interface BillingFacts {
  founder: boolean;
  interval: Interval | null;
  amount_cents: number | null;
}

/** `billing` in the guild billing read: null before the guild has a plan. */
export interface GuildPlanView extends BillingFacts {
  scope: 'plan' | 'officer_tools';
  tier: GuildTier;
  status: string;
  current_period_end: string | null;
  cancel_at_period_end: boolean;
  trial_end: string | null;
  /** Empty for a plan granted through the CLI: there is no billing contact to name. */
  billed_by: string;
  you_are_billing_contact: boolean;
}

export interface SeatHolder {
  user_id: number;
  character_key: string;
  name: string;
  class: string;
}

export interface SeatReturned {
  name: string;
  at: string;
}

/** A member who pays for their own Premium and has been offered this seat; they accept it first. */
export interface SeatOfferee {
  name: string;
  class: string;
  offered_at: string;
}

export interface SeatView {
  slot: number;
  state: SeatState;
  holder: SeatHolder | null;
  given_at: string | null;
  starts_at: string | null;
  movable_at: string | null;
  returned: SeatReturned | null;
  /** Waiting on a paying recipient; the holder, if any, keeps the seat until they accept. */
  offeree: SeatOfferee | null;
}

export interface TransferView {
  to_name: string;
  /** True when the reader is the receiver. */
  to_you: boolean;
  from_name: string;
  at: string;
  state: TransferState;
  starts_at: string | null;
  amount_cents: number | null;
  interval: Interval | null;
  founder: boolean;
}

export interface TransferCandidate {
  user_id: number;
  name: string;
}

/** A verified member who holds no seat here; `pays_premium` marks one who would be offered, not given, the seat. */
export interface SeatCandidate {
  user_id: number;
  name: string;
  class: string;
  pays_premium: boolean;
}

/** What `GET /v1/guilds/{id}/billing` answers. `can` holds `billing`, `seats` and `transfer` as the reader earns them. */
export interface GuildBillingSettings {
  billing: GuildPlanView | null;
  /** Top level because `billing` is null before a plan. */
  trial_used: boolean;
  /** The live count of distinct verified accounts, leader included. */
  verified_count: number;
  seats: SeatView[];
  transfer: TransferView | null;
  transfer_candidates: TransferCandidate[];
  can: GuildBillingCan[];
}

export interface UpgradePreview {
  amount_due_cents: number;
  next_amount_cents: number;
  interval: Interval;
  renews_at: string;
  /** Unix seconds the amounts were priced at; the upgrade must send it back so the charge matches. */
  proration_date: number;
}

/** `me.seat`: the Premium seat this account holds. */
export interface MeSeat {
  guild_name: string;
  given_by: string;
  since: string;
  starts_at: string | null;
}

/** `me.seat_offer`: a seat offered to an account that pays for its own Premium. */
export interface MeSeatOffer {
  guild_id: number;
  guild_name: string;
  from_name: string;
  offered_at: string;
  own_ends_at: string | null;
  own_amount_cents: number | null;
  own_interval: Interval | null;
}

/** `me.ended_seat`: shown for 30 days after a seat that replaced the account's own plan ends. */
export interface MeEndedSeat {
  guild_name: string;
  ended_at: string;
  resubscribe_url: string;
  amount_cents: number | null;
  interval: Interval | null;
}
