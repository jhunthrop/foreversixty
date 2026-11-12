// web/src/lib/home/home-view.ts
// The home's shared facts and words (design/specs/2026-11-06-home.md §3, kept by layout B
// 2026-11-07): the signed-out copy, the meta line, the band and gain words, the upgrade lines and
// the guild links. Pure: the islands read them, the unit tests pin them. Layout B's own blocks
// are in home-b-view.ts.
import type { Me, MeCharacter } from '../account/api';
import { ctaCopy } from '../addon/cta-copy';
import { bisHref } from '../character-profile';
import { guildHref, guildRaidHref } from '../characters';
import { scoreUnitWord, type SlotScoreUnit } from '../bis/tank-view';
import type { GainFigure } from '../ui/primitives';
import type { SlotUpgrade } from './upgrades';
import { SLOT_LABELS } from '../planner/types';
import {
  sharedHomeCheckAgain,
  sharedHomeCheckAgainSub,
  sharedHomeGetIn,
  sharedHomeHowToStart,
  sharedHomeNoCharacterYet,
  sharedHomeOpenBis,
  sharedHomePaste,
  sharedHomePasteAction,
  sharedHomePickClass,
  sharedHomeSetupAction,
  sharedHomeSignIn,
} from '../copy/shared/home.generated';

import { oneDecimal } from '../format/number';
import { dayMonth, utcDayMonth } from '../dates';
import { titleCase } from '../text';
/** How many upgrades home lists (round 5 mock: five in the framed rail). */
export const UPGRADES_SHOWN = 5;
const MAX_LEVEL = 60;
const BAND_SPAN = 9;

export const homeCopy = {
  pickClass: sharedHomePickClass,
  howToStart: sharedHomeHowToStart,
  steps: [{ title: sharedHomeSignIn }, { title: ctaCopy.installAddonAndCompanion }],
  openBis: sharedHomeOpenBis,
  getIn: sharedHomeGetIn,
  /** The signed-in home's head while the account read is in flight (a screen reader's heading). */
  loadingTitle: 'Loading your character',
  checkAgain: sharedHomeCheckAgain,
  noCharacterYet: sharedHomeNoCharacterYet,
  checkAgainSub: sharedHomeCheckAgainSub,
  paste: sharedHomePaste,
  pasteAction: sharedHomePasteAction,
  getAddon: ctaCopy.getAddonAndCompanion,
  setupAction: sharedHomeSetupAction,
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
export function shortDate(iso: string): string {
  return /^\d{4}-\d{2}-\d{2}$/.test(iso) ? utcDayMonth(iso) : dayMonth(new Date(iso));
}

/** The head's one line, as the client's character pane reads it: `Level 60 Human Fury Warrior`.
 *  Unknown facts are left out, never a placeholder. */
export function whoLine(character: MeCharacter): string {
  // The API may send the class in lower case ("hunter"); the page names it as a title.
  return [
    character.level === undefined ? '' : `Level ${character.level}`,
    character.race ?? '',
    ...[character.spec, character.class].map((part) => titleCase(part ?? '')),
  ]
    .filter((part) => part !== '')
    .join(' ');
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

/** An upgrade row's one line under the name, as the client and Wowhead word it: the slot (a ring
 *  is a Finger, whichever of the two) and where the pick comes from (a profession says crafted). */
export function upgradeLine(upgrade: Pick<SlotUpgrade, 'slot' | 'pick'>): string {
  const slot = SLOT_LABELS[upgrade.slot].replace(/ \d$/, '');
  return `${slot} · ${upgrade.pick.source}`;
}

/** `+5.1` */
export function gainText(dps: number): string {
  return `+${oneDecimal(dps)}`;
}

/** An upgrade row's gain at its right: the figure in the band's own unit, the word the paperdoll
 *  uses (`+5.1 DPS`, `+5.1 HPS`, a tank's `+5.1 score`); undefined with no measured gain. */
export function upgradeGain(
  upgrade: Pick<SlotUpgrade, 'gainDps'>,
  unit: SlotScoreUnit,
): GainFigure | undefined {
  if (upgrade.gainDps === null) return undefined;
  return { value: gainText(upgrade.gainDps), unit: scoreUnitWord(unit), tone: 'kill' };
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
