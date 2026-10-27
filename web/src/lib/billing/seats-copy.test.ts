import { describe, expect, it } from 'vitest';
import { specString } from '../../test-support/spec-table';
import { seatsCopy as copy } from './seats-copy';

const SPEC = new URL('../../../../design/specs/2026-10-27-launch-pricing.md', import.meta.url);
const spec = (label: string, index: number, fill: Record<string, string> = {}): string =>
  specString(SPEC, '### 6.2', label, index, fill);

describe('Seats row', () => {
  it('every string', () => {
    expect(copy.heading).toBe(spec('Seats', 0));
    expect(copy.given(1, 3)).toBe(spec('Seats', 1, { given: '1', total: '3' }));
    expect(copy.explainer).toBe(spec('Seats', 2));
    expect(copy.givenOn('2026-12-05')).toBe(spec('Seats', 3, { 'd MMM': '5 Dec' }));
    expect(copy.canMoveFrom('2026-12-20')).toBe(spec('Seats', 4, { 'd MMM': '20 Dec' }));
    expect(copy.canMoveNow).toBe(spec('Seats', 5));
    expect(copy.move).toBe(spec('Seats', 6));
    expect(copy.moving).toBe(spec('Seats', 7));
    expect(copy.waiting).toBe(spec('Seats', 8));
    expect(copy.offeredOn('2026-12-05')).toBe(spec('Seats', 9, { 'd MMM': '5 Dec' }));
    expect(copy.withdraw).toBe(spec('Seats', 10));
    expect(copy.withdrawing).toBe(spec('Seats', 11));
    expect(copy.accepted('2027-01-09')).toBe(spec('Seats', 12, { 'd MMM': '9 Jan' }));
    expect(copy.notGiven).toBe(spec('Seats', 13));
    expect(copy.give).toBe(spec('Seats', 14));
    expect(copy.giving).toBe(spec('Seats', 15));
    expect(copy.freedSince('2026-12-05', 'Aela')).toBe(spec('Seats', 16, { 'd MMM': '5 Dec', Name: 'Aela' }));
    expect(copy.movingTo('Aela', '2026-12-05')).toBe(spec('Seats', 17, { Name: 'Aela', 'd MMM': '5 Dec' }));
  });
});

describe('Seat picker row', () => {
  it('every string', () => {
    const p = copy.picker;
    const who = { Name: 'Aela', Class: 'Hunter' };
    expect(p.giveTitle).toBe(spec('Seat picker', 0));
    expect(p.moveTitle).toBe(spec('Seat picker', 1));
    expect(p.option('Aela', 'Hunter', false)).toBe(spec('Seat picker', 2, who));
    expect(p.option('Aela', 'Hunter', true)).toBe(spec('Seat picker', 3, who));
    expect(p.option('Mirelle', 'priest', true)).toBe('Mirelle · Priest · pays for Premium');
    expect(p.paysNote).toBe(spec('Seat picker', 4));
    expect(p.keepsNote('Aela')).toBe(spec('Seat picker', 5, who));
    expect(p.give).toBe(spec('Seat picker', 6));
    expect(p.move).toBe(spec('Seat picker', 7));
    expect(p.cancel).toBe(spec('Seat picker', 8));
    expect(p.everyoneHasPremium).toBe(spec('Seat picker', 9));
    expect(p.saveFailed).toBe(spec('Seat picker', 10));
    expect(p.lockedError('2026-12-20')).toBe(spec('Seat picker', 11, { 'd MMM': '20 Dec' }));
    expect(p.lockedError('2026-12-20')).toBe('That seat can move from 20 Dec.');
    expect(p.notMember('Aela')).toBe(spec('Seat picker', 13, who));
  });
});

describe('Who pays and Receiver rows', () => {
  it('who pays', () => {
    const w = copy.whoPays;
    expect(w.heading).toBe(spec('Who pays', 0));
    expect(w.help).toBe(spec('Who pays', 1));
    expect(w.transfer).toBe(spec('Who pays', 2));
    expect(w.transferTo).toBe(spec('Who pays', 3));
    expect(w.offerNote('Aela', '2027-01-10', 12000, 'yearly')).toBe(
      spec('Who pays', 4, { Name: 'Aela', date: '10 Jan 2027', price: '$120 a year' }),
    );
    expect(w.send).toBe(spec('Who pays', 5));
    expect(w.sending).toBe(spec('Who pays', 6));
    expect(w.cancel).toBe(spec('Who pays', 7));
    expect(w.offered('Aela', '2026-12-05')).toBe(spec('Who pays', 8, { Name: 'Aela', 'd MMM': '5 Dec' }));
    expect(w.withdraw).toBe(spec('Who pays', 9));
    expect(w.declined('Aela')).toBe(spec('Who pays', 10, { Name: 'Aela' }));
    expect(w.noCandidates).toBe(spec('Who pays', 11));
    expect(w.failed).toBe(spec('Who pays', 12));
    expect(w.staysWithGuild('Aela')).toBe(spec('Who pays', 13, { Name: 'Aela' }));
    expect(w.inProgressOffer('Aela')).toBe(spec('Who pays', 14, { Name: 'Aela' }));
    expect(w.inProgress).toBe(spec('Who pays', 15));
  });

  it('receiver, with and without the founder clause', () => {
    const r = copy.receiver;
    const fill = { Name: 'Aela', date: '10 Jan 2027', price: '$120 a year' };
    const clause = spec('Receiver', 1);
    expect(clause).toBe(', the founder price');
    expect(r.offer('Aela', '2027-01-10', 12000, 'yearly', true)).toBe(spec('Receiver', 0, fill));
    expect(r.offer('Aela', '2027-01-10', 12000, 'yearly', true)).toContain(`$120 a year${clause}.`);
    expect(r.offer('Aela', '2027-01-10', 15000, 'yearly', false)).toBe(
      spec('Receiver', 0, { ...fill, price: '$150 a year' }).replace(clause, ''),
    );
    expect(r.take).toBe(spec('Receiver', 2));
    expect(r.opening).toBe(spec('Receiver', 3));
    expect(r.notNow).toBe(spec('Receiver', 4));
    expect(r.staysWithGuild).toBe(spec('Receiver', 5));
  });
});

describe('instants read on the viewer calendar', () => {
  const unlocks = '2026-12-20T20:00:00Z';

  it('a seat that unlocks at 20:00 UTC is the 20th in Los Angeles and the 21st in Auckland', () => {
    expect(copy.canMoveFrom(unlocks, 'America/Los_Angeles')).toBe('Can move from 20 Dec');
    expect(copy.canMoveFrom(unlocks, 'Pacific/Auckland')).toBe('Can move from 21 Dec');
    expect(copy.picker.lockedError(unlocks, 'Pacific/Auckland')).toBe('That seat can move from 21 Dec.');
    const offer = copy.receiver.offer(
      'Aela',
      '2026-12-31T20:00:00Z',
      12000,
      'yearly',
      true,
      'Pacific/Auckland',
    );
    expect(offer).toContain('billed from 1 Jan 2027');
  });
});
