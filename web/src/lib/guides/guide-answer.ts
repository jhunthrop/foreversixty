// web/src/lib/guides/guide-answer.ts
// The signed-in answer on /guides, a class landing and a spec guide (guides spec 2026-11-08
// §4.4): the header selector's character (`selectedCharacter`, never the account's first,
// bug B1) and its next link. On the guide of its own spec below 60 the next step is the
// talents at its level, the ladder's next point and its best in slot; at 60 its best in slot
// and the sim; anywhere else "Your guide: <Spec> <Class> ›". Pure: the mount feeds it the
// account read, the pointer and the page.
import type { Me, MeCharacter } from '../account/api';
import { answerCopy } from '../account/answer-copy';
import { addonCopy } from '../addon/copy';
import { selectedCharacter } from '../account/hero-character';
import type { CurrentCharacter } from '../current-character';
import { withRuleSpec } from '../character-spec';
import { descriptorLines } from '../character-header/descriptor';
import { classSlugFromName } from '../report/tree-sizes';
import { guidesCopy } from './copy';
import { bandForLevel, MAX_LEVEL } from './guide-character';
import classes from '../../data/classes.json';

/** What the page tells the answer about itself. */
export interface GuideAnswerPage {
  /** The class a landing or a spec guide is about; absent on /guides. */
  classSlug?: string;
  /** The spec a guide is about. */
  specSlug?: string;
  /** The guide's ladder (`guideLadderSteps`, level and talent name): the next point's line. */
  ladder?: readonly { level: number; name: string | null }[];
}

export interface AnswerLink {
  label: string;
  href: string;
}

/** The character the answer names, and what it means for the page under it. */
export interface GuideCharacter {
  name: string;
  classSlug: string;
  specSlug: string;
  className: string;
  specName: string;
  level: number | undefined;
  faction: 'alliance' | 'horde' | undefined;
  /** "Level 34" */
  levelLine: string | null;
  /** "Dwarf Survival Hunter" */
  identity: string;
}

export type GuideAnswer =
  /** Signed out, or a pointer that names no account character: no answer at all. */
  | { state: 'hidden' }
  | { state: 'none'; message: string; link: AnswerLink }
  | { state: 'unknown'; message: string; linkLabel: string }
  | { state: 'error'; message: string; retry: string }
  | {
      state: 'ready';
      character: GuideCharacter;
      /** The ladder's next point ("Next point at 35: Expose Prey"), on its own spec below 60. */
      next?: string;
      links: AnswerLink[];
    };

const SETUP_HREF = '/setup';

function classNameOf(slug: string): string | undefined {
  return classes.find((c) => c.slug === slug)?.name;
}

/** The answer's character, or undefined while its class or spec is not known. */
export function guideCharacterOf(raw: MeCharacter): GuideCharacter | undefined {
  const character = withRuleSpec(raw);
  if (character.class === undefined || character.spec === undefined) return undefined;
  const lines = descriptorLines(character);
  const classSlug = classSlugFromName(character.class);
  return {
    name: character.name,
    classSlug,
    specSlug: classSlugFromName(character.spec),
    className: classNameOf(classSlug) ?? character.class,
    specName: character.spec,
    level: character.level,
    faction: character.faction,
    levelLine: lines.level,
    identity: lines.identity,
  };
}

function linksFor(character: GuideCharacter, page: GuideAnswerPage): { links: AnswerLink[]; next?: string } {
  const own = page.classSlug === character.classSlug && page.specSlug === character.specSlug;
  const guideHref = `/guides/${character.classSlug}/${character.specSlug}`;
  const bisHref = `/bis/${character.classSlug}/${character.specSlug}`;
  if (!own) {
    return {
      links: [
        { label: guidesCopy.signedInCallout(character.specName, character.className), href: guideHref },
      ],
    };
  }
  const level = character.level ?? MAX_LEVEL;
  if (level >= MAX_LEVEL) {
    return {
      links: [
        { label: `${answerCopy.bis} ›`, href: bisHref },
        { label: `${answerCopy.sim} ›`, href: '/sim' },
      ],
    };
  }
  const point = page.ladder?.find((step) => step.level > level && step.name !== null);
  return {
    links: [
      { label: `${guidesCopy.talentsAtLevel(level)} ›`, href: '#leveling' },
      { label: `${answerCopy.bis} ›`, href: `${bisHref}?band=${bandForLevel(level)}` },
    ],
    ...(point?.name ? { next: guidesCopy.nextPointAt(point.level, point.name) } : {}),
  };
}

/** The answer for one read: `me` null is signed out; `failed` a read that threw. */
export function guideAnswerFrom(
  me: Pick<Me, 'characters' | 'main_character_key'> | null,
  current: CurrentCharacter | null,
  page: GuideAnswerPage,
  failed = false,
): GuideAnswer {
  if (failed) return { state: 'error', message: answerCopy.error, retry: answerCopy.retry };
  if (me === null) return { state: 'hidden' };
  if (me.characters.length === 0) {
    return {
      state: 'none',
      message: guidesCopy.noCharacter,
      link: { label: addonCopy.setupPageTitle, href: SETUP_HREF },
    };
  }
  const raw = selectedCharacter(current, me);
  if (raw === null) return { state: 'hidden' };
  const character = guideCharacterOf(raw);
  if (character === undefined) {
    return { state: 'unknown', message: answerCopy.unknown(raw.name), linkLabel: answerCopy.unknownLink };
  }
  return { state: 'ready', character, ...linksFor(character, page) };
}
