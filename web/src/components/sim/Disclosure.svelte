<!-- web/src/components/sim/Disclosure.svelte -->
<!-- Task 4, Ruling 3 (task-4-brief.md round 2, controller): the one show/hide primitive
     every tappable "more" control in this lane builds on. Task 4's own "what's in it" is
     the first caller; Task 7 lands a help note on the same Buffs control this task already
     occupies, and two bespoke toggles on one control would be the duplication the ruling
     exists to head off -- so nothing here is specific to buffs, presets or any other
     caller's content.

     Not a native <details>: SettingsSheet.svelte and BuffPanel.svelte already use one each
     for a disclosure that never needs a second one beside it, and that native element has
     no aria-controls and no Escape-to-close. This is a button plus a conditionally-rendered
     panel instead, wired by hand to the three things the ruling names -- aria-expanded,
     aria-controls and Escape -- so a control that gains a second Disclosure (Task 7) gets
     the same behaviour from the same element, not a second toggle re-deriving it. -->
<script lang="ts">
  import type { Snippet } from 'svelte';

  let {
    label,
    id,
    open = $bindable(false),
    disabled = false,
    triggerClass = '',
    panelClass = '',
    triggerTestId,
    panelTestId,
    children,
  }: {
    /** The trigger's own visible text and accessible name. */
    label: string;
    /** Unique within the page; becomes the panel's id and the trigger's `aria-controls`. */
    id: string;
    open?: boolean;
    disabled?: boolean;
    triggerClass?: string;
    panelClass?: string;
    triggerTestId?: string;
    panelTestId?: string;
    children: Snippet;
  } = $props();

  const panelId = $derived(`${id}-panel`);
  let triggerEl: HTMLButtonElement | undefined;

  function toggle(): void {
    open = !open;
  }

  /**
   * Escape closes, from anywhere on the page while the panel is open -- not a keydown
   * handler on the panel itself, which svelte-check's a11y rule flags on a non-interactive
   * element (the panel's own content can be a plain list, not a focusable region) and which
   * would miss the key firing while focus is still on the trigger button. `<svelte:window>`
   * is the one place a keydown belongs when what it closes is not the element with focus.
   */
  function onWindowKeydown(event: KeyboardEvent): void {
    if (event.key !== 'Escape') return;
    open = false;
    triggerEl?.focus();
  }
</script>

<svelte:window onkeydown={open ? onWindowKeydown : undefined} />

<div class="flex flex-col gap-1">
  <button
    bind:this={triggerEl}
    type="button"
    {disabled}
    class={`flex min-h-11 min-w-11 items-center gap-1 text-left ${triggerClass}`}
    aria-expanded={open}
    aria-controls={panelId}
    data-testid={triggerTestId}
    onclick={toggle}
  >
    {label}
  </button>
  {#if open}
    <div id={panelId} class={panelClass} data-testid={panelTestId} role="group" aria-label={label}>
      {@render children()}
    </div>
  {/if}
</div>
