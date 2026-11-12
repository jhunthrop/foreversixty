<!-- web/src/components/ui/CharacterAnswer.svelte -->
<!-- "Your character" as a PageHeader answer (pages spec 2026-11-08 §3A; the guides' answer as a
     primitive): the label (12 px caps), then, painted after load by a page script
     (lib/guides/guide-answer-mount.ts paintGuideAnswer: the same data-field contract), the crest
     40 left of the name 18/22 in the class colour and a Small level line, then its links; or a
     Small message and a link or a retry. Reserved at its final size (380 x 96 from 900) while
     the account loads when a session is hinted (the skeleton); no answer column signed out.
     `crest="end"` (the guides) puts the crest at the outer right edge of the right-aligned box
     from 900, the name and level line right-aligned beside it. With `testid`, the detail, next
     and action fields carry `<testid>-detail`, `-next` and `-action`. `figure` (the /bis landing,
     design/specs/2026-11-07-bis-landing.md) puts a Figure over a Meter between the name and the
     links, on one line with the Meter, reserved in the skeleton with the 44 px button the page
     paints into its links. -->
<script lang="ts">
  import type { Snippet } from 'svelte';
  import { LINK } from '../../lib/ui/primitives';

  let {
    label,
    id,
    page,
    testid,
    crest = 'start',
    figure,
  }: {
    label: string;
    /** The element the page script paints. */
    id: string;
    /** What the page tells its script (JSON in `data-page`). */
    page?: unknown;
    testid?: string;
    /** Which side of the name the crest sits on from 900: left (`start`) or the outer right edge (`end`). */
    crest?: 'start' | 'end';
    /** A Figure (and its Meter) the page script paints under the name: `data-field="figure"`. */
    figure?: Snippet;
  } = $props();

  const sub = (name: string): string | undefined => (testid === undefined ? undefined : `${testid}-${name}`);
</script>

<div
  class="ui-cans"
  {id}
  data-signed-in-only
  data-state="loading"
  data-page={page === undefined ? undefined : JSON.stringify(page)}
  data-crest={crest}
  data-figure={figure === undefined ? undefined : ''}
  data-testid={testid}
>
  <span class="sr-only">{label}</span>
  <div class="ui-cans-sk" data-answer-skeleton aria-hidden="true">
    <span class="ui-cans-row">
      <span class="skeleton-block ui-cans-sk-lead"></span>
      <span class="ui-cans-text">
        <span class="skeleton-block ui-cans-sk-name"></span>
        <span class="skeleton-block ui-cans-sk-detail"></span>
      </span>
    </span>
    {#if figure}<span class="ui-cans-sk-fig">
        <span class="skeleton-block ui-cans-sk-num"></span>
        <span class="skeleton-block ui-cans-sk-bar"></span>
      </span>{/if}
    <span class="skeleton-block ui-cans-sk-link"></span>
  </div>
  <div class="ui-cans-body" data-answer-body hidden>
    <span class="ui-cans-row" data-field="who">
      <img class="ui-cans-crest crest" data-field="crest" src="" alt="" width="40" height="40" />
      <span class="ui-cans-text">
        <b class="ui-cans-name" data-field="name"></b>
        <span class="text-small text-muted" data-field="detail" data-testid={sub('detail')}></span>
      </span>
    </span>
    {#if figure}<div class="ui-cans-fig" data-field="figure" hidden>{@render figure()}</div>{/if}
    <span class="text-small text-muted" data-field="message" hidden></span>
    <span class="ui-cans-links">
      <span class="text-small text-muted ui-cans-next" data-field="next" data-testid={sub('next')} hidden
      ></span>
      <span class="ui-cans-links" data-field="links"></span>
      <button type="button" class={LINK} data-field="action" data-testid={sub('action')} hidden></button>
    </span>
  </div>
</div>

<style>
  .ui-cans,
  .ui-cans-sk,
  .ui-cans-body {
    display: flex;
    flex-direction: column;
    gap: 6px;
    min-width: 0;
  }
  .ui-cans[data-state='hidden'],
  .ui-cans [hidden] {
    display: none;
  }
  /* No session hint, or nothing to name: the header keeps no answer column at all. */
  :global(html:not([data-session='1']) .ui-page-header-answer:has(> .ui-cans)),
  :global(.ui-page-header-answer:has(> .ui-cans[data-state='hidden'])) {
    display: none;
  }
  .ui-cans-row {
    display: flex;
    align-items: center;
    gap: 12px;
  }
  .ui-cans-text {
    display: flex;
    flex-direction: column;
    min-width: 0;
  }
  .ui-cans-sk .ui-cans-text {
    gap: 4px;
  }
  .ui-cans-sk-lead {
    width: 40px;
    height: 40px;
    border-radius: 999px;
  }
  .ui-cans-sk-name {
    width: 150px;
    height: 22px;
  }
  .ui-cans-sk-detail {
    width: 190px;
    height: 16px;
  }
  .ui-cans-sk-link {
    width: 220px;
    height: 20px;
  }
  /* The figure (and its Meter and one short line): the count, the bar beside it from 900 (under it
     on a phone), the line under both. The who row is a screen reader's: the header's selector and
     the page's own "Yours" say whose it is. */
  .ui-cans[data-figure] [data-field='who'],
  .ui-cans[data-figure] .ui-cans-sk > .ui-cans-row {
    position: absolute;
    width: 1px;
    height: 1px;
    overflow: hidden;
    clip-path: inset(50%);
    white-space: nowrap;
  }
  .ui-cans-sk-fig,
  .ui-cans-fig {
    display: flex;
    flex-wrap: wrap;
    align-items: center;
    gap: 6px 14px;
    width: 100%;
    max-width: 380px;
  }
  .ui-cans-fig :global(.ui-meter),
  .ui-cans-sk-bar {
    flex: 1 0 100%;
    height: 6px;
  }
  /* The empty track reads on the band. */
  .ui-cans-fig :global(.ui-meter.ui-meter) {
    background: rgb(255 255 255 / 0.16);
  }
  .ui-cans-fig :global([data-field='note']) {
    flex: 1 0 100%;
  }
  .ui-cans-sk-num {
    width: 120px;
    height: 32px;
  }
  .ui-cans[data-figure] .ui-cans-sk-link {
    width: 240px;
    height: 44px;
  }
  /* The answer is left-aligned under the title below 900 px: so is its figure. */
  @media (max-width: 899px) {
    .ui-cans-fig :global(.ui-figure) {
      align-items: flex-start;
      text-align: left;
    }
  }
  /* The figure's reserve below 900: the filled answer (figure, Meter, line, button) is 120 px, so
     the skeleton holds it and the page below never moves. */
  .ui-cans[data-figure] {
    min-height: 120px;
  }
  .ui-cans-crest {
    width: 40px;
    height: 40px;
  }
  .ui-cans-name {
    font: 700 18px/22px var(--font-body);
  }
  .ui-cans-links {
    display: flex;
    flex-wrap: wrap;
    align-items: center;
    gap: 4px 16px;
  }
  .ui-cans :global([data-field='links'] a:not(.button-label)) {
    color: var(--color-gold);
    font-size: 15px;
    line-height: 24px;
    font-weight: 600;
  }
  .ui-cans :global([data-field='links'] a[data-size='small']) {
    font-size: 13px;
    line-height: 20px;
  }
  .ui-cans :global([data-field='links'] a:not(.button-label):hover) {
    color: var(--color-gold-hover);
  }
  .ui-cans :global([data-field='links'] a:focus-visible) {
    outline: 2px solid var(--color-gold);
    outline-offset: 2px;
  }
  @media (min-width: 900px) {
    .ui-cans {
      min-width: 380px;
      min-height: 72px;
      align-items: flex-end;
      text-align: right;
    }
    /* With a figure: the figure and its line at the right, the button beside the figure's line. */
    .ui-cans[data-figure] {
      min-height: 56px;
    }
    .ui-cans[data-figure]:not(:has([data-field='figure'][hidden])) .ui-cans-sk,
    .ui-cans[data-figure]:not(:has([data-field='figure'][hidden])) .ui-cans-body {
      display: grid;
      grid-template-columns: auto auto;
      column-gap: 20px;
      row-gap: 6px;
      align-items: start;
    }
    .ui-cans[data-figure] .ui-cans-sk > *,
    .ui-cans[data-figure] .ui-cans-body > * {
      grid-column: 2;
    }
    .ui-cans[data-figure] .ui-cans-sk > :last-child,
    .ui-cans[data-figure] .ui-cans-body > .ui-cans-links {
      grid-column: 1;
      grid-row: 1;
      margin-top: -2px;
    }
    .ui-cans[data-figure] .ui-cans-links {
      padding-top: 0;
    }
    .ui-cans-fig :global(.ui-meter),
    .ui-cans-sk-bar {
      flex: 1 1 140px;
    }
    /* The figure's line holds its count and Meter side by side; the short line runs under both. */
    .ui-cans[data-figure] .ui-cans-fig,
    .ui-cans[data-figure] .ui-cans-sk-fig {
      width: 360px;
    }
    .ui-cans-fig :global([data-field='note']) {
      text-align: right;
    }
    .ui-cans-sk,
    .ui-cans-body {
      align-items: flex-end;
    }
    .ui-cans-row,
    .ui-cans-links {
      justify-content: flex-end;
    }
    .ui-cans-text {
      text-align: left;
    }
    .ui-cans-sk-fig,
    .ui-cans-fig {
      flex-direction: row-reverse;
    }
    .ui-cans-next {
      flex-basis: 100%;
      text-align: right;
    }
    .ui-cans[data-crest='end'] .ui-cans-row {
      flex-direction: row-reverse;
    }
    .ui-cans[data-crest='end'] .ui-cans-text {
      align-items: flex-end;
      text-align: right;
    }
  }
</style>
