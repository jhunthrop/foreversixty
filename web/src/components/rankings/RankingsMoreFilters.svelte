<!-- web/src/components/rankings/RankingsMoreFilters.svelte -->
<!-- The rankings bar's Filters group (logs spec 2026-11-07 section 7, decision K2) on the shared
     FilterPopover: Region, Phase, Faction, Today only and every Metric, behind one compact button
     from 768 px ("Filters", "Filters (2)" when any is set). Below 768 px the same fields sit in the
     expanded control stack. Every choice lands in the URL through `onpatch`. -->
<script lang="ts">
  import { factionOptions, metricOptions, phaseOptions, regionOptions } from '../../lib/rankings/options';
  import { rankingsCopy } from '../../lib/rankings/copy';
  import type { RankingsState } from '../../lib/rankings/url';
  import type { SelectOption } from '../../lib/ui/listbox-model';
  import FilterPopover from '../ui/FilterPopover.svelte';
  import Select from '../ui/Select.svelte';
  import Switch from '../ui/Switch.svelte';

  let {
    filters,
    defaultPhase = '',
    onpatch,
  }: {
    filters: Pick<RankingsState, 'region' | 'phase' | 'faction' | 'since' | 'metric'>;
    /** The phase the board starts on (the current one): only a different choice counts as set. */
    defaultPhase?: string;
    onpatch: (next: Partial<RankingsState>) => void;
  } = $props();

  interface Field {
    key: 'region' | 'phase' | 'faction' | 'metric';
    label: string;
    options: readonly SelectOption[];
  }
  const FIELDS: readonly Field[] = [
    { key: 'region', label: rankingsCopy.regionLabel, options: regionOptions() },
    { key: 'phase', label: rankingsCopy.phaseLabel, options: phaseOptions() },
    { key: 'faction', label: rankingsCopy.factionLabel, options: factionOptions() },
    // Every metric (Damage taken, Execution too); the bar's segmented control offers the two common ones.
    { key: 'metric', label: rankingsCopy.metricLabel, options: metricOptions() },
  ];
  const COUNTED: ReadonlySet<Field['key']> = new Set(['region', 'phase', 'faction']);

  const setCount = $derived(
    FIELDS.filter(
      (field) => COUNTED.has(field.key) && filters[field.key] !== (field.key === 'phase' ? defaultPhase : ''),
    ).length + (filters.since === '' ? 0 : 1),
  );
  const label = $derived(
    setCount === 0 ? rankingsCopy.filtersButton : `${rankingsCopy.filtersButton} (${setCount})`,
  );
</script>

{#snippet fields(prefix: string)}
  {#each FIELDS as field (field.key)}
    <div class="flex flex-col gap-1">
      <span class="text-label text-muted font-bold" id={`${prefix}-${field.key}-label`}>{field.label}</span>
      <Select
        id={`${prefix}-${field.key}`}
        label={field.label}
        size="dense"
        value={filters[field.key]}
        options={field.options}
        testid={`${prefix}-${field.key}`}
        onchange={(event) => onpatch({ [field.key]: (event.currentTarget as HTMLSelectElement).value })}
      />
    </div>
  {/each}
  <div class="flex items-center gap-3">
    <Switch
      labelledby={`${prefix}-today-label`}
      checked={filters.since !== ''}
      testid={`${prefix}-today`}
      onchange={(next) => onpatch({ since: next ? 'today' : '' })}
    />
    <span class="text-[14px] font-semibold" id={`${prefix}-today-label`}>{rankingsCopy.todayOnly}</span>
  </div>
{/snippet}

<FilterPopover
  label={rankingsCopy.filtersButton}
  buttonLabel={label}
  id="rankings-filters-panel"
  testid="rankings-more-filters"
  buttonTestid="rankings-filters-button"
  popoverTestid="rankings-filters-popover"
  focusSelector="button, select"
>
  {#snippet panel()}
    <div class="flex flex-col gap-3 p-4">
      {@render fields('rankings-pop')}
      <button
        type="button"
        class="text-gold self-start text-[14px] font-semibold"
        data-testid="rankings-filters-clear"
        onclick={() => onpatch({ region: '', phase: defaultPhase, faction: '', since: '' })}
        >Clear filters</button
      >
    </div>
  {/snippet}
  {#snippet stack()}{@render fields('rankings-stack')}{/snippet}
</FilterPopover>
