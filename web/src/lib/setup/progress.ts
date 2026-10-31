// web/src/lib/setup/progress.ts
// What /setup says under each step once the account is signed in, and which of the three
// progress chips are done. Pure reads of GET /v1/me (the account's name and the characters the
// addon has synced) and GET /v1/devices (the companions signed in), so the page never claims a
// step the account does not show.
import type { Device, Me, MeCharacter } from '../account/api';
import { addonCopy } from '../addon/copy';
import { relativeTime } from '../dates';

/** What the three steps know about the visitor: `me` null is signed out, `devices` null is not read yet. */
export interface SetupState {
  me: Me | null;
  devices: readonly Device[] | null;
}

/** A step's line under its action: `done` draws the check, otherwise the line is muted waiting text. */
export interface StepLine {
  done: boolean;
  text: string;
}

/** The newest device by last contact (or, never contacted, by creation). */
function newest(devices: readonly Device[]): Device {
  const stamp = (device: Device): number => new Date(device.last_seen_at ?? device.created_at).getTime();
  return [...devices].sort((a, b) => stamp(b) - stamp(a))[0];
}

/** Step 1: "Signed in as hunthrop#1894"; a name-less account (an email link) says its email. */
export function signInLine(me: Me | null): StepLine | null {
  if (me === null) return null;
  const name = me.user.battletag ?? me.user.email ?? 'your account';
  return { done: true, text: addonCopy.setupSignedInAs(name) };
}

/** The character the addon synced most recently, with when: null while no character carries an addon build. */
export function newestAddonCharacter(
  characters: readonly MeCharacter[],
): { character: MeCharacter; capturedAt: string } | null {
  let best: { character: MeCharacter; capturedAt: string } | null = null;
  for (const character of characters) {
    if (character.build?.source !== 'addon') continue;
    const capturedAt = character.build.captured_at;
    if (best === null || new Date(capturedAt).getTime() > new Date(best.capturedAt).getTime()) {
      best = { character, capturedAt };
    }
  }
  return best;
}

/** Step 2: "Bow Jackzon arrived from the addon 1 hour ago", or what to do to make one arrive. */
export function addonLine(me: Me | null, now: Date = new Date()): StepLine | null {
  if (me === null) return null;
  const arrived = newestAddonCharacter(me.characters);
  if (arrived === null) return { done: false, text: addonCopy.setupAddonWaiting };
  return {
    done: true,
    text: addonCopy.setupAddonArrived(
      arrived.character.name,
      relativeTime(new Date(arrived.capturedAt), now),
    ),
  };
}

/** Step 3: "Companion signed in on JH-DESKTOP · last seen 3 minutes ago", or that none is. */
export function companionLine(
  me: Me | null,
  devices: readonly Device[] | null,
  now: Date = new Date(),
): StepLine | null {
  if (me === null || devices === null) return null;
  if (devices.length === 0) return { done: false, text: addonCopy.setupCompanionWaiting };
  const device = newest(devices);
  return {
    done: true,
    text: addonCopy.setupCompanionLine(
      device.name,
      relativeTime(new Date(device.last_seen_at ?? device.created_at), now),
    ),
  };
}

/** Which <html> hint each step's chip and tag follow until the reads answer: the session cookie,
 *  the stored character, or none (the companion has no hint). */
export const STEP_HINTS = ['session', 'pointer', 'none'] as const;

/** The three chips: signed in, an addon character arrived, a companion signed in; null until it is known. */
export function chipsDone(state: SetupState): [boolean, boolean, boolean] {
  return [
    state.me !== null,
    state.me !== null && newestAddonCharacter(state.me.characters) !== null,
    state.devices !== null && state.devices.length > 0,
  ];
}

/** "Windows. Mac is not supported yet." while the release carries no Mac bundle; nothing once it does. */
export function companionPlatformsLine(platforms: readonly string[]): string | null {
  if (platforms.some((platform) => platform.toLowerCase().startsWith('mac'))) return null;
  return addonCopy.setupCompanionPlatforms(platforms.join(' and '));
}
