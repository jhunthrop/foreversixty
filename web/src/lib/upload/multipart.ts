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
export const UPLOAD_CANCELLED = 'The upload was cancelled';

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
 *
 * The signal aborts the request in flight rather than at the next part boundary. A part is
 * 64 MiB, which on a home uplink is minutes: checking the signal between parts only, as
 * this did before, means Cancel appears to do nothing for as long as the part takes and
 * the bytes keep going up in the meantime.
 */
export function putPart(
  url: string,
  body: Blob,
  onBytes: (loaded: number) => void,
  make: XhrFactory = () => new XMLHttpRequest(),
  signal?: AbortSignal,
): Promise<string> {
  return new Promise((resolve, reject) => {
    if (signal?.aborted === true) {
      reject(new AccountError(UPLOAD_CANCELLED, 0));
      return;
    }
    const xhr = make();
    const stopListening = (): void => signal?.removeEventListener('abort', onAbort);
    function onAbort(): void {
      xhr.abort();
    }
    signal?.addEventListener('abort', onAbort, { once: true });
    xhr.open('PUT', url);
    xhr.upload.onprogress = (event): void => onBytes(event.loaded);
    xhr.onload = (): void => {
      stopListening();
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
    xhr.onabort = (): void => {
      stopListening();
      reject(new AccountError(UPLOAD_CANCELLED, 0));
    };
    xhr.onerror = (): void => {
      stopListening();
      reject(new AccountError(UPLOAD_FAILED, 0));
    };
    xhr.send(body);
  });
}

export interface UploadOptions {
  makeXhr?: XhrFactory;
  retryDelayMs?: number;
  signal?: AbortSignal;
  /**
   * Starts a fresh upload when R2 refuses an expired signature. Injected rather than
   * called directly so the tests can drive it, and because only the caller still holds the
   * File this module was handed as a bare Blob.
   */
  restart?: () => Promise<CreatedUpload>;
}

/** What uploadParts finished against, which is not always the upload it was handed. */
export interface UploadedParts {
  uploadId: string;
  etags: PartEtag[];
}

const ATTEMPTS_PER_PART = 3;

/**
 * An expired signature, which is the one failure retrying the same url cannot fix: the
 * part urls last an hour (r2.URLTTL) and a 4 GiB night on a home uplink takes longer.
 * R2 answers 403 for a signature that no longer verifies, and 401 is included for the same
 * class of answer from anything in front of it.
 */
function isExpiredSignature(error: unknown): boolean {
  return error instanceof AccountError && (error.status === 401 || error.status === 403);
}

function aborted(signal: AbortSignal | undefined): boolean {
  return signal?.aborted === true;
}

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
 * A part that fails is retried, not the upload: the ones already done keep their ETags.
 * The one failure that is not retried here is an expired signature, which retrying the
 * same url cannot fix; uploadParts handles that.
 */
async function putEveryPart(
  file: Blob,
  created: CreatedUpload,
  onProgress: (progress: UploadProgress) => void,
  options: UploadOptions,
): Promise<PartEtag[]> {
  const { makeXhr, retryDelayMs = 1000, signal } = options;
  const etags: PartEtag[] = [];
  let completedBytes = 0;

  for (const [index, part] of created.parts.entries()) {
    if (aborted(signal)) throw new AccountError(UPLOAD_CANCELLED, 0);

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
          signal,
        );
        break;
      } catch (error) {
        lastError = error;
        // Neither an expired signature nor a cancellation is worth two more attempts at
        // the same url: the first needs a fresh upload, the second needs nothing at all.
        if (isExpiredSignature(error) || aborted(signal)) throw error;
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

/**
 * The whole file, and the recovery from the one failure that cannot be retried in place.
 *
 * The signed part urls last an hour and MAX_UPLOAD_BYTES is 4 GiB, so an upload on a home
 * uplink can outlive its own signatures; before this, every part already sent was thrown
 * away and the visitor was told part N failed three times.
 *
 * The recovery is a restart rather than a resume, because `POST /v1/uploads` is not a
 * re-signing endpoint: api/internal/reports/uploads.go's `start` mints a new id, a new
 * object key and a new R2 multipart upload every time, so the parts already sent belong to
 * an object the new upload cannot complete with. Once, and only once: a second expiry
 * means the file cannot be uploaded in an hour on this connection, and restarting forever
 * would be a loop rather than a recovery. The upload that was actually finished comes back
 * with the ETags, because it is the one `completeUpload` has to be told about.
 */
export async function uploadParts(
  file: Blob,
  created: CreatedUpload,
  onProgress: (progress: UploadProgress) => void,
  options: UploadOptions = {},
): Promise<UploadedParts> {
  try {
    return { uploadId: created.upload_id, etags: await putEveryPart(file, created, onProgress, options) };
  } catch (error) {
    if (options.restart === undefined || aborted(options.signal) || !isExpiredSignature(error)) throw error;
    const fresh = await options.restart();
    return {
      uploadId: fresh.upload_id,
      etags: await putEveryPart(file, fresh, onProgress, { ...options, restart: undefined }),
    };
  }
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
