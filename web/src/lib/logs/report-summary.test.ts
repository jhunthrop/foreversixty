import { describe, expect, it } from 'vitest';
import {
  accountReportSummary,
  creditedName,
  formatReportLength,
  ownReportRow,
  ownReportSummary,
  publicReportRow,
  publicReportSummary,
  recentReportSummary,
  reportStatusLabel,
} from './report-summary';

const MINUTE = 60_000;

describe('formatReportLength', () => {
  it('prints minutes under an hour', () => {
    expect(formatReportLength(48 * MINUTE)).toBe('48 m');
  });

  it('prints hours and a padded minute from an hour on', () => {
    expect(formatReportLength(65 * MINUTE)).toBe('1 h 05 m');
    expect(formatReportLength(130 * MINUTE)).toBe('2 h 10 m');
    expect(formatReportLength(120 * MINUTE)).toBe('2 h 00 m');
  });

  it('rounds to the nearest minute', () => {
    expect(formatReportLength(48 * MINUTE + 31_000)).toBe('49 m');
  });
});

describe('ownReportSummary', () => {
  it('reads zone, bosses and length', () => {
    expect(ownReportSummary({ zone: 'Blackrock Depths', killCount: 4, durationMs: 130 * MINUTE })).toBe(
      'Blackrock Depths · 4 bosses · 2 h 10 m',
    );
  });

  it('counts one boss and zero bosses correctly', () => {
    expect(ownReportSummary({ zone: "Onyxia's Lair", killCount: 1, durationMs: 48 * MINUTE })).toBe(
      "Onyxia's Lair · 1 boss · 48 m",
    );
    expect(ownReportSummary({ zone: "Onyxia's Lair", killCount: 0, durationMs: 65 * MINUTE })).toBe(
      "Onyxia's Lair · 0 bosses · 1 h 05 m",
    );
  });

  it('leaves the length out when the API did not send one', () => {
    expect(ownReportSummary({ zone: 'Molten Core', killCount: 4 })).toBe('Molten Core · 4 bosses');
  });
});

describe('publicReportSummary', () => {
  it('reads zone and kills, then the guild when there is one', () => {
    expect(publicReportSummary({ zone: 'Sanguine Depths', killCount: 12 })).toBe(
      'Sanguine Depths · 12 kills',
    );
    expect(publicReportSummary({ killCount: 1, guildName: 'The Last Watch' })).toBe(
      '1 kill · The Last Watch',
    );
  });
});

describe('accountReportSummary', () => {
  it('reads zone and fights, with no bosses part when none died', () => {
    expect(accountReportSummary({ zone: 'Darnassus', fightCount: 0, killCount: 0 })).toBe(
      'Darnassus · 0 fights',
    );
  });

  it('adds bosses before fights once there are any', () => {
    expect(accountReportSummary({ zone: 'Blackrock Depths', fightCount: 14, killCount: 3 })).toBe(
      'Blackrock Depths · 3 bosses · 14 fights',
    );
    expect(accountReportSummary({ zone: 'Ragefire Chasm', fightCount: 1, killCount: 1 })).toBe(
      'Ragefire Chasm · 1 boss · 1 fight',
    );
  });

  it('says so when the zone is not known yet', () => {
    expect(accountReportSummary({ zone: '', fightCount: 2, killCount: 0 })).toBe('No zone yet · 2 fights');
  });
});

describe('reportStatusLabel', () => {
  it('names live and processing, and says nothing once complete', () => {
    expect(reportStatusLabel('live')).toBe('Live');
    expect(reportStatusLabel('processing')).toBe('Processing');
    expect(reportStatusLabel('uploading')).toBe('Uploading');
    expect(reportStatusLabel('complete')).toBe('');
    // MUST 4 (spec 2026-11-07): a finished report carries no Tag, "ready" included.
    expect(reportStatusLabel('ready')).toBe('');
    expect(reportStatusLabel('')).toBe('');
  });
});

describe('creditedName', () => {
  it('reads the name from the API object, spaced', () => {
    expect(creditedName({ key: 'us/pvp/bow-jackzon', name: 'Bow Jackzon' })).toBe('Bow Jackzon');
  });

  it('is null when the report has no logging character', () => {
    for (const value of [undefined, null, '', { key: '', name: '' }]) expect(creditedName(value)).toBeNull();
  });
});

describe('recentReportSummary (W.8)', () => {
  it('reads zone, bosses, fights and guild in the "Your reports" shape', () => {
    expect(
      recentReportSummary({
        zone: 'Blackrock Depths',
        killCount: 3,
        fightCount: 14,
        guildName: 'Lifecycle Test',
      }),
    ).toBe('Blackrock Depths · 3 bosses · 14 fights · Lifecycle Test');
  });

  it('leaves out the bosses when there are none and the zone when unknown', () => {
    expect(recentReportSummary({ zone: 'Darnassus', killCount: 0, fightCount: 1 })).toBe(
      'Darnassus · 1 fight',
    );
    expect(recentReportSummary({ killCount: 1, fightCount: 4 })).toBe('1 boss · 4 fights');
  });
});

describe('publicReportRow (W.9)', () => {
  const FACTS = { title: 'Blackrock Depths', zone: 'Blackrock Depths', killCount: 3, fightCount: 14 };

  it('titles a credited report with the character and leads the line under it with the zone', () => {
    expect(
      publicReportRow({ ...FACTS, loggingCharacter: { key: 'us/pvp/bow-jackzon', name: 'Bow Jackzon' } }),
    ).toEqual({
      title: 'Bow Jackzon',
      summary: 'Blackrock Depths · 3 bosses · 14 fights',
    });
  });

  it('with nobody credited, titles it with the zone and does not say the zone twice', () => {
    expect(publicReportRow(FACTS)).toEqual({ title: 'Blackrock Depths', summary: '3 bosses · 14 fights' });
    expect(publicReportRow({ ...FACTS, title: '' })).toEqual({
      title: 'Blackrock Depths',
      summary: '3 bosses · 14 fights',
    });
  });

  it('keeps the zone on the line when the report has a title of its own', () => {
    expect(publicReportRow({ ...FACTS, title: 'Tuesday night', guildName: 'Lifecycle Test' })).toEqual({
      title: 'Tuesday night',
      summary: 'Blackrock Depths · 3 bosses · 14 fights · Lifecycle Test',
    });
  });

  it('never repeats the title anywhere on the line', () => {
    const row = publicReportRow({ ...FACTS, loggingCharacter: 'Bow Jackzon' });
    expect(row.summary).not.toContain(row.title);
    expect(publicReportRow(FACTS).summary).not.toContain(FACTS.zone);
  });
});

describe('ownReportRow (logs spec 2026-11-07 section 5)', () => {
  const base = {
    title: '',
    zone: '',
    startedAt: '2026-09-17T20:54:38.383Z',
    fightCount: 42,
    killCount: 0,
    bossCount: 0,
    loggingCharacter: null,
  };

  it('falls back from the title to the zone, then to the night it logged', () => {
    expect(ownReportRow(base).title).toBe('17 Sep 2026');
    expect(ownReportRow({ ...base, zone: 'Molten Core' }).title).toBe('Molten Core');
    expect(ownReportRow({ ...base, title: 'Tuesday raid', zone: 'Molten Core' }).title).toBe('Tuesday raid');
  });

  it('reads the pulls, the bosses or Trash only, and the credit when there is one (directions b1)', () => {
    expect(ownReportRow(base).sub).toBe('42 pulls · Trash only');
    expect(
      ownReportRow({
        ...base,
        title: 'Tuesday',
        zone: 'Molten Core',
        bossCount: 1,
        fightCount: 1,
        loggingCharacter: 'Soren',
      }).sub,
    ).toBe('Molten Core · 1 pull · 1 boss · Soren');
    expect(ownReportRow({ ...base, zone: 'Molten Core' }).sub).toBe('42 pulls · Trash only');
  });
});
