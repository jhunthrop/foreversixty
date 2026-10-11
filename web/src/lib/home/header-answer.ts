// web/src/lib/home/header-answer.ts
// What the home signed-in panel hands the shared character header (components/character/
// CharacterHeader.svelte, spec 2026-10-11 section 4): the answer "Upgrades in your band" as two
// figures at the same size, the count and the gain, from the same comparison "Your upgrades"
// shows below it (`cachedUpgradesFor`), and the one action, which jumps to that table. Below
// level 20 the action opens the first band's list instead.
import type { MeCharacter } from '../account/api';
import { bisCopy } from '../bis/copy';
import { bisPageHref } from '../bis/hover';
import { DEFAULT_LANDING_BAND, FIRST_BAND_LEVEL } from '../bis/landing-bands';
import { landingCopy } from '../bis/landing-copy';
import { scoreUnitWord } from '../bis/tank-view';
import type {
  ActionContext,
  AnswerSpec,
  FiguresAnswer,
  HeaderAction,
  HeaderAnswer,
} from '../character-header/types';
import { homeHeroCardsCopy, homePanelCopy } from '../home-panel-copy';
import { SPEC_BY_KEY } from '../sim/specs';
import { cachedUpgradesFor, specKeyForCharacter, type CachedUpgrades } from './upgrades-loader';

export function upgradesNote(band: number, notSimCheckedCount: number): string {
  const base = homePanelCopy.headerNote(bisCopy.bandRangeLabel(band));
  if (notSimCheckedCount === 0) return base;
  return `${base} · ${homePanelCopy.headerNotSimChecked(notSimCheckedCount)}`;
}

/** The count and, when any slot has a measured gain, the gain at the same size beside it. */
export function upgradesAnswer(cached: CachedUpgrades): HeaderAnswer {
  const { upgrades, totalGainDps, notSimCheckedCount, scoreUnit } = cached.result;
  if (upgrades.length === 0) {
    return { kind: 'sentence', text: homePanelCopy.headerAllMatch(bisCopy.bandRangeLabel(cached.band)) };
  }
  const gain =
    totalGainDps > 0 ? [{ value: `+${totalGainDps.toFixed(1)}`, unit: scoreUnitWord(scoreUnit) }] : [];
  const count = { value: String(upgrades.length), unit: upgrades.length === 1 ? 'upgrade' : 'upgrades' };
  const figures = (gain.length === 0 ? [count] : [count, gain[0]!]) as FiguresAnswer['figures'];
  return {
    kind: 'figures',
    label: homePanelCopy.headerAnswerLabel,
    figures,
    note: { text: upgradesNote(cached.band, notSimCheckedCount) },
  };
}

async function loadUpgradesAnswer(character: MeCharacter): Promise<HeaderAnswer> {
  if (specKeyForCharacter(character) === undefined) {
    return { kind: 'sentence', text: homeHeroCardsCopy.pickASpecLine };
  }
  const cached = await cachedUpgradesFor(character);
  if (cached === null) return { kind: 'sentence', text: homeHeroCardsCopy.bestInSlotNoListYet };
  return upgradesAnswer(cached);
}

export const homeAnswerSpec: AnswerSpec = {
  levelFloor: FIRST_BAND_LEVEL,
  needsGear: true,
  strings: {
    belowLevelFloor: landingCopy.under20,
    noGear: {
      sentence: homePanelCopy.headerNoGear,
      link: { label: landingCopy.noSyncLink, href: landingCopy.noSyncHref },
    },
    failed: homePanelCopy.headerFailed,
  },
  load: loadUpgradesAnswer,
};

/** Jumps to "Your upgrades"; under level 20 there is no table, so it opens the first band's list. */
export function homeAction({ character, belowLevelFloor }: ActionContext): HeaderAction {
  if (!belowLevelFloor) return { label: homePanelCopy.headerAction, href: '#upgrades' };
  const specKey = specKeyForCharacter(character);
  const row = specKey === undefined ? undefined : SPEC_BY_KEY[specKey];
  if (specKey === undefined || row === undefined) {
    return { label: landingCopy.fallbackAction, href: '/bis' };
  }
  return {
    label: landingCopy.leadButton(row.name, DEFAULT_LANDING_BAND),
    href: bisPageHref(specKey, character.faction ?? 'alliance', DEFAULT_LANDING_BAND),
  };
}
