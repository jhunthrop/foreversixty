import { render } from 'svelte/server';
import { describe, expect, it } from 'vitest';
import type { GuildRowFacts } from '../../lib/account/guild-row-line';
import AccountGuildRowStep from './AccountGuildRowStep.svelte';

const GUILD: GuildRowFacts = {
  id: 7,
  name: 'Olympus',
  verified: true,
  verified_count: 12,
  subscribe: { step: 'ready' },
};
const html = (guild: GuildRowFacts, part: 'plan' | 'mark' | 'line'): string =>
  render(AccountGuildRowStep, { props: { guild, part } }).body;

describe('AccountGuildRowStep', () => {
  it('marks an unverified viewer', () => {
    expect(html({ ...GUILD, verified: false }, 'mark')).toContain('Not verified');
    expect(html(GUILD, 'mark')).not.toContain('Not verified');
  });

  it('links a no-plan guild to the checkout step', () => {
    const out = html(GUILD, 'line');
    expect(out).toContain('Olympus has no plan.');
    expect(out).toContain('See the next step');
    expect(out).toContain('href="/premium/checkout?plan=guild&amp;interval=yearly#guild-7"');
    expect(html({ ...GUILD, verified_count: 28 }, 'line')).toContain('plan=raid_guild');
  });

  it('tags the plan tier, or No plan when the guild has none', () => {
    expect(
      html({ ...GUILD, plan: { tier: 'raid_guild' }, subscribe: { step: 'on_plan' } }, 'plan'),
    ).toContain('Raid guild');
    expect(html(GUILD, 'plan')).toContain('No plan');
    expect(html({ ...GUILD, subscribe: null }, 'plan')).not.toContain('Tag');
    expect(html({ ...GUILD, subscribe: null }, 'plan')).not.toContain('account-guild-plan-tag');
  });

  it('draws nothing in the line on a plan or without a step', () => {
    expect(html({ ...GUILD, subscribe: { step: 'on_plan' } }, 'line')).not.toContain('no plan');
    expect(html({ ...GUILD, subscribe: null }, 'line')).not.toContain('no plan');
  });
});
