// web/src/lib/guild/readiness-view.ts
// Guild control-centre spec §4.E: the raid leader's pre-pull board. Every cell's text and
// the worst-first sort are pure functions over `GuildReadinessRow` (api.ts, the contract's
// own shape) so GuildReadiness.svelte never re-derives a number the API already sent.
import type { GuildReadinessRow } from './api';
import { SLOT_LABELS, type Slot } from '../planner/types';
import { counted } from '../plural';

/** Spec §4.E: `readinessScore = failCount * 100 + gearGainDps`, so a raider failing more
 *  checks always sorts above one failing fewer, ties broken by the bigger gear gain. The
 *  API already sends `failing`; this never recomputes it from the row's own fields. */
export function readinessScore(row: GuildReadinessRow): number {
  return row.failing * 100 + (row.gear_gap?.gain_dps ?? 0);
}

/** Worst-first, stable otherwise -- never mutates `rows`. */
export function sortReadinessWorstFirst(rows: readonly GuildReadinessRow[]): GuildReadinessRow[] {
  return [...rows].sort((a, b) => readinessScore(b) - readinessScore(a));
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
 *  "15 upgrades · not simmed yet", never "+0 DPS". */
export function gearGapLabel(row: GuildReadinessRow): string {
  const gap = row.gear_gap;
  if (gap === null || gap.upgrades === null) return withheldWord(row, 'gear');
  if (gap.upgrades === 0) return NO_UPGRADES;
  const upgrades = counted(gap.upgrades, 'upgrade');
  const gain = Math.round(gap.gain_dps ?? 0);
  if (gain <= 0 && gap.not_sim_checked > 0) return `${upgrades} · ${NOT_SIMMED_YET}`;
  if (gap.not_sim_checked > 0) return `${upgrades} · +${gain} DPS · ${gap.not_sim_checked} ${NOT_SIMMED_YET}`;
  return `${upgrades} · +${gain} DPS`;
}

/** Humanises an enchant slot key ("main_hand" -> "Main hand") via the shared slot-label
 *  table every other item surface already uses (web/src/lib/planner/types.ts's
 *  `SLOT_LABELS`) -- never a second, drifting slot-name table for this one cell. Falls
 *  back to a bare capitalised slug for a key `SLOT_LABELS` does not carry, the same
 *  fail-open leniency `capitalize` alone already gave an unrecognised value. */
function enchantSlotLabel(slot: string): string {
  return Object.prototype.hasOwnProperty.call(SLOT_LABELS, slot)
    ? SLOT_LABELS[slot as Slot]
    : capitalize(slot);
}

/** Spec §4.E's Enchants cell, full text -- `checked: false` means no gear consent, read
 *  the same as Gear gap's own consent gate (the stricter `gear_bags` floor applies to
 *  Consumables only, enchants share Gear gap's `gear` floor). Every missing slot, never
 *  truncated -- the Enchants cell's own tooltip (`title`) reads this; the pill itself
 *  reads `enchantPillText` below, which truncates a long list the pill cannot hold.
 */
export function enchantLabel(row: GuildReadinessRow): string {
  if (!row.enchants.checked) return withheldWord(row, 'gear');
  if (row.enchants.missing_slots.length === 0) return ALL_ENCHANTED;
  return row.enchants.missing_slots.map(enchantSlotLabel).join(', ');
}

/** The Enchants pill can hold at most this many slot names before it reads as a wall of
 *  text ("MAIN_HAND, CHEST, BACK, FEET" overflowing the column, the defect this guards
 *  against) -- past this, the pill reads "<first two> +N more" and the cell's own
 *  `title` attribute (`enchantLabel`, above) carries the full list. */
const ENCHANT_PILL_MAX_SLOTS = 2;

/** The Enchants pill's own visible text: `enchantLabel`'s full list, truncated to
 *  `ENCHANT_PILL_MAX_SLOTS` slots with a "+N more" suffix once a row fails that many
 *  enchant checks at once. */
export function enchantPillText(row: GuildReadinessRow): string {
  if (!row.enchants.checked) return withheldWord(row, 'gear');
  if (row.enchants.missing_slots.length === 0) return ALL_ENCHANTED;
  const labels = row.enchants.missing_slots.map(enchantSlotLabel);
  if (labels.length <= ENCHANT_PILL_MAX_SLOTS) return labels.join(', ');
  const shown = labels.slice(0, ENCHANT_PILL_MAX_SLOTS).join(', ');
  return `${shown} +${labels.length - ENCHANT_PILL_MAX_SLOTS} more`;
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

/** Spec §4.E's Item level cell: signed delta against the roster's own median, `—` when the
 *  character carries no item level at all and the roster-only consent explains why. */
export function itemLevelLabel(row: GuildReadinessRow): string {
  if (row.item_level === null) return row.consent === 'roster' ? NOT_SHARED : '—';
  const delta = row.item_level_delta ?? 0;
  const sign = delta >= 0 ? '+' : '';
  return `${row.item_level} (${sign}${delta})`;
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
 * §4.A.2/§6's own template clauses -- the Readiness tab's own long-form fail list, used to
 * build a Nudge message when the API has not sent `nudge_text` (a fixture, or an older
 * officer response shape).
 */
export function readinessFails(row: GuildReadinessRow): string[] {
  const fails: string[] = [];
  const gap = row.gear_gap;
  if (gap !== null && gap.upgrades !== null && gap.upgrades >= 3) {
    fails.push(`${gap.upgrades} gear upgrades waiting (${Math.round(gap.gain_dps ?? 0)} DPS)`);
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
  return `${row.name}: ${fails.join(', ')} -- check before Thursday.`;
}

function capitalize(value: string): string {
  return value.length === 0 ? value : value[0].toUpperCase() + value.slice(1);
}
