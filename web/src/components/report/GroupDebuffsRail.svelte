<!-- web/src/components/report/GroupDebuffsRail.svelte -->
<!-- The Dispels tab's rail under This pull (round 5, b1-report-dispels): Debuffs on your group, a
     Row per debuff that landed on the scope's players (its icon and name), how many players it hit
     and how often it landed, with the Debuffs › link. What a dispeller could have taken off, beside
     what they did. Nothing landed draws nothing. -->
<script lang="ts">
  import type { SpellIconOf } from '../../lib/guild/report-spell-icons';
  import { groupDebuffs } from '../../lib/report/exchange-rails';
  import type { AuraTrack } from '../../lib/report/types';
  import Row from '../ui/Row.svelte';
  import Section from '../ui/Section.svelte';
  import SpellIcon from './SpellIcon.svelte';

  let {
    auras,
    players,
    spellIcon,
    tabHref,
  }: {
    auras: readonly AuraTrack[];
    players: ReadonlySet<string>;
    spellIcon: SpellIconOf;
    tabHref: string;
  } = $props();

  const rows = $derived(groupDebuffs(auras, players));
</script>

{#if rows.length > 0}
  <Section title="Debuffs on your group" action={{ label: 'Debuffs', href: tabHref }} panel testid="group-debuffs-rail">
    {#each rows as row (row.name)}
      <Row
        title={row.name}
        sub={`on ${row.players} ${row.players === 1 ? 'player' : 'players'}`}
        density="slim"
        open
      >
        {#snippet lead()}<SpellIcon url={spellIcon(row.spellId, row.name)} />{/snippet}
        {#snippet trail()}<span class="text-strong text-[14px] font-semibold tabular-nums">{row.landed}×</span
          >{/snippet}
      </Row>
    {/each}
  </Section>
{/if}
