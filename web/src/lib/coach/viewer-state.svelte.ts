// web/src/lib/coach/viewer-state.svelte.ts
// Must be called while a component initialises (it registers an effect).
// The viewer, read once per page from the shared `/v1/me` query. Components create one and read
// `.viewer`; it is the signed-out viewer until the answer lands, so nothing coach-shaped paints
// early. A failed request reads as signed out: the coach never blocks the report.
import { fetchMeOnce } from '../account/api';
import { SIGNED_OUT_VIEWER, viewerFromMe, type CoachViewer } from './viewer';

export interface CoachViewerState {
  readonly viewer: CoachViewer;
  /** False until `/v1/me` has answered, so a slot can wait for it. */
  readonly resolved: boolean;
}

export function createCoachViewerState(apiBase?: string): CoachViewerState {
  let viewer = $state<CoachViewer>(SIGNED_OUT_VIEWER);
  let resolved = $state(false);
  // In an effect, so a server render (and the report's SSR tests) never asks.
  $effect(() => {
    let current = true;
    fetchMeOnce(apiBase)
      .then((me) => {
        if (current) viewer = viewerFromMe(me);
      })
      .catch(() => {
        if (current) viewer = SIGNED_OUT_VIEWER;
      })
      .finally(() => {
        if (current) resolved = true;
      });
    return () => {
      current = false;
    };
  });
  return {
    get viewer() {
      return viewer;
    },
    get resolved() {
      return resolved;
    },
  };
}
