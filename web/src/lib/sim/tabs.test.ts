// @vitest-environment jsdom
// web/src/lib/sim/tabs.test.ts
import { afterEach, describe, expect, it } from 'vitest';
import { KIND_TITLES, simCopy } from './copy';
import { SIM_TABS, SIM_TAB_ATTR, syncTabHrefs, tabHref, tabStateFor } from './tabs';
import { defaultSimState } from './url';

describe('SIM_TABS', () => {
  it('carries exactly the six entries task-1-brief.md names, in its order', () => {
    expect(SIM_TABS).toEqual([
      { id: 'quick-sim', label: 'Quick Sim', href: '/sim' },
      { id: 'gear', label: 'Top Gear', href: '/sim/gear' },
      { id: 'drops', label: 'Droptimizer', href: '/sim/drops' },
      { id: 'talents', label: 'Talents', href: '/sim/talents' },
      { id: 'weights', label: 'Weights', href: '/sim/weights' },
      { id: 'specs', label: 'Spec support', href: '/sim/specs' },
    ]);
  });

  it('reuses KIND_TITLES for gear and drops, so a page title and its tab can never disagree', () => {
    const gear = SIM_TABS.find((tab) => tab.id === 'gear');
    const drops = SIM_TABS.find((tab) => tab.id === 'drops');
    expect(gear?.label).toBe(KIND_TITLES.gear);
    expect(drops?.label).toBe(KIND_TITLES.drops);
  });

  it('shows "Weights" for the tool whose page title is KIND_TITLES.weights ("Stat weights")', () => {
    const weights = SIM_TABS.find((tab) => tab.id === 'weights');
    expect(weights?.label).toBe(simCopy.tabWeights);
    expect(weights?.label).not.toBe(KIND_TITLES.weights);
  });
});

describe('tabHref', () => {
  it('is the bare path with no character loaded', () => {
    expect(tabHref('/sim/gear', defaultSimState())).toBe('/sim/gear');
  });

  it('carries the state query lib/sim/url.ts builds, never a hand-built one', () => {
    const state = { ...defaultSimState(), source: 'addon' as const, ref: 'FS1:abc' };
    expect(tabHref('/sim/gear', state)).toBe('/sim/gear?source=addon&ref=FS1%3Aabc');
  });
});

describe('tabStateFor', () => {
  it('is the default (empty) state for no character', () => {
    expect(tabStateFor(null)).toEqual(defaultSimState());
  });

  it('carries a loaded character’s own source and ref, nothing else', () => {
    const state = tabStateFor({ kind: 'armory', ref: 'region/ruleset/slug' });
    expect(state).toEqual({ ...defaultSimState(), source: 'armory', ref: 'region/ruleset/slug' });
  });
});

describe('syncTabHrefs', () => {
  afterEach(() => {
    document.body.innerHTML = '';
  });

  function buildStrip(): HTMLElement {
    const nav = document.createElement('nav');
    for (const tab of SIM_TABS) {
      const anchor = document.createElement('a');
      anchor.href = tab.href;
      anchor.setAttribute(SIM_TAB_ATTR, tab.id);
      nav.append(anchor);
    }
    document.body.append(nav);
    return nav;
  }

  it('rewrites every tab to carry the given state’s query', () => {
    const nav = buildStrip();
    syncTabHrefs({ ...defaultSimState(), source: 'build', ref: 'b1' }, document);

    for (const tab of SIM_TABS) {
      const anchor = nav.querySelector<HTMLAnchorElement>(`[${SIM_TAB_ATTR}="${tab.id}"]`);
      expect(anchor?.getAttribute('href')).toBe(`${tab.href}?source=build&ref=b1`);
    }
  });

  it('leaves every href bare for the default (no character) state', () => {
    const nav = buildStrip();
    syncTabHrefs(defaultSimState(), document);

    for (const tab of SIM_TABS) {
      const anchor = nav.querySelector<HTMLAnchorElement>(`[${SIM_TAB_ATTR}="${tab.id}"]`);
      expect(anchor?.getAttribute('href')).toBe(tab.href);
    }
  });

  it('skips a tab missing from the given root instead of throwing', () => {
    const empty = document.createElement('div');
    expect(() => syncTabHrefs(defaultSimState(), empty)).not.toThrow();
  });
});
