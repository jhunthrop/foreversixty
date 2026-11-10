<!-- web/src/components/guild/RaidDisclosure.svelte -->
<!-- A section closed by default (attendance spec R-A4): on a `done` raid, Notes and the Sign-ups
     roster sit below Attendance as native disclosures. The summary is a 44 px row holding the h2,
     an optional mono count and a 16 px chevron that turns when the section opens. -->
<script lang="ts">
  import Icon from '../ui/Icon.svelte';
  import { ICON_CHEVRON_DOWN } from '../../lib/ui/ui-icons';
  import type { Snippet } from 'svelte';

  let {
    heading,
    headingId,
    count = null,
    headingClass = 'text-strong text-[18px] leading-6 font-bold',
    testid,
    children,
  }: {
    heading: string;
    headingId: string;
    /** `14/40` beside the Sign-ups heading. */
    count?: string | null;
    /** The heading's look; a section title unless the disclosure is one quiet line. */
    headingClass?: string;
    testid: string;
    children: Snippet;
  } = $props();
</script>

<details class="group flex flex-col" data-testid={testid}>
  <summary
    class="focus-visible:outline-gold flex min-h-11 cursor-pointer list-none items-center gap-3 focus-visible:outline-2 focus-visible:outline-offset-2 [&::-webkit-details-marker]:hidden"
    data-testid={`${testid}-summary`}
  >
    <h2 id={headingId} class={headingClass}>{heading}</h2>
    {#if count !== null}<span class="text-muted numeric text-[13px]">{count}</span>{/if}
    <Icon icon={ICON_CHEVRON_DOWN} size={16} class="text-muted transition-transform group-open:rotate-180" />
  </summary>
  <div class="pt-3">{@render children()}</div>
</details>
