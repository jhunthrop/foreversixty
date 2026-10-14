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

// A local-time instant, so "today in the viewer's date" is 2026-12-22 in any time zone; the
// snapshot's own received_at is derived from it (three minutes earlier).
const NOW = new Date(2026, 11, 22, 21, 46);
const RECEIVED = new Date(NOW.getTime() - 3 * 60_000).toISOString();

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
    addonSnapshot: home.addon_snapshot ? { ...home.addon_snapshot, received_at: RECEIVED } : null,
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

  it('a matched row shows the level, the spec tagged "per addon" only where it differs, "gap N DPS", the build and the seen time', () => {
    const byName = rows(withSnapshot);
    // Zulmara: the site says Marksmanship, the addon says Survival.
    const zulmara = byName.get('Zulmara') ?? '';
    expect(zulmara).toContain('data-testid="guild-roster-ingame"');
    expect(text(zulmara)).toContain('Level 60 Survival per addon · gap 12 DPS · build: Survival raid · seen 21:40');
    expect(zulmara).toMatch(/class="[^"]*text-muted[^"]*"[^>]*>per addon</);
    // Kraggor: the addon agrees with the site's Protection, so the level shows with no tag.
    const kraggor = text(byName.get('Kraggor') ?? '');
    expect(kraggor).toContain('Level 60 · gap 4 DPS · build: Protection raid · seen 21:38');
    expect(kraggor).not.toContain('per addon');
    expect(kraggor).not.toContain('(addon)');
    expect(withSnapshot).not.toContain('(addon)');
  });

  it("the seen time hovers \"Seen at HH:MM on the member's clock\"", () => {
    const zulmara = rows(withSnapshot).get('Zulmara') ?? '';
    expect(zulmara).toContain(`title="${guildRosterAddonCopy.seenHover('21:40')}"`);
    expect(guildRosterAddonCopy.seenHover('21:40')).toBe("Seen at 21:40 on the member's clock");
  });

  it('a seen time that is not today names its day', () => {
    expect(text(rows(withSnapshot).get('Duskstrider') ?? '')).toContain('seen Mon 23:12');
  });

  it('a consent-withheld row (level, spec, gap, build all null) shows only the seen time, hovering "Details withheld by the member"', () => {
    const dusk = rows(withSnapshot).get('Duskstrider') ?? '';
    expect(dusk).toContain('data-testid="guild-roster-ingame"');
    expect(dusk).not.toContain('guild-roster-ingame-gap');
    expect(dusk).not.toContain('per addon');
    const cell = dusk.slice(
      dusk.indexOf('data-testid="guild-roster-ingame"'),
      dusk.indexOf('guild-roster-handoff'),
    );
    expect(cell).not.toContain('DPS');
    expect(cell).toContain('title="Details withheld by the member"');
    expect(text(dusk)).not.toContain('· seen');
  });

  it('a row the snapshot did not carry shows "—", hovering "Not seen by the addon in the last snapshot", on desktop only', () => {
    const windtalon = rows(withSnapshot).get('Windtalon') ?? '';
    const cell = /<span[^>]*data-testid="guild-roster-ingame-none"[^>]*>[^<]*</.exec(windtalon)?.[0] ?? '';
    expect(cell).toContain(`title="${guildRosterAddonCopy.notSeenHover}"`);
    expect(guildRosterAddonCopy.notSeenHover).toBe('Not seen by the addon in the last snapshot');
    expect(cell).toContain('>—<');
    expect(windtalon).not.toContain('data-testid="guild-roster-ingame"');
    // A phone shows no lone "—": the whole slot is hidden below md.
    const slot = /<div[^>]*data-testid="guild-roster-ingame-slot"[^>]*>/.exec(windtalon)?.[0] ?? '';
    expect(slot).toMatch(/class="[^"]*\bhidden\b[^"]*lg:block/);
    // A matched row's slot is shown on phones.
    const matched = /<div[^>]*data-testid="guild-roster-ingame-slot"[^>]*>/.exec(rows(withSnapshot).get('Zulmara') ?? '')?.[0] ?? '';
    expect(matched).not.toMatch(/\bhidden\b/);
  });

  it('gives the In game text a fixed slot at the end of the row, before an actions group that never wraps', () => {
    for (const [name, part] of rows(withSnapshot)) {
      const slotAt = part.indexOf('data-testid="guild-roster-ingame-slot"');
      const actionsAt = part.indexOf('data-testid="guild-roster-actions"');
      expect(slotAt, name).toBeGreaterThan(-1);
      expect(actionsAt, `${name} actions`).toBeGreaterThan(slotAt);
      const slot = /<div[^>]*data-testid="guild-roster-ingame-slot"[^>]*>/.exec(part)?.[0] ?? '';
      expect(slot).toContain('lg:basis-[260px]');
      expect(slot).toContain('lg:shrink-0');
      expect(slot).toContain('min-w-0');
      const actions = /<div[^>]*data-testid="guild-roster-actions"[^>]*>/.exec(part)?.[0] ?? '';
      expect(actions).toContain('flex-nowrap');
      expect(actions).toContain('shrink-0');
      expect(actions).toContain('lg:w-[190px]');
      expect(actions).not.toMatch(/\bxl:/);
      expect(slot).not.toMatch(/\bxl:/);
      // Nothing of the row sits between the slot and the actions but the slot's own content.
      expect(part.slice(actionsAt)).not.toContain('data-testid="guild-roster-ingame-slot"');
    }
  });

  it('labels the build name in the In game slot from the copy file', () => {
    expect(guildRosterAddonCopy.build('Survival raid')).toBe('build: Survival raid');
    const zulmara = rows(withSnapshot).get('Zulmara') ?? '';
    expect(text(zulmara)).toContain('build: Survival raid');
    expect(text(zulmara)).not.toMatch(/gap 12 DPS · Survival raid/);
  });

  it('the In game header sits over the same slot', () => {
    const header = /<button[^>]*data-testid="guild-roster-sort-addon_gap"[^>]*>/.exec(withSnapshot)?.[0] ?? '';
    expect(header).toContain('lg:w-[260px]');
  });

  it('puts a dot between the spec and "alt of {main}" rather than running them together', () => {
    const alt = [...rows(withSnapshot).values()].find((part) => part.includes('alt of '));
    expect(alt, 'the fixture has an alt row').toBeDefined();
    expect(text(alt ?? '')).toMatch(/[A-Za-z] · alt of [A-Z]/);
    expect(text(alt ?? '')).not.toMatch(/[a-z]alt of/);
    const spec = /guild-roster-spec[^>]*>/.exec(alt ?? '');
    expect(spec).not.toBeNull();
  });

  it('names the column for a screen reader on every row', () => {
    for (const part of rows(withSnapshot).values()) {
      expect(part).toContain(`<span class="sr-only">${guildRosterAddonCopy.srLabel}</span>`);
    }
  });

  it('puts the snapshot line above the table: who sent it, when, and how many members were seen', () => {
    const home = buildMockHome(VIEWER_MEMBER);
    const seen = home.roster.filter((r) => r.addon !== null).length;
    const line = /data-testid="guild-roster-addon-snapshot"[^>]*>([^<]*)</.exec(withSnapshot)?.[1].trim();
    expect(line).toBe(`Roster from the addon · sent by Zulmara 3 minutes ago · ${seen} of ${home.roster.length} members seen`);
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
