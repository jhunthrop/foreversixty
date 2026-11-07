<!-- web/src/components/ui/Section.svelte -->
<!-- An open section (web layout spec 2026-11-05 §3A.2): a text heading (the body face 700, normal
     case, --color-strong) over a --color-line-soft hairline, an optional muted meta and one action
     link at the right, then the body on S0 with no box. Sections sit --section-gap apart (Page);
     a Section is never inside a Plate. `reading` sets the 22 px reading-page head. -->
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
  <div class="border-line-soft flex flex-wrap items-baseline gap-x-3 gap-y-1 border-b pb-[10px]">
    <svelte:element
      this={level === 3 ? 'h3' : 'h2'}
      id={focusId ?? (id ? `${id}-title` : undefined)}
      tabindex={focusId ? -1 : undefined}
      class={[
        'text-strong font-bold',
        reading ? 'text-section md:text-section-reading' : 'text-section',
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
  <div class="ui-section-body pt-3">{@render children()}</div>
</section>
