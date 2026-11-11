<!-- web/src/components/home/HomeCharacter.svelte -->
<!-- The signed-in home for one character (round 5, design/mocks/directions/png/g-home-*.png): the
     title band on the camp-road art (crest, name in the class colour, one line, the sim on import),
     the raid plaque (the guild's next raid, else the first raids), Latest on Forever with Your
     upgrades beside it, then Guild sims with Your spec this week beside it (or, with no guild, Your
     spec this week alone). Each read has its own state;
     none blocks another. -->
<script lang="ts">
  import { untrack } from 'svelte';
  import type { Me, MeCharacter } from '../../lib/account/api';
  import { fetchGuildHome, type GuildHome } from '../../lib/guild/api';
  import { fetchBisFileStrict } from '../../lib/bis/hover';
  import { cachedUpgradesFor, homeSpecKeyFor, type CachedUpgrades } from '../../lib/home/upgrades-loader';
  import { bisLink, guildLinks, whoLine } from '../../lib/home/home-view';
  import { homeBCopy, homeBHrefs, simFigure } from '../../lib/home/home-b-view';
  import { EMPTY_HOME, fetchHomeAnswer, type HomeAnswer } from '../../lib/home/home-answer';
  import { signupRole } from '../../lib/home/plaque';
  import { factsCopy, guildSummaryFacts, standingFacts } from '../../lib/home/home-facts';
  import { bandOf, type FeedRow } from '../../lib/home/feed';
  import type { RaidFact } from '../../lib/home/raid-facts';
  import { classSlugFromName } from '../../lib/report/tree-sizes';
  import type { ClassSlug } from '../../lib/ui/primitives';
  import { LINK } from '../../lib/ui/primitives';
  import activeBuild from '../../data/active-build.json';
  import Crest from '../ui/Crest.svelte';
  import Figure from '../ui/Figure.svelte';
  import PageHeader from '../ui/PageHeader.svelte';
  import HomeFeed from './HomeFeed.svelte';
  import HomeFirstRaids from './HomeFirstRaids.svelte';
  import HomeGrid from './HomeGrid.svelte';
  import HomeGuildSims from './HomeGuildSims.svelte';
  import HomePlaqueSkeleton from './HomePlaqueSkeleton.svelte';
  import HomeSpecWeek from './HomeSpecWeek.svelte';
  import HomeUpgrades from './HomeUpgrades.svelte';
  import HomeRaidPlaque from './HomeRaidPlaque.svelte';
  import HomeFacts from './HomeFacts.svelte';

  let {
    me,
    character,
    floor,
    raidsOpen,
    raids,
  }: { me: Me; character: MeCharacter; floor: FeedRow[]; raidsOpen: string | null; raids: RaidFact[] } =
    $props();

  const CREST_PX = 48;

  let upgrades = $state<CachedUpgrades | null>(null);
  let answer = $state<HomeAnswer>(EMPTY_HOME);
  let answered = $state(false);
  let guildHome = $state<GuildHome | null>(null);
  /** The guild read is in flight: the plaque holds its place until it knows which plaque to draw. */
  let guildPending = $state(false);

  $effect(() => {
    void character.key;
    untrack(() => {
      upgrades = null;
      guildHome = null;
      guildPending = false;
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
      if (guildId !== undefined) {
        guildPending = true;
        fetchGuildHome(guildId)
          .then((home) => {
            if (target === character) guildHome = home;
          })
          .catch(() => {
            if (target === character) guildHome = null;
          })
          .finally(() => {
            if (target === character) guildPending = false;
          });
      }
    });
  });

  const classSlug = $derived(classSlugFromName(character.class ?? '') as ClassSlug);
  const bis = $derived(bisLink(character));
  const raid = $derived(guildHome?.next_raid ?? null);
  const links = $derived(guildLinks(me, character.guild?.id ?? -1, raid?.id ?? null));
  const sim = $derived(answered ? simFigure(answer.import_sim) : null);
  const specKey = $derived(homeSpecKeyFor(character));
  const band = $derived(bandOf(character.level));
  const standing = $derived(
    guildHome === null ? [] : standingFacts(guildHome.standing, character.guild, guildHome.attendance_source),
  );
  const guildFacts = $derived(
    guildHome === null ? [] : [...guildSummaryFacts(guildHome.summary), ...standing],
  );
</script>

<div data-testid="home-signed-in">
  <PageHeader title={character.name} tone={{ classSlug }} art="hero" testid="home-head">
    {#snippet crest()}<Crest cls={character.class ?? ''} size={CREST_PX} />{/snippet}
    {#snippet meta()}
      <span
        >{whoLine(character)}{#if character.guild}&nbsp;·
          <span class="whitespace-nowrap">&lt;{character.guild.name}&gt;</span
          >{/if}{#if character.item_level !== undefined}&nbsp;·
          <span class="whitespace-nowrap">{factsCopy.itemLevel(character.item_level)}</span>{/if}</span
      >
    {/snippet}
    {#snippet answer()}
      <div class="home-answer" data-testid="home-answer">
        {#if sim !== null}
          <span data-testid="home-figure-sim">
            <Figure value={sim.value} unit={sim.unit} tone={sim.tone} align="end" />
          </span>
          {#if sim.link}<a class="{LINK} text-[13px]" href={sim.link.href}>{sim.link.label} ›</a>{/if}
        {/if}
      </div>
    {/snippet}
  </PageHeader>
  <HomeGrid>
    {#snippet lead()}
      {#if guildPending}
        <HomePlaqueSkeleton />
      {:else if raid !== null && guildHome !== null && character.guild}
        <HomeRaidPlaque
          guildId={character.guild.id}
          guild={character.guild.name}
          {raid}
          needs={guildHome.standing?.needs_before_next_raid ?? []}
          {standing}
          characterKey={character.key}
          role={signupRole(raid.own, character)}
          raidHref={links.raid}
        />
      {:else}
        <HomeFirstRaids
          {raidsOpen}
          {raids}
          viewer={character.class ? { slug: classSlug, name: character.class } : undefined}
          build={activeBuild.build}
        >
          {#snippet foot()}
            {#if character.guild === undefined}
              {homeBCopy.findGuildSub} <a class={LINK} href={homeBHrefs.guilds}>{homeBCopy.guildsLink} ›</a>
            {:else if guildHome !== null}
              {homeBCopy.guildNone(character.guild.name)}
              {#if links.guild !== null}· <a class={LINK} href={links.guild}>{homeBCopy.raidPageLink} ›</a
                >{/if}
              {#if guildFacts.length > 0}<div class="mt-3">
                  <HomeFacts facts={guildFacts} testid="home-guild-facts" />
                </div>{/if}
            {/if}
          {/snippet}
        </HomeFirstRaids>
      {/if}
    {/snippet}
    {#snippet rail()}<HomeUpgrades {upgrades} {bis} build={activeBuild.build} />{/snippet}
    {#snippet foot()}
      {#if answer.guild_sims !== null}
        <div class="home-foot-pair">
          <HomeGuildSims sims={answer.guild_sims} />
          {#if answer.spec_week !== null}<HomeSpecWeek week={answer.spec_week} movers={answer.movers} />{/if}
        </div>
      {:else if answer.spec_week !== null}
        <HomeSpecWeek week={answer.spec_week} movers={answer.movers} />
      {/if}
    {/snippet}
    <HomeFeed
      signedIn
      spec={specKey}
      {band}
      raidOnPlaque={raid !== null}
      guild={character.guild?.name}
      {floor}
      build={activeBuild.build}
    />
  </HomeGrid>
</div>

<style>
  /* The head's answer column sets the side (left under the title on a phone, right from 900). */
  .home-answer {
    display: flex;
    flex-direction: column;
    align-items: inherit;
    gap: 2px;
  }
  /* Guild sims with Your spec this week beside it, on the columns of the feed and the rail above. */
  .home-foot-pair {
    display: grid;
    grid-template-columns: minmax(0, 1fr);
    align-items: start;
    gap: var(--gap);
  }
  @media (min-width: 1100px) {
    .home-foot-pair {
      grid-template-columns: minmax(0, 1fr) var(--home-rail);
    }
  }
</style>
