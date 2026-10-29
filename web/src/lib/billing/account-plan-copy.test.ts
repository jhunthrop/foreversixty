import { describe, expect, it } from 'vitest';
import { specString } from '../../test-support/spec-table';
import { accountPlanCopy as copy } from './account-plan-copy';

const SPEC = new URL('../../../../design/specs/2026-10-27-launch-pricing.md', import.meta.url);
const spec = (label: string, index: number, fill: Record<string, string> = {}): string =>
  specString(SPEC, '### 6.3', label, index, fill);

describe('account page rows', () => {
  it('Offer', () => {
    const o = copy.offer;
    expect(o.title('Olympus')).toBe(spec('Offer', 0, { Guild: 'Olympus' }));
    expect(o.from('Aela', '2026-12-05')).toBe(spec('Offer', 1, { Name: 'Aela', 'd MMM': '5 Dec' }));
    expect(o.consequence('2027-02-01', 4000, 'yearly')).toBe(
      spec('Offer', 2, { date: '1 Feb 2027', price: '$40 a year' }),
    );
    expect(o.accept).toBe(spec('Offer', 3));
    expect(o.accepting).toBe(spec('Offer', 4));
    expect(o.decline).toBe(spec('Offer', 5));
    expect(o.accepted('2027-02-01')).toBe(spec('Offer', 6, { date: '1 Feb 2027' }));
    expect(o.declined).toBe(spec('Offer', 7));
    expect(o.failed).toBe(spec('Offer', 8));
  });

  it('Plan line', () => {
    const p = copy.planLine;
    expect(p.renews('2027-02-01', 500, 'monthly')).toBe(
      spec('Plan line', 0, { date: '1 Feb 2027', price: '$5 a month' }),
    );
    expect(p.endsThenSeat('2027-02-01', 'Olympus')).toBe(
      spec('Plan line', 1, { date: '1 Feb 2027', Guild: 'Olympus' }),
    );
    expect(p.seat('Olympus')).toBe(spec('Plan line', 2, { Guild: 'Olympus' }));
    expect(p.givenBy('Aela')).toBe(spec('Plan line', 3, { Name: 'Aela' }));
  });

  it('Founder, Seat ended and Guild row', () => {
    expect(copy.founder.pill).toBe(spec('Founder', 0));
    expect(copy.founder.title).toBe(spec('Founder', 1));
    expect(copy.seatEnded.text('Olympus', '2026-12-05')).toBe(
      spec('Seat ended', 0, { Guild: 'Olympus', 'd MMM': '5 Dec' }),
    );
    expect(copy.seatEnded.action(4000, 'yearly')).toBe(spec('Seat ended', 1, { price: '$40 a year' }));
    expect(copy.guildRow.youHoldSeat).toBe(spec('Guild row', 0));
    expect(copy.guildRow.active).toBe(spec('Guild row', 1));
  });
});

describe('instants read on the viewer calendar', () => {
  it('an instant near midnight UTC lands on the viewer day', () => {
    const ends = '2027-01-01T05:00:00Z';
    expect(copy.offer.accepted(ends, 'America/Los_Angeles')).toContain('ends 31 Dec 2026');
    expect(copy.offer.accepted(ends, 'Pacific/Auckland')).toContain('ends 1 Jan 2027');
    expect(copy.seatEnded.text('Olympus', ends, 'America/Los_Angeles')).toContain('ended 31 Dec.');
  });
});
