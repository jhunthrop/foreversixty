<!-- web/src/components/Character.svelte -->
<!-- Every ranked fight this character has, their best per encounter, and the builds they
     have been seen in at pull. The page exists for characters who have never signed in:
     a public report is enough to have one, which is the spec's position. -->
<script lang="ts">
  import CharacterFightTable from './character/CharacterFightTable.svelte';
  import CharacterIdentity from './character/CharacterIdentity.svelte';
  import HistoryMount from './character/HistoryMount.svelte';
  import HistoryReserve from './character/HistoryReserve.svelte';
  import ProfileDescriptor from './character/ProfileDescriptor.svelte';
  import CharacterNotFound from './CharacterNotFound.svelte';
  import CoachHistoryList from './coach/CoachHistoryList.svelte';
  import CharacterUnlogged from './CharacterUnlogged.svelte';
  import { CHARACTER_LOADING_MIN_H } from '../lib/character-layout';
  import { fetchMeOnce, type MeCharacter } from '../lib/account/api';
  import { viewerFromMe } from '../lib/coach/viewer';
  import { syncView } from '../lib/character-header/sync';
  import { characterKey, nameFromSlug, parseCharacterPath, type CharacterPath } from '../lib/characters';
  import {
    encounterSlug,
    fetchCharacter,
    fetchCharacterRating,
    RankingsError,
    type CharacterFight,
    type CharacterPage,
  } from '../lib/rankings/api';
  import { formatDayMonthYear } from '../lib/dates';
  import {
    bestRepeatsHistory,
    bisHref,
    characterBuildsCopy,
    fightRowId,
    profileHeader,
    resolvedProfileSpec,
    seenBuildSpec,
    seenSplits,
  } from '../lib/character-profile';
  import { onCurrentCharacterChange, onCurrentCharacterSwitch, readCurrent } from '../lib/current-character';
  import { factionOfRace } from '../lib/race-faction';
  import { sourcePill } from '../lib/sim/sources';
  import { foreverOf, hasScore, ownCharacter } from '../lib/score/score';
  import { bandForLevel } from '../lib/bis/hover';
  import { slotsHaveGear } from '../lib/addon-export';
  import { characterPrompt, showsNoDataPage } from '../lib/character-prompt';
  import { GEAR_ANCHOR, gearSpecLine } from '../lib/gear/panel';
  import { classSlugFromName } from '../lib/report/tree-sizes';
  import CharacterGearPanel from './character/CharacterGearPanel.svelte';
  import CharacterHandoffLinks from './CharacterHandoffLinks.svelte';
  import CharacterInGame from './CharacterInGame.svelte';
  import CharacterRatingPanel from './CharacterRatingPanel.svelte';
  import LoadError from './ui/LoadError.svelte';
  import Skeleton from './ui/Skeleton.svelte';

  let { path = null }: { path?: CharacterPath | null } = $props();

  const resolved = $derived(
    path ?? (typeof window === 'undefined' ? null : parseCharacterPath(window.location.pathname)),
  );

  let data = $state<CharacterPage | null>(null);
  /** `unlogged`: the API has no record of this character (a 404). The page still answers --
   *  the name from the address, "no logs yet" where the score would be, and how to get the
   *  character onto the site -- since a character search lands here for anyone. */
  let status = $state<'loading' | 'ready' | 'failed' | 'missing' | 'unlogged'>('loading');
  let error = $state('');
  let attempt = $state(0);

  /**
   * This page has no filter or fight-switching UI -- unlike Rankings.svelte or the report
   * island, `resolved` is derived once from a prop that never changes after mount (the
   * shell passes none) and never re-fires this effect. There is therefore no second,
   * differently-addressed fetch that can land after this one and need dropping. The guard
   * below is kept anyway, at the same cost as Rankings.svelte's `state !== requested`
   * check, so this effect stays correct if that ever stops being true (a future edit that
   * threads a live `path` prop through, say) rather than relying on today's absence of a
   * trigger.
   */
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
        if (thrown instanceof RankingsError && thrown.status === 404) {
          status = 'unlogged';
          return;
        }
        status = 'failed';
        error = thrown instanceof Error ? thrown.message : 'That character did not load.';
      });
  });

  /** "Holy Priest · 31/20/0 · <Olympus XXVII> · Classic Beta PvP", "Level 60 Night Elf" and the
   *  chip, from the answer. */
  const header = $derived(data === null || resolved === null ? null : profileHeader(data, resolved, own));
  /** The spec's best-in-slot page at this band: the Gear section's footer link (L8.6 ruling 2),
   *  and where the score card's "Most to gain: Gear" goes for an API that sends no `gear`. */
  const specBis = $derived(
    data === null
      ? null
      : bisHref(
          data.character.class,
          resolvedProfileSpec(data, own),
          data.character.level,
          data.character.race,
        ),
  );
  /** Under "Builds seen at pull": every distinct split, newest first (fix round 5 item 4). */
  const buildsFootnote = $derived(
    data === null ? null : characterBuildsCopy.footnote(seenSplits(data.builds_seen)),
  );

  /** The name a 404 page shows: the address's own slug, readable. */
  const unloggedName = $derived(resolved === null ? '' : nameFromSlug(decodeURIComponent(resolved.slug)));

  /** Whether the character is the viewer's own (signed in and theirs, or their active
   *  character from the header selector), and the viewer's own /v1/me row for it -- off the
   *  same cached /v1/me every island on the page shares. */
  let own = $state<MeCharacter | null>(null);
  let owned = $state(false);
  /** /v1/me has answered (or failed), so `owned` is the settled answer, not the first guess. */
  let ownerKnown = $state(false);
  /** The viewer's plan includes History (`me.entitlements.history`). */
  let historyEntitled = $state(false);
  /** The viewer holds the log coach (`me.entitlements.log_coach`): the owner's coach history. */
  let coachEntitled = $state(false);
  /** The latest rated fight and its rating, from the rating panel once it loads (L8.5 fix round
   *  2 item 3): "Every ranked fight" marks that row "rated 65". */
  let rated = $state<{ report_id: string; fight_index: number; overall: number } | null>(null);

  /** The header selector's current character is this page's. */
  function isActive(requested: CharacterPath): boolean {
    const current = readCurrent();
    return (
      current?.source === 'armory' &&
      current.ref === characterKey(requested.region, requested.ruleset, requested.slug)
    );
  }
  // A choice in the header re-reads `owned` in place (the account's own rows have not changed).
  $effect(() =>
    onCurrentCharacterChange(() => {
      if (resolved !== null) owned = isActive(resolved) || own !== null;
    }),
  );
  // Your character's page moves to the newly chosen character's (character-selector/follow-page.ts);
  // a character you looked up stays.
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
    ownerKnown = false;
    historyEntitled = false;
    coachEntitled = false;
    if (requested === null) return;
    owned = isActive(requested);
    void fetchMeOnce()
      .then((me) => {
        if (resolved !== requested) return;
        own = ownCharacter(me?.characters, requested);
        historyEntitled = me?.entitlements?.history === true;
        coachEntitled = viewerFromMe(me).premium;
        owned = isActive(requested) || own !== null;
        ownerKnown = true;
      })
      .catch(() => {
        own = null;
        ownerKnown = true;
      });
  });

  /** The clock the owner's sync line reads, ticking each minute while the tab lives (L8.4 fix
   *  round 3 item 14), so "Addon · 4 minutes ago" does not stand still. */
  let now = $state(new Date());
  $effect(() => {
    const tick = setInterval(() => {
      now = new Date();
    }, 60_000);
    return () => clearInterval(tick);
  });

  /** The owner's sync line (fix round 2 item 5), the Character header's own: "Addon · 4
   *  minutes ago", the stale words in ember past the nav spec's limit; none for a visitor. The
   *  profile reads no gear, so the line stands whenever a build exists (fix round 3 item 3). */
  const sync = $derived(own === null ? null : syncView(own, now, { needsGear: false }));
  /** The export carried no gear and no talents (W.3): all 17 slots empty and no talent points
   *  spent in the character's newest build. The gear section says "not read yet" instead of
   *  seventeen "Empty" rows. A character with no export at all (`gear: []`) is not this. */
  const gearUnread = $derived(
    data !== null &&
      data.gear !== undefined &&
      data.gear.length > 0 &&
      !slotsHaveGear(data.gear) &&
      (own?.build?.talents === undefined || !own.build.talents.points.some((points) => points > 0)),
  );
  /** The one primary prompt and the one quiet line an under-filled page asks (lib/character-prompt.ts). */
  const prompt = $derived(
    data === null
      ? null
      : characterPrompt({
          hasScore: hasScore(foreverOf(data)),
          hasGear: data.gear !== undefined && data.gear.length > 0,
          owned,
        }),
  );
  /** Nothing on the page for a viewer who does not hold the character: the one-sentence page (P.1). */
  const bareForVisitor = $derived(
    data !== null &&
      showsNoDataPage({
        hasScore: hasScore(foreverOf(data)),
        hasGear: data.gear !== undefined && data.gear.length > 0,
        hasFights: data.best.length > 0 || data.history.length > 0,
        owned: false,
      }),
  );
  /** For a bare page a visitor would see, whether the character has a rating after all (the rating panel's own
   *  endpoint, which the page's answer does not carry): the one-sentence page is only for nothing at all. */
  let ratingProbe = $state<'idle' | 'pending' | 'none' | 'has'>('idle');
  $effect(() => {
    const requested = resolved;
    if (requested === null || !bareForVisitor || !ownerKnown || owned) {
      ratingProbe = 'idle';
      return;
    }
    ratingProbe = 'pending';
    void fetchCharacterRating(requested)
      .then((rating) => {
        if (resolved !== requested) return;
        ratingProbe = rating.sample_size > 0 ? 'has' : 'none';
      })
      .catch(() => {
        if (resolved !== requested) return;
        ratingProbe = 'none';
      });
  });
  /** No score yet and gear to show: the gear card leads, above the empty score card (W.9). */
  const gearFirst = $derived(
    data !== null && !hasScore(foreverOf(data)) && data.gear !== undefined && data.gear.length > 0,
  );
</script>

{#if status === 'missing'}
  <p class="text-[14px]" data-testid="character-missing">
    That is not a character address. They look like <code class="font-mono"
      >/character/us/hardcore/elyra-duskvale</code
    >.
  </p>
{:else if status === 'loading'}
  <Skeleton lines={6} rowHeight="h-4" minHeight={CHARACTER_LOADING_MIN_H} testid="character-skeleton" />
  <HistoryReserve afterSkeleton />
{:else if status === 'failed'}
  <LoadError message={error} onRetry={() => (attempt += 1)} testid="character-error" />
{:else if status === 'unlogged' && resolved !== null}
  <!-- A character the viewer owns is known to them (their own export, no fights yet), so it keeps
       the empty card with its upload action; any other name is simply not on the site. -->
  {#if owned}
    <CharacterUnlogged path={resolved} name={unloggedName} {owned} />
  {:else}
    <CharacterNotFound path={resolved} name={unloggedName} />
  {/if}
{:else if data !== null && resolved !== null && bareForVisitor && (!ownerKnown || (!owned && ratingProbe !== 'none' && ratingProbe !== 'has'))}
  <!-- Whether this is the viewer's own character decides which page they get: wait for /v1/me. -->
  <Skeleton lines={6} rowHeight="h-4" minHeight={CHARACTER_LOADING_MIN_H} testid="character-skeleton" />
  <HistoryReserve afterSkeleton />
{:else if data !== null && resolved !== null && bareForVisitor && !owned && ratingProbe === 'none'}
  <CharacterNotFound path={resolved} name={data.character.name} known />
{:else if data !== null && resolved !== null}
  <div class="reveal flex flex-col gap-[22px] md:gap-8" data-testid="character" id="character">
    <!-- At every width (fix round 3 item 4, L8.4 fix round 2 item 7): the header first (-3),
         then the score card (-2, ForeverScoreCard) with the Gear section under it (-2, after it
         in the page: L8.6): the character is the page's subject. Then the rest in page order. -->
    <header class="-order-3 flex flex-col gap-4 lg:flex-row lg:items-end lg:justify-between">
      <div class="flex flex-col gap-1">
        <CharacterIdentity
          character={{
            key: characterKey(resolved.region, resolved.ruleset, data.character.name),
            region: data.character.region,
            ruleset: data.character.ruleset,
            name: data.character.name,
            class: data.character.class,
            avatar_url: data.character.avatar_url,
            render_url: data.character.render_url,
          }}
          size="lg"
          descriptor="none"
          heading
          align="start"
          supporter={data.supporter === true}
          nameTestid="character-name"
        >
          {#snippet below()}
            {#if header !== null}
              <ProfileDescriptor {header} {sync} />
            {/if}
          {/snippet}
        </CharacterIdentity>
        <!-- "No export yet · paste it here" only for an owner the site holds no build for:
             with one, the sync line above already says when it came. -->
        <CharacterHandoffLinks path={resolved} showNeedsExport={owned && own?.build === undefined} buttons />
      </div>
      {#if data.character.render_url !== undefined}
        <img
          class="hidden max-h-[280px] w-auto object-contain lg:block lg:max-h-[360px]"
          src={data.character.render_url}
          alt=""
          loading="lazy"
          data-testid="character-render"
        />
      {/if}
    </header>

    <CharacterRatingPanel
      path={resolved}
      forever={foreverOf(data)}
      scoreHistory={data.score_history ?? []}
      bisHref={specBis}
      gearHref={data.gear === undefined ? null : `#${GEAR_ANCHOR}`}
      {owned}
      compactEmpty={gearFirst}
      name={data.character.name}
      fights={data.history}
      onrated={(next) => (rated = next)}
    />

    <HistoryMount
      path={resolved}
      {ownerKnown}
      own={own !== null}
      entitled={historyEntitled}
      hasScore={hasScore(foreverOf(data))}
    />

    <!-- L8.6: the Gear section, right under the score card (-order-2, after it in the page), the
         anchor its "Most to gain: Gear" opens. An API that sends no `gear` draws nothing. -->
    {#if data.gear !== undefined}
      <CharacterGearPanel
        gear={data.gear}
        total={data.gear_gap_total ?? null}
        classSlug={data.character.class === undefined || data.character.class === ''
          ? null
          : classSlugFromName(data.character.class)}
        bisHref={specBis}
        spec={resolvedProfileSpec(data, own)}
        specLine={gearSpecLine(data.gear_spec, data.gear_spec_source)}
        band={typeof data.character.level === 'number' ? bandForLevel(data.character.level) : null}
        unread={gearUnread}
        first={gearFirst}
        planner={prompt?.primary === 'planner'}
        capturedAt={own?.build?.captured_at ?? null}
        identity={{
          name: data.character.name,
          classSlug:
            data.character.class === undefined || data.character.class === ''
              ? ''
              : classSlugFromName(data.character.class),
          race: data.character.race ?? '',
          spec: resolvedProfileSpec(data, own) ?? '',
          level: typeof data.character.level === 'number' ? data.character.level : null,
          faction: factionOfRace(data.character.race),
          source: sourcePill({ kind: 'addon', ref: '', captured_at: own?.build?.captured_at ?? '' }),
        }}
        {owned}
      />
    {/if}

    <CharacterInGame
      forever={foreverOf(data)}
      name={data.character.name}
      className={data.character.class}
      guild={data.character.guild?.name ?? null}
      level={data.character.level ?? null}
      race={data.character.race ?? null}
      faction={factionOfRace(data.character.race)}
      {owned}
    />

    <!-- Fix round 1 item 4: a section with no rows is absent (spec principle 2). "Best per
         encounter" is left out when it would only repeat every ranked fight. The two tables
         share one component (character/CharacterFightTable.svelte). -->
    {#if data.best.length > 0 && !bestRepeatsHistory(data.best, data.history)}
      <CharacterFightTable
        title="Best per encounter"
        rows={data.best}
        testid="character-best"
        bossHref={(row: CharacterFight) => `/rankings/${encounterSlug(row.encounter)}`}
        rowKey={(row: CharacterFight) => `${row.encounter_id}-${row.difficulty}-${row.metric}`}
        executionTestid="character-best-execution"
      />
    {/if}

    {#if data.history.length > 0}
      <CharacterFightTable
        title="Every ranked fight"
        rows={data.history}
        testid="character-history"
        bossHref={(row: CharacterFight) => `/reports/${row.report_id}?fight=${row.fight_index}`}
        rowKey={(row: CharacterFight, index: number) => `${row.report_id}-${row.fight_index}-${index}`}
        executionTestid="character-execution"
        rowId={fightRowId}
        {rated}
      />
    {/if}

    {#if data.builds_seen.length > 0}
      <section class="flex flex-col gap-2">
        <h2 class="section-title">Builds seen at pull</h2>
        <ul class="flex flex-col" data-testid="character-builds">
          {#each data.builds_seen as build (`${build.talent_split}-${build.first_seen}`)}
            <li
              class="border-line-soft flex min-h-11 flex-wrap items-center gap-3 border-b px-2 py-2 text-[14px]"
            >
              <span class="numeric font-semibold">{build.talent_split}</span>
              {#if seenBuildSpec(data.character.class, build) !== ''}<span class="text-muted text-[13px]"
                  >{seenBuildSpec(data.character.class, build)}</span
                >{/if}
              <span class="text-muted numeric text-[13px]"
                >first seen {formatDayMonthYear(new Date(build.first_seen))}</span
              >
            </li>
          {/each}
        </ul>
        <!-- L8.4 fix round 4 item 4: in plain words. Fix round 5 item 4: naming every distinct
             split seen, newest first; none to name, no note. -->
        {#if buildsFootnote !== null}
          <p class="text-muted text-[12px]" data-testid="character-builds-footnote">{buildsFootnote}</p>
        {/if}
      </section>
    {/if}

    <!-- The coach's stored reviews of this character (log coach spec 3.5): the owner on Premium
         only, last on the page so its late arrival moves nothing above it. -->
    {#if own !== null && coachEntitled}
      <CoachHistoryList path={resolved} />
    {/if}
  </div>
{/if}
