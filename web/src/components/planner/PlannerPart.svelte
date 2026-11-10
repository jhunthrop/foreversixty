<!-- web/src/components/planner/PlannerPart.svelte -->
<!-- One part of the planner's take-it-with-you plate (h-audit 2026-11-09): In the game, Share and
     Import were three thin plates stacked in the rail, so they share one panel, each after a
     hairline under its own small heading. -->
<script lang="ts">
  import type { Snippet } from 'svelte';

  let { title, testid, children }: { title: string; testid?: string; children: Snippet } = $props();
</script>

<div class="pl-part" data-testid={testid}>
  <h3 class="pl-part-title">{title}</h3>
  {@render children()}
</div>

<style>
  .pl-part {
    margin-top: 16px;
    padding-top: 16px;
    border-top: 1px solid var(--color-line-soft);
  }
  .pl-part-title {
    margin-bottom: 10px;
    font: 600 15px/20px var(--font-body);
    color: var(--color-strong);
  }
</style>
