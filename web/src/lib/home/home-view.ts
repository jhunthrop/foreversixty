// web/src/lib/home/home-view.ts
// The signed-in home's facts and words (design/specs/2026-11-06-home.md §1 and §3): the meta
// line, the one next-step line, the talents, sim, logs, guild and set-up rows. Pure: the island
// (components/home/HomeSignedIn.svelte) reads them, the unit tests pin them.
import type { Me, MeCharacter, MyReport } from '../account/api';
import { bisHref } from '../character-profile';
import { guildHref, guildRaidHref } from '../characters';
import type { NextRaidView } from '../guild/api';
import { pointsEarnedAt } from '../leveling/leveling-plan';
import { FIRST_POINT_LEVEL } from '../planner/types';
import { armorySimHref } from '../sim/url';
import type { SimListRow } from '../sim/types';
import type { SlotUpgrade } from './upgrades';
import { slotSourceLine } from './source-line';

/** How many upgrades home lists before "All N on best in slot" (ruling H3). */
export const UPGRADES_SHOWN = 3;
/** How many of the character's own logs the rail lists (ruling H7). */
export const LOGS_SHOWN = 3;
/** An addon export older than this warns in Set up: the upgrades are read from it (ruling H7). */
export const STALE_EXPORT_DAYS = 7;
const DAY_MS = 24 * 60 * 60 * 1000;
const MAX_LEVEL = 60;
const BAND_SPAN = 9;

export const homeCopy = {
  signedOutTitle: 'Play your class better.',
  signedOutMeta: 'Pick your class.',
  signIn: 'Sign in with Battle.net',
  pickClass: 'Pick your class',
  example: (name: string, first: number, last: number) => `${name}, level ${first} to ${last}`,
  exampleAction: 'Best in slot',
  exampleCaption: 'the set',
  exampleLabel: 'Best in slot by level',
  howToStart: 'How to start',
  steps: [
    {
      title: 'Sign in with Battle.net',
      sub: 'Imports your characters with their level, gear and item level.',
    },
    { title: 'Install the addon', sub: "Rates every item in your bags with your spec's weights." },
  ],
  firstRaids: (date: string) => `First raids ${date}`,
  answerLabel: (band: string) => `Upgrades at ${band}`,
  openBis: 'Open best in slot',
  upgrades: 'Your upgrades',
  upgradesMeta: (band: string) => (band === '60' ? 'Level 60' : `Levels ${band}`),
  upgradesAll: (n: number) => `All ${n} on best in slot`,
  upgradesFailed: 'We could not load your upgrades.',
  upgradesNone: (band: string) => `You wear every pick for ${band === '60' ? 'level 60' : `levels ${band}`}.`,
  vsWorn: (name: string) => `vs your ${name}`,
  vsEmpty: 'empty slot',
  nextRaid: 'Next raid',
  raidPage: 'Raid page',
  talents: 'Talents',
  talentsTitle: (n: number) => `${n} talent ${n === 1 ? 'point' : 'points'} to spend`,
  talentsSub: (level: number, earned: number, spent: number) =>
    `Level ${level} earns ${earned}. You have spent ${spent}.`,
  simulator: 'Simulator',
  openSim: 'Open the simulator',
  simmed: (date: string) => `Simmed ${date}`,
  firstSim: 'Run your first sim',
  firstSimSub: 'Your gear and talents, against a training dummy.',
  guild: 'Guild',
  attendance: (present: number, nights: number) => `${present} of ${nights} raids`,
  logs: 'Your logs',
  allLogs: 'All logs',
  noLogs: 'Upload a log',
  noLogsSub: 'The companion uploads your logs as you play.',
  bosses: (n: number) => `${n} ${n === 1 ? 'boss' : 'bosses'}`,
  setUp: 'Set up',
  bnet: 'Battle.net import',
  bnetDone: (date: string) => `Imported ${date}`,
  bnetNone: 'Not imported yet',
  addon: 'Addon export',
  addonDone: (date: string) => `Exported ${date}`,
  addonStale: (date: string) => `Exported ${date}. Your upgrades are read from it.`,
  addonNone: 'No export yet. Your upgrades are read from it.',
  addonFix: 'Get the addon',
  noCharacter: 'No character yet. Battle.net lists a character after you play it and log out.',
  signedInAs: (tag: string) => `Signed in as ${tag}`,
  getIn: 'Get your character in',
  checkAgain: 'Check again',
  checkAgainSub: 'Battle.net lists a character after you play it and log out.',
  paste: 'Paste an addon export',
  pasteSub: 'Your gear and talents, straight from the game.',
  getAddon: 'Get the addon',
  getAddonSub: "Rates every item in your bags with your spec's weights.",
  openRaid: 'Open the raid',
  openPlanner: 'Open the planner',
} as const;

export const homeHrefs = {
  paste: '/setup#paste',
  addon: '/setup',
  logs: '/logs',
  logsUpload: '/logs/upload',
  sim: '/sim',
  guide: (classSlug: string) => `/guides/${classSlug}`,
  report: (id: string) => `/reports/${id}`,
  simResult: (id: string) => `/sim/${id}`,
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
  const who = [character.spec, character.class].filter(Boolean).join(' ');
  const level = character.level === undefined ? who : `Level ${character.level} ${who}`.trim();
  const faction = character.faction === undefined ? '' : capitalise(character.faction);
  const origin = [faction, character.race ?? ''].filter(Boolean).join(' ');
  const itemLevel = character.item_level === undefined ? '' : `Item level ${character.item_level}`;
  return [level, origin, itemLevel].filter((fact) => fact !== '');
}

/** Points earned and not spent, from the addon's talent export; null when it has none. */
export function talentsToSpend(
  character: MeCharacter,
): { toSpend: number; earned: number; spent: number } | null {
  const talents = character.build?.talents;
  if (talents === undefined || character.level === undefined || character.level < FIRST_POINT_LEVEL)
    return null;
  const earned = pointsEarnedAt(character.level);
  const spent = talents.points.reduce((sum, count) => sum + count, 0);
  return { toSpend: Math.max(0, earned - spent), earned, spent };
}

export interface HomeLink {
  label: string;
  href: string;
}

/** The one next-step line (ruling H2), most urgent first: the raid sign-up, talent points, the
 *  top upgrade. Null when nothing waits and nothing is known. */
export function nextStep(input: {
  character: MeCharacter;
  upgrades: readonly SlotUpgrade[] | null;
  raid: NextRaidView | null;
  raidHref: string | null;
  raidDay: string;
}): { text: string; action: HomeLink | null } | null {
  const { character, upgrades, raid, raidHref, raidDay } = input;
  const state = raid?.own?.state ?? null;
  if (raid !== null && state !== 'in') {
    const where = state === null ? 'you are not signed up' : `you are ${capitalise(state)}`;
    return {
      text: `Next: ${raid.title} is on ${raidDay} and ${where}. Confirm your sign-up.`,
      action: raidHref === null ? null : { label: homeCopy.openRaid, href: raidHref },
    };
  }
  const talents = talentsToSpend(character);
  const top = upgrades?.[0]?.pick;
  if (talents !== null && talents.toSpend > 0) {
    const then = top === undefined ? '' : `, then get ${top.item_name}`;
    return {
      text: `Next: spend your ${talents.toSpend} talent ${talents.toSpend === 1 ? 'point' : 'points'}${then}.`,
      action: { label: homeCopy.openPlanner, href: homeHrefs.planner(character.key) },
    };
  }
  const bis = bisLink(character);
  if (top !== undefined) {
    return { text: `Next: get ${top.item_name} from ${top.source}.`, action: bis };
  }
  if (upgrades !== null) return { text: 'Nothing waiting.', action: null };
  return null;
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

export function simRow(
  sim: SimListRow | undefined,
  character: MeCharacter,
): { title: string; sub: string; href: string; dps: string | null } {
  if (sim === undefined) {
    return {
      title: homeCopy.firstSim,
      sub: homeCopy.firstSimSub,
      href: armorySimHref(character.key),
      dps: null,
    };
  }
  return {
    title: sim.title,
    sub: homeCopy.simmed(shortDate(sim.created_at)),
    href: homeHrefs.simResult(sim.sim_id),
    dps: sim.dps.toLocaleString('en-US', { minimumFractionDigits: 1, maximumFractionDigits: 1 }),
  };
}

/** The character's own logs, newest first (a log credited to another character is left out). */
export function ownLogs(reports: readonly MyReport[], character: MeCharacter): MyReport[] {
  return reports
    .filter((report) => {
      const who = report.logging_character;
      if (who === null || who === undefined) return true;
      const key = typeof who === 'string' ? who : who.key;
      return key === undefined || key === character.key;
    })
    .slice(0, LOGS_SHOWN);
}

export function logSub(report: MyReport): string {
  const parts = [report.zone];
  if (report.boss_count !== undefined) parts.push(homeCopy.bosses(report.boss_count));
  parts.push(shortDate(report.started_at ?? report.created_at));
  return parts.join(' · ');
}

export interface SetUpCheck {
  key: string;
  tone: 'ok' | 'warn';
  title: string;
  sub: string;
  fix: HomeLink | null;
}

/** Set up (ruling H7): the Battle.net import and the addon export, with their age. */
export function setUpChecks(me: Me, character: MeCharacter, now: Date = new Date()): SetUpCheck[] {
  const imported = me.bnet_imported_at;
  const bnet: SetUpCheck = {
    key: 'bnet',
    tone: imported === undefined ? 'warn' : 'ok',
    title: homeCopy.bnet,
    sub: imported === undefined ? homeCopy.bnetNone : homeCopy.bnetDone(shortDate(imported)),
    fix: null,
  };
  const build = character.build;
  const captured = build?.source === 'addon' ? build.captured_at : undefined;
  const fix = { label: homeCopy.addonFix, href: homeHrefs.addon };
  let addon: SetUpCheck;
  if (captured === undefined) {
    addon = { key: 'addon', tone: 'warn', title: homeCopy.addon, sub: homeCopy.addonNone, fix };
  } else if (now.getTime() - new Date(captured).getTime() > STALE_EXPORT_DAYS * DAY_MS) {
    addon = {
      key: 'addon',
      tone: 'warn',
      title: homeCopy.addon,
      sub: homeCopy.addonStale(shortDate(captured)),
      fix,
    };
  } else {
    addon = {
      key: 'addon',
      tone: 'ok',
      title: homeCopy.addon,
      sub: homeCopy.addonDone(shortDate(captured)),
      fix: null,
    };
  }
  return [bnet, addon];
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

function capitalise(word: string): string {
  return word.charAt(0).toUpperCase() + word.slice(1);
}
