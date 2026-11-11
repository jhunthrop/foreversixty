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
import { CHICAGO, NOW, between, doneRaid, plannedRaid, text } from '../test-support/raid-fixtures';
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
  saveSignup: never,
  setStatus: never,
  updateRaid: never,
  loadMatchCandidates: never,
  putMatch: never,
  deleteMatch: never,
  putOverride: never,
  deleteOverride: never,
  postReport: never,
  putCompTarget: never,
  deleteCompTarget: never,
  pingRaid: never,
  postRaidNow: never,
  putAssignment: never,
  deleteAssignment: never,
  postAssignments: never,
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
  it('draws the officer row for officers only', () => {
    expect(html()).not.toContain('data-testid="raid-officer-row"');
    const officer = html({ role: 'officer' });
    expect(officer).toContain('data-testid="raid-officer-row"');
    // The calendar is free (owner ruling 2026-11-09): no plan gate on the officer row.
    expect(html({ role: 'officer', officerToolsEnded: true })).not.toContain('id="officer-tools-gate-raids"');
  });

  it('draws the tabs as links to the guild page, Raids active', () => {
    const body = html();
    expect(body).toContain(`href="${GUILD_PATH}#roster"`);
    expect(body).toMatch(/aria-current="page"[^>]*data-tab-id="raids"/);
  });
});

describe('the body', () => {
  it('draws the notes lists with their bullets and numbers', () => {
    const body = ready(plannedRaid({ notes_html: '<ul><li>Cloak</li></ul>' }));
    expect(body).toMatch(/class="[^"]*\[(?:&|&amp;)_ul\]:list-disc[^"]*"[^>]*data-testid="raid-notes-prose"/);
    expect(body).toMatch(
      /class="[^"]*\[(?:&|&amp;)_ol\]:list-decimal[^"]*"[^>]*data-testid="raid-notes-prose"/,
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

  it('puts a cancelled race line above the cancelled strip, focusable', () => {
    const body = ready(plannedRaid({ status: 'cancelled' }), { raceLine: raidCopy.races.cancelled });
    expect(body).toMatch(/role="alert"[^>]*tabindex="-1"[^>]*data-testid="raid-race-line"/);
    expect(body.indexOf('raid-race-line')).toBeLessThan(body.indexOf('raid-signup-strip'));
    expect(text(between(body, 'raid-race-line', 'raid-signup-strip'))).toBe(raidCopy.races.cancelled);
    expect(html()).not.toContain('raid-race-line');
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

  it('is the report: the Done tag in the head, the match line under it, no sign-up count or night line', () => {
    const body = doneHtml();
    expect(body).not.toContain('23 of 40 signed up');
    expect(body).not.toContain('raid-night-line');
    expect(text(between(body, 'raid-match-line', 'raid-report'))).toContain(
      `${attendanceCopy.match.loggedAsNight} ${attendanceCopy.match.openNight}`,
    );
    expect(text(between(body, 'raid-status-chip', 'raid-when'))).toContain('Done');
  });

  it('shows members the lines without Not this log, and officers with it', () => {
    expect(doneHtml()).not.toContain('raid-not-this-log');
    expect(doneHtml({ role: 'officer' })).toContain('raid-not-this-log');
  });

  it('draws no count line and no report on a done raid without a night', () => {
    const body = doneHtml({}, doneRaid({ night: null }));
    expect(body).not.toContain('raid-count');
    expect(body).not.toContain('id="report"');
    expect(body).toContain('raid-match-line');
  });

  it('keeps an empty polite live region, which says the log was matched once the poll flips', () => {
    expect(doneHtml()).toMatch(/role="status"[^>]*data-testid="raid-live-region">\s*<\/p>/);
    const raised = doneHtml({ logMatched: true });
    expect(text(between(raised, 'raid-live-region', 'raid-match-line'))).toContain(attendanceCopy.liveRegion);
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
    expect(body).toMatch(/aria-current="page"[^>]*data-tab-id="raids"/);
    expect(body.indexOf('raid-band-skeleton')).toBeLessThan(body.indexOf('data-tab-id="raids"'));
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
