import { describe, expect, it } from 'vitest';
import { characterHistoryFixture } from '../../fixtures/character-history';
import { collapseScorePoints } from './history-chart';
import {
  COLLAPSED_ROWS,
  MAX_ROWS,
  canStep,
  initialSelection,
  killListFights,
  killMarkKeys,
  listWindow,
  markedKeys,
  nearestIndex,
  scoreListFights,
  scoreMarkKeys,
  select,
  shownIndex,
  step,
  withHover,
} from './history-select';
import type { HistoryFight } from './history-api';

const history = characterHistoryFixture();

function rows(count: number): HistoryFight[] {
  return Array.from({ length: count }, (_, i) => ({
    ...history.fights[0],
    report_id: `r${i}`,
    fight_index: 1,
  }));
}

describe('selection', () => {
  it('starts on the newest point', () => {
    expect(initialSelection(8)).toEqual({ selected: 7, hover: null });
    expect(initialSelection(0).selected).toBe(0);
  });

  it('steps one at a time, jumps to the ends, and stays inside', () => {
    const start = initialSelection(5);
    expect(step(start, 'earlier', 5).selected).toBe(3);
    expect(step(start, 'later', 5).selected).toBe(4);
    expect(step(start, 'first', 5).selected).toBe(0);
    expect(step(step(start, 'first', 5), 'newest', 5).selected).toBe(4);
    expect(step(step(start, 'first', 5), 'earlier', 5).selected).toBe(0);
  });

  it('disables the step buttons at the first and newest point', () => {
    const first = select(initialSelection(5), 0, 5);
    expect(canStep(first, 'earlier', 5)).toBe(false);
    expect(canStep(first, 'later', 5)).toBe(true);
    expect(canStep(initialSelection(5), 'later', 5)).toBe(false);
  });

  it('returns new values and never mutates', () => {
    const before = initialSelection(5);
    const after = select(before, 1, 5);
    expect(before.selected).toBe(4);
    expect(after).not.toBe(before);
  });

  it('previews on hover without moving the selection, and restores on leave', () => {
    const hovered = withHover(initialSelection(5), 1, 5);
    expect(shownIndex(hovered)).toBe(1);
    expect(hovered.selected).toBe(4);
    expect(shownIndex(withHover(hovered, null, 5))).toBe(4);
  });

  it('finds the point nearest an x', () => {
    expect(nearestIndex([10, 50, 200], 60)).toBe(1);
    expect(nearestIndex([10, 50, 200], 120)).toBe(1);
    expect(nearestIndex([10, 50, 200], 500)).toBe(2);
  });
});

describe('marked rows', () => {
  it('marks every fight of the selected score day', () => {
    const days = collapseScorePoints(history.score);
    const marks = markedKeys(scoreMarkKeys(days), initialSelection(days.length));
    expect(marks).toEqual(new Set(['rpt000000008:7', 'rpt000000009:7']));
  });

  it('marks the one selected kill, keyed report_id:fight_index', () => {
    const encounter = history.encounters[0];
    const marks = markedKeys(killMarkKeys(encounter), select(initialSelection(5), 0, 5));
    expect(marks).toEqual(new Set(['rpt000000004:7']));
  });

  it('follows the selection, not the hover', () => {
    const days = collapseScorePoints(history.score);
    const hovered = withHover(initialSelection(days.length), 0, days.length);
    expect(markedKeys(scoreMarkKeys(days), hovered)).toEqual(new Set(['rpt000000008:7', 'rpt000000009:7']));
  });
});

describe('the lists', () => {
  it('holds the fights behind the score, and the kills of one boss', () => {
    expect(scoreListFights(history)).toHaveLength(9);
    expect(killListFights(history, history.encounters[0]).map((f) => f.report_id)).toEqual([
      'rpt000000009',
      'rpt000000007',
      'rpt000000006',
      'rpt000000005',
      'rpt000000004',
    ]);
  });

  it('shows 8 rows and a Show all button above that', () => {
    const window = listWindow(rows(31), false, new Set());
    expect(window.rows).toHaveLength(COLLAPSED_ROWS);
    expect(window).toMatchObject({ shownTotal: 31, button: 'show-all', needsExpand: false });
  });

  it('shows no button with 8 rows or fewer', () => {
    expect(listWindow(rows(8), false, new Set()).button).toBe('none');
  });

  it('expands in place and offers Show fewer', () => {
    const window = listWindow(rows(31), true, new Set());
    expect(window.rows).toHaveLength(31);
    expect(window.button).toBe('show-fewer');
  });

  it('stops at 100 rows and names 100', () => {
    const window = listWindow(rows(130), true, new Set());
    expect(window.rows).toHaveLength(MAX_ROWS);
    expect(window.shownTotal).toBe(MAX_ROWS);
  });

  it('asks to expand when a marked row is hidden', () => {
    const all = rows(31);
    expect(listWindow(all, false, new Set(['r20:1'])).needsExpand).toBe(true);
    expect(listWindow(all, false, new Set(['r2:1'])).needsExpand).toBe(false);
  });
});
