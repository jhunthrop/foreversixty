// web/src/lib/sim/cold-paste.ts
// The /sim hero's cold-paste box (persona review 2026-10-06, §5/§8 item 1): a visitor
// signed out, or signed in with no tracked character, pastes an addon export or an unsaved
// build's FS1 code straight into the hero and the page runs it -- no Battle.net round trip,
// no queue, because the engine runs as wasm in this browser (engine.ts) and the decode is
// local. "Decodes" means exactly what it says: the identical grammar AddonPasteBox.svelte
// already proves on /addon (planner/fs1.ts's decodeFS1), reached the same way the switcher
// below reaches it for "From a build"'s own unsaved-code case (build-input.ts's
// parseBuildInput). A bare saved-build id is not accepted here -- it has no local decode at
// all, only a server round trip, and that path already exists unchanged on the "From a
// build" card beneath this box.
import { decodeFS1 } from '../planner/fs1';
import { parseBuildInput } from './build-input';

export type ColdPasteResult =
  | { ok: true; code: string }
  | {
      ok: false;
      /** Null for a still-empty box: that is not yet a mistake, so Run stays disabled
       *  with nothing said about it. */
      message: string | null;
    };

const INVALID_MESSAGE = 'That does not look like an addon export or a build code.';

export function validateColdPaste(raw: string): ColdPasteResult {
  const trimmed = raw.trim();
  if (trimmed === '') return { ok: false, message: null };
  const input = parseBuildInput(trimmed);
  if (input === null || input.kind !== 'code') return { ok: false, message: INVALID_MESSAGE };
  const decoded = decodeFS1(input.code);
  return decoded.ok ? { ok: true, code: input.code } : { ok: false, message: decoded.message };
}
