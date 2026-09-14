// web/src/lib/report/format.ts
// Every number the report prints goes through here, so a duration reads the same in the
// fight selector, the chart axis and a death's timestamp. Numbers render in --font-mono
// with `tabular`, per design/DESIGN-SYSTEM.md.

/** Under a minute, tenths tell you more than a leading zero does. */
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
  return GROUPED.format(Math.round(safe));
}

/** 1.25 stays 1.25, 12.50 becomes 12.5, 3.00 becomes 3. */
function trimZero(value: number): string {
  return value.toFixed(2).replace(/\.?0+$/, '');
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
  'warrior', 'paladin', 'hunter', 'rogue', 'priest', 'shaman', 'mage', 'warlock', 'druid',
]);

export function classColorVar(className: string | undefined): string {
  const slug = (className ?? '').toLowerCase().replace(/\s+/g, '-');
  return CLASS_TOKENS.has(slug) ? `var(--color-class-${slug})` : 'var(--color-text)';
}

/**
 * Parse percentiles use the item-rarity scale. This is the one place the design system's
 * "never repurpose the game's colours" rule is deliberately bent, and it is bent because
 * every player already reads a purple parse as epic: reusing the scale is what makes the
 * number legible at a glance. The two AA-safe text variants are used for rare and epic,
 * as the rarity text classes do elsewhere.
 */
/**
 * The window can leave a summary figure in one of three states (see window.ts's module
 * doc and scopeSummary): exact, scaled from the whole fight by the window's share (a
 * per-ability or per-target split, cast started/succeeded/failed, threat), or not
 * rescoped at all -- returned unchanged from the whole fight because the summary never
 * tracked it over time (interrupts, dispels, resource gained/spent/zero_ms, aura
 * max_stacks, roster activity_pct and active_ms). The two non-exact cases are different
 * lies to avoid, so they get different marks: `~` means "close, scaled proportionally",
 * `†` (a dagger) means "not this window's number at all, the whole fight's". Every
 * table that renders one of these figures uses these two functions rather than inlining
 * its own glyph, so the convention -- and its title text -- reads the same everywhere.
 */
export function approximateMark(scaled: boolean): string {
  return scaled ? '~' : '';
}

export function approximateTitle(scaled: boolean): string | undefined {
  return scaled
    ? 'Split across abilities and targets in proportion to the window, not measured directly in it.'
    : undefined;
}

export function wholeFightMark(stale: boolean): string {
  return stale ? '†' : '';
}

export function wholeFightTitle(stale: boolean): string | undefined {
  return stale ? 'The whole fight’s figure: the summary does not track this over time.' : undefined;
}

export function percentileToken(percentile: number): string {
  const p = Number.isFinite(percentile) ? percentile : 0;
  if (p >= 99) return 'var(--color-rarity-legendary)';
  if (p >= 95) return 'var(--color-rarity-epic-text)';
  if (p >= 75) return 'var(--color-rarity-rare-text)';
  if (p >= 50) return 'var(--color-rarity-uncommon)';
  if (p >= 25) return 'var(--color-rarity-common)';
  return 'var(--color-rarity-poor)';
}
