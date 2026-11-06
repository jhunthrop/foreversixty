// web/src/lib/ui/votes.ts
// The Votes primitive's rules (web layout spec 2026-11-05 §7.3): a yes Tag, a no Tag and the count
// of who can vote ("2 yes · 1 no of 4 officers"), the viewer's own vote marked. Pressing your own
// vote again takes it back; pressing the other changes it.

export type VoteChoice = 'yes' | 'no';

/** "of 4 officers", "of 1 officer". */
export function votersLine(voters: number, noun = 'officer'): string {
  return `of ${voters} ${voters === 1 ? noun : `${noun}s`}`;
}

/** The whole tally in words, for a screen reader and a title. */
export function votesLine(yes: number, no: number, voters: number, noun = 'officer'): string {
  return `${yes} yes · ${no} no ${votersLine(voters, noun)}`;
}

/** The vote a press sends: the pressed choice, or null to take the viewer's own back. */
export function nextVote(mine: VoteChoice | null, pressed: VoteChoice): VoteChoice | null {
  return mine === pressed ? null : pressed;
}
