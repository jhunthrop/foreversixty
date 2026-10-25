<!-- web/src/components/character/HistoryFights.svelte -->
<!-- One fights list of the History section (experience spec 2026-10-24-character-history 3.7):
     the shipped fight table's rows without its heading, eight at first, `Show all N` for up to
     a hundred, the selected point's rows marked. Selecting a point whose rows are hidden opens
     the list and scrolls the first marked row into view. -->
<script lang="ts">
  import { tick, untrack } from 'svelte';
  import { historyCopy, type ListKind } from '../../lib/character/history-copy';
  import type { HistoryFight } from '../../lib/character/history-api';
  import { scrollBehavior } from '../../lib/character/history-layout';
  import { listWindow } from '../../lib/character/history-select';
  import CharacterFightTable from './CharacterFightTable.svelte';

  let {
    fights,
    label,
    kind,
    marked,
    lead = 'boss',
    testid,
  }: {
    fights: readonly HistoryFight[];
    /** The list's own label: "Fights behind the score", "Kills of Golemagg". */
    label: string;
    kind: ListKind;
    /** `report_id:fight_index` keys the selected point marks. */
    marked: ReadonlySet<string>;
    lead?: 'boss' | 'date';
    testid: string;
  } = $props();

  let expanded = $state(false);
  let host: HTMLDivElement | undefined = $state();

  const view = $derived(listWindow(fights, expanded, marked));

  const reportHref = (row: HistoryFight): string => `/reports/${row.report_id}?fight=${row.fight_index}`;
  const rowKey = (row: HistoryFight, index: number): string => `${row.report_id}-${row.fight_index}-${index}`;
  const visibilityNote = (row: HistoryFight): string | null => historyCopy.visibilityWord(row.visibility);

  // A selection that marks hidden rows opens the list and brings the first of them into view.
  // Only the selection triggers it, so `Show fewer` stays collapsed until the next selection.
  // Keyed on the marked rows' own identity, not the Set's: a rebuilt chart or a hover that leaves
  // the committed point alone hands over an equal Set and must not re-open a collapsed list.
  const markedSignature = $derived([...marked].sort().join(','));
  $effect(() => {
    void markedSignature;
    const opens = untrack(() => listWindow(fights, false, marked).needsExpand);
    if (!opens) return;
    expanded = true;
    // After the opened rows are in the DOM and painted, so the scroll measures the final list.
    void tick().then(() =>
      requestAnimationFrame(() =>
        host
          ?.querySelector('[data-marked="true"]')
          ?.scrollIntoView({ block: 'nearest', behavior: scrollBehavior() }),
      ),
    );
  });
</script>

<div class="flex flex-col" data-testid={testid} bind:this={host}>
  <CharacterFightTable
    title={label}
    rows={view.rows}
    testid={`${testid}-rows`}
    bossHref={reportHref}
    {rowKey}
    executionTestid={`${testid}-execution`}
    heading={false}
    columns="no-spec"
    {lead}
    selected={marked}
    note={visibilityNote}
  />
  {#if view.button !== 'none'}
    <button
      type="button"
      class="text-text hover:text-gold-hover active:text-gold-hover inline-flex min-h-11 items-center self-start text-[13px] font-semibold underline underline-offset-4"
      aria-expanded={expanded}
      data-testid={`${testid}-toggle`}
      onclick={() => (expanded = !expanded)}
    >
      {view.button === 'show-all' ? historyCopy.showAll(view.shownTotal, kind) : historyCopy.showFewer}
    </button>
  {/if}
</div>
