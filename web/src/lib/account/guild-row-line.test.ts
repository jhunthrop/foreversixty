import { describe, expect, it } from 'vitest';
import { specCell } from '../../test-support/spec-table';
import { accountPlanCopy } from '../billing/account-plan-copy';
import type { GuildCoveredView } from '../billing/contract';
import { guildRowNextStep, notVerifiedMark, type GuildRowFacts } from './guild-row-line';

const SPEC = new URL('../../../../design/specs/2026-10-27-subscribe-and-join.md', import.meta.url);

function row(extra: Partial<GuildRowFacts> = {}): GuildRowFacts {
  return {
    id: 7,
    name: 'Olympus',
    verified: true,
    verified_count: 12,
    subscribe: { step: 'ready' },
    ...extra,
  };
}

describe('copy', () => {
  it('matches spec 9.3', () => {
    const [notVerified, noPlan, next] = specCell(SPEC, '### 9.3', 'Account guild row');
    expect(accountPlanCopy.guildRow.notVerified).toBe(notVerified);
    expect(accountPlanCopy.guildRow.noPlan('{Guild}')).toBe(noPlan);
    expect(accountPlanCopy.guildRow.nextStep).toBe(next);
  });
});

describe('notVerifiedMark', () => {
  it('shows only for an unverified viewer', () => {
    expect(notVerifiedMark(row({ verified: false }))).toBe('Not verified');
    expect(notVerifiedMark(row())).toBeNull();
  });
});

describe('guildRowNextStep', () => {
  it('links a no-plan guild to its checkout card', () => {
    expect(guildRowNextStep(row())).toEqual({
      text: 'Olympus has no plan.',
      action: 'See the next step',
      label: 'See the next step for Olympus',
      href: '/premium/checkout?plan=guild&interval=yearly#guild-7',
    });
  });
  it('picks raid_guild at 28 verified', () => {
    expect(guildRowNextStep(row({ verified_count: 28 }))?.href).toBe(
      '/premium/checkout?plan=raid_guild&interval=yearly#guild-7',
    );
  });
  it.each(['frozen', 'checkout_open'] as const)('still shows the line on %s', (step) => {
    expect(guildRowNextStep(row({ subscribe: { step } }))).not.toBeNull();
  });
  it('stays quiet on a plan, a covered guild, or an API with no step', () => {
    expect(guildRowNextStep(row({ subscribe: { step: 'on_plan' } }))).toBeNull();
    expect(guildRowNextStep(row({ covered: { tier: 'guild' } as GuildCoveredView }))).toBeNull();
    expect(guildRowNextStep(row({ subscribe: null }))).toBeNull();
    expect(guildRowNextStep(row({ subscribe: undefined }))).toBeNull();
  });
});
