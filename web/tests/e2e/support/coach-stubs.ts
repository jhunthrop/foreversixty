// web/tests/e2e/support/coach-stubs.ts
// What the coach specs and captures share: the viewers (`/v1/me` answers), the coach route stub, and
// one fixture pack per panel state of the experience spec (design/specs/2026-10-28-log-coach.md
// section 4.2), all built from the worked example (a Fury warrior, "Bloodthirst: 9 casts lost").
// The report is the e2e fixture night (src/fixtures/report): Baelgrim (warrior) and Sunwick
// (priest) are in fight 3, Thalgrit (also in the pull) is the officer's pick.
import type { Page, Route } from '@playwright/test';
import { meAddonFixture } from '../../../src/fixtures/me-addon';
import type { Me } from '../../../src/lib/account/me-types';
import { premiumEntitlements } from '../../../src/test-support/billing-fixtures';
import {
  coachBusyReview,
  coachModelReview,
  coachNotReviewable,
  coachPreparing,
  coachRulesReview,
  coachTeaser,
} from '../../../src/test-support/coach-fixtures';
import type {
  CoachAnchor,
  CoachFull,
  CoachRecent,
  CoachTeaser,
  LastRaid,
} from '../../../src/lib/coach/types';
import { envelope } from './guild-billing';
import { plantSessionHint, routeMeCell, SIGNED_OUT } from './subscribe-stubs';

export const REPORT = '/reports/fixture2abcd';
export const FIGHT = 3;
export const GUID = {
  baelgrim: 'Player-4184-000000A1',
  sunwick: 'Player-4184-000000A2',
  thalgrit: 'Player-4184-000000A4',
};
export const KEY = {
  baelgrim: 'us/hardcore/baelgrim',
  sunwick: 'us/hardcore/sunwick',
  thalgrit: 'us/hardcore/thalgrit',
};
export const PANEL_URL = (guid: string = GUID.baelgrim): string =>
  `${REPORT}?fight=${FIGHT}&tab=coach&source=${guid}`;
export const ROWS_URL = `${REPORT}?fight=${FIGHT}&tab=damage-done`;
const COACH_ROUTE = /\/v1\/reports\/fixture2abcd\/fights\/3\/coach/;
const RATINGS_ROUTE = '**/v1/reports/fixture2abcd/fights/3/ratings';
const RECENT_ROUTE = '**/v1/me/coach/recent';

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

/** Signs the visitor in as `kind` (the cookie hint plus `/v1/me`); returns the cell to swap the account later. */
export async function signInAs(
  page: Page,
  kind: ViewerKind,
  owned: readonly (keyof typeof OWN)[] = ['baelgrim'],
): Promise<void> {
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
    ['Thalgrit', KEY.thalgrit, 'Warrior'],
    ['Baelgrim', KEY.baelgrim, 'Warrior'],
    ['Sunwick', KEY.sunwick, 'Priest'],
  ].map(([name, key, cls]) => ({
    player_key: key,
    player_name: name,
    class: cls,
    spec: 'Fury',
    role: 'dps',
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

export async function routeRecent(page: Page, recent: CoachRecent | number): Promise<void> {
  await page.route(RECENT_ROUTE, (route) =>
    typeof recent === 'number'
      ? route.fulfill(envelope({ code: 'internal', message: 'api sentence' }, recent))
      : route.fulfill(envelope(recent)),
  );
}

// ---- the worked example -------------------------------------------------------------------------

const BLOODTHIRST = 23881;
const WHIRLWIND = 1680;
const BATTLE_SHOUT = 6673;
const SUNDER_ARMOR = 11597;
const HEAL = 2060;
const DEATH_AT_MS = 200_000;

/** One anchor of every kind the chips can link to, with the report tab each opens. */
export const ANCHORS: Readonly<Record<string, CoachAnchor>> = {
  cast: { kind: 'cast', spell_id: BLOODTHIRST },
  cast2: { kind: 'cast', spell_id: WHIRLWIND },
  buffs: { kind: 'buffs' },
  aura: { kind: 'aura', spell_id: BATTLE_SHOUT },
  debuff: { kind: 'aura', spell_id: SUNDER_ARMOR, debuff: true },
  death: { kind: 'death', at_ms: DEATH_AT_MS },
  time: { kind: 'time', at_ms: 83_000, to_ms: 91_000 },
  resource: { kind: 'resource', at_ms: 120_000 },
  healing: { kind: 'healing', spell_id: HEAL },
};

/** The URL fragment each anchor id must land on. */
export const ANCHOR_LANDS: Readonly<Record<string, RegExp>> = {
  cast: /tab=casts.*ability=23881|ability=23881.*tab=casts/,
  cast2: /tab=casts.*ability=1680|ability=1680.*tab=casts/,
  buffs: /tab=buffs/,
  aura: /tab=buffs/,
  debuff: /tab=debuffs/,
  death: /tab=deaths/,
  time: /view=timelines/,
  resource: /view=timelines/,
  healing: /tab=healing/,
};

const WHY_LINK = 'The Bloodthirst gap at 1:23 includes your death at 3:20.';

/** The model's review of the worked example, with a Fix first, three Then items, Doing well and Couldn't check. */
export function workedReview(overrides: Partial<CoachFull> = {}): CoachFull {
  return coachModelReview({
    player_key: KEY.baelgrim,
    report_id: 'fixture2abcd',
    fight_index: FIGHT,
    review: {
      opening: 'A kill with one habit to fix: Bloodthirst sat unused while you had the rage for it.',
      fixFirst: {
        finding: 'core_casts',
        title: 'Bloodthirst: 9 casts lost',
        why: 'Bloodthirst went unused 9 times while you had the rage. The longest gaps were 1:23 and 2:41.',
        drill: 'core_casts',
        drillText: 'For the next pull, press Bloodthirst the moment it comes off cooldown.',
        anchors: ['cast', 'time'],
        links: [],
      },
      thenFix: [
        {
          finding: 'consumables',
          title: 'No flask at the pull',
          why: 'You pulled with no flask.',
          anchors: ['buffs', 'aura'],
          links: [],
        },
        {
          finding: 'deaths',
          title: 'Death at 3:20',
          why: 'Boss damage killed you at 3:20 with Last Stand ready.',
          anchors: ['death', 'resource'],
          links: [],
        },
        {
          finding: 'whirlwind',
          title: 'Whirlwind: 4 casts lost',
          why: 'Whirlwind sat ready for four casts in the first phase.',
          anchors: ['cast2', 'debuff'],
          links: [],
        },
      ],
      doingWell: [
        { strength: 'adherence', text: '91% of your casts were on the written rotation.', anchors: ['cast'] },
        { strength: 'uptime', text: 'Battle Shout was up for the whole pull.', anchors: ['aura'] },
      ],
      couldntCheck: [
        "Spell ranks: this log doesn't tell us your level.",
        "Defensive cooldowns: the coach doesn't know your spec's list yet.",
      ],
    },
    anchors: ANCHORS,
    basis_labels: ['your log', 'top Fury Warriors on Warden Kelthas (5)'],
    ...overrides,
  });
}

/** The same review as the rule-only rendering: the pack's own words, with the computed link lines. */
export function workedRules(overrides: Partial<CoachFull> = {}): CoachFull {
  const base = workedReview({ source: 'rules', persona: { label: 'Your Fury review', status: 'full' } });
  const fix = base.review.fixFirst;
  return {
    ...base,
    ...overrides,
    review: {
      ...base.review,
      fixFirst: fix === null ? null : { ...fix, links: [{ id: 'death_in_gap', sentence: WHY_LINK }] },
    },
  };
}

export const LINK_SENTENCE = WHY_LINK;

// ---- one pack per panel state -------------------------------------------------------------------

export type StateId =
  | 'ready-model'
  | 'ready-rules'
  | 'writing'
  | 'busy'
  | 'preparing'
  | 'no-findings'
  | 'not-reviewable'
  | 'teaser-two'
  | 'teaser-one'
  | 'teaser-empty'
  | 'officer'
  | 'starter'
  | 'healer'
  | 'raid-buffs-only'
  | 'error'
  | 'not-available';

const empty = (review: CoachFull): CoachFull => ({
  ...review,
  review: { ...review.review, fixFirst: null, thenFix: [] },
});

const healerReview = (): CoachFull => {
  const base = workedReview({ spec: 'Holy', role: 'healer', player_key: KEY.sunwick });
  return {
    ...base,
    review: {
      ...base.review,
      opening: 'A kill where Flash Heal carried the throughput.',
      fixFirst: {
        finding: 'overheal',
        title: 'Heal: 38% overheal',
        why: 'More than a third of your Heal never landed on a hurt target.',
        drill: 'overheal',
        drillText: 'Hold Heal until the tank is under 70% and use Flash Heal for the spikes.',
        anchors: ['healing'],
        links: [],
      },
      thenFix: [],
    },
  };
};

export const STATES: Readonly<Record<StateId, CoachStub>> = {
  'ready-model': { panel: workedReview() },
  'ready-rules': { panel: workedRules() },
  writing: { panel: workedRules({ status: 'writing' }) },
  busy: { panel: { ...coachBusyReview(), ...pickReview(workedRules()) } },
  preparing: { panel: coachPreparing() },
  'no-findings': { panel: empty(workedReview()) },
  'not-reviewable': { panel: coachNotReviewable() },
  'teaser-two': {
    panel: coachTeaser({
      finding_count: 2,
      fix_first: { title: 'Bloodthirst: 9 casts lost', code: 'core_casts', spell_id: BLOODTHIRST },
    }),
  },
  'teaser-one': {
    panel: coachTeaser({
      finding_count: 1,
      fix_first: { title: 'Bloodthirst: 9 casts lost', code: 'core_casts', spell_id: BLOODTHIRST },
    }),
  },
  'teaser-empty': { panel: coachTeaser({ finding_count: 0, fix_first: null }) },
  officer: { panel: workedReview({ access: 'officer' }) },
  starter: { panel: workedReview({ persona: { label: 'Your Fury coach', status: 'starter' } }) },
  healer: { panel: healerReview() },
  'raid-buffs-only': {
    panel: workedReview({
      review: { ...workedReview().review, fixFirst: null },
    }),
  },
  error: { panel: 500 },
  'not-available': { panel: 404 },
};

function pickReview(full: CoachFull): Pick<CoachFull, 'review' | 'anchors' | 'basis_labels' | 'player_key'> {
  return {
    review: full.review,
    anchors: full.anchors,
    basis_labels: full.basis_labels,
    player_key: KEY.baelgrim,
  };
}

/** The not-reviewable answer for a spec the rules cannot read: the title, the notice, then the reason. */
export const NOT_REVIEWABLE_WITH_REASON: CoachFull = workedRules({
  status: 'not_reviewable',
  review: {
    opening: '',
    fixFirst: null,
    thenFix: [],
    doingWell: [],
    couldntCheck: ["We couldn't tell this player's spec from the log."],
  },
});

// ---- the account line ---------------------------------------------------------------------------

const LAST_RAID: LastRaid = {
  zone: "Onyxia's Lair",
  fought_at: '2026-10-25T01:30:00Z',
  report_id: 'fixture2abcd',
  finding_count: 3,
  first: { stem: 'Bloodthirst casts lost', on: 5, of: 8 },
  url: '/reports/fixture2abcd?fight=3&tab=coach',
};

export type AccountLineId = 'premium' | 'free' | 'pending' | 'clean' | 'none' | 'error';

export const RECENT: Readonly<
  Record<AccountLineId, { viewer: 'premium' | 'free'; recent: CoachRecent | number }>
> = {
  premium: { viewer: 'premium', recent: { last_raid: LAST_RAID, pending: false, reviews: [] } },
  free: { viewer: 'free', recent: { last_raid: LAST_RAID, pending: false, reviews: [] } },
  pending: {
    viewer: 'premium',
    recent: { last_raid: { ...LAST_RAID, finding_count: 0, first: null }, pending: true, reviews: [] },
  },
  clean: {
    viewer: 'premium',
    recent: { last_raid: { ...LAST_RAID, finding_count: 0, first: null }, pending: false, reviews: [] },
  },
  none: { viewer: 'premium', recent: { last_raid: null, pending: false, reviews: [] } },
  error: { viewer: 'premium', recent: 500 },
};

/** The two stubs a Premium or free viewer needs on `/account`. */
export async function routeAccountDevices(page: Page): Promise<void> {
  await page.route('**/v1/devices', (route) => route.fulfill(envelope([])));
}

export { coachRulesReview };
