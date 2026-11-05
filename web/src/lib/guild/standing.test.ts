// web/src/lib/guild/standing.test.ts
import { describe, expect, it } from 'vitest';
import { levelingLine, needsLine, standingSentence } from './standing';
import type { GuildHomeStanding, GuildReadinessRow } from './api';

function standing(overrides: Partial<GuildHomeStanding> = {}): GuildHomeStanding {
  return {
    spec: 'Marksmanship',
    class: 'hunter',
    same_spec_count: 3,
    rank_by_item_level: 2,
    item_level: 66,
    needs_before_next_raid: [],
    ...overrides,
  };
}

describe('standingSentence', () => {
  it('ranks by item level among same-spec peers', () => {
    expect(standingSentence(standing())).toBe(
      '2 of 3 Marksmanship Hunters this week by item level (66 ilvl)',
    );
  });

  it('says nothing when there is exactly one ranked peer: no "the only" gloss', () => {
    expect(standingSentence(standing({ same_spec_count: 1, rank_by_item_level: 1 }))).toBe('');
  });

  it('prints the spec as a player says it, never the camel-cased export name', () => {
    expect(standingSentence(standing({ spec: 'BeastMastery', same_spec_count: 2 }))).toBe(
      '2 of 2 Beast Mastery Hunters this week by item level (66 ilvl)',
    );
  });

  it('says nothing when there are no ranked peers at all', () => {
    expect(standingSentence(standing({ same_spec_count: 0, rank_by_item_level: 0 }))).toBe('');
  });
});

function readinessRow(overrides: Partial<GuildReadinessRow> = {}): GuildReadinessRow {
  return {
    character_key: 'us/pvp/rootbound',
    name: 'Rootbound',
    class: 'druid',
    spec: 'Balance',
    consent: 'gear_bags',
    level: 60,
    gear_gap: { upgrades: 0, gain_dps: 0, not_sim_checked: 0, unit: 'dps' },
    enchants: { missing_slots: [], checked: true },
    consumables: { state: 'stocked' },
    talent_points_unspent: 0,
    talent_points_read: true,
    item_level: 60,
    item_level_delta: 0,
    logged_at: '2026-11-05T00:00:00Z',
    failing: 0,
    no_export: false,
    synced_before_cutoff: null,
    ...overrides,
  };
}

describe('needsLine', () => {
  it('names the character and its fails, from the row the board renders', () => {
    const row = readinessRow({ talent_points_unspent: 2, failing: 1 });
    expect(needsLine(row, 'before Thursday')).toEqual({
      kind: 'needs',
      text: 'What Rootbound needs before Thursday: 2 unspent talent points.',
    });
    expect(needsLine(row, 'before the next raid')?.text).toBe(
      'What Rootbound needs before the next raid: 2 unspent talent points.',
    );
  });

  it('every check passing is the character, never a guild-wide claim', () => {
    expect(needsLine(readinessRow(), 'before Thursday')).toEqual({
      kind: 'clear',
      text: 'Rootbound passes every readiness check.',
    });
  });

  it('a leveling character is told when checks start, never green', () => {
    expect(needsLine(readinessRow({ level: 52 }), 'before Thursday')).toEqual({
      kind: 'leveling',
      text: 'Rootbound is level 52. Rank and readiness checks start at 60.',
    });
  });

  it('says nothing for a row with nothing read', () => {
    const unread = readinessRow({
      gear_gap: null,
      enchants: { missing_slots: [], checked: false },
      consumables: { state: 'unknown' },
      talent_points_read: false,
    });
    expect(needsLine(unread, 'before Thursday')).toBeNull();
  });
});

describe('levelingLine', () => {
  it('says the level once, and when rank and readiness start', () => {
    expect(levelingLine('Rootbound', 52)).toBe(
      'Rootbound is level 52. Rank and readiness checks start at 60.',
    );
  });
});
