// web/src/components/guild/report/report-parts.test.ts
// The raid report's parts on their own (web layout spec 2026-11-05 §7.2 "Raid report"), server
// rendered: the pulls Plate, one pull's deaths, DPS and HPS, consumables by pull, attendance with
// the officer fix, loot as ItemRows, the Discord summary and Your night. Opening a pull is the
// e2e's (guild-attendance.spec.ts): the server render has no media query or clicks.
import { render } from 'svelte/server';
import type { Component } from 'svelte';
import { describe, expect, it } from 'vitest';
import type { RaidReportPull } from '../../../lib/guild/attendance-contract';
import { attendanceCopy } from '../../../lib/guild/attendance-copy';
import { reportNightCopy as copy } from '../../../lib/guild/report-night-copy';
import { doneRaid, raidNight, text, viewerCharacter } from '../../../test-support/raid-fixtures';
import { attendanceRow, raidReport } from '../../../test-support/report-fixtures';
import PullDeaths from './PullDeaths.svelte';
import ReportAttendance from './ReportAttendance.svelte';
import ReportConsumables from './ReportConsumables.svelte';
import ReportDiscord from './ReportDiscord.svelte';
import ReportLoot from './ReportLoot.svelte';
import ReportOutput from './ReportOutput.svelte';
import ReportPulls from './ReportPulls.svelte';
import ReportYourNight from './ReportYourNight.svelte';

function html<P extends Record<string, unknown>>(component: Component<P>, props: P): string {
  return render(component, { props }).body;
}

const report = raidReport();
const noop = (): void => {};
const count = (body: string, testid: string): number =>
  body.match(new RegExp(`data-testid="${testid}"`, 'g'))?.length ?? 0;

describe('ReportPulls', () => {
  const body = html(ReportPulls, { report });

  it('is one Plate with a row per pull and its result tag', () => {
    expect(count(body, 'raid-pulls')).toBe(1);
    expect(count(body, 'raid-pull-row')).toBe(3);
    const tags = body
      .split('data-testid="raid-pull-result"')
      .slice(1)
      .map((part) => text(`<x ${part}`).split(' ')[0]);
    expect(tags).toEqual(['Clear', 'Wipe', 'Kill']);
  });

  it('makes only the pulls with something to show a button, closed until pressed', () => {
    expect(count(body, 'raid-pull-open')).toBe(2);
    expect(body).toMatch(/aria-expanded="false"[^>]*data-testid="raid-pull-open">Onyxia, pull 1</);
    expect(body).not.toContain('raid-pull-inspector');
  });

  it('says so when the log has no pulls', () => {
    expect(html(ReportPulls, { report: raidReport({ pulls: [] }) })).toContain(copy.pulls.none);
  });
});

describe('PullDeaths', () => {
  const iconOf = (id: number | null): string | null => (id === 18435 ? '/data/b/icons/flame.webp' : null);
  const wipe = report.pulls[1];

  it('writes who, when, the killing blow with its icon and the hits before it', () => {
    const body = html(PullDeaths, { pull: wipe, iconOf });
    expect(count(body, 'raid-pull-death')).toBe(2);
    expect(text(body)).toContain('Pyrewisp 1:42 in');
    expect(text(body)).toContain('Flame Breath from Onyxia , 4,210 (1,050 overkill)');
    expect(body).toContain('src="/data/b/icons/flame.webp"');
    expect(count(body, 'raid-pull-last-hit')).toBe(1);
    expect(text(body)).toContain('Melee from Onyxian Whelp , 310');
  });

  it('says the log showed no hit, and a pull not read yet says that in one muted line', () => {
    expect(html(PullDeaths, { pull: wipe, iconOf })).toContain(copy.pulls.noHit);
    const unread: RaidReportPull = { ...wipe, death_causes: null };
    const body = html(PullDeaths, { pull: unread, iconOf });
    expect(body).toMatch(/class="text-muted"[^>]*data-testid="raid-pull-not-read">/);
    expect(body).not.toContain('text-wipe');
  });
});

describe('ReportOutput', () => {
  it('lists each raider with the guild unit, never the word parse', () => {
    const body = html(ReportOutput, { parses: report.parses, officer: true });
    expect(count(body, 'raid-output-row')).toBe(2);
    expect(text(body)).toContain('812 DPS');
    expect(text(body)).toContain('1,204 HPS');
    expect(text(body).toLowerCase()).not.toContain('parse');
  });

  it('says why there is nothing, in the viewer’s words', () => {
    expect(html(ReportOutput, { parses: [], officer: true })).toContain(copy.output.officerEmpty);
    expect(html(ReportOutput, { parses: [], officer: false })).toContain(copy.output.memberEmpty);
  });
});

describe('ReportConsumables', () => {
  it('draws the grid, its legend and the unknown pull in one muted line, never red', () => {
    const body = html(ReportConsumables, { byPull: report.consumables.by_pull });
    expect(body).toContain('raid-consumables-grid');
    expect(body).toContain(copy.consumables.legend);
    expect(body).toMatch(/class="text-muted[^"]*"[^>]*data-testid="raid-consumables-unknown"/);
    expect(text(body)).toContain('Onyxia, pull 1: the log carried no character sheet');
    expect(body).not.toContain('This log carries no buff data');
  });

  it('says there is nothing to read with no boss pull', () => {
    expect(html(ReportConsumables, { byPull: [] })).toContain(copy.consumables.none);
  });
});

describe('ReportAttendance', () => {
  const night = raidNight();
  const fixed = attendanceRow({
    character_key: 'k2',
    name: 'Menda',
    outcome: 'present',
    auto_outcome: 'no_show',
    override: { status: 'bench', by_name: 'Kraggor', at: '2026-12-11T06:00:00Z' },
  });
  const props = { night, rows: [attendanceRow(), fixed], readOnly: false, gateId: 'gate', onFix: noop };

  it('gives a member the summary sentence only', () => {
    const body = html(ReportAttendance, { ...props, officer: false });
    expect(text(body)).toContain('22 of 23 came');
    expect(body).not.toContain('raid-attendance-row');
  });

  it('gives an officer every row with the fix, the auto value muted under a fix and Back to the log', () => {
    const body = html(ReportAttendance, { ...props, officer: true });
    expect(count(body, 'raid-attendance-row')).toBe(2);
    expect(body).toContain('raid-attendance-fix-k2');
    expect(text(body)).toContain('Bench, set by Kraggor');
    expect(text(body)).toContain('The log said Not in the log');
    expect(count(body, 'raid-attendance-back')).toBe(1);
  });

  it('disables the fix in officer-read-only and names the gate', () => {
    const body = html(ReportAttendance, { ...props, officer: true, readOnly: true });
    expect(body).toMatch(/disabled[^>]*aria-describedby="gate"[^>]*data-testid="raid-attendance-back"/);
  });

  it('shows a refused fix under its row', () => {
    const body = html(ReportAttendance, { ...props, officer: true, lines: { k2: 'That did not save.' } });
    expect(body).toMatch(/role="alert"[^>]*data-testid="raid-attendance-line">That did not save./);
  });
});

describe('ReportLoot', () => {
  it('draws every award as an ItemRow with the winner, the boss only when the night had several', () => {
    const body = html(ReportLoot, { groups: report.loot, officer: true, lootHref: '/g#loot' });
    expect(count(body, 'raid-loot-row')).toBe(1);
    expect(text(body)).toContain('Head of Onyxia');
    expect(text(body)).toContain('to Kraggor');
    expect(text(body)).not.toContain('Onyxia · to Kraggor');
    const twoBosses = [...report.loot, { ...report.loot[0], encounter_id: 663, boss: 'Lucifron' }];
    expect(text(html(ReportLoot, { groups: twoBosses, officer: true, lootHref: '/g#loot' }))).toContain(
      'Onyxia · to Kraggor',
    );
  });

  it('points an officer at the Loot tab when nothing was awarded', () => {
    const officer = html(ReportLoot, { groups: [], officer: true, lootHref: '/g#loot' });
    expect(officer).toContain('href="/g#loot"');
    expect(html(ReportLoot, { groups: [], officer: false, lootHref: '/g#loot' })).toContain(
      attendanceCopy.loot.memberEmpty,
    );
  });
});

describe('ReportDiscord', () => {
  const base = {
    title: "Onyxia's Lair",
    report,
    night: raidNight(),
    connected: true,
    settingsHref: '/g/settings',
    officerTools: true,
    posting: false,
    postLine: null,
    viewerTimezone: 'America/Chicago',
    locale: 'en-GB',
    copyText: () => '',
    onPost: noop,
  };

  it('previews the post and offers Post to Discord before the first post', () => {
    const body = html(ReportDiscord, base);
    expect(text(body)).toContain("Report: Onyxia's Lair");
    expect(text(body)).toContain('14 deaths');
    expect(body).toMatch(/data-testid="raid-report-post"[^>]*>(<!--[^>]*-->)*Post to Discord/);
  });

  it('after a post shows its time and Post again', () => {
    const body = html(ReportDiscord, { ...base, report: raidReport({ posted_at: '2026-12-11T05:58:00Z' }) });
    expect(body).toMatch(/data-testid="raid-report-post"[^>]*>(<!--[^>]*-->)*Post again/);
    expect(body).toContain('raid-report-posted');
  });

  it('points at Settings without a channel, and disables Post in officer-read-only', () => {
    const none = html(ReportDiscord, { ...base, connected: false });
    expect(none).not.toContain('raid-report-post"');
    expect(none).toContain('href="/g/settings"');
    expect(html(ReportDiscord, { ...base, officerTools: false })).toMatch(
      /disabled[^>]*data-testid="raid-report-post"/,
    );
  });
});

describe('ReportYourNight', () => {
  it('writes the sentence, the facts and the loot for each own character', () => {
    const raid = doneRaid({ attendance: [attendanceRow({ own: true })], report });
    const body = html(ReportYourNight, { raid, locale: 'en-GB' });
    expect(count(body, 'raid-your-night-row')).toBe(1);
    expect(body).toContain('raid-your-night-sentence');
    expect(text(body)).toContain('Head of Onyxia');
  });

  it('keeps a muted row for a character waiting for approval', () => {
    const raid = doneRaid({
      attendance: [],
      report,
      viewer: { ...doneRaid().viewer, characters: [viewerCharacter({ verified: false })] },
    });
    expect(html(ReportYourNight, { raid, locale: 'en-GB' })).toContain('raid-your-night-waiting');
  });
});
