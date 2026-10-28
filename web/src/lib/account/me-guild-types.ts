// The guild standing and billing-source fields `GET /v1/me` adds (spec
// design/specs/2026-10-27-subscribe-and-join.md section 8). Kept out of api.ts, which is at
// its line cap; `MeGuild`, `EntitlementsView` and `BillingView` extend these.
import type {
  GuildClaimView,
  GuildCoveredView,
  GuildSubscribeView,
  GuildVerifyView,
} from '../billing/contract';

/**
 * Every field is optional: an older API sends none of them, and `normalizeMeGuild` turns
 * that into "unknown" (`subscribe: null`), which the web draws as an identity-only card.
 */
export interface MeGuildStanding {
  /**
   * The ruleset's display label. The API shipped it as `realm`; a review is renaming it
   * `ruleset_label`. Both are typed and `normalizeMeGuild` reads either into `ruleset_label`.
   */
  ruleset_label?: string;
  realm?: string;
  crest_url?: string | null;
  faction?: 'alliance' | 'horde' | null;
  roster_count?: number | null;
  claim?: GuildClaimView | null;
  verify?: GuildVerifyView | null;
  covered?: GuildCoveredView | null;
  subscribe?: GuildSubscribeView | null;
}

/** `entitlements.billing.source`: absent until the API sends it (Task 6). */
export interface BillingSourceHint {
  source?: 'stripe' | 'grant';
}

/** `entitlements.personal_portal`: absent until the API sends it (Task 6). */
export interface PersonalPortalHint {
  personal_portal?: boolean;
}
