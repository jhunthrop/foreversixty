// web/src/lib/sim/action-mark.ts
// The one icon decision of ActionMark.svelte: the action's own, then the buff it is.
import activeBuild from '../../data/active-build.json';
import { dataUrl } from '../planner/load';
import type { ActionVisuals } from './action-visuals';

/** The rage an ability hands back when it misses: the engine's, with no client icon (R1's one exception). */
export const REFUND_KEY = 'other:refund';

export function resolveIconUrl(action: string | null, visuals: ActionVisuals, fallback: string | null): string | null {
  const own = action === null ? null : visuals(action).icon;
  if (own !== null) return own;
  return fallback === null || fallback === '' ? null : dataUrl(activeBuild.build, `icons/${fallback}.webp`);
}
