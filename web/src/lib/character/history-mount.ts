// web/src/lib/character/history-mount.ts
// What the character page puts under the Forever Score card (experience spec
// 2026-10-24-character-history 3.1 and 3.9), decided in one place from four facts.

export type HistoryMount =
  /** The owner's charts: an owner on a plan with history, for a character with a score. */
  | 'charts'
  /** The free owner's one sentence. */
  | 'sentence'
  /** /v1/me has not answered: the section's reserved place, drawn only under the pre-paint hint. */
  | 'pending'
  | 'none';

export interface HistoryMountFacts {
  /** /v1/me has answered (or failed), so `own` is settled. */
  ownerKnown: boolean;
  /** The signed-in viewer's account holds this character. */
  own: boolean;
  /** `me.entitlements.history`. */
  entitled: boolean;
  hasScore: boolean;
}

export function historyMount(facts: HistoryMountFacts): HistoryMount {
  if (!facts.hasScore) return 'none';
  if (!facts.ownerKnown) return 'pending';
  if (!facts.own) return 'none';
  return facts.entitled ? 'charts' : 'sentence';
}
