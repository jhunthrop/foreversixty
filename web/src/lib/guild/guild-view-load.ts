// web/src/lib/guild/guild-view-load.ts
// The first read of every view under a guild page that is not the page itself (recruit,
// applications, a member profile): the guild by its path (its id, name and mark) and who is
// asking (/v1/me, the account's membership and rank there). The server's own 403 stays the gate;
// this only picks which view to draw.
import type { Me, MeGuild } from '../account/me-types';
import { characterSlug, type CharacterPath } from '../characters';
import { guildMarkSrc } from './mark';
import type { GuildPage } from '../rankings/api';

export interface GuildViewGuild {
  id: number;
  name: string;
  faction: 'alliance' | 'horde' | null;
  mark: string | null;
}

export type GuildViewRole = 'signed-out' | 'visitor' | 'unverified' | 'member' | 'officer';

export interface GuildViewContext {
  guild: GuildViewGuild;
  me: Me | null;
  membership: MeGuild | null;
  role: GuildViewRole;
}

export interface GuildViewDeps {
  loadGuild: (path: CharacterPath) => Promise<GuildPage>;
  loadMe: () => Promise<Me | null>;
}

/** The account's line for this guild in /v1/me, matched by realm and slug. */
export function membershipOf(me: Me | null, path: CharacterPath): MeGuild | null {
  return (
    me?.guilds.find(
      (g) => g.region === path.region && g.ruleset === path.ruleset && characterSlug(g.name) === path.slug,
    ) ?? null
  );
}

/** Signed out, not in the guild, waiting for approval, a member, or an officer (or the guild master). */
export function viewRole(me: Me | null, membership: MeGuild | null): GuildViewRole {
  if (me === null) return 'signed-out';
  if (membership === null) return 'visitor';
  if (!membership.verified) return 'unverified';
  return membership.rank === 'officer' || membership.rank === 'leader' ? 'officer' : 'member';
}

export async function loadGuildView(path: CharacterPath, deps: GuildViewDeps): Promise<GuildViewContext> {
  const [page, me] = await Promise.all([deps.loadGuild(path), deps.loadMe().catch(() => null)]);
  const { guild } = page;
  const membership = membershipOf(me, path);
  return {
    guild: { id: guild.id, name: guild.name, faction: guild.faction ?? null, mark: guildMarkSrc(guild) },
    me,
    membership,
    role: viewRole(me, membership),
  };
}
