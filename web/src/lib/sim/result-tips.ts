// web/src/lib/sim/result-tips.ts
// The saved result's tooltip keys and how each resolves: the page's spells from the class bundle
// and its one explanatory note (build ruling R7's Crit header) through the same shared panel.
import { talentTipModel, type GameTipModel, type TalentTipEntry } from '../tooltips/game-tip';
import { damageCopy } from './result-copy';
import type { SpellInfo } from './spell-info';

/** The Crit column header's note. */
export const CRIT_NOTE_TIP = 'note:crit';

const CRIT_NOTE: GameTipModel = { name: damageCopy.headers.crit, lines: [], description: damageCopy.critTip };

const SPELL_TIP = /^spell:(\d+)$/;
const TALENT_TIP = /^talent:(\d+):(\d+)$/;

let talentTips: ReadonlyMap<string, TalentTipEntry> = new Map();

/** The Talents card's own chunk hands over its trees' text once it has read them: the page's host
 *  loads with the first paint and cannot know them sooner. */
export function setResultTalentTips(entries: ReadonlyMap<string, TalentTipEntry>): void {
  talentTips = entries;
}

/** The model for a `data-game-tip` key on this page; null when the build carries nothing for it. */
export function resultTip(key: string, spells: SpellInfo): GameTipModel | null {
  if (key === CRIT_NOTE_TIP) return CRIT_NOTE;
  const talent = TALENT_TIP.exec(key);
  if (talent !== null) {
    const entry = talentTips.get(talent[1]!);
    return entry === undefined ? null : talentTipModel(entry, Number(talent[2]));
  }
  const match = SPELL_TIP.exec(key);
  return match === null ? null : (spells.fact(Number(match[1]))?.tip ?? null);
}
