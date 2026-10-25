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
