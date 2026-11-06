// web/src/lib/sim/settings-rows.ts
// This run's inputs as data (sim spec 2026-11-06 §4.5 "This run"; spec 2026-11-04-sim-results
// 3.12): the key and value rows with their small lines, and the
// buff, debuff and consumable chips with the rule that files a buff under "No effect on your
// damage". Pure, so the figures are pinned by a test.
import { raceRows } from '../planner/reference';
import { hasNoDamageEffect, type BuffEffectRow, type SpecWeights } from './buff-effect';
import { IMPROVED_SUFFIX } from './buff-id';
import { CATALOGUE } from './buffs';
import { durationLabel } from './encounter';
import { grouped } from './result-copy';
import { settingsCopy } from './settings-copy';
import { humanise } from './humanise';
import type { SimResult } from './types';

/** A simbuffs.json row: its name and icon, and the data lane's effect fields (spec 11.4). */
export interface SettingsBuffRow extends BuffEffectRow {
  name: string;
  icon: string;
}

export type SettingsBuffs = Readonly<Record<string, SettingsBuffRow>>;

export interface SettingsRow {
  id: 'fight' | 'target' | 'execute' | 'race' | 'runs' | 'apm';
  label: string;
  value: string;
  /** The small muted line under the value. */
  sub?: string;
}

export interface SettingsChip {
  id: string;
  name: string;
  icon: string | null;
  /** The muted reason after the name: "replaced by Expose Armor". */
  reason: string | null;
}

export interface SettingsChipGroup {
  id: 'player' | 'target' | 'consumables';
  label: string;
  /** Every chip in the group, however it is filed. */
  count: number;
  chips: SettingsChip[];
  /** Buffs with no effect on this spec's damage, drawn under their own sub-label. */
  noEffect: SettingsChip[];
}

/**
 * Actions per minute: player button presses, which are a cast row whose key is not `other:*` and
 * has no `/<tag>` suffix (a tag is the queue press or a second row of one action).
 */
export function actionsPerMinute(result: SimResult): number {
  const minutes = result.summary.duration_ms / 60_000;
  if (minutes <= 0) return 0;
  const presses = result.summary.casts
    .filter((row) => !row.spell_name.startsWith('other:') && !row.spell_name.includes('/'))
    .reduce((sum, row) => sum + row.succeeded, 0);
  return presses / minutes;
}

/** "Level 63 boss" over "3,731 armor, 1,471 after debuffs"; a part the run does not know is left out. */
export function targetFact(result: SimResult): { value: string; sub?: string } {
  const { encounter } = result.request;
  const stats = result.target_stats;
  const level = stats?.level ?? encounter.target_level ?? 63;
  const armor = encounter.target_armor ?? stats?.armor;
  const after = stats?.armor_after_debuffs;
  const parts = [
    ...(armor === undefined ? [] : [settingsCopy.armor(grouped(armor))]),
    ...(after === undefined ? [] : [settingsCopy.afterDebuffs(grouped(after))]),
  ];
  const value = settingsCopy.target(level);
  return parts.length === 0 ? { value } : { value, sub: parts.join(', ') };
}

export function settingsRows(result: SimResult): SettingsRow[] {
  const { encounter, character } = result.request;
  const { keys } = settingsCopy;
  const seconds = encounter.duration_sec;
  const percent = Math.round(encounter.variation * 100);
  const low = durationLabel(Math.round(seconds * (1 - encounter.variation)));
  const high = durationLabel(Math.round(seconds * (1 + encounter.variation)));
  const race = raceRows.find((row) => row.slug === character.race)?.name;
  const length = durationLabel(seconds);
  return [
    {
      id: 'fight',
      label: keys.fight,
      value: settingsCopy.fight(length, percent),
      sub: settingsCopy.range(low, high),
    },
    { id: 'target', label: keys.target, ...targetFact(result) },
    ...(encounter.execute_ratio > 0
      ? [
          {
            id: 'execute' as const,
            label: keys.execute,
            value: settingsCopy.execute(Math.round(encounter.execute_ratio * 100)),
            // The request's own share of the fight, no claim about boss health (spec flag A7).
            sub: settingsCopy.executeLength(
              durationLabel(Math.round(seconds * encounter.execute_ratio)),
              length,
            ),
          },
        ]
      : []),
    ...(race === undefined ? [] : [{ id: 'race' as const, label: keys.race, value: race }]),
    { id: 'runs', label: keys.runs, value: grouped(result.iterations_run) },
    { id: 'apm', label: keys.apm, value: actionsPerMinute(result).toFixed(1) },
  ];
}

const baseId = (id: string): string =>
  id.endsWith(IMPROVED_SUFFIX) ? id.slice(0, -IMPROVED_SUFFIX.length) : id;

/** A chip's words: the file's name, its hand, its grade. A raw key never shows. */
export function chipName(id: string, buffs: SettingsBuffs): string {
  const base = baseId(id);
  const grade = id === base ? '' : settingsCopy.improved;
  const name = buffs[base]?.name ?? humanise(base.slice(base.indexOf(':') + 1));
  if (base.startsWith('main_hand_')) return `${name} ${settingsCopy.mainHand}`;
  if (base.startsWith('off_hand_')) return `${name} ${settingsCopy.offHand}`;
  return `${name}${grade}`;
}

const DEBUFF_IDS: ReadonlySet<string> = new Set(
  CATALOGUE.filter((row) => row.group === 'debuffs').map((row) => row.id),
);

/** The file's `target` flag; a build without it falls back to the engine's own Debuffs message. */
function landsOnTarget(id: string, buffs: SettingsBuffs): boolean {
  const base = baseId(id);
  return buffs[base]?.target ?? DEBUFF_IDS.has(base);
}

/** The debuff that holds over this one: another on the list, same group, lower rank. */
function replacedBy(id: string, list: readonly string[], buffs: SettingsBuffs): string | undefined {
  const own = buffs[baseId(id)]?.exclusive_with;
  if (own === undefined || own === null) return undefined;
  return list.find((other) => {
    const theirs = buffs[baseId(other)]?.exclusive_with;
    return theirs?.group === own.group && theirs.rank < own.rank;
  });
}

function toChip(id: string, list: readonly string[], buffs: SettingsBuffs, buildUrl: IconUrl): SettingsChip {
  const base = baseId(id);
  const winner = replacedBy(id, list, buffs);
  const icon = buffs[base]?.icon;
  return {
    id,
    name: chipName(id, buffs),
    icon: icon === undefined || icon === '' ? null : buildUrl(icon),
    reason: winner === undefined ? null : settingsCopy.replacedBy(chipName(winner, buffs)),
  };
}

export type IconUrl = (stem: string) => string;

function groupOf(
  id: SettingsChipGroup['id'],
  ids: readonly string[],
  buffs: SettingsBuffs,
  weights: SpecWeights,
  schools: readonly string[],
  iconUrl: IconUrl,
): SettingsChipGroup {
  const chips = ids.map((buff) => toChip(buff, ids, buffs, iconUrl));
  const quiet = (chip: SettingsChip): boolean => hasNoDamageEffect(buffs[baseId(chip.id)], weights, schools);
  return {
    id,
    label: settingsCopy.groups[id],
    count: ids.length,
    chips: id === 'consumables' ? chips : chips.filter((chip) => !quiet(chip)),
    noEffect: id === 'consumables' ? [] : chips.filter(quiet),
  };
}

/** On you, On the target, Consumables, in the request's own order. */
export function settingsChipGroups(
  result: SimResult,
  buffs: SettingsBuffs,
  weights: SpecWeights,
  schools: readonly string[],
  iconUrl: IconUrl,
): SettingsChipGroup[] {
  const { buffs: requested, consumes } = result.request.character;
  const onTarget = requested.filter((id) => landsOnTarget(id, buffs));
  const onPlayer = requested.filter((id) => !landsOnTarget(id, buffs));
  return [
    groupOf('player', onPlayer, buffs, weights, schools, iconUrl),
    groupOf('target', onTarget, buffs, weights, schools, iconUrl),
    groupOf('consumables', consumes, buffs, weights, schools, iconUrl),
  ];
}
