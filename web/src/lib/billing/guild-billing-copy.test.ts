import { describe, expect, it } from 'vitest';
import { fillSlots, specKeyCell, specString } from '../../test-support/spec-table';
import type { GuildBillingView } from '../account/api';
import { billedLine, guildBillingCopy as copy, guildPlanLine } from './guild-billing-copy';
import { FOUNDER_WINDOW_CLOSES, GUILD_TRIAL_DAYS } from './plans';

const SPEC = new URL('../../../../design/specs/2026-10-27-launch-pricing.md', import.meta.url);
const spec = (label: string, index: number, fill: Record<string, string> = {}): string =>
  specString(SPEC, '### 6.2', label, index, fill);

describe('plan line', () => {
  it('trial', () => {
    const line = copy.planLine({
      kind: 'trial',
      tier: 'guild',
      firstCharge: '2026-12-11',
      amountCents: 12000,
      interval: 'yearly',
    });
    expect(line).toBe('Guild · trial · first charge 11 Dec 2026: $120');
    expect(line).toBe(spec('Plan line', 0, { date: '11 Dec 2026', amount: '$120' }));
  });

  it('renews, for both tiers', () => {
    const fill = { date: '31 Dec 2026', price: '$150 a year' };
    const active = (tier: 'guild' | 'raid_guild', amountCents: number): string =>
      copy.planLine({ kind: 'active', tier, renews: '2026-12-31', amountCents, interval: 'yearly' });
    expect(active('guild', 15000)).toBe(spec('Plan line', 1, fill));
    expect(active('raid_guild', 25000)).toBe(spec('Plan line', 2, { ...fill, price: '$250 a year' }));
  });

  it('ends, ends then transfers, and lapsed', () => {
    expect(copy.planLine({ kind: 'ending', tier: 'guild', ends: '2026-12-31' })).toBe(
      spec('Plan line', 3, { date: '31 Dec 2026' }),
    );
    expect(
      copy.planLine({ kind: 'ending-transfer', tier: 'guild', ends: '2026-12-31', toName: 'Aela' }),
    ).toBe(spec('Plan line', 4, { date: '31 Dec 2026', Name: 'Aela' }));
    expect(copy.planLine({ kind: 'lapsed', ended: '2026-12-31' })).toBe(
      spec('Plan line', 5, { date: '31 Dec 2026' }),
    );
    expect(copy.founderPill).toBe(spec('Plan line', 6));
  });
});

describe('line 2', () => {
  it('every row', () => {
    expect(copy.line2({ kind: 'trial', cancelBefore: '2026-12-11' })).toBe(
      spec('Line 2', 0, { 'd MMM': '11 Dec' }),
    );
    expect(copy.line2({ kind: 'billed-to-you' })).toBe(spec('Line 2', 1));
    expect(copy.line2({ kind: 'billed-by', name: 'Aela' })).toBe(spec('Line 2', 2, { Name: 'Aela' }));
    expect(copy.line2({ kind: 'ending' })).toBe(spec('Line 2', 3));
    expect(copy.line2({ kind: 'read-only' })).toBe(spec('Line 2', 4));
  });
});

describe('chooser and size line', () => {
  it('chooser rows', () => {
    expect(copy.chooser.tierName('guild')).toBe(spec('Chooser', 0));
    expect(copy.chooser.tierName('raid_guild')).toBe(spec('Chooser', 1));
    expect(copy.chooser.tierLine('guild')).toBe(spec('Chooser', 2));
    expect(copy.chooser.tierLine('raid_guild')).toBe(spec('Chooser', 3));
    expect(copy.chooser.verifiedToday(28)).toBe(spec('Chooser', 4, { n: '28' }));
    expect(copy.chooser.startTrial).toBe(spec('Chooser', 5));
    expect(copy.chooser.subscribe('guild', 'yearly', true)).toBe(
      spec('Chooser', 6, { price: '$120 a year' }),
    );
    expect(copy.chooser.subscribe('raid_guild', 'monthly', false)).toBe(
      spec('Chooser', 6, { price: '$25 a month' }),
    );
    expect(copy.chooser.subscribeAgain).toBe(spec('Chooser', 7));
    // The trial line is B5's `trial` row, which supersedes the §6.1 sentence.
    expect(copy.chooser.trialLine).toBe(
      fillSlots(specKeyCell(SPEC, '### B5. Copy, verbatim', 'trial')[0] ?? '', {
        'GUILD_TRIAL_DAYS + 1': String(GUILD_TRIAL_DAYS + 1),
        GUILD_TRIAL_DAYS: String(GUILD_TRIAL_DAYS),
      }),
    );
    expect(copy.chooser.founderLine).toBe(spec('Chooser', 8));
    expect(copy.chooser.opening).toBe(spec('Chooser', 9));
    expect(copy.chooser.failed).toBe(spec('Chooser', 10));
  });

  it('size line', () => {
    expect(copy.size.line(28)).toBe(spec('Size line', 0, { n: '28' }));
    expect(copy.size.move).toBe(spec('Size line', 1));
  });
});

describe('upgrade confirm', () => {
  it('every row', () => {
    expect(copy.upgrade.sentence(8750, 'yearly', '2027-03-01', 25000)).toBe(
      spec('Upgrade confirm', 0, {
        prorated: '$87.50',
        'year or month': 'year',
        date: '1 Mar 2027',
        price: '$250 a year',
      }),
    );
    expect(copy.upgrade.sentence(500, 'monthly', '2027-03-01', 2500)).toContain('rest of this month');
    expect(copy.upgrade.commit).toBe(spec('Upgrade confirm', 1));
    expect(copy.upgrade.committing).toBe(spec('Upgrade confirm', 2));
    expect(copy.upgrade.keep).toBe(spec('Upgrade confirm', 3));
    expect(copy.upgrade.done).toBe(spec('Upgrade confirm', 4));
    expect(copy.upgrade.failed).toBe(spec('Upgrade confirm', 5));
    expect(copy.upgrade.unconfirmed).toBe(spec('Upgrade confirm', 6));
    expect(copy.upgrade.unconfirmedStale).toBe(spec('Upgrade confirm', 7));
    expect(copy.upgrade.previewExpired).toBe(spec('Upgrade confirm', 8));
  });

  it('the read-only officer line', () => {
    expect(copy.readOnly).toBe(spec('Read-only officer', 0));
  });
});

describe('records and gate', () => {
  it('records rows', () => {
    const rows = [
      copy.records.heading,
      copy.records.body,
      copy.records.attendance,
      copy.records.loot,
      copy.records.preparing,
      copy.records.failed,
    ];
    rows.forEach((text, index) => expect(text).toBe(spec('Records', index)));
    expect(copy.records.fileName('olympus', 'attendance')).toBe(
      spec('Records', 6, { 'guild-slug': 'olympus' }),
    );
    expect(copy.records.fileName('olympus', 'loot')).toBe(spec('Records', 7, { 'guild-slug': 'olympus' }));
  });

  it('gate row', () => {
    expect(copy.gate.text).toBe(spec('Gate, Guild tier', 0));
    expect(copy.gate.action).toBe(spec('Gate, Guild tier', 1));
  });
});

describe('moved plan and billed lines', () => {
  const billing = (overrides: Partial<GuildBillingView> = {}): GuildBillingView => ({
    tier: 'guild',
    you_hold_seat: false,
    status: 'active',
    current_period_end: '2026-12-31T00:00:00Z',
    cancel_at_period_end: false,
    scope: 'plan',
    billed_by: 'hunthrop#1894',
    you_are_billing_contact: false,
    ...overrides,
  });

  it('prints the plan line by who bills and whether it ends', () => {
    expect(guildPlanLine(billing(), 'UTC')).toBe('Guild plan · renews 31 Dec 2026');
    expect(guildPlanLine(billing({ cancel_at_period_end: true }), 'UTC')).toBe(
      'Guild plan · ends 31 Dec 2026',
    );
    expect(guildPlanLine(billing({ billed_by: '' }), 'UTC')).toBe('Guild plan · until 31 Dec 2026');
    expect(guildPlanLine(billing({ current_period_end: null }), 'UTC')).toBe('Guild plan');
    expect(guildPlanLine(billing({ scope: 'officer_tools', billed_by: '' }), 'UTC')).toBe(
      'Officer tools · until 31 Dec 2026',
    );
    expect(guildPlanLine(billing({ scope: 'officer_tools', current_period_end: null }), 'UTC')).toBe(
      'Officer tools',
    );
  });

  it('reads a Stripe period end on the viewer calendar and an operator grant as its stored day', () => {
    const periodEnd = billing({ current_period_end: '2027-01-01T05:00:00Z' });
    expect(guildPlanLine(periodEnd, 'America/Los_Angeles')).toBe('Guild plan · renews 31 Dec 2026');
    expect(guildPlanLine(periodEnd, 'Pacific/Auckland')).toBe('Guild plan · renews 1 Jan 2027');
    const grant = billing({ billed_by: '', current_period_end: '2026-12-31T00:00:00Z' });
    expect(guildPlanLine(grant, 'Pacific/Auckland')).toBe('Guild plan · until 31 Dec 2026');
  });

  it('says who is billed', () => {
    expect(billedLine(billing())).toBe(spec('Line 2', 2, { Name: 'hunthrop#1894' }));
    expect(billedLine(billing({ you_are_billing_contact: true }))).toBe(spec('Line 2', 1));
    expect(billedLine(billing({ billed_by: '' }))).toBe('Granted by Forever Sixty.');
  });
});

describe('founder line and gate spec', () => {
  it('the closing day in the line is the 6th in the founder window own zone', () => {
    const lastSecond = new Date(Date.parse(FOUNDER_WINDOW_CLOSES) - 1000);
    const day = new Intl.DateTimeFormat('en-GB', {
      timeZone: 'America/Los_Angeles',
      day: 'numeric',
      month: 'long',
    }).format(lastSecond);
    expect(copy.chooser.founderLine).toContain(`until ${day} 2027,`);
  });

  it('reads the scoped plan line from the officer tools gate spec', () => {
    const gate = new URL('../../../../design/specs/2026-10-25-officer-tools-gate.md', import.meta.url);
    const expected = specString(gate, '## 7. Copy, verbatim', 'Settings plan line, scoped', 0, {
      'd Mon yyyy': '12 Dec 2026',
    });
    const scoped: GuildBillingView = {
      tier: 'guild',
      you_hold_seat: false,
      status: 'active',
      current_period_end: '2026-12-12T00:00:00Z',
      cancel_at_period_end: false,
      scope: 'officer_tools',
      billed_by: '',
      you_are_billing_contact: false,
    };
    expect(guildPlanLine(scoped, 'UTC')).toBe(expected);
  });
});

describe('the lead-in', () => {
  it('starts the sentence', () => {
    expect(copy.upgrade.sentence(600, 'monthly', '2026-12-01T12:00:00Z', 2500)).toContain(
      copy.upgrade.leadIn,
    );
  });
});
