<!-- web/src/components/rankings/FilterSelect.svelte -->
<!-- One filter: a small label over a 44 px select, the same control for every filter on
     /rankings (design/specs/2026-10-11-rankings.md section 5: rest, hover, focus-visible, 44 px
     everywhere). The chevron is drawn here so every select has the same one. `grow` lets the
     Encounter, the filter every answer needs, run wider than the rest. -->
<script lang="ts">
  import type { FilterOption } from '../../lib/rankings/options';

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
    disabled = false,
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
    /** Read while disabled: the sentence that says why. */
    describedBy?: string;
  } = $props();
</script>

<label
  class="flex min-w-0 flex-col gap-1 {ownRow ? 'col-span-full min-[1100px]:col-auto' : ''}"
  style:flex={`${grow} 1 0`}
  for={id}
  data-testid={testid === undefined ? undefined : `${testid}-field`}
>
  <span class={labelHidden ? 'sr-only' : 'label text-muted'}>{label}</span>
  <span class="relative block">
    <select
      {id}
      class="border-line-warm bg-raised text-text rounded-control hover:border-gold-hover h-11 w-full min-w-0 cursor-pointer appearance-none truncate border pr-8 pl-3 text-[13px] transition-colors duration-[120ms] disabled:cursor-not-allowed"
      {value}
      {disabled}
      aria-label={labelHidden ? label : undefined}
      aria-describedby={disabled ? describedBy : undefined}
      data-testid={testid}
      onchange={(event) => onchange((event.currentTarget as HTMLSelectElement).value)}
    >
      {#each options as option (option.value)}
        <option value={option.value}>{option.label}</option>
      {/each}
    </select>
    <svg
      class="text-muted pointer-events-none absolute top-1/2 right-3 -translate-y-1/2"
      width="12"
      height="12"
      viewBox="0 0 24 24"
      fill="none"
      stroke="currentColor"
      stroke-width="3"
      stroke-linecap="round"
      aria-hidden="true"><path d="M6 9l6 6 6-6" /></svg
    >
  </span>
</label>
