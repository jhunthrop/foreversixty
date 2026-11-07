<!-- web/src/components/ui/Disclosure.svelte -->
<!-- A summary line that opens its body in place (web layout spec 2026-11-05 §3A.2, the character
     page's parts, history and coach history): a native details element, so find-in-page and the
     keyboard work without script; a 44 px summary in the body face 600 with a caret, no box. The
     body is rendered only once opened, so a lazy section (charts, a list that reads the API) costs
     nothing until the reader asks for it; `eager` renders it from the start (a static page with no
     Svelte runtime, where the native element opens it: the content pages' phone contents). -->
<script lang="ts">
  import type { Snippet } from 'svelte';

  let {
    summary,
    meta,
    open = $bindable(false),
    eager = false,
    testid,
    summaryTestid,
    children,
  }: {
    summary: string;
    /** Muted words after the summary. */
    meta?: string;
    open?: boolean;
    /** Render the body before it is opened (static markup, opened by the browser). */
    eager?: boolean;
    testid?: string;
    /** The summary's own testid (a test that presses the trigger and reads its words). */
    summaryTestid?: string;
    children: Snippet;
  } = $props();
</script>

<details class="ui-disclosure" bind:open data-testid={testid}>
  <summary
    data-testid={summaryTestid}
    class="text-nav focus-visible:outline-gold flex min-h-11 cursor-pointer list-none items-center gap-2 text-[14px] leading-5 font-semibold select-none focus-visible:outline-2 focus-visible:-outline-offset-2 [&::-webkit-details-marker]:hidden"
  >
    {summary}{#if meta}<span class="text-small text-muted ml-1 font-normal">{meta}</span>{/if}
  </summary>
  {#if open || eager}<div class="pt-1 pb-2">{@render children()}</div>{/if}
</details>

<style>
  /* The caret is drawn, not written (the mock library's Disclosure): no text node to read. */
  .ui-disclosure > summary::before {
    content: '\25B8';
    display: inline-block;
    font-size: 11px;
    color: var(--color-muted);
  }
  .ui-disclosure[open] > summary::before {
    transform: rotate(90deg);
  }
</style>
