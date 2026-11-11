<!-- web/src/components/guild/progression/BestByRole.svelte -->
<!-- The guild's best kill figure on one boss for each role (the progression read's best_by_role):
     the tank, the healer and the damage dealer, in class colour with their spec. The Progression
     tab's rail. -->
<script lang="ts">
  import type { GuildParseRef, GuildProgressionEncounter } from '../../../lib/guild/api';
  import { outputFigure, outputUnit } from '../../../lib/guild/output-figure';
  import { SEP } from '../../../lib/guild/raid-copy';
  import { specClassName } from '../../../lib/guild/spec-name';
  import { classColorVar } from '../../../lib/report/format';
  import Crest from '../../ui/Crest.svelte';
  import Section from '../../ui/Section.svelte';
  import StatRow from '../../ui/StatRow.svelte';
  import StatTable from '../../ui/StatTable.svelte';

  let { encounter }: { encounter: GuildProgressionEncounter } = $props();

  const ROLE_WORD = { tank: 'Tank', healer: 'Healer', dps: 'Damage' } as const;
  const roles = $derived(
    (['tank', 'healer', 'dps'] as const).flatMap((role) => {
      const parse: GuildParseRef | undefined = encounter.best_by_role[role];
      return parse === undefined ? [] : [{ role, parse }];
    }),
  );
</script>

{#if roles.length > 0}
  <Section title={`Best on ${encounter.name}`} panel flush testid="guild-progression-best-parse">
    <StatTable columns={[{ id: 'figure', label: 'Best' }]} bar={false} dense>
      {#each roles as { role, parse } (role)}
        <StatRow
          bar={false}
          figures={[{ value: `${outputFigure(parse.value)} ${outputUnit(parse.metric)}` }]}
        >
          {#snippet icon()}<Crest cls={parse.class} size={28} />{/snippet}
          {#snippet name()}<span style:color={classColorVar(parse.class)}>{parse.name}</span>{/snippet}
          {#snippet sub()}{ROLE_WORD[role]}{#if parse.spec}{SEP}{specClassName(
                parse.spec,
                parse.class,
              )}{/if}{/snippet}
        </StatRow>
      {/each}
    </StatTable>
  </Section>
{/if}
