<!-- web/src/components/ui/Section.svelte -->
<!-- A section (round 5): its heading in Marcellus on the ground, cream, mixed case ("Latest on
     Forever"), an optional muted meta and one action link at the right, then its body. `boxed` puts
     the body on the panel (ui/panel.css) 12 px under the heading, as every list and table in the
     mocks sits: class and quality colours only clear contrast on the panel, so a page lane boxes every
     list, table and row set. A boxed body that already holds a panel (a Plate, a grid of panels) is
     not boxed again. Open (the default until each page lane moves its sections), the body sits on the
     ground under the heading's hairline. Sections sit --section-gap apart (Page); a Section is never
     inside a Plate. `reading` sets the 24 px reading-page head.
     `panel` (the logs and report mocks' prims.panel) puts the heading row inside the panel: one
     Plate (its `frame`: plain, framed or lead) holds the head and the body, the body padded 20 px,
     or edge to edge with `flush` for a table. Page's `panels` sets them 22 px apart. -->
<script lang="ts">
  import type { Snippet } from 'svelte';
  import { sectionRegion } from '../../lib/ui/regions';
  import { LINK, type PanelFrame } from '../../lib/ui/primitives';
  import Plate from './Plate.svelte';
  import SectionTitle from './SectionTitle.svelte';

  let {
    title,
    level = 2,
    meta,
    action,
    reading = false,
    metaTestid,
    id,
    anchor,
    testid,
    focusId,
    titleTestid,
    head,
    boxed = false,
    panel = false,
    frame = 'plain',
    flush = false,
    children,
  }: {
    title: string;
    level?: 2 | 3;
    meta?: string;
    /** One link at the head's right. `arrow`: the link reads `›` alone (its label is its accessible
     *  name), for a plate whose title already names where it goes ("Buffs" › the Buffs tab). */
    action?: { label: string; href: string; arrow?: boolean };
    reading?: boolean;
    /** The meta line's testid, for a test that reads the line alone. */
    metaTestid?: string;
    /** Names the section (aria-labelledby its heading: a landmark) and is its link target. */
    id?: string;
    /** A link target only (`/logs#upload`): the section stays an unnamed block. */
    anchor?: string;
    testid?: string;
    /** The heading's own id, which also makes it a programmatic focus target (tabindex -1): a
     *  panel that moves focus to its heading once it settles (the report's Coach tab). */
    focusId?: string;
    /** The heading's testid, for a test that reads the title alone. */
    titleTestid?: string;
    /** Extra content in the head row, right-aligned (a toggle, a Tag). */
    head?: Snippet;
    /** The body on the panel (every list, table and row set); off, it sits on the ground under a hairline. */
    boxed?: boolean;
    /** The heading row inside the panel with the body: one Plate around both. */
    panel?: boolean;
    /** `panel` only: the Plate's frame (two or three framed panels a page, one lead). */
    frame?: PanelFrame;
    /** `panel` only: the body runs edge to edge (a table that pads its own first and last cells). */
    flush?: boolean;
    children: Snippet;
  } = $props();
</script>

{#snippet headRow()}
  <div
    class={[
      'flex flex-wrap items-baseline gap-x-3 gap-y-1',
      panel ? 'ui-section-panel-head' : 'pb-[10px]',
      !boxed && !panel && 'ui-section-rule',
    ]}
  >
    <SectionTitle
      {level}
      {reading}
      id={focusId ?? (id ? `${id}-title` : undefined)}
      focusable={focusId !== undefined}
      testid={titleTestid}
      {title}
    />
    {#if meta}<span class="text-small text-muted" data-testid={metaTestid}>{meta}</span>{/if}
    {#if head || action}
      <span class="ml-auto flex items-center gap-4">
        {#if head}{@render head()}{/if}
        <!-- On a touch screen the link's hit area is 44 px tall: 12 px of padding above and below,
             taken back by the same negative margin, so the head row does not grow. -->
        {#if action?.arrow}<a
            class="{LINK} -my-3 inline-flex min-h-11 min-w-11 items-center justify-end text-[18px]"
            href={action.href}
            aria-label={action.label}
            title={action.label}>›</a
          >{:else if action}<a
            class="{LINK} text-[14px] [@media(pointer:coarse)]:-my-3 [@media(pointer:coarse)]:inline-block [@media(pointer:coarse)]:py-3"
            href={action.href}>{action.label} ›</a
          >{/if}
      </span>
    {/if}
  </div>
{/snippet}

<section
  id={id ?? anchor}
  class={['ui-section min-w-0 scroll-mt-[calc(var(--bar-h)+12px)]', panel && 'ui-section-panel']}
  data-testid={testid}
  data-region={sectionRegion(title)}
  aria-labelledby={id ? (focusId ?? `${id}-title`) : undefined}
>
  {#if panel}
    <Plate {frame}>
      {@render headRow()}
      <div class={['ui-section-body', flush ? 'ui-section-panel-flush' : 'ui-section-panel-body']}>
        {@render children()}
      </div>
    </Plate>
  {:else}
    {@render headRow()}
    <div class={['ui-section-body', boxed ? 'ui-surface ui-section-box' : 'pt-3']}>{@render children()}</div>
  {/if}
</section>

<style>
  .ui-section-rule {
    border-bottom: 1px solid rgb(255 255 255 / 0.14);
  }
  .ui-section-box {
    margin-top: 12px;
    padding: 8px 20px;
  }
  /* A body that brings its own panel is not boxed twice. */
  .ui-section-box:has(:global(.ui-surface), :global([data-region='plate'])) {
    padding: 0;
    border: 0;
    border-radius: 0;
    background: none;
    box-shadow: none;
  }
  /* The panel form (prims.panel): the head row and the body inside one Plate. */
  .ui-section-panel-head {
    padding: 16px 20px 10px;
  }
  .ui-section-panel-body {
    padding: 0 20px 16px;
  }
  .ui-section-panel-flush {
    padding-bottom: 6px;
  }
  @media (max-width: 767px) {
    .ui-section-box {
      padding: 4px 14px;
    }
    .ui-section-panel-head {
      padding: 14px 14px 8px;
    }
    .ui-section-panel-body {
      padding: 0 14px 14px;
    }
  }
</style>
