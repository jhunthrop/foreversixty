// web/src/lib/bis/follow-selector-client.ts
// Carries out `decideBisFollow` on a BiS spec page: reads the page's own state from the DOM,
// resolves the selected character, then either updates faction and band the way the band tabs
// do (the faction radio and the `#band-<faction>-<band>` hash) or moves to the character's page.
import { heroCharacter } from '../account/hero-character';
import { fetchMeOnce } from '../account/api';
import { specKeyForCharacter } from '../home/upgrades-loader';
import { onCurrentCharacterChange, readCurrent } from '../current-character';
import { classSlugFromName } from '../report/tree-sizes';
import { decideBisFollow, type BisPageState, type SelectedTarget } from './follow-selector';
import { MIN_BAND } from './hover';
import type { Faction } from './types';

const FACTION_PARAM = 'faction';
const BAND_HASH = /^#band-(?:alliance|horde)-(\d+)$/;

function currentFaction(): Faction {
  const horde = document.getElementById('faction-horde');
  return horde instanceof HTMLInputElement && horde.checked ? 'horde' : 'alliance';
}

function currentBand(): number {
  const match = BAND_HASH.exec(window.location.hash);
  return match === null ? MIN_BAND : Number(match[1]);
}

function readPageState(root: HTMLElement): BisPageState {
  return { specKey: root.dataset.bisSpec ?? '', faction: currentFaction(), band: currentBand() };
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

/** The band tabs' own mechanism: the faction radio, then the band's hash (which `:target` reads). */
function showBand(faction: Faction, band: number): void {
  const radio = document.getElementById(`faction-${faction}`);
  if (radio instanceof HTMLInputElement) radio.checked = true;
  const url = new URL(window.location.href);
  url.searchParams.set(FACTION_PARAM, faction);
  window.history.replaceState(null, '', url);
  window.location.hash = `band-${faction}-${band}`;
}

async function follow(root: HTMLElement): Promise<void> {
  const target = await selectedTarget();
  if (target === null) return;
  const action = decideBisFollow(readPageState(root), target);
  if (action.kind === 'navigate') window.location.assign(action.href);
  else if (action.kind === 'update') showBand(action.faction, action.band);
}

export function mountBisFollow(root: HTMLElement): void {
  onCurrentCharacterChange((change) => {
    if (!change.fromPageLoad) void follow(root);
  });
}
