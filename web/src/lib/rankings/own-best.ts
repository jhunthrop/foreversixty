// web/src/lib/rankings/own-best.ts
// The Your rank panel's facts (directions mock b1-rankings-board): the viewer's best kill on this
// board, the value, Executed, the length and the day, read from the row the board placed. None
// when the board did not place the viewer.
import { formatDate } from '../dates';
import { formatAmount } from '../report/format';
import { executionLabel } from '../sim/execution';
import type { Fact } from '../ui/facts';
import type { RankingsPage } from './api';
import { killLength } from './rank-row-view';

export function ownBestFacts(page: Pick<RankingsPage, 'character_rank'>, board: { unit: string }): Fact[] {
  const row = page.character_rank?.row;
  if (row === undefined) return [];
  return [
    { id: 'value', label: board.unit || 'Value', value: formatAmount(Math.round(row.value)) },
    { id: 'executed', label: 'Executed', value: executionLabel(row.execution_score ?? null) },
    { id: 'length', label: 'Length', value: killLength(row.duration_ms) },
    { id: 'killed', label: 'Killed', value: formatDate(new Date(row.fought_at)) },
  ];
}
