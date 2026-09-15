// web/src/lib/report/format.ts
// Every number the report prints goes through here, so a duration reads the same in the
// fight selector, the chart axis and a death's timestamp. Numbers render in --font-mono
// with `tabular`, per design/DESIGN-SYSTEM.md.

/** Under a minute, tenths tell you more than a leading zero does. */
/** Like formatDuration but never drops the tenths: "1:01.4". For event lists. */
export function formatDurationPrecise(ms: number): string {
  const safe = Number.isFinite(ms) && ms > 0 ? ms : 0;
  if (safe < 60_000) return `${(safe / 1000).toFixed(1)}s`;
  const minutes = Math.floor(safe / 60_000);
  const seconds = (safe - minutes * 60_000) / 1000;
  return `${minutes}:${seconds.toFixed(1).padStart(4, '0')}`;
}

/** The game's spell school mask as a word; combined schools are joined. */
export function schoolName(mask: number | undefined): string {
  if (mask === undefined || mask <= 0) return '';
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

export function formatDuration(ms: number): string {
  const safe = Number.isFinite(ms) && ms > 0 ? ms : 0;
  if (safe < 60_000) return `${(safe / 1000).toFixed(1)}s`;
  const totalSeconds = Math.round(safe / 1000);
  const minutes = Math.floor(totalSeconds / 60);
  const seconds = totalSeconds % 60;
  return `${minutes}:${String(seconds).padStart(2, '0')}`;
}

/**
 * The engine writes UTC instants. The report shows them in UTC too rather than the
 * viewer's zone: a raid leader comparing a log against a teammate's screenshot needs the
 * same clock on both, and the log file itself is the shared reference.
 */
/** The calendar day of an instant, in UTC like formatClock: "26 Sep 2026". */
export function formatDate(iso: string): string {
  const at = new Date(iso);
  if (Number.isNaN(at.getTime())) return '';
  return new Intl.DateTimeFormat('en-GB', {
    day: 'numeric',
    month: 'short',
    year: 'numeric',
    timeZone: 'UTC',
  }).format(at);
}

export function formatClock(iso: string): string {
  const at = new Date(iso);
  if (Number.isNaN(at.getTime())) return '--:--:--';
  const pad = (n: number): string => String(n).padStart(2, '0');
  return `${pad(at.getUTCHours())}:${pad(at.getUTCMinutes())}:${pad(at.getUTCSeconds())}`;
}

const GROUPED = new Intl.NumberFormat('en-US');

export function formatAmount(n: number): string {
  const safe = Number.isFinite(n) ? n : 0;
  const abs = Math.abs(safe);
  if (abs >= 1_000_000_000) return `${trimZero(safe / 1_000_000_000)}B`;
  if (abs >= 1_000_000) return `${trimZero(safe / 1_000_000)}M`;
  // From 100k the grouped form is seven characters against a neighbour's "1.09M", and a
  // column mixing the two reads the longer string as the bigger number.
  if (abs >= 100_000) return `${(safe / 1000).toFixed(1).replace(/\.0$/, '')}k`;
  return GROUPED.format(Math.round(safe));
}

/** 1.25 stays 1.25, 12.50 becomes 12.5, 3.00 becomes 3. */
function trimZero(value: number): string {
  return value.toFixed(2).replace(/\.?0+$/, '');
}

/** A fight's outcome word for a list: Kill, Wipe, Live, or how much trash died. */
export function outcomeLabel(fight: {
  kind: string;
  kill: boolean;
  in_progress: boolean;
  npc_kills: number;
}): string {
  if (fight.in_progress) return 'Live';
  if (fight.kind !== 'encounter') return `${fight.npc_kills} killed`;
  return fight.kill ? 'Kill' : 'Wipe';
}

export function formatPerSecond(total: number, ms: number): string {
  if (!Number.isFinite(total) || !Number.isFinite(ms) || ms <= 0) return '0.0';
  return (total / (ms / 1000)).toFixed(1);
}

export function formatPercent(value: number): string {
  return `${(Number.isFinite(value) ? value : 0).toFixed(1)}%`;
}

/**
 * Class colours are tokens, never hex literals: design/DESIGN-SYSTEM.md owns the values
 * and src/styles/tokens.css holds them. A class the token set does not name -- the engine
 * can emit an inferred class from a future expansion -- gets the body colour rather than
 * a guess.
 */
const CLASS_TOKENS = new Set([
  'warrior',
  'paladin',
  'hunter',
  'rogue',
  'priest',
  'shaman',
  'mage',
  'warlock',
  'druid',
  'monk',
]);

export function classColorVar(className: string | undefined): string {
  const slug = (className ?? '').toLowerCase().replace(/\s+/g, '-');
  return CLASS_TOKENS.has(slug) ? `var(--color-class-${slug})` : 'var(--color-text)';
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

/**
 * Parse percentiles use the item-rarity scale. This is the one place the design system's
 * "never repurpose the game's colours" rule is deliberately bent, and it is bent because
 * every player already reads a purple parse as epic: reusing the scale is what makes the
 * number legible at a glance. The two AA-safe text variants are used for rare and epic,
 * as the rarity text classes do elsewhere.
 */
export function percentileToken(percentile: number): string {
  const p = Number.isFinite(percentile) ? percentile : 0;
  if (p >= 99) return 'var(--color-rarity-legendary)';
  if (p >= 95) return 'var(--color-rarity-epic-text)';
  if (p >= 75) return 'var(--color-rarity-rare-text)';
  if (p >= 50) return 'var(--color-rarity-uncommon)';
  if (p >= 25) return 'var(--color-rarity-common)';
  return 'var(--color-rarity-poor)';
}
