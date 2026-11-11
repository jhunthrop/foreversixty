// web/src/lib/setup/progress.ts
// What /setup says under each step once the account is signed in (lib/setup/steps.ts draws the
// rows from these lines). Pure reads of GET /v1/me (the account's name and the characters the
// addon has synced), GET /v1/devices (the companions signed in, and the addons each found in its
// game folder) and the addon feed's newest version, so the page never claims a step the account
// does not show.
import type { Device, Me, MeCharacter } from '../account/api';
import { addonCopy } from '../addon/copy';
import { relativeTime } from '../dates';
import { versionBehind } from './release';

/** The folder whose TOC carries the addon's version (the companion's addonupdater.Addon.Main). */
const ADDON_FOLDER = 'ForeverSixty';

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

/** The addon as the newest companion that has scanned for it found it: on which PC, at which version.
 *  `installed` is null when every companion that scanned found none; the whole answer is null while
 *  no companion has scanned (an older companion, or none signed in), so nothing is claimed. */
export function companionAddon(
  devices: readonly Device[] | null,
): { device: Device; version: string | null } | null {
  const scanned = (devices ?? []).filter((device) => device.addons !== undefined);
  if (scanned.length === 0) return null;
  const withAddon = scanned.filter((device) => device.addons?.some((a) => a.folder === ADDON_FOLDER));
  if (withAddon.length === 0) return { device: newest(scanned), version: null };
  const device = newest(withAddon);
  const version = device.addons?.find((a) => a.folder === ADDON_FOLDER)?.version ?? '';
  return { device, version };
}

/** A step line plus, for an installed addon behind the feed, the line that says so, and whether a
 *  companion has looked in the game folder at all. */
export interface AddonStatus extends StepLine {
  behind?: string;
  scanned: boolean;
}

/** The addon step: installed on a PC (and whether behind `latest`), else arrived from the addon,
 *  else what it does. Never a fault: a companion that found no addon is a plain to-do. */
export function addonLine(
  me: Me | null,
  devices: readonly Device[] | null,
  latest: string | null,
  now: Date = new Date(),
): AddonStatus | null {
  if (me === null) return null;
  const found = companionAddon(devices);
  if (found !== null && found.version !== null) {
    const status: AddonStatus = {
      done: true,
      text: addonCopy.setupAddonInstalled(found.device.name, found.version),
      scanned: true,
    };
    return latest !== null && versionBehind(found.version, latest)
      ? { ...status, behind: addonCopy.setupAddonBehind(latest) }
      : status;
  }
  const arrived = newestAddonCharacter(me.characters);
  if (arrived !== null) {
    return {
      done: true,
      text: addonCopy.setupAddonArrived(
        arrived.character.name,
        relativeTime(new Date(arrived.capturedAt), now),
      ),
      scanned: found !== null,
    };
  }
  return found === null
    ? { done: false, text: addonCopy.setupAddonNote, scanned: false }
    : { done: false, text: addonCopy.setupAddonNotOn(found.device.name), scanned: true };
}

/** The companion step: "Companion signed in on JH-DESKTOP · last seen 3 minutes ago", or that none is. */
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
