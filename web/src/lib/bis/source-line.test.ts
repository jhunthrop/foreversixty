import { describe, expect, it } from 'vitest';
import { instanceLevels } from './load';
import { fullSourceLine, levelsSuffix } from './source-line';

const NO_LEVELS = new Map<string, { min: number; max: number }>();

describe('fullSourceLine', () => {
  it('prints a drop with boss, instance levels from the curated table and chance', () => {
    const levels = new Map([['The Deadmines', { min: 17, max: 24 }]]);
    expect(
      fullSourceLine(
        { kind: 'dungeon', instance: 'The Deadmines', boss: 'Cookie', dropChance: 40.2 },
        { instanceLevels: levels },
      ),
    ).toBe('Cookie · The Deadmines (17-24) · 40%');
  });

  it("prefers the source's own level range over the curated table", () => {
    const levels = new Map([['The Deadmines', { min: 17, max: 24 }]]);
    expect(
      fullSourceLine(
        { kind: 'dungeon', instance: 'The Deadmines', levelMin: 16, levelMax: 22 },
        { instanceLevels: levels },
      ),
    ).toBe('The Deadmines (16-22)');
  });

  it('adds the recipe skill only when the data names it', () => {
    expect(fullSourceLine({ kind: 'crafted', name: 'leatherworking' }, { instanceLevels: NO_LEVELS })).toBe(
      'Crafted · Leatherworking',
    );
    expect(
      fullSourceLine({ kind: 'crafted', name: 'leatherworking', skill: 100 }, { instanceLevels: NO_LEVELS }),
    ).toBe('Crafted · Leatherworking (skill 100)');
  });

  it('adds the quest zone only when the data names it', () => {
    const base = {
      kind: 'quest',
      questName: 'The Defias Brotherhood',
      faction: 'alliance',
      minLevel: 14,
    } as const;
    expect(fullSourceLine(base, { instanceLevels: NO_LEVELS })).toBe(
      'Quest: The Defias Brotherhood · requires level 14',
    );
    expect(fullSourceLine({ ...base, zone: 'Westfall' }, { instanceLevels: NO_LEVELS })).toBe(
      'Quest: The Defias Brotherhood · Westfall · requires level 14',
    );
  });

  it("names a vendor's place and a price that is not gold, only when the data names them", () => {
    expect(fullSourceLine({ kind: 'vendor', npc: 'Pix Xizzix' }, { instanceLevels: NO_LEVELS })).toBe(
      'Vendor: Pix Xizzix',
    );
    expect(
      fullSourceLine(
        {
          kind: 'vendor',
          npc: 'Pix Xizzix',
          place: 'Booty Bay',
          cost: { amount: 100, currency: 'Tarnished Undermine Reals' },
        },
        { instanceLevels: NO_LEVELS },
      ),
    ).toBe('Vendor: Pix Xizzix, Booty Bay · 100 Tarnished Undermine Reals');
  });

  it("names a reputation's standing and where to buy", () => {
    const base = { kind: 'rep', faction: 'Theramore Expeditionary Force', standing: 'exalted' } as const;
    expect(fullSourceLine(base, { instanceLevels: NO_LEVELS })).toBe(
      'Theramore Expeditionary Force reputation · Exalted',
    );
    expect(fullSourceLine({ ...base, place: 'The Barrens' }, { instanceLevels: NO_LEVELS })).toBe(
      'Theramore Expeditionary Force reputation · Exalted · The Barrens',
    );
  });

  it('notes the auction house on a bind-on-equip world drop', () => {
    expect(fullSourceLine({ kind: 'world_drop' }, { instanceLevels: NO_LEVELS, bonding: 'equip' })).toBe(
      'World drop · BoE · on the AH',
    );
  });
});

describe('instanceLevels', () => {
  it('reads dungeon ranges from the curated zone table', () => {
    const levels = instanceLevels();
    expect(levels.get('Ragefire Chasm')).toEqual({ min: 13, max: 18 });
  });

  it('is empty when the file is missing', () => {
    expect(instanceLevels('/nonexistent/events.json').size).toBe(0);
  });
});

describe('levelsSuffix', () => {
  it('prints a range, a single level once, and nothing when unknown', () => {
    expect(levelsSuffix({ min: 15, max: 24 })).toBe(' (15-24)');
    expect(levelsSuffix({ min: 60, max: 60 })).toBe(' (60)');
    expect(levelsSuffix(undefined)).toBe('');
  });
});
