<!-- web/src/components/coach/CoachHistoryList.svelte -->
<!-- The character page's coach history (experience spec 3.5, 4.3 and 10.3): the Coach tab's
     History link lands here (`#coach`). 56 px rows, newest first, each a link to that pull's Coach
     tab; twenty, then a 44 px `Show older` that reads the next page. Mounted only for the owner on
     Premium (Character.svelte); absent while the first page is read and with no history. -->
<script lang="ts">
  import { untrack } from 'svelte';
  import { fetchCoachHistory } from '../../lib/coach/api';
  import { COACH_SEPARATOR, coachCopy } from '../../lib/coach/copy';
  import {
    HISTORY_LOADING,
    historyRow,
    historyShown,
    withFailure,
    withOlderLoading,
    withPage,
    type HistoryListState,
  } from '../../lib/coach/history-list';
  import { scrollBehavior } from '../../lib/character/history-layout';
  import type { CharacterPath } from '../../lib/characters';
  import LoadError from '../ui/LoadError.svelte';

  let { path, apiBase = undefined }: { path: CharacterPath; apiBase?: string } = $props();

  /** The anchor the Coach tab's History link ends in. */
  const SECTION_ID = 'coach';

  let list = $state<HistoryListState>(HISTORY_LOADING);
  let section: HTMLElement | undefined = $state();
  const rows = $derived(list.items.map(historyRow));

  function read(cursor: string | undefined): void {
    void fetchCoachHistory(path.region, path.ruleset, path.slug, cursor, apiBase)
      .then((page) => {
        list = withPage(list, page);
      })
      .catch(() => {
        list = withFailure(list);
      });
  }

  function loadFirst(): void {
    list = HISTORY_LOADING;
    read(undefined);
  }

  function showOlder(): void {
    if (list.cursor === null || list.olderLoading) return;
    list = withOlderLoading(list);
    read(list.cursor);
  }

  $effect(() => {
    untrack(loadFirst);
  });

  // The section mounts after /v1/me, the page and the first history page: the browser's own
  // anchor jump has already missed it.
  let landed = false;
  $effect(() => {
    if (section === undefined || landed) return;
    landed = true;
    if (window.location.hash === `#${SECTION_ID}`) section.scrollIntoView({ behavior: scrollBehavior() });
  });

  const ROW =
    'border-line-soft hover:bg-card-top focus-visible:outline-gold group flex min-h-14 items-center gap-2 border-b px-2 py-2 text-[13px] leading-[18px] focus-visible:-outline-offset-2 focus-visible:outline-2';
  const TAG = 'border-line-soft text-muted rounded-control shrink-0 border px-1.5 text-[12px] leading-[18px]';
</script>

{#if historyShown(list)}
  <section bind:this={section} id={SECTION_ID} class="flex flex-col gap-2" data-testid="coach-history-list">
    <h2 class="section-title text-[18px]">{coachCopy.titleGeneric}</h2>
    {#if list.status === 'error'}
      <LoadError message={coachCopy.error} onRetry={loadFirst} testid="coach-history-error" />
    {:else}
      <ul class="border-line-soft flex flex-col border-t">
        {#each rows as row (row.key)}
          <li>
            <a href={row.href} class={ROW} data-testid="coach-history-row">
              <span class="min-w-0 flex-1">
                <span class="text-muted">{row.lead.join(COACH_SEPARATOR)}{COACH_SEPARATOR}</span><span
                  class="text-strong font-semibold group-hover:underline group-hover:underline-offset-[3px]"
                  data-testid="coach-history-title">{row.title}</span
                >
              </span>
              {#if row.wipe}<span class={TAG} data-testid="coach-history-wipe">{coachCopy.historyWipe}</span
                >{/if}
              {#if row.logOnly}<span class={TAG} data-testid="coach-history-rules"
                  >{coachCopy.historyRules}</span
                >{/if}
              <span class="text-muted pl-1" aria-hidden="true">›</span>
            </a>
          </li>
        {/each}
      </ul>
      {#if list.cursor !== null}
        <button
          type="button"
          class="text-text hover:text-strong focus-visible:outline-gold disabled:text-muted inline-flex min-h-11 items-center self-start px-2 text-[13px] underline underline-offset-[3px] focus-visible:outline-2"
          disabled={list.olderLoading}
          aria-busy={list.olderLoading}
          onclick={showOlder}
          data-testid="coach-history-older">{coachCopy.historyOlder}</button
        >
      {/if}
      {#if list.olderFailed}
        <p class="text-muted text-[13px]" role="alert" data-testid="coach-history-older-error">
          {coachCopy.error}
        </p>
      {/if}
    {/if}
  </section>
{/if}
