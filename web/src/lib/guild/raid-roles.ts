// web/src/lib/guild/raid-roles.ts
// The raid role a spec plays (spec 3.3: the role select is prefilled from the spec), derived from
// the generated SPECS so the role has one source. Specs that share a name across classes (Holy,
// Restoration, Protection) share a role.
import { SPECS } from '../sim/specs';
import type { GuildRaidRole } from './api';
import { specLookupKey } from './spec-name';

/** Spec display name (`Feral Bear`) → raid role. */
export const SPEC_ROLE: Readonly<Record<string, GuildRaidRole>> = Object.fromEntries(
  SPECS.map((spec) => [spec.name, spec.role]),
);

const BY_KEY = new Map<string, GuildRaidRole>(
  SPECS.flatMap((spec) => [
    [specLookupKey(spec.name), spec.role],
    [specLookupKey(spec.spec), spec.role],
  ]),
);

const ROLE_ORDER: readonly GuildRaidRole[] = ['tank', 'healer', 'dps'];

/** The roles a class plays in Classic, tank, healer, DPS order (a mage only DPS); every role for
 *  an unknown class, which the API does not judge either. */
export function rolesForClass(cls: string): GuildRaidRole[] {
  const slug = cls.trim().toLowerCase();
  const played = new Set(SPECS.filter((spec) => spec.class_slug === slug).map((spec) => spec.role));
  return played.size === 0 ? [...ROLE_ORDER] : ROLE_ORDER.filter((role) => played.has(role));
}

/** The role for a spec name (any case or spacing, so a log's `BeastMastery` reads) or spec key
 *  (`druid-feral-bear`); `null` when unknown. */
export function roleForSpec(spec: string): GuildRaidRole | null {
  return BY_KEY.get(specLookupKey(spec)) ?? null;
}
