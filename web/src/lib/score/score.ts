// web/src/lib/score/score.ts
// The Forever Score as the site shows it (spec 2026-10-11 §4.6): the fields
// GET /v1/characters/{region}/{ruleset}/{name} carries (api/openapi.yaml `ForeverScore`), the
// formula's ceilings and tier lines (api/internal/rating/score_v2.go, the formula's only
// home), and the reads every surface shares. Like the addon's Score.lua, nothing here
// computes a score or works out a tier: the API sends both.
//
// Labels and tier colours come from docs/contracts/fixtures/score-labels.json, the fixture
// addon/tests/score_labels_spec.lua pins to Locale.lua, so the site and the tooltip use the
// same words and the same client item-quality colours.
import labels from '../../../../docs/contracts/fixtures/score-labels.json';
import type { MeCharacter } from '../account/api';
import { characterSlug, type CharacterPath } from '../characters';
import { scoreCopy } from './copy';

export type ScoreComponentKey = 'gear' | 'experience' | 'parses' | 'attendance' | 'discipline';

export type ScoreComponents = Record<ScoreComponentKey, number>;

export interface ScoreExperience {
  dungeons: number;
  raid_kills: number;
  alts: number;
}

/** Every field is null when the character has no rated fight in a public report, or its
 *  account is anonymous: "no logs yet", never a zero. */
export interface ForeverScore {
  score: number | null;
  tier: string | null;
  /** The newest rated fight's date, YYYY-MM-DD. */
  as_of: string | null;
  components: ScoreComponents | null;
  experience: ScoreExperience | null;
}

/** One date of a character's score history (api/openapi.yaml `ScorePoint`): the score the
 *  rater recorded on that UTC date's last rated fight. */
export interface ScorePoint {
  /** YYYY-MM-DD, UTC. */
  date: string;
  score: number;
  tier: string;
}

/** The most points each component can give (score_v2.go's *Max constants). */
export const COMPONENT_MAX: Readonly<ScoreComponents> = {
  gear: 500,
  experience: 700,
  parses: 600,
  attendance: 200,
  discipline: 100,
};

export type Tier = keyof typeof labels.tiers;

/** The lowest score of each tier, lowest tier first (score_v2.go's *From constants;
 *  provisional there, so provisional here). */
export const TIER_FROM: Readonly<Record<Tier, number>> = {
  common: 0,
  uncommon: 500,
  rare: 1000,
  epic: 1500,
  legendary: 2000,
};

const TIERS = Object.keys(TIER_FROM) as Tier[];

function isTier(tier: string | null | undefined): tier is Tier {
  return typeof tier === 'string' && Object.hasOwn(labels.tiers, tier);
}

/** "below 500", "500 to 999", ..., "2000 and up": one per tier, lowest first. */
export function tierRanges(): { tier: Tier; range: string }[] {
  return TIERS.map((tier, index) => {
    const from = TIER_FROM[tier];
    const next = TIERS[index + 1];
    const range =
      next === undefined
        ? scoreCopy.tierFrom(from)
        : from === 0
          ? scoreCopy.tierBelow(TIER_FROM[next])
          : scoreCopy.tierBetween(from, TIER_FROM[next] - 1);
    return { tier, range };
  });
}

/** The client's item-quality colour for a tier, or null for a tier the site does not know
 *  (shown uncoloured, as Score.tierColor does in game). */
export function tierColor(tier: string | null | undefined): string | null {
  return isTier(tier) ? labels.tiers[tier].color : null;
}

export function tierLabel(tier: string | null | undefined): string | null {
  return isTier(tier) ? labels.tiers[tier].label : null;
}

/**
 * The colour for a tier's name written at body size. Rare and epic are too dark to read as
 * small text on the site's ground (3.9:1), so words use design/DESIGN-SYSTEM.md's lightened
 * variants (tokens.css `--color-rarity-rare-text` / `-epic-text`); the big number, bars,
 * borders and the tier strip keep the exact quality colour from tierColor.
 */
const TIER_TEXT: Readonly<Partial<Record<Tier, string>>> = { rare: '#3d94f0', epic: '#b866f5' };

export function tierTextColor(tier: string | null | undefined): string | null {
  return isTier(tier) ? (TIER_TEXT[tier] ?? labels.tiers[tier].color) : null;
}

/** The most a score can be: the five ceilings added up. */
export const SCORE_MAX = Object.values(COMPONENT_MAX).reduce((sum, max) => sum + max, 0);

const COMPONENT_KEYS = Object.keys(COMPONENT_MAX) as ScoreComponentKey[];

/**
 * The part with the most points left to earn: the largest gap between a component and its
 * ceiling (ties go to the earlier part, in the addon's order). Null when every part is full
 * or there are no parts.
 */
export function mostToGain(
  components: ScoreComponents | null | undefined,
): { key: ScoreComponentKey; gap: number } | null {
  return componentGaps(components)[0] ?? null;
}

/**
 * Every part's distance to its ceiling, largest first (L8.4 fix round 5 item 5), ties in the
 * addon's order, so the first is the part "Most to gain" names. A part at its ceiling, or with
 * no number, is left out; empty when there are no parts.
 */
export function componentGaps(
  components: ScoreComponents | null | undefined,
): { key: ScoreComponentKey; gap: number }[] {
  if (components == null) return [];
  const gaps: { key: ScoreComponentKey; gap: number }[] = [];
  for (const key of COMPONENT_KEYS) {
    const points = components[key];
    if (typeof points !== 'number') continue;
    const gap = COMPONENT_MAX[key] - points;
    if (gap > 0) gaps.push({ key, gap });
  }
  // Array.prototype.sort is stable, so equal gaps keep the addon's order.
  return gaps.sort((a, b) => b.gap - a.gap);
}

/**
 * The one colour for the part to work on (fix round 2 item 2): the bar "Most to gain" names,
 * never a tier colour, so it cannot be read as a tier. The design system's `--gold-hover`
 * amber, which the "+88" cue beside the bar repeats in words.
 */
export const NEEDS_WORK_COLOR = '#f5d27a';

/** The five tiers as the strip draws them: label, exact colour, name colour, short range. */
export function tierStrip(): {
  tier: Tier;
  label: string;
  color: string;
  textColor: string;
  range: string;
}[] {
  return TIERS.map((tier, index) => {
    const next = TIERS[index + 1];
    return {
      tier,
      label: labels.tiers[tier].label,
      color: labels.tiers[tier].color,
      textColor: TIER_TEXT[tier] ?? labels.tiers[tier].color,
      range: scoreCopy.tierShort(TIER_FROM[tier], next === undefined ? null : TIER_FROM[next] - 1),
    };
  });
}

/**
 * The line under the score (L8.4 fix round 4 item 1): the points to the next tier's line and
 * that tier, "170 to Legendary (2000)", or "Top tier" at or past the top tier's line. From the
 * score and the tier lines alone, so it reads the same wherever the tier came from.
 */
export function nextTierLine(score: number): string {
  const next = TIERS.find((tier) => TIER_FROM[tier] > score);
  return next === undefined
    ? scoreCopy.topTier
    : scoreCopy.toNextTier(TIER_FROM[next] - score, labels.tiers[next].label, TIER_FROM[next]);
}

/** The score fields off a character page; a page that predates them reads as no score. */
export function foreverOf(page: Partial<ForeverScore> | null | undefined): ForeverScore {
  return {
    score: page?.score ?? null,
    tier: page?.tier ?? null,
    as_of: page?.as_of ?? null,
    components: page?.components ?? null,
    experience: page?.experience ?? null,
  };
}

export function hasScore(
  forever: ForeverScore | null | undefined,
): forever is ForeverScore & { score: number } {
  return typeof forever?.score === 'number';
}

/** Compat.shortDate's month names, so the site's "as of 4 Nov" is the tooltip's. */
const MONTHS = ['Jan', 'Feb', 'Mar', 'Apr', 'May', 'Jun', 'Jul', 'Aug', 'Sep', 'Oct', 'Nov', 'Dec'];

/** "2026-11-04" reads "4 Nov" (the addon's Compat.shortDate); anything else is null. */
export function shortDate(text: string | null | undefined): string | null {
  const match = /^\d{4}-(\d{2})-(\d{2})$/.exec(text ?? '');
  if (match === null) return null;
  const month = MONTHS[Number(match[1]) - 1];
  const day = Number(match[2]);
  if (month === undefined || day < 1 || day > 31) return null;
  return `${day} ${month}`;
}

/** "as of 4 Nov" for a score's date, or null when it has none. */
export function asOfText(forever: ForeverScore | null | undefined): string | null {
  const date = shortDate(forever?.as_of);
  return date === null ? null : scoreCopy.asOf(date);
}

/** The signed-in viewer's own row for the character at `path`, if it is theirs. */
export function ownCharacter(
  characters: readonly MeCharacter[] | null | undefined,
  path: CharacterPath,
): MeCharacter | null {
  return (
    (characters ?? []).find(
      (character) =>
        character.region === path.region &&
        character.ruleset === path.ruleset &&
        characterSlug(character.name) === decodeSlug(path.slug),
    ) ?? null
  );
}

function decodeSlug(slug: string): string {
  try {
    return decodeURIComponent(slug);
  } catch {
    return slug;
  }
}
