<!-- web/src/components/coach/CoachHistoryList.svelte -->
<!-- The character page's coach history (experience spec 3.5, 4.3 and 10.3): the Coach tab's
     History link lands here (`#coach`), inside the page's `Coach history` Disclosure. Rows,
     newest first (the pull's facts muted under its title, Wipe and Log only as Tags), each a link to
     that pull's Coach tab; twenty, then a ghost `Show older` that reads the next page. Mounted only for the owner on
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
  import { ROW_ITEM } from '../../lib/ui/primitives';
  import LoadError from '../ui/LoadError.svelte';
  import LoadOlder from '../ui/LoadOlder.svelte';
  import Row from '../ui/Row.svelte';
  import Tag from '../ui/Tag.svelte';

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
</script>

{#if historyShown(list)}
  <!-- It opens inside the page's `Coach history` Section, so it names itself only for assistive tech. -->
  <section
    bind:this={section}
    id={SECTION_ID}
    class="flex scroll-mt-[calc(var(--bar-h)+12px)] flex-col"
    aria-label={coachCopy.titleGeneric}
    data-testid="coach-history-list"
  >
    {#if list.status === 'error'}
      <LoadError message={coachCopy.error} onRetry={loadFirst} testid="coach-history-error" />
    {:else}
      <ul class="flex flex-col">
        {#each rows as row (row.key)}
          <li class={ROW_ITEM}>
            <Row
              title={row.title}
              sub={row.lead.join(COACH_SEPARATOR)}
              href={row.href}
              open
              testid="coach-history-row"
            >
              {#snippet titleContent()}<span data-testid="coach-history-title">{row.title}</span>{/snippet}
              {#snippet trail()}
                {#if row.wipe}<Tag tone="danger" testid="coach-history-wipe">{coachCopy.historyWipe}</Tag
                  >{/if}
                {#if row.logOnly}<Tag
                    tone="muted"
                    title={coachCopy.historyRulesTitle}
                    testid="coach-history-rules">{coachCopy.historyRules}</Tag
                  >{/if}
                <span class="text-muted" aria-hidden="true">›</span>
              {/snippet}
            </Row>
          </li>
        {/each}
      </ul>
      {#if list.cursor !== null}
        <LoadOlder
          label={coachCopy.historyOlder}
          loading={list.olderLoading}
          failed={list.olderFailed}
          failedText={coachCopy.historyOlderError}
          onLoad={showOlder}
          testid="coach-history-older"
        />
      {/if}
    {/if}
  </section>
{/if}
