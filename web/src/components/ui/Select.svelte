<!-- web/src/components/ui/Select.svelte -->
<!-- The one dropdown (design/specs/2026-10-30-controls-charts.md §A). It renders the native
     <select> with every option and its data: the source of truth for value, form name, bind
     and onchange at every pointer. Over it sits the styled face (the picked label, its icon and
     class colour, the chevron), and the native control lies transparent on top, so a phone or
     coarse pointer taps straight into its own picker (D3). On a fine pointer listbox.ts loads
     after first paint and puts a trigger and our own list (top layer, Floating UI) over the
     same face, so the swap moves nothing (§A11.1). Sizes: default 52, compact 44, dense 36
     (44 below 768 or on a coarse pointer). Static Astro pages use Select.astro. -->
<script lang="ts" module>
  export type SelectSize = 'default' | 'compact' | 'dense';
</script>

<script lang="ts">
  import { selectCopy } from '../../lib/ui/copy';
  import {
    FOOT_EVENT,
    crestRing,
    labelColour,
    optionAttributes,
    pickedOption,
    qualityBorder,
    type SelectOption,
  } from '../../lib/ui/listbox-model';
  import { attachListbox, type AttachedField } from '../../lib/ui/select-boot';
  import Spinner from './Spinner.svelte';

  let {
    value = $bindable(''),
    values = $bindable([]),
    options,
    id,
    name,
    label,
    placeholder,
    size = 'default',
    searchable = false,
    searchPlaceholder,
    multiple = false,
    loading = false,
    busy = false,
    emptyLabel,
    align = 'start',
    disabled = false,
    required = false,
    error,
    invalid = false,
    describedBy,
    testid,
    class: className = '',
    fieldLabel,
    footLabel,
    footCount,
    onfoot,
    oncoarseopen,
    staticField = false,
    focusOnMount = false,
    onchange,
  }: {
    value?: string;
    /** The picked values of a `multiple` field. */
    values?: string[];
    options: readonly SelectOption[];
    id: string;
    /** The form field name, for a select inside a plain form. */
    name?: string;
    /** The accessible name when no visible `<label for>` exists. */
    label?: string;
    placeholder?: string;
    size?: SelectSize;
    /** Adds the filter field (§A4.4); `searchPlaceholder` is then required. */
    searchable?: boolean;
    searchPlaceholder?: string;
    /** The multi-select (§A4.5): always searchable, toggles stay open. */
    multiple?: boolean;
    loading?: boolean;
    /** A save of this field is in flight: `aria-busy` alone, the field keeps its look. */
    busy?: boolean;
    /** Shown in the field when `options` is empty; required wherever the list can be empty. */
    emptyLabel?: string;
    /** `end` for a field in a right slot: the list opens bottom-end (§A3.7). */
    align?: 'start' | 'end';
    disabled?: boolean;
    required?: boolean;
    error?: string;
    /** Invalid with the reason drawn by the caller (read through `describedBy`), no error line. */
    invalid?: boolean;
    /** Ids of lines that describe the field (read with the error line). */
    describedBy?: string;
    testid?: string;
    /** Classes on the outer box: its width in a row. */
    class?: string;
    /** A `multiple` field's closed text (`Top 8 of 40`, §A4.5). */
    fieldLabel?: string;
    /** A `multiple` field's foot button (`Top 8`) and count (`9 drawn`). */
    footLabel?: string;
    footCount?: string;
    onfoot?: () => void;
    /** A `multiple` field on a coarse pointer has no native picker: the caller opens its sheet. */
    oncoarseopen?: () => void;
    /** Select.astro's server-rendered field: its script boots it, not this component. */
    staticField?: boolean;
    /** Takes focus when it mounts (a picker that opens in place of a button). */
    focusOnMount?: boolean;
    onchange?: (event: Event) => void;
  } = $props();

  /** Fails fast on a searchable field with no filter placeholder (§A3 props). */
  function assertProps(): void {
    if (searchable && searchPlaceholder === undefined) {
      throw new Error(`Select ${id}: a searchable field needs searchPlaceholder.`);
    }
  }
  assertProps();

  const errorId = $derived(`${id}-error`);
  const showsError = $derived(error !== undefined || invalid);
  const describedIds = $derived(
    [describedBy, error === undefined ? undefined : errorId].filter(Boolean).join(' ') || undefined,
  );
  const empty = $derived(options.length === 0 && emptyLabel !== undefined);
  const blocked = $derived(disabled || loading || empty);
  const picked = $derived(multiple ? undefined : pickedOption(options, value));
  const showsPlaceholder = $derived(!multiple && picked === undefined && placeholder !== undefined);
  const faceText = $derived(
    loading
      ? selectCopy.loading
      : empty
        ? (emptyLabel ?? '')
        : multiple
          ? (fieldLabel ?? '')
          : (picked?.label ?? placeholder ?? ''),
  );
  const faceColour = $derived(loading || empty || picked === undefined ? undefined : labelColour(picked));

  /** Consecutive options sharing a group become one <optgroup>, in the order given. */
  const runs = $derived.by(() => {
    const out: { group: string | undefined; options: SelectOption[] }[] = [];
    for (const option of options) {
      const last = out.at(-1);
      if (last !== undefined && last.group === option.group) last.options.push(option);
      else out.push({ group: option.group, options: [option] });
    }
    return out;
  });

  let field: AttachedField | null = null;

  function listbox(node: HTMLElement): (() => void) | undefined {
    if (staticField) return undefined;
    if (focusOnMount) node.querySelector<HTMLElement>('.sel-native')?.focus();
    const attached = attachListbox(node);
    field = attached;
    const onFoot = (): void => onfoot?.();
    node.addEventListener(FOOT_EVENT, onFoot);
    return () => {
      node.removeEventListener(FOOT_EVENT, onFoot);
      attached.destroy();
      field = null;
    };
  }

  // The enhanced trigger and an open list follow every change the page makes.
  $effect(() => {
    void [options, value, values, blocked, busy, showsError, faceText, footCount, describedIds];
    field?.refresh();
  });

  function onMultipleChange(event: Event): void {
    const select = event.currentTarget as HTMLSelectElement;
    values = [...select.selectedOptions].map((option) => option.value);
    onchange?.(event);
  }
</script>

{#snippet optionList(list: readonly SelectOption[])}
  {#each list as option (option.value)}
    <option
      value={option.value}
      disabled={option.disabled}
      aria-label={option.ariaLabel}
      {...optionAttributes(option)}>{option.label}</option
    >
  {/each}
{/snippet}

{#snippet allOptions()}
  {#if placeholder !== undefined && !multiple}<option value="">{placeholder}</option>{/if}
  {#each runs as run, index (index)}
    {#if run.group === undefined}
      {@render optionList(run.options)}
    {:else}
      <optgroup label={run.group}>{@render optionList(run.options)}</optgroup>
    {/if}
  {/each}
{/snippet}

<div class={['sel-root', className]}>
  <div
    class="sel"
    data-listbox
    data-size={size}
    data-align={align}
    data-searchable={searchable || multiple ? 'true' : undefined}
    data-search-placeholder={searchPlaceholder}
    data-multiple={multiple ? 'true' : undefined}
    data-placeholder={placeholder !== undefined && !multiple ? 'true' : undefined}
    data-loading={loading ? 'true' : undefined}
    data-empty={empty ? 'true' : undefined}
    data-disabled={disabled ? 'true' : undefined}
    data-blocked={blocked ? 'true' : undefined}
    data-error={showsError ? 'true' : undefined}
    data-static={staticField ? 'true' : undefined}
    data-foot-label={footLabel}
    data-foot-count={footCount}
    {@attach listbox}
  >
    {#if multiple}
      <select
        class="sel-native"
        {id}
        {name}
        multiple
        hidden
        bind:value={values}
        disabled={blocked}
        aria-label={label}
        data-testid={testid}
        onchange={onMultipleChange}
      >
        {@render allOptions()}
      </select>
      <button
        type="button"
        class="sel-native"
        disabled={blocked}
        aria-label={label === undefined ? faceText : `${label}: ${faceText}`}
        aria-haspopup="dialog"
        data-testid={testid === undefined ? undefined : `${testid}-sheet`}
        onclick={() => oncoarseopen?.()}
      ></button>
    {:else}
      <select
        class="sel-native"
        {id}
        {name}
        bind:value
        disabled={blocked}
        {required}
        aria-label={label}
        aria-invalid={showsError ? 'true' : undefined}
        aria-describedby={describedIds}
        aria-busy={loading || busy ? 'true' : undefined}
        data-testid={testid}
        {onchange}
      >
        {@render allOptions()}
      </select>
    {/if}
    <!-- aria-disabled marks a disabled field's dimmed face as the disabled control's text. -->
    <span class="sel-face" aria-hidden="true" aria-disabled={disabled ? 'true' : undefined}>
      {#if picked?.icon !== undefined && !loading}
        <span
          class="sel-icon"
          data-kind={picked.icon.kind}
          style:--sel-quality={qualityBorder(picked.icon)}
          style:--sel-ring={crestRing(picked.icon, picked.colour)}
        >
          <img src={picked.icon.src} alt="" width="20" height="20" decoding="async" />
        </span>
      {/if}
      <span class="sel-label" class:is-placeholder={showsPlaceholder} style:color={faceColour}
        >{faceText}</span
      >
      {#if loading}
        <span class="sel-spin"><Spinner /></span>
      {:else if !empty}
        <svg
          class="sel-chev"
          width="16"
          height="16"
          viewBox="0 0 24 24"
          fill="none"
          stroke="currentColor"
          stroke-width="2.5"
          stroke-linecap="round"
          stroke-linejoin="round"
          aria-hidden="true"
          data-testid="select-chevron"><path d="M6 9l6 6 6-6" /></svg
        >
      {/if}
    </span>
  </div>
  {#if error !== undefined}
    <p id={errorId} class="sel-error" role="alert" data-testid="select-error">{error}</p>
  {/if}
</div>
