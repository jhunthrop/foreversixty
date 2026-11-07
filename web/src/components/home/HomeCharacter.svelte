<!-- web/src/components/home/HomeCharacter.svelte -->
<!-- The signed-in home for one character, direction B (spec 2026-11-07 §1): the compact head with its
     three figures (the sim on import, what the upgrades add, the levels to the next band), the
     countdowns (the first raids, the guild's next event), then "Forever right now" with the rail
     beside it. Each read has its own state; none blocks another. -->
<script lang="ts">
  import { untrack } from 'svelte';
  import type { Me, MeCharacter } from '../../lib/account/api';
  import { fetchGuildHome, type GuildHome } from '../../lib/guild/api';
  import { fetchBisFileStrict } from '../../lib/bis/hover';
  import { cachedUpgradesFor, homeSpecKeyFor, type CachedUpgrades } from '../../lib/home/upgrades-loader';
  import { bisLink, guildLinks, metaFacts } from '../../lib/home/home-view';
  import { bandFigure, gainFigure, simFigure, type HeadFigure } from '../../lib/home/home-b-view';
  import { EMPTY_HOME, fetchHomeAnswer, type HomeAnswer } from '../../lib/home/home-answer';
  import { bandOf, type FeedRow } from '../../lib/home/feed';
  import { classSlugFromName } from '../../lib/report/tree-sizes';
  import type { ClassSlug } from '../../lib/ui/primitives';
  import { LINK } from '../../lib/ui/primitives';
  import activeBuild from '../../data/active-build.json';
  import ClassCrestRing from '../character/ClassCrestRing.svelte';
  import Figure from '../ui/Figure.svelte';
  import Page from '../ui/Page.svelte';
  import PageHeader from '../ui/PageHeader.svelte';
  import HomeCountdowns from './HomeCountdowns.svelte';
  import HomeFeed from './HomeFeed.svelte';
  import HomeRail from './HomeRail.svelte';

  let {
    me,
    character,
    floor,
    raidsOpen,
  }: { me: Me; character: MeCharacter; floor: FeedRow[]; raidsOpen: string | null } = $props();

  const CREST_PX = 56;

  let upgrades = $state<CachedUpgrades | null>(null);
  let answer = $state<HomeAnswer>(EMPTY_HOME);
  let answered = $state(false);
  let guildHome = $state<GuildHome | null>(null);

  $effect(() => {
    void character.key;
    untrack(() => {
      upgrades = null;
      guildHome = null;
      answer = EMPTY_HOME;
      answered = false;
      const target = character;
      cachedUpgradesFor(target, fetchBisFileStrict)
        .then((found) => {
          if (target === character) upgrades = found;
        })
        .catch(() => {
          if (target === character) upgrades = null;
        });
      fetchHomeAnswer(target.key)
        .then((found) => {
          if (target === character) answer = found;
        })
        .catch(() => {
          if (target === character) answer = EMPTY_HOME;
        })
        .finally(() => {
          if (target === character) answered = true;
        });
      const guildId = target.guild?.id;
      if (guildId !== undefined)
        fetchGuildHome(guildId)
          .then((home) => (guildHome = home))
          .catch(() => (guildHome = null));
    });
  });

  const classSlug = $derived(classSlugFromName(character.class ?? '') as ClassSlug);
  const bis = $derived(bisLink(character));
  const list = $derived(upgrades?.result.upgrades ?? []);
  const raid = $derived(guildHome?.next_raid ?? null);
  const links = $derived(guildLinks(me, character.guild?.id ?? -1, raid?.id ?? null));
  const onBis = $derived(upgrades === null ? null : upgrades.result.alreadyBis.length);
  const figures = $derived(
    [
      answered ? simFigure(answer.import_sim) : null,
      gainFigure(upgrades?.result.totalGainDps ?? 0, list.length),
      bandFigure(character.level, onBis),
    ].filter((figure): figure is HeadFigure => figure !== null),
  );
  const specKey = $derived(homeSpecKeyFor(character));
  const band = $derived(bandOf(character.level));
</script>

<Page layout="split" testid="home-signed-in">
  {#snippet header()}
    <PageHeader title={character.name} tone={{ classSlug }} testid="home-head">
      {#snippet crest()}<ClassCrestRing characterClass={character.class ?? ''} size={CREST_PX} />{/snippet}
      {#snippet meta()}
        {#each metaFacts(character) as fact (fact)}<span>{fact}</span>{/each}
        {#if character.guild}<span>&lt;{character.guild.name}&gt;</span>{/if}
      {/snippet}
      {#snippet answer()}
        <div class="home-figures" data-testid="home-answer">
          {#each figures as figure (figure.key)}
            <span class="home-figure" data-testid={`home-figure-${figure.key}`}>
              <Figure
                value={figure.value}
                unit={figure.unit}
                label={figure.label}
                caption={figure.caption}
                tone={figure.tone}
                align="end"
              />
              {#if figure.link}<a class="{LINK} text-[13px]" href={figure.link.href}>{figure.link.label} ›</a
                >{/if}
            </span>
          {/each}
        </div>
      {/snippet}
    </PageHeader>
  {/snippet}
  {#snippet bar()}
    <HomeCountdowns
      {raidsOpen}
      build={activeBuild.build}
      guild={character.guild?.name ?? null}
      {raid}
      raidHref={links.raid}
      guildHref={links.guild}
    />
  {/snippet}
  {#snippet rail()}
    <HomeRail {upgrades} {bis} {answer} guildHref={links.guild} build={activeBuild.build} />
  {/snippet}
  <HomeFeed
    signedIn
    spec={specKey}
    {band}
    inGuild={character.guild !== undefined}
    {floor}
    build={activeBuild.build}
  />
</Page>

<style>
  .home-figures {
    display: flex;
    flex-wrap: wrap;
    justify-content: flex-end;
    gap: 12px 36px;
  }
  .home-figure {
    display: flex;
    flex-direction: column;
    align-items: flex-end;
    gap: 2px;
  }
</style>
