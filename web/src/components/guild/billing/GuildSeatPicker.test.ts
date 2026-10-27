import { render } from 'svelte/server';
import { describe, expect, it } from 'vitest';
import type { SeatCandidate } from '../../../lib/billing/contract';
import { SEAT_CANDIDATES } from '../../../test-support/billing-fixtures';
import { text } from '../../../test-support/raid-fixtures';
import GuildSeatPicker from './GuildSeatPicker.svelte';

interface Over {
  mode?: 'give' | 'move';
  candidates?: readonly SeatCandidate[] | null;
  error?: string;
}
const html = ({ mode = 'give', candidates = SEAT_CANDIDATES, error = '' }: Over = {}): string =>
  render(GuildSeatPicker, {
    props: {
      slot: 1,
      mode,
      holderName: 'Lightbrand',
      candidates,
      busy: false,
      error,
      onSubmit: () => {},
      onCancel: () => {},
    },
  }).body;

describe('GuildSeatPicker', () => {
  it('lists members by name with the payer marked', () => {
    const body = text(html());
    expect(body).toContain('Give this seat to');
    expect(body).toContain('Kraggor · Warrior');
    expect(body).toContain('Mirelle · Priest · pays for Premium');
    expect(body.indexOf('Kraggor')).toBeLessThan(body.indexOf('Mirelle'));
  });

  it('shows the payer note only when the chosen member pays', () => {
    expect(html()).not.toContain('pays-note');
    expect(html({ candidates: [SEAT_CANDIDATES[1]!] })).toContain('pays-note');
  });

  it('a move names who keeps Premium', () => {
    expect(text(html({ mode: 'move' }))).toContain('Lightbrand keeps Premium until the seat moves.');
  });

  it('an empty picker says every member already has Premium', () => {
    const body = html({ candidates: [] });
    expect(text(body)).toContain('Every verified member already has Premium from this guild.');
    expect(body).not.toContain('<select');
  });

  it('reserves the room with a Skeleton while loading', () => {
    expect(html({ candidates: null })).toContain('guild-seat-1-loading');
  });

  it('prints a refusal line as an alert', () => {
    expect(html({ error: 'That seat can move from 2 Dec.' })).toContain('role="alert"');
  });
});
