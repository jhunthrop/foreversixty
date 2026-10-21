import { describe, expect, it } from 'vitest';
import type { GuildRaidFight, GuildRaidRow } from './api';
import { bossIdsOf, raidTally, raidTallyLine, reportTallyLine } from './raid-view';

function fight(
  index: number,
  name: string | null,
  encounterId: number | null,
  kill: boolean,
): GuildRaidFight {
  return { index, name, encounter_id: encounterId, kill, duration_ms: 60000, deaths: 0, players: 5 };
}

function night(fights: GuildRaidFight[], wipes = 0): GuildRaidRow {
  return {
    id: 'r1',
    title: 'Onyxia',
    zone: "Onyxia's Lair",
    created_at: '2026-10-20T19:00:00Z',
    duration_ms: 1320000,
    fight_count: fights.length,
    kill_count: fights.filter((f) => f.kill).length,
    wipe_count: wipes,
    raiders: 5,
    deaths: 0,
    top_parse: null,
    fights,
    present: [],
  };
}

const BOSSES = new Set([1084]);

describe('raidTallyLine', () => {
  it('counts bosses and says the trash separately', () => {
    const row = night([
      fight(0, 'Onyxian Whelp', 2001, true),
      fight(1, 'Onyxian Whelp', 2001, true),
      fight(2, 'Onyxia', 1084, true),
    ]);
    expect(raidTally(row, BOSSES)).toEqual({ bosses: ['Onyxia'], bossKills: 1, trashPacks: 2, wipes: 0 });
    expect(raidTallyLine(row, BOSSES)).toBe('1 boss kill (Onyxia) · 2 trash packs · 0 wipes');
  });

  it('leaves the trash out when there was none and pluralises the bosses', () => {
    const row = night([fight(0, 'Onyxia', 1084, true), fight(1, 'Onyxia', 1084, true)], 1);
    expect(raidTallyLine(row, BOSSES)).toBe('2 boss kills (Onyxia) · 1 wipe');
  });

  it('reads 0 boss kills on a night with no boss down', () => {
    const row = night([fight(0, 'Onyxian Whelp', 2001, true)]);
    expect(raidTallyLine(row, BOSSES)).toBe('0 boss kills · 1 trash pack · 0 wipes');
  });

  it('falls back to the API figure under its own name while the tier list is unknown', () => {
    const row = night([fight(0, 'Onyxia', 1084, true)]);
    expect(raidTallyLine(row, null)).toBe('1 kill · 0 wipes');
    expect(reportTallyLine(4, 1)).toBe('4 kills · 1 wipe');
  });
});

describe('bossIdsOf', () => {
  it('is null until the progression page is known', () => {
    expect(bossIdsOf(null)).toBeNull();
  });
});
