<!-- web/src/components/character/HistoryMount.svelte -->
<!-- What the character page shows under the Forever Score card (experience spec
     2026-10-24-character-history 3.1 and 3.9): the owner's History section (loaded lazily, so
     no other viewer's bundle carries it), the free owner's one sentence, or nothing. Until
     /v1/me has answered, the section's place is held (3.11) when the pre-paint hint says the
     viewer will get it. -->
<script lang="ts">
  import { historyCopy } from '../../lib/character/history-copy';
  import { PREMIUM_LINKS } from '../../lib/billing/premium-links';
  import { historyMount, type HistoryMountFacts } from '../../lib/character/history-mount';
  import type { CharacterPath } from '../../lib/characters';
  import HistoryReserve from './HistoryReserve.svelte';
  import { createLazyComponent } from '../../lib/report/lazy-component.svelte';
  import type { Component } from 'svelte';

  let { path, ...facts }: { path: CharacterPath } & HistoryMountFacts = $props();

  const mount = $derived(historyMount(facts));
  const section = createLazyComponent<{ path: CharacterPath }>(
    () => import('./CharacterHistory.svelte') as Promise<{ default: Component<{ path: CharacterPath }> }>,
  );

  $effect(() => {
    if (mount === 'charts') section.load();
  });
</script>

{#if mount === 'charts'}
  {#if section.current}
    <section.current {path} />
  {:else}
    <!-- The chunk is on its way: the reserved box, drawn only under the hint. -->
    <HistoryReserve />
  {/if}
{:else if mount === 'pending'}
  <HistoryReserve />
{:else if mount === 'sentence'}
  <p class="text-muted -order-2 flex min-h-11 items-center text-[13px]" data-testid="history-free-sentence">
    <!-- One line of flow: the link stays inline with its space and stop; padding gives it the
         44px target without making it a block. -->
    <span
      >{historyCopy.freeOwnerSentence}&#32;<span class="whitespace-nowrap"
        ><a
          class="text-text hover:text-gold-hover py-3.5 underline underline-offset-2"
          href={PREMIUM_LINKS.history}
          data-testid="history-see-plans">{historyCopy.freeOwnerLink}</a
        >.</span
      ></span
    >
  </p>
{/if}
