<!-- web/src/components/sim/tools/ToolSection.svelte -->
<!-- One block of a sim tool page (web layout spec 2026-11-05 §3A.2, layout L3): an open ui/Section
     with its title in sentence case, its body optionally on the page's one ui/Plate. The composer
     passes `plate` on exactly one block per page (the primary object); every other block stays open. -->
<script lang="ts">
  import type { Snippet } from 'svelte';
  import Plate from '../../ui/Plate.svelte';
  import Section from '../../ui/Section.svelte';

  let {
    title,
    testid,
    plate = false,
    meta,
    head,
    children,
  }: {
    title: string;
    testid?: string;
    /** The body sits on the page's one Plate. */
    plate?: boolean;
    meta?: string;
    head?: Snippet;
    children: Snippet;
  } = $props();
</script>

<Section {title} {testid} {meta} {head}>
  {#if plate}
    <Plate flush={false}>{@render children()}</Plate>
  {:else}
    {@render children()}
  {/if}
</Section>
