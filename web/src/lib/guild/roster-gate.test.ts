import { describe, expect, it } from 'vitest';
import { rosterGateCopy, rosterGateState } from './roster-gate';

describe('rosterGateState', () => {
  it('is a visitor when signed out, whatever else is true', () => {
    expect(rosterGateState({ signedIn: false, member: false })).toBe('visitor');
    expect(rosterGateState({ signedIn: false, member: true })).toBe('visitor');
  });

  it('is unverified for a signed-in member the roster is still closed to', () => {
    expect(rosterGateState({ signedIn: true, member: true })).toBe('unverified');
  });

  it('is a stranger for a signed-in account with no character in the guild', () => {
    expect(rosterGateState({ signedIn: true, member: false })).toBe('stranger');
  });
});

describe('rosterGateCopy', () => {
  it('tells a visitor whom the roster is for and which account to sign in with', () => {
    expect(rosterGateCopy.visitor).toBe(
      'Roster is for members. Sign in with the account your addon syncs to.',
    );
  });
});
