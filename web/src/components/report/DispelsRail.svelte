<!-- web/src/components/report/DispelsRail.svelte -->
<!-- The Debuffs tab's rail under This pull (density pass 2026-11-09): Dispelled, a Row per debuff
     taken off (the removed aura's icon and name), who took it off whom with which spell, and how
     many times, with the Dispels › link. Read from the scoped summary's dispels (summary.json
     dispels[]: extra_spell_* is what was removed). Nothing dispelled draws nothing. -->
<script lang="ts">
  import { splitUnitName } from '../../lib/characters';
  import type { SpellIconOf } from '../../lib/guild/report-spell-icons';
  import type { ExchangeRow } from '../../lib/report/types';
  import Row from '../ui/Row.svelte';
  import Section from '../ui/Section.svelte';
  import SpellIcon from './SpellIcon.svelte';

  let {
    dispels,
    spellIcon,
    tabHref,
  }: { dispels: readonly ExchangeRow[]; spellIcon: SpellIconOf; tabHref: string } = $props();

  const rows = $derived(
    dispels
      .filter((row) => row.kind === 'dispel' && row.count > 0)
      .toSorted((a, b) => b.count - a.count || a.extra_spell_name.localeCompare(b.extra_spell_name)),
  );
  const who = (name: string): string => splitUnitName(name).name;
</script>

{#if rows.length > 0}
  <Section title="Dispelled" action={{ label: 'Dispels', href: tabHref }} panel testid="dispels-rail">
    {#each rows as row (`${row.source_guid}-${row.target_guid}-${row.spell_id}-${row.extra_spell_id}`)}
      <Row
        title={row.extra_spell_name || `Spell #${row.extra_spell_id}`}
        sub={`${row.spell_name} · ${who(row.source_name)} on ${who(row.target_name)}`}
        density="slim"
        open
      >
        {#snippet lead()}<SpellIcon url={spellIcon(row.extra_spell_id, row.extra_spell_name)} />{/snippet}
        {#snippet trail()}<span class="text-strong text-[14px] font-semibold tabular-nums">{row.count}</span
          >{/snippet}
      </Row>
    {/each}
  </Section>
{/if}
