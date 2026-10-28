// The wire shapes of `GET /v1/me`: pure types, kept out of api.ts so the request code there
// stays under the file cap. api.ts re-exports every name, so importers are unchanged.
import type {
  BillingFacts,
  GuildClaimView,
  GuildCoveredView,
  GuildSubscribeView,
  GuildTier,
  GuildVerifyView,
  MeEndedSeat,
  MeSeat,
  MeSeatOffer,
} from '../billing/contract';

/**
 * The standing fields `GET /v1/me` adds to each guild entry (spec
 * design/specs/2026-10-27-subscribe-and-join.md section 8). Every one is optional: an older
 * API sends none, and `normalizeMeGuild` turns that into "unknown" (`subscribe: null`), which
 * the web draws as an identity-only card.
 */
export interface MeGuildStanding {
  /** The guild page's slug; absent on an older API, so the hand-off link is hidden then. */
  slug?: string;
  /** The ruleset's display label (api auth.Guild.RulesetLabel). */
  ruleset_label?: string;
  crest_url?: string | null;
  faction?: GuildFaction | null;
  roster_count?: number | null;
  claim?: GuildClaimView | null;
  verify?: GuildVerifyView | null;
  covered?: GuildCoveredView | null;
  subscribe?: GuildSubscribeView | null;
}

export const GUILD_FACTIONS = ['alliance', 'horde'] as const;
export type GuildFaction = (typeof GUILD_FACTIONS)[number];

/**
 * A character's own guild line (spec 2026-09-22 §6): distinct from `MeGuild`, the
 * account-level "My guilds" entry below -- this is what `GET /v1/me`'s `characters[].guild`
 * carries, with no `region`/`ruleset`/`consent`/`plan` of its own.
 */
export interface MeCharacterGuild {
  id: number;
  name: string;
  rank?: string;
  rank_index?: number;
  verified: boolean;
  /** Raid mode only: raids this character attended of raids signed up for or attended. */
  attendance?: { present: number; nights: number };
}

export interface MeCharacter {
  key: string;
  region: string;
  ruleset: string;
  name: string;
  class?: string;
  /** Spec 2026-09-22 §6: omitted when the Battle.net import/refresh has not learned it yet. */
  realm?: string;
  level?: number;
  faction?: 'alliance' | 'horde';
  /** "Night Elf"; from the Battle.net import, omitted until it has run. */
  race?: string;
  gender?: 'male' | 'female';
  /** Equipped item level, from the Battle.net character profile when it answered. */
  item_level?: number;
  /** Blizzard's own render-CDN images (design/DESIGN-SYSTEM.md's "Character row"/hero
   *  band); omitted when the Battle.net import has not captured them. */
  avatar_url?: string;
  render_url?: string;
  /** `'bnet'` or `'export'` -- which path last wrote this row. Optional: older/stubbed
   *  fixtures written before the Battle.net import landed carry no such field. */
  source?: string;
  /** The site's newest export for this character, when it has one (spec 2026-09-22 §5):
   *  `'addon'` for a paste or companion push, `'blizzard'` for one built from the
   *  Battle.net profile. Omitted exactly when the site holds no build for the character --
   *  never inferred from anything else on this row.
   *
   *  Home rebuild spec §3.B: `gear`/`talents`/`level`/`data_build` are the addon export's
   *  own worn-gear and talent snapshot (the API lane's extension to this contract, landed
   *  alongside this lane) -- all optional, and omitted together on a `build` row the addon
   *  export predates (an older paste, or a Battle.net-only `build`, which never carries
   *  worn items at all). Every reader treats a missing `gear`/`talents` as "not synced yet",
   *  never an empty set. */
  build?: {
    source: 'addon' | 'blizzard';
    captured_at: string;
    /** Export slot name (`addon/ForeverSixty/Export.lua`'s own `INVENTORY_SLOTS`, identical
     *  vocabulary to `planner/types.ts`'s `SLOTS` -- no mapping needed, see that file's
     *  `Slot` type) -> the worn item's id. Omitted entirely for a slot with nothing
     *  equipped, the same "absent means empty" convention `Gear` (`planner/types.ts`)
     *  already uses. */
    gear?: Record<string, number>;
    /** Three per-tree rank strings in tree order (`"503200000"`, one digit per talent
     *  cell, the identical convention a leveling-BiS band's own `talents` string uses per
     *  segment -- `lib/home/talent-delta.ts` compares the two directly) and the point sum
     *  per tree. */
    talents?: { trees: string[]; points: number[] };
    /** The character's level at the moment this build was captured -- distinct from the
     *  character row's own top-level `level` (Battle.net's, refreshed on its own cadence);
     *  this one is the addon's own snapshot, so a gear/talent comparison against a BiS band
     *  can use the figure that actually matches the equipped items above. */
    level?: number;
    /** The game build the addon captured this snapshot against (e.g. `"1.60.1.70009"`),
     *  for a future reader that needs to know which item/talent table the ids above resolve
     *  against; this lane's own BiS/item lookups use the site's own active build
     *  (`data/active-build.json`) instead, the same build every other BiS comparison on the
     *  site reads against. */
    data_build?: string;
    /** The player's own median gap between syncs, in seconds, from the sync history; null
     *  with fewer than three syncs on record. The selector's stale rule reads it
     *  (`lib/character-selector/stale.ts`). */
    median_sync_gap_sec?: number | null;
    /** Why the last Battle.net refresh failed; null or omitted when it did not. Non-null
     *  is what makes the selector's row "failed" with a Retry. */
    sync_error?: string | null;
  };
  /** Omitted when the character has no `guild_characters` row at all. */
  guild?: MeCharacterGuild;
  /**
   * The talent spec or primary tree name ("Marksmanship", "Fury", "Holy") -- the API lane's
   * own extension to this contract, landed alongside this one (home rebuild spec §3.B):
   * `level`, `race`, `faction` and `realm` now arrive for an addon-only character too, not
   * only one that has been through the Battle.net import, but `spec` can still be absent
   * (an addon export with no learned talent, or a character the addon has not captured a
   * build for at all) -- every reader treats a missing value as "not known yet" and omits
   * the clause, never a placeholder string.
   */
  spec?: string;
}

export interface MeGuild extends MeGuildStanding {
  id: number;
  region: string;
  ruleset: string;
  name: string;
  rank?: string;
  /** Spec section 2.6: GET /v1/me gains Consent and Verified per Guild entry. */
  consent?: 'roster' | 'gear' | 'gear_bags';
  verified: boolean;
  verified_count: number;
  /**
   * Non-null only for a verified officer/leader of a guild with an active guild-plan
   * entitlement (spec 1.4) -- billing detail, gated tighter than membership alone. Optional
   * because the live API does not send it yet (the API lane has not landed); every reader
   * treats a missing key the same as null.
   */
  plan?: GuildBillingView | null;
  /**
   * What the account may do in this guild (`approve`, `remove`, `settings`, `award`,
   * `billing`, `release`), from the API's one rule (api `guilds.Capabilities`). The page draws
   * a control exactly when its word is here, never from the rank. Optional so a response from
   * before it reads as an empty list.
   */
  can?: string[];
}

/** Spec 1.4: every Can() feature pre-resolved for the caller, plus their own billing state. */
export interface EntitlementsView {
  /** Absent until the API sends it. */
  personal_portal?: boolean;
  server_sims: boolean;
  retention: boolean;
  multi_compare: boolean;
  history: boolean;
  notifications: boolean;
  officer_views: boolean;
  roster_check: boolean;
  supporter_mark: boolean;
  billing: BillingView | null;
}

export interface BillingView extends BillingFacts {
  /** Absent until the API sends it; a live Stripe row with none reads as stripe. */
  source?: 'stripe' | 'grant';
  plan: string;
  status: string;
  current_period_end: string | null;
  cancel_at_period_end: boolean;
}

export interface GuildBillingView {
  tier: GuildTier;
  you_hold_seat: boolean;
  /** `officer_tools` is a scoped grant for Award, Undo and Nudge only; `plan` is the guild plan. */
  scope: 'plan' | 'officer_tools';
  status: string;
  current_period_end: string | null;
  cancel_at_period_end: boolean;
  /** Empty for a plan granted through the CLI, never through Stripe -- there is no billing
   *  contact to name (api/internal/auth/handler.go's attachGuildPlans). A reader must not
   *  show "Billed by" with nothing after it. */
  billed_by: string;
  you_are_billing_contact: boolean;
}

/** A server run that finished while the member was away, from GET /v1/me. */
export interface FinishedSim {
  sim_id: string;
  spec: string;
  title: string;
  dps: number;
  finished_at: string;
}

export interface Me {
  user: {
    id: number;
    battletag: string | null;
    email: string | null;
    role: string;
    anonymize: boolean;
    /**
     * The premium flag, set by hand until payments exist. Optional and kept only for
     * backward compatibility with e2e fixtures written before the entitlements API landed
     * (web lane ruling A, docs/superpowers/plans/2026-09-21-pay-web.md) -- new code reads
     * `entitlements.server_sims` via `effectiveServerSims`, never this field directly.
     */
    premium?: boolean;
  };
  characters: MeCharacter[];
  guilds: MeGuild[];
  /** Optional: absent until the API lane ships spec 1.4's block. See effectiveServerSims. */
  entitlements?: EntitlementsView;
  /** RFC3339; spec 2026-09-22 §6, omitted when this account has never imported from
   *  Battle.net. */
  bnet_imported_at?: string;
  /** The account's chosen main character's key; omitted until chosen. Every other
   *  character is an alt. */
  main_character_key?: string;
  /** Unseen finished server runs, newest first, at most 3. Absent on an older API. */
  finished_sims?: FinishedSim[];
  seat?: MeSeat | null;
  seat_offer?: MeSeatOffer | null;
  ended_seat?: MeEndedSeat | null;
  /** Absent on an older API. */
  server_sims?: { runs_this_month: number };
  /** Keyed by channel; `discord` is null when nothing is saved. Absent on an older API. */
  notifications?: { connect?: boolean; discord: DiscordNotifications | null };
  /** The deployment has the Discord bot on: the account page draws the link row. Absent on an older API. */
  discord_link_available?: boolean;
  /** The Discord account this account is linked to; null or absent when unlinked. */
  discord_link?: { display_name: string } | null;
  /** The name of a guild of this account that signs up in Discord, when the account is unlinked. */
  discord_link_suggested?: string | null;
}

/** Why posts stopped while a webhook is kept (spec 2026-10-25-discord-notifications §10). */
export type DiscordDisabledReason = 'failed five times' | 'plan lapsed';

/** The three event flags a member can choose; the first two are what the Premium row counts. */
export interface DiscordEvents {
  sim_finished: boolean;
  upload_parsed: boolean;
  /** Whether a log whose visibility is private is posted too (needs `upload_parsed`). */
  upload_private: boolean;
}

/** How the webhook was saved: pasted by the member, or created by the Discord connect flow. */
export type DiscordSource = 'pasted' | 'connected';

/** The saved webhook as `/v1/me` and `PUT /v1/notifications/discord` answer it: never the URL. */
export interface DiscordNotifications {
  channel_id: string;
  /** Missing on an older API, which reads as `pasted`. */
  source?: DiscordSource;
  /** The Discord server's id when known (the `Open in Discord` link); null otherwise. */
  guild_id?: string | null;
  /** Discord's own name for the webhook; null when Discord returned none. Member-controlled text. */
  name: string | null;
  events: DiscordEvents;
  disabled_reason: DiscordDisabledReason | null;
  /** Discord's own error text with any URL or token removed. */
  last_error: string | null;
  /** Consecutive failed posts, 0 to 5. */
  failures: number;
}
