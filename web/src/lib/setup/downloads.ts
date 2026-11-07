// web/src/lib/setup/downloads.ts
// Which companion build a visitor is offered. companion.json lists the builds the release feed
// carries (the owner adds the Mac entry once the first signed and notarized Mac build is in the
// feed), and this is the one place that turns a visitor's platform into one of them: Windows gets
// the Windows build, a Mac gets the Mac build when there is one and the export route when not.
import companion from '../../data/companion.json' with { type: 'json' };
import type { VisitorPlatform } from './visitor-platform';

export interface CompanionDownload {
  /** The word on the button: "Download for Windows", "Download for Mac". */
  platform: string;
  /** The file's name in the feed's companion manifest. */
  asset: string;
  /** The signing and first-run line under the download. */
  firstRun: string;
}

export const COMPANION_DOWNLOADS: readonly CompanionDownload[] = companion.downloads;

/** Every build's platform word, in companion.json's order: what lib/setup/steps.ts reads. */
export const COMPANION_PLATFORMS: readonly string[] = COMPANION_DOWNLOADS.map((entry) => entry.platform);

/** "Mac" or "macOS": the platform word of a Mac build. */
export function isMacPlatform(platform: string): boolean {
  return platform.toLowerCase().startsWith('mac');
}

export function isMacDownload(download: CompanionDownload): boolean {
  return isMacPlatform(download.platform);
}

/** The build this visitor downloads, or null on a Mac while the feed carries no Mac build. */
export function downloadFor(
  platform: VisitorPlatform,
  downloads: readonly CompanionDownload[] = COMPANION_DOWNLOADS,
): CompanionDownload | null {
  if (platform === 'mac') return downloads.find(isMacDownload) ?? null;
  return downloads.find((download) => !isMacDownload(download)) ?? downloads[0] ?? null;
}
