<!-- web/src/components/ui/Page.svelte -->
<!-- The page composition (web layout spec 2026-11-05 §3.7, §3A.2): the header, the control bar,
     then the body in one of four layouts inside Base.astro's one .page-column (this never declares
     a container of its own). `ledger` (BiS, tiers, rankings): one column. `split` (sim, guild,
     reports): main and a 340 px rail from 1024, --rail-w from 1760. `reading` (guides): a full-width
     `top` (the plate), then the ReadingSections, a 200 px contents rail left of them from 1760. `overview` (home): a 12-column grid.
     `tocFrom` (pages spec 2026-11-08 §3A): a document with no lanes shows its rail from 1024 (the
     legal pages, Data sources), or never (`0`: About, one block of prose); `tocLabel` names it.
     `form` (/connect; setup spec 2026-11-07 section 3A): one 560 px column at the page column's left.
     `workbench` (planner; planner spec 2026-11-08 section 3A): main and a 340 px rail from 1280
     (--rail-w from 1760); below 1280 main and rail give their regions to one flow ordered by each
     ui/Ordered's `order`, two columns from 1024 where an Ordered `span` takes both.
     Sections in main and rail are --section-gap apart; nothing here draws a surface. -->
<script lang="ts">
  import type { Snippet } from 'svelte';
  import type { PageLayout } from '../../lib/ui/primitives';

  let {
    layout = 'ledger',
    header,
    bar,
    toc,
    top,
    rail,
    stickyRail = false,
    panels = false,
    railFirst = false,
    splitFrom = 1024,
    tocFrom,
    tocLabel = 'Contents',
    testid,
    children,
  }: {
    layout?: PageLayout;
    /** The PageHeader. */
    header?: Snippet;
    /** The ControlBar (full-bleed, sticky under the site header). */
    bar?: Snippet;
    /** `reading` only: the contents rail (from 1760, left of the sections). */
    toc?: Snippet;
    /** `reading` only: the full-width block above the sections and the contents rail (the plate). */
    top?: Snippet;
    /** `split`, `reading` and `workbench`: the rail (the lane on a reading page). */
    rail?: Snippet;
    /** The rail sticks below the bar from 1024. */
    stickyRail?: boolean;
    /** Main and rail hold Section `panel`s (the logs and report mocks): 22 px apart, the columns 28. */
    panels?: boolean;
    /** `split` only: below 1024, where main and rail stack, the rail comes first (the tier list's
     *  spec panel above the list on a phone). */
    railFirst?: boolean;
    /** `split` only: the width main and rail sit side by side from; 1280 for a page whose main
     *  column needs the room below it (the tier list), stacked and full width under it. */
    splitFrom?: 1024 | 1280;
    /** `reading` only, a document (the changelog, the content pages): the width the contents rail
     *  shows from (1760, 1024), or 0 for never, with the document's section gap above it. The
     *  guides leave it out (the rail from 1760 under their plate). */
    tocFrom?: 1760 | 1024 | 0;
    /** `reading` only: the contents rail's accessible name. */
    tocLabel?: string;
    testid?: string;
    children: Snippet;
  } = $props();
</script>

{#if header}{@render header()}{/if}
{#if bar}{@render bar()}{/if}
<div
  class="ui-page"
  data-layout={layout}
  data-panels={panels ? 'true' : undefined}
  data-split-from={layout === 'split' && splitFrom !== 1024 ? splitFrom : undefined}
  data-toc-from={layout === 'reading' ? tocFrom : undefined}
  data-testid={testid}
>
  {#if layout === 'reading'}
    {#if top}<div class="ui-page-top">{@render top()}</div>{/if}
    {#if toc && tocFrom !== 0}<nav class="ui-page-toc" aria-label={tocLabel}>{@render toc()}</nav>{/if}
    <div class="ui-page-main">{@render children()}</div>
  {:else if layout === 'split' || layout === 'workbench'}
    <div class="ui-page-main">{@render children()}</div>
    {#if rail}<aside
        class="ui-page-rail"
        data-sticky={stickyRail ? 'true' : undefined}
        data-first={railFirst && layout === 'split' ? 'true' : undefined}
      >
        {@render rail()}
      </aside>{/if}
  {:else}
    <div class="ui-page-main">{@render children()}</div>
  {/if}
</div>

<style>
  .ui-page {
    display: grid;
    gap: var(--gap);
    grid-template-columns: minmax(0, 1fr);
    align-items: start;
    padding-top: 20px;
    padding-inline: var(--page-inset);
  }
  .ui-page-main,
  .ui-page-rail {
    display: flex;
    flex-direction: column;
    gap: var(--section-gap);
    min-width: 0;
  }
  .ui-page-rail {
    gap: 32px;
  }
  .ui-page[data-panels] {
    --section-gap: 22px;
    column-gap: 28px;
    padding-top: 28px;
  }
  .ui-page[data-panels] > .ui-page-rail {
    gap: var(--section-gap);
  }
  .ui-page-toc {
    display: none;
  }
  @media (max-width: 1023px) {
    .ui-page-rail[data-first='true'] {
      order: -1;
    }
  }
  @media (min-width: 1024px) and (max-width: 1279px) {
    .ui-page[data-layout='split'][data-split-from='1280'] {
      grid-template-columns: minmax(0, 1fr);
    }
    .ui-page[data-split-from='1280'] > .ui-page-rail[data-first='true'] {
      order: -1;
    }
  }
  /* L2 Reading (guides spec 2026-11-08 §3A): the plate full width, then the sections, the
     contents rail beside them from 1760. */
  .ui-page[data-layout='reading'] {
    row-gap: var(--section-gap);
  }
  .ui-page-top {
    grid-column: 1 / -1;
    min-width: 0;
  }
  .ui-page[data-layout='form'] > .ui-page-main {
    /* The form measure (spec 3A, the proposed --measure-form at 560): the full width on a phone. */
    width: 100%;
    max-width: 560px;
  }
  .ui-page[data-layout='overview'] > .ui-page-main {
    display: grid;
    grid-template-columns: repeat(12, minmax(0, 1fr));
    gap: var(--section-gap) var(--gap);
  }
  .ui-page[data-layout='overview'] > .ui-page-main > :global(*) {
    grid-column: 1 / -1;
  }
  /* L3 Workbench (planner spec 2026-11-08 section 3A): one flow below 1280, ordered by --o. */
  .ui-page[data-layout='workbench'] {
    gap: var(--section-gap) var(--gap);
  }
  .ui-page[data-layout='workbench'] > .ui-page-main,
  .ui-page[data-layout='workbench'] > .ui-page-rail {
    display: contents;
  }
  .ui-page[data-layout='workbench'] :global([data-order]) {
    order: var(--o, 0);
    min-width: 0;
  }
  @media (min-width: 1024px) and (max-width: 1279px) {
    .ui-page[data-layout='workbench'] {
      grid-template-columns: repeat(2, minmax(0, 1fr));
    }
    .ui-page[data-layout='workbench'] :global([data-span]) {
      grid-column: 1 / -1;
    }
  }
  @media (min-width: 1280px) {
    .ui-page[data-layout='workbench'] {
      grid-template-columns: minmax(0, 1fr) 340px;
    }
    .ui-page[data-layout='workbench'] > .ui-page-main,
    .ui-page[data-layout='workbench'] > .ui-page-rail {
      display: flex;
    }
    .ui-page[data-layout='workbench'] > .ui-page-rail {
      gap: 40px;
    }
  }
  @media (min-width: 1024px) {
    .ui-page[data-layout='split'] {
      grid-template-columns: minmax(0, 1fr) 340px;
      padding-top: 28px;
    }
    .ui-page-rail[data-sticky='true'] {
      position: sticky;
      top: calc(var(--bar-h) + 20px);
    }
    .ui-page[data-layout='overview'] > .ui-page-main > :global([data-span='8']) {
      grid-column: span 8;
    }
    .ui-page[data-layout='overview'] > .ui-page-main > :global([data-span='4']) {
      grid-column: span 4;
    }
  }
  /* A document (pages spec 2026-11-08 §3A): the rail at the left from 1024, sticky 24 under the
     top; with no rail (`0`) one column at every width. */
  .ui-page[data-toc-from] {
    padding-top: var(--section-gap);
  }
  /* A document's section head: its hairline stops at the reading measure (no rule over empty
     space), and an arrival at its anchor lands 24 px under the top (no bar on a document). */
  .ui-page:is([data-toc-from='1024'], [data-toc-from='0']) :global(.ui-rsec) {
    scroll-margin-top: 24px;
  }
  .ui-page:is([data-toc-from='1024'], [data-toc-from='0']) :global(.ui-rsec > :first-child) {
    max-width: var(--measure-reading);
  }
  @media (min-width: 1024px) {
    .ui-page[data-toc-from='1024'] {
      grid-template-columns: var(--toc-w) minmax(0, 1fr);
    }
    .ui-page[data-toc-from='1024'] > .ui-page-toc {
      display: block;
      position: sticky;
      top: 24px;
    }
  }
  @media (min-width: 1760px) {
    .ui-page[data-toc-from='0'] {
      grid-template-columns: minmax(0, 1fr);
    }
  }
  @media (min-width: 1760px) {
    .ui-page[data-layout='split'],
    .ui-page[data-layout='workbench'] {
      grid-template-columns: minmax(0, 1fr) var(--rail-w);
    }
    .ui-page[data-layout='reading'] {
      grid-template-columns: var(--toc-w) minmax(0, 1fr);
    }
    .ui-page-toc {
      display: block;
      position: sticky;
      top: calc(var(--bar-h) + 20px);
    }
  }
</style>
