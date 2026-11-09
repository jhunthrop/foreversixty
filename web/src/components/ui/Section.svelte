<!-- web/src/components/ui/Section.svelte -->
<!-- A section (round 5): its heading in Marcellus on the ground, cream, mixed case ("Latest on
     Forever"), an optional muted meta and one action link at the right, then its body. `boxed` puts
     the body on the panel (ui/panel.css) 12 px under the heading, as every list and table in the
     mocks sits: class and quality colours only clear contrast on the panel, so a page lane boxes every
     list, table and row set. A boxed body that already holds a panel (a Plate, a grid of panels) is
     not boxed again. Open (the default until each page lane moves its sections), the body sits on the
     ground under the heading's hairline. Sections sit --section-gap apart (Page); a Section is never
     inside a Plate. `reading` sets the 24 px reading-page head. -->
<script lang="ts">
  import type { Snippet } from 'svelte';
  import { sectionRegion } from '../../lib/ui/regions';
  import { LINK } from '../../lib/ui/primitives';

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
    children,
  }: {
    title: string;
    level?: 2 | 3;
    meta?: string;
    action?: { label: string; href: string };
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
    children: Snippet;
  } = $props();
</script>

<section
  id={id ?? anchor}
  class="ui-section min-w-0 scroll-mt-[calc(var(--bar-h)+12px)]"
  data-testid={testid}
  data-region={sectionRegion(title)}
  aria-labelledby={id ? (focusId ?? `${id}-title`) : undefined}
>
  <div class={['flex flex-wrap items-baseline gap-x-3 gap-y-1 pb-[10px]', !boxed && 'ui-section-rule']}>
    <svelte:element
      this={level === 3 ? 'h3' : 'h2'}
      id={focusId ?? (id ? `${id}-title` : undefined)}
      tabindex={focusId ? -1 : undefined}
      class={[
        'ui-section-title',
        reading ? 'text-section-reading' : 'text-section',
        focusId && 'outline-none',
      ]}
      data-testid={titleTestid}>{title}</svelte:element
    >
    {#if meta}<span class="text-small text-muted" data-testid={metaTestid}>{meta}</span>{/if}
    {#if head || action}
      <span class="ml-auto flex items-center gap-4">
        {#if head}{@render head()}{/if}
        <!-- On a touch screen the link's hit area is 44 px tall: 12 px of padding above and below,
             taken back by the same negative margin, so the head row does not grow. -->
        {#if action}<a
            class="{LINK} text-[14px] [@media(pointer:coarse)]:-my-3 [@media(pointer:coarse)]:inline-block [@media(pointer:coarse)]:py-3"
            href={action.href}>{action.label} ›</a
          >{/if}
      </span>
    {/if}
  </div>
  <div class={['ui-section-body', boxed ? 'ui-surface ui-section-box' : 'pt-3']}>{@render children()}</div>
</section>

<style>
  .ui-section-rule {
    border-bottom: 1px solid rgb(255 255 255 / 0.14);
  }
  .ui-section-title {
    font-family: var(--font-heading);
    font-weight: 400;
    letter-spacing: var(--tracking-heading);
    color: var(--color-cream);
    text-shadow: 0 1px 2px rgb(0 0 0 / 0.45);
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
  @media (max-width: 767px) {
    .ui-section-box {
      padding: 4px 14px;
    }
  }
</style>
