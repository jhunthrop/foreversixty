// web/src/lib/guild/readiness-view.ts
// Guild control-centre spec §4.E: the raid leader's pre-pull board. Every cell's text and
// the worst-first sort are pure functions over `GuildReadinessRow` (api.ts, the contract's
// own shape) so GuildReadiness.svelte never re-derives a number the API already sent.
import type { GuildReadinessRow } from './api';
import { SLOT_LABELS, type Slot } from '../planner/types';
import { counted } from '../plural';

/** The level a character raids at: every readiness check is measured for it. */
export const RAID_LEVEL = 60;

/** Spec §4.E: `readinessScore = failCount * 100 + gearGainDps`, so a raider failing more
 *  checks always sorts above one failing fewer, ties broken by the bigger gear gain. The
 *  API already sends `failing`; this never recomputes it from the row's own fields. */
export function readinessScore(row: GuildReadinessRow): number {
  return row.failing * 100 + (row.gear_gap?.gain_dps ?? 0);
}

/** Worst-first, a signed character with no export after that, the leveling characters last, highest
 *  level first then by name (week-one spec §3.5) -- never mutates `rows`. */
export function sortReadinessWorstFirst(rows: readonly GuildReadinessRow[]): GuildReadinessRow[] {
  return [...rows].sort((a, b) => {
    const leveling = Number(isLeveling(a)) - Number(isLeveling(b));
    if (leveling !== 0) return leveling;
    if (isLeveling(a)) return (b.level ?? 0) - (a.level ?? 0) || a.name.localeCompare(b.name);
    return Number(a.no_export) - Number(b.no_export) || readinessScore(b) - readinessScore(a);
  });
}

/** The viewer's own row pins first (member and officer, spec §4.E) -- applied after the
 *  worst-first sort, never instead of it, so the rest of the table stays worst-first. */
export function pinReadinessOwnRow(
  rows: readonly GuildReadinessRow[],
  myCharacterKey: string | null,
): GuildReadinessRow[] {
  if (myCharacterKey === null) return [...rows];
  return [...rows].sort(
    (a, b) => Number(b.character_key === myCharacterKey) - Number(a.character_key === myCharacterKey),
  );
}

/** A character under 60 is not graded: level-60 raid needs are not a level 52's to chase. */
export function isLeveling(row: Pick<GuildReadinessRow, 'level'>): boolean {
  return row.level !== null && row.level !== undefined && row.level < RAID_LEVEL;
}

/** `Level 52 · 8 levels to 60`, the whole of a leveling character's row. */
export function levelingLabel(level: number): string {
  return `Level ${level} · ${counted(RAID_LEVEL - level, 'level')} to ${RAID_LEVEL}`;
}

export const READINESS_EMPTY = 'No raiders on the board yet. Approve members on the Roster tab.';

/** The board's two groups: the graded rows, then the leveling ones under their rule. */
export function readinessGroups(rows: readonly GuildReadinessRow[]): {
  graded: GuildReadinessRow[];
  leveling: GuildReadinessRow[];
} {
  return {
    graded: rows.filter((row) => !isLeveling(row)),
    leveling: rows.filter((row) => isLeveling(row)),
  };
}

/** Whether the board read anything at all for this row: a row with nothing shared passes nothing. */
function anyCheckRead(row: GuildReadinessRow): boolean {
  return (
    row.gear_gap !== null ||
    row.enchants.checked ||
    row.consumables.state !== 'unknown' ||
    row.talent_points_read === true
  );
}

/** A graded row with something read and nothing failing: the viewer's `Ready`, the needs line's clear. */
export function passesEveryCheck(row: GuildReadinessRow): boolean {
  return !isLeveling(row) && !row.no_export && row.failing === 0 && anyCheckRead(row);
}

/** Humanises an enchant slot key ("main_hand" -> "Main Hand") via the shared slot-label
 *  table every other item surface already uses (web/src/lib/planner/types.ts's
 *  `SLOT_LABELS`) -- never a second, drifting slot-name table for this one cell. Falls
 *  back to a bare capitalised slug for a key `SLOT_LABELS` does not carry, the same
 *  fail-open leniency `capitalize` alone already gave an unrecognised value. */
/** A slot key from the API (`main_hand`, `finger_1` or `finger1`) in the client's words. */
export function enchantSlotLabel(raw: string): string {
  const slot = raw.replace(/^(finger|trinket)_(\d)$/, '$1$2');
  const label = Object.prototype.hasOwnProperty.call(SLOT_LABELS, slot) ? SLOT_LABELS[slot as Slot] : slot;
  // The client's own tooltip casing ("Main Hand"), the words the loot cards and the bot use.
  return label
    .split(/[\s_]+/)
    .map((word) => capitalize(word))
    .join(' ');
}

/**
 * The checks this one character fails right now, worst-first phrasing matching spec
 * §4.A.2/§6's own template clauses: the needs line reads it, so the line and the board can never
 * disagree, and it builds a Nudge message when the API has not sent `nudge_text`. A gain prints
 * only for a damage spec's sim-checked upgrades.
 */
export function readinessFails(row: GuildReadinessRow): string[] {
  const fails: string[] = [];
  const gap = row.gear_gap;
  if (gap !== null && gap.upgrades !== null && gap.upgrades > 0) {
    const gain = gap.unit === 'dps' ? Math.round(gap.gain_dps ?? 0) : 0;
    const upgrades = `${gap.upgrades} gear upgrade${gap.upgrades === 1 ? '' : 's'} waiting`;
    fails.push(gain > 0 ? `${upgrades} (+${gain} DPS)` : upgrades);
  }
  if (row.enchants.checked && row.enchants.missing_slots.length > 0) {
    fails.push(`no enchant: ${row.enchants.missing_slots.map(enchantSlotLabel).join(', ')}`);
  }
  if (row.consumables.state === 'short') fails.push('bags short on consumables');
  if (row.talent_points_unspent > 0) {
    fails.push(
      `${row.talent_points_unspent} unspent talent point${row.talent_points_unspent > 1 ? 's' : ''}`,
    );
  }
  return fails;
}

/**
 * Spec §4.E's Nudge button: copy-to-clipboard only, "there is no network path from the
 * site to the addon." Prefers the API's own officer-only `nudge_text` (the server's exact
 * wording); falls back to a short message built from `readinessFails` so the button still
 * works against a fixture or an API response that omits it.
 */
export function nudgeText(row: GuildReadinessRow): string {
  if (row.nudge_text !== undefined) return row.nudge_text;
  const fails = readinessFails(row);
  if (fails.length === 0) return `${row.name}: every readiness check passes.`;
  return `${row.name}: ${fails.join('; ')}.`;
}

function capitalize(value: string): string {
  return value.length === 0 ? value : value[0].toUpperCase() + value.slice(1);
}

/** The board by exception (batch 3): what a raider is short of, in a raid leader's words, worst
 *  first: gear, enchants, consumables, talents, then the raid's own asks (attunement, flask and
 *  food) and an export missing. Empty for a raider short of nothing. `raid` is the
 *  next raid's board header when the board is cut to it. */
export function readinessShortfalls(
  row: GuildReadinessRow,
  raid: { attunement?: boolean } | null = null,
): string[] {
  if (row.no_export) return [NO_EXPORT];
  const out: string[] = [];
  const upgrades = row.gear_gap?.upgrades ?? 0;
  if (upgrades > 0) out.push(counted(upgrades, 'upgrade'));
  const missing = row.enchants.checked ? row.enchants.missing_slots.length : 0;
  if (missing > 0) out.push(`${counted(missing, 'enchant')} missing`);
  // What the bags lack, named when the export says (a flask or raid elixir, buff food); short
  // without a name only when it does not.
  const flaskFood = row.flask_food;
  const lacks: string[] = [];
  if (flaskFood != null && !flaskFood.flask_or_elixir) lacks.push('no flask or raid elixir');
  if (flaskFood != null && !flaskFood.food) lacks.push('no buff food');
  if (lacks.length > 0) out.push(...lacks);
  else if (row.consumables.state === 'short') out.push('consumables short');
  if (row.talent_points_unspent > 0)
    out.push(`${counted(row.talent_points_unspent, 'talent point')} unspent`);
  if (raid?.attunement === true && row.attuned === false) out.push('not attuned');
  return out;
}

/** The board's three lists: who is short of something (worst first), who is ready, and who is still
 *  leveling. A raider the site could read nothing for is short of their export. */
export function readinessLists(
  rows: readonly GuildReadinessRow[],
  raid: { attunement?: boolean } | null = null,
): {
  short: { row: GuildReadinessRow; shortfalls: string[] }[];
  ready: GuildReadinessRow[];
  leveling: GuildReadinessRow[];
} {
  const { graded, leveling } = readinessGroups(sortReadinessWorstFirst(rows));
  const short: { row: GuildReadinessRow; shortfalls: string[] }[] = [];
  const ready: GuildReadinessRow[] = [];
  for (const row of graded) {
    const shortfalls = readinessShortfalls(row, raid);
    if (shortfalls.length > 0) short.push({ row, shortfalls });
    else if (anyCheckRead(row)) ready.push(row);
    else short.push({ row, shortfalls: [NOTHING_SHARED] });
  }
  return { short, ready, leveling };
}

/** A graded raider with nothing shared: Gear, bags and talents all unread. */
export const NOTHING_SHARED = 'nothing shared';
export const NO_EXPORT = 'no export yet';
/** The board's toggle: the next raid's In and Late sign-ups alone. */
export const NEXT_RAID_ONLY = 'Next raid only';
/** The ready panel's meta: what ready means. */
export const READY_MEANS = 'Best in slot or no upgrades, enchanted, stocked, every point spent';

/** A signed tank's fire resistance on the board (a raid that asks for it), or not shared. */
export function tankFireLine(fire: number | null | undefined): string {
  return typeof fire === 'number' ? `${fire} fire resistance` : 'Fire resistance not shared';
}
