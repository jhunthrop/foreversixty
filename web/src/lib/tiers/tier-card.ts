// web/src/lib/tiers/tier-card.ts
// The tier list's link preview, as data (lib/og-tier.ts paints it): the page's own list on the
// default faction, grouped S to D on the damage list and ranked on the tank and healer lists, with
// the page's meta line and its Updated stamp. Pure, so the card's words are tested without a render.
import { createHash } from 'node:crypto';
import type { BisRole } from '../bis/types';
import { updatedLabel } from '../dates';
import { DEFAULT_FACTION } from './tier-controls';
import { FACTION_LABELS, ROLE_TABS, formatMetric, specFullName, tiersCopy } from './tier-copy';
import { groupByTier, roleHasTiers, type TierRow } from './tier-list';
import type { TierData } from './tier-load';
import type { TierLetter } from './tier-rules';

export interface TierCardSpec {
  /** "Fury Warrior" */
  name: string;
  classSlug: string;
  /** The row's rank as the page prints it ("2", "=2"); the damage list leads with its letter instead. */
  rank: string;
  /** The sorted figure with its column's unit ("812.3 HPS"), on a ranked list only. */
  figure?: string;
}

export interface TierCardGroup {
  /** S to D on the damage list; `null` on the tank and healer lists, which carry no letters. */
  letter: TierLetter | null;
  specs: TierCardSpec[];
}

export interface TierCardModel {
  title: string;
  /** "Damage · Alliance": the page's role tab and the faction its numbers are for. */
  role: string;
  /** "Best in slot, raid buffs · vs a level 63 boss · 3 min · Updated 12 Nov" */
  meta: string;
  groups: TierCardGroup[];
}

/** The hash length of the card's cache-busting query: enough to never collide across deploys. */
const VERSION_LENGTH = 10;

/** The figure column's head, which names its unit ("HPS", "Taken / s"). */
const FIGURE_HEAD = 4;

function cardSpec(row: TierRow): TierCardSpec {
  return { name: specFullName(row), classSlug: row.classSlug, rank: tiersCopy.rankLabel(row) };
}

function rankedSpec(row: TierRow): TierCardSpec {
  const unit = tiersCopy.heads[row.role][FIGURE_HEAD];
  return { ...cardSpec(row), figure: `${formatMetric(row.role, row.metric)} ${unit}` };
}

export function tierCardModel(data: TierData, role: BisRole): TierCardModel {
  const rows = data.lists[DEFAULT_FACTION][role];
  const groups: TierCardGroup[] = roleHasTiers(role)
    ? groupByTier(rows).map((group) => ({ letter: group.letter, specs: group.rows.map(cardSpec) }))
    : [{ letter: null, specs: rows.map(rankedSpec) }];
  const meta = tiersCopy.metaLine({
    role,
    healerBasis: data.healerBasis,
    healSeconds: data.healSeconds,
    raidFightSeconds: data.raidFightSeconds,
  });
  const updated = data.updated === null ? '' : updatedLabel(data.updated);
  return {
    title: tiersCopy.title,
    role: `${ROLE_TABS[role]} · ${FACTION_LABELS[DEFAULT_FACTION]}`,
    meta: updated === '' ? meta : `${meta} · ${updated}`,
    groups,
  };
}

/** The card's version: it changes whenever anything the card draws changes, so the page's
 *  og:image URL changes with it and Reddit and Discord fetch the new card. */
export function tierCardVersion(model: TierCardModel): string {
  return createHash('sha256').update(JSON.stringify(model)).digest('hex').slice(0, VERSION_LENGTH);
}
