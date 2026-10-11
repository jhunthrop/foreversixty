import { describe, expect, it } from 'vitest';
import type { MeCharacter } from '../account/api';
import type { CharacterFight } from '../rankings/api';
import { formatKillFigure, logsAnswer, logsAnswerCopy, logsAnswerSpec } from './latest-kill';

const NOW = new Date('2026-10-10T12:00:00Z');

function fight(overrides: Partial<CharacterFight> = {}): CharacterFight {
  return {
    encounter: 'Onyxia',
    encounter_id: 1084,
    difficulty: 1,
    metric: 'dps',
    value: 612.4,
    fought_at: '2026-10-07T10:00:00Z',
    report_id: 'rep1',
    fight_index: 3,
    execution_score: null,
    ...overrides,
  };
}

const NEWER_REPORT = { id: 'rep2', createdAt: '2026-10-09T10:00:00Z' };
const OLDER_REPORT = { id: 'rep0', createdAt: '2026-10-06T21:00:00Z' };

describe('formatKillFigure', () => {
  it('keeps one decimal under 1,000 and goes whole with a separator at or above', () => {
    expect(formatKillFigure(612.44)).toBe('612.4');
    expect(formatKillFigure(999.96)).toBe('1,000');
    expect(formatKillFigure(1250.4)).toBe('1,250');
  });
});

describe('logsAnswer: the latest kill against the last report', () => {
  it('reads the kill, its age and the link to the fight', () => {
    const answer = logsAnswer([fight()], null, NOW);
    expect(answer).toEqual({
      kind: 'figures',
      label: 'Latest kill',
      figures: [{ value: '612.4', unit: 'DPS' }],
      note: {
        text: 'Onyxia kill · 3 days ago',
        link: { label: 'Open the fight', href: '/reports/rep1?fight=3' },
      },
    });
  });

  it('adds the report line, linked, only when the newest report is newer than the kill', () => {
    const answer = logsAnswer([fight()], NEWER_REPORT, NOW);
    expect(answer).toMatchObject({
      secondNote: { text: 'Last report 1 day ago', href: '/reports/rep2' },
    });
  });

  it('shows no report line when the report is older than the kill', () => {
    expect(logsAnswer([fight()], OLDER_REPORT, NOW)).not.toHaveProperty('secondNote');
  });

  it('prints a date for a kill older than two weeks', () => {
    const answer = logsAnswer([fight({ fought_at: '2026-09-15T20:00:00Z' })], null, NOW);
    expect(answer).toMatchObject({ note: { text: 'Onyxia kill · Sep 15' } });
  });

  it('labels a healer kill HPS and skips a damage-taken row', () => {
    const answer = logsAnswer(
      [fight({ metric: 'damage_taken' }), fight({ metric: 'hps', value: 801 })],
      null,
      NOW,
    );
    expect(answer).toMatchObject({ figures: [{ value: '801.0', unit: 'HPS' }] });
  });

  it('says there are no logs when there is neither a kill nor a report', () => {
    expect(logsAnswer([], null, NOW)).toEqual({ kind: 'sentence', text: logsAnswerCopy.empty });
  });

  it('names the report when there is no kill yet but a report exists', () => {
    expect(logsAnswer([], NEWER_REPORT, NOW)).toEqual({
      kind: 'sentence',
      text: 'No ranked kill yet. Last report 1 day ago.',
    });
  });
});

describe('logsAnswerSpec', () => {
  const character = { key: 'us/normal/zulmara', name: 'Zulmara' } as MeCharacter;

  it('loads the history and the newest report together', async () => {
    const spec = logsAnswerSpec({
      history: () => Promise.resolve([fight()]),
      newestReport: () => Promise.resolve(NEWER_REPORT),
      now: () => NOW,
    });
    await expect(spec.load(character)).resolves.toMatchObject({ kind: 'figures' });
    expect(spec.strings.failed).toBe(logsAnswerCopy.failed);
  });

  it('rejects, so the header shows its failed state, when the history cannot load', async () => {
    const spec = logsAnswerSpec({
      history: () => Promise.reject(new Error('down')),
      newestReport: () => Promise.resolve(null),
      now: () => NOW,
    });
    await expect(spec.load(character)).rejects.toThrow('down');
  });
});
