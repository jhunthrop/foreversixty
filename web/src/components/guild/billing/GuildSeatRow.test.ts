import { render } from 'svelte/server';
import { describe, expect, it } from 'vitest';
import type { SeatView } from '../../../lib/billing/contract';
import { emptySeat, heldSeat, offeredSeat, pendingSeat } from '../../../test-support/billing-fixtures';
import { text } from '../../../test-support/raid-fixtures';
import GuildSeatRow from './GuildSeatRow.svelte';

const NOON = { given_at: '2026-11-02T12:00:00Z', movable_at: '2026-12-02T12:00:00Z' };
const held = (): SeatView => heldSeat(1, NOON);
const NOW = new Date('2026-11-20T12:00:00Z');
const html = (seat: SeatView, manage = true, now = NOW): string =>
  render(GuildSeatRow, { props: { guildId: 7, seat, manage, now, onChanged: () => Promise.resolve() } }).body;

describe('GuildSeatRow', () => {
  it('Held and locked: Move seat is aria-disabled and its date is visible text', () => {
    const body = html(held());
    expect(body).toContain('aria-disabled="true"');
    expect(body).toContain('aria-describedby="guild-seat-1-note"');
    expect(text(body)).toContain('Can move from 2 Dec');
    expect(text(body)).toContain('Given 2 Nov');
  });

  it('Held past its lock reads Can move now with a live button', () => {
    const body = html(held(), true, new Date('2026-12-05T00:00:00Z'));
    expect(text(body)).toContain('Can move now');
    expect(body).not.toContain('aria-disabled="true"');
  });

  it('Offered: Waiting pill, the payer line and Withdraw', () => {
    const body = html(offeredSeat());
    expect(text(body)).toContain('Waiting');
    expect(text(body)).toContain('They pay for Premium, so they accept it first.');
    expect(body).toContain('guild-seat-1-withdraw');
  });

  it('Pending: the holder, a Waiting pill, the offer wording, and Withdraw beside Move seat', () => {
    const body = html(pendingSeat(1, NOON));
    expect(text(body)).toContain('Kraggor');
    expect(text(body)).toContain('Mirelle');
    expect(text(body)).toContain('Waiting');
    expect(text(body)).toContain('Offered 10 Nov. They pay for Premium, so they accept it first.');
    expect(body).toContain('guild-seat-1-move');
    expect(body).toContain('aria-disabled="true"');
    expect(body).toContain('guild-seat-1-withdraw');
    expect(body).not.toContain('guild-seat-1-give');
  });

  it('Pending on an empty seat: the offeree stands in, with Withdraw only', () => {
    const body = html(
      emptySeat(1, { offeree: { name: 'Mirelle', class: 'priest', offered_at: '2026-11-10T12:00:00Z' } }),
    );
    expect(text(body)).toContain('Mirelle');
    expect(body).toContain('guild-seat-1-withdraw');
    expect(body).not.toContain('guild-seat-1-move');
  });

  it('Accepted: the start line, then the lock date', () => {
    const body = html(heldSeat(1, { ...NOON, state: 'accepted', starts_at: '2026-12-22T12:00:00Z' }));
    expect(text(body)).toContain('Accepted · starts 22 Dec, when their own Premium ends');
    expect(text(body)).toContain('Can move from 2 Dec');
  });

  it('Empty: Not given and Give seat', () => {
    const body = html(emptySeat(2));
    expect(text(body)).toContain('Not given');
    expect(body).toContain('guild-seat-2-give');
  });

  it('Returned: the free-since line', () => {
    expect(text(html(emptySeat(1, { returned: { name: 'Kraggor', at: '2026-11-20T12:00:00Z' } })))).toContain(
      'Free since 20 Nov: Kraggor left the guild.',
    );
  });

  it('draws no right-hand column for a reader without seats', () => {
    const body = html(held(), false);
    expect(body).not.toContain('guild-seat-1-move');
    expect(text(body)).toContain('Can move from 2 Dec');
  });
});
