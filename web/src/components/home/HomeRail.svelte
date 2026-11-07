<!-- web/src/components/home/HomeRail.svelte -->
<!-- Home B's side column (spec 2026-11-07 R1 to R4): Your upgrades (the top three), Guild sims (the
     guild's ladder for the viewer's role, band and preset), Your spec this week (the grade and its
     move) and Top movers. A block with nothing true to say is left out. -->
<script lang="ts">
  import type { Item } from '../../lib/planner/types';
  import { SLOT_LABELS } from '../../lib/planner/types';
  import type { GuildSimRow, HomeAnswer } from '../../lib/home/home-answer';
  import type { CachedUpgrades } from '../../lib/home/upgrades-loader';
  import { UPGRADES_SHOWN, gainText, upgradeLines, type HomeLink } from '../../lib/home/home-view';
  import {
    GUILD_SIMS_SHOWN,
    characterLink,
    deltaText,
    figureText,
    guildSimsMeta,
    homeBCopy,
    homeBHrefs,
    moversMeta,
    moversTitle,
    specWeekView,
  } from '../../lib/home/home-b-view';
  import type { ClassSlug } from '../../lib/ui/primitives';
  import ClassCrestRing from '../character/ClassCrestRing.svelte';
  import Figure from '../ui/Figure.svelte';
  import GradeTile from '../ui/GradeTile.svelte';
  import ItemRow from '../ui/ItemRow.svelte';
  import Row from '../ui/Row.svelte';
  import Section from '../ui/Section.svelte';

  let {
    upgrades,
    bis,
    answer,
    guildHref,
    build,
  }: {
    upgrades: CachedUpgrades | null;
    bis: HomeLink | null;
    answer: HomeAnswer;
    guildHref: string | null;
    build: string;
  } = $props();

  const CREST_PX = 28;
  const items = $derived(new Map<number, Item>((upgrades?.items ?? []).map((item) => [item.id, item])));
  const list = $derived((upgrades?.result.upgrades ?? []).slice(0, UPGRADES_SHOWN));
  const total = $derived(upgrades?.result.totalGainDps ?? 0);
  const sims = $derived(answer.guild_sims);
  const ladder = $derived(sims === null ? [] : sims.rows.slice(0, GUILD_SIMS_SHOWN));
  const youInLadder = $derived(ladder.some((row) => row.you));
  const week = $derived(answer.spec_week === null ? null : specWeekView(answer.spec_week));
  const movers = $derived(answer.movers);
</script>

{#snippet simRow(row: GuildSimRow)}
  <Row
    title={row.name}
    titleTone={{ classSlug: row.class as ClassSlug }}
    rank={`#${row.rank}`}
    href={characterLink(row.key, row.name)}
    density="slim"
    selected={row.you}
    open
    testid="home-guild-sim"
  >
    {#snippet lead()}<ClassCrestRing characterClass={row.class} size={CREST_PX} />{/snippet}
    {#snippet trail()}
      {#if deltaText(row.delta) !== null}<span class="text-kill text-[13px] font-bold"
          >{deltaText(row.delta)}</span
        >{/if}
      <Figure value={figureText(row.value)} unit={sims?.unit} size="sm" align="end" />
    {/snippet}
  </Row>
{/snippet}

{#if list.length > 0}
  <Section
    title={homeBCopy.upgrades}
    meta={`${gainText(total)} DPS`}
    action={bis === null ? undefined : { label: homeBCopy.bisLink, href: bis.href }}
    testid="home-upgrades"
  >
    {#each list as upgrade (upgrade.slot)}
      {@const pick = items.get(upgrade.pick.item_id)}
      <ItemRow
        item={{
          id: upgrade.pick.item_id,
          name: upgrade.pick.item_name,
          quality: pick?.quality ?? 1,
          icon: pick?.icon,
        }}
        {build}
        iconSize={32}
        source={`${SLOT_LABELS[upgrade.slot]} · ${upgradeLines(upgrade).source}`}
        gain={upgrade.gainDps === null
          ? undefined
          : { value: gainText(upgrade.gainDps), unit: 'DPS', tone: 'kill' }}
        testid="home-upgrade"
      />
    {/each}
  </Section>
{/if}

{#if sims !== null}
  <Section
    title={homeBCopy.guildSims}
    meta={guildSimsMeta(sims)}
    action={guildHref === null ? undefined : { label: homeBCopy.guildPage, href: guildHref }}
    testid="home-guild-sims"
  >
    {#each ladder as row (row.key)}{@render simRow(row)}{/each}
    {#if sims.you !== null && !youInLadder}{@render simRow(sims.you)}{:else if !youInLadder}
      <Row title={homeBCopy.youNoSim} href={homeBHrefs.sim} density="slim" open testid="home-guild-sim-you">
        {#snippet trail()}<span class="text-gold text-[14px] font-semibold">{homeBCopy.runSim} ›</span
          >{/snippet}
      </Row>
    {/if}
  </Section>
{/if}

{#if answer.spec_week !== null && week !== null}
  <Section
    title={homeBCopy.specWeek}
    action={{ label: homeBCopy.tierList, href: homeBHrefs.tiers }}
    testid="home-spec-week"
  >
    <Row title={week.title} sub={week.sub} open>
      {#snippet lead()}
        {#if answer.spec_week?.grade}<GradeTile letter={answer.spec_week.grade} testid="home-grade" />{/if}
      {/snippet}
    </Row>
  </Section>
{/if}

{#if movers !== null && movers.rows.length > 0}
  <Section
    title={moversTitle(movers)}
    meta={moversMeta(movers)}
    action={{ label: homeBCopy.rankings, href: homeBHrefs.rankings }}
    testid="home-movers"
  >
    {#each movers.rows as row (row.key)}
      <Row
        title={row.name}
        titleTone={{ classSlug: row.class as ClassSlug }}
        rank={`#${row.rank}`}
        href={characterLink(row.key, row.name)}
        density="slim"
        open
        testid="home-mover"
      >
        {#snippet lead()}<ClassCrestRing characterClass={row.class} size={CREST_PX} />{/snippet}
        {#snippet trail()}
          {#if deltaText(row.delta) !== null}<span class="text-kill text-[13px] font-bold"
              >{deltaText(row.delta)}</span
            >{/if}
        {/snippet}
      </Row>
    {/each}
  </Section>
{/if}
