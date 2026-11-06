// web/src/lib/guild/report-attendance-view.test.ts
// The report rail's attendance rows (web layout spec 2026-11-05 §7.2, §7.4): the effective night,
// and under a fix who set it and what the log said; the attendance groups' order.
import { describe, expect, it } from 'vitest';
import { attendanceRow } from '../../test-support/report-fixtures';
import { fixOptions, fixRowModel, fixRows, logOutcomeText } from './report-attendance-view';

const FIX = { by_name: 'Kraggor', at: '2026-12-11T06:00:00Z' };

describe('logOutcomeText', () => {
  it('reads what the log said, before any fix', () => {
    expect(logOutcomeText(attendanceRow({ auto_outcome: 'no_show', outcome: 'present' }))).toBe(
      'Not in the log',
    );
    expect(logOutcomeText(attendanceRow({ auto_outcome: 'absent', signed: 'bench' }))).toBe('Bench');
    expect(logOutcomeText(attendanceRow({ auto_outcome: 'absent', signed: 'out' }))).toBe('Out');
    expect(logOutcomeText(attendanceRow({ auto_outcome: 'no_show', auto_excused: true }))).toBe(
      'Not in the log · Excused',
    );
  });
});

describe('fixRowModel', () => {
  it('has no auto line and no value without a fix', () => {
    expect(fixRowModel(attendanceRow())).toMatchObject({ effective: 'Present', auto: null, value: '' });
  });

  it('reads the fix and who set it, the log muted under it', () => {
    const row = attendanceRow({
      outcome: 'present',
      auto_outcome: 'no_show',
      override: { status: 'bench', ...FIX },
    });
    expect(fixRowModel(row)).toMatchObject({
      effective: 'Bench, set by Kraggor',
      auto: 'The log said Not in the log',
      value: 'bench',
    });
  });
});

describe('fixRows', () => {
  it('puts the rows an officer is likely to fix first', () => {
    const rows = fixRows([
      attendanceRow({ character_key: 'a', name: 'Ann' }),
      attendanceRow({ character_key: 'b', name: 'Bo', outcome: 'no_show', auto_outcome: 'no_show' }),
    ]);
    expect(rows.map((m) => m.row.name)).toEqual(['Bo', 'Ann']);
  });

  it('offers present, absent, bench and excused', () => {
    expect(fixOptions().map((o) => o.label)).toEqual(['Present', 'Absent', 'Bench', 'Excused']);
  });
});
