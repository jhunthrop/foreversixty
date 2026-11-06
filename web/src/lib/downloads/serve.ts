// web/src/lib/downloads/serve.ts
// The Worker's answer to /dl/... (paths.ts has the layout). The release workflows put every
// signed file and the family's latest.json into the bucket under downloads/; this serves
// them, and turns "the latest addon zip" or "the latest companion build" into a redirect to
// that file's own immutable path, read from the manifest. Only names and tags that paths.ts
// accepts are ever looked up, and every redirect is a path on this site.
//
// Until a family's first release reaches the bucket, its "latest" paths fall back to a 302
// under RELEASES_URL (wrangler.jsonc), the release host the site used before. The var is
// deleted once every family has published here (docs/runbooks/repo-private.md).
import companion from '../../data/companion.json';
import { latestZip, parseManifest, type Manifest } from './manifest';
import {
  MANIFEST_NAME,
  parseDownloadPath,
  releaseFilePath,
  storageKey,
  type DownloadRequest,
  type Family,
} from './paths';

export interface DownloadObject {
  body: ReadableStream | null;
  httpEtag: string;
  size: number;
  text?(): Promise<string>;
}

export interface DownloadBucket {
  get(key: string): Promise<DownloadObject | null>;
}

export interface DownloadEnv {
  /** The bucket the release workflows publish into (the LOGS binding, under downloads/). */
  LOGS?: DownloadBucket;
  /** The pre-feed release host, e.g. https://host/owner/repo/releases; unset once the feed is live. */
  RELEASES_URL?: string;
}

/** A minute: a new release reaches the button and the updater quickly. */
const LATEST_CACHE = 'public, max-age=60';
/** A published file never changes under its tag. */
const IMMUTABLE_CACHE = 'public, max-age=31536000, immutable';

const CONTENT_TYPES: Record<string, string> = {
  '.exe': 'application/vnd.microsoft.portable-executable',
  '.zip': 'application/zip',
  '.json': 'application/json; charset=utf-8',
  '.minisig': 'text/plain; charset=utf-8',
  '.txt': 'text/plain; charset=utf-8',
};

const COMPANION_ASSETS: readonly string[] = companion.downloads.map((entry) => entry.asset);

function contentType(name: string): string {
  const dot = name.lastIndexOf('.');
  return (dot >= 0 && CONTENT_TYPES[name.slice(dot)]) || 'application/octet-stream';
}

function refuse(status: number, message: string): Response {
  return new Response(message, {
    status,
    headers: { 'content-type': 'text/plain; charset=utf-8', 'cache-control': 'no-store' },
  });
}

function redirect(location: string): Response {
  return new Response(null, { status: 302, headers: { location, 'cache-control': LATEST_CACHE } });
}

async function readManifest(env: DownloadEnv, family: Family): Promise<Manifest | null> {
  const object = await env.LOGS?.get(storageKey(family, MANIFEST_NAME));
  if (!object?.text) return null;
  return parseManifest(await object.text(), family);
}

/** Where the pre-feed host kept the same thing, or null when it has no equivalent. */
export function legacyLocation(releasesUrl: string, request: DownloadRequest): string | null {
  if (request.kind === 'latest-zip') return `${releasesUrl}?q=${request.family}-v`;
  if (
    request.kind === 'latest-file' &&
    request.family === 'companion' &&
    COMPANION_ASSETS.includes(request.name)
  ) {
    return `${releasesUrl}/latest/download/${request.name}`;
  }
  return null;
}

async function serveFile(
  request: Request,
  env: DownloadEnv,
  key: string,
  name: string,
  cache: string,
): Promise<Response> {
  const object = await env.LOGS?.get(key);
  if (object === null || object === undefined) return refuse(404, 'No such download');
  const headers = new Headers({
    'content-type': contentType(name),
    'cache-control': cache,
    'content-length': String(object.size),
    etag: object.httpEtag,
  });
  if (!name.endsWith('.json')) headers.set('content-disposition', `attachment; filename="${name}"`);
  return new Response(request.method === 'HEAD' ? null : object.body, { status: 200, headers });
}

async function serveLatest(
  env: DownloadEnv,
  request: DownloadRequest & { kind: 'latest-zip' | 'latest-file' },
): Promise<Response> {
  const manifest = await readManifest(env, request.family);
  if (manifest !== null) {
    const asset =
      request.kind === 'latest-zip'
        ? latestZip(manifest)
        : manifest.assets.find((candidate) => candidate.name === request.name);
    if (asset === undefined) return refuse(404, 'The latest release has no such file');
    return redirect(releaseFilePath(request.family, manifest.tag_name, asset.name));
  }
  const legacy = env.RELEASES_URL ? legacyLocation(env.RELEASES_URL, request) : null;
  return legacy === null ? refuse(404, 'No release has been published yet') : redirect(legacy);
}

export async function serveDownload(request: Request, env: DownloadEnv): Promise<Response> {
  if (request.method !== 'GET' && request.method !== 'HEAD') return refuse(405, 'Method not allowed');
  const target = parseDownloadPath(new URL(request.url).pathname);
  if (target === null) return refuse(404, 'No such download');
  switch (target.kind) {
    case 'manifest':
      return serveFile(request, env, storageKey(target.family, MANIFEST_NAME), MANIFEST_NAME, LATEST_CACHE);
    case 'file':
      return serveFile(
        request,
        env,
        storageKey(target.family, target.tag, target.name),
        target.name,
        IMMUTABLE_CACHE,
      );
    case 'latest-zip':
    case 'latest-file':
      return serveLatest(env, target);
  }
}
