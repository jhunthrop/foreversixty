import { describe, expect, it } from 'vitest';
import { BillingApiError } from '../billing/api';
import { emptySeat, heldSeat, offeredSeat, returnedSeat } from '../../test-support/billing-fixtures';
import { givenCount, isSeatLocked, seatErrorLine, seatRowKind, sortedByName } from './seat-view';

const NOW = new Date('2026-11-20T12:00:00Z');

describe('seat-view', () => {
  it('names the five row states', () => {
    expect(seatRowKind(emptySeat(1))).toBe('empty');
    expect(seatRowKind(returnedSeat())).toBe('returned');
    expect(seatRowKind(heldSeat())).toBe('held');
    expect(seatRowKind(offeredSeat())).toBe('offered');
    expect(seatRowKind(heldSeat(1, { state: 'accepted' }))).toBe('accepted');
  });

  it('counts the seats that name someone', () => {
    expect(givenCount([heldSeat(1), emptySeat(2), offeredSeat(3)])).toBe(2);
  });

  it('locks while movable_at is ahead and frees on or after it', () => {
    expect(isSeatLocked(heldSeat(), NOW)).toBe(true);
    expect(isSeatLocked(heldSeat(1, { movable_at: '2026-11-20T12:00:00Z' }), NOW)).toBe(false);
    expect(isSeatLocked(heldSeat(1, { movable_at: null }), NOW)).toBe(false);
  });

  it('sorts candidates by name without touching the input', () => {
    const input = [
      { user_id: 2, name: 'Mirelle', class: 'priest', pays_premium: true },
      { user_id: 1, name: 'Kraggor', class: 'warrior', pays_premium: false },
    ];
    expect(sortedByName(input).map((c) => c.name)).toEqual(['Kraggor', 'Mirelle']);
    expect(input[0]?.name).toBe('Mirelle');
  });

  it('reads each refusal as its line', () => {
    const locked = new BillingApiError('x', 409, 'seat_locked', { movable_at: '2026-12-02T12:00:00Z' });
    expect(seatErrorLine(locked, 'Kraggor', 'UTC')).toBe('That seat can move from 2 Dec.');
    const gone = new BillingApiError('x', 409, 'not_member');
    expect(seatErrorLine(gone, 'Kraggor')).toBe('Kraggor is no longer a verified member.');
    expect(seatErrorLine(new BillingApiError('x', 500), 'Kraggor')).toBe('That did not save. Try again.');
    expect(seatErrorLine(new Error('x'), 'Kraggor')).toBe('That did not save. Try again.');
  });
});
