<!-- web/src/components/Character.svelte -->
<!-- The character page (design/specs/2026-11-06-character.md), built from the ui/ primitives and
     the shared Paperdoll: the head (crest, the name in the class colour, the facts, the upgrade DPS
     available), the owner's Next line, the worn paperdoll on the page's one Plate, then the Forever
     Score, the best fight per encounter, the builds seen and the owner's extras as open Sections.
     A name the site does not hold gets the not-found page (C10). The page exists for characters
     who have never signed in: a public report is enough to have one. -->
<script lang="ts">
  import activeBuild from '../data/active-build.json';
  import { fetchMeOnce, type MeCharacter } from '../lib/account/api';
  import { slotsHaveGear } from '../lib/addon-export';
  import { bandForLevel } from '../lib/bis/hover';
  import { scoreUnitWord, slotScoreUnitFor } from '../lib/bis/tank-view';
  import { pickCells, wornCells } from '../lib/character/doll';
  import { loadGearContext, type GearContext } from '../lib/character/gear-context';
  import {
    bestRows,
    buildRows,
    characterHrefs,
    characterNextStep,
    characterPageCopy as copy,
    headFacts,
    pageSpec,
    plannerHref,
    topUpgrade,
    upgradeTotal,
  } from '../lib/character/page-view';
  import { STALE_SEPARATOR, syncText, syncView } from '../lib/character-header/sync';
  import { bisHref } from '../lib/character-profile';
  import { characterKey, nameFromSlug, parseCharacterPath, type CharacterPath } from '../lib/characters';
  import { onCurrentCharacterChange, onCurrentCharacterSwitch, readCurrent } from '../lib/current-character';
  import { viewerFromMe } from '../lib/coach/viewer';
  import { premiumBlockCopy } from '../lib/billing/copy';
  import { bandPricer, UNPRICED } from '../lib/gear/gain-pricer';
  import { notReadCopy } from '../lib/gear/copy';
  import { paperdollNote, unknownCount } from '../lib/gear/paperdoll';
  import { slotsFromCharacterGear } from '../lib/gear/paperdoll-slots';
  import { fetchItemTooltipModel } from '../lib/items/lookup';
  import { factionOfRace } from '../lib/race-faction';
  import { encounterSlug, fetchCharacter, RankingsError, type CharacterPage } from '../lib/rankings/api';
  import { classSlugFromName } from '../lib/report/tree-sizes';
  import { foreverOf, hasScore, ownCharacter } from '../lib/score/score';
  import { signedFigure, LINK, type ClassSlug } from '../lib/ui/primitives';
  import { REGION } from '../lib/ui/regions';
  import BisTooltipHost from './BisTooltipHost.svelte';
  import CharacterGear from './character/CharacterGear.svelte';
  import CharacterMissing from './character/CharacterMissing.svelte';
  import CharacterOwnerSections from './character/CharacterOwnerSections.svelte';
  import CharacterScore from './character/CharacterScore.svelte';
  import ClassCrestRing from './character/ClassCrestRing.svelte';
  import Figure from './ui/Figure.svelte';
  import LoadError from './ui/LoadError.svelte';
  import Notice from './ui/Notice.svelte';
  import Page from './ui/Page.svelte';
  import Paperdoll from './ui/Paperdoll.svelte';
  import PageHeader from './ui/PageHeader.svelte';
  import Plate from './ui/Plate.svelte';
  import Row from './ui/Row.svelte';
  import Section from './ui/Section.svelte';
  import Tag from './ui/Tag.svelte';

  let { path = null }: { path?: CharacterPath | null } = $props();

  const build = activeBuild.build;
  const CREST_PX = 56;
  const resolved = $derived(
    path ?? (typeof window === 'undefined' ? null : parseCharacterPath(window.location.pathname)),
  );

  let data = $state<CharacterPage | null>(null);
  /** `unlogged`: the API has no record of this character (a 404). */
  let status = $state<'loading' | 'ready' | 'failed' | 'missing' | 'unlogged'>('loading');
  let attempt = $state(0);

  $effect(() => {
    void attempt;
    const requested = resolved;
    if (requested === null) {
      status = 'missing';
      return;
    }
    status = 'loading';
    void fetchCharacter(requested)
      .then((result) => {
        if (resolved !== requested) return;
        data = result;
        status = 'ready';
      })
      .catch((thrown: unknown) => {
        if (resolved !== requested) return;
        status = thrown instanceof RankingsError && thrown.status === 404 ? 'unlogged' : 'failed';
      });
  });

  /** The viewer's own record of this character, their plan, and whether /v1/me has answered. */
  let own = $state<MeCharacter | null>(null);
  let owned = $state(false);
  let signedIn = $state(false);
  let historyEntitled = $state(false);
  let coachEntitled = $state(false);

  function isActive(requested: CharacterPath): boolean {
    const current = readCurrent();
    return (
      current?.source === 'armory' &&
      current.ref === characterKey(requested.region, requested.ruleset, requested.slug)
    );
  }
  $effect(() =>
    onCurrentCharacterChange(() => {
      if (resolved !== null) owned = isActive(resolved) || own !== null;
    }),
  );
  // Your character's page moves to the newly chosen character's; a character you looked up stays.
  $effect(() =>
    onCurrentCharacterSwitch((change) => {
      void import('../lib/character-selector/follow-page-run').then(({ followCharacterPage }) =>
        followCharacterPage(change),
      );
    }),
  );
  $effect(() => {
    const requested = resolved;
    own = null;
    owned = false;
    historyEntitled = false;
    coachEntitled = false;
    if (requested === null) return;
    owned = isActive(requested);
    void fetchMeOnce()
      .then((me) => {
        if (resolved !== requested) return;
        signedIn = me !== null && me !== undefined;
        own = ownCharacter(me?.characters, requested);
        historyEntitled = me?.entitlements?.history === true;
        coachEntitled = viewerFromMe(me).premium;
        owned = isActive(requested) || own !== null;
      })
      .catch(() => {
        own = null;
      });
  });

  /** The owner's sync line, re-read each minute while the tab lives. */
  let now = $state(new Date());
  $effect(() => {
    const tick = setInterval(() => (now = new Date()), 60_000);
    return () => clearInterval(tick);
  });
  const sync = $derived(own === null ? null : syncView(own, now, { needsGear: false }));

  const classSlug = $derived(
    data?.character.class ? (classSlugFromName(data.character.class) as ClassSlug) : null,
  );
  const level = $derived(typeof data?.character.level === 'number' ? data.character.level : null);
  const faction = $derived(data === null ? null : factionOfRace(data.character.race));
  const forever = $derived(foreverOf(data));
  const scored = $derived(hasScore(forever));
  const gear = $derived(data?.gear ?? []);
  const specBis = $derived(
    data === null
      ? null
      : bisHref(data.character.class, pageSpec(data, own), data.character.level, data.character.race),
  );

  /** The items file and the band (C2, C4); undefined while they load. */
  let context = $state<GearContext | undefined>(undefined);
  $effect(() => {
    const page = data;
    context = undefined;
    if (page === null) return;
    if ((page.gear ?? []).length === 0) {
      context = { table: null, band: null };
      return;
    }
    const input = { build, classSlug, gearSpec: page.gear_spec ?? null, level, faction };
    void loadGearContext(input).then((loaded) => {
      if (data === page) context = loaded;
    });
  });
  const table = $derived(context?.table ?? null);
  const band = $derived(context?.band ?? null);
  const pricer = $derived(
    table !== null && band !== null ? bandPricer(band, (id) => table.get(id)?.item) : UNPRICED,
  );
  const slots = $derived(slotsFromCharacterGear(gear, table, pricer));
  const unit = $derived(band === null ? 'DPS' : scoreUnitWord(slotScoreUnitFor(band)));
  /** The export carried no gear and no talents (W.3): the band's picks instead of empty slots. */
  const unread = $derived(
    gear.length > 0 &&
      !slotsHaveGear(gear) &&
      (own?.build?.talents === undefined || !own.build.talents.points.some((points) => points > 0)),
  );
  const total = $derived(context === undefined || unread ? 0 : upgradeTotal(slots));
  const cells = $derived(
    unread ? pickCells(gear, table, classSlug ?? '') : wornCells(slots, classSlug ?? '', owned),
  );
  const note = $derived(
    context === undefined || unread ? null : paperdollNote(unknownCount(slots), table === null),
  );
  const step = $derived.by(() => {
    if (own === null || context === undefined) return null;
    if (unread)
      return { text: `${notReadCopy.lead} ${notReadCopy.steps}`, action: null, tone: 'neutral' as const };
    const top = topUpgrade(slots);
    const pick = top === null ? null : gear.find((slot) => slot.slot === top.slot)?.pick;
    const next = characterNextStep({
      own,
      top: pick ? { name: pick.name, source: pick.source } : null,
      bisHref: specBis,
    });
    return { ...next, tone: 'gold' as const };
  });
  const bandLevel = $derived(level === null ? null : bandForLevel(level));
  /** Nothing to show a viewer who does not hold the character: the not-found page (C10). */
  const bare = $derived(
    data !== null &&
      !scored &&
      gear.length === 0 &&
      data.best.length === 0 &&
      data.history.length === 0 &&
      !owned,
  );
  const unloggedName = $derived(resolved === null ? '' : nameFromSlug(decodeURIComponent(resolved.slug)));
</script>

{#snippet headSkeleton()}
  <div class="character-head-skeleton" data-region={REGION.head} role="status" aria-busy="true">
    <span class="skeleton-block h-14 w-14 rounded-full" aria-hidden="true"></span>
    <span class="skeleton-block h-8 w-[280px] max-w-full" aria-hidden="true"></span>
  </div>
{/snippet}

{#if status === 'missing'}
  <Page layout="ledger" testid="character-missing">
    {#snippet header()}<PageHeader title={copy.findTitle} />{/snippet}
    <p class="text-[15px]">
      <a class={LINK} href="/characters#lookup" data-testid="character-missing-lookup">{copy.lookup} ›</a>
    </p>
  </Page>
{:else if status === 'loading'}
  <Page layout="ledger" testid="character-skeleton">
    {#snippet header()}{@render headSkeleton()}{/snippet}
    <Section title={copy.gear}>
      <Plate label={copy.gear}>
        <Paperdoll cells={[]} mode="worn" {build} loading label={copy.gear} testid="character-doll" />
      </Plate>
    </Section>
  </Page>
{:else if status === 'failed'}
  <Page layout="ledger" testid="character-failed">
    {#snippet header()}<PageHeader title={unloggedName} />{/snippet}
    <Plate flush={false} label="Character">
      <LoadError message={copy.failed} onRetry={() => (attempt += 1)} testid="character-error" />
    </Plate>
  </Page>
{:else if resolved !== null && (status === 'unlogged' || bare)}
  <CharacterMissing path={resolved} name={data?.character.name ?? unloggedName} {signedIn} />
{:else if data !== null && resolved !== null}
  <Page layout="ledger" testid="character">
    {#snippet header()}
      <PageHeader
        title={data.character.name}
        tone={classSlug === null ? 'strong' : { classSlug }}
        testid="character-head"
      >
        {#snippet crest()}<ClassCrestRing
            characterClass={data.character.class ?? ''}
            size={CREST_PX}
          />{/snippet}
        {#snippet meta()}
          {#each headFacts(data, resolved, own) as fact (fact.key)}
            {#if fact.href}<a class={LINK} href={fact.href} data-testid={`character-fact-${fact.key}`}
                >{fact.text}</a
              >{:else}<span data-testid={`character-fact-${fact.key}`}>{fact.text}</span>{/if}
          {/each}
          {#if data.supporter === true}<Tag testid="supporter-chip">{premiumBlockCopy.supporterChip}</Tag
            >{/if}
          {#if sync !== null}<span data-testid="character-sync"
              >{syncText(sync)}{#if sync.staleTail !== null}{STALE_SEPARATOR}<span class="text-wipe"
                  >{sync.staleTail}</span
                >{/if}</span
            >{/if}
        {/snippet}
        {#snippet answer()}
          {#if total > 0}
            <span class="character-answer-figures">
              {#if scored}<Figure
                  value={(forever.score ?? 0).toLocaleString('en-US')}
                  label={copy.scoreLabel}
                  size="sm"
                  align="end"
                />{/if}
              <Figure
                value={signedFigure(total)}
                {unit}
                label={copy.answerLabel}
                size="big"
                align="end"
                testid="character-answer"
              />
            </span>
            <span class="character-answer-links">
              <a class={LINK} href={characterHrefs.sim}>{copy.openSim} ›</a>
              <a class={LINK} href={plannerHref(own)}>{copy.openPlanner} ›</a>
            </span>
          {/if}
        {/snippet}
      </PageHeader>
    {/snippet}
    {#snippet bar()}
      {#if step !== null}
        <div class="character-summary" data-region={REGION.summary} data-testid="character-next-step">
          <Notice text={step.text} action={step.action ?? undefined} tone={step.tone} />
        </div>
      {/if}
    {/snippet}
    {#if gear.length > 0}
      <CharacterGear
        title={unread ? notReadCopy.bestInSlot(bandLevel) : copy.gear}
        meta={unread || bandLevel === null ? undefined : copy.gearLevel(bandLevel)}
        {cells}
        {gear}
        {band}
        {table}
        {unit}
        {build}
        bisHref={specBis}
        loading={context === undefined}
        {note}
      />
    {/if}
    <CharacterScore path={resolved} {forever} {owned} fought={data.best.length > 0} />
    {#if data.best.length > 0}
      <Section
        title={copy.best}
        action={{ label: copy.allFights(data.history.length), href: characterHrefs.logs }}
        testid="character-best"
      >
        {#each bestRows(data.best) as row (row.key)}
          <Row
            title={row.boss}
            sub={row.sub}
            href={characterHrefs.encounter(encounterSlug(row.encounter))}
            open
          >
            {#snippet trail()}
              {#if row.percentile !== null}<span
                  class="w-8 text-right text-[15px] leading-5 font-bold tabular-nums"
                  style:color={row.color}
                  aria-label={`${row.percentile} percentile`}>{row.percentile}</span
                >{/if}
              <Figure value={row.value} unit={row.unit} size="sm" align="end" />
            {/snippet}
          </Row>
        {/each}
      </Section>
    {/if}
    {#if data.builds_seen.length > 0}
      <Section title={copy.builds} testid="character-builds">
        {#each buildRows(data.character.class, data.builds_seen) as row (row.key)}
          <Row title={row.split} sub={row.sub} href={plannerHref(own)} open />
        {/each}
      </Section>
    {/if}
    {#if own !== null && scored}
      <CharacterOwnerSections path={resolved} history={historyEntitled} coach={coachEntitled} />
    {/if}
  </Page>
  {#if classSlug !== null && gear.length > 0}
    <BisTooltipHost
      {build}
      modelsElementId="gear-tooltip-models"
      resolve={(id) => fetchItemTooltipModel(build, classSlug ?? '', id)}
    />
  {/if}
{/if}

<style>
  .character-head-skeleton {
    display: flex;
    align-items: center;
    gap: 16px;
    min-height: 96px;
    padding: 20px var(--page-inset) 16px;
  }
  .character-summary {
    padding-inline: var(--page-inset);
  }
  .character-answer-figures {
    display: flex;
    align-items: flex-end;
    justify-content: flex-end;
    gap: 24px;
  }
  .character-answer-links {
    display: flex;
    flex-wrap: wrap;
    justify-content: flex-end;
    gap: 16px;
    font-size: 14px;
  }
</style>
