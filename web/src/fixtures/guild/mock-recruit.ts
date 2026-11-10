// web/src/fixtures/guild/mock-recruit.ts
// Recruitment's API answers for tests and the mocked e2e (api/openapi.yaml GuildRecruitCard,
// GuildOfficerApplications, GuildMemberProfile), over the mock guild (mock-guild.ts) so the
// roster, account keys and officers agree. Test data only; never shipped.
import type {
  ApplicantCharacter,
  ApplicationDecision,
  GuildApplication,
  MemberNote,
  MemberProfile,
  OfficerApplication,
  OfficerApplications,
  RecruitCard,
} from '../../lib/guild/recruit-api';
import { GUILD, GUILD_ID, buildMockReadiness, buildMockRoster } from './mock-guild';

export const APPLICANT_KEY = 'us/pvp/ashenveil';
export const SECOND_APPLICANT_KEY = 'us/pvp/quillbright';
export const MEMBER_ACCOUNT = 'u:1001';

export function buildMockRecruitCard(open = true): RecruitCard {
  return {
    guild: {
      id: GUILD_ID,
      name: GUILD.name,
      region: GUILD.region,
      ruleset: GUILD.ruleset,
      faction: GUILD.faction,
      crest_url: null,
      path: '/guild/us/pvp/olympus-xxvii',
    },
    members: 24,
    open,
    pitch: 'Two nights a week, flasks on every pull, loot by council with the numbers shown.',
    raid_nights: {
      text: 'Tue and Thu, 8 to 11 pm server time',
      upcoming: open ? ['2026-12-10T02:00:00Z', '2026-12-15T02:00:00Z'] : [],
      timezone: 'America/Chicago',
    },
    progress: { name: 'Tier 1', raids: ["Onyxia's Lair"], named_encounters: 1, down: 1, pulls_total: 14 },
    wanted: [
      {
        class: 'paladin',
        spec: 'paladin-holy',
        spec_name: 'Holy Paladin',
        note: 'Two for the heal rotation',
      },
      { class: 'warrior', spec: 'warrior-protection', spec_name: 'Protection Warrior', note: 'Off tank' },
      { class: 'mage', spec: null, spec_name: null, note: '' },
    ],
    officers: [
      { name: 'Kraggor', class: 'warrior', spec: 'Protection', rank: 'leader' },
      { name: 'Vexlash', class: 'rogue', spec: 'Combat', rank: 'officer' },
    ],
    updated_at: '2026-12-01T18:00:00Z',
  };
}

function applicant(
  key: string,
  name: string,
  cls: string,
  specName: string,
  ilvl: number,
): ApplicantCharacter {
  return {
    character_key: key,
    name,
    class: cls,
    spec: null,
    spec_name: specName,
    level: 60,
    item_level: ilvl,
    best_parse: {
      metric: cls === 'paladin' ? 'hps' : 'dps',
      value: 465.4,
      percentile: 91,
      encounter: 'Onyxia',
      report_id: 'abc',
      fight_index: 3,
    },
  };
}

export function buildMockApplication(status: GuildApplication['status'] = 'pending'): GuildApplication {
  return {
    id: 41,
    guild_id: GUILD_ID,
    status,
    created_at: '2026-12-02T19:00:00Z',
    decided_at: status === 'pending' ? null : '2026-12-03T19:00:00Z',
    message: 'Two tiers healing in a guild that broke up. I can raid Tue and Thu.',
    availability: 'Tue, Thu',
    discord_dm: false,
    character: applicant(APPLICANT_KEY, 'Ashenveil', 'paladin', 'Holy Paladin', 63),
  };
}

function officerRow(
  base: GuildApplication,
  votes: OfficerApplication['votes'],
  extra: Partial<OfficerApplication> = {},
): OfficerApplication {
  return { ...base, votes, decided_by: null, role_outcome: null, dm_outcome: null, ...extra };
}

export function buildMockOfficerApplications(): OfficerApplications {
  const first = buildMockApplication();
  const second: GuildApplication = {
    ...buildMockApplication(),
    id: 42,
    created_at: '2026-12-04T19:00:00Z',
    message: 'New to Forever, level 60 since launch.',
    availability: '',
    discord_dm: true,
    character: applicant(SECOND_APPLICANT_KEY, 'Quillbright', 'warlock', 'Destruction Warlock', 60),
  };
  const decided: GuildApplication = {
    ...buildMockApplication('accepted'),
    id: 39,
    character: applicant('us/pvp/hexmarrow', 'Hexmarrow', 'priest', 'Holy Priest', 61),
  };
  return {
    rows: [
      officerRow(first, { yes: 2, no: 0, voters: 3, mine: 'yes' }),
      officerRow(second, { yes: 1, no: 1, voters: 3, mine: null }),
      officerRow(
        decided,
        { yes: 3, no: 0, voters: 3, mine: 'yes' },
        {
          decided_by: { name: 'Kraggor', class: 'warrior', character_key: 'us/pvp/kraggor' },
          role_outcome: 'granted',
          dm_outcome: 'not_opted_in',
        },
      ),
    ],
    pending: 2,
    discord_role_id: '123456789012345678',
  };
}

/** Accept's answer: the role granted, no DM (the applicant left the opt-in off). */
export function buildMockDecision(app: OfficerApplication, accept = true): ApplicationDecision {
  return {
    application: { ...app, status: accept ? 'accepted' : 'declined', decided_at: '2026-12-05T19:00:00Z' },
    discord_role: accept ? { done: true, reason: null } : { done: false, reason: 'not_accepted' },
    discord_dm: app.discord_dm ? { done: true, reason: null } : { done: false, reason: 'not_opted_in' },
  };
}

export function buildMockNotes(): MemberNote[] {
  return [
    {
      id: 7,
      body: 'Ask about the Thursday conflict.',
      author: { name: 'Kraggor', class: 'warrior', character_key: 'us/pvp/kraggor' },
      created_at: '2026-12-05T18:00:00Z',
      may_delete: true,
    },
    {
      id: 6,
      body: 'Always brings flasks.',
      author: { name: 'Vexlash', class: 'rogue', character_key: 'us/pvp/vexlash' },
      created_at: '2026-12-02T18:00:00Z',
      may_delete: false,
    },
  ];
}

/** A member's profile; with `officer`, the notes an officer's read carries. */
export function buildMockProfile(name = 'Zulmara', officer = false): MemberProfile {
  const member = buildMockRoster().find((row) => row.name === name) ?? null;
  const readiness = buildMockReadiness().rows.find((row) => row.name === name) ?? null;
  return {
    account_key: member?.account_key ?? MEMBER_ACCOUNT,
    member,
    characters:
      member === null
        ? []
        : [{ character_key: member.character_key, name, class: member.class ?? null, is_main: true }],
    attendance: {
      source: 'raids',
      last_4: { present: 3, nights: 4 },
      last_30: { present: 25, nights: 30 },
      nights: [
        { raid_id: 9, starts_at: '2026-11-24T02:00:00Z', outcome: 'present', excused: false },
        { raid_id: 10, starts_at: '2026-11-26T02:00:00Z', outcome: 'present', excused: false },
        { raid_id: 11, starts_at: '2026-12-01T02:00:00Z', outcome: 'absent', excused: false },
        { raid_id: 12, starts_at: '2026-12-03T02:00:00Z', outcome: 'late', excused: false },
      ],
      streak: 1,
    },
    loot: {
      count: 1,
      items: [
        {
          item_id: 16939,
          name: "Dragonstalker's Helm",
          icon: 'inv_helmet_05',
          quality: 4,
          raid_id: 11,
          awarded_at: '2026-12-01T04:00:00Z',
          ms_os: 'ms',
          reason: 'Best in slot, every night this month',
          votes_yes: 3,
          votes_no: 1,
          votes_officers: 4,
        },
      ],
      guild_median: 2,
      nights_since: 1,
    },
    readiness,
    upgrades: [
      {
        slot: 'hands',
        item_id: 16852,
        name: "Giantstalker's Gloves",
        source: 'Molten Core: Shazzrah',
        gain_dps: 11.4,
      },
      { slot: 'waist', item_id: 16851, name: "Giantstalker's Belt", source: 'Molten Core', gain_dps: null },
    ],
    ...(officer ? { notes: buildMockNotes() } : {}),
  };
}
