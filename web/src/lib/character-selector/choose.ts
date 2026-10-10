// web/src/lib/character-selector/choose.ts
// Choosing a character (spec decisions 8, 9 and 10): write the one pointer, remember the
// choice for the list order, tell every island, and on the two pages whose state lives in
// the URL point the URL at the new character so a refresh keeps it.
import type { MeCharacter } from '../account/api';
import { pointerForCharacter } from '../account/main-character';
import {
  CURRENT_CHARACTER_CHANGED,
  plannerHrefFor,
  simHrefFor,
  writeCurrent,
  type CurrentCharacter,
} from '../current-character';
import { SIM_TABS } from '../sim/tabs';
import { recordChosenKey } from './order';

export interface ChooseDeps {
  storage?: Storage;
  dispatch: (event: Event) => void;
  pathname: string;
  replaceUrl: (href: string) => void;
  now: () => Date;
}

/** The URL a refresh on `pathname` should land on for `pointer`, or null for any other page. */
export function urlForChosenPointer(pathname: string, pointer: CurrentCharacter): string | null {
  if (pathname === '/planner') return plannerHrefFor(pointer);
  const tab = SIM_TABS.find((entry) => entry.href === pathname);
  return tab === undefined ? null : simHrefFor(pointer, tab.id);
}

/**
 * Writes `pointer` as the current character. `accountKey` is the account character's key when
 * the pointer names one, so the list remembers the choice. Planner and simulator state lives
 * in the URL, so on those pages the URL is replaced first (no reload, no history entry) and
 * the page then re-renders in place from `CURRENT_CHARACTER_CHANGED`, as every other page does.
 */
export function choosePointer(pointer: CurrentCharacter, accountKey: string | null, deps: ChooseDeps): void {
  writeCurrent(pointer, deps.storage);
  if (accountKey !== null) recordChosenKey(accountKey, deps.storage);
  const href = urlForChosenPointer(deps.pathname, pointer);
  if (href !== null) deps.replaceUrl(href);
  deps.dispatch(new Event(CURRENT_CHARACTER_CHANGED));
}

export function chooseCharacter(character: MeCharacter, deps: ChooseDeps): void {
  const pointer = { ...pointerForCharacter(character), savedAt: deps.now().toISOString() };
  choosePointer(pointer, character.key, deps);
}

/** The browser wiring of `ChooseDeps`, built at the call site so tests can inject their own. */
export function browserChooseDeps(): ChooseDeps {
  return {
    dispatch: (event) => window.dispatchEvent(event),
    pathname: window.location.pathname,
    replaceUrl: (href) => window.history.replaceState(null, '', href),
    now: () => new Date(),
  };
}
