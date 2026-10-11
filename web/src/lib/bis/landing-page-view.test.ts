// web/src/lib/bis/landing-page-view.test.ts
import { describe, expect, it } from 'vitest';
import { landingPageView } from './landing-page-view';

describe('landingPageView', () => {
  it('signed out: 20 to 29, no you mark, no faction, nothing is Yours, the pick-a-class heading', () => {
    expect(landingPageView(null, null)).toEqual({
      selectedBand: 20,
      youBand: undefined,
      faction: undefined,
      yoursSpecKey: undefined,
      gridHeading: 'Pick a class and spec',
    });
  });

  it('signed out, the viewer’s own pick wins', () => {
    expect(landingPageView(50, null).selectedBand).toBe(50);
  });

  it('with a character: their band is selected and marked, their spec is Yours, their faction rides the links', () => {
    expect(landingPageView(null, { level: 47, faction: 'horde', specKey: 'warrior-fury' })).toEqual({
      selectedBand: 40,
      youBand: 40,
      faction: 'horde',
      yoursSpecKey: 'warrior-fury',
      gridHeading: 'Another class or spec',
    });
  });

  it('keeps the you mark on the character’s band while the viewer selects another', () => {
    const view = landingPageView(60, { level: 47, faction: 'alliance', specKey: 'warrior-fury' });
    expect(view.selectedBand).toBe(60);
    expect(view.youBand).toBe(40);
  });

  it('a character under 20 selects 20 to 29 with no you mark', () => {
    const view = landingPageView(null, { level: 12, faction: 'alliance', specKey: 'warrior-fury' });
    expect(view.selectedBand).toBe(20);
    expect(view.youBand).toBeUndefined();
  });

  it('a character with no spec yet marks no row', () => {
    expect(landingPageView(null, { level: 30, faction: 'alliance', specKey: undefined }).yoursSpecKey).toBeUndefined();
  });
});
