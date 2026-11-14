// web/src/lib/guides/hit-cap-line.ts
// The Stat priority block's headline, worded the way a guide gives it: "Hit cap: 16% ·
// 6% from Precision, 10% still to find". The numbers are the same ones the BiS page's rail
// prints (`hitCapLine`), so the two pages cannot disagree; only the sentence differs. Each
// source the published figure counts is named: the BiS gear and buffs, and each talent of the
// weights character's own build that adds to it (`talent-hit.ts`), never typed per guide.
import { bisCopy } from '../bis/copy';
import { expertiseIsGained } from '../bis/expertise-cap';
import { formatCapPercent } from '../bis/hit-cap';
import type { BisExpertiseToCap, BisHitToCap, BisMeleeHitToCap, BisSpellHitToCap } from '../bis/types';
import { guidesCopy } from './copy';
import type { TalentHit, TalentHitScope } from './talent-hit';

export interface GuideHitCapLine {
  /** The whole line, cap and breakdown joined: "Hit cap: 16% · 6% from Precision, 10% still to find". */
  text: string;
  title: string;
  /** The cap alone: "Hit cap: 16%", "Hit cap: reached (weapon attacks), 16% (spells)". */
  cap: string;
  /** The breakdown after the cap, when there is one: "6% from Precision, 10% still to find". The
   *  page sets it on its own muted line, so a long melee-and-spells cap no longer wraps to three
   *  lines at phone width (round 20). */
  detail: string | undefined;
}

/** The two parts and their join, so `text` can never disagree with `cap` and `detail`. */
function capLine(cap: string, detail: string | undefined, title: string): GuideHitCapLine {
  return { text: guidesCopy.capWithDetail(cap, detail), title, cap, detail };
}

// The band publishes how far its best-in-slot character is FROM each cap (`to_cap`, floored at 0)
// and what that character already has (`baseline`), never the cap itself (2026-11-02 launch
// defect: a tank with 8% expertise read "Expertise cap: 0% (dodge), 6% (parry)"). The cap a
// guide names is the two added together; a distance of 0 with a baseline means the character
// is already past it, so the cap reads "reached" rather than a number the data cannot pin down.

/** What a baseline counts besides talents (`sim/cmd/leveling-bis/hitcap.go`, `expertisecap.go`). */
const HIT_GEAR_SOURCES = 'gear and buffs';
const EXPERTISE_GEAR_SOURCES = 'gear';
/** Every source at once, for a baseline its talents do not reconcile with (never guessed apart). */
const HIT_ALL_SOURCES = 'gear, buffs and talents';
const EXPERTISE_ALL_SOURCES = 'gear and talents';

/** The talent scopes each baseline counts (talent-hit.ts). */
const MELEE_BASELINE: readonly TalentHitScope[] = ['hit'];
const SPELL_BASELINE: readonly TalentHitScope[] = ['hit', 'spell'];
/** A caster with no school figure: the engine reads only the spells the rotation casts, so a
 *  school talent is then on every one of them (Shadow Focus for a Shadow priest) and in the
 *  baseline. */
const SPELL_BASELINE_ONE_SCHOOL: readonly TalentHitScope[] = [...SPELL_BASELINE, 'school'];
const EXPERTISE_BASELINE: readonly TalentHitScope[] = ['expertise'];

/** Float noise in the published figures (8.9 + 0.1 is 8.999...). */
const EPSILON = 0.05;

/** The cap as the guide names it: `baseline + toCap`, or "reached" past it. */
function capText(baseline: number, toCap: number): string {
  return toCap <= 0 && baseline > 0 ? guidesCopy.capReached : `${formatCapPercent(baseline + toCap)}%`;
}

function sum(talents: readonly TalentHit[]): number {
  return talents.reduce((total, talent) => total + talent.percent, 0);
}

function talentParts(talents: readonly TalentHit[]): string[] {
  return talents.map((talent) => guidesCopy.capFrom(formatCapPercent(talent.percent), talent.name));
}

/** The baseline's parts: the gear (and buffs) share, then each talent. When the talents add
 *  up to more than the baseline the two cannot be told apart, so it reads as one figure. */
function baselineParts(
  baseline: number,
  talents: readonly TalentHit[],
  stat: string,
  gearSources: string,
  allSources: string,
): string[] {
  const gear = baseline - sum(talents);
  if (gear < -EPSILON) {
    return baseline > 0 ? [guidesCopy.capOnSet(formatCapPercent(baseline), stat, allSources)] : [];
  }
  return [
    ...(gear > EPSILON ? [guidesCopy.capOnSet(formatCapPercent(gear), stat, gearSources)] : []),
    ...talentParts(talents),
  ];
}

/** "6% still to find" when `toFind` is above 0 (worded by `words`), else nothing. */
function stillToFind(
  toFind: number,
  words: (percent: string) => string = guidesCopy.capStillToFind,
): string | undefined {
  return toFind > 0 ? words(formatCapPercent(toFind)) : undefined;
}

/** The split after the cap: where the hit comes from, then what is still short. Empty when
 *  nothing counts toward the cap yet (the cap is then all gear, and the cap alone says it). */
function splitText(sources: readonly string[], short: string | undefined): string | undefined {
  if (sources.length === 0) return undefined;
  return [...sources, ...(short === undefined ? [] : [short])].join(', ');
}

/** What a melee band is still short of: the weapon cap first, then, past it, the spell cap
 *  of a spec that also casts damaging spells. */
function meleeStillToFind(hitToCap: BisMeleeHitToCap): string | undefined {
  if (hitToCap.specials > 0 || hitToCap.spell === undefined) return stillToFind(hitToCap.specials);
  return stillToFind(hitToCap.spell, guidesCopy.capStillToFindForSpells);
}

/** The melee caps with what each is for; one bare cap when there is only the weapon's. */
function meleeCaps(hitToCap: BisMeleeHitToCap): { cap: string; label: string }[] {
  const { baseline, white, spell } = hitToCap;
  const specials = capText(baseline, hitToCap.specials);
  return [
    { cap: specials, label: white === undefined ? guidesCopy.capLabelWeapon : guidesCopy.capLabelSpecials },
    ...(white === undefined ? [] : [{ cap: capText(baseline, white), label: guidesCopy.capLabelWhite }]),
    ...(spell === undefined ? [] : [{ cap: capText(baseline, spell), label: guidesCopy.capLabelSpells }]),
  ];
}

function inScope(talents: readonly TalentHit[], scopes: readonly TalentHitScope[]): TalentHit[] {
  return talents.filter((talent) => scopes.includes(talent.scope));
}

/** A school bonus's parts, each saying which spells get it: its talents when they account
 *  for it, else "talents". */
function schoolParts(bonus: number, schools: readonly string[], talents: readonly TalentHit[]): string[] {
  if (bonus <= EPSILON) return [];
  const names = schools.join(' and ');
  const school = inScope(talents, ['school']);
  const sources = Math.abs(sum(school) - bonus) <= EPSILON ? school : [{ name: 'talents', percent: bonus }];
  return sources.map((talent) =>
    guidesCopy.capFromOnSchools(formatCapPercent(talent.percent), talent.name, names),
  );
}

function spellHitLine(hitToCap: BisSpellHitToCap, talents: readonly TalentHit[]): GuideHitCapLine {
  const school = hitToCap.school;
  const toFind = school === undefined ? hitToCap.spell : school.to_cap;
  const sources = [
    ...baselineParts(
      hitToCap.baseline,
      inScope(talents, school === undefined ? SPELL_BASELINE_ONE_SCHOOL : SPELL_BASELINE),
      'hit',
      HIT_GEAR_SOURCES,
      HIT_ALL_SOURCES,
    ),
    ...(school === undefined ? [] : schoolParts(school.hit - hitToCap.baseline, school.names, talents)),
  ];
  return capLine(
    guidesCopy.hitToCap(capText(hitToCap.baseline, hitToCap.spell)),
    splitText(sources, stillToFind(toFind)),
    bisCopy.spellHitToCapTitle,
  );
}

/** `talents`: the weights character's own (`talentHitsOf` over the band's talents). */
export function guideHitCapLine(
  hitToCap: BisHitToCap | null | undefined,
  talents: readonly TalentHit[] = [],
): GuideHitCapLine | undefined {
  if (hitToCap === null || hitToCap === undefined) return undefined;
  if (hitToCap.kind === 'spell') return spellHitLine(hitToCap, talents);
  const split = splitText(
    baselineParts(
      hitToCap.baseline,
      inScope(talents, MELEE_BASELINE),
      'hit',
      HIT_GEAR_SOURCES,
      HIT_ALL_SOURCES,
    ),
    meleeStillToFind(hitToCap),
  );
  const caps = meleeCaps(hitToCap);
  return capLine(
    caps.length === 1
      ? guidesCopy.hitToCap(caps[0].cap)
      : guidesCopy.hitToCapEach(caps.map(({ cap, label }) => guidesCopy.capLabelled(cap, label)).join(', ')),
    split,
    hitToCap.spell === undefined ? bisCopy.hitToCapTitle : bisCopy.hitToCapWithSpellsTitle,
  );
}

/** The guide's expertise line ("Expertise cap: 6.5% (dodge)"); the BiS rail keeps its own wording. */
export function guideExpertiseCapLine(
  expertiseToCap: BisExpertiseToCap | null | undefined,
  talents: readonly TalentHit[] = [],
): GuideHitCapLine | undefined {
  if (expertiseToCap === null || expertiseToCap === undefined) return undefined;
  if (!expertiseIsGained(expertiseToCap)) return undefined;
  const { baseline, dodge, parry } = expertiseToCap;
  return capLine(
    guidesCopy.expertiseCap(
      capText(baseline, dodge),
      parry === undefined ? undefined : capText(baseline, parry),
    ),
    splitText(
      baselineParts(
        baseline,
        inScope(talents, EXPERTISE_BASELINE),
        'expertise',
        EXPERTISE_GEAR_SOURCES,
        EXPERTISE_ALL_SOURCES,
      ),
      undefined,
    ),
    bisCopy.expertiseToCapTitle,
  );
}
