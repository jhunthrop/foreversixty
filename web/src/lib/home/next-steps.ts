// Pure formatting for the "Next steps" grid's four signed-in cards (spec 2026-09-24 §2.3).
// Each function decides what one card's live element says; HomeNextSteps.svelte decides
// only when to show it (loading/ready/empty), so the two Svelte components that need one of
// these facts (the hero's rating figure, this grid's rating card) can never format the same
// number two different ways.
import { headlineOf, kindOf } from '../sim/history';
import { simCopy } from '../sim/copy';
import { MAX_POINTS } from '../planner/types';
import { talentPointsFromString } from '../sim/character';
import { formatDate } from '../dates';
import type { SimListRow } from '../sim/types';
import type { MyReport } from '../account/api';
import type { CharacterRating } from '../rating/types';

/** "Top Gear · +41 DPS from Vis'kag" — the visitor's own most recent saved sim. Not claimed
 *  to be for one character: GET /v1/sims?mine=1 carries no character key per row (Ruling 3,
 *  2026-09-24-home-v2.md), only spec/dps/title, so this names the sim rather than a
 *  character match the API cannot confirm. */
export function simCardLine(row: SimListRow): string {
  return `${simCopy.kindLabel[kindOf(row)]} · ${headlineOf(row)}`;
}

/** "24 of 51 points" from a talent split string ("31/0/20", GET .../sim-input's `talents`
 *  field) via the same digit-sum function CharacterStrip.svelte/SavedSim.svelte already use
 *  for a saved sim's own "51 points" figure (Ruling 2). Empty when the character has no
 *  recorded talents yet -- the caller's cue to show the empty state instead. */
export function plannerPointsLabel(talents: string): string {
  if (talents === '') return '';
  return `${talentPointsFromString(talents)} of ${MAX_POINTS} points`;
}

/** "<title>, <date>" for the visitor's latest report. */
export function logsCardLine(report: MyReport): string {
  return `${report.title}, ${formatDate(new Date(report.created_at))}`;
}

/** The hero's rating figure, two decimals -- empty until a real sample exists, matching
 *  HomeAccountPanel's own pre-existing "never shown until it resolves with a real sample"
 *  rule (spec 2026-09-23 §2 item 2), now the one place both the hero and this grid's
 *  Rankings card read it from. */
export function ratingCardValue(rating: CharacterRating | null): string {
  if (rating === null || rating.sample_size === 0 || rating.latest === null) return '';
  return rating.latest.overall.toFixed(2);
}
