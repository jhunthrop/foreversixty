// web/src/lib/guild/claim-row.ts
// The band's claim row (spec 2026-11-05 §3.1): who sees it and which form. Pure, so the states are
// pinned without a page. It reads `/v1/me`, which an unverified guild master has too (`/home` does
// not answer them), so the row draws for them on every tab.
import type { ClaimState } from './api';

export type ClaimRowKind = 'guild-master' | 'officer' | 'pending';

export function claimRowKind(
  rank: string | null | undefined,
  state: ClaimState | null | undefined,
): ClaimRowKind | null {
  const officerRank = rank === 'leader' || rank === 'officer';
  if (!officerRank) return null;
  if (state === 'pending') return 'pending';
  if (state !== 'unclaimed') return null;
  return rank === 'leader' ? 'guild-master' : 'officer';
}
