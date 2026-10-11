// web/src/lib/character-header/answer-plan.ts
// Which presentation the answer slot takes for a character: the page's level-floor sentence,
// the page's no-gear sentence and link, or a real answer to fetch. Pure, so the order of the
// checks is pinned by a test: below the floor wins over missing gear (a level 12 has no list).
import type { MeCharacter } from '../account/api';
import { hasGearOnFile } from './sync';
import type { AnswerSpec } from './types';

export type AnswerPlan = 'below-floor' | 'no-gear' | 'load';

export function planAnswer(
  character: MeCharacter,
  spec: Pick<AnswerSpec, 'levelFloor' | 'needsGear'>,
): AnswerPlan {
  if (spec.levelFloor !== undefined && character.level !== undefined && character.level < spec.levelFloor) {
    return 'below-floor';
  }
  if (spec.needsGear === true && !hasGearOnFile(character)) return 'no-gear';
  return 'load';
}
