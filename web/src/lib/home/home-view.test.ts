// The signed-in home's facts and words (design/specs/2026-11-06-home.md §1).
import { describe, expect, it } from 'vitest';
import type { Me, MeCharacter } from '../account/api';
import { bandWords, gainText, guildLinks, shortDate, upgradeGain, upgradeLine, whoLine } from './home-view';
import type { SlotUpgrade } from './upgrades';

function hunter(over: Partial<MeCharacter> = {}): MeCharacter {
  return {
    key: 'us/normal/felsnap',
    region: 'us',
    ruleset: 'normal',
    name: 'Felsnap',
    class: 'Hunter',
    spec: 'Beast Mastery',
    race: 'Dwarf',
    level: 34,
    faction: 'alliance',
    item_level: 31,
    build: {
      source: 'addon',
      captured_at: '2026-11-04T19:30:00Z',
      talents: { trees: ['0', '0', '0'], points: [23, 0, 0] },
    },
    ...over,
  };
}

function me(over: Partial<Me> = {}): Me {
  return {
    user: { id: 1, battletag: 'A#1', email: null, role: 'user', anonymize: false },
    characters: [hunter()],
    guilds: [],
    bnet_imported_at: '2026-11-06T09:00:00Z',
    ...over,
  };
}

const UPGRADE = {
  slot: 'main_hand',
  pick: {
    slot: 'main_hand',
    item_id: 1,
    item_name: 'Felbough Arm',
    source: 'Fel Ancient',
    source_kind: 'world',
  },
  wornItemId: 2,
  wornItemName: 'Armor Piercer',
  wornUnknown: false,
  gainDps: 1.26,
  notSimChecked: false,
} as unknown as SlotUpgrade;

describe('home view', () => {
  it('words a band and a gain', () => {
    expect(bandWords(30)).toBe('30 to 39');
    expect(bandWords(60)).toBe('60');
    expect(gainText(7.84)).toBe('+7.8');
  });

  it('reads a date-only day as that calendar day', () => {
    expect(shortDate('2026-12-09')).toBe('9 Dec');
  });

  it('reads the character in one line, the unknown left out', () => {
    expect(whoLine(hunter({ class: 'hunter', spec: 'Beast Mastery' }))).toBe(
      'Level 34 Dwarf Beast Mastery Hunter',
    );
    expect(whoLine(hunter({ race: undefined }))).toBe('Level 34 Beast Mastery Hunter');
    expect(whoLine(hunter({ level: undefined, race: undefined }))).toBe('Beast Mastery Hunter');
  });

  it('names an upgrade by its slot and where it comes from, as the game words it', () => {
    expect(upgradeLine(UPGRADE)).toBe('Main hand · Fel Ancient');
    const ring = {
      ...UPGRADE,
      slot: 'finger1',
      pick: { ...UPGRADE.pick, source: 'Blacksmithing', source_kind: 'crafted' },
    };
    expect(upgradeLine(ring as typeof UPGRADE)).toBe('Finger · Blacksmithing');
  });

  it("links the guild and its raid from the account's guild entry", () => {
    const account = me({
      guilds: [
        { id: 2, region: 'us', ruleset: 'normal', name: 'Olympus XXVII', verified: true, verified_count: 1 },
      ],
    });
    const links = guildLinks(account, 2, 41);
    expect(links.raid).toBe(`${links.guild}/raids/41`);
    expect(guildLinks(account, 9, null)).toEqual({ guild: null, raid: null });
  });
});

describe('upgradeGain', () => {
  it('words the row in the band unit the paperdoll uses: DPS, HPS, a tank score', () => {
    expect(upgradeGain({ gainDps: 1182.64 }, 'tank_score')).toEqual({
      value: '+1,182.6',
      unit: 'score',
      tone: 'kill',
    });
    expect(upgradeGain({ gainDps: 12.34 }, 'hps')).toEqual({ value: '+12.3', unit: 'HPS', tone: 'kill' });
    expect(upgradeGain({ gainDps: 5.06 }, 'dps')).toEqual({ value: '+5.1', unit: 'DPS', tone: 'kill' });
  });

  it('has no figure for a slot with no measured gain', () => {
    expect(upgradeGain({ gainDps: null }, 'dps')).toBeUndefined();
  });
});
