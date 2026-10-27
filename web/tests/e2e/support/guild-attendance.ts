// web/tests/e2e/support/guild-attendance.ts
// Stubs for attendance and the raid-night report (design/specs/2026-10-27-guild-attendance.md §9),
// shared by tests/e2e/guild-attendance.spec.ts and the captures in tests/captures/guild-attendance/.
// It sits on top of guild-raids.ts's RaidStub (the raid, its sign-ups, /v1/me, /home) and adds the
// G2 routes: the raid view with its night shaped for the viewer (R-A1), match candidates, set and
// clear the match, Excused, Post to Discord (counting the `report_ready` posts it would send), the
// readiness board with `raid=`, and the roster's attendance source. Routes registered here run
// before RaidStub's, and hand anything they do not own back to it.
import type { Page, Route } from '@playwright/test';
import {
  GUILD,
  GUILD_ID,
  buildMockHome,
  buildMockReadiness,
  buildMockRoster,
} from '../../../src/fixtures/guild/mock-guild';
import type { GuildRaidView, GuildReadinessPage } from '../../../src/lib/guild/api';
import type {
  AttendanceSource,
  MatchCandidates,
  RaidAttendanceRow,
  RaidReport,
} from '../../../src/lib/guild/attendance-contract';
import {
  MATCHED_REPORT,
  allAttendance,
  fixtureCandidates,
  fixtureMatch,
  fixtureReport,
  nightOf,
  rosterAttendance,
} from './guild-attendance-data';
import {
  MEMBER_ALT,
  ONYXIA_ID,
  RaidStub,
  keyOf,
  stubGuildRaids,
  type RaidStubOptions,
  type RecordedWrite,
} from './guild-raids';
import { readinessFails } from '../../../src/lib/guild/readiness-view';
import { envelope } from './officer-tools';

export * from './guild-attendance-data';

/** A log matched to a raid: what the API keeps for its night. */
export interface StoredNight {
  reportId: string;
  how: 'auto' | 'hand';
  attendance: RaidAttendanceRow[];
  report: RaidReport;
}

export function storedNight(extra: Partial<StoredNight> = {}): StoredNight {
  return {
    reportId: MATCHED_REPORT,
    how: 'auto',
    attendance: allAttendance(),
    report: fixtureReport(),
    ...extra,
  };
}

interface Refusal {
  code: string;
  status: number;
}

const REFUSAL_LINES: Record<string, string> = {
  report_matched: 'That log is matched to another raid.',
  guild_plan_required: 'Running raids is part of the guild plan.',
  already_posted: 'This report was already posted.',
  post_failed: 'The post reached no channel.',
};

export interface AttendanceStubOptions extends RaidStubOptions {
  /** The matched night per raid id; a raid without one has no log yet. */
  nights?: Record<number, StoredNight>;
  candidates?: MatchCandidates;
  /** The roster's `attendance_source`; left out, the field is absent (report mode). */
  attendanceSource?: AttendanceSource;
  /** Holds each Excused write this long before it answers, so two ticks overlap. */
  excuseDelayMs?: number;
}

/** The ids of the readiness board's stale and missing exports on the next-raid filter. */
export const STALE_EXPORT = 'Windtalon';
export const NO_EXPORT = 'Shadowquill';

type ReadinessRow = GuildReadinessPage['rows'][number];

/** The row's failing count, from the checks its cells show (readinessFails). */
function withShownFailing(row: ReadinessRow): ReadinessRow {
  return { ...row, failing: readinessFails(row).length };
}

/** A signed character with no export: the API sends no gear, enchants or item level for it. */
function noExport(row: ReadinessRow): ReadinessRow {
  return {
    ...row,
    failing: 0,
    no_export: true,
    synced_before_cutoff: null,
    gear_gap: null,
    enchants: { missing_slots: [], checked: false },
    item_level: null,
    item_level_delta: null,
  };
}

export class AttendanceStub {
  readonly raids: RaidStub;
  nights: Map<number, StoredNight>;
  candidates: MatchCandidates;
  attendanceSource: AttendanceSource | undefined;
  excuseDelayMs: number;
  /** Every `report_ready` the API would post to Discord. */
  reportReady = 0;
  readonly readinessReads: (string | null)[] = [];
  refuseNext: { match?: Refusal; clear?: Refusal; excuse?: Refusal; post?: Refusal } = {};

  constructor(raids: RaidStub, options: AttendanceStubOptions) {
    this.raids = raids;
    this.nights = new Map(Object.entries(options.nights ?? {}).map(([id, n]) => [Number(id), n]));
    this.candidates = options.candidates ?? fixtureCandidates();
    this.attendanceSource = options.attendanceSource;
    this.excuseDelayMs = options.excuseDelayMs ?? 0;
    for (const id of this.nights.keys()) this.markDone(id);
  }

  get writes(): RecordedWrite[] {
    return this.raids.writes;
  }

  private markDone(id: number): void {
    this.raids.update(id, (raid) => ({ ...raid, status: 'done' }));
  }

  /** The poll test: the parse matches the log while the page is open. */
  matchFromElsewhere(id: number, night: StoredNight = storedNight()): void {
    this.nights.set(id, night);
    this.markDone(id);
  }

  /** The raid view as this viewer reads it: officers every row, a member their own (R-A1). */
  view(id: number): GuildRaidView {
    const base = this.raids.view(id);
    const night = this.nights.get(id);
    if (night === undefined) return base;
    const own = new Set(this.raids.ownKeys);
    const rows = night.attendance.map((row) => ({ ...row, own: own.has(row.character_key) }));
    const officer = this.raids.viewer === 'officer';
    const mine = <T extends { character_key: string }>(list: T[]): T[] =>
      officer ? list : list.filter((row) => own.has(row.character_key));
    const { report } = night;
    return {
      ...base,
      status: 'done',
      pre_done_status: 'planned',
      match: fixtureMatch({ report_id: night.reportId, how: night.how }),
      night: nightOf(night.attendance, report),
      attendance: mine(rows),
      report: {
        ...report,
        deaths: mine(report.deaths),
        consumables: { ...report.consumables, rows: mine(report.consumables.rows) },
        moves: { ...report.moves, rows: mine(report.moves.rows) },
      },
    };
  }

  takeRefusal(kind: keyof AttendanceStub['refuseNext']): Refusal | undefined {
    const found = this.refuseNext[kind];
    this.refuseNext = { ...this.refuseNext, [kind]: undefined };
    return found;
  }

  setMatch(id: number, reportId: string): GuildRaidView {
    this.nights.set(id, storedNight({ reportId, how: 'hand' }));
    this.markDone(id);
    return this.view(id);
  }

  clearMatch(id: number): GuildRaidView {
    this.nights.delete(id);
    this.raids.update(id, (raid) => ({ ...raid, status: 'planned' }));
    return this.view(id);
  }

  setExcused(id: number, key: string, excused: boolean): GuildRaidView {
    const night = this.nights.get(id);
    if (night === undefined) throw new Error(`guild-attendance stub: raid ${id} has no night`);
    const attendance = night.attendance.map((row) => (row.character_key === key ? { ...row, excused } : row));
    this.nights.set(id, { ...night, attendance });
    return this.view(id);
  }

  /** POST …/report/post: once per raid (R-A6); the second press is 409 `already_posted`. */
  post(id: number): { posted_at: string } | null {
    const night = this.nights.get(id);
    if (night === undefined || night.report.posted_at !== null) return null;
    const postedAt = '2026-12-11T06:05:00Z';
    this.nights.set(id, { ...night, report: { ...night.report, posted_at: postedAt } });
    this.reportReady += 1;
    return { posted_at: postedAt };
  }

  /** The readiness board: the guild's with `next_raid`, or one raid's In and Late sign-ups. */
  readiness(raidParam: string | null): GuildReadinessPage {
    const seeded = buildMockReadiness();
    // The failing count is what the board's own cells show (the sort's key), never a separate seed.
    const board = { ...seeded, rows: seeded.rows.map((row) => withShownFailing(row)) };
    const next = this.raids.nextRaid();
    if (raidParam === null) {
      return {
        ...board,
        next_raid:
          next === undefined
            ? null
            : { id: next.id, title: next.title, zone_name: next.zone_name, starts_at: next.starts_at },
      };
    }
    const raid = this.raids.raid(Number(raidParam));
    const signed = raid.signups.filter((s) => s.state === 'in' || s.state === 'late');
    const cutoff = new Date(Date.parse(raid.starts_at) - 7 * 86_400_000).toISOString();
    const rows = signed.map((s) => {
      const found = board.rows.find((r) => r.character_key === s.character_key);
      if (found === undefined) throw new Error(`guild-attendance stub: no readiness row for ${s.name}`);
      if (s.name === NO_EXPORT) return noExport(found);
      if (s.name === STALE_EXPORT) {
        // A stale export that still names a real upgrade (the board never reads `+0 DPS`).
        const gap = found.gear_gap === null ? null : { ...found.gear_gap, gain_dps: 14 };
        const stale = {
          ...found,
          gear_gap: gap,
          logged_at: '2026-12-01T18:00:00Z',
          synced_before_cutoff: true,
        };
        // A sync before the raid's cutoff is one more failing check (spec 3.7).
        return { ...stale, failing: stale.failing + 1 };
      }
      return { ...found, synced_before_cutoff: false };
    });
    return {
      ...board,
      rows,
      raid: { id: raid.id, title: raid.title, starts_at: raid.starts_at, cutoff_at: cutoff },
    };
  }

  /** `/home` with the roster's attendance source: raid mode adds the 30-raid window per row. */
  home() {
    const viewer = this.raids.viewer;
    const name = this.raids.verifiedMember ? (viewer === 'officer' ? 'Kraggor' : 'Zulmara') : null;
    const home = buildMockHome(name, GUILD.faction, null, true, {
      officerTools: this.raids.officerTools,
      nextRaid: this.raids.nextRaid(),
    });
    if (this.attendanceSource !== 'raids') {
      return this.attendanceSource === undefined ? home : { ...home, attendance_source: 'reports' as const };
    }
    return {
      ...home,
      attendance_source: 'raids' as const,
      // One attendance read for the roster and the guild line (WP 4).
      roster: home.roster.map((row) => {
        if (row.attendance === undefined) return row;
        const { last8, last30 } = rosterAttendance(row.name);
        return { ...row, attendance: last8, attendance_30: last30 };
      }),
    };
  }
}

function refused(found: Refusal) {
  return {
    status: found.status,
    contentType: 'application/json',
    body: JSON.stringify({
      ok: false,
      data: null,
      error: { code: found.code, message: REFUSAL_LINES[found.code] ?? 'Refused.' },
      request_id: 'req-test',
    }),
  };
}

const G2_PATH = new RegExp(`/v1/guilds/${GUILD_ID}/raids/(\\d+)(?:/(.+))?$`);
const wait = (ms: number): Promise<void> => new Promise((resolve) => setTimeout(resolve, ms));

async function answerG2(stub: AttendanceStub, route: Route): Promise<void> {
  const request = route.request();
  const url = new URL(request.url());
  const found = G2_PATH.exec(url.pathname);
  if (found === null) return route.fallback();
  const id = Number(found[1]);
  const rest = found[2];
  const method = request.method();
  const record = (): void => {
    stub.writes.push({ method, path: url.pathname, body: method === 'GET' ? null : request.postDataJSON() });
  };
  if (method === 'GET' && rest === undefined) {
    if (!stub.raids.verifiedMember) return route.fallback();
    stub.raids.raidReads.push(id);
    return route.fulfill(envelope(stub.view(id)));
  }
  if (method === 'GET' && rest === 'match-candidates') return route.fulfill(envelope(stub.candidates));
  if (rest === 'match' && (method === 'PUT' || method === 'DELETE')) {
    record();
    const refusal = stub.takeRefusal(method === 'PUT' ? 'match' : 'clear');
    if (refusal !== undefined) return route.fulfill(refused(refusal));
    const body = method === 'PUT' ? (request.postDataJSON() as { report_id: string }) : null;
    return route.fulfill(envelope(body === null ? stub.clearMatch(id) : stub.setMatch(id, body.report_id)));
  }
  if (method === 'PUT' && rest?.startsWith('attendance/') === true) {
    record();
    await wait(stub.excuseDelayMs);
    const refusal = stub.takeRefusal('excuse');
    if (refusal !== undefined) return route.fulfill(refused(refusal));
    const key = rest.slice('attendance/'.length).split('/').map(decodeURIComponent).join('/');
    const { excused } = request.postDataJSON() as { excused: boolean };
    return route.fulfill(envelope(stub.setExcused(id, key, excused)));
  }
  if (method === 'POST' && rest === 'report/post') {
    record();
    const refusal = stub.takeRefusal('post');
    if (refusal !== undefined) return route.fulfill(refused(refusal));
    const posted = stub.post(id);
    if (posted === null) return route.fulfill(refused({ code: 'already_posted', status: 409 }));
    return route.fulfill(envelope(posted));
  }
  return route.fallback();
}

/** stubGuildRaids plus the G2 routes; the raid is guild-raids.ts's Onyxia's Lair (id 7). */
export async function stubGuildAttendance(
  page: Page,
  options: AttendanceStubOptions,
): Promise<AttendanceStub> {
  const raids = await stubGuildRaids(page, options);
  const stub = new AttendanceStub(raids, options);
  await page.route(`**/v1/guilds/${GUILD_ID}/home`, (route) => route.fulfill(envelope(stub.home())));
  await page.route(new RegExp(`/v1/guilds/${GUILD_ID}/readiness(?:\\?.*)?$`), (route) => {
    const raid = new URL(route.request().url()).searchParams.get('raid');
    stub.readinessReads.push(raid);
    return route.fulfill(envelope(stub.readiness(raid)));
  });
  await page.route(G2_PATH, (route) => answerG2(stub, route));
  return stub;
}

/** The fixture night on the Onyxia raid, as most tests and captures start. */
export const DONE: Record<number, StoredNight> = { [ONYXIA_ID]: storedNight() };

/** The member's no-show variant: Duskstrider signed In and is not in the log. */
export function memberNoShow(): Record<number, StoredNight> {
  const report = fixtureReport();
  const alt = keyOf(MEMBER_ALT);
  const away = <T extends { character_key: string }>(rows: T[]): T[] =>
    rows.filter((r) => r.character_key !== alt);
  return {
    [ONYXIA_ID]: storedNight({
      attendance: allAttendance({ [MEMBER_ALT]: { outcome: 'no_show', pulls: 0 } }),
      // Not in the log: no deaths, consumables or moves of the alt's (the API reads them from it).
      report: {
        ...report,
        deaths: away(report.deaths),
        consumables: { ...report.consumables, rows: away(report.consumables.rows) },
        moves: { ...report.moves, rows: away(report.moves.rows) },
      },
    }),
  };
}

/** The roster rows' keys of the readiness filter: the raid's In and Late sign-ups. */
export function signedKeys(stub: AttendanceStub, id = ONYXIA_ID): string[] {
  return stub.raids
    .raid(id)
    .signups.filter((s) => s.state === 'in' || s.state === 'late')
    .map((s) => s.character_key);
}

export { buildMockRoster, keyOf };
