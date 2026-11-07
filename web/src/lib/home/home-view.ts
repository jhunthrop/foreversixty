// web/src/lib/home/home-view.ts
// The home's shared facts and words (design/specs/2026-11-06-home.md §3, kept by layout B
// 2026-11-07): the signed-out copy, the meta line, the band and gain words, the upgrade lines and
// the guild links. Pure: the islands read them, the unit tests pin them. Layout B's own blocks
// are in home-b-view.ts.
import type { Me, MeCharacter } from '../account/api';
import { ctaCopy } from '../addon/cta-copy';
import { bisHref } from '../character-profile';
import { guildHref, guildRaidHref } from '../characters';
import type { SlotUpgrade } from './upgrades';
import { slotSourceLine } from './source-line';
import { sharedText } from '../copy/shared';

/** How many upgrades home lists (ruling R1). */
export const UPGRADES_SHOWN = 3;
const MAX_LEVEL = 60;
const BAND_SPAN = 9;

export const homeCopy = {
  signedOutTitle: sharedText('shared.home.signedOutTitle'),
  signedOutMeta: sharedText('shared.home.signedOutMeta'),
  signIn: sharedText('shared.home.signIn'),
  pickClass: sharedText('shared.home.pickClass'),
  howToStart: sharedText('shared.home.howToStart'),
  steps: [
    {
      title: sharedText('shared.home.signIn'),
      sub: sharedText('shared.home.stepSignInSub'),
    },
    {
      title: ctaCopy.installAddonAndCompanion,
      sub: sharedText('shared.home.stepAddonSub'),
    },
  ],
  firstRaids: (date: string) => sharedText('shared.home.firstRaids', { date }),
  openBis: sharedText('shared.home.openBis'),
  vsWorn: (name: string) => sharedText('shared.home.vsWorn', { name }),
  vsEmpty: sharedText('shared.home.vsEmpty'),
  noCharacter: sharedText('shared.home.noCharacter'),
  signedInAs: (tag: string) => sharedText('shared.home.signedInAs', { tag }),
  getIn: sharedText('shared.home.getIn'),
  checkAgain: sharedText('shared.home.checkAgain'),
  checkAgainSub: sharedText('shared.home.checkAgainSub'),
  paste: sharedText('shared.home.paste'),
  pasteSub: sharedText('shared.home.pasteSub'),
  getAddon: ctaCopy.getAddonAndCompanion,
  getAddonSub: sharedText('shared.home.stepAddonSub'),
} as const;

export const homeHrefs = {
  paste: '/setup#paste',
  addon: '/setup',
  guide: (classSlug: string) => `/guides/${classSlug}`,
  planner: (key: string) => `/planner?character=${encodeURIComponent(key)}`,
} as const;

/** A band as words: `30 to 39`, `60`. */
export function bandWords(band: number): string {
  return band >= MAX_LEVEL ? String(MAX_LEVEL) : `${band} to ${band + BAND_SPAN}`;
}

/** `4 Nov`, in the viewer's own calendar; a date with no time (`2026-12-09`) is that calendar day. */
export function shortDate(iso: string, locale = 'en-GB'): string {
  const dateOnly = /^\d{4}-\d{2}-\d{2}$/.test(iso);
  return new Date(iso).toLocaleDateString(locale, {
    day: 'numeric',
    month: 'short',
    ...(dateOnly ? { timeZone: 'UTC' } : {}),
  });
}

/** The head's one fact line: level, spec and class; faction and race; item level. Unknown facts
 *  are left out, never a placeholder. */
export function metaFacts(character: MeCharacter): string[] {
  // The API may send the class in lower case ("hunter"); the page names it as a title.
  const who = [character.spec, character.class]
    .filter(Boolean)
    .map((part) => titleCase(part ?? ''))
    .join(' ');
  const level = character.level === undefined ? who : `Level ${character.level} ${who}`.trim();
  const faction = character.faction === undefined ? '' : capitalise(character.faction);
  const origin = [faction, character.race ?? ''].filter(Boolean).join(' ');
  const itemLevel = character.item_level === undefined ? '' : `Item level ${character.item_level}`;
  return [level, origin, itemLevel].filter((fact) => fact !== '');
}

export interface HomeLink {
  label: string;
  href: string;
}

/** The character's best-in-slot page at its band, or null when its class or spec is unknown. */
export function bisLink(character: MeCharacter): HomeLink | null {
  const href = bisHref(character.class, character.spec ?? null, character.level, character.race);
  return href === null ? null : { label: homeCopy.openBis, href };
}

/** An upgrade row's source and its basis caption. */
export function upgradeLines(upgrade: SlotUpgrade): { source: string; caption: string } {
  return {
    source: slotSourceLine(upgrade.pick),
    caption: upgrade.wornItemName === undefined ? homeCopy.vsEmpty : homeCopy.vsWorn(upgrade.wornItemName),
  };
}

/** `+5.1` */
export function gainText(dps: number): string {
  return `+${dps.toLocaleString('en-US', { minimumFractionDigits: 1, maximumFractionDigits: 1 })}`;
}

/** The character's guild page and the raid page, from the account's guild entry. */
export function guildLinks(
  me: Me,
  guildId: number,
  raidId: number | null,
): { guild: string | null; raid: string | null } {
  const entry = me.guilds.find((guild) => guild.id === guildId);
  if (entry === undefined) return { guild: null, raid: null };
  const guild = guildHref(entry.region, entry.ruleset, entry.name);
  return { guild, raid: raidId === null ? guild : guildRaidHref(guild, raidId) };
}

function titleCase(words: string): string {
  return words.split(' ').map(capitalise).join(' ');
}

function capitalise(word: string): string {
  return word.charAt(0).toUpperCase() + word.slice(1);
}
