// web/src/lib/planner/share.ts
// Saving a build. The API answers in the Phase 0 envelope { ok, data, error, request_id };
// 201 is a new id, 200 is the same build saved before, and both carry { id, url }.
import { API_BASE_URL, SITE_BASE_URL } from './config';
import type { BuildDraft } from './types';

export interface SavedBuild {
  id: string;
  url: string;
}

export type SaveOutcome =
  { ok: true; build: SavedBuild } | { ok: false; message: string; fields: Record<string, string> };

/** Copied from the spec's error handling section. */
export const RATE_LIMIT_MESSAGE = 'Too many saves from this connection; try again in an hour.';
export const SAVE_FAILED_MESSAGE = 'The build could not be saved; try again.';

interface Envelope {
  ok: boolean;
  data: SavedBuild | null;
  error: { message?: string; fields?: Record<string, string> } | null;
  request_id: string;
}

function failed(message: string, fields: Record<string, string> = {}): SaveOutcome {
  return { ok: false, message, fields };
}

export function shareUrlFor(id: string): string {
  return `${SITE_BASE_URL}/b/${id}`;
}

export function cardUrlFor(id: string): string {
  return `${shareUrlFor(id)}/card.png`;
}

export async function saveBuild(draft: BuildDraft, apiBase: string = API_BASE_URL): Promise<SaveOutcome> {
  let response: Response;
  try {
    response = await fetch(`${apiBase}/v1/builds`, {
      method: 'POST',
      headers: { 'content-type': 'application/json', accept: 'application/json' },
      body: JSON.stringify(draft),
    });
  } catch {
    return failed(SAVE_FAILED_MESSAGE);
  }

  if (response.status === 429) return failed(RATE_LIMIT_MESSAGE);

  let envelope: Envelope;
  try {
    envelope = (await response.json()) as Envelope;
  } catch {
    return failed(SAVE_FAILED_MESSAGE);
  }

  if (response.ok && envelope.data) return { ok: true, build: envelope.data };
  return failed(envelope.error?.message ?? SAVE_FAILED_MESSAGE, envelope.error?.fields ?? {});
}
