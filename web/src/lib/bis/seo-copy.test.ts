// web/src/lib/bis/seo-copy.test.ts
import { describe, expect, it } from 'vitest';
import { bisSeoCopy } from './seo-copy';

const fury = { specName: 'Fury', className: 'Warrior', preRaid: true };
const furyRaids = { ...fury, preRaid: false };

describe('bisSeoCopy', () => {
  it('puts the spec before the class, as players search, and never an em dash', () => {
    const all = [
      bisSeoCopy.specTitle(fury, 20),
      bisSeoCopy.bandTitle(fury, 40),
      bisSeoCopy.bandTitle(fury, 60),
      bisSeoCopy.heading(fury, 20),
      bisSeoCopy.specDescription(fury, 20),
      bisSeoCopy.bandDescription(fury, 50),
    ];
    for (const text of all) {
      expect(text.startsWith('Fury Warrior')).toBe(true);
      expect(text).not.toContain('—');
    }
    expect(bisSeoCopy.specTitle(fury, 20)).toBe(
      'Fury Warrior BiS: Pre-Raid and Leveling 20 to 60, WoW Forever',
    );
    expect(bisSeoCopy.bandTitle(fury, 40)).toBe('Fury Warrior BiS, Levels 40 to 49, WoW Forever');
  });

  it('says Pre-Raid only while the level 60 list holds no raid drop', () => {
    expect(bisSeoCopy.heading(fury, 60)).toBe('Fury Warrior Pre-Raid BiS');
    expect(bisSeoCopy.heading(furyRaids, 60)).toBe('Fury Warrior BiS, Level 60');
    expect(bisSeoCopy.specTitle(furyRaids, 20)).not.toContain('Pre-Raid');
    expect(bisSeoCopy.bandTitle(furyRaids, 60)).toBe('Fury Warrior BiS, Level 60, WoW Forever');
  });

  it('names the band in every band h1 below 60', () => {
    expect(bisSeoCopy.heading(fury, 20)).toBe('Fury Warrior BiS, Levels 20 to 29');
  });
});
