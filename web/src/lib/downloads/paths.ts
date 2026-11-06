// web/src/lib/downloads/paths.ts
// foreversixty.gg/dl/...: the one address every download on the site links to. The page never
// names where a file is hosted; the Worker answers each path with a redirect (redirect.ts), so
// moving the files is a Worker change, not a page change.
export const DOWNLOAD_PREFIX = '/dl/';

/** The addon's zip, for a hand install. */
export const ADDON_DOWNLOAD_PATH = `${DOWNLOAD_PREFIX}addon`;

/** The nightly data addon, for a hand install. */
export const DATA_DOWNLOAD_PATH = `${DOWNLOAD_PREFIX}data`;

const COMPANION_PREFIX = `${DOWNLOAD_PREFIX}companion/`;

/** One companion build, by the asset name companion.json lists. */
export function companionDownloadPath(asset: string): string {
  return `${COMPANION_PREFIX}${asset}`;
}

export type DownloadTarget = { kind: 'addon' } | { kind: 'data' } | { kind: 'companion'; asset: string };

/** What a /dl/ path asks for, or null when it names nothing the site offers. */
export function parseDownloadPath(
  pathname: string,
  companionAssets: readonly string[],
): DownloadTarget | null {
  if (pathname === ADDON_DOWNLOAD_PATH) return { kind: 'addon' };
  if (pathname === DATA_DOWNLOAD_PATH) return { kind: 'data' };
  if (!pathname.startsWith(COMPANION_PREFIX)) return null;
  const asset = pathname.slice(COMPANION_PREFIX.length);
  return companionAssets.includes(asset) ? { kind: 'companion', asset } : null;
}
