// web/src/lib/tiers/tier-card.ts
// The tier list's link preview, as data (lib/og-tier.ts paints it): the page's own list on the
// default faction, grouped S to D on the damage list and ranked on the tank and healer lists, with
// the page's meta line and its Updated stamp. Pure, so the card's words are tested without a render.
import { createHash } from 'node:crypto';
import type { BisRole } from '../bis/types';
import { updatedLabel } from '../dates';
import { DEFAULT_FACTION } from './tier-controls';
import { ROLE_TABS, specFullName, tiersCopy } from './tier-copy';
import { groupByTier, roleHasTiers, type TierRow } from './tier-list';
import type { TierData } from './tier-load';
import type { TierLetter } from './tier-rules';

export interface TierCardSpec {
  /** "Fury Warrior" */
  name: string;
  classSlug: string;
  /** The row's rank as the page prints it ("2", "=2"); the damage list leads with its letter instead. */
  rank: string;
}

export interface TierCardGroup {
  /** S to D on the damage list; `null` on the tank and healer lists, which carry no letters. */
  letter: TierLetter | null;
  specs: TierCardSpec[];
}

export interface TierCardModel {
  title: string;
  /** "Damage", "Tanks", "Healers": the page's role tab. */
  role: string;
  /** "Best in slot, raid buffs · vs a level 63 boss · 3 min · Updated 12 Nov" */
  meta: string;
  groups: TierCardGroup[];
}

/** The hash length of the card's cache-busting query: enough to never collide across deploys. */
const VERSION_LENGTH = 10;

function cardSpec(row: TierRow): TierCardSpec {
  return { name: specFullName(row), classSlug: row.classSlug, rank: tiersCopy.rankLabel(row) };
}

export function tierCardModel(data: TierData, role: BisRole): TierCardModel {
  const rows = data.lists[DEFAULT_FACTION][role];
  const groups: TierCardGroup[] = roleHasTiers(role)
    ? groupByTier(rows).map((group) => ({ letter: group.letter, specs: group.rows.map(cardSpec) }))
    : [{ letter: null, specs: rows.map(cardSpec) }];
  const meta = tiersCopy.metaLine({
    role,
    healerBasis: data.healerBasis,
    healSeconds: data.healSeconds,
    raidFightSeconds: data.raidFightSeconds,
  });
  const updated = data.updated === null ? '' : updatedLabel(data.updated);
  return {
    title: tiersCopy.title,
    role: ROLE_TABS[role],
    meta: updated === '' ? meta : `${meta} · ${updated}`,
    groups,
  };
}

/** The card's version: it changes whenever anything the card draws changes, so the page's
 *  og:image URL changes with it and Reddit and Discord fetch the new card. */
export function tierCardVersion(model: TierCardModel): string {
  return createHash('sha256').update(JSON.stringify(model)).digest('hex').slice(0, VERSION_LENGTH);
}
