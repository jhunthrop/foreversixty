// web/src/lib/bis/landing-answer-view.ts
// The /bis landing's head answer as one pure decision (design/specs/2026-11-07-bis-landing.md
// section 3): the header selector's character (crest, name, level and identity), then how many of
// its range's best-in-slot pieces it wears as a Figure over a Meter, or the one sentence in its
// place (under 20, no gear on file, no list, a count that failed), then the one sentence-case
// button to its own spec at its own range. The mount (landing-answer-mount.ts) reads the account
// and the count and paints what this returns.
import type { Me, MeCharacter } from '../account/api';
import { answerCopy } from '../account/answer-copy';
import { selectedCharacter } from '../account/hero-character';
import type { CurrentCharacter } from '../current-character';
import { descriptorLines } from '../character-header/descriptor';
import { isStaleAddonSync } from '../character-header/sync';
import type { AnswerPlan } from '../character-header/answer-plan';
import type { HeaderAction, HeaderAnswer } from '../character-header/types';
import { classSlugFromName } from '../report/tree-sizes';
import { landingCopy } from './landing-copy';

export interface AnswerLink {
  label: string;
  href: string;
}

export interface AnswerCharacter {
  name: string;
  /** Lower-case class slug: the crest and the class colour. */
  classSlug: string;
  /** "Hunter": the colour lookup's key. */
  className: string;
  /** "Level 24 · Dwarf Beast Mastery Hunter"; either half alone when the other is not known. */
  line: string;
}

export type AnswerBody =
  | {
      kind: 'count';
      worn: number;
      unit: string;
      /** The caption under the figure; the stale wording when the addon sync is old. */
      caption: string;
      /** 0 to 1. */
      bar: number;
      /** "Trinkets have no pick at this level range." */
      note?: string;
    }
  | { kind: 'sentence'; text: string; link?: AnswerLink }
  | { kind: 'failed'; text: string; retry: string };

export type LandingAnswer =
  /** Signed out, or a pointer that names no account character: no answer column. */
  | { state: 'hidden' }
  | { state: 'none'; message: string; link: AnswerLink }
  | { state: 'error'; message: string; retry: string }
  | { state: 'ready'; character: AnswerCharacter; body: AnswerBody; action: AnswerLink };

/** Which character the answer names: null signed out or when the pointer names none. */
export function landingAnswerCharacter(
  me: Pick<Me, 'characters' | 'main_character_key'> | null,
  current: CurrentCharacter | null,
): MeCharacter | null {
  return me === null ? null : selectedCharacter(current, me);
}

export function answerCharacterOf(character: MeCharacter): AnswerCharacter {
  const lines = descriptorLines(character);
  const className = character.class ?? '';
  return {
    name: character.name,
    classSlug: className === '' ? '' : classSlugFromName(className),
    className,
    line: [lines.level, lines.identity].filter((part) => part !== null && part !== '').join(' · '),
  };
}

/** The body for a plan that never fetched (under 20, no gear) or for what the fetch returned. */
export function answerBodyOf(
  plan: AnswerPlan,
  fetched: HeaderAnswer | 'failed' | undefined,
  stale: boolean,
): AnswerBody {
  if (plan === 'below-floor' || plan === 'load-below-floor')
    return { kind: 'sentence', text: landingCopy.under20 };
  if (plan === 'no-gear') {
    return {
      kind: 'sentence',
      text: landingCopy.noSync,
      link: { label: landingCopy.noSyncLink, href: landingCopy.noSyncHref },
    };
  }
  if (fetched === undefined || fetched === 'failed') {
    return { kind: 'failed', text: landingCopy.countFailed, retry: landingCopy.retry };
  }
  if (fetched.kind === 'figures') {
    const [figure] = fetched.figures;
    return {
      kind: 'count',
      worn: Number(figure.value),
      unit: figure.unit ?? '',
      caption: stale ? (fetched.staleLabel ?? fetched.label) : fetched.label,
      bar: fetched.bar ?? 0,
      ...(fetched.note === undefined ? {} : { note: fetched.note.text }),
    };
  }
  return { kind: 'sentence', text: fetched.kind === 'sentence' ? fetched.text : landingCopy.countFailed };
}

export interface ReadyInput {
  character: MeCharacter;
  plan: AnswerPlan;
  fetched?: HeaderAnswer | 'failed';
  action: HeaderAction;
  now: Date;
}

export function readyAnswer({ character, plan, fetched, action, now }: ReadyInput): LandingAnswer {
  return {
    state: 'ready',
    character: answerCharacterOf(character),
    body: answerBodyOf(plan, fetched, isStaleAddonSync(character, now)),
    action: { label: action.label, href: action.href },
  };
}

/** The answer before any character is chosen: signed out, no characters, or the read failed. */
export function emptyAnswer(me: Pick<Me, 'characters'> | null, failed: boolean): LandingAnswer {
  if (failed) return { state: 'error', message: answerCopy.error, retry: answerCopy.retry };
  if (me === null) return { state: 'hidden' };
  if (me.characters.length === 0) {
    return {
      state: 'none',
      message: landingCopy.noCharacter,
      link: { label: landingCopy.setupLink, href: landingCopy.setupHref },
    };
  }
  return { state: 'hidden' };
}
