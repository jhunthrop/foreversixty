<!-- Its styles travel with its code (css="injected"): it renders only after the island hydrates, and a
     page inlines only the styles of what its server render drew. -->
<svelte:options css="injected" />

<!-- web/src/components/planner/PlannerPart.svelte -->
<!-- One part of the planner's take-it-with-you plate (h-audit 2026-11-09): In the game, Share and
     Import were three thin plates stacked in the rail, so they share one panel, each after a
     hairline under its own small heading. -->
<script lang="ts">
  import type { Snippet } from 'svelte';
  import { RULE_TOP } from '../../lib/ui/primitives';

  let { title, testid, children }: { title: string; testid?: string; children: Snippet } = $props();
</script>

<div class="pl-part {RULE_TOP}" data-testid={testid}>
  <h3 class="pl-part-title">{title}</h3>
  {@render children()}
</div>

<style>
  .pl-part {
    margin-top: 16px;
    padding-top: 16px;
  }
  .pl-part:first-child {
    margin-top: 0;
    padding-top: 0;
    border-top: 0;
  }
  .pl-part-title {
    margin-bottom: 10px;
    font: 600 15px/20px var(--font-body);
    color: var(--color-strong);
  }
</style>
