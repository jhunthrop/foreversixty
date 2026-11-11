// web/src/components/GuildRaidDone.test.ts
// The done raid page, the raid report (web layout spec 2026-11-05 §7.2): the answer figures in the
// head; `Your night` for a member with an own row and never for an officer; the pulls, DPS and HPS
// and consumables for everyone (the API scopes a member's to their own); the officers' attendance
// rows with the fix and the Discord summary; a matched log that turned private (night, report and
// attendance null) draws none of it and no error.
import { render } from 'svelte/server';
import { describe, expect, it } from 'vitest';
import type { GuildRaidView } from '../lib/guild/api';
import { RaidPageController } from '../lib/guild/raid-page-controller.svelte';
import type { RaidPageDeps, RaidPageSeed } from '../lib/guild/raid-page-types';
import { CHICAGO, NOW, between, doneRaid, text } from '../test-support/raid-fixtures';
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

  it('heads with the answer figures: kills, wipes, attempts, deaths', () => {
    expect(text(between(body, 'raid-report-figures', 'guild-tabs'))).toBe(
      'Kills 1 Wipes 3 Attempts 4 Deaths 14',
    );
  });

  it('shows Your night, the pulls, DPS and HPS, consumables, attendance and loot; no strip, no Discord', () => {
    for (const shown of [
      'raid-your-night"',
      'raid-pulls"',
      'raid-output"',
      'raid-consumables"',
      'raid-attendance"',
      'raid-loot"',
      'id="report"',
    ]) {
      expect(body).toContain(shown);
    }
    for (const hidden of ['raid-signup-strip', 'raid-discord"', 'raid-report-copy', 'raid-attendance-row']) {
      expect(body).not.toContain(hidden);
    }
  });

  it('puts Your night before the pulls', () => {
    expect(body.indexOf('raid-your-night"')).toBeLessThan(body.indexOf('raid-pulls"'));
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

  it('has no Your night, the attendance fix behind Fix attendance, and the Discord summary', () => {
    expect(body).not.toContain('raid-your-night"');
    expect(body).toContain('raid-attendance-fix-open');
    expect(body).not.toContain('data-testid="raid-attendance-row"');
    expect(body).toContain('raid-discord"');
    expect(body).toContain('raid-report-copy');
  });

  it('opens the fix on its own once a row has been fixed: every raider, each with its fix', () => {
    const fixed = page(
      DONE({
        attendance: [
          MINE,
          attendanceRow({
            character_key: 'k2',
            name: 'Menda',
            override: { status: 'bench', by_name: 'Kraggor', at: '2026-12-11T06:00:00Z' },
          }),
        ],
      }),
      OFFICER,
    );
    expect(fixed.match(/data-testid="raid-attendance-row"/g)).toHaveLength(2);
    expect(fixed).toContain('raid-attendance-fix-k2');
  });

  it('keeps the rail after the main column: attendance, loot, Discord', () => {
    const order = ['raid-pulls"', 'raid-attendance"', 'raid-loot"', 'raid-discord"'].map((id) =>
      body.indexOf(id),
    );
    expect([...order].sort((a, b) => a - b)).toEqual(order);
  });
});

describe('a matched log that turned private or flagged', () => {
  const body = page(doneRaid({ night: null, report: null, attendance: null }), OFFICER);

  it('keeps the match line and draws no figures, report, attendance, Your night or error', () => {
    expect(text(body)).toContain('Logged as a raid night.');
    expect(body).not.toContain('raid-report-figures');
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
