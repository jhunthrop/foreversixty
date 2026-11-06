// web/src/lib/downloads/redirect.ts
// The Worker's answer to /dl/...: a 302 to wherever the file is hosted today. The host is the
// Worker's RELEASES_URL var (wrangler.jsonc), never a string in the site's pages, so the built
// site carries no link to it and a move is a deploy of the Worker alone. Only the names
// paths.ts parses are answered; anything else is a 404, never an open redirect.
import companion from '../../data/companion.json';
import { parseDownloadPath, type DownloadTarget } from './paths';

export interface DownloadEnv {
  /** The release listing the downloads live under, e.g. https://host/owner/repo/releases. */
  RELEASES_URL?: string;
}

const COMPANION_ASSETS = companion.downloads.map((entry) => entry.asset);

/** A minute: a new release reaches the button quickly, and a click never waits on a miss. */
const REDIRECT_CACHE = 'public, max-age=60';

/** Where a target lives under the release listing. */
export function releaseLocation(releasesUrl: string, target: DownloadTarget): string {
  switch (target.kind) {
    // Not releases/latest: that is the companion's release (companion.yml keeps --latest),
    // which carries no addon zip. The addon's releases are tagged addon-v<version>.
    case 'addon':
      return `${releasesUrl}?q=addon-v`;
    // The data pack is released nightly under its own data-v tag.
    case 'data':
      return `${releasesUrl}?q=data-v`;
    case 'companion':
      return `${releasesUrl}/latest/download/${target.asset}`;
  }
}

function refuse(status: number, message: string): Response {
  return new Response(message, { status, headers: { 'content-type': 'text/plain; charset=utf-8' } });
}

export function serveDownload(request: Request, env: DownloadEnv): Response {
  if (request.method !== 'GET' && request.method !== 'HEAD') return refuse(405, 'Method not allowed');
  const target = parseDownloadPath(new URL(request.url).pathname, COMPANION_ASSETS);
  if (target === null) return refuse(404, 'No such download');
  if (!env.RELEASES_URL) return refuse(503, 'Downloads are not configured');
  return new Response(null, {
    status: 302,
    headers: { location: releaseLocation(env.RELEASES_URL, target), 'cache-control': REDIRECT_CACHE },
  });
}
