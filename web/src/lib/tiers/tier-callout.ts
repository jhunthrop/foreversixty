// web/src/lib/tiers/tier-callout.ts
// The signed-in answer on the tier list: the visitor's own spec, ranked. The page embeds one
// ready-made view per spec and faction, and a small script picks the one that matches the
// visitor's selected character, so the text is written once, here, and tested.
import classes from '../../data/classes.json';
import { classCrestSrc } from '../class-crest';
import { classSlugFromName } from '../report/tree-sizes';
import { classColorVar } from '../report/format';
import type { BisRole } from '../bis/types';
import { ROLE_LABELS, ROLE_PATHS, ordinal, tiersCopy } from './tier-copy';
import type { TierRow } from './tier-list';

/** What the page needs to answer for one spec. Every field is final text or a path. */
export interface TierSpecView {
  role: BisRole;
  classSlug: string;
  crestSrc: string;
  colorVar: string;
  /** "Fury Warrior" */
  fullName: string;
  /** "Fury Warrior is 3rd of 8 DPS specs." */
  title: string;
  /** "4.8% behind Arms Warrior." */
  detail: string;
  /** The spec's BiS page, the callout button's target. */
  bisHref: string;
  /** Cross-role pointer text, "Fury Warrior is on the DPS list. See where it stands →". */
  pointer: string;
  /** The role's tab path, the pointer's target. */
  rolePath: string;
}

export type TierSpecViews = Record<string, TierSpecView>;

/** The key the script looks a character up by: `<class slug>/<spec slug>`. */
export function specViewKey(classSlug: string, specSlug: string): string {
  return `${classSlug}/${specSlug}`;
}

/** The lookup key for a signed-in character, or `undefined` while its class or spec is not
 *  known yet (never a guess at either). */
export function specViewKeyForCharacter(character: { class?: string; spec?: string }): string | undefined {
  if (character.class === undefined || character.spec === undefined) return undefined;
  return specViewKey(classSlugFromName(character.class), classSlugFromName(character.spec));
}

function classNameOf(classSlug: string): string {
  const found = classes.find((c) => c.slug === classSlug);
  if (found === undefined) throw new Error(`tier list: no class named ${classSlug} in classes.json`);
  return found.name;
}

function fullNameOf(row: Pick<TierRow, 'name' | 'classSlug'>): string {
  return `${row.name} ${classNameOf(row.classSlug)}`;
}

function detailFor(row: TierRow, top: TierRow): string {
  if (row.rank === 1) return row.role === 'tank' ? 'Takes the least damage.' : 'Level with the top spec.';
  const figure = row.gapPercent.toFixed(1);
  return row.role === 'tank'
    ? `${figure}% more damage taken than ${fullNameOf(top)}.`
    : `${figure}% behind ${fullNameOf(top)}.`;
}

/** Where a row stands, in words: no place on a list not ranked yet, "tied" on a shared rank. */
function placeTitle(row: TierRow, fullName: string, count: number, role: BisRole): string {
  if (!row.ranked) {
    return `${fullName} is on the ${ROLE_LABELS[role]} list; ranks appear once the next ranking publishes the measured error.`;
  }
  const tied = row.role === 'healer' && row.tie !== null ? 'tied ' : '';
  return `${fullName} is ${tied}${ordinal(row.rank)} of ${count} ${ROLE_LABELS[role]} specs.`;
}

/** One view per row of one faction's lists. */
export function tierSpecViews(lists: Record<BisRole, readonly TierRow[]>): TierSpecViews {
  const views: TierSpecViews = {};
  for (const [role, rows] of Object.entries(lists) as [BisRole, readonly TierRow[]][]) {
    const top = rows[0];
    if (top === undefined) continue;
    for (const row of rows) {
      const fullName = fullNameOf(row);
      views[specViewKey(row.classSlug, row.specSlug)] = {
        role,
        classSlug: row.classSlug,
        crestSrc: classCrestSrc(row.classSlug),
        colorVar: classColorVar(classNameOf(row.classSlug)),
        fullName,
        title: placeTitle(row, fullName, rows.length, role),
        detail: detailFor(row, top),
        bisHref: `/bis/${row.classSlug}/${row.specSlug}`,
        pointer: `${fullName} is on the ${ROLE_LABELS[role]} list. ${tiersCopy.pointerCta}`,
        rolePath: ROLE_PATHS[role],
      };
    }
  }
  return views;
}
