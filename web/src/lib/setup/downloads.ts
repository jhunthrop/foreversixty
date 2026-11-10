// web/src/lib/setup/downloads.ts
// The companion builds /setup offers, and which one is a visitor's own. companion.json lists the
// builds the release feed carries (the feed's Worker serves only those names, lib/downloads/serve.ts);
// this turns a visitor's system (lib/setup/visitor-platform.ts) into one of them, and the feed's
// manifest (lib/setup/release.ts) adds the version and size.
import companion from '../../data/companion.json' with { type: 'json' };
import type { VisitorPlatform } from './visitor-platform';

export type CompanionOs = 'windows' | 'mac';

export interface CompanionDownload {
  /** The system the build runs on. */
  os: CompanionOs;
  /** The word on the button: "Download for Windows", "Download for Mac". */
  platform: string;
  /** The file's name in the feed's companion manifest. */
  asset: string;
  /** The name the browser saves it as (the Worker's Content-Disposition, the link's download). */
  saveAs: string;
  /** The signing and first-run line under the download. */
  firstRun: string;
}

export const COMPANION_DOWNLOADS: readonly CompanionDownload[] = companion.downloads.map((entry) => {
  if (entry.os !== 'windows' && entry.os !== 'mac') {
    throw new Error(`companion.json: "${entry.os}" is not windows or mac`);
  }
  return { ...entry, os: entry.os };
});

/** The visitor's own build: null on a phone and when the system is unknown (both are offered). */
export function downloadFor(
  platform: VisitorPlatform,
  downloads: readonly CompanionDownload[] = COMPANION_DOWNLOADS,
): CompanionDownload | null {
  if (platform !== 'windows' && platform !== 'mac') return null;
  return downloads.find((download) => download.os === platform) ?? null;
}

/** Every build but the visitor's own: the "Other platforms" list, or all of them when there is no own build. */
export function otherDownloads(
  platform: VisitorPlatform,
  downloads: readonly CompanionDownload[] = COMPANION_DOWNLOADS,
): CompanionDownload[] {
  const own = downloadFor(platform, downloads);
  return downloads.filter((download) => download !== own);
}
