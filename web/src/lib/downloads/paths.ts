// web/src/lib/downloads/paths.ts
// foreversixty.gg/dl/...: the one address every download on the site links to, and the
// release feed the companion updates from. The page never names where a file is stored; the
// Worker answers each path (serve.ts), so moving the files is a Worker change, not a page one.
//
//   /dl/<family>/latest.json          the family's manifest (docs/runbooks/repo-private.md has the shape)
//   /dl/<family>/<tag>/<name>         one published file, immutable
//   /dl/<family>/<name>               the latest release's file of that name (a redirect)
//   /dl/addon, /dl/data               the latest zip of the family (a redirect)
export const DOWNLOAD_PREFIX = '/dl/';

/** Every family the feed publishes, by its key, with the tag prefix its releases carry. */
export const FAMILIES = {
  companion: 'companion-v',
  addon: 'addon-v',
  data: 'data-v',
} as const;

export type Family = keyof typeof FAMILIES;

export const MANIFEST_NAME = 'latest.json';

/** The addon's zip, for a hand install. */
export const ADDON_DOWNLOAD_PATH = `${DOWNLOAD_PREFIX}addon`;

/** The nightly data addon's zip, for a hand install. */
export const DATA_DOWNLOAD_PATH = `${DOWNLOAD_PREFIX}data`;

/** The latest release's file of this name, e.g. one companion build by the name companion.json lists. */
export function latestFilePath(family: Family, name: string): string {
  return `${DOWNLOAD_PREFIX}${family}/${name}`;
}

/** One companion build, by the asset name companion.json lists. */
export function companionDownloadPath(asset: string): string {
  return latestFilePath('companion', asset);
}

/** A published file's own, immutable path. */
export function releaseFilePath(family: Family, tag: string, name: string): string {
  return `${DOWNLOAD_PREFIX}${family}/${tag}/${name}`;
}

/** Where the Worker's bucket keeps a path's bytes: the same layout under downloads/. */
export function storageKey(family: Family, ...rest: string[]): string {
  return ['downloads', family, ...rest].join('/');
}

export type DownloadRequest =
  | { kind: 'latest-zip'; family: Family }
  | { kind: 'manifest'; family: Family }
  | { kind: 'latest-file'; family: Family; name: string }
  | { kind: 'file'; family: Family; tag: string; name: string };

const FILE_NAME = /^[A-Za-z0-9][A-Za-z0-9._-]{0,127}$/;
const TAG_TAIL = /^[0-9][0-9A-Za-z.-]{0,40}$/;

function isFamily(value: string): value is Family {
  return Object.hasOwn(FAMILIES, value);
}

export function isFileName(name: string): boolean {
  return FILE_NAME.test(name) && !name.includes('..');
}

export function isTagOf(family: Family, tag: string): boolean {
  const prefix = FAMILIES[family];
  return tag.startsWith(prefix) && TAG_TAIL.test(tag.slice(prefix.length)) && !tag.includes('..');
}

/** What a /dl/ path asks for, or null when it is not a shape the feed serves. */
export function parseDownloadPath(pathname: string): DownloadRequest | null {
  if (!pathname.startsWith(DOWNLOAD_PREFIX)) return null;
  const parts = pathname.slice(DOWNLOAD_PREFIX.length).split('/');
  const [family, second, third] = parts;
  if (!isFamily(family) || parts.length > 3) return null;
  if (parts.length === 1) return family === 'companion' ? null : { kind: 'latest-zip', family };
  if (parts.length === 2) {
    if (second === MANIFEST_NAME) return { kind: 'manifest', family };
    return isFileName(second) ? { kind: 'latest-file', family, name: second } : null;
  }
  return isTagOf(family, second) && isFileName(third)
    ? { kind: 'file', family, tag: second, name: third }
    : null;
}
