<!-- web/src/components/ui/Dated.svelte -->
<!-- One entry of a dated log (pages spec 2026-11-08 §3A): the date (a link to the entry, its
     permalink) over its Tags in a 120 px column, the entry (title 18/24, the body at the reading
     measure) and its lane (sources, an action, a caveat), side by side from 1280, stacked below.
     `folded` (ruling CL1, an entry before launch): a title row (date, title, tags) that opens in
     place onto the body and lane. A hairline between entries; no box. -->
<script lang="ts">
  import type { Snippet } from 'svelte';

  let {
    anchor,
    iso,
    date,
    title,
    folded = false,
    show,
    tags,
    lane,
    testid,
    children,
  }: {
    /** The entry's id: its permalink target. */
    anchor: string;
    /** The date as `2026-11-05`. */
    iso: string;
    /** The date as the page prints it (`5 Nov`). */
    date: string;
    title: string;
    folded?: boolean;
    /** The filter side it belongs to (`data-show`), for a page that filters by CSS. */
    show?: string;
    tags?: Snippet;
    lane?: Snippet;
    testid?: string;
    /** The body. */
    children: Snippet;
  } = $props();
</script>

{#if folded}
  <details class="ui-dated ui-dated--f" id={anchor} data-show={show} data-testid={testid}>
    <summary class="ui-dated-s">
      <time class="ui-dated-fd" datetime={iso}>{date}</time>
      <h3 class="ui-dated-ft">{title}</h3>
      {#if tags}<span class="ui-dated-k">{@render tags()}</span>{/if}
    </summary>
    <div class="ui-dated-o">
      <div class="ui-dated-b">{@render children()}</div>
      {#if lane}<div class="ui-dated-l">{@render lane()}</div>{/if}
    </div>
  </details>
{:else}
  <article class="ui-dated ui-dated--o" id={anchor} data-show={show} data-testid={testid}>
    <div class="ui-dated-w">
      <a class="ui-dated-d" href={`#${anchor}`}><time datetime={iso}>{date}</time></a>
      {#if tags}{@render tags()}{/if}
    </div>
    <div class="ui-dated-b">
      <h3 class="ui-dated-t">{title}</h3>
      {@render children()}
    </div>
    {#if lane}<div class="ui-dated-l">{@render lane()}</div>{/if}
  </article>
{/if}

<style>
  .ui-dated {
    border-bottom: 1px solid var(--color-line-soft);
    scroll-margin-top: 72px;
  }
  .ui-dated:last-child {
    border-bottom: 0;
  }
  .ui-dated--o {
    display: grid;
    grid-template-columns: minmax(0, 1fr);
    gap: 10px var(--gap);
    padding: 24px 0;
  }
  .ui-dated--o:first-child {
    padding-top: 20px;
  }
  .ui-dated-w {
    display: flex;
    flex-wrap: wrap;
    align-items: center;
    gap: 6px 10px;
  }
  .ui-dated-d {
    font-weight: 700;
    font-size: 15px;
    line-height: 20px;
    font-variant-numeric: tabular-nums;
    color: var(--color-strong);
  }
  .ui-dated-d:hover {
    color: var(--color-gold-hover);
  }
  .ui-dated-d:focus-visible,
  .ui-dated-s:focus-visible {
    outline: 2px solid var(--color-gold);
    outline-offset: 2px;
  }
  .ui-dated-b {
    min-width: 0;
    max-width: var(--measure-reading);
  }
  .ui-dated-t {
    margin: 0 0 8px;
    font-weight: 700;
    font-size: 18px;
    line-height: 24px;
    color: var(--color-strong);
  }
  .ui-dated-l {
    min-width: 0;
    display: flex;
    flex-direction: column;
    gap: 10px;
  }
  .ui-dated-s {
    list-style: none;
    cursor: pointer;
    display: grid;
    grid-template-columns: auto minmax(0, 1fr);
    grid-template-areas: 'd k' 'h h';
    align-items: center;
    gap: 4px 10px;
    min-height: 44px;
    padding: 14px 0;
  }
  .ui-dated--f:first-child > .ui-dated-s {
    padding-top: 20px;
  }
  .ui-dated-s::-webkit-details-marker {
    display: none;
  }
  .ui-dated-fd {
    grid-area: d;
    font-weight: 600;
    font-size: 14px;
    line-height: 20px;
    font-variant-numeric: tabular-nums;
    color: var(--color-muted);
  }
  .ui-dated-ft {
    grid-area: h;
    display: flex;
    align-items: baseline;
    gap: 8px;
    margin: 0;
    font-weight: 700;
    font-size: 16px;
    line-height: 22px;
    color: var(--color-strong);
  }
  .ui-dated-ft::before {
    content: '\25B8';
    flex: none;
    font-size: 11px;
    color: var(--color-muted);
  }
  .ui-dated--f[open] .ui-dated-ft::before {
    content: '\25BE';
  }
  .ui-dated-s:hover .ui-dated-ft {
    color: var(--color-gold-hover);
  }
  .ui-dated-k {
    grid-area: k;
    display: flex;
    gap: 6px;
  }
  .ui-dated-o {
    display: grid;
    grid-template-columns: minmax(0, 1fr);
    gap: 10px var(--gap);
    padding: 0 0 24px;
  }
  @media (min-width: 1280px) {
    .ui-dated--o,
    .ui-dated-s,
    .ui-dated-o {
      grid-template-columns: 120px var(--measure-reading) minmax(0, 1fr);
      column-gap: var(--gap);
    }
    .ui-dated--o .ui-dated-w {
      flex-direction: column;
      align-items: flex-start;
      gap: 8px;
    }
    .ui-dated--o .ui-dated-l {
      padding-top: 2px;
    }
    .ui-dated-s {
      grid-template-areas: 'd h k';
    }
    .ui-dated-o > .ui-dated-b {
      grid-column: 2;
    }
  }
  @media (prefers-reduced-motion: reduce) {
    .ui-dated-d {
      transition: none;
    }
  }
</style>
