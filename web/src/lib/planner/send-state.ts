// web/src/lib/planner/send-state.ts
// The pure half of "Send this build to the addon": where the send stands, and the card's one
// action for that state. Kept apart from `send-build.ts` so the planner island ships only
// this at boot; the account API, the save and the queue load with the first press.
import type { MeCharacter } from '../account/api';
import type { HeaderAction } from '../character-header/types';
import { classDisplayName } from '../sim/spec-label';
import { plannerCharacterCardCopy as copy } from './copy';

export type SendState =
  | { kind: 'idle' }
  /** Another player's build asked "Send X's build to Y?" and waits for the yes. */
  | { kind: 'confirming' }
  | { kind: 'sending' }
  | { kind: 'sent' }
  | { kind: 'sign-in' }
  | { kind: 'no-device' }
  /** `reason`: what the API said when it refused the save (its first field's message). */
  | { kind: 'failed'; reason?: string };

/** What the planner knows about the build on screen, beyond its points. */
export interface SendBuild {
  /** The build's class slug ("paladin"). */
  classSlug?: string;
  /** Its display name ("Paladin"). */
  className?: string;
  /** Set when the build is somebody else's: whose it is ('' for a linked build with no name). */
  owner?: string;
  /** The planner's "Sim this build" link, offered under the send button. */
  simHref?: string;
}

/**
 * The card's action: a real send when a character is current and the build is that character's
 * own (same class, opened from it or edited from it); a disabled button naming the mismatch for
 * another class's build; a confirmed send for another player's build of the same class;
 * otherwise a way to get the addon.
 */
export function sendActionFor(
  character: MeCharacter | null,
  state: SendState,
  pointsSpent = 1,
  characterSpent: number | null = null,
  build: SendBuild = {},
): HeaderAction {
  if (character === null) {
    return { label: copy.getTheAddon, href: copy.getTheAddonHref };
  }
  const secondary =
    build.simHref === undefined
      ? {}
      : { secondary: { label: copy.simThisBuild, href: build.simHref, testid: 'planner-sim-link' } };
  const base = { kind: 'picker' as const, href: copy.getTheAddonHref, ...secondary };
  const characterClass = character.class?.toLowerCase();
  // Never send another class's build to a character: the button names why and stays off.
  if (characterClass !== undefined && build.classSlug !== undefined && build.classSlug !== characterClass) {
    return {
      ...base,
      label: copy.otherClass(
        build.className ?? build.classSlug,
        character.name,
        classDisplayName(characterClass),
      ),
      disabled: true,
    };
  }
  const others = build.owner !== undefined;
  // An empty build is never sent. A partial one is: a player may send a plan on purpose.
  if (pointsSpent <= 0 && state.kind !== 'sending') {
    return {
      ...base,
      label: copy.sendBuildTo(character.name),
      disabled: true,
      note: { tone: 'muted' as const, text: copy.spendFirst },
    };
  }
  switch (state.kind) {
    case 'confirming':
      return {
        ...base,
        label: copy.confirmYes,
        secondary: { label: copy.confirmNo, testid: 'planner-send-cancel' },
        note: { tone: 'muted' as const, text: copy.confirmQuestion(build.owner ?? '', character.name) },
      };
    case 'sending':
      return { ...base, label: copy.sending, disabled: true };
    case 'sent':
      return {
        ...base,
        label: copy.sent,
        disabled: true,
        note: { tone: 'ok' as const, text: copy.sentLine(character.name), detail: copy.sentDetail },
      };
    case 'sign-in':
      return {
        ...base,
        label: copy.sendBuildTo(character.name),
        note: {
          tone: 'error' as const,
          text: copy.signIn,
          link: { label: copy.signInLink, href: copy.signInHref },
        },
      };
    case 'no-device':
      return {
        ...base,
        label: copy.sendBuildTo(character.name),
        note: {
          tone: 'error' as const,
          text: copy.noDevice,
          link: { label: copy.noDeviceLink, href: copy.noDeviceHref },
        },
      };
    case 'failed':
      return {
        ...base,
        label: copy.sendBuildTo(character.name),
        note: {
          tone: 'error' as const,
          text: state.reason === undefined ? copy.failed : copy.refused(state.reason),
        },
      };
    case 'idle':
      return others
        ? { ...base, label: copy.sendOthersBuildTo(character.name, characterSpent) }
        : {
            ...base,
            label: copy.sendBuildTo(character.name),
            ...shorterNote(character, pointsSpent, characterSpent),
          };
  }
}

/** Points a character has spent, from its synced talents; null when the site holds none. */
export function characterPointsSpent(character: MeCharacter | null): number | null {
  return character?.build?.talents?.points.reduce((sum, n) => sum + n, 0) ?? null;
}

/** The muted "this build is shorter than the character's spend" note, or nothing to add. */
function shorterNote(
  character: MeCharacter,
  pointsSpent: number,
  characterSpent: number | null,
): { note?: { tone: 'muted'; text: string } } {
  if (characterSpent === null || pointsSpent >= characterSpent) return {};
  return {
    note: { tone: 'muted', text: copy.shorterThanSpend(pointsSpent, character.name, characterSpent) },
  };
}
