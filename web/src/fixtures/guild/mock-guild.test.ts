// web/src/fixtures/guild/mock-guild.test.ts
// Sanity checks that the fixture's shapes satisfy the API contract and stay internally
// consistent -- this fixture is read by unit tests, e2e specs and the capture script, so a
// silent shape drift here would be invisible everywhere else.
import { describe, expect, it } from 'vitest';
import {
  GUILD_ID,
  VIEWER_MEMBER,
  VIEWER_OFFICER,
  buildMockHome,
  buildMockLoot,
  buildMockProgression,
  buildMockReadiness,
  buildMockRaids,
  buildMockRoster,
} from './mock-guild';

describe('buildMockRoster', () => {
  it('has 24 raiders, matching the spec board caption', () => {
    expect(buildMockRoster()).toHaveLength(24);
  });

  it('every row has a unique character_key', () => {
    const keys = buildMockRoster().map((r) => r.character_key);
    expect(new Set(keys).size).toBe(keys.length);
  });
});

describe('buildMockHome', () => {
  it('reads officer role and standing for the officer viewer', () => {
    const home = buildMockHome(VIEWER_OFFICER);
    expect(home.viewer?.role).toBe('officer');
    expect(home.standing).not.toBeNull();
  });

  it('reads member role for the member viewer', () => {
    const home = buildMockHome(VIEWER_MEMBER);
    expect(home.viewer?.role).toBe('member');
  });

  it('reads public role with no standing for a signed-out visitor', () => {
    const home = buildMockHome(null);
    expect(home.viewer?.role).toBe('public');
    expect(home.standing).toBeNull();
  });

  it('pending holds exactly the unverified rows', () => {
    const home = buildMockHome(VIEWER_OFFICER);
    expect(home.pending?.every((r) => !r.verified)).toBe(true);
    expect(home.pending?.length).toBeGreaterThan(0);
  });

  it('guild id matches every other endpoint fixture', () => {
    expect(buildMockHome(VIEWER_OFFICER).guild.id).toBe(GUILD_ID);
  });
});

describe('buildMockRaids', () => {
  it('names Onyxia and only Onyxia as a boss, per §12.1', () => {
    const raids = buildMockRaids();
    for (const row of raids.rows) {
      for (const fight of row.fights) {
        if (row.zone === "Onyxia's Lair") expect(fight.name).toBe('Onyxia');
        else expect(fight.name).toBeNull();
      }
    }
  });

  it('has 8 raid nights', () => {
    expect(buildMockRaids().rows).toHaveLength(8);
  });
});

describe('buildMockProgression', () => {
  it('names one encounter down, and two unnamed zones', () => {
    const progression = buildMockProgression();
    expect(progression.tier.down).toBe(1);
    expect(progression.encounters).toHaveLength(1);
    expect(progression.unnamed.map((u) => u.zone).sort()).toEqual(['Barrow Deeps', 'Hyjal Summit']);
  });
});

describe('buildMockReadiness', () => {
  it('only includes verified rows, each with a failing count', () => {
    const readiness = buildMockReadiness();
    expect(readiness.rows.every((r) => typeof r.failing === 'number')).toBe(true);
    expect(readiness.rows).toHaveLength(buildMockRoster().filter((r) => r.verified).length);
  });

  it('gates gear_gap/enchants on consent, matching the roster row’s own consent', () => {
    const readiness = buildMockReadiness();
    const roster = buildMockRoster();
    for (const row of readiness.rows) {
      const rosterRow = roster.find((r) => r.character_key === row.character_key);
      const hasGear = rosterRow?.consent === 'gear' || rosterRow?.consent === 'gear_bags';
      expect(row.gear_gap !== null).toBe(hasGear);
      expect(row.enchants.checked).toBe(hasGear);
    }
  });
});

describe('buildMockLoot', () => {
  it('names only Onyxia as the killed encounter', () => {
    const loot = buildMockLoot();
    expect(loot.encounters).toEqual([{ encounter_id: 1084, name: 'Onyxia', zone: "Onyxia's Lair", killed: true }]);
  });

  it('every awarded item has exactly one candidate carrying the award', () => {
    const loot = buildMockLoot();
    for (const item of loot.items) {
      if (item.awarded_to === null) continue;
      expect(item.candidates.some((c) => c.character_key === item.awarded_to?.character_key)).toBe(true);
    }
  });
});
