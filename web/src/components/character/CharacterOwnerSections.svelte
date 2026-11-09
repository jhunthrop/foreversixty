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
  /** The account page's and the coach tab's links land here (`#history`, `#coach`): the address opens
   *  its Disclosure, whether the page opened on it or the hash changed after; the opened section
   *  (CharacterHistory, CoachHistoryList) carries the id and scrolls itself into view. */
  const HISTORY_ANCHOR = 'history';
  const COACH_ANCHOR = 'coach';
  $effect(() => {
    const openFromHash = (): void => {
      const hash = window.location.hash.slice(1);
      if (hash === HISTORY_ANCHOR && history) historyOpen = true;
      if (hash === COACH_ANCHOR && coach) coachOpen = true;
    };
    openFromHash();
    window.addEventListener('hashchange', openFromHash);
    return () => window.removeEventListener('hashchange', openFromHash);
  });
  $effect(() => {
    if (historyOpen) charts.load();
  });
  $effect(() => {
    if (coachOpen) reviews.load();
  });
</script>

<Section title={copy.history} panel testid="character-owner-history">
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
  <Section title={copy.coach} panel testid="character-coach-section">
    <Disclosure summary={copy.coachLead} bind:open={coachOpen} testid="character-coach-toggle">
      {#if reviews.current}<reviews.current {path} />{/if}
    </Disclosure>
  </Section>
{/if}
