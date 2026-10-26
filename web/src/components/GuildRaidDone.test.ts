// web/src/components/GuildRaidDone.test.ts
// The done raid page (attendance spec 3.3 to 3.5, R-A1, R-A4): `Your night` replaces the strip for a
// member with an own row and never for an officer; the Report reads by rank; Attendance is the
// officers'; a matched log that turned private (night, report and attendance null) draws neither
// section nor an error.
import { render } from 'svelte/server';
import { describe, expect, it } from 'vitest';
import type { GuildRaidView } from '../lib/guild/api';
import { RaidPageController } from '../lib/guild/raid-page-controller.svelte';
import type { RaidPageDeps, RaidPageSeed } from '../lib/guild/raid-page-types';
import { CHICAGO, NOW, doneRaid, text } from '../test-support/raid-fixtures';
import { attendanceRow, raidReport } from '../test-support/report-fixtures';
import GuildRaid from './GuildRaid.svelte';

const ROUTE = { guild: { region: 'us', ruleset: 'pvp', slug: 'olympus-xxvii' }, raidId: '7' } as const;
const GUILD = { id: 2, name: 'Olympus XXVII', faction: 'horde' as const, mark: null };
const never = <T>(): Promise<T> => new Promise<T>(() => {});
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

function page(raid: GuildRaidView, seed: Partial<RaidPageSeed> = {}): string {
  const controller = new RaidPageController(DEPS, ROUTE, {
    guild: GUILD,
    access: 'member',
    load: 'ready',
    raid,
    ...seed,
  });
  return render(GuildRaid, {
    props: { path: ROUTE, controller, now: NOW, viewerTimezone: CHICAGO },
  }).body;
}

const OFFICER = { role: 'officer' as const };
const MINE = attendanceRow({ own: true });
const DONE = (extra: Partial<GuildRaidView> = {}): GuildRaidView =>
  doneRaid({ attendance: [MINE], report: raidReport(), ...extra });

describe('a done raid, member', () => {
  const body = page(DONE());

  it('shows Your night in the strip slot, no strip, the Report without officer blocks, no Attendance', () => {
    expect(body).toContain('data-testid="raid-your-night"');
    expect(body).not.toContain('data-testid="raid-signup-strip"');
    expect(body).toContain('id="report"');
    expect(body).toContain('data-testid="raid-kill-order"');
    expect(body).toContain('data-testid="raid-loot"');
    for (const hidden of [
      'raid-deaths"',
      'raid-moves"',
      'raid-consumables"',
      'id="attendance"',
      'raid-report-copy',
    ]) {
      expect(body).not.toContain(hidden);
    }
  });

  it('puts the Report before the notes', () => {
    expect(body.indexOf('id="report"')).toBeLessThan(body.indexOf('data-testid="raid-notes"'));
  });

  it('draws no Your night for a member with no own row', () => {
    expect(page(DONE({ attendance: [attendanceRow({ own: false })] }))).not.toContain('raid-your-night"');
  });
});

describe('a done raid, officer', () => {
  const body = page(
    DONE({ attendance: [MINE, attendanceRow({ character_key: 'k2', name: 'Menda' })] }),
    OFFICER,
  );

  it('has no Your night, every Report block and the Attendance section after it', () => {
    expect(body).not.toContain('raid-your-night"');
    for (const block of ['raid-deaths"', 'raid-moves"', 'raid-consumables"', 'raid-report-copy']) {
      expect(body).toContain(block);
    }
    expect(body.indexOf('id="report"')).toBeLessThan(body.indexOf('id="attendance"'));
    expect(body.indexOf('id="attendance"')).toBeLessThan(body.indexOf('data-testid="raid-notes"'));
  });

  it('washes the officer own row in the attendance rows', () => {
    expect(body).toContain('background:rgba(229,185,85,.06)');
  });
});

describe('a matched log that turned private or flagged', () => {
  const body = page(doneRaid({ night: null, report: null, attendance: null }), OFFICER);

  it('keeps the match line and draws no Report, Attendance, Your night or error', () => {
    expect(text(body)).toContain('Logged as a raid night.');
    expect(body).not.toContain('id="report"');
    expect(body).not.toContain('id="attendance"');
    expect(body).not.toContain('raid-your-night"');
    expect(body).not.toContain('role="alert"');
  });

  it('is the same for a member', () => {
    const member = page(doneRaid({ night: null, report: null, attendance: null }));
    expect(member).not.toContain('id="report"');
    expect(member).not.toContain('raid-your-night"');
  });
});
