import { describe, expect, it } from 'vitest';
import { buildMockReadiness } from '../../fixtures/guild/mock-guild';
import type { GuildReadinessPage } from './api';
import {
  attendanceSourceOf,
  raidsChangeAction,
  readinessTabView,
  type TabFetchStatus,
} from './guild-page-decisions';

const RAID = { id: 7, title: 'Onyxia', starts_at: '2026-12-10T20:00:00Z', cutoff_at: '2026-12-03T20:00:00Z' };
const all = buildMockReadiness();
const mine = (rows: number): GuildReadinessPage => ({ ...all, raid: RAID, rows: all.rows.slice(0, rows) });

describe('readinessTabView', () => {
  it('shows nothing before the first read', () => {
    expect(readinessTabView(null, true, null)).toEqual({
      rows: [],
      raid: null,
      raidId: null,
      guildTimezone: null,
    });
  });

  it('keeps the toggle across a re-read that replaces the board', () => {
    const first = readinessTabView({ all, raid: mine(1) }, true, null);
    const reread = readinessTabView({ all, raid: mine(2) }, true, null);
    expect(first.rows).toHaveLength(1);
    expect(reread.rows).toHaveLength(2);
    expect(readinessTabView({ all, raid: mine(2) }, false, null).rows).toBe(all.rows);
  });

  it('shows the whole board and no toggle when the filtered read is gone', () => {
    const view = readinessTabView({ all, raid: null }, true, { id: 7, guild_timezone: 'UTC' });
    expect(view).toEqual({ rows: all.rows, raid: null, raidId: null, guildTimezone: null });
  });

  it('takes the guild timezone from home only when it names the same raid', () => {
    const board = { all, raid: mine(1) };
    const paris = (id: number) => readinessTabView(board, true, { id, guild_timezone: 'Europe/Paris' });
    expect(paris(7).guildTimezone).toBe('Europe/Paris');
    expect(paris(7).raidId).toBe(7);
    expect(paris(8).guildTimezone).toBeNull();
    expect(readinessTabView(board, true, null).guildTimezone).toBeNull();
  });
});

describe('raidsChangeAction', () => {
  const cases: Array<[boolean, TabFetchStatus, boolean, string]> = [
    [true, 'loading', false, 'reread'],
    [true, 'ready', true, 'reread'],
    [true, 'ready', false, 'reset'],
    [true, 'idle', true, 'reset'],
    [false, 'loading', true, 'reset'],
    [false, 'ready', true, 'reset'],
  ];
  it.each(cases)('guild %s, %s, tab open %s -> %s', (hasGuild, status, open, expected) => {
    expect(raidsChangeAction(hasGuild, status, open)).toBe(expected);
  });
});

describe('attendanceSourceOf', () => {
  it('defaults to reports', () => {
    expect(attendanceSourceOf(null)).toBe('reports');
    expect(attendanceSourceOf({})).toBe('reports');
    expect(attendanceSourceOf({ attendance_source: 'raids' })).toBe('raids');
  });
});
