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

/**
 * When the site last saw this character's addon export, for the signed-in viewer who owns it:
 * `/v1/me`'s `build.captured_at` on the matching character, when that newest build came from
 * the addon (a paste or a companion push). Null for anyone else's character, for a build
 * built from Battle.net, and for no build at all.
 */
export function lastAddonExportAt(
  characters: readonly MeCharacter[] | null | undefined,
  path: CharacterPath,
): string | null {
  const own = (characters ?? []).find(
    (character) =>
      character.region === path.region &&
      character.ruleset === path.ruleset &&
      characterSlug(character.name) === decodeSlug(path.slug),
  );
  return own?.build?.source === 'addon' ? own.build.captured_at : null;
}

function decodeSlug(slug: string): string {
  try {
    return decodeURIComponent(slug);
  } catch {
    return slug;
  }
}

/** "Last seen in game 4 Nov 14:03", in the viewer's own clock; null for an unreadable time. */
export function lastSeenText(capturedAt: string): string | null {
  const when = new Date(capturedAt);
  if (Number.isNaN(when.getTime())) return null;
  const hours = String(when.getHours()).padStart(2, '0');
  const minutes = String(when.getMinutes()).padStart(2, '0');
  return scoreCopy.lastSeen(`${when.getDate()} ${MONTHS[when.getMonth()]} ${hours}:${minutes}`);
}
