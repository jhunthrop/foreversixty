<!-- web/src/components/ui/Page.svelte -->
<!-- The page composition (web layout spec 2026-11-05 §3.7, §3A.2): the header, the control bar,
     then the body in one of four layouts inside Base.astro's one .page-column (this never declares
     a container of its own). `ledger` (BiS, tiers, rankings): one column. `split` (sim, guild,
     reports): main and a 340 px rail from 1024, --rail-w from 1760. `reading` (guides): a 200 px
     contents rail, the 720 px measure and a lane from 1760. `overview` (home): a 12-column grid.
     Sections in main and rail are --section-gap apart; nothing here draws a surface. -->
<script lang="ts">
  import type { Snippet } from 'svelte';
  import type { PageLayout } from '../../lib/ui/primitives';

  let {
    layout = 'ledger',
    header,
    bar,
    toc,
    rail,
    stickyRail = false,
    testid,
    children,
  }: {
    layout?: PageLayout;
    /** The PageHeader. */
    header?: Snippet;
    /** The ControlBar (full-bleed, sticky under the site header). */
    bar?: Snippet;
    /** `reading` only: the contents rail. */
    toc?: Snippet;
    /** `split` and `reading`: the rail (the lane on a reading page). */
    rail?: Snippet;
    /** The rail sticks below the bar from 1024. */
    stickyRail?: boolean;
    testid?: string;
    children: Snippet;
  } = $props();
</script>

{#if header}{@render header()}{/if}
{#if bar}{@render bar()}{/if}
<div class="ui-page" data-layout={layout} data-testid={testid}>
  {#if layout === 'split' || layout === 'reading'}
    {#if layout === 'reading' && toc}<nav class="ui-page-toc" aria-label="Contents">{@render toc()}</nav>{/if}
    <div class="ui-page-main">{@render children()}</div>
    {#if rail}<aside class="ui-page-rail" data-sticky={stickyRail ? 'true' : undefined}>
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
  .ui-page-toc {
    display: none;
  }
  .ui-page[data-layout='overview'] > .ui-page-main {
    display: grid;
    grid-template-columns: repeat(12, minmax(0, 1fr));
    gap: var(--section-gap) var(--gap);
  }
  .ui-page[data-layout='overview'] > .ui-page-main > :global(*) {
    grid-column: 1 / -1;
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
    .ui-page[data-layout='reading'] {
      grid-template-columns: minmax(0, var(--measure-reading)) minmax(0, 1fr);
    }
    .ui-page[data-layout='overview'] > .ui-page-main > :global([data-span='8']) {
      grid-column: span 8;
    }
    .ui-page[data-layout='overview'] > .ui-page-main > :global([data-span='4']) {
      grid-column: span 4;
    }
  }
  @media (min-width: 1760px) {
    .ui-page[data-layout='split'] {
      grid-template-columns: minmax(0, 1fr) var(--rail-w);
    }
    .ui-page[data-layout='reading'] {
      grid-template-columns: var(--toc-w) minmax(0, var(--measure-reading)) minmax(440px, 640px);
    }
    .ui-page-toc {
      display: block;
      position: sticky;
      top: calc(var(--bar-h) + 20px);
    }
  }
</style>
