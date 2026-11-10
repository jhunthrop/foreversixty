// web/src/lib/report/format.test.ts
import { describe, expect, it } from 'vitest';
import {
  classColorVar,
  formatClock,
  formatPerSecond,
  joinFacts,
  outcomeLabel,
  parseTitle,
  TRASH_PARSE,
  percentileToken,
  phaseReached,
  schoolName,
  schoolToken,
  formatAmountLike,
  formatLanded,
  formatShare,
} from './format';
import { clockDuration, clockDurationPrecise } from '../format/duration';
import { compact, ordinal, percent } from '../format/number';

describe('report formatting', () => {
  it('formats durations as minutes and seconds, with tenths under a minute', () => {
    expect(clockDuration(6400)).toBe('0:06.4');
    expect(clockDuration(40000)).toBe('0:40');
    expect(clockDuration(107000)).toBe('1:47');
    expect(clockDuration(3_723_000)).toBe('1:02:03');
    expect(clockDuration(0)).toBe('0:00');
  });

  it('formats wall-clock times from the engine’s ISO strings', () => {
    expect(formatClock('2026-09-26T20:12:00Z')).toBe('20:12:00');
    expect(formatClock('2026-09-26T20:10:04.6Z')).toBe('20:10:04');
    expect(formatClock('not a time')).toBe('--:--:--');
  });

  it('groups amounts, and abbreviates from a million up', () => {
    expect(compact(0)).toBe('0');
    expect(compact(1484)).toBe('1,484');
    expect(compact(99_999)).toBe('99,999');
    expect(compact(239_230)).toBe('239.2k');
    expect(compact(999_999)).toBe('1000k');
    expect(compact(1_250_000)).toBe('1.25M');
    expect(compact(12_500_000)).toBe('12.5M');
    expect(compact(1_250_000_000)).toBe('1.25B');
  });

  it('divides by the window, not the fight, and never by zero', () => {
    expect(formatPerSecond(4400, 40000)).toBe('110.0');
    expect(formatPerSecond(1484, 6400)).toBe('231.9');
    expect(formatPerSecond(100, 0)).toBe('0.0');
  });

  it('formats percentages to one decimal', () => {
    expect(percent(7.5)).toBe('7.5%');
    expect(percent(100)).toBe('100.0%');
  });

  it('maps a class name to its token, and anything unknown to the body colour', () => {
    expect(classColorVar('Warrior')).toBe('var(--color-class-warrior)');
    expect(classColorVar('Death Knight')).toBe('var(--color-text)');
    expect(classColorVar(undefined)).toBe('var(--color-text)');
  });

  it('maps a parse percentile onto the ladder the community reads', () => {
    expect(percentileToken(3)).toBe('var(--color-parse-grey-text)');
    expect(percentileToken(40)).toBe('var(--color-parse-green)');
    expect(percentileToken(60)).toBe('var(--color-parse-blue)');
    expect(percentileToken(80)).toBe('var(--color-parse-purple)');
    expect(percentileToken(96)).toBe('var(--color-parse-orange)');
    expect(percentileToken(99)).toBe('var(--color-parse-pink)');
    expect(percentileToken(100)).toBe('var(--color-parse-gold)');
  });

  it('floors a percentile onto the scale: 99.53 is pink, 94.9 orange is not, gold only at 100 (B4)', () => {
    expect(percentileToken(99.53)).toBe('var(--color-parse-pink)');
    expect(percentileToken(99.99)).toBe('var(--color-parse-pink)');
    expect(percentileToken(94.9)).toBe('var(--color-parse-purple)');
    expect(percentileToken(24.6)).toBe('var(--color-parse-grey-text)');
  });

  it('keeps tenths past a minute in the precise form, and names spell schools', () => {
    expect(clockDurationPrecise(61_400)).toBe('1:01.4');
    expect(clockDurationPrecise(6400)).toBe('0:06.4');
    expect(schoolName(1)).toBe('Physical');
    expect(schoolName(36)).toBe('Fire/Shadow');
    expect(schoolName(undefined)).toBe('Physical');
  });

  it('says how far a wipe got when the log showed the boss', () => {
    const base = { kind: 'encounter', kill: false, in_progress: false };
    expect(outcomeLabel({ ...base, boss_health_pct: 23.4 })).toBe('Wipe 23%');
    expect(outcomeLabel({ ...base, boss_health_pct: -1 })).toBe('Wipe');
    expect(outcomeLabel({ ...base, kill: true, boss_health_pct: 0 })).toBe('Kill');
    const trash = { kind: 'trash', kill: false, in_progress: false };
    expect(outcomeLabel({ ...trash, mob_deaths: 177, group_kills: 26 })).toBe(
      '177 mobs died · your group killed 26',
    );
    expect(outcomeLabel({ ...trash, mob_deaths: 1, group_kills: 0 })).toBe(
      '1 mob died · your group killed 0',
    );
    // A trash line from before engine 0.6 has only the double-counted npc_kills: no figure.
    expect(outcomeLabel({ ...trash, npc_kills: 115 } as never)).toBe('');
    expect(joinFacts('Trash · 3:26', '', '0:09')).toBe('Trash · 3:26 · 0:09');
    expect(outcomeLabel({ ...trash, in_progress: true, mob_deaths: 1, group_kills: 1 })).toBe('Live');
    expect(parseTitle(80, 12)).toContain('80th percentile among 12 ranked kills');
    expect(parseTitle(100, 1)).toContain('first of one');
    expect(parseTitle('none')).toContain('no parse');
    expect(ordinal(1)).toBe('1st');
    expect(ordinal(12)).toBe('12th');
    expect(ordinal(23)).toBe('23rd');
  });

  it('colours an ability by its first school, and melee as physical', () => {
    expect(schoolToken(undefined)).toBe('var(--color-school-physical)');
    expect(schoolToken(32)).toBe('var(--color-school-shadow)');
    expect(schoolToken(36)).toBe('var(--color-school-fire)');
  });
});

describe('the phase a fight reached', () => {
  it('is the last phase that started, and nothing when there are none', () => {
    expect(
      phaseReached([
        { name: 'Phase 1', start_ms: 0, end_ms: 14_000 },
        { name: 'Phase 2', start_ms: 14_000, end_ms: 60_000 },
      ]),
    ).toBe('Phase 2');
    expect(phaseReached([])).toBe('');
    expect(phaseReached(undefined)).toBe('');
  });

  it('a wipe’s outcome says which phase it got to; a kill’s does not need one', () => {
    const base = { kind: 'encounter', kill: false, in_progress: false };
    expect(outcomeLabel({ ...base, boss_health_pct: 59 }, 'Phase 3')).toBe('Wipe 59% · in Phase 3');
    expect(outcomeLabel({ ...base, boss_health_pct: 59 })).toBe('Wipe 59%');
    expect(outcomeLabel({ ...base, kill: true }, 'Phase 3')).toBe('Kill');
  });
});

describe('formatAmountLike', () => {
  it('reads a column in one scale once its largest figure is abbreviated', () => {
    expect(formatAmountLike(80_993, 239_200)).toBe('81k');
    expect(formatAmountLike(1_500, 239_200)).toBe('1.5k');
    expect(formatAmountLike(239_200, 239_200)).toBe('239.2k');
    expect(formatAmountLike(80_993, 90_000)).toBe('80,993');
  });
});

describe('the trash parse', () => {
  it('is a muted dash that says parses are for bosses, not an error', () => {
    expect(TRASH_PARSE).toBe('—');
    expect(parseTitle(TRASH_PARSE)).toBe('No parse for trash.');
    expect(parseTitle(TRASH_PARSE)).not.toMatch(/could not be reached/);
  });
});

// Round 7 (wow-player): a sim row is a per-fight average, so a wand a Fire mage fires once
// in a hundred fights read "3 damage, 0 hits, 0.0%" and looked broken.
describe('per-fight averages under one', () => {
  it('reads a row that landed damage on fewer than one hit a fight as "<1"', () => {
    expect(formatLanded(0, 3)).toBe('<1');
    expect(formatLanded(0, 0)).toBe('0');
    expect(formatLanded(29, 2048)).toBe('29');
  });

  it('reads a share above zero that rounds to 0.0% as "<0.1%"', () => {
    expect(formatShare(0.0025)).toBe('<0.1%');
    expect(formatShare(0)).toBe('0.0%');
    expect(formatShare(7.8)).toBe('7.8%');
  });
});
