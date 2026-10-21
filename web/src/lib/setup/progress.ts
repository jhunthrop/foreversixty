// web/src/lib/setup/progress.ts
// What /setup says under each step once it is done. Sign-in and the addon read from the
// <html> hints before any script runs (global.css); the companion has no hint, only the
// devices the account holds, so its line comes from this one pure read of GET /v1/devices.
import type { Device } from '../account/api';
import { addonCopy } from '../addon/copy';
import { relativeTime } from '../dates';

/** The newest device by last contact (or, never contacted, by creation). */
function newest(devices: readonly Device[]): Device {
  const stamp = (device: Device): number => new Date(device.last_seen_at ?? device.created_at).getTime();
  return [...devices].sort((a, b) => stamp(b) - stamp(a))[0];
}

/** The companion step's one-line done state: null until the account holds a connected device. */
export function companionDoneLine(devices: readonly Device[], now: Date = new Date()): string | null {
  if (devices.length === 0) return null;
  const device = newest(devices);
  return addonCopy.setupCompanionDoneLine(
    device.name,
    relativeTime(new Date(device.last_seen_at ?? device.created_at), now),
  );
}

/** "Windows now, macOS next." while the release carries no Mac bundle; nothing once it does. */
export function companionPlatformsLine(platforms: readonly string[]): string | null {
  if (platforms.some((platform) => platform.toLowerCase().startsWith('mac'))) return null;
  return addonCopy.setupCompanionPlatforms(platforms.join(' and '));
}
