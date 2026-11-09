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

export const LEVELING_SUFFIX = ` · checks start at ${RAID_LEVEL}`;
export const LEVELING_GROUP = 'LEVELING';
/** The own row's tag with nothing to fix, and the Overview's value for a clean raider. */
export const READINESS_READY = 'Ready';
/** The Overview's Readiness card with no one to grade yet. */
export const READINESS_STARTS_AT_60 = `Readiness starts when someone reaches ${RAID_LEVEL}.`;
export const READINESS_EMPTY = 'No raiders on the board yet. Approve members on the Roster tab.';
export const AVG_ITEM_LEVEL = 'Avg item level';
export const AVG_ITEM_LEVEL_TOOLTIP =
  'Your average item level, and how far it is from the middle of your level-60 raiders.';

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

export const NOT_SHARED = 'Not shared';
export const NOT_CHECKED_YET = 'Not checked yet';
export const NOT_READ_YET = 'Not read yet';
export const NO_UPGRADES = 'No upgrades';
export const NOT_SIMMED_YET = 'not simmed yet';
export const ALL_ENCHANTED = 'All enchanted';

/** A cell the character's consent keeps from the site reads "Not shared" (consent is roster
 *  only), and one the site could have read but has not yet reads "Not checked yet". The words
 *  name the real state per cell; "consent needed" told a player nothing about which. */
function withheldWord(row: GuildReadinessRow, sharedAt: 'gear' | 'gear_bags'): string {
  const shares = sharedAt === 'gear' ? row.consent !== 'roster' : row.consent === 'gear_bags';
  return shares ? NOT_CHECKED_YET : NOT_SHARED;
}

/** Spec §4.E's Gear gap cell. `gear`/`gear_bags` consent only -- `gear_gap === null` (the
 *  API's fail-closed shape for `roster` consent) reads "Not shared", never a leaked figure; a
 *  member who shares gear whose band has no entry yet reads "Not checked yet". The gain is
 *  printed only when a simulation produced it: upgrades with no simulated gain read
 *  "15 upgrades · not simmed yet", never "+0 DPS". A healer's or tank's gap has no unit and is the
 *  upgrade count alone, never a DPS figure (week-one spec §3.5). */
export function gearGapLabel(row: GuildReadinessRow): string {
  const gap = row.gear_gap;
  if (gap === null || gap.upgrades === null) return withheldWord(row, 'gear');
  if (gap.upgrades === 0) return NO_UPGRADES;
  const upgrades = counted(gap.upgrades, 'upgrade');
  if (gap.unit !== 'dps') return upgrades;
  const gain = Math.round(gap.gain_dps ?? 0);
  if (gain <= 0 && gap.not_sim_checked > 0) return `${upgrades} · ${NOT_SIMMED_YET}`;
  if (gain <= 0) return upgrades;
  if (gap.not_sim_checked > 0) return `${upgrades} · +${gain} DPS · ${gap.not_sim_checked} ${NOT_SIMMED_YET}`;
  return `${upgrades} · +${gain} DPS`;
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

/** Spec §4.E's Enchants cell, full text -- `checked: false` means no gear consent, read
 *  the same as Gear gap's own consent gate (the stricter `gear_bags` floor applies to
 *  Consumables only, enchants share Gear gap's `gear` floor). Every missing slot, never
 *  truncated: the Readiness board's Enchants pill reads it in full and wraps (guild loop
 *  round 3, web4, after `Main hand, Chest and 2 more slots` clipped at 1440).
 */
export function enchantLabel(row: GuildReadinessRow): string {
  if (!row.enchants.checked) return withheldWord(row, 'gear');
  if (row.enchants.missing_slots.length === 0) return ALL_ENCHANTED;
  return row.enchants.missing_slots.map(enchantSlotLabel).join(', ');
}

/** The slots past the named ones, in words: ` and 2 more slots`, never `+2`, which reads as an
 *  enchant's own `+2` (guild round 2). Empty at 0. */
export function moreSlots(rest: number): string {
  return rest > 0 ? ` and ${counted(rest, 'more slot', 'more slots')}` : '';
}

/** Spec §4.E's Consumables cell -- `gear_bags` consent only, stricter than Gear gap/
 *  Enchants' `gear` floor: a member who shares gear but not bags reads "Not shared" here even
 *  though the other two cells are read; one who shares bags the export did not carry reads
 *  "Not checked yet". */
export function consumablesLabel(row: GuildReadinessRow): string {
  if (row.consumables.state === 'unknown') return withheldWord(row, 'gear_bags');
  return row.consumables.state === 'stocked' ? 'Stocked' : 'Short';
}

/** Spec §4.E's Talent points cell. The API sends 0 both for "nothing unspent" and for "the
 *  export carried no talents", so it also says which (`talent_points_read`): a 0 it did not
 *  read is "Not read yet", never a 0 a level 60 would be ashamed of. A response from before
 *  that field reads "—" for 0, claiming nothing. */
export function talentPointsLabel(row: GuildReadinessRow): string {
  if (row.talent_points_unspent > 0) return `${row.talent_points_unspent} unspent`;
  if (row.talent_points_read === false) return NOT_READ_YET;
  return row.talent_points_read === true ? 'None unspent' : '—';
}

/** Spec §4.E's Item level cell: the character's average item level, `—` when it carries none and
 *  `Not shared` when the roster-only consent explains why. */
export function itemLevelLabel(row: GuildReadinessRow): string {
  if (row.item_level === null) return row.consent === 'roster' ? NOT_SHARED : '—';
  return String(row.item_level);
}

/** `1 above the raid median`, `2 below the raid median`, `at the raid median`; empty with no delta. */
export function medianCaption(row: GuildReadinessRow): string {
  const delta = row.item_level_delta;
  if (delta === null || delta === undefined) return '';
  if (delta === 0) return 'at the raid median';
  return `${Math.abs(delta)} ${delta > 0 ? 'above' : 'below'} the raid median`;
}

/** The phone chip: `Item level 60 · 1 below median`, or the cell's words with no figure. */
export function itemLevelChip(row: GuildReadinessRow): string {
  if (row.item_level === null) return itemLevelLabel(row);
  const caption = medianCaption(row).replace(' the raid median', ' median');
  return caption === '' ? `Item level ${row.item_level}` : `Item level ${row.item_level} · ${caption}`;
}

/** The slot as the client's Character panel writes it: `Main Hand`, `Chest`. */
function clientSlotName(slot: string): string {
  return enchantSlotLabel(slot)
    .split(' ')
    .map((word) => capitalize(word))
    .join(' ');
}

/** How many slots the Overview's enchant line names before `and {n} more slots`. */
const WORST_CHECK_SLOTS = 2;

/** The Overview's row: the worst failing check in words, and how many more fail (week-one spec
 *  §3.3). Null for a row failing nothing. */
export function worstCheck(row: GuildReadinessRow): { text: string; more: number } | null {
  if (row.failing <= 0) return null;
  const more = row.failing - 1;
  const missing = row.enchants.checked ? row.enchants.missing_slots : [];
  if (missing.length > 0) {
    const named = missing.slice(0, WORST_CHECK_SLOTS).map(clientSlotName).join(', ');
    const rest = missing.length - WORST_CHECK_SLOTS;
    return { text: `No enchant: ${named}${moreSlots(rest)}`, more };
  }
  if (row.talent_points_unspent > 0) {
    return { text: `${counted(row.talent_points_unspent, 'talent point')} unspent`, more };
  }
  if (row.consumables.state === 'short') return { text: 'Bags short on consumables', more };
  const upgrades = row.gear_gap?.upgrades ?? 0;
  return upgrades > 0 ? { text: `${counted(upgrades, 'upgrade')} waiting`, more } : null;
}

/** Which visual treatment each cell takes (GuildReadiness.svelte's own pill/colour
 *  choices) -- kept separate from the `*Label` text functions above so a row-level
 *  "consent needed" grey, "Short"/"missing" red, "Stocked"/"All enchanted" green and
 *  "N unspent" gold read consistently without the label text itself encoding colour. */
export type ReadinessCellKind = 'consent' | 'warn' | 'ok' | 'neutral';

export function gearGapKind(row: GuildReadinessRow): ReadinessCellKind {
  return row.gear_gap === null || row.gear_gap.upgrades === null ? 'consent' : 'neutral';
}

export function enchantKind(row: GuildReadinessRow): ReadinessCellKind {
  if (!row.enchants.checked) return 'consent';
  return row.enchants.missing_slots.length === 0 ? 'ok' : 'warn';
}

export function consumablesKind(row: GuildReadinessRow): ReadinessCellKind {
  if (row.consumables.state === 'unknown') return 'consent';
  return row.consumables.state === 'stocked' ? 'ok' : 'warn';
}

export function talentPointsKind(row: GuildReadinessRow): ReadinessCellKind {
  return row.talent_points_unspent > 0 ? 'warn' : 'neutral';
}

/** The muted treatment (no pill, nothing to flag) for a cell that says why it has no figure. */
export function isWithheldWord(label: string): boolean {
  return label === NOT_SHARED || label === NOT_CHECKED_YET || label === NOT_READ_YET;
}

/** `'ok'` at or above the roster median (spec: "green at or above"), `'warn'` below,
 *  `'consent'` when the character carries no item level at all. */
export function itemLevelKind(row: GuildReadinessRow): ReadinessCellKind {
  if (row.item_level === null) return 'consent';
  return (row.item_level_delta ?? 0) >= 0 ? 'ok' : 'warn';
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
/** The ready panel's meta: what ready means. */
export const READY_MEANS = 'Best in slot or no upgrades, enchanted, stocked, every point spent';

/** A signed tank's fire resistance on the board (a raid that asks for it), or not shared. */
export function tankFireLine(fire: number | null | undefined): string {
  return typeof fire === 'number' ? `${fire} fire resistance` : 'Fire resistance not shared';
}
