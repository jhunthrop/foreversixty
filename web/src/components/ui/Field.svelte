<!-- web/src/components/ui/Field.svelte -->
<!-- A text field (web layout spec 2026-11-05 §3A.2): a label in 12 px caps over a 40 px input with
     the 1 px --color-line edge and the field radius; an error sentence under it in --color-wipe,
     tied to the input by aria-describedby. Lookup, paste and login forms use it through FormRow. -->
<script lang="ts">
  let {
    label,
    value = $bindable(''),
    name,
    id,
    type = 'text',
    placeholder,
    error,
    grow = false,
    autocomplete,
    inputmode,
    required = false,
    maxlength,
    rows,
    hint,
    testid,
  }: {
    label: string;
    value?: string;
    name?: string;
    id: string;
    type?: 'text' | 'search' | 'email' | 'url' | 'date';
    /** The most characters the field takes. */
    maxlength?: number;
    placeholder?: string;
    error?: string;
    grow?: boolean;
    autocomplete?: HTMLInputElement['autocomplete'];
    /** The phone keyboard: `numeric` for a count. */
    inputmode?: 'numeric' | 'text';
    required?: boolean;
    /** A multi-line field (a message) this many lines tall. */
    rows?: number;
    /** One muted line under the field. */
    hint?: string;
    testid?: string;
  } = $props();

  const CONTROL = 'rounded-field bg-bg text-strong placeholder:text-muted w-full border px-3 text-[14px]';
  const FOCUS = 'focus-visible:outline-gold focus-visible:outline-2 focus-visible:outline-offset-2';
  const describedBy = $derived(
    [error ? `${id}-error` : '', hint ? `${id}-hint` : ''].filter((part) => part !== '').join(' ') ||
      undefined,
  );
</script>

<label class={['ui-field flex min-w-0 flex-col gap-[6px]', grow ? 'min-w-[220px] flex-1' : '']} for={id}>
  <span class="text-label tracking-label text-muted font-bold uppercase">{label}</span>
  {#if rows === undefined}
    <input
      {id}
      {name}
      {type}
      {placeholder}
      {autocomplete}
      {inputmode}
      {required}
      {maxlength}
      bind:value
      class={[CONTROL, 'h-10', FOCUS, error ? 'border-wipe' : 'border-line']}
      aria-invalid={error ? 'true' : undefined}
      aria-describedby={describedBy}
      data-testid={testid}
    />
  {:else}
    <textarea
      {id}
      {name}
      {placeholder}
      {required}
      {maxlength}
      {rows}
      bind:value
      class={[CONTROL, 'resize-y py-2 leading-5', FOCUS, error ? 'border-wipe' : 'border-line']}
      aria-invalid={error ? 'true' : undefined}
      aria-describedby={describedBy}
      data-testid={testid}></textarea>
  {/if}
  {#if hint}<span id={`${id}-hint`} class="text-muted text-[12px] leading-4">{hint}</span>{/if}
  {#if error}<span id={`${id}-error`} class="text-wipe text-[13px] leading-5">{error}</span>{/if}
</label>
