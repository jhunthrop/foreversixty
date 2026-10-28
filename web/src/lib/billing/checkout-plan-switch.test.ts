// web/src/lib/billing/checkout-plan-switch.test.ts
import { describe, expect, it, vi } from 'vitest';
import { replacePlanInUrl, withPlan } from './checkout-plan-switch';

describe('withPlan', () => {
  it('rewrites only the plan and keeps every other param', () => {
    expect(withPlan('?plan=guild&interval=yearly&guild_id=7', 'raid_guild')).toBe(
      '?plan=raid_guild&interval=yearly&guild_id=7',
    );
  });

  it('adds the plan when the search has none', () => {
    expect(withPlan('', 'guild')).toBe('?plan=guild');
  });
});

describe('replacePlanInUrl', () => {
  it('calls replaceState with the new plan, keeping the path and hash, and never navigates', () => {
    const replaceState = vi.fn();
    const assign = vi.fn();
    const location = {
      pathname: '/premium/checkout',
      search: '?plan=guild&interval=yearly',
      hash: '#guild-3',
      assign,
    };
    replacePlanInUrl('raid_guild', location, { replaceState, state: { kept: true } });
    expect(replaceState).toHaveBeenCalledWith(
      { kept: true },
      '',
      '/premium/checkout?plan=raid_guild&interval=yearly#guild-3',
    );
    expect(assign).not.toHaveBeenCalled();
  });
});
