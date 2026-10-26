// web/src/lib/guild/focus-now.ts
// A Svelte action that focuses its element when it mounts: race and refusal lines (raids spec 4
// "Races") and the cancel confirm's Keep it take focus the moment they appear.

export function focusOnMount(node: HTMLElement): void {
  node.focus();
}

/** Selects a textarea's text when it mounts: the copy fallback, so the officer can press copy. */
export function selectOnMount(node: HTMLTextAreaElement): void {
  node.focus();
  node.select();
}
