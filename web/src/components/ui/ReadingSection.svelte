<!-- web/src/components/ui/ReadingSection.svelte -->
<!-- A Section on a reading page (guides spec 2026-11-08 §3A): the head at 22/28 (20/26 below 768)
     over a --color-line-soft hairline, then the prose at the reading measure and its lane (the
     section's data: lists, tables, trees) beside it from 1280 (`720 | minmax(0, 640)`). Below
     1280 one column, the lane first (its DOM order: data before the why, G4). A section with no
     lane stays at the measure. Nothing here draws a surface. From 1280 a lane that runs well past
     the prose wraps its trailing rows into `.ui-rd-more` under the prose (lib/ui/reading-wrap.ts,
     mounted by the page), so the left column never stands empty beside a tall lane. -->
<script lang="ts">
  import type { Snippet } from 'svelte';
  import { sectionRegion } from '../../lib/ui/regions';

  let {
    title,
    anchor,
    lane,
    proseFirst = false,
    testid,
    children,
  }: {
    title: string;
    /** The section's link target (`#rotation-and-priority`). */
    anchor?: string;
    /** The section's data, beside the prose. */
    lane?: Snippet;
    /** Below 1280 the prose comes before the lane (a section whose lane is a long list the
     *  prose introduces: the guide's Best in slot). */
    proseFirst?: boolean;
    testid?: string;
    /** The prose. */
    children: Snippet;
  } = $props();
</script>

<section
  id={anchor}
  class="ui-rsec min-w-0 scroll-mt-[calc(var(--bar-h)+12px)]"
  data-region={sectionRegion(title)}
  data-testid={testid}
>
  <div class="border-line-soft border-b pb-[10px]">
    <h2 class="ui-rsec-title">{title}</h2>
  </div>
  <div class="ui-rd" data-solo={lane ? undefined : ''} data-prose-first={proseFirst ? '' : undefined}>
    {#if lane}<div class="ui-rd-lane">{@render lane()}</div>{/if}
    <div class="ui-rd-prose">
      {@render children()}
      {#if lane}<div class="ui-rd-more"></div>{/if}
    </div>
  </div>
</section>

<style>
  .ui-rsec-title {
    margin: 0;
    font-weight: 700;
    font-size: 20px;
    line-height: 26px;
    color: var(--color-strong);
  }
  .ui-rd {
    display: grid;
    grid-template-columns: minmax(0, var(--measure-reading));
    grid-template-areas: 'lane' 'prose';
    gap: 24px var(--gap);
    padding-top: 20px;
  }
  .ui-rd[data-solo] {
    grid-template-areas: 'prose';
  }
  .ui-rd-prose {
    grid-area: prose;
    min-width: 0;
  }
  .ui-rd-lane {
    grid-area: lane;
    min-width: 0;
    display: flex;
    flex-direction: column;
    gap: 28px;
  }
  /* The lane's rows wrapped under the prose (lib/ui/reading-wrap.ts): columns of a lane's width. */
  .ui-rd-more {
    display: none;
  }
  .ui-rd-more:not(:empty) {
    display: block;
    margin-top: 24px;
    padding-top: 16px;
    border-top: 1px solid var(--color-line-soft);
    columns: 340px;
    column-gap: 28px;
  }
  .ui-rd-more :global([data-wrapped]) {
    break-inside: avoid;
  }
  @media (max-width: 1279.98px) {
    .ui-rd[data-prose-first] {
      grid-template-areas: 'prose' 'lane';
    }
  }
  @media (min-width: 768px) {
    .ui-rsec-title {
      font-size: 22px;
      line-height: 28px;
    }
  }
  @media (min-width: 1280px) {
    .ui-rd {
      grid-template-columns: var(--measure-reading) minmax(0, 640px);
      grid-template-areas: 'prose lane';
      align-items: start;
    }
    .ui-rd[data-lane-sticky] > .ui-rd-lane {
      position: sticky;
      top: calc(var(--bar-h) + 12px);
    }
    .ui-rd[data-solo] {
      grid-template-columns: var(--measure-reading);
    }
  }
</style>
