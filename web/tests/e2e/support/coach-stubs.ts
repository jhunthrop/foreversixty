// web/tests/e2e/support/coach-stubs.ts
// What the coach specs and captures share: the viewers (`/v1/me` answers), the coach route stub,
// the Onyxia pull the panel is read on, and one answer per panel state of the experience spec
// (design/specs/2026-10-28-log-coach.md section 4.2). Every answer is the coach's own output:
// src/fixtures/coach/*.json, written by api/internal/coach's TestWebFixtures_Write from the real
// rules, rule files, renderer and route shaping (and held there by TestWebFixtures_NoDrift).
// Nothing here types a review sentence.
//
// The report is the e2e fixture night (src/fixtures/report), re-read as the spec's worked pull:
// fight 3 is Onyxia (40), a 3:52 kill in Onyxia's Lair, Baelgrim (the Fury warrior) and Sunwick
// (the Holy priest) are in it, and Morrowlyn (a second Fury warrior) is the officer's pick.
import { readFileSync } from 'node:fs';
import path from 'node:path';
import { fileURLToPath } from 'node:url';
import type { Page, Route } from '@playwright/test';
import { meAddonFixture } from '../../../src/fixtures/me-addon';
import type { Me } from '../../../src/lib/account/me-types';
import { premiumEntitlements } from '../../../src/test-support/billing-fixtures';
import type { CoachFull, CoachRecent, CoachTeaser, LastRaid } from '../../../src/lib/coach/types';
import { envelope } from './guild-billing';
import { plantSessionHint, routeMeCell, SIGNED_OUT } from './subscribe-stubs';

export const REPORT = '/reports/fixture2abcd';
export const FIGHT = 3;
export const GUID = {
  baelgrim: 'Player-4184-000000A1',
  sunwick: 'Player-4184-000000A2',
  morrowlyn: 'Player-4184-000000A3',
};
export const KEY = {
  baelgrim: 'us/hardcore/baelgrim',
  sunwick: 'us/hardcore/sunwick',
  morrowlyn: 'us/hardcore/morrowlyn',
};
export const PANEL_URL = (guid: string = GUID.baelgrim): string =>
  `${REPORT}?fight=${FIGHT}&tab=coach&source=${guid}`;
export const ROWS_URL = `${REPORT}?fight=${FIGHT}&tab=damage-done`;
const COACH_ROUTE = /\/v1\/reports\/fixture2abcd\/fights\/3\/coach/;
const RATINGS_ROUTE = '**/v1/reports/fixture2abcd/fights/3/ratings';
export const RECENT_ROUTE = '**/v1/me/coach/recent';

// ---- the coach's own output ---------------------------------------------------------------------

const FIXTURES = path.resolve(path.dirname(fileURLToPath(import.meta.url)), '../../../src/fixtures/coach');

/** One file of src/fixtures/coach, fresh on every call so a caller may spread it. */
export function coachFixture<T = CoachFull>(name: string): T {
  return JSON.parse(readFileSync(path.join(FIXTURES, `${name}.json`), 'utf8')) as T;
}

// ---- the pull: Onyxia (40), 3:52 ----------------------------------------------------------------

/** The worked pull's facts (design spec §5.3), which the fixtures were built for. */
export const ONYXIA = {
  zone: "Onyxia's Lair",
  boss: 'Onyxia',
  size: 40,
  difficulty: 9,
  durationMs: 232_000,
} as const;
const FIXTURE_ZONE = 'Sanguine Depths';
const RETRY_MS = 38_000; // fight 4 becomes a second, short Onyxia pull after the kill

interface MetaFight {
  index: number;
  name: string;
  zone?: string;
  size?: number;
  difficulty?: number;
  duration_ms: number;
  start: string;
  end: string;
  kill: boolean;
}
interface ReportMeta {
  title?: string;
  zone?: string;
  fights: MetaFight[];
}

const at = (iso: string, ms: number): string => new Date(Date.parse(iso) + ms).toISOString();

/** The report's meta with fights 3 and 4 as Onyxia pulls: the 3:52 kill, then a short wipe. */
export function onyxiaMeta<T extends ReportMeta>(meta: T): T {
  const kill = meta.fights.find((f) => f.index === FIGHT);
  const killEnd = kill === undefined ? '' : at(kill.start, ONYXIA.durationMs);
  const fights = meta.fights.map((fight) => {
    const zoned = { ...fight, ...(fight.zone === undefined ? {} : { zone: ONYXIA.zone }) };
    if (fight.index === FIGHT)
      return {
        ...zoned,
        name: ONYXIA.boss,
        size: ONYXIA.size,
        difficulty: ONYXIA.difficulty,
        duration_ms: ONYXIA.durationMs,
        end: killEnd,
      };
    if (fight.index === FIGHT + 1 && killEnd !== '') {
      const start = at(killEnd, RETRY_MS);
      return {
        ...zoned,
        name: ONYXIA.boss,
        size: ONYXIA.size,
        difficulty: ONYXIA.difficulty,
        kill: false,
        start,
        end: at(start, fight.duration_ms),
      };
    }
    return zoned;
  });
  return {
    ...meta,
    ...(meta.zone === undefined ? {} : { zone: ONYXIA.zone }),
    ...(meta.title === undefined ? {} : { title: meta.title.replace(FIXTURE_ZONE, ONYXIA.zone) }),
    fights,
  };
}

/** The two warriors of the pull are Fury, as the worked review is. */
const FURY = new Set([GUID.baelgrim, GUID.morrowlyn]);
const FIXTURE_BOSS = 'Warden Kelthas';

/** Every unit record of the two Fury players, wherever the file carries one, as a Fury warrior. */
function asFury(value: unknown): unknown {
  if (Array.isArray(value)) return value.map(asFury);
  if (value === null || typeof value !== 'object') return value;
  const record = Object.fromEntries(Object.entries(value).map(([key, inner]) => [key, asFury(inner)]));
  if (!FURY.has(String(record.guid))) return record;
  return {
    ...record,
    ...('class' in record ? { class: 'Warrior' } : {}),
    ...('spec' in record ? { spec: 'Fury' } : {}),
  };
}

/** Fight 3's summary as the Onyxia kill: 3:52 long, the boss named, the warriors Fury. */
export function onyxiaSummary(summary: object): object {
  const named = JSON.parse(JSON.stringify(summary).replaceAll(FIXTURE_BOSS, ONYXIA.boss)) as object;
  return { ...(asFury(named) as object), duration_ms: ONYXIA.durationMs };
}

const ENTITIES: Readonly<Record<string, string>> = {
  '&quot;': '"',
  '&#39;': "'",
  '&lt;': '<',
  '&gt;': '>',
  '&amp;': '&',
};
const unescapeHtml = (s: string): string => s.replace(/&(quot|#39|lt|gt|amp);/g, (m) => ENTITIES[m] ?? m);
const escapeHtml = (s: string): string =>
  s
    .replace(/&/g, '&amp;')
    .replace(/"/g, '&quot;')
    .replace(/'/g, '&#39;')
    .replace(/</g, '&lt;')
    .replace(/>/g, '&gt;');

/** The report page with its inlined meta, and every place it names the zone, read as Onyxia's Lair. */
function onyxiaPage(html: string): string {
  return html
    .replace(/data-report="([^"]*)"/, (_m, raw: string) => {
      const meta = onyxiaMeta(JSON.parse(unescapeHtml(raw)) as ReportMeta);
      return `data-report="${escapeHtml(JSON.stringify(meta))}"`;
    })
    .replaceAll(FIXTURE_ZONE, escapeHtml(ONYXIA.zone));
}

/** Serves the fixture night as the worked Onyxia pull: the page, its report file and fight 3's summary. */
export async function routeOnyxiaPull(page: Page): Promise<void> {
  await page.route(/\/reports\/fixture2abcd(\?|$)/, async (route) => {
    if (route.request().resourceType() !== 'document') return route.fallback();
    const response = await route.fetch();
    await route.fulfill({ response, body: onyxiaPage(await response.text()) });
  });
  await page.route('**/logs-data/reports/fixture2abcd/report.json*', async (route) => {
    const response = await route.fetch();
    await route.fulfill({ response, json: asFury(onyxiaMeta((await response.json()) as ReportMeta)) });
  });
  await page.route('**/logs-data/reports/fixture2abcd/fights/3/summary.json*', async (route) => {
    const response = await route.fetch();
    await route.fulfill({ response, json: onyxiaSummary((await response.json()) as object) });
  });
}

// ---- viewers ------------------------------------------------------------------------------------

export type ViewerKind = 'premium' | 'free' | 'officer' | 'stranger' | 'signedOut';

const OWN = {
  baelgrim: { key: KEY.baelgrim, region: 'us', ruleset: 'hardcore', name: 'Baelgrim', class: 'Warrior' },
  sunwick: { key: KEY.sunwick, region: 'us', ruleset: 'hardcore', name: 'Sunwick', class: 'Priest' },
};

function meFor(kind: Exclude<ViewerKind, 'signedOut'>, owned: readonly (keyof typeof OWN)[]): Me {
  const characters = kind === 'officer' || kind === 'stranger' ? [] : owned.map((own) => OWN[own]);
  const entitlements = {
    ...premiumEntitlements({ officer_views: kind === 'officer' }),
    ...(kind === 'free' || kind === 'stranger' ? { log_coach: false } : {}),
  };
  return { ...meAddonFixture, characters: characters as unknown as Me['characters'], entitlements };
}

/** Signs the visitor in as `kind` (the cookie hint plus `/v1/me`) and serves the Onyxia pull. */
export async function signInAs(
  page: Page,
  kind: ViewerKind,
  owned: readonly (keyof typeof OWN)[] = ['baelgrim'],
): Promise<void> {
  await routeOnyxiaPull(page);
  if (kind === 'signedOut') {
    await page.route('**/v1/me', (route) => route.fulfill(SIGNED_OUT));
    return;
  }
  await plantSessionHint(page);
  await routeMeCell(page, meFor(kind, owned));
}

// ---- the coach route ----------------------------------------------------------------------------

/** A body for the coach route: an answer object, or a number for an HTTP error status. */
export type CoachAnswer = CoachFull | CoachTeaser | number;

export interface CoachStub {
  /** The answer to the panel's request (no `line`). */
  panel: CoachAnswer | (() => CoachAnswer);
  /** The answer to the strip's `line=1` request; defaults to `panel`. */
  line?: CoachAnswer | (() => CoachAnswer);
  /** Held until released: the loading state. */
  gate?: Promise<void>;
}

const resolve = <T>(value: T | (() => T)): T => (typeof value === 'function' ? (value as () => T)() : value);

function fulfil(route: Route, answer: CoachAnswer): Promise<void> {
  return typeof answer === 'number'
    ? route.fulfill(envelope({ code: 'internal', message: 'api sentence' }, answer))
    : route.fulfill(envelope(answer));
}

/** Records the requests it saw (`line` true for the strip's form) so a spec can assert who asked. */
export async function routeCoach(
  page: Page,
  stub: CoachStub,
): Promise<{ lines: number[]; panels: number[] }> {
  const seen = { lines: [] as number[], panels: [] as number[] };
  await page.route(COACH_ROUTE, async (route) => {
    const url = new URL(route.request().url());
    const isLine = url.searchParams.get('line') === '1';
    (isLine ? seen.lines : seen.panels).push(Date.now());
    if (stub.gate !== undefined) await stub.gate;
    const answer = isLine && stub.line !== undefined ? resolve(stub.line) : resolve(stub.panel);
    await fulfil(route, answer);
  });
  return seen;
}

/** The pull's rating rows, which is where an officer's pick finds the character key. */
export async function routeRatingKeys(page: Page): Promise<void> {
  const players = [
    ['Morrowlyn', KEY.morrowlyn, 'Warrior'],
    ['Baelgrim', KEY.baelgrim, 'Warrior'],
    ['Sunwick', KEY.sunwick, 'Priest'],
  ].map(([name, key, cls]) => ({
    player_key: key,
    player_name: name,
    class: cls,
    spec: cls === 'Priest' ? 'Holy' : 'Fury',
    role: cls === 'Priest' ? 'healer' : 'dps',
    overall: 70,
    overall_uncapped: 70,
    overall_capped: false,
    basis: 'percentile',
    components: [],
  }));
  await page.route(RATINGS_ROUTE, (route) =>
    route.fulfill(
      envelope({ fight_index: FIGHT, kill: true, kill_time_band: 'typical', model_version: 'v1', players }),
    ),
  );
}

export async function routeRecent(
  page: Page,
  recent: CoachRecent | number,
  gate?: Promise<void>,
): Promise<void> {
  await page.route(RECENT_ROUTE, async (route) => {
    if (gate !== undefined) await gate;
    await (typeof recent === 'number'
      ? route.fulfill(envelope({ code: 'internal', message: 'api sentence' }, recent))
      : route.fulfill(envelope(recent)));
  });
}

// ---- one answer per panel state -----------------------------------------------------------------

export type StateId =
  | 'ready-model'
  | 'ready-rules'
  | 'writing'
  | 'busy'
  | 'preparing'
  | 'no-findings'
  | 'not-reviewable'
  | 'not-reviewable-short'
  | 'not-reviewable-spec'
  | 'teaser-several'
  | 'teaser-one'
  | 'teaser-empty'
  | 'teaser-no-first'
  | 'officer'
  | 'starter'
  | 'healer'
  | 'raid-buffs-only'
  | 'error'
  | 'not-available';

/** Each state's coach output, by fixture file; `error` and `not-available` are HTTP statuses. */
const STATE_FILES: Readonly<Record<StateId, string | number>> = {
  'ready-model': 'fury-model',
  'ready-rules': 'fury-rules',
  writing: 'fury-writing',
  busy: 'fury-busy',
  preparing: 'fury-preparing',
  'no-findings': 'fury-no-findings',
  'not-reviewable': 'fury-not-reviewable-trash',
  'not-reviewable-short': 'fury-not-reviewable-short',
  'not-reviewable-spec': 'not-reviewable-spec',
  'teaser-several': 'fury-teaser-several',
  'teaser-one': 'fury-teaser-one',
  'teaser-empty': 'fury-teaser-empty',
  'teaser-no-first': 'fury-teaser-no-first',
  officer: 'fury-officer',
  starter: 'fury-starter',
  healer: 'holy-model',
  'raid-buffs-only': 'fury-raid-buffs-only',
  error: 500,
  'not-available': 404,
};

export const STATES: Readonly<Record<StateId, CoachStub>> = Object.fromEntries(
  Object.entries(STATE_FILES).map(([id, file]) => [
    id,
    { panel: typeof file === 'number' ? file : coachFixture<CoachFull | CoachTeaser>(file) },
  ]),
) as Record<StateId, CoachStub>;

/** The writing answer until `ready()` says so, then the model's review: the A5 swap. */
export function writingThenModel(ready: () => boolean): CoachStub {
  return { panel: () => (ready() ? coachFixture('fury-model') : coachFixture('fury-writing')) };
}

/** The rule-only review's first link line, the pack's sentence word for word. */
export const LINK_SENTENCE: string = coachFixture('fury-rules').review.fixFirst?.links[0]?.sentence ?? '';

/** The worked review's Fix first title, as the pack words it. */
export const FIX_FIRST_TITLE: string = coachFixture('fury-model').review.fixFirst?.title ?? '';

// ---- the account line ---------------------------------------------------------------------------

const lastRaid = (file: string): LastRaid => coachFixture<LastRaid>(file);
const recentOf = (last_raid: LastRaid | null, pending = false): CoachRecent => ({
  last_raid,
  pending,
  reviews: [],
});

export type AccountLineId = 'premium' | 'premium-one' | 'free' | 'pending' | 'clean' | 'none' | 'error';

export const RECENT: Readonly<
  Record<AccountLineId, { viewer: 'premium' | 'free'; recent: CoachRecent | number }>
> = {
  premium: { viewer: 'premium', recent: recentOf(lastRaid('last-raid-molten-core')) },
  'premium-one': { viewer: 'premium', recent: recentOf(lastRaid('last-raid-onyxia')) },
  free: { viewer: 'free', recent: recentOf(lastRaid('last-raid-molten-core')) },
  pending: { viewer: 'premium', recent: recentOf(lastRaid('last-raid-clean'), true) },
  clean: { viewer: 'premium', recent: recentOf(lastRaid('last-raid-clean')) },
  none: { viewer: 'premium', recent: recentOf(null) },
  error: { viewer: 'premium', recent: 500 },
};

/** The two stubs a Premium or free viewer needs on `/account`. */
export async function routeAccountDevices(page: Page): Promise<void> {
  await page.route('**/v1/devices', (route) => route.fulfill(envelope([])));
}
