<!-- web/src/components/ui/PageHeader.svelte -->
<!-- The page head, once per page (web layout spec 2026-11-05 §3.3, §3A.2, §5.3): a crest, the Title
     in Marcellus 40/46 (32/38 below 768) in the class, spec or faction colour (owner decision 8: the
     faction colour where a page has no class, the strong colour where it has neither), one meta line
     of facts and Tags, and at the right the page's answer (a Figure and its links) or its actions.
     A TabStrip may sit under it. No eyebrow, no job paragraph, no band art. -->
<script lang="ts">
  import type { Snippet } from 'svelte';
  import { titleColor, type TitleTone } from '../../lib/ui/primitives';

  let {
    title,
    tone = 'strong',
    crest,
    meta,
    answer,
    actions,
    tabs,
    testid,
  }: {
    title: string;
    tone?: TitleTone;
    crest?: Snippet;
    /** The one fact line: strings and Tags, separated by space. */
    meta?: Snippet;
    /** The page's answer: a Figure with its label and links. */
    answer?: Snippet;
    /** The page's buttons, when it has no answer. */
    actions?: Snippet;
    tabs?: Snippet;
    testid?: string;
  } = $props();
</script>

<header class="ui-page-header" data-testid={testid}>
  <div class="ui-page-header-row">
    <div class="flex min-w-0 items-center gap-4">
      {#if crest}{@render crest()}{/if}
      <div class="min-w-0">
        <h1 class="ui-title" style:color={titleColor(tone)} data-testid="page-title">{title}</h1>
        {#if meta}<div class="text-small text-muted mt-[6px] flex flex-wrap items-center gap-x-4 gap-y-[6px]">
            {@render meta()}
          </div>{/if}
      </div>
    </div>
    {#if answer}<div class="ui-page-header-answer">{@render answer()}</div>
    {:else if actions}<div class="flex flex-wrap items-center gap-[10px]">{@render actions()}</div>{/if}
  </div>
  {#if tabs}<div class="mt-2">{@render tabs()}</div>{/if}
</header>

<style>
  .ui-page-header {
    padding: 16px var(--page-inset) 14px;
  }
  .ui-page-header-row {
    display: grid;
    gap: 16px;
    grid-template-columns: minmax(0, 1fr);
    align-items: end;
  }
  .ui-title {
    font-family: var(--font-title);
    font-weight: 400;
    font-size: var(--text-title-phone);
    line-height: var(--text-title-phone--line-height);
    letter-spacing: 0.01em;
    overflow-wrap: anywhere;
    margin: 0;
  }
  .ui-page-header-answer {
    display: flex;
    flex-direction: column;
    gap: 4px;
  }
  @media (min-width: 768px) {
    .ui-title {
      font-size: var(--text-title);
      line-height: var(--text-title--line-height);
    }
  }
  @media (min-width: 900px) {
    .ui-page-header {
      padding-block: 20px 16px;
    }
    .ui-page-header-row {
      grid-template-columns: minmax(0, 1fr) auto;
    }
    .ui-page-header-answer {
      align-items: flex-end;
      text-align: right;
    }
  }
</style>
