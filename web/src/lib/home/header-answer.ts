// web/src/lib/home/header-answer.ts
// What the home signed-in panel hands the shared character header (components/character/
// CharacterHeader.svelte, spec 2026-10-11 section 4): the answer "Upgrades in your band" as two
// figures at the same size, the count and the gain, from the same comparison "Your upgrades"
// shows below it (`cachedUpgradesFor`), and the one action, which jumps to that table. Below
// level 20 the answer is the next talent point and the action opens the leveling guide
// (design/specs/2026-11-04-low-level-home.md).
import type { MeCharacter } from '../account/api';
import { bisCopy } from '../bis/copy';
import { FIRST_BAND_LEVEL } from '../bis/landing-bands';
import { landingCopy } from '../bis/landing-copy';
import { scoreUnitWord } from '../bis/tank-view';
import type {
  ActionContext,
  AnswerSpec,
  FiguresAnswer,
  HeaderAction,
  HeaderAnswer,
} from '../character-header/types';
import { homeHeroCardsCopy, homePanelCopy, levelingHomeCopy } from '../home-panel-copy';
import {
  levelingGuideHref,
  loadLevelingPlans,
  specNameOf,
  talentStepFor,
  type LevelingPlan,
} from '../leveling/leveling-plan';
import { FIRST_POINT_LEVEL } from '../planner/types';
import {
  cachedUpgradesFor,
  homeSpecKeyFor,
  readsAsLevelingSpec,
  type CachedUpgrades,
} from './upgrades-loader';

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
  const specKey = homeSpecKeyFor(character);
  if (specKey === undefined) return { kind: 'sentence', text: homeHeroCardsCopy.bestInSlotNoListYet };
  const cached = await cachedUpgradesFor(character);
  if (cached === null) return { kind: 'sentence', text: homeHeroCardsCopy.bestInSlotNoListYet };
  const answer = upgradesAnswer(cached);
  const specName = specNameOf(specKey);
  if (!readsAsLevelingSpec(character) || specName === undefined || answer.kind !== 'figures') return answer;
  const note = answer.note?.text ?? '';
  return { ...answer, note: { text: `${note} · ${levelingHomeCopy.comparedAs(specName)}` } };
}

/** The plan a level 1 to 19 character follows: its own spec's when it has one, else its class's leveling spec. */
export async function levelingPlanFor(character: MeCharacter): Promise<LevelingPlan | undefined> {
  const specKey = homeSpecKeyFor(character);
  if (specKey === undefined) return undefined;
  return (await loadLevelingPlans())[specKey];
}

/** Below 20: when the first point comes, or which talent this level's point goes into. */
export async function loadLevelingAnswer(character: MeCharacter): Promise<HeaderAnswer> {
  const level = character.level ?? 1;
  if (level < FIRST_POINT_LEVEL)
    return { kind: 'sentence', text: levelingHomeCopy.headerFirstPoint(FIRST_POINT_LEVEL) };
  const plan = await levelingPlanFor(character);
  if (plan === undefined) return { kind: 'sentence', text: levelingHomeCopy.headerNoPlan };
  const step = talentStepFor(plan, level, character.build?.talents);
  const next = plan.points[step.nextIndex];
  if (next === undefined) return { kind: 'sentence', text: levelingHomeCopy.orderDone };
  if (step.toSpend === 0)
    return { kind: 'sentence', text: levelingHomeCopy.headerNextPoint(step.nextPointLevel) };
  if (character.build?.talents === undefined) {
    return { kind: 'sentence', text: levelingHomeCopy.headerPointsToSpend(step.toSpend, next.name) };
  }
  return { kind: 'sentence', text: levelingHomeCopy.headerLevelTalent(next.level, next.name) };
}

export const homeAnswerSpec: AnswerSpec = {
  levelFloor: FIRST_BAND_LEVEL,
  needsGear: true,
  strings: {
    belowLevelFloor: levelingHomeCopy.headerNoPlan,
    noGear: {
      sentence: homePanelCopy.headerNoGear,
      link: { label: landingCopy.noSyncLink, href: landingCopy.noSyncHref },
    },
    failed: homePanelCopy.headerFailed,
  },
  load: loadUpgradesAnswer,
  loadBelowFloor: loadLevelingAnswer,
};

/** Jumps to "Your upgrades"; under level 20 it opens the leveling spec's guide at its Leveling section. */
export function homeAction({ character, belowLevelFloor }: ActionContext): HeaderAction {
  if (!belowLevelFloor) return { label: homePanelCopy.headerAction, href: '#upgrades' };
  const specKey = homeSpecKeyFor(character);
  const specName = specKey === undefined ? undefined : specNameOf(specKey);
  const href = specKey === undefined ? undefined : levelingGuideHref(specKey);
  if (specName === undefined || href === undefined) {
    return { label: levelingHomeCopy.openLevelingGuide, href: '/guides' };
  }
  return { label: levelingHomeCopy.levelAs(specName), href };
}
