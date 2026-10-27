// web/tests/e2e/support/guild-attendance-data.ts
// The attendance lane's fixture night (design/specs/2026-10-27-guild-attendance.md §9): the
// Onyxia's Lair raid of Thu 10 Dec from guild-raids.ts, logged as 3 wipes then a kill. Every row
// comes from that raid's sign-ups (ONYXIA_SIGNUPS) and the seeded mock roster, so the attendance,
// the report and the night line agree with each other. Shared by guild-attendance.spec.ts and the
// captures in tests/captures/guild-attendance/.
import { buildMockRoster } from '../../../src/fixtures/guild/mock-guild';
import type {
  MatchCandidates,
  RaidAttendanceRow,
  RaidMatch,
  RaidNight,
  RaidReport,
} from '../../../src/lib/guild/attendance-contract';
import { MEMBER_ALT, ONYXIA_SIGNUPS, keyOf } from './guild-raids';

/** The raid ended at 11:00 pm Chicago (05:00 UTC); one hour later the matcher's window is open. */
export const AFTER_END = new Date('2026-12-11T06:00:00Z');
/** Past end + 6 h: the page asks for an upload instead. */
export const AFTER_WINDOW = new Date('2026-12-11T12:00:00Z');
export const MATCHED_REPORT = 'abc123def';
export const SECOND_REPORT = 'zz4400kk';
export const TAKEN_REPORT = 'tt9911rr';
export const PREVIOUS_RAID = { id: 6, title: "Onyxia's Lair", starts_at: '2026-12-04T02:00:00Z' };

const ROSTER = new Map(buildMockRoster().map((row) => [row.name, row]));
const PULLS = 4;
const SIGNED = new Map(ONYXIA_SIGNUPS.map((s) => [s.name, s.state]));

/** name: how the log saw them. Absent names signed In or Late read `no_show`. */
interface Seen {
  outcome: RaidAttendanceRow['outcome'];
  pulls: number;
  offsetS?: number;
  cameAs?: string;
}

const SEEN: Record<string, Seen> = {
  Kraggor: { outcome: 'present', pulls: PULLS },
  Lightbrand: { outcome: 'present', pulls: PULLS },
  Mendwhisper: { outcome: 'late', pulls: 3, offsetS: 1320 },
  Earthhoof: { outcome: 'present', pulls: PULLS },
  'Obnoxious Yell': { outcome: 'present', pulls: PULLS },
  Sunderfel: { outcome: 'present', pulls: PULLS },
  Windtalon: { outcome: 'present', pulls: PULLS },
  Felsnap: { outcome: 'present', pulls: PULLS, cameAs: 'Mirestalk' },
  Vexlash: { outcome: 'present', pulls: PULLS },
  Shadowquill: { outcome: 'present', pulls: PULLS },
  Hexbramble: { outcome: 'no_show', pulls: 0 },
  Pyrewisp: { outcome: 'present', pulls: PULLS },
  [MEMBER_ALT]: { outcome: 'present', pulls: PULLS },
  // Signed Late and came 9 min in, as she said: present (WP 9 ruling).
  Arcanemoor: { outcome: 'present', pulls: PULLS, offsetS: 540 },
  Soulgrave: { outcome: 'present', pulls: PULLS },
  Frostnettle: { outcome: 'absent', pulls: 0 },
  Thornhide: { outcome: 'present', pulls: 2 },
};

function rosterRow(name: string) {
  const row = ROSTER.get(name);
  if (row === undefined) throw new Error(`guild-attendance data: no roster row "${name}"`);
  return row;
}

function attendanceRow(name: string, seen: Seen): RaidAttendanceRow {
  const row = rosterRow(name);
  const signed = SIGNED.get(name) ?? null;
  return {
    character_key: row.character_key,
    name,
    class: row.class ?? '',
    spec: row.spec ?? '',
    signed,
    role: signed === null ? null : (row.role ?? 'dps'),
    outcome: seen.outcome,
    came_as: seen.cameAs ?? null,
    first_seen_offset_s: seen.outcome === 'absent' || seen.outcome === 'no_show' ? null : (seen.offsetS ?? 0),
    pulls_present: seen.pulls,
    pulls_total: PULLS,
    excused: false,
    own: false,
  };
}

/** A verified raider who has signed up for and attended no logged raid yet: `No raids yet`. */
export const NEW_RECRUIT = 'Grimvow';

/** The roster attendance the API reads once for both the roster chip and the account's guild line
 *  (WP 4): raids attended over raids counted, last 8 and last 30. One map, so the two agree. */
interface Window {
  present: number;
  nights: number;
}

export function rosterAttendance(name: string): { last8: Window; last30: Window } {
  if (name === NEW_RECRUIT) return { last8: { present: 0, nights: 0 }, last30: { present: 0, nights: 0 } };
  const seed = [...name].reduce((sum, ch) => sum + ch.charCodeAt(0), 0);
  const raided = name in SEEN && SEEN[name].outcome !== 'absent' && SEEN[name].outcome !== 'no_show';
  const present8 = raided ? 6 + (seed % 3) : 3 + (seed % 3);
  return {
    last8: { present: Math.min(present8, 8), nights: 8 },
    last30: { present: 18 + (seed % 10), nights: 30 },
  };
}

/** Every attendance row, as an officer reads them. */
export function allAttendance(change: Partial<Record<string, Partial<Seen>>> = {}): RaidAttendanceRow[] {
  return Object.entries(SEEN).map(([name, seen]) => attendanceRow(name, { ...seen, ...change[name] }));
}

/** The night's figures, counted from the rows (never typed apart from them). */
export function nightOf(rows: readonly RaidAttendanceRow[], report: RaidReport): RaidNight {
  const attending = rows.filter((r) => r.signed === 'in' || r.signed === 'late');
  const came = attending.filter((r) => r.outcome === 'present' || r.outcome === 'late');
  return {
    signed: attending.length,
    came: came.length,
    late: attending.filter((r) => r.outcome === 'late').length,
    not_in_log: attending.filter((r) => r.outcome === 'no_show').length,
    joined: rows.filter((r) => !attending.includes(r) && r.outcome !== 'absent').length,
    excused: rows.filter((r) => r.excused).length,
    kills: report.kill_order.filter((k) => k.killed).length,
    wipes: report.kill_order.reduce((sum, k) => sum + k.wipes_before, 0),
    deaths: report.deaths.reduce((sum, d) => sum + d.on_wipes + d.on_kills, 0),
    duration_ms: 9_660_000,
  };
}

export function fixtureMatch(extra: Partial<RaidMatch> = {}): RaidMatch {
  return {
    report_id: MATCHED_REPORT,
    zone_name: "Onyxia's Lair",
    uploaded_at: '2026-12-11T05:42:00Z',
    uploaded_by_name: 'Kraggor',
    pulls: PULLS,
    kills: 1,
    duration_ms: 9_660_000,
    how: 'auto',
    ...extra,
  };
}

const ref = (name: string) => {
  const row = rosterRow(name);
  return { character_key: row.character_key, name, class: row.class ?? '' };
};

/** Eleven raiders died, so Deaths shows `Show all 11 raiders` (spec 7: 10 before Show all). */
const DEATHS: [string, number, number][] = [
  ['Pyrewisp', 3, 1],
  ['Windtalon', 3, 0],
  ['Vexlash', 2, 0],
  ['Kraggor', 1, 0],
  ['Arcanemoor', 1, 0],
  ['Sunderfel', 1, 0],
  ['Lightbrand', 1, 0],
  ['Earthhoof', 1, 0],
  ['Shadowquill', 1, 0],
  ['Felsnap', 1, 0],
  ['Mendwhisper', 1, 0],
];

/** id, name, icon (the API's loot table: Head of Onyxia, a quest item, has none), awardee. */
const LOOT: [number, string, string, string][] = [
  [18423, 'Head of Onyxia', '', 'Kraggor'],
  [17075, "Vis'kag the Bloodletter", 'inv_sword_18', 'Vexlash'],
  [16914, 'Netherwind Crown', 'inv_helmet_70', 'Pyrewisp'],
  [18813, 'Ring of Binding', 'inv_jewelry_ring_13', 'Thornhide'],
];

const MOVES: [string, 'dps' | 'hps', number, number][] = [
  ['Pyrewisp', 'dps', 238, 291],
  ['Lightbrand', 'hps', 402, 443],
  ['Windtalon', 'dps', 251, 219],
  ['Vexlash', 'dps', 270, 289],
  [MEMBER_ALT, 'dps', 180, 196],
];

/** Every raider in the log (the API lists no one else): Before, flask, food, potions. */
type Consumed = ['stocked' | 'short' | 'unknown', boolean, boolean, number];
const CONSUMED: Record<string, Consumed> = {
  Windtalon: ['short', false, true, 1],
  Pyrewisp: ['short', true, true, 0],
  Felsnap: ['short', false, true, 2],
  Kraggor: ['stocked', true, true, 3],
  Lightbrand: ['stocked', true, true, 2],
  Mendwhisper: ['stocked', true, true, 1],
  Vexlash: ['stocked', true, true, 2],
  Sunderfel: ['stocked', true, false, 2],
  [MEMBER_ALT]: ['unknown', false, true, 2],
  Earthhoof: ['unknown', true, true, 1],
};
const DEFAULT_CONSUMED: Consumed = ['stocked', true, true, 1];

function consumableRows() {
  return Object.entries(SEEN)
    .filter(([, seen]) => seen.outcome === 'present' || seen.outcome === 'late')
    .map(([name]) => {
      const [before, flask, food, potions] = CONSUMED[name] ?? DEFAULT_CONSUMED;
      return { ...ref(name), before, flask, food, potions };
    });
}

export function fixtureReport(extra: Partial<RaidReport> = {}): RaidReport {
  return {
    kill_order: [
      {
        encounter_id: 1084,
        name: 'Onyxia',
        fight_index: 4,
        killed: true,
        kill_ms: 252_000,
        wipes_before: 3,
        deaths: 2,
      },
    ],
    unnamed: null,
    deaths: DEATHS.map(([name, w, k]) => ({ ...ref(name), on_wipes: w, on_kills: k })),
    consumables: {
      log_has_buffs: true,
      rows: consumableRows(),
    },
    loot: [
      {
        encounter_id: 1084,
        boss: 'Onyxia',
        awards: LOOT.map(([item_id, item_name, icon, name]) => ({
          item_id,
          item_name,
          quality: 4,
          icon,
          ...ref(name),
          by_name: 'Kraggor',
          at: '2026-12-11T04:30:00Z',
        })),
      },
    ],
    moves: {
      previous: PREVIOUS_RAID,
      rows: MOVES.map(([name, metric, before, after]) => ({
        ...ref(name),
        metric,
        encounter: 'Onyxia',
        before,
        after,
      })),
    },
    posted_at: null,
    ...extra,
  };
}

/** §9 capture 7: no buff data, no previous raid, no loot. */
export function bareReport(): RaidReport {
  const full = fixtureReport();
  return {
    ...full,
    consumables: { log_has_buffs: false, rows: full.consumables.rows },
    loot: [],
    moves: { previous: null, rows: [] },
  };
}

export function fixtureCandidates(extra: Partial<MatchCandidates> = {}): MatchCandidates {
  return {
    zone_known: true,
    rows: [
      {
        report_id: MATCHED_REPORT,
        zone_name: "Onyxia's Lair",
        uploaded_at: '2026-12-11T05:42:00Z',
        uploaded_by_name: 'Kraggor',
        pulls: PULLS,
        kills: 1,
        duration_ms: 9_660_000,
        closest: true,
        matched_raid: null,
      },
      {
        report_id: SECOND_REPORT,
        zone_name: "Onyxia's Lair",
        uploaded_at: '2026-12-11T06:20:00Z',
        uploaded_by_name: 'Lightbrand',
        pulls: 2,
        kills: 0,
        duration_ms: 1_440_000,
        closest: false,
        matched_raid: null,
      },
      {
        report_id: TAKEN_REPORT,
        zone_name: "Onyxia's Lair",
        uploaded_at: '2026-12-04T05:05:00Z',
        uploaded_by_name: 'Kraggor',
        pulls: 6,
        kills: 1,
        duration_ms: 6_000_000,
        closest: false,
        matched_raid: PREVIOUS_RAID,
      },
    ],
    ...extra,
  };
}

/** The copied report for the fixture night (spec 4 "Copy for Discord"), typed out in full so the
 *  e2e compares the clipboard with the spec's lines rather than with the function that writes them. */
export function expectedDiscordText(reportUrl: string): string {
  return [
    "Onyxia's Lair · Thu 10 Dec · Olympus XXVII",
    '1 kill, 3 wipes, 2 h 41 min. 13 of 14 came.',
    'Kills: Onyxia 4:12 (after 3 wipes)',
    'Most deaths: Pyrewisp 4, Windtalon 3, Vexlash 2',
    "Loot: Onyxia: Head of Onyxia to Kraggor; Vis'kag the Bloodletter to Vexlash; Netherwind Crown to Pyrewisp; Ring of Binding to Thornhide",
    `Biggest changes: Pyrewisp DPS +53, Lightbrand HPS +41, Vexlash DPS +19, ${MEMBER_ALT} DPS +16`,
    `Report: ${reportUrl}`,
  ].join('\n');
}

export { keyOf };
