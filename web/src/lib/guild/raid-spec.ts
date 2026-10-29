// web/src/lib/guild/raid-spec.ts
// The raid page's spec ask (guided-cards spec P.10.7): which specs a character can pick, the one
// it saved, whether a sign-up state asks, and the spec's own icon. One catalogue (SPECS) feeds
// the options, the role the pick signs up with and the icon.
import activeBuild from '../../data/active-build.json';
import { dataUrl } from '../planner/load';
import { SPECS, type Spec } from '../sim/specs';
import type { GuildRaidViewerCharacter, GuildSignupState } from './api';

/** The sign-up states that carry a spec; Out never asks. */
const SPEC_STATES: readonly GuildSignupState[] = ['in', 'late', 'bench'];

/** The class's specs in catalogue order. */
export function specsForClass(classSlug: string): readonly Spec[] {
  return SPECS.filter((spec) => spec.class_slug === classSlug);
}

/** The character's saved raid spec, else `null`. */
export function savedSpec(character: GuildRaidViewerCharacter): Spec | null {
  const slug = character.raid_spec;
  if (slug === null || slug === undefined) return null;
  return specsForClass(character.class).find((spec) => spec.spec === slug) ?? null;
}

/** The spec of the class with this slug, else `null` (a forged value is refused, never sent). */
export function specOfClass(classSlug: string, slug: string): Spec | null {
  return specsForClass(classSlug).find((spec) => spec.spec === slug) ?? null;
}

/** The select shows for In, Late and Bench when the class has specs to pick from. */
export function showsSpecSelect(classSlug: string, state: GuildSignupState | null): boolean {
  return state !== null && SPEC_STATES.includes(state) && specsForClass(classSlug).length > 0;
}

/** A press of In, Late or Bench waits for the pick when the character has no saved spec. */
export function pressAsksSpec(character: GuildRaidViewerCharacter, state: GuildSignupState): boolean {
  return showsSpecSelect(character.class, state) && savedSpec(character) === null;
}

/** The spec's talent-tab icon, from the active build. */
export function specIconSrc(spec: Spec): string {
  return dataUrl(activeBuild.build, `icons/${spec.icon}.webp`);
}
