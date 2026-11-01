// web/src/lib/guides/hit-cap-line.ts
// The Stat priority block's headline, worded the way a guide gives it: "Hit cap: 16% ·
// 6% from Precision, 10% from gear". The numbers are the same ones the BiS page's rail
// prints (`hitCapLine`), so the two pages cannot disagree; only the sentence differs. The talent
// that supplies the school hit is named from the build, never typed per guide.
import { bisCopy } from '../bis/copy';
import { formatCapPercent } from '../bis/hit-cap';
import type { BisExpertiseToCap, BisHitToCap } from '../bis/types';
import type { DecodedTreeView } from './build-tree';
import { guidesCopy } from './copy';

export interface GuideHitCapLine {
  text: string;
  title: string;
}

/** The talents that add hit to some of a caster's spells (`BisSpellHitToCap.school`). */
export const SCHOOL_HIT_TALENTS: readonly string[] = ['Precision', "Nature's Reach", 'Suppression'];

/** The school-hit talent a build spends points in, if any. */
export function schoolHitTalent(trees: readonly DecodedTreeView[]): string | undefined {
  for (const tree of trees) {
    for (const talent of tree.talents) {
      if (talent.rank > 0 && SCHOOL_HIT_TALENTS.includes(talent.name)) return talent.name;
    }
  }
  return undefined;
}

export function guideHitCapLine(
  hitToCap: BisHitToCap | null | undefined,
  talentName?: string,
): GuideHitCapLine | undefined {
  if (hitToCap === null || hitToCap === undefined) return undefined;
  if (hitToCap.kind === 'spell') {
    const full = formatCapPercent(hitToCap.spell);
    const school = hitToCap.school;
    return {
      text:
        school === undefined
          ? guidesCopy.hitToCap(full)
          : guidesCopy.hitToCap(full, {
              talentPercent: formatCapPercent(hitToCap.spell - school.to_cap),
              talent: talentName ?? 'talents',
              gearPercent: formatCapPercent(school.to_cap),
            }),
      title: bisCopy.spellHitToCapTitle,
    };
  }
  const specials = formatCapPercent(hitToCap.specials);
  return {
    text:
      hitToCap.white === undefined
        ? guidesCopy.hitToCap(specials)
        : guidesCopy.hitToCapWhite(specials, formatCapPercent(hitToCap.white)),
    title: bisCopy.hitToCapTitle,
  };
}

/** The guide's expertise line ("Expertise cap: 6.5% (dodge)"); the BiS rail keeps its own wording. */
export function guideExpertiseCapLine(
  expertiseToCap: BisExpertiseToCap | null | undefined,
): GuideHitCapLine | undefined {
  if (expertiseToCap === null || expertiseToCap === undefined) return undefined;
  const parry = expertiseToCap.parry === undefined ? undefined : formatCapPercent(expertiseToCap.parry);
  return {
    text: guidesCopy.expertiseCap(formatCapPercent(expertiseToCap.dodge), parry),
    title: bisCopy.expertiseToCapTitle,
  };
}
