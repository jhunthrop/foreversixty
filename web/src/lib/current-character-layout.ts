// web/src/lib/current-character-layout.ts
// The flex gap SimView.svelte's and ToolsView.svelte's own root `<div>` both use between
// their blocks. A static shell's pre-hydration markup reuses the identical value, so its
// blocks sit where the hydrated root puts them instead of behind a literal that can drift.
export const VIEW_GAP = 'gap-[22px] md:gap-8';
