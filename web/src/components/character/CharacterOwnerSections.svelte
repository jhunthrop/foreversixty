<!-- web/src/components/character/CharacterOwnerSections.svelte -->
<!-- The owner's extras below the fights (character spec 2026-11-06 C8), one History panel since the
     h-audit review (2026-11-09: two panels each holding one closed line read thin): the score and
     parse charts collapsed with the fight count beside them, their chunk read only when opened (a
     free owner gets one line to the plans instead), then the coach's reviews, collapsed, Premium only. -->
<script lang="ts">
  import type { Component } from 'svelte';
  import type { CharacterPath } from '../../lib/characters';
  import { characterHrefs, characterPageCopy as copy, dayMonth } from '../../lib/character/page-view';
  import type { ScorePoint } from '../../lib/score/score';
  import { createLazyComponent } from '../../lib/report/lazy-component.svelte';
  import { LINK } from '../../lib/ui/primitives';
  import Disclosure from '../ui/Disclosure.svelte';
  import Row from '../ui/Row.svelte';
  import Section from '../ui/Section.svelte';

  let {
    path,
    history,
    coach,
    fights,
    latest = null,
    inPanel = false,
  }: {
    path: CharacterPath;
    history: boolean;
    coach: boolean;
    fights: number;
    /** The newest score on file: the History summary names it ("Score 1,830 on 6 Nov"). */
    latest?: ScorePoint | null;
    /** Drawn inside the Best per encounter panel, under its fights, instead of a History panel of its own. */
    inPanel?: boolean;
  } = $props();

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

{#snippet body()}
  {#if history}
    <Disclosure
      summary={latest === null
        ? copy.historyLead
        : copy.historyLatest(latest.score, dayMonth(`${latest.date}T00:00:00Z`))}
      meta={copy.fightCount(fights)}
      bind:open={historyOpen}
      testid="character-history-toggle"
    >
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
  {#if coach}
    <div data-testid="character-coach-section">
      <Disclosure summary={copy.coachLead} bind:open={coachOpen} testid="character-coach-toggle">
        {#if reviews.current}<reviews.current {path} />{/if}
      </Disclosure>
    </div>
  {/if}
{/snippet}

<!-- In the fights panel (`inPanel`, h-audit 2026-11-09): under the best fights, no plate of its own; with no
     best fights, its own History panel. -->
{#if inPanel}
  <div class="border-line-soft mt-2 border-t pt-1" data-testid="character-owner-history">
    {@render body()}
  </div>
{:else}
  <Section title={copy.history} panel testid="character-owner-history">{@render body()}</Section>
{/if}
