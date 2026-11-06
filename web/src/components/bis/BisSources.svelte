<!-- web/src/components/bis/BisSources.svelte -->
<!-- The control bar's Sources group (web layout spec 2026-11-05 §5.11): a multi-select reading
     "All sources", "7 of 8 sources" or "Dungeon only", each kind with its count (slots with a
     ranked option from that source) and one footer line. Below 768 px the same choice is a row
     of toggle chips in the expanded control stack. Turning a kind off re-picks; nothing hides. -->
<script lang="ts">
  import Select from '../ui/Select.svelte';
  import { SOURCE_FILTER_KINDS } from '../../lib/bis/source-filter';
  import { sourceKindLabel, sourcesFieldLabel } from '../../lib/bis/source-choice';
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

  const FOOT = 'Counts are slots with a ranked option from that source. Turning one off re-picks each slot; nothing is hidden.';

  const options = $derived(
    SOURCE_FILTER_KINDS.map((kind) => ({
      value: kind,
      label: sourceKindLabel(kind),
      meta: String(counts[kind]),
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

<div class="ui-cbar-group bis-sources" data-kind="select" data-testid="bis-sources">
  <span class="ui-cbar-label text-label tracking-label text-muted font-bold uppercase" id="bis-sources-label"
    >Sources</span
  >
  <div class="bis-sources-select">
    <Select
      id="bis-sources-select"
      label="Sources"
      size="dense"
      fit
      multiple
      values={[...on]}
      fieldLabel={sourcesFieldLabel(on, band)}
      footCount={FOOT}
      {options}
      testid="bis-sources-select"
      onchange={(event) =>
        onchange(ordered([...(event.currentTarget as HTMLSelectElement).selectedOptions].map((o) => o.value)))}
    />
  </div>
  <div class="bis-sources-chips" role="group" aria-labelledby="bis-sources-label">
    {#each SOURCE_FILTER_KINDS as kind (kind)}
      <button
        type="button"
        class="bis-chip"
        aria-pressed={on.includes(kind)}
        data-testid={`bis-source-chip-${kind}`}
        onclick={() => toggle(kind)}
      >
        {sourceKindLabel(kind)} <span class="bis-chip-count">{counts[kind]}</span>
      </button>
    {/each}
    <p class="bis-sources-foot text-muted">{FOOT}</p>
  </div>
</div>

<style>
  .bis-sources-chips {
    display: flex;
    flex-wrap: wrap;
    gap: 6px;
  }
  .bis-sources-select {
    display: none;
  }
  .bis-chip {
    display: inline-flex;
    align-items: center;
    gap: 6px;
    min-height: 32px;
    padding: 0 10px;
    border: 1px solid color-mix(in srgb, currentColor 40%, transparent);
    border-radius: var(--radius-control, 6px);
    color: var(--color-muted);
    font-size: 13px;
    font-weight: 600;
  }
  .bis-chip[aria-pressed='true'] {
    border-color: var(--color-gold);
    color: var(--color-gold);
  }
  .bis-chip:focus-visible {
    outline: 2px solid var(--color-gold);
    outline-offset: 2px;
  }
  .bis-chip-count {
    font-variant-numeric: tabular-nums;
    opacity: 0.8;
  }
  .bis-sources-foot {
    flex-basis: 100%;
    font-size: 12px;
    line-height: 16px;
  }
  @media (min-width: 768px) {
    .bis-sources-chips {
      display: none;
    }
    .bis-sources-select {
      display: block;
    }
  }
</style>
