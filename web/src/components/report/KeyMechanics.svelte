<!-- web/src/components/report/KeyMechanics.svelte -->
<!-- The Summary rail's Mechanics on a boss pull (owner review 2026-11-09): the boss's mechanics at
     a glance, with the Mechanics › link. With a table: its rows, the ones that went wrong first,
     each with how often it landed where it should not ("3 hits", or "Avoided"). Without one (a
     boss no table covers yet): the boss abilities that hit players, by damage, with how many
     they hit. At most five rows; the tab has the rest. -->
<script lang="ts">
  import { mechanicTable, unclassifiedAbilities } from '../../lib/report/mechanics';
  import { formatAmount } from '../../lib/report/format';
  import type { SpellIconOf } from '../../lib/guild/report-spell-icons';
  import type { Summary } from '../../lib/report/types';
  import Row from '../ui/Row.svelte';
  import Section from '../ui/Section.svelte';
  import SpellIcon from '../ui/SpellIcon.svelte';

  let { summary, spellIcon, tabHref }: { summary: Summary; spellIcon: SpellIconOf; tabHref: string } =
    $props();

  const SHOWN = 5;
  /** Each ability's school, from what it did to players: the icon chain's school step for a boss
   *  ability the build has no icon for, so the same spell draws the same icon here and in the tables. */
  const schoolOf = $derived(
    new Map(
      summary.damage_taken.flatMap((actor) =>
        actor.abilities.map((ability) => [ability.spell_id, ability.school] as const),
      ),
    ),
  );
  const plural = (n: number, one: string, many = `${one}s`): string => `${n} ${n === 1 ? one : many}`;

  interface Line {
    key: string;
    spellId: number;
    name: string;
    figure: string;
    sub: string;
    bad: boolean;
  }

  const lines = $derived.by((): Line[] => {
    if (summary.mechanics?.table_found === true)
      return mechanicTable(summary.mechanics.rows, summary.roster)
        .slice(0, SHOWN)
        .map(({ row, tally }) => ({
          key: `${row.spell_id}-${row.name}`,
          spellId: row.spell_id,
          name: row.name,
          figure: tally.failed > 0 ? plural(tally.failed, 'hit') : 'Avoided',
          sub: tally.who,
          bad: tally.failed > 0 || tally.killed,
        }));
    return unclassifiedAbilities(summary)
      .filter((ability) => ability.damage > 0)
      .slice(0, SHOWN)
      .map((ability) => ({
        key: String(ability.spell_id),
        spellId: ability.spell_id,
        name: ability.name,
        figure: formatAmount(ability.damage),
        sub: `${plural(ability.players, 'player')} hit`,
        bad: false,
      }));
  });
</script>

{#if lines.length > 0}
  <Section
    title="Mechanics"
    action={{ label: 'Mechanics', href: tabHref, arrow: true }}
    panel
    testid="summary-mechanics"
  >
    {#each lines as line (line.key)}
      <Row title={line.name} sub={line.sub} open density="slim">
        {#snippet lead()}<SpellIcon
            socket
            tip={line.spellId}
            size={24}
            url={spellIcon(line.spellId, line.name, schoolOf.get(line.spellId))}
          />{/snippet}
        {#snippet trail()}<span
            class={['text-[14px] font-semibold tabular-nums', line.bad ? 'text-wipe' : 'text-strong']}
            >{line.figure}</span
          >{/snippet}
      </Row>
    {/each}
  </Section>
{/if}
