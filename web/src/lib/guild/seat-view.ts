// web/src/lib/guild/seat-view.ts
// What the seats list decides from a seat read: which of the five §3.2.5 states a row is in, whether
// its 30-day lock still runs, and the error line a refused give or move reads as.
import { BillingApiError } from '../billing/api';
import type { SeatCandidate, SeatView } from '../billing/contract';
import { seatsCopy } from '../billing/seats-copy';

export type SeatRowKind = 'empty' | 'returned' | 'held' | 'offered' | 'accepted' | 'pending';

/** `pending`: a paying recipient has been offered the seat while the holder (if any) keeps it. */
export function seatRowKind(seat: SeatView): SeatRowKind {
  if (seat.offeree !== null) return 'pending';
  if (seat.state === 'empty') return seat.returned === null ? 'empty' : 'returned';
  return seat.state;
}

/** Seats that name someone: held, offered or accepted. */
export function givenCount(seats: readonly SeatView[]): number {
  return seats.filter((seat) => seat.state !== 'empty').length;
}

export function isSeatLocked(seat: SeatView, now: Date): boolean {
  return seat.movable_at !== null && new Date(seat.movable_at).getTime() > now.getTime();
}

export function sortedByName(candidates: readonly SeatCandidate[]): SeatCandidate[] {
  return [...candidates].sort((a, b) => a.name.localeCompare(b.name));
}

/** The refusals that mean the read this row drew from is stale: reload it. */
const STALE_READ_CODES: ReadonlySet<string> = new Set([
  'seat_locked',
  'not_member',
  'seat_count',
  'offer_pending',
]);

export function staleReadRefusal(error: unknown): boolean {
  return error instanceof BillingApiError && error.code !== null && STALE_READ_CODES.has(error.code);
}

/**
 * Each 409 reads as its line. `seat_locked` names the date from the refusal, else from the seat's own
 * `movable_at` (`knownMovableAt`); with neither it falls to the generic line.
 */
export function seatErrorLine(
  error: unknown,
  chosenName: string,
  knownMovableAt: string | null = null,
  zone?: string,
): string {
  if (!(error instanceof BillingApiError)) return seatsCopy.picker.saveFailed;
  switch (error.code) {
    case 'seat_locked': {
      const movableAt = error.fields?.movable_at ?? knownMovableAt;
      return movableAt === null ? seatsCopy.picker.saveFailed : seatsCopy.picker.lockedError(movableAt, zone);
    }
    case 'not_member':
      return seatsCopy.picker.notMember(chosenName);
    case 'seat_count':
      return seatsCopy.picker.seatCount;
    case 'offer_pending':
      return seatsCopy.picker.offerPending;
    default:
      return seatsCopy.picker.saveFailed;
  }
}
