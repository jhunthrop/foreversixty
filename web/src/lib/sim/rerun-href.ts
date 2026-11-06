// web/src/lib/sim/rerun-href.ts
// "Open in the sim" on a saved run (spec 2026-11-06-sim §4.5): /sim with the run's own source,
// or, for a source that stored no ref (an addon paste, a manual code), the stored character as a
// code. A link, so it works the moment the page paints, before the island mounts.
import { codeForCharacterSpec } from './character';
import type { SimResult } from './types';
import { defaultSimState, simSearch, withSimState } from './url';

export function rerunHref(result: SimResult, treeVersion: string): string {
  const { kind, ref } = result.request.source;
  const target =
    ref !== ''
      ? withSimState(defaultSimState(), { source: kind, ref })
      : withSimState(defaultSimState(), {
          code: codeForCharacterSpec(result.request.character, treeVersion),
        });
  return `/sim${simSearch(target)}`;
}
