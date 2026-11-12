<!-- web/src/components/report/ReportFilterPopover.svelte -->
<!-- The report bar's Filters (logs spec 2026-11-07 6.1) on the shared FilterPopover: one button
     from 768 px ("Filters", "Filters (2)" while any switch is on) that opens the switches.
     Below 768 the same switches sit in the open control stack. Boss damage only, Players only,
     Count overkill, Ignore events while dead, and Clear filters. -->
<script lang="ts">
  import { reportCopy } from '../../lib/report/copy';
  import { DEFAULT_FILTERS, type ReportFilters } from '../../lib/report/filters';
  import FilterPopover from '../ui/FilterPopover.svelte';

  let {
    filters,
    showBossOnly,
    afterDeathAvailable,
    onChange,
  }: {
    filters: ReportFilters;
    /** False on Healing, where "boss damage only" has nothing to apply to. */
    showBossOnly: boolean;
    /** False over the whole night, where a death's span cannot be read across pulls. */
    afterDeathAvailable: boolean;
    onChange: (filters: ReportFilters) => void;
  } = $props();

  type Switch = 'bossOnly' | 'playersOnly' | 'countOverkill' | 'ignoreAfterDeath';
  const SWITCHES: readonly { key: Switch; label: string }[] = [
    { key: 'bossOnly', label: 'Boss damage only' },
    { key: 'playersOnly', label: 'Players only' },
    { key: 'countOverkill', label: 'Count overkill' },
    { key: 'ignoreAfterDeath', label: 'Ignore events while dead' },
  ];
  const shownSwitches = $derived(SWITCHES.filter((item) => item.key !== 'bossOnly' || showBossOnly));
  const onCount = $derived(shownSwitches.filter((item) => filters[item.key]).length);
  const label = $derived(onCount === 0 ? reportCopy.filters : `${reportCopy.filters} (${onCount})`);
</script>

{#snippet switches(prefix: string)}
  {#each shownSwitches as item (item.key)}
    {@const disabled = item.key === 'ignoreAfterDeath' && !afterDeathAvailable}
    <label class="flex min-h-11 items-center gap-2 text-[14px] md:min-h-9">
      <input
        type="checkbox"
        checked={filters[item.key] && !disabled}
        {disabled}
        title={disabled ? 'Over the whole night a death’s span cannot be read; open a pull' : undefined}
        data-testid={item.key === 'bossOnly' ? `${prefix}filter-boss` : undefined}
        onchange={(event) =>
          onChange({ ...filters, [item.key]: (event.currentTarget as HTMLInputElement).checked })}
      />
      {item.label}
    </label>
  {/each}
  <button
    type="button"
    class="text-gold min-h-11 self-start text-[14px] font-semibold md:min-h-9"
    onclick={() => onChange({ ...DEFAULT_FILTERS, target: filters.target, ability: filters.ability })}
    >{reportCopy.clearFilters}</button
  >
  <p class="text-small text-muted">
    A pet’s damage and healing count on its owner’s row, as the engine records it. Split it in Queries.
  </p>
{/snippet}

<FilterPopover
  label={reportCopy.filters}
  buttonLabel={label}
  id="report-filters-panel"
  testid="filter-bar"
  buttonTestid="report-filters-button"
  popoverTestid="report-filters-popover"
  focusSelector="input"
>
  {#snippet panel()}<div class="flex flex-col gap-1 p-4">{@render switches('')}</div>{/snippet}
  {#snippet stack()}<div class="flex flex-col">{@render switches('stack-')}</div>{/snippet}
</FilterPopover>
