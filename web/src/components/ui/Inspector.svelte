<!-- web/src/components/ui/Inspector.svelte -->
<!-- The detail panel for one selected thing (web layout spec 2026-11-05 §4.5): a head, then labelled
     sections on hairlines. `side` (1024 and up): a column inside the Plate, --inspector-w wide, one
     --color-line rule at its left, sticky under the bar and scrolling within the viewport. `inline`
     (below 1024): the same content expanded under the tapped row with a 3 px gold bar joining it to
     the row; the caller hides the head there, since the row above already shows the item. -->
<script lang="ts" module>
  import type { Snippet } from 'svelte';
  import { REGION } from '../../lib/ui/regions';

  export interface InspectorSection {
    /** 12 px caps label over the section ("Where to get it"). */
    label: string;
    body: Snippet;
    testid?: string;
  }
</script>

<script lang="ts">
  let {
    selection,
    head,
    sections,
    placement = 'side',
    id,
    testid,
  }: {
    /** What is selected, for the panel's accessible name ("Head detail"). */
    selection: string;
    head?: Snippet;
    sections: readonly InspectorSection[];
    placement?: 'side' | 'inline';
    id?: string;
    testid?: string;
  } = $props();
</script>

<aside
  {id}
  class="ui-inspector min-w-0"
  data-placement={placement}
  data-region={placement === 'side' ? REGION.inspector : REGION.inspectorInline}
  aria-label={`${selection} detail`}
  aria-live="polite"
  tabindex={placement === 'side' ? 0 : undefined}
  data-testid={testid}
>
  {#if head}<div class="ui-inspector-head">{@render head()}</div>{/if}
  {#each sections as section (section.label)}
    <section class="ui-inspector-section" data-testid={section.testid}>
      <h3 class="text-label tracking-label text-muted font-bold uppercase">{section.label}</h3>
      <div class="mt-1 text-[14px] leading-[22px]">{@render section.body()}</div>
    </section>
  {/each}
</aside>

<style>
  .ui-inspector[data-placement='side'] {
    width: var(--inspector-w);
    border-left: 1px solid var(--color-line);
    position: sticky;
    top: calc(var(--bar-h) + 4px);
    max-height: calc(100vh - 72px);
    overflow-y: auto;
    align-self: start;
  }
  .ui-inspector[data-placement='inline'] {
    box-shadow: inset 3px 0 0 var(--color-gold);
    background: var(--color-selected);
  }
  .ui-inspector-head,
  .ui-inspector-section {
    padding: 14px 16px;
  }
  .ui-inspector-section + .ui-inspector-section,
  .ui-inspector-head + .ui-inspector-section {
    border-top: 1px solid var(--color-line-soft);
  }
</style>
