<!-- web/src/components/report/FilterBar.svelte -->
<script lang="ts">
  import {
    DEFAULT_FILTERS,
    abilityOptions,
    targetOptionId,
    targetOptions,
    type ReportFilters,
  } from '../../lib/report/filters';
  import activeBuild from '../../data/active-build.json';
  import { splitUnitName } from '../../lib/characters';
  import { loadSpellIconUrls, type IconUrl } from '../../lib/coach/names';
  import type { Actor } from '../../lib/report/types';
  import { selectCopy } from '../../lib/ui/copy';
  import { iconOption } from '../../lib/ui/select-options';
  import Select from '../ui/Select.svelte';
  import { reportCopy } from '../../lib/report/copy';

  let {
    filters,
    actors,
    onChange,
    showBossOnly = true,
    afterDeathAvailable = true,
  }: {
    filters: ReportFilters;
    actors: Actor[];
    onChange: (filters: ReportFilters) => void;
    /** False on the Healing tab, where "boss damage only" has nothing to apply to. */
    showBossOnly?: boolean;
    /** False over the whole night, where a death's span cannot be read across pulls. */
    afterDeathAvailable?: boolean;
  } = $props();

  const abilities = $derived(abilityOptions(actors));
  const targets = $derived(targetOptions(actors));
  /** The option the filter's value names, which over a night is an enemy's name, not a GUID. */
  const picked = $derived(targetOptionId(actors, filters.target));
  /** A target list longer than this gets the filter field (§A10 row 6). */
  const SEARCHABLE_OVER = 15;
  /** Spell icons for the Ability list, loaded once from the build's icon table (§A10 row 7). */
  let iconUrl = $state<IconUrl>(() => null);
  $effect(() => {
    void loadSpellIconUrls(activeBuild.build).then((urls) => {
      iconUrl = urls;
    });
  });
  const targetChoices = $derived([
    { value: '', label: reportCopy.everyTarget },
    ...targets.map((option) => ({ value: option.id, label: splitUnitName(option.name).name })),
  ]);
  const abilityChoices = $derived([
    { value: '', label: reportCopy.everyAbility },
    ...abilities.map((option) => iconOption(option.id, option.name, iconUrl(Number(option.id)), 'spell')),
  ]);
  const toggle = 'flex min-h-11 items-center gap-2 text-[13px] md:min-h-0';
  // The 44px hit target for a checkbox is the label around it, not the box: a click
  // anywhere in the label toggles the input, and `toggle` above gives every one of them
  // min-h-11 on phone. The box keeps its native size; its gold comes from global.css's
  // site-wide checkbox rule.
</script>

<div class="flex flex-wrap items-center gap-x-4 gap-y-2" data-testid="filter-bar">
  <div class="flex items-center gap-2">
    <label class="quiet-label" for="filter-target">{reportCopy.target}</label>
    <Select
      id="filter-target"
      class="w-[200px]"
      size="dense"
      value={picked}
      options={targetChoices}
      searchable={targets.length > SEARCHABLE_OVER}
      searchPlaceholder={selectCopy.filterTargets}
      testid="filter-target"
      onchange={(event) => onChange({ ...filters, target: (event.currentTarget as HTMLSelectElement).value })}
    />
  </div>

  <div class="flex items-center gap-2">
    <label class="quiet-label" for="filter-ability">{reportCopy.ability}</label>
    <Select
      id="filter-ability"
      class="w-[220px]"
      size="dense"
      value={filters.ability === null ? '' : String(filters.ability)}
      options={abilityChoices}
      searchable
      searchPlaceholder={selectCopy.filterAbilities}
      testid="filter-ability"
      onchange={(event) => {
        const raw = (event.currentTarget as HTMLSelectElement).value;
        onChange({ ...filters, ability: raw === '' ? null : Number(raw) });
      }}
    />
  </div>

  {#if showBossOnly}
    <label class={toggle}>
      <input
        type="checkbox"
        checked={filters.bossOnly}
        data-testid="filter-boss"
        onchange={(event) =>
          onChange({ ...filters, bossOnly: (event.currentTarget as HTMLInputElement).checked })}
      />
      Boss damage only
    </label>
  {/if}
  <label class={toggle}>
    <input
      type="checkbox"
      checked={filters.playersOnly}
      onchange={(event) =>
        onChange({ ...filters, playersOnly: (event.currentTarget as HTMLInputElement).checked })}
    />
    Players only
  </label>
  <label class={toggle}>
    <input
      type="checkbox"
      checked={filters.countOverkill}
      onchange={(event) =>
        onChange({ ...filters, countOverkill: (event.currentTarget as HTMLInputElement).checked })}
    />
    Count overkill
  </label>
  <label class={toggle}>
    <input
      type="checkbox"
      checked={filters.ignoreAfterDeath && afterDeathAvailable}
      disabled={!afterDeathAvailable}
      title={afterDeathAvailable
        ? 'Leave out what each player did, and took, while dead: from a death to their first cast after it'
        : 'Over the whole night a death’s span cannot be read; open a pull'}
      onchange={(event) =>
        onChange({ ...filters, ignoreAfterDeath: (event.currentTarget as HTMLInputElement).checked })}
    />
    Ignore events while dead{#if !afterDeathAvailable}
      <span class="text-muted text-[11px]">(one pull at a time)</span>{/if}
  </label>

  <button
    type="button"
    class="text-nav inline-flex min-h-11 items-center text-[12px] font-bold tracking-[0.06em] uppercase md:min-h-0"
    onclick={() => onChange(DEFAULT_FILTERS)}
  >
    Clear filters
  </button>

  <p class="text-muted basis-full text-[12px]">
    A pet’s damage and healing count on its owner’s row, as the engine records it. Split it in Queries.
  </p>
</div>
