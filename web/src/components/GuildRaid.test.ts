// web/src/components/GuildRaid.test.ts
// The raid page rendered with svelte/server from a seeded controller (raids spec 3.2, 4 "Page
// states", R1): the band lines, the officer row for officers only, the tabs as links to the guild
// page, the notes, the locked, ended, cancelled and done states, the poll pause line, the
// not-visible page for every viewer who is not a verified member, and the loading reserve.
import { render } from 'svelte/server';
import { afterEach, describe, expect, it, vi } from 'vitest';
import type { GuildRaidView } from '../lib/guild/api';
import { guildClaimCopy } from '../lib/guild/copy';
import { raidCopy } from '../lib/guild/raid-copy';
import { RAID_ROWS_KEY } from '../lib/guild/raid-page';
import {
  RaidPageController,
  type RaidPageDeps,
  type RaidPageSeed,
} from '../lib/guild/raid-page-controller.svelte';
import { CHICAGO, NEW_YORK, NOW, between, plannedRaid, text } from '../test-support/raid-fixtures';
import GuildRaid from './GuildRaid.svelte';

const ROUTE = { guild: { region: 'us', ruleset: 'pvp', slug: 'olympus-xxvii' }, raidId: '7' } as const;
const GUILD = { id: 2, name: 'Olympus XXVII', faction: 'horde' as const, mark: null };
const GUILD_PATH = '/guild/us/pvp/olympus-xxvii';

function never<T>(): Promise<T> {
  return new Promise<T>(() => {});
}

const DEPS: RaidPageDeps = {
  loadGuild: never,
  loadMe: never,
  loadHome: never,
  fetchRaid: never,
  refetchRaid: never,
  putSignup: never,
  setStatus: never,
  updateRaid: never,
  loadZones: never,
  isVisible: () => true,
  now: () => NOW,
  storage: null,
};

function html(seed: Partial<RaidPageSeed> = {}, viewerTimezone = CHICAGO): string {
  const controller = new RaidPageController(DEPS, ROUTE, {
    guild: GUILD,
    access: 'member',
    load: 'ready',
    raid: plannedRaid(),
    ...seed,
  });
  return render(GuildRaid, { props: { path: ROUTE, controller, now: NOW, viewerTimezone } }).body;
}

function ready(raid: GuildRaidView, extra: Partial<RaidPageSeed> = {}): string {
  return html({ raid, ...extra });
}

afterEach(() => {
  vi.unstubAllGlobals();
});

describe('the band', () => {
  it('links back to the guild Raids tab and reads eyebrow, chip and title', () => {
    const body = html();
    expect(body).toContain(`href="${GUILD_PATH}#raids"`);
    expect(text(between(body, 'raid-back-link', 'raid-status-chip'))).toContain(
      raidCopy.backLink('Olympus XXVII'),
    );
    expect(text(between(body, 'raid-status-chip', 'raid-title'))).toContain('Planned');
    expect(text(between(body, 'raid-title', 'raid-when'))).toBe("Onyxia's Lair");
  });

  it('follows a title that differs from the zone with the zone', () => {
    const body = ready(plannedRaid({ title: 'Ony farm' }));
    expect(text(between(body, 'raid-title', 'raid-when'))).toBe("Ony farm · Onyxia's Lair");
  });

  it('reads the when line with the countdown, and guild time only when the offsets differ', () => {
    const chicago = text(between(html(), 'raid-when', 'raid-count'));
    expect(chicago).toBe('Thursday 10 December · 8:00 pm to 11:00 pm · in 2 days 8 h');
    const newYork = text(between(html({}, NEW_YORK), 'raid-when', 'raid-count'));
    expect(newYork).toBe(
      'Thursday 10 December · 9:00 pm to 12:00 am · 8:00 pm CST guild time · in 2 days 8 h',
    );
  });

  it('reads the count line with the bar, and no bar while the size is unknown', () => {
    const body = html();
    expect(text(between(body, 'raid-count', 'guild-tabs'))).toBe(
      '23 of 40 signed up · 3 tanks · 7 healers · 13 DPS',
    );
    expect(body).toContain('width: 57%');
    expect(ready(plannedRaid({ size: null }))).not.toContain('bg-gold absolute');
  });

  it('draws the officer row for officers only', () => {
    expect(html()).not.toContain('data-testid="raid-officer-row"');
    const officer = html({ role: 'officer' });
    expect(officer).toContain('data-testid="raid-officer-row"');
    expect(html({ role: 'officer', officerToolsEnded: true })).toContain('id="officer-tools-gate-raids"');
  });

  it('draws the tabs as links to the guild page, Raids active', () => {
    const body = html();
    expect(body).toContain(`href="${GUILD_PATH}#roster"`);
    expect(body).toMatch(/aria-selected="true"[^>]*data-guild-tab="raids"/);
  });

  it('shows the poll pause line after three failed polls', () => {
    expect(html()).not.toContain(raidCopy.poll);
    expect(text(between(html({ pollPaused: true }), 'raid-poll-paused', 'guild-tabs'))).toBe(raidCopy.poll);
  });
});

describe('the body', () => {
  it('orders the strip, the notes and the roster', () => {
    const body = html();
    const strip = body.indexOf('data-testid="raid-signup-strip"');
    const notes = body.indexOf('data-testid="raid-notes"');
    expect(strip).toBeGreaterThan(0);
    expect(notes).toBeGreaterThan(strip);
    expect(body.indexOf('data-testid="raid-roster"')).toBeGreaterThan(notes);
  });

  it('renders the API notes, or the empty line for members and officers', () => {
    expect(ready(plannedRaid({ notes_html: '<p><strong>Flasks</strong></p>' }))).toContain(
      '<strong>Flasks</strong>',
    );
    expect(text(between(html(), 'raid-notes-empty', 'raid-roster'))).toBe(raidCopy.notesEmpty.member);
    expect(text(between(html({ role: 'officer' }), 'raid-notes-empty', 'raid-roster'))).toBe(
      raidCopy.notesEmpty.officer,
    );
  });
});

describe('page states', () => {
  it('puts the locked line above disabled segments', () => {
    const body = ready(plannedRaid({ status: 'locked' }));
    expect(text(between(body, 'raid-strip-state', 'never'))).toContain(raidCopy.raidStates.locked);
    expect(body).toMatch(/<button[^>]*role="radio"[^>]* disabled/);
  });

  it('says an ended raid has ended', () => {
    expect(html()).not.toContain(raidCopy.raidStates.ended);
    const controller = new RaidPageController(DEPS, ROUTE, {
      guild: GUILD,
      access: 'member',
      load: 'ready',
      raid: plannedRaid(),
    });
    const afterEnd = new Date('2026-12-11T06:00:00Z');
    const ended = render(GuildRaid, {
      props: { path: ROUTE, controller, now: afterEnd, viewerTimezone: CHICAGO },
    }).body;
    expect(text(ended)).toContain(raidCopy.raidStates.ended);
    expect(text(between(ended, 'raid-status-chip', 'raid-title'))).toContain('Ended');
  });

  it('replaces the strip on a cancelled raid and mutes the roster', () => {
    const body = ready(plannedRaid({ status: 'cancelled' }));
    expect(text(between(body, 'raid-strip-state', 'raid-notes'))).toBe(raidCopy.raidStates.cancelled);
    expect(body).toMatch(/class="[^"]*opacity-60[^"]*"[^>]*data-testid="raid-roster"/);
  });

  it('puts a cancelled race line above the cancelled strip, focusable', () => {
    const body = ready(plannedRaid({ status: 'cancelled' }), { raceLine: raidCopy.races.cancelled });
    expect(body).toMatch(/role="alert"[^>]*tabindex="-1"[^>]*data-testid="raid-race-line"/);
    expect(body.indexOf('raid-race-line')).toBeLessThan(body.indexOf('raid-signup-strip'));
    expect(text(between(body, 'raid-race-line', 'raid-signup-strip'))).toBe(raidCopy.races.cancelled);
    expect(html()).not.toContain('raid-race-line');
  });

  it('holds the officer roster selects on a cancelled or done raid', () => {
    const signups = [
      {
        character_key: 'us/pvp/bolt',
        name: 'Bolt',
        class: 'warrior',
        spec: 'Fury',
        role: 'dps' as const,
        state: 'in' as const,
        note: '',
        own: false,
      },
    ];
    for (const status of ['cancelled', 'done'] as const) {
      const body = ready(plannedRaid({ status, signups }), { role: 'officer' });
      expect(/<select[^>]*data-testid="raid-roster-select-Bolt"[^>]*>/.exec(body)?.[0]).toMatch(/ disabled/);
    }
  });

  it('logs a done raid with the link to the night', () => {
    const body = ready(plannedRaid({ status: 'done' }));
    expect(text(between(body, 'raid-done', 'guild-tabs'))).toBe(
      `${raidCopy.raidStates.done} ${raidCopy.raidStates.openNight}`,
    );
  });

  it('says the raid did not load, with Try again, even before access is known', () => {
    expect(text(html({ load: 'failed', access: 'loading', raid: null }))).toContain(raidCopy.loadErrors.raid);
    const body = html({ load: 'failed' });
    expect(text(between(body, 'raid-load-error', 'never'))).toContain(raidCopy.loadErrors.raid);
  });
});

describe('not visible (R1)', () => {
  it('asks a signed-out visitor to sign in, under the guild band', () => {
    const body = html({ access: 'signed-out', raid: null });
    expect(body).toContain('data-testid="raid-sign-in-prompt"');
    expect(text(body)).toContain(raidCopy.notVisible.members('Olympus XXVII'));
    expect(body).not.toContain('raid-signup-strip');
  });

  it('tells a signed-in stranger the site knows no character of theirs here', () => {
    const body = text(html({ access: 'not-member', raid: null }));
    expect(body).toContain(raidCopy.notVisible.members('Olympus XXVII'));
    expect(body).toContain(guildClaimCopy.notAMember);
  });

  it('tells an unverified member to wait for approval', () => {
    const body = text(html({ access: 'unverified', raid: null }));
    expect(body).toContain(raidCopy.notVisible.unverified);
    expect(body).not.toContain(raidCopy.notVisible.members('Olympus XXVII'));
  });
});

describe('loading', () => {
  it('reserves the band, one strip row and the roster, and says Loading the raid', () => {
    const body = html({ load: 'loading', raid: null });
    expect(body).toContain(raidCopy.page.loading);
    expect(body).toContain('min-h-[300px] md:min-h-[232px]');
    expect(body).toContain('--strip-rows: 1');
    expect(body).toContain('min-h-[560px] md:min-h-[420px]');
  });

  it('reserves the strip rows last drawn', () => {
    vi.stubGlobal('localStorage', { getItem: (key: string) => (key === RAID_ROWS_KEY ? '3' : null) });
    expect(html({ load: 'loading', raid: null })).toContain('--strip-rows: 3');
  });
});
