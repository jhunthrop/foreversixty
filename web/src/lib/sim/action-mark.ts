// web/src/lib/sim/action-mark.ts
// The one icon decision of ActionMark.svelte: the action's own, then the buff it is.
import activeBuild from '../../data/active-build.json';
import { parseActionKey } from './action-names';
import type { ActionVisuals } from './action-visuals';
import { iconUrl } from '../ui/assets';

/** The rage an ability hands back when it misses: the engine's, with no client icon (R1's one exception). */
export const REFUND_KEY = 'other:refund';

export function resolveIconUrl(
  action: string | null,
  visuals: ActionVisuals,
  fallback: string | null,
): string | null {
  const own = action === null ? null : visuals(action).icon;
  if (own !== null) return own;
  return iconUrl(activeBuild.build, fallback);
}

/** The tooltip an action opens: a spell the build has a tooltip for, or an item (their ids). */
export function actionTip(
  action: string | null,
  spellKnown: (spellId: number) => boolean,
): { spell: number | undefined; item: number | undefined } {
  const parsed = action === null ? null : parseActionKey(action);
  return {
    spell: parsed?.kind === 'spell' && spellKnown(parsed.id) ? parsed.id : undefined,
    item: parsed?.kind === 'item' ? parsed.id : undefined,
  };
}
