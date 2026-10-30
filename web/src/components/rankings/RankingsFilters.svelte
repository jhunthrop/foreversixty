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
  import { selectCopy } from '../../lib/ui/copy';
  import type { SelectOption } from '../../lib/ui/listbox-model';
  import Select from '../ui/Select.svelte';

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

  /** One filter field: its label over a 44 px Select; `grow` lets the Encounter run wider,
   *  `ownRow` gives it a row of its own below the wide layout (rankings spec section 5). */
  interface Field {
    id: string;
    label: string;
    value: string;
    options: readonly SelectOption[];
    testid: string;
    onpick: (value: string) => void;
    grow?: number;
    ownRow?: boolean;
    disabled?: boolean;
    searchable?: boolean;
  }
  /** An encounter list longer than this gets the filter field (§A10 row 14). */
  const SEARCHABLE_OVER = 15;

  let open = $state(false);
  const guildKinds = GUILD_KINDS.map((kind) => ({ value: kind.id, label: kind.label }));
</script>

{#snippet field(f: Field)}
  <div
    class="flex min-w-0 flex-col gap-1 {f.ownRow ? 'col-span-full min-[1100px]:col-auto' : ''}"
    style:flex={`${f.grow ?? 1} 1 0`}
    data-testid="{f.testid}-field"
  >
    <label class="quiet-label" for={f.id}>{f.label}</label>
    <Select
      id={f.id}
      size="compact"
      value={f.value}
      options={f.options}
      disabled={f.disabled}
      searchable={f.searchable}
      searchPlaceholder={selectCopy.filterEncounters}
      testid={f.testid}
      onchange={(event) => f.onpick((event.currentTarget as HTMLSelectElement).value)}
    />
  </div>
{/snippet}

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
      {@render field({
        id: 'rankings-class',
        label: rankingsCopy.classLabel,
        value: filters.class,
        options: classOptions(),
        testid: 'filter-class',
        onpick: (value) => onpatch({ class: value, spec: '' }),
      })}
      {@render field({
        id: 'rankings-spec',
        label: rankingsCopy.specLabel,
        value: filters.spec,
        options: specOptions(filters.class),
        testid: 'filter-spec',
        onpick: (value) => onpatch({ spec: value }),
      })}
    {/if}
    {@render field({
      id: 'rankings-encounter',
      label: rankingsCopy.encounterLabel,
      value: noEncounters ? '' : encounterSlug,
      options: noEncounters
        ? [{ value: '', label: encounterPickerCopy.noEncountersOption }]
        : encounterOptions(encounters, encounterSlug),
      disabled: noEncounters,
      testid: 'filter-encounter',
      grow: 1.3,
      ownRow: !showClassAndSpec,
      searchable: encounters.length > SEARCHABLE_OVER,
      onpick: onencounter,
    })}
    {@render field({
      id: 'rankings-ruleset',
      label: rankingsCopy.rulesetLabel,
      value: filters.ruleset,
      options: rulesetOptions(),
      testid: 'filter-ruleset',
      onpick: (value) => onpatch({ ruleset: value }),
    })}
    {@render field({
      id: 'rankings-region',
      label: rankingsCopy.regionLabel,
      value: filters.region,
      options: regionOptions(),
      testid: 'filter-region',
      onpick: (value) => onpatch({ region: value }),
    })}
    {@render field({
      id: 'rankings-phase',
      label: rankingsCopy.phaseLabel,
      value: filters.phase,
      options: phaseOptions(),
      testid: 'filter-phase',
      onpick: (value) => onpatch({ phase: value }),
    })}
    {@render field({
      id: 'rankings-faction',
      label: rankingsCopy.factionLabel,
      value: filters.faction,
      options: factionOptions(),
      testid: 'filter-faction',
      onpick: (value) => onpatch({ faction: value }),
    })}
    {#if filters.board === 'character'}
      {@render field({
        id: 'rankings-metric',
        label: rankingsCopy.metricLabel,
        value: filters.metric,
        options: metricOptions(),
        testid: 'filter-metric',
        onpick: (value) => onpatch({ metric: value }),
      })}
    {:else}
      {@render field({
        id: 'rankings-kind',
        label: rankingsCopy.boardLabel,
        value: filters.kind,
        options: guildKinds,
        testid: 'filter-kind',
        onpick: (value) => onpatch({ kind: value }),
      })}
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
