// web/src/lib/guild/readiness-view.ts
// Guild control-centre spec §4.E: the raid leader's pre-pull board. Every cell's text and
// the worst-first sort are pure functions over `GuildReadinessRow` (api.ts, the contract's
// own shape) so GuildReadiness.svelte never re-derives a number the API already sent.
import type { GuildReadinessRow } from './api';

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

export const NO_GEAR_CONSENT = 'no gear consent';
export const CONSENT_NEEDED = 'consent needed';
export const ALL_ENCHANTED = 'All enchanted';

/** Spec §4.E's Gear gap cell. `gear`/`gear_bags` consent only -- `gear_gap === null` (the
 *  API's fail-closed shape for `roster` consent) always reads the honest "no gear consent"
 *  line, never a leaked figure. */
export function gearGapLabel(row: GuildReadinessRow): string {
  const gap = row.gear_gap;
  if (gap === null || gap.upgrades === null) return NO_GEAR_CONSENT;
  const gain = Math.round(gap.gain_dps ?? 0);
  return `${gap.upgrades} upgrade${gap.upgrades === 1 ? '' : 's'} · +${gain} DPS`;
}

/** Spec §4.E's Enchants cell -- `checked: false` means no gear consent, read the same as
 *  Gear gap's own consent gate (the stricter `gear_bags` floor applies to Consumables
 *  only, enchants share Gear gap's `gear` floor). */
export function enchantLabel(row: GuildReadinessRow): string {
  if (!row.enchants.checked) return CONSENT_NEEDED;
  if (row.enchants.missing_slots.length === 0) return ALL_ENCHANTED;
  return row.enchants.missing_slots.map(capitalize).join(', ');
}

/** Spec §4.E's Consumables cell -- `gear_bags` consent only, stricter than Gear gap/
 *  Enchants' `gear` floor (a `gear`-only consent reads "consent needed" here even though
 *  it already unlocks the other two cells). */
export function consumablesLabel(row: GuildReadinessRow): string {
  if (row.consumables.state === 'unknown') return CONSENT_NEEDED;
  return row.consumables.state === 'stocked' ? 'Stocked' : 'Short';
}

/** Spec §4.E's Talent points cell. */
export function talentPointsLabel(row: GuildReadinessRow): string {
  return row.talent_points_unspent > 0 ? `${row.talent_points_unspent} unspent` : '0';
}

/** Spec §4.E's Item level cell: signed delta against the roster's own median, `—` when the
 *  character carries no item level at all (no gear consent). */
export function itemLevelLabel(row: GuildReadinessRow): string {
  if (row.item_level === null) return '—';
  const delta = row.item_level_delta ?? 0;
  const sign = delta >= 0 ? '+' : '';
  return `${row.item_level} (${sign}${delta})`;
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
    fails.push(`no enchant: ${row.enchants.missing_slots.map(capitalize).join(', ')}`);
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
