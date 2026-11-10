<!-- Its styles travel with its code (css="injected"): it renders only after the island hydrates, and a
     page inlines only the styles of what its server render drew. -->
<svelte:options css="injected" />

<script lang="ts">
  let {
    value = $bindable(''),
    label,
    placeholder,
    invalid = false,
    short = false,
    readonly = false,
    id,
    describedby,
    testid,
    oninput,
  }: {
    value?: string;
    /** The field's accessible name. */
    label: string;
    placeholder?: string;
    invalid?: boolean;
    /** One line, 44 px. */
    short?: boolean;
    readonly?: boolean;
    id?: string;
    /** The id of the message that explains an invalid code. */
    describedby?: string;
    testid?: string;
    oninput?: (event: Event & { currentTarget: HTMLTextAreaElement }) => void;
  } = $props();
</script>

<!-- web/src/components/ui/CodeField.svelte -->
<!-- The paste box for a literal code a player copies out of the game or the addon (planner spec
     2026-11-08 section 3A, the mock's CodeField): mono 13/20 (the one face for literal codes),
     68 px tall (`short`: 44 px, one line), the dark field ground, the --wipe border when the code
     does not read. -->

<textarea
  class="codef"
  class:bad={invalid}
  class:short
  bind:value
  {id}
  {placeholder}
  {readonly}
  rows={short ? 1 : 2}
  aria-label={label}
  aria-invalid={invalid ? 'true' : undefined}
  aria-describedby={describedby}
  spellcheck="false"
  autocomplete="off"
  data-testid={testid}
  {oninput}></textarea>

<style>
  .codef {
    display: block;
    box-sizing: border-box;
    width: 100%;
    min-height: 68px;
    margin: 0;
    padding: 10px 12px;
    border: 1px solid var(--color-line);
    border-radius: var(--radius-field);
    background: var(--color-bg);
    color: var(--color-strong);
    font: 500 13px/20px var(--font-mono);
    overflow-wrap: anywhere;
    white-space: pre-wrap;
    resize: vertical;
  }
  .codef::placeholder {
    color: var(--color-muted);
  }
  .codef:focus-visible {
    outline: 2px solid var(--color-gold);
    outline-offset: 2px;
  }
  .codef.bad {
    border-color: var(--color-wipe);
  }
  .codef.short {
    min-height: 44px;
    height: 44px;
    padding-block: 11px;
    resize: none;
    overflow: hidden;
  }
  .codef[readonly] {
    color: var(--color-text);
    resize: none;
  }
</style>
