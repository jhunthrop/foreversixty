<!-- web/src/components/ui/Sheet.svelte -->
<!-- The overlay sheet (sim spec 2026-11-06 §5.2, ruling S7; mock ws_prims `Sheet`): a dialog over a
     scrim with a title and a 44 px close control, a scrolling body and a foot of actions. From
     768 px a 480 px panel at the right, full height; below, a bottom sheet of at most 88vh with
     16 px top corners. It is fixed, so the page under it never moves. Focus goes to the title on
     open, stays inside, and returns to the control that opened it; Escape and the scrim close. -->
<script lang="ts">
  import type { Snippet } from 'svelte';
  import { tick } from 'svelte';
  import { trapTab } from '../../lib/ui/focus-trap';

  let {
    open,
    title,
    closeLabel = 'Close',
    onclose,
    testid,
    children,
    foot,
  }: {
    open: boolean;
    title: string;
    closeLabel?: string;
    onclose: () => void;
    testid?: string;
    children: Snippet;
    foot?: Snippet;
  } = $props();

  let panel = $state<HTMLElement | undefined>(undefined);
  let heading = $state<HTMLElement | undefined>(undefined);
  /** The control focused when the sheet opened, given focus back on close. */
  let opener: HTMLElement | null = null;

  $effect(() => {
    if (!open) return;
    opener = document.activeElement instanceof HTMLElement ? document.activeElement : null;
    void tick().then(() => heading?.focus());
    return () => {
      opener?.focus();
      opener = null;
    };
  });

  function onkeydown(event: KeyboardEvent): void {
    if (event.key === 'Escape') {
      event.preventDefault();
      onclose();
      return;
    }
    if (panel !== undefined) trapTab(event, panel, heading);
  }
</script>

{#if open}
  <div class="ui-sheet-scrim" aria-hidden="true" onclick={onclose}></div>
  <div
    bind:this={panel}
    class="ui-sheet bg-raised border-line"
    role="dialog"
    aria-modal="true"
    aria-labelledby={testid ? `${testid}-title` : 'ui-sheet-title'}
    data-testid={testid}
    tabindex="-1"
    {onkeydown}
  >
    <div class="ui-sheet-head border-line-soft">
      <h2
        bind:this={heading}
        id={testid ? `${testid}-title` : 'ui-sheet-title'}
        class="text-strong m-0 flex-1 text-[20px] leading-6 font-bold outline-none"
        tabindex="-1"
      >
        {title}
      </h2>
      <button
        type="button"
        class="ui-sheet-close text-nav"
        aria-label={closeLabel}
        onclick={onclose}
        data-testid={testid ? `${testid}-close` : undefined}>×</button
      >
    </div>
    <div class="ui-sheet-body">{@render children()}</div>
    {#if foot}<div class="ui-sheet-foot border-line-soft">{@render foot()}</div>{/if}
  </div>
{/if}

<style>
  .ui-sheet-scrim {
    position: fixed;
    inset: 0;
    background: #05070bcc;
    z-index: 60;
  }
  .ui-sheet {
    position: fixed;
    z-index: 61;
    display: flex;
    flex-direction: column;
    right: 0;
    left: 0;
    bottom: 0;
    max-height: 88vh;
    border-width: 1px;
    border-style: solid;
    border-radius: 16px 16px 0 0;
    outline: none;
  }
  .ui-sheet-head {
    display: flex;
    align-items: center;
    gap: 12px;
    padding: 16px 24px;
    border-bottom-width: 1px;
    border-bottom-style: solid;
  }
  .ui-sheet-close {
    width: 44px;
    height: 44px;
    display: inline-flex;
    align-items: center;
    justify-content: center;
    font-size: 22px;
    border-radius: 4px;
  }
  .ui-sheet-close:focus-visible {
    outline: 2px solid var(--color-gold);
    outline-offset: 2px;
  }
  .ui-sheet-body {
    flex: 1;
    overflow: auto;
    padding: 8px 24px 16px;
  }
  .ui-sheet-foot {
    display: flex;
    gap: 12px;
    justify-content: flex-end;
    padding: 16px 24px;
    border-top-width: 1px;
    border-top-style: solid;
  }
  @media (max-width: 767px) {
    .ui-sheet-foot > :global(*) {
      flex: 1;
    }
  }
  @media (min-width: 768px) {
    .ui-sheet {
      top: 0;
      left: auto;
      width: 480px;
      max-height: none;
      border-radius: 0;
      border-width: 0 0 0 1px;
    }
  }
</style>
