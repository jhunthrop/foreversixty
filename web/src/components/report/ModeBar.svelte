<!-- web/src/components/report/ModeBar.svelte -->
<!-- Modes, then views, then source scope, then the twelve tabs -- the order spec section 4
     sets, which is also Warcraft Logs' order, so nobody relearns it. Replay is not rendered
     until it exists (the mode stays in url.ts): a "later" chip on the tab row read as a
     roadmap pill on the one page meant to impress.

     Over the whole night the mode row stays up rather than hiding: Analyze's tables and
     Mechanics both have a meaning across pulls, and hiding the row hid the way back out
     of Mechanics. Compare and Rankings are one pull's, so they are disabled there. -->
<script lang="ts">
  import { unitLabel } from '../../lib/report/names.svelte';
  import { selectCopy } from '../../lib/ui/copy';
  import type { SelectOption } from '../../lib/ui/listbox-model';
  import { classOption } from '../../lib/ui/select-options';
  import MoreMenu from './MoreMenu.svelte';
  import Select from '../ui/Select.svelte';
  import { reportCopy } from '../../lib/report/copy';
  import { createCoachViewerState } from '../../lib/coach/viewer-state.svelte';
  import { coachTabListed, ownedKey } from '../../lib/coach/viewer';
  import {
    FIRST_TIER_TABS,
    MODES,
    SOURCE_ENEMIES,
    SOURCE_FRIENDLIES,
    TABS,
    VIEWS,
    type Mode,
    type ModeOption,
    type ReportState,
    type Tab,
    type View,
  } from '../../lib/report/url';

  // Destructured under a different local name than the prop: this component also needs
  // Svelte's own `$state` rune (for the tab row's fade below), and a local
  // binding literally named `state` makes every `$state` call after it ambiguous with a
  // store subscription of that binding -- Svelte warns, and the rune stops working.
  let {
    state: reportState,
    roster,
    onPatch,
    nightMode = false,
  }: {
    state: ReportState;
    roster: { guid: string; name: string; class?: string; role?: string }[];
    onPatch: (patch: Partial<ReportState>) => void;
    /** The whole night: only Analyze's tables and Mechanics have a meaning over it. */
    nightMode?: boolean;
  } = $props();

  const pill =
    'inline-flex min-h-11 items-center px-3 text-[12px] font-bold tracking-[0.06em] uppercase md:min-h-9';
  // The selected mode and tab carry a gold rule under the word: text-strong against
  // text-nav is a two-step difference in grey, which is not enough to find the active
  // tab in a row of twelve without reading every one.
  const underline = 'border-b-2 -mb-px';

  /** A mode that is built, and that means something over whatever is selected. */
  function available(option: ModeOption): boolean {
    if (!option.enabled) return false;
    return !nightMode || option.id === 'analyze' || option.id === 'mechanics';
  }

  // Desktop's category row promotes the six tabs raiders reach for first (design review
  // 2026-09-26 finding 3) and folds the rest under a "More" button: TABS stays the single
  // source for every id and label, this just splits it in two without copying either.
  const firstTierTabs = FIRST_TIER_TABS.map((id) => TABS.find((tab) => tab.id === id)!);
  const moreTabIds = new Set<Tab>(FIRST_TIER_TABS);
  // The Coach tab is listed only for a signed-in owner of a character in the report, or an officer.
  const coachViewer = createCoachViewerState();
  const coachListed = $derived(coachTabListed(coachViewer.viewer, roster));
  const listedTabs = $derived(TABS.filter((tab) => tab.id !== 'coach' || coachListed));
  const moreTabs = $derived(listedTabs.filter((tab) => !moreTabIds.has(tab.id)));
  const moreActive = $derived(moreTabs.some((tab) => tab.id === reportState.tab));
  const activeMoreLabel = $derived(
    moreTabs.find((tab) => tab.id === reportState.tab)?.label ?? reportCopy.more,
  );

  /** A list longer than this gets the filter field (§A10 row 1). */
  const SEARCHABLE_OVER = 15;
  /** The roster's roles as the Source list groups them, in this order (§A10 row 1). */
  const ROLE_GROUPS: readonly { role: string; group: string }[] = [
    { role: 'tank', group: reportCopy.sourceGroups.tank },
    { role: 'healer', group: reportCopy.sourceGroups.healer },
    { role: 'dps', group: reportCopy.sourceGroups.dps },
  ];
  const groupOf = (role: string | undefined): string =>
    ROLE_GROUPS.find((entry) => entry.role === role)?.group ?? reportCopy.sourceGroups.dps;
  const sourceOptions = $derived.by((): SelectOption[] => {
    const players = ROLE_GROUPS.flatMap(({ group }) =>
      roster
        .filter((unit) => groupOf(unit.role) === group)
        .map((unit) =>
          classOption({
            value: unit.guid,
            label: unitLabel(unit.name).text,
            className: unit.class,
            group,
            you: coachViewer.viewer.signedIn && ownedKey(coachViewer.viewer, unit.name) !== null,
          }),
        ),
    );
    return [
      { value: SOURCE_FRIENDLIES, label: reportCopy.allFriendlies },
      { value: SOURCE_ENEMIES, label: reportCopy.allEnemies },
      ...players,
    ];
  });
  const tabOptions = $derived(listedTabs.map((tab) => ({ value: tab.id, label: tab.label })));

  let tabRow = $state<HTMLElement | null>(null);
  let fadeRight = $state(false);
  function measureFade(): void {
    if (tabRow === null) return;
    fadeRight = tabRow.scrollLeft + tabRow.clientWidth < tabRow.scrollWidth - 1;
  }
  $effect(() => {
    if (tabRow === null) return;
    measureFade();
    const watch = typeof ResizeObserver === 'undefined' ? null : new ResizeObserver(measureFade);
    watch?.observe(tabRow);
    return () => watch?.disconnect();
  });

  function selectTab(tab: Tab): void {
    onPatch({ tab });
  }
</script>

<div class="flex flex-col gap-3" data-testid="mode-bar">
  <!-- On a phone the row scrolls sideways; a 24px fade on the right edge says there is more
       while it does, and goes once the last tab is in view. -->
  <div class="relative">
    <div
      bind:this={tabRow}
      onscroll={measureFade}
      role="tablist"
      aria-label="Mode"
      class="border-line-soft flex flex-nowrap items-center overflow-x-auto border-b"
    >
      {#each MODES as option (option.id)}
        {#if option.id === 'replay'}
          <!-- Not drawn until Replay exists. -->
        {:else}
          <button
            type="button"
            role="tab"
            class="{pill} {underline} shrink-0 whitespace-nowrap"
            class:border-gold={reportState.mode === option.id}
            class:border-transparent={reportState.mode !== option.id}
            class:text-strong={available(option) && reportState.mode === option.id}
            class:text-nav={available(option) && reportState.mode !== option.id}
            class:text-muted={!available(option)}
            aria-selected={reportState.mode === option.id}
            disabled={!available(option)}
            data-testid={`mode-${option.id}`}
            onclick={() => available(option) && onPatch({ mode: option.id as Mode })}
          >
            {option.label}{#if !available(option)}<span class="text-muted ml-2 lowercase">one pull’s</span
              >{/if}
          </button>
        {/if}
      {/each}
    </div>
    {#if fadeRight}
      <div
        class="from-bg/0 to-bg pointer-events-none absolute inset-y-0 right-0 w-6 bg-gradient-to-r md:hidden"
        data-testid="mode-fade"
        aria-hidden="true"
      ></div>
    {/if}
  </div>

  <!-- Row 2: the view segmented control, the source scope and the table tabs, one flex-wrap
       group -- spec 2026-09-25 §6 collapses the old three rows to two. Each piece still
       renders only when it means something for the current mode, so a Compare/Rankings/
       Mechanics pull shows only its own Source select, and a night carries no view control
       at all (nightMode's own gate below). -->
  {#if (reportState.mode === 'analyze' && !nightMode) || (nightMode && reportState.mode !== 'mechanics') || (reportState.mode === 'analyze' && (reportState.view === 'tables' || nightMode))}
    <div class="flex flex-wrap items-center gap-2" data-testid="mode-bar-context">
      {#if reportState.mode === 'analyze' && !nightMode}
        <div role="tablist" aria-label="View" class="flex flex-wrap items-center gap-1">
          {#each VIEWS as option (option.id)}
            <button
              type="button"
              role="tab"
              class="{pill} rounded-control border"
              class:border-gold={reportState.view === option.id}
              class:bg-card-top={reportState.view === option.id}
              class:border-line-soft={reportState.view !== option.id}
              class:text-strong={reportState.view === option.id}
              class:text-nav={reportState.view !== option.id}
              aria-selected={reportState.view === option.id}
              data-testid={`view-${option.id}`}
              onclick={() => onPatch({ view: option.id as View })}
            >
              {option.label}
            </button>
          {/each}
        </div>
      {/if}

      {#if (reportState.mode === 'analyze' && !nightMode) || (nightMode && reportState.mode !== 'mechanics')}
        <div class="flex items-center gap-2">
          <label class="quiet-label" for="report-source">{reportCopy.source}</label>
          <Select
            id="report-source"
            class="w-[220px]"
            size="dense"
            value={reportState.source}
            options={sourceOptions}
            searchable={roster.length > SEARCHABLE_OVER}
            searchPlaceholder={selectCopy.filterPlayers}
            testid="source-scope"
            onchange={(event) => onPatch({ source: (event.currentTarget as HTMLSelectElement).value })}
          />
        </div>
      {/if}

      {#if reportState.mode === 'analyze' && (reportState.view === 'tables' || nightMode)}
        <!-- Below lg: a native select, one tap to any of the thirteen tabs. At lg and up:
             the existing wrapping pill row -- both read reportState.tab and both call the same
             onPatch, so the URL and the fight-switch races above never have to know which
             control fired. -->
        <!-- Visible, like Source's own label above: a bare select next to "Source" read as
             a second copy of it (finding 3) -- naming this one "Table" is the whole fix. -->
        <div class="flex items-center gap-2 lg:hidden">
          <label class="quiet-label" for="report-tab-select">{reportCopy.table}</label>
          <Select
            id="report-tab-select"
            class="w-[180px]"
            size="dense"
            value={reportState.tab}
            options={tabOptions}
            testid="tab-select"
            onchange={(event) => onPatch({ tab: (event.currentTarget as HTMLSelectElement).value as Tab })}
          />
        </div>
        <!-- The six tabs raiders reach for first (finding 3), stuck under the page header
             once scrolled past (finding 1/6) so a table switch never asks for a scroll back
             up. bg-bg keeps the rows sliding under it from showing through, and z-20 keeps them
             under it; the More menu itself is in the top layer and needs no z-index. -->
        <div
          role="tablist"
          aria-label="Table"
          class="border-line-soft bg-bg relative z-20 hidden flex-wrap border-b md:sticky md:top-0 lg:flex"
        >
          {#each firstTierTabs as tab (tab.id)}
            <button
              type="button"
              role="tab"
              class="{pill} {underline} shrink-0 whitespace-nowrap"
              class:border-gold={reportState.tab === tab.id}
              class:border-transparent={reportState.tab !== tab.id}
              class:text-strong={reportState.tab === tab.id}
              class:text-nav={reportState.tab !== tab.id}
              aria-selected={reportState.tab === tab.id}
              data-testid={`tab-${tab.id}`}
              onclick={() => selectTab(tab.id)}
            >
              {tab.label}
            </button>
          {/each}
          <!-- The rest, behind one button (finding 3), on the shared popover core: the menu
               renders in the top layer, so no chart or sticky row paints over it (§C). -->
          <MoreMenu
            tabs={moreTabs}
            current={reportState.tab}
            label={activeMoreLabel}
            menuLabel={reportCopy.more}
            active={moreActive}
            buttonClass="{pill} {underline} shrink-0 whitespace-nowrap"
            onselect={selectTab}
          />
        </div>
      {/if}
    </div>
  {/if}
</div>
