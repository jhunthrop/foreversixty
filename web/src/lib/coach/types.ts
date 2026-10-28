// web/src/lib/coach/types.ts
// The coach API's answers, mirroring the Go DTOs of Tasks 30 and 31 (api/internal/coach/dto.go).
// api.ts normalises every answer into these shapes, so nothing downstream reads a raw field.
// states.ts maps an answer onto the states the experience spec's section 5.1 names.

export type CoachStatus = 'ready' | 'writing' | 'preparing' | 'not_reviewable';
export type CoachSource = 'model' | 'rules';

export interface CoachTeaserFirst {
  title: string;
  /** The finding type, e.g. `core_casts`. */
  code: string;
  /** The pack's first anchor spell, for the icon; null when the finding names no spell. */
  spell_id: number | null;
}

export interface CoachTeaser {
  status: CoachStatus;
  access: 'teaser';
  finding_count: number;
  fix_first: CoachTeaserFirst | null;
}

export type CoachAnchorKind = 'cast' | 'aura' | 'death' | 'resource' | 'time' | 'buffs' | 'healing';

export interface CoachAnchor {
  kind: CoachAnchorKind;
  spell_id?: number;
  at_ms?: number;
  to_ms?: number;
  death?: string;
  /** Sent by the API when an aura anchor is a debuff; absent means the Buffs tab. */
  debuff?: boolean;
}

/** A computed link between two findings, as the rule-only rendering carries it. */
export interface CoachLink {
  id: string;
  sentence: string | null;
}

export interface FixItem {
  finding: string;
  title: string;
  why: string;
  drill: string;
  drillText: string;
  anchors: string[];
  links: CoachLink[];
}

export interface ThenItem {
  finding: string;
  title: string;
  why: string;
  anchors: string[];
  links: CoachLink[];
}

export interface WellItem {
  strength: string;
  text: string;
  anchors: string[];
}

export interface CoachReview {
  opening: string;
  fixFirst: FixItem | null;
  thenFix: ThenItem[];
  doingWell: WellItem[];
  couldntCheck: string[];
}

export interface CoachFull {
  report_id: string;
  fight_index: number;
  player_key: string;
  spec: string;
  role: string;
  size: number;
  kill: boolean;
  status: CoachStatus;
  source: CoachSource;
  busy: boolean;
  access: 'full' | 'officer';
  persona: { label: string; status: 'full' | 'starter' };
  reviewed_at: string;
  review: CoachReview;
  anchors: Record<string, CoachAnchor>;
  basis_labels: string[];
}

export interface LastRaid {
  zone: string;
  fought_at: string;
  report_id: string;
  finding_count: number;
  first: { stem: string; on: number; of: number } | null;
  url: string;
}

export interface CoachHistoryItem {
  report_id: string;
  fight_index: number;
  player_key: string;
  spec: string;
  boss: string;
  kill: boolean;
  source: CoachSource;
  generated_at: string;
  fix_first: { title: string } | null;
}

export interface CoachRecent {
  last_raid: LastRaid | null;
  pending: boolean;
  reviews: CoachHistoryItem[];
}

export interface CoachHistoryPage {
  reviews: CoachHistoryItem[];
  /** The keyset cursor for the next page; null on the last. */
  next_cursor: string | null;
}

export interface NightLine {
  stem: string;
  on: number;
  of: number;
  finding_count: number;
  fight_index: number;
}
