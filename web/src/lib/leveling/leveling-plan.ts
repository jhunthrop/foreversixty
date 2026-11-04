// web/src/lib/leveling/leveling-plan.ts
// What a level 1 to 19 character is shown instead of "Pick a spec"
// (design/specs/2026-11-04-low-level-home.md): the class's leveling spec as its guide names it,
// and the guide's first talent points in the order they are spent. Browser safe; the plans are
// computed at build time from each guide's own build (pages/data/leveling-plans.json.ts) so the
// home page and the guide's Leveling strip always name the same talents.
import type { MeCharacter } from '../account/api';
import { FIRST_POINT_LEVEL, BASE_LEVEL } from '../planner/types';
import { classSlugFromName } from '../report/tree-sizes';
import { SPECS } from '../sim/specs';

/** The spec each class's guide says to level as (its Leveling sections, checked 2026-11-04). */
export const LEVELING_SPEC_BY_CLASS: Readonly<Record<string, string>> = {
  druid: 'druid-feral',
  hunter: 'hunter-beast-mastery',
  mage: 'mage-frost',
  paladin: 'paladin-retribution',
  priest: 'priest-shadow',
  rogue: 'rogue-combat',
  shaman: 'shaman-enhancement',
  warlock: 'warlock-affliction',
  warrior: 'warrior-arms',
};

/** Where the plans are published. */
export const LEVELING_PLANS_URL = '/data/leveling-plans.json';

/** One point of a guide's leveling order: the talent it goes into and that talent's rank after it. */
export interface LevelingPoint {
  /** The character level the point is spent at (10 for the first). */
  level: number;
  talentId: number;
  name: string;
  icon: string;
  /** The talent's rank once this point is in. */
  rank: number;
  maxRank: number;
  /** Where the talent's digit sits in a talent export: tree position, then cell in tier/column order. */
  tree: number;
  cell: number;
}

/** One spec's leveling plan, as the guide gives it. */
export interface LevelingPlan {
  specKey: string;
  classSlug: string;
  specName: string;
  /** `/guides/hunter/beast-mastery#leveling`. */
  guideHref: string;
  /** The planner with the guide's 20 to 29 build: the talent order through level 20. */
  plannerHref: string;
  /** The spec guide's tooltip bundle, which names every class talent. */
  tipsUrl: string;
  /** The points from level 10 up to the first gear range, in spend order. */
  points: LevelingPoint[];
}

export type LevelingPlans = Readonly<Record<string, LevelingPlan>>;

/** The class guide's own page, for "Other specs". */
export function classGuideHref(classSlug: string): string {
  return `/guides/${classSlug}`;
}

/** The leveling spec key for a character's class; undefined for a class the site does not know. */
export function levelingSpecKey(character: Pick<MeCharacter, 'class'>): string | undefined {
  if (character.class === undefined) return undefined;
  return LEVELING_SPEC_BY_CLASS[classSlugFromName(character.class)];
}

/** The spec guide's Leveling section: `/guides/hunter/beast-mastery#leveling`. */
export function levelingGuideHref(specKey: string): string | undefined {
  const spec = SPECS.find((row) => row.spec === specKey);
  return spec === undefined ? undefined : `/guides/${spec.class_slug}/${spec.spec_slug}#leveling`;
}

/** The display name of a spec key ("Beast Mastery"). */
export function specNameOf(specKey: string): string | undefined {
  return SPECS.find((row) => row.spec === specKey)?.name;
}

/** Talent points a character of `level` has earned. */
export function pointsEarnedAt(level: number): number {
  return Math.max(0, level - BASE_LEVEL);
}

export interface TalentStep {
  /** Points the character can spend now; when the export is unknown, the points earned. */
  toSpend: number;
  /** Index into `plan.points` of the next point to spend; -1 when the plan is used up. */
  nextIndex: number;
  /** The level the next point arrives at, when nothing is waiting to be spent. */
  nextPointLevel: number;
}

function wornRank(trees: readonly string[], point: LevelingPoint): number {
  return Number.parseInt(trees[point.tree]?.[point.cell] ?? '0', 10) || 0;
}

/**
 * Where a character stands in the plan. Without a talent export the guide's order is assumed
 * followed, so the next point is the one this level earns. With one, the next point is the
 * first in the order the character has not matched, and the points to spend are what the
 * level earns minus what is spent.
 */
export function talentStepFor(
  plan: LevelingPlan,
  level: number,
  talents: { trees: readonly string[]; points: readonly number[] } | undefined,
): TalentStep {
  const earned = pointsEarnedAt(level);
  if (talents === undefined) {
    const nextIndex = earned === 0 ? 0 : Math.min(earned - 1, plan.points.length - 1);
    return { toSpend: earned, nextIndex, nextPointLevel: FIRST_POINT_LEVEL };
  }
  const spent = talents.points.reduce((sum, count) => sum + count, 0);
  const nextIndex = plan.points.findIndex((point) => wornRank(talents.trees, point) < point.rank);
  return { toSpend: Math.max(0, earned - spent), nextIndex, nextPointLevel: FIRST_POINT_LEVEL + spent };
}

let plansPromise: Promise<LevelingPlans> | null = null;

/** The published plans, fetched once per page; an empty map when they cannot be read. */
export function loadLevelingPlans(fetcher: typeof fetch = fetch): Promise<LevelingPlans> {
  plansPromise ??= fetcher(LEVELING_PLANS_URL)
    .then((response) => (response.ok ? (response.json() as Promise<LevelingPlans>) : {}))
    .catch(() => {
      plansPromise = null;
      return {};
    });
  return plansPromise;
}
