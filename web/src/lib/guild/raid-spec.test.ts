// web/src/lib/guild/raid-spec.test.ts
// The raid page's spec ask rules (guided-cards spec P.10.7): the class's specs in catalogue order,
// the saved spec, and which presses wait for the pick.
import { describe, expect, it } from 'vitest';
import { viewerCharacter } from '../../test-support/raid-fixtures';
import {
  pressAsksSpec,
  savedSpec,
  showsSpecSelect,
  specIconSrc,
  specOfClass,
  specsForClass,
} from './raid-spec';

const PALADIN = { class: 'paladin', spec: 'Holy' } as const;

describe('specsForClass', () => {
  it('lists the class in catalogue order, with the full names', () => {
    expect(specsForClass('druid').map((spec) => spec.name)).toEqual([
      'Balance',
      'Feral',
      'Feral Bear',
      'Restoration',
    ]);
    expect(specsForClass('hunter').map((spec) => spec.name)).toContain('Beast Mastery');
  });

  it('is empty for an unknown class', () => {
    expect(specsForClass('death-knight')).toEqual([]);
  });
});

describe('savedSpec and specOfClass', () => {
  it('reads the saved slug, and null while unset, absent or foreign to the class', () => {
    expect(savedSpec(viewerCharacter({ ...PALADIN, raid_spec: 'paladin-holy' }))?.name).toBe('Holy');
    expect(savedSpec(viewerCharacter({ ...PALADIN, raid_spec: null }))).toBeNull();
    expect(savedSpec(viewerCharacter(PALADIN))).toBeNull();
    expect(savedSpec(viewerCharacter({ ...PALADIN, raid_spec: 'priest-holy' }))).toBeNull();
  });

  it('refuses a slug of another class', () => {
    expect(specOfClass('paladin', 'paladin-retribution')?.name).toBe('Retribution');
    expect(specOfClass('paladin', 'priest-holy')).toBeNull();
  });
});

describe('asking', () => {
  it('shows for In, Late and Bench, never for Out or unsigned', () => {
    expect(['in', 'late', 'bench', 'out', null].map((s) => showsSpecSelect('paladin', s as never))).toEqual([
      true,
      true,
      true,
      false,
      false,
    ]);
  });

  it('holds a press only while no spec is saved', () => {
    expect(pressAsksSpec(viewerCharacter({ ...PALADIN, raid_spec: null }), 'in')).toBe(true);
    expect(pressAsksSpec(viewerCharacter({ ...PALADIN, raid_spec: 'paladin-holy' }), 'in')).toBe(false);
    expect(pressAsksSpec(viewerCharacter({ ...PALADIN, raid_spec: null }), 'out')).toBe(false);
  });
});

describe('specIconSrc', () => {
  it('is the spec talent-tab icon from the active build', () => {
    const holy = specOfClass('paladin', 'paladin-holy');
    expect(holy).not.toBeNull();
    expect(specIconSrc(holy!)).toMatch(/^\/data\/[\d.]+\/icons\/[a-z0-9_]+\.webp$/);
  });
});
