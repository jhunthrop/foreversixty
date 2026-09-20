import { describe, expect, it } from 'vitest';
import generated from '../../data/generated/sim-ids.json';
import { selectedIn } from './buffs';
import { simCopy } from './copy';
import {
  BUFF_PRESETS,
  CASTER_CONSUMABLES,
  DURATIONS,
  MAX_DURATION_SEC,
  MAX_TARGETS,
  MAX_TARGET_ARMOR,
  MAX_VARIATION,
  MIN_DURATION_SEC,
  PHYSICAL_CONSUMABLES,
  PRESET_BUFFS,
  RAID_BUFFS,
  TARGET_ARMOR_BY_LEVEL,
  TARGET_LEVELS,
  TARGET_TYPES,
  VARIATIONS,
  defaultSettings,
  durationLabel,
  executePhaseOn,
  presetConsumables,
  settingsLabel,
  styleIdOf,
  targetArmorField,
  withDummy,
  withDuration,
  withExecutePhase,
  withPreset,
  withSpecForPreset,
  withStyle,
  withTargetArmor,
  withTargetLevel,
  withTargetType,
  withTargets,
  withVariation,
} from './settings';
import { DEFAULT_ENCOUNTER } from './types';

const PHYSICAL = 'attack_power';
const CASTER = 'spell_power';

describe('defaultSettings', () => {
  it("is the contract's EncounterSpec defaults, raid-buffed, on Patchwerk", () => {
    const settings = defaultSettings(PHYSICAL);
    expect(settings.encounter).toEqual({
      duration_sec: 180,
      variation: 0.2,
      targets: 1,
      execute_ratio: 0.25,
      profile: '',
      style: 'patchwerk',
      target_level: 63,
      target_armor: 0,
      target_type: '',
      dummy: false,
    });
    expect(settings.preset).toBe('raid-buffed');
    expect(settings.buffs.length).toBeGreaterThan(0);
    expect(settings.cooldowns).toEqual([]);
  });

  it('carries the physical consumables for attack_power and the caster ones for spell_power', () => {
    expect(defaultSettings(PHYSICAL).consumables).toEqual([...PHYSICAL_CONSUMABLES]);
    expect(defaultSettings(CASTER).consumables).toEqual([...CASTER_CONSUMABLES]);
  });

  it('falls back to the physical set for a reference_stat it does not recognise, rather than an empty preset', () => {
    expect(defaultSettings('').consumables).toEqual([...PHYSICAL_CONSUMABLES]);
    expect(defaultSettings('unknown-stat').consumables).toEqual([...PHYSICAL_CONSUMABLES]);
  });
});

/**
 * Task 4 (dps-minmaxer D16 BLOCKER): the old preset was five buffs and two consumables,
 * +6% over Solo. This is the full standard set instead -- the brief's 54 buff ids, copied
 * verbatim from sim/adapter/testdata/warrior-fury.request.json, plus the 9 physical and 7
 * caster consumable ids the controller verified against the real engine (8b2169e61): every
 * one of the 70 accepted, no ErrUnknownBuff, no ErrUnknownConsume, no ErrAmbiguousConsume.
 */
describe('the Raid-buffed preset', () => {
  it('is exactly the brief’s 54 buff ids, never split by spec (Ruling 2)', () => {
    expect(RAID_BUFFS).toHaveLength(54);
    expect(PRESET_BUFFS['raid-buffed']).toBe(RAID_BUFFS);
  });

  it('splits the consumable half by reference_stat: 9 physical ids, 7 caster ids', () => {
    expect(PHYSICAL_CONSUMABLES).toHaveLength(9);
    expect(CASTER_CONSUMABLES).toHaveLength(7);
    expect(presetConsumables('raid-buffed', PHYSICAL)).toBe(PHYSICAL_CONSUMABLES);
    expect(presetConsumables('raid-buffed', CASTER)).toBe(CASTER_CONSUMABLES);
    // Never both at once -- the engine's Consumes has one field per effect.
    expect(PHYSICAL_CONSUMABLES.some((id) => CASTER_CONSUMABLES.includes(id))).toBe(false);
  });

  it('carries every buff and consumable id in the synced catalogue -- an id the engine cannot map is ErrUnknownBuff/ErrUnknownConsume and fails the whole run', () => {
    const published = new Set([
      ...generated.buffs.map((row) => row.id),
      ...generated.consumables.map((row) => row.id),
    ]);
    for (const id of [...RAID_BUFFS, ...PHYSICAL_CONSUMABLES, ...CASTER_CONSUMABLES]) {
      expect(published.has(id), id).toBe(true);
    }
  });

  it('solo stays self-only', () => {
    expect(PRESET_BUFFS.solo).toEqual([]);
    expect(presetConsumables('solo', PHYSICAL)).toEqual([]);
    expect(presetConsumables('solo', CASTER)).toEqual([]);
  });

  /**
   * The findings' own counters (task-4-brief.md): "RAID BUFFS 4/30, ... ON THE TARGET
   * 0/24, ... WEAPON OILS AND STONES 0/48" -- the group counts `selectedIn` (buffs.ts)
   * reports for the Custom panel, over the settings this preset now writes. This is no
   * longer 4, 0 and 0: the fix landed in the counters the panel actually renders, not only
   * in a list nobody reads a length off.
   */
  it('reports real totals in the group counters the findings quoted as 4/30, 0/24 and 0/48', () => {
    const settings = defaultSettings(PHYSICAL);
    const selection = { buffs: settings.buffs, consumables: settings.consumables };
    expect(selectedIn(selection, 'raid-buffs').length).toBeGreaterThan(4);
    expect(selectedIn(selection, 'debuffs').length).toBeGreaterThan(0);
    expect(selectedIn(selection, 'weapon-imbue').length).toBeGreaterThan(0);
    expect(selectedIn(selection, 'world-buffs').length).toBeGreaterThan(0);
    // No buff id is lost between groups: every group's members sum back to the full list.
    const grouped = [
      ...selectedIn(selection, 'raid-buffs'),
      ...selectedIn(selection, 'party-buffs'),
      ...selectedIn(selection, 'player-buffs'),
      ...selectedIn(selection, 'world-buffs'),
      ...selectedIn(selection, 'debuffs'),
    ];
    expect(new Set(grouped)).toEqual(new Set(RAID_BUFFS));
  });
});

describe('DURATIONS and durationLabel', () => {
  it('runs from twenty seconds to ten minutes', () => {
    expect(DURATIONS[0]).toBe(MIN_DURATION_SEC);
    expect(MIN_DURATION_SEC).toBe(20);
    expect(DURATIONS.at(-1)).toBe(MAX_DURATION_SEC);
    expect(MAX_DURATION_SEC).toBe(600);
    expect(DURATIONS).toContain(180);
    expect(DURATIONS.slice(0, 4)).toEqual([20, 30, 45, 60]);
    // Past a minute the step is thirty seconds all the way to ten minutes.
    const past = DURATIONS.slice(3);
    expect(past.every((seconds, i) => i === 0 || seconds - past[i - 1] === 30)).toBe(true);
  });

  it('reads as a clock, not as seconds', () => {
    expect(durationLabel(20)).toBe('0:20');
    expect(durationLabel(180)).toBe('3:00');
    expect(durationLabel(600)).toBe('10:00');
  });
});

describe('the setters never mutate and always clamp', () => {
  it('keeps duration inside twenty seconds to ten minutes', () => {
    const base = defaultSettings(PHYSICAL);
    expect(withDuration(base, 5).encounter.duration_sec).toBe(MIN_DURATION_SEC);
    expect(withDuration(base, 9999).encounter.duration_sec).toBe(MAX_DURATION_SEC);
    expect(base.encounter.duration_sec).toBe(180);
  });

  it('keeps targets between one and ten', () => {
    const base = defaultSettings(PHYSICAL);
    expect(withTargets(base, 0).encounter.targets).toBe(1);
    expect(withTargets(base, 99).encounter.targets).toBe(MAX_TARGETS);
  });

  it('keeps variation between none and thirty per cent, in five-point steps', () => {
    const base = defaultSettings(PHYSICAL);
    expect(VARIATIONS).toEqual([0, 0.05, 0.1, 0.15, 0.2, 0.25, 0.3]);
    expect(MAX_VARIATION).toBe(0.3);
    expect(withVariation(base, -1).encounter.variation).toBe(0);
    expect(withVariation(base, 5).encounter.variation).toBe(MAX_VARIATION);
    expect(withVariation(base, 0.1).encounter.variation).toBe(0.1);
  });

  it('keeps the target level between sixty and sixty-three', () => {
    const base = defaultSettings(PHYSICAL);
    expect(TARGET_LEVELS).toEqual([60, 61, 62, 63]);
    expect(withTargetLevel(base, 42).encounter.target_level).toBe(60);
    expect(withTargetLevel(base, 99).encounter.target_level).toBe(63);
    expect(withTargetLevel(base, 61).encounter.target_level).toBe(61);
  });

  it('keeps target armor non-negative and bounded, and zero means the engine’s preset', () => {
    const base = defaultSettings(PHYSICAL);
    expect(withTargetArmor(base, -10).encounter.target_armor).toBe(0);
    expect(withTargetArmor(base, 999999).encounter.target_armor).toBe(MAX_TARGET_ARMOR);
    expect(withTargetArmor(base, 3731).encounter.target_armor).toBe(3731);
  });

  it('publishes contract A8’s armor preset for each level, so the control can name the figure', () => {
    expect(TARGET_ARMOR_BY_LEVEL).toEqual({ 60: 3300, 61: 3444, 62: 3588, 63: 3731 });
    expect(TARGET_LEVELS.every((level) => TARGET_ARMOR_BY_LEVEL[level] > 0)).toBe(true);
  });

  // tank MAJOR, review.md:227-229: armor 0 has to DISPLAY as the preset it silently means,
  // not as a blank field indistinguishable from a typed 0. `targetArmorField` is the pure
  // decision SettingsSheet.svelte renders its input value off of; the wire value (0) is
  // untouched -- only what the field shows changes.
  describe('targetArmorField', () => {
    it('shows the level 63 preset when armor is 0', () => {
      expect(targetArmorField({ ...DEFAULT_ENCOUNTER, target_level: 63, target_armor: 0 })).toEqual({
        value: '3731',
        preset: 3731,
        level: 63,
      });
    });

    it('shows the level 60 preset when armor is 0', () => {
      expect(targetArmorField({ ...DEFAULT_ENCOUNTER, target_level: 60, target_armor: 0 })).toEqual({
        value: '3300',
        preset: 3300,
        level: 60,
      });
    });

    it('shows an overridden armor value as itself, alongside the level’s own preset', () => {
      expect(targetArmorField({ ...DEFAULT_ENCOUNTER, target_level: 63, target_armor: 5000 })).toEqual({
        value: '5000',
        preset: 3731,
        level: 63,
      });
    });

    it('falls back to level 63 when target_level is absent', () => {
      const { target_level: _targetLevel, ...withoutLevel } = {
        ...DEFAULT_ENCOUNTER,
        target_armor: 0,
      };
      expect(targetArmorField(withoutLevel)).toEqual({ value: '3731', preset: 3731, level: 63 });
    });

    /**
     * Fix round, Minor 1: `level` rides on `targetArmorField`'s own return value now,
     * so SettingsSheet.svelte reads the SAME level `preset` was computed from, rather than
     * a second `settings.encounter.target_level ?? DEFAULT_TARGET_LEVEL` lookup that only
     * agreed with it by coincidence -- the same "two lookups agreeing by accident" shape an
     * earlier fix round already removed for `preset` itself.
     */
    it('returns the level preset was computed from, not a second independent lookup', () => {
      expect(targetArmorField({ ...DEFAULT_ENCOUNTER, target_level: 60, target_armor: 0 }).level).toBe(60);
    });
  });

  it('accepts only the contract’s target types, and the empty string for “any”', () => {
    const base = defaultSettings(PHYSICAL);
    expect(TARGET_TYPES).toContain('undead');
    expect(withTargetType(base, 'undead').encounter.target_type).toBe('undead');
    expect(withTargetType(base, 'gnome').encounter.target_type).toBe('');
    expect(withTargetType(base, '').encounter.target_type).toBe('');
  });

  it('turns the dummy on and off', () => {
    const base = defaultSettings(PHYSICAL);
    expect(withDummy(base, true).encounter.dummy).toBe(true);
    expect(withDummy(withDummy(base, true), false).encounter.dummy).toBe(false);
  });

  it('turns the execute phase off by zeroing the ratio, and back on to the default', () => {
    const base = defaultSettings(PHYSICAL);
    const off = withExecutePhase(base, false);
    expect(off.encounter.execute_ratio).toBe(0);
    expect(executePhaseOn(off)).toBe(false);
    expect(withExecutePhase(off, true).encounter.execute_ratio).toBe(0.25);
  });

  it('swaps the whole buff and consumable list with the preset, keyed off reference_stat, and leaves it alone for custom', () => {
    const base = defaultSettings(PHYSICAL);
    const solo = withPreset(base, 'solo', PHYSICAL);
    expect(solo.buffs).toEqual([]);
    expect(solo.consumables).toEqual([]);
    const caster = withPreset(base, 'raid-buffed', CASTER);
    expect(caster.consumables).toEqual([...CASTER_CONSUMABLES]);
    expect(caster.buffs).toEqual([...RAID_BUFFS]);
    const custom = withPreset(base, 'custom', PHYSICAL);
    expect(custom.buffs).toEqual(base.buffs);
    expect(custom.consumables).toEqual(base.consumables);
    expect(custom.preset).toBe('custom');
  });

  it('offers exactly the three presets the design names', () => {
    expect(BUFF_PRESETS.map((row) => row.id)).toEqual(['raid-buffed', 'solo', 'custom']);
  });
});

/**
 * "Custom opens pre-ticked with the current preset" (task-4-brief.md): switching to
 * Custom never touches the lists, and switching back to a static preset restores it --
 * both directions a store ever takes through this one function.
 */
describe('withPreset(..., "custom")', () => {
  it('carries every tick across rather than resetting to empty', () => {
    const raid = defaultSettings(PHYSICAL);
    const custom = withPreset(raid, 'custom', PHYSICAL);
    expect(custom.buffs).toEqual(raid.buffs);
    expect(custom.buffs.length).toBeGreaterThan(0);
    expect(custom.consumables).toEqual(raid.consumables);
    expect(custom.consumables.length).toBeGreaterThan(0);
  });

  it('leaving custom and returning to raid-buffed restores the preset exactly', () => {
    const raid = defaultSettings(PHYSICAL);
    const custom = withPreset(raid, 'custom', PHYSICAL);
    const backToRaid = withPreset(custom, 'raid-buffed', PHYSICAL);
    expect(backToRaid.buffs).toEqual(raid.buffs);
    expect(backToRaid.consumables).toEqual(raid.consumables);
  });
});

/**
 * The store's own re-apply rule (task-4-brief.md's ambiguity, resolved): a preset is
 * spec-dependent now, so `withSpecForPreset` is what a store calls after adopting a new
 * character -- never a raw `defaultSettings()` again, which would also reset the encounter,
 * the title and everything else the player already touched.
 */
describe('withSpecForPreset', () => {
  it('re-applies the current preset for the new spec’s reference_stat', () => {
    // warrior-fury has reference_stat attack_power (specs.ts); mage-frost spell_power.
    const forWarrior = defaultSettings(PHYSICAL);
    expect(forWarrior.consumables).toEqual([...PHYSICAL_CONSUMABLES]);
    const forMage = withSpecForPreset(forWarrior, 'mage-frost');
    expect(forMage.consumables).toEqual([...CASTER_CONSUMABLES]);
    expect(forMage.preset).toBe('raid-buffed');
    const backToWarrior = withSpecForPreset(forMage, 'warrior-fury');
    expect(backToWarrior.consumables).toEqual([...PHYSICAL_CONSUMABLES]);
  });

  it('never touches a custom selection -- a mage-to-warrior switch keeps the player’s own ticks', () => {
    const custom = withPreset(defaultSettings(PHYSICAL), 'custom', PHYSICAL);
    const afterSwitch = withSpecForPreset(custom, 'mage-frost');
    expect(afterSwitch).toEqual(custom);
  });

  it('falls back to the physical set for an unknown spec rather than an empty preset', () => {
    const forMage = withSpecForPreset(defaultSettings(CASTER), 'mage-frost');
    expect(forMage.consumables).toEqual([...CASTER_CONSUMABLES]);
    const forUnknown = withSpecForPreset(forMage, 'not-a-real-spec');
    expect(forUnknown.consumables).toEqual([...PHYSICAL_CONSUMABLES]);
  });
});

describe('the setters guard against non-finite input', () => {
  // An emptied number input parses to NaN (or, briefly, Infinity); clamping that down to
  // the minimum would pick a value nobody asked for, and JSON.stringify(NaN) is "null" on
  // a wire field the contract requires as a plain number. Either way the setting has to
  // stay where it was, not snap anywhere.
  it('leaves duration where it was rather than producing NaN', () => {
    const moved = withDuration(defaultSettings(PHYSICAL), 300);
    expect(withDuration(moved, NaN).encounter.duration_sec).toBe(300);
    expect(withDuration(moved, Infinity).encounter.duration_sec).toBe(300);
  });

  it('leaves targets where they were rather than producing NaN', () => {
    const moved = withTargets(defaultSettings(PHYSICAL), 4);
    expect(withTargets(moved, NaN).encounter.targets).toBe(4);
    expect(withTargets(moved, Infinity).encounter.targets).toBe(4);
  });

  it('leaves variation where it was rather than producing NaN', () => {
    const moved = withVariation(defaultSettings(PHYSICAL), 0.1);
    expect(withVariation(moved, NaN).encounter.variation).toBe(0.1);
    expect(withVariation(moved, Infinity).encounter.variation).toBe(0.1);
  });

  it('leaves the target level where it was rather than producing NaN', () => {
    const moved = withTargetLevel(defaultSettings(PHYSICAL), 61);
    expect(withTargetLevel(moved, NaN).encounter.target_level).toBe(61);
    expect(withTargetLevel(moved, Infinity).encounter.target_level).toBe(61);
  });

  it('leaves target armor where it was rather than producing NaN', () => {
    const moved = withTargetArmor(defaultSettings(PHYSICAL), 3731);
    expect(withTargetArmor(moved, NaN).encounter.target_armor).toBe(3731);
    expect(withTargetArmor(moved, Infinity).encounter.target_armor).toBe(3731);
  });

  it('never lets a non-finite value reach the wire as null', () => {
    const settings = withDuration(defaultSettings(PHYSICAL), NaN);
    const wire = JSON.parse(JSON.stringify(settings.encounter)) as Record<string, unknown>;
    expect(wire.duration_sec).toBe(180);
  });
});

describe('the style and the controls beside it', () => {
  it('writes the style’s fields through applyFightStyle', () => {
    const cleave = withStyle(defaultSettings(PHYSICAL), 'cleave-5');
    expect(cleave.encounter.targets).toBe(5);
    expect(styleIdOf(cleave)).toBe('cleave-5');
  });

  it('detaches from the style the moment a style-owned field is set by hand', () => {
    const cleave = withStyle(defaultSettings(PHYSICAL), 'cleave-3');
    expect(styleIdOf(withTargets(cleave, 4))).toBe('');
    expect(styleIdOf(withExecutePhase(cleave, false))).toBe('');
    expect(styleIdOf(withDummy(cleave, true))).toBe('');
    // Fight length, variation, target level, armor and type are the player's, not the
    // style's: changing one keeps the style's name on the run.
    expect(styleIdOf(withDuration(cleave, 300))).toBe('cleave-3');
    expect(styleIdOf(withVariation(cleave, 0))).toBe('cleave-3');
    expect(styleIdOf(withTargetLevel(cleave, 60))).toBe('cleave-3');
    expect(styleIdOf(withTargetArmor(cleave, 2000))).toBe('cleave-3');
    expect(styleIdOf(withTargetType(cleave, 'undead'))).toBe('cleave-3');
  });

  it('names the detached state', () => {
    expect(simCopy.styleCustom).toBeTruthy();
  });

  // Fix round 1 (reviewer Important): the read-only TARGETS control keys off a non-empty
  // `targets_over_time`, not off `style`, precisely because `detached()` must NOT clear the
  // ramp when a style-owned checkbox is toggled -- doing so would silently drop the
  // dungeon pull's 1->5 ramp from the run the moment a player ticks "dummy" or flips
  // execute phase. This pins that: the ramp survives detachment.
  it('leaves the dungeon pull’s target-count timeline intact when a style-owned field detaches the encounter by hand', () => {
    const dungeon = withStyle(defaultSettings(PHYSICAL), 'dungeon');
    const detached = withDummy(dungeon, true);
    expect(styleIdOf(detached)).toBe('');
    expect(detached.encounter.targets_over_time).toEqual(dungeon.encounter.targets_over_time);
  });
});

describe('settingsLabel', () => {
  it('is the one line a saved sim is titled with, unchanged by the new controls', () => {
    expect(settingsLabel(defaultSettings(PHYSICAL))).toBe('Raid-buffed, 3:00, single target');
    expect(settingsLabel(withTargets(withPreset(defaultSettings(PHYSICAL), 'solo', PHYSICAL), 4))).toBe(
      'Solo, 3:00, 4 targets',
    );
  });
});
