<!-- web/src/components/report/TargetsRail.svelte -->
<!-- The actor tabs' rail Section after This pull (logs spec 2026-11-07 6.2): Targets on Damage
     done, Healed on Healing, Sources on Damage taken. The scope's rows summed per unit name (a
     trash pack's six "Gluttonous Tick" GUIDs are one row), the top eight as Rows with their share
     under the name and the amount at the right (directions b1), then "and {n} more". -->
<script lang="ts">
  import EmptyState from '../ui/EmptyState.svelte';
  import { splitUnitName } from '../../lib/characters';
  import { compact } from '../../lib/format/number';
  import type { Actor } from '../../lib/report/types';
  import Row from '../ui/Row.svelte';
  import Section from '../ui/Section.svelte';

  let { title, actors }: { title: 'Targets' | 'Healed' | 'Sources'; actors: readonly Actor[] } = $props();

  const SHOWN = 8;
  const units = $derived.by(() => {
    // eslint-disable-next-line svelte/prefer-svelte-reactivity
    const sums = new Map<string, number>();
    for (const actor of actors)
      for (const pair of actor.targets) {
        // Same-named units summed (a trash pack's GUIDs are one row), by the log's full name.
        sums.set(pair.name, (sums.get(pair.name) ?? 0) + pair.total);
      }
    // A unit healed only into overhealing keeps its row at 0: it was healed, nothing landed.
    return [...sums.entries()].sort((a, b) => b[1] - a[1]);
  });
  const total = $derived(units.reduce((sum, [, amount]) => sum + amount, 0));
</script>

<Section {title} testid="report-targets" panel>
  {#if units.length === 0}
    <EmptyState inset message="Nothing in this window." />
  {:else}
    <div>
      {#each units.slice(0, SHOWN) as [name, amount] (name)}
        <Row
          title={splitUnitName(name).name}
          sub={`${total === 0 ? '0.0' : ((amount / total) * 100).toFixed(1)}%`}
          open
        >
          {#snippet trail()}<span class="text-strong text-[15px] font-semibold tabular-nums"
              >{compact(amount)}</span
            >{/snippet}
        </Row>
      {/each}
    </div>
    {#if units.length > SHOWN}
      <p class="text-small text-muted pt-[10px]">
        and {units.length - SHOWN} more
      </p>
    {/if}
  {/if}
</Section>
