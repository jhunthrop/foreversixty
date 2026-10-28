// web/src/lib/account/setup-next.ts
// Where /setup's `next` may send the visitor back to: only a guild claim page on this site.
// `safeNextPath` answers its fallback (never null) for anything off-site, so an empty fallback
// is the "no target" signal here.
import { parseGuildClaimPath, type CharacterPath } from '../characters';
import { safeNextPath } from './safe-next';

export interface SetupReturnTarget {
  /** The same-site path and query to link back to. */
  href: string;
  /** The guild whose claim page it is. */
  guild: CharacterPath;
}

const NO_TARGET = '';

export function setupReturnTarget(rawNext: string | null): SetupReturnTarget | null {
  const href = safeNextPath(rawNext, NO_TARGET);
  if (href === NO_TARGET) return null;
  const guild = parseGuildClaimPath(href.split(/[?#]/, 1)[0] ?? NO_TARGET);
  return guild === null ? null : { href, guild };
}
