<!-- web/src/components/ui/PageHead.svelte -->
<!-- The page head (site alignment spec 2026-10-29 §3.1, S2), lifted from /account's head row:
     the gold-ruled eyebrow over the page's one h1, an optional job line and meta line, an
     optional 64 px crest at the left (`lead`), the right slot (`aside`: the identity block or the
     page's one control, bottom-aligned from 1024, stacked under the title below it), and a
     `below` row for tabs or controls. It sits on the page ground: no band, never a card (D1).
     PageHead.astro is the static twin; both read lib/ui/page-head.ts. -->
<script lang="ts">
  import type { Snippet } from 'svelte';
  import { PAGE_HEAD, pageHeadTitleClass } from '../../lib/ui/page-head';

  let {
    eyebrow,
    title,
    titleClass,
    titleColor,
    titleId,
    job,
    lead,
    meta,
    aside,
    below,
    testid,
  }: {
    /** The page's section name, drawn uppercase (ACCOUNT, GUIDES). */
    eyebrow: string;
    title: string;
    /** Extra classes on the h1. */
    titleClass?: string;
    /** A class colour for the h1 (a CSS colour or var()). */
    titleColor?: string;
    titleId?: string;
    /** One line under the h1 that says what the page does. */
    job?: string;
    /** The 64 px crest at the title's left. */
    lead?: Snippet;
    /** The 13 px muted line under the job line (the Updated stamp). */
    meta?: Snippet;
    /** The right slot: the identity block or the page's one control. */
    aside?: Snippet;
    /** Tabs or a controls row, 16 px under the head. */
    below?: Snippet;
    testid?: string;
  } = $props();
</script>

{#snippet titleBlock()}
  <div class={PAGE_HEAD.titleBlock}>
    <p class={PAGE_HEAD.eyebrow}><span class={PAGE_HEAD.eyebrowRule} aria-hidden="true"></span>{eyebrow}</p>
    <h1 class={pageHeadTitleClass(titleClass)} id={titleId} style:color={titleColor}>{title}</h1>
    {#if job !== undefined}<p class={PAGE_HEAD.job} data-testid="page-head-job">{job}</p>{/if}
    {#if meta !== undefined}<p class={PAGE_HEAD.meta} data-testid="page-head-meta">{@render meta()}</p>{/if}
  </div>
{/snippet}

<div class={PAGE_HEAD.root} data-testid={testid}>
  <div class={PAGE_HEAD.row}>
    {#if lead !== undefined}
      <div class={PAGE_HEAD.withLead}>{@render lead()}{@render titleBlock()}</div>
    {:else}
      {@render titleBlock()}
    {/if}
    {#if aside !== undefined}<div class={PAGE_HEAD.aside} data-testid="page-head-aside">
        {@render aside()}
      </div>{/if}
  </div>
  {@render below?.()}
</div>
