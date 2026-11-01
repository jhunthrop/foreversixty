import { readFileSync } from 'node:fs';
import path from 'node:path';
import { fileURLToPath } from 'node:url';
import { describe, expect, it } from 'vitest';
import generated from '../../data/generated/sim-ids.json';
import { selectedIn } from './buffs';
import { simCopy } from './copy';
import {
  BUFF_PRESETS,
  DEFAULT_TARGET_LEVEL,
  defaultSettings,
  defaultTargetLevel,
  durationLabel,
  DURATIONS,
  encounterFor,
  executePhaseOn,
  MAX_DURATION_SEC,
  MAX_TARGET_ARMOR,
  MAX_TARGETS,
  MAX_VARIATION,
  MIN_DURATION_SEC,
  PRESET_BUFFS,
  presetConsumables,
  RAID_BUFFS,
  settingsLabel,
  styleIdOf,
  TARGET_ARMOR_BY_LEVEL,
  TARGET_LEVELS,
  TARGET_TYPES,
  targetArmorField,
  targetLevelFor,
  targetTypeShown,
  VARIATIONS,
  withDummy,
  withDuration,
  withExecutePhase,
  withPreset,
  withSpecForPreset,
  withStyle,
  withTargetArmor,
  withTargetLevel,
  withTargets,
  withTargetType,
  withVariation,
} from './settings';
import { DEFAULT_ENCOUNTER } from './types';

const CURATED_PRESETS = path.resolve(
  path.dirname(fileURLToPath(import.meta.url)),
  '../../../../data/curated/presets.json',
);
const RAID_CONSUMES = path.resolve(
  path.dirname(fileURLToPath(import.meta.url)),
  '../../../../sim/request/raid-consumes.json',
);
const FURY = 'warrior-fury';
const FIRE = 'mage-fire';

const MAX = 60;
const LADDER_BANDS = [20, 30, 40, 50, 60] as const;

type LevelRows = { from_level: number; consumes: string[] }[];

/** The raid set per spec and level, read straight from the Go-generated file (whose rows
 *  sim/request's TestRaidConsumesAreThePublishedSetAtEveryLevel holds to the published setup). */
function publishedRows(): Record<string, LevelRows> {
  return (JSON.parse(readFileSync(RAID_CONSUMES, 'utf8')) as { specs: Record<string, LevelRows> }).specs;
}

function publishedAt(spec: string, level: number): string[] {
  return publishedRows()
    [spec]!.filter((row) => row.from_level <= level)
    .at(-1)!.consumes;
}

describe('defaultSettings', () => {
  it("is the contract's EncounterSpec defaults, raid-buffed, on Patchwerk", () => {
    const settings = defaultSettings(FURY, MAX);
    expect(settings.encounter).toEqual({
      duration_sec: 180,
      variation: 0.2,
      targets: 1,
      execute_ratio: 0.25,
      profile: '',
      style: 'patchwerk',
      target_level: 63,
      // Absent, not 0: the default fight carries no armor override at all (Defect 3).
      // The published numbers' target (sim/leveling RankerEncounter): no creature type, so
      // a rogue's Murder or a slaying enchant reads the same here as on the tier list.
      target_type: 'unknown',
      dummy: false,
    });
    expect(settings.encounter.target_armor).toBeUndefined();
    expect(settings.preset).toBe('raid-buffed');
    expect(settings.buffs.length).toBeGreaterThan(0);
    expect(settings.cooldowns).toEqual([]);
  });

  it('carries each spec’s published level-60 raid consumables (round 5 ruling 3), so page and published DPS agree', () => {
    const specs = Object.keys(publishedRows());
    expect(specs.length).toBeGreaterThanOrEqual(27);
    for (const spec of specs) {
      for (const level of LADDER_BANDS) {
        expect(defaultSettings(spec, level).consumables, `${spec} ${level}`).toEqual(
          publishedAt(spec, level),
        );
      }
    }
  });

  it('follows the kit’s level gates: Instant Poisons from 20, Deadly from 30, Rockbiter before Windfury', () => {
    expect(defaultSettings('rogue-combat', 25).consumables).toContain('main_hand_imbue:instant_poison');
    expect(defaultSettings('rogue-combat', 25).consumables).not.toContain('main_hand_imbue:deadly_poison');
    expect(defaultSettings('rogue-combat', 19).consumables.some((id) => id.endsWith('_poison'))).toBe(false);
    expect(defaultSettings('shaman-enhancement', 29).consumables).toContain(
      'main_hand_imbue:rockbiter_weapon',
    );
    expect(defaultSettings('shaman-enhancement', 30).consumables).toContain(
      'main_hand_imbue:windfury_weapon',
    );
  });

  it('carries only consumables the character’s level can use (the client’s RequiredLevel)', () => {
    const fury40 = defaultSettings('warrior-fury', 40).consumables;
    for (const id of [
      'elixir_of_the_mongoose',
      'mighty_rage_potion',
      'main_hand_imbue:elemental_sharpening_stone',
    ]) {
      expect(fury40).not.toContain(id);
    }
    expect(defaultSettings('warrior-fury', 60).consumables).toContain('elixir_of_the_mongoose');
  });

  it('gives rogues their poisons and Survival its ranged set, not a shared physical list', () => {
    const combat = defaultSettings('rogue-combat', MAX).consumables;
    expect(combat).toContain('main_hand_imbue:deadly_poison');
    expect(combat).toContain('off_hand_imbue:instant_poison');
    expect(combat).not.toContain('main_hand_imbue:elemental_sharpening_stone');
    const survival = defaultSettings('hunter-survival', MAX).consumables;
    expect(survival).not.toContain('flask_of_the_titans');
    expect(survival).not.toContain('mighty_rage_potion');
  });

  it('carries no consumables before a character, or for a spec the published set does not know', () => {
    expect(defaultSettings().consumables).toEqual([]);
    expect(defaultSettings('not-a-real-spec').consumables).toEqual([]);
  });

  it('carries Power Infusion in no default, for any spec or preset (round 5 ruling 2)', () => {
    for (const spec of [...Object.keys(publishedRows()), '']) {
      for (const preset of BUFF_PRESETS.map((row) => row.id)) {
        const settings = withPreset(defaultSettings(spec, MAX), preset, spec, MAX);
        expect([...settings.buffs, ...settings.consumables], `${spec} ${preset}`).not.toContain(
          'power_infusions',
        );
      }
    }
    expect(RAID_BUFFS).not.toContain('power_infusions');
  });

  it('still offers Power Infusion as a player buff the player can tick', () => {
    expect(
      generated.buffs.some((row) => row.id === 'power_infusions' && row.message === 'IndividualBuffs'),
    ).toBe(true);
  });
});

/**
 * Task 4 (dps-minmaxer D16 BLOCKER): the old preset was five buffs and two consumables,
 * +6% over Solo. This is the full standard set instead -- the brief's 54 buff ids less the
 * eight world buffs Forever does not have (2026-09-28), 46, copied verbatim from
 * sim/adapter/testdata/warrior-fury.request.json, plus the 9 physical and 7
 * caster consumable ids the controller verified against the real engine (8b2169e61): every
 * one of the 70 accepted, no ErrUnknownBuff, no ErrUnknownConsume, no ErrAmbiguousConsume.
 */
describe('the Raid-buffed preset', () => {
  it('is the curated raid preset’s buffs and debuffs, the list every published level-60 number is simmed under, never split by spec', () => {
    const curated = JSON.parse(readFileSync(CURATED_PRESETS, 'utf8')) as {
      presets: { raid: { buffs: { id: string }[]; debuffs: { id: string }[] } };
    };
    const raid = curated.presets.raid;
    expect([...RAID_BUFFS]).toEqual([...raid.buffs.map((b) => b.id), ...raid.debuffs.map((d) => d.id)]);
    expect(PRESET_BUFFS['raid-buffed']).toBe(RAID_BUFFS);
  });

  it('carries Heart of the Lion and none of the buffs the raid preset refuses (round 3 C3)', () => {
    expect(RAID_BUFFS).toContain('heart_of_the_lion');
    for (const refused of ['blessing_of_kings', 'sanctity_aura', 'winters_chill', 'shadow_weaving']) {
      expect(RAID_BUFFS, refused).not.toContain(refused);
    }
  });

  it('reads its consumable half per spec from the published set', () => {
    expect(presetConsumables('raid-buffed', FURY, MAX)).toEqual(publishedAt(FURY, MAX));
    expect(presetConsumables('raid-buffed', FIRE, MAX)).toEqual(publishedAt(FIRE, MAX));
  });

  it('carries every buff and consumable id in the synced catalogue -- an id the engine cannot map is ErrUnknownBuff/ErrUnknownConsume and fails the whole run', () => {
    const published = new Set([
      ...generated.buffs.map((row) => row.id),
      ...generated.consumables.map((row) => row.id),
    ]);
    for (const id of [
      ...RAID_BUFFS,
      ...Object.values(publishedRows()).flatMap((rows) => rows.flatMap((row) => row.consumes)),
    ]) {
      expect(published.has(id), id).toBe(true);
    }
  });

  it('solo stays self-only', () => {
    expect(PRESET_BUFFS.solo).toEqual([]);
    expect(presetConsumables('solo', FURY, MAX)).toEqual([]);
    expect(presetConsumables('solo', FIRE, MAX)).toEqual([]);
  });

  /**
   * The findings' own counters (task-4-brief.md): "RAID BUFFS 4/30, ... ON THE TARGET
   * 0/24, ... WEAPON OILS AND STONES 0/48" -- the group counts `selectedIn` (buffs.ts)
   * reports for the Custom panel, over the settings this preset now writes. This is no
   * longer 4, 0 and 0: the fix landed in the counters the panel actually renders, not only
   * in a list nobody reads a length off.
   */
  it('reports real totals in the group counters the findings quoted as 4/30, 0/24 and 0/48', () => {
    const settings = defaultSettings(FURY, MAX);
    const selection = { buffs: settings.buffs, consumables: settings.consumables };
    expect(selectedIn(selection, 'raid-buffs').length).toBeGreaterThan(4);
    expect(selectedIn(selection, 'debuffs').length).toBeGreaterThan(0);
    expect(selectedIn(selection, 'weapon-imbue').length).toBeGreaterThan(0);
    // No buff id is lost between groups: every group's members sum back to the full list.
    const grouped = [
      ...selectedIn(selection, 'raid-buffs'),
      ...selectedIn(selection, 'party-buffs'),
      ...selectedIn(selection, 'player-buffs'),
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
    const base = defaultSettings(FURY, MAX);
    expect(withDuration(base, 5).encounter.duration_sec).toBe(MIN_DURATION_SEC);
    expect(withDuration(base, 9999).encounter.duration_sec).toBe(MAX_DURATION_SEC);
    expect(base.encounter.duration_sec).toBe(180);
  });

  it('keeps targets between one and ten', () => {
    const base = defaultSettings(FURY, MAX);
    expect(withTargets(base, 0).encounter.targets).toBe(1);
    expect(withTargets(base, 99).encounter.targets).toBe(MAX_TARGETS);
  });

  it('keeps variation between none and thirty per cent, in five-point steps', () => {
    const base = defaultSettings(FURY, MAX);
    expect(VARIATIONS).toEqual([0, 0.05, 0.1, 0.15, 0.2, 0.25, 0.3]);
    expect(MAX_VARIATION).toBe(0.3);
    expect(withVariation(base, -1).encounter.variation).toBe(0);
    expect(withVariation(base, 5).encounter.variation).toBe(MAX_VARIATION);
    expect(withVariation(base, 0.1).encounter.variation).toBe(0.1);
  });

  it('keeps the target level between sixty and sixty-three', () => {
    const base = defaultSettings(FURY, MAX);
    expect(TARGET_LEVELS).toEqual([60, 61, 62, 63]);
    expect(withTargetLevel(base, 42).encounter.target_level).toBe(60);
    expect(withTargetLevel(base, 99).encounter.target_level).toBe(63);
    expect(withTargetLevel(base, 61).encounter.target_level).toBe(61);
  });

  // 2026-09-21 result-page review, Defect 3: blank must mean "the level's preset" and a
  // typed 0 must mean zero armor, all the way through -- `null` is `withTargetArmor`'s own
  // "clear the override" (settings.ts's own doc comment); a number, including 0, is always
  // an explicit override, never silently swapped for the preset.
  it('clears the override to the level’s preset on null, and keeps a numeric override non-negative and bounded', () => {
    const base = defaultSettings(FURY, MAX);
    const overridden = withTargetArmor(base, 3731);
    expect(overridden.encounter.target_armor).toBe(3731);
    expect(withTargetArmor(overridden, null).encounter.target_armor).toBeUndefined();
    expect(withTargetArmor(base, -10).encounter.target_armor).toBe(0);
    expect(withTargetArmor(base, 999999).encounter.target_armor).toBe(MAX_TARGET_ARMOR);
    expect(withTargetArmor(base, 3731).encounter.target_armor).toBe(3731);
  });

  it('an explicit 0 is zero armor, not the preset', () => {
    const base = defaultSettings(FURY, MAX);
    expect(withTargetArmor(base, 0).encounter.target_armor).toBe(0);
  });

  it('publishes contract A8’s armor preset for each level, so the control can name the figure', () => {
    expect(TARGET_ARMOR_BY_LEVEL[60]).toBe(3300);
    expect(TARGET_ARMOR_BY_LEVEL[61]).toBe(3444);
    expect(TARGET_ARMOR_BY_LEVEL[62]).toBe(3588);
    expect(TARGET_ARMOR_BY_LEVEL[63]).toBe(3731);
    expect(TARGET_LEVELS.every((level) => TARGET_ARMOR_BY_LEVEL[level] > 0)).toBe(true);
  });

  // Level-aware sim design (2026-09-27), design step 5: extended down through level 1 for
  // a levelling character's target (character level + 3). Below 60 the figure is the one the
  // engine runs (sim/api envelope.go armorBelowMaxLevel: the level-60 preset scaled by
  // level), so the field never names an armor the run does not use.
  it('extends the armor table down through level 1 with the engine’s own below-60 rule', () => {
    expect(Object.keys(TARGET_ARMOR_BY_LEVEL)).toHaveLength(63);
    expect(TARGET_ARMOR_BY_LEVEL[41]).toBe(2255);
    expect(TARGET_ARMOR_BY_LEVEL[36]).toBe(1980);
    expect(TARGET_ARMOR_BY_LEVEL[1]).toBe(55);
    for (let level = 1; level < 63; level += 1) {
      expect(TARGET_ARMOR_BY_LEVEL[level + 1]).toBeGreaterThanOrEqual(TARGET_ARMOR_BY_LEVEL[level]);
    }
  });

  it('names the armor for the target level the character actually fights', () => {
    // A level-33 character's default target is 36, not the level-63 boss.
    expect(targetArmorField(DEFAULT_ENCOUNTER, 33)).toMatchObject({ level: 36, preset: 1980 });
  });

  describe('defaultTargetLevel', () => {
    it('is 3 above the character’s own level', () => {
      expect(defaultTargetLevel(1)).toBe(4);
      expect(defaultTargetLevel(30)).toBe(33);
      expect(defaultTargetLevel(60)).toBe(63);
    });

    it('is what DEFAULT_TARGET_LEVEL already was, applied to MAX_LEVEL', () => {
      expect(DEFAULT_TARGET_LEVEL).toBe(defaultTargetLevel(60));
      expect(DEFAULT_TARGET_LEVEL).toBe(63);
    });
  });

  // 2026-09-21 result-page review, Defect 3: the field must read BLANK when there is no
  // override (its `placeholder`, SettingsSheet.svelte's own concern, shows the preset as a
  // hint), and read exactly the override -- including "0" -- when there is one. `0` must
  // never again be swapped for the preset text: that collapsed "leave it blank" and "type
  // 0" into the identical request.
  describe('targetArmorField', () => {
    it('reads blank, with the level 63 preset available for a hint, when there is no override', () => {
      const { target_armor: _armor, ...withoutOverride } = { ...DEFAULT_ENCOUNTER, target_level: 63 };
      expect(targetArmorField(withoutOverride)).toEqual({ value: '', preset: 3731, level: 63 });
    });

    it('reads blank with the level 60 preset when there is no override', () => {
      const { target_armor: _armor, ...withoutOverride } = { ...DEFAULT_ENCOUNTER, target_level: 60 };
      expect(targetArmorField(withoutOverride)).toEqual({ value: '', preset: 3300, level: 60 });
    });

    it('reads an explicit 0 as "0", never as the preset', () => {
      expect(targetArmorField({ ...DEFAULT_ENCOUNTER, target_level: 63, target_armor: 0 })).toEqual({
        value: '0',
        preset: 3731,
        level: 63,
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
      const { target_level: _targetLevel, target_armor: _armor, ...withoutLevel } = { ...DEFAULT_ENCOUNTER };
      expect(targetArmorField(withoutLevel)).toEqual({ value: '', preset: 3731, level: 63 });
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

  it('accepts only the contract’s target types; anything else is the default, no creature type', () => {
    const base = defaultSettings(FURY, MAX);
    expect(TARGET_TYPES).toContain('undead');
    expect(withTargetType(base, 'undead').encounter.target_type).toBe('undead');
    expect(withTargetType(base, 'gnome').encounter.target_type).toBe('unknown');
    expect(withTargetType(base, '').encounter.target_type).toBe('unknown');
  });

  it('shows an empty stored type as the humanoid the engine fights for it (sim/request request.go)', () => {
    expect(targetTypeShown({ ...DEFAULT_ENCOUNTER, target_type: '' })).toBe('humanoid');
    expect(targetTypeShown({ ...DEFAULT_ENCOUNTER, target_type: undefined })).toBe('humanoid');
    expect(targetTypeShown({ ...DEFAULT_ENCOUNTER, target_type: 'beast' })).toBe('beast');
    expect(targetTypeShown(DEFAULT_ENCOUNTER)).toBe('unknown');
  });

  it('turns the dummy on and off', () => {
    const base = defaultSettings(FURY, MAX);
    expect(withDummy(base, true).encounter.dummy).toBe(true);
    expect(withDummy(withDummy(base, true), false).encounter.dummy).toBe(false);
  });

  it('turns the execute phase off by zeroing the ratio, and back on to the default', () => {
    const base = defaultSettings(FURY, MAX);
    const off = withExecutePhase(base, false);
    expect(off.encounter.execute_ratio).toBe(0);
    expect(executePhaseOn(off)).toBe(false);
    expect(withExecutePhase(off, true).encounter.execute_ratio).toBe(0.25);
  });

  it('swaps the whole buff and consumable list with the preset, keyed off the spec, and leaves it alone for custom', () => {
    const base = defaultSettings(FURY, MAX);
    const solo = withPreset(base, 'solo', FURY, MAX);
    expect(solo.buffs).toEqual([]);
    expect(solo.consumables).toEqual([]);
    const caster = withPreset(base, 'raid-buffed', FIRE, MAX);
    expect(caster.consumables).toEqual(publishedAt(FIRE, MAX));
    expect(caster.buffs).toEqual([...RAID_BUFFS]);
    const custom = withPreset(base, 'custom', FURY, MAX);
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
    const raid = defaultSettings(FURY, MAX);
    const custom = withPreset(raid, 'custom', FURY, MAX);
    expect(custom.buffs).toEqual(raid.buffs);
    expect(custom.buffs.length).toBeGreaterThan(0);
    expect(custom.consumables).toEqual(raid.consumables);
    expect(custom.consumables.length).toBeGreaterThan(0);
  });

  it('leaving custom and returning to raid-buffed restores the preset exactly', () => {
    const raid = defaultSettings(FURY, MAX);
    const custom = withPreset(raid, 'custom', FURY, MAX);
    const backToRaid = withPreset(custom, 'raid-buffed', FURY, MAX);
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
  it('re-applies the current preset for the new spec', () => {
    const forWarrior = defaultSettings(FURY, MAX);
    expect(forWarrior.consumables).toEqual(publishedAt(FURY, MAX));
    const forMage = withSpecForPreset(forWarrior, 'mage-frost', MAX);
    expect(forMage.consumables).toEqual(publishedAt('mage-frost', MAX));
    expect(forMage.preset).toBe('raid-buffed');
    const forRogue = withSpecForPreset(forMage, 'rogue-combat', 25);
    expect(forRogue.consumables).toEqual(publishedAt('rogue-combat', 25));
  });

  it('never touches a custom selection -- a mage-to-warrior switch keeps the player’s own ticks', () => {
    const custom = withPreset(defaultSettings(FURY, MAX), 'custom', FURY, MAX);
    const afterSwitch = withSpecForPreset(custom, 'mage-frost', MAX);
    expect(afterSwitch).toEqual(custom);
  });

  it('carries no consumables for an unknown spec rather than another spec’s set', () => {
    const forMage = withSpecForPreset(defaultSettings(FIRE, MAX), 'mage-frost', MAX);
    expect(withSpecForPreset(forMage, 'not-a-real-spec', MAX).consumables).toEqual([]);
  });
});

describe('the setters guard against non-finite input', () => {
  // An emptied number input parses to NaN (or, briefly, Infinity); clamping that down to
  // the minimum would pick a value nobody asked for, and JSON.stringify(NaN) is "null" on
  // a wire field the contract requires as a plain number. Either way the setting has to
  // stay where it was, not snap anywhere.
  it('leaves duration where it was rather than producing NaN', () => {
    const moved = withDuration(defaultSettings(FURY, MAX), 300);
    expect(withDuration(moved, NaN).encounter.duration_sec).toBe(300);
    expect(withDuration(moved, Infinity).encounter.duration_sec).toBe(300);
  });

  it('leaves targets where they were rather than producing NaN', () => {
    const moved = withTargets(defaultSettings(FURY, MAX), 4);
    expect(withTargets(moved, NaN).encounter.targets).toBe(4);
    expect(withTargets(moved, Infinity).encounter.targets).toBe(4);
  });

  it('leaves variation where it was rather than producing NaN', () => {
    const moved = withVariation(defaultSettings(FURY, MAX), 0.1);
    expect(withVariation(moved, NaN).encounter.variation).toBe(0.1);
    expect(withVariation(moved, Infinity).encounter.variation).toBe(0.1);
  });

  it('leaves the target level where it was rather than producing NaN', () => {
    const moved = withTargetLevel(defaultSettings(FURY, MAX), 61);
    expect(withTargetLevel(moved, NaN).encounter.target_level).toBe(61);
    expect(withTargetLevel(moved, Infinity).encounter.target_level).toBe(61);
  });

  it('leaves target armor where it was rather than producing NaN', () => {
    const moved = withTargetArmor(defaultSettings(FURY, MAX), 3731);
    expect(withTargetArmor(moved, NaN).encounter.target_armor).toBe(3731);
    expect(withTargetArmor(moved, Infinity).encounter.target_armor).toBe(3731);
  });

  it('never lets a non-finite value reach the wire as null', () => {
    const settings = withDuration(defaultSettings(FURY, MAX), NaN);
    const wire = JSON.parse(JSON.stringify(settings.encounter)) as Record<string, unknown>;
    expect(wire.duration_sec).toBe(180);
  });
});

describe('the style and the controls beside it', () => {
  it('writes the style’s fields through applyFightStyle', () => {
    const cleave = withStyle(defaultSettings(FURY, MAX), 'cleave-5');
    expect(cleave.encounter.targets).toBe(5);
    expect(styleIdOf(cleave)).toBe('cleave-5');
  });

  it('detaches from the style the moment a style-owned field is set by hand', () => {
    const cleave = withStyle(defaultSettings(FURY, MAX), 'cleave-3');
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
    const dungeon = withStyle(defaultSettings(FURY, MAX), 'dungeon');
    const detached = withDummy(dungeon, true);
    expect(styleIdOf(detached)).toBe('');
    expect(detached.encounter.targets_over_time).toEqual(dungeon.encounter.targets_over_time);
  });
});

describe('settingsLabel', () => {
  it('is the one line a saved sim is titled with, unchanged by the new controls', () => {
    expect(settingsLabel(defaultSettings(FURY, MAX))).toBe('Raid-buffed, 3:00, single target');
    expect(settingsLabel(withTargets(withPreset(defaultSettings(FURY, MAX), 'solo', FURY, MAX), 4))).toBe(
      'Solo, 3:00, 4 targets',
    );
  });
});

describe('targetLevelFor', () => {
  it('reads a stored 60..63 as an offset above the character, so a level-15 export fights 15..18', () => {
    expect(targetLevelFor({ ...DEFAULT_ENCOUNTER, target_level: 63 }, 15)).toBe(18);
    expect(targetLevelFor({ ...DEFAULT_ENCOUNTER, target_level: 60 }, 15)).toBe(15);
    expect(targetLevelFor({ ...DEFAULT_ENCOUNTER, target_level: 61 }, 38)).toBe(39);
  });

  it('is unchanged at the cap and defaults to three above when nothing is stored', () => {
    expect(targetLevelFor({ ...DEFAULT_ENCOUNTER, target_level: 63 }, 60)).toBe(63);
    expect(targetLevelFor({ ...DEFAULT_ENCOUNTER, target_level: undefined }, 38)).toBe(41);
  });

  it('treats a stored level outside the offset range as the default', () => {
    expect(targetLevelFor({ ...DEFAULT_ENCOUNTER, target_level: 70 }, 38)).toBe(41);
    expect(targetLevelFor({ ...DEFAULT_ENCOUNTER, target_level: 12 }, 38)).toBe(41);
  });

  it('encounterFor resolves target_level and leaves the rest alone', () => {
    const resolved = encounterFor({ ...DEFAULT_ENCOUNTER, target_level: 63, targets: 2 }, 15);
    expect(resolved.target_level).toBe(18);
    expect(resolved.targets).toBe(2);
  });
});
