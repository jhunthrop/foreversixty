// web/src/lib/scroll-fade.test.ts
// @vitest-environment jsdom
import { describe, expect, it } from 'vitest';
import { revealInStrip, updateScrollFade } from './scroll-fade';

function rect(left: number, right: number): DOMRect {
  return {
    left,
    right,
    top: 0,
    bottom: 0,
    width: right - left,
    height: 0,
    x: left,
    y: 0,
    toJSON: () => ({}),
  };
}

function strip(edge: [number, number], first: [number, number], last: [number, number]): HTMLElement {
  const wrapper = document.createElement('div');
  const scroller = document.createElement('div');
  const a = document.createElement('button');
  const b = document.createElement('button');
  scroller.append(a, b);
  wrapper.append(scroller);
  scroller.getBoundingClientRect = () => rect(...edge);
  a.getBoundingClientRect = () => rect(...first);
  b.getBoundingClientRect = () => rect(...last);
  return scroller;
}

describe('updateScrollFade', () => {
  it('marks only the clipped end', () => {
    const scroller = strip([0, 100], [0, 50], [50, 140]);
    updateScrollFade(scroller);
    expect(scroller.parentElement?.hasAttribute('data-fade-end')).toBe(true);
    expect(scroller.parentElement?.hasAttribute('data-fade-start')).toBe(false);
  });

  it('marks the start once scrolled, and nothing when every tab fits', () => {
    const scrolled = strip([0, 100], [-30, 40], [40, 100]);
    updateScrollFade(scrolled);
    expect(scrolled.parentElement?.hasAttribute('data-fade-start')).toBe(true);
    expect(scrolled.parentElement?.hasAttribute('data-fade-end')).toBe(false);

    const fits = strip([0, 100], [0, 50], [50, 100]);
    updateScrollFade(fits);
    expect(fits.parentElement?.hasAttribute('data-fade-start')).toBe(false);
    expect(fits.parentElement?.hasAttribute('data-fade-end')).toBe(false);
  });
});

describe('revealInStrip', () => {
  const tabAt = (edge: [number, number], tab: [number, number], scrollLeft = 0) => {
    const scroller = strip(edge, tab, tab);
    scroller.scrollLeft = scrollLeft;
    const target = scroller.lastElementChild as HTMLElement;
    target.scrollIntoView = () => {
      throw new Error('scrollIntoView moves the page; the strip scrolls itself');
    };
    revealInStrip(scroller, target);
    return scroller.scrollLeft;
  };

  it('scrolls the strip right by what the tab overhangs its end', () => {
    expect(tabAt([0, 100], [80, 140])).toBe(40);
  });

  it('scrolls the strip left by what the tab overhangs its start', () => {
    expect(tabAt([0, 100], [-30, 40], 50)).toBe(20);
  });

  it('leaves a tab that is already in the strip where it is', () => {
    expect(tabAt([0, 100], [10, 60], 5)).toBe(5);
  });
});
