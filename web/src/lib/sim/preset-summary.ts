// web/src/lib/sim/preset-summary.ts
// Task 4 (dps-minmaxer, tank, newcomer, raid-leader): "what's in it" is one click away from
// the Buffs control, so a player never has to switch to Custom just to see what Raid-buffed
// actually applies. This groups a static preset's ids exactly the way the Custom panel
// groups them -- buffs.ts's own `BUFF_GROUPS` order and `rowsIn`'s id-alphabetical rows --
// so the disclosure can never disagree with what flipping to Custom shows, and names each
// id through buff-names.ts's own `buffLabel`, which already humanises an id the build's
// table has no row for rather than printing it bare (the raid-leader review's own finding).
import { buffLabel, type BuffNames } from './buff-names';
import { BUFF_GROUPS, rowsIn, type BuffGroupId } from './buffs';
import { presetConsumables, PRESET_BUFFS, type BuffPresetId } from './settings';

export interface PresetSummaryRow {
  id: string;
  label: string;
}

export interface PresetSummaryGroup {
  group: BuffGroupId;
  rows: PresetSummaryRow[];
}

/**
 * The preset's own ids, named and grouped. Solo carries none, so it always answers `[]`.
 * An id the synced catalogue does not carry (a build ahead of the IDS.md this ran
 * `sync:sim-ids` against) is silently absent from every group rather than shown ungrouped --
 * `settings.test.ts`'s own catalogue check is what keeps that from happening in practice.
 */
export function presetSummary(
  preset: Exclude<BuffPresetId, 'custom'>,
  referenceStat: string,
  names: BuffNames | null,
): PresetSummaryGroup[] {
  const applied = new Set([...PRESET_BUFFS[preset], ...presetConsumables(preset, referenceStat)]);
  return BUFF_GROUPS.map((group) => ({
    group,
    rows: rowsIn(group)
      .filter((row) => applied.has(row.id))
      .map((row) => ({ id: row.id, label: buffLabel(row.id, names) })),
  })).filter((entry) => entry.rows.length > 0);
}
