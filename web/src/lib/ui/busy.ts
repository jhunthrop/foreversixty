// Retired (owner ruling 2026-11-10: a busy button shows the spinner plus a busy word, through
// ui/Button's `busy` and ui/ConfirmButtons). The dimmed look below survives only for lane D's
// guild/billing/GuildSeatRow.svelte, and goes when that file moves to Button busy.
export const BUSY_CLASS = 'opacity-60 cursor-progress';
