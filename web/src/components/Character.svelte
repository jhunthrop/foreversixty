<!-- web/src/components/Character.svelte -->
<!-- The character page (design/specs/2026-11-06-character.md), built from the ui/ primitives and
     the shared Paperdoll: the head (crest, the name in the class colour, the facts, the upgrade DPS
     available), the owner's Next line, the worn paperdoll on the page's one Plate, then the Forever
     Score, the best fight per encounter, the builds seen and the owner's extras as open Sections.
     A name the site does not hold gets the not-found page (C10). The page exists for characters
     who have never signed in: a public report is enough to have one. -->
<script lang="ts">
  import activeBuild from '../data/active-build.json';
  import { fetchMeOnce, type Me, type MeCharacter } from '../lib/account/api';
  import { slotsHaveGear } from '../lib/addon-export';
  import { bandForLevel } from '../lib/bis/hover';
  import { scoreUnitWord, slotScoreUnitFor } from '../lib/bis/tank-view';
  import { pickCells, wornCells } from '../lib/character/doll';
  import { loadGearContext, type GearContext } from '../lib/character/gear-context';
  import {
    bestRows,
    buildRows,
    characterHrefs,
    characterWaiting,
    characterPageCopy as copy,
    headFacts,
    pageSpec,
    plannerHref,
    topUpgrade,
    upgradeTotal,
    type FightRowView,
    type WaitingFact,
  } from '../lib/character/page-view';
  import { upgradeRows } from '../lib/character/upgrades';
  import { STALE_SEPARATOR, syncText, syncView } from '../lib/character-header/sync';
  import { bisHref, specClassLabel } from '../lib/character-profile';
  import { characterKey, nameFromSlug, parseCharacterPath, type CharacterPath } from '../lib/characters';
  import { onCurrentCharacterChange, onCurrentCharacterSwitch, readCurrent } from '../lib/current-character';
  import { viewerFromMe } from '../lib/coach/viewer';
  import { simViewedCharacterHref } from '../lib/handoff-links';
  import { button as buttonClass } from '../lib/ui/button';
  import { bandPricer, UNPRICED } from '../lib/gear/gain-pricer';
  import { notReadCopy } from '../lib/gear/copy';
  import { uiCopy } from '../lib/ui/copy';
  import { paperdollNote, unknownCount } from '../lib/gear/paperdoll';
  import { slotsFromCharacterGear } from '../lib/gear/paperdoll-slots';
  import { fetchItemTooltipModel } from '../lib/items/lookup';
  import { factionOfRace } from '../lib/race-faction';
  import { encounterSlug, fetchCharacter, RankingsError, type CharacterPage } from '../lib/rankings/api';
  import { classSlugFromName } from '../lib/report/tree-sizes';
  import { foreverOf, hasScore, ownCharacter } from '../lib/score/score';
  import { signedFigure, LINK, PLATE_SURFACE, type ClassSlug } from '../lib/ui/primitives';
  import { REGION } from '../lib/ui/regions';
  import BisTooltipHost from './BisTooltipHost.svelte';
  import CharacterActivity from './character/CharacterActivity.svelte';
  import CharacterGear from './character/CharacterGear.svelte';
  import CharacterMissing from './character/CharacterMissing.svelte';
  import CharacterOwnerSections from './character/CharacterOwnerSections.svelte';
  import CharacterScore from './character/CharacterScore.svelte';
  import CharacterStanding from './character/CharacterStanding.svelte';
  import ClassCrestRing from './character/ClassCrestRing.svelte';
  import Figure from './ui/Figure.svelte';
  import LoadError from './ui/LoadError.svelte';
  import Notice from './ui/Notice.svelte';
  import Page from './ui/Page.svelte';
  import PaperdollFrame from './ui/PaperdollFrame.svelte';
  import PageHeader from './ui/PageHeader.svelte';
  import Plate from './ui/Plate.svelte';
  import ItemIcon from './ui/ItemIcon.svelte';
  import Ordered from './ui/Ordered.svelte';
  import Row from './ui/Row.svelte';
  import Section from './ui/Section.svelte';
  import Spark from './ui/Spark.svelte';

  let { path = null }: { path?: CharacterPath | null } = $props();

  const build = activeBuild.build;
  const CREST_PX = 56;
  /** The band's crest: 44 px on a phone, 56 from 768 (the mock's title band). */
  const CREST_CLASS = 'size-11 md:size-14';
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
  /** The viewer's /v1/me, for the owner's standing line (its guild links). */
  let meRead = $state<Me | null>(null);
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
        meRead = me ?? null;
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
  /** The owner's waiting facts; an export that carried nothing reads one neutral line instead. */
  const waiting = $derived.by(() => {
    if (own === null || context === undefined || unread) return [];
    const top = topUpgrade(slots);
    const pick = top === null ? null : gear.find((slot) => slot.slot === top.slot)?.pick;
    const ref = top?.upgrade?.kind === 'gain' ? top.upgrade.pick : null;
    return characterWaiting({
      own,
      top: pick
        ? { name: pick.name, source: pick.source, icon: ref?.icon ?? null, quality: ref?.quality ?? 1 }
        : null,
      bisHref: specBis,
    });
  });
  const notReadLine = $derived(own !== null && context !== undefined && unread);
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

<!-- The upgrade's "Upgrade:" lead drops on a phone, so the item and its source fit one line. -->
{#snippet waitingText(fact: WaitingFact)}<span class="text-gold block truncate font-semibold"
    >{#if fact.lead !== undefined}<span class="max-md:hidden">{fact.lead}</span>{/if}{fact.text}</span
  >{/snippet}

{#snippet trendCells(row: FightRowView)}{#if row.since !== ''}<span
      class="text-muted text-[13px] whitespace-nowrap tabular-nums">{row.since}</span
    >{/if}{#if row.trend.length > 1}<span class="block w-12 shrink-0" aria-hidden="true"
      ><Spark values={row.trend} /></span
    >{/if}{/snippet}

{#snippet headSkeleton()}
  <div role="status" class="contents">
    <PageHeader title={unloggedName} pending={unloggedName === ''} testid="character-head-loading">
      {#snippet crest()}<span class="skeleton-block block size-14 rounded-full" aria-hidden="true"
        ></span>{/snippet}
      {#snippet meta()}<span class="skeleton-block block h-4 w-[320px] max-w-full" aria-hidden="true"
        ></span>{/snippet}
    </PageHeader>
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
  <Page layout="split" panels splitFrom={1280} testid="character-skeleton">
    {#snippet header()}{@render headSkeleton()}{/snippet}
    <Section title={copy.gear} panel frame="framed">
      <div class="flex justify-center" role="status" aria-busy="true">
        <span class="sr-only">{uiCopy.loading}</span>
        <div class="w-full max-w-[640px]">
          <PaperdollFrame cells={[]} {build} title="" label={copy.gear} testid="character-doll" />
        </div>
      </div>
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
  <!-- A workbench (h-audit 2026-11-09): main and the rail side by side from 1280; below it one flow,
       the gear, then Score and Standing, then the fights (each ui/Ordered), the rail sticky beside main. -->
  <Page layout="workbench" panels stickyRail testid="character">
    {#snippet header()}
      <PageHeader
        title={data.character.name}
        tone={classSlug === null ? 'strong' : { classSlug }}
        testid="character-head"
      >
        {#snippet crest()}<ClassCrestRing
            characterClass={data.character.class ?? ''}
            size={CREST_PX}
            class={CREST_CLASS}
          />{/snippet}
        {#snippet meta()}
          {#each headFacts(data, resolved, own) as fact (fact.key)}
            {#if fact.href}<a class={LINK} href={fact.href} data-testid={`character-fact-${fact.key}`}
                >{fact.text}</a
              >{:else}<span data-testid={`character-fact-${fact.key}`}>{fact.text}</span>{/if}
          {/each}
          {#if sync !== null}<span data-testid="character-sync"
              >{syncText(sync)}{#if sync.staleTail !== null}{STALE_SEPARATOR}<span class="text-wipe"
                  >{sync.staleTail}</span
                >{/if}</span
            >{/if}
        {/snippet}
        {#snippet answer()}
          <span class="flex flex-wrap items-center gap-x-5 gap-y-2">
            {#if !unread}<a
                class={buttonClass('secondary', 'sm')}
                href={simViewedCharacterHref(characterKey(resolved.region, resolved.ruleset, resolved.slug))}
                data-testid="character-sim">{copy.simThis}</a
              >{/if}
            {#if total > 0}<span class="shrink-0 whitespace-nowrap"
                ><Figure
                  value={signedFigure(total)}
                  unit={copy.answerUnit(unit)}
                  align="end"
                  testid="character-answer"
                /></span
              >{/if}
          </span>
        {/snippet}
      </PageHeader>
    {/snippet}
    {#snippet bar()}
      {#if notReadLine || waiting.length > 0}
        <div class="character-summary" data-region={REGION.summary}>
          {#if notReadLine}<div data-testid="character-next-step">
              <Notice text={`${notReadCopy.lead} ${notReadCopy.steps}`} tone="neutral" />
            </div>{:else if waiting.length > 0}
            <!-- The waiting facts on a plate (h-audit 2026-11-09): one row each, the upgrade led by its
                 item in its quality ring, one line that truncates on a phone. -->
            <div class="{PLATE_SURFACE} mt-4 w-full px-4" data-testid="character-next-step">
              {#each waiting as fact (fact.key)}
                {@const item = fact.item}
                {#if item !== undefined}
                  <Row
                    title={fact.text}
                    href={fact.href ?? undefined}
                    density="slim"
                    open
                    testid={`character-waiting-${fact.key}`}
                  >
                    {#snippet lead()}<ItemIcon
                        icon={item.icon ?? undefined}
                        quality={item.quality}
                        {build}
                        size={28}
                        eager
                      />{/snippet}
                    {#snippet titleContent()}{@render waitingText(fact)}{/snippet}
                  </Row>
                {:else}
                  <Row
                    title={fact.text}
                    href={fact.href ?? undefined}
                    density="slim"
                    open
                    testid={`character-waiting-${fact.key}`}
                  >
                    {#snippet titleContent()}{@render waitingText(fact)}{/snippet}
                  </Row>
                {/if}
              {/each}
            </div>{/if}
        </div>
      {/if}
    {/snippet}
    {#snippet rail()}
      {#if data.score_hidden !== true}
        <Ordered order={2}
          ><CharacterScore
            path={resolved}
            {forever}
            {owned}
            fought={data.best.length > 0}
            history={data.score_history ?? []}
          /></Ordered
        >
      {/if}
      {#if own !== null && meRead !== null}
        <Ordered order={3}
          ><CharacterStanding me={meRead} {own} specName={pageSpec(data, own)} {now} /></Ordered
        >
      {/if}
      {#if data.builds_seen.length > 0}
        <Ordered order={6}
          ><Section title={copy.builds} panel testid="character-builds">
            {#each buildRows(data.character.class, data.builds_seen) as row (row.key)}
              <Row title={row.split} sub={row.sub} href={plannerHref(own)} density="slim" open />
            {/each}
          </Section></Ordered
        >
      {/if}
      <!-- The rail carries the activity beside the gear and fights (h-audit 2026-11-09): main runs
           long with the paperdoll, so the rail takes the feed instead of ending under Builds seen. -->
      <Ordered order={7}
        ><CharacterActivity
          characterKey={characterKey(resolved.region, resolved.ruleset, resolved.slug)}
          characterName={data.character.name}
          {signedIn}
          {build}
          {now}
        /></Ordered
      >
    {/snippet}
    {#if gear.length > 0}
      <Ordered order={1} span
        ><CharacterGear
          sectionTitle={unread ? notReadCopy.bestInSlot(bandLevel) : copy.gear}
          title={specClassLabel(data.character.class, pageSpec(data, own))}
          meta={level === null ? undefined : copy.levelMeta(level)}
          avatar={data.character.avatar_url ?? own?.avatar_url ?? null}
          {cells}
          {gear}
          {band}
          {table}
          upgrades={unread ? [] : upgradeRows(slots, classSlug ?? '')}
          {unit}
          {build}
          bisHref={specBis}
          loading={context === undefined}
          {note}
        /></Ordered
      >
    {/if}
    {#if data.best.length > 0}
      <Ordered order={4} span
        ><Section
          title={copy.best}
          panel
          action={{ label: copy.allFights(data.history.length), href: characterHrefs.logs }}
          testid="character-best"
        >
          {#each bestRows(data.best, data.history, pageSpec(data, own)) as row (row.key)}
            <Row
              title={row.boss}
              sub={row.sub}
              href={characterHrefs.encounter(encounterSlug(row.encounter))}
              density="slim"
              open
            >
              <!-- The delta and the spark: in the trail from 768, their own line under the boss below it
                   (h-audit 2026-11-09: at 390 they wrapped the row to three lines against the parse). -->
              {#snippet subContent()}{row.sub}{#if row.since !== '' || row.trend.length > 1}<span
                    class="mt-1 flex items-center gap-3 md:hidden">{@render trendCells(row)}</span
                  >{/if}{/snippet}
              {#snippet trail()}
                <span class="flex items-center gap-3 max-md:hidden">{@render trendCells(row)}</span>
                {#if row.percentile !== null}<span
                    class="w-8 text-right text-[15px] leading-5 font-bold tabular-nums"
                    style:color={row.color}
                    aria-label={`${row.percentile} percentile`}>{row.percentile}</span
                  >{/if}
                <Figure value={row.value} unit={row.unit} size="sm" align="end" />
              {/snippet}
            </Row>
          {/each}
          {#if own !== null && scored}
            <CharacterOwnerSections
              path={resolved}
              history={historyEntitled}
              coach={coachEntitled}
              fights={data.history.length}
              inPanel
            />
          {/if}
        </Section></Ordered
      >
    {/if}
    {#if own !== null && scored && data.best.length === 0}
      <Ordered order={5} span
        ><CharacterOwnerSections
          path={resolved}
          history={historyEntitled}
          coach={coachEntitled}
          fights={data.history.length}
        /></Ordered
      >
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
  /* The waiting facts and the standing links share one wrapping row. */
  .character-summary {
    display: flex;
    flex-wrap: wrap;
    align-items: baseline;
    column-gap: 20px;
    padding-inline: var(--page-inset);
  }
</style>
