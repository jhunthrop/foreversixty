// web/src/lib/guides/hit-cap-line.ts
// The Stat priority block's headline, worded the way a guide gives it: "Hit cap: 16% ·
// 6% from Precision, 10% still to find". The numbers are the same ones the BiS page's rail
// prints (`hitCapLine`), so the two pages cannot disagree; only the sentence differs. The talent
// that supplies the school hit is named from the build, never typed per guide.
import { bisCopy } from '../bis/copy';
import { expertiseIsGained } from '../bis/expertise-cap';
import { formatCapPercent } from '../bis/hit-cap';
import type { BisExpertiseToCap, BisHitToCap, BisSpellHitToCap } from '../bis/types';
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

// The band publishes how far its best-in-slot character is FROM each cap (`to_cap`, floored at 0)
// and what that character already has from its set, talents and buffs (`baseline`), never the
// cap itself (2026-11-02 launch defect: a tank with 8%
// expertise read "Expertise cap: 0% (dodge), 6% (parry)"). The cap a guide names is the two
// added together; a distance of 0 with a baseline means the character is already past it,
// so the cap reads "reached" rather than a number the data cannot pin down.

/** What each published `baseline` counts (`sim/cmd/leveling-bis/hitcap.go`, `expertisecap.go`). */
const MELEE_HIT_SOURCES = 'gear, talents and buffs';
const SPELL_HIT_SOURCES = 'gear and buffs';
const EXPERTISE_SOURCES = 'gear and talents';

/** The cap as the guide names it: `baseline + toCap`, or "reached" past it. */
function capText(baseline: number, toCap: number): string {
  return toCap <= 0 && baseline > 0 ? guidesCopy.capReached : `${formatCapPercent(baseline + toCap)}%`;
}

/** The split after the cap: what the band's best-in-slot set already has, what a school talent
 *  adds, and what is still short. Empty when the character starts at 0 and no talent adds hit (the cap is then
 *  all gear, and the cap alone says it). */
function splitText(
  baseline: number,
  gear: number,
  sources: string,
  talent?: { name: string; percent: number },
): string | undefined {
  if (baseline <= 0 && talent === undefined) return undefined;
  const parts = [
    ...(baseline > 0 ? [guidesCopy.capOnSet(formatCapPercent(baseline), 'hit', sources)] : []),
    ...(talent === undefined ? [] : [guidesCopy.capFrom(formatCapPercent(talent.percent), talent.name)]),
    ...(gear > 0 ? [guidesCopy.capStillToFind(formatCapPercent(gear))] : []),
  ];
  return parts.join(', ');
}

function spellHitLine(hitToCap: BisSpellHitToCap, talentName: string | undefined): GuideHitCapLine {
  const school = hitToCap.school;
  const talent =
    school === undefined
      ? undefined
      : { name: talentName ?? 'talents', percent: hitToCap.spell - school.to_cap };
  const gear = school === undefined ? hitToCap.spell : school.to_cap;
  return {
    text: guidesCopy.hitToCap(
      capText(hitToCap.baseline, hitToCap.spell),
      splitText(hitToCap.baseline, gear, SPELL_HIT_SOURCES, talent),
    ),
    title: bisCopy.spellHitToCapTitle,
  };
}

export function guideHitCapLine(
  hitToCap: BisHitToCap | null | undefined,
  talentName?: string,
): GuideHitCapLine | undefined {
  if (hitToCap === null || hitToCap === undefined) return undefined;
  if (hitToCap.kind === 'spell') return spellHitLine(hitToCap, talentName);
  const specials = capText(hitToCap.baseline, hitToCap.specials);
  const split = splitText(hitToCap.baseline, hitToCap.specials, MELEE_HIT_SOURCES);
  return {
    text:
      hitToCap.white === undefined
        ? guidesCopy.hitToCap(specials, split)
        : guidesCopy.hitToCapWhite(specials, capText(hitToCap.baseline, hitToCap.white), split),
    title: bisCopy.hitToCapTitle,
  };
}

/** The guide's expertise line ("Expertise cap: 6.5% (dodge)"); the BiS rail keeps its own wording. */
export function guideExpertiseCapLine(
  expertiseToCap: BisExpertiseToCap | null | undefined,
): GuideHitCapLine | undefined {
  if (expertiseToCap === null || expertiseToCap === undefined) return undefined;
  if (!expertiseIsGained(expertiseToCap)) return undefined;
  const { baseline, dodge, parry } = expertiseToCap;
  return {
    text: guidesCopy.expertiseCap(
      capText(baseline, dodge),
      parry === undefined ? undefined : capText(baseline, parry),
      baseline > 0
        ? guidesCopy.capOnSet(formatCapPercent(baseline), 'expertise', EXPERTISE_SOURCES)
        : undefined,
    ),
    title: bisCopy.expertiseToCapTitle,
  };
}
