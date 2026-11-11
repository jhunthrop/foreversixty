<!-- web/src/components/bis/BisSources.svelte -->
<!-- The control bar's Sources group (web layout spec 2026-11-05 §5.11) on the shared FilterPopover:
     from 768 px a compact button ("All sources", "Sources (6)": the kinds that are on) opens the
     checklist, each kind with its count (slots with a ranked option from that source), and one
     footer line. Below 768 px the same choice is a row of toggle chips (ui/Chips) in the expanded
     control stack. Turning a kind off re-picks; nothing hides. -->
<script lang="ts">
  import FilterPopover from '../ui/FilterPopover.svelte';
  import Chips from '../ui/Chips.svelte';
  import { SOURCE_FILTER_KINDS } from '../../lib/bis/source-filter';
  import { sourceKindLabel, sourcesButtonLabel, sourcesFieldLabel } from '../../lib/bis/source-choice';
  import type { SourceFilterKind } from '../../lib/bis/types';

  let {
    on,
    counts,
    band,
    onchange,
  }: {
    on: readonly SourceFilterKind[];
    counts: Readonly<Record<SourceFilterKind, number>>;
    band: number;
    onchange: (next: SourceFilterKind[]) => void;
  } = $props();

  const LABEL = 'Sources';
  const FOOT =
    'Counts are slots with a ranked option from that source. Turning one off re-picks each slot; nothing is hidden.';

  const chips = $derived(
    SOURCE_FILTER_KINDS.map((kind) => ({
      id: kind,
      label: sourceKindLabel(kind),
      on: on.includes(kind),
      count: counts[kind],
    })),
  );

  function ordered(kinds: Iterable<string>): SourceFilterKind[] {
    const set = new Set(kinds);
    return SOURCE_FILTER_KINDS.filter((kind) => set.has(kind));
  }

  function toggle(kind: SourceFilterKind): void {
    onchange(on.includes(kind) ? on.filter((other) => other !== kind) : ordered([...on, kind]));
  }
</script>

<FilterPopover
  label={LABEL}
  buttonLabel={sourcesButtonLabel(on)}
  buttonAriaLabel={`${LABEL}: ${sourcesFieldLabel(on, band)}`}
  id="bis-sources-panel"
  testid="bis-sources"
  buttonTestid="bis-sources-button"
  popoverTestid="bis-sources-popover"
  width={280}
  focusSelector="input"
>
  {#snippet panel()}
    <fieldset class="bis-sources-list">
      <legend class="sr-only">{LABEL}</legend>
      {#each SOURCE_FILTER_KINDS as kind (kind)}
        <label class="bis-sources-row" data-testid={`bis-source-check-${kind}`}>
          <input type="checkbox" checked={on.includes(kind)} onchange={() => toggle(kind)} />
          <span class="bis-sources-name">{sourceKindLabel(kind)}</span>
          <span class="bis-sources-count">{counts[kind]}</span>
        </label>
      {/each}
    </fieldset>
  {/snippet}
  {#snippet foot()}{FOOT}{/snippet}
  {#snippet stack()}
    <Chips {chips} label={LABEL} testid="bis-source-chip" onToggle={(id) => toggle(id as SourceFilterKind)} />
    <p class="text-muted text-[12px] leading-4">{FOOT}</p>
  {/snippet}
</FilterPopover>

<style>
  .bis-sources-list {
    margin: 0;
    padding: 6px 0;
    border: 0;
  }
  .bis-sources-row {
    display: flex;
    align-items: center;
    gap: 10px;
    min-height: 36px;
    padding: 0 14px;
    font-size: 14px;
    font-weight: 600;
    cursor: pointer;
  }
  .bis-sources-row:hover {
    background: var(--color-hover);
  }
  .bis-sources-row input {
    width: 16px;
    height: 16px;
    accent-color: var(--color-gold);
  }
  .bis-sources-name {
    flex: 1;
  }
  .bis-sources-count {
    color: var(--color-muted);
    font-variant-numeric: tabular-nums;
    font-size: 13px;
  }
</style>
