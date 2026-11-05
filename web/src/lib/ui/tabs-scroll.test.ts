// The tab strip's sideways scroll (spec 2026-11-05 §3.6): the previous tab shows 48 px, never a
// sliver of the one before it, and the first tab rests at 0.
import { describe, expect, it } from 'vitest';
import { TAB_PEEK_PX, TAB_STRIP, scrollTabIntoView } from './tabs';

function fakes(offsetLeft: number, hasPrevious: boolean): { strip: HTMLElement; tab: HTMLElement } {
  return {
    strip: { scrollLeft: 999 } as unknown as HTMLElement,
    tab: { offsetLeft, previousElementSibling: hasPrevious ? {} : null } as unknown as HTMLElement,
  };
}

describe('scrollTabIntoView', () => {
  it('leaves 48 px of the previous tab beside the current one', () => {
    const { strip, tab } = fakes(300, true);
    scrollTabIntoView(strip, tab);
    expect(strip.scrollLeft).toBe(300 - TAB_PEEK_PX);
    expect(TAB_PEEK_PX).toBe(48);
  });

  it('rests the first tab at 0', () => {
    const { strip, tab } = fakes(0, false);
    scrollTabIntoView(strip, tab);
    expect(strip.scrollLeft).toBe(0);
  });

  it('never asks for a negative scroll on a second tab near the start', () => {
    const { strip, tab } = fakes(30, true);
    scrollTabIntoView(strip, tab);
    expect(strip.scrollLeft).toBe(0);
  });

  it('the strip declares the same peek for the browser', () => {
    expect(TAB_STRIP.strip).toContain('[scroll-padding-inline:48px]');
  });
});
