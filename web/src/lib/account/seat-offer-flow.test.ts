import { describe, expect, it, vi } from 'vitest';
import { ME_SEAT_OFFER } from '../../test-support/billing-fixtures';
import { answerOffer, isAnswered, isBusy, offerStatusLine } from './seat-offer-flow';

describe('answerOffer', () => {
  it('calls with the guild and reports the answer', async () => {
    const call = vi.fn().mockResolvedValue(undefined);
    expect(await answerOffer(call, 501, 'accepted')).toBe('accepted');
    expect(call).toHaveBeenCalledWith(501);
    expect(await answerOffer(call, 501, 'declined')).toBe('declined');
  });

  it('a refused call is failed, not thrown', async () => {
    const call = vi.fn().mockRejectedValue(new Error('boom'));
    expect(await answerOffer(call, 501, 'accepted')).toBe('failed');
  });
});

describe('offerStatusLine', () => {
  it('accepting reads when the own Premium ends and the seat starts', () => {
    expect(offerStatusLine('accepted', ME_SEAT_OFFER, 'UTC')).toBe(
      'Accepted. Your Premium ends 20 Dec 2026; the seat starts then.',
    );
  });

  it('declining reads its status', () => {
    expect(offerStatusLine('declined', ME_SEAT_OFFER)).toBe('Declined. Your Premium carries on as before.');
  });

  it('an accepted offer with no own end date still reads a status', () => {
    const offer = { ...ME_SEAT_OFFER, own_ends_at: null };
    expect(offerStatusLine('accepted', offer)).toBe('Accepted. The seat starts when your Premium ends.');
  });

  it('an open, busy or failed offer has no status line', () => {
    for (const outcome of ['idle', 'accepting', 'declining', 'failed'] as const) {
      expect(offerStatusLine(outcome, ME_SEAT_OFFER)).toBe('');
    }
  });
});

describe('offer outcome flags', () => {
  it('busy while an answer is in flight, so a second click cannot send it again', () => {
    expect(isBusy('accepting')).toBe(true);
    expect(isBusy('declining')).toBe(true);
    expect(isBusy('idle')).toBe(false);
    expect(isBusy('failed')).toBe(false);
  });

  it('answered for good only after a successful accept or decline', () => {
    expect(isAnswered('accepted')).toBe(true);
    expect(isAnswered('declined')).toBe(true);
    expect(isAnswered('failed')).toBe(false);
    expect(isAnswered('idle')).toBe(false);
  });
});
