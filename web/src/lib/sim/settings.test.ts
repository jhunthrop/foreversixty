import { describe, expect, it } from 'vitest';
import {
  BUFF_PRESETS,
  DURATIONS,
  MAX_DURATION_SEC,
  MAX_TARGETS,
  MIN_DURATION_SEC,
  defaultSettings,
  durationLabel,
  executePhaseOn,
  settingsLabel,
  withDuration,
  withExecutePhase,
  withPreset,
  withTargets,
} from './settings';

describe('defaultSettings', () => {
  it("is the contract's EncounterSpec defaults, raid-buffed", () => {
    const settings = defaultSettings();
    expect(settings.encounter).toEqual({
      duration_sec: 180,
      variation: 0.2,
      targets: 1,
      execute_ratio: 0.25,
      profile: '',
    });
    expect(settings.preset).toBe('raid-buffed');
    expect(settings.buffs.length).toBeGreaterThan(0);
  });
});

describe('DURATIONS and durationLabel', () => {
  it('runs from one to eight minutes in thirty-second steps', () => {
    expect(DURATIONS[0]).toBe(MIN_DURATION_SEC);
    expect(DURATIONS.at(-1)).toBe(MAX_DURATION_SEC);
    expect(DURATIONS).toContain(180);
    expect(DURATIONS.every((seconds, i) => i === 0 || seconds - DURATIONS[i - 1] === 30)).toBe(true);
  });

  it('reads as a clock, not as seconds', () => {
    expect(durationLabel(60)).toBe('1:00');
    expect(durationLabel(180)).toBe('3:00');
    expect(durationLabel(210)).toBe('3:30');
    expect(durationLabel(480)).toBe('8:00');
  });
});

describe('the setters never mutate and always clamp', () => {
  it('keeps duration inside one to eight minutes', () => {
    const base = defaultSettings();
    expect(withDuration(base, 30).encounter.duration_sec).toBe(MIN_DURATION_SEC);
    expect(withDuration(base, 9999).encounter.duration_sec).toBe(MAX_DURATION_SEC);
    expect(base.encounter.duration_sec).toBe(180);
  });

  it('keeps targets between one and ten', () => {
    const base = defaultSettings();
    expect(withTargets(base, 0).encounter.targets).toBe(1);
    expect(withTargets(base, 99).encounter.targets).toBe(MAX_TARGETS);
  });

  it('turns the execute phase off by zeroing the ratio, and back on to the default', () => {
    const base = defaultSettings();
    const off = withExecutePhase(base, false);
    expect(off.encounter.execute_ratio).toBe(0);
    expect(executePhaseOn(off)).toBe(false);
    expect(executePhaseOn(withExecutePhase(off, true))).toBe(true);
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

describe('settingsLabel', () => {
  it('is the one line a saved sim is titled with', () => {
    expect(settingsLabel(defaultSettings())).toBe('Raid-buffed, 3:00, single target');
    expect(settingsLabel(withTargets(withPreset(defaultSettings(), 'solo'), 4))).toBe(
      'Solo, 3:00, 4 targets',
    );
  });
});
