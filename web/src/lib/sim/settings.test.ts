import { describe, expect, it } from 'vitest';
import { simCopy } from './copy';
import {
  BUFF_PRESETS,
  DURATIONS,
  MAX_DURATION_SEC,
  MAX_TARGETS,
  MAX_TARGET_ARMOR,
  MAX_VARIATION,
  MIN_DURATION_SEC,
  TARGET_ARMOR_BY_LEVEL,
  TARGET_LEVELS,
  TARGET_TYPES,
  VARIATIONS,
  defaultSettings,
  durationLabel,
  executePhaseOn,
  settingsLabel,
  styleIdOf,
  targetArmorField,
  withDummy,
  withDuration,
  withExecutePhase,
  withPreset,
  withStyle,
  withTargetArmor,
  withTargetLevel,
  withTargetType,
  withTargets,
  withVariation,
} from './settings';
import { DEFAULT_ENCOUNTER } from './types';

describe('defaultSettings', () => {
  it("is the contract's EncounterSpec defaults, raid-buffed, on Patchwerk", () => {
    const settings = defaultSettings();
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
    const base = defaultSettings();
    expect(withDuration(base, 5).encounter.duration_sec).toBe(MIN_DURATION_SEC);
    expect(withDuration(base, 9999).encounter.duration_sec).toBe(MAX_DURATION_SEC);
    expect(base.encounter.duration_sec).toBe(180);
  });

  it('keeps targets between one and ten', () => {
    const base = defaultSettings();
    expect(withTargets(base, 0).encounter.targets).toBe(1);
    expect(withTargets(base, 99).encounter.targets).toBe(MAX_TARGETS);
  });

  it('keeps variation between none and thirty per cent, in five-point steps', () => {
    const base = defaultSettings();
    expect(VARIATIONS).toEqual([0, 0.05, 0.1, 0.15, 0.2, 0.25, 0.3]);
    expect(MAX_VARIATION).toBe(0.3);
    expect(withVariation(base, -1).encounter.variation).toBe(0);
    expect(withVariation(base, 5).encounter.variation).toBe(MAX_VARIATION);
    expect(withVariation(base, 0.1).encounter.variation).toBe(0.1);
  });

  it('keeps the target level between sixty and sixty-three', () => {
    const base = defaultSettings();
    expect(TARGET_LEVELS).toEqual([60, 61, 62, 63]);
    expect(withTargetLevel(base, 42).encounter.target_level).toBe(60);
    expect(withTargetLevel(base, 99).encounter.target_level).toBe(63);
    expect(withTargetLevel(base, 61).encounter.target_level).toBe(61);
  });

  it('keeps target armor non-negative and bounded, and zero means the engine’s preset', () => {
    const base = defaultSettings();
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
    const base = defaultSettings();
    expect(TARGET_TYPES).toContain('undead');
    expect(withTargetType(base, 'undead').encounter.target_type).toBe('undead');
    expect(withTargetType(base, 'gnome').encounter.target_type).toBe('');
    expect(withTargetType(base, '').encounter.target_type).toBe('');
  });

  it('turns the dummy on and off', () => {
    const base = defaultSettings();
    expect(withDummy(base, true).encounter.dummy).toBe(true);
    expect(withDummy(withDummy(base, true), false).encounter.dummy).toBe(false);
  });

  it('turns the execute phase off by zeroing the ratio, and back on to the default', () => {
    const base = defaultSettings();
    const off = withExecutePhase(base, false);
    expect(off.encounter.execute_ratio).toBe(0);
    expect(executePhaseOn(off)).toBe(false);
    expect(withExecutePhase(off, true).encounter.execute_ratio).toBe(0.25);
  });

  it('swaps the whole buff and consumable list with the preset, and leaves it alone for custom', () => {
    const base = defaultSettings();
    const solo = withPreset(base, 'solo');
    expect(solo.buffs).toEqual([]);
    expect(solo.consumables).toEqual([]);
    const custom = withPreset(base, 'custom');
    expect(custom.buffs).toEqual(base.buffs);
    expect(custom.preset).toBe('custom');
  });

  it('offers exactly the three presets the design names', () => {
    expect(BUFF_PRESETS.map((row) => row.id)).toEqual(['raid-buffed', 'solo', 'custom']);
  });
});

describe('the setters guard against non-finite input', () => {
  // An emptied number input parses to NaN (or, briefly, Infinity); clamping that down to
  // the minimum would pick a value nobody asked for, and JSON.stringify(NaN) is "null" on
  // a wire field the contract requires as a plain number. Either way the setting has to
  // stay where it was, not snap anywhere.
  it('leaves duration where it was rather than producing NaN', () => {
    const moved = withDuration(defaultSettings(), 300);
    expect(withDuration(moved, NaN).encounter.duration_sec).toBe(300);
    expect(withDuration(moved, Infinity).encounter.duration_sec).toBe(300);
  });

  it('leaves targets where they were rather than producing NaN', () => {
    const moved = withTargets(defaultSettings(), 4);
    expect(withTargets(moved, NaN).encounter.targets).toBe(4);
    expect(withTargets(moved, Infinity).encounter.targets).toBe(4);
  });

  it('leaves variation where it was rather than producing NaN', () => {
    const moved = withVariation(defaultSettings(), 0.1);
    expect(withVariation(moved, NaN).encounter.variation).toBe(0.1);
    expect(withVariation(moved, Infinity).encounter.variation).toBe(0.1);
  });

  it('leaves the target level where it was rather than producing NaN', () => {
    const moved = withTargetLevel(defaultSettings(), 61);
    expect(withTargetLevel(moved, NaN).encounter.target_level).toBe(61);
    expect(withTargetLevel(moved, Infinity).encounter.target_level).toBe(61);
  });

  it('leaves target armor where it was rather than producing NaN', () => {
    const moved = withTargetArmor(defaultSettings(), 3731);
    expect(withTargetArmor(moved, NaN).encounter.target_armor).toBe(3731);
    expect(withTargetArmor(moved, Infinity).encounter.target_armor).toBe(3731);
  });

  it('never lets a non-finite value reach the wire as null', () => {
    const settings = withDuration(defaultSettings(), NaN);
    const wire = JSON.parse(JSON.stringify(settings.encounter)) as Record<string, unknown>;
    expect(wire.duration_sec).toBe(180);
  });
});

describe('the style and the controls beside it', () => {
  it('writes the style’s fields through applyFightStyle', () => {
    const cleave = withStyle(defaultSettings(), 'cleave-5');
    expect(cleave.encounter.targets).toBe(5);
    expect(styleIdOf(cleave)).toBe('cleave-5');
  });

  it('detaches from the style the moment a style-owned field is set by hand', () => {
    const cleave = withStyle(defaultSettings(), 'cleave-3');
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
    const dungeon = withStyle(defaultSettings(), 'dungeon');
    const detached = withDummy(dungeon, true);
    expect(styleIdOf(detached)).toBe('');
    expect(detached.encounter.targets_over_time).toEqual(dungeon.encounter.targets_over_time);
  });
});

describe('settingsLabel', () => {
  it('is the one line a saved sim is titled with, unchanged by the new controls', () => {
    expect(settingsLabel(defaultSettings())).toBe('Raid-buffed, 3:00, single target');
    expect(settingsLabel(withTargets(withPreset(defaultSettings(), 'solo'), 4))).toBe(
      'Solo, 3:00, 4 targets',
    );
  });
});
