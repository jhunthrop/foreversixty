// web/src/lib/planner/send-build.ts
// "Send this build to the addon" on the planner's character card: saves the build when it has
// no id yet, checks the account has a companion to carry it, and queues it in the addon inbox.
// The pure half (`sendActionFor`) turns where the send stands into the card's one action.
import { AccountError, queueBuildToAddon, refreshDevices, type MeCharacter } from '../account/api';
import type { HeaderAction } from '../character-header/types';
import { plannerCharacterCardCopy as copy } from './copy';
import { saveBuild } from './share';
import type { BuildDraft } from './types';

export type SendOutcome =
  | { kind: 'sent'; buildId: string }
  | { kind: 'sign-in' }
  | { kind: 'no-device' }
  | { kind: 'failed' };

export interface SendDeps {
  save: typeof saveBuild;
  devices: typeof refreshDevices;
  queue: typeof queueBuildToAddon;
}

const LIVE: SendDeps = { save: saveBuild, devices: refreshDevices, queue: queueBuildToAddon };

/**
 * Sends `draft` to the addon inbox of `characterKey`. A `known` build id (this exact build was
 * saved by an earlier press) skips the save; the API dedupes an identical save anyway.
 */
export async function sendBuildToAddon(
  draft: BuildDraft,
  characterKey: string,
  known: string | null = null,
  deps: SendDeps = LIVE,
): Promise<SendOutcome> {
  try {
    // First, because a 401 here is the cleanest "not signed in" and an empty list means the
    // inbox would be written and never read.
    if ((await deps.devices()).length === 0) return { kind: 'no-device' };
    let buildId = known;
    if (buildId === null) {
      const saved = await deps.save(draft);
      if (!saved.ok) return { kind: 'failed' };
      buildId = saved.build.id;
    }
    await deps.queue(buildId, characterKey);
    return { kind: 'sent', buildId };
  } catch (error) {
    if (error instanceof AccountError && error.status === 401) return { kind: 'sign-in' };
    return { kind: 'failed' };
  }
}

export type SendState =
  | { kind: 'idle' }
  | { kind: 'sending' }
  | { kind: 'sent' }
  | { kind: 'sign-in' }
  | { kind: 'no-device' }
  | { kind: 'failed' };

/** The card's action: a real send when a character is current, otherwise a way to get the addon. */
export function sendActionFor(character: MeCharacter | null, state: SendState): HeaderAction {
  if (character === null) {
    return { label: copy.getTheAddon, href: copy.getTheAddonHref };
  }
  const base = { kind: 'picker' as const, href: copy.getTheAddonHref };
  switch (state.kind) {
    case 'sending':
      return { ...base, label: copy.sending, disabled: true };
    case 'sent':
      return {
        ...base,
        label: copy.sent,
        disabled: true,
        note: { tone: 'ok' as const, text: copy.sentLine(character.name) },
      };
    case 'sign-in':
      return {
        ...base,
        label: copy.sendBuildToAddon,
        note: { tone: 'error' as const, text: copy.signIn, link: { label: copy.signInLink, href: copy.signInHref } },
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
      return { ...base, label: copy.sendBuildToAddon };
  }
}
