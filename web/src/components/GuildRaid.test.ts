// web/src/components/GuildRaid.test.ts
// The raid page rendered with svelte/server from a seeded controller (raids spec 3.2, 4 "Page
// states", R1): the band lines, the officer row for officers only, the tabs as links to the guild
// page, the notes, the locked, ended, cancelled and done states, the poll pause line, the
// not-visible page for every viewer who is not a verified member, and the loading reserve.
import { render } from 'svelte/server';
import { afterEach, describe, expect, it, vi } from 'vitest';
import type { GuildRaidView } from '../lib/guild/api';
import { attendanceCopy } from '../lib/guild/attendance-copy';
import { guildClaimCopy } from '../lib/guild/copy';
import { raidCopy } from '../lib/guild/raid-copy';
import { RAID_BAND_LINES_MIN, RAID_STRIP_RESERVE } from '../lib/guild/raid-layout';
import { RAID_ROWS_KEY } from '../lib/guild/reserve-keys';
import { RaidPageController } from '../lib/guild/raid-page-controller.svelte';
import type { RaidPageDeps, RaidPageSeed } from '../lib/guild/raid-page-types';
import { CHICAGO, NEW_YORK, NOW, between, doneRaid, plannedRaid, text } from '../test-support/raid-fixtures';
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
  loadMatchCandidates: never,
  putMatch: never,
  deleteMatch: never,
  putExcused: never,
  postReport: never,
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

  it('keeps the spaces around every separator, which the md-up inline layout shows as is', () => {
    const body = html({ raid: plannedRaid({ title: 'Ony farm' }) }, NEW_YORK);
    const markup = (from: string, to: string): string =>
      between(body, from, to)
        .replace(/<!--[\s\S]*?-->/g, '')
        .replace(/<[^>]+>/g, '');
    expect(markup('raid-when', 'raid-count').trim()).toBe(
      'Thursday 10 December · 9:00 pm to 12:00 am · 8:00 pm CST guild time · in 2 days 8 h',
    );
    expect(markup('raid-count', 'guild-tabs')).toContain('23 of 40 signed up · 3 tanks');
    expect(markup('raid-title', 'raid-when').trim()).toBe("Ony farm · Onyxia's Lair");
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
      '23 of 40 signed up · 3 tanks · 7 healers · 13 DPS · 2 late',
    );
    expect(body).toContain('width: 57%');
    expect(ready(plannedRaid({ size: null }))).not.toContain('bg-gold absolute');
  });

  it('marks a raid still ahead in gold, and a cancelled or started one muted', () => {
    const live = between(html(), 'raid-when', 'guild-tabs');
    expect(live).toMatch(
      /whitespace-nowrap text-gold">(<!--[^>]*-->)*in (<!--[^>]*-->)*<span class="numeric">2<\/span>/,
    );
    expect(live).toContain('bg-gold absolute');
    const cancelled = between(ready(plannedRaid({ status: 'cancelled' })), 'raid-when', 'guild-tabs');
    expect(cancelled).not.toContain('text-gold');
    expect(cancelled).not.toContain('bg-gold');
    expect(cancelled).toMatch(
      /whitespace-nowrap text-muted">(<!--[^>]*-->)*in (<!--[^>]*-->)*<span class="numeric">2<\/span>/,
    );
    expect(cancelled).toContain('bg-muted absolute');
    expect(cancelled).toMatch(/text-muted numeric text-\[15px\]">23 of 40 signed up/);
    const controller = new RaidPageController(DEPS, ROUTE, {
      guild: GUILD,
      access: 'member',
      load: 'ready',
      raid: plannedRaid(),
    });
    const started = render(GuildRaid, {
      props: { path: ROUTE, controller, now: new Date('2026-12-11T02:30:00Z'), viewerTimezone: CHICAGO },
    }).body;
    expect(between(started, 'raid-when', 'guild-tabs')).toMatch(
      /whitespace-nowrap text-muted">(<!--[^>]*-->)*started (<!--[^>]*-->)*<span class="numeric">30<\/span>/,
    );
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

  it('draws the notes lists with their bullets and numbers', () => {
    const body = ready(plannedRaid({ notes_html: '<ul><li>Cloak</li></ul>' }));
    expect(body).toMatch(/class="[^"]*\[(?:&|&amp;)_ul\]:list-disc[^"]*"[^>]*data-testid="raid-notes-prose"/);
    expect(body).toMatch(
      /class="[^"]*\[(?:&|&amp;)_ol\]:list-decimal[^"]*"[^>]*data-testid="raid-notes-prose"/,
    );
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

  it('says an ended raid ended with no log yet, in the band', () => {
    expect(html()).not.toContain(attendanceCopy.ended.noLog);
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
    expect(text(between(ended, 'raid-ended-line', 'guild-tabs'))).toContain(
      `${attendanceCopy.ended.noLog}${attendanceCopy.ended.matchSoon}`,
    );
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

  it('says the raid did not load, with Try again, even before access is known', () => {
    expect(text(html({ load: 'failed', access: 'loading', raid: null }))).toContain(raidCopy.loadErrors.raid);
    const body = html({ load: 'failed' });
    expect(text(between(body, 'raid-load-error', 'never'))).toContain(raidCopy.loadErrors.raid);
  });
});

describe('a raid with a matched log', () => {
  const AFTER = new Date('2026-12-11T12:00:00Z');

  function doneHtml(seed: Partial<RaidPageSeed> = {}, raid = doneRaid()): string {
    const controller = new RaidPageController(DEPS, ROUTE, {
      guild: GUILD,
      access: 'member',
      load: 'ready',
      raid,
      ...seed,
    });
    return render(GuildRaid, {
      props: { path: ROUTE, controller, now: AFTER, viewerTimezone: CHICAGO },
    }).body;
  }

  it('puts the night line where the count line was, with the match line under it', () => {
    const body = doneHtml();
    expect(body).not.toContain('23 of 40 signed up');
    expect(text(between(body, 'raid-night-line', 'raid-match-line'))).toBe(
      '22 of 23 came · 2 joined · 1 kill · 3 wipes · 14 deaths · 2 h 41 min',
    );
    expect(text(between(body, 'raid-match-line', 'guild-tabs'))).toContain(
      `${attendanceCopy.match.loggedAsNight} ${attendanceCopy.match.openNight}`,
    );
    expect(text(between(body, 'raid-status-chip', 'raid-title'))).toContain('Done');
  });

  it('shows members the lines without Not this log, and officers with it', () => {
    expect(doneHtml()).not.toContain('raid-not-this-log');
    expect(doneHtml({ role: 'officer' })).toContain('raid-not-this-log');
  });

  it('keeps the count line on a raid without a night', () => {
    const body = doneHtml({}, doneRaid({ night: null }));
    expect(text(between(body, 'raid-count', 'raid-match-line'))).toContain('23 of 40 signed up');
  });

  it('keeps an empty polite live region, which says the log was matched once the poll flips', () => {
    expect(doneHtml()).toMatch(/role="status"[^>]*data-testid="raid-live-region">\s*<\/p>/);
    const raised = doneHtml({ logMatched: true });
    expect(text(between(raised, 'raid-live-region', 'raid-title'))).toContain(attendanceCopy.liveRegion);
    expect(raised).toMatch(/class="sr-only"[^>]*role="status"/);
  });

  it('draws the ended line, not the match line, with Match a log for an officer, past a raid`s end', () => {
    const controller = new RaidPageController(DEPS, ROUTE, {
      guild: GUILD,
      access: 'member',
      load: 'ready',
      raid: plannedRaid(),
      role: 'officer',
    });
    const body = render(GuildRaid, {
      props: { path: ROUTE, controller, now: AFTER, viewerTimezone: CHICAGO },
    }).body;
    expect(text(body)).toContain(attendanceCopy.ended.noLog);
    expect(body).toContain('raid-match-open');
    expect(body).not.toContain('raid-match-line');
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
  it('draws the back link, the eyebrow and the tabs for real while the raid loads', () => {
    const body = html({ load: 'loading', raid: null });
    expect(text(between(body, 'raid-back-link', 'raid-band-skeleton'))).toContain(
      `${raidCopy.backLink('Olympus XXVII')} ${raidCopy.page.eyebrow}`,
    );
    expect(body).toContain(`href="${GUILD_PATH}#raids"`);
    expect(body).toMatch(/aria-selected="true"[^>]*data-guild-tab="raids"/);
    expect(body.indexOf('raid-band-skeleton')).toBeLessThan(body.indexOf('data-guild-tab="raids"'));
  });

  it('reserves the band lines, one strip row and the roster, and says Loading the raid', () => {
    const body = html({ load: 'loading', raid: null });
    expect(body).toContain(raidCopy.page.loading);
    expect(body).toContain(RAID_BAND_LINES_MIN);
    expect(body).toContain('--strip-rows: 1');
    expect(body).toContain(RAID_STRIP_RESERVE);
    expect(body).toContain('min-h-[560px] md:min-h-[420px]');
  });

  it('reserves the strip rows last drawn', () => {
    vi.stubGlobal('localStorage', { getItem: (key: string) => (key === RAID_ROWS_KEY ? '3' : null) });
    expect(html({ load: 'loading', raid: null })).toContain('--strip-rows: 3');
  });
});
