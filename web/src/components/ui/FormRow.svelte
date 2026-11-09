<!-- web/src/components/ui/FormRow.svelte -->
<!-- Labelled Fields and Selects in one wrapping row, then one submit Button (web layout spec
     2026-11-05 §3A.2, §6.1 item 1): a form is not a filter, so it is not a ControlBar. Below 768 every
     field and the button take the full width, unless `inline`: one field and its button stay on one
     row at every width (the sim's paste box, "FS1… | Run"). -->
<script lang="ts">
  import type { Snippet } from 'svelte';

  let {
    onsubmit,
    action,
    method,
    label,
    inline = false,
    testid,
    children,
    submit,
  }: {
    onsubmit?: (event: SubmitEvent) => void;
    action?: string;
    method?: 'get' | 'post';
    /** The form's accessible name. */
    label: string;
    /** One field and its button on one row at every width; the field takes the room. */
    inline?: boolean;
    testid?: string;
    /** The Fields and Selects. */
    children: Snippet;
    /** The one submit Button. */
    submit: Snippet;
  } = $props();
</script>

<form
  class={['ui-form-row flex items-end gap-3', inline ? 'flex-nowrap' : 'flex-wrap']}
  data-inline={inline ? 'true' : undefined}
  {action}
  {method}
  {onsubmit}
  aria-label={label}
  data-testid={testid}
>
  {@render children()}
  <span class="ui-form-row-submit">{@render submit()}</span>
</form>

<style>
  @media (max-width: 767px) {
    .ui-form-row:not([data-inline]) > :global(*) {
      width: 100%;
    }
    .ui-form-row:not([data-inline]) .ui-form-row-submit > :global(*) {
      width: 100%;
    }
    .ui-form-row[data-inline] > :global(.ui-field) {
      flex: 1 1 0;
      min-width: 0;
    }
  }
</style>
