// web/src/lib/planner/send-state.ts
// The pure half of "Send this build to the addon": where the send stands, and the card's one
// action for that state. Kept apart from `send-build.ts` so the planner island ships only
// this at boot; the account API, the save and the queue load with the first press.
import type { MeCharacter } from '../account/api';
import type { HeaderAction } from '../character-header/types';
import { plannerCharacterCardCopy as copy } from './copy';

export type SendState =
  | { kind: 'idle' }
  | { kind: 'sending' }
  | { kind: 'sent' }
  | { kind: 'sign-in' }
  | { kind: 'no-device' }
  | { kind: 'failed' };

/** The card's action: a real send when a character is current, otherwise a way to get the addon. */
export function sendActionFor(
  character: MeCharacter | null,
  state: SendState,
  pointsSpent = 1,
  characterSpent: number | null = null,
): HeaderAction {
  if (character === null) {
    return { label: copy.getTheAddon, href: copy.getTheAddonHref };
  }
  const base = { kind: 'picker' as const, href: copy.getTheAddonHref };
  // An empty build is never sent. A partial one is: a player may send a plan on purpose.
  if (pointsSpent <= 0 && state.kind !== 'sending') {
    return {
      ...base,
      label: copy.sendBuildToAddon,
      disabled: true,
      note: { tone: 'muted' as const, text: copy.spendFirst },
    };
  }
  switch (state.kind) {
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
        label: copy.sendBuildToAddon,
        note: {
          tone: 'error' as const,
          text: copy.signIn,
          link: { label: copy.signInLink, href: copy.signInHref },
        },
      };
    case 'no-device':
      return {
        ...base,
        label: copy.sendBuildToAddon,
        note: {
          tone: 'error' as const,
          text: copy.noDevice,
          link: { label: copy.noDeviceLink, href: copy.noDeviceHref },
        },
      };
    case 'failed':
      return { ...base, label: copy.sendBuildToAddon, note: { tone: 'error' as const, text: copy.failed } };
    case 'idle':
      return {
        ...base,
        label: copy.sendBuildToAddon,
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
