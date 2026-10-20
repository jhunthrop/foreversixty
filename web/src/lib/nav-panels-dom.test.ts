// @vitest-environment jsdom
// web/src/lib/nav-panels-dom.test.ts
// The header's hover panels (spec 2026-10-15-grouped-header.md section 4) against the markup
// contract Header.astro renders: hover delays, click toggle, touch first tap, keyboard.
import { afterEach, beforeEach, describe, expect, it, vi } from 'vitest';
import { applySessionGuild, wireNavPanels, type NavPanels } from './nav-panels-dom';

const OPEN_DELAY = 120;
const CLOSE_DELAY = 180;

const group = (key: string, rows: string): string => `
  <div class="site-group" data-group="${key}">
    <a class="site-door" data-door="${key}" href="/${key}" aria-haspopup="true" aria-expanded="false" aria-controls="site-panel-${key}">${key}</a>
    <button type="button" class="site-door-toggle" aria-expanded="false" aria-controls="site-panel-${key}" data-door-toggle="${key}">v</button>
    <div id="site-panel-${key}" class="site-panel" role="menu" hidden>${rows}</div>
  </div>`;

const row = (href: string, extra = ''): string =>
  `<a class="site-panel-row" role="menuitem" href="${href}" ${extra}><span data-guild-name>Your guild</span></a>`;

function mount(): HTMLElement {
  document.body.innerHTML = `<header data-site-header>
    ${group('tools', row('/planner') + row('/bis'))}
    ${group('logs', row('/logs') + row('/reports') + row('/rankings'))}
    ${group(
      'guild',
      row('/guild', 'data-guild-tab="overview"') +
        '<a class="site-panel-row" role="menuitem" href="/guild#roster" data-guild-tab="roster" data-guildless-label="Find a guild" data-guildless-href="/guild" data-guildless-description="Look one up"><span class="site-panel-title">Roster</span><span class="site-panel-desc">Everyone the addon has seen</span></a>' +
        '<a class="site-panel-foot" role="menuitem" href="/guild#settings" data-guild-tab="settings" data-guildless-label="Claim your guild" data-guildless-href="/guild#claim">Guild settings</a>',
    )}
  </header><p id="outside">outside</p>`;
  return document.querySelector('header') as HTMLElement;
}

const door = (key: string): HTMLAnchorElement => document.querySelector(`[data-door="${key}"]`)!;
const toggle = (key: string): HTMLButtonElement => document.querySelector(`[data-door-toggle="${key}"]`)!;
const panel = (key: string): HTMLElement => document.getElementById(`site-panel-${key}`)!;
const groupEl = (key: string): HTMLElement => document.querySelector(`[data-group="${key}"]`)!;
const isOpen = (key: string): boolean =>
  !panel(key).hidden && door(key).getAttribute('aria-expanded') === 'true';
const press = (target: Element, name: string): KeyboardEvent => {
  const event = new KeyboardEvent('keydown', { key: name, bubbles: true, cancelable: true });
  target.dispatchEvent(event);
  return event;
};
const click = (target: Element): MouseEvent => {
  const event = new MouseEvent('click', { bubbles: true, cancelable: true });
  target.dispatchEvent(event);
  return event;
};
const enter = (el: Element): void => void el.dispatchEvent(new MouseEvent('pointerenter'));
const leave = (el: Element): void => void el.dispatchEvent(new MouseEvent('pointerleave'));

let panels: NavPanels;
let touch = false;

beforeEach(() => {
  vi.useFakeTimers();
  touch = false;
  panels = wireNavPanels(mount(), { isTouch: () => touch });
});

afterEach(() => {
  panels.destroy();
  vi.useRealTimers();
});

describe('hover', () => {
  it('opens 120ms after the pointer enters a door and closes 180ms after it leaves', () => {
    enter(groupEl('tools'));
    vi.advanceTimersByTime(OPEN_DELAY - 1);
    expect(isOpen('tools')).toBe(false);
    vi.advanceTimersByTime(1);
    expect(isOpen('tools')).toBe(true);
    leave(groupEl('tools'));
    vi.advanceTimersByTime(CLOSE_DELAY - 1);
    expect(isOpen('tools')).toBe(true);
    vi.advanceTimersByTime(1);
    expect(isOpen('tools')).toBe(false);
  });

  it('never opens when the pointer only crosses a door', () => {
    enter(groupEl('tools'));
    vi.advanceTimersByTime(50);
    leave(groupEl('tools'));
    vi.advanceTimersByTime(1000);
    expect(isOpen('tools')).toBe(false);
  });

  it('keeps the panel open when the pointer returns before it closes', () => {
    panels.open('tools');
    leave(groupEl('tools'));
    vi.advanceTimersByTime(100);
    enter(groupEl('tools'));
    vi.advanceTimersByTime(1000);
    expect(isOpen('tools')).toBe(true);
  });

  it('closes the first panel when a second door opens', () => {
    panels.open('tools');
    leave(groupEl('tools'));
    enter(groupEl('logs'));
    vi.advanceTimersByTime(OPEN_DELAY);
    expect(isOpen('logs')).toBe(true);
    expect(isOpen('tools')).toBe(false);
  });

  it('ignores a touch pointer, which opens through the tap instead', () => {
    const event = new MouseEvent('pointerenter');
    Object.defineProperty(event, 'pointerType', { value: 'touch' });
    groupEl('tools').dispatchEvent(event);
    vi.advanceTimersByTime(1000);
    expect(isOpen('tools')).toBe(false);
  });
});

describe('click and touch', () => {
  it('toggles on the chevron without navigating', () => {
    expect(click(toggle('tools')).defaultPrevented).toBe(true);
    expect(isOpen('tools')).toBe(true);
    click(toggle('tools'));
    expect(isOpen('tools')).toBe(false);
  });

  it('lets a click on the door label navigate', () => {
    expect(click(door('tools')).defaultPrevented).toBe(false);
    expect(isOpen('tools')).toBe(false);
  });

  it('opens on a first tap with no hover and navigates on the second', () => {
    touch = true;
    expect(click(door('guild')).defaultPrevented).toBe(true);
    expect(isOpen('guild')).toBe(true);
    expect(click(door('guild')).defaultPrevented).toBe(false);
  });

  it('closes on a click outside the groups', () => {
    panels.open('logs');
    click(document.getElementById('outside')!);
    expect(isOpen('logs')).toBe(false);
  });
});

describe('keyboard', () => {
  it('opens on ArrowDown from the door and focuses the first row', () => {
    door('logs').focus();
    expect(press(door('logs'), 'ArrowDown').defaultPrevented).toBe(true);
    expect(isOpen('logs')).toBe(true);
    expect(document.activeElement).toBe(panel('logs').querySelector('a'));
  });

  it('opens on Enter from a closed door, and lets Enter navigate once it is open', () => {
    door('logs').focus();
    expect(press(door('logs'), 'Enter').defaultPrevented).toBe(true);
    expect(isOpen('logs')).toBe(true);
    door('logs').focus();
    expect(press(door('logs'), 'Enter').defaultPrevented).toBe(false);
  });

  it('walks the rows with ArrowDown and ArrowUp', () => {
    door('logs').focus();
    press(door('logs'), 'ArrowDown');
    const rows = [...panel('logs').querySelectorAll('a')];
    press(document.activeElement!, 'ArrowDown');
    press(document.activeElement!, 'ArrowDown');
    expect(document.activeElement).toBe(rows[2]);
    press(document.activeElement!, 'ArrowUp');
    expect(document.activeElement).toBe(rows[1]);
  });

  it('closes on Escape and returns focus to the door', () => {
    door('logs').focus();
    press(door('logs'), 'ArrowDown');
    press(document.activeElement!, 'Escape');
    expect(isOpen('logs')).toBe(false);
    expect(document.activeElement).toBe(door('logs'));
  });

  it('closes when focus leaves the group', () => {
    door('logs').focus();
    press(door('logs'), 'ArrowDown');
    const rows = [...panel('logs').querySelectorAll('a')];
    rows.at(-1)!.dispatchEvent(new FocusEvent('focusout', { bubbles: true, relatedTarget: door('guild') }));
    expect(isOpen('logs')).toBe(false);
  });
});

describe('applySessionGuild', () => {
  const hrefs = (): (string | null)[] =>
    [...document.querySelectorAll('[data-guild-tab]')].map((el) => el.getAttribute('href'));
  const nameText = (): string | null | undefined => document.querySelector('[data-guild-name]')?.textContent;
  const rosterText = (): string | null | undefined =>
    document.querySelector('[data-guild-tab="roster"] .site-panel-title')?.textContent;
  const footText = (): string | null | undefined =>
    document.querySelector('.site-panel-foot')?.textContent?.trim();
  const olympus = { name: 'OLYMPUS XXVII', href: '/guild/us/pvp/olympus-xxvii' };

  it('rewrites the guild rows to the session guild and names it', () => {
    applySessionGuild(document.querySelector('header')!, olympus);
    expect(hrefs()).toEqual([
      '/guild/us/pvp/olympus-xxvii',
      '/guild/us/pvp/olympus-xxvii#roster',
      '/guild/us/pvp/olympus-xxvii/settings',
    ]);
    expect(nameText()).toBe('OLYMPUS XXVII');
    expect(rosterText()).toBe('Roster');
    expect(footText()).toBe('Guild settings');
  });

  it('offers Find a guild and Claim your guild, not dead links, with no guild', () => {
    applySessionGuild(document.querySelector('header')!, null);
    expect(hrefs()).toEqual(['/guild', '/guild', '/guild#claim']);
    expect(rosterText()).toBe('Find a guild');
    expect(document.querySelector('[data-guild-tab="roster"] .site-panel-desc')?.textContent).toBe(
      'Look one up',
    );
    expect(footText()).toBe('Claim your guild');
  });

  it('puts the member rows back when a guild arrives after the guildless pass', () => {
    const header = document.querySelector('header')!;
    applySessionGuild(header, null);
    applySessionGuild(header, olympus);
    expect(hrefs()).toEqual([
      '/guild/us/pvp/olympus-xxvii',
      '/guild/us/pvp/olympus-xxvii#roster',
      '/guild/us/pvp/olympus-xxvii/settings',
    ]);
    expect(rosterText()).toBe('Roster');
    expect(document.querySelector('[data-guild-tab="roster"] .site-panel-desc')?.textContent).toBe(
      'Everyone the addon has seen',
    );
    expect(footText()).toBe('Guild settings');
  });

  it('restores the default rows and name for no guild', () => {
    const header = document.querySelector('header')!;
    applySessionGuild(header, olympus);
    applySessionGuild(header, null);
    expect(hrefs()).toEqual(['/guild', '/guild', '/guild#claim']);
    expect(nameText()).toBe('Your guild');
  });
});
