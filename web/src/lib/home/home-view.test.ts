// The signed-in home's facts and words (design/specs/2026-11-06-home.md §1).
import { describe, expect, it } from 'vitest';
import type { Me, MeCharacter, MyReport } from '../account/api';
import type { NextRaidView } from '../guild/api';
import {
  bandWords,
  gainText,
  guildLinks,
  logSub,
  metaFacts,
  nextStep,
  ownLogs,
  setUpView,
  shortDate,
  simRow,
  talentsToSpend,
  upgradeLines,
} from './home-view';
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

const RAID = {
  id: 41,
  title: "Onyxia's Lair",
  starts_at: '2026-12-10T02:00:00Z',
  own: { character_key: 'k', state: 'late', role: 'dps' },
} as unknown as NextRaidView;

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

  it('counts the talent points waiting, and nothing without an export', () => {
    expect(talentsToSpend(hunter())).toEqual({ toSpend: 2, earned: 25, spent: 23 });
    expect(talentsToSpend(hunter({ build: undefined }))).toBeNull();
  });

  it('puts the raid sign-up first, then talents, then the top upgrade', () => {
    const raid = nextStep({
      character: hunter(),
      upgrades: [UPGRADE],
      raid: RAID,
      raidHref: '/r',
      raidDay: 'Wed 9 Dec',
    });
    expect(raid?.text).toBe("Next: Onyxia's Lair is on Wed 9 Dec and you are Late. Confirm your sign-up.");
    expect(raid?.action?.href).toBe('/r');
    const talents = nextStep({
      character: hunter(),
      upgrades: [UPGRADE],
      raid: null,
      raidHref: null,
      raidDay: '',
    });
    expect(talents?.text).toBe('Next: spend your 2 talent points, then get Felbough Arm.');
    const spent = hunter({
      build: { source: 'addon', captured_at: '', talents: { trees: [], points: [25, 0, 0] } },
    });
    const gear = nextStep({ character: spent, upgrades: [UPGRADE], raid: null, raidHref: null, raidDay: '' });
    expect(gear?.text).toBe('Next: get Felbough Arm from Fel Ancient.');
    expect(gear?.action?.href).toContain('/bis/hunter/beast-mastery');
    expect(nextStep({ character: spent, upgrades: [], raid: null, raidHref: null, raidDay: '' })?.text).toBe(
      'Nothing waiting.',
    );
    expect(
      nextStep({ character: spent, upgrades: null, raid: null, raidHref: null, raidDay: '' }),
    ).toBeNull();
  });

  it('says what an upgrade is measured against', () => {
    expect(upgradeLines(UPGRADE)).toEqual({ source: 'Fel Ancient', caption: 'vs your Armor Piercer' });
    expect(upgradeLines({ ...UPGRADE, wornItemName: undefined }).caption).toBe('empty slot');
  });

  it('offers the first sim, or the newest with its DPS', () => {
    expect(simRow(undefined, hunter()).title).toBe('Run your first sim');
    const row = simRow(
      {
        sim_id: 's1',
        spec: 'x',
        dps: 61.2,
        engine_version: '',
        created_at: '2026-11-04T20:10:00Z',
        title: 'T',
      },
      hunter(),
    );
    expect(row).toEqual({ title: 'T', sub: 'Simmed 4 Nov', href: '/sim/s1', dps: '61.2' });
  });

  it("keeps the character's own logs, three at most", () => {
    const log = (id: string, key?: string): MyReport =>
      ({
        id,
        title: id,
        zone: 'Z',
        created_at: '2026-11-05T21:00:00Z',
        boss_count: 5,
        logging_character: key === undefined ? null : { key },
      }) as MyReport;
    const logs = [log('a'), log('b', 'other'), log('c', 'us/normal/felsnap'), log('d'), log('e')];
    expect(ownLogs(logs, hunter()).map((r) => r.id)).toEqual(['a', 'c', 'd']);
    expect(logSub(log('a'))).toBe('Z · 5 bosses · 5 Nov');
  });

  it('reads Battle.net as the gear source and offers the addon as an optional extra', () => {
    const view = setUpView(me(), hunter());
    expect(view.checks.map((c) => [c.key, c.tone])).toEqual([['bnet', 'ok']]);
    expect(view.addonOptional).toBe(true);
    // No export at all is no warning while Battle.net has imported the gear.
    expect(setUpView(me(), hunter({ build: undefined })).checks.map((c) => c.tone)).toEqual(['ok']);
  });

  it('falls back to the addon export, and warns only with no gear source at all', () => {
    const addonOnly = setUpView(
      me({ bnet_imported_at: undefined }),
      hunter({ build: { source: 'addon', captured_at: '2026-11-04T19:30:00Z', gear: { head: 1 } } }),
    );
    expect(addonOnly.checks.map((c) => [c.key, c.tone])).toEqual([['addon', 'ok']]);
    expect(addonOnly.addonOptional).toBe(false);
    const none = setUpView(me({ bnet_imported_at: undefined }), hunter({ build: undefined }));
    expect(none.checks.map((c) => [c.key, c.tone])).toEqual([['none', 'warn']]);
    expect(none.checks[0].fix?.href).toBe('/setup');
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
