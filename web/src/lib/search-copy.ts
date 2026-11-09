// web/src/lib/search-copy.ts
// The <title> and meta description of the pages people find through search, written for the words
// they type (design/marketing/2026-11-08-marketing-plan.md, Search): spec before class, "WoW
// Forever" in the title, the brand short at the end (lib/seo.ts pageTitle adds it when it fits).
// The page's own heading keeps its own words; these only go to Base.astro. No em dashes.
// The BiS pages' own are lib/bis/seo-copy.ts (they need the band and the ranked data).
import type { BisRole } from './bis/types';

const GAME = 'WoW Forever';
const GAME_LONG = 'World of Warcraft: Forever';

const TIER_TITLES: Readonly<Record<BisRole, string>> = {
  dps: `${GAME} Tier List (DPS, Tank, Healer)`,
  tank: `${GAME} Tank Tier List`,
  healer: `${GAME} Healer Tier List`,
};

const TIER_DESCRIPTIONS: Readonly<Record<BisRole, string>> = {
  dps: `Every DPS spec in ${GAME} ranked at level 60 by the simulator, with each spec's gap to the top, its guide and its best in slot. Tank and healer lists too.`,
  tank: `Every tank spec in ${GAME} ranked at level 60 by the simulator, with each spec's gap to the top, its guide and its best in slot.`,
  healer: `Every healer spec in ${GAME} ranked at level 60 by the simulator, with each spec's gap to the top, its guide and its best in slot.`,
};

export const searchCopy = {
  sim: {
    title: `${GAME} DPS Simulator: Sim Your Character Free`,
    /** The tool's own name, for its WebApplication node. */
    appName: 'Forever Sixty Simulator',
    description: `Sim your ${GAME_LONG} character free, in your browser: your DPS and why, what each stat is worth, the best gear and talents for you, from a simulator built on Forever's own class changes.`,
  },
  planner: {
    title: `${GAME} Talent Calculator`,
    appName: 'Forever Sixty Talent Calculator',
    description: `The ${GAME_LONG} talent calculator for every class: spend points in order, see Forever's changed talents, and share the build with a short link.`,
  },
  tiers: {
    title: (role: BisRole): string => TIER_TITLES[role],
    description: (role: BisRole): string => TIER_DESCRIPTIONS[role],
  },
  bisIndex: {
    title: `${GAME} BiS Lists for Every Class and Spec`,
    description: `Best in slot gear for every class and spec in ${GAME_LONG}, from leveling 20 to 60 to level 60, ranked by the simulator, with where each item comes from.`,
  },
  guidesIndex: {
    title: `${GAME} Class Guides for Every Spec`,
    description: `Class and spec guides for ${GAME_LONG}: talents for leveling and raids, rotation, stat priority and best in slot for every spec.`,
  },
  /** "Beast Mastery Hunter Guide, WoW Forever". */
  specGuideTitle: (fullName: string): string => `${fullName} Guide, ${GAME}`,
  /** "Hunter Class Guide, WoW Forever". */
  classGuideTitle: (className: string): string => `${className} Class Guide, ${GAME}`,
  /** The breadcrumb's name for a spec guide: "Beast Mastery Hunter guide". */
  specGuideCrumb: (fullName: string): string => `${fullName} guide`,
} as const;
