// web/src/lib/guides/hit-cap-line.ts
// The Stat priority block's hit line, worded the way a guide gives it: "Hit cap: 16% (13% from
// school)". The numbers are the same ones the BiS page's rail prints (`hitCapLine`), so the two
// pages cannot disagree; only the sentence differs.
import { bisCopy } from '../bis/copy';
import { formatCapPercent } from '../bis/hit-cap';
import type { BisHitToCap } from '../bis/types';
import { guidesCopy } from './copy';

export interface GuideHitCapLine {
  text: string;
  title: string;
}

export function guideHitCapLine(hitToCap: BisHitToCap | null | undefined): GuideHitCapLine | undefined {
  if (hitToCap === null || hitToCap === undefined) return undefined;
  if (hitToCap.kind === 'spell') {
    const school = hitToCap.school === undefined ? undefined : formatCapPercent(hitToCap.school.to_cap);
    return {
      text: guidesCopy.hitCapLine(formatCapPercent(hitToCap.spell), school, undefined),
      title: bisCopy.spellHitToCapTitle,
    };
  }
  const white = hitToCap.white === undefined ? undefined : formatCapPercent(hitToCap.white);
  return {
    text: guidesCopy.hitCapLine(formatCapPercent(hitToCap.specials), undefined, white),
    title: bisCopy.hitToCapTitle,
  };
}
