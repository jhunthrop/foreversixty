// web/src/lib/guild/roster-links.test.ts
// current-character.ts's plannerHrefFor/simHrefFor produce identical output to
// handoff-links.ts's plannerCodeHref/simCodeHref for a 'code'-sourced character (both
// build `/planner?code=<ref>` and, for the default quick-sim tab, `/sim?code=<ref>` with
// every other sim-state field at its default) -- confirmed by reading both modules before
// writing this. This module is still the one place that prefers the richer
// current-character.ts functions, per spec section 4.1 and CharacterHandoffLinks.svelte's
// own compatibility rule, so a future change to either module's behaviour needs to change
// only this file.
import { describe, expect, it } from 'vitest';
import { rosterHoldsHandoff, rosterPlannerHref, rosterSimHref } from './roster-links';
import { rosterProfileHref } from './roster-profile';

describe('rosterHoldsHandoff', () => {
  it('holds the slot only for shared gear whose export the simulator reads', () => {
    expect(rosterHoldsHandoff({ consent: 'gear', sim_export: true })).toBe(true);
    expect(rosterHoldsHandoff({ consent: 'gear_bags', sim_export: true })).toBe(true);
    expect(rosterHoldsHandoff({ consent: 'gear', sim_export: false })).toBe(false);
    expect(rosterHoldsHandoff({ consent: 'roster', sim_export: true })).toBe(false);
  });

  it('keeps the slot, for its lookup to decide, when the API predates sim_export', () => {
    expect(rosterHoldsHandoff({ consent: 'gear' })).toBe(true);
    expect(rosterHoldsHandoff({ consent: 'roster' })).toBe(false);
  });
});

const CODE = 'FS1:1.60.1.69893:warrior:human:0/0/0:';

describe('rosterPlannerHref', () => {
  it('builds a code-carrying planner link', () => {
    expect(rosterPlannerHref(CODE)).toBe(`/planner?code=${encodeURIComponent(CODE)}`);
  });
});

describe('rosterSimHref', () => {
  it('builds a code-carrying simulator link at the default quick-sim tab', () => {
    expect(rosterSimHref(CODE)).toBe(`/sim?code=${encodeURIComponent(CODE)}`);
  });
});

describe('rosterProfileHref', () => {
  const guildPath = '/guild/us/pvp/olympus-xxvii';
  const row = { character_key: 'us/pvp/zulmara', account_key: 'u:1001', verified: true };

  it('opens anyone for an officer, by account, and only their own row for a member', () => {
    const officer = { officer: true, myCharacterKey: 'us/pvp/kraggor', guildPath };
    expect(rosterProfileHref(row, officer)).toBe(`${guildPath}/members/u%3A1001`);
    const member = { officer: false, myCharacterKey: 'us/pvp/zulmara', guildPath };
    expect(rosterProfileHref(row, member)).toBe(`${guildPath}/members/me`);
    expect(rosterProfileHref({ ...row, character_key: 'us/pvp/other' }, member)).toBeNull();
  });

  it('links no unverified row, no row without an account and nothing before the guild path', () => {
    const officer = { officer: true, myCharacterKey: null, guildPath };
    expect(rosterProfileHref({ ...row, verified: false }, officer)).toBeNull();
    expect(rosterProfileHref({ ...row, account_key: undefined }, officer)).toBeNull();
    expect(rosterProfileHref(row, { ...officer, guildPath: '' })).toBeNull();
  });
});
