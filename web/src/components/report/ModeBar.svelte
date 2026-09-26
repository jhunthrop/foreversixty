<!-- web/src/components/report/ModeBar.svelte -->
<!-- Modes, then views, then source scope, then the twelve tabs -- the order spec section 4
     sets, which is also Warcraft Logs' order, so nobody relearns it. Replay is rendered
     disabled with "later" rather than hidden: the spec defers it, and a missing control
     reads as a missing feature.

     Over the whole night the mode row stays up rather than hiding: Analyze's tables and
     Mechanics both have a meaning across pulls, and hiding the row hid the way back out
     of Mechanics. Compare and Rankings are one pull's, so they are disabled there. -->
<script lang="ts">
  import { splitUnitName } from '../../lib/characters';
  import { classColorVar } from '../../lib/report/format';
  import { reportCopy } from '../../lib/report/copy';
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
  // Svelte's own `$state` rune (for the "More" menu's open flag below), and a local
  // binding literally named `state` makes every `$state` call after it ambiguous with a
  // store subscription of that binding -- Svelte warns, and the rune stops working.
  let {
    state: reportState,
    roster,
    onPatch,
    nightMode = false,
  }: {
    state: ReportState;
    roster: { guid: string; name: string; class?: string }[];
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
  const moreTabs = TABS.filter((tab) => !moreTabIds.has(tab.id));
  const moreActive = $derived(moreTabs.some((tab) => tab.id === reportState.tab));
  const activeMoreLabel = $derived(
    moreTabs.find((tab) => tab.id === reportState.tab)?.label ?? reportCopy.more,
  );
  /** Closed once a tab is picked from it, the way a native select closes on its own. */
  let moreOpen = $state(false);

  function selectTab(tab: Tab): void {
    onPatch({ tab });
    moreOpen = false;
  }
</script>

<div class="flex flex-col gap-3" data-testid="mode-bar">
  <div
    role="tablist"
    aria-label="Mode"
    class="border-line-soft flex flex-nowrap items-center overflow-x-auto border-b"
  >
    {#each MODES as option (option.id)}
      {#if option.id === 'replay'}
        <!-- A trailing muted chip, not another tab: Replay is not built yet, and giving it
             the same underline treatment as the four real modes made it read as a fifth
             one that happened to be broken. shrink-0 keeps it from being the reason the
             row wraps on a phone -- the row scrolls sideways instead (finding 4). -->
        <button
          type="button"
          role="tab"
          class="{pill} border-line-soft bg-card-top text-muted ml-2 shrink-0 gap-1 rounded-full border px-3 whitespace-nowrap normal-case"
          aria-selected="false"
          disabled
          data-testid={`mode-${option.id}`}
        >
          {option.label}<span class="lowercase">{option.note}</span>
        </button>
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
          {option.label}{#if !available(option)}<span class="text-muted ml-2 lowercase">one pull’s</span>{/if}
        </button>
      {/if}
    {/each}
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
        <label class="text-muted label flex items-center gap-2" for="report-source">
          Source
          <select
            id="report-source"
            class="border-line-warm bg-raised rounded-control text-text h-11 px-2 text-[13px] md:h-9"
            value={reportState.source}
            data-testid="source-scope"
            onchange={(event) => onPatch({ source: event.currentTarget.value })}
          >
            <option value={SOURCE_FRIENDLIES}>All friendlies</option>
            <option value={SOURCE_ENEMIES}>All enemies</option>
            {#each roster as unit (unit.guid)}
              <option value={unit.guid} style={`color: ${classColorVar(unit.class)}`}>
                {splitUnitName(unit.name).name}
              </option>
            {/each}
          </select>
        </label>
      {/if}

      {#if reportState.mode === 'analyze' && (reportState.view === 'tables' || nightMode)}
        <!-- Below lg: a native select, one tap to any of the thirteen tabs. At lg and up:
             the existing wrapping pill row -- both read reportState.tab and both call the same
             onPatch, so the URL and the fight-switch races above never have to know which
             control fired. -->
        <!-- Visible, like Source's own label above: a bare select next to "Source" read as
             a second copy of it (finding 3) -- naming this one "Table" is the whole fix. -->
        <label class="text-muted label flex items-center gap-2 lg:hidden" for="report-tab-select">
          {reportCopy.table}
          <select
            id="report-tab-select"
            class="border-line-warm bg-raised rounded-control text-text h-11 px-2 text-[13px]"
            value={reportState.tab}
            data-testid="tab-select"
            onchange={(event) => onPatch({ tab: event.currentTarget.value as Tab })}
          >
            {#each TABS as tab (tab.id)}<option value={tab.id}>{tab.label}</option>{/each}
          </select>
        </label>
        <!-- The six tabs raiders reach for first (finding 3), stuck under the page header
             once scrolled past (finding 1/6) so a table switch never asks for a scroll back
             up. bg-bg keeps the rows sliding under it from showing through. -->
        <div
          role="tablist"
          aria-label="Table"
          class="border-line-soft bg-bg hidden flex-wrap border-b md:sticky md:top-0 lg:flex"
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
          <!-- The rest, behind one button (finding 3): a `<details>` needs no keyboard or
               click-outside plumbing of its own to open, close and be reachable by Tab. -->
          <details class="relative" bind:open={moreOpen} data-testid="tab-more-menu">
            <summary
              class="{pill} {underline} text-nav shrink-0 cursor-pointer list-none border-transparent whitespace-nowrap select-none [&::-webkit-details-marker]:hidden"
              class:border-gold={moreActive}
              class:text-strong={moreActive}
              data-testid="tab-more"
            >
              {activeMoreLabel}
            </summary>
            <div
              role="tablist"
              aria-label={reportCopy.more}
              class="border-line bg-raised absolute top-full left-0 z-10 flex min-w-[160px] flex-col border shadow-lg"
            >
              {#each moreTabs as tab (tab.id)}
                <button
                  type="button"
                  role="tab"
                  class="{pill} justify-start whitespace-nowrap"
                  class:text-strong={reportState.tab === tab.id}
                  class:text-nav={reportState.tab !== tab.id}
                  aria-selected={reportState.tab === tab.id}
                  data-testid={`tab-${tab.id}`}
                  onclick={() => selectTab(tab.id)}
                >
                  {tab.label}
                </button>
              {/each}
            </div>
          </details>
        </div>
      {/if}
    </div>
  {/if}
</div>
