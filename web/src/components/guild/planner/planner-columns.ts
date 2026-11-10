// web/src/components/guild/planner/planner-columns.ts
// The planner's figure columns (item level, ready, rating, 30 nights), one track list shared by the
// roster's head line and every sign-up row so the figures sit under their heads.
export const PLANNER_FACT_COLUMNS =
  'grid-cols-[72px_64px_48px_64px] items-center gap-x-3 text-right tabular-nums';

export const PLANNER_FACT_HEADS: readonly string[] = ['Item level', 'Ready', 'Rating', '30 nights'];
