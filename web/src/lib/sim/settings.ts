// durationLabel and encounterLabel live in encounter.ts, not here: the Worker's unfurl copy
// needs them too and it is built by a group that runs in parallel with this one. They are
// re-exported so every caller has one import for "the settings vocabulary".
import { durationLabel, encounterLabel } from './encounter';
import { applyFightStyle, DEFAULT_STYLE_ID, type FightStyleId } from './styles';
import { DEFAULT_ENCOUNTER, TARGET_TYPE_IDS, type CooldownSpec, type EncounterSpec } from './types';

export { durationLabel, encounterLabel };

// Design 4.1 and contract A3, which ratifies both numbers on the Go side too
// (`api.MinDurationSec = 20`, `api.MaxDurationSec = 600`). The old floor was a minute and
// the old ceiling eight.
export const MIN_DURATION_SEC = 20;
export const MAX_DURATION_SEC = 600;
export const DURATION_STEP_SEC = 30;
export const MAX_TARGETS = 10;

/**
 * Three short lengths for the openers and the burst windows people actually ask about,
 * then thirty-second steps to ten minutes. A uniform step from twenty seconds would put
 * twenty options under a minute, which is a select nobody can use on a phone.
 */
export const DURATIONS: readonly number[] = [
  20,
  30,
  45,
  ...Array.from(
    { length: (MAX_DURATION_SEC - 60) / DURATION_STEP_SEC + 1 },
    (_, i) => 60 + i * DURATION_STEP_SEC,
  ),
];

export const MAX_VARIATION = 0.3;
export const VARIATIONS: readonly number[] = [0, 0.05, 0.1, 0.15, 0.2, 0.25, 0.3];

export const TARGET_LEVELS: readonly number[] = [60, 61, 62, 63];
export const DEFAULT_TARGET_LEVEL = 63;

/**
 * Contract A8: the engine's own 3,731 at level 63 and a linear fall to the level-60
 * figure. `target_armor: 0` still means "the level's preset" and is what the request
 * carries by default -- these numbers exist so the override control can say what it is
 * overriding rather than showing an empty field. A better source replaces the three
 * interior numbers on the Go side and here together.
 */
export const TARGET_ARMOR_BY_LEVEL: Record<number, number> = {
  60: 3300,
  61: 3444,
  62: 3588,
  63: 3731,
};

/** A generous bound on the override field; the presets above are far below it. */
export const MAX_TARGET_ARMOR = 20_000;
export const TARGET_TYPES: readonly string[] = TARGET_TYPE_IDS;

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
  /** Cooldown timing rows (contract 1.7). Empty means "everything on cooldown". */
  cooldowns: CooldownSpec[];
}

export function defaultSettings(): SimSettings {
  return {
    encounter: applyFightStyle({ ...DEFAULT_ENCOUNTER }, DEFAULT_STYLE_ID),
    preset: 'raid-buffed',
    buffs: [...PRESET_BUFFS['raid-buffed']],
    consumables: [...PRESET_CONSUMABLES['raid-buffed']],
    cooldowns: [],
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

/**
 * Clamps to [low, high], except a non-finite `value` -- NaN or +/-Infinity, the shape an
 * emptied number input takes once parsed -- falls back to `current` instead. Snapping an
 * emptied field to the minimum would silently pick a value nobody asked for; a clamped
 * NaN would reach `JSON.stringify` as `null` on a field the contract requires as a plain
 * `number`. Either way the setting has to stay where it was until a real number replaces
 * it, which is what every numeric setter below relies on this for.
 */
function clamp(value: number, low: number, high: number, current: number): number {
  if (!Number.isFinite(value)) return current;
  return Math.min(high, Math.max(low, value));
}

/**
 * A style owns targets, the execute ratio, movement, the target-count timeline and the
 * dummy flag. Setting any of those by hand means the encounter is no longer the style's,
 * so the label goes -- a run that says "Cleave, 3 targets" while simming four targets
 * would be a lie in the saved sim's own title.
 */
function detached(encounter: EncounterSpec, overrides: Partial<EncounterSpec>): EncounterSpec {
  return { ...encounter, style: '', ...overrides };
}

/** The style this encounter still is, or "" once a style-owned field was changed by hand. */
export function styleIdOf(settings: SimSettings): string {
  return settings.encounter.style ?? '';
}

export function withStyle(settings: SimSettings, id: FightStyleId): SimSettings {
  return { ...settings, encounter: applyFightStyle(settings.encounter, id) };
}

export function withDuration(settings: SimSettings, seconds: number): SimSettings {
  const clamped = clamp(
    Math.round(seconds),
    MIN_DURATION_SEC,
    MAX_DURATION_SEC,
    settings.encounter.duration_sec,
  );
  return { ...settings, encounter: { ...settings.encounter, duration_sec: clamped } };
}

export function withTargets(settings: SimSettings, targets: number): SimSettings {
  const clamped = clamp(Math.round(targets), 1, MAX_TARGETS, settings.encounter.targets);
  return { ...settings, encounter: detached(settings.encounter, { targets: clamped }) };
}

export function withVariation(settings: SimSettings, variation: number): SimSettings {
  const clamped = clamp(variation, 0, MAX_VARIATION, settings.encounter.variation);
  return { ...settings, encounter: { ...settings.encounter, variation: clamped } };
}

export function withTargetLevel(settings: SimSettings, level: number): SimSettings {
  const clamped = clamp(
    Math.round(level),
    TARGET_LEVELS[0],
    TARGET_LEVELS[TARGET_LEVELS.length - 1],
    settings.encounter.target_level ?? DEFAULT_TARGET_LEVEL,
  );
  return { ...settings, encounter: { ...settings.encounter, target_level: clamped } };
}

export function withTargetArmor(settings: SimSettings, armor: number): SimSettings {
  const clamped = clamp(Math.round(armor), 0, MAX_TARGET_ARMOR, settings.encounter.target_armor ?? 0);
  return { ...settings, encounter: { ...settings.encounter, target_armor: clamped } };
}

/**
 * What the target-armor field should DISPLAY (tank MAJOR, review.md:227-229): `0` is the
 * contract's "use the level's preset", not an empty override, so the field shows the
 * preset number itself rather than going blank -- and follows the target-level select when
 * that changes. The wire value is untouched: `target_armor` stays 0 in settings state until
 * the player types something else (settings.ts:40-51 above); this only decides what the
 * input reads.
 */
export function targetArmorField(encounter: EncounterSpec): { value: string; preset: number } {
  const level = encounter.target_level ?? DEFAULT_TARGET_LEVEL;
  const preset = TARGET_ARMOR_BY_LEVEL[level] ?? TARGET_ARMOR_BY_LEVEL[DEFAULT_TARGET_LEVEL];
  const armor = encounter.target_armor ?? 0;
  return { value: String(armor === 0 ? preset : armor), preset };
}

/** An id outside the contract's vocabulary reads as "any", never as itself. */
export function withTargetType(settings: SimSettings, type: string): SimSettings {
  const known = TARGET_TYPES.includes(type) ? (type as EncounterSpec['target_type']) : '';
  return { ...settings, encounter: { ...settings.encounter, target_type: known } };
}

export function withDummy(settings: SimSettings, on: boolean): SimSettings {
  return { ...settings, encounter: detached(settings.encounter, { dummy: on }) };
}

export function withExecutePhase(settings: SimSettings, on: boolean): SimSettings {
  return {
    ...settings,
    encounter: detached(settings.encounter, {
      execute_ratio: on ? DEFAULT_ENCOUNTER.execute_ratio : 0,
    }),
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
