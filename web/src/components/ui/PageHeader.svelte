<!-- web/src/components/ui/PageHeader.svelte -->
<!-- The compact title band, once per page (round 5, design/specs/2026-11-09-directions.md: "no big
     headers"): the painted blue band continued from the floating bar above it (nav.css .site-top),
     and in it a crest, the Title in Marcellus 26/30 (21/25 below 768) in cream or the class, spec or
     faction colour, an optional Lead sentence directly under the title, then one meta line of facts and Tags, and at the right the
     page's answer, its actions, or both. A notice line and a TabStrip may sit under it on the band.
     `art="hero"` (home only) puts the owner's camp-road art behind it under the night veil instead;
     Base.astro's `art="hero"` starts the same art under the bar. No eyebrow, no job paragraph.

     The one head contract (design/reviews/2026-11-09-sim-head), every page alike:
     - the same top padding and title row; the crest, title and meta are top-aligned with the answer;
     - the answer is a Figure (or a who-line), at most one short line under it (the column holds a
       short measure, so a long sentence wraps instead of running across the band), and its
       actions to its left from 900 px, centred on its last line, never a row of their own;
     - a notice (a finished server sim, a sync state) is the `notice` line under the row, one style;
     - a way back (a raid's guild) is the `crumb` line on the band above the title, never above the band;
     - the TabStrip sits under the row on one rule, so tabs and a control bar start at one height. -->
<script lang="ts">
  import type { Snippet } from 'svelte';
  import { REGION } from '../../lib/ui/regions';
  import { titleColor, type TitleTone } from '../../lib/ui/primitives';
  import type { BandArt } from '../../lib/ui/primitives';

  let {
    title,
    tone = 'strong',
    crumb,
    crest,
    meta,
    lead,
    answer,
    actions,
    notice,
    tabs,
    titleTestid = 'page-title',
    pending = false,
    secondHead = false,
    testid,
    art = 'band',
  }: {
    title: string;
    tone?: TitleTone;
    /** A link back up (the raid's guild), small, on the band above the title. */
    crumb?: Snippet;
    crest?: Snippet;
    /** The one fact line: strings and Tags, separated by space. */
    meta?: Snippet;
    /** The page's one sentence under the meta line (the Lead type, 17/26, max 62ch). */
    lead?: string;
    /** The page's answer: a Figure with its label and links. */
    answer?: Snippet;
    /** The page's buttons and links: beside the answer from 900 px, or alone at the right. */
    actions?: Snippet;
    /** One line under the row: a state the page reports (a finished server sim). */
    notice?: Snippet;
    tabs?: Snippet;
    /** The title's testid, for a page whose tests name its heading ("report-title"). */
    titleTestid?: string;
    /** The title is not known yet (an account read in flight): its skeleton holds the row, the
     *  heading is for a screen reader only, and the head reads busy. */
    pending?: boolean;
    /** A head that holds a place beside the page's own server-rendered head (home's signed-in
     *  skeleton under the signed-out page): its pending title is screen-reader text, not a
     *  second h1. */
    secondHead?: boolean;
    testid?: string;
    /** The band's art: the painted blue band (every page), or home's camp-road hero. */
    art?: BandArt;
    /** Never rendered: an Astro caller's whitespace between named slots arrives as children. */
    children?: Snippet;
  } = $props();
</script>

<div class="ui-band ui-band-end full-bleed" data-art={art} style:--band-y="var(--site-top-h, 80px)">
  <div class="page-column">
    <header
      class="ui-page-header"
      data-region={REGION.head}
      data-testid={testid}
      aria-busy={pending ? 'true' : undefined}
    >
      {#if crumb}<div class="ui-page-header-crumb text-small text-muted">{@render crumb()}</div>{/if}
      <div class="ui-page-header-row">
        <div class="flex min-w-0 items-center gap-4" data-region={REGION.headTitle}>
          {#if crest}{@render crest()}{/if}
          <div class="min-w-0">
            {#if pending}<span class="skeleton-block ui-title-skeleton" aria-hidden="true"></span>
              {#if secondHead}<p class="sr-only" data-testid={titleTestid}>{title}</p>
              {:else}<h1 class="sr-only" data-testid={titleTestid}>{title}</h1>{/if}
            {:else}<h1 class="ui-title" style:color={titleColor(tone)} data-testid={titleTestid}>
                {title}
              </h1>{/if}
            {#if lead}<p class="ui-page-header-lead">{lead}</p>{/if}
            {#if meta}<div
                data-region={REGION.headMeta}
                class="text-small text-muted mt-[6px] flex flex-wrap items-center gap-x-4 gap-y-[6px]"
              >
                {@render meta()}
              </div>{/if}
          </div>
        </div>
        {#if answer || actions}<div class="ui-page-header-side">
            {#if actions}<div class="ui-page-header-actions">
                {@render actions()}
              </div>{/if}
            {#if answer}<div class="ui-page-header-answer" data-region={REGION.headAnswer}>
                {@render answer()}
              </div>{/if}
          </div>{/if}
      </div>
      {#if notice}<div class="ui-page-header-notice text-small text-muted">
          {@render notice()}
        </div>{/if}
      {#if tabs}<div class="ui-page-header-tabs">{@render tabs()}</div>{/if}
    </header>
  </div>
</div>

<style>
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
    align-items: start;
  }
  /* Below 900 the side's parts are the row's own items: the answer under the title, the actions
     under the answer, and a hidden answer (signed out) leaves no empty row behind. */
  .ui-page-header-side {
    display: contents;
  }
  @media (max-width: 899px) {
    .ui-page-header-answer {
      order: 1;
    }
    .ui-page-header-actions {
      order: 2;
    }
  }
  .ui-page-header-actions {
    display: flex;
    flex-wrap: wrap;
    align-items: center;
    gap: 10px 16px;
  }
  .ui-page-header-notice {
    margin-top: 10px;
  }
  /* A notice slot the page fills after load (a finished sim) takes no room while it is empty. */
  .ui-page-header-notice:not(:has(:not(:empty, astro-slot, astro-island))) {
    margin-top: 0;
  }
  /* The crumb takes the row's top padding: a crumbed head is no taller than another by its 44 px target. */
  .ui-page-header-crumb {
    margin: -12px 0 0;
    display: flex;
    align-items: center;
    min-height: 32px;
  }
  .ui-page-header-crumb :global(a) {
    display: inline-flex;
    align-items: center;
    min-height: 32px;
    font-size: 14px;
    color: var(--color-muted);
    text-decoration: none;
  }
  .ui-page-header-crumb :global(a:hover) {
    color: var(--color-text);
    text-decoration: underline;
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
  .ui-title-skeleton {
    display: block;
    width: 280px;
    max-width: 100%;
    height: var(--text-title-phone--line-height);
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
  @media (min-width: 900px) {
    .ui-page-header-row {
      grid-template-columns: minmax(0, 1fr) auto;
    }
    /* The actions sit beside the answer, on its last line: they never add a row. */
    .ui-page-header-side {
      display: flex;
      flex-direction: row;
      align-items: flex-end;
      gap: 24px;
      min-width: 0;
    }
    .ui-page-header-actions {
      justify-content: flex-end;
    }
    /* Beside an answer, the actions centre on its last line (a 24 px caption row, 36 px buttons). */
    .ui-page-header-side:has(> .ui-page-header-answer) > .ui-page-header-actions {
      margin-bottom: -6px;
    }
    .ui-page-header-answer {
      align-items: flex-end;
      text-align: right;
      /* A short measure: a sentence in the answer wraps here, never across the band. */
      max-width: var(--head-answer-measure, 26rem);
    }
  }
</style>
