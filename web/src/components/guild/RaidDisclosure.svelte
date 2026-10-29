<!-- web/src/components/guild/RaidDisclosure.svelte -->
<!-- A section closed by default (attendance spec R-A4): on a `done` raid, Notes and the Sign-ups
     roster sit below Attendance as native disclosures. The summary is a 44 px row holding the h2,
     an optional mono count and a 16 px chevron that turns when the section opens. -->
<script lang="ts">
  import type { Snippet } from 'svelte';

  let {
    heading,
    headingId,
    count = null,
    testid,
    children,
  }: {
    heading: string;
    headingId: string;
    /** `14/40` beside the Sign-ups heading. */
    count?: string | null;
    testid: string;
    children: Snippet;
  } = $props();
</script>

<details class="group flex flex-col" data-testid={testid}>
  <summary
    class="focus-visible:outline-gold flex min-h-11 cursor-pointer list-none items-center gap-3 focus-visible:outline-2 focus-visible:outline-offset-2 [&::-webkit-details-marker]:hidden"
    data-testid={`${testid}-summary`}
  >
    <h2 id={headingId} class="section-title">{heading}</h2>
    {#if count !== null}<span class="text-muted numeric text-[13px]">{count}</span>{/if}
    <svg
      class="text-muted size-4 shrink-0 transition-transform group-open:rotate-180"
      viewBox="0 0 16 16"
      aria-hidden="true"><path d="M4 6l4 4 4-4" fill="none" stroke="currentColor" stroke-width="1.5" /></svg
    >
  </summary>
  <div class="pt-3">{@render children()}</div>
</details>
