<!-- web/src/components/ui/Spinner.svelte -->
<!-- The 14 px ring a busy button shows in place of its glyph (polish spec 4.2, Loading). Still
     under reduced motion: the open ring alone says "working". -->
<svg
  class="shrink-0 animate-spin motion-reduce:animate-none"
  width="14"
  height="14"
  viewBox="0 0 24 24"
  fill="none"
  stroke="currentColor"
  stroke-width="3"
  stroke-linecap="round"
  aria-hidden="true"
  data-testid="spinner"><path d="M21 12a9 9 0 11-9-9" /></svg
>
