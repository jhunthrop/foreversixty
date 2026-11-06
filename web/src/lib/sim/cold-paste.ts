// web/src/lib/sim/cold-paste.ts
// The /sim paste Plate's local check (spec 2026-11-06-sim §5, "Paste plate"): Run stays disabled
// until the paste decodes, and a paste that does not decode says so under the field. The box takes
// an addon export or an unsaved build's FS1 code, bare or wrapped in a planner link's ?code=
// (build-input.ts's parseBuildInput), and decodes it with the planner's own grammar (fs1.ts). A
// bare saved-build id has no local decode; "From a build" under the Plate loads that one.
import { decodeFS1 } from '../planner/fs1';
import { parseBuildInput } from './build-input';
import { simPageCopy } from './page-copy';

export type ColdPasteResult =
  | { ok: true; code: string }
  | {
      ok: false;
      /** Null for a still-empty box: that is not yet a mistake, so Run stays disabled
       *  with nothing said about it. */
      message: string | null;
    };

export function validateColdPaste(raw: string): ColdPasteResult {
  const trimmed = raw.trim();
  if (trimmed === '') return { ok: false, message: null };
  const input = parseBuildInput(trimmed);
  if (input === null || input.kind !== 'code') return { ok: false, message: simPageCopy.paste.invalid };
  const decoded = decodeFS1(input.code);
  return decoded.ok ? { ok: true, code: input.code } : { ok: false, message: decoded.message };
}
