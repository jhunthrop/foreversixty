// web/src/lib/race-faction.test.ts
// L8.4 fix round 3 item 15: the tooltip mock's faction line comes from the race, through one
// fixed table (the Classic races, which the six Forever pairs only recombine with classes, as
// addon/ForeverSixty/Rules.lua lists them), so a visitor and the owner see the same mock.
import { describe, expect, it } from 'vitest';
import { factionOfRace, RACE_FACTION } from './race-faction';

describe('factionOfRace', () => {
  it('maps the eight races Rules.lua knows to their faction', () => {
    expect(RACE_FACTION).toEqual({
      human: 'alliance',
      dwarf: 'alliance',
      nightelf: 'alliance',
      gnome: 'alliance',
      orc: 'horde',
      undead: 'horde',
      tauren: 'horde',
      troll: 'horde',
    });
  });

  it('reads the display name, the client token and the slug alike', () => {
    expect(factionOfRace('Night Elf')).toBe('alliance');
    expect(factionOfRace('NightElf')).toBe('alliance');
    expect(factionOfRace('night-elf')).toBe('alliance');
    expect(factionOfRace('Scourge')).toBe('horde');
    expect(factionOfRace('Undead')).toBe('horde');
    expect(factionOfRace('Tauren')).toBe('horde');
  });

  it('is null when the race is unknown or not one of them', () => {
    expect(factionOfRace(null)).toBeNull();
    expect(factionOfRace(undefined)).toBeNull();
    expect(factionOfRace('')).toBeNull();
    expect(factionOfRace('Blood Elf')).toBeNull();
  });
});
