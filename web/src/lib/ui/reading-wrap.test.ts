// @vitest-environment jsdom
import { afterEach, beforeEach, describe, expect, it, vi } from 'vitest';
import { movableTail, rewrapAround, unwrapSection, wrapSection, wrapUnits } from './reading-wrap';

const ROW_PX = 50;
const PROSE_PX = 100;

/** A section: a lane of `rows` rows in a labelled block (the list marked), the prose beside it. */
function section(rows: number, extra = ''): HTMLElement {
  document.body.innerHTML = `
    <div class="ui-rd" data-reading>
      <div class="ui-rd-lane" data-reading-lane>
        <div class="block" data-testid="block"><h3 class="head">Key talents</h3>
          <div class="list" data-wrap-list>${Array.from({ length: rows }, (_, i) => `<div class="row" data-testid="r${i}">${i}</div>`).join('')}</div>
        </div>${extra}
      </div>
      <div class="ui-rd-prose" data-reading-prose><p>Why</p><div class="ui-rd-more" data-reading-more></div></div>
    </div>`;
  return document.querySelector<HTMLElement>('.ui-rd')!;
}

/** Heights: the lane is its rows; the prose column is the prose and the rows wrapped under it. */
function bottomOf(element: Element): number {
  const count = (root: Element | null): number => root?.querySelectorAll('.row, .note').length ?? 0;
  if (element.classList.contains('ui-rd-lane')) return count(element) * ROW_PX;
  if (element.classList.contains('ui-rd-prose'))
    return (
      Number((element as HTMLElement).dataset.px ?? PROSE_PX) +
      count(element.querySelector('.ui-rd-more')) * ROW_PX
    );
  return 0;
}

beforeEach(() => {
  vi.spyOn(Element.prototype, 'getBoundingClientRect').mockImplementation(function (this: Element) {
    return { bottom: bottomOf(this), height: bottomOf(this) } as DOMRect;
  });
});
afterEach(() => vi.restoreAllMocks());

const rowIds = (root: Element): string[] =>
  [...root.querySelectorAll('.row')].map((row) => row.textContent ?? '');

describe('reading-wrap', () => {
  it('lists the rows under every marked list, opening the blocks that hold one', () => {
    const rd = section(3);
    const units = wrapUnits(rd.querySelector('.ui-rd-lane')!);
    expect(units.map((unit) => unit.className)).toEqual(['head', 'row', 'row', 'row']);
    expect(movableTail(units).map((unit) => unit.textContent)).toEqual(['2', '1', '0']);
  });

  it('wraps the trailing rows under the prose until the columns end within a row', () => {
    const rd = section(20);
    wrapSection(rd, 1440);
    const lane = rowIds(rd.querySelector('.ui-rd-lane')!);
    const more = rowIds(rd.querySelector('.ui-rd-more')!);
    expect([...lane, ...more]).toEqual(Array.from({ length: 20 }, (_, i) => String(i)));
    expect(Math.abs(lane.length * ROW_PX - (PROSE_PX + more.length * ROW_PX))).toBeLessThanOrEqual(ROW_PX);
    // The mirror keeps the list's classes and drops its test id; the block's label stays put.
    expect(rd.querySelector('.ui-rd-more .block .list .row')).not.toBeNull();
    expect(rd.querySelectorAll('[data-testid="block"]')).toHaveLength(1);
    expect(rd.querySelectorAll('.head')).toHaveLength(1);
  });

  it('leaves a section alone under 1280 or under a 200 px gap', () => {
    const rd = section(20);
    wrapSection(rd, 1279);
    expect(rd.querySelector('.ui-rd-more')!.children).toHaveLength(0);
    const short = section(5);
    wrapSection(short, 1440);
    expect(short.querySelector('.ui-rd-more')!.children).toHaveLength(0);
  });

  it('puts every row back in order on unwrap', () => {
    const rd = section(20);
    wrapSection(rd, 1440);
    unwrapSection(rd);
    expect(rowIds(rd.querySelector('.ui-rd-lane')!)).toEqual(Array.from({ length: 20 }, (_, i) => String(i)));
    expect(rd.querySelector('.ui-rd-more')!.children).toHaveLength(0);
  });

  it('never moves an island, nor anything above it', () => {
    const rd = section(20, '<div class="note"><astro-island></astro-island></div>');
    wrapSection(rd, 1440);
    expect(rd.querySelector('.ui-rd-more')!.children).toHaveLength(0);
  });

  it('re-wraps around a lane that swaps its rows', () => {
    const rd = section(20);
    vi.stubGlobal('innerWidth', 1440);
    wrapSection(rd, 1440);
    rewrapAround(rd.querySelector('.list')!, () => {
      rd.querySelector('.list')!.innerHTML = Array.from(
        { length: 12 },
        (_, i) => `<div class="row">n${i}</div>`,
      ).join('');
    });
    const all = [...rowIds(rd.querySelector('.ui-rd-lane')!), ...rowIds(rd.querySelector('.ui-rd-more')!)];
    expect(all).toEqual(Array.from({ length: 12 }, (_, i) => `n${i}`));
    vi.unstubAllGlobals();
  });

  it('pins a short lane beside long prose while it fits the viewport, and lets it go on unwrap', () => {
    const rd = section(4);
    rd.querySelector<HTMLElement>('.ui-rd-prose')!.dataset.px = '900';
    wrapSection(rd, 1440, 800);
    expect(rd.hasAttribute('data-lane-sticky')).toBe(true);
    unwrapSection(rd);
    expect(rd.hasAttribute('data-lane-sticky')).toBe(false);
    wrapSection(rd, 1440, 300);
    expect(rd.hasAttribute('data-lane-sticky')).toBe(false);
    wrapSection(rd, 1279, 800);
    expect(rd.hasAttribute('data-lane-sticky')).toBe(false);
  });
});
