<!-- web/src/components/guild/FactLine.svelte -->
<!-- A line of guild facts as sentences (P.1): the body face, " · " with a space each side, each fact
     one nowrap piece carrying its own trailing dot, so a wrap moves whole facts and never leaves a
     dot at the start of a line or splits "1 wipe". The plain spaces between facts are the only
     breaks; they are mustaches because Svelte drops a bare space at the end of an {#each} body. -->
<script lang="ts">
  import { FACT_TAIL, asFacts, type Fact } from '../../lib/guild/fact-line';

  let {
    facts,
    class: className = '',
    testid = undefined,
  }: { facts: readonly (Fact | string)[]; class?: string; testid?: string } = $props();

  const items = $derived(asFacts(facts));
</script>

<!-- eslint-disable svelte/no-useless-mustaches -->
<span class={className} data-testid={testid}
  >{#each items as fact, i (i)}<span class="whitespace-nowrap"
      ><span
        class={fact.class ?? ''}
        style={fact.title === undefined ? '' : 'cursor:help'}
        title={fact.title}
        data-testid={fact.testid}>{fact.text}</span
      >{#if i < items.length - 1}<span class="text-muted" aria-hidden="true">{FACT_TAIL}</span>{/if}</span
    >{#if i < items.length - 1}{' '}{/if}{/each}</span
>
<!-- eslint-enable svelte/no-useless-mustaches -->
