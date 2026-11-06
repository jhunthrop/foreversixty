<!-- web/src/components/character/CharacterOwnerSections.svelte -->
<!-- The owner's extras below the fights (character spec 2026-11-06 C8): History (score over time,
     then parse per boss) collapsed, its charts' chunk read only when opened; a free owner gets one
     line to the plans instead. Coach history last, collapsed, Premium only. -->
<script lang="ts">
  import type { Component } from 'svelte';
  import type { CharacterPath } from '../../lib/characters';
  import { characterHrefs, characterPageCopy as copy } from '../../lib/character/page-view';
  import { createLazyComponent } from '../../lib/report/lazy-component.svelte';
  import { LINK } from '../../lib/ui/primitives';
  import Disclosure from '../ui/Disclosure.svelte';
  import Row from '../ui/Row.svelte';
  import Section from '../ui/Section.svelte';

  let { path, history, coach }: { path: CharacterPath; history: boolean; coach: boolean } = $props();

  let historyOpen = $state(false);
  let coachOpen = $state(false);
  const charts = createLazyComponent<{ path: CharacterPath }>(
    () => import('./CharacterHistory.svelte') as Promise<{ default: Component<{ path: CharacterPath }> }>,
  );
  const reviews = createLazyComponent<{ path: CharacterPath }>(
    () =>
      import('../coach/CoachHistoryList.svelte') as Promise<{ default: Component<{ path: CharacterPath }> }>,
  );
  $effect(() => {
    if (historyOpen) charts.load();
  });
  $effect(() => {
    if (coachOpen) reviews.load();
  });
</script>

<Section title={copy.history} testid="character-history-section">
  {#if history}
    <Disclosure summary={copy.historyLead} bind:open={historyOpen} testid="character-history-toggle">
      {#if charts.current}<charts.current {path} />{/if}
    </Disclosure>
  {:else}
    <Row title={copy.historyLockTitle} sub={copy.historyLockSub} open testid="history-free-sentence">
      {#snippet trail()}<a
          class="{LINK} text-[14px]"
          href={characterHrefs.plans}
          data-testid="history-see-plans">{copy.seePlans} ›</a
        >{/snippet}
    </Row>
  {/if}
</Section>
{#if coach}
  <Section title={copy.coach} testid="character-coach-section">
    <Disclosure summary={copy.coachLead} bind:open={coachOpen} testid="character-coach-toggle">
      {#if reviews.current}<reviews.current {path} />{/if}
    </Disclosure>
  </Section>
{/if}
