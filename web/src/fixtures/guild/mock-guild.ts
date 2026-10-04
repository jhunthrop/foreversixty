// web/src/fixtures/guild/mock-guild.ts
// The guild control-centre's one shared mock guild: a direct TypeScript port of
// design/mocks/gen_guild.py's own roster, raid nights and seeded derived-stat functions,
// so this fixture, the unit/e2e tests that stub the API with it, and the acceptance boards
// gen_guild.py renders can never quietly drift apart (gen_guild.py's own module docstring:
// "no two boards can disagree"). `seed()` uses the exact same md5-prefix algorithm as the
// Python file, so every derived number (rating, attendance, parses, gear gap, ...) is
// byte-identical between the two.
//
// Node-only (`node:crypto`): this module is never imported by a page component or bundled
// for the browser -- only by vitest unit tests and Playwright spec files, both of which run
// in Node. A Svelte/Astro component that needs this guild's shape imports the API types
// from `../../lib/guild/api` and gets its data from a real fetch (or, in a test, from
// `page.route`'s response built with this fixture).
import { createHash } from 'node:crypto';
import type {
  GuildHome,
  GuildLootPage,
  GuildProgressionPage,
  GuildReadinessPage,
  GuildRaidsPage,
  GuildRosterRow,
} from '../../lib/guild/api';

export const GUILD_ID = 2024;
export const GUILD = { id: GUILD_ID, name: 'Olympus XXVII', region: 'us', ruleset: 'pvp' } as const;

type MockRosterSeed = [
  name: string,
  cls: string,
  spec: string,
  rank: 'leader' | 'officer' | 'member',
  verified: boolean,
  logged: boolean,
  consent: 'roster' | 'gear' | 'gear_bags',
  ilvl: number | null,
];

// gen_guild.py's own ROSTER, verbatim.
const ROSTER_SEED: MockRosterSeed[] = [
  ['Kraggor', 'warrior', 'Protection', 'leader', true, true, 'gear_bags', 68],
  ['Obnoxious Yell', 'warrior', 'Fury', 'officer', true, true, 'gear_bags', 66],
  ['Sunderfel', 'warrior', 'Arms', 'officer', true, true, 'gear', 61],
  ['Grimtotem', 'warrior', 'Fury', 'member', true, false, 'gear', 58],
  ['Zulmara', 'hunter', 'Marksmanship', 'member', true, true, 'gear_bags', 64],
  ['Windtalon', 'hunter', 'Marksmanship', 'member', true, true, 'gear_bags', 70],
  ['Duskstrider', 'hunter', 'Marksmanship', 'member', true, false, 'gear', 55],
  ['Felsnap', 'hunter', 'Beast Mastery', 'member', true, true, 'gear', 60],
  ['Vexlash', 'rogue', 'Combat', 'officer', true, true, 'gear_bags', 65],
  ['Shadowquill', 'rogue', 'Assassination', 'member', true, true, 'gear', 59],
  ['Nixthrottle', 'rogue', 'Combat', 'member', true, false, 'roster', null],
  ['Hexbramble', 'warlock', 'Affliction', 'member', true, true, 'gear', 63],
  ['Soulgrave', 'warlock', 'Destruction', 'member', true, true, 'gear', 57],
  ['Pyrewisp', 'mage', 'Fire', 'member', true, true, 'gear_bags', 69],
  ['Frostnettle', 'mage', 'Frost', 'member', true, false, 'gear', 56],
  ['Arcanemoor', 'mage', 'Arcane', 'member', true, true, 'gear', 62],
  ['Lightbrand', 'priest', 'Holy', 'member', true, true, 'gear_bags', 67],
  ['Grimvow', 'priest', 'Shadow', 'member', true, true, 'roster', null],
  ['Mendwhisper', 'priest', 'Holy', 'member', true, false, 'gear', 60],
  ['Earthhoof', 'druid', 'Restoration', 'member', true, true, 'gear_bags', 65],
  ['Thornhide', 'druid', 'Feral', 'member', true, true, 'gear', 58],
  ['Rootgall', 'shaman', 'Restoration', 'member', false, true, 'gear', 61],
  ['Stormtusk', 'shaman', 'Elemental', 'member', false, false, 'gear', 55],
  ['Fulmintide', 'shaman', 'Enhancement', 'member', false, true, 'gear', 59],
];

const SPEC_ROLE: Record<string, 'tank' | 'healer' | 'dps'> = {
  Protection: 'tank',
  Fury: 'dps',
  Arms: 'dps',
  Marksmanship: 'dps',
  'Beast Mastery': 'dps',
  Combat: 'dps',
  Assassination: 'dps',
  Affliction: 'dps',
  Destruction: 'dps',
  Fire: 'dps',
  Frost: 'dps',
  Arcane: 'dps',
  Holy: 'healer',
  Shadow: 'dps',
  Restoration: 'healer',
  Feral: 'dps',
  Elemental: 'dps',
  Enhancement: 'dps',
};

const ALT_OF: Record<string, string> = { Duskstrider: 'Zulmara', Grimtotem: 'Kraggor' };

const PROFESSIONS_BY_CLASS: Record<string, [string, string]> = {
  warrior: ['Blacksmithing', 'Mining'],
  hunter: ['Leatherworking', 'Skinning'],
  rogue: ['Engineering', 'Mining'],
  warlock: ['Tailoring', 'Enchanting'],
  mage: ['Tailoring', 'Enchanting'],
  priest: ['Tailoring', 'Enchanting'],
  druid: ['Herbalism', 'Alchemy'],
  shaman: ['Leatherworking', 'Enchanting'],
};

export const VIEWER_MEMBER = 'Zulmara';
export const VIEWER_OFFICER = 'Kraggor';

// gen_guild.py's own NIGHTS, verbatim -- the real first tier, all nights after 9 Dec 2026.
const NIGHTS: [zone: string, date: string, durationMin: number, pulls: number, kills: number][] = [
  ['Barrow Deeps', '2026-12-10', 118, 6, 0],
  ['Hyjal Summit', '2026-12-11', 96, 5, 0],
  ['Barrow Deeps', '2026-12-13', 134, 7, 0],
  ["Onyxia's Lair", '2026-12-15', 41, 3, 1],
  ['Hyjal Summit', '2026-12-16', 102, 6, 0],
  ['Barrow Deeps', '2026-12-18', 151, 8, 0],
  ["Onyxia's Lair", '2026-12-20', 22, 2, 1],
  ['Hyjal Summit', '2026-12-22', 109, 7, 0],
];

const ONYXIA_PULLS: [date: string, pull: number, durationSec: number, result: 'kill' | 'wipe'][] = [
  ['2026-12-15', 1, 95, 'wipe'],
  ['2026-12-15', 2, 210, 'wipe'],
  ['2026-12-15', 3, 268, 'kill'],
  ['2026-12-20', 1, 240, 'wipe'],
  ['2026-12-20', 2, 251, 'kill'],
];

const ONYXIA_DROPS: [item: string, slot: string, classes: string[]][] = [
  ['Helm of Wrath', 'head', ['warrior']],
  ["Dragonstalker's Helm", 'head', ['hunter']],
  ['Netherwind Crown', 'head', ['mage']],
  ['Halo of Transcendence', 'head', ['priest']],
  ['Nemesis Skullcap', 'head', ['warlock']],
  ['Stormrage Cover', 'head', ['druid']],
  ["Eskhandar's Collar", 'neck', ['warrior', 'hunter', 'rogue']],
  ["Vis'kag the Bloodletter", 'one_hand', ['rogue', 'warrior']],
];

const ONYXIA_ENCOUNTER_ID = 1084;

function seed(name: string, salt: string): number {
  const hash = createHash('md5').update(`${name}:${salt}`).digest('hex');
  return parseInt(hash.slice(0, 8), 16) / 0xffffffff;
}

function rowByName(name: string): MockRosterSeed {
  const found = ROSTER_SEED.find((r) => r[0] === name);
  if (found === undefined) throw new Error(`mock-guild: unknown raider "${name}"`);
  return found;
}

function roleFor(name: string): 'tank' | 'healer' | 'dps' {
  return SPEC_ROLE[rowByName(name)[2]];
}

const ILVLS = ROSTER_SEED.map((r) => r[7]).filter((v): v is number => v !== null);
const MEDIAN_ILVL = [...ILVLS].sort((a, b) => a - b)[Math.floor(ILVLS.length / 2)];

export function ratingFor(name: string) {
  const comps = ['Output', 'Survival', 'Mechanics', 'Utility', 'Preparation', 'Activity'].map(
    (c) => [c.toLowerCase(), Math.round(50 + seed(name, c) * 45)] as const,
  );
  const overall = Math.round(comps.reduce((sum, [, v]) => sum + v, 0) / comps.length);
  return { overall, ...Object.fromEntries(comps) } as {
    overall: number;
    output: number;
    survival: number;
    mechanics: number;
    utility: number;
    preparation: number;
    activity: number;
  };
}

export function attendanceFor(name: string): { present: number; nights: number } {
  return { present: 5 + Math.floor(seed(name, 'attendance') * 4), nights: 8 };
}

function parsesFor(name: string): { metric: string; best: number; avg: number } {
  const metric = roleFor(name) === 'healer' ? 'hps' : 'dps';
  const best = 230 + seed(name, 'best') * 180;
  const avg = best * (0.72 + seed(name, 'avg') * 0.18);
  return { metric, best, avg };
}

export function gearGapFor(name: string): { upgrades: number | null; gain: number | null } {
  const ilvl = rowByName(name)[7];
  if (ilvl === null) return { upgrades: null, gain: null };
  const deficit = Math.max(0, 70 - ilvl);
  const upgrades = Math.min(7, Math.round(deficit / 3) + Math.floor(seed(name, 'upgrades') * 2));
  const gain = Math.round(deficit * (2.0 + seed(name, 'gain') * 2.4) * 10) / 10;
  return { upgrades, gain };
}

const ENCHANT_SLOTS = ['weapon', 'chest', 'cloak', 'boots'];

export function missingEnchantsFor(name: string): string[] {
  const n = Math.floor(seed(name, 'enchant-count') * 3);
  const pool = [...ENCHANT_SLOTS].sort((a, b) => seed(name, `enchant-${a}`) - seed(name, `enchant-${b}`));
  return pool.slice(0, n);
}

export function consumablesOkFor(name: string): boolean {
  return seed(name, 'consumables') > 0.35;
}

export function unspentPointsFor(name: string): number {
  return Math.floor(seed(name, 'talent-points') * 3.2);
}

function professionsFor(name: string): [string, string] {
  return PROFESSIONS_BY_CLASS[rowByName(name)[1]];
}

/** `guild_characters.user_id` grouping (EXISTS): an alt shares its main's own `account_key`
 *  (`gen_guild.py`'s own `ALT_OF`), so the Roster tab's "alt of {main}" tag has a real
 *  account to group on rather than a fixture-only coincidence of name. */
function accountKeyFor(name: string): string {
  const mainName = ALT_OF[name] ?? name;
  return `u:${1000 + ROSTER_SEED.findIndex((r) => r[0] === mainName)}`;
}

function toRosterRow(name: string): GuildRosterRow {
  const [raiderName, cls, spec, rank, verified, logged, consent, ilvl] = rowByName(name);
  const parses = parsesFor(name);
  const hasGear = consent === 'gear' || consent === 'gear_bags';
  return {
    character_key: `us/pvp/${raiderName.toLowerCase().replace(/[^a-z0-9]+/g, '-')}`,
    region: 'us',
    ruleset: 'pvp',
    name: raiderName,
    class: cls,
    spec,
    role: roleFor(name),
    rank,
    verified,
    logged_recently: logged,
    logged_at: logged ? '2026-12-22T21:40:00Z' : '2026-12-18T19:00:00Z',
    item_level: ilvl ?? undefined,
    consent,
    attendance: attendanceFor(name),
    best_parse: {
      metric: parses.metric,
      value: Math.round(parses.best * 10) / 10,
      percentile: null,
      encounter: 'Onyxia',
      report_id: 'fixtureguildraid7',
      fight_index: 2,
    },
    rating: ratingFor(name),
    professions: professionsFor(name),
    account_key: accountKeyFor(name),
    last_report_at: logged ? '2026-12-22T21:40:00Z' : '2026-12-15T19:00:00Z',
    may_approve: !verified,
    may_remove: hasGear && rank !== 'leader',
  };
}

export function buildMockRoster(): GuildRosterRow[] {
  return ROSTER_SEED.map((r) => toRosterRow(r[0]));
}

export function buildMockHome(viewerName: string | null): GuildHome {
  const roster = buildMockRoster();
  const pending = roster.filter((r) => !r.verified);
  const verifiedCount = roster.filter((r) => r.verified).length;
  const belowFloor = roster.filter((r) => (r.rating?.overall ?? 100) < 55).length;
  const viewerRow = viewerName === null ? null : (roster.find((r) => r.name === viewerName) ?? null);
  const role: 'public' | 'member' | 'officer' =
    viewerName === null ? 'public' : viewerName === VIEWER_OFFICER ? 'officer' : 'member';

  const standing =
    viewerRow === null
      ? null
      : (() => {
          const peers = roster.filter(
            (r) => r.class === viewerRow.class && r.spec === viewerRow.spec && r.item_level !== undefined,
          );
          const ranked = [...peers].sort((a, b) => (b.item_level ?? 0) - (a.item_level ?? 0));
          const rank = ranked.findIndex((r) => r.character_key === viewerRow.character_key) + 1;
          const fails = readinessFailsFor(viewerRow.name);
          return {
            spec: viewerRow.spec ?? '',
            class: viewerRow.class ?? '',
            same_spec_count: peers.length,
            rank_by_item_level: rank,
            item_level: viewerRow.item_level ?? 0,
            needs_before_next_raid: fails.slice(0, 3),
          };
        })();

  return {
    guild: GUILD,
    viewer: {
      role,
      character_key: viewerRow?.character_key ?? null,
      verified: viewerRow?.verified ?? false,
    },
    claim: { state: 'claimed', since: '2026-10-01T00:00:00Z', frozen: false, claimed_by_name: 'Kraggor' },
    summary: {
      raiders: verifiedCount,
      waiting_for_approval: pending.length,
      below_rating_floor: belowFloor,
      named_encounters_down: 1,
      pulls_this_tier: NIGHTS.reduce((sum, n) => sum + n[3], 0),
      updated_at: '2026-12-22T21:40:00Z',
    },
    standing,
    reports: NIGHTS.slice(-3)
      .reverse()
      .map(([zone, date], i) => ({
        id: `fixtureguildraid${i}`,
        title: `${zone} · raid night`,
        zone,
        created_at: `${date}T19:00:00Z`,
        fight_count: 6,
        kill_count: zone === "Onyxia's Lair" ? 1 : 0,
      })),
    roster,
    pending,
  };
}

export function readinessFailsFor(name: string): string[] {
  const fails: string[] = [];
  const gap = gearGapFor(name);
  if (gap.upgrades !== null && gap.upgrades >= 3) {
    fails.push(`${gap.upgrades} gear upgrades waiting (${gap.gain} DPS)`);
  }
  const missing = missingEnchantsFor(name);
  if (missing.length > 0)
    fails.push(`no enchant: ${missing.map((s) => s[0].toUpperCase() + s.slice(1)).join(', ')}`);
  if (!consumablesOkFor(name)) fails.push('bags short on consumables');
  const pts = unspentPointsFor(name);
  if (pts > 0) fails.push(`${pts} unspent talent point${pts > 1 ? 's' : ''}`);
  return fails;
}

export function buildMockRaids(): GuildRaidsPage {
  const rows = NIGHTS.slice()
    .reverse()
    .map(([zone, date, durationMin, pulls, kills], i) => {
      const isOnyxia = zone === "Onyxia's Lair";
      const present = buildMockRoster()
        .filter((r) => r.verified)
        .slice(0, 20)
        .map((r) => ({ character_key: r.character_key, name: r.name, class: r.class ?? '' }));
      const fights = isOnyxia
        ? ONYXIA_PULLS.filter((p) => p[0] === date).map(([, , d, result], idx) => ({
            index: idx,
            name: 'Onyxia',
            encounter_id: ONYXIA_ENCOUNTER_ID,
            kill: result === 'kill',
            duration_ms: d * 1000,
            deaths: result === 'kill' ? 1 : 3,
            players: 20,
          }))
        : Array.from({ length: pulls }, (_, idx) => ({
            index: idx,
            name: null,
            encounter_id: null,
            kill: false,
            duration_ms: 90000,
            deaths: 2,
            players: 20,
          }));
      return {
        id: `fixtureguildraid${NIGHTS.length - 1 - i}`,
        title: `${zone} · raid night`,
        zone,
        created_at: `${date}T19:00:00Z`,
        duration_ms: durationMin * 60 * 1000,
        fight_count: pulls,
        kill_count: kills,
        wipe_count: pulls - kills,
        raiders: present.length,
        deaths: fights.reduce((sum, f) => sum + f.deaths, 0),
        top_parse: { name: 'Pyrewisp', class: 'mage', metric: 'dps', value: 358.2 },
        fights,
        present,
      };
    });
  return { rows, next_cursor: null };
}

export function buildMockProgression(): GuildProgressionPage {
  const barrowPulls = NIGHTS.filter((n) => n[0] === 'Barrow Deeps').reduce((sum, n) => sum + n[3], 0);
  const hyjalPulls = NIGHTS.filter((n) => n[0] === 'Hyjal Summit').reduce((sum, n) => sum + n[3], 0);
  const onyxiaPullsByNight = ['2026-12-15', '2026-12-20'].map((date, i) => ({
    report_id: `fixtureguildraid${i}`,
    date,
    pulls: ONYXIA_PULLS.filter((p) => p[0] === date).length,
    killed: ONYXIA_PULLS.some((p) => p[0] === date && p[3] === 'kill'),
  }));
  return {
    tier: {
      name: 'First tier',
      raids: ['Barrow Deeps', 'Hyjal Summit', "Onyxia's Lair"],
      named_encounters: 1,
      down: 1,
    },
    encounters: [
      {
        encounter_id: ONYXIA_ENCOUNTER_ID,
        name: 'Onyxia',
        zone: "Onyxia's Lair",
        first_kill_at: '2026-12-15T19:45:00Z',
        pulls: ONYXIA_PULLS.length,
        kills: ONYXIA_PULLS.filter((p) => p[3] === 'kill').length,
        best_kill_ms: 251 * 1000,
        pulls_by_night: onyxiaPullsByNight,
        deaths_per_pull: 1.4,
        best_by_role: {
          dps: {
            name: 'Pyrewisp',
            class: 'mage',
            spec: 'Fire',
            metric: 'dps',
            value: 358.2,
            report_id: 'fixtureguildraid0',
            fight_index: 1,
          },
          healer: {
            name: 'Lightbrand',
            class: 'priest',
            spec: 'Holy',
            metric: 'hps',
            value: 301.9,
            report_id: 'fixtureguildraid0',
            fight_index: 2,
          },
        },
      },
    ],
    unnamed: [
      {
        zone: 'Barrow Deeps',
        pulls: barrowPulls,
        nights: NIGHTS.filter((n) => n[0] === 'Barrow Deeps').length,
      },
      {
        zone: 'Hyjal Summit',
        pulls: hyjalPulls,
        nights: NIGHTS.filter((n) => n[0] === 'Hyjal Summit').length,
      },
    ],
  };
}

export function buildMockReadiness(): GuildReadinessPage {
  const rows = buildMockRoster()
    .filter((r) => r.verified)
    .map((r) => {
      const gap = gearGapFor(r.name);
      const hasGear = r.consent === 'gear' || r.consent === 'gear_bags';
      const hasBags = r.consent === 'gear_bags';
      const missing = hasGear ? missingEnchantsFor(r.name) : [];
      const fails = readinessFailsFor(r.name);
      return {
        character_key: r.character_key,
        name: r.name,
        class: r.class ?? '',
        spec: r.spec ?? '',
        consent: r.consent,
        gear_gap: hasGear ? { upgrades: gap.upgrades, gain_dps: gap.gain, not_sim_checked: 0 } : null,
        enchants: { missing_slots: missing, checked: hasGear },
        consumables: {
          state: hasBags
            ? consumablesOkFor(r.name)
              ? ('stocked' as const)
              : ('short' as const)
            : ('unknown' as const),
        },
        talent_points_unspent: unspentPointsFor(r.name),
        item_level: r.item_level ?? null,
        item_level_delta: r.item_level === undefined ? null : r.item_level - MEDIAN_ILVL,
        logged_at: r.logged_at ?? '2026-12-18T19:00:00Z',
        failing: fails.length,
        nudge_text: `${r.name}: ${fails.length === 0 ? 'every readiness check passes.' : `${fails.join(', ')} -- check before Thursday.`}`,
      };
    });
  return { median_item_level: MEDIAN_ILVL, generated_at: '2026-12-22T21:40:00Z', rows };
}

export function buildMockLoot(): GuildLootPage {
  const roster = buildMockRoster();
  const items = ONYXIA_DROPS.slice(0, 6).map(([name, slot, classes], i) => {
    const eligible = roster.filter(
      (r) => r.verified && classes.includes(r.class ?? '') && gearGapFor(r.name).gain !== null,
    );
    const candidates = [...eligible]
      .sort((a, b) => (gearGapFor(b.name).gain ?? 0) - (gearGapFor(a.name).gain ?? 0))
      .slice(0, 4)
      .map((r) => ({
        character_key: r.character_key,
        name: r.name,
        class: r.class ?? '',
        spec: r.spec ?? '',
        gain_dps: gearGapFor(r.name).gain ?? 0,
        not_sim_checked: false,
        attendance: attendanceFor(r.name),
        already_equivalent: gearGapFor(r.name).upgrades === 0,
      }));
    const isAwarded = i === 0 && candidates.length > 0;
    return {
      item_id: 16955 + i,
      name,
      icon: 'inv_helmet_71',
      quality: 4,
      slot,
      awarded_to: isAwarded
        ? {
            character_key: candidates[0].character_key,
            name: candidates[0].name,
            at: '2026-12-15T20:10:00Z',
            by_name: 'Kraggor',
          }
        : null,
      candidates,
    };
  });
  return {
    encounters: [{ encounter_id: ONYXIA_ENCOUNTER_ID, name: 'Onyxia', zone: "Onyxia's Lair", killed: true }],
    selected: ONYXIA_ENCOUNTER_ID,
    items,
  };
}
