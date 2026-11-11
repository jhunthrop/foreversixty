// web/src/lib/setup/release.ts
// The companion's newest release from the site's own feed, /dl/companion/latest.json (written by
// tools/publish-feed.sh): its version, and each file's address and size, so a download button can
// link straight to the file and say what it is. The page works without it: until the manifest
// answers, a button links to the feed's latest-file redirect for the same name. The addon's own
// manifest, /dl/addon/latest.json, gives the newest addon version, so /setup can say an installed
// addon is behind; without it no such claim is made.
import {
  DOWNLOAD_PREFIX,
  FAMILIES,
  MANIFEST_NAME,
  companionDownloadPath,
  type Family,
} from '../downloads/paths';
import { addonCopy } from '../addon/copy';

const manifestPath = (family: Family): string => `${DOWNLOAD_PREFIX}${family}/${MANIFEST_NAME}`;
export const COMPANION_MANIFEST_PATH = manifestPath('companion');
export const ADDON_MANIFEST_PATH = manifestPath('addon');

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

/** The version a manifest's tag names for `family` ("addon-v2.0.37" is 2.0.37), or null. */
export function tagVersion(manifest: unknown, family: Family): string | null {
  if (typeof manifest !== 'object' || manifest === null) return null;
  const { tag_name: tag } = manifest as { tag_name?: unknown };
  const prefix = FAMILIES[family];
  if (typeof tag !== 'string' || !tag.startsWith(prefix) || tag.length === prefix.length) return null;
  return tag.slice(prefix.length);
}

/** The manifest's companion release, or null for anything that is not one. */
export function parseRelease(manifest: unknown): CompanionRelease | null {
  const version = tagVersion(manifest, 'companion');
  const { assets } = (manifest ?? {}) as { assets?: unknown };
  if (version === null || !Array.isArray(assets)) return null;
  const files = assets.flatMap((asset: ManifestAsset) => {
    const file = fileOf(asset);
    return file === null ? [] : [file];
  });
  return { version, files: Object.fromEntries(files) };
}

/** One manifest read: its parsed body, or null when it does not answer. */
async function readManifest<T>(
  path: string,
  parse: (manifest: unknown) => T | null,
  fetchFn: typeof fetch,
): Promise<T | null> {
  try {
    const response = await fetchFn(path);
    return response.ok ? parse(await response.json()) : null;
  } catch {
    return null;
  }
}

/** Reads the manifest; null when it does not answer or is not a companion release. */
export function fetchCompanionRelease(fetchFn: typeof fetch = fetch): Promise<CompanionRelease | null> {
  return readManifest(COMPANION_MANIFEST_PATH, parseRelease, fetchFn);
}

/** The newest addon release's version; null when the feed does not answer with one. */
export function fetchAddonVersion(fetchFn: typeof fetch = fetch): Promise<string | null> {
  return readManifest(ADDON_MANIFEST_PATH, (manifest) => tagVersion(manifest, 'addon'), fetchFn);
}

const DOTTED = /^\d+(\.\d+)*$/;

/** Whether dotted version `latest` is past `installed` (2.0.37 past 2.0.9). False when either is not dotted numbers. */
export function versionBehind(installed: string, latest: string): boolean {
  if (!DOTTED.test(installed) || !DOTTED.test(latest)) return false;
  const a = installed.split('.').map(Number);
  const b = latest.split('.').map(Number);
  for (let i = 0; i < Math.max(a.length, b.length); i += 1) {
    const diff = (b[i] ?? 0) - (a[i] ?? 0);
    if (diff !== 0) return diff > 0;
  }
  return false;
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
