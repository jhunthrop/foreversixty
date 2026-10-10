// web/src/lib/bis/follow-selector.ts
// What a BiS spec page does when the header selector changes character (owner rule: every page
// follows the switcher). A pure decision -- stay, update band/faction in place, or go to the
// character's own spec page -- so the page script only has to carry it out.
import { SPECS } from '../sim/specs';
import { bandForLevel, bisPageHref } from './hover';
import type { Faction } from './types';

/** Where the open page is. */
export interface BisPageState {
  /** The spec key, "hunter-beast-mastery". */
  specKey: string;
  faction: Faction;
  band: number;
}

/** What the selector now points at. Every field but the class may be unknown. */
export interface SelectedTarget {
  /** Empty for a pasted export that names no class. */
  classSlug: string;
  specKey?: string;
  level?: number;
  faction?: Faction;
}

export type BisFollowAction =
  { kind: 'stay' } | { kind: 'update'; faction: Faction; band: number } | { kind: 'navigate'; href: string };

/** The class's first spec, the page /bis lists first for it. */
export function defaultSpecKeyFor(classSlug: string): string | undefined {
  return SPECS.find((row) => row.class_slug === classSlug)?.spec;
}

export function decideBisFollow(page: BisPageState, target: SelectedTarget): BisFollowAction {
  if (target.classSlug === '') return { kind: 'stay' };
  const specKey = target.specKey ?? defaultSpecKeyFor(target.classSlug);
  if (specKey === undefined) return { kind: 'stay' };
  const faction = target.faction ?? page.faction;
  const band = target.level === undefined ? page.band : bandForLevel(target.level);
  if (specKey !== page.specKey) return { kind: 'navigate', href: bisPageHref(specKey, faction, band) };
  if (faction === page.faction && band === page.band) return { kind: 'stay' };
  return { kind: 'update', faction, band };
}
