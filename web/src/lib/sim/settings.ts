// durationLabel and encounterLabel live in encounter.ts, not here: the Worker's unfurl copy
// needs them too and it is built by a group that runs in parallel with this one. They are
// re-exported so every caller has one import for "the settings vocabulary".
import { durationLabel, encounterLabel } from './encounter';
import { DEFAULT_ENCOUNTER, type EncounterSpec } from './types';

export { durationLabel, encounterLabel };

export const MIN_DURATION_SEC = 60;
export const MAX_DURATION_SEC = 480;
export const DURATION_STEP_SEC = 30;
export const MAX_TARGETS = 10;

export const DURATIONS: readonly number[] = Array.from(
  { length: (MAX_DURATION_SEC - MIN_DURATION_SEC) / DURATION_STEP_SEC + 1 },
  (_, i) => MIN_DURATION_SEC + i * DURATION_STEP_SEC,
);

export const BUFF_PRESETS = [
  { id: 'raid-buffed', label: 'Raid-buffed' },
  { id: 'solo', label: 'Solo' },
  { id: 'custom', label: 'Custom' },
] as const;
export type BuffPresetId = (typeof BUFF_PRESETS)[number]['id'];

// Buff and consumable ids, not spell ids and not kebab case. The engine publishes the
// vocabulary as sim/request/IDS.md, generated from its own protobuf descriptors: a buff id
// is a field name of Debuffs, IndividualBuffs, PartyBuffs or RaidBuffs, and a consumable id
// is an enum value name or a client item id spelled "item:<id>", either optionally
// qualified by its field and an imbue always ("off_hand_imbue:shadow_oil"). An id the
// engine cannot map is ErrUnknownBuff, ErrUnknownConsume or ErrAmbiguousConsume and fails
// the whole run, so every id below is checked against IDS.md in Step 5a.
//
// This is a preset, not a picker, and the preset is the only list this file holds. The full
// vocabulary is several hundred entries and it lives in IDS.md; copying it into web/ would
// be a second copy to keep in step with the descriptors that generate it. The settings bar
// (Task 12) offers this preset plus a custom list, and surfaces the engine's own error
// verbatim when an id in that list does not resolve.
export const PRESET_BUFFS: Record<Exclude<BuffPresetId, 'custom'>, string[]> = {
  'raid-buffed': [
    'blessing_of_kings',
    'battle_shout',
    'gift_of_the_wild',
    'power_word_fortitude',
    'arcane_brilliance',
  ],
  solo: [],
};

export const PRESET_CONSUMABLES: Record<Exclude<BuffPresetId, 'custom'>, string[]> = {
  'raid-buffed': ['flask_of_supreme_power', 'elixir_of_the_mongoose'],
  solo: [],
};

export interface SimSettings {
  encounter: EncounterSpec;
  preset: BuffPresetId;
  buffs: string[];
  consumables: string[];
}

export function defaultSettings(): SimSettings {
  return {
    encounter: { ...DEFAULT_ENCOUNTER },
    preset: 'raid-buffed',
    buffs: [...PRESET_BUFFS['raid-buffed']],
    consumables: [...PRESET_CONSUMABLES['raid-buffed']],
  };
}

export function withPreset(settings: SimSettings, preset: BuffPresetId): SimSettings {
  if (preset === 'custom') return { ...settings, preset };
  return {
    ...settings,
    preset,
    buffs: [...PRESET_BUFFS[preset]],
    consumables: [...PRESET_CONSUMABLES[preset]],
  };
}

export function withDuration(settings: SimSettings, seconds: number): SimSettings {
  const clamped = Math.min(MAX_DURATION_SEC, Math.max(MIN_DURATION_SEC, Math.round(seconds)));
  return { ...settings, encounter: { ...settings.encounter, duration_sec: clamped } };
}

export function withTargets(settings: SimSettings, targets: number): SimSettings {
  const clamped = Math.min(MAX_TARGETS, Math.max(1, Math.round(targets)));
  return { ...settings, encounter: { ...settings.encounter, targets: clamped } };
}

export function withExecutePhase(settings: SimSettings, on: boolean): SimSettings {
  return {
    ...settings,
    encounter: { ...settings.encounter, execute_ratio: on ? DEFAULT_ENCOUNTER.execute_ratio : 0 },
  };
}

export function executePhaseOn(settings: SimSettings): boolean {
  return settings.encounter.execute_ratio > 0;
}

export function settingsLabel(settings: SimSettings): string {
  const preset = BUFF_PRESETS.find((row) => row.id === settings.preset)?.label ?? 'Custom';
  const targets =
    settings.encounter.targets === 1 ? 'single target' : `${settings.encounter.targets} targets`;
  return `${preset}, ${durationLabel(settings.encounter.duration_sec)}, ${targets}`;
}
