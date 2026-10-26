// web/src/lib/guild/match-picker.test.ts
// The match picker's state (attendance spec 3.2, 4): load, `Closest` preselected unless another raid
// holds it, the 20-row cap, a taken row that cannot be chosen, the write and the lost race, and
// the keyboard model over the free rows.
import { describe, expect, it, vi } from 'vitest';
import type { MatchCandidate, MatchCandidates } from './attendance-contract';
import { PICKER_MAX_ROWS } from './attendance-copy';
import { MatchPicker, pickerKeyAction, type MatchOutcome } from './match-picker.svelte';

function candidate(id: string, extra: Partial<MatchCandidate> = {}): MatchCandidate {
  return {
    report_id: id,
    zone_name: "Onyxia's Lair",
    uploaded_at: '2026-12-11T05:42:00Z',
    uploaded_by_name: 'Kraggor',
    pulls: 9,
    kills: 1,
    duration_ms: 9_660_000,
    closest: false,
    matched_raid: null,
    ...extra,
  };
}

const OTHER_RAID = { id: 3, title: 'Ony farm', starts_at: '2026-12-03T02:00:00Z' };

function picker(
  page: MatchCandidates,
  match: (id: string) => Promise<MatchOutcome> = () => Promise.resolve({ ok: true }),
): MatchPicker {
  return new MatchPicker({ load: () => Promise.resolve(page), match });
}

describe('MatchPicker load', () => {
  it('preselects the Closest row', async () => {
    const p = picker({ zone_known: true, rows: [candidate('a', { closest: true }), candidate('b')] });
    await p.start();
    expect(p.load).toBe('ready');
    expect(p.selected).toBe('a');
    expect(p.canMatch).toBe(true);
  });

  it('preselects nothing when another raid holds the Closest one', async () => {
    const p = picker({
      zone_known: true,
      rows: [candidate('a', { closest: true, matched_raid: OTHER_RAID }), candidate('b')],
    });
    await p.start();
    expect(p.selected).toBeNull();
    expect(p.canMatch).toBe(false);
    expect(p.freeIds).toEqual(['b']);
  });

  it('keeps the unknown-zone flag and caps the list', async () => {
    const rows = Array.from({ length: PICKER_MAX_ROWS + 5 }, (_, i) => candidate(`r${i}`));
    const p = picker({ zone_known: false, rows });
    await p.start();
    expect(p.zoneKnown).toBe(false);
    expect(p.rows).toHaveLength(PICKER_MAX_ROWS);
  });

  it('fails the load and retries', async () => {
    const load = vi
      .fn<() => Promise<MatchCandidates>>()
      .mockRejectedValueOnce(new Error('down'))
      .mockResolvedValueOnce({ zone_known: true, rows: [] });
    const p = new MatchPicker({ load, match: () => Promise.resolve({ ok: true }) });
    await p.start();
    expect(p.load).toBe('failed');
    await p.start();
    expect(p.load).toBe('ready');
  });

  it('cannot select a taken row', async () => {
    const p = picker({
      zone_known: true,
      rows: [candidate('a', { matched_raid: OTHER_RAID }), candidate('b')],
    });
    await p.start();
    p.select('a');
    expect(p.selected).toBeNull();
    p.select('b');
    expect(p.selected).toBe('b');
  });
});

describe('MatchPicker submit', () => {
  it('matches the selected log', async () => {
    const match = vi.fn(() => Promise.resolve<MatchOutcome>({ ok: true }));
    const p = picker({ zone_known: true, rows: [candidate('a', { closest: true })] }, match);
    await p.start();
    await expect(p.submit()).resolves.toBe(true);
    expect(match).toHaveBeenCalledWith('a');
    expect(p.busy).toBe(false);
  });

  it('shows a refusal in the status line and keeps the picker open', async () => {
    const refused: MatchOutcome = { ok: false, code: 'x', message: 'That did not save; try again.' };
    const p = picker({ zone_known: true, rows: [candidate('a', { closest: true })] }, () =>
      Promise.resolve(refused),
    );
    await p.start();
    await expect(p.submit()).resolves.toBe(false);
    expect(p.line).toBe('That did not save; try again.');
    expect(p.selected).toBe('a');
  });

  it('turns the row matched-elsewhere on a lost race and re-reads who holds it', async () => {
    const lost: MatchOutcome = {
      ok: false,
      code: 'report_matched',
      message: 'That log was just matched to another raid.',
    };
    const load = vi
      .fn<() => Promise<MatchCandidates>>()
      .mockResolvedValueOnce({ zone_known: true, rows: [candidate('a', { closest: true }), candidate('b')] })
      .mockResolvedValueOnce({
        zone_known: true,
        rows: [candidate('a', { closest: true, matched_raid: OTHER_RAID }), candidate('b')],
      });
    const p = new MatchPicker({ load, match: () => Promise.resolve(lost) });
    await p.start();
    await p.submit();
    expect(p.line).toBe(lost.message);
    expect(p.selected).toBeNull();
    expect(p.freeIds).toEqual(['b']);
    expect(p.rows[0].matched_raid).toEqual(OTHER_RAID);
    expect(p.load).toBe('ready');
  });

  it('keeps the row disabled even when the re-read fails', async () => {
    const lost: MatchOutcome = { ok: false, code: 'report_matched', message: 'lost' };
    const load = vi
      .fn<() => Promise<MatchCandidates>>()
      .mockResolvedValueOnce({ zone_known: true, rows: [candidate('a', { closest: true })] })
      .mockRejectedValueOnce(new Error('down'));
    const p = new MatchPicker({ load, match: () => Promise.resolve(lost) });
    await p.start();
    await p.submit();
    expect(p.freeIds).toEqual([]);
    expect(p.load).toBe('ready');
  });

  it('does nothing with no selection', async () => {
    const match = vi.fn(() => Promise.resolve<MatchOutcome>({ ok: true }));
    const p = picker({ zone_known: true, rows: [candidate('a')] }, match);
    await p.start();
    await expect(p.submit()).resolves.toBe(false);
    expect(match).not.toHaveBeenCalled();
  });
});

describe('pickerKeyAction', () => {
  const free = ['a', 'c', 'd'];

  it('moves across the free rows and wraps', () => {
    expect(pickerKeyAction('ArrowDown', 'a', free)).toEqual({ kind: 'move', reportId: 'c' });
    expect(pickerKeyAction('ArrowUp', 'a', free)).toEqual({ kind: 'move', reportId: 'd' });
    expect(pickerKeyAction('ArrowRight', 'd', free)).toEqual({ kind: 'move', reportId: 'a' });
  });

  it('selects the focused row with Space or Enter', () => {
    expect(pickerKeyAction(' ', 'c', free)).toEqual({ kind: 'select', reportId: 'c' });
    expect(pickerKeyAction('Enter', 'c', free)).toEqual({ kind: 'select', reportId: 'c' });
  });

  it('ignores other keys and rows that cannot be chosen', () => {
    expect(pickerKeyAction('Tab', 'a', free)).toBeNull();
    expect(pickerKeyAction(' ', 'b', free)).toBeNull();
  });
});
