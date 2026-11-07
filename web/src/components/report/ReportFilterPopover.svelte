<!-- web/src/components/report/ReportFilterPopover.svelte -->
<!-- The report bar's Filters (logs spec 2026-11-07 6.1): one button from 768 px ("Filters",
     "Filters (2)" while any switch is on) that opens the switches on the shared overlay surface,
     the BiS Sources pattern; Escape or a click outside closes it and focus returns to the button.
     Below 768 the same switches sit in the open control stack. Boss damage only, Players only,
     Count overkill, Ignore events while dead, and Clear filters. -->
<script lang="ts">
  import { reportCopy } from '../../lib/report/copy';
  import { DEFAULT_FILTERS, type ReportFilters } from '../../lib/report/filters';
  import { button as buttonClass } from '../../lib/ui/button';
  import { popoverSupported } from '../../lib/ui/listbox-model';
  import type { PopoverHandle } from '../../lib/ui/popover';
  import Overlay from '../ui/Overlay.svelte';

  let {
    filters,
    showBossOnly,
    afterDeathAvailable,
    onChange,
  }: {
    filters: ReportFilters;
    /** False on Healing, where "boss damage only" has nothing to apply to. */
    showBossOnly: boolean;
    /** False over the whole night, where a death's span cannot be read across pulls. */
    afterDeathAvailable: boolean;
    onChange: (filters: ReportFilters) => void;
  } = $props();

  const PANEL_ID = 'report-filters-panel';
  type Switch = 'bossOnly' | 'playersOnly' | 'countOverkill' | 'ignoreAfterDeath';
  const SWITCHES: readonly { key: Switch; label: string }[] = [
    { key: 'bossOnly', label: 'Boss damage only' },
    { key: 'playersOnly', label: 'Players only' },
    { key: 'countOverkill', label: 'Count overkill' },
    { key: 'ignoreAfterDeath', label: 'Ignore events while dead' },
  ];
  const shownSwitches = $derived(SWITCHES.filter((item) => item.key !== 'bossOnly' || showBossOnly));
  const onCount = $derived(shownSwitches.filter((item) => filters[item.key]).length);
  const label = $derived(onCount === 0 ? reportCopy.filters : `${reportCopy.filters} (${onCount})`);

  let wrap = $state<HTMLElement | null>(null);
  let trigger = $state<HTMLButtonElement | null>(null);
  let panel = $state<HTMLElement | null>(null);
  let open = $state(false);
  let fallback = $state(false);
  let handle: PopoverHandle | null = null;

  async function show(): Promise<void> {
    if (open || wrap === null || trigger === null || panel === null) return;
    open = true;
    fallback = !popoverSupported();
    if (!fallback) {
      const { openPopover } = await import('../../lib/ui/popover');
      if (!open) return;
      handle = openPopover({
        anchor: trigger,
        panel,
        container: wrap,
        align: 'start',
        matchWidth: false,
        onDismiss: () => {
          handle = null;
          open = false;
        },
      });
    }
    panel.querySelector<HTMLElement>('input')?.focus();
  }

  function hide(refocus: boolean): void {
    if (!open) return;
    handle?.close();
    handle = null;
    open = false;
    if (refocus) trigger?.focus();
  }

  function onPanelKey(event: KeyboardEvent): void {
    if (event.key !== 'Escape') return;
    event.preventDefault();
    event.stopPropagation();
    hide(true);
  }
</script>

{#snippet switches(prefix: string)}
  {#each shownSwitches as item (item.key)}
    {@const disabled = item.key === 'ignoreAfterDeath' && !afterDeathAvailable}
    <label class="flex min-h-11 items-center gap-2 text-[14px] md:min-h-9">
      <input
        type="checkbox"
        checked={filters[item.key] && !disabled}
        {disabled}
        title={disabled ? 'Over the whole night a death’s span cannot be read; open a pull' : undefined}
        data-testid={item.key === 'bossOnly' ? `${prefix}filter-boss` : undefined}
        onchange={(event) =>
          onChange({ ...filters, [item.key]: (event.currentTarget as HTMLInputElement).checked })}
      />
      {item.label}
    </label>
  {/each}
  <button
    type="button"
    class="text-gold min-h-11 self-start text-[14px] font-semibold md:min-h-9"
    onclick={() => onChange({ ...DEFAULT_FILTERS, target: filters.target, ability: filters.ability })}
    >{reportCopy.clearFilters}</button
  >
  <p class="text-small text-muted">
    A pet’s damage and healing count on its owner’s row, as the engine records it. Split it in Queries.
  </p>
{/snippet}

<div class="report-filters" data-testid="filter-bar">
  <div class="report-filters-pop" bind:this={wrap}>
    <button
      bind:this={trigger}
      type="button"
      class="{buttonClass('secondary', 'sm')} report-filters-button"
      aria-haspopup="dialog"
      aria-expanded={open}
      aria-controls={PANEL_ID}
      data-testid="report-filters-button"
      onclick={() => (open ? hide(true) : void show())}>{label}</button
    >
    <!-- svelte-ignore a11y_no_static_element_interactions -->
    <div
      bind:this={panel}
      id={PANEL_ID}
      class="report-filters-panel"
      popover="manual"
      data-fallback={fallback && open ? 'true' : undefined}
      onkeydown={onPanelKey}
    >
      <Overlay kind="menu" role="dialog" label={reportCopy.filters} testid="report-filters-popover">
        <div class="flex flex-col gap-1 p-4">{@render switches('')}</div>
      </Overlay>
    </div>
  </div>
  <div class="report-filters-stack">{@render switches('stack-')}</div>
</div>

<style>
  .report-filters-pop {
    display: none;
    position: relative;
  }
  .report-filters-stack {
    display: flex;
    flex-direction: column;
  }
  .report-filters-button {
    height: 40px;
    padding: 0 12px;
    color: var(--color-strong);
    font-size: 14px;
    font-weight: 600;
  }
  .report-filters-panel {
    position: fixed;
    inset: auto;
    margin: 0;
    padding: 0;
    border: 0;
    background: transparent;
    overflow: visible;
    width: 300px;
    color: var(--color-strong);
    display: none;
  }
  .report-filters-panel:popover-open {
    display: block;
  }
  .report-filters-panel[data-fallback='true'] {
    display: block;
    position: absolute;
    top: calc(100% + 4px);
    left: 0;
    z-index: 40;
  }
  @media (min-width: 768px) {
    .report-filters-pop {
      display: block;
    }
    .report-filters-stack {
      display: none;
    }
  }
</style>
