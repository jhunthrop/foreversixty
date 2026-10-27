// web/src/lib/guild/report-view.test.ts
// The report's orderings and caps, the copy state machine and the report URL.
import { afterEach, beforeEach, describe, expect, it, vi } from 'vitest';
import { attendanceCopy } from './attendance-copy';
import { ReportCopy } from './report-copy.svelte';
import { biggestMoves, consumableOrder, deathOrder, lootOrder, reportUrlOf } from './report-view';
import { raidReport } from '../../test-support/report-fixtures';

const death = (name: string, cls: string, wipes: number, kills: number) => ({
  character_key: name,
  name,
  class: cls,
  on_wipes: wipes,
  on_kills: kills,
});

describe('deathOrder', () => {
  it('puts the most deaths first, then the game order', () => {
    const rows = [death('Zed', 'mage', 1, 0), death('Abe', 'warrior', 1, 0), death('Mo', 'rogue', 2, 2)];
    expect(deathOrder(rows).map((r) => r.name)).toEqual(['Mo', 'Abe', 'Zed']);
  });
});

describe('consumableOrder', () => {
  const row = (name: string, before: 'stocked' | 'short' | 'unknown', used: boolean) => ({
    character_key: name,
    name,
    class: 'mage',
    before,
    flask: used,
    food: false,
    potions: 0,
  });
  it('reads Short with nothing used first, then Short, then the rest', () => {
    const rows = [row('A', 'stocked', false), row('B', 'short', true), row('C', 'short', false)];
    expect(consumableOrder(rows, true).map((r) => r.name)).toEqual(['C', 'B', 'A']);
  });
});

describe('consumableOrder without buff data', () => {
  it('ranks every Short row alike, then the game order', () => {
    const row = (name: string, before: 'stocked' | 'short') => ({
      character_key: name,
      name,
      class: 'mage',
      before,
      flask: null,
      food: null,
      potions: null,
    });
    const rows = [row('B', 'short'), row('A', 'stocked'), row('A2', 'short')];
    expect(consumableOrder(rows, false).map((r) => r.name)).toEqual(['A2', 'B', 'A']);
  });
});

describe('lootOrder', () => {
  it('sorts by quality, then name', () => {
    const award = (item_name: string, quality: number) => ({
      ...raidReport().loot[0].awards[0],
      item_name,
      quality,
    });
    const sorted = lootOrder([award('B', 3), award('A', 3), award('Z', 4)]);
    expect(sorted.map((a) => a.item_name)).toEqual(['Z', 'A', 'B']);
  });
});

describe('biggestMoves', () => {
  it('keeps the five largest absolute moves and drops flat ones', () => {
    const move = (name: string, before: number, after: number) => ({
      ...raidReport().moves.rows[0],
      character_key: name,
      name,
      before,
      after,
    });
    const rows = [
      move('a', 100, 101),
      move('b', 100, 160),
      move('c', 100, 40),
      move('d', 100, 100),
      move('e', 100, 130),
      move('f', 100, 120),
      move('g', 100, 110),
    ];
    expect(biggestMoves(rows).map((r) => r.name)).toEqual(['b', 'c', 'e', 'f', 'g']);
  });
});

describe('reportUrlOf', () => {
  it('is the raid page with the report anchor', () => {
    expect(reportUrlOf({ origin: 'https://foreversixty.gg', pathname: '/guild/a/b/c/raids/7' })).toBe(
      'https://foreversixty.gg/guild/a/b/c/raids/7#report',
    );
  });
});

describe('ReportCopy', () => {
  beforeEach(() => vi.useFakeTimers());
  afterEach(() => vi.useRealTimers());

  it('says Copied, then goes quiet after the visible time', async () => {
    const writeText = vi.fn(() => Promise.resolve());
    const copy = new ReportCopy(writeText);
    await copy.copy('text');
    expect(copy.state).toBe('copied');
    expect(writeText).toHaveBeenCalledWith('text');
    vi.advanceTimersByTime(4000);
    expect(copy.state).toBe('idle');
  });

  it('keeps the text for the fallback when the clipboard refuses', async () => {
    const copy = new ReportCopy(() => Promise.reject(new Error('no')));
    await copy.copy('text');
    expect(copy.state).toBe('failed');
    expect(copy.text).toBe('text');
    expect(attendanceCopy.errors.copyFailed).toContain('Could not copy');
  });
});
