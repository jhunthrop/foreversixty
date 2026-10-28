// web/src/lib/coach/types.ts
// The coach API's answers, mirroring the Go DTOs of Tasks 30 and 31 (api/internal/coach/dto.go).
// api.ts normalises every answer into these shapes, so nothing downstream reads a raw field.
// states.ts maps an answer onto the states the experience spec's section 5.1 names.

export type CoachStatus = 'ready' | 'writing' | 'preparing' | 'not_reviewable';
export type CoachSource = 'model' | 'rules';

export interface CoachTeaserFirst {
  readonly title: string;
  /** The finding type, e.g. `core_casts`. */
  readonly code: string;
  /** The pack's first anchor spell, for the icon; null when the finding names no spell. */
  readonly spell_id: number | null;
}

export interface CoachTeaser {
  readonly status: CoachStatus;
  readonly access: 'teaser';
  readonly finding_count: number;
  readonly fix_first: CoachTeaserFirst | null;
  /** The reviewed player's spec, when the API sends it (it does not for older answers). */
  readonly spec?: string;
}

export type CoachAnchorKind = 'cast' | 'aura' | 'death' | 'resource' | 'time' | 'buffs' | 'healing';

export interface CoachAnchor {
  readonly kind: CoachAnchorKind;
  readonly spell_id?: number;
  /** The build's name for `spell_id` (the class spell data), which the web's name tables may lack. */
  readonly spell_name?: string;
  readonly at_ms?: number;
  readonly to_ms?: number;
  readonly death?: string;
  /** Sent by the API when an aura anchor is a debuff; absent means the Buffs tab. */
  readonly debuff?: boolean;
  /** A debuff aura anchor's NPC target (never a player): the AuraTable row's key. */
  readonly target_guid?: string;
}

/** A computed link between two findings, as the rule-only rendering carries it. */
export interface CoachLink {
  readonly id: string;
  readonly sentence: string | null;
}

export interface FixItem {
  readonly finding: string;
  readonly title: string;
  readonly why: string;
  readonly drill: string;
  readonly drillText: string;
  readonly anchors: readonly string[];
  readonly links: readonly CoachLink[];
}

export interface ThenItem {
  readonly finding: string;
  readonly title: string;
  readonly why: string;
  readonly anchors: readonly string[];
  readonly links: readonly CoachLink[];
}

export interface WellItem {
  readonly strength: string;
  readonly text: string;
  readonly anchors: readonly string[];
}

export interface CoachReview {
  readonly opening: string;
  readonly fixFirst: FixItem | null;
  readonly thenFix: readonly ThenItem[];
  readonly doingWell: readonly WellItem[];
  readonly couldntCheck: readonly string[];
}

/** An account's model budget is spent for the day or the month; `until` is the UTC day it lifts. */
export interface CoachCap {
  readonly period: 'day' | 'month';
  /** `YYYY-MM-DD`. */
  readonly until: string;
}

export interface CoachFull {
  readonly report_id: string;
  readonly fight_index: number;
  readonly player_key: string;
  readonly spec: string;
  readonly role: string;
  readonly size: number;
  readonly kill: boolean;
  readonly status: CoachStatus;
  readonly source: CoachSource;
  readonly busy: boolean;
  /** The account cap behind `busy`, when it is the account's own; null for the site's cap. */
  readonly cap: CoachCap | null;
  readonly access: 'full' | 'officer';
  readonly persona: { readonly label: string; readonly status: 'full' | 'starter' };
  readonly reviewed_at: string;
  readonly review: CoachReview;
  readonly anchors: Readonly<Record<string, CoachAnchor>>;
  readonly basis_labels: readonly string[];
}

export interface LastRaid {
  readonly zone: string;
  readonly fought_at: string;
  readonly report_id: string;
  readonly finding_count: number;
  readonly first: { readonly stem: string; readonly on: number; readonly of: number } | null;
  readonly url: string;
}

export interface CoachHistoryItem {
  readonly report_id: string;
  readonly fight_index: number;
  readonly player_key: string;
  readonly spec: string;
  readonly boss: string;
  readonly kill: boolean;
  readonly source: CoachSource;
  readonly generated_at: string;
  readonly fix_first: { readonly title: string } | null;
}

export interface CoachRecent {
  readonly last_raid: LastRaid | null;
  readonly pending: boolean;
  readonly reviews: readonly CoachHistoryItem[];
}

export interface CoachHistoryPage {
  readonly reviews: readonly CoachHistoryItem[];
  /** The keyset cursor for the next page; null on the last. */
  readonly next_cursor: string | null;
}

export interface NightLine {
  readonly stem: string;
  readonly on: number;
  readonly of: number;
  readonly finding_count: number;
  readonly fight_index: number;
}
