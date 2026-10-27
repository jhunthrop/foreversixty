import { render } from 'svelte/server';
import { describe, expect, it } from 'vitest';
import type { SeatView } from '../../../lib/billing/contract';
import { emptySeat, heldSeat } from '../../../test-support/billing-fixtures';
import { text } from '../../../test-support/raid-fixtures';
import GuildSeats from './GuildSeats.svelte';

const NOW = new Date('2026-11-20T12:00:00Z');
const html = (seats: SeatView[], manage = true): string =>
  render(GuildSeats, { props: { guildId: 7, seats, manage, now: NOW, onChanged: () => Promise.resolve() } })
    .body;

describe('GuildSeats', () => {
  it('heads the list with its count and the helper', () => {
    const body = text(html([heldSeat(1), emptySeat(2), emptySeat(3)]));
    expect(body).toContain('Premium seats');
    expect(body).toContain('1 of 3 given');
    expect(body).toContain('A seat can move once every 30 days.');
  });

  it('counts out of the plan seat count, never the rows sent', () => {
    const rows = [
      heldSeat(1),
      heldSeat(2),
      emptySeat(3),
      emptySeat(4, { returned: { name: 'A', at: '2026-11-19' } }),
    ];
    const body = render(GuildSeats, {
      props: {
        guildId: 7,
        seats: rows,
        total: 3,
        manage: true,
        now: NOW,
        onChanged: () => Promise.resolve(),
      },
    }).body;
    expect(text(body)).toContain('2 of 3 given');
  });

  it('draws nothing when the plan has no seats', () => {
    expect(html([])).not.toContain('guild-seats');
  });
});
