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

/**
 * What a BiS spec page does at load when the URL names no band (owner 2026-11-05: "default to
 * the current level range of their selected character"). Unlike a selector change it never
 * leaves the page: a player who opened another spec's list on purpose keeps it, and only the
 * band and faction move to the selected character's when it is this spec's.
 */
export function decideBisDefault(page: BisPageState, target: SelectedTarget): BisFollowAction {
  if (target.classSlug === '') return { kind: 'stay' };
  const specKey = target.specKey ?? defaultSpecKeyFor(target.classSlug);
  if (specKey !== page.specKey) return { kind: 'stay' };
  const faction = target.faction ?? page.faction;
  const band = target.level === undefined ? page.band : bandForLevel(target.level);
  if (faction === page.faction && band === page.band) return { kind: 'stay' };
  return { kind: 'update', faction, band };
}

/** `?band=60` (or any level, `?band=34`) as the band it opens, or null when absent or not a number. */
export function bandFromParam(value: string | null): number | null {
  if (value === null || !/^\d{1,2}$/.test(value)) return null;
  return bandForLevel(Number(value));
}
