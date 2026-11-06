// web/src/lib/bis/follow-selector-run.ts
// Carries out `decideBisFollow` on a BiS spec page: reads the page's own state from the DOM,
// resolves the selected character, then either updates faction and band the way the band tabs
// do (the faction radio and the `#band-<faction>-<band>` hash) or moves to the character's page.
// Loaded by follow-selector-client.ts on the first pointer change, not with the page: it pulls in
// the account API, the hero-character lookup and the upgrades loader.
import { heroCharacter } from '../account/hero-character';
import { fetchMeOnce } from '../account/api';
import { specKeyForCharacter } from '../home/upgrades-loader';
import { readCurrent } from '../current-character';
import { classSlugFromName } from '../report/tree-sizes';
import { decideBisDefault, decideBisFollow, type BisPageState, type SelectedTarget } from './follow-selector';
import { MIN_BAND } from './hover';
import { sendBisShow } from './show-event';
import type { Faction } from './types';

/** The page's island publishes the shown faction and band on its root (`data-bis-faction`,
 *  `data-bis-band`) and listens for `bis:show`, so this module never reaches into its markup. */

function readPageState(root: HTMLElement): BisPageState {
  const faction: Faction = root.dataset.bisFaction === 'horde' ? 'horde' : 'alliance';
  const band = Number(root.dataset.bisBand);
  return {
    specKey: root.dataset.bisSpec ?? '',
    faction,
    band: Number.isFinite(band) && band > 0 ? band : MIN_BAND,
  };
}

async function selectedTarget(): Promise<SelectedTarget | null> {
  const current = readCurrent();
  if (current === null) return null;
  const me = await fetchMeOnce();
  const character = me === null ? null : heroCharacter(current, me.characters);
  if (character === null) return { classSlug: current.classSlug };
  return {
    classSlug: classSlugFromName(character.class ?? ''),
    specKey: specKeyForCharacter(character),
    level: character.level,
    faction: character.faction,
  };
}

function showBand(root: HTMLElement, faction: Faction, band: number): void {
  sendBisShow(root, { faction, band });
}

export async function follow(root: HTMLElement): Promise<void> {
  const target = await selectedTarget();
  if (target === null) return;
  const action = decideBisFollow(readPageState(root), target);
  if (action.kind === 'navigate') window.location.assign(action.href);
  else if (action.kind === 'update') showBand(root, action.faction, action.band);
}

/** At load with no band in the URL: the selected character's band and faction, never another page. */
export async function followDefault(root: HTMLElement): Promise<void> {
  const target = await selectedTarget();
  if (target === null) return;
  const action = decideBisDefault(readPageState(root), target);
  if (action.kind === 'update') showBand(root, action.faction, action.band);
}
