// web/tests/e2e/support/guild-raids.ts
// Stubs for guild raids and sign-ups (design/specs/2026-10-26-guild-raids.md §9), shared by
// tests/e2e/guild-raids.spec.ts and the acceptance captures in tests/captures/guild-raids/.
// `RaidStub` is a small stand-in for the API: it keeps the raids, answers each read as the
// viewer would see it (own flags, counts, `next_raid`), applies sign-ups with ruling R3 (one
// character per account at In or Late), records every write's body, and refuses on demand.
// The guild is the seeded mock guild (src/fixtures/guild/mock-guild.ts); a fixture build
// prerenders no raid route, so the raid URL is served the guild page's shell, as the Worker does.
import type { Page, Route } from '@playwright/test';
import {
  GUILD,
  GUILD_ID,
  VIEWER_MEMBER,
  VIEWER_OFFICER,
  buildMockGuildPage,
  buildMockHome,
  buildMockLoot,
  buildMockProgression,
  buildMockRaids,
  buildMockReadiness,
  buildMockRoster,
} from '../../../src/fixtures/guild/mock-guild';
import type {
  GuildRaidRole,
  GuildRaidStatus,
  GuildRaidDiscord,
  GuildRaidDiscordPost,
  GuildRaidView,
  GuildRaidsPage,
  GuildSignupState,
  GuildZoneView,
  NextRaidView,
} from '../../../src/lib/guild/api';
import { instantToWallClock, wallClockToInstant } from '../../../src/lib/guild/raid-time';
import type { RaidInput } from '../../../src/lib/guild/raids-api';
import { SPECS } from '../../../src/lib/sim/specs';
import { CHICAGO, NOW, NO_LOG_YET, plannedRaid } from '../../../src/test-support/raid-fixtures';
import { GUILD_URL, envelope, refusal } from './officer-tools';

export { CHICAGO, NOW, GUILD_URL };

/** The member's alt (the mock roster's `ALT_OF`): two own characters, so R3 has one to move. */
export const MEMBER_ALT = 'Duskstrider';
export const ONYXIA_ID = 7;
export const BARROW_ID = 8;
export const HYJAL_ID = 9;
/** The id the stub gives the first raid an officer plans. */
export const FIRST_NEW_ID = 100;

export const ZONES: GuildZoneView[] = [
  { key: 'raid:barrow-deeps', name: 'Barrow Deeps', size: null },
  { key: 'raid:hyjal-summit', name: 'Hyjal Summit', size: null },
  { key: 'raid:onyxias-lair', name: "Onyxia's Lair", size: 40 },
];

/** Who is looking: a verified member with an alt, the officer, or one of R1's three outsiders. */
export type RaidViewer = 'member' | 'officer' | 'visitor' | 'outsider' | 'unverified';

export const raidUrl = (id: number): string => `${GUILD_URL}/raids/${id}`;
export const keyOf = (name: string): string => `us/pvp/${name.toLowerCase().replace(/[^a-z0-9]+/g, '-')}`;

const ROSTER = new Map(buildMockRoster().map((row) => [row.name, row]));

function rosterRow(name: string) {
  const row = ROSTER.get(name);
  if (row === undefined) throw new Error(`guild-raids stub: no roster row "${name}"`);
  return row;
}

export interface SignupSeed {
  name: string;
  state: GuildSignupState;
  note?: string;
  role?: GuildRaidRole;
}

interface StoredSignup {
  character_key: string;
  name: string;
  class: string;
  spec: string;
  role: GuildRaidRole;
  state: GuildSignupState;
  note: string;
}

/** A raid as the stub keeps it: the view's fields without the viewer-shaped ones. */
export interface StoredRaid {
  id: number;
  zone: string;
  zone_name: string;
  title: string;
  starts_at: string;
  duration_minutes: number;
  status: GuildRaidStatus;
  size: number | null;
  notes_html: string;
  notes: string;
  signups: StoredSignup[];
}

function signup(seed: SignupSeed): StoredSignup {
  const row = rosterRow(seed.name);
  return {
    character_key: row.character_key,
    name: row.name,
    class: row.class ?? '',
    spec: row.spec ?? '',
    role: seed.role ?? row.role ?? 'dps',
    state: seed.state,
    note: seed.note ?? '',
  };
}

/** Onyxia's Lair, Thu 10 Dec 2026 8:00 pm Chicago: 13 In, 1 Late, 1 Bench, 1 Out; the member's
 *  alt is In and the member is not signed up. */
export const ONYXIA_SIGNUPS: SignupSeed[] = [
  { name: 'Kraggor', state: 'in', note: 'Bringing the Onyxia Scale Cloaks.' },
  { name: 'Lightbrand', state: 'in' },
  { name: 'Mendwhisper', state: 'in' },
  { name: 'Earthhoof', state: 'in', note: 'Can raid lead if needed.' },
  { name: 'Obnoxious Yell', state: 'in' },
  { name: 'Sunderfel', state: 'in' },
  { name: 'Windtalon', state: 'in' },
  { name: 'Felsnap', state: 'in' },
  { name: 'Vexlash', state: 'in' },
  { name: 'Shadowquill', state: 'in' },
  { name: 'Hexbramble', state: 'in' },
  { name: 'Pyrewisp', state: 'in' },
  { name: MEMBER_ALT, state: 'in' },
  { name: 'Arcanemoor', state: 'late', note: 'There by 8:30.' },
  { name: 'Soulgrave', state: 'bench' },
  { name: 'Frostnettle', state: 'out', note: 'Away this week.' },
];

const NOTES_HTML =
  '<p>Bring <strong>fire resistance</strong> and flasks.</p><ul><li>Onyxia Scale Cloak on everyone</li><li>Whelps: AoE on the side groups</li></ul><p>Strategy: <a href="https://foreversixty.gg/guides">the guide</a>.</p>';
const NOTES_MD =
  'Bring **fire resistance** and flasks.\n\n- Onyxia Scale Cloak on everyone\n- Whelps: AoE on the side groups\n\nStrategy: [the guide](https://foreversixty.gg/guides).';

export function storedRaid(extra: Partial<StoredRaid> & { seeds?: SignupSeed[] } = {}): StoredRaid {
  const base = plannedRaid();
  const { seeds, ...rest } = extra;
  return {
    id: base.id,
    zone: 'raid:onyxias-lair',
    zone_name: base.zone_name,
    title: base.title,
    starts_at: base.starts_at,
    duration_minutes: base.duration_minutes,
    status: 'planned',
    size: base.size,
    notes_html: NOTES_HTML,
    notes: NOTES_MD,
    signups: (seeds ?? ONYXIA_SIGNUPS).map(signup),
    ...rest,
  };
}

/** The Raids tab's three planned raids (§9 capture 1): Onyxia, Barrow Deeps (no size, R6) and a
 *  cancelled Hyjal Summit. */
export function threePlanned(): StoredRaid[] {
  return [
    storedRaid(),
    storedRaid({
      id: BARROW_ID,
      zone: 'raid:barrow-deeps',
      zone_name: 'Barrow Deeps',
      title: 'Barrow Deeps: first clear',
      starts_at: '2026-12-12T02:00:00Z',
      size: null,
      notes_html: '',
      notes: '',
      seeds: ONYXIA_SIGNUPS.slice(0, 9).filter((s) => s.name !== MEMBER_ALT),
    }),
    storedRaid({
      id: HYJAL_ID,
      zone: 'raid:hyjal-summit',
      zone_name: 'Hyjal Summit',
      title: 'Hyjal Summit',
      starts_at: '2026-12-14T02:00:00Z',
      status: 'cancelled',
      size: null,
      notes_html: '',
      notes: '',
      seeds: ONYXIA_SIGNUPS.slice(0, 5),
    }),
  ];
}

export interface RaidStubOptions {
  viewer: RaidViewer;
  /** `home.viewer.officer_tools`; left out, the field is absent. */
  officerTools?: boolean;
  raids?: StoredRaid[];
  /** `discord.available`: the deployment has Discord connect on; left out, it is. */
  discordAvailable?: boolean;
  discordConnected?: boolean;
  /** `discord.post_default`: the guild's channel posts raids by default. */
  postDefault?: boolean;
  /** `discord.remind_default`; left out, it is on. */
  remindDefault?: boolean;
  /** `discord.channel_name`: the bot's channel; left out, a webhook. */
  channelName?: string | null;
  /** `discord_post` on every raid view; left out, none is scheduled. */
  discordPost?: GuildRaidDiscordPost | null;
  guildTimezone?: string | null;
  /** The viewer's characters have no saved raid spec, so a sign-up asks for one (spec P.10.7);
   *  left out, each already has the spec its roster row names. */
  askSpec?: boolean;
}

export interface RecordedWrite {
  method: string;
  path: string;
  body: unknown;
}

interface Refusal {
  code: string;
  status: number;
  message?: string;
  fields?: Record<string, string>;
}

const REFUSAL_LINES: Record<string, string> = {
  raid_locked: 'This raid is locked.',
  raid_cancelled: 'This raid was cancelled.',
  raid_cap: 'A guild plans at most 12 raids ahead.',
  guild_plan_required: 'Planning and running raids are part of the guild plan.',
  conflict: 'The raid moved on.',
};

const OWN_NAMES: Record<RaidViewer, string[]> = {
  member: [VIEWER_MEMBER, MEMBER_ALT],
  officer: [VIEWER_OFFICER],
  visitor: [],
  outsider: [],
  unverified: [VIEWER_MEMBER],
};

const ON_THE_ROSTER: GuildSignupState[] = ['in', 'late'];
const DAY_MS = 86_400_000;

function counts(signups: readonly StoredSignup[]): GuildRaidView['counts'] {
  const n = (state: GuildSignupState) => signups.filter((s) => s.state === state).length;
  const inRole = (role: GuildRaidRole) => signups.filter((s) => s.state === 'in' && s.role === role).length;
  return {
    in: n('in'),
    late: n('late'),
    bench: n('bench'),
    out: n('out'),
    tanks: inRole('tank'),
    healers: inRole('healer'),
    dps: inRole('dps'),
  };
}

const ROLE_ORDER: GuildRaidRole[] = ['tank', 'healer', 'dps'];
const CLASS_ORDER = ['warrior', 'paladin', 'hunter', 'rogue', 'priest', 'shaman', 'mage', 'warlock', 'druid'];

function apiOrder(a: StoredSignup, b: StoredSignup): number {
  return (
    ROLE_ORDER.indexOf(a.role) - ROLE_ORDER.indexOf(b.role) ||
    CLASS_ORDER.indexOf(a.class) - CLASS_ORDER.indexOf(b.class) ||
    a.name.localeCompare(b.name)
  );
}

function addDays(date: string, days: number): string {
  return new Date(Date.parse(`${date}T00:00:00Z`) + days * DAY_MS).toISOString().slice(0, 10);
}

/** The API's repeat rule (R4, R7): the same guild wall time on each of the next weeks. */
export function repeatStarts(startsAt: string, timezone: string, weeks: number): string[] {
  const { date, time } = instantToWallClock(new Date(startsAt), timezone);
  return Array.from({ length: weeks + 1 }, (_, week) =>
    wallClockToInstant(addDays(date, week * 7), time, timezone).toISOString(),
  );
}

export class RaidStub {
  viewer: RaidViewer;
  officerTools: boolean | undefined;
  raids: StoredRaid[];
  discordAvailable: boolean;
  discordConnected: boolean;
  postDefault: boolean;
  remindDefault: boolean;
  channelName: string | null;
  discordPost: GuildRaidDiscordPost | null;
  guildTimezone: string | null;
  readonly writes: RecordedWrite[] = [];
  /** Every raid read, by id, so R1's tests can say none was made. */
  readonly raidReads: number[] = [];
  /** The next sign-up, create or status write answers this instead (once). */
  refuseNext: { signup?: Refusal; create?: Refusal; status?: Refusal } = {};
  private nextId = FIRST_NEW_ID;
  /** The saved raid spec per character key (a catalogue slug); written by a sign-up's `spec`. */
  readonly raidSpecs = new Map<string, string>();

  constructor(options: RaidStubOptions) {
    this.viewer = options.viewer;
    if (options.askSpec !== true) this.seedRaidSpecs();
    this.officerTools = options.officerTools;
    this.raids = options.raids ?? [storedRaid()];
    this.discordAvailable = options.discordAvailable ?? true;
    this.discordConnected = options.discordConnected ?? false;
    this.postDefault = options.postDefault ?? false;
    this.remindDefault = options.remindDefault ?? true;
    this.channelName = options.channelName ?? null;
    this.discordPost = options.discordPost ?? null;
    this.guildTimezone = options.guildTimezone === undefined ? CHICAGO : options.guildTimezone;
  }

  get ownKeys(): string[] {
    return OWN_NAMES[this.viewer].map(keyOf);
  }

  get verifiedMember(): boolean {
    return this.viewer === 'member' || this.viewer === 'officer';
  }

  raid(id: number): StoredRaid {
    const raid = this.raids.find((r) => r.id === id);
    if (raid === undefined) throw new Error(`guild-raids stub: no raid ${id}`);
    return raid;
  }

  /** Replaces one raid (the poll test changes the sign-ups between two reads). */
  update(id: number, change: (raid: StoredRaid) => StoredRaid): void {
    this.raids = this.raids.map((r) => (r.id === id ? change(r) : r));
  }

  /** Each own character starts with the spec its roster row names. */
  private seedRaidSpecs(): void {
    for (const name of OWN_NAMES[this.viewer]) {
      const row = rosterRow(name);
      const spec = SPECS.find((s) => s.class_slug === row.class && s.name === row.spec);
      if (spec !== undefined) this.raidSpecs.set(row.character_key, spec.spec);
    }
  }

  /** The raid view as this viewer reads it. */
  view(id: number): GuildRaidView {
    const raid = this.raid(id);
    const own = new Set(this.ownKeys);
    const signups = [...raid.signups].sort(apiOrder);
    const signed = new Set(signups.map((s) => s.character_key));
    const notSigned = buildMockRoster()
      .filter((r) => r.verified && !signed.has(r.character_key))
      .map((r) => ({
        character_key: r.character_key,
        name: r.name,
        class: r.class ?? '',
        spec: r.spec ?? '',
      }));
    return {
      ...NO_LOG_YET,
      ...raid,
      guild_timezone: this.guildTimezone,
      counts: counts(raid.signups),
      signups: signups.map((s) => ({ ...s, own: own.has(s.character_key) })),
      not_signed: notSigned,
      viewer: {
        characters: OWN_NAMES[this.viewer].map((name) => {
          const row = rosterRow(name);
          const mine = raid.signups.find((s) => s.character_key === row.character_key);
          return {
            character_key: row.character_key,
            name: row.name,
            class: row.class ?? '',
            spec: row.spec ?? '',
            verified: row.verified,
            state: mine?.state ?? null,
            role: mine?.role ?? row.role ?? 'dps',
            raid_spec: this.raidSpecs.get(row.character_key) ?? null,
          };
        }),
      },
      discord: this.discordView(),
      discord_post: this.discordPost,
    };
  }

  /** The guild channel as the raid view carries it. */
  discordView(): GuildRaidDiscord {
    return {
      available: this.discordAvailable,
      connected: this.discordConnected,
      post_default: this.postDefault,
      remind_default: this.remindDefault,
      channel_name: this.channelName,
    };
  }

  /** The guild's next planned or locked raid, as `/home` carries it (raids spec §9 contract). */
  nextRaid(): NextRaidView | undefined {
    const next = this.raids
      .filter((r) => r.status === 'planned' || r.status === 'locked')
      .sort((a, b) => Date.parse(a.starts_at) - Date.parse(b.starts_at))[0];
    if (next === undefined || !this.verifiedMember) return undefined;
    const view = this.view(next.id);
    const current = keyOf(OWN_NAMES[this.viewer][0]);
    const mine = view.signups.filter((s) => s.own);
    const own = mine.find((s) => s.character_key === current) ?? mine[0];
    const { in: inCount, late, tanks, healers, dps } = view.counts;
    return {
      id: next.id,
      zone: next.zone,
      zone_name: next.zone_name,
      title: next.title,
      starts_at: next.starts_at,
      status: next.status === 'locked' ? 'locked' : 'planned',
      size: next.size,
      guild_timezone: this.guildTimezone,
      counts: { in: inCount, late, tanks, healers, dps },
      own:
        own === undefined
          ? null
          : { character_key: own.character_key, state: own.state, role: own.role, note: own.note },
    };
  }

  raidsPage(): GuildRaidsPage {
    const page = buildMockRaids();
    if (!this.verifiedMember) return page;
    return {
      ...page,
      planned: this.raids.map((r) => this.view(r.id)),
      zones: ZONES,
      guild_timezone: this.guildTimezone,
      discord: this.discordView(),
    };
  }

  /** PUT …/signups/{key}: R3 moves the account's other In or Late character to Bench. */
  applySignup(
    id: number,
    key: string,
    body: { state: GuildSignupState; role: GuildRaidRole; note: string; spec?: string },
  ): { raid: GuildRaidView; moved: string | null } {
    const own = this.ownKeys;
    if (body.spec !== undefined) this.raidSpecs.set(key, body.spec);
    let moved: string | null = null;
    this.update(id, (raid) => {
      const others = raid.signups
        .filter((s) => s.character_key !== key)
        .map((s) => {
          const bumped =
            ON_THE_ROSTER.includes(body.state) &&
            own.includes(key) &&
            own.includes(s.character_key) &&
            ON_THE_ROSTER.includes(s.state);
          if (!bumped) return s;
          moved = s.name;
          return { ...s, state: 'bench' as const };
        });
      const name = [...ROSTER.values()].find((r) => r.character_key === key)?.name;
      if (name === undefined) throw new Error(`guild-raids stub: no roster row for ${key}`);
      return {
        ...raid,
        signups: [...others, { ...signup({ name, state: body.state, role: body.role }), note: body.note }],
      };
    });
    return { raid: this.view(id), moved };
  }

  /** POST …/raids: one raid, or four at the same guild wall time with `repeat_weeks: 3`. */
  applyCreate(body: RaidInput): GuildRaidView[] {
    const zone = ZONES.find((z) => z.key === body.zone);
    if (zone === undefined) throw new Error(`guild-raids stub: unknown zone ${body.zone}`);
    this.guildTimezone = body.timezone;
    this.postDefault = body.post_to_discord;
    this.remindDefault = body.remind_day_before;
    const made = repeatStarts(body.starts_at, body.timezone, body.repeat_weeks ?? 0).map((startsAt) =>
      storedRaid({
        id: this.nextId++,
        zone: zone.key,
        zone_name: zone.name,
        title: body.title === '' ? zone.name : body.title,
        starts_at: startsAt,
        duration_minutes: body.duration_minutes,
        size: zone.size,
        notes: body.notes,
        notes_html: body.notes === '' ? '' : `<p>${body.notes}</p>`,
        seeds: [],
      }),
    );
    this.raids = [...this.raids, ...made];
    return made.map((r) => this.view(r.id));
  }

  applyStatus(id: number, action: string): GuildRaidView {
    const status: Record<string, GuildRaidStatus> = {
      lock: 'locked',
      unlock: 'planned',
      cancel: 'cancelled',
    };
    this.update(id, (raid) => ({ ...raid, status: status[action] }));
    return this.view(id);
  }

  takeRefusal(kind: 'signup' | 'create' | 'status'): Refusal | undefined {
    const found = this.refuseNext[kind];
    this.refuseNext = { ...this.refuseNext, [kind]: undefined };
    return found;
  }
}

function refused(found: Refusal) {
  return {
    status: found.status,
    contentType: 'application/json',
    body: JSON.stringify({
      ok: false,
      data: null,
      error: {
        code: found.code,
        message: found.message ?? REFUSAL_LINES[found.code] ?? 'Refused.',
        ...(found.fields === undefined ? {} : { fields: found.fields }),
      },
      request_id: 'req-test',
    }),
  };
}

const SIGNED_OUT = {
  status: 401,
  contentType: 'application/json',
  body: '{"ok":false,"data":null,"error":null,"request_id":"r"}',
};

/** `/v1/me` per viewer: the member carries both characters, the unverified member's row is not
 *  verified yet, and an outsider is signed in with no row for this guild. */
export function raidsMe(viewer: RaidViewer) {
  const names = OWN_NAMES[viewer];
  const rank = viewer === 'officer' ? 'officer' : 'member';
  return {
    user: { id: 7, battletag: 'Fixture#1234', email: null, role: 'user', anonymize: false, premium: false },
    characters: names.map((name) => ({
      key: keyOf(name),
      region: 'us',
      ruleset: 'pvp',
      name,
      class: rosterRow(name).class,
    })),
    guilds:
      viewer === 'outsider'
        ? []
        : [
            {
              id: GUILD_ID,
              region: GUILD.region,
              ruleset: GUILD.ruleset,
              name: GUILD.name,
              rank,
              verified: viewer !== 'unverified',
              can: viewer === 'officer' ? ['approve', 'remove', 'settings', 'award'] : [],
            },
          ],
  };
}

function homeFor(stub: RaidStub) {
  const name = stub.verifiedMember ? OWN_NAMES[stub.viewer][0] : null;
  return buildMockHome(name, GUILD.faction, null, true, {
    officerTools: stub.officerTools,
    nextRaid: stub.nextRaid(),
  });
}

const RAIDS_PATH = new RegExp(`/v1/guilds/${GUILD_ID}/raids(?:/(\\d+)(?:/(.+))?)?(?:\\?.*)?$`);

async function answerRaids(stub: RaidStub, route: Route): Promise<void> {
  const request = route.request();
  const method = request.method();
  const match = RAIDS_PATH.exec(new URL(request.url()).pathname + new URL(request.url()).search);
  if (match === null) return route.fallback();
  const [, idText, rest] = match;
  const body = method === 'GET' ? null : (request.postDataJSON() as unknown);
  if (method !== 'GET') stub.writes.push({ method, path: new URL(request.url()).pathname, body });
  if (idText === undefined) {
    if (method === 'GET') return route.fulfill(envelope(stub.raidsPage()));
    const refusedCreate = stub.takeRefusal('create');
    if (refusedCreate !== undefined) return route.fulfill(refused(refusedCreate));
    return route.fulfill(envelope({ raids: stub.applyCreate(body as RaidInput) }, 201));
  }
  const id = Number(idText);
  if (method === 'GET' && rest === undefined) {
    stub.raidReads.push(id);
    if (!stub.verifiedMember) return route.fulfill(refusal('not_found', 'Not found.', 404));
    return route.fulfill(envelope(stub.view(id)));
  }
  if (method === 'PUT' && rest?.startsWith('signups/') === true) {
    const refusedSignup = stub.takeRefusal('signup');
    if (refusedSignup !== undefined) return route.fulfill(refused(refusedSignup));
    const key = rest.slice('signups/'.length).split('/').map(decodeURIComponent).join('/');
    const signupBody = body as { state: GuildSignupState; role: GuildRaidRole; note: string; spec?: string };
    return route.fulfill(envelope(stub.applySignup(id, key, signupBody)));
  }
  if (method === 'POST' && (rest === 'lock' || rest === 'unlock' || rest === 'cancel')) {
    const refusedStatus = stub.takeRefusal('status');
    if (refusedStatus !== undefined) return route.fulfill(refused(refusedStatus));
    return route.fulfill(envelope(stub.applyStatus(id, rest)));
  }
  return route.fulfill(refusal('not_found', 'Not found.', 404));
}

/** Serves the guild page's prerendered shell at a raid URL (the Worker's job in production). */
async function serveRaidShell(route: Route): Promise<void> {
  if (route.request().resourceType() !== 'document') return route.fallback();
  const shell = new URL(GUILD_URL, route.request().url()).toString();
  const response = await route.fetch({ url: shell });
  return route.fulfill({ response });
}

/** Routes the guild's public page, /v1/me, /home, the raids endpoints, the three depth tabs and
 *  the raid route's shell; answers the stub that holds the raids and records the writes. */
export async function stubGuildRaids(page: Page, options: RaidStubOptions): Promise<RaidStub> {
  const stub = new RaidStub(options);
  await page.route(`**/guild/${GUILD.region}/${GUILD.ruleset}/olympus-xxvii/raids/*`, serveRaidShell);
  await page.route(`**/v1/guilds/${GUILD.region}/${GUILD.ruleset}/olympus-xxvii`, (route) =>
    route.fulfill(envelope(buildMockGuildPage())),
  );
  await page.route('**/v1/me', (route) =>
    stub.viewer === 'visitor' ? route.fulfill(SIGNED_OUT) : route.fulfill(envelope(raidsMe(stub.viewer))),
  );
  await page.route(`**/v1/guilds/${GUILD_ID}/home`, (route) => route.fulfill(envelope(homeFor(stub))));
  await page.route(RAIDS_PATH, (route) => answerRaids(stub, route));
  await page.route(`**/v1/guilds/${GUILD_ID}/progression`, (route) =>
    route.fulfill(envelope(buildMockProgression())),
  );
  await page.route(`**/v1/guilds/${GUILD_ID}/readiness`, (route) =>
    route.fulfill(envelope(buildMockReadiness())),
  );
  await page.route(`**/v1/guilds/${GUILD_ID}/loot*`, (route) => route.fulfill(envelope(buildMockLoot())));
  return stub;
}

/** The writes of one method whose path ends with `suffix`. */
export function writesTo(stub: RaidStub, method: string, suffix: RegExp): RecordedWrite[] {
  return stub.writes.filter((w) => w.method === method && suffix.test(w.path));
}

/** The test id of a strip row's part for a character (`signup-segments-us/pvp/zulmara`). */
export const stripId = (
  part: 'row' | 'segments' | 'role' | 'note' | 'status' | 'spec',
  name: string,
): string => `signup-${part}-${keyOf(name)}`;
