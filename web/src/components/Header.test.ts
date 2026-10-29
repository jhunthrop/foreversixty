import { experimental_AstroContainer as AstroContainer } from 'astro/container';
import { JSDOM } from 'jsdom';
import { describe, expect, it } from 'vitest';
import classes from '../data/classes.json';
import communityLinks from '../data/links.json';
import { PREMIUM_NAV_ITEM } from '../lib/nav';
import Header from './Header.astro';

async function renderDoc(path: string): Promise<Document> {
  return new JSDOM(await renderHeader(path)).window.document;
}

async function renderHeader(path: string): Promise<string> {
  const container = await AstroContainer.create();
  return container.renderToString(Header, { props: { path } });
}

describe('Header', () => {
  it('ends the doors with the Premium tag after Get set up: a plain link with no panel (spec 2026-10-29-premium-entry 3)', async () => {
    const doc = await renderDoc('/');
    const nav = doc.getElementById('primary-nav')!;
    const last = nav.lastElementChild!;
    expect(last.previousElementSibling?.getAttribute('data-door')).toBe('setup');
    expect(last.getAttribute('data-door')).toBe('premium');
    expect(last.getAttribute('href')).toBe(PREMIUM_NAV_ITEM.href);
    expect(last.hasAttribute('aria-haspopup')).toBe(false);
    expect(last.hasAttribute('aria-current')).toBe(false);
    const tag = last.querySelector('.pill.pill-site');
    expect(tag?.textContent?.trim()).toBe(PREMIUM_NAV_ITEM.label);
    expect(last.querySelectorAll('svg, i')).toHaveLength(0);
    expect(PREMIUM_NAV_ITEM).toEqual({ label: 'Premium', href: '/premium' });
  });

  it('carries the pre-paint hint script inside the header, before the doors (so /b/{id} has it)', async () => {
    const doc = await renderDoc('/');
    const header = doc.querySelector('[data-site-header]')!;
    const script = [...header.querySelectorAll('script:not([src])')].find((s) =>
      s.textContent?.includes('fs.premiumSource'),
    );
    expect(script).toBeDefined();
    expect(script?.textContent).toContain('premiumNav');
    const nav = doc.getElementById('primary-nav')!;
    expect(script!.compareDocumentPosition(nav) & 4).toBe(4);
  });

  it('marks the Premium tag current on /premium and its checkout', async () => {
    for (const path of ['/premium', '/premium/checkout']) {
      const doc = await renderDoc(path);
      expect(doc.querySelector('[data-door="premium"]')?.getAttribute('aria-current')).toBe('page');
    }
  });

  it('renders the four group doors in order, each a real link with its popup attributes, then Get set up as a plain door', async () => {
    const doc = await renderDoc('/');
    const doors = [...doc.querySelectorAll('a.site-door')];
    expect(doors.map((door) => [door.textContent?.trim(), door.getAttribute('href')])).toEqual([
      ['Tools', '/planner'],
      ['Logs', '/logs'],
      ['Guild', '/guild'],
      ['Guides', '/guides'],
      ['Get set up', '/setup'],
      ['Premium', '/premium'],
    ]);
    const setup = doors.at(-2)!;
    expect(setup.getAttribute('data-door')).toBe('setup');
    expect(setup.hasAttribute('aria-haspopup')).toBe(false);
    expect(setup.classList.contains('site-door-setup')).toBe(true);
    for (const door of doors.slice(0, -2)) {
      const key = door.getAttribute('data-door');
      expect(door.getAttribute('aria-haspopup')).toBe('true');
      expect(door.getAttribute('aria-expanded')).toBe('false');
      expect(door.getAttribute('aria-controls')).toBe(`site-panel-${key}`);
    }
    expect(doc.querySelector('a.site-door[data-door="guild"]')?.getAttribute('data-guild-tab')).toBe(
      'overview',
    );
    const toggles = [...doc.querySelectorAll('button.site-door-toggle')];
    expect(toggles.map((t) => t.getAttribute('aria-label'))).toEqual([
      'Open the Tools list',
      'Open the Logs list',
      'Open the Guild list',
      'Open the Guides list',
    ]);
  });

  it('server-renders every panel hidden with its rows and hrefs', async () => {
    const doc = await renderDoc('/');
    const rowsOf = (key: string) =>
      [...doc.querySelectorAll(`#site-panel-${key} a.site-panel-row`)].map((row) => [
        row.querySelector('.site-panel-title')?.textContent?.trim(),
        row.getAttribute('href'),
      ]);
    for (const key of ['tools', 'logs', 'guild', 'guides']) {
      const panel = doc.getElementById(`site-panel-${key}`);
      expect(panel?.hasAttribute('hidden')).toBe(true);
      expect(panel?.getAttribute('role')).toBe('menu');
    }
    expect(rowsOf('tools')).toEqual([
      ['Planner', '/planner'],
      ['BiS', '/bis'],
      ['Simulator', '/sim'],
    ]);
    expect(rowsOf('logs')).toEqual([
      ['Upload a log', '/logs'],
      ['Your reports', '/reports'],
      ['Rankings', '/rankings'],
      ['Tier list', '/tiers'],
    ]);
    expect(rowsOf('guild')).toEqual([
      ['Your guild', '/guild'],
      ['Roster', '/guild#roster'],
      ['Loot council', '/guild#loot'],
      ['Guild logs', '/guild#raids'],
      ['Forever Sixty Discord', communityLinks.discordInvite],
    ]);
    expect(doc.querySelector('#site-panel-guild [data-guild-name]')?.textContent).toBe('Your guild');
    expect(doc.querySelector('#site-panel-guild .site-panel-foot')?.getAttribute('href')).toBe(
      '/guild#settings',
    );
    expect(doc.querySelector('[data-testid="site-panel-row-upload-a-log"]')).not.toBeNull();
  });

  it('carries what the Roster and Guild settings doors say without a guild, for the header script', async () => {
    const doc = await renderDoc('/');
    const roster = doc.querySelector('#site-panel-guild [data-guild-tab="roster"]')!;
    expect(roster.getAttribute('data-guildless-label')).toBe('Find a guild');
    expect(roster.getAttribute('data-guildless-href')).toBe('/guild');
    const foot = doc.querySelector('#site-panel-guild .site-panel-foot')!;
    expect(foot.getAttribute('data-guildless-label')).toBe('Claim your guild');
    expect(foot.getAttribute('data-guildless-href')).toBe('/guild#claim');
    expect(
      doc.querySelector('#site-panel-guild [data-guild-tab="loot"]')?.hasAttribute('data-guildless-label'),
    ).toBe(false);
  });

  it('renders the Guides grid from classes.json with crests and spec names', async () => {
    const doc = await renderDoc('/');
    const tiles = [...doc.querySelectorAll('#site-panel-guides a[href^="/guides/"]')];
    expect(tiles.map((tile) => tile.getAttribute('href'))).toEqual(classes.map((c) => `/guides/${c.slug}`));
    expect(tiles[0].querySelector('img')?.getAttribute('src')).toBe('/icons/hd/crests/warrior.webp');
    expect(tiles[0].textContent).toContain('Warrior');
    expect(tiles[0].textContent).toContain('Arms · Fury · Protection');
    expect(doc.querySelector('#site-panel-guides .site-panel-foot')?.textContent?.trim()).toBe('All guides');
  });

  it('marks the group whose row is the current page', async () => {
    const doc = await renderDoc('/tiers/healer');
    expect(doc.querySelector('a.site-door[data-door="logs"]')?.getAttribute('aria-current')).toBe('page');
    expect(doc.querySelector('a.site-door[data-door="tools"]')?.hasAttribute('aria-current')).toBe(false);
  });

  it('lists the four groups as headed sections in the phone Menu', async () => {
    const doc = await renderDoc('/');
    const headings = [...doc.querySelectorAll('#site-menu h2.site-menu-heading')];
    expect(headings.map((h) => h.textContent?.trim())).toEqual(['Tools', 'Logs', 'Guild', 'Guides']);
    expect(doc.querySelectorAll('#site-menu .site-menu-list a[href^="/guides/"]').length).toBe(9);
  });

  it('renders no Reference disclosure, and Premium only as the one tag (never a door)', async () => {
    const html = await renderHeader('/');
    const doc = await renderDoc('/');
    expect(html).not.toContain('<details');
    expect(doc.querySelectorAll('a[href^="/premium"]')).toHaveLength(1);
    expect(doc.querySelector('a[href="/premium"] > .pill-site')).not.toBeNull();
    expect(html).not.toContain('>Classes<');
    expect(html).not.toContain('>The addon<');
  });

  it('has no Discord link in the bar or the Menu; the Guild panel carries the one row', async () => {
    const doc = await renderDoc('/');
    expect(doc.querySelector('.site-discord, .site-menu-discord, [data-testid="discord-link"]')).toBeNull();
    const rows = [...doc.querySelectorAll('a[href*="discord.gg"]')];
    expect(rows).toHaveLength(2);
    for (const row of rows) {
      expect(row.closest('#site-panel-guild, [data-menu-list="guild"]')).not.toBeNull();
      expect(row.textContent).toContain('Forever Sixty Discord');
    }
  });

  it('renders a collapsed Menu button that controls the menu panel holding the doors', async () => {
    const html = await renderHeader('/');
    expect(html).toMatch(/<button[^>]*aria-expanded="false"[^>]*aria-controls="site-menu"/);
    expect(html).toContain('id="site-menu"');
    expect(html.indexOf('id="site-menu"')).toBeLessThan(html.indexOf('id="primary-nav"'));
  });

  it('never says Leveling BiS and keeps the selector slot in the bar', async () => {
    const html = await renderHeader('/');
    expect(html).not.toContain('Leveling BiS');
    expect(html).toContain('site-right');
    expect(html.indexOf('site-menu"')).toBeLessThan(html.indexOf('primary-nav'));
  });

  it('renders the phone nav as a fixed-height wrapping grid, not a horizontally scrolling row', async () => {
    const html = await renderHeader('/');
    expect(html).not.toContain('overflow-x-auto');
    expect(html).not.toContain('nav-scroll-fade');
  });

  it('carries a compact character lookup on every page: a bar button keyed to "/" and a Menu row (Forever Score fix round 2 item 7)', async () => {
    const html = await renderHeader('/ratings');
    expect(html).toMatch(/<button[^>]*class="site-search-button"[^>]*aria-keyshortcuts="\/"/);
    expect(html).toMatch(/class="site-search-panel" hidden/);
    expect(html).toContain('data-testid="site-search-input"');
    expect(html).toContain('data-testid="site-menu-search-input"');
    expect(html.match(/placeholder="Character name"/g)).toHaveLength(2);
    // The Menu's lookup comes before the doors.
    expect(html.indexOf('site-menu-search')).toBeLessThan(html.indexOf('id="primary-nav"'));
  });

  it('keeps each lookup status live region present while empty, closing its gap without hiding it (fix round 4 item 2)', async () => {
    const html = await renderHeader('/ratings');
    const regions = [
      ...html.matchAll(/<p id="site-search-(bar|menu)-status" class="([^"]*)" role="status"><\/p>/g),
    ];
    expect(regions).toHaveLength(2);
    for (const [, , cls] of regions) {
      expect(cls.split(' ')).toContain('empty:mt-0');
      expect(cls).not.toMatch(/hidden|sr-only|display/);
    }
  });
});
