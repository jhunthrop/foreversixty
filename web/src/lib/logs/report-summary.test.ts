import { describe, expect, it } from 'vitest';
import { formatReportLength, ownReportSummary, publicReportSummary } from './report-summary';

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
