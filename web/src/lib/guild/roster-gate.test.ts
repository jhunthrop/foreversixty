import { describe, expect, it } from 'vitest';
import { claimEntryFor, rosterGateCopy, rosterGateState } from './roster-gate';

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

describe('claimEntryFor', () => {
  const base = {
    signedIn: true,
    canManage: false,
    membership: { rank: 'member' },
    claimState: 'claimed' as const,
  };

  it('tells a member of a claimed guild who holds it', () => {
    expect(claimEntryFor(base)).toEqual({ kind: 'claimedBy' });
    expect(claimEntryFor({ ...base, claimState: 'contested' })).toEqual({ kind: 'claimedBy' });
  });

  it('leaves an unclaimed or pending guild to the claim row and the Before raids open card', () => {
    for (const claimState of ['unclaimed', 'pending', null] as const) {
      expect(claimEntryFor({ ...base, claimState })).toEqual({ kind: 'none' });
    }
  });

  it('says nothing to a stranger, a signed-out visitor or a verified officer', () => {
    expect(claimEntryFor({ ...base, membership: null })).toEqual({ kind: 'none' });
    expect(claimEntryFor({ ...base, signedIn: false })).toEqual({ kind: 'none' });
    expect(claimEntryFor({ ...base, canManage: true, membership: { rank: 'officer' } })).toEqual({
      kind: 'none',
    });
  });
});
