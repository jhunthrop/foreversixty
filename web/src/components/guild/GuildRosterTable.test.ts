// web/src/components/guild/GuildRosterTable.test.ts
// Task L8.3: the Roster tab's "In game" column and the snapshot line above the table,
// rendered with svelte/server (initial state: the default item-level sort, no effects). The
// e2e spec guild-centre.spec.ts clicks the sort and reads the hovers in a browser.
import { render } from 'svelte/server';
import { describe, expect, it } from 'vitest';
import { ADDON_SNAPSHOT, VIEWER_MEMBER, buildMockHome } from '../../fixtures/guild/mock-guild';
import type { GuildHome } from '../../lib/guild/api';
import { guildRosterAddonCopy } from '../../lib/guild/copy';
import GuildRosterTable from './GuildRosterTable.svelte';

const NOW = new Date('2026-12-22T21:46:00Z');

function props(home: GuildHome) {
  return {
    roster: home.roster,
    pending: home.pending ?? [],
    officer: false,
    myCharacterKey: home.viewer?.character_key ?? null,
    rosterBusy: null,
    rosterActionError: '',
    frozen: false,
    onApprove: () => {},
    onRemove: () => {},
    onApproveAll: () => {},
    rowPath: () => ({ region: 'us' as const, ruleset: 'pvp' as const, slug: 'x' }),
    addonSnapshot: home.addon_snapshot,
    now: NOW,
  };
}

function text(html: string): string {
  return html
    .replace(/<!--[\s\S]*?-->/g, '')
    .replace(/<[^>]+>/g, ' ')
    .replace(/&#39;/g, "'")
    .replace(/&amp;/g, '&')
    .replace(/\s+/g, ' ')
    .trim();
}

/** Each verified row's markup, in rendered order, keyed by the row's name. */
function rows(html: string): Map<string, string> {
  const body = html.slice(html.indexOf('data-testid="guild-roster-rows"'));
  const parts = body.split('guild-roster-row ').slice(1);
  const byName = new Map<string, string>();
  for (const part of parts) {
    const name = /data-testid="guild-roster-name"[^>]*>([^<]+)</.exec(part)?.[1].trim() ?? '';
    byName.set(name.replace(' (you)', ''), part);
  }
  return byName;
}

const withSnapshot = render(GuildRosterTable, { props: props(buildMockHome(VIEWER_MEMBER)) }).body;
const withoutSnapshot = render(GuildRosterTable, {
  props: props(buildMockHome(VIEWER_MEMBER, undefined, null, false)),
}).body;

describe('GuildRosterTable: the In game column (L8.3)', () => {
  it('adds an "In game" sort header after the existing columns, hovering "as the addon reports it"', () => {
    const header = withSnapshot.slice(withSnapshot.indexOf('data-testid="guild-roster-sort-header"'));
    const labels = [...header.matchAll(/data-testid="guild-roster-sort-(?!header")([a-z_]+)"/g)].map(
      (m) => m[1],
    );
    expect(labels).toEqual(['rank', 'ilvl', 'rating', 'attendance', 'seen', 'addon_gap']);
    const tag = /<button[^>]*data-testid="guild-roster-sort-addon_gap"[^>]*>/.exec(header)?.[0] ?? '';
    expect(tag).toContain(`title="${guildRosterAddonCopy.columnHover}"`);
    expect(guildRosterAddonCopy.columnHover).toContain('as the addon reports it');
    expect(text(header)).toContain(guildRosterAddonCopy.column);
  });

  it('keeps the default sort on item level', () => {
    const header = withSnapshot.slice(withSnapshot.indexOf('data-testid="guild-roster-sort-header"'));
    const active = /data-testid="guild-roster-sort-ilvl"[^>]*>([^<]*)</.exec(header)?.[1] ?? '';
    expect(active).toContain('▾');
    const names = [...rows(withSnapshot).keys()];
    // Zulmara (the viewer) pinned first, then item level descending: Windtalon 70 leads.
    expect(names.slice(0, 3)).toEqual(['Zulmara', 'Windtalon', 'Pyrewisp']);
  });

  it('a matched row shows the level (and the spec where it differs) tagged "(addon)", the gap in DPS, the build and the seen time', () => {
    const byName = rows(withSnapshot);
    // Zulmara: the site says Marksmanship, the addon says Survival.
    const zulmara = byName.get('Zulmara') ?? '';
    expect(zulmara).toContain('data-testid="guild-roster-ingame"');
    expect(text(zulmara)).toContain('Level 60 Survival (addon) · 12 DPS · Survival raid · seen 21:40');
    expect(zulmara).toMatch(/class="[^"]*text-muted[^"]*"[^>]*>\(addon\)</);
    // Kraggor: the addon agrees with the site's Protection, so only the level shows.
    expect(text(byName.get('Kraggor') ?? '')).toContain(
      'Level 60 (addon) · 4 DPS · Protection raid · seen 21:38',
    );
  });

  it('a consent-withheld row (level, spec, gap, build all null) shows only "seen HH:MM"', () => {
    const dusk = rows(withSnapshot).get('Duskstrider') ?? '';
    expect(dusk).toContain('data-testid="guild-roster-ingame"');
    expect(dusk).not.toContain('guild-roster-ingame-gap');
    expect(dusk).not.toContain('(addon)');
    const cell = dusk.slice(
      dusk.indexOf('data-testid="guild-roster-ingame"'),
      dusk.indexOf('guild-roster-handoff'),
    );
    expect(cell).not.toContain('DPS');
    expect(text(dusk)).toContain('seen 23:12');
    expect(text(dusk)).not.toContain('· seen');
  });

  it('a row the snapshot did not carry shows "—", hovering "Not seen by the addon in the last snapshot"', () => {
    const windtalon = rows(withSnapshot).get('Windtalon') ?? '';
    const cell = /<span[^>]*data-testid="guild-roster-ingame-none"[^>]*>[^<]*</.exec(windtalon)?.[0] ?? '';
    expect(cell).toContain(`title="${guildRosterAddonCopy.notSeenHover}"`);
    expect(guildRosterAddonCopy.notSeenHover).toBe('Not seen by the addon in the last snapshot');
    expect(cell).toContain('>—<');
    expect(windtalon).not.toContain('data-testid="guild-roster-ingame"');
  });

  it('names the column for a screen reader on every row', () => {
    for (const part of rows(withSnapshot).values()) {
      expect(part).toContain(`<span class="sr-only">${guildRosterAddonCopy.srLabel}</span>`);
    }
  });

  it('puts the snapshot line above the table: when, by whom, how many', () => {
    const line = /data-testid="guild-roster-addon-snapshot"[^>]*>([^<]*)</.exec(withSnapshot)?.[1].trim();
    expect(line).toBe('Roster from the addon · received 3 minutes ago by Zulmara · 4 members');
    expect(withSnapshot.indexOf('guild-roster-addon-snapshot')).toBeLessThan(
      withSnapshot.indexOf('data-testid="guild-roster-rows"'),
    );
    expect(ADDON_SNAPSHOT.by_character).toBe('us/pvp/zulmara');
  });

  it('with no snapshot: the null line, and every row "—"', () => {
    const line = /data-testid="guild-roster-addon-snapshot"[^>]*>([^<]*)</.exec(withoutSnapshot)?.[1].trim();
    expect(line).toBe(guildRosterAddonCopy.noSnapshot);
    expect(withoutSnapshot).not.toContain('data-testid="guild-roster-ingame"');
    const byName = rows(withoutSnapshot);
    for (const part of byName.values()) expect(part).toContain('data-testid="guild-roster-ingame-none"');
  });

  it('a home from before the addendum (no fields at all) reads as no snapshot', () => {
    const home = buildMockHome(VIEWER_MEMBER, undefined, null, false);
    const bare: GuildHome = {
      ...home,
      roster: home.roster.map(({ addon: _addon, ...row }) => row),
    };
    delete bare.addon_snapshot;
    const html = render(GuildRosterTable, { props: { ...props(bare), addonSnapshot: undefined } }).body;
    expect(text(html)).toContain(guildRosterAddonCopy.noSnapshot);
    expect(html).toContain('data-testid="guild-roster-ingame-none"');
  });
});
