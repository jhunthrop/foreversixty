import { describe, expect, it } from 'vitest';
import { claimEntryFor, claimEntryNotes, rosterGateCopy, rosterGateState } from './roster-gate';

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
    claimedBy: 'hunthrop#1894',
  };

  it('tells a member of a claimed guild who holds it, muted, never "Run this guild?"', () => {
    expect(claimEntryFor(base)).toEqual({ kind: 'note', text: 'Claimed by hunthrop#1894' });
    expect(claimEntryFor({ ...base, claimState: 'contested' })).toEqual({
      kind: 'note',
      text: 'Claimed by hunthrop#1894',
    });
  });

  it('says nothing to a member when the claimant is not named', () => {
    expect(claimEntryFor({ ...base, claimedBy: null })).toEqual({ kind: 'none' });
  });

  it('tells a member of an unclaimed guild that an officer can claim it', () => {
    expect(claimEntryFor({ ...base, claimState: 'unclaimed' })).toEqual({
      kind: 'note',
      text: claimEntryNotes.unclaimedForMember,
    });
  });

  it('offers "Claim it" only to an officer rank in an unclaimed guild', () => {
    expect(claimEntryFor({ ...base, claimState: 'unclaimed', membership: { rank: 'officer' } })).toEqual({
      kind: 'claim',
    });
    // An officer rank the site cannot read the claim for yet (not verified) is offered the page,
    // which states the claim itself.
    expect(claimEntryFor({ ...base, claimState: null, membership: { rank: 'leader' } })).toEqual({
      kind: 'claim',
    });
    expect(claimEntryFor({ ...base, claimState: null })).toEqual({ kind: 'none' });
  });

  it('says nothing to a stranger, a signed-out visitor or a verified officer', () => {
    expect(claimEntryFor({ ...base, membership: null })).toEqual({ kind: 'none' });
    expect(claimEntryFor({ ...base, signedIn: false })).toEqual({ kind: 'none' });
    expect(claimEntryFor({ ...base, canManage: true, membership: { rank: 'officer' } })).toEqual({
      kind: 'none',
    });
  });
});
