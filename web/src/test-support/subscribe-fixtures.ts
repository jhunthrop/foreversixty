// web/src/test-support/subscribe-fixtures.ts
// Typed `/v1/me` fixtures for subscribe-and-join (spec design/specs/2026-10-27-subscribe-and-join.md
// section 8): the owner's account as production holds it, one fixture per subscribe step, one per
// premium source, and the empty and older-API shapes. Each is a full `Me`, so `npm run check` proves
// it against the contract; e2e stubs it with `routeMe`.
import type { Page } from '@playwright/test';
import type { BillingView, Me, MeCharacter, MeGuild } from '../lib/account/me-types';
import type {
  BillingFacts,
  GuildClaimView,
  GuildCoveredView,
  GuildSubscribeView,
  SubscribeStep,
} from '../lib/billing/contract';
import { NO_PRICE_FACTS, mockEntitlements } from './billing-fixtures';

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

export function coveredView(overrides: Partial<GuildCoveredView> = {}): GuildCoveredView {
  return {
    tier: 'guild',
    state: 'active',
    ends_at: '2026-12-01T00:00:00Z',
    billed_by: 'Thornwick',
    you_pay: false,
    you_hold_seat: false,
    seats_given: 0,
    seats_total: 3,
    ...overrides,
  };
}

/** A `/v1/me` guild entry on a given subscribe step; the owner's OLYMPUS XXVII unless overridden. */
export function stepGuild(overrides: Partial<MeGuild> & { step?: SubscribeStep } = {}): MeGuild {
  const { step = 'ask_leader', ...rest } = overrides;
  return {
    ...OLYMPUS,
    region: 'us',
    ruleset: 'forever',
    ruleset_label: 'Forever',
    verified: false,
    verified_count: 0,
    crest_url: null,
    rank: 'member',
    claim: UNCLAIMED,
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
    key: `us/forever/${name.toLowerCase()}`,
    region: 'us',
    ruleset: 'forever',
    name,
    class: 'Hunter',
    ...(guildId === undefined
      ? {}
      : { guild: { id: guildId, name: 'G', rank: 'member', verified: false, ...guild } }),
  };
}

const OWNER_CHARACTER: MeCharacter = {
  key: 'us/forever/obnoxious-yell',
  region: 'us',
  ruleset: 'forever',
  name: 'Obnoxious Yell',
  class: 'Warrior',
  spec: 'Fury',
  level: 21,
  faction: 'horde',
  guild: { id: OLYMPUS.id, name: OLYMPUS.name, rank: 'member', verified: false },
};

export const GRANT_OPEN_ENDED: BillingView = {
  ...NO_PRICE_FACTS,
  source: 'grant',
  plan: 'premium',
  status: 'active',
  current_period_end: null,
  cancel_at_period_end: false,
};

interface MeParts {
  characters?: MeCharacter[];
  guilds?: MeGuild[];
  billing?: BillingView | null;
  personalPortal?: boolean;
  extra?: Partial<Me>;
}

/** A signed-in `Me`: the owner's character by default, entitlements carrying the given billing row. */
export function buildMe({
  characters,
  guilds = [],
  billing = null,
  personalPortal = false,
  extra,
}: MeParts): Me {
  return {
    user: { id: 78, battletag: 'hunthrop#1894', email: null, role: 'user', anonymize: false },
    characters: characters ?? [OWNER_CHARACTER],
    guilds,
    main_character_key: OWNER_CHARACTER.key,
    entitlements: mockEntitlements({ billing, personal_portal: personalPortal }),
    ...extra,
  };
}

// --- the owner ----------------------------------------------------------------------------------

/** Capture 1: the owner exactly as production holds the account. */
export const OWNER_MEMBER: Me = buildMe({
  guilds: [stepGuild(), stepGuild({ ...MILLENNIALS })],
  billing: GRANT_OPEN_ENDED,
});

/** Capture 1b: after an addon save that makes the owner guild master of OLYMPUS XXVII. */
export const OWNER_GUILD_MASTER: Me = buildMe({
  guilds: [stepGuild({ step: 'claim', rank: 'guild master' }), stepGuild({ ...MILLENNIALS })],
  billing: GRANT_OPEN_ENDED,
});

/** The same after the owner claims it: OLYMPUS XXVII is ready with the trial unused. */
export const OWNER_GUILD_MASTER_CLAIMED: Me = buildMe({
  guilds: [
    stepGuild({
      rank: 'guild master',
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

function ready(verified: number, tier: 'guild' | 'raid_guild'): MeGuild {
  return stepGuild({
    rank: 'officer',
    verified: true,
    verified_count: verified,
    claim: { state: 'claimed', yours: true, pending_yours: false, holder_name: 'Obnoxious Yell' },
    covered: coveredView({ tier, state: 'lapsed', ends_at: null, billed_by: '' }),
    subscribe: subscribeView('ready', { trial_used: false }),
  });
}

/** An `on_plan` guild; `covered` overrides pick the variant. */
function onPlan(covered: Partial<GuildCoveredView>): MeGuild {
  return stepGuild({
    rank: 'officer',
    verified: true,
    verified_count: 12,
    claim: OFFICER_CLAIM,
    covered: coveredView(covered),
    subscribe: subscribeView('on_plan'),
  });
}

const stepFixture = (guild: MeGuild): Me => buildMe({ guilds: [guild] });

/** One `Me` per subscribe step, the guild varying only in what the step needs. */
export const STEP_FIXTURES: Record<SubscribeStep, Me> = {
  on_plan: stepFixture(onPlan({ you_pay: true })),
  checkout_open: stepFixture(stepGuild({ step: 'checkout_open', rank: 'officer', verified: true })),
  frozen: stepFixture(stepGuild({ step: 'frozen', rank: 'officer', verified: true, claim: OFFICER_CLAIM })),
  ready: stepFixture(ready(12, 'guild')),
  claimed_ask: stepFixture(stepGuild({ step: 'claimed_ask', claim: OFFICER_CLAIM, rank: 'officer' })),
  claim_pending_yours: stepFixture(
    stepGuild({
      step: 'claim_pending_yours',
      claim: { ...PENDING_CLAIM, pending_yours: true },
    }),
  ),
  confirm_claim: stepFixture(
    stepGuild({ step: 'confirm_claim', claim: { ...PENDING_CLAIM, holder_name: 'Thornwick' } }),
  ),
  claim_pending_ask: stepFixture(stepGuild({ step: 'claim_pending_ask', claim: PENDING_CLAIM })),
  claim: stepFixture(stepGuild({ step: 'claim', rank: 'guild master' })),
  claim_link_bnet: stepFixture(stepGuild({ step: 'claim_link_bnet', rank: 'guild master' })),
  start_claim: stepFixture(stepGuild({ step: 'start_claim', rank: 'officer' })),
  claim_blocked: stepFixture(
    stepGuild({
      rank: 'guild master',
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
  stepGuild({ step: 'claimed_ask', rank: 'officer', verified: true, claim: OFFICER_CLAIM }),
);
export const CLAIMED_ASK_MEMBER: Me = stepFixture(stepGuild({ step: 'claimed_ask', claim: MEMBER_CLAIM }));

/** Both `claim_blocked` reasons. */
export const CLAIM_BLOCKED_RATE_LIMITED: Me = STEP_FIXTURES.claim_blocked;
export const CLAIM_BLOCKED_HOLDS_OTHER: Me = stepFixture(
  stepGuild({
    rank: 'guild master',
    subscribe: subscribeView('claim_blocked', { reason: 'holds_other', other_guild_name: MILLENNIALS.name }),
  }),
);

/** The four `on_plan` variants: the viewer pays, holds a seat, someone else pays, and the plan is ending. */
export const ON_PLAN_YOU_PAY: Me = STEP_FIXTURES.on_plan;
export const ON_PLAN_SEAT: Me = stepFixture(onPlan({ you_hold_seat: true, seats_given: 1 }));
export const ON_PLAN_OTHER_PAYS: Me = stepFixture(onPlan({}));
export const ON_PLAN_ENDING: Me = stepFixture(
  onPlan({ you_pay: true, state: 'ending', ends_at: '2026-12-15T00:00:00Z' }),
);

/** `ready` on a Raid guild with 12 verified, and on a Guild with 28 (the size line offers Raid guild). */
export const READY_RAID_GUILD_12: Me = stepFixture(ready(12, 'raid_guild'));
export const READY_GUILD_28: Me = stepFixture(ready(28, 'guild'));

/** One `ready` (trial unused) and one `on_plan` where the viewer is a member with seats free. */
export const MIXED: Me = buildMe({
  guilds: [
    ready(12, 'guild'),
    { ...onPlan({ seats_given: 1 }), ...MILLENNIALS, rank: 'member', verified: true },
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

const SEAT_ROW = {
  guild_name: OLYMPUS.name,
  given_by: 'Thornwick',
  since: '2026-11-10T00:00:00Z',
  starts_at: '2026-11-10T00:00:00Z',
};

export const PREMIUM_OWN_FOUNDER: Me = premiumMe(
  stripeBilling({ founder: true, interval: 'yearly', amount_cents: 5000 }),
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
export const PREMIUM_FALLBACK: Me = buildMe({
  extra: {
    user: { id: 78, battletag: 'hunthrop#1894', email: null, role: 'user', anonymize: false, premium: true },
  },
});
export const PREMIUM_FREE: Me = premiumMe(null);

// --- empty and older shapes ---------------------------------------------------------------------

/** No characters at all. */
export const EMPTY_A: Me = buildMe({ characters: [] });
/** Characters, none in a guild. */
export const EMPTY_B: Me = buildMe({ characters: [meCharacter('Obnoxious Yell')] });
/** An older API: guild entries carry no standing fields, so `subscribe` is absent. */
export const OLDER_ME: Me = buildMe({
  guilds: [
    { ...OLYMPUS, region: 'us', ruleset: 'forever', rank: 'member', verified: false, verified_count: 0 },
  ],
});

// --- e2e ----------------------------------------------------------------------------------------

/** Stub `GET /v1/me` with a fixture, in the API's envelope (the pattern of support/guild-billing.ts). */
export async function routeMe(page: Page, me: Me): Promise<void> {
  await page.route('**/v1/me', (route) =>
    route.fulfill({
      status: 200,
      contentType: 'application/json',
      body: JSON.stringify({ ok: true, data: me, error: null, request_id: 'req-test' }),
    }),
  );
}
