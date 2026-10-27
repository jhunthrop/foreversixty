// web/src/lib/billing/seats-api.ts
// The Premium seat routes (launch pricing spec section 4). A refusal is a `BillingApiError`
// whose `code` is `seat_locked` (with `fields.movable_at`), `not_member`, `no_plan` or `conflict`.
import { invalidateGuild } from '../guild/api';
import { API_BASE_URL } from '../planner/config';
import { call, callVoid } from './api';
import type { SeatCandidate, SeatView } from './contract';

function seatsPath(guildId: number): string {
  return `/v1/guilds/${guildId}/seats`;
}

/** Verified members who hold no seat here; a payer's read. */
export async function seatCandidates(
  guildId: number,
  apiBase: string = API_BASE_URL,
): Promise<SeatCandidate[]> {
  const data = await call<{ candidates: SeatCandidate[] }>(`${seatsPath(guildId)}/candidates`, apiBase);
  return data.candidates;
}

async function seatAction(
  guildId: number,
  slot: number,
  action: 'give' | 'move' | 'withdraw',
  body: unknown,
  apiBase: string,
): Promise<SeatView> {
  const seat = await call<SeatView>(`${seatsPath(guildId)}/${slot}/${action}`, apiBase, {
    method: 'POST',
    body,
  });
  invalidateGuild(guildId, apiBase);
  return seat;
}

/** A payer recipient is offered the seat; anyone else holds it at once. */
export function giveSeat(
  guildId: number,
  slot: number,
  userId: number,
  apiBase: string = API_BASE_URL,
): Promise<SeatView> {
  return seatAction(guildId, slot, 'give', { user_id: userId }, apiBase);
}

/** Refused with 409 `seat_locked` while the 30-day lock runs. */
export function moveSeat(
  guildId: number,
  slot: number,
  userId: number,
  apiBase: string = API_BASE_URL,
): Promise<SeatView> {
  return seatAction(guildId, slot, 'move', { user_id: userId }, apiBase);
}

/** Only an offered seat can be withdrawn. */
export function withdrawSeat(
  guildId: number,
  slot: number,
  apiBase: string = API_BASE_URL,
): Promise<SeatView> {
  return seatAction(guildId, slot, 'withdraw', {}, apiBase);
}

async function offerAction(guildId: number, action: 'accept' | 'decline', apiBase: string): Promise<void> {
  await callVoid(`${seatsPath(guildId)}/offer/${action}`, apiBase, { method: 'POST', body: {} });
  invalidateGuild(guildId, apiBase);
}

export function acceptSeatOffer(guildId: number, apiBase: string = API_BASE_URL): Promise<void> {
  return offerAction(guildId, 'accept', apiBase);
}

export function declineSeatOffer(guildId: number, apiBase: string = API_BASE_URL): Promise<void> {
  return offerAction(guildId, 'decline', apiBase);
}
