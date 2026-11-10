<!-- web/src/components/sim/tools/ToolSection.svelte -->
<!-- One block of a sim tool page (web layout spec 2026-11-05 §3A.2, layout L3): a ui/Section panel
     (round 5), its title inside the panel in sentence case. Every block is the same plain panel:
     one border, no inner frame (h-audit 2026-11-09). -->
<script lang="ts">
  import type { Snippet } from 'svelte';
  import Section from '../../ui/Section.svelte';

  let {
    title,
    testid,
    meta,
    head,
    children,
  }: {
    title: string;
    testid?: string;
    meta?: string;
    head?: Snippet;
    children: Snippet;
  } = $props();
</script>

<Section {title} {testid} {meta} {head} panel>
  {@render children()}
</Section>
