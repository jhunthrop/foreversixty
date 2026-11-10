<!-- web/src/components/rankings/RankingsMoreFilters.svelte -->
<!-- The rankings bar's Filters group (logs spec 2026-11-07 section 7, decision K2): Region, Phase,
     Faction, Today only and every Metric, behind one compact button from 768 px ("Filters", "Filters (2)" when
     any is set) that opens a popover on the shared overlay surface (the BiS Sources pattern);
     Escape or a click outside closes it and focus returns to the button. Below 768 px the same
     fields sit in the expanded control stack. Every choice lands in the URL through `onpatch`. -->
<script lang="ts">
  import Icon from '../ui/Icon.svelte';
  import { ICON_CHEVRON_DOWN } from '../../lib/ui/ui-icons';
  import { factionOptions, metricOptions, phaseOptions, regionOptions } from '../../lib/rankings/options';
  import { rankingsCopy } from '../../lib/rankings/copy';
  import type { RankingsState } from '../../lib/rankings/url';
  import { button as buttonClass } from '../../lib/ui/button';
  import { popoverSupported } from '../../lib/ui/listbox-model';
  import type { PopoverHandle } from '../../lib/ui/popover';
  import type { SelectOption } from '../../lib/ui/listbox-model';
  import Overlay from '../ui/Overlay.svelte';
  import Select from '../ui/Select.svelte';
  import Switch from '../ui/Switch.svelte';

  let {
    filters,
    defaultPhase = '',
    onpatch,
  }: {
    filters: Pick<RankingsState, 'region' | 'phase' | 'faction' | 'since' | 'metric'>;
    /** The phase the board starts on (the current one): only a different choice counts as set. */
    defaultPhase?: string;
    onpatch: (next: Partial<RankingsState>) => void;
  } = $props();

  const PANEL_ID = 'rankings-filters-panel';

  interface Field {
    key: 'region' | 'phase' | 'faction' | 'metric';
    label: string;
    options: readonly SelectOption[];
  }
  const FIELDS: readonly Field[] = [
    { key: 'region', label: rankingsCopy.regionLabel, options: regionOptions() },
    { key: 'phase', label: rankingsCopy.phaseLabel, options: phaseOptions() },
    { key: 'faction', label: rankingsCopy.factionLabel, options: factionOptions() },
    // Every metric (Damage taken, Execution too); the bar's segmented control offers the two common ones.
    { key: 'metric', label: rankingsCopy.metricLabel, options: metricOptions() },
  ];
  const COUNTED: ReadonlySet<Field['key']> = new Set(['region', 'phase', 'faction']);

  let wrap = $state<HTMLElement | null>(null);
  let trigger = $state<HTMLButtonElement | null>(null);
  let panel = $state<HTMLElement | null>(null);
  let open = $state(false);
  let fallback = $state(false);
  let handle: PopoverHandle | null = null;

  const setCount = $derived(
    FIELDS.filter(
      (field) => COUNTED.has(field.key) && filters[field.key] !== (field.key === 'phase' ? defaultPhase : ''),
    ).length + (filters.since === '' ? 0 : 1),
  );
  const label = $derived(
    setCount === 0 ? rankingsCopy.filtersButton : `${rankingsCopy.filtersButton} (${setCount})`,
  );

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
    panel.querySelector<HTMLElement>('button, select')?.focus();
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

{#snippet fields(prefix: string)}
  {#each FIELDS as field (field.key)}
    <div class="flex flex-col gap-1">
      <span class="text-label text-muted font-bold" id={`${prefix}-${field.key}-label`}>{field.label}</span>
      <Select
        id={`${prefix}-${field.key}`}
        label={field.label}
        size="dense"
        value={filters[field.key]}
        options={field.options}
        testid={`${prefix}-${field.key}`}
        onchange={(event) => onpatch({ [field.key]: (event.currentTarget as HTMLSelectElement).value })}
      />
    </div>
  {/each}
  <div class="flex items-center gap-3">
    <Switch
      labelledby={`${prefix}-today-label`}
      checked={filters.since !== ''}
      testid={`${prefix}-today`}
      onchange={(next) => onpatch({ since: next ? 'today' : '' })}
    />
    <span class="text-[14px] font-semibold" id={`${prefix}-today-label`}>{rankingsCopy.todayOnly}</span>
  </div>
{/snippet}

<div class="ui-cbar-group rankings-more" data-kind="select" data-testid="rankings-more-filters">
  <span class="ui-cbar-label text-label text-muted font-bold">{rankingsCopy.filtersButton}</span>
  <div class="rankings-more-pop" bind:this={wrap}>
    <button
      bind:this={trigger}
      type="button"
      class="{buttonClass('secondary', 'sm', true)} rankings-more-button"
      aria-haspopup="dialog"
      aria-expanded={open}
      aria-controls={PANEL_ID}
      data-testid="rankings-filters-button"
      onclick={() => (open ? hide(true) : void show())}
    >
      <span>{label}</span>
      <Icon icon={ICON_CHEVRON_DOWN} size={14} />
    </button>
    <!-- svelte-ignore a11y_no_static_element_interactions -->
    <div
      bind:this={panel}
      id={PANEL_ID}
      class="rankings-more-panel"
      popover="manual"
      data-fallback={fallback && open ? 'true' : undefined}
      onkeydown={onPanelKey}
    >
      <Overlay kind="menu" role="dialog" label={rankingsCopy.filtersButton} testid="rankings-filters-popover">
        <div class="flex flex-col gap-3 p-4">
          {@render fields('rankings-pop')}
          <button
            type="button"
            class="text-gold self-start text-[14px] font-semibold"
            data-testid="rankings-filters-clear"
            onclick={() => onpatch({ region: '', phase: defaultPhase, faction: '', since: '' })}
            >Clear filters</button
          >
        </div>
      </Overlay>
    </div>
  </div>
  <div class="rankings-more-stack">{@render fields('rankings-stack')}</div>
</div>

<style>
  .rankings-more {
    display: flex;
    flex-direction: column;
    gap: 4px;
    min-width: 0;
  }
  .rankings-more-pop {
    display: none;
    position: relative;
  }
  .rankings-more-stack {
    display: flex;
    flex-direction: column;
    gap: 12px;
  }
  .rankings-more-button {
    display: inline-flex;
    align-items: center;
    gap: 8px;
    height: 40px;
    padding: 0 12px;
    color: var(--color-strong);
    font-size: 14px;
    font-weight: 600;
    white-space: nowrap;
  }
  .rankings-more-button:hover,
  .rankings-more-button[aria-expanded='true'] {
    border-color: var(--color-gold);
  }
  .rankings-more-panel {
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
  .rankings-more-panel:popover-open {
    display: block;
  }
  .rankings-more-panel[data-fallback='true'] {
    display: block;
    position: absolute;
    top: calc(100% + 4px);
    left: 0;
    z-index: 40;
  }
  @media (min-width: 768px) {
    .rankings-more-pop {
      display: block;
    }
    .rankings-more-stack {
      display: none;
    }
    .rankings-more > .ui-cbar-label {
      display: none;
    }
  }
</style>
