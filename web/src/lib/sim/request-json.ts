// web/src/lib/sim/request-json.ts
// Design 8: "Advanced is a JSON editor, not a script language". SimulationCraft input is
// Raidbots' escape hatch; ours is the request envelope itself.
//
// Nothing here decides whether a request is valid: `api.SimRequest.Validate` runs in the
// wasm (contract 10.2, `SimPool.validate`) and the drawer shows whatever it says. This
// module turns a request into text, text back into a request, a request back into the
// page's own settings, and -- `checkRequest` below -- the three shapes the drawer's own
// question ("can this run?") can come back in. None of the three is a rule written here.
import { simCopy } from './copy';
import type { RequestValidation, RequestValidationError } from './engine';
import { defaultSettings, type SimSettings } from './settings';
import type { SimRequest } from './types';

/**
 * A generous bound on a hand-edited request. A real one is a couple of kilobytes; a bulk
 * request with four hundred candidates is under sixty. This exists so a pasted megabyte
 * is refused before `JSON.parse` is asked to do any work on it -- the same reason
 * `fs1.ts` and `url.ts` each carry one.
 */
export const MAX_REQUEST_CHARS = 262_144;

/** Indented, and newline-terminated so a textarea's last line is editable. */
export function formatRequest(request: SimRequest): string {
  return `${JSON.stringify(request, null, 2)}\n`;
}

export type ParsedRequest = { ok: true; request: SimRequest } | { ok: false; message: string };

export function parseRequest(text: string): ParsedRequest {
  if (text.length > MAX_REQUEST_CHARS) return { ok: false, message: simCopy.requestTooLong };
  let parsed: unknown;
  try {
    parsed = JSON.parse(text);
  } catch (error) {
    // The parser's own message names the line and column, which is the only thing that
    // says where the typo is. Ours says what kind of thing went wrong.
    return {
      ok: false,
      message: `${simCopy.requestNotJson} ${error instanceof Error ? error.message : ''}`.trim(),
    };
  }
  if (typeof parsed !== 'object' || parsed === null || Array.isArray(parsed)) {
    return { ok: false, message: simCopy.requestNotObject };
  }
  return { ok: true, request: parsed as SimRequest };
}

/**
 * `base`, overwritten by every key of `patch` that carries a *defined* value. A key
 * `patch` carries with an explicit `undefined` -- which a hand-built request missing an
 * optional field can, even though `JSON.parse` output never does -- is treated as absent
 * rather than as a deliberate blank, so `settingsFromRequest` below can fall back to the
 * default for it instead of rendering a `<select>` on nothing.
 */
function withDefined<T extends object>(base: T, patch: T): T {
  const merged = { ...base };
  for (const key of Object.keys(patch) as (keyof T)[]) {
    const value = patch[key];
    if (value !== undefined) merged[key] = value;
  }
  return merged;
}

/**
 * The page's settings, rebuilt from a request. Every encounter field the request does not
 * carry falls back to the default rather than staying undefined: the controls are
 * `<select>`s over closed vocabularies and an undefined value renders as a blank option.
 *
 * The preset is always "custom". A pasted buff list is nobody's preset, and filing it
 * under one would mean the next preset change silently discarded it.
 */
export function settingsFromRequest(request: SimRequest): SimSettings {
  const base = defaultSettings();
  return {
    encounter: withDefined(base.encounter, request.encounter),
    preset: 'custom',
    buffs: [...(request.character.buffs ?? [])],
    consumables: [...(request.character.consumes ?? [])],
    cooldowns: (request.character.cooldowns ?? []).map((row) => ({ ...row, at_sec: [...row.at_sec] })),
  };
}

/** The drawer's own question, answered. Three ways to fail, one way to pass. */
export type RequestCheck =
  | { status: 'parse-error'; message: string }
  | { status: 'invalid'; errors: RequestValidationError[] }
  | { status: 'validate-error'; message: string }
  | { status: 'valid'; request: SimRequest };

/**
 * "Can this run?", as one function over its inputs: parse the text, then ask the engine.
 * `validate` is `store.validateRequest` (`SimPool.validate` underneath), injected so this
 * stays a pure function rather than reaching for a pool itself -- which is also what makes
 * it testable without one, and is why `RequestDrawer.svelte` calls this rather than
 * carrying the same three-way branch inline: the component renders whichever `status`
 * comes back, and never decides one for itself.
 *
 * `pool.validate` REJECTS, rather than resolving `{ok:false}`, for an envelope the engine
 * cannot even parse into a request (`worker.test.ts`'s own "rejects validate... for a
 * malformed envelope", Task 5's hardening of the pool and the fake-worker lane). That is
 * caught here and reported as `validate-error`, never as `valid`: a rejected promise is
 * not a pass, whatever the drawer's own `parseRequest` already let through as syntactic
 * JSON.
 */
export async function checkRequest(
  text: string,
  validate: (json: string) => Promise<RequestValidation>,
): Promise<RequestCheck> {
  const parsed = parseRequest(text);
  if (!parsed.ok) return { status: 'parse-error', message: parsed.message };
  let validation: RequestValidation;
  try {
    validation = await validate(text);
  } catch (error) {
    return {
      status: 'validate-error',
      message: error instanceof Error ? error.message : simCopy.requestValidateFailed,
    };
  }
  return validation.ok
    ? { status: 'valid', request: parsed.request }
    : { status: 'invalid', errors: validation.errors };
}
