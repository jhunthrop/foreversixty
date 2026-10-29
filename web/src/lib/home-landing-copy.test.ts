// web/src/lib/home-landing-copy.test.ts
import { describe, expect, it } from 'vitest';
import { homeCompanionRow, homeGetSetUpCopy, homeHeroCopy } from './home-landing-copy';

function nonEmpty(value: string): boolean {
  return value.trim().length > 0;
}

describe('home-landing-copy', () => {
  it('carries a non-empty hero eyebrow, headline and addon-alt link', () => {
    expect(nonEmpty(homeHeroCopy.eyebrow)).toBe(true);
    expect(nonEmpty(homeHeroCopy.headline)).toBe(true);
    expect(homeHeroCopy.addonAltHref).toBe('/setup#paste');
  });

  it('collapses the addon and companion row into one "Get set up" card', () => {
    expect(homeCompanionRow).toHaveLength(1);
    expect(homeCompanionRow[0].title).toBe('Get set up');
    expect(homeCompanionRow[0].href).toBe('/setup');
    expect(homeCompanionRow[0].buttonLabel).toBe('Open setup');
    expect(homeCompanionRow[0].description).toBe(
      'Sign in, install the addon, and run the companion — one page.',
    );
  });

  it('carries the exact collapsed "Get set up" sentence, interpolating the real sync time', () => {
    // Home rebuild spec §3.B.6: a function, not a fixed string, so this card's sync time
    // can never disagree with the hero's own sync line.
    expect(homeGetSetUpCopy.bothDone('4 minutes ago')).toBe(
      'Signed in and addon linked · synced 4 minutes ago → Set up the companion',
    );
    expect(nonEmpty(homeGetSetUpCopy.addonRemaining)).toBe(true);
    expect(nonEmpty(homeGetSetUpCopy.companionRemaining)).toBe(true);
  });
});
