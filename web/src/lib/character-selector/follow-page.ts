// web/src/lib/character-selector/follow-page.ts
// How a page named by its address follows the header selector (owner, 2026-11-01: "when I
// switch characters it doesn't switch guilds"). A guild or character page is "yours" when it
// shows the previously-current character's guild or character; switching characters then moves
// it to the new one's. A page the visitor looked up stays put. Pure, so the decision is tested
// without a page; follow-page-run.ts reads the pointer and the account and carries it out.
import type { Me } from '../account/api';
import { selectedCharacter } from '../account/hero-character';
import type { CurrentCharacter } from '../current-character';
import {
  characterHref,
  guildHref,
  isGuildLandingPath,
  parseCharacterPath,
  parseGuildClaimPath,
  parseGuildPath,
  parseGuildRaidPath,
  parseGuildSettingsPath,
} from '../characters';
import type { SessionGuild } from './session-guild';

/** The find-or-claim landing a guildless character's guild view becomes. */
export const GUILD_LANDING_HREF = '/guild';

/** Where the page is: its path and its hash (a tab or a section, `#roster`). */
export interface PagePlace {
  readonly pathname: string;
  readonly hash: string;
}

export interface GuildFollowInput extends PagePlace {
  /** The previously-current character's guild; null when it had none. */
  readonly previous: SessionGuild | null;
  /** The newly-current character's guild; null when it has none. */
  readonly next: SessionGuild | null;
}

export interface CharacterFollowInput extends PagePlace {
  /** The previously-current character's page; null when it has none. */
  readonly previous: string | null;
  /** The newly-current character's page; null when it has none. */
  readonly next: string | null;
}

/** Two page paths name the same page however the address spells them (escapes, case). */
function samePage(a: string, b: string): boolean {
  return decodedPath(a) === decodedPath(b);
}

function decodedPath(path: string): string {
  try {
    return decodeURIComponent(path).toLowerCase();
  } catch {
    return path.toLowerCase();
  }
}

/** The guild a guild-page address shows, and the subpage the new guild keeps (null: none, and no tab). */
interface ShownGuild {
  readonly href: string;
  readonly suffix: string | null;
}

function shownGuild(pathname: string): ShownGuild | null {
  const overview = parseGuildPath(pathname);
  if (overview !== null)
    return { href: guildHref(overview.region, overview.ruleset, overview.slug), suffix: '' };
  const settings = parseGuildSettingsPath(pathname);
  if (settings !== null)
    return { href: guildHref(settings.region, settings.ruleset, settings.slug), suffix: '/settings' };
  const claim = parseGuildClaimPath(pathname);
  if (claim !== null) return { href: guildHref(claim.region, claim.ruleset, claim.slug), suffix: '/claim' };
  // A raid's id is its own guild's, so the new guild opens on its page.
  const raid = parseGuildRaidPath(pathname)?.guild;
  if (raid !== undefined) return { href: guildHref(raid.region, raid.ruleset, raid.slug), suffix: null };
  return null;
}

/**
 * Where a guild view goes when the current character changes, or null to stay. The landing
 * opens the new character's guild (the landing is a guildless character's guild view). A guild
 * page moves only when it shows the previous character's guild: to the new character's guild,
 * keeping the subpage and the tab, or to the landing for a guildless character.
 */
export function guildFollowHref(input: GuildFollowInput): string | null {
  if (isGuildLandingPath(input.pathname)) return input.next?.href ?? null;
  const shown = shownGuild(input.pathname);
  if (shown === null || input.previous === null || !samePage(shown.href, input.previous.href)) return null;
  if (input.next === null) return GUILD_LANDING_HREF;
  if (samePage(input.next.href, input.previous.href)) return null;
  if (shown.suffix === null) return input.next.href;
  return `${input.next.href}${shown.suffix}${input.hash}`;
}

/**
 * Where a character page goes when the current character changes, or null to stay: only the
 * previous character's own page moves, to the new character's page with the same section.
 */
export function characterFollowHref(input: CharacterFollowInput): string | null {
  const shown = parseCharacterPath(input.pathname);
  if (shown === null || input.previous === null || input.next === null) return null;
  if (!samePage(characterHref(shown.region, shown.ruleset, shown.slug), input.previous)) return null;
  if (samePage(input.next, input.previous)) return null;
  return `${input.next}${input.hash}`;
}

/**
 * The character page a pointer names: an armory pointer's own key (signed in or not), else the
 * account character `selectedCharacter` resolves (the main with no pointer); null otherwise.
 */
export function pointerCharacterHref(me: Me | null, pointer: CurrentCharacter | null): string | null {
  if (pointer?.source === 'armory') return `/character/${pointer.ref}`;
  if (me === null) return null;
  const character = selectedCharacter(pointer, {
    characters: me.characters ?? [],
    main_character_key: me.main_character_key,
  });
  return character === null ? null : characterHref(character.region, character.ruleset, character.name);
}
