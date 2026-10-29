<!-- web/src/components/ui/Select.svelte -->
<!-- The styled select (polish spec 4.2): the native element with `appearance: none`, so the
     keyboard and phone pickers stay native, a 52 px field (44 px with `compact`, for a filter
     bar), the secondary ground and border and SelectChevron at the right. Options are passed as
     data; `error` draws the wipe border and the field's error line, `placeholder` a muted first
     option. Bind `value`. -->
<script lang="ts" module>
  export interface SelectOption {
    value: string;
    label: string;
  }
</script>

<script lang="ts">
  import SelectChevron from './SelectChevron.svelte';

  let {
    value = $bindable(''),
    options,
    id,
    name,
    label,
    placeholder,
    compact = false,
    disabled = false,
    error,
    describedBy,
    testid,
  }: {
    value?: string;
    options: readonly SelectOption[];
    id: string;
    /** The form field name, for a select inside a plain form. */
    name?: string;
    /** The accessible name when no visible `<label for>` exists. */
    label?: string;
    placeholder?: string;
    compact?: boolean;
    disabled?: boolean;
    error?: string;
    /** Ids of lines that describe the field (read with the error line). */
    describedBy?: string;
    testid?: string;
  } = $props();

  const errorId = $derived(`${id}-error`);
  const describedIds = $derived(
    [describedBy, error === undefined ? undefined : errorId].filter(Boolean).join(' ') || undefined,
  );
</script>

<div class="flex flex-col gap-1">
  <div class="relative">
    <select
      {id}
      {name}
      bind:value
      {disabled}
      aria-label={label}
      aria-invalid={error === undefined ? undefined : 'true'}
      aria-describedby={describedIds}
      data-testid={testid}
      class="rounded-control border-line-card text-strong hover:border-gold/30 focus-visible:outline-gold w-full cursor-pointer appearance-none border bg-white/[.045] pr-10 pl-3 text-[15px] font-semibold transition-colors duration-[120ms] focus-visible:outline-2 focus-visible:outline-offset-2 disabled:cursor-not-allowed disabled:opacity-45 {compact
        ? 'h-11'
        : 'h-[52px]'} {error === undefined ? '' : 'border-wipe'} {value === '' && placeholder !== undefined
        ? 'text-muted'
        : ''}"
    >
      {#if placeholder !== undefined}<option value="" class="text-muted">{placeholder}</option>{/if}
      {#each options as option (option.value)}
        <option value={option.value} class="bg-raised text-strong">{option.label}</option>
      {/each}
    </select>
    <SelectChevron />
  </div>
  {#if error !== undefined}
    <p id={errorId} class="text-wipe text-[13px]" role="alert" data-testid="select-error">{error}</p>
  {/if}
</div>
