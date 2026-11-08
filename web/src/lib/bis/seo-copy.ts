// web/src/lib/bis/seo-copy.ts
// What a BiS spec page and each of its band pages (lib/bis/band-path.ts) tell a search engine and
// say in their h1: spec before class, the way players search ("Fury Warrior BiS"), the band in
// words, "Pre-Raid" only while the level 60 list holds no raid drop (page-server.ts bisPreRaid).
import { bandWords } from './band-path';
import { MAX_BAND } from './band-levels';

export interface BisSeoSubject {
  specName: string;
  className: string;
  /** The level 60 list holds no raid drop. */
  preRaid: boolean;
}

const GAME = 'WoW Forever';

function who({ specName, className }: BisSeoSubject): string {
  return `${specName} ${className}`;
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
  /** The h1, naming the band shown: "Fury Warrior BiS, Levels 20 to 29", "Fury Warrior Pre-Raid BiS". */
  heading: (subject: BisSeoSubject, band: number): string =>
    band >= MAX_BAND && subject.preRaid
      ? `${who(subject)} Pre-Raid BiS`
      : `${who(subject)} BiS, ${bandWords(band)}`,
  specDescription: (subject: BisSeoSubject, firstBand: number): string =>
    `${who(subject)} best in slot gear for ${GAME}, from leveling ${firstBand} to ${MAX_BAND} to ${subject.preRaid ? 'pre-raid ' : ''}level ${MAX_BAND}: every slot ranked by the simulator, with where each item comes from.`,
  bandDescription: (subject: BisSeoSubject, band: number): string =>
    `${who(subject)} best in slot gear at ${bandWords(band).toLowerCase()} in ${GAME}: every slot ranked by the simulator, with the dungeon, quest or vendor each item comes from.`,
  /** The breadcrumb's name for the spec's BiS: "Fury Warrior BiS". */
  crumb: (subject: BisSeoSubject): string => `${who(subject)} BiS`,
} as const;
