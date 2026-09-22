// web/src/lib/rating/copy.ts
// Every visible string the rating surfaces use, in the site's honest-copy voice: no
// exclamation marks, no "AI", no claim the engine does not make. Strings quoted directly
// from spec §6.4/§1.5 are copied verbatim; the rest follow their tone.

/** Spec §6.4's own worked-example threshold for showing the character trend sparkline —
 *  a display-only choice (Ruling 7), not a server-side minimum. */
export const MIN_TREND_SAMPLES = 5;

export const ratingCopy = {
  tabLabel: 'Rating',
  headingFor: (playerName: string): string => `${playerName}’s rating`,
  explainLink: 'How is this calculated?',
  panelHeading: 'Performance rating',
  panelMore: 'Rating tab',
  panelEmpty: 'No ratings for this fight yet.',
  /** Under half the role's weight was measurable, so there is no overall. */
  insufficient: 'Not rated',
  insufficientNote: (reason: string): string =>
    `Not enough is measured on this fight to rate this player: ${reason}. The parts below are what was measured.`,
  tabEmpty:
    'No ratings for this fight yet. Ratings are computed after a report finishes processing, and some encounters do not have a curated table yet.',
  fetchFailed: 'Ratings did not load.',
  nightNotOnePull: 'Ratings are one pull’s. Pick a boss pull from the list to see one.',
  enemiesHaveNone: 'Ratings are for your raid, not the enemy.',
  noMatchingRow: 'This player has no rating for this fight.',
  /** Spec §1.2's two exact basis sentences. */
  percentileBasis: (spec: string, klass: string, role: string, pct: number, n: number): string =>
    `Compared with other ${spec} ${roleNoun(klass, role)}s on this fight (${n} logs).`,
  absoluteBasis: 'Measured against the encounter’s own numbers — not enough logs yet to compare players.',
  /** Spec §7.4's fixed excluded-reason strings, by the `reason` code the API returns. */
  excludedReason: (componentLabel: string, reason: string): string => {
    switch (reason) {
      case 'no_mechanics_table':
        return `${componentLabel} — not scored. This fight’s encounter has no curated mechanics table yet.`;
      case 'spec_not_modeled':
        return `${componentLabel} — not scored. The simulator does not model this spec yet; scores return once it does.`;
      case 'no_utility_table':
        return `${componentLabel} — not scored. No curated utility table exists for this spec yet.`;
      case 'no_consumable_catalogue':
        return `${componentLabel} — not scored. No consumable catalogue exists for this role yet.`;
      case 'unclaimed_and_unranked':
        return `${componentLabel} — not scored. This character is not signed in, and the bracket has too few logs to compare against.`;
      default:
        return `${componentLabel} — not scored.`;
    }
  },
  threatNotModeled: 'Threat — not modeled yet. This does not count for or against Utility.',
  /** Spec §6.4's exact capped-score sentence. */
  cappedNote:
    'capped from a higher weighted average — a costly avoidable death outweighs the rest of the fight. See Survival.',
  /** Spec §1.5's exact rule sentence, published on /ratings. */
  cappedRule:
    'An avoidable death early in a fight caps the overall score at 40, because nothing else in the fight makes up for it.',
  trendTooFew: (have: number): string =>
    `Not enough rated fights yet to show a trend (${have} of ${MIN_TREND_SAMPLES} needed).`,
  characterEmpty: 'No rated fights yet.',
} as const;

const COMPONENT_LABELS: Record<string, string> = {
  output: 'Output',
  survival: 'Survival',
  mechanics: 'Mechanics',
  utility: 'Utility',
  preparation: 'Preparation',
  activity: 'Activity',
};

export function componentLabel(name: string): string {
  return COMPONENT_LABELS[name] ?? name;
}

/** "Fury" + "Warrior" -> "Fury Warriors"; every vanilla class name pluralises with a bare
 *  "s", so no irregular table is needed. `role` is unused today (the bracket key includes
 *  it, but the copy the spec quotes names the class, not the role) and kept as a parameter
 *  so a future copy revision that needs it has no signature to change. */
// eslint-disable-next-line @typescript-eslint/no-unused-vars -- kept for the reason above
function roleNoun(klass: string, _role: string): string {
  return `${klass}s`;
}
