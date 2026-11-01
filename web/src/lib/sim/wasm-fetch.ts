// web/src/lib/sim/wasm-fetch.ts
// sim.wasm is published gzipped (`make publish-wasm` writes sim.wasm.gz): the raw module is
// within a few percent of Cloudflare's 25 MiB per-file static-asset limit, and a deploy that
// carries one file over it is refused whole. The gzip is about a sixth of the size and is the
// same bytes the visitor downloaded before (the edge compressed the raw file on the fly), so
// the page inflates it itself with DecompressionStream and hands the module to
// instantiateStreaming, which still compiles while it downloads.
//
// The body is sniffed rather than trusted: a host that labels the file Content-Encoding: gzip
// (astro preview does, so every local and e2e build takes that path) makes fetch inflate it already, and inflating a module a second time would throw.

const GZIP_MAGIC = [0x1f, 0x8b] as const;
const WASM_MAGIC = [0x00, 0x61, 0x73, 0x6d] as const;

function startsWith(bytes: Uint8Array, magic: readonly number[]): boolean {
  return magic.every((byte, index) => bytes[index] === byte);
}

/** Reads until at least `count` bytes have arrived (or the body ends) and returns them all. */
async function readHead(reader: ReadableStreamDefaultReader<Uint8Array>, count: number): Promise<Uint8Array> {
  const chunks: Uint8Array[] = [];
  let length = 0;
  while (length < count) {
    const { done, value } = await reader.read();
    if (done) break;
    chunks.push(value);
    length += value.length;
  }
  const head = new Uint8Array(length);
  let offset = 0;
  for (const chunk of chunks) {
    head.set(chunk, offset);
    offset += chunk.length;
  }
  return head;
}

/**
 * The body again, whole: the bytes already read, then the rest of the reader. Typed as
 * BufferSource chunks because that is what DecompressionStream's writable side takes.
 */
function rejoin(
  head: Uint8Array,
  reader: ReadableStreamDefaultReader<Uint8Array>,
): ReadableStream<BufferSource> {
  return new ReadableStream<BufferSource>({
    start(controller) {
      controller.enqueue(head);
    },
    async pull(controller) {
      const { done, value } = await reader.read();
      if (done) controller.close();
      else controller.enqueue(value);
    },
    cancel(reason) {
      return reader.cancel(reason);
    },
  });
}

/**
 * Turns the fetched engine file into a response `WebAssembly.instantiateStreaming` accepts:
 * the module itself, typed application/wasm, whether the file arrived gzipped or already
 * inflated. Throws with the URL on a failed fetch or a body that is neither.
 */
export async function wasmResponse(response: Response): Promise<Response> {
  const source = response.url || 'the sim engine';
  if (!response.ok) throw new Error(`${source} answered ${response.status}`);
  if (response.body === null) throw new Error(`${source} returned an empty body`);
  const reader = response.body.getReader();
  const head = await readHead(reader, WASM_MAGIC.length);
  if (head.length === 0) throw new Error(`${source} returned an empty body`);
  const whole = rejoin(head, reader);
  const headers = { 'content-type': 'application/wasm' };
  if (startsWith(head, GZIP_MAGIC)) {
    return new Response(whole.pipeThrough(new DecompressionStream('gzip')), { headers });
  }
  if (startsWith(head, WASM_MAGIC)) return new Response(whole, { headers });
  await reader.cancel();
  throw new Error(`${source} is neither a gzip file nor a WebAssembly module`);
}
