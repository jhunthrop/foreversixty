// web/src/lib/character-prompt.ts
// What an under-filled character page asks the visitor to do (W.10, UX review r6 F6-5, player
// review P6-7): one primary prompt and at most one quiet secondary, so a page nobody has logged
// does not carry five calls to action for one next step.
//   - No score yet: the primary is "Get the companion" (the score fills in from its logs).
//   - A score but no gear on file: the primary is the planner.
//   - No gear on file: the one quiet line under Gear says how to get gear on the page (sync the
//     addon; a visitor may also sign in to claim the character). With gear it says nothing.

export interface CharacterPromptInput {
  /** The character has a Forever Score. */
  hasScore: boolean;
  /** The page holds gear for the character (at least one slot row). */
  hasGear: boolean;
  /** The viewer's own character. */
  owned: boolean;
}

export interface CharacterPrompt {
  primary: 'companion' | 'planner' | null;
  /** The quiet line under Gear: sync, or sync and claim for a visitor; none with gear on file. */
  secondary: 'sync' | 'sync-or-claim' | null;
}

export function characterPrompt({ hasScore, hasGear, owned }: CharacterPromptInput): CharacterPrompt {
  return {
    primary: !hasScore ? 'companion' : !hasGear ? 'planner' : null,
    secondary: hasGear ? null : owned ? 'sync' : 'sync-or-claim',
  };
}

export interface CharacterNoDataInput {
  hasScore: boolean;
  hasGear: boolean;
  /** Any ranked fight, best or history. */
  hasFights: boolean;
  /** The viewer holds this character on their account. */
  owned: boolean;
}

/**
 * A character the viewer does not hold, with no score, no gear and no fights, gets one sentence
 * ("<Name> has no Forever Sixty data yet."), the lookup box and one quiet addon link, not the
 * empty score bars and the in-game tooltip lookalike (P.1, player review S3-3). The signed-in
 * holder keeps the full set: those empty cards are their to-do list.
 */
export function showsNoDataPage({ hasScore, hasGear, hasFights, owned }: CharacterNoDataInput): boolean {
  return !owned && !hasScore && !hasGear && !hasFights;
}
