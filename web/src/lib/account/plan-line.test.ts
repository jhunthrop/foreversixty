import { readFileSync } from 'node:fs';
import { describe, expect, it } from 'vitest';
import {
  FOUNDER_PREMIUM_FACTS,
  ME_ENDED_SEAT,
  ME_SEAT,
  mockEntitlements,
} from '../../test-support/billing-fixtures';
import type { BillingView, GuildBillingView, Me } from './api';
import { premiumSourceCopy } from '../billing/premium-source-copy';
import { specCell, specString } from '../../test-support/spec-table';
import { accountPlanLine, guildRowPlanLine, onPremium, premiumSource } from './plan-line';

const ZONE = 'UTC';
const NOW = new Date('2026-11-12T12:00:00Z');

const OWN: BillingView = {
  plan: 'premium',
  status: 'active',
  current_period_end: '2026-12-20T00:00:00Z',
  cancel_at_period_end: false,
  founder: false,
  interval: 'monthly',
  amount_cents: 600,
};

function meWith(extra: Partial<Me>, billing: BillingView | null = null): Me {
  return {
    user: { id: 7, battletag: 'Fixture#1234', email: null, role: 'user', anonymize: false },
    characters: [],
    guilds: [],
    entitlements: mockEntitlements({ billing }),
    ...extra,
  };
}

describe('accountPlanLine', () => {
  it('own Premium renews with its price and keeps Manage billing', () => {
    const out = accountPlanLine(meWith({}, OWN), NOW, ZONE);
    expect(out).toEqual({
      line: 'Premium · renews 20 Dec 2026 · $6 a month',
      line2: null,
      head: {
        name: 'Premium',
        when: { label: 'Renews', date: '20 Dec 2026' },
        extra: '$6 a month',
      },
      founder: false,
      billingButton: 'manage',
      ended: null,
    });
  });

  it('a founder price carries the Founder flag', () => {
    const out = accountPlanLine(meWith({}, { ...OWN, ...FOUNDER_PREMIUM_FACTS }), NOW, ZONE);
    expect(out.founder).toBe(true);
  });

  it('own Premium that stops renewing reads its end', () => {
    const out = accountPlanLine(meWith({}, { ...OWN, cancel_at_period_end: true }), NOW, ZONE);
    expect(out.line).toBe('Premium · ends 20 Dec 2026');
  });

  it('own Premium ending into a seat names the seat and keeps Manage billing', () => {
    const out = accountPlanLine(meWith({ seat: ME_SEAT }, OWN), NOW, ZONE);
    expect(out.line).toBe('Premium · ends 20 Dec 2026 · then a seat from The Last Watch');
    expect(out.billingButton).toBe('manage');
    expect(out.line2).toBeNull();
  });

  it('a seat holder reads who gave it and has no Manage billing', () => {
    const seat = { ...ME_SEAT, starts_at: null };
    const out = accountPlanLine(meWith({ seat }), NOW, ZONE);
    expect(out.line).toBe('Premium · a seat from The Last Watch');
    expect(out.line2).toBe(
      "Given by Thornwick. It lasts while the guild's plan does; nothing is billed to you.",
    );
    expect(out.billingButton).toBeNull();
    expect(out.founder).toBe(false);
  });

  it('a seat whose start has passed reads as a seat even beside a lapsed own plan', () => {
    const lapsed = { ...OWN, status: 'canceled' };
    const out = accountPlanLine(meWith({ seat: ME_SEAT }, lapsed), new Date('2026-12-21T00:00:00Z'), ZONE);
    expect(out.line).toBe('Premium · a seat from The Last Watch');
    expect(out.billingButton).toBeNull();
  });

  it('free reads Not on Premium with no ended seat', () => {
    expect(accountPlanLine(meWith({}), NOW, ZONE)).toEqual({
      line: 'Not on Premium.',
      line2: null,
      head: { name: 'Free', when: null, extra: null },
      founder: false,
      billingButton: null,
      ended: null,
    });
  });

  it('an ended seat offers the comeback link at the price they paid', () => {
    const out = accountPlanLine(meWith({ ended_seat: ME_ENDED_SEAT }), NOW, ZONE);
    expect(out.line).toBe('Not on Premium.');
    expect(out.ended).toEqual({
      text: 'Your seat from The Last Watch ended 18 Nov. ',
      action: 'Come back at $6 a month',
      href: '/premium/checkout?plan=premium&interval=monthly&comeback=9',
    });
  });

  it('an ended seat with no remembered price says Come back', () => {
    const ended = { ...ME_ENDED_SEAT, amount_cents: null, interval: null };
    expect(accountPlanLine(meWith({ ended_seat: ended }), NOW, ZONE).ended?.action).toBe('Come back');
  });
});

describe('guildRowPlanLine', () => {
  const plan: GuildBillingView = {
    tier: 'guild',
    you_hold_seat: false,
    scope: 'plan',
    status: 'active',
    current_period_end: '2026-12-12T00:00:00Z',
    cancel_at_period_end: false,
    billed_by: 'Fixture#1234',
    you_are_billing_contact: false,
  };

  it('reads the verb and date alone, whatever the tier (the tag names it)', () => {
    expect(guildRowPlanLine(plan, ZONE)).toBe('Renews 12 Dec 2026');
    expect(guildRowPlanLine({ ...plan, tier: 'raid_guild' }, ZONE)).toBe('Renews 12 Dec 2026');
    expect(guildRowPlanLine({ ...plan, cancel_at_period_end: true }, ZONE)).toBe('Ends 12 Dec 2026');
  });

  it('reads Active when the plan has no date', () => {
    expect(guildRowPlanLine({ ...plan, current_period_end: null }, ZONE)).toBe('Active');
  });

  it('adds the seat note when the account holds one', () => {
    expect(guildRowPlanLine({ ...plan, you_hold_seat: true }, ZONE)).toBe(
      'Renews 12 Dec 2026 · you hold a seat',
    );
  });
});

describe('onPremium: an account on Premium never sees the free See plans link', () => {
  it('own plan, a seat and a resolved grant are all Premium', () => {
    expect(onPremium(meWith({}, OWN))).toBe(true);
    expect(onPremium(meWith({ seat: ME_SEAT }))).toBe(true);
    expect(onPremium(meWith({ entitlements: mockEntitlements({ server_sims: true }) }))).toBe(true);
    expect(onPremium(GRANT_OPEN)).toBe(true);
  });

  it('a free account and no account are not', () => {
    expect(onPremium(meWith({}))).toBe(false);
    expect(onPremium(null)).toBe(false);
    expect(onPremium(meWith({}, { ...OWN, status: 'canceled' }))).toBe(false);
  });
});

const SPEC = new URL('../../../../design/specs/2026-10-27-subscribe-and-join.md', import.meta.url);

/** Line 2 (the third column) of a section 4.3 row, by the row's first-column label. */
function line2Cell(label: string): string[] {
  const spec = readFileSync(SPEC, 'utf8');
  const rows = spec.slice(spec.indexOf('### 4.3')).split('\n');
  const row = rows.find((line) => line.startsWith(`| ${label} |`));
  if (row === undefined) throw new Error(`4.3 has no row "${label}"`);
  return [...(row.split('|')[3] ?? '').matchAll(/`([^`]+)`/g)].map((m) => m[1] ?? '');
}

const STRIPE: BillingView = { ...OWN, source: 'stripe' };
const GRANT_OPEN = meWith(
  {},
  { ...OWN, source: 'grant', current_period_end: null, amount_cents: null, interval: null },
);
const GRANT_UNTIL = meWith(
  {},
  { ...OWN, source: 'grant', current_period_end: '2027-03-31T00:00:00Z', amount_cents: null, interval: null },
);
const SEAT = meWith({ seat: { ...ME_SEAT, starts_at: null } });
const FALLBACK = meWith({ entitlements: mockEntitlements({ server_sims: true }) });

function withPortal(me: Me): Me {
  return { ...me, entitlements: mockEntitlements({ ...me.entitlements, personal_portal: true }) };
}

describe('premiumSource (spec 4.1)', () => {
  it.each([
    ['own renewing', meWith({}, STRIPE), 'own'],
    ['own without a source (older API)', meWith({}, OWN), 'own'],
    ['own ending', meWith({}, { ...STRIPE, cancel_at_period_end: true }), 'own'],
    ['past due', meWith({}, { ...STRIPE, status: 'past_due' }), 'own'],
    ['trialing', meWith({}, { ...STRIPE, status: 'trialing' }), 'own'],
    ['own then seat', meWith({ seat: ME_SEAT }, STRIPE), 'own_then_seat'],
    ['grant open-ended', GRANT_OPEN, 'grant'],
    ['grant until', GRANT_UNTIL, 'grant'],
    ['seat', SEAT, 'seat'],
    ['lapsed own plus a seat', meWith({ seat: ME_SEAT }, { ...STRIPE, status: 'canceled' }), 'seat'],
    ['fallback', FALLBACK, 'fallback'],
    ['free', meWith({}), 'none'],
    ['lapsed grant', meWith({}, { ...OWN, source: 'grant', status: 'canceled' }), 'none'],
  ] as const)('%s', (_name, me, want) => {
    expect(premiumSource(me, NOW)).toBe(want);
  });
});

describe('accountPlanLine by source (spec 4.3, 9.2)', () => {
  it('grant, open-ended: no longer reads Not on Premium', () => {
    const out = accountPlanLine(GRANT_OPEN, NOW, ZONE);
    expect(out.line).toBe(specString(SPEC, '### 9.2', 'Grant lines', 0));
    expect(out.line2).toBe(specString(SPEC, '### 9.2', 'Grant lines', 2));
    expect(out.line).toBe(specCell(SPEC, '### 4.3', '`grant`, open-ended')[0]);
    expect(out.line2).toBe(line2Cell('`grant`, open-ended')[0]);
    expect(out.billingButton).toBeNull();
    expect(out.founder).toBe(false);
  });

  it('grant until a date', () => {
    const out = accountPlanLine(GRANT_UNTIL, NOW, ZONE);
    expect(out.line).toBe(specString(SPEC, '### 9.2', 'Grant lines', 1, { date: '31 Mar 2027' }));
    expect(out.line2).toBe(specString(SPEC, '### 9.2', 'Grant lines', 3));
  });

  it('own rows read section 4.3 verbatim', () => {
    const end = '2026-12-22T00:00:00Z';
    const renewing = accountPlanLine(
      meWith({}, { ...STRIPE, amount_cents: 500, current_period_end: end }),
      NOW,
      ZONE,
    );
    expect(renewing.line).toBe(specCell(SPEC, '### 4.3', '`own`, renewing')[0]);
    const ending = accountPlanLine(
      meWith({}, { ...STRIPE, cancel_at_period_end: true, current_period_end: end }),
      NOW,
      ZONE,
    );
    expect(ending.line).toBe(specCell(SPEC, '### 4.3', '`own`, ending')[0]);
    const seat = { ...ME_SEAT, guild_name: 'Olympus XXVII' };
    const then = accountPlanLine(meWith({ seat }, { ...STRIPE, current_period_end: end }), NOW, ZONE);
    expect(then.line).toBe(specCell(SPEC, '### 4.3', '`own_then_seat`')[0]);
  });

  it('seat reads section 4.3 verbatim', () => {
    const seat = { ...ME_SEAT, starts_at: null, guild_name: 'Olympus XXVII', given_by: 'Kraggor' };
    const out = accountPlanLine(meWith({ seat }), NOW, ZONE);
    expect(out.line).toBe(specCell(SPEC, '### 4.3', '`seat`')[0]);
    expect(out.line2).toBe(line2Cell('`seat`')[0]);
  });

  it('fallback reads Premium with no button', () => {
    const out = accountPlanLine(withPortal(FALLBACK), NOW, ZONE);
    expect(out.line).toBe(specCell(SPEC, '### 4.3', '`fallback`')[0]);
    expect(out.billingButton).toBeNull();
  });

  it('Manage billing for own and own_then_seat; Billing history for the rest with a portal', () => {
    const cases: [Me, string | null][] = [
      [meWith({}, STRIPE), 'manage'],
      [meWith({ seat: ME_SEAT }, STRIPE), 'manage'],
      [withPortal(SEAT), 'history'],
      [withPortal(GRANT_OPEN), 'history'],
      [withPortal(meWith({})), 'history'],
      [SEAT, null],
      [GRANT_UNTIL, null],
      [meWith({}), null],
    ];
    for (const [me, want] of cases) expect(accountPlanLine(me, NOW).billingButton).toBe(want);
  });

  it('past due keeps the renewing line', () => {
    const out = accountPlanLine(meWith({}, { ...STRIPE, status: 'past_due' }), NOW, ZONE);
    expect(out.line).toContain('renews');
    expect(out.line2).toBeNull();
  });

  it('own then seat has no line 2 and keeps the founder flag', () => {
    const founder = { ...STRIPE, ...FOUNDER_PREMIUM_FACTS };
    const out = accountPlanLine(meWith({ seat: ME_SEAT }, founder), NOW, ZONE);
    expect(out.line2).toBeNull();
    expect(out.founder).toBe(true);
  });

  it('only the free plan carries an ended seat', () => {
    const ended = { ended_seat: ME_ENDED_SEAT };
    const sources = [
      meWith(ended, STRIPE),
      { ...GRANT_OPEN, ...ended },
      { ...SEAT, ...ended },
      { ...FALLBACK, ...ended },
    ];
    for (const me of sources) expect(accountPlanLine(me, NOW, ZONE).ended).toBeNull();
    expect(accountPlanLine(meWith(ended), NOW, ZONE).ended).not.toBeNull();
  });

  it('a source-less row without a price is a grant, never own', () => {
    const bare = { ...OWN, amount_cents: null, interval: null };
    const open = accountPlanLine(meWith({}, { ...bare, current_period_end: null }), NOW, ZONE);
    expect(open.line).toBe(premiumSourceCopy.grant.noEnd);
    expect(open.billingButton).toBeNull();
    const dated = accountPlanLine(
      meWith({}, { ...bare, current_period_end: '2027-03-31T00:00:00Z' }),
      NOW,
      ZONE,
    );
    expect(dated.line).toBe('Premium · granted until 31 Mar 2027');
    expect(dated.billingButton).toBeNull();
  });

  it('a source-less priced row is own with Manage billing', () => {
    expect(premiumSource(meWith({}, OWN), NOW)).toBe('own');
    expect(accountPlanLine(meWith({}, OWN), NOW, ZONE).billingButton).toBe('manage');
  });
});

describe('premiumSourceCopy (spec 9.2)', () => {
  it('matches the spec table', () => {
    expect(premiumSourceCopy.panelTitle).toBe(specString(SPEC, '### 9.2', 'Panel title'));
    expect(premiumSourceCopy.fallback).toBe(specString(SPEC, '### 9.2', 'Fallback line'));
    expect(Object.values(premiumSourceCopy.buttons)).toEqual(specCell(SPEC, '### 9.2', 'Buttons'));
    expect(Object.values(premiumSourceCopy.premiumCard)).toEqual(
      specCell(SPEC, '### 9.2', '`/premium` card line'),
    );
  });
});
