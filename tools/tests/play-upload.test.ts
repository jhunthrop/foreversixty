/**
 * tools/play-upload.ts against a stub of the Play Developer API (no network).
 *   node --test tools/tests/play-upload.test.ts
 */
import assert from 'node:assert/strict';
import { test } from 'node:test';
import { API_ROOT, decideUpload, type Fetch, PlayError, trackBody, uploadToTrack, type UploadRequest } from '../play-upload.ts';

interface Seen {
  method: string;
  url: string;
  headers: Record<string, string>;
  body: unknown;
}

const PACKAGE = 'gg.foreversixty.app';
const BASE = `${API_ROOT}/androidpublisher/v3/applications/${PACKAGE}`;
const UPLOAD = `${API_ROOT}/upload/androidpublisher/v3/applications/${PACKAGE}`;
const SESSION = 'https://upload.example/session-1';
const TOKEN = 'stub-token';
const BUNDLE = new Uint8Array([1, 2, 3, 4]);
const MAPPING = new Uint8Array([9, 9]);
const VERSION_CODE = 18;
const MAPPING_URL = `${UPLOAD}/edits/edit-1/apks/${VERSION_CODE}/deobfuscationFiles/proguard?uploadType=media`;
const DRAFT_REFUSAL = 'Only releases with status draft may be created on draft app.';

type Answer = (seen: Seen) => Response | undefined;

const json = (body: unknown, status = 200): Response =>
  new Response(JSON.stringify(body), { status, headers: { 'Content-Type': 'application/json' } });
const refusal = (message: string, status = 400): Response => json({ error: { code: status, message } }, status);

/** A stub that records every request and answers like Play unless an override answers first. */
function stub(override: Answer = () => undefined): { fetch: Fetch; seen: Seen[] } {
  const seen: Seen[] = [];
  const fetch: Fetch = async (url, init) => {
    const request: Seen = { method: init.method ?? 'GET', url, headers: init.headers as Record<string, string>, body: init.body };
    seen.push(request);
    const answered = override(request);
    if (answered) return answered;
    if (url === `${BASE}/edits` && request.method === 'POST') return json({ id: 'edit-1' });
    if (url.startsWith(`${UPLOAD}/edits/edit-1/bundles`)) return new Response('{}', { status: 200, headers: { Location: SESSION } });
    if (url === SESSION) return json({ versionCode: VERSION_CODE, sha256: 'abc' });
    if (url === MAPPING_URL) return json({ deobfuscationFile: { symbolType: 'proguard' } });
    if (url === `${BASE}/edits/edit-1/tracks/internal` && request.method === 'GET') return json({ track: 'internal' });
    if (url === `${BASE}/edits/edit-1/tracks/internal`) return json(JSON.parse(String(init.body)));
    if (url === `${BASE}/edits/edit-1:commit`) return json({ id: 'edit-1' });
    if (url === `${BASE}/edits/edit-1` && request.method === 'DELETE') return new Response(null, { status: 204 });
    return refusal(`unexpected ${request.method} ${url}`, 404);
  };
  return { fetch, seen };
}

const request = (status: UploadRequest['status'], extra: Partial<UploadRequest> = {}): UploadRequest => ({
  token: TOKEN,
  packageName: PACKAGE,
  track: 'internal',
  status,
  bundle: BUNDLE,
  mapping: null,
  notes: '',
  ...extra,
});
const quiet = (): void => {};
const steps = (seen: Seen[]): string[] => seen.map((s) => `${s.method} ${s.url}`);
const trackPuts = (seen: Seen[]): { releases: { status: string; releaseNotes?: unknown }[] }[] =>
  seen.filter((s) => s.method === 'PUT' && s.url.endsWith('/tracks/internal')).map((s) => JSON.parse(String(s.body)));

test('the edits flow: insert, resumable upload, mapping, track with notes, commit, each with the bearer', async () => {
  const { fetch, seen } = stub();
  const result = await uploadToTrack(request('completed', { mapping: MAPPING, notes: '- Raid sign-up.\n' }), fetch, quiet);
  assert.deepEqual(result, { outcome: 'committed', editId: 'edit-1', versionCode: VERSION_CODE, status: 'completed' });
  assert.deepEqual(steps(seen), [
    `POST ${BASE}/edits`,
    `POST ${UPLOAD}/edits/edit-1/bundles?uploadType=resumable`,
    `PUT ${SESSION}`,
    `GET ${BASE}/edits/edit-1/tracks/internal`,
    `POST ${MAPPING_URL}`,
    `PUT ${BASE}/edits/edit-1/tracks/internal`,
    `POST ${BASE}/edits/edit-1:commit`,
  ]);
  for (const s of seen) assert.equal(s.headers.Authorization, `Bearer ${TOKEN}`);
  assert.equal(seen[1].headers['X-Upload-Content-Length'], String(BUNDLE.byteLength));
  assert.equal(seen[2].body, BUNDLE);
  assert.equal(seen[4].body, MAPPING);
  assert.deepEqual(trackPuts(seen)[0], {
    track: 'internal',
    releases: [{ versionCodes: [String(VERSION_CODE)], status: 'completed', releaseNotes: [{ language: 'en-US', text: '- Raid sign-up.' }] }],
  });
});

test('no mapping and no notes: no mapping upload, no releaseNotes', async () => {
  const { fetch, seen } = stub();
  await uploadToTrack(request('draft'), fetch, quiet);
  assert.equal(seen.filter((s) => s.url === MAPPING_URL).length, 0);
  assert.deepEqual(trackPuts(seen)[0].releases[0], { versionCodes: [String(VERSION_CODE)], status: 'draft' });
});

test('trackBody leaves blank notes out', () => {
  assert.deepEqual(JSON.parse(trackBody('internal', 3, 'draft', '  \n')), { track: 'internal', releases: [{ versionCodes: ['3'], status: 'draft' }] });
});

test('the first-review refusal: the track again as draft, committed', async () => {
  let commits = 0;
  const { fetch, seen } = stub((s) => (s.url.endsWith(':commit') && commits++ === 0 ? refusal(DRAFT_REFUSAL) : undefined));
  const lines: string[] = [];
  const result = await uploadToTrack(request('completed'), fetch, (line) => lines.push(line));
  assert.equal(result.outcome, 'committed-as-draft');
  assert.equal(result.status, 'draft');
  assert.deepEqual(trackPuts(seen).map((t) => t.releases[0].status), ['completed', 'draft']);
  assert.ok(lines.some((line) => line.includes('draft app')));
});

test('a draft refusal on a draft request is not retried', async () => {
  const { fetch, seen } = stub((s) => (s.url.endsWith(':commit') ? refusal(DRAFT_REFUSAL) : undefined));
  await assert.rejects(uploadToTrack(request('draft'), fetch, quiet), PlayError);
  assert.equal(trackPuts(seen).length, 1);
});

test('any other refusal deletes the edit and fails with the step and message, never the token', async () => {
  const { fetch, seen } = stub((s) => (s.url === SESSION ? refusal('All uploaded bundles must be signed.', 403) : undefined));
  await assert.rejects(uploadToTrack(request('completed'), fetch, quiet), (error: unknown) => {
    assert.ok(error instanceof PlayError);
    assert.equal(error.step, 'upload bundle');
    assert.equal(error.status, 403);
    assert.doesNotMatch(error.message, new RegExp(TOKEN));
    return true;
  });
  assert.equal(seen.at(-1)?.method, 'DELETE');
});

test('a refused mapping deletes the edit and fails', async () => {
  const { fetch, seen } = stub((s) => (s.url === MAPPING_URL ? refusal('bad mapping') : undefined));
  await assert.rejects(uploadToTrack(request('draft', { mapping: MAPPING }), fetch, quiet), (e: unknown) => e instanceof PlayError && e.step === 'upload mapping');
  assert.equal(seen.at(-1)?.method, 'DELETE');
  assert.equal(seen.filter((s) => s.url.endsWith(':commit')).length, 0);
});

test('a 403 on insert (the service account not invited in Play Console) fails before any edit exists', async () => {
  const { fetch, seen } = stub((s) => (s.url === `${BASE}/edits` ? refusal('The caller does not have permission', 403) : undefined));
  await assert.rejects(uploadToTrack(request('draft'), fetch, quiet), (e: unknown) => e instanceof PlayError && e.status === 403 && e.step === 'insert edit');
  assert.equal(seen.length, 1);
});

const trackHolding = (releases: unknown): Answer => (s) =>
  s.method === 'GET' && s.url.endsWith('/tracks/internal') ? json({ track: 'internal', releases }) : undefined;

test('a track at an equal or higher code: the edit is deleted, nothing set or committed', async () => {
  for (const code of ['18', '30']) {
    const { fetch, seen } = stub(trackHolding([{ versionCodes: [code], status: 'completed' }]));
    const result = await uploadToTrack(request('completed', { mapping: MAPPING }), fetch, quiet);
    assert.equal(result.outcome, 'track-already-ahead');
    assert.equal(result.trackVersionCode, Number(code));
    assert.equal(seen.at(-1)?.method, 'DELETE');
    assert.equal(trackPuts(seen).length, 0);
    assert.equal(seen.filter((s) => s.url === MAPPING_URL).length, 0);
  }
});

test('a 404 reading a track with no release yet is an empty track; a 500 is an error', async () => {
  const empty = stub((s) => (s.method === 'GET' ? refusal('Track not found', 404) : undefined));
  assert.equal((await uploadToTrack(request('draft'), empty.fetch, quiet)).outcome, 'committed');
  const broken = stub((s) => (s.method === 'GET' ? refusal('backend error', 500) : undefined));
  await assert.rejects(uploadToTrack(request('draft'), broken.fetch, quiet), (e: unknown) => e instanceof PlayError && e.step === 'read track');
});

test('every request carries a timeout signal', async () => {
  const { fetch, seen } = stub();
  const signals: unknown[] = [];
  const watching: Fetch = (url, init) => {
    signals.push(init.signal);
    return fetch(url, init);
  };
  await uploadToTrack(request('draft', { mapping: MAPPING }), watching, quiet);
  assert.equal(signals.length, seen.length);
  for (const signal of signals) assert.ok(signal instanceof AbortSignal);
});

test('decideUpload', () => {
  assert.equal(decideUpload([], 18), 'upload');
  assert.equal(decideUpload([10, 17], 18), 'upload');
  assert.equal(decideUpload([18], 18), 'skip');
  assert.equal(decideUpload([10, 30], 18), 'skip');
});
