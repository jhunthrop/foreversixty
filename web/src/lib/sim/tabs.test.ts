// @vitest-environment jsdom
// web/src/lib/sim/tabs.test.ts
import { afterEach, describe, expect, it } from 'vitest';
import { KIND_TITLES, simCopy } from './copy';
import { SIM_TABS, SIM_TAB_ATTR, syncTabHrefs, tabHref, tabStateFor } from './tabs';
import { defaultSimState, MAX_CODE } from './url';

describe('SIM_TABS', () => {
  it('carries exactly the six entries task-1-brief.md names, in its order', () => {
    expect(SIM_TABS).toEqual([
      { id: 'quick-sim', label: 'Quick Sim', href: '/sim', supportsCode: true },
      { id: 'gear', label: 'Top Gear', href: '/sim/gear', supportsCode: false },
      { id: 'drops', label: 'Droptimizer', href: '/sim/drops', supportsCode: false },
      { id: 'talents', label: 'Talents', href: '/sim/talents', supportsCode: false },
      { id: 'weights', label: 'Weights', href: '/sim/weights', supportsCode: false },
      { id: 'specs', label: 'Spec support', href: '/sim/specs', supportsCode: true },
    ]);
  });

  // Fix round 1, Finding A: only SimView.svelte's store (`store.svelte.ts`'s
  // `fromPlannerCode`) ever reads `?code=`; ToolsView.svelte / bulk-store.svelte.ts never
  // have. `syncTabHrefs` relies on exactly this split to know which tabs a fallback code is
  // safe to offer.
  it('marks only the two SimView-served tabs as able to bootstrap from ?code=', () => {
    const supportsCode = SIM_TABS.filter((tab) => tab.supportsCode).map((tab) => tab.id);
    expect(supportsCode).toEqual(['quick-sim', 'specs']);
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

  it('ignores a fallback code when the source already carries a usable ref', () => {
    const state = tabStateFor({ kind: 'build', ref: 'b1' }, 'FS1:1:warrior:orc:0/0/0:');
    expect(state).toEqual({ ...defaultSimState(), source: 'build', ref: 'b1' });
  });

  // Fix round 1, Finding A: an addon paste or a `?code=` link both stamp `ref: ''`
  // (sources.ts, store.svelte.ts's `fromPlannerCode`) -- `?source=addon` alone would be an
  // inert query that looks like it restores the character and does not, so the tab strip
  // falls back to a fresh FS1 code the caller derives (SimView.svelte's `fallbackTabCode`,
  // bulk-store.svelte.ts's `characterCode`).
  it('falls back to a `code` state when the source has no ref to give it', () => {
    const state = tabStateFor({ kind: 'addon', ref: '' }, 'FS1:1:warrior:orc:0/0/0:head=1');
    expect(state).toEqual({ ...defaultSimState(), code: 'FS1:1:warrior:orc:0/0/0:head=1' });
  });

  it('is bare, not an inert `?source=addon`, when there is no fallback code either', () => {
    expect(tabStateFor({ kind: 'addon', ref: '' })).toEqual(defaultSimState());
    expect(tabStateFor({ kind: 'manual', ref: '' }, '')).toEqual(defaultSimState());
    expect(tabStateFor({ kind: 'manual', ref: '' }, null)).toEqual(defaultSimState());
  });

  // A link this function would not itself honour (parseSimState silently drops an
  // over-budget code, url.ts's own `bounded`) must never be emitted as if it worked.
  it('is bare rather than emitting a code past MAX_CODE', () => {
    const tooLong = 'FS1:'.padEnd(MAX_CODE + 1, '0');
    expect(tooLong.length).toBeGreaterThan(MAX_CODE);
    expect(tabStateFor({ kind: 'addon', ref: '' }, tooLong)).toEqual(defaultSimState());
  });

  it('accepts a code exactly at MAX_CODE', () => {
    const atLimit = 'FS1:'.padEnd(MAX_CODE, '0');
    expect(atLimit.length).toBe(MAX_CODE);
    expect(tabStateFor({ kind: 'manual', ref: '' }, atLimit)).toEqual({
      ...defaultSimState(),
      code: atLimit,
    });
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

  function hrefOf(nav: HTMLElement, id: string): string | null {
    return nav.querySelector<HTMLAnchorElement>(`[${SIM_TAB_ATTR}="${id}"]`)?.getAttribute('href') ?? null;
  }

  it('rewrites every tab to carry a ref-bearing source’s query', () => {
    const nav = buildStrip();
    syncTabHrefs({ kind: 'build', ref: 'b1' }, null, document);

    for (const tab of SIM_TABS) expect(hrefOf(nav, tab.id)).toBe(`${tab.href}?source=build&ref=b1`);
  });

  it('leaves every href bare for no character at all', () => {
    const nav = buildStrip();
    syncTabHrefs(null, null, document);

    for (const tab of SIM_TABS) expect(hrefOf(nav, tab.id)).toBe(tab.href);
  });

  // Fix round 1, Finding A: the fallback code is real ("Quick Sim", "Spec support"), never
  // an inert query the destination cannot consume ("Top Gear", "Droptimizer", "Talents",
  // "Weights").
  it('offers a fallback code only to the tabs whose own destination can bootstrap from it', () => {
    const nav = buildStrip();
    syncTabHrefs({ kind: 'addon', ref: '' }, 'FS1:1:warrior:orc:0/0/0:head=1', document);

    expect(hrefOf(nav, 'quick-sim')).toBe('/sim?code=FS1%3A1%3Awarrior%3Aorc%3A0%2F0%2F0%3Ahead%3D1');
    expect(hrefOf(nav, 'specs')).toBe('/sim/specs?code=FS1%3A1%3Awarrior%3Aorc%3A0%2F0%2F0%3Ahead%3D1');
    for (const id of ['gear', 'drops', 'talents', 'weights']) expect(hrefOf(nav, id)).toBe(`/sim/${id}`);
  });

  it('is bare on every tab for a ref-less source with no fallback code either', () => {
    const nav = buildStrip();
    syncTabHrefs({ kind: 'manual', ref: '' }, null, document);

    for (const tab of SIM_TABS) expect(hrefOf(nav, tab.id)).toBe(tab.href);
  });

  it('skips a tab missing from the given root instead of throwing', () => {
    const empty = document.createElement('div');
    expect(() => syncTabHrefs(null, null, empty)).not.toThrow();
  });
});
