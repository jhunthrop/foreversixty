<!-- web/src/components/report/ReportActions.svelte -->
<!-- The band's Report actions below 768 (logs spec 2026-11-07 6.1): a 44 px "···" button at the
     answer's top right that opens a menu on the shared overlay surface (ui/Overlay, the menu kind)
     with the pull's two actions, Sim this fight and Copy link, so a phone reaches them without
     scrolling to This pull. Escape or a click outside closes it and focus returns to the button. -->
<script lang="ts">
  import { reportCopy } from '../../lib/report/copy';
  import { CopyState } from '../../lib/ui/copy-state.svelte';
  import { simCopy } from '../../lib/sim/copy';
  import Overlay from '../ui/Overlay.svelte';

  let { simHref }: { simHref: string | null } = $props();

  let open = $state(false);
  const copier = new CopyState();
  $effect(() => () => copier.dispose());
  /** The menu closes once its Copied has shown (lib/ui/copy-feedback.ts's two seconds). */
  let closeAfterCopy = false;
  $effect(() => {
    if (copier.state !== 'idle' || !closeAfterCopy) return;
    closeAfterCopy = false;
    close(true);
  });
  let root = $state<HTMLElement | null>(null);

  function close(refocus: boolean): void {
    open = false;
    if (refocus) root?.querySelector<HTMLElement>('[data-report-actions-toggle]')?.focus();
  }

  async function copyLink(): Promise<void> {
    await copier.copy(window.location.href);
    closeAfterCopy = copier.state === 'copied';
  }

  function onKey(event: KeyboardEvent): void {
    if (open && event.key === 'Escape') {
      event.preventDefault();
      close(true);
    }
  }

  $effect(() => {
    if (!open) return;
    const outside = (event: PointerEvent): void => {
      if (root !== null && !root.contains(event.target as Node)) close(false);
    };
    document.addEventListener('pointerdown', outside);
    root?.querySelector<HTMLElement>('[role="menuitem"]')?.focus();
    return () => document.removeEventListener('pointerdown', outside);
  });
</script>

<!-- svelte-ignore a11y_no_static_element_interactions -->
<div class="relative md:hidden" bind:this={root} onkeydown={onKey} data-testid="report-actions">
  <button
    type="button"
    class="text-strong flex h-11 w-11 items-center justify-center text-[20px] leading-none font-bold"
    aria-haspopup="menu"
    aria-expanded={open}
    aria-label={reportCopy.actionsLabel}
    data-report-actions-toggle
    data-testid="report-actions-toggle"
    onclick={() => (open = !open)}>···</button
  >
  {#if open}
    <div class="absolute top-[calc(100%+4px)] right-0 z-40 min-w-[200px]">
      <Overlay kind="menu" label={reportCopy.actionsLabel}>
        <div class="flex flex-col py-1">
          {#if simHref !== null}
            <a role="menuitem" class="report-action" href={simHref} data-testid="report-actions-sim"
              >{simCopy.simThisFight}</a
            >
          {/if}
          <button
            type="button"
            role="menuitem"
            class="report-action"
            data-testid="report-actions-copy"
            onclick={() => void copyLink()}>{copier.label(reportCopy.copyLink)}</button
          >
        </div>
      </Overlay>
    </div>
  {/if}
</div>

<style>
  .report-action {
    display: flex;
    align-items: center;
    min-height: 44px;
    padding: 0 16px;
    font-size: 14px;
    font-weight: 600;
    text-align: left;
    color: var(--color-nav);
  }
  .report-action:hover,
  .report-action:focus-visible {
    color: var(--color-strong);
    outline: none;
  }
</style>
