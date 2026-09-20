// durationLabel and encounterLabel live in encounter.ts, not here: the Worker's unfurl copy
// needs them too and it is built by a group that runs in parallel with this one. They are
// re-exported so every caller has one import for "the settings vocabulary".
import { durationLabel, encounterLabel } from './encounter';
import { referenceStatOf } from './spec-label';
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
// the whole run, so every id below is checked against IDS.md (settings.test.ts, over the
// same generated catalogue buffs.ts builds its panel from).
//
// This is a preset, not a picker, and the preset is the only list this file holds. The full
// vocabulary is several hundred entries and it lives in IDS.md; copying it into web/ would
// be a second copy to keep in step with the descriptors that generate it. The settings bar
// offers this preset plus a custom list, and surfaces the engine's own error verbatim when
// an id in that list does not resolve.
//
// Task 4 (dps-minmaxer D16 BLOCKER): the old preset was five buffs and two consumables,
// +6% over Solo where vanilla's real answer for a melee is +60% to +120%. This is the full
// standard set instead, copied verbatim from sim/adapter/testdata/warrior-fury.request.json
// (identical in mage-frost.request.json, so the buff half never splits by spec -- Ruling 2).
export const RAID_BUFFS: readonly string[] = [
  'curse_of_elements',
  'curse_of_recklessness',
  'curse_of_shadow',
  'curse_of_weakness',
  'demoralizing_roar',
  'demoralizing_shout',
  'expose_armor',
  'faerie_fire',
  'improved_scorch',
  'improved_shadow_bolt',
  'insect_swarm',
  'judgement_of_light',
  'judgement_of_the_crusader',
  'judgement_of_wisdom',
  'scorpid_sting',
  'shadow_weaving',
  'stormstrike',
  'sunder_armor',
  'thunder_clap',
  'winters_chill',
  'blessing_of_kings',
  'blessing_of_might',
  'blessing_of_sanctuary',
  'blessing_of_wisdom',
  'fengus_ferocity',
  'moldars_moxie',
  'rallying_cry_of_the_dragonslayer',
  'sayges_fortune',
  'slipkiks_savvy',
  'songflower_serenade',
  'spirit_of_zandalar',
  'warchiefs_blessing',
  'arcane_brilliance',
  'battle_shout',
  'blood_pact',
  'devotion_aura',
  'divine_spirit',
  'fire_resistance_aura',
  'fire_resistance_totem',
  'frost_resistance_aura',
  'frost_resistance_totem',
  'gift_of_the_wild',
  'grace_of_air_totem',
  'leader_of_the_pack',
  'mana_spring_totem',
  'moonkin_aura',
  'nature_resistance_totem',
  'power_word_fortitude',
  'retribution_aura',
  'sanctity_aura',
  'shadow_protection',
  'strength_of_earth_totem',
  'thorns',
  'trueshot_aura',
];

export const PRESET_BUFFS: Record<Exclude<BuffPresetId, 'custom'>, readonly string[]> = {
  'raid-buffed': RAID_BUFFS,
  solo: [],
};

/**
 * Consumables split by role (task-4-brief.md, Ruling 2): the engine's `Consumes` has one
 * field per effect, so a caster's flask and a warrior's flask can never both apply at once.
 * The off-hand imbue in `PHYSICAL_CONSUMABLES` is included even for a character with no
 * off-hand weapon -- the engine ignores an imbue it has nowhere to apply, so this file does
 * not branch on gear.
 */
export const PHYSICAL_CONSUMABLES: readonly string[] = [
  'flask_of_the_titans',
  'elixir_of_the_mongoose',
  'juju_power',
  'winterfall_firewater',
  'food_bless_sunfruit',
  'main_hand_imbue:elemental_sharpening_stone',
  'off_hand_imbue:elemental_sharpening_stone',
  'ground_scorpok_assay',
  'mighty_rage_potion',
];

export const CASTER_CONSUMABLES: readonly string[] = [
  'flask_of_supreme_power',
  'greater_arcane_elixir',
  'food_nightfin_soup',
  'main_hand_imbue:brilliant_wizard_oil',
  'cerebral_cortex_compound',
  'mageblood_potion',
  'major_mana_potion',
];

/**
 * The Raid-buffed preset's consumable half, keyed off `Spec.reference_stat` (specs.ts) --
 * `'spell_power'` is caster, anything else physical. Physical is the deliberate fallback
 * for a spec `reference_stat` does not resolve, since an unrecognised spec still needs an
 * answer from `defaultSettings()`: an empty preset that still claims "Raid-buffed" would be
 * worse than defaulting to the wrong half, and every spec the site simulates but one
 * (druid-balance and the other caster specs) is physical to begin with.
 */
export function presetConsumables(
  preset: Exclude<BuffPresetId, 'custom'>,
  referenceStat: string,
): readonly string[] {
  if (preset === 'solo') return [];
  return referenceStat === 'spell_power' ? CASTER_CONSUMABLES : PHYSICAL_CONSUMABLES;
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
 * `referenceStat` is `Spec.reference_stat` ('attack_power' | 'spell_power'), or the result
 * of `spec-label.ts`'s `referenceStatOf` for a caller that only has the spec slug -- both
 * this and `withPreset` stay pure over it rather than resolving a spec themselves, so a
 * store can call them from wherever it already knows the answer (live-dps.svelte.ts and
 * SharePanel.svelte know it from the character in hand; the sim/bulk stores own the
 * re-apply, below).
 */
export function defaultSettings(referenceStat: string): SimSettings {
  return {
    encounter: applyFightStyle({ ...DEFAULT_ENCOUNTER }, DEFAULT_STYLE_ID),
    preset: 'raid-buffed',
    buffs: [...PRESET_BUFFS['raid-buffed']],
    consumables: [...presetConsumables('raid-buffed', referenceStat)],
    cooldowns: [],
  };
}

export function withPreset(settings: SimSettings, preset: BuffPresetId, referenceStat: string): SimSettings {
  if (preset === 'custom') return { ...settings, preset };
  return {
    ...settings,
    preset,
    buffs: [...PRESET_BUFFS[preset]],
    consumables: [...presetConsumables(preset, referenceStat)],
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
export function withSpecForPreset(settings: SimSettings, spec: string): SimSettings {
  if (settings.preset === 'custom') return settings;
  return withPreset(settings, settings.preset, referenceStatOf(spec));
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
