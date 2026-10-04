// web/src/lib/reports/device-status.ts
// Logs landing spec (2026-10-04) §4.D.1's one honest status line per paired device, built
// entirely from the existing `Device.last_seen_at` field -- no new API call, no present-
// tense "Connected" claim (that needs the heartbeat endpoint §9 names and leaves out of
// this pass). Pulled out of Account.svelte as a plain function so the two copies (seen,
// never seen) are tested without mounting the component.
import { relativeTime } from '../dates';
import type { Device } from '../account/api';

/** `{device.name} · {device.platform} · last seen {relative}` when `last_seen_at` is set,
 *  or `· paired, not seen yet` when it is null -- verbatim, spec §6. */
export function deviceStatusLine(
  device: Pick<Device, 'name' | 'platform' | 'last_seen_at'>,
  now: Date = new Date(),
): string {
  const seen =
    device.last_seen_at === null
      ? 'paired, not seen yet'
      : `last seen ${relativeTime(new Date(device.last_seen_at), now)}`;
  return `${device.name} · ${device.platform} · ${seen}`;
}
