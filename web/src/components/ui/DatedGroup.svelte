<!-- web/src/components/ui/DatedGroup.svelte -->
<!-- A reading section of dated entries (pages spec 2026-11-08 §3A, a month of the changelog): the
     reading head (22/28, 20/26 below 768) over a --color-line-soft hairline with its count at the
     right, then its Dated entries. Region `section-<slug of its title>`; no box. -->
<script lang="ts">
  import type { Snippet } from 'svelte';
  import { sectionRegion } from '../../lib/ui/regions';

  let {
    title,
    anchor,
    meta,
    metaContent,
    attrs = {},
    testid,
    children,
  }: {
    title: string;
    /** The section's link target (`#2026-11`). */
    anchor: string;
    /** The count at the right of the head (`3 changes`). */
    meta?: string;
    /** Replaces `meta` (a count per filter, one shown by CSS). */
    metaContent?: Snippet;
    /** Data attributes a page filters the group by (`data-game="1"`). */
    attrs?: Readonly<Record<`data-${string}`, string>>;
    testid?: string;
    /** The Dated entries. */
    children: Snippet;
  } = $props();
</script>

<section
  id={anchor}
  class="ui-dgroup min-w-0 scroll-mt-[calc(var(--bar-h)+12px)]"
  data-region={sectionRegion(title)}
  data-testid={testid}
  {...attrs}
>
  <div class="ui-dgroup-h border-line-soft border-b pb-[10px]">
    <h2 class="ui-dgroup-title">{title}</h2>
    {#if metaContent}<span class="text-small text-muted ml-auto">{@render metaContent()}</span>
    {:else if meta}<span class="text-small text-muted ml-auto">{meta}</span>{/if}
  </div>
  <div class="ui-dgroup-list">{@render children()}</div>
</section>

<style>
  .ui-dgroup-h {
    display: flex;
    align-items: baseline;
    gap: 12px;
  }
  .ui-dgroup-title {
    margin: 0;
    font-weight: 700;
    font-size: 20px;
    line-height: 26px;
    color: var(--color-strong);
  }
  .ui-dgroup-list {
    display: flex;
    flex-direction: column;
  }
  @media (min-width: 768px) {
    .ui-dgroup-title {
      font-size: 22px;
      line-height: 28px;
    }
  }
</style>
