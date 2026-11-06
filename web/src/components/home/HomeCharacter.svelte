<!-- web/src/components/home/HomeCharacter.svelte -->
<!-- The signed-in home for one character (home spec 2026-11-06 §2): the head (crest, the name in the
     class colour, the facts, the upgrades figure), the one next-step line, then the next raid, the
     upgrades Plate, talents and the last sim, with the rail beside. Each read has its own state;
     only the upgrades show a failure, the others fall back to their empty row. -->
<script lang="ts">
  import { untrack } from 'svelte';
  import { listMyReports, type Me, type MeCharacter, type MyReport } from '../../lib/account/api';
  import { fetchGuildHome, type GuildHome } from '../../lib/guild/api';
  import { fetchBisFileStrict } from '../../lib/bis/hover';
  import { dateLabel } from '../../lib/guild/raid-time';
  import { raidTimeLine } from '../../lib/guild/overview-role-view';
  import { cachedUpgradesFor, homeSpecKeyFor, type CachedUpgrades } from '../../lib/home/upgrades-loader';
  import { invalidate } from '../../lib/data/query';
  import { upgradesCacheKey } from '../../lib/home/upgrades-loader';
  import {
    bandWords,
    bisLink,
    gainText,
    guildLinks,
    homeCopy,
    homeHrefs,
    metaFacts,
    nextStep,
    ownLogs,
    simRow,
    talentsToSpend,
  } from '../../lib/home/home-view';
  import { listMySims } from '../../lib/sim/api';
  import type { SimListRow } from '../../lib/sim/types';
  import { classSlugFromName } from '../../lib/report/tree-sizes';
  import type { ClassSlug } from '../../lib/ui/primitives';
  import { LINK } from '../../lib/ui/primitives';
  import activeBuild from '../../data/active-build.json';
  import ClassCrestRing from '../character/ClassCrestRing.svelte';
  import Checklist from '../ui/Checklist.svelte';
  import Figure from '../ui/Figure.svelte';
  import Page from '../ui/Page.svelte';
  import PageHeader from '../ui/PageHeader.svelte';
  import Row from '../ui/Row.svelte';
  import Section from '../ui/Section.svelte';
  import Tag from '../ui/Tag.svelte';
  import HomeRail from './HomeRail.svelte';
  import HomeSummary from './HomeSummary.svelte';
  import HomeUpgrades from './HomeUpgrades.svelte';

  let { me, character }: { me: Me; character: MeCharacter } = $props();

  const CREST_PX = 56;
  const now = new Date();
  const viewerTz = Intl.DateTimeFormat().resolvedOptions().timeZone;

  let upgradesStatus = $state<'loading' | 'ready' | 'failed' | 'none'>('loading');
  let upgrades = $state<CachedUpgrades | null>(null);
  let sims = $state<SimListRow[]>([]);
  let reports = $state<MyReport[]>([]);
  let guildHome = $state<GuildHome | null>(null);

  function loadUpgrades(): void {
    upgradesStatus = 'loading';
    const target = character;
    cachedUpgradesFor(target, fetchBisFileStrict)
      .then((answer) => {
        if (target !== character) return;
        upgrades = answer;
        upgradesStatus = answer === null ? 'none' : 'ready';
      })
      .catch(() => {
        if (target === character) upgradesStatus = 'failed';
      });
  }

  function retryUpgrades(): void {
    invalidate(upgradesCacheKey(character));
    loadUpgrades();
  }

  $effect(() => {
    void character.key;
    untrack(() => {
      upgrades = null;
      guildHome = null;
      loadUpgrades();
      listMySims()
        .then((page) => (sims = page.rows))
        .catch(() => (sims = []));
      listMyReports()
        .then((page) => (reports = page.rows))
        .catch(() => (reports = []));
      const guildId = character.guild?.id;
      if (guildId !== undefined)
        fetchGuildHome(guildId)
          .then((home) => (guildHome = home))
          .catch(() => (guildHome = null));
    });
  });

  const classSlug = $derived(classSlugFromName(character.class ?? '') as ClassSlug);
  const bis = $derived(bisLink(character));
  const list = $derived(upgradesStatus === 'ready' ? (upgrades?.result.upgrades ?? []) : null);
  const total = $derived(upgrades?.result.totalGainDps ?? 0);
  const raid = $derived(guildHome?.next_raid ?? null);
  const links = $derived(guildLinks(me, character.guild?.id ?? -1, raid?.id ?? null));
  const raidTz = $derived(raid?.guild_timezone ?? viewerTz);
  const raidDay = $derived(raid === null ? '' : dateLabel(new Date(raid.starts_at), raidTz));
  const step = $derived(nextStep({ character, upgrades: list, raid, raidHref: links.raid, raidDay }));
  const talents = $derived(talentsToSpend(character));
  const specKey = $derived(homeSpecKeyFor(character));
  const sim = $derived(simRow(sims.find((row) => row.spec === specKey) ?? sims[0], character));
  const logs = $derived(ownLogs(reports, character));
  const needs = $derived(
    (guildHome?.standing?.needs_before_next_raid ?? []).map((need, index) => ({
      key: String(index),
      tone: 'warn' as const,
      title: need,
      sub: '',
      fix: null,
    })),
  );
</script>

<Page layout="split" testid="home-signed-in">
  {#snippet header()}
    <PageHeader title={character.name} tone={{ classSlug }} testid="home-head">
      {#snippet crest()}<ClassCrestRing characterClass={character.class ?? ''} size={CREST_PX} />{/snippet}
      {#snippet meta()}
        {#each metaFacts(character) as fact (fact)}<span>{fact}</span>{/each}
      {/snippet}
      {#snippet answer()}
        {#if list !== null && list.length > 0 && upgrades !== null}
          <Figure
            value={gainText(total)}
            unit="DPS"
            label={homeCopy.answerLabel(bandWords(upgrades.band))}
            align="end"
            testid="home-answer"
          />
          {#if bis !== null}<a class={LINK} href={bis.href}>{bis.label} ›</a>{/if}
        {/if}
      {/snippet}
    </PageHeader>
  {/snippet}
  {#snippet bar()}
    <HomeSummary text={step?.text} action={step?.action} />
  {/snippet}
  {#snippet rail()}
    <HomeRail {me} {character} {logs} guildHref={links.guild} {now} />
  {/snippet}
  {#if raid !== null}
    <Section
      title={homeCopy.nextRaid}
      action={links.raid === null ? undefined : { label: homeCopy.raidPage, href: links.raid }}
      testid="home-raid"
    >
      <Row
        title={raid.title}
        sub={`${raidTimeLine(raid.starts_at, raidTz, viewerTz)} · ${raid.counts.in} in, ${raid.counts.late} late`}
        href={links.raid ?? undefined}
        open
      >
        {#snippet trail()}
          {#if raid.own !== null}<Tag tone={raid.own.state === 'in' ? 'success' : 'gold'}
              >{raid.own.state.charAt(0).toUpperCase() + raid.own.state.slice(1)}</Tag
            >{/if}
        {/snippet}
      </Row>
      {#if needs.length > 0}<Checklist items={needs} testid="home-raid-needs" />{/if}
    </Section>
  {/if}
  <HomeUpgrades status={upgradesStatus} {upgrades} {bis} build={activeBuild.build} onRetry={retryUpgrades} />
  {#if talents !== null && talents.toSpend > 0}
    <Section title={homeCopy.talents} testid="home-talents">
      <Row
        title={homeCopy.talentsTitle(talents.toSpend)}
        sub={homeCopy.talentsSub(character.level ?? 0, talents.earned, talents.spent)}
        href={homeHrefs.planner(character.key)}
        open
      />
    </Section>
  {/if}
  <Section
    title={homeCopy.simulator}
    action={{ label: homeCopy.openSim, href: homeHrefs.sim }}
    testid="home-sim"
  >
    <Row title={sim.title} sub={sim.sub} href={sim.href} open>
      {#snippet trail()}
        {#if sim.dps !== null}<Figure value={sim.dps} unit="DPS" size="sm" align="end" />{/if}
      {/snippet}
    </Row>
  </Section>
</Page>
