<!-- web/src/components/ui/StatRow.svelte -->
<!-- One row of a StatTable (sim spec 2026-11-06 §4.7): the icon, the name over its muted sub line,
     the Meter, then the figure cells, on the table's grid. The first figure is the primary one
     (700, strong) and keeps its `phone` line under it below 1024, where the other figures and the
     bar hide. The icon and the bar are decorative (the name and the figures say it all); `total`
     is the 44 px muted foot row, which keeps an empty bar cell so its figures stay under their
     heads. Parts of the sub line marked `ui-st-phone` show only on a phone. -->
<script lang="ts">
  import type { Snippet } from 'svelte';
  import type { StatFigure } from '../../lib/ui/stat-table';
  import './stat-table.css';

  let {
    icon,
    name,
    sub,
    figures,
    meter,
    bar = true,
    phone,
    total = false,
    testid,
  }: {
    icon?: Snippet;
    name: Snippet;
    sub?: Snippet;
    figures: readonly StatFigure[];
    /** The Meter in the bar column. */
    meter?: Snippet;
    /** The table has a bar column (keep in step with StatTable's `bar`). */
    bar?: boolean;
    /** The primary figure's phone line ("14.6%" under the DPS). */
    phone?: string;
    total?: boolean;
    testid?: string;
    /** Never rendered: an Astro caller's whitespace between named slots arrives as children. */
    children?: Snippet;
  } = $props();
</script>

<div class="ui-st-r ui-st-c" role="row" data-total={total ? 'true' : undefined} data-testid={testid}>
  <span class="ui-st-icon" aria-hidden="true"
    >{#if icon}{@render icon()}{/if}</span
  >
  <div class="ui-st-name" role="cell">
    {@render name()}
    {#if sub}<div class="ui-st-sub">{@render sub()}</div>{/if}
  </div>
  {#if bar}<span class="ui-st-bar" aria-hidden="true"
      >{#if meter}{@render meter()}{/if}</span
    >{/if}
  {#each figures as figure, index (index)}
    <span
      class="ui-st-f"
      role="cell"
      data-primary={index === 0 ? 'true' : undefined}
      data-muted={figure.muted === true ? 'true' : undefined}
      data-tone={figure.tone}
      data-testid={figure.testid}
      >{figure.value}{#if index === 0 && phone !== undefined}<span class="ui-st-phone">{phone}</span
        >{/if}</span
    >
  {/each}
</div>
