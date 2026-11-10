<!-- web/src/components/guild/dense/OfficerRail.svelte -->
<!-- The officer Overview's rail (dense spec 5.1): roles still short if everyone answers with the open
     applications that fit them, the last night's loot with the top candidate on each open drop, and
     the raiders slipping on attendance with their last nights. The buff cover keeps its own panel. -->
<script lang="ts">
  import activeBuild from '../../../data/active-build.json';
  import type { OfficerApplication } from '../../../lib/guild/recruit-api';
  import type { LootNight } from '../../../lib/guild/loot-council-api';
  import { GROUP_WORD } from '../../../lib/guild/dense/roster-board';
  import type { RecruitGap, SlipRow } from '../../../lib/guild/dense/week-board';
  import { lootIconOf } from '../../../lib/guild/loot-icon';
  import ClassCrestRing from '../../character/ClassCrestRing.svelte';
  import Button from '../../ui/Button.svelte';
  import ItemIcon from '../../ui/ItemIcon.svelte';
  import Row from '../../ui/Row.svelte';
  import Section from '../../ui/Section.svelte';
  import AttendanceStrip from './AttendanceStrip.svelte';

  let {
    gaps,
    applications,
    applicationsHref,
    recruitHref,
    night,
    nightLabel,
    lootHref,
    slipping,
  }: {
    gaps: readonly RecruitGap[];
    applications: readonly OfficerApplication[];
    applicationsHref: string;
    recruitHref: string;
    /** The newest loot night; null before one. */
    night: LootNight | null;
    nightLabel: string;
    lootHref: string;
    slipping: readonly SlipRow[];
  } = $props();

  const build = activeBuild.build;
  const pending = $derived(applications.filter((a) => a.status === 'pending'));
</script>

{#if gaps.length > 0}
  <Section title="Recruit for the gaps" panel testid="guild-week-recruit">
    {#each gaps as gap (gap.group)}
      <Row title={GROUP_WORD[gap.group]} sub={`Short ${gap.short} if all answer`} open>
        {#snippet trail()}<Button tone="secondary" size="sm" href={recruitHref}>Post the role</Button>{/snippet}
      </Row>
    {/each}
    {#each pending as app (app.id)}
      <Row
        title={app.character.name}
        titleTone={{ classSlug: (app.character.class ?? 'warrior') as never }}
        sub={[app.character.spec_name ?? '', app.character.item_level === null ? '' : `${app.character.item_level} ilvl`]
          .filter((p) => p !== '')
          .join(' · ')}
        open
      >
        {#snippet lead()}<ClassCrestRing characterClass={app.character.class ?? ''} size={32} />{/snippet}
        {#snippet trail()}<Button tone="secondary" size="sm" href={applicationsHref}>Open application</Button>{/snippet}
      </Row>
    {/each}
  </Section>
{/if}

{#if night !== null && night.drops.length > 0}
  <Section title={`Loot, ${nightLabel}`} panel action={{ label: 'Loot', href: lootHref }} testid="guild-week-loot">
    {#each night.drops as drop (drop.item_id)}
      <Row
        title={drop.name}
        sub={drop.award === null ? 'Open' : drop.award.name}
        open
      >
        {#snippet lead()}<ItemIcon icon={lootIconOf(drop) ?? undefined} quality={drop.quality} {build} size={32} />{/snippet}
        {#snippet trail()}
          {#if drop.award === null}<Button tone="secondary" size="sm" href={lootHref}>Council</Button>{/if}
        {/snippet}
      </Row>
    {/each}
  </Section>
{/if}

{#if slipping.length > 0}
  <Section title="Slipping" panel action={{ label: 'Roster', href: '#roster' }} testid="guild-week-slipping">
    {#each slipping as row (row.name)}
      <Row title={row.name} titleTone={{ classSlug: row.cls as never }} sub={`${row.present} of ${row.nights} nights`} open>
        {#snippet trail()}<AttendanceStrip cells={row.cells} />{/snippet}
      </Row>
    {/each}
    <p class="pt-3"><AttendanceStrip legend /></p>
  </Section>
{/if}
