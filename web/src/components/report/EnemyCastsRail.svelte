<!-- web/src/components/report/EnemyCastsRail.svelte -->
<!-- The Interrupts tab's rail under This pull (round 5, b1-report-interrupts): Enemy casts, a Row
     per enemy spell cast in the window (its icon and name), who cast it and its cast time, and how
     many went off; the casts with a cast bar first, since those are the ones a kick stops. Read from
     the unscoped summary's casts (the Source scope keeps only friendly rows). -->
<script lang="ts">
  import { splitUnitName } from '../../lib/characters';
  import type { SpellIconOf } from '../../lib/guild/report-spell-icons';
  import { enemyCasts } from '../../lib/report/exchange-rails';
  import type { CastRow } from '../../lib/report/types';
  import Row from '../ui/Row.svelte';
  import Section from '../ui/Section.svelte';
  import SpellIcon from './SpellIcon.svelte';

  let {
    casts,
    players,
    spellIcon,
  }: { casts: readonly CastRow[]; players: ReadonlySet<string>; spellIcon: SpellIconOf } = $props();

  const rows = $derived(enemyCasts(casts, players));
</script>

{#if rows.length > 0}
  <Section title="Enemy casts" panel testid="enemy-casts-rail">
    {#each rows as row (row.spellId)}
      <Row
        title={row.name}
        sub={`${row.casters.map((name) => splitUnitName(name).name).join(', ')} · ${row.castLabel}`}
        density="slim"
        open
      >
        {#snippet lead()}<SpellIcon url={spellIcon(row.spellId, row.name)} />{/snippet}
        {#snippet trail()}<span class="text-strong text-[14px] font-semibold tabular-nums">{row.count}</span
          >{/snippet}
      </Row>
    {/each}
  </Section>
{/if}
