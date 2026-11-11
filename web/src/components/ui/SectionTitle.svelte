<!-- web/src/components/ui/SectionTitle.svelte -->
<!-- The one section heading (round 5, spec 2026-11-05 §3.3; component audit 2026-11-10 cluster 15):
     Marcellus on the ground, cream, mixed case, 22 px (24 px on a reading page). ui/Section draws
     its head with it; a heading that sits where a Section cannot (a stop inside the BiS itinerary,
     the head of a rail Frame) takes it on its own, so no page writes its own heading face. -->
<script lang="ts">
  import type { Snippet } from 'svelte';

  let {
    title,
    level = 2,
    reading = false,
    id,
    focusable = false,
    class: extra = '',
    testid,
    element = $bindable(),
    after,
  }: {
    title: string;
    level?: 2 | 3;
    /** The 24 px reading-page head. */
    reading?: boolean;
    id?: string;
    /** A programmatic focus target (tabindex -1, no outline). */
    focusable?: boolean;
    /** Layout only (padding, a page's own hook class). */
    class?: string;
    testid?: string;
    /** The heading element, for a caller that moves focus to it. */
    element?: HTMLElement;
    /** Muted words after the title (a stop's level range). */
    after?: Snippet;
  } = $props();
</script>

<svelte:element
  this={level === 3 ? 'h3' : 'h2'}
  bind:this={element}
  {id}
  tabindex={focusable ? -1 : undefined}
  class={[
    'ui-section-title',
    reading ? 'text-section-reading' : 'text-section',
    focusable && 'outline-none',
    extra,
  ]}
  data-testid={testid}>{title}{@render after?.()}</svelte:element
>

<style>
  .ui-section-title {
    font-family: var(--font-heading);
    font-weight: 400;
    letter-spacing: var(--tracking-heading);
    color: var(--color-cream);
    text-shadow: 0 1px 2px rgb(0 0 0 / 0.45);
  }
</style>
