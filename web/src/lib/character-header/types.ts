// web/src/lib/character-header/types.ts
// The shared character header's page-facing API (see CharacterHeader.svelte's header comment):
// what a page hands in. The identity and every state's presentation are fixed by the component;
// a page supplies only its answer's words and its one action.
import type { MeCharacter } from '../account/api';

export type HeaderEmphasis = 'primary' | 'secondary';
/** `band`: full width, three columns from 1024 px. `card`: always stacked (a side card on a tool page). */
export type HeaderLayout = 'band' | 'card';
/** What the header shows for a visitor with no character: the shared band, or nothing. */
export type SignedOutMode = 'band' | 'none';

/** The page's one action. `link` opens a page or jumps to an anchor; `picker` runs the page's `onPick`. */
export type HeaderAction =
  | { kind?: 'link'; label: string; href: string }
  | {
      kind: 'picker';
      label: string;
      href: string;
      /** The button is shown but cannot be pressed (a send that already went out). */
      disabled?: boolean;
      /** A line under the button: what the press did, or why it could not. */
      note?: HeaderNote;
      /** A second, quieter button under the first: a link out ("Sim this build") or a press
       *  (`onSecondary`, the "Cancel" of a confirm). */
      secondary?: HeaderSecondary;
    };

/** The picker action's second button. A link carries `href`; without one it calls `onSecondary`. */
export interface HeaderSecondary {
  label: string;
  href?: string;
  testid?: string;
}

/** The one status line a picker action can carry under its button. */
export interface HeaderNote {
  text: string;
  tone: 'ok' | 'error' | 'muted';
  /** A second, quieter line under the first ("The companion checks every 10 minutes ..."). */
  detail?: string;
  /** A link ending the line ("Sign in", "Run the companion"). */
  link?: { label: string; href: string };
}

export interface ActionContext {
  character: MeCharacter;
  /** The character is below the page's level floor, so the action points at the first band. */
  belowLevelFloor: boolean;
}

export interface AnswerFigure {
  value: string;
  /** The muted unit after the value; the component adds the space between them. */
  unit?: string;
  /** A parse-ladder colour, only where the page's board uses that ladder. */
  color?: string;
}

export interface AnswerLink {
  label: string;
  href: string;
}

export interface AnswerNote {
  text: string;
  /** At most one gold inline link; its arrow never wraps. */
  link?: AnswerLink;
}

export interface FiguresAnswer {
  kind: 'figures';
  label: string;
  /** The label when the sync is stale and the page's answer reads gear. */
  staleLabel?: string;
  figures: readonly [AnswerFigure] | readonly [AnswerFigure, AnswerFigure];
  note?: AnswerNote;
  /** The second fact, only where a page has two (Logs). */
  secondNote?: { text: string; href: string };
  /** 0 to 1; only when the figure is a fraction of a whole. */
  bar?: number;
}

/** A figure line (one or two figures at the same size), or one sentence, or nothing to show. */
export type HeaderAnswer = FiguresAnswer | { kind: 'sentence'; text: string } | { kind: 'none' };

export interface AnswerStrings {
  /** Shown instead of the answer below the level floor. */
  belowLevelFloor?: string;
  /** Shown instead of the answer when no gear is on file, on a page that reads gear. */
  noGear?: { sentence: string; link: AnswerLink };
  /** The sentence when the answer failed to load; the component adds `Try again`. */
  failed: string;
}

export interface AnswerSpec {
  /** The answer starts at this level (BiS and home: 20). Omit for a page with no floor. */
  levelFloor?: number;
  /** The answer reads worn gear, so a character with no gear on file gets `strings.noGear`. */
  needsGear?: boolean;
  strings: AnswerStrings;
  /** The loading box reserves a bar under the figure, for a page whose answer draws one. */
  placeholderBar?: boolean;
  /** Fetches the answer for a character that passed the floor and gear checks. A rejection is the failed state. */
  load(character: MeCharacter): Promise<HeaderAnswer>;
}
