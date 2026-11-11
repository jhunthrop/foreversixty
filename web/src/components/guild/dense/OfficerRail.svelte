<!-- web/src/components/guild/dense/OfficerRail.svelte -->
<!-- The officer Overview's rail (dense spec 5.1): roles still short if everyone answers with the open
     applications that fit them, the last night's loot with the top candidate on each open drop, and
     the raiders slipping on attendance with their last nights. The buff cover keeps its own panel. -->
<script lang="ts">
  import activeBuild from '../../../data/active-build.json';
  import type { OfficerApplication } from '../../../lib/guild/recruit-api';
  import type { LootCouncilCandidate, LootNight } from '../../../lib/guild/loot-council-api';
  import { classColorVar, parseText } from '../../../lib/report/format';
  import ItemRow from '../../ui/ItemRow.svelte';
  import { GROUP_WORD } from '../../../lib/guild/dense/roster-board';
  import type { RecruitGap, SlipRow } from '../../../lib/guild/dense/week-board';
  import Crest from '../../ui/Crest.svelte';
  import Button from '../../ui/Button.svelte';
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
    topOf = new Map(),
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
    /** Each open drop's leading candidate (the council's first). */
    topOf?: ReadonlyMap<number, LootCouncilCandidate>;
  } = $props();

  const build = activeBuild.build;
  const pending = $derived(applications.filter((a) => a.status === 'pending'));
</script>

{#if gaps.length > 0}
  <Section title="Recruit for the gaps" panel testid="guild-week-recruit">
    {#each gaps as gap (gap.group)}
      <Row title={GROUP_WORD[gap.group]} sub={`${gap.short} seat${gap.short === 1 ? '' : 's'} open`} open>
        {#snippet trail()}<Button tone="ghost" size="sm" href={recruitHref}>Post the role</Button>{/snippet}
      </Row>
    {/each}
    {#each pending as app (app.id)}
      {@const parse = app.character.best_parse}
      <Row
        title={app.character.name}
        titleTone={{ classSlug: (app.character.class ?? 'warrior') as never }}
        sub={[
          app.character.spec_name ?? '',
          parse === null || parse.percentile === null ? '' : `parse ${parseText(parse.percentile)}`,
        ]
          .filter((p) => p !== '')
          .join(' · ')}
        href={applicationsHref}
        open
      >
        {#snippet lead()}<Crest cls={app.character.class ?? ''} size={32} />{/snippet}
      </Row>
    {/each}
  </Section>
{/if}

{#if night !== null && night.drops.length > 0}
  <Section
    title={`Loot, ${nightLabel}`}
    panel
    action={{ label: 'Loot', href: lootHref }}
    testid="guild-week-loot"
  >
    {#each night.drops as drop (drop.item_id)}
      {@const top = topOf.get(drop.item_id) ?? null}
      <ItemRow
        item={{
          id: drop.item_id,
          name: drop.name,
          quality: drop.quality,
          icon: drop.icon,
        }}
        {build}
        iconSize={32}
        testid={`guild-week-drop-${drop.item_id}`}
      >
        {#snippet sourceContent()}
          {#if drop.award !== null}<b style:color={classColorVar(drop.award.class)}>{drop.award.name}</b
            >{:else if top !== null}Open · <b style:color={classColorVar(top.class)}>{top.name}</b
            >{#if top.gain_dps !== null}{` +${top.gain_dps} ${(top.gain_unit ?? 'dps').toUpperCase()}`}{/if}{:else}Open{/if}
        {/snippet}
      </ItemRow>
    {/each}
  </Section>
{/if}

{#if slipping.length > 0}
  <Section title="Slipping" panel action={{ label: 'Roster', href: '#roster' }} testid="guild-week-slipping">
    {#each slipping as row (row.name)}
      <Row title={row.name} titleTone={{ classSlug: row.cls as never }} open>
        {#snippet lead()}<Crest cls={row.cls} size={28} />{/snippet}
        {#snippet trail()}<AttendanceStrip cells={row.cells} />{/snippet}
      </Row>
    {/each}
    <p class="pt-3"><AttendanceStrip legend /></p>
  </Section>
{/if}
