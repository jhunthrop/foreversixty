<!-- web/src/components/ui/PlanHead.svelte -->
<!-- A plan column's head (premium spec 2026-11-08 section 3A): the name in Barlow 700 22/28 with an
     optional Tag, the audience in Small, the price as a big Figure (gold on the one plan the page
     sells, strong on the others), then its small lines: muted, or gold for the founder line. The
     column's one Button sits under it, outside this block. -->
<script lang="ts" module>
  export interface PlanHeadLine {
    text: string;
    tone?: 'muted' | 'gold';
    testid?: string;
  }
</script>

<script lang="ts">
  import type { Snippet } from 'svelte';
  import Figure from './Figure.svelte';

  let {
    name,
    audience,
    value,
    unit,
    lines = [],
    gold = false,
    level = 3,
    tag,
    price,
    testid,
  }: {
    name: string;
    audience: string;
    value: string;
    unit?: string;
    lines?: readonly PlanHeadLine[];
    /** The plan the page sells: the price in gold. */
    gold?: boolean;
    level?: 2 | 3;
    /** A Tag beside the name (`Your plan`, a guild's name). */
    tag?: Snippet;
    /** Replaces the price Figure (a static page that draws both intervals in fixed cells). */
    price?: Snippet;
    testid?: string;
  } = $props();
</script>

<div class="ui-plh" data-testid={testid}>
  <svelte:element this={`h${level}`} class="ui-plh-name">{name}{#if tag}{@render tag()}{/if}</svelte:element>
  <span class="ui-plh-aud">{audience}</span>
  <div class="ui-plh-fig">
    {#if price}{@render price()}{:else}<Figure {value} {unit} tone={gold ? 'gold' : 'strong'} size="big" />{/if}
  </div>
  {#each lines as line, i (i)}
    <span class={['ui-plh-line', line.tone === 'gold' && 'is-gold']} data-testid={line.testid}>{line.text}</span>
  {/each}
</div>

<style>
  .ui-plh {
    display: flex;
    flex-direction: column;
    gap: 4px;
    min-width: 0;
  }
  .ui-plh-name {
    display: flex;
    align-items: center;
    gap: 10px;
    margin: 0;
    font: 700 22px/28px var(--font-body);
    color: var(--color-strong);
  }
  .ui-plh-aud {
    font-size: 14px;
    line-height: 20px;
    color: var(--color-muted);
  }
  .ui-plh-fig {
    margin-top: 8px;
  }
  .ui-plh-line {
    font-size: 13px;
    line-height: 20px;
    color: var(--color-muted);
  }
  .ui-plh-line.is-gold {
    color: var(--color-gold);
  }
</style>
