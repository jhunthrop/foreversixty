// web/src/lib/guild/roster-gate.ts
// What the guild page says in place of the roster to a viewer who cannot read it yet. The
// roster (`GET /v1/guilds/{id}/home`) is for verified members of the guild, so everyone else
// gets one honest sentence under the tabs, never an empty table: a visitor is told whom the
// roster is for, a member still waiting on an officer is told what unlocks it. Pure, so the
// states are pinned without a page.

export type RosterGateState = 'visitor' | 'stranger' | 'unverified';

export interface RosterGateInput {
  /** `/v1/me` answered with a session. */
  readonly signedIn: boolean;
  /** The viewer has a character in this guild, whether or not an officer has verified it. */
  readonly member: boolean;
}

export const rosterGateCopy = {
  heading: 'Roster',
  visitor: 'Roster is for members. Sign in with the account your addon syncs to.',
  stranger: 'Roster is for members of this guild. None of your characters is on it.',
  unverified:
    'Your character is on this guild, waiting for an officer to verify it. The roster opens as soon as they do.',
} as const;

export function rosterGateState({ signedIn, member }: RosterGateInput): RosterGateState {
  if (!signedIn) return 'visitor';
  return member ? 'unverified' : 'stranger';
}

/** What the band says about claiming beyond the claim row, for this viewer: a verified member of a
 *  claimed guild reads who holds it, muted; everyone else nothing (the row, or the settings link). */
export type ClaimEntry = { kind: 'none' } | { kind: 'claimedBy' };

export interface ClaimEntryInput {
  signedIn: boolean;
  /** A verified officer or leader: the Guild settings link in the band is theirs. */
  canManage: boolean;
  /** This account's entry in the guild, when it has one. */
  membership: { rank: string } | null;
  /** The guild's claim state; null while the viewer cannot read it (no member home). */
  claimState: 'unclaimed' | 'pending' | 'claimed' | 'contested' | null;
}

export function claimEntryFor(input: ClaimEntryInput): ClaimEntry {
  const { signedIn, canManage, membership, claimState } = input;
  if (!signedIn || canManage || membership === null) return { kind: 'none' };
  return claimState === 'claimed' || claimState === 'contested' ? { kind: 'claimedBy' } : { kind: 'none' };
}
