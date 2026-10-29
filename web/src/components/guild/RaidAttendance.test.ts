// web/src/components/guild/RaidAttendance.test.ts
// Attendance rendered with svelte/server (attendance spec 3.5): the summary with zero clauses left
// out and Excused added, the five groups in order (empty ones dropped), full rows for the first three
// and compact rows for the rest, Excused only on Not-in-the-log and Late rows, the read-only state,
// the row's refusal line, and the outcome words.
import { render } from 'svelte/server';
import { describe, expect, it } from 'vitest';
import type { RaidAttendanceRow } from '../../lib/guild/api';
import { raidNight, text } from '../../test-support/raid-fixtures';
import { attendanceRow } from '../../test-support/report-fixtures';
import RaidAttendance from './RaidAttendance.svelte';

const GATE = 'officer-tools-gate-raids';

function view(
  rows: RaidAttendanceRow[],
  extra: {
    night?: ReturnType<typeof raidNight>;
    readOnly?: boolean;
    lines?: Record<string, string>;
    busy?: Record<string, boolean>;
  } = {},
): string {
  return render(RaidAttendance, {
    props: {
      night: extra.night ?? raidNight(),
      rows,
      readOnly: extra.readOnly ?? false,
      gateId: GATE,
      busy: extra.busy ?? {},
      lines: extra.lines ?? {},
      onExcuse: () => {},
    },
  }).body;
}

const row = (name: string, extra: Partial<RaidAttendanceRow> = {}): RaidAttendanceRow =>
  attendanceRow({ character_key: `us-pvp-${name.toLowerCase()}`, name, ...extra });

const ALL = [
  row('Grimtotem', { outcome: 'no_show', pulls_present: 0, class: 'shaman' }),
  row('Slowpoke', { outcome: 'late', first_seen_offset_s: 1320, pulls_present: 2 }),
  row('Wanderer', { signed: 'bench', outcome: 'present', pulls_present: 3 }),
  row('Kraggor', { own: true }),
  row('Alt', { came_as: 'Menda' }),
  row('Benchwarmer', { signed: 'bench', outcome: 'absent', pulls_present: 0 }),
];

describe('the summary', () => {
  it('reads every clause, then Excused when any row is excused', () => {
    expect(text(view(ALL, { night: raidNight({ excused: 1 }) }))).toContain(
      '22 of 23 came · 1 late · 1 not in the log · 2 more joined · 1 excused',
    );
  });

  it('leaves a zero clause out', () => {
    const body = view(ALL, { night: raidNight({ late: 0, joined: 0, not_in_log: 1 }) });
    const summary = text(body.slice(body.indexOf('raid-attendance-summary')));
    expect(
      summary.startsWith('"> 22 of 23 came · 1 not in the log') ||
        summary.includes('22 of 23 came · 1 not in the log'),
    ).toBe(true);
    expect(summary).not.toContain('late ·');
    expect(summary).not.toContain('joined');
  });
});

describe('the groups', () => {
  const body = view(ALL);

  it('come in the spec order and an empty one is left out', () => {
    const marks = [
      'Signed up, not in the log',
      'Late',
      'Joined from the bench or without a sign-up',
      'Present',
      'Bench and out',
    ].map((label) => body.indexOf(`${label} <span`));
    expect(marks.every((i) => i >= 0)).toBe(true);
    expect(marks).toEqual([...marks].sort((a, b) => a - b));
    expect(view([row('Kraggor')])).not.toContain('Bench and out');
  });

  it('give the first three groups full rows and the last two compact rows', () => {
    expect(body.match(/data-testid="raid-attendance-row"/g)).toHaveLength(3);
    expect(body.match(/data-testid="raid-attendance-compact-row"/g)).toHaveLength(3);
  });

  it('never prints the word no-show', () => {
    expect(text(body).toLowerCase()).not.toContain('no-show');
  });
});

describe('Excused labels and phone rows', () => {
  it('names each checkbox with its raider', () => {
    expect(view(ALL)).toContain('aria-label="Excused, Grimtotem"');
  });

  it('puts outcome and pulls in one line below md', () => {
    const body = view([row('Grimtotem', { outcome: 'no_show' })]);
    expect(body).toMatch(/md:contents[^>]*>\s*<span[^>]*data-testid="raid-attendance-outcome"/);
  });
});

describe('Excused', () => {
  it('shows on Not in the log and Late rows only', () => {
    const body = view(ALL);
    expect(body.match(/data-testid="raid-attendance-excused"/g)).toHaveLength(2);
    expect(view([row('Kraggor'), row('Wanderer', { signed: 'bench' })])).not.toContain(
      'raid-attendance-excused',
    );
  });

  it('is a 44 px label with a checked box when excused', () => {
    const body = view([row('Grimtotem', { outcome: 'no_show', excused: true })]);
    expect(body).toContain('min-h-11');
    expect(body).toMatch(
      /checked[^>]*data-testid="raid-attendance-excused"|data-testid="raid-attendance-excused"[^>]*checked/,
    );
  });

  it('is disabled in officer-read-only and names the gate sentence', () => {
    const body = view(ALL, { readOnly: true });
    expect(body).toContain(`aria-describedby="${GATE}"`);
    expect(body).toMatch(/disabled/);
  });

  it('marks a saving row busy and shows its refusal line', () => {
    const grim = ALL[0];
    const body = view(ALL, {
      busy: { [grim.character_key]: true },
      lines: { [grim.character_key]: 'That did not save. Try again.' },
    });
    expect(body).toContain('aria-busy="true"');
    expect(text(body)).toContain('That did not save. Try again.');
  });
});

describe('the outcome cell', () => {
  const flat = text(view(ALL));

  it('reads Not in the log, the late minutes, Came as and Present', () => {
    expect(flat).toContain('Not in the log');
    expect(flat).toContain('Late, first seen 22 min after the first pull');
    expect(flat).toContain('as Menda');
  });

  it('shows pulls with the ratio and the compact row ratio, text colour when pulls were missed', () => {
    // Full rows read `2 of 4` under the `Boss pulls` head (WP 8).
    expect(flat).toContain('Boss pulls 2 of 4');
    expect(view(ALL)).toContain('data-testid="raid-attendance-heads"');
    const compact = view([row('Wanderer', { pulls_present: 3 }), row('Full')]);
    expect(compact).toContain('text-text numeric text-[13px]">3/4');
    expect(compact).toContain('text-muted numeric text-[13px]">4/4');
  });

  it('says Bench or Out, muted, for a bench-and-out row instead of a pull count', () => {
    expect(flat).toMatch(/Benchwarmer Bench$/);
    expect(flat).not.toContain('Benchwarmer 0/4');
  });

  it('washes the own row', () => {
    expect(view(ALL)).toContain('background:rgba(229,185,85,.06)');
  });
});
