<!-- web/src/components/character/HistoryFights.svelte -->
<!-- One fights list of the History section (experience spec 2026-10-24-character-history 3.7):
     the shipped fight table's rows without its heading, eight at first, `Show all N` for up to
     a hundred, the selected point's rows marked. Selecting a point whose rows are hidden opens
     the list and scrolls the first marked row into view. -->
<script lang="ts">
  import { tick, untrack } from 'svelte';
  import { historyCopy, type ListKind } from '../../lib/character/history-copy';
  import type { HistoryFight } from '../../lib/character/history-api';
  import { afterLayout, opensOnSelection, scrollRowIntoView } from '../../lib/character/history-layout';
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

  // A selection that marks hidden rows opens the list and jumps the first of them into view, once.
  // Only a committed selection reaches `marked` (a hover preview never does), so `Show fewer`
  // stays collapsed until the next selection. Keyed on the marked rows' identity, not the Set's:
  // a rebuilt chart hands over an equal Set and must not re-open a collapsed list.
  const markedSignature = $derived([...marked].sort().join(','));
  let scrolledFor = '';
  $effect(() => {
    const signature = markedSignature;
    const opens = untrack(() => listWindow(fights, false, marked).needsExpand);
    const previous = scrolledFor;
    scrolledFor = signature;
    if (!opensOnSelection(previous, signature, opens)) return;
    expanded = true;
    void afterLayout(() => scrollRowIntoView(host), {
      tick,
      frame: (run) => requestAnimationFrame(run),
    });
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
