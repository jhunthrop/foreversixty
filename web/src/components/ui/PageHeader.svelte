<!-- web/src/components/ui/PageHeader.svelte -->
<!-- The compact title band, once per page (round 5, design/specs/2026-11-09-directions.md: "no big
     headers"): the painted blue band continued from the floating bar above it (nav.css .site-top),
     and in it a crest, the Title in Marcellus 26/30 (21/25 below 768) in cream or the class, spec or
     faction colour, one meta line of facts and Tags, an optional Lead sentence, and at the right the
     page's answer (a Figure and its unit) or its actions. A TabStrip may sit under it on the band.
     `art="hero"` (home only) puts the owner's camp-road art behind it under the night veil instead;
     Base.astro's `art="hero"` starts the same art under the bar. No eyebrow, no job paragraph. -->
<script lang="ts">
  import type { Snippet } from 'svelte';
  import { REGION } from '../../lib/ui/regions';
  import { titleColor, type TitleTone } from '../../lib/ui/primitives';
  import type { BandArt } from '../../lib/ui/primitives';

  let {
    title,
    tone = 'strong',
    crest,
    meta,
    lead,
    answer,
    actions,
    tabs,
    minRowHeight,
    titleTestid = 'page-title',
    testid,
    art = 'band',
  }: {
    title: string;
    tone?: TitleTone;
    crest?: Snippet;
    /** The one fact line: strings and Tags, separated by space. */
    meta?: Snippet;
    /** The page's one sentence under the meta line (the Lead type, 17/26, max 62ch). */
    lead?: string;
    /** The page's answer: a Figure with its label and links. */
    answer?: Snippet;
    /** The page's buttons, when it has no answer. */
    actions?: Snippet;
    tabs?: Snippet;
    /** A CSS length: the row's floor from 1024 px, so a page whose answer changes height
     *  (one spec's figure, another's four) keeps its title and controls at the same y. The
     *  title block sits at the bottom of that area. */
    minRowHeight?: string;
    /** The title's testid, for a page whose tests name its heading ("report-title"). */
    titleTestid?: string;
    testid?: string;
    /** The band's art: the painted blue band (every page), or home's camp-road hero. */
    art?: BandArt;
    /** Never rendered: an Astro caller's whitespace between named slots arrives as children. */
    children?: Snippet;
  } = $props();
</script>

<div class="ui-band ui-band-end full-bleed" data-art={art} style:--band-y="var(--site-top-h, 80px)">
  <div class="page-column">
    <header class="ui-page-header" data-region={REGION.head} data-testid={testid}>
      <div class="ui-page-header-row" style:--row-min={minRowHeight}>
        <div class="flex min-w-0 items-center gap-4" data-region={REGION.headTitle}>
          {#if crest}{@render crest()}{/if}
          <div class="min-w-0">
            <h1 class="ui-title" style:color={titleColor(tone)} data-testid={titleTestid}>{title}</h1>
            {#if meta}<div
                data-region={REGION.headMeta}
                class="text-small text-muted mt-[6px] flex flex-wrap items-center gap-x-4 gap-y-[6px]"
              >
                {@render meta()}
              </div>{/if}
            {#if lead}<p class="ui-page-header-lead">{lead}</p>{/if}
          </div>
        </div>
        {#if answer}<div class="ui-page-header-answer" data-region={REGION.headAnswer}>
            {@render answer()}
          </div>
        {:else if actions}<div class="flex flex-wrap items-center gap-[10px]">{@render actions()}</div>{/if}
      </div>
      {#if tabs}<div class="ui-page-header-tabs">{@render tabs()}</div>{/if}
    </header>
  </div>
</div>

<style>
  /* Home's hero: the title sits at the foot of the art. */
  [data-art='hero'] .ui-page-header-row {
    min-height: 150px;
  }
  .ui-page-header {
    padding: 16px var(--page-inset) 14px;
  }
  .ui-page-header-tabs {
    margin-top: 12px;
    border-top: 1px solid rgb(255 255 255 / 0.14);
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
    letter-spacing: var(--tracking-heading);
    overflow-wrap: anywhere;
    margin: 0;
  }
  .ui-page-header-lead {
    margin: 10px 0 0;
    max-width: 62ch;
    font-size: 16px;
    line-height: 24px;
    color: var(--color-text);
  }
  .ui-page-header-answer {
    display: flex;
    flex-direction: column;
    gap: 4px;
  }
  @media (min-width: 768px) {
    .ui-page-header-lead {
      font-size: 17px;
      line-height: 26px;
    }
    .ui-title {
      font-size: var(--text-title);
      line-height: var(--text-title--line-height);
    }
  }
  @media (min-width: 1024px) {
    .ui-page-header-row {
      min-height: var(--row-min, auto);
    }
  }
  @media (min-width: 900px) {
    .ui-page-header-row {
      grid-template-columns: minmax(0, 1fr) auto;
    }
    .ui-page-header-answer {
      align-items: flex-end;
      text-align: right;
    }
  }
</style>
