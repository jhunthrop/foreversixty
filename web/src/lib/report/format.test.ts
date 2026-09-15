// web/src/lib/report/format.test.ts
import { describe, expect, it } from 'vitest';
import {
  classColorVar,
  formatAmount,
  formatClock,
  formatDuration,
  formatDurationPrecise,
  formatPercent,
  formatPerSecond,
  ordinal,
  outcomeLabel,
  parseTitle,
  percentileToken,
  schoolName,
  schoolToken,
} from './format';

describe('report formatting', () => {
  it('formats durations as minutes and seconds, with tenths under a minute', () => {
    expect(formatDuration(6400)).toBe('6.4s');
    expect(formatDuration(40000)).toBe('40.0s');
    expect(formatDuration(107000)).toBe('1:47');
    expect(formatDuration(3_723_000)).toBe('62:03');
    expect(formatDuration(0)).toBe('0.0s');
  });

  it('formats wall-clock times from the engine’s ISO strings', () => {
    expect(formatClock('2026-09-26T20:12:00Z')).toBe('20:12:00');
    expect(formatClock('2026-09-26T20:10:04.6Z')).toBe('20:10:04');
    expect(formatClock('not a time')).toBe('--:--:--');
  });

  it('groups amounts, and abbreviates from a million up', () => {
    expect(formatAmount(0)).toBe('0');
    expect(formatAmount(1484)).toBe('1,484');
    expect(formatAmount(99_999)).toBe('99,999');
    expect(formatAmount(239_230)).toBe('239.2k');
    expect(formatAmount(999_999)).toBe('1000k');
    expect(formatAmount(1_250_000)).toBe('1.25M');
    expect(formatAmount(12_500_000)).toBe('12.5M');
    expect(formatAmount(1_250_000_000)).toBe('1.25B');
  });

  it('divides by the window, not the fight, and never by zero', () => {
    expect(formatPerSecond(4400, 40000)).toBe('110.0');
    expect(formatPerSecond(1484, 6400)).toBe('231.9');
    expect(formatPerSecond(100, 0)).toBe('0.0');
  });

  it('formats percentages to one decimal', () => {
    expect(formatPercent(7.5)).toBe('7.5%');
    expect(formatPercent(100)).toBe('100.0%');
  });

  it('maps a class name to its token, and anything unknown to the body colour', () => {
    expect(classColorVar('Warrior')).toBe('var(--color-class-warrior)');
    expect(classColorVar('Death Knight')).toBe('var(--color-text)');
    expect(classColorVar(undefined)).toBe('var(--color-text)');
  });

  it('maps a parse percentile onto the ladder the community reads', () => {
    expect(percentileToken(3)).toBe('var(--color-parse-grey)');
    expect(percentileToken(40)).toBe('var(--color-parse-green)');
    expect(percentileToken(60)).toBe('var(--color-parse-blue)');
    expect(percentileToken(80)).toBe('var(--color-parse-purple)');
    expect(percentileToken(96)).toBe('var(--color-parse-orange)');
    expect(percentileToken(99)).toBe('var(--color-parse-pink)');
    expect(percentileToken(100)).toBe('var(--color-parse-gold)');
  });

  it('keeps tenths past a minute in the precise form, and names spell schools', () => {
    expect(formatDurationPrecise(61_400)).toBe('1:01.4');
    expect(formatDurationPrecise(6400)).toBe('6.4s');
    expect(schoolName(1)).toBe('Physical');
    expect(schoolName(36)).toBe('Fire/Shadow');
    expect(schoolName(undefined)).toBe('Physical');
  });

  it('says how far a wipe got when the log showed the boss', () => {
    const base = { kind: 'encounter', kill: false, in_progress: false, npc_kills: 0 };
    expect(outcomeLabel({ ...base, boss_health_pct: 23.4 })).toBe('Wipe 23%');
    expect(outcomeLabel({ ...base, boss_health_pct: -1 })).toBe('Wipe');
    expect(outcomeLabel({ ...base, kill: true, boss_health_pct: 0 })).toBe('Kill');
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
