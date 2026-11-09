<!-- web/src/components/sim/tools/ToolSection.svelte -->
<!-- One block of a sim tool page (web layout spec 2026-11-05 §3A.2, layout L3): a ui/Section panel
     (round 5), its title inside the panel in sentence case. The composer passes `plate` on exactly
     one block per page (the primary object), which takes the framed panel; every other block is a
     plain panel. -->
<script lang="ts">
  import type { Snippet } from 'svelte';
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
    /** The page's primary object: the framed panel. */
    plate?: boolean;
    meta?: string;
    head?: Snippet;
    children: Snippet;
  } = $props();
</script>

<Section {title} {testid} {meta} {head} panel frame={plate ? 'framed' : 'plain'}>
  {@render children()}
</Section>
