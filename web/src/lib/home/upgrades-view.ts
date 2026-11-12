// web/src/lib/home/upgrades-view.ts
// Your upgrades' own words (home units lane, 2026-11-11): each row's gain and the panel's foot in
// the band's own unit (DPS, HPS, a tank's score). A leaf of the home graph on purpose: it is the
// only home module that reads bis/tank-view, whose static graph carries the whole BiS copy table.
// home-view and home-facts are shared with the Character page and GuildShell's lazy week-one card
// (through home-b-view), so an import of tank-view there put two more chunk names in GuildShell's
// preload map and tripped its island budget (scripts/check-island-size.mjs). Only HomeUpgrades, in
// the home island that already loads tank-view through upgrades-loader, imports this module.
import { scoreUnitWord, type SlotScoreUnit } from '../bis/tank-view';
import type { GainFigure } from '../ui/primitives';
import { figureText } from './home-b-view';
import { factsCopy, type HomeFact } from './home-facts';
import { UPGRADES_SHOWN, gainText } from './home-view';
import type { SlotUpgrade, UpgradesResult } from './upgrades';

/** An upgrade row's gain at its right: the figure in the band's own unit, the word the paperdoll
 *  uses (`+5.1 DPS`, `+5.1 HPS`, a tank's `+5.1 score`); undefined with no measured gain. */
export function upgradeGain(
  upgrade: Pick<SlotUpgrade, 'gainDps'>,
  unit: SlotScoreUnit,
): GainFigure | undefined {
  if (upgrade.gainDps === null) return undefined;
  return { value: gainText(upgrade.gainDps), unit: scoreUnitWord(unit), tone: 'kill' };
}

/** Your upgrades' foot: `+52.1 DPS · 5 best in slot` (See all N beside it counts the upgrades); each
 *  clause only when it is above zero. The gap is every measured upgrade's, so with more upgrades
 *  than the panel's rows it says so (`+1,058.4 score across all 17`), never read as the rows' sum. */
export function upgradesFacts(result: UpgradesResult | undefined): HomeFact[] {
  if (result === undefined || result.upgrades.length === 0) return [];
  const facts: HomeFact[] = [];
  if (result.totalGainDps > 0) {
    const unit = scoreUnitWord(result.scoreUnit);
    const count = result.upgrades.length;
    facts.push({
      n: `+${figureText(result.totalGainDps)}`,
      text: count > UPGRADES_SHOWN ? factsCopy.acrossAll(unit, count) : unit,
    });
  }
  if (result.alreadyBis.length > 0) {
    facts.push({ n: String(result.alreadyBis.length), text: factsCopy.slotsBis });
  }
  return facts;
}
