// web/src/lib/og-pages.ts
// The share card of every static page. Base.astro points og:image at /og<path>.png for every
// page, so a static page missing here ships a link preview that 404s; og-pages.test.ts walks
// src/pages and fails on any static page without a card. Content pages, guides and BiS specs
// get theirs from their collections in src/pages/og/[...path].png.ts.
import type { OgContent } from './og';
import { scoreCopy } from './score/copy';
import { KIND_TITLES } from './sim/copy';
import { tiersCopy } from './tiers/tier-copy';

const SITE_KICKER = 'Forever Sixty';
const SIM_KICKER = 'Simulator';
const TIERS_KICKER = 'Tier list · Level 60';

/** Card path (the page path without its leading slash; the home page is 'index') to its card. */
export const STATIC_OG_PAGES: Readonly<Record<string, OgContent>> = {
  index: { title: 'Forever Sixty', kicker: 'World of Warcraft: Forever' },
  guides: { title: 'Guides by class', kicker: 'Leveling · talents · gear' },
  changelog: { title: 'What changed', kicker: 'Dated and sourced' },
  planner: { title: 'Build planner', kicker: 'Talents · order · share' },
  bis: { title: 'BiS', kicker: 'Best gear, levels 20 to 60' },
  'b-unavailable': { title: 'Build unavailable', kicker: SITE_KICKER },
  tiers: { title: tiersCopy.pageTitle('dps'), kicker: TIERS_KICKER },
  'tiers/tank': { title: tiersCopy.pageTitle('tank'), kicker: TIERS_KICKER },
  'tiers/healer': { title: tiersCopy.pageTitle('healer'), kicker: TIERS_KICKER },
  sim: { title: 'Simulator', kicker: 'Your character · your gear' },
  'sim/drops': { title: KIND_TITLES.drops, kicker: SIM_KICKER },
  'sim/gear': { title: KIND_TITLES.gear, kicker: SIM_KICKER },
  'sim/talents': { title: KIND_TITLES.talents, kicker: SIM_KICKER },
  'sim/weights': { title: KIND_TITLES.weights, kicker: SIM_KICKER },
  'sim/specs': { title: 'Spec support', kicker: SIM_KICKER },
  rankings: { title: 'Rankings', kicker: SITE_KICKER },
  ratings: { title: scoreCopy.explainerPageTitle, kicker: SITE_KICKER },
  premium: { title: 'Premium', kicker: SITE_KICKER },
  'premium/checkout': { title: 'Premium', kicker: SITE_KICKER },
  guild: { title: 'Guild tools', kicker: 'Sign-ups · attendance · Discord' },
  logs: { title: 'Logs and coaching', kicker: SITE_KICKER },
  'reports/sample': { title: 'Sample report', kicker: 'Logs and coaching' },
  character: { title: 'Character', kicker: SITE_KICKER },
  characters: { title: 'Look up a character', kicker: SITE_KICKER },
  setup: { title: 'Get set up', kicker: 'Addon · companion · apps' },
  connect: { title: 'Connect your companion', kicker: SITE_KICKER },
  account: { title: 'Your account', kicker: SITE_KICKER },
  login: { title: 'Sign in', kicker: SITE_KICKER },
};
