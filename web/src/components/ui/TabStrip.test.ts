// TabStrip and its static twin (site alignment spec 2026-10-29 D3, S7): states, a11y and the
// keyboard. The key handling is lib/ui/tabs.ts's tabIdForKey, which the component calls as is;
// a real browser presses the keys in tests/e2e/tab-strip.spec.ts.
import { experimental_AstroContainer as AstroContainer } from 'astro/container';
import { JSDOM } from 'jsdom';
import { render } from 'svelte/server';
import { describe, expect, it } from 'vitest';
import { TAB_STRIP, tabClass, tabIdForKey, tabIndexForKey, type TabItem } from '../../lib/ui/tabs';
import TabStripAstro from './TabStrip.astro';
import TabStrip from './TabStrip.svelte';

const TABS: readonly TabItem[] = [
  { id: 'overview', label: 'Overview', testid: 'tab-overview' },
  { id: 'roster', label: 'Roster', suffix: ' · read-only' },
  { id: 'raids', label: 'Raids', disabled: true },
  { id: 'loot', label: 'Loot', href: '#loot' },
];

function strip(props: Record<string, unknown>): Document {
  const body = render(TabStrip, {
    props: { tabs: TABS, active: 'roster', label: 'Guild sections', ...props },
  }).body;
  return new JSDOM(body).window.document;
}

const select = (): void => undefined;

describe('tabIndexForKey', () => {
  const plain = [{}, {}, {}, {}];
  it.each([
    ['ArrowRight', 0, 1],
    ['ArrowRight', 3, 0],
    ['ArrowLeft', 0, 3],
    ['ArrowLeft', 2, 1],
    ['Home', 2, 0],
    ['End', 1, 3],
  ])('%s from %i goes to %i, wrapping', (key, from, to) => {
    expect(tabIndexForKey(key, from, plain)).toBe(to);
  });

  it('skips a disabled tab both ways and at the ends', () => {
    const tabs = [{ disabled: true }, {}, { disabled: true }, {}, { disabled: true }];
    expect(tabIndexForKey('ArrowRight', 1, tabs)).toBe(3);
    expect(tabIndexForKey('ArrowLeft', 3, tabs)).toBe(1);
    expect(tabIndexForKey('ArrowRight', 3, tabs)).toBe(1);
    expect(tabIndexForKey('Home', 3, tabs)).toBe(1);
    expect(tabIndexForKey('End', 1, tabs)).toBe(3);
  });

  it('lets every other key through, and has nowhere to go when nothing is enabled', () => {
    expect(tabIndexForKey('Enter', 0, plain)).toBeNull();
    expect(tabIndexForKey('ArrowDown', 0, plain)).toBeNull();
    expect(tabIndexForKey('ArrowRight', 0, [{ disabled: true }])).toBeNull();
  });

  it('names the tab by id from the current one', () => {
    expect(tabIdForKey('ArrowRight', TABS, 'roster')).toBe('loot');
    expect(tabIdForKey('End', TABS, 'overview')).toBe('loot');
    expect(tabIdForKey('ArrowLeft', TABS, 'unknown')).toBe('loot');
  });
});

describe('TabStrip (in-page tabs)', () => {
  it('is a labelled horizontal tablist of tabs', () => {
    const doc = strip({ onSelect: select });
    const list = doc.querySelector('[role=tablist]');
    expect(list?.getAttribute('aria-label')).toBe('Guild sections');
    expect(list?.getAttribute('aria-orientation')).toBe('horizontal');
    expect(list?.className).toBe(TAB_STRIP.strip);
    expect(doc.querySelectorAll('[role=tab]')).toHaveLength(4);
    expect(doc.querySelector('nav')).toBeNull();
  });

  it('selects the current tab and makes it the one tab stop (roving tabindex)', () => {
    const tabs = [...strip({ onSelect: select }).querySelectorAll('[role=tab]')];
    expect(tabs.map((tab) => tab.getAttribute('aria-selected'))).toEqual(['false', 'true', 'false', 'false']);
    expect(tabs.map((tab) => tab.getAttribute('tabindex'))).toEqual(['-1', '0', '-1', '-1']);
    expect(tabs[1].className).toBe(tabClass(true, false));
    expect(tabs[0].className).toBe(tabClass(false, false));
  });

  it('draws a disabled tab muted, aria-disabled and out of the tab order', () => {
    const raids = strip({ onSelect: select }).querySelector('[data-tab-id=raids]');
    expect(raids?.getAttribute('aria-disabled')).toBe('true');
    expect(raids?.getAttribute('tabindex')).toBe('-1');
    expect(raids?.className).toBe(tabClass(false, true));
  });

  it('keeps a hash tab a real link and a plain tab a button', () => {
    const doc = strip({ onSelect: select });
    expect(doc.querySelector('[data-tab-id=loot]')?.tagName).toBe('A');
    expect(doc.querySelector('[data-tab-id=loot]')?.getAttribute('href')).toBe('#loot');
    expect(doc.querySelector('[data-tab-id=overview]')?.tagName).toBe('BUTTON');
    expect(doc.querySelector('[data-tab-id=overview]')?.getAttribute('type')).toBe('button');
  });

  it('wires each tab to its panel through idPrefix', () => {
    const tab = strip({ onSelect: select, idPrefix: 'guild' }).querySelector('[data-tab-id=roster]');
    expect(tab?.id).toBe('guild-tab-roster');
    expect(tab?.getAttribute('aria-controls')).toBe('guild-panel-roster');
  });

  it('renders the suffix muted and the per-tab test id', () => {
    const doc = strip({ onSelect: select });
    expect(doc.querySelector('[data-tab-id=roster] span')?.className).toBe(TAB_STRIP.suffix);
    expect(doc.querySelector('[data-testid=tab-overview]')).not.toBeNull();
  });

  it('names the strip itself, and carries both edge fades, decorative, driven by the wrapper', () => {
    const doc = strip({ onSelect: select, testid: 'strip' });
    const list = doc.querySelector('[data-testid=strip]');
    expect(list?.getAttribute('role')).toBe('tablist');
    const wrapper = list?.parentElement;
    expect(wrapper?.className).toBe(TAB_STRIP.wrapper);
    expect(wrapper?.querySelector('[data-testid=strip-fade-start]')?.getAttribute('aria-hidden')).toBe(
      'true',
    );
    expect(wrapper?.querySelector('[data-testid=strip-fade-end]')?.className).toBe(TAB_STRIP.fadeEnd);
  });

  it('names the fades generically without a test id, and passes a tab its own attributes', () => {
    const doc = strip({ onSelect: select, tabs: [{ id: 'a', label: 'A', attrs: { 'data-sim-tab': 'a' } }] });
    expect(doc.querySelector('[data-testid=tab-strip-fade-end]')).not.toBeNull();
    expect(doc.querySelector('[data-sim-tab=a]')?.getAttribute('role')).toBe('tab');
  });
});

describe('TabStrip (link tabs)', () => {
  const links: readonly TabItem[] = [
    { id: 'quick', label: 'Quick sim', href: '/sim', attrs: { 'data-sim-tab': 'quick' } },
    { id: 'gear', label: 'Top gear', href: '/sim/gear' },
    { id: 'soon', label: 'Soon' },
  ];

  it('is a labelled nav of links with the current one aria-current, no tab roles', () => {
    const doc = strip({ tabs: links, active: 'gear', label: 'Simulator tools' });
    const nav = doc.querySelector('nav');
    expect(nav?.getAttribute('aria-label')).toBe('Simulator tools');
    expect(doc.querySelector('[role=tablist], [role=tab]')).toBeNull();
    expect(doc.querySelector('a[aria-current=page]')?.getAttribute('href')).toBe('/sim/gear');
    expect(doc.querySelectorAll('a[aria-current]')).toHaveLength(1);
    expect(doc.querySelector('[data-tab-id=soon]')?.tagName).toBe('SPAN');
  });

  it('matches the static twin', async () => {
    const container = await AstroContainer.create();
    const html = await container.renderToString(TabStripAstro, {
      props: { tabs: links, active: 'quick', label: 'Simulator tools', testid: 'sim-tabs' },
    });
    const doc = new JSDOM(html).window.document;
    expect(doc.querySelector('nav')?.getAttribute('aria-label')).toBe('Simulator tools');
    expect(doc.querySelector('nav')?.className).toBe(TAB_STRIP.strip);
    const current = doc.querySelector('a[aria-current=page]');
    expect(current?.getAttribute('href')).toBe('/sim');
    expect(current?.className).toBe(tabClass(true, false));
    expect(doc.querySelector('a[href="/sim/gear"]')?.className).toBe(tabClass(false, false));
    expect(doc.querySelector('[data-tab-id=soon]')?.getAttribute('aria-disabled')).toBe('true');
    expect(doc.querySelector('nav')?.getAttribute('data-testid')).toBe('sim-tabs');
    expect(doc.querySelector('[data-testid=sim-tabs-fade-end]')).not.toBeNull();
    expect(doc.querySelector('[data-sim-tab=quick]')?.getAttribute('href')).toBe('/sim');
  });
});
