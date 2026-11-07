// The signed-in home's facts and words (design/specs/2026-11-06-home.md §1).
import { describe, expect, it } from 'vitest';
import type { Me, MeCharacter } from '../account/api';
import { bandWords, gainText, guildLinks, metaFacts, shortDate, upgradeLines } from './home-view';
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

  it('lists the known facts and leaves the unknown out', () => {
    expect(metaFacts(hunter({ class: 'hunter', spec: 'Beast Mastery' }))[0]).toBe(
      'Level 34 Beast Mastery Hunter',
    );
    expect(metaFacts(hunter())).toEqual(['Level 34 Beast Mastery Hunter', 'Alliance Dwarf', 'Item level 31']);
    expect(metaFacts(hunter({ item_level: undefined, race: undefined, faction: undefined }))).toEqual([
      'Level 34 Beast Mastery Hunter',
    ]);
  });

  it('says what an upgrade is measured against', () => {
    expect(upgradeLines(UPGRADE)).toEqual({ source: 'Fel Ancient', caption: 'vs your Armor Piercer' });
    expect(upgradeLines({ ...UPGRADE, wornItemName: undefined }).caption).toBe('empty slot');
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
