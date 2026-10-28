// web/src/lib/guild/my-guild.ts
// The one rule for "which of the account's guilds is this page's guild": region + ruleset +
// slugified name, not region + ruleset alone, which many guilds share.
import { characterSlug, type CharacterPath } from '../characters';
import type { Me, MeGuild } from '../account/api';

export function findMyGuild(me: Me | null | undefined, path: CharacterPath): MeGuild | null {
  return (
    me?.guilds.find(
      (g) => g.region === path.region && g.ruleset === path.ruleset && characterSlug(g.name) === path.slug,
    ) ?? null
  );
}
