// web/src/lib/setup/release.ts
// The companion's newest release from the site's own feed, /dl/companion/latest.json (written by
// tools/publish-feed.sh): its version, and each file's address and size, so a download button can
// link straight to the file and say what it is. The page works without it: until the manifest
// answers, a button links to the feed's latest-file redirect for the same name.
import { DOWNLOAD_PREFIX, MANIFEST_NAME, companionDownloadPath } from '../downloads/paths';
import { addonCopy } from '../addon/copy';

export const COMPANION_MANIFEST_PATH = `${DOWNLOAD_PREFIX}companion/${MANIFEST_NAME}`;
const TAG_PREFIX = 'companion-v';

export interface ReleaseFile {
  /** The file's own immutable path on this site, /dl/companion/<tag>/<name>. */
  path: string;
  size: number;
}

export interface CompanionRelease {
  version: string;
  files: Readonly<Record<string, ReleaseFile>>;
}

interface ManifestAsset {
  name?: unknown;
  browser_download_url?: unknown;
  size?: unknown;
}

function fileOf(asset: ManifestAsset): [string, ReleaseFile] | null {
  const { name, browser_download_url: url, size } = asset;
  if (typeof name !== 'string' || typeof url !== 'string' || typeof size !== 'number') return null;
  let path: string;
  try {
    path = new URL(url).pathname; // the same file on whichever host serves this page
  } catch {
    return null;
  }
  return path.startsWith(`${DOWNLOAD_PREFIX}companion/`) ? [name, { path, size }] : null;
}

/** The manifest's companion release, or null for anything that is not one. */
export function parseRelease(manifest: unknown): CompanionRelease | null {
  if (typeof manifest !== 'object' || manifest === null) return null;
  const { tag_name: tag, assets } = manifest as { tag_name?: unknown; assets?: unknown };
  if (typeof tag !== 'string' || !tag.startsWith(TAG_PREFIX) || !Array.isArray(assets)) return null;
  const files = assets.flatMap((asset: ManifestAsset) => {
    const file = fileOf(asset);
    return file === null ? [] : [file];
  });
  return { version: tag.slice(TAG_PREFIX.length), files: Object.fromEntries(files) };
}

/** Reads the manifest; null when it does not answer or is not a companion release. */
export async function fetchCompanionRelease(fetchFn: typeof fetch = fetch): Promise<CompanionRelease | null> {
  try {
    const response = await fetchFn(COMPANION_MANIFEST_PATH);
    return response.ok ? parseRelease(await response.json()) : null;
  } catch {
    return null;
  }
}

/** Where a build's button links: the release's own file once the manifest answered, else the latest-file redirect. */
export function downloadHref(asset: string, release: CompanionRelease | null): string {
  return release?.files[asset]?.path ?? companionDownloadPath(asset);
}

const MB = 1_000_000;

/** "16.8 MB": decimal megabytes, as the operating systems' own download views count them. */
export function formatSize(bytes: number): string {
  return `${(bytes / MB).toFixed(1)} MB`;
}

/** "Version 0.2.37 · 16.8 MB" under a button, or null until the manifest names the file. */
export function releaseMeta(asset: string, release: CompanionRelease | null): string | null {
  const file = release?.files[asset];
  if (release === null || file === undefined) return null;
  return addonCopy.setupCompanionFileMeta(release.version, formatSize(file.size));
}
