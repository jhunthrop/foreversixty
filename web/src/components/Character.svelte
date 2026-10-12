<!-- web/src/components/Character.svelte -->
<!-- Every ranked fight this character has, their best per encounter, and the builds they
     have been seen in at pull. The page exists for characters who have never signed in:
     a public report is enough to have one, which is the spec's position. -->
<script lang="ts">
  import CurrentCharacterBar from './CurrentCharacterBar.svelte';
  import CharacterFightTable from './character/CharacterFightTable.svelte';
  import CharacterIdentity from './character/CharacterIdentity.svelte';
  import ProfileDescriptor from './character/ProfileDescriptor.svelte';
  import CharacterUnlogged from './CharacterUnlogged.svelte';
  import { CHARACTER_LOADING_MIN_H } from '../lib/character-layout';
  import { fetchMeOnce, type MeCharacter } from '../lib/account/api';
  import { syncView } from '../lib/character-header/sync';
  import { characterKey, nameFromSlug, parseCharacterPath, type CharacterPath } from '../lib/characters';
  import {
    encounterSlug,
    fetchCharacter,
    RankingsError,
    type CharacterFight,
    type CharacterPage,
  } from '../lib/rankings/api';
  import { formatDayMonthYear } from '../lib/dates';
  import {
    bestRepeatsHistory,
    bisHref,
    characterBuildsCopy,
    newestSplit,
    profileHeader,
    profileSpec,
  } from '../lib/character-profile';
  import { readCurrent } from '../lib/current-character';
  import { factionOfRace } from '../lib/race-faction';
  import { foreverOf, ownCharacter } from '../lib/score/score';
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
  const header = $derived(data === null || resolved === null ? null : profileHeader(data, resolved));

  /** The name a 404 page shows: the address's own slug, readable. */
  const unloggedName = $derived(resolved === null ? '' : nameFromSlug(decodeURIComponent(resolved.slug)));

  /** Whether the character is the viewer's own (signed in and theirs, or their active
   *  character from the header selector), and the viewer's own /v1/me row for it -- off the
   *  same cached /v1/me every island on the page shares. */
  let own = $state<MeCharacter | null>(null);
  let owned = $state(false);

  $effect(() => {
    const requested = resolved;
    own = null;
    owned = false;
    if (requested === null) return;
    const current = readCurrent();
    const active =
      current?.source === 'armory' &&
      current.ref === characterKey(requested.region, requested.ruleset, requested.slug);
    owned = active;
    void fetchMeOnce()
      .then((me) => {
        if (resolved !== requested) return;
        own = ownCharacter(me?.characters, requested);
        owned = active || own !== null;
      })
      .catch(() => {
        own = null;
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
</script>

{#if status === 'missing'}
  <p class="text-[14px]" data-testid="character-missing">
    That is not a character address. They look like <code class="font-mono"
      >/character/us/hardcore/elyra-duskvale</code
    >.
  </p>
{:else if status === 'loading'}
  <Skeleton lines={6} rowHeight="h-4" minHeight={CHARACTER_LOADING_MIN_H} testid="character-skeleton" />
{:else if status === 'failed'}
  <LoadError message={error} onRetry={() => (attempt += 1)} testid="character-error" />
{:else if status === 'unlogged' && resolved !== null}
  <CharacterUnlogged path={resolved} name={unloggedName} {owned} />
{:else if data !== null && resolved !== null}
  <div class="reveal flex flex-col gap-[22px] md:gap-8" data-testid="character" id="character">
    <!-- At every width (fix round 3 item 4, L8.4 fix round 2 item 7): a signed-in chip first
         (-4), then the header (-3) and the score card (-2, ForeverScoreCard), then the visitor's
         sign-in band (-1): the character is the page's subject. Then the rest in page order. -->
    <CurrentCharacterBar spine profile />
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
        <CharacterHandoffLinks path={resolved} showNeedsExport={owned && own?.build === undefined} />
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
      bisHref={bisHref(data.character.class, profileSpec(data), data.character.level, data.character.race)}
      {owned}
      name={data.character.name}
    />

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
      />
    {/if}

    {#if data.builds_seen.length > 0}
      <section class="flex flex-col gap-2">
        <h2 class="section-title text-[18px]">Builds seen at pull</h2>
        <ul class="flex flex-col" data-testid="character-builds">
          {#each data.builds_seen as build (`${build.talent_split}-${build.first_seen}`)}
            <li
              class="border-line-soft flex min-h-11 flex-wrap items-center gap-3 border-b px-2 py-2 text-[14px]"
            >
              <span class="tabular font-mono font-semibold">{build.talent_split}</span>
              {#if build.spec}<span class="text-muted text-[13px]">{build.spec}</span>{/if}
              <span class="text-muted tabular font-mono text-[13px]"
                >first seen {formatDayMonthYear(new Date(build.first_seen))}</span
              >
            </li>
          {/each}
        </ul>
        <!-- L8.4 fix round 4 item 4: in plain words, naming the newest split. -->
        <p class="text-muted text-[12px]" data-testid="character-builds-footnote">
          {characterBuildsCopy.footnote(newestSplit(data.builds_seen) ?? data.builds_seen[0].talent_split)}
        </p>
      </section>
    {/if}
  </div>
{/if}
