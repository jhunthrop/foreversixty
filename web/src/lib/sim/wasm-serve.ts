// web/src/lib/sim/wasm-serve.ts
// The Worker's answer for the engine module (wrangler.jsonc routes `/_sim/*/sim.wasm.gz`
// here). `make publish-wasm` writes two compressions of one module beside each other:
// sim.wasm.gz and sim.wasm.br, about 28% smaller. A browser that accepts brotli is sent the
// .br bytes under the .gz's URL with Content-Encoding: br, so it decodes them natively and
// the page (wasm-fetch.ts) receives the module itself, which it already accepts. Every
// other request, and any miss on the .br, gets the .gz exactly as the static assets serve it.
//
// Why not static assets alone: Workers static assets publish no precompressed variants and
// document no way to label a file's own encoding, and DecompressionStream('brotli') is not
// in Chrome, so the page cannot inflate a plain .br itself. A Worker response can carry
// pre-encoded bytes: `encodeBody: 'manual'` tells the runtime the body is already in the
// Content-Encoding it names (developers.cloudflare.com/workers/runtime-apis/response).
//
// `astro preview` and every Playwright run serve dist/ with no Worker, so they keep taking
// the .gz path, the one the real-engine CI job proves.

/** The module's published path: the artifact directory, then the gzip's file name. */
export const SIM_WASM_PATH = /^\/_sim\/[0-9a-f]{7,40}(?:-[0-9a-f]{12})?\/sim\.wasm\.gz$/;

/** The directory is named by the bytes (version.ts), so a response never goes stale. */
const IMMUTABLE = 'public, max-age=31536000, immutable';

/** The same URL answers with two encodings, so every cache must key on the request's. */
const VARY = 'Accept-Encoding';

export interface AssetFetcher {
  fetch(request: Request): Promise<Response>;
}

/** The request's Request.cf, as far as this module reads it. */
interface CfProperties {
  clientAcceptEncoding?: string;
}

/**
 * The browser's own Accept-Encoding. Cloudflare rewrites the header to a canonical value
 * before a Worker runs and keeps the visitor's original on request.cf.
 */
function clientAcceptEncoding(request: Request): string {
  const cf = (request as Request & { cf?: CfProperties }).cf;
  return cf?.clientAcceptEncoding ?? request.headers.get('accept-encoding') ?? '';
}

/** True when the header lists br with a non-zero quality. */
export function acceptsBrotli(acceptEncoding: string): boolean {
  return acceptEncoding.split(',').some((entry) => {
    const [coding, ...params] = entry.trim().toLowerCase().split(';');
    if (coding.trim() !== 'br') return false;
    const q = params.map((p) => p.trim()).find((p) => p.startsWith('q='));
    return q === undefined || Number(q.slice(2)) > 0;
  });
}

/** The .gz from the static assets, with the cache headers this route owns. */
async function serveGzip(request: Request, assets: AssetFetcher): Promise<Response> {
  const asset = await assets.fetch(request);
  if (!asset.ok && asset.status !== 304) return asset;
  const headers = new Headers(asset.headers);
  headers.set('cache-control', IMMUTABLE);
  headers.set('vary', VARY);
  return new Response(asset.body, { status: asset.status, headers });
}

/**
 * The .br asset, fetched with no encoding of its own so its body is the stored bytes; a
 * conditional request keeps its If-None-Match, so a revalidation is still a 304.
 */
function brotliAssetRequest(request: Request): Request {
  const headers = new Headers();
  const etag = request.headers.get('if-none-match');
  if (etag !== null) headers.set('if-none-match', etag);
  headers.set('accept-encoding', 'identity');
  return new Request(request.url.replace(/\.gz$/, '.br'), { method: request.method, headers });
}

export async function serveSimWasm(request: Request, assets: AssetFetcher): Promise<Response> {
  if (!acceptsBrotli(clientAcceptEncoding(request))) return serveGzip(request, assets);
  const asset = await assets.fetch(brotliAssetRequest(request));
  // A missing .br (an artifact published before this route) or a body the asset layer
  // re-encoded itself cannot be sent as stored brotli: the .gz is always right.
  if ((!asset.ok && asset.status !== 304) || asset.headers.has('content-encoding')) {
    await asset.body?.cancel();
    return serveGzip(request, assets);
  }
  const headers = new Headers({
    'content-type': 'application/wasm',
    'content-encoding': 'br',
    'cache-control': IMMUTABLE,
    vary: VARY,
  });
  const etag = asset.headers.get('etag');
  if (etag !== null) headers.set('etag', etag);
  const length = asset.headers.get('content-length');
  if (length !== null) headers.set('content-length', length);
  const body = request.method === 'HEAD' || asset.status === 304 ? null : asset.body;
  // encodeBody is the Workers runtime's own ResponseInit field; Node and the DOM types
  // do not know it, and ignore it.
  const init = { status: asset.status, headers, encodeBody: 'manual' } as ResponseInit;
  return new Response(body, init);
}
