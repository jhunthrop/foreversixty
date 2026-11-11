// web/src/lib/ui/primitives.ts
// The flat system's shared types and pure rules (web layout spec 2026-11-05 section 3A): the
// props the ui/ primitives take and the colour rules they apply, so a page passes a class, a
// faction or an item quality and never a colour (3A.2 "applied by the primitive from a prop").
// Item quality colours stay in lib/planner/items.ts (rarityClassFor, rarityBorderColorFor).

/** A playable class's slug, as the --color-class-* tokens name it. */
export type ClassSlug =
  'warrior' | 'paladin' | 'hunter' | 'rogue' | 'priest' | 'shaman' | 'mage' | 'warlock' | 'druid';

export type FactionSlug = 'alliance' | 'horde';

/**
 * Whose colour a page title takes (spec 3.3, owner decision 8): a class or spec page its class's,
 * a guild its faction's, any other page the strong text colour.
 */
export type TitleTone = { classSlug: ClassSlug } | { faction: FactionSlug } | 'strong';

/** The nine playable classes in the game's own order: the one class list (component audit
 *  cluster 26). Names come from lib/sim/spec-label.ts `classDisplayName`, colours from the
 *  --color-class-* tokens. */
export const CLASS_SLUGS: readonly ClassSlug[] = [
  'warrior',
  'paladin',
  'hunter',
  'rogue',
  'priest',
  'shaman',
  'mage',
  'warlock',
  'druid',
];

const CLASS_SLUG_SET: ReadonlySet<string> = new Set<string>(CLASS_SLUGS);

export function isClassSlug(value: string): value is ClassSlug {
  return CLASS_SLUG_SET.has(value);
}

/** A class's colour as its token (`var(--color-class-mage)`); an unknown slug reads the strong
 *  colour. The tokens are the one source: data/classes.json's `color` is for the Discord and
 *  mock tooling only. */
export function classColor(classSlug: string): string {
  return isClassSlug(classSlug) ? `var(--color-class-${classSlug})` : 'var(--color-strong)';
}

/** A row or title in a character's class colour; an unknown class reads the strong colour. */
export function classTone(classSlug: string): TitleTone {
  const slug = classSlug.toLowerCase();
  return isClassSlug(slug) ? { classSlug: slug } : 'strong';
}

/** The art behind a title band (ui/PageHeader, Base.astro): the painted blue band, or home's hero. */
export type BandArt = 'band' | 'hero';

/** The CSS colour a title tone paints with, always a token: a heading's cream when it has no
 *  class or faction (round 5). */
export function titleColor(tone: TitleTone): string {
  if (tone === 'strong') return 'var(--color-cream)';
  if ('classSlug' in tone) return `var(--color-class-${tone.classSlug})`;
  return `var(--color-${tone.faction})`;
}

/** The outcome tones a Figure, Meter or Notice reads. */
export type FigureTone = 'strong' | 'gold' | 'kill' | 'wipe' | 'muted' | 'gain';

export const TONE_TEXT: Readonly<Record<FigureTone, string>> = {
  strong: 'text-strong',
  gold: 'text-gold',
  kill: 'text-kill',
  wipe: 'text-wipe',
  muted: 'text-muted',
  /** A DPS gain on a list row (round 5 `--color-gain`). */
  gain: 'text-gain',
};

/** The Page primitive's compositions (spec 3.7; `form`: setup spec 2026-11-07 section 3A;
 *  `workbench`: planner spec 2026-11-08 section 3A). */
export type PageLayout = 'ledger' | 'split' | 'reading' | 'overview' | 'form' | 'workbench';

/** One choice in a ControlBar group (spec 3.4). */
export interface ControlOption {
  value: string;
  label: string;
  /** A count beside the label ("5"), muted. */
  count?: number;
  /** An icon URL (a spec icon at 22 px). */
  icon?: string;
  /** A select option's second line in the open list (the sim's fight styles), muted. */
  description?: string;
  /** A level with no data yet: shown, not selectable, with a reason in its title. */
  disabled?: boolean;
  disabledReason?: string;
  /** The option navigates (a path segment such as the spec) instead of writing a query param. */
  href?: string;
  /** `href` is the option's own URL (a BiS band page), but the island applies the choice in place
   *  and moves the address itself: the link is for crawlers, a new tab and a script-less visit. */
  inPlace?: boolean;
  /** A class name or slug: the option leads with the ringed class crest (ruling P10), not `icon`. */
  crest?: string;
}

export interface ControlGroup {
  /** The group label: 12 px caps from 1280, and above the control when the phone bar is open. */
  label: string;
  /** The query parameter it reads and writes (`band`, `faction`, `race`, `view`). */
  param: string;
  kind: 'segmented' | 'select';
  options: readonly ControlOption[];
  value: string;
  /** The value the URL leaves out: choosing it removes the parameter (`data-default`). */
  defaultValue?: string;
  /** A select's width from 768 px, in px, instead of its longest label: one long, rare option
   *  (a race name) truncates in the field, never in its list, so the bar keeps to one row. */
  fieldWidth?: number;
  /** A segmented group of icons (the class picker, ruling P10): the labels show from 1280 and in
   *  the open phone stack, else each radio keeps its label as its accessible name. */
  compact?: boolean;
}

/** The phone bar's one line: each group's chosen label, joined with a middle dot. */
export function controlSummary(groups: readonly ControlGroup[]): string {
  return groups
    .map((group) => group.options.find((option) => option.value === group.value)?.label)
    .filter((label): label is string => label !== undefined && label !== '')
    .join(' · ');
}

/** A Row's link xor action (3A.2 rule): a row that navigates never holds a button. */
export type RowTarget = { href: string } | { action: true } | null;

export function rowElement(href: string | undefined, item = false): 'a' | 'li' | 'div' {
  if (href !== undefined && href !== '') return 'a';
  return item ? 'li' : 'div';
}

/** A signed figure the way the site prints a gain: `+1.8`, `−1.8` (true minus), `0.0`. */
export function signedFigure(value: number, digits = 1): string {
  const rounded = Number(value.toFixed(digits));
  if (rounded === 0) return (0).toFixed(digits);
  const text = Math.abs(rounded).toFixed(digits);
  return rounded > 0 ? `+${text}` : `−${text}`;
}

/** The flat system's link (spec 3.5): gold 600, no underline, `›` after it where it navigates. */
export const LINK =
  'text-gold font-semibold hover:text-gold-hover active:text-gold-hover rounded-field ' +
  'focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-gold';

/** A Row wrapped in a list item: the Row's own hairline is lost (each Row is its item's last
 *  child), so the item carries it, none under the last. */
export const ROW_ITEM = 'border-line-soft border-b last:border-b-0';

/** The Row's hairline over a part of a panel that is not a Row (a footnote line, a second list). */
export const RULE_TOP = 'border-line-soft border-t';

/** The Row's hairline under a part of a panel that is not a Row (an opened night's detail). */
export const RULE_BOTTOM = 'border-line-soft border-b';

/** The few fields an item row or icon needs; every surface's item maps to this. */
export interface ItemRef {
  id: number;
  name: string;
  quality: number;
  /** The icon file stem; absent, the dark socket. */
  icon?: string;
}

/** A gain or lead shown at a row's right: the figure, its unit line and its tone. */
export interface GainFigure {
  value: string;
  /** The unit after the figure (`DPS`). */
  unit?: string;
  caption?: string;
  /** What the figure is measured against, under it: `vs <item>`, `next: −1.7 vs <item>`. One line
   *  that may ellipsize (its full words in the title); the item carries its tooltip hook. */
  basis?: GainBasis;
  tone?: FigureTone;
}

export interface GainBasis {
  prefix: string;
  item?: ItemRef;
}

/** A basis line's full words, for its title when the column ellipsizes it. */
/** The middle dot every source line joins its facts with. */
export const FACT_SEPARATOR = ' · ';

/** A source line's facts, each kept whole on screen (coordinator ruling, 2026-11-09): every fact
 *  but the last carries its trailing dot, so the line wraps only between facts and "100 Tarnished
 *  Undermine Reals" or "The Barrens" never splits. Joined with single spaces the parts read back
 *  as the line itself. */
export function factSegments(line: string): string[] {
  const facts = line.split(FACT_SEPARATOR);
  return facts.map((fact, index) => (index < facts.length - 1 ? `${fact} ·` : fact));
}

export function basisText(basis: GainBasis): string {
  return basis.item ? `${basis.prefix} ${basis.item.name}` : basis.prefix;
}

/** The Plate's surface (ui/Plate.svelte), for a static skeleton that draws the page's plate before
 *  its island mounts. The only place outside the component that may name it. */
/** Round 5's panel: the card brown-black, a black hairline inside a warm outer ring, the 10 px
 *  radius (ui/panel.css draws the ring and the framed and lead variants). */
export const PLATE_SURFACE = 'ui-plate ui-surface min-w-0';

/** A panel's emphasis (round 5: two or three featured panels a page, flourishes on the lead only):
 *  `plain` the rounded panel, `framed` the cream inset line, `lead` the line and corner flourishes. */
export type PanelFrame = 'plain' | 'framed' | 'lead';

/** The class a frame adds to a surface. */
export function frameClass(frame: PanelFrame): string {
  if (frame === 'lead') return 'ui-framed ui-lead';
  return frame === 'framed' ? 'ui-framed' : '';
}
