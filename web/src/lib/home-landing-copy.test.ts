// web/src/lib/home-landing-copy.test.ts
import { describe, expect, it } from 'vitest';
import { homeCompanionRow, homeGetSetUpCopy, homeHeroCopy, homeProductPanels } from './home-landing-copy';

function nonEmpty(value: string): boolean {
  return value.trim().length > 0;
}

describe('home-landing-copy', () => {
  it('carries a non-empty hero eyebrow, headline and addon-alt link', () => {
    expect(nonEmpty(homeHeroCopy.eyebrow)).toBe(true);
    expect(nonEmpty(homeHeroCopy.headline)).toBe(true);
    expect(homeHeroCopy.addonAltHref).toBe('/setup#paste');
  });

  it('lists exactly the five product panels, Guides last, each fully populated', () => {
    expect(homeProductPanels.map((p) => p.label)).toEqual([
      'Planner',
      'Simulator',
      'Logs',
      'Rankings',
      'Guides',
    ]);
    for (const panel of homeProductPanels) {
      expect(nonEmpty(panel.sentence)).toBe(true);
      expect(nonEmpty(panel.linkLabel)).toBe(true);
      expect(panel.href.startsWith('/')).toBe(true);
    }
  });

  it('collapses the addon and companion row into one "Get set up" card', () => {
    expect(homeCompanionRow).toHaveLength(1);
    expect(homeCompanionRow[0].title).toBe('Get set up');
    expect(homeCompanionRow[0].href).toBe('/setup');
  });

  it('carries the exact collapsed "Get set up" sentence and the two remaining-step words', () => {
    expect(homeGetSetUpCopy.bothDone).toBe('Signed in and addon linked · Set up the companion →');
    expect(nonEmpty(homeGetSetUpCopy.addonRemaining)).toBe(true);
    expect(nonEmpty(homeGetSetUpCopy.companionRemaining)).toBe(true);
  });
});
