<!-- web/src/components/report/FilterBar.svelte -->
<script lang="ts">
  import {
    DEFAULT_FILTERS,
    abilityOptions,
    targetOptions,
    type ReportFilters,
  } from '../../lib/report/filters';
  import { splitUnitName } from '../../lib/characters';
  import type { Actor } from '../../lib/report/types';

  let {
    filters,
    actors,
    onChange,
  }: { filters: ReportFilters; actors: Actor[]; onChange: (filters: ReportFilters) => void } = $props();

  const abilities = $derived(abilityOptions(actors));
  const targets = $derived(targetOptions(actors));
  const select = 'border-line-warm bg-raised rounded-control text-text h-11 px-2 text-[13px] md:h-9';
  const toggle = 'flex min-h-11 items-center gap-2 text-[13px] md:min-h-0';
  // The 44px hit target for a checkbox is the label around it, not the box: a click
  // anywhere in the label toggles the input, and `toggle` above gives every one of them
  // min-h-11 on phone. The box keeps its native size and gains only the gold accent the
  // chart's own range inputs use. (A range input is the other case -- it is dragged, so
  // its own box has to be the handle; TimeChart.svelte carries that reasoning.)
  const check = 'accent-gold';
</script>

<div class="flex flex-wrap items-center gap-x-4 gap-y-2" data-testid="filter-bar">
  <label class="label text-muted flex items-center gap-2" for="filter-target">
    Target
    <select
      id="filter-target"
      class={select}
      value={filters.target}
      onchange={(event) => onChange({ ...filters, target: (event.currentTarget as HTMLSelectElement).value })}
    >
      <option value="">Every target</option>
      {#each targets as option (option.id)}
        <option value={option.id}>{splitUnitName(option.name).name}</option>
      {/each}
    </select>
  </label>

  <label class="label text-muted flex items-center gap-2" for="filter-ability">
    Ability
    <select
      id="filter-ability"
      class={select}
      value={filters.ability === null ? '' : String(filters.ability)}
      data-testid="filter-ability"
      onchange={(event) => {
        const raw = (event.currentTarget as HTMLSelectElement).value;
        onChange({ ...filters, ability: raw === '' ? null : Number(raw) });
      }}
    >
      <option value="">Every ability</option>
      {#each abilities as option (option.id)}
        <option value={option.id}>{option.name}</option>
      {/each}
    </select>
  </label>

  <label class={toggle}>
    <input
      type="checkbox"
      class={check}
      checked={filters.bossOnly}
      data-testid="filter-boss"
      onchange={(event) =>
        onChange({ ...filters, bossOnly: (event.currentTarget as HTMLInputElement).checked })}
    />
    Boss damage only
  </label>
  <label class={toggle}>
    <input
      type="checkbox"
      class={check}
      checked={filters.playersOnly}
      onchange={(event) =>
        onChange({ ...filters, playersOnly: (event.currentTarget as HTMLInputElement).checked })}
    />
    Players only
  </label>
  <label class={toggle}>
    <input
      type="checkbox"
      class={check}
      checked={filters.countOverkill}
      onchange={(event) =>
        onChange({ ...filters, countOverkill: (event.currentTarget as HTMLInputElement).checked })}
    />
    Count overkill
  </label>
  <label class={toggle}>
    <input
      type="checkbox"
      class={check}
      checked={filters.ignoreAfterDeath}
      onchange={(event) =>
        onChange({ ...filters, ignoreAfterDeath: (event.currentTarget as HTMLInputElement).checked })}
    />
    Ignore events after a death
  </label>

  <button
    type="button"
    class="text-nav inline-flex min-h-11 items-center text-[12px] font-bold tracking-[0.06em] uppercase md:min-h-0"
    onclick={() => onChange(DEFAULT_FILTERS)}
  >
    Clear filters
  </button>

  <p class="text-muted basis-full text-[12px]">
    Pet damage is counted on its owner's row, as the engine records it. Split it in Queries.
  </p>
</div>
