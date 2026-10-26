// web/src/lib/guild/raid-roles.ts
// The raid role a spec plays (spec 3.3: the role select is prefilled from the spec), derived from
// the generated SPECS so the role has one source. Specs that share a name across classes (Holy,
// Restoration, Protection) share a role.
import { SPECS } from '../sim/specs';
import type { GuildRaidRole } from './api';

/** Spec display name (`Feral Bear`) → raid role. */
export const SPEC_ROLE: Readonly<Record<string, GuildRaidRole>> = Object.fromEntries(
  SPECS.map((spec) => [spec.name, spec.role]),
);

const BY_LOWER = new Map<string, GuildRaidRole>(
  SPECS.flatMap((spec) => [
    [spec.name.toLowerCase(), spec.role],
    [spec.spec.toLowerCase(), spec.role],
  ]),
);

/** The role for a spec name (any case) or spec key (`druid-feral-bear`); `null` when unknown. */
export function roleForSpec(spec: string): GuildRaidRole | null {
  return BY_LOWER.get(spec.trim().toLowerCase()) ?? null;
}
