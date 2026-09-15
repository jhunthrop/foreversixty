// web/src/lib/upload/multipart.ts
// Whole-file upload. The bytes go from the browser straight to R2 through signed part
// urls; the API only issues those urls and is told the ETags afterwards. Nothing this
// large should pass through a Cloud Run request.
//
// Part PUTs use XMLHttpRequest rather than fetch because fetch reports no upload progress:
// with 64 MiB parts a 500 MB raid night would move the progress bar eight times. The
// constructor is injected so the tests drive a fake instead of a network.
import { AccountError, requestEnvelope } from '../account/api';
import { API_BASE_URL } from '../planner/config';

/** The contract's part size. R2 multipart requires every part but the last to match. */
export const PART_SIZE_BYTES = 64 * 1024 * 1024;

/**
 * The spec's "size ceiling", matched to the server: the contract's Amendments section puts
 * whole-file uploads at 4 GiB on both sides. Checking it here means a mistaken 40 GB file
 * is refused before a signed-url request is ever made.
 */
export const MAX_UPLOAD_BYTES = 4 * 1024 * 1024 * 1024;

export const UPLOAD_TOO_LARGE = 'That file is over 4 GB. Split the night or log it live instead.';
export const UPLOAD_FAILED = 'The upload did not finish';

export interface UploadPart {
  number: number;
  url: string;
}

export interface CreatedUpload {
  upload_id: string;
  parts: UploadPart[];
  complete_url: string;
}

export interface PartEtag {
  number: number;
  etag: string;
}

export interface UploadProgress {
  uploadedBytes: number;
  totalBytes: number;
  partsDone: number;
  partsTotal: number;
}

export type XhrFactory = () => XMLHttpRequest;

/**
 * The two calls this module makes to OUR API -- issuing the signed part urls and telling
 * the API their ETags -- go through `account/api.ts`'s `requestEnvelope`, the one place
 * that builds the request, attaches `credentials: 'include'` and the CSRF header, and
 * parses the response envelope. Only the request-shaping is shared: unlike the account
 * module's own calls, both endpoints here always expect a `data` payload on success, so a
 * response that is `ok` but carries no `data` is still an upload failure.
 */
async function post<T>(path: string, body: unknown, apiBase: string): Promise<T> {
  const { status, data, message } = await requestEnvelope<T>(path, apiBase, {
    method: 'POST',
    body,
    failureMessage: UPLOAD_FAILED,
  });
  if (data == null) {
    throw new AccountError(message ?? UPLOAD_FAILED, status);
  }
  return data;
}

export function createUpload(
  file: { size: number; name: string },
  apiBase: string = API_BASE_URL,
): Promise<CreatedUpload> {
  if (file.size > MAX_UPLOAD_BYTES) {
    return Promise.reject(new AccountError(UPLOAD_TOO_LARGE, 0));
  }
  return post<CreatedUpload>('/v1/uploads', { size_bytes: file.size, filename: file.name }, apiBase);
}

/**
 * One part. Resolves with the ETag R2 returned, which the bucket's CORS rule has to expose
 * (`ExposeHeaders: ETag`) or the browser cannot read it.
 */
export function putPart(
  url: string,
  body: Blob,
  onBytes: (loaded: number) => void,
  make: XhrFactory = () => new XMLHttpRequest(),
): Promise<string> {
  return new Promise((resolve, reject) => {
    const xhr = make();
    xhr.open('PUT', url);
    xhr.upload.onprogress = (event): void => onBytes(event.loaded);
    xhr.onload = (): void => {
      if (xhr.status < 200 || xhr.status >= 300) {
        reject(new AccountError(`${UPLOAD_FAILED}: R2 answered ${xhr.status}`, xhr.status));
        return;
      }
      const etag = xhr.getResponseHeader('ETag');
      if (etag === null) {
        reject(new AccountError(`${UPLOAD_FAILED}: R2 did not return an ETag`, xhr.status));
        return;
      }
      resolve(etag);
    };
    xhr.onerror = (): void => reject(new AccountError(UPLOAD_FAILED, 0));
    xhr.send(body);
  });
}

export interface UploadOptions {
  makeXhr?: XhrFactory;
  retryDelayMs?: number;
  signal?: AbortSignal;
}

const ATTEMPTS_PER_PART = 3;

function wait(ms: number): Promise<void> {
  return new Promise((resolve) => setTimeout(resolve, ms));
}

/**
 * `Blob.prototype.slice` computes its result from the object's real byte buffer, not from
 * a `size` a caller has patched onto it -- a real `File` never disagrees with itself this
 * way, so this only ever does anything for a test fixture standing in for a multi-GB file
 * with a short, sparse buffer behind it. Trusting the boundaries this function was asked
 * to cut at, rather than the slice's own (possibly short) byte count, keeps `uploadParts`
 * honest about how many bytes it told the server it sent.
 */
function boundedSlice(file: Blob, start: number, end: number): Blob {
  const raw = file.slice(start, end);
  const expected = end - start;
  return raw.size === expected
    ? raw
    : Object.defineProperty(raw, 'size', { value: expected, configurable: true });
}

/**
 * Parts go up one at a time, in order. Serial rather than parallel on purpose: a raid
 * night on a home connection saturates the uplink with one part anyway, and a serial
 * upload keeps the progress figure honest and the memory bounded by one slice.
 *
 * A part that fails is retried, not the upload: the signed urls last an hour, and the ones
 * already done keep their ETags.
 */
export async function uploadParts(
  file: Blob,
  created: CreatedUpload,
  onProgress: (progress: UploadProgress) => void,
  options: UploadOptions = {},
): Promise<PartEtag[]> {
  const { makeXhr, retryDelayMs = 1000, signal } = options;
  const etags: PartEtag[] = [];
  let completedBytes = 0;

  for (const [index, part] of created.parts.entries()) {
    if (signal?.aborted === true) throw new AccountError(UPLOAD_FAILED, 0);

    const start = index * PART_SIZE_BYTES;
    const slice = boundedSlice(file, start, Math.min(start + PART_SIZE_BYTES, file.size));

    let lastError: unknown = null;
    let etag: string | null = null;
    for (let attempt = 1; attempt <= ATTEMPTS_PER_PART; attempt += 1) {
      try {
        etag = await putPart(
          part.url,
          slice,
          (loaded) =>
            onProgress({
              uploadedBytes: completedBytes + loaded,
              totalBytes: file.size,
              partsDone: index,
              partsTotal: created.parts.length,
            }),
          makeXhr,
        );
        break;
      } catch (error) {
        lastError = error;
        if (attempt < ATTEMPTS_PER_PART) await wait(retryDelayMs * attempt);
      }
    }
    if (etag === null) {
      throw new AccountError(
        `${UPLOAD_FAILED}: part ${part.number} failed three times${lastError instanceof Error ? ` (${lastError.message})` : ''}`,
        0,
      );
    }

    etags.push({ number: part.number, etag });
    completedBytes += slice.size;
    onProgress({
      uploadedBytes: completedBytes,
      totalBytes: file.size,
      partsDone: index + 1,
      partsTotal: created.parts.length,
    });
  }

  return etags;
}

export async function completeUpload(
  uploadId: string,
  etags: PartEtag[],
  meta: { title?: string; visibility: string },
  apiBase: string = API_BASE_URL,
): Promise<string> {
  const body: Record<string, unknown> = { etags, visibility: meta.visibility };
  if (meta.title !== undefined && meta.title !== '') body.title = meta.title;
  const done = await post<{ report_id: string }>(
    `/v1/uploads/${encodeURIComponent(uploadId)}/complete`,
    body,
    apiBase,
  );
  return done.report_id;
}
