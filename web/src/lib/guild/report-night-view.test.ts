// web/src/lib/guild/report-night-view.test.ts
// The raid report's view (web layout spec 2026-11-05 §7.2): the answer figures, the pull rows and
// their causes, the death models, DPS and HPS by kill, the consumables grid (unknown is never red)
// and the Discord preview (nobody named for dying).
import { describe, expect, it } from 'vitest';
import { raidNight } from '../../test-support/raid-fixtures';
import { raidReport } from '../../test-support/report-fixtures';
import type { RaidPullConsumables, RaidReportDeath, RaidReportHit } from './attendance-contract';
import {
  abilityName,
  causeSummary,
  consumablesModel,
  deathModels,
  defaultPullKey,
  discordPreview,
  headlineFigures,
  hitModel,
  outputGroups,
  pullRows,
  raiderDeaths,
} from './report-night-view';
import { pullTitle, reportNightCopy as copy } from './report-night-copy';
import { killTimeLabel } from './attendance-copy';

const report = raidReport();

const hit = (extra: Partial<RaidReportHit> = {}): RaidReportHit => ({
  at_ms: 60_000,
  spell_id: 18584,
  ability: 'Deep Breath',
  source: 'Onyxia',
  amount: 5200,
  overkill: 0,
  ...extra,
});

const death = (blow: RaidReportHit | null): RaidReportDeath => ({
  character_key: 'k',
  name: 'Zed',
  class: 'mage',
  at_ms: 61_000,
  killing_blow: blow,
  last_hits: blow === null ? [] : [blow],
});

describe('headlineFigures', () => {
  it('reads kills, wipes, attempts, deaths in order, kills green and wipes red only when there are any', () => {
    expect(headlineFigures(report).map((f) => [f.label, f.value, f.tone])).toEqual([
      ['Kills', '1', 'kill'],
      ['Wipes', '3', 'wipe'],
      ['Attempts', '4', 'strong'],
      ['Deaths', '14', 'strong'],
    ]);
    const clean = raidReport({ headline: { kills: 0, wipes: 0, attempts: 0, deaths: 0 } });
    expect(headlineFigures(clean).map((f) => f.tone)).toEqual(['strong', 'strong', 'strong', 'strong']);
  });
});

describe('pullRows', () => {
  const rows = pullRows(report);

  it('titles a boss pull with its number and keeps a trash pack its name', () => {
    expect(rows.map((r) => r.title)).toEqual(['Onyxian Whelp pack', 'Onyxia, pull 1', 'Onyxia, pull 2']);
  });

  it('writes the time, the deaths and what killed them, and the result tag', () => {
    expect(rows[0]).toMatchObject({ sub: 'Trash · 1:00', tag: { text: 'Clear', tone: 'muted' } });
    expect(rows[1]).toMatchObject({
      sub: '4:00 · 2 deaths · died to Flame Breath',
      tag: { text: 'Wipe', tone: 'danger' },
    });
    expect(rows[2]).toMatchObject({ sub: '1:42 · 1 death', tag: { text: 'Kill', tone: 'success' } });
  });

  it('opens a pull with causes or with deaths not read yet, never one with nothing to show', () => {
    expect(rows.map((r) => r.expandable)).toEqual([false, true, true]);
  });

  it('opens the wide inspector on the first wipe with deaths', () => {
    expect(defaultPullKey(rows)).toBe('2');
    expect(defaultPullKey(rows.filter((r) => r.pull.result !== 'wipe'))).toBe('4');
    expect(defaultPullKey([rows[0]])).toBeNull();
  });
});

describe('causes', () => {
  it('counts the killing blows, most first, and calls a swing Melee', () => {
    const causes = [death(hit()), death(hit()), death(hit({ spell_id: 1, ability: '' })), death(null)];
    expect(causeSummary(causes)).toBe('died to Deep Breath ×2, Melee');
    expect(causeSummary([death(null)])).toBeNull();
    expect(abilityName(hit({ spell_id: 1, ability: 'Auto Attack' }))).toBe(copy.pulls.melee);
  });

  it('writes a hit with its amount and overkill, and no spell id for a swing', () => {
    expect(hitModel(hit({ amount: 4210, overkill: 1050 }))).toEqual({
      ability: 'Deep Breath',
      spellId: 18584,
      source: 'Onyxia',
      clock: '1:00',
      amount: '4,210 (1,050 overkill)',
    });
    expect(hitModel(hit({ spell_id: 1, ability: 'Melee' })).spellId).toBeNull();
    expect(hitModel(hit()).amount).toBe('5,200');
  });

  it('keeps the killing blow apart from the hits before it', () => {
    const [pyre, krag] = deathModels(report.pulls[1].death_causes ?? []);
    expect(pyre).toMatchObject({ name: 'Pyrewisp', at: '1:42 in' });
    expect(pyre.killingBlow?.ability).toBe('Flame Breath');
    expect(pyre.before.map((h) => h.ability)).toEqual(['Melee']);
    expect(krag).toMatchObject({ killingBlow: null, before: [] });
  });
});

describe('outputGroups', () => {
  it('groups by kill in the API order, the guild unit, the rating only when computed', () => {
    const [group] = outputGroups(report.parses);
    expect(group.boss).toBe('Onyxia');
    expect(group.rows.map((r) => [r.name, r.figure, r.unit, r.sub])).toEqual([
      ['Pyrewisp', '812', 'DPS', 'DPS · rated 96 of 100'],
      ['Menda', '1,204', 'HPS', 'Healer'],
    ]);
  });
});

describe('consumablesModel', () => {
  const model = consumablesModel(report.consumables.by_pull);

  it('is raider by boss pull, Pull n for one boss', () => {
    expect(model.columns).toEqual(['Pull 1', 'Pull 2']);
    expect(model.rows.map((r) => r.label)).toEqual(['Pyrewisp']);
  });

  it('marks an unknown pull ? in the none tone, never red, and says why below', () => {
    const [unknown, read] = model.rows[0].cells;
    expect(unknown).toMatchObject({ tone: 'none', text: '?' });
    expect(read).toMatchObject({ tone: 'warn', text: '✓✓·' });
    expect(read.label).toBe('Up: Flask of Supreme Power, Runn Tum Tuber Surprise, 1 potion');
    expect(model.unknown).toEqual([
      'Onyxia, pull 1: the log carried no character sheet, so nothing can be said.',
    ]);
  });

  it('reds a proven miss, greens all three, and names the boss in the column with several', () => {
    const pull = (name: string, flask: boolean | null, weapon: boolean | null): RaidPullConsumables => ({
      fight_index: 1,
      name,
      attempt: 1,
      reason: null,
      rows: [
        {
          character_key: 'k',
          name: 'Zed',
          class: 'mage',
          flask,
          food: true,
          weapon_enchant: weapon,
          potions: 0,
          brought: [],
        },
      ],
    });
    const two = consumablesModel([pull('Lucifron', false, true), pull('Magmadar', true, true)]);
    expect(two.columns).toEqual(['Lucifron, pull 1', 'Magmadar, pull 1']);
    expect(two.rows[0].cells.map((c) => c.tone)).toEqual(['miss', 'ok']);
    expect(two.unknown).toEqual([]);
  });

  it('draws no rows when every pull is unknown', () => {
    const none = consumablesModel([
      { fight_index: 2, name: 'Onyxia', attempt: 1, reason: 'not_read', rows: null },
    ]);
    expect(none.rows).toEqual([]);
    expect(none.unknown).toEqual(['Onyxia, pull 1: not read from the log yet.']);
  });

  it('says one reason once, naming every pull it covers', () => {
    const unread = (attempt: number): RaidPullConsumables => ({
      fight_index: attempt,
      name: 'Onyxia',
      attempt,
      reason: 'no_character_sheet',
      rows: null,
    });
    expect(consumablesModel([unread(1), unread(2), unread(3)]).unknown).toEqual([
      'Onyxia, pull 1, Onyxia, pull 2 and Onyxia, pull 3: the log carried no character sheet, so nothing can be said.',
    ]);
  });
});

describe('raiderDeaths', () => {
  it('lists each raider who died in the API order, each death with its pull and time', () => {
    const rows = raiderDeaths(report);
    expect(rows.map((row) => row.name)).toEqual(report.deaths.map((row) => row.name));
    rows.forEach((model, index) => {
      const row = report.deaths[index];
      const total = row.on_wipes + row.on_kills;
      expect(model.count).toBe(total > 1 ? copy.deaths.count(row.on_wipes, row.on_kills) : '');
    });
    const pull = report.pulls.find((p) => (p.death_causes ?? []).length > 0)!;
    const death = pull.death_causes![0];
    const model = rows.find((row) => row.key === death.character_key)!;
    expect(model.deaths).toContain(`${pullTitle(pull.name, pull.attempt)} at ${killTimeLabel(death.at_ms)}`);
  });

  it('words the counts as a raider would', () => {
    expect(copy.deaths.count(2, 1)).toBe('2 on wipes · 1 on a kill');
    expect(copy.deaths.count(1, 0)).toBe('1 on a wipe');
    expect(copy.deaths.count(0, 3)).toBe('3 on kills');
  });
});

describe('discordPreview', () => {
  it("is the embed's fields, the death count and the best damage, and names nobody for dying", () => {
    const lines = discordPreview("Onyxia's Lair", report, raidNight());
    expect(lines).toEqual([
      "Report: Onyxia's Lair",
      'Onyxia 1:42',
      '3 wipes · 2 h 41 min · 22 of 23 came · 14 deaths',
      'Pyrewisp 812 DPS on Onyxia',
      'Head of Onyxia to Kraggor',
    ]);
    expect(lines.join(' ')).not.toMatch(/Most deaths|died/);
  });

  it('says no kills on a night without one', () => {
    expect(
      discordPreview('Onyxia', raidReport({ kill_order: [], parses: [], loot: [] }), raidNight())[1],
    ).toBe(copy.discord.noKills);
  });
});
