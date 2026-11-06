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

const CLASS_SLUGS: ReadonlySet<string> = new Set<ClassSlug>([
  'warrior',
  'paladin',
  'hunter',
  'rogue',
  'priest',
  'shaman',
  'mage',
  'warlock',
  'druid',
]);

export function isClassSlug(value: string): value is ClassSlug {
  return CLASS_SLUGS.has(value);
}

/** A row or title in a character's class colour; an unknown class reads the strong colour. */
export function classTone(classSlug: string): TitleTone {
  const slug = classSlug.toLowerCase();
  return isClassSlug(slug) ? { classSlug: slug } : 'strong';
}

/** The CSS colour a title tone paints with, always a token. */
export function titleColor(tone: TitleTone): string {
  if (tone === 'strong') return 'var(--color-strong)';
  if ('classSlug' in tone) return `var(--color-class-${tone.classSlug})`;
  return `var(--color-${tone.faction})`;
}

/** The outcome tones a Figure, Meter or Notice reads. */
export type FigureTone = 'strong' | 'gold' | 'kill' | 'wipe' | 'muted';

export const TONE_TEXT: Readonly<Record<FigureTone, string>> = {
  strong: 'text-strong',
  gold: 'text-gold',
  kill: 'text-kill',
  wipe: 'text-wipe',
  muted: 'text-muted',
};

/** The Page primitive's four compositions (spec 3.7). */
export type PageLayout = 'ledger' | 'split' | 'reading' | 'overview';

/** One choice in a ControlBar group (spec 3.4). */
export interface ControlOption {
  value: string;
  label: string;
  /** A count beside the label ("5"), muted. */
  count?: number;
  /** An icon URL (a spec icon at 22 px). */
  icon?: string;
  /** A level with no data yet: shown, not selectable, with a reason in its title. */
  disabled?: boolean;
  disabledReason?: string;
  /** The option navigates (a path segment such as the spec) instead of writing a query param. */
  href?: string;
}

export interface ControlGroup {
  /** The group label: 12 px caps from 1280, and above the control when the phone bar is open. */
  label: string;
  /** The query parameter it reads and writes (`band`, `faction`, `race`, `view`). */
  param: string;
  kind: 'segmented' | 'select';
  options: readonly ControlOption[];
  value: string;
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

export function rowElement(href: string | undefined): 'a' | 'div' {
  return href === undefined || href === '' ? 'div' : 'a';
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
  caption?: string;
  tone?: FigureTone;
}
