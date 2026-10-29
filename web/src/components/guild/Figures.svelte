<!-- web/src/components/guild/Figures.svelte -->
<!-- A copy string with its figures in mono and its words in the body face (attendance spec 3.1,
     3.3); the string never breaks inside (`2 h 41 min` stays on one line). -->
<script lang="ts">
  import { figureParts } from '../../lib/guild/figures';

  let { text, class: extra = '' }: { text: string; class?: string } = $props();
  const parts = $derived(figureParts(text));
</script>

<span class="whitespace-nowrap {extra}"
  >{#each parts as part, index (index)}{#if part.figure}<span class="numeric">{part.text}</span
      >{:else}{part.text}{/if}{/each}</span
>
