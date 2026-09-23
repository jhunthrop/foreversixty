// web/src/lib/home-landing-copy.test.ts
import { describe, expect, it } from 'vitest';
import {
  homeCompanionRow,
  homeGuildPanelCopy,
  homeHeroCopy,
  homeProductPanels,
  homeReferenceCopy,
} from './home-landing-copy';

function nonEmpty(value: string): boolean {
  return value.trim().length > 0;
}

describe('home-landing-copy', () => {
  it('carries a non-empty hero eyebrow, headline and addon-alt link', () => {
    expect(nonEmpty(homeHeroCopy.eyebrow)).toBe(true);
    expect(nonEmpty(homeHeroCopy.headline)).toBe(true);
    expect(homeHeroCopy.addonAltHref).toBe('/addon#paste');
  });

  it('lists exactly the four product panels, in spec order, each fully populated', () => {
    expect(homeProductPanels.map((p) => p.label)).toEqual(['Planner', 'Simulator', 'Logs', 'Rankings']);
    for (const panel of homeProductPanels) {
      expect(nonEmpty(panel.sentence)).toBe(true);
      expect(nonEmpty(panel.linkLabel)).toBe(true);
      expect(panel.href.startsWith('/')).toBe(true);
    }
  });

  it('carries the reference band sentence', () => {
    expect(homeReferenceCopy.sentence).toBe('Every fact dated and sourced.');
  });

  it('formats the guild progression sentence honestly from real counts', () => {
    expect(homeGuildPanelCopy.progressionOf(3, 12)).toBe('3 of 12 bosses down');
    expect(homeGuildPanelCopy.progressionOf(0, 0)).toBe('0 of 0 bosses down');
  });

  it('points the claim CTA at Battle.net sign-in, since no guild directory page exists', () => {
    expect(homeGuildPanelCopy.claimHref).toMatch(/\/v1\/auth\/battlenet\/start\?next=/);
  });

  it('lists exactly the addon and companion cards', () => {
    expect(homeCompanionRow.map((c) => c.title)).toEqual(['The addon', 'The companion']);
    expect(homeCompanionRow.map((c) => c.href)).toEqual(['/addon', '/logs#companion']);
  });
});
