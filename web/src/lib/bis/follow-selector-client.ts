// web/src/lib/bis/follow-selector-client.ts
// Makes a BiS spec page follow the header character selector. The page only has to listen: it
// never reads the pointer at load (the server-rendered page already shows the default faction
// and band), so the work -- the account API, the character lookup, the upgrades loader, about
// 35 modules -- loads on the first pointer change (follow-selector-run.ts) instead of with the
// page. Loaded eagerly it started fetching just before the page's first paint, and the
// Lighthouse trace then counted all of it toward the BiS LCP, which swung between ~1.95 s and
// ~2.7 s run to run against a 1.9 s budget.
import { onCurrentCharacterChange } from '../current-character-events';

export function mountBisFollow(root: HTMLElement): () => void {
  return onCurrentCharacterChange((change) => {
    if (change.fromPageLoad) return;
    void import('./follow-selector-run').then(({ follow }) => follow(root));
  });
}
