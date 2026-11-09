// web/src/lib/report/role-icon.ts
// A roster row's role as the game's own icon (owner review 2026-11-09: role icons on the Players
// table): the tank's Defensive Stance, the healer's Flash Heal, damage's Dual Wield, from the
// build's icon files (web/src/data/spell-icon-fallbacks.json `roles`, published at sync).
import activeBuild from '../../data/active-build.json';
import fallbacks from '../../data/spell-icon-fallbacks.json';
import { dataUrl } from '../planner/load';

export type Role = 'tank' | 'healer' | 'dps';

const ROLE_LABEL: Readonly<Record<Role, string>> = { tank: 'Tank', healer: 'Healer', dps: 'Damage' };
const STEMS: Readonly<Record<string, string>> = fallbacks.roles;

/** A known role, or null for a row the log could not place. */
export function roleOf(value: string | undefined): Role | null {
  return value === 'tank' || value === 'healer' || value === 'dps' ? value : null;
}

/** The role's icon url and its spoken name. */
export function roleIcon(role: Role, build: string = activeBuild.build): { url: string; label: string } {
  return { url: dataUrl(build, `icons/${STEMS[role]}.webp`), label: ROLE_LABEL[role] };
}
