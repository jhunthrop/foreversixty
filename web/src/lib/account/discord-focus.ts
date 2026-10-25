// web/src/lib/account/discord-focus.ts
// Where the Discord block should put keyboard focus once the page has re-read `/v1/me`. A change
// reloads the whole account page (it swaps to its loading state and back), which remounts the
// block, so the request is kept here rather than in the block's own state.
let pending: string | null = null;

export const discordFocus = {
  /** Remember the selector to focus (or clear it with null). */
  set(selector: string | null): void {
    pending = selector;
  },
  /** The selector still waiting to be focused, if any. */
  get(): string | null {
    return pending;
  },
};
