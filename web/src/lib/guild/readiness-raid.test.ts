import { describe, expect, it } from 'vitest';
import { buildMockReadiness } from '../../fixtures/guild/mock-guild';
import type { GuildReadinessPage, GuildReadinessRow } from './api';
import { loadReadinessBoard, readinessRows, syncCell, toggleVisible } from './readiness-raid';

const NEXT = { id: 7, title: "Onyxia's Lair", zone_name: "Onyxia's Lair", starts_at: '2026-12-10T20:00:00Z' };
const RAID = { id: 7, title: NEXT.title, starts_at: NEXT.starts_at, cutoff_at: '2026-12-03T20:00:00Z' };

function page(extra: Partial<GuildReadinessPage> = {}): GuildReadinessPage {
  return { ...buildMockReadiness(), ...extra };
}

function row(extra: Partial<GuildReadinessRow>): GuildReadinessRow {
  return { ...buildMockReadiness().rows[0], ...extra };
}

describe('loadReadinessBoard', () => {
  it('fetches the filtered board for the next raid', async () => {
    const calls: Array<number | undefined> = [];
    const fetchPage = (_id: number, raidId?: number): Promise<GuildReadinessPage> => {
      calls.push(raidId);
      return Promise.resolve(raidId === undefined ? page({ next_raid: NEXT }) : page({ raid: RAID }));
    };
    const board = await loadReadinessBoard(1, fetchPage);
    expect(calls).toEqual([undefined, 7]);
    expect(board.raid?.raid).toEqual(RAID);
    expect(toggleVisible(board)).toBe(true);
  });

  it('skips the filtered read when no raid is planned', async () => {
    const calls: Array<number | undefined> = [];
    const fetchPage = (_id: number, raidId?: number): Promise<GuildReadinessPage> => {
      calls.push(raidId);
      return Promise.resolve(page({ next_raid: null }));
    };
    const board = await loadReadinessBoard(1, fetchPage);
    expect(calls).toEqual([undefined]);
    expect(toggleVisible(board)).toBe(false);
  });

  it('falls back to the plain board when the filtered read fails', async () => {
    const fetchPage = (_id: number, raidId?: number): Promise<GuildReadinessPage> =>
      raidId === undefined ? Promise.resolve(page({ next_raid: NEXT })) : Promise.reject(new Error('404'));
    const board = await loadReadinessBoard(1, fetchPage);
    expect(board.raid).toBeNull();
    expect(toggleVisible(board)).toBe(false);
  });
});

describe('readinessRows', () => {
  const all = page({ next_raid: NEXT });
  const filtered = page({ raid: RAID, rows: all.rows.slice(0, 2) });
  const board = { all, raid: filtered };

  it('is the raid board while the toggle is on', () => {
    expect(readinessRows(board, true)).toEqual(filtered.rows);
  });

  it('is the guild board while it is off or there is no filtered board', () => {
    expect(readinessRows(board, false)).toEqual(all.rows);
    expect(readinessRows({ all, raid: null }, true)).toEqual(all.rows);
  });
});

describe('syncCell', () => {
  it('names a missing export', () => {
    expect(syncCell(row({ no_export: true }), RAID)).toEqual({ kind: 'no_export', text: 'No export yet' });
  });

  it('names how many days before the raid a stale sync was', () => {
    const stale = row({ no_export: false, synced_before_cutoff: true, logged_at: '2026-12-01T20:00:00Z' });
    expect(syncCell(stale, RAID)).toEqual({ kind: 'stale', text: 'Synced 9 days before the raid' });
  });

  it('leaves a fresh sync and the plain board to the day word', () => {
    expect(syncCell(row({ no_export: false, synced_before_cutoff: false }), RAID).kind).toBe('day');
    expect(syncCell(row({ no_export: false, synced_before_cutoff: null }), null).kind).toBe('day');
  });
});
