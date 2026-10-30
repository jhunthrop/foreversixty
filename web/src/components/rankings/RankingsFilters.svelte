<!-- web/src/components/rankings/RankingsFilters.svelte -->
<!-- The one filter panel (design/specs/2026-10-11-rankings.md section 3, decisions 2 and 8):
     Encounter first and widest, then Ruleset, Region, Phase, Faction, Metric and Today only;
     Class and Spec lead the row only where the page is not aimed at a character (signed out, or
     after "Another class or spec"). From 1100 px it is one row, from 600 px three columns, and
     on a phone it collapses to a one-line summary with a 44 px Filters button that opens the
     two-column panel. Every filter lands in the URL through `onpatch`. -->
<script lang="ts">
  import { encounterPickerCopy, rankingsCopy } from '../../lib/rankings/copy';
  import {
    classOptions,
    encounterOptions,
    factionOptions,
    metricOptions,
    phaseOptions,
    regionOptions,
    rulesetOptions,
    specOptions,
  } from '../../lib/rankings/options';
  import type { EncounterOption } from '../../lib/rankings/api';
  import { GUILD_KINDS, type RankingsState } from '../../lib/rankings/url';
  import { SECONDARY_BUTTON_FIXED } from '../../lib/planner/styles';
  import Switch from '../ui/Switch.svelte';
  import FilterSelect from './FilterSelect.svelte';

  let {
    filters,
    encounters,
    encounterSlug,
    summary,
    showClassAndSpec,
    noEncounters,
    onpatch,
    onencounter,
  }: {
    filters: RankingsState;
    encounters: readonly EncounterOption[];
    encounterSlug: string;
    /** The phone's one-line collapsed form of these filters (lib/rankings/scope.ts). */
    summary: string;
    showClassAndSpec: boolean;
    /** The encounter list has loaded and is empty, and none is named: the select is disabled. */
    noEncounters: boolean;
    onpatch: (next: Partial<RankingsState>) => void;
    onencounter: (slug: string) => void;
  } = $props();

  let open = $state(false);
  const guildKinds = GUILD_KINDS.map((kind) => ({ value: kind.id, label: kind.label }));
</script>

<div class="flex flex-col gap-2" data-testid="rankings-filters">
  <!-- Phone only: the summary wraps rather than truncating, and the button opens the panel. -->
  <div
    class="bg-raised border-line rounded-panel flex min-h-[52px] items-center gap-2.5 border py-1 pr-1.5 pl-3.5 min-[600px]:hidden"
    data-testid="rankings-summary"
  >
    <span class="text-text min-w-0 flex-1 py-2 text-[13px] leading-[1.35]" data-testid="rankings-summary-line"
      >{summary}</span
    >
    <button
      type="button"
      class="{SECONDARY_BUTTON_FIXED} border-line-warm text-text gap-1.5 px-3.5"
      aria-expanded={open}
      aria-controls="rankings-filter-panel"
      data-testid="rankings-filters-toggle"
      onclick={() => (open = !open)}
    >
      {rankingsCopy.filtersButton}
      <svg
        class={open ? 'rotate-180' : ''}
        width="12"
        height="12"
        viewBox="0 0 24 24"
        fill="none"
        stroke="currentColor"
        stroke-width="3"
        stroke-linecap="round"
        aria-hidden="true"><path d="M6 9l6 6 6-6" /></svg
      >
    </button>
  </div>

  <div
    id="rankings-filter-panel"
    class="grid-cols-2 items-end gap-2.5 min-[600px]:grid min-[600px]:grid-cols-3 min-[1100px]:flex min-[1100px]:gap-3 {open
      ? 'grid'
      : 'hidden'}"
    data-testid="rankings-filter-panel"
  >
    {#if showClassAndSpec}
      <FilterSelect
        id="rankings-class"
        label="Class"
        value={filters.class}
        options={classOptions()}
        testid="filter-class"
        onchange={(value) => onpatch({ class: value, spec: '' })}
      />
      <FilterSelect
        id="rankings-spec"
        label="Spec"
        value={filters.spec}
        options={specOptions(filters.class)}
        testid="filter-spec"
        onchange={(value) => onpatch({ spec: value })}
      />
    {/if}
    <FilterSelect
      id="rankings-encounter"
      label={rankingsCopy.encounterLabel}
      value={noEncounters ? '' : encounterSlug}
      options={noEncounters
        ? [{ value: '', label: encounterPickerCopy.noEncountersOption }]
        : encounterOptions(encounters, encounterSlug)}
      disabled={noEncounters}
      testid="filter-encounter"
      grow={1.3}
      ownRow={!showClassAndSpec}
      onchange={onencounter}
    />
    <FilterSelect
      id="rankings-ruleset"
      label="Ruleset"
      value={filters.ruleset}
      options={rulesetOptions()}
      testid="filter-ruleset"
      onchange={(value) => onpatch({ ruleset: value })}
    />
    <FilterSelect
      id="rankings-region"
      label="Region"
      value={filters.region}
      options={regionOptions()}
      testid="filter-region"
      onchange={(value) => onpatch({ region: value })}
    />
    <FilterSelect
      id="rankings-phase"
      label="Phase"
      value={filters.phase}
      options={phaseOptions()}
      testid="filter-phase"
      onchange={(value) => onpatch({ phase: value })}
    />
    <FilterSelect
      id="rankings-faction"
      label="Faction"
      value={filters.faction}
      options={factionOptions()}
      testid="filter-faction"
      onchange={(value) => onpatch({ faction: value })}
    />
    {#if filters.board === 'character'}
      <FilterSelect
        id="rankings-metric"
        label="Metric"
        value={filters.metric}
        options={metricOptions()}
        testid="filter-metric"
        onchange={(value) => onpatch({ metric: value })}
      />
    {:else}
      <FilterSelect
        id="rankings-kind"
        label="Board"
        value={filters.kind}
        options={guildKinds}
        testid="filter-kind"
        onchange={(value) => onpatch({ kind: value })}
      />
    {/if}
    <div class="flex min-h-11 items-center gap-3">
      <span id="rankings-today-label" class="text-text text-[13px]">{rankingsCopy.todayOnly}</span>
      <Switch
        checked={filters.since === 'today'}
        labelledby="rankings-today-label"
        testid="filter-today"
        onchange={(next) => onpatch({ since: next ? 'today' : '' })}
      />
    </div>
  </div>
</div>
