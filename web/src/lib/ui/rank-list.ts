// web/src/lib/ui/rank-list.ts
// The ranked list's data shapes and its one grid (tiers spec 2026-11-06 §4.4, owner decision T6):
// RankList, ColumnHeads, RankGroup and RankRow share one column template, so the heads and every
// row line up. Written for the tier list first; the log rankings and the guild Top guilds read it next.

/** A figure in a row: the sorted one first, then any extra columns (a tank's health and threat). */
export interface RankFigure {
  /** Final text: "1,063.2". */
  value: string;
  /** "DPS", "taken / s": 12 px muted after the value; left out where the column head names it. */
  unit?: string;
  /** The column's name, read by assistive tech and shown beside an extra figure on a phone. */
  label: string;
  /** The best of its column: green, with `bestTitle` so colour is not the only signal. */
  best?: boolean;
  bestTitle?: string;
  /** The figure opens something of its own (a kill's execution opens compare mode on that fight). */
  href?: string;
  /** The figure's own explanation, on hover. */
  title?: string;
  /** The link's accessible name where the value alone does not say where it goes. */
  ariaLabel?: string;
  /** No visible label on a phone: the figure reads alone (a percentile in its parse colour). */
  bare?: boolean;
  /** The value's own colour (a percentile on the parse scale); the strong text without. */
  color?: string;
  /** Drawn below 768 only: a fact the wide row has a column for (the tier list's vs #1). */
  phoneOnly?: boolean;
  testid?: string;
}

/** One next step in a row's link strip. The first `primary` link is gold, the rest nav. */
export interface RankLink {
  label: string;
  href: string;
  primary?: boolean;
  ariaLabel?: string;
  title?: string;
  testid?: string;
}

/** The gap column: "−8.2%", or "Top" on the leader (`leader` sets it strong). */
export interface RankGap {
  text: string;
  /** The column's name for assistive tech ("Gap to the top"). */
  label: string;
  leader?: boolean;
  /** The gap in words, on hover ("Takes 53% more damage a second than Protection Warrior"). */
  title?: string;
}

/** A short line under the bar or the name, with the reason in its title ("≈ tie with Arms"). */
export interface RankNote {
  text: string;
  title?: string;
}

/** A column head. Numeric heads align to the end, over their figures. */
export interface RankHead {
  label: string;
  align?: 'start' | 'end';
}

export interface RankListShape {
  /** How many figure columns each row has (1 to 4). */
  figures: number;
  /** A figure column's width in px (112 by default; the log rankings' four columns take 96). */
  figureWidth?: number;
  /** How many links each row ends with (0 to 4). */
  links: number;
}

const MAX_FIGURES = 4;
const MAX_LINKS = 4;
/** Each link column width fits that many links at 600 14 px, 16 px apart. */
const LINK_COLUMN_PX: readonly number[] = [0, 56, 106, 156, 206];
const FIGURE_COLUMN_PX = 112;

/** The list's grid variables, as an inline style. Fails fast on a shape the grid cannot hold. */
export function rankListStyle(shape: RankListShape): string {
  const { figures, links, figureWidth = FIGURE_COLUMN_PX } = shape;
  if (!Number.isInteger(figures) || figures < 1 || figures > MAX_FIGURES) {
    throw new Error(`RankList: figures must be 1 to ${MAX_FIGURES}, got ${figures}`);
  }
  if (!Number.isInteger(links) || links < 0 || links > MAX_LINKS) {
    throw new Error(`RankList: links must be 0 to ${MAX_LINKS}, got ${links}`);
  }
  return `--rl-n: ${figures}; --rl-fig: ${figureWidth}px; --rl-links: ${LINK_COLUMN_PX[links]}px`;
}
