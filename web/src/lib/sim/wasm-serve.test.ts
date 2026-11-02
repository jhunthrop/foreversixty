import { brotliCompressSync, brotliDecompressSync } from 'node:zlib';
import { describe, expect, it, vi } from 'vitest';
import { acceptsBrotli, serveSimWasm, SIM_WASM_PATH } from './wasm-serve';

const URL_GZ = 'https://foreversixty.gg/_sim/4d3129cae-cc1844615a14/sim.wasm.gz';
const MODULE = new Uint8Array([0x00, 0x61, 0x73, 0x6d, 0x01, 0x00, 0x00, 0x00]);
const BROTLI = brotliCompressSync(MODULE);
const GZIP_BYTES = new Uint8Array([0x1f, 0x8b, 0x08, 0x00]);

type AssetFetch = (request: Request) => Promise<Response>;

/** The static assets as the two published files, each with its own etag. */
function assets(files: { br?: boolean } = { br: true }) {
  return {
    fetch: vi.fn<AssetFetch>(async (request) => {
      const path = new URL(request.url).pathname;
      if (path.endsWith('.br') && files.br) {
        return new Response(BROTLI, {
          headers: { etag: '"br-etag"', 'content-type': 'application/x-brotli' },
        });
      }
      if (path.endsWith('.gz')) {
        return new Response(GZIP_BYTES, {
          headers: { etag: '"gz-etag"', 'content-type': 'application/gzip' },
        });
      }
      return new Response('not found', { status: 404 });
    }),
  };
}

function get(acceptEncoding: string | null, init: RequestInit = {}): Request {
  const headers = new Headers(init.headers);
  if (acceptEncoding !== null) headers.set('accept-encoding', acceptEncoding);
  return new Request(URL_GZ, { ...init, headers });
}

describe('SIM_WASM_PATH', () => {
  it('matches the published module under an artifact directory and nothing else', () => {
    expect(SIM_WASM_PATH.test('/_sim/4d3129cae-cc1844615a14/sim.wasm.gz')).toBe(true);
    expect(SIM_WASM_PATH.test('/_sim/4d3129cae/sim.wasm.gz')).toBe(true);
    expect(SIM_WASM_PATH.test('/_sim/4d3129cae-cc1844615a14/sim.js')).toBe(false);
    expect(SIM_WASM_PATH.test('/_sim/4d3129cae-cc1844615a14/sim.wasm.br')).toBe(false);
    expect(SIM_WASM_PATH.test('/_sim/../secrets/sim.wasm.gz')).toBe(false);
  });
});

describe('acceptsBrotli', () => {
  it('reads br among the codings and honours q=0', () => {
    expect(acceptsBrotli('gzip, deflate, br, zstd')).toBe(true);
    expect(acceptsBrotli('br;q=0.5')).toBe(true);
    expect(acceptsBrotli('gzip, br;q=0')).toBe(false);
    expect(acceptsBrotli('gzip')).toBe(false);
    expect(acceptsBrotli('')).toBe(false);
  });
});

describe('serveSimWasm', () => {
  it('sends a brotli-capable browser the .br bytes as Content-Encoding: br application/wasm', async () => {
    const env = assets();
    const response = await serveSimWasm(get('gzip, deflate, br'), env);
    expect(response.status).toBe(200);
    expect(response.headers.get('content-encoding')).toBe('br');
    expect(response.headers.get('content-type')).toBe('application/wasm');
    expect(response.headers.get('cache-control')).toBe('public, max-age=31536000, immutable');
    expect(response.headers.get('vary')).toBe('Accept-Encoding');
    expect(response.headers.get('etag')).toBe('"br-etag"');
    // The stored bytes, untouched: the browser's own decoder turns them into the module.
    const body = new Uint8Array(await response.arrayBuffer());
    expect(new Uint8Array(brotliDecompressSync(body))).toEqual(MODULE);
    const asked = env.fetch.mock.calls[0][0];
    expect(new URL(asked.url).pathname).toBe('/_sim/4d3129cae-cc1844615a14/sim.wasm.br');
    expect(asked.headers.get('accept-encoding')).toBe('identity');
  });

  it("reads the visitor's own Accept-Encoding from request.cf over Cloudflare's rewrite", async () => {
    const request = Object.assign(get('gzip'), { cf: { clientAcceptEncoding: 'br, gzip' } });
    const response = await serveSimWasm(request, assets());
    expect(response.headers.get('content-encoding')).toBe('br');
  });

  it('sends a browser without br the gzip as the static assets serve it', async () => {
    const response = await serveSimWasm(get('gzip, deflate'), assets());
    expect(response.headers.get('content-encoding')).toBeNull();
    expect(response.headers.get('content-type')).toBe('application/gzip');
    expect(response.headers.get('etag')).toBe('"gz-etag"');
    expect(response.headers.get('vary')).toBe('Accept-Encoding');
    expect(response.headers.get('cache-control')).toBe('public, max-age=31536000, immutable');
    expect(new Uint8Array(await response.arrayBuffer())).toEqual(GZIP_BYTES);
  });

  it('falls back to the gzip when the artifact has no .br', async () => {
    const response = await serveSimWasm(get('br'), assets({ br: false }));
    expect(response.headers.get('content-encoding')).toBeNull();
    expect(new Uint8Array(await response.arrayBuffer())).toEqual(GZIP_BYTES);
  });

  it('falls back to the gzip when the asset layer encoded the .br itself', async () => {
    const env = {
      fetch: vi.fn<AssetFetch>(async (request) =>
        request.url.endsWith('.br')
          ? new Response(BROTLI, { headers: { 'content-encoding': 'gzip' } })
          : new Response(GZIP_BYTES, { headers: { 'content-type': 'application/gzip' } }),
      ),
    };
    const response = await serveSimWasm(get('br'), env);
    expect(response.headers.get('content-encoding')).toBeNull();
    expect(new Uint8Array(await response.arrayBuffer())).toEqual(GZIP_BYTES);
  });

  it('passes a revalidation through as a 304 with the brotli headers', async () => {
    const env = {
      fetch: vi.fn<AssetFetch>(
        async () => new Response(null, { status: 304, headers: { etag: '"br-etag"' } }),
      ),
    };
    const response = await serveSimWasm(get('br', { headers: { 'if-none-match': '"br-etag"' } }), env);
    expect(response.status).toBe(304);
    expect(response.headers.get('content-encoding')).toBe('br');
    expect(env.fetch.mock.calls[0][0].headers.get('if-none-match')).toBe('"br-etag"');
  });

  it('answers HEAD with the headers and no body', async () => {
    const response = await serveSimWasm(get('br', { method: 'HEAD' }), assets());
    expect(response.headers.get('content-encoding')).toBe('br');
    expect(response.body).toBeNull();
  });
});
