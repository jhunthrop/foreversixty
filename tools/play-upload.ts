/**
 * tools/play-upload.ts: the signed release bundle to a Play track through the Play Developer
 * API's edits flow, with a short-lived access token. tools/release-apps.sh runs it; CI mints the
 * token by GitHub federation as the Play service account (vars.PLAY_SA), a local release with
 * `gcloud auth print-access-token --impersonate-service-account`. No JSON key exists or may
 * exist. No dependency beyond fetch:
 *   1. insert an edit,
 *   2. upload the bundle (a resumable upload: one request after the session),
 *   3. upload the R8 mapping for that version code, when one is given,
 *   4. set the track to that version code with the release status and notes,
 *   5. commit the edit.
 * Before the track is set it is read: when its highest version code is equal to or above this
 * build's, the edit is deleted and the run says so (exit 0), so a release never rolls the track
 * back. An app that has not passed its first review takes only draft releases: when Play refuses
 * a completed release for that reason, the track is set again as draft in the same edit and
 * committed, and the run says so (exit 0). Any other refusal deletes the edit and exits 1. Every
 * request has a timeout and none is retried.
 *
 *   PLAY_ACCESS_TOKEN=... node tools/play-upload.ts --package gg.foreversixty.app \
 *     --bundle app-release.aab [--mapping mapping.txt] [--notes notes.txt] \
 *     [--track internal] [--status draft|completed]
 *
 * The token is read from the environment only and never printed; the API's answers are.
 */
import { readFile } from 'node:fs/promises';
import { parseArgs } from 'node:util';

export const API_ROOT = 'https://androidpublisher.googleapis.com';
const API_PATH = '/androidpublisher/v3/applications';
const UPLOAD_PATH = '/upload/androidpublisher/v3/applications';
const BINARY_TYPE = 'application/octet-stream';
const NOTES_LANGUAGE = 'en-US';
/** Play's words when an app with no reviewed release is sent a release that is not a draft. */
const DRAFT_APP_REFUSAL = /draft app/i;
/** A JSON call's bound, and a file's (a large body on a slow runner). */
const CALL_TIMEOUT_MS = 60_000;
const FILE_TIMEOUT_MS = 5 * 60_000;

export type ReleaseStatus = 'completed' | 'draft';
export type Outcome = 'committed' | 'committed-as-draft' | 'track-already-ahead';

export interface UploadRequest {
  token: string;
  packageName: string;
  track: string;
  status: ReleaseStatus;
  bundle: Uint8Array<ArrayBuffer>;
  /** The R8 mapping, uploaded as the version's deobfuscation file; null sends none. */
  mapping: Uint8Array<ArrayBuffer> | null;
  /** The release notes Play shows testers; empty sends none. */
  notes: string;
}

export interface UploadResult {
  outcome: Outcome;
  editId: string;
  versionCode: number;
  status: ReleaseStatus | null;
  /** Only on a skip: the highest code the track already holds. */
  trackVersionCode?: number;
}

export type Fetch = (url: string, init: RequestInit) => Promise<Response>;
export type Log = (line: string) => void;

/** A refusal from the API, with its status and its own message (never the request's headers). */
export class PlayError extends Error {
  readonly status: number;
  readonly step: string;
  constructor(step: string, status: number, message: string) {
    super(`${step}: HTTP ${status}: ${message}`);
    this.step = step;
    this.status = status;
  }
}

async function errorMessage(response: Response): Promise<string> {
  const text = await response.text();
  try {
    const parsed = JSON.parse(text) as { error?: { message?: string } };
    return parsed.error?.message ?? text;
  } catch {
    return text;
  }
}

class Client {
  // Plain fields, not parameter properties: Node runs this file by stripping types only.
  private readonly fetch: Fetch;
  private readonly token: string;
  readonly base: string;
  readonly uploadBase: string;

  constructor(fetch: Fetch, token: string, packageName: string) {
    this.fetch = fetch;
    this.token = token;
    const name = encodeURIComponent(packageName);
    this.base = `${API_ROOT}${API_PATH}/${name}`;
    this.uploadBase = `${API_ROOT}${UPLOAD_PATH}/${name}`;
  }

  async call(step: string, url: string, init: RequestInit = {}, timeoutMs = CALL_TIMEOUT_MS): Promise<Response> {
    const headers = { ...(init.headers as Record<string, string>), Authorization: `Bearer ${this.token}` };
    const response = await this.fetch(url, { ...init, headers, signal: AbortSignal.timeout(timeoutMs) });
    if (!response.ok) throw new PlayError(step, response.status, await errorMessage(response));
    return response;
  }

  async json<T>(step: string, url: string, init: RequestInit = {}, timeoutMs = CALL_TIMEOUT_MS): Promise<T> {
    return (await (await this.call(step, url, init, timeoutMs)).json()) as T;
  }
}

const JSON_HEADERS = { 'Content-Type': 'application/json' };

async function insertEdit(client: Client): Promise<string> {
  const edit = await client.json<{ id: string }>('insert edit', `${client.base}/edits`, {
    method: 'POST',
    headers: JSON_HEADERS,
    body: '{}',
  });
  return edit.id;
}

async function uploadBundle(client: Client, editId: string, bundle: Uint8Array<ArrayBuffer>): Promise<number> {
  const session = await client.call(
    'start bundle upload',
    `${client.uploadBase}/edits/${editId}/bundles?uploadType=resumable`,
    {
      method: 'POST',
      headers: {
        ...JSON_HEADERS,
        'X-Upload-Content-Type': BINARY_TYPE,
        'X-Upload-Content-Length': String(bundle.byteLength),
      },
      body: '{}',
    },
  );
  const location = session.headers.get('location');
  if (!location) throw new PlayError('start bundle upload', session.status, 'no upload location in the answer');
  const uploaded = await client.json<{ versionCode: number }>(
    'upload bundle',
    location,
    { method: 'PUT', headers: { 'Content-Type': BINARY_TYPE }, body: bundle },
    FILE_TIMEOUT_MS,
  );
  return uploaded.versionCode;
}

async function uploadMapping(client: Client, editId: string, versionCode: number, mapping: Uint8Array<ArrayBuffer>): Promise<void> {
  await client.call(
    'upload mapping',
    `${client.uploadBase}/edits/${editId}/apks/${versionCode}/deobfuscationFiles/proguard?uploadType=media`,
    { method: 'POST', headers: { 'Content-Type': BINARY_TYPE }, body: mapping },
    FILE_TIMEOUT_MS,
  );
}

interface TrackAnswer {
  releases?: { versionCodes?: string[] }[];
}

/** A track with no release yet answers 404: an empty track. Any other refusal stays an error. */
async function readTrack(client: Client, editId: string, track: string): Promise<TrackAnswer> {
  try {
    return await client.json<TrackAnswer>('read track', `${client.base}/edits/${editId}/tracks/${encodeURIComponent(track)}`);
  } catch (error) {
    if (error instanceof PlayError && error.status === 404) return {};
    throw error;
  }
}

/** Every version code on the track's releases; a release without codes contributes none. */
async function trackVersionCodes(client: Client, editId: string, track: string): Promise<number[]> {
  const answer = await readTrack(client, editId, track);
  return (answer.releases ?? [])
    .flatMap((release) => release.versionCodes ?? [])
    .map(Number)
    .filter(Number.isInteger);
}

export type TrackDecision = 'upload' | 'skip';

/** Upload only a build newer than everything on the track. */
export function decideUpload(trackCodes: readonly number[], buildCode: number): TrackDecision {
  return trackCodes.length > 0 && Math.max(...trackCodes) >= buildCode ? 'skip' : 'upload';
}

/** The track body Play takes: one release of this version code, with its notes when there are any. */
export function trackBody(track: string, versionCode: number, status: ReleaseStatus, notes: string): string {
  const text = notes.trim();
  const release = {
    versionCodes: [String(versionCode)],
    status,
    ...(text ? { releaseNotes: [{ language: NOTES_LANGUAGE, text }] } : {}),
  };
  return JSON.stringify({ track, releases: [release] });
}

async function setTrack(client: Client, editId: string, request: UploadRequest, versionCode: number, status: ReleaseStatus): Promise<void> {
  await client.call(`set ${request.track} track`, `${client.base}/edits/${editId}/tracks/${encodeURIComponent(request.track)}`, {
    method: 'PUT',
    headers: JSON_HEADERS,
    body: trackBody(request.track, versionCode, status, request.notes),
  });
}

async function commitEdit(client: Client, editId: string): Promise<void> {
  await client.call('commit edit', `${client.base}/edits/${editId}:commit`, { method: 'POST' });
}

/** True when Play refused a completed release because the app has not passed its first review. */
export function isDraftAppRefusal(error: unknown): boolean {
  return error instanceof PlayError && error.status === 400 && DRAFT_APP_REFUSAL.test(error.message);
}

async function releaseWithFallback(client: Client, editId: string, request: UploadRequest, versionCode: number, log: Log): Promise<Outcome> {
  try {
    await setTrack(client, editId, request, versionCode, request.status);
    await commitEdit(client, editId);
    return 'committed';
  } catch (error) {
    if (request.status === 'draft' || !isDraftAppRefusal(error)) throw error;
    log(`Play refused a completed release (${(error as PlayError).message}); setting the ${request.track} track as draft instead.`);
    await setTrack(client, editId, request, versionCode, 'draft');
    await commitEdit(client, editId);
    return 'committed-as-draft';
  }
}

async function deleteEdit(client: Client, editId: string, log: Log): Promise<void> {
  try {
    await client.call('delete edit', `${client.base}/edits/${editId}`, { method: 'DELETE' });
    log(`Edit ${editId} deleted.`);
  } catch (error) {
    log(`Edit ${editId} was not deleted (${(error as Error).message}); Play expires it.`);
  }
}

/** The whole flow; the edit is deleted when any step after its insert fails. */
export async function uploadToTrack(request: UploadRequest, fetch: Fetch, log: Log): Promise<UploadResult> {
  const client = new Client(fetch, request.token, request.packageName);
  const editId = await insertEdit(client);
  log(`Edit ${editId} inserted for ${request.packageName}.`);
  try {
    const versionCode = await uploadBundle(client, editId, request.bundle);
    log(`Bundle uploaded as version code ${versionCode}.`);
    const onTrack = await trackVersionCodes(client, editId, request.track);
    if (decideUpload(onTrack, versionCode) === 'skip') {
      await deleteEdit(client, editId, log);
      return { outcome: 'track-already-ahead', editId, versionCode, status: null, trackVersionCode: Math.max(...onTrack) };
    }
    if (request.mapping) {
      await uploadMapping(client, editId, versionCode, request.mapping);
      log(`Mapping uploaded for version code ${versionCode}.`);
    }
    const outcome = await releaseWithFallback(client, editId, request, versionCode, log);
    const status = outcome === 'committed' ? request.status : 'draft';
    log(`Edit ${editId} committed: version code ${versionCode} on ${request.track} as ${status}.`);
    return { outcome, editId, versionCode, status };
  } catch (error) {
    await deleteEdit(client, editId, log);
    throw error;
  }
}

export function parseStatus(value: string): ReleaseStatus {
  if (value === 'completed' || value === 'draft') return value;
  throw new Error(`--status must be completed or draft, not '${value}'`);
}

function required(value: string | undefined, name: string): string {
  if (!value) throw new Error(`${name} is required`);
  return value;
}

async function optionalFile(path: string | undefined): Promise<Uint8Array<ArrayBuffer> | null> {
  return path ? new Uint8Array(await readFile(path)) : null;
}

async function main(): Promise<void> {
  const { values } = parseArgs({
    options: {
      bundle: { type: 'string' },
      package: { type: 'string' },
      mapping: { type: 'string' },
      notes: { type: 'string' },
      track: { type: 'string', default: 'internal' },
      status: { type: 'string', default: 'draft' },
    },
  });
  const request: UploadRequest = {
    token: required(process.env.PLAY_ACCESS_TOKEN, 'PLAY_ACCESS_TOKEN'),
    packageName: required(values.package, '--package'),
    track: values.track,
    status: parseStatus(values.status),
    bundle: new Uint8Array(await readFile(required(values.bundle, '--bundle'))),
    mapping: await optionalFile(values.mapping),
    notes: values.notes ? await readFile(values.notes, 'utf8') : '',
  };
  const result = await uploadToTrack(request, globalThis.fetch, (line) => console.log(line));
  if (result.outcome === 'track-already-ahead') {
    console.log(
      `::notice title=Play track is ahead::The ${request.track} track holds version code ${result.trackVersionCode}, not below this build's ${result.versionCode}, so nothing was committed.`,
    );
  }
  if (result.outcome === 'committed-as-draft') {
    console.log(
      `::notice title=Play release is a draft::The app has not passed its first review, so version code ${result.versionCode} is a draft on ${request.track}; roll it out in Play Console.`,
    );
  }
}

if (import.meta.main) {
  main().catch((error: unknown) => {
    console.error(`::error title=Play upload::${(error as Error).message}`);
    process.exit(1);
  });
}
