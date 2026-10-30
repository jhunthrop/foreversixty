// web/src/components/guild/GuildRosterTable.test.ts
// Task L8.3: the Roster tab's "In game" column and the snapshot line above the table,
// rendered with svelte/server (initial state: the default item-level sort, no effects). The
// e2e spec guild-centre.spec.ts clicks the sort and reads the hovers in a browser.
import { render } from 'svelte/server';
import { describe, expect, it } from 'vitest';
import {
  ADDON_SNAPSHOT,
  VIEWER_MEMBER,
  VIEWER_OFFICER,
  buildMockHome,
} from '../../fixtures/guild/mock-guild';
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
    members: home.members,
    counts: home.counts,
    can: [] as string[],
    myCharacterKey: home.viewer?.character_key ?? null,
    rosterBusy: null,
    rosterActionError: '',
    frozen: false,
    onApprove: () => {},
    onRemove: () => {},
    onDecline: () => {},
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
  it('sorts by chips under a SORT label: no header row, no "In game" heading', () => {
    expect(withSnapshot).not.toContain('guild-roster-sort-header');
    expect(withSnapshot).not.toContain('columnheader');
    const sort = withSnapshot.slice(withSnapshot.indexOf('data-testid="guild-roster-sort"'));
    const row = sort.slice(0, sort.indexOf('data-testid="guild-roster-addon-snapshot"'));
    expect(text(row)).toContain('Sort Rank Item level');
    const labels = [...row.matchAll(/data-testid="guild-roster-sort-([a-z_]+)"/g)].map((m) => m[1]);
    expect(labels).toEqual(['rank', 'ilvl', 'rating', 'attendance', 'seen']);
    expect(labels).not.toContain('addon_gap');
    expect(text(row)).not.toContain(guildRosterAddonCopy.column);
  });

  it('hides the Rating chip when no row has a rating', () => {
    const home = buildMockHome(VIEWER_MEMBER);
    const unrated = {
      ...props(home),
      roster: home.roster.map(({ rating: _rating, ...rest }) => rest),
    };
    const html = render(GuildRosterTable, { props: unrated }).body;
    expect(html).not.toContain('guild-roster-sort-rating');
    expect(html).toContain('guild-roster-sort-attendance');
  });

  it('keeps the default sort on item level, the active chip showing its arrow and aria-sort', () => {
    const tag = /<button[^>]*data-testid="guild-roster-sort-ilvl"[^>]*>[^<]*</.exec(withSnapshot)?.[0] ?? '';
    expect(tag).toContain('▾');
    expect(tag).toContain('aria-sort="descending"');
    expect(/<button[^>]*data-testid="guild-roster-sort-rank"[^>]*>/.exec(withSnapshot)?.[0]).not.toContain(
      'aria-sort',
    );
    const names = [...rows(withSnapshot).keys()];
    // Zulmara (the viewer) pinned first, then item level descending: Windtalon 70 leads.
    expect(names.slice(0, 3)).toEqual(['Zulmara', 'Windtalon', 'Pyrewisp']);
  });

  it('a matched row reads Level, Gear gap, the gear set and seen in that order; the spec has its own line', () => {
    const byName = rows(withSnapshot);
    // Zulmara: the site says Marksmanship, the addon says Survival.
    const zulmara = byName.get('Zulmara') ?? '';
    expect(zulmara).toContain('data-testid="guild-roster-ingame"');
    expect(text(zulmara)).toContain(
      'Level 60 · Gear gap 12 DPS · Gear set: Survival (raid) · seen 6 minutes ago',
    );
    expect(text(zulmara)).toContain('Survival Hunter in game · Marksmanship on the site');
    // Kraggor: the addon agrees with the site's Protection, so the line says only "in game".
    const kraggor = text(byName.get('Kraggor') ?? '');
    expect(kraggor).toContain('Level 60 · Gear gap 4 DPS · Gear set: Protection (raid) · seen 8 minutes ago');
    expect(kraggor).toContain('Protection Warrior in game');
    expect(kraggor).not.toContain('on the site');
    expect(kraggor).not.toContain('(addon)');
    expect(withSnapshot).not.toContain('(addon)');
  });

  it("the viewer's own row compares the snapshot with the site's character record", () => {
    const html = render(GuildRosterTable, {
      props: { ...props(buildMockHome(VIEWER_MEMBER)), mySiteSpec: 'Beast Mastery' },
    }).body;
    expect(text(rows(html).get('Zulmara') ?? '')).toContain(
      'Survival Hunter in game · Beast Mastery on the site',
    );
    // Other rows still read the roster's own spec.
    expect(text(rows(html).get('Kraggor') ?? '')).toContain('Protection Warrior in game');
  });

  it('the seen time hovers "Seen at HH:MM on the member\'s clock"', () => {
    const zulmara = rows(withSnapshot).get('Zulmara') ?? '';
    expect(zulmara).toContain(`title="${guildRosterAddonCopy.seenHover('21:40')}"`);
    expect(guildRosterAddonCopy.seenHover('21:40')).toBe("Seen at 21:40 on the member's clock");
  });

  it('an older seen time reads as how long ago it was', () => {
    expect(text(rows(withSnapshot).get('Duskstrider') ?? '')).toContain('seen 22 hours ago');
  });

  it('a consent-withheld row (level, spec, gap, build all null) shows only the seen time, hovering "Details withheld by the member"', () => {
    const dusk = rows(withSnapshot).get('Duskstrider') ?? '';
    expect(dusk).toContain('data-testid="guild-roster-ingame"');
    expect(dusk).not.toContain('guild-roster-ingame-gap');
    expect(dusk).not.toContain(' in game');
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
    const matched =
      /<div[^>]*data-testid="guild-roster-ingame-slot"[^>]*>/.exec(
        rows(withSnapshot).get('Zulmara') ?? '',
      )?.[0] ?? '';
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

  it('words the gear set as "Gear set: Survival (raid)" and never "build:" or "per addon"', () => {
    const zulmara = text(rows(withSnapshot).get('Zulmara') ?? '');
    expect(zulmara).toContain('Gear set: Survival (raid)');
    expect(zulmara).not.toContain('build:');
    expect(zulmara).not.toContain('per addon');
  });

  it('names the unit on every stat chip and keeps the same fact order on every row', () => {
    for (const [name, html] of rows(withSnapshot)) {
      const line = text(html);
      expect(line, name).not.toMatch(/\d rating\b/);
      expect(line, name).not.toMatch(/\bilvl\b/);
      const level = line.indexOf('Level 60');
      const gap = line.indexOf('Gear gap');
      const set = line.indexOf('Gear set:');
      const seen = line.indexOf('seen ');
      const present = [level, gap, set, seen].filter((at) => at >= 0);
      expect(present, name).toEqual([...present].sort((a, b) => a - b));
    }
    const kraggor = text(rows(withSnapshot).get('Kraggor') ?? '');
    expect(kraggor).toMatch(/Rating \d+/);
    expect(kraggor).toMatch(/Item level \d+/);
    expect(kraggor).toContain('Protection Warrior');
  });

  it('gives every row the same fixed name block and rank slot at lg, so the stats line up', () => {
    for (const [name, part] of rows(withSnapshot)) {
      expect(part, name).toContain('lg:w-[240px]');
      expect(part, name).toContain('lg:w-[72px]');
      const at = part.indexOf('data-testid="guild-roster-spec"');
      const subtitle = part.slice(part.lastIndexOf('<span', at), at + 400);
      expect(subtitle, name).toContain('lg:truncate');
      expect(subtitle, name).toContain('title="');
    }
  });

  it('drops the Gear gap fact when it is 0', () => {
    const home = buildMockHome(VIEWER_MEMBER);
    const zeroed = {
      ...props(home),
      roster: home.roster.map((row) =>
        row.addon === undefined || row.addon === null
          ? row
          : { ...row, addon: { ...row.addon, gear_gap: 0 } },
      ),
    };
    const html = render(GuildRosterTable, { props: zeroed }).body;
    expect(text(html)).not.toContain('Gear gap 0');
    expect(html).not.toContain('guild-roster-ingame-gap');
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

  it('puts the snapshot line above the table: who sent it and when (the counts have their own line)', () => {
    const line = /data-testid="guild-roster-addon-snapshot"[^>]*>([^<]*)</.exec(withSnapshot)?.[1].trim();
    expect(line).toBe('Roster from the addon · sent by Zulmara 3 minutes ago');
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

describe('GuildRosterTable: everyone the addon saw, with a word beside each', () => {
  const home = buildMockHome(VIEWER_OFFICER);
  const officerAsk = (extra: Record<string, unknown> = {}) =>
    render(GuildRosterTable, {
      props: {
        ...props(home),
        can: ['approve', 'remove', 'settings', 'award'],
        ...extra,
      },
    }).body;
  const asOfficer = officerAsk();

  it('puts the count line above the filters: on the roster, in game without an account, waiting, not seen', () => {
    const line = /data-testid="guild-roster-counts"[^>]*>([^<]*)</.exec(asOfficer)?.[1].trim();
    expect(line).toBe(
      `${home.roster.length + 2} on the roster · 1 in game without an account · ${home.pending?.length} waiting for approval · 1 not seen in game yet`,
    );
    expect(asOfficer.indexOf('guild-roster-counts')).toBeLessThan(
      asOfficer.indexOf('data-testid="guild-roster-filters"'),
    );
    expect(text(asOfficer)).toContain('Loot and readiness list members with gear on the site.');
  });

  it('leaves a zero part out of the count line, and derives the counts from the account rows when the home sends none', () => {
    const bare = render(GuildRosterTable, {
      props: { ...props(home), members: undefined, counts: undefined, can: [] },
    }).body;
    const line = /data-testid="guild-roster-counts"[^>]*>([^<]*)</.exec(bare)?.[1].trim();
    expect(line).toBe(`${home.roster.length} on the roster · ${home.pending?.length} waiting for approval`);
  });

  it('lists the names with no account row after the account rows, each with a rank word and a state word', () => {
    const entries = [...asOfficer.matchAll(/data-testid="guild-roster-entry" data-state="([a-z_]+)"/g)].map(
      (m) => m[1],
    );
    expect(entries).toEqual(['no_account', 'not_seen']);
    const section = asOfficer.slice(asOfficer.indexOf('data-testid="guild-roster-entries"'));
    expect(text(section)).toContain('Brightwick');
    expect(text(section)).toContain('No account yet');
    expect(text(section)).toContain('Driftwood');
    expect(text(section)).toContain('Not seen in game yet');
    // The snapshot's level and a relative seen time, never a clock.
    expect(text(section)).toMatch(/Level 58 · seen 2 hours ago/);
    expect(asOfficer.indexOf('data-testid="guild-roster-rows"')).toBeLessThan(
      asOfficer.indexOf('data-testid="guild-roster-entries"'),
    );
    // A name with no account has neither Approve nor Remove.
    expect(section).not.toContain('guild-roster-approve');
    expect(section).not.toContain('guild-roster-remove');
  });

  it('gives every account row a rank word, Member included', () => {
    const byName = rows(asOfficer);
    const rank = (name: string) =>
      text(/data-testid="guild-roster-rank"[^>]*>([^<]*)</.exec(byName.get(name) ?? '')?.[1] ?? '');
    expect(rank('Kraggor')).toBe('Leader');
    expect(rank('Sunderfel')).toBe('Officer');
    expect(rank('Zulmara')).toBe('Member');
  });

  it('prints no rank for a name only the addon saw: the state word alone, "Rank not known" on hover', () => {
    const section = asOfficer.slice(asOfficer.indexOf('data-testid="guild-roster-entries"'));
    const noAccount = section.split('data-testid="guild-roster-entry" ')[1];
    expect(noAccount).toContain('data-state="no_account"');
    expect(noAccount).not.toContain('data-testid="guild-roster-entry-rank"');
    expect(text(noAccount)).not.toContain('Member');
    expect(text(noAccount)).toContain('No account yet');
    expect(noAccount).toMatch(/title="[^"]*Rank not known until the addon sends ranks/);
    // The invited member who is not yet seen did join as a member, and says so.
    const invited = section.split('data-testid="guild-roster-entry" ')[2];
    expect(text(invited)).toContain('Member');
  });

  it('says once above the list that ranks show once the addon sends them, only when a row has none', () => {
    expect(text(asOfficer)).toContain('Ranks show once the addon sends them.');
    expect(asOfficer.split('Ranks show once the addon sends them.').length - 1).toBe(1);
    const none = render(GuildRosterTable, {
      props: { ...props(home), members: undefined, counts: undefined, can: [] },
    }).body;
    expect(none).not.toContain('guild-roster-ranks-pending');
  });
  it('stacks each account row as chips, one-line professions, then in-game facts, with Remove on the name row', () => {
    const zulmara = rows(asOfficer).get('Zulmara') ?? '';
    const chipsAt = zulmara.indexOf('data-testid="guild-roster-facts"');
    const profAt = zulmara.indexOf('data-testid="guild-roster-professions"');
    const factsAt = zulmara.indexOf('data-testid="guild-roster-ingame-slot"');
    expect(chipsAt).toBeGreaterThan(0);
    expect(profAt).toBeGreaterThan(chipsAt);
    expect(factsAt).toBeGreaterThan(profAt);
    expect(zulmara).toMatch(/class="[^"]*truncate[^"]*"[^>]*data-testid="guild-roster-professions"/);
  });
  it('gives the viewer no Remove on their own row, guild master or officer, and keeps it on everyone else', () => {
    for (const viewer of [VIEWER_OFFICER, 'Sunderfel']) {
      const own = buildMockHome(viewer);
      const html = render(GuildRosterTable, {
        props: { ...props(own), can: ['approve', 'remove', 'settings', 'award'] },
      }).body;
      const byName = rows(html);
      expect(byName.get(viewer), viewer).toBeDefined();
      expect(byName.get(viewer), viewer).not.toContain('data-testid="guild-roster-remove"');
      expect(byName.get('Zulmara'), viewer).toContain('data-testid="guild-roster-remove"');
    }
  });

  it('the waiting list names what Approve does, shows the rank the addon sent and the state word, and offers Decline', () => {
    const list = asOfficer.slice(asOfficer.indexOf('data-testid="guild-roster-waiting"'));
    expect(text(list)).toContain(
      'Approving lets this character see the roster, loot and readiness; officers can also remove members and change settings.',
    );
    const waitingRows = asOfficer
      .slice(
        asOfficer.indexOf('data-testid="guild-roster-unverified-list"'),
        asOfficer.indexOf('data-testid="guild-roster-addon-snapshot"'),
      )
      .split('guild-roster-row ')
      .slice(1);
    expect(waitingRows.length).toBe(home.pending?.length);
    for (const part of waitingRows) {
      // The card's head says it; the per-row state pill is gone (site alignment §5.10).
      expect(text(part)).not.toContain('Waiting for approval');
      expect(part).toMatch(/data-testid="guild-roster-pending-rank"[^>]*>(Member|Officer|Leader)</);
      expect(part).toContain('data-testid="guild-roster-decline"');
      expect(part).toContain('data-testid="guild-roster-approve"');
    }
  });

  it('draws Approve, Decline and Remove from the capability list, never from the rank', () => {
    // The viewer is the fixture's leader, but the account's list is empty: no controls at all.
    const none = officerAsk({ can: [] });
    for (const id of ['approve', 'approve-all', 'decline', 'remove']) {
      expect(none, id).not.toContain(`data-testid="guild-roster-${id}"`);
    }
    // approve alone: Approve and Approve all, no Remove or Decline.
    const approveOnly = officerAsk({ can: ['approve'] });
    expect(approveOnly).toContain('data-testid="guild-roster-approve-all"');
    expect(approveOnly).toContain('data-testid="guild-roster-approve"');
    expect(approveOnly).not.toContain('data-testid="guild-roster-remove"');
    expect(approveOnly).not.toContain('data-testid="guild-roster-decline"');
    // remove alone: Remove (only where the API says may_remove) and Decline, no Approve.
    const removeOnly = officerAsk({ can: ['remove'] });
    expect(removeOnly).not.toContain('data-testid="guild-roster-approve"');
    expect(removeOnly).toContain('data-testid="guild-roster-remove"');
    expect(removeOnly).toContain('data-testid="guild-roster-decline"');
    // A row the API marks may_remove: false never gets Remove, whatever the list says.
    const noRemoves = officerAsk({
      roster: home.roster.map((r) => ({ ...r, may_remove: false })),
      pending: (home.pending ?? []).map((r) => ({ ...r, may_remove: false })),
      can: ['approve', 'remove'],
    });
    expect(noRemoves).not.toContain('data-testid="guild-roster-remove"');
    expect(noRemoves).not.toContain('data-testid="guild-roster-decline"');
  });

  it('a member (no capabilities) sees the table without any control, and still the count line', () => {
    const member = render(GuildRosterTable, { props: props(buildMockHome(VIEWER_MEMBER)) }).body;
    expect(member).toContain('data-testid="guild-roster-counts"');
    expect(member).not.toContain('guild-roster-waiting');
    expect(member).not.toContain('guild-roster-remove');
  });

  it('adds a "No account" filter chip beside "Verified only"', () => {
    const filters = asOfficer.slice(
      asOfficer.indexOf('data-testid="guild-roster-filters"'),
      asOfficer.indexOf('data-testid="guild-roster-sort"'),
    );
    expect(filters.indexOf('guild-roster-filter-verified')).toBeLessThan(
      filters.indexOf('guild-roster-filter-noaccount'),
    );
    expect(text(filters)).toContain('No account');
  });
});

describe('GuildRosterTable: the Supporter chip', () => {
  it('draws the chip once, on the account row whose payload says supporter', () => {
    const home = buildMockHome(VIEWER_MEMBER);
    const roster = home.roster.map((row, index) => ({ ...row, supporter: index === 0 }));
    const html = render(GuildRosterTable, { props: { ...props(home), roster } }).body;
    expect(html.split('data-testid="supporter-chip"')).toHaveLength(2);
    const plain = render(GuildRosterTable, { props: props(home) }).body;
    expect(plain).not.toContain('supporter-chip');
  });

  it('draws the chip on a pending account row, and never on a name-only row', () => {
    const home = buildMockHome(VIEWER_OFFICER);
    const pending = (home.pending ?? []).map((row) => ({ ...row, supporter: true }));
    const roster = home.roster.map((row) => ({ ...row, supporter: true }));
    const html = render(GuildRosterTable, { props: { ...props(home), pending, roster } }).body;
    const waiting = html.slice(
      html.indexOf('id="guild-roster-unverified"'),
      html.indexOf('data-testid="guild-roster-rows"'),
    );
    expect(pending.length).toBeGreaterThan(0);
    expect(waiting.split('data-testid="supporter-chip"')).toHaveLength(pending.length + 1);
    const entries = html.slice(html.indexOf('data-testid="guild-roster-entries"'));
    expect(entries).toContain('guild-roster-entries');
    expect(entries).not.toContain('supporter-chip');
  });
});

describe('GuildRosterTable attendance source', () => {
  const home = buildMockHome(VIEWER_MEMBER);
  const raidHome = {
    ...home,
    roster: home.roster.map((r) => ({
      ...r,
      attendance: { present: 7, nights: 8 },
      attendance_30: { present: 26, nights: 30 },
    })),
  };

  it('reads the two-window chip and the one footnote in raid mode', () => {
    const body = text(
      render(GuildRosterTable, { props: { ...props(raidHome), attendanceSource: 'raids' } }).body,
    );
    expect(body).toContain('Came 7 of 8 \u00b7 26 of 30 raids');
    expect(body.split('Came: raids attended of raids signed up for or attended').length - 1).toBe(1);
  });

  it('keeps report mode on nights with no footnote', () => {
    const body = text(render(GuildRosterTable, { props: props(raidHome) }).body);
    expect(body).toContain('7 of 8 nights');
    expect(body).not.toContain('Came: raids attended');
  });
});
