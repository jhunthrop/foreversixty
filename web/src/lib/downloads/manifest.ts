// web/src/lib/downloads/manifest.ts
// A family's latest.json, as tools/publish-feed.sh writes it (docs/runbooks/repo-private.md).
// Its field names match a GitHub release object on purpose, so the companion decodes either.
// The Worker trusts only the names and the tag: every redirect it builds is a path on this
// site, never a URL read out of the file.
import { isFileName, isTagOf, type Family } from './paths';

export interface ManifestAsset {
  name: string;
  browser_download_url: string;
  size: number;
  sha256: string;
}

export interface Manifest {
  schema: 1;
  family: Family;
  tag_name: string;
  published_at: string;
  assets: ManifestAsset[];
}

function isAsset(value: unknown): value is ManifestAsset {
  if (typeof value !== 'object' || value === null) return false;
  const asset = value as Record<string, unknown>;
  return (
    typeof asset.name === 'string' && isFileName(asset.name) && typeof asset.browser_download_url === 'string'
  );
}

/** The manifest, or null when the text is not one for this family (a bad upload is a miss, not a crash). */
export function parseManifest(text: string, family: Family): Manifest | null {
  let value: unknown;
  try {
    value = JSON.parse(text);
  } catch {
    return null;
  }
  if (typeof value !== 'object' || value === null) return null;
  const manifest = value as Record<string, unknown>;
  if (manifest.schema !== 1 || manifest.family !== family) return null;
  if (typeof manifest.tag_name !== 'string' || !isTagOf(family, manifest.tag_name)) return null;
  if (!Array.isArray(manifest.assets) || !manifest.assets.every(isAsset)) return null;
  return value as Manifest;
}

/** The family's installable zip: the one .zip the release carries. */
export function latestZip(manifest: Manifest): ManifestAsset | undefined {
  return manifest.assets.find((asset) => asset.name.endsWith('.zip'));
}
