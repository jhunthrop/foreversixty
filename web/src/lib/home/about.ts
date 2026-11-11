// web/src/lib/home/about.ts
// Home's server-rendered words for a signed-out visitor (seo-fix 2026-11-11, lane G2): what the site
// does for a Forever player, every spec's BiS by role, the first raid tier and what the rankings
// hold, and the client build the data is read from. Search engines read about 90 words on home
// before this; the islands above it draw the rest only in a browser. Every count, name and date is
// read from the build's data (about-load.ts); the sentences keep only the stable why. Pure: the
// component renders the parts, the unit tests pin them.
import { FIRST_RAIDS } from '../first-raids';
import { SPECS, type SpecRole } from '../sim/specs';
import { utcDayMonth } from '../dates';

/** A run of text, a link when it has an href. */
export interface Part {
  text: string;
  href?: string;
  /** A spec's link is drawn in its class colour. */
  classSlug?: string;
}

export interface AboutFacts {
  /** The client build the data is read from, and the day the pipeline last read it (ISO). */
  build: string;
  readAt: string | null;
  /** The newest BiS ranking (ISO), or null when no spec has a file. */
  bisRankedAt: string | null;
  /** The specs with a BiS file, and the bands every file covers. */
  bisSpecs: ReadonlySet<string>;
  bands: { first: number; last: number } | null;
  /** The first tier's opening day (ISO date) while it is ahead; null once it is open. */
  raidsOpen: string | null;
  /** The most-logged ranked encounters, most kills first. */
  ranked: readonly { name: string; slug: string }[];
}

export interface AboutView {
  paragraphs: Part[][];
  roles: { role: SpecRole; label: string; specs: Part[] }[];
  bisMeta: string | undefined;
}

export const aboutCopy = {
  title: "What's here",
  bisTitle: 'Leveling BiS by spec',
  roleLabels: { tank: 'Tanks', healer: 'Healers', dps: 'Damage' } as const satisfies Record<SpecRole, string>,
} as const;

/** How many ranked encounters the rankings sentence names. */
export const RANKED_NAMED = 4;
const ROLE_ORDER: readonly SpecRole[] = ['tank', 'healer', 'dps'];

const link = (text: string, href: string): Part => ({ text, href });
const text = (value: string): Part => ({ text: value });

/** `a`, `a and b`, `a, b and c`, as parts. */
export function listParts(items: readonly Part[]): Part[] {
  return items.flatMap((item, index) => {
    if (index === 0) return [item];
    return [text(index === items.length - 1 ? ' and ' : ', '), item];
  });
}

function gearParagraph(facts: AboutFacts): Part[] {
  const bands =
    facts.bands === null ? '' : ` for each ten-level band from ${facts.bands.first} to ${facts.bands.last}`;
  return [
    text(`Every spec has `),
    link('leveling BiS', '/bis'),
    text(`${bands}, with the quest, dungeon or vendor each piece comes from, a `),
    link('guide', '/guides'),
    text(" written for Forever's changes, and talent builds you can open in the "),
    link('planner', '/planner'),
    text('.'),
  ];
}

function simParagraph(): Part[] {
  return [
    text('The '),
    link('simulator', '/sim'),
    text(' runs your character in the browser on the engine the BiS lists are ranked with: '),
    ...listParts([
      link('Top Gear', '/sim/gear'),
      link('stat weights', '/sim/weights'),
      link('Talent compare', '/sim/talents'),
      link('the Droptimizer', '/sim/drops'),
    ]),
    text('. The same engine orders every spec on the '),
    link('tier list', '/tiers'),
    text(' at 60, for '),
    ...listParts([
      link('damage', '/tiers'),
      link('healing', '/tiers/healer'),
      link('tanking', '/tiers/tank'),
    ]),
    text('. The '),
    link('addon and the companion', '/setup'),
    text(
      ' bring your character over with its gear and talents, and upload your combat logs as you play, so your boss kills reach the ',
    ),
    link('rankings', '/rankings'),
    text(' on their own.'),
  ];
}

function raidParagraph(facts: AboutFacts): Part[] {
  const raids = listParts(FIRST_RAIDS.map((raid) => text(raid.name)));
  const opening =
    facts.raidsOpen === null
      ? [text('The first raid tier is open: '), ...raids, text('.')]
      : [
          text(`The first raid tier opens on ${utcDayMonth(facts.raidsOpen.slice(0, 10))} with `),
          ...raids,
          text('.'),
        ];
  if (facts.ranked.length === 0) return opening;
  const bosses = facts.ranked.slice(0, RANKED_NAMED).map((boss) => link(boss.name, `/rankings/${boss.slug}`));
  return [...opening, text(' The most logged bosses on the rankings: '), ...listParts(bosses), text('.')];
}

function buildParagraph(facts: AboutFacts): Part[] {
  const read =
    facts.readAt === null ? '' : `, read from the game files on ${utcDayMonth(facts.readAt.slice(0, 10))}`;
  const ranked =
    facts.bisRankedAt === null ? '' : ` BiS lists ranked ${utcDayMonth(facts.bisRankedAt.slice(0, 10))}.`;
  return [text(`Data from client build ${facts.build}${read}.${ranked}`)];
}

/** Every spec with a BiS file, by role, each its full name linking to its BiS page. */
function bisRoles(facts: AboutFacts): AboutView['roles'] {
  return ROLE_ORDER.map((role) => ({
    role,
    label: aboutCopy.roleLabels[role],
    specs: SPECS.filter((spec) => spec.role === role && facts.bisSpecs.has(spec.spec)).map((spec) => ({
      ...link(spec.full_name, `/bis/${spec.class_slug}/${spec.spec_slug}`),
      classSlug: spec.class_slug,
    })),
  })).filter((group) => group.specs.length > 0);
}

export function aboutView(facts: AboutFacts): AboutView {
  return {
    paragraphs: [gearParagraph(facts), simParagraph(), raidParagraph(facts), buildParagraph(facts)],
    roles: bisRoles(facts),
    bisMeta: facts.bands === null ? undefined : `Levels ${facts.bands.first} to ${facts.bands.last}`,
  };
}
