import { describe, expect, it } from 'vitest';
import { specCell, specString } from '../../test-support/spec-table';
import type { RaidNight } from './attendance-contract';
import {
  attendanceCopy as c,
  attendanceSummary,
  cameGuildLine,
  PICKER_WINDOW_DAYS,
  cameAsOutcome,
  durationLabel,
  lateOutcome,
  endedHint,
  killTimeLabel,
  matchedElsewhereParts,
  moveDelta,
  nightLine,
  nextRaidToggle,
  pickerEmpty,
  pickerFigures,
  pickerHelper,
  raidsAttended,
  rosterChip,
  showingParts,
  stampLabel,
  syncedLine,
  uploadedLine,
} from './attendance-copy';

const SPEC = new URL('../../../../design/specs/2026-10-27-guild-attendance.md', import.meta.url);
const HEADING = '### 6.1 On the page';
const cell = (label: string): string[] => specCell(SPEC, HEADING, label);
const str = (label: string, index = 0, fill: Record<string, string> = {}): string =>
  specString(SPEC, HEADING, label, index, fill);

const CHICAGO = 'America/Chicago';
const NEW_YORK = 'America/New_York';
const NIGHT: RaidNight = {
  signed: 23,
  came: 22,
  late: 1,
  not_in_log: 1,
  joined: 2,
  excused: 0,
  kills: 1,
  wipes: 3,
  deaths: 14,
  duration_ms: 161 * 60_000,
};

describe('constants, verbatim from spec 6.1', () => {
  it('match line and ended line', () => {
    expect([c.match.loggedAsNight, c.match.openNight]).toEqual([str('Match line', 0), str('Match line', 1)]);
    expect(c.match.notThisLog).toBe(str('Match line', 4));
    expect(c.match.byTime).toBe('matched by time');
    expect(c.match.byOfficer).toBe('matched by an officer');
    expect(c.ended.noLog).toBe(str('Ended, no log', 0));
    expect(c.ended.matchSoon).toBe(str('Ended, no log', 1));
    expect(`${c.ended.uploadLead}${c.ended.uploadLink}${c.ended.uploadTail}`).toBe(str('Ended, no log', 2));
    expect(c.ended.matchALog).toBe(str('Ended, no log', 3));
  });

  it('clear confirm', () => {
    expect([c.clearConfirm.prompt, c.clearConfirm.confirm, c.clearConfirm.keep]).toEqual(
      cell('Clear confirm'),
    );
  });

  it('picker', () => {
    expect(c.picker.head).toBe(str('Picker', 0));
    expect(c.picker.closest).toBe(str('Picker', 5));
    expect(c.picker.match).toBe(str('Picker', 7));
    expect(c.picker.cancel).toBe(str('Picker', 9));
    expect(pickerHelper('Onyxia', true)).toBe(str('Picker', 1, { Zone: 'Onyxia' }));
    expect(pickerHelper('Onyxia', false)).toBe(str('Picker', 2, { Zone: 'Onyxia' }));
    expect(pickerEmpty('Onyxia')).toBe(str('Picker', 10, { Zone: 'Onyxia' }));
    expect(c.picker.loadFailed).toBe(str('Picker', 11));
    expect(c.picker.matchedElsewhere).toBe(str('Picker', 12));
    expect(pickerFigures(9, 1)).toBe('9 pulls · 1 kill');
  });

  it('your night facts and report words', () => {
    expect(c.yourNight.head).toBe(str('Your night facts', 0));
    expect(c.yourNight.excused).toBe(str('Your night facts', 4));
    expect(c.yourNight.noLoot).toBe(str('Your night facts', 5));
    expect(c.report.head).toBe(str('Report', 0));
    expect(c.report.copy).toBe(str('Report', 1));
    expect(c.report.copied).toBe(str('Report', 2));
    expect(c.report.post).toBe(str('Report', 3));
    expect(c.report.posting).toBe(str('Report', 4));
    expect([c.report.kills, c.report.deaths, c.report.consumables, c.report.loot]).toEqual(
      cell('Report').slice(6, 10),
    );
    expect(c.report.firstPull).toBe(str('Kill order', 1));
    expect(c.deaths.none).toBe(str('Deaths', 2));
    expect(c.consumables.noBuffData).toBe(cell('Consumables').at(-1));
    expect(c.loot.memberEmpty).toBe(str('Loot', 2));
    expect(`${c.loot.officerEmptyLead}${c.loot.officerEmptyLink}${c.loot.officerEmptyTail}`).toBe(
      str('Loot', 1),
    );
    expect(c.liveRegion).toBe(str('Live region'));
    expect(c.changelog).toBe(str('Changelog'));
  });

  it('errors', () => {
    const e = c.errors;
    expect([e.buttons, e.rows, e.post, e.reportLoad, e.officerToolsEnded, e.copyFailed]).toEqual(
      cell('Errors'),
    );
  });

  it('attendance groups, roster and readiness', () => {
    const g = c.attendance.groups;
    expect(c.attendance.head).toBe(str('Attendance', 0));
    expect([g.noShow, g.late, g.joined, g.present, g.benchOut]).toEqual(cell('Attendance').slice(2, 7));
    expect(c.attendance.notInLog).toBe(str('Attendance', 8));
    expect(c.attendance.excused).toBe(str('Attendance', 11));
    expect(c.roster.noNights).toBe(str('Roster', 4));
    expect(c.readiness.noExport).toBe(str('Readiness', 3));
    expect(c.readiness.nobodySigned).toBe(str('Readiness', 4));
    expect(c.consumables.state.unknown).toBe(str('Consumables', 4));
  });
});

describe('night line', () => {
  it('reads the spec example and leaves out zero clauses', () => {
    expect(nightLine(NIGHT)).toBe('22 of 23 came · 2 joined · 1 kill · 3 wipes · 14 deaths · 2 h 41 min');
    expect(nightLine({ ...NIGHT, joined: 0, kills: 0, wipes: 0, deaths: 0, duration_ms: 0 })).toBe(
      '22 of 23 came',
    );
    expect(nightLine({ ...NIGHT, kills: 2, deaths: 1 })).toContain('2 kills · 3 wipes · 1 death ·');
  });

  it("says the bot's sentence when nobody signed up beforehand, never 0 of 0 came", () => {
    const unsigned = {
      ...NIGHT,
      signed: 0,
      came: 0,
      joined: 5,
      kills: 0,
      wipes: 0,
      deaths: 0,
      duration_ms: 0,
    };
    expect(nightLine(unsigned)).toBe('5 came, nobody signed up beforehand');
    expect(nightLine(unsigned)).toBe(str('Night line', 2, { n: '5' }));
    expect(nightLine({ ...unsigned, joined: 1 })).toBe('1 came, nobody signed up beforehand');
    expect(nightLine({ ...unsigned, joined: 0 })).toBe(str('Night line', 3));
    expect(nightLine({ ...unsigned, kills: 1 })).toBe('5 came, nobody signed up beforehand · 1 kill');
    expect(attendanceSummary({ ...NIGHT, signed: 0, came: 0, late: 0, not_in_log: 0, joined: 5 })).toBe(
      '5 came, nobody signed up beforehand',
    );
  });

  it('formats durations and kill times', () => {
    expect(durationLabel(161 * 60_000)).toBe('2 h 41 min');
    expect(durationLabel(120 * 60_000)).toBe('2 h');
    expect(durationLabel(41 * 60_000)).toBe('41 min');
    expect(durationLabel(10_000)).toBe('1 min');
    expect(killTimeLabel(102_000)).toBe('1:42');
    expect(killTimeLabel(3_729_000)).toBe('1:02:09');
  });
});

describe('match line', () => {
  const match = { uploaded_at: '2026-12-11T05:42:00Z', uploaded_by_name: 'Kraggor', how: 'auto' as const };

  it('names the upload, in guild time only when offsets differ', () => {
    expect(uploadedLine(match, CHICAGO, CHICAGO)).toBe(
      'Log uploaded Thu 10 Dec, 11:42 pm by Kraggor · matched by time',
    );
    expect(uploadedLine({ ...match, how: 'hand' }, CHICAGO, NEW_YORK)).toBe(
      'Log uploaded Thu 10 Dec, 11:42 pm (12:42 am guild time) by Kraggor · matched by an officer',
    );
    expect(stampLabel(new Date(match.uploaded_at), CHICAGO, CHICAGO)).toBe('Thu 10 Dec, 11:42 pm');
  });

  it('switches the ended hint at end + 6 h', () => {
    const end = new Date('2026-12-11T05:00:00Z');
    expect(endedHint(end, new Date('2026-12-11T10:59:00Z')).kind).toBe('soon');
    expect(endedHint(end, new Date('2026-12-11T11:01:00Z')).kind).toBe('upload');
  });

  it('writes a matched-elsewhere line with the title apart', () => {
    const parts = matchedElsewhereParts(
      { id: 3, title: 'Onyxia A', starts_at: '2026-12-04T02:00:00Z' },
      CHICAGO,
    );
    expect(`${parts.lead}${parts.title}${parts.tail}`).toBe(
      'Matched to Onyxia A · Thu 3 Dec. Clear it there first.',
    );
  });
});

describe('outcome words', () => {
  it('writes the attendance outcomes', () => {
    expect(lateOutcome(22)).toBe(str('Attendance', 9, { m: '22' }));
    expect(cameAsOutcome('Grimtotem')).toBe(str('Attendance', 10, { Alt: 'Grimtotem' }));
  });

  it('writes the window constants into prose', () => {
    expect(pickerHelper('Onyxia', true)).toContain(`${PICKER_WINDOW_DAYS} days`);
  });
});

describe('report words', () => {
  it('changes', () => {
    expect(moveDelta(238, 291)).toEqual({ text: '+53', rising: true, size: 53 });
    expect(moveDelta(291, 291)).toEqual({ text: '+0', rising: false, size: 0 });
    expect(moveDelta(291, 279)).toEqual({ text: '−12', rising: false, size: 12 });
  });
});

describe('attendance, roster and readiness', () => {
  it('summarises, leaving out zero clauses', () => {
    expect(attendanceSummary(NIGHT)).toBe('22 of 23 came · 1 late · 1 not in the log · 2 more joined');
    expect(attendanceSummary({ ...NIGHT, late: 0, not_in_log: 0, joined: 0, excused: 2 })).toBe(
      '22 of 23 came · 2 excused',
    );
  });

  it('roster chip and lines', () => {
    // The bot's chip (api guilds attendanceChip), one form on every surface.
    expect(rosterChip({ present: 7, nights: 8 }, { present: 26, nights: 30 })).toBe(
      '7 of 8 raid nights · 26 of the last 30',
    );
    expect(rosterChip({ present: 3, nights: 3 }, { present: 3, nights: 3 })).toBe('3 of 3 raid nights');
    expect(rosterChip({ present: 1, nights: 1 }, undefined)).toBe('1 of 1 raid night');
    expect(raidsAttended(7, 8)).toBe('7 of 8 raid nights');
    expect(raidsAttended(1, 1)).toBe('1 of 1 raid night');
    expect(cameGuildLine(7, 8)).toBe(' · came 7 of 8 raid nights');
  });

  it('one zero rule: no eligible night reads No raid nights yet, any night the real ratio', () => {
    expect(raidsAttended(0, 0)).toBe('No raid nights yet');
    expect(raidsAttended(0, 1)).toBe('0 of 1 raid night');
    expect(rosterChip({ present: 0, nights: 0 }, { present: 0, nights: 0 })).toBe('No raid nights yet');
    expect(rosterChip({ present: 0, nights: 0 }, undefined)).toBe('No raid nights yet');
    expect(rosterChip({ present: 0, nights: 0 }, { present: 3, nights: 5 })).toBe('3 of 5 raid nights');
    expect(cameGuildLine(0, 0)).toBe(' · No raid nights yet');
    expect(cameGuildLine(0, 1)).toBe(' · came 0 of 1 raid night');
  });

  it('readiness lines', () => {
    expect(nextRaidToggle('Thu 10 Dec')).toBe('Next raid · Thu 10 Dec');
    const p = showingParts(23, "Onyxia's Lair", 'Thu 10 Dec, 8:00 pm');
    expect(`${p.lead}${p.title}${p.tail}`).toBe(
      "Showing the 23 signed In or Late for Onyxia's Lair, Thu 10 Dec, 8:00 pm.",
    );
    expect(syncedLine('27 Oct')).toBe(str('Readiness', 2, { '27 Oct': '27 Oct' }));
    expect(syncedLine('27 Oct')).toBe('Synced 27 Oct');
  });
});
