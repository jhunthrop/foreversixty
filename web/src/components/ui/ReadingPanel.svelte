<!-- web/src/components/ui/ReadingPanel.svelte -->
<!-- A reading page's sections on one panel (round 5, directions mock b2-guide-spec): the
     ReadingSections inside it lose their head hairline and take a hairline between them instead,
     each padded 22/24 (18/16 below 768), the heading in Marcellus cream; from 1280 the prose takes
     the room and the lane is a 380 px column behind a hairline. Class, quality and faction colours
     in the lanes sit on the panel, where they clear contrast. Holds ReadingSections only. -->
<script lang="ts">
  import type { Snippet } from 'svelte';
  import { PLATE_SURFACE } from '../../lib/ui/primitives';

  let {
    testid,
    children,
  }: {
    testid?: string;
    children: Snippet;
  } = $props();
</script>

<div class="{PLATE_SURFACE} ui-reading-panel" data-testid={testid}>{@render children()}</div>

<style>
  .ui-reading-panel {
    min-width: 0;
  }
  .ui-reading-panel > :global(.ui-rsec) {
    padding: 22px 24px;
  }
  .ui-reading-panel > :global(.ui-rsec + .ui-rsec) {
    border-top: 1px solid var(--color-line-soft);
  }
  .ui-reading-panel > :global(.ui-rsec > :first-child) {
    padding-bottom: 0;
    border-bottom: 0;
  }
  .ui-reading-panel :global(.ui-rsec-title) {
    font-family: var(--font-title);
    font-weight: 400;
    font-size: var(--text-section-reading);
    line-height: var(--text-section-reading--line-height);
    letter-spacing: var(--tracking-heading);
    color: var(--color-cream);
  }
  .ui-reading-panel :global(.ui-rd) {
    padding-top: 12px;
  }
  @media (max-width: 767px) {
    .ui-reading-panel > :global(.ui-rsec) {
      padding: 18px 16px;
    }
  }
  @media (min-width: 1280px) {
    .ui-reading-panel :global(.ui-rd:not([data-solo])) {
      grid-template-columns: minmax(0, 1fr) minmax(0, 380px);
      column-gap: 28px;
    }
    .ui-reading-panel :global(.ui-rd:not([data-solo]) > .ui-rd-lane) {
      padding-left: 18px;
      border-left: 1px solid var(--color-line-soft);
    }
  }
</style>
