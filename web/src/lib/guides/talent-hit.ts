// web/src/lib/guides/talent-hit.ts
// The hit and expertise a band's weights character gets from its talents, so the Stat
// priority caption can name each talent beside the gear and buffs (2026-11-04: Arcane read
// "6% hit from BiS gear and buffs" when all 6% was Precision).
//
// Read from the client's own talent text at the rank taken (`talents/<class>.json`), sorted by
// what the ranker counts in each published figure (`sim/cmd/leveling-bis/hitcap.go`,
// `expertisecap.go`):
// - 'hit': "Improves your chance to hit by 3%" (Precision, Suppression, Tidal Focus,
//   Surefooted, Nature's Reach) is the unit's hit stat, in every melee and spell baseline.
// - 'spell': "...chance to hit with all spells by 6%" (the mage's Precision) is in a caster's
//   baseline only.
// - 'school': "...chance to hit with Shadow spells by 5%" (Shadow Focus) is in the school
//   figure (`BisSpellHitToCap.school`) when one is published, else in the baseline (every
//   spell the rotation casts carries it).
// - 'expertise': "Reduces the chance for your attacks to be Dodged or Parried by 2%" (Weapon
//   Expertise) is in the expertise baseline.
// A talent that adds hit to some abilities only (Furious Precision's off hand, Survival
// Tactics' traps) is in none of them, so it is not named.
import { decodeBandTrees, type DecodedTreeView } from './build-tree';

export type TalentHitScope = 'hit' | 'spell' | 'school' | 'expertise';

export interface TalentHit {
  name: string;
  scope: TalentHitScope;
  /** Percent at the rank taken. */
  percent: number;
}

const PERCENT = String.raw`(\d+(?:\.\d+)?)%`;

const SCOPE_PATTERNS: readonly { scope: TalentHitScope; pattern: RegExp }[] = [
  {
    scope: 'hit',
    pattern: new RegExp(
      String.raw`(?:improves your chance to hit|increases your hit chance) by ${PERCENT}`,
      'i',
    ),
  },
  { scope: 'spell', pattern: new RegExp(String.raw`chance to hit with all spells by ${PERCENT}`, 'i') },
  {
    scope: 'school',
    pattern: new RegExp(
      String.raw`chance to hit with (?:Arcane|Fire|Frost|Holy|Nature|Shadow) spells by ${PERCENT}`,
      'i',
    ),
  },
  {
    scope: 'expertise',
    pattern: new RegExp(String.raw`chance for your attacks to be dodged or parried by ${PERCENT}`, 'i'),
  },
];

/** What one rank's text adds, or undefined when it adds to no published figure. */
export function talentHitOf(name: string, description: string): TalentHit | undefined {
  for (const { scope, pattern } of SCOPE_PATTERNS) {
    const match = pattern.exec(description);
    if (match !== null) return { name, scope, percent: Number(match[1]) };
  }
  return undefined;
}

/** Every talent in `trees` with a rank taken that adds hit or expertise, at that rank. */
export function talentHitsOf(trees: readonly DecodedTreeView[]): TalentHit[] {
  return trees.flatMap((tree) =>
    tree.talents.flatMap((talent) => {
      const description = talent.rank > 0 ? talent.ranks[talent.rank - 1]?.description : undefined;
      const hit = description === undefined ? undefined : talentHitOf(talent.name, description);
      return hit === undefined ? [] : [hit];
    }),
  );
}

/** The hit and expertise talents of a band's weights character (`BisBand.talents`). */
export function bandTalentHits(build: string, classSlug: string, bandTalents: string): TalentHit[] {
  return talentHitsOf(decodeBandTrees(build, classSlug, bandTalents) ?? []);
}
