// web/src/test-support/subscribe-fixtures.ts
// Typed `/v1/me` fixtures for subscribe-and-join (spec design/specs/2026-10-27-subscribe-and-join.md
// section 8): the owner's account as production holds it, one fixture per subscribe step, one per
// premium source, and the empty and older-API shapes. Each is a full `Me`, so `npm run check` proves
// it against the contract; e2e stubs it through tests/e2e/support/checkout-guild.ts.
import type { BillingView, Me, MeCharacter, MeGuild } from '../lib/account/me-types';
import type {
  BillingFacts,
  GuildClaimView,
  GuildCoveredView,
  GuildSubscribeView,
  SubscribeStep,
} from '../lib/billing/contract';
import { PLANS } from '../lib/billing/plans';
import { NO_PRICE_FACTS, mockEntitlements, premiumEntitlements } from './billing-fixtures';

/** Freezes a fixture and everything under it, so a test cannot mutate what other tests share. */
function deepFreeze<T>(value: T): T {
  if (typeof value === 'object' && value !== null && !Object.isFrozen(value)) {
    Object.values(value).forEach(deepFreeze);
    Object.freeze(value);
  }
  return value;
}

// --- builders -----------------------------------------------------------------------------------

export const OLYMPUS = { id: 2, name: 'OLYMPUS XXVII', slug: 'olympus-xxvii' } as const;
export const MILLENNIALS = { id: 3, name: 'Employed Millennials', slug: 'employed-millennials' } as const;

const UNCLAIMED: GuildClaimView = {
  state: 'unclaimed',
  yours: false,
  pending_yours: false,
  holder_name: null,
};

export function subscribeView(
  step: SubscribeStep,
  overrides: Partial<GuildSubscribeView> = {},
): GuildSubscribeView {
  return { step, trial_used: null, available_at: null, reason: null, other_guild_name: null, ...overrides };
}

/** A plan's coverage; `seats_total` follows the tier (Guild 1, Raid guild 3) unless overridden. */
export function coveredView(overrides: Partial<GuildCoveredView> = {}): GuildCoveredView {
  const tier = overrides.tier ?? 'guild';
  return {
    tier,
    state: 'active',
    ends_at: '2026-12-01T00:00:00Z',
    billed_by: 'Thornwick',
    you_pay: false,
    you_hold_seat: false,
    seats_given: 0,
    seats_total: PLANS[tier].seats,
    ...overrides,
  };
}

/** A `/v1/me` guild entry on a given subscribe step; the owner's OLYMPUS XXVII (Horde, as production
 *  holds it) unless overridden. The owner's character in it reads its rank and verify mark from here. */
export function stepGuild(overrides: Partial<MeGuild> & { step?: SubscribeStep } = {}): MeGuild {
  const { step = 'ask_leader', ...rest } = overrides;
  return {
    ...OLYMPUS,
    region: 'us',
    ruleset: 'pvp',
    ruleset_label: 'PvP',
    verified: false,
    verified_count: 0,
    crest_url: null,
    faction: 'horde',
    rank: 'member',
    claim: { ...UNCLAIMED },
    verify: { has_approver: false },
    covered: null,
    subscribe: subscribeView(step),
    ...rest,
  };
}

export function meCharacter(
  name: string,
  guildId?: number,
  guild: Partial<NonNullable<MeCharacter['guild']>> = {},
): MeCharacter {
  return {
    key: `us/pvp/${name.toLowerCase()}`,
    region: 'us',
    ruleset: 'pvp',
    name,
    class: 'Hunter',
    ...(guildId === undefined
      ? {}
      : { guild: { id: guildId, name: 'G', rank: 'member', verified: false, ...guild } }),
  };
}

const OWNER_BARE: MeCharacter = {
  key: 'us/pvp/obnoxious-yell',
  region: 'us',
  ruleset: 'pvp',
  name: 'Obnoxious Yell',
  class: 'Warrior',
  spec: 'Fury',
  level: 21,
  faction: 'horde',
};
const OWNER_CHARACTER: MeCharacter = {
  ...OWNER_BARE,
  guild: { id: OLYMPUS.id, name: OLYMPUS.name, rank: 'member', verified: false },
};
/** The owner's character in Employed Millennials, as production holds it. */
const OWNER_SECOND_BARE: MeCharacter = {
  key: 'us/pvp/bow-jackzon',
  region: 'us',
  ruleset: 'pvp',
  name: 'Bow Jackzon',
  class: 'Hunter',
  spec: 'Beast Mastery',
  level: 24,
  faction: 'horde',
};

/** `bare` in `guild`: the character's rank and verify mark are the guild entry's, as production moves them together. */
function inGuild(bare: MeCharacter, guild: MeGuild): MeCharacter {
  return { ...bare, guild: { id: guild.id, name: guild.name, rank: guild.rank, verified: guild.verified } };
}

/**
 * The owner's characters for a set of guild entries: Obnoxious Yell in OLYMPUS XXVII and Bow Jackzon
 * in Employed Millennials, each carrying that entry's rank and verify mark. Without an OLYMPUS XXVII
 * entry the main keeps its production membership there; Bow Jackzon appears with Employed Millennials.
 */
function ownerCharacters(guilds: readonly MeGuild[]): MeCharacter[] {
  const olympus = guilds.find((guild) => guild.id === OLYMPUS.id);
  const millennials = guilds.find((guild) => guild.id === MILLENNIALS.id);
  const main = olympus === undefined ? OWNER_CHARACTER : inGuild(OWNER_BARE, olympus);
  return millennials === undefined ? [main] : [main, inGuild(OWNER_SECOND_BARE, millennials)];
}

export const GRANT_OPEN_ENDED: BillingView = deepFreeze({
  ...NO_PRICE_FACTS,
  source: 'grant',
  plan: 'premium',
  status: 'active',
  current_period_end: null,
  cancel_at_period_end: false,
});

interface MeParts {
  characters?: MeCharacter[];
  guilds?: MeGuild[];
  billing?: BillingView | null;
  personalPortal?: boolean;
  serverSims?: boolean;
  extra?: Partial<Me>;
}

/** True when the parts give the account Premium from any source (a billing row, a seat, or the fallback). */
function hasPremium({ billing, serverSims, extra }: MeParts): boolean {
  return billing != null || serverSims === true || extra?.seat != null;
}

/**
 * A signed-in `Me`: the owner's characters in the given guilds by default, entitlements carrying
 * the given billing row. Any Premium source turns every Premium feature on, as the API does.
 */
export function buildMe({
  characters,
  guilds = [],
  billing = null,
  personalPortal = false,
  serverSims = false,
  extra,
}: MeParts): Me {
  return deepFreeze(
    structuredClone({
      user: { id: 78, battletag: 'hunthrop#1894', email: null, role: 'user', anonymize: false },
      characters: characters ?? ownerCharacters(guilds),
      guilds,
      main_character_key: OWNER_CHARACTER.key,
      entitlements: hasPremium({ billing, serverSims, extra })
        ? premiumEntitlements({ billing, personal_portal: personalPortal })
        : mockEntitlements({ billing, personal_portal: personalPortal, server_sims: serverSims }),
      ...extra,
    }),
  );
}

// --- the owner ----------------------------------------------------------------------------------

/** Capture 1: the owner exactly as production holds the account. */
export const OWNER_MEMBER: Me = buildMe({
  guilds: [stepGuild(), stepGuild({ ...MILLENNIALS })],
  billing: GRANT_OPEN_ENDED,
});

/** Capture 1b: after an addon save that makes the owner leader of OLYMPUS XXVII. */
export const OWNER_GUILD_MASTER: Me = buildMe({
  guilds: [stepGuild({ step: 'claim', rank: 'leader' }), stepGuild({ ...MILLENNIALS })],
  billing: GRANT_OPEN_ENDED,
});

/** The same after the owner claims it: OLYMPUS XXVII is ready with the trial unused. */
export const OWNER_GUILD_MASTER_CLAIMED: Me = buildMe({
  guilds: [
    stepGuild({
      rank: 'leader',
      verified: true,
      verified_count: 1,
      claim: { state: 'claimed', yours: true, pending_yours: false, holder_name: 'Obnoxious Yell' },
      subscribe: subscribeView('ready', { trial_used: false }),
    }),
    stepGuild({ ...MILLENNIALS }),
  ],
  billing: GRANT_OPEN_ENDED,
});

// --- one guild per step -------------------------------------------------------------------------

const OFFICER_CLAIM: GuildClaimView = {
  state: 'claimed',
  yours: false,
  pending_yours: false,
  holder_name: 'Thornwick',
};
const MEMBER_CLAIM: GuildClaimView = { ...OFFICER_CLAIM, holder_name: null };
const PENDING_CLAIM: GuildClaimView = {
  state: 'pending',
  yours: false,
  pending_yours: false,
  holder_name: null,
};

/** The viewer holds the claim, so they are its verified guild master (claiming verifies the claimant). */
function ready(verified: number): MeGuild {
  return stepGuild({
    rank: 'leader',
    verified: true,
    verified_count: verified,
    claim: { state: 'claimed', yours: true, pending_yours: false, holder_name: 'Obnoxious Yell' },
    covered: null,
    subscribe: subscribeView('ready', { trial_used: false }),
  });
}

/** A claimed guild whose plan lapsed: the trial is spent, the card says "Subscribe again". */
function readyLapsed(tier: 'guild' | 'raid_guild'): MeGuild {
  return {
    ...ready(12),
    covered: coveredView({ tier, state: 'lapsed', ends_at: null, billed_by: '' }),
    subscribe: subscribeView('ready', { trial_used: true }),
  };
}

/** An `on_plan` guild; `covered` overrides pick the variant. The payer is the verified guild master
 *  holding the claim; anyone else is a verified officer, so the payer's name may reach them (§8). */
function onPlan(covered: Partial<GuildCoveredView>): MeGuild {
  const youPay = covered.you_pay === true;
  return stepGuild({
    rank: youPay ? 'leader' : 'officer',
    verified: true,
    verified_count: 12,
    claim: youPay ? { ...OFFICER_CLAIM, yours: true, holder_name: OWNER_BARE.name } : OFFICER_CLAIM,
    verify: { has_approver: true },
    covered: coveredView(covered),
    subscribe: subscribeView('on_plan'),
  });
}

/** A claimed guild someone else holds: the holder is verified, so it counts and can approve. */
function claimedByOther(step: SubscribeStep, overrides: Partial<MeGuild> = {}): MeGuild {
  return stepGuild({
    step,
    verified_count: 9,
    claim: MEMBER_CLAIM,
    verify: { has_approver: true },
    ...overrides,
  });
}

const stepFixture = (guild: MeGuild, extra?: Partial<Me>): Me => buildMe({ guilds: [guild], extra });

const SEAT_ROW = {
  guild_name: OLYMPUS.name,
  given_by: 'Thornwick',
  since: '2026-11-10T00:00:00Z',
  starts_at: '2026-11-10T00:00:00Z',
};

/** One `Me` per subscribe step, the guild varying only in what the step needs. */
export const STEP_FIXTURES: Record<SubscribeStep, Me> = {
  on_plan: stepFixture(onPlan({ you_pay: true })),
  checkout_open: stepFixture(
    claimedByOther('checkout_open', { rank: 'officer', verified: true, claim: OFFICER_CLAIM }),
  ),
  frozen: stepFixture(claimedByOther('frozen', { rank: 'officer', verified: true, claim: OFFICER_CLAIM })),
  ready: stepFixture(ready(12)),
  claimed_ask: stepFixture(claimedByOther('claimed_ask')),
  claim_pending_yours: stepFixture(
    stepGuild({
      step: 'claim_pending_yours',
      claim: { ...PENDING_CLAIM, pending_yours: true },
    }),
  ),
  confirm_claim: stepFixture(stepGuild({ step: 'confirm_claim', claim: PENDING_CLAIM })),
  claim_pending_ask: stepFixture(stepGuild({ step: 'claim_pending_ask', claim: PENDING_CLAIM })),
  claim: stepFixture(stepGuild({ step: 'claim', rank: 'leader' })),
  claim_link_bnet: stepFixture(stepGuild({ step: 'claim_link_bnet', rank: 'leader' })),
  start_claim: stepFixture(stepGuild({ step: 'start_claim', rank: 'officer' })),
  claim_blocked: stepFixture(
    stepGuild({
      rank: 'leader',
      subscribe: subscribeView('claim_blocked', {
        reason: 'rate_limited',
        available_at: '2026-11-26T00:00:00Z',
      }),
    }),
  ),
  ask_leader: stepFixture(stepGuild()),
};

/** `claimed_ask` as the officer sees the holder's name, and as a member, who does not. */
export const CLAIMED_ASK_OFFICER: Me = stepFixture(
  claimedByOther('claimed_ask', { rank: 'officer', verified: true, claim: OFFICER_CLAIM }),
);
export const CLAIMED_ASK_MEMBER: Me = stepFixture(claimedByOther('claimed_ask'));

/** Both `claim_blocked` reasons. */
export const CLAIM_BLOCKED_RATE_LIMITED: Me = STEP_FIXTURES.claim_blocked;
export const CLAIM_BLOCKED_HOLDS_OTHER: Me = stepFixture(
  stepGuild({
    rank: 'leader',
    subscribe: subscribeView('claim_blocked', { reason: 'holds_other', other_guild_name: MILLENNIALS.name }),
  }),
);

/** The `on_plan` variants: the viewer pays, holds a seat (so has Premium through it), someone else
 *  pays, every seat is given (a Raid guild's three), and the plan is ending. */
export const ON_PLAN_YOU_PAY: Me = STEP_FIXTURES.on_plan;
export const ON_PLAN_SEAT: Me = stepFixture(
  { ...onPlan({ you_hold_seat: true, seats_given: 1 }), rank: 'member' },
  { seat: SEAT_ROW },
);
export const ON_PLAN_OTHER_PAYS: Me = stepFixture(onPlan({}));
export const ON_PLAN_ALL_SEATS_GIVEN: Me = stepFixture(onPlan({ tier: 'raid_guild', seats_given: 3 }));
export const ON_PLAN_ENDING: Me = stepFixture(
  onPlan({ you_pay: true, state: 'ending', ends_at: '2026-12-15T00:00:00Z' }),
);

/** `ready` on a Raid guild with 12 verified, and on a Guild with 28 (the size line offers Raid guild). */
export const READY_RAID_GUILD_12: Me = stepFixture(ready(12));
export const READY_GUILD_28: Me = stepFixture(ready(28));
/** A lapsed Raid guild plan: the trial is spent and the card says "Subscribe again". */
export const READY_LAPSED: Me = stepFixture(readyLapsed('raid_guild'));

/** One `ready` (trial unused) and one Raid guild `on_plan` where the viewer is a member with seats free. */
export const MIXED: Me = buildMe({
  guilds: [
    ready(12),
    { ...onPlan({ tier: 'raid_guild', seats_given: 1 }), ...MILLENNIALS, rank: 'member', verified: true },
  ],
});

// --- one per premium source ---------------------------------------------------------------------

const OWN_FACTS: BillingFacts = { founder: false, interval: 'monthly', amount_cents: 600 };

function stripeBilling(overrides: Partial<BillingView> = {}): BillingView {
  return {
    ...OWN_FACTS,
    source: 'stripe',
    plan: 'premium',
    status: 'active',
    current_period_end: '2026-12-01T00:00:00Z',
    cancel_at_period_end: false,
    ...overrides,
  };
}

const premiumMe = (billing: BillingView | null, personalPortal = false, extra?: Partial<Me>): Me =>
  buildMe({ billing, personalPortal, extra });

/** The founder yearly price Stripe bills ($40), not the list price. */
const CENTS_PER_DOLLAR = 100;
const FOUNDER_YEARLY_CENTS = PLANS.premium.founderYearly * CENTS_PER_DOLLAR;

export const PREMIUM_OWN_FOUNDER: Me = premiumMe(
  stripeBilling({ founder: true, interval: 'yearly', amount_cents: FOUNDER_YEARLY_CENTS }),
  true,
);
export const PREMIUM_OWN_ENDING: Me = premiumMe(stripeBilling({ cancel_at_period_end: true }), true);
export const PREMIUM_PAST_DUE: Me = premiumMe(stripeBilling({ status: 'past_due' }), true);
export const PREMIUM_OWN_THEN_SEAT: Me = premiumMe(stripeBilling({ cancel_at_period_end: true }), true, {
  seat: { ...SEAT_ROW, starts_at: '2026-12-01T00:00:00Z' },
});
export const PREMIUM_SEAT_WITH_PORTAL: Me = premiumMe(null, true, { seat: SEAT_ROW });
export const PREMIUM_SEAT_NO_PORTAL: Me = premiumMe(null, false, { seat: SEAT_ROW });
export const PREMIUM_GRANT_OPEN_ENDED: Me = premiumMe(GRANT_OPEN_ENDED);
export const PREMIUM_GRANT_UNTIL: Me = premiumMe({
  ...GRANT_OPEN_ENDED,
  current_period_end: '2027-03-31T00:00:00Z',
});
/** Premium that arrives only through the server's fallback (no billing row of any kind). */
export const PREMIUM_FALLBACK: Me = buildMe({ serverSims: true });
export const PREMIUM_FREE: Me = premiumMe(null);

// --- empty and older shapes ---------------------------------------------------------------------

/** No characters at all. */
export const EMPTY_A: Me = buildMe({ characters: [] });
/** Characters, none in a guild. */
export const EMPTY_B: Me = buildMe({ characters: [OWNER_BARE] });
/** An older API: guild entries carry no standing fields, so `subscribe` is absent. */
export const OLDER_ME: Me = buildMe({
  guilds: [
    {
      id: OLYMPUS.id,
      name: OLYMPUS.name,
      region: 'us',
      ruleset: 'pvp',
      rank: 'member',
      verified: false,
      verified_count: 0,
    },
  ],
});
