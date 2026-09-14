<!-- web/src/components/report/ModeBar.svelte -->
<!-- Modes, then views, then source scope, then the twelve tabs -- the order spec section 4
     sets, which is also Warcraft Logs' order, so nobody relearns it. Mechanics and Replay
     are rendered disabled with "later" rather than hidden: the spec defers them, and a
     missing control reads as a missing feature. -->
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
    type ReportState,
    type Tab,
    type View,
  } from '../../lib/report/url';

  let {
    state,
    roster,
    onPatch,
  }: {
    state: ReportState;
    roster: { guid: string; name: string; class?: string }[];
    onPatch: (patch: Partial<ReportState>) => void;
  } = $props();

  const pill =
    'inline-flex min-h-11 items-center px-3 text-[12px] font-bold tracking-[0.06em] uppercase md:min-h-9';
</script>

<div class="flex flex-col gap-3" data-testid="mode-bar">
  <div role="tablist" aria-label="Mode" class="border-line-soft flex flex-wrap items-center border-b">
    {#each MODES as option (option.id)}
      <button
        type="button"
        role="tab"
        class={pill}
        class:text-strong={option.enabled && state.mode === option.id}
        class:text-nav={option.enabled && state.mode !== option.id}
        class:text-muted={!option.enabled}
        aria-selected={state.mode === option.id}
        disabled={!option.enabled}
        data-testid={`mode-${option.id}`}
        onclick={() => option.enabled && onPatch({ mode: option.id as Mode })}
      >
        {option.label}{#if option.note}<span class="text-muted ml-2 lowercase">{option.note}</span>{/if}
      </button>
    {/each}
  </div>

  {#if state.mode === 'analyze'}
    <div role="tablist" aria-label="View" class="flex flex-wrap items-center gap-1">
      {#each VIEWS as option (option.id)}
        <button
          type="button"
          role="tab"
          class="{pill} rounded-control border"
          class:border-line-warm-strong={state.view === option.id}
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

  {#if state.mode === 'analyze' && state.view === 'tables'}
    <div
      role="tablist"
      aria-label="Table"
      class="flex flex-nowrap overflow-x-auto md:flex-wrap"
    >
      {#each TABS as tab (tab.id)}
        <button
          type="button"
          role="tab"
          class="{pill} shrink-0 whitespace-nowrap"
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
