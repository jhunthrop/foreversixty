import { describe, expect, it } from 'vitest';
import { landingPageView } from './landing-page-view';

const NONE = { url: null, stored: null };

describe('landingPageView', () => {
  it('signed out: 20 to 29, no own range, no faction, nothing is Yours', () => {
    expect(landingPageView(NONE, null)).toEqual({
      selectedBand: 20,
      youBand: undefined,
      faction: undefined,
      yoursSpecKey: undefined,
    });
  });

  it('signed out, the address bar wins, then the stored pick', () => {
    expect(landingPageView({ url: 50, stored: 30 }, null).selectedBand).toBe(50);
    expect(landingPageView({ url: null, stored: 30 }, null).selectedBand).toBe(30);
  });

  it('with a character: their band is selected, their spec is Yours, their faction rides the links', () => {
    expect(landingPageView(NONE, { level: 47, faction: 'horde', specKey: 'warrior-fury' })).toEqual({
      selectedBand: 40,
      youBand: 40,
      faction: 'horde',
      yoursSpecKey: 'warrior-fury',
    });
  });

  it('a stored 60 never beats a level 24 character (owner 2026-11-07)', () => {
    const view = landingPageView(
      { url: null, stored: 60 },
      { level: 24, faction: 'alliance', specKey: 'hunter-beast-mastery' },
    );
    expect(view.selectedBand).toBe(20);
  });

  it('keeps the own range while the address bar selects another', () => {
    const view = landingPageView(
      { url: 60, stored: null },
      { level: 47, faction: 'alliance', specKey: 'warrior-fury' },
    );
    expect(view.selectedBand).toBe(60);
    expect(view.youBand).toBe(40);
  });

  it('a character under 20 selects 20 to 29 with no own range', () => {
    const view = landingPageView(NONE, { level: 12, faction: 'alliance', specKey: 'warrior-fury' });
    expect(view.selectedBand).toBe(20);
    expect(view.youBand).toBeUndefined();
  });

  it('a character with no spec yet marks no row', () => {
    expect(
      landingPageView(NONE, { level: 30, faction: 'alliance', specKey: undefined }).yoursSpecKey,
    ).toBeUndefined();
  });
});
