// web/src/lib/report/format.ts
// Every number the report prints goes through here, so a duration reads the same in the
// fight selector, the chart axis and a death's timestamp. Numbers render in the body face
// with tabular numerals (the `numeric` utility), per design/DESIGN-SYSTEM.md.

import { killCounts, killCountsLabel } from './kill-counts';
import type { FightEntry, Phase, Taunt } from './types';
import { compact, ordinal, percent, thousands } from '../format/number';
import { CLASS_SLUGS } from '../ui/primitives';
import { uiCopy } from '../ui/copy';

/**
 * An `<a>` is inline: its own box is only as tall as its text, not the `min-h-11` row it
 * sits in. Every row link across Character.svelte, Guild.svelte, Rankings.svelte and
 * ReportView.svelte needs this same fix, so it lives once here rather than as four copies
 * of the identical string.
 */
export const rowLink = 'inline-flex min-h-11 items-center';

/**
 * Everything that tells one taunt from another: two tanks can taunt two adds on the
 * same millisecond, and over a night the same instant recurs in every pull, which the
 * label separates. Shared by the taunt list and the timeline marks, which key the same
 * events off the same summary.
 */
export function tauntKey(taunt: Taunt): string {
  return `${taunt.label ?? ''}|${taunt.at_ms}|${taunt.source_guid}|${taunt.target_guid}|${taunt.spell_id}`;
}

/** The colour token for a school mask's first school; melee (no school) is physical. */
export function schoolToken(mask: number | undefined): string {
  const bits: [number, string][] = [
    [1, 'physical'],
    [2, 'holy'],
    [4, 'fire'],
    [8, 'nature'],
    [16, 'frost'],
    [32, 'shadow'],
    [64, 'arcane'],
  ];
  if (mask === undefined || mask <= 0) return 'var(--color-school-physical)';
  for (const [bit, name] of bits) if ((mask & bit) !== 0) return `var(--color-school-${name})`;
  return 'var(--color-school-physical)';
}

/** The game's spell school mask as a word; combined schools are joined. */
export function schoolName(mask: number | undefined): string {
  // A melee swing carries no school in the log; it is physical.
  if (mask === undefined || mask <= 0) return 'Physical';
  const names: string[] = [];
  const table: [number, string][] = [
    [1, 'Physical'],
    [2, 'Holy'],
    [4, 'Fire'],
    [8, 'Nature'],
    [16, 'Frost'],
    [32, 'Shadow'],
    [64, 'Arcane'],
  ];
  for (const [bit, name] of table) if ((mask & bit) !== 0) names.push(name);
  return names.join('/');
}

/**
 * An amount in the scale its column is read in: when the column's largest figure is
 * abbreviated (239.2k), every figure in it is (81k, not 80,993), or a five-digit row reads
 * bigger than the six-digit rows above it. Below that the grouped form stays.
 */
export function formatAmountLike(n: number, columnMax: number): string {
  const safe = Number.isFinite(n) ? n : 0;
  if (Math.abs(columnMax) >= 100_000 && Math.abs(safe) < 100_000) return thousands(safe);
  return compact(safe);
}

/**
 * The phase a fight ended in: the last one that started, since the engine writes them in
 * order and runs the last one to the fight's end. '' when the encounter has no phases.
 */
export function phaseReached(phases: readonly Phase[] | undefined): string {
  return phases === undefined || phases.length === 0 ? '' : phases[phases.length - 1].name;
}

/** Facts joined by " · ", leaving out the empty ones (a trash pull with no kill count). */
export function joinFacts(...parts: readonly string[]): string {
  return parts.filter((part) => part !== '').join(' · ');
}

/**
 * A fight's outcome for a list: Kill, Wipe, Live, or for trash how many mobs died and how many
 * the group killed. '' for trash indexed before engine 0.6.0, which has no true count.
 */
export function outcomeLabel(
  fight: Pick<FightEntry, 'kind' | 'kill' | 'in_progress' | 'mob_deaths' | 'group_kills' | 'boss_health_pct'>,
  phase = '',
): string {
  if (fight.in_progress) return 'Live';
  if (fight.kind !== 'encounter') {
    const counts = killCounts(fight);
    return counts === null ? '' : killCountsLabel(counts);
  }
  if (fight.kill) return 'Kill';
  // "Wipe 23%": how far the pull got, which is what separates one wipe from the next.
  // With phases curated for the boss, how far is also which phase it died in.
  const health = fight.boss_health_pct;
  const wipe = health !== undefined && health >= 0 ? `Wipe ${Math.round(health)}%` : 'Wipe';
  return phase === '' ? wipe : `${wipe} · in ${phase}`;
}

export function formatPerSecond(total: number, ms: number): string {
  if (!Number.isFinite(total) || !Number.isFinite(ms) || ms <= 0) return '0.0';
  return (total / (ms / 1000)).toFixed(1);
}

/**
 * A row's landed count. A sim row is a per-fight average rounded to a whole count, so an
 * action that dealt damage in a few fights out of a hundred (a Fire mage's wand) has a count
 * that rounds to 0 beside a real amount; that reads "<1", never "0".
 */
export function formatLanded(landed: number, amount: number): string {
  return landed === 0 && amount > 0 ? '<1' : String(landed);
}

/** A share of a total; a share above zero too small for one decimal reads "<0.1%", not "0.0%". */
export function formatShare(value: number): string {
  return value > 0 && value < 0.05 ? '<0.1%' : percent(value);
}

/**
 * Class colours are tokens, never hex literals: design/DESIGN-SYSTEM.md owns the values
 * and src/styles/tokens.css holds them. A class the token set does not name -- the engine
 * can emit an inferred class from a future expansion -- gets the body colour rather than
 * a guess.
 */
/** `Death Knight` -> `death-knight`: the one class name to slug rule (the API sends names). */
export function classSlugFromName(className: string): string {
  return className.toLowerCase().replace(/\s+/g, '-');
}

const CLASS_TOKENS: ReadonlySet<string> = new Set<string>([...CLASS_SLUGS, 'monk']);

/** A known class's slug (`death-knight`), else undefined. */
export function classSlug(className: string | undefined): string | undefined {
  const slug = classSlugFromName(className ?? '');
  return CLASS_TOKENS.has(slug) ? slug : undefined;
}

/** A known class's colour token name (`--color-class-mage`), for a dropdown option's colour. */
export function classColorToken(className: string | undefined): string | undefined {
  const slug = classSlug(className);
  return slug === undefined ? undefined : `--color-class-${slug}`;
}

export function classColorVar(className: string | undefined): string {
  const token = classColorToken(className);
  return token === undefined ? 'var(--color-text)' : `var(${token})`;
}

/**
 * A figure that is not what the window (window.ts's `scopeSummary`) or the active filters
 * (filters.ts's `applyActorFilters`) claim it is falls into one of two lies to avoid, and
 * every table marks them differently rather than inlining its own glyph:
 *
 *   `~` (scaled)      A per-ability or per-target split rescaled by a ratio -- the
 *                      window's share of the actor's total (window.ts), or a filter's
 *                      target share (filters.ts's `targetShare`). Close, proportionally,
 *                      not measured directly. Never applies to an actor's own `effective`
 *                      figure: window.ts computes that independently from the one-second
 *                      series, so it is exact regardless of the window, and untouched by
 *                      filters unless a target or boss filter is active (see
 *                      ReportView.svelte's `approximate` derivation).
 *   `†` (whole fight)  A figure the summary never rescopes at all under a window --
 *                      interrupts, dispels, resource gained/spent/zero_ms, aura
 *                      max_stacks, roster activity_pct and active_ms. Not "close", simply
 *                      the whole fight's number regardless of the window.
 *
 * Each mark has a title (for a sighted hover) and an aria-label composer (so the meaning
 * reaches assistive tech too: a `title` on an element that already has visible text is
 * not reliably exposed as an accessible name or description, and is unreachable on
 * touch). Compose the full text once -- `approximateAriaLabel(scaled, formatAmount(x))`
 * -- and set it as the element's `aria-label`, keeping `title` only as a bonus for mouse
 * hover. Task 12 should reuse all of these rather than inventing new ones.
 */
export function approximateMark(scaled: boolean): string {
  return scaled ? '~' : '';
}

export function approximateTitle(scaled: boolean): string | undefined {
  return scaled
    ? 'Split across abilities and targets in proportion to the window and any active filter, not measured directly.'
    : undefined;
}

export function approximateAriaLabel(scaled: boolean, text: string): string {
  return scaled ? `approximately ${text}` : text;
}

export function wholeFightMark(stale: boolean): string {
  return stale ? '†' : '';
}

export function wholeFightTitle(stale: boolean): string | undefined {
  return stale ? 'The whole fight’s figure: the summary does not track this over time.' : undefined;
}

export function wholeFightAriaLabel(stale: boolean, text: string): string {
  return stale ? `${text}, the whole fight’s figure` : text;
}

/** A Parse cell's fallback on a trash fight: not computed, which is not a failure. */
export const TRASH_PARSE = '—';

/** The hover text behind a Parse cell: the number's meaning, or why the cell is empty. */
export function parseTitle(value: number | string, ranked = 0): string {
  if (typeof value === 'number') {
    const among = ranked > 0 ? `${ranked} ranked ${ranked === 1 ? 'kill' : 'kills'}` : 'the ranked kills';
    const note = ranked === 1 ? '; the only kill of its kind ranked so far, so it stands first of one' : '';
    return `${ordinal(Math.floor(value))} percentile among ${among} of this boss by this spec on this ruleset${note}`;
  }
  if (value === TRASH_PARSE) return 'No parse for trash.';
  if (value === 'wipe') return 'A wipe is not ranked';
  if (value === '–') return 'Nothing of this spec has been ranked on this boss yet';
  if (value === '?')
    return `The rankings could not be reached, so this parse is unknown; the figure beside it is right. ${uiCopy.retry} from the line above the table`;
  if (value === 'none') return 'Damage taken has no parse';
  if (value === '…') return 'Asking the rankings where this sits';
  if (value === 'window') return 'A parse is for the whole fight; clear the window to read it';
  if (value === 'filter')
    return 'A parse is for the whole table; clear the target, boss or ability filter to read it';
  if (value === 'night')
    return 'A parse is per kill; the night as a whole has none. Pick a pull for its parses';
  return '';
}

/**
 * Parse percentiles use the item-rarity scale. This is the one place the design system's
 * "never repurpose the game's colours" rule is deliberately bent, and it is bent because
 * every player already reads a purple parse as epic: reusing the scale is what makes the
 * number legible at a glance.
 *
 * The ladder every log reader already knows: grey under 25, green from 25, blue from 50, purple
 * from 75, orange from 95, pink at 99, gold at 100. Floored, never rounded, for the colour and
 * the number alike (Warcraft Logs' own rule, owner ruling 2026-11-10): 98.6 reads "98" in
 * orange, 99.53 is pink, and only a whole 100 is gold.
 */
export type ParseTier = 'grey' | 'green' | 'blue' | 'purple' | 'orange' | 'pink' | 'gold';

const PARSE_LADDER: readonly (readonly [number, ParseTier])[] = [
  [100, 'gold'],
  [99, 'pink'],
  [95, 'orange'],
  [75, 'purple'],
  [50, 'blue'],
  [25, 'green'],
];

const PARSE_TEXT_TOKEN: Readonly<Record<ParseTier, string>> = {
  grey: 'var(--color-parse-grey-text)',
  green: 'var(--color-parse-green)',
  blue: 'var(--color-rarity-rare-text)',
  purple: 'var(--color-rarity-epic-text)',
  orange: 'var(--color-parse-orange)',
  pink: 'var(--color-parse-pink)',
  gold: 'var(--color-parse-gold)',
};

function flooredParse(percentile: number): number {
  return Number.isFinite(percentile) ? Math.floor(percentile) : 0;
}

/** The ladder step a percentile sits on. */
export function parseTier(percentile: number): ParseTier {
  const p = flooredParse(percentile);
  return PARSE_LADDER.find(([floor]) => p >= floor)?.[1] ?? 'grey';
}

/** The figure a parse cell prints: the floored whole percentile, "98". */
export function parseText(percentile: number): string {
  return String(flooredParse(percentile));
}

/** A parse read aloud: "98th percentile". */
export function parseLabel(percentile: number): string {
  return `${ordinal(flooredParse(percentile))} percentile`;
}

/** The colour a parse figure is printed in. */
export function percentileToken(percentile: number): string {
  // Every caller colours text, so the three steps whose raw colour fails WCAG AA as small text
  // read their text variants: the grey's own, and the rare and epic text colours the rarity
  // classes use (tokens.css), the same blue and purple a player reads.
  return PARSE_TEXT_TOKEN[parseTier(percentile)];
}

/** Everything a parse cell needs: the floored figure, its colour and its screen-reader name. */
export interface ParseFigure {
  text: string;
  color: string;
  label: string;
}

export function parseFigure(percentile: number): ParseFigure {
  return { text: parseText(percentile), color: percentileToken(percentile), label: parseLabel(percentile) };
}
