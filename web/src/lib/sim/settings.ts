// durationLabel and encounterLabel live in encounter.ts, not here: the Worker's unfurl copy
// needs them too and it is built by a group that runs in parallel with this one. They are
// re-exported so every caller has one import for "the settings vocabulary".
import { MAX_LEVEL } from './character';
import { durationLabel, encounterLabel } from './encounter';
import { applyFightStyle, DEFAULT_STYLE_ID, type FightStyleId } from './styles';
import raidPreset from '../../data/generated/raid-preset.json';
import { raidConsumesAt } from './raid-consumes';
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

/**
 * The level-aware sim design (2026-09-27), design step 5: the default target is 3 above
 * the character's own level, not a flat 63 -- 63 was always just this rule applied to a
 * character pinned at MAX_LEVEL, which every character used to be. `sim/api/envelope.go`
 * computes the same offset on the Go side once lane T2 wires a request's default there.
 */
export function defaultTargetLevel(characterLevel: number): number {
  return characterLevel + 3;
}

/**
 * The plain constant every caller with no character in scope still uses (`DEFAULT_ENCOUNTER`,
 * the settings bar's own fallback before a character is picked) -- `defaultTargetLevel`
 * applied to MAX_LEVEL, unchanged in value (63) from before this design.
 */
export const DEFAULT_TARGET_LEVEL = defaultTargetLevel(MAX_LEVEL);

/** How far above the character the target may be: 0 (same level) to this (a boss). */
export const MAX_TARGET_LEVEL_OFFSET = defaultTargetLevel(0);

/**
 * A stored `target_level` is an offset above the character, written as the absolute level
 * of a character at the cap: 60 is "the same level", 63 is "a boss, three above" (the
 * settings picker and every saved setting predate characters below 60). Resolved for the
 * character actually being simmed, so a level-15 export fights a level-18 target rather
 * than a level-63 one the engine refuses. Anything outside the offset range, including a
 * setting saved by hand, reads as the default (three above).
 */
export function targetLevelFor(encounter: EncounterSpec, characterLevel: number): number {
  const stored = encounter.target_level ?? DEFAULT_TARGET_LEVEL;
  const offset = stored - MAX_LEVEL;
  if (!Number.isInteger(offset) || offset < 0 || offset > MAX_TARGET_LEVEL_OFFSET) {
    return defaultTargetLevel(characterLevel);
  }
  return characterLevel + offset;
}

/** The encounter as the request sends it: `target_level` resolved for this character. */
export function encounterFor(encounter: EncounterSpec, characterLevel: number): EncounterSpec {
  return { ...encounter, target_level: targetLevelFor(encounter, characterLevel) };
}

/**
 * Contract A8's own rule -- the engine's boss preset, 3,731 at level 63, and a linear fall
 * to the level-60 figure, 3,300 (sim/api/envelope.go's TargetArmorByLevel) -- and below 60
 * the engine's own estimate for a levelling character's target (design step 5: "the target
 * armour table extended by level"): the level-60 figure scaled by level
 * (envelope.go's armorBelowMaxLevel, rounded). No per-level source for a levelling mob's
 * armour exists yet, so both sides derive it from the one number; this file mirrors the
 * engine's rule so the field never names an armor the run does not use.
 */
const ARMOR_LEVEL_60 = 3300;
const ARMOR_BOSS_LEVEL = 63;
const ARMOR_BOSS = 3731;
function armorForLevel(level: number): number {
  if (level < MAX_LEVEL) return Math.round((ARMOR_LEVEL_60 * level) / MAX_LEVEL);
  const raised = ((level - 60) * (ARMOR_BOSS - ARMOR_LEVEL_60)) / (ARMOR_BOSS_LEVEL - 60);
  return Math.ceil(ARMOR_LEVEL_60 + raised);
}
export const TARGET_ARMOR_BY_LEVEL: Record<number, number> = Object.fromEntries(
  Array.from({ length: MAX_LEVEL + 3 }, (_, index) => index + 1).map((level) => [
    level,
    armorForLevel(level),
  ]),
);

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
// the whole run, so every id below is checked against IDS.md (settings.test.ts, over the
// same generated catalogue buffs.ts builds its panel from).
//
// The Raid-buffed preset's buff and debuff half is the curated raid preset
// (data/curated/presets.json), synced into src/data/generated/raid-preset.json by
// scripts/sync-raid-preset.mjs. The BiS ranker, the tier list and every sim tool run their
// level-60 numbers under that same preset, so the page sims a character with the buffs and
// debuffs the published number was simmed with. The engine applies the faction-gated ones
// (Blessing of Might and Wisdom for the Alliance; Strength of Earth and Mana Spring Totem
// for the Horde) to the side that can receive them, so one list serves both factions.
// Every id is checked against IDS.md (settings.test.ts), and the curated file's own loader
// (sim/request/preset.go) refuses an id the engine cannot map.
export const RAID_BUFFS: readonly string[] = [...raidPreset.buffs, ...raidPreset.debuffs];

export const PRESET_BUFFS: Record<Exclude<BuffPresetId, 'custom'>, readonly string[]> = {
  'raid-buffed': RAID_BUFFS,
  solo: [],
};

/**
 * The preset's consumable half for one spec key ("rogue-combat") at the character's level,
 * from the raid consumables table (raid-consumes.ts; load it first).
 * Solo carries none. A spec the raid set does not know (no character yet, or a malformed
 * one) carries none either, rather than another spec's set: the request then sims what it
 * says it does.
 */
export function presetConsumables(
  preset: Exclude<BuffPresetId, 'custom'>,
  spec: string,
  level: number,
): readonly string[] {
  if (preset === 'solo') return [];
  return raidConsumesAt(spec, level);
}

export interface SimSettings {
  encounter: EncounterSpec;
  preset: BuffPresetId;
  buffs: string[];
  consumables: string[];
  /** Cooldown timing rows (contract 1.7). Empty means "everything on cooldown". */
  cooldowns: CooldownSpec[];
}

/**
 * `spec` is the character's spec key ("warrior-fury") and `level` its level; both default to
 * "no character yet" (no spec, so no consumables). This and `withPreset` stay pure over them,
 * so a store calls them from wherever it already knows the character (live-dps.svelte.ts and
 * SharePanel.svelte from the character in hand; the sim/bulk stores own the re-apply, below).
 */
export function defaultSettings(spec = '', level: number = MAX_LEVEL): SimSettings {
  return {
    encounter: applyFightStyle({ ...DEFAULT_ENCOUNTER }, DEFAULT_STYLE_ID),
    preset: 'raid-buffed',
    buffs: [...PRESET_BUFFS['raid-buffed']],
    consumables: [...presetConsumables('raid-buffed', spec, level)],
    cooldowns: [],
  };
}

export function withPreset(
  settings: SimSettings,
  preset: BuffPresetId,
  spec: string,
  level: number,
): SimSettings {
  if (preset === 'custom') return { ...settings, preset };
  return {
    ...settings,
    preset,
    buffs: [...PRESET_BUFFS[preset]],
    consumables: [...presetConsumables(preset, spec, level)],
  };
}

/**
 * The store's own re-apply rule (task-4-brief.md's ambiguity, resolved): a preset is
 * spec-dependent now, so a store calls this after adopting a new character, and only this
 * -- never `defaultSettings()` again -- so a run in progress, a typed title or any other
 * settings field the player already touched survives a character swap. A no-op for Custom:
 * switching from a mage to a warrior must not silently discard a player's own ticks, only
 * the wizard oil a *preset* would have carried. Pure, like `withPreset` itself.
 */
export function withSpecForPreset(settings: SimSettings, spec: string, level: number): SimSettings {
  if (settings.preset === 'custom') return settings;
  return withPreset(settings, settings.preset, spec, level);
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

/**
 * `armor === null` clears the override: the request omits `target_armor` entirely and the
 * engine uses the level's preset. Any number, INCLUDING 0, is an explicit override -- 0 is
 * a real request (an unarmoured target), not a second way to spell "no override"
 * (2026-09-21 result-page review, Defect 3: a bare `number` here could not represent
 * "cleared" at all, so the settings bar had nothing to call `withTargetArmor` with that
 * meant anything other than some number, and a cleared field and a typed 0 collapsed into
 * the same request). Out-of-range or non-finite input clamps against whatever override is
 * already set, falling back to 0 when there is none yet -- the same defensive shape every
 * other numeric control on this bar already has.
 */
export function withTargetArmor(settings: SimSettings, armor: number | null): SimSettings {
  if (armor === null) {
    return { ...settings, encounter: { ...settings.encounter, target_armor: undefined } };
  }
  const clamped = clamp(Math.round(armor), 0, MAX_TARGET_ARMOR, settings.encounter.target_armor ?? 0);
  return { ...settings, encounter: { ...settings.encounter, target_armor: clamped } };
}

/**
 * What the target-armor field should DISPLAY. `target_armor` ABSENT (`undefined`) is "no
 * override": the field reads blank, and its `placeholder` -- SettingsSheet.svelte's own
 * concern, not this function's -- shows the preset as a hint rather than as real content
 * (2026-09-21 result-page review, Defect 3; tank MAJOR, review.md:227-229 wanted the
 * preset visible some way, which the placeholder still gives it, without making 0
 * unreachable as a real value). Present, including `0`, DISPLAYS as exactly what it is:
 * the player's own override, never silently swapped for the preset. `preset` and `level`
 * ride along either way, for the help text and the placeholder.
 */
export function targetArmorField(
  encounter: EncounterSpec,
  characterLevel: number = MAX_LEVEL,
): { value: string; preset: number; level: number } {
  const level = targetLevelFor(encounter, characterLevel);
  const preset = TARGET_ARMOR_BY_LEVEL[level] ?? TARGET_ARMOR_BY_LEVEL[DEFAULT_TARGET_LEVEL];
  // `level` rides on the return value so a caller (SettingsSheet.svelte's `targetArmorNote`)
  // reads the SAME level `preset` was computed from, rather than a second, independent
  // `encounter.target_level ?? DEFAULT_TARGET_LEVEL` lookup that only agrees with it by
  // coincidence -- fix round, Minor 1, the same shape an earlier round already removed for
  // `preset` itself.
  return { value: encounter.target_armor === undefined ? '' : String(encounter.target_armor), preset, level };
}

/** An id outside the contract's vocabulary reads as the default, never as itself. */
export function withTargetType(settings: SimSettings, type: string): SimSettings {
  const known = TARGET_TYPES.includes(type)
    ? (type as EncounterSpec['target_type'])
    : DEFAULT_ENCOUNTER.target_type;
  return { ...settings, encounter: { ...settings.encounter, target_type: known } };
}

/**
 * The type the engine fights, for the target-type control. A request stored before the
 * field existed carries none, and sim/request fights a Humanoid for it, so the control
 * says Humanoid rather than a choice the run is not making.
 */
export function targetTypeShown(encounter: EncounterSpec): string {
  const type = encounter.target_type ?? '';
  return type === '' ? 'humanoid' : type;
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
