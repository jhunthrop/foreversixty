// web/src/lib/held-submit.ts
// A lookup form's real submit handler (lib/character-search-dom.ts) is imported after first
// paint (lib/after-paint.ts). A visitor who types and presses Enter before that import lands
// would otherwise get the browser's native GET to the current page with the fields in the
// query string. This eager, import-free hold keeps the submit from leaving, remembers it, and
// replays it once the handler is wired, so the first Enter always lands on the profile.

/** Holds `form`'s submits until the returned release runs; release replays a held submit. */
export function holdSubmitUntilWired(form: HTMLFormElement): () => void {
  let held = false;
  let submitter: HTMLElement | null = null;
  const hold = (event: SubmitEvent): void => {
    event.preventDefault();
    held = true;
    submitter = event.submitter;
  };
  form.addEventListener('submit', hold);
  return () => {
    form.removeEventListener('submit', hold);
    if (held) form.requestSubmit(submitter ?? undefined);
  };
}
