// web/src/components/guild/planner/RaidPlanner.test.ts
// The raid planner rendered through GuildRaid with svelte/server from a seeded controller (web
// layout spec 2026-11-05 §7.2): the head, the control bar, the own sign-up before the one Plate
// of sign-up Rows, the officer-only reads (Not signed up with Ping, Bench and swaps, the comp edit)
// and the rail's comp, cover and notes.
import { render } from 'svelte/server';
import { describe, expect, it } from 'vitest';
import type { GuildRaidSignup, GuildRaidView } from '../../../lib/guild/api';
import { notesEmptyOf, raidCopy } from '../../../lib/guild/raid-copy';
import { RaidPageController } from '../../../lib/guild/raid-page-controller.svelte';
import type { RaidPageSeed } from '../../../lib/guild/raid-page-types';
import { deps } from '../../../test-support/raid-controller-fakes';
import { CHICAGO, NEW_YORK, NOW, between, plannedRaid, text } from '../../../test-support/raid-fixtures';
import GuildRaid from '../../GuildRaid.svelte';

const ROUTE = { guild: { region: 'us', ruleset: 'pvp', slug: 'olympus-xxvii' }, raidId: '7' } as const;
const GUILD = { id: 2, name: 'Olympus XXVII', faction: 'horde' as const, mark: null };

function signup(name: string, extra: Partial<GuildRaidSignup> = {}): GuildRaidSignup {
  return {
    character_key: `us/pvp/${name.toLowerCase()}`,
    name,
    class: 'mage',
    spec: 'Frost',
    spec_key: 'mage-frost',
    group: 'ranged',
    role: 'dps',
    state: 'in',
    note: '',
    own: false,
    ...extra,
  };
}

const SIGNUPS = [
  signup('Frostquill', { note: '20 min late maybe' }),
  signup('Cinderglass', { state: 'bench', spec: 'Fire', spec_key: 'mage-fire' }),
];
const COVER = [
  {
    key: 'kings',
    name: 'Blessing of Kings',
    kind: 'buff' as const,
    spell_id: 20217,
    icon: 'spell_magic_magearmor',
    need: 1,
    per: 'raid' as const,
    providers_in: 0,
    providers_late: 0,
    providers_bench: 0,
    likely: 0,
    state: 'missing' as const,
  },
];

function full(extra: Partial<GuildRaidView> = {}): GuildRaidView {
  return plannedRaid({
    signups: SIGNUPS,
    comp_target: { tanks: 2, healers: 8, melee: 10, ranged: 20, source: 'default' },
    cover: COVER,
    not_signed_raiders: [
      {
        character_key: 'us/pvp/voidcall',
        name: 'Voidcall',
        class: 'priest',
        account_key: 'a',
        discord_linked: false,
      },
    ],
    ...extra,
  });
}

function html(seed: Partial<RaidPageSeed> = {}, now = NOW, viewerTimezone = CHICAGO): string {
  const controller = new RaidPageController(deps(), ROUTE, {
    guild: GUILD,
    access: 'member',
    load: 'ready',
    raid: full(),
    ...seed,
  });
  return render(GuildRaid, { props: { path: ROUTE, controller, now, viewerTimezone } }).body;
}

const officer = (extra: Partial<RaidPageSeed> = {}): string => html({ role: 'officer', ...extra });

describe('the head', () => {
  it('links back, tags the status and titles the zone in the faction colour', () => {
    const body = html();
    expect(body).toContain('href="/guild/us/pvp/olympus-xxvii#raids"');
    expect(text(between(body, 'raid-status-chip', 'raid-when'))).toBe('Planned');
    expect(body).toMatch(/data-testid="page-title" style="color: var\(--color-horde\);">Onyxia's Lair/);
    expect(body).not.toContain('raid-title-zone');
    expect(text(between(html({ raid: full({ title: 'Ony farm' }) }), 'raid-title-zone', 'raid-when'))).toBe(
      "Onyxia's Lair",
    );
  });

  it('reads the guild time, the viewer’s own when it differs, the countdown and the counts', () => {
    expect(text(between(html(), 'raid-when', 'raid-count'))).toBe(
      'Thursday 10 December · 8 pm to 11 pm CST · in 2 days, 8 hours',
    );
    expect(text(between(html({}, NOW, NEW_YORK), 'raid-when', 'raid-count'))).toContain('9 pm EST your time');
    expect(text(between(html(), 'raid-count', 'guild-tabs'))).toBe(
      '23 of 40 signed up · 21 In · 2 Late · 1 Bench',
    );
  });

  it('mutes the countdown of a cancelled or started raid', () => {
    expect(between(html(), 'raid-when', 'raid-count')).toContain('text-gold');
    expect(between(html({ raid: full({ status: 'cancelled' }) }), 'raid-when', 'raid-count')).not.toContain(
      'text-gold',
    );
    const started = html({}, new Date('2026-12-11T02:30:00Z'));
    expect(between(started, 'raid-when', 'raid-count')).not.toContain('text-gold');
  });

  it('gives officers the officer row and the guild tabs as links', () => {
    expect(html()).not.toContain('data-testid="raid-officer-row"');
    expect(officer()).toContain('data-testid="raid-officer-row"');
    expect(html()).toContain('href="/guild/us/pvp/olympus-xxvii#roster"');
  });
});

describe('the control bar', () => {
  it('groups by Role first, offering Status and Class', () => {
    const bar = between(html(), 'raid-planner-bar', 'raid-signup-strip');
    expect(bar).toMatch(/aria-checked="true"[^>]*data-testid="control-group-role"/);
    expect(bar).toContain('data-testid="control-group-status"');
    expect(bar).toContain('data-testid="control-group-class"');
    expect(bar).toContain('Group by role');
  });
});

describe('the body', () => {
  it('puts the own sign-up, then the sign-ups Plate, then the rail', () => {
    const body = html();
    const strip = body.indexOf('data-testid="raid-signup-strip"');
    const roster = body.indexOf('data-testid="raid-roster"');
    expect(strip).toBeGreaterThan(0);
    expect(roster).toBeGreaterThan(strip);
    expect(body.indexOf('data-testid="guild-overview-comp"')).toBeGreaterThan(roster);
    expect(body.indexOf('data-testid="raid-notes"')).toBeGreaterThan(body.indexOf('guild-overview-cover'));
    expect(body.match(/data-region="plate"/g)?.length).toBeGreaterThan(0);
  });

  it('draws the Raid-Helper groups with each row’s spec and note, and no Tag for the default In', () => {
    const body = html();
    const ranged = text(between(body, 'raid-roster-ranged', 'raid-roster-bench'));
    expect(ranged).toContain('Ranged 1');
    expect(ranged).toContain('Frostquill Frost Mage · DPS 20 min late maybe');
    expect(ranged).not.toContain('maybe In');
    expect(text(between(body, 'raid-roster-melee', 'raid-roster-ranged'))).toContain('No melee yet.');
    expect(body).not.toContain('raid-roster-state-us/pvp/frostquill');
  });

  it('gives officers the in-row control instead of the Tag', () => {
    const body = officer();
    expect(body).toMatch(/aria-checked="true"[^>]*data-testid="raid-roster-state-us\/pvp\/frostquill-in"/);
    expect(body).not.toContain('data-testid="raid-roster-state"');
  });

  it('disables the officer control of a cancelled raid; with no plan it stays live (free calendar)', () => {
    const cancelled = officer({ raid: full({ status: 'cancelled' }) });
    expect(cancelled).toMatch(/class="opacity-60" data-testid="raid-roster"/);
    expect(
      /<button[^>]*data-testid="raid-roster-state-us\/pvp\/frostquill-in"[^>]*>/.exec(cancelled)?.[0],
    ).toMatch(/ disabled/);
    const noPlan = officer({ officerToolsField: false });
    expect(
      /<button[^>]*data-testid="raid-roster-state-us\/pvp\/frostquill-in"[^>]*>/.exec(noPlan)?.[0],
    ).not.toMatch(/ disabled/);
  });

  it('says nobody has signed up yet', () => {
    expect(html({ raid: full({ signups: [], zone_kind: 'event' }) })).toContain('Nobody has signed up yet.');
  });

  it('shows the poll pause line', () => {
    expect(html()).not.toContain(raidCopy.poll);
    expect(html({ pollPaused: true })).toContain(raidCopy.poll);
  });
});

describe('officer sections', () => {
  it('lists who has not signed up, with Ping on Discord waiting for the raid post', () => {
    const section = between(officer(), 'raid-not-signed', 'guild-overview-comp');
    expect(text(section)).toContain('Voidcall Priest · DPS · No Discord linked');
    expect(section).toMatch(/disabled[^>]*data-testid="raid-ping-discord"/);
    expect(text(section)).toContain('Post the raid to Discord first, then ping under it.');
    expect(html()).not.toContain('raid-not-signed');
  });

  it('lets the ping go once the raid is posted', () => {
    const posted = full({
      discord_post: { posted_at: '2026-12-07T18:00:00Z', post_at: null, reminded_at: null, remind: false },
    });
    const section = between(officer({ raid: posted }), 'raid-not-signed', 'guild-overview-comp');
    expect(/<button[^>]*data-testid="raid-ping-discord"/.exec(section)?.[0]).not.toMatch(
      /\sdisabled(=|\s|>)/,
    );
  });

  it('lists the bench with Swap in, officers only', () => {
    const bench = text(between(officer(), 'raid-bench', 'guild-overview-cover'));
    expect(bench).toContain('Cinderglass Fire Mage · DPS Swap in');
    expect(html()).not.toContain('data-testid="raid-bench"');
  });

  it('offers officers the comp edit, members the tally alone', () => {
    expect(officer()).toContain('data-testid="raid-comp-edit"');
    expect(html()).not.toContain('raid-comp-edit');
    expect(html()).toContain('data-testid="guild-comp-grid"');
  });
});

describe('the rail', () => {
  it('draws the notes, or the empty line, and no cover', () => {
    // The cover is the Overview's (batch 3): the planner's rail is the comp, the jobs and the notes.
    expect(html()).not.toContain('guild-overview-cover');
    expect(text(between(html(), 'raid-notes-empty', 'never'))).toContain(notesEmptyOf('raid', false));
    expect(html({ raid: full({ notes_html: '<p><strong>Flasks</strong></p>' }) })).toContain(
      '<strong>Flasks</strong>',
    );
  });

  it('leaves the comp and bench off a guild event', () => {
    const event = html({ role: 'officer', raid: full({ zone_kind: 'event', comp_target: null }) });
    expect(event).not.toContain('guild-overview-comp');
    expect(event).not.toContain('data-testid="raid-bench"');
  });
});
