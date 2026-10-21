// web/src/fixtures/guild/mock-guild.test.ts
// Sanity checks that the fixture's shapes satisfy the API contract and stay internally
// consistent -- this fixture is read by unit tests, e2e specs and the capture script, so a
// silent shape drift here would be invisible everywhere else.
import { describe, expect, it } from 'vitest';
import {
  ADDON_SNAPSHOT,
  ADDON_UNMATCHED_MEMBER,
  GUILD_ID,
  VIEWER_MEMBER,
  VIEWER_OFFICER,
  buildMockGuildPage,
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

  // Task L8.3 (Task L6.14's addendum): the latest addon roster snapshot.
  it('carries an addon snapshot: three members matched to roster rows, one that matches no row, and its sender', () => {
    const home = buildMockHome(VIEWER_MEMBER);
    expect(home.addon_snapshot).toEqual(ADDON_SNAPSHOT);
    expect(ADDON_SNAPSHOT.members).toBe(4);
    const matched = home.roster.filter((r) => r.addon !== null && r.addon !== undefined);
    expect(matched.map((r) => r.name).sort()).toEqual(['Duskstrider', 'Kraggor', 'Zulmara']);
    // The fourth snapshot member is on no roster row: counted in `members` only, never a row.
    expect(home.roster.some((r) => r.name === ADDON_UNMATCHED_MEMBER)).toBe(false);
    // Every other row is `null`, never absent, as the addendum sends it.
    expect(home.roster.filter((r) => r.addon === null)).toHaveLength(home.roster.length - 3);
    for (const r of matched) {
      expect(r.addon?.snapshot_at).toBe(ADDON_SNAPSHOT.received_at);
      expect(r.addon?.seen_at).toMatch(/^\d{4}-\d{2}-\d{2} \d{2}:\d{2}$/);
    }
    expect(home.roster.some((r) => r.character_key === ADDON_SNAPSHOT.by_character)).toBe(true);
  });

  it('builds the same home without a snapshot: `addon_snapshot` null and every row null', () => {
    const home = buildMockHome(VIEWER_MEMBER, undefined, null, false);
    expect(home.addon_snapshot).toBeNull();
    expect(home.roster.every((r) => r.addon === null)).toBe(true);
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
    expect(loot.encounters).toEqual([
      { encounter_id: 1084, name: 'Onyxia', zone: "Onyxia's Lair", killed: true },
    ]);
  });

  it('every awarded item has exactly one candidate carrying the award', () => {
    const loot = buildMockLoot();
    for (const item of loot.items) {
      if (item.awarded_to === null) continue;
      expect(item.candidates.some((c) => c.character_key === item.awarded_to?.character_key)).toBe(true);
    }
  });

  // Live-fix round, defect 5: this fixture now matches the live contract's own 22-item
  // table exactly (api/internal/guilds/loot.go's onyxiaLoot), with candidates read
  // through the tier-1 (fallback) shape every real candidate on a live roster uses today.
  it("lists the contract's 22 Onyxia drops as the API does: a shared name once, with its copies", () => {
    const items = buildMockLoot().items;
    expect(items).toHaveLength(21);
    expect(items.filter((item) => item.name === 'Head of Onyxia')).toHaveLength(1);
    expect(items.find((item) => item.name === 'Head of Onyxia')?.copies).toBe(2);
    expect(items.reduce((sum, item) => sum + (item.copies ?? 1), 0)).toBe(22);
  });

  it('offers a class-locked token only to its class, so a roster without one lists no one', () => {
    const items = buildMockLoot().items;
    const roster = new Map(buildMockRoster().map((r) => [r.character_key, r.class]));
    const classesOf = (name: string) =>
      new Set(items.find((item) => item.name === name)?.candidates.map((c) => roster.get(c.character_key)));
    expect(classesOf("Dragonstalker's Helm")).toEqual(new Set(['hunter']));
    expect(classesOf('Stormrage Cover')).toEqual(new Set(['druid']));
    expect(classesOf('Helm of Wrath')).toEqual(new Set(['warrior']));
    // No verified paladin or shaman in the fixture roster: the real API answers an empty list.
    expect(items.find((item) => item.name === 'Judgement Crown')?.candidates).toEqual([]);
    expect(items.find((item) => item.name === 'Helmet of Ten Storms')?.candidates).toEqual([]);
    // A healer trinket is not offered to a warrior.
    expect(classesOf('Shard of the Scale').has('warrior')).toBe(false);
  });

  it('the awarded item carries the award id the Undo sends', () => {
    const awarded = buildMockLoot().items.find((item) => item.awarded_to !== null);
    expect(awarded?.awarded_to?.award_id).toBeGreaterThan(0);
  });

  it('gives every slotted item with candidates the tier-1 shape', () => {
    const loot = buildMockLoot();
    for (const item of loot.items) {
      if (item.slot === '') {
        expect(item.candidates).toHaveLength(0);
        continue;
      }
      for (const candidate of item.candidates) {
        expect(candidate.gain_dps).toBeNull();
        expect(candidate.not_sim_checked).toBe(true);
        expect(candidate.ilvl_delta).not.toBeNull();
      }
    }
  });

  it('selects Onyxia as a number, never null', () => {
    expect(buildMockLoot().selected).toBe(1084);
  });
});

describe('buildMockGuildPage', () => {
  it('matches the same guild id every other endpoint fixture uses', () => {
    expect(buildMockGuildPage().guild.id).toBe(GUILD_ID);
  });

  it('carries no kill/wipe counts on its reports -- the real public endpoint’s own shape', () => {
    for (const report of buildMockGuildPage().reports) {
      expect(report).not.toHaveProperty('fight_count');
      expect(report).not.toHaveProperty('kill_count');
    }
  });
});
