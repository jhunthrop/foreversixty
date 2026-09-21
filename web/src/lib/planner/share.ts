// web/src/lib/planner/share.ts
// Saving a build. The API answers in the Phase 0 envelope { ok, data, error, request_id };
// 201 is a new id, 200 is the same build saved before, and both carry { id, url }.
import { API_BASE_URL } from './config';
import { plannerCopy } from './copy';
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
/**
 * Shown when the draft cannot even be composed: its class or race is not one the loaded
 * reference data has, so there is no id to post. The store repairs both on every write it
 * owns, so this is the panel's backstop against a future writer that forgets -- never a
 * button left reading "Saving" with nothing said.
 */
export const UNSAVABLE_BUILD_MESSAGE =
  'This build’s class and race are not ones the site knows, so it cannot be saved.';

interface Envelope {
  ok: boolean;
  data: SavedBuild | null;
  error: { message?: string; fields?: Record<string, string> } | null;
  request_id: string;
}

function failed(message: string, fields: Record<string, string> = {}): SaveOutcome {
  return { ok: false, message, fields };
}

/**
 * The preview card the API renders for a saved build, derived from the `url` the save
 * returned rather than from a second copy of the origin held here. The contract builds that
 * url as `PUBLIC_BASE_URL + "/b/" + id`, so on any deployment whose base url is not this
 * bundle's idea of the site -- a preview domain, a staging API -- a locally composed card
 * url would point at a host that has no such image.
 */
export function cardUrlFor(buildUrl: string): string {
  return `${buildUrl}/card.png`;
}

/**
 * What sharing this draft actually posts publicly, named for the confirm step -- one place
 * that lists it, so the confirm's copy can never drift from `POST /v1/builds`'s own body
 * (`class_id`, `race_id`, `point_order`, `gear`, and `title` when set). Class/race, talent
 * order and gear are always listed: `store.toDraft()` always fills `gear` (`{}` when
 * nothing is equipped, never omitted), so the request always carries that key, the same as
 * it always carries the class and the talent order. Title is listed only when the draft
 * has one -- the one field `toDraft()` genuinely omits rather than sends empty.
 * `includeSim` is the caller's own "Include a sim on the card" checkbox, effectively
 * checked -- ticked and a live estimate is actually ready to run from -- since a sim result
 * is only ever saved (`attachSim` -> `saveSim`) when both hold.
 */
export function sharedFields(draft: BuildDraft, includeSim: boolean): string[] {
  const fields: string[] = [
    plannerCopy.shareFieldClassRace,
    plannerCopy.shareFieldTalents,
    plannerCopy.shareFieldGear,
  ];
  if (draft.title !== undefined && draft.title.length > 0) {
    fields.push(plannerCopy.shareFieldTitle(draft.title));
  }
  if (includeSim) fields.push(plannerCopy.shareFieldSim);
  return fields;
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
