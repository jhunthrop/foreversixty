// web/src/lib/guild/mark.ts
// The one helper that decides what image represents a guild (docs/contracts/2026-10-05-
// guild-crest-api.md, Web section: "Every place that later shows a guild mark reads the
// same `crest_url ?? faction logo` rule from one helper, never two copies"). An uploaded
// crest wins outright; the faction's flat logo is the default; neither (no crest and no
// known faction) is `null`, which FactionCrest's own `src` override reads as "render
// nothing," matching the pre-crest neutral-band behaviour the header already shipped.
import { factionLogoSrc, type Faction } from '../faction-mark';

/** The subset of `GuildSummary`/`GuildPage.guild` this helper needs -- structural, so
 *  either shape (lib/guild/api.ts's `GuildSummary`, lib/rankings/api.ts's inline
 *  `GuildPage.guild`) passes without a cast. */
export interface GuildMarkSource {
  crest_url?: string | null;
  faction?: Faction | null;
}

export function guildMarkSrc(guild: GuildMarkSource | null | undefined): string | null {
  if (guild == null) return null;
  if (guild.crest_url != null) return guild.crest_url;
  return guild.faction === 'alliance' || guild.faction === 'horde' ? factionLogoSrc(guild.faction) : null;
}
