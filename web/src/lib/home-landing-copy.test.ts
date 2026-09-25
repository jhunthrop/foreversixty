// web/src/lib/home-landing-copy.test.ts
import { describe, expect, it } from 'vitest';
import { homeCompanionRow, homeGuildPanelCopy, homeHeroCopy, homeProductPanels } from './home-landing-copy';

function nonEmpty(value: string): boolean {
  return value.trim().length > 0;
}

describe('home-landing-copy', () => {
  it('carries a non-empty hero eyebrow, headline and addon-alt link', () => {
    expect(nonEmpty(homeHeroCopy.eyebrow)).toBe(true);
    expect(nonEmpty(homeHeroCopy.headline)).toBe(true);
    expect(homeHeroCopy.addonAltHref).toBe('/setup#paste');
  });

  it('lists exactly the four product panels, in spec order, each fully populated', () => {
    expect(homeProductPanels.map((p) => p.label)).toEqual(['Planner', 'Simulator', 'Logs', 'Rankings']);
    for (const panel of homeProductPanels) {
      expect(nonEmpty(panel.sentence)).toBe(true);
      expect(nonEmpty(panel.linkLabel)).toBe(true);
      expect(panel.href.startsWith('/')).toBe(true);
    }
  });

  it('formats the guild progression sentence honestly from real counts', () => {
    expect(homeGuildPanelCopy.progressionOf(3, 12)).toBe('3 of 12 bosses down');
    expect(homeGuildPanelCopy.progressionOf(0, 0)).toBe('0 of 0 bosses down');
  });

  it('points the claim CTA at Battle.net sign-in, since no guild directory page exists', () => {
    expect(homeGuildPanelCopy.claimHref).toMatch(/\/v1\/auth\/battlenet\/start\?next=/);
  });

  it('collapses the addon and companion row into one "Get set up" card', () => {
    expect(homeCompanionRow).toHaveLength(1);
    expect(homeCompanionRow[0].title).toBe('Get set up');
    expect(homeCompanionRow[0].href).toBe('/setup');
  });
});
