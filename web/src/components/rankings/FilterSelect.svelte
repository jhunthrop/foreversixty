<!-- web/src/components/rankings/FilterSelect.svelte -->
<!-- One filter: a small label over a 44 px select, the same control for every filter on
     /rankings (design/specs/2026-10-11-rankings.md section 5: rest, hover, focus-visible, 44 px
     everywhere). The chevron is drawn here so every select has the same one. `grow` lets the
     Encounter, the filter every answer needs, run wider than the rest. -->
<script lang="ts">
  import type { FilterOption } from '../../lib/rankings/options';
  import SelectChevron from '../ui/SelectChevron.svelte';

  let {
    id,
    label,
    value,
    options,
    onchange,
    testid,
    grow = 1,
    ownRow = false,
    labelHidden = false,
    compact = false,
    busy = false,
    disabled = false,
    invalid = false,
    groups,
    describedBy,
  }: {
    id: string;
    label: string;
    value: string;
    options: readonly FilterOption[];
    onchange: (value: string) => void;
    testid?: string;
    grow?: number;
    /** Takes a row of its own in the grid below the wide layout (the Encounter, when no
     *  Class and Spec lead the panel). */
    ownRow?: boolean;
    /** Keeps the label for screen readers only (the raid strip's Role select, raids spec 3.3). */
    labelHidden?: boolean;
    disabled?: boolean;
    /** The value was refused: `aria-invalid`, read with `describedBy`'s sentence. */
    invalid?: boolean;
    /** Options under `<optgroup>` headings, drawn after `options` (a null label draws bare). */
    groups?: readonly { label: string | null; options: readonly FilterOption[] }[];
    /** Read while disabled or invalid: the sentence that says why. */
    describedBy?: string;
    /** 32 px from `md`, 44 px below (the raid roster's officer select, raids spec 3.7). */
    compact?: boolean;
    /** The row's save is in flight. */
    busy?: boolean;
  } = $props();
</script>

<label
  class="flex min-w-0 flex-col gap-1 {ownRow ? 'col-span-full min-[1100px]:col-auto' : ''}"
  style:flex={`${grow} 1 0`}
  for={id}
  data-testid={testid === undefined ? undefined : `${testid}-field`}
>
  <span class={labelHidden ? 'sr-only' : 'label text-muted'}>{label}</span>
  <span class="relative block has-[select:disabled]:opacity-50">
    <select
      {id}
      class="border-line-card text-strong rounded-control hover:border-gold/30 focus-visible:outline-gold bg-white/[.045] {compact
        ? 'h-11 md:h-8'
        : 'h-11'} w-full min-w-0 cursor-pointer appearance-none truncate border pr-10 pl-3 text-[13px] transition-colors duration-[120ms] focus-visible:outline-2 focus-visible:outline-offset-2 disabled:cursor-not-allowed"
      {value}
      {disabled}
      aria-busy={busy ? 'true' : undefined}
      aria-label={labelHidden ? label : undefined}
      aria-invalid={invalid ? 'true' : undefined}
      aria-describedby={disabled || invalid ? describedBy : undefined}
      data-testid={testid}
      onchange={(event) => onchange((event.currentTarget as HTMLSelectElement).value)}
    >
      {#snippet choices(list: readonly FilterOption[])}
        {#each list as option (option.value)}
          <option value={option.value} disabled={option.disabled} aria-label={option.ariaLabel}
            >{option.label}</option
          >
        {/each}
      {/snippet}
      {@render choices(options)}
      {#each groups ?? [] as group (group.label ?? '')}
        {#if group.label === null}
          {@render choices(group.options)}
        {:else}
          <optgroup label={group.label}>{@render choices(group.options)}</optgroup>
        {/if}
      {/each}
    </select>
    <SelectChevron />
  </span>
</label>
