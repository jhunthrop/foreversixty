import { gzipSync } from 'node:zlib';
import { describe, expect, it } from 'vitest';
import { wasmResponse } from './wasm-fetch';

// The first eight bytes of every WebAssembly module: "\0asm" and version 1.
const WASM_MODULE = new Uint8Array([0x00, 0x61, 0x73, 0x6d, 0x01, 0x00, 0x00, 0x00, 0x2a, 0x2b]);

/** A response whose body arrives in `chunkSize`-byte pieces, the way a network read does. */
function chunkedResponse(bytes: Uint8Array, chunkSize: number, init?: ResponseInit): Response {
  const body = new ReadableStream<Uint8Array>({
    start(controller) {
      for (let offset = 0; offset < bytes.length; offset += chunkSize) {
        controller.enqueue(bytes.slice(offset, offset + chunkSize));
      }
      controller.close();
    },
  });
  return new Response(body, init);
}

async function bytesOf(response: Response): Promise<Uint8Array> {
  return new Uint8Array(await response.arrayBuffer());
}

describe('wasmResponse', () => {
  it('inflates the gzip file the deploy publishes into the module, typed for instantiateStreaming', async () => {
    const response = await wasmResponse(new Response(gzipSync(WASM_MODULE, { level: 9 })));
    expect(response.headers.get('content-type')).toBe('application/wasm');
    expect(await bytesOf(response)).toEqual(WASM_MODULE);
  });

  it('inflates a body whose first network chunk is a single byte', async () => {
    const response = await wasmResponse(chunkedResponse(gzipSync(WASM_MODULE), 1));
    expect(await bytesOf(response)).toEqual(WASM_MODULE);
  });

  it('passes a module through untouched when something upstream already decoded the gzip', async () => {
    // A host that labels the file Content-Encoding: gzip makes fetch inflate it itself; a
    // second inflate would throw on bytes that are already the module.
    const response = await wasmResponse(chunkedResponse(WASM_MODULE, 3));
    expect(response.headers.get('content-type')).toBe('application/wasm');
    expect(await bytesOf(response)).toEqual(WASM_MODULE);
  });

  it('names the URL and status of a failed fetch instead of handing a 404 page to the compiler', async () => {
    const failed = new Response('not found', { status: 404 });
    Object.defineProperty(failed, 'url', { value: 'https://example.test/_sim/x/sim.wasm.gz' });
    await expect(wasmResponse(failed)).rejects.toThrow(
      'https://example.test/_sim/x/sim.wasm.gz answered 404',
    );
  });

  it('refuses a body that is neither gzip nor a module', async () => {
    await expect(wasmResponse(new Response('<!doctype html>'))).rejects.toThrow(
      'is neither a gzip file nor a WebAssembly module',
    );
  });

  it('refuses an empty body', async () => {
    await expect(wasmResponse(new Response(new Uint8Array()))).rejects.toThrow('empty');
  });
});
