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
    id,
    testid,
    head,
    children,
  }: {
    title: string;
    level?: 2 | 3;
    meta?: string;
    action?: { label: string; href: string };
    reading?: boolean;
    id?: string;
    testid?: string;
    /** Extra content in the head row, right-aligned (a toggle, a Tag). */
    head?: Snippet;
    children: Snippet;
  } = $props();
</script>

<section
  {id}
  class="ui-section min-w-0 scroll-mt-[calc(var(--bar-h)+12px)]"
  data-testid={testid}
  data-region={sectionRegion(title)}
  aria-labelledby={id ? `${id}-title` : undefined}
>
  <div class="border-line-soft flex flex-wrap items-baseline gap-x-3 gap-y-1 border-b pb-[10px]">
    <svelte:element
      this={level === 3 ? 'h3' : 'h2'}
      id={id ? `${id}-title` : undefined}
      class={['text-strong font-bold', reading ? 'text-section md:text-section-reading' : 'text-section']}
      >{title}</svelte:element
    >
    {#if meta}<span class="text-small text-muted">{meta}</span>{/if}
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
