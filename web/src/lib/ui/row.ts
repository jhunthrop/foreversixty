// web/src/lib/ui/row.ts
// The settings row's class strings (polish spec 4.2), shared by ui/SettingsRow.svelte and by the
// one row that must be a different element (the Discord reconnect link, which keeps its opening
// watchdog): 56 px minimum, 20 px side padding, 14 px gap, a 1 px top rule, a hover ground only
// when the row is interactive.
export const ROW =
  'border-line-soft flex min-h-14 w-full flex-wrap items-center gap-x-[14px] gap-y-1 border-t px-5 py-[10px] text-left';

export const ROW_INTERACTIVE =
  'cursor-pointer transition-colors duration-[120ms] hover:bg-row-hover active:bg-gold/10 focus-visible:outline-2 focus-visible:-outline-offset-2 focus-visible:outline-gold';

/** The label, sub and value classes of a row's text. */
export const ROW_LABEL = 'text-[15px] font-semibold [overflow-wrap:anywhere]';
export const ROW_SUB = 'text-muted text-[13px]';
