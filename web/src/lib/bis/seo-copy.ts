// web/src/lib/bis/seo-copy.ts
// What a BiS spec page and each of its band pages (lib/bis/band-path.ts) tell a search engine and
// say in their h1: spec before class, the way players search ("Fury Warrior BiS"), the band in
// words, "Pre-Raid" only while the level 60 list holds no raid drop (page-server.ts bisPreRaid).
import { bandWords } from './band-path';
import { MAX_BAND } from './band-levels';

export interface BisSeoSubject {
  /** The spec as a player names it in full, the catalogue's `full_name` ("Fury Warrior"). */
  fullName: string;
  /** The level 60 list holds no raid drop. */
  preRaid: boolean;
}

const GAME = 'WoW Forever';

function who({ fullName }: BisSeoSubject): string {
  return fullName;
}

export const bisSeoCopy = {
  /** The spec's own page, the hub for every band: "Fury Warrior BiS: Pre-Raid and Leveling 20 to 60, WoW Forever". */
  specTitle: (subject: BisSeoSubject, firstBand: number): string =>
    `${who(subject)} BiS: ${subject.preRaid ? 'Pre-Raid' : `Level ${MAX_BAND}`} and Leveling ${firstBand} to ${MAX_BAND}, ${GAME}`,
  /** A band page: "Fury Warrior BiS, Levels 40 to 49, WoW Forever"; at 60 "Fury Warrior Pre-Raid BiS, Level 60, WoW Forever". */
  bandTitle: (subject: BisSeoSubject, band: number): string =>
    band >= MAX_BAND && subject.preRaid
      ? `${who(subject)} Pre-Raid BiS, ${bandWords(band)}, ${GAME}`
      : `${who(subject)} BiS, ${bandWords(band)}, ${GAME}`,
  /** The h1, one line at 1440 for every spec (the longest, "Marksmanship Hunter BiS", is 470 px of
   *  the head's 520): "Fury Warrior BiS". The band is named under it (`meta`) and in the <title>. */
  heading: (subject: BisSeoSubject): string => `${who(subject)} BiS`,
  /** The head's meta line under the h1, naming the band shown: "Best in slot · Levels 40 to 49",
   *  "Pre-raid best in slot · Level 60". */
  meta: (subject: BisSeoSubject, band: number): string =>
    `${band >= MAX_BAND && subject.preRaid ? 'Pre-raid best in slot' : 'Best in slot'} · ${bandWords(band)}`,
  /** The band page's breadcrumb name: "Fury Warrior BiS, Levels 40 to 49". */
  bandCrumb: (subject: BisSeoSubject, band: number): string =>
    band >= MAX_BAND && subject.preRaid
      ? `${who(subject)} Pre-Raid BiS`
      : `${who(subject)} BiS, ${bandWords(band)}`,
  /** The level 60 pre-raid list's own page (/60/pre-raid): "Fury Warrior Pre-Raid BiS, WoW Forever". */
  preRaidTitle: (subject: BisSeoSubject): string => `${who(subject)} Pre-Raid BiS, ${GAME}`,
  preRaidDescription: (subject: BisSeoSubject): string =>
    `${who(subject)} pre-raid best in slot gear at level ${MAX_BAND} in ${GAME}: no raid or high-rank PvP pieces, every slot ranked by the simulator, with the dungeon, quest or vendor each item comes from.`,
  preRaidCrumb: (subject: BisSeoSubject): string => `${who(subject)} Pre-Raid BiS`,
  specDescription: (subject: BisSeoSubject, firstBand: number): string =>
    `${who(subject)} best in slot gear for ${GAME}, from leveling ${firstBand} to ${MAX_BAND} to ${subject.preRaid ? 'pre-raid ' : ''}level ${MAX_BAND}: every slot ranked by the simulator, with where each item comes from.`,
  bandDescription: (subject: BisSeoSubject, band: number): string =>
    `${who(subject)} best in slot gear at ${bandWords(band).toLowerCase()} in ${GAME}: every slot ranked by the simulator, with the dungeon, quest or vendor each item comes from.`,
} as const;
