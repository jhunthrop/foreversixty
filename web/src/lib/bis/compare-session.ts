// web/src/lib/bis/compare-session.ts
// The BiS page's signed-in half, imported only after mount: once a session hint exists it reads
// the header's selected character (the same /v1/me the header loaded, R3) and, when that
// character is this class, builds the doll's compare data for the shown band. It re-runs when the
// selected character changes or the page shows another band. Signed out, it does nothing.
import { fetchMeOnce } from '../account/api';
import { selectedCharacter } from '../account/hero-character';
import { onCurrentCharacterChange, readCurrent } from '../current-character';
import { afterLoad, hasSessionHint } from '../page-load';
import { comparable, loadDollCharacter } from './compare-loader';
import type { DollCharacter } from './doll-view';
import type { BisPresetId, Faction } from './types';

export interface ComparePage {
  build: string;
  spec: string;
  classSlug: string;
}

export interface ShownBand {
  band: number;
  faction: Faction;
  preset: BisPresetId;
}

export interface CompareSession {
  /** The page now shows this band: rebuild the compare data for it. */
  show(band: ShownBand): void;
  stop(): void;
}

const IDLE_SESSION: CompareSession = { show: () => {}, stop: () => {} };

/** Calls `onCharacter` with the compare data (or undefined) for each band the page shows. */
export function watchCompare(
  page: ComparePage,
  first: ShownBand,
  onCharacter: (character: DollCharacter | undefined, race: string | undefined) => void,
): CompareSession {
  if (!hasSessionHint()) return IDLE_SESSION;
  let stopped = false;
  let shown = first;
  let ready = false;
  const run = async (): Promise<void> => {
    const band = shown;
    const me = await fetchMeOnce().catch(() => null);
    const character = me === null ? null : selectedCharacter(readCurrent(), me);
    if (stopped || band !== shown) return;
    if (!comparable(character, page.classSlug)) {
      onCharacter(undefined, undefined);
      return;
    }
    const loaded = await loadDollCharacter(character, { ...page, ...band }).catch(() => null);
    if (!stopped && band === shown)
      onCharacter(loaded ?? undefined, character.race?.toLowerCase().replace(/\s+/g, '-'));
  };
  afterLoad(() => {
    ready = true;
    void run();
  });
  const stopChange = onCurrentCharacterChange(() => {
    if (ready) void run();
  });
  return {
    show(band) {
      const same = band.band === shown.band && band.faction === shown.faction && band.preset === shown.preset;
      shown = band;
      if (!same && ready) void run();
    },
    stop() {
      stopped = true;
      stopChange();
    },
  };
}
