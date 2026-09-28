// web/src/lib/guild/claim-view.ts
// What the claim page says to THIS viewer, as one pure rule. GuildClaim.svelte used to
// read the claim state only from GET .../settings, which is officer-only, so a plain
// member -- the common visitor, sent here by "Claim this guild" -- got a heading and a
// blank page (found live, 2026-09-28, /guild/us/pvp/olympus-xxvii/claim). The state now
// comes from settings when the viewer may read it, else from the member home, and the
// page always says who can act and what to do next.
import type { Me } from '../account/api';
import type { ClaimStateView, GuildHome, GuildSettingsData } from './api';

export type ClaimViewer =
  /** No session: the state is not readable, the sign-in prompt is the whole page. */
  | 'signed-out'
  /** Signed in, but no character of theirs is in this guild as far as the site knows. */
  | 'stranger'
  /** A member below officer rank: can read the state, cannot claim or confirm. */
  | 'member'
  /** Officer or leader rank in this guild: may claim or confirm. */
  | 'eligible';

export interface ClaimView {
  viewer: ClaimViewer;
  /** The viewer's rank in this guild, when the site knows one. */
  rank: string | null;
  /** The claim state as far as this viewer may read it; null when nothing answered. */
  claim: ClaimStateView | null;
}

export interface ClaimViewInput {
  me: Me | null;
  /** The viewer's entry in `me.guilds` for this guild, matched by the page. */
  membership: { rank?: string } | null;
  settings: GuildSettingsData | null;
  home: GuildHome | null;
}

export function claimView({ me, membership, settings, home }: ClaimViewInput): ClaimView {
  const claim = settings?.claim ?? home?.claim ?? null;
  if (me === null) return { viewer: 'signed-out', rank: null, claim };
  if (membership === null) return { viewer: 'stranger', rank: null, claim };
  const rank = membership.rank ?? null;
  const eligible = rank === 'officer' || rank === 'leader';
  return { viewer: eligible ? 'eligible' : 'member', rank, claim };
}
