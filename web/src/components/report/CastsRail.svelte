<!-- web/src/components/report/CastsRail.svelte -->
<!-- The Casts tab's rail under This pull (density pass 2026-11-09): By player, a Row per player
     in scope (crest, name in the class colour), their own casts per minute as the figure, the casts,
     the failed casts and the spells used in the sub line, busiest first. Pets' and totems' casts
     are left out. Read from the scoped summary's casts (summary.json casts[]:
     owner_guid, succeeded, failed). -->
<script lang="ts">
  import { splitUnitName } from '../../lib/characters';
  import { castsByPlayer } from '../../lib/report/casts-by-player';
  import { classSlugOf } from '../../lib/report/planner-link';
  import type { CastRow } from '../../lib/report/types';
  import { classTone } from '../../lib/ui/primitives';
  import Crest from '../ui/Crest.svelte';
  import Row from '../ui/Row.svelte';
  import Section from '../ui/Section.svelte';
  import { counted } from '../../lib/format/number';

  let {
    casts,
    players,
    names,
    classOf,
    durationMs,
  }: {
    casts: readonly CastRow[];
    players: ReadonlySet<string>;
    names: ReadonlyMap<string, string>;
    classOf: Map<string, string>;
    durationMs: number;
  } = $props();

  const rows = $derived(castsByPlayer(casts, players, durationMs));
</script>

{#if rows.length > 0}
  <Section title="By player" meta="casts per minute" testid="casts-rail" panel>
    {#each rows as row (row.guid)}
      {@const slug = classSlugOf(classOf.get(row.guid)) ?? ''}
      <Row
        title={splitUnitName(names.get(row.guid) ?? row.guid).name}
        titleTone={classTone(slug)}
        sub={[
          counted(row.casts, 'cast'),
          row.failed > 0 ? `${row.failed} failed` : '',
          counted(row.spells, 'spell'),
        ]
          .filter((part) => part !== '')
          .join(' · ')}
        density="slim"
        open
        testid={`casts-rail-${row.guid}`}
      >
        {#snippet lead()}<Crest cls={slug} size={24} />{/snippet}
        {#snippet trail()}<span class="text-strong text-[14px] font-semibold tabular-nums"
            >{row.perMinute.toFixed(1)}</span
          >{/snippet}
      </Row>
    {/each}
  </Section>
{/if}
