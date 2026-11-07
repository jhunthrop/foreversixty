<!-- web/src/components/ui/Facts.svelte -->
<!-- A key/value list (sim spec 2026-11-06 §4.5 "This run", ruling S8): the key muted at the left,
     the value strong and right-aligned with tabular figures, an optional small muted line under the
     value, a hairline between pairs. A value can be markup (`content`: a control) or literal text a player types (`code`). A `dl`, so each pair reads as a term and its definition. -->
<script lang="ts">
  import type { Fact } from '../../lib/ui/facts';
  import Code from './Code.svelte';

  let { items, testid }: { items: readonly Fact[]; testid?: string } = $props();
</script>

<dl class="ui-facts m-0 grid grid-cols-[minmax(0,auto)_minmax(0,1fr)]" data-testid={testid}>
  {#each items as item (item.id)}
    <dt class="ui-facts-cell text-muted pr-4">{item.label}</dt>
    <dd
      class={[
        'ui-facts-cell m-0 text-right tabular-nums',
        item.muted ? 'text-muted' : 'text-strong font-semibold',
      ]}
      data-testid={item.testid}
    >
      {#if item.content}{@render item.content()}{:else if item.code !== undefined}<Code
          text={item.code}
        />{:else}{item.value ?? ''}{/if}{#if item.sub}<small class="text-muted block text-[13px] font-normal"
          >{item.sub}</small
        >{/if}
    </dd>
  {/each}
</dl>

<style>
  .ui-facts-cell {
    padding: 10px 0;
    border-bottom: 1px solid var(--color-line-soft);
    font-size: 14px;
    line-height: 20px;
  }
</style>
