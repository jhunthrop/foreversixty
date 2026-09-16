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
  import {
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

  let {
    state,
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
</script>

<div class="flex flex-col gap-3" data-testid="mode-bar">
  <div role="tablist" aria-label="Mode" class="border-line-soft flex flex-wrap items-center border-b">
    {#each MODES as option (option.id)}
      <button
        type="button"
        role="tab"
        class="{pill} {underline}"
        class:border-gold={state.mode === option.id}
        class:border-transparent={state.mode !== option.id}
        class:text-strong={available(option) && state.mode === option.id}
        class:text-nav={available(option) && state.mode !== option.id}
        class:text-muted={!available(option)}
        aria-selected={state.mode === option.id}
        disabled={!available(option)}
        data-testid={`mode-${option.id}`}
        onclick={() => available(option) && onPatch({ mode: option.id as Mode })}
      >
        {option.label}{#if option.note}<span class="text-muted ml-2 lowercase">{option.note}</span
          >{:else if !available(option)}<span class="text-muted ml-2 lowercase">one pull’s</span>{/if}
      </button>
    {/each}
  </div>

  {#if state.mode === 'analyze' && !nightMode}
    <!-- The tablist holds only tabs: the Source picker beside it is a sibling, since a
         label inside a tablist is a child the role does not allow. -->
    <div class="flex flex-wrap items-center gap-1">
      <div role="tablist" aria-label="View" class="flex flex-wrap items-center gap-1">
        {#each VIEWS as option (option.id)}
          <button
            type="button"
            role="tab"
            class="{pill} rounded-control border"
            class:border-gold={state.view === option.id}
            class:bg-card-top={state.view === option.id}
            class:border-line-soft={state.view !== option.id}
            class:text-strong={state.view === option.id}
            class:text-nav={state.view !== option.id}
            aria-selected={state.view === option.id}
            data-testid={`view-${option.id}`}
            onclick={() => onPatch({ view: option.id as View })}
          >
            {option.label}
          </button>
        {/each}
      </div>

      <label class="text-muted label ml-auto flex items-center gap-2" for="report-source">
        Source
        <select
          id="report-source"
          class="border-line-warm bg-raised rounded-control text-text h-11 px-2 text-[13px] md:h-9"
          value={state.source}
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
    </div>
  {/if}

  {#if nightMode && state.mode !== 'mechanics'}
    <label class="text-muted label flex items-center gap-2" for="report-source">
      Source
      <select
        id="report-source"
        class="border-line-warm bg-raised rounded-control text-text h-11 px-2 text-[13px] md:h-9"
        value={state.source}
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
  {#if state.mode === 'analyze' && (state.view === 'tables' || nightMode)}
    <div role="tablist" aria-label="Table" class="border-line-soft flex flex-wrap border-b">
      {#each TABS as tab (tab.id)}
        <button
          type="button"
          role="tab"
          class="{pill} {underline} shrink-0 whitespace-nowrap"
          class:border-gold={state.tab === tab.id}
          class:border-transparent={state.tab !== tab.id}
          class:text-strong={state.tab === tab.id}
          class:text-nav={state.tab !== tab.id}
          aria-selected={state.tab === tab.id}
          data-testid={`tab-${tab.id}`}
          onclick={() => onPatch({ tab: tab.id as Tab })}
        >
          {tab.label}
        </button>
      {/each}
    </div>
  {/if}
</div>
