// web/src/lib/guild/officer-tools.ts
// Officer tools gate §3.1/§4: whether the officer tools are on for the viewer. Only an officer can
// be in officer-read-only; a moderator's Award is refused by rank and a member never sees the
// tools, so both keep today's view.
import type { GuildViewerRole } from './api';

/** A `guild_plan_required` refusal: which boss's item, so a remount or another boss cannot show it. */
export interface GuildPlanRefusal {
  encounterId: number;
  itemId: number;
  kind: 'award' | 'undo';
}

/**
 * False only for an officer whose guild lacks the tools: the API says so (`officer_tools: false`)
 * or an Award/Undo was refused with `guild_plan_required` after the page loaded. A missing field
 * (an older API) counts as on: the server's 403 is the real gate.
 */
export function officerToolsOn(
  role: GuildViewerRole,
  officerToolsField: boolean | undefined,
  ended: boolean,
): boolean {
  if (role !== 'officer') return true;
  return (officerToolsField ?? true) && !ended;
}

const PLAN_REQUIRED = 'guild_plan_required';
const RAID_GUILD_REQUIRED = 'raid_guild_required';

/** Which gate an Award/Undo refusal code ends: the plan's (`guild_plan_required`), Raid guild's
 *  (`raid_guild_required`, a Guild-tier guild), or none (any other error). */
export function refusalGate(code: string): 'plan' | 'raid-guild' | null {
  if (code === PLAN_REQUIRED) return 'plan';
  return code === RAID_GUILD_REQUIRED ? 'raid-guild' : null;
}

/**
 * Launch pricing spec §3.2.8: Loot council and the readiness board belong to Raid guild. False
 * only for an officer whose guild lacks it: the API says so (`raid_guild_tools: false`) or an
 * Award/Undo was refused with `raid_guild_required` after the page loaded. The API always sends
 * the field, so there is no fallback here.
 */
export function raidGuildToolsOn(
  role: GuildViewerRole,
  raidGuildToolsField: boolean,
  ended: boolean,
): boolean {
  if (role !== 'officer') return true;
  return raidGuildToolsField && !ended;
}
