<!-- web/src/components/Guild.svelte -->
<!-- Guild control-centre spec v2 (design/specs/2026-10-04-guild-page.md): the page becomes
     a seven-tab control center -- Overview, Roster, Raids, Progression, Readiness, Loot,
     Settings -- one URL, hash-anchored (#roster, #raids, ...). Round 1's whole page (header,
     officer tools, flat roster list, raid nights, progression, roster bests) is kept where
     it is still true and replaced where the spec names a new tab for it: the flat roster
     list and its duplicated unverified note are gone (now GuildRosterTable.svelte, full
     depth, §4.B); the raid-nights/progression/roster-bests lists are now the Raids and
     Progression tabs (GuildRaids.svelte, GuildProgression.svelte); the officer tools strip
     and standing line stay exactly where v1 put them -- directly under the header facts
     line, above the tab strip itself (spec §4.0: "the tab strip sits directly under the
     role-based line") -- so they are visible on every tab, not just Overview. -->
<script lang="ts">
  import { fetchMeOnce } from '../lib/account/api';
  import { lookupAddonExport } from '../lib/addon-export';
  import {
    characterSlug,
    guildClaimHref,
    guildSettingsHref,
    parseGuildPath,
    rulesetLabel,
    type CharacterPath,
    type Region,
    type Ruleset,
  } from '../lib/characters';
  import {
    approveAllRoster,
    approveCharacter,
    awardLoot,
    contestClaim,
    fetchGuildHome,
    fetchGuildLoot,
    fetchGuildProgression,
    fetchGuildReadiness,
    fetchGuildRaids,
    removeCharacter,
    type GuildHome,
    type GuildLootItem,
    type GuildLootPage,
    type GuildProgressionPage,
    type GuildRaidsPage,
    type GuildReadinessPage,
    type GuildRosterRow,
    type GuildViewerRole,
  } from '../lib/guild/api';
  import { guildHomeCopy } from '../lib/guild/copy';
  import { rosterPlannerHref } from '../lib/guild/roster-links';
  import { needsBeforeThursdaySentence, standingSentence } from '../lib/guild/standing';
  import { createLazyComponent, type LazyLoadState } from '../lib/report/lazy-component.svelte';
  import { classColorVar } from '../lib/report/format';
  import { fetchGuild, type GuildPage } from '../lib/rankings/api';
  import GuildOfficerStrip from './guild/GuildOfficerStrip.svelte';
  import SignInPrompt from './SignInPrompt.svelte';
  import GuildOverview from './guild/GuildOverview.svelte';
  import GuildTabs, { tabFromHash, tabsForRole, type GuildTabId } from './guild/GuildTabs.svelte';
  import GuildStatus from './GuildStatus.svelte';
  import EmptyState from './ui/EmptyState.svelte';
  import Skeleton from './ui/Skeleton.svelte';
  import { GUILD_LOADING } from '../lib/guild/layout';

  // Spec §8: "never all seven tabs' data fetched on first paint" extends to their own JS --
  // Overview (and the tab strip itself) load eagerly with GuildShell's own entry chunk;
  // each of the other six tabs is its own dynamically imported chunk, fetched the first
  // time its tab is actually opened (the same `createLazyComponent` idiom the report
  // island's non-landing modes already use). Found necessary during this round's own
  // build: with all seven tabs statically imported, the GuildShell island's entry chunk
  // exceeded `scripts/check-island-size.mjs`'s 16KB gzip budget.
  const rosterLazy = createLazyComponent(() => import('./guild/GuildRosterTable.svelte'));
  const raidsLazy = createLazyComponent(() => import('./guild/GuildRaids.svelte'));
  const progressionLazy = createLazyComponent(() => import('./guild/GuildProgression.svelte'));
  const readinessLazy = createLazyComponent(() => import('./guild/GuildReadiness.svelte'));
  const lootLazy = createLazyComponent(() => import('./guild/GuildLoot.svelte'));
  const settingsLazy = createLazyComponent(() => import('./guild/GuildSettingsTab.svelte'));

  let { path = null }: { path?: CharacterPath | null } = $props();

  const resolved = $derived(
    path ?? (typeof window === 'undefined' ? null : parseGuildPath(window.location.pathname)),
  );

  let data = $state<GuildPage | null>(null);
  let status = $state<'loading' | 'ready' | 'failed' | 'missing'>('loading');
  let error = $state('');
  let attempt = $state(0);

  $effect(() => {
    void attempt;
    const requested = resolved;
    if (requested === null) {
      status = 'missing';
      return;
    }
    status = 'loading';
    void fetchGuild(requested)
      .then((result) => {
        if (resolved !== requested) return;
        data = result;
        status = 'ready';
      })
      .catch((thrown: unknown) => {
        if (resolved !== requested) return;
        status = 'failed';
        error = thrown instanceof Error ? thrown.message : 'That guild did not load.';
      });
  });

  let home = $state<GuildHome | null>(null);
  let homeStatus = $state<'idle' | 'loading' | 'ready'>('idle');
  /** True once /v1/me has answered (signed in or not), so the public sign-in prompt never
   *  flashes for a member during the first render. */
  let meResolved = $state(false);
  let rosterBusy = $state<string | null>(null);
  let rosterActionError = $state('');
  let contestBusy = $state(false);
  let contestError = $state('');
  let showContestConfirm = $state(false);

  let myGuildMembership = $state<{ rank: string; verified: boolean } | null>(null);
  let myCharacterKeys = $state<Set<string>>(new Set());
  let isModerator = $state(false);
  /** The signed-in account's own battletag (fix round 1): the only signal available to
   *  tell "this viewer's own account is the one holding the claim" from "a verified
   *  officer of a different account" -- `GuildHome.claim.claimed_by_name` is a battletag
   *  string, never an account id, so that is what this compares against. */
  let myBattletag = $state<string | null>(null);

  /** Same "fail toward the public-only view" rule as before this rebuild: a stranger, a
   *  signed-out visitor, or a /v1/me failure all land on 'idle', which renders the public
   *  tabs only -- never a visible error for this section. */
  $effect(() => {
    const requested = resolved;
    if (requested === null) return;
    homeStatus = 'loading';
    void fetchMeOnce()
      .then((result) => {
        if (resolved !== requested) return;
        const membership = result?.guilds.find(
          (g) =>
            g.region === requested.region &&
            g.ruleset === requested.ruleset &&
            characterSlug(g.name) === requested.slug,
        );
        myCharacterKeys = new Set((result?.characters ?? []).map((c) => c.key));
        isModerator = result?.user.role === 'moderator' || result?.user.role === 'admin';
        myBattletag = result?.user.battletag ?? null;
        meResolved = true;
        if (membership === undefined) {
          homeStatus = 'idle';
          return null;
        }
        myGuildMembership = { rank: membership.rank ?? 'member', verified: membership.verified };
        return fetchGuildHome(membership.id).then((page) => {
          if (resolved !== requested) return;
          home = page;
          homeStatus = 'ready';
        });
      })
      .catch(() => {
        if (resolved !== requested) return;
        meResolved = true;
        homeStatus = 'idle';
      });
  });

  const canManage = $derived(
    myGuildMembership !== null &&
      myGuildMembership.verified &&
      (myGuildMembership.rank === 'officer' || myGuildMembership.rank === 'leader'),
  );
  const isOfficerOrLeader = $derived(
    myGuildMembership !== null &&
      (myGuildMembership.rank === 'officer' || myGuildMembership.rank === 'leader'),
  );

  /** Control-centre contract addition: `home.viewer.role` is the server's own answer,
   *  preferred outright when present; a home response from before this contract landed
   *  (or a fixture using the old shape) falls back to the pre-existing `/v1/me`-membership
   *  derivation, so this page behaves identically either way. */
  const role = $derived<GuildViewerRole>(
    home?.viewer?.role ?? (home === null ? 'public' : canManage ? 'officer' : 'member'),
  );
  const guildId = $derived(home?.guild.id ?? data?.guild.id ?? null);
  const myCharacterKey = $derived(
    home?.viewer?.character_key ??
      (home !== null
        ? (home.roster.find((r) => myCharacterKeys.has(r.character_key))?.character_key ?? null)
        : null),
  );
  /** The standing line's own name prefix (spec board precedent: "{name} · {rank clause}").
   *  `GuildHomeStanding` (the contract) carries no name of its own -- this is the viewer's
   *  own roster row, looked up by `myCharacterKey`, never invented. */
  const myRosterRow = $derived(
    home !== null && myCharacterKey !== null
      ? (home.roster.find((r) => r.character_key === myCharacterKey) ?? null)
      : null,
  );

  /** Fix round 1: "Contest this claim" must never reach the account that already holds
   *  the claim -- a verified officer whose own battletag matches `claim.claimed_by_name`
   *  is that account, regardless of which of their own characters they are viewing from.
   *  `undefined`/`null` on either side (a pending claim carries no `claimed_by_name`, a
   *  signed-out or /v1/me-failed visitor carries no battletag) reads as "not the
   *  claimant," never a false positive that hides the control from someone who should see
   *  it. */
  const isClaimantAccount = $derived(
    myBattletag !== null &&
      home?.claim.claimed_by_name !== undefined &&
      myBattletag === home.claim.claimed_by_name,
  );

  /** Fix round 1, item 3: the standing line's "Gear in the planner" link, built the exact
   *  same way `GuildRosterHandoff.svelte` builds it per roster row -- one lookup, for the
   *  viewer's own character only, never a second mechanism (tenet 10). `requestedKey`
   *  guards against a stale resolution landing after `myRosterRow` has already moved on to
   *  a different character (the same guard `GuildRosterHandoff` itself uses via its own
   *  `result.path !== requested` check, adapted here since this lookup is keyed by
   *  character rather than by path identity). */
  let viewerPlannerHref = $state<string | null>(null);
  $effect(() => {
    const row = myRosterRow;
    if (row === null || row.consent === 'roster') {
      viewerPlannerHref = null;
      return;
    }
    const requestedKey = row.character_key;
    void lookupAddonExport(rowPath(row)).then((result) => {
      if (myRosterRow?.character_key !== requestedKey) return;
      viewerPlannerHref = result.code !== null ? rosterPlannerHref(result.code) : null;
    });
  });

  async function onApprove(row: GuildRosterRow): Promise<void> {
    if (home === null) return;
    rosterBusy = row.character_key;
    rosterActionError = '';
    try {
      await approveCharacter(home.guild.id, row.region, row.ruleset, characterSlug(row.name));
      home = {
        ...home,
        roster: home.roster.map((r) =>
          r.character_key === row.character_key ? { ...r, verified: true } : r,
        ),
        pending: (home.pending ?? []).filter((r) => r.character_key !== row.character_key),
      };
    } catch (thrown) {
      rosterActionError = thrown instanceof Error ? thrown.message : 'That did not work; try again';
    } finally {
      rosterBusy = null;
    }
  }

  async function onApproveAll(): Promise<void> {
    if (home === null) return;
    rosterBusy = 'approve-all';
    rosterActionError = '';
    try {
      const result = await approveAllRoster(home.guild.id);
      const approved = new Set(result.approved);
      home = {
        ...home,
        roster: home.roster.map((r) => (approved.has(r.character_key) ? { ...r, verified: true } : r)),
        pending: (home.pending ?? []).filter((r) => !approved.has(r.character_key)),
      };
    } catch (thrown) {
      rosterActionError = thrown instanceof Error ? thrown.message : 'That did not work; try again';
    } finally {
      rosterBusy = null;
    }
  }

  async function onRemove(row: GuildRosterRow): Promise<void> {
    if (home === null) return;
    rosterBusy = row.character_key;
    rosterActionError = '';
    try {
      await removeCharacter(home.guild.id, row.region, row.ruleset, characterSlug(row.name));
      home = { ...home, roster: home.roster.filter((r) => r.character_key !== row.character_key) };
    } catch (thrown) {
      rosterActionError = thrown instanceof Error ? thrown.message : 'That did not work; try again';
    } finally {
      rosterBusy = null;
    }
  }

  async function onContest(): Promise<void> {
    if (home === null) return;
    contestBusy = true;
    contestError = '';
    try {
      await contestClaim(home.guild.id);
      showContestConfirm = false;
    } catch (thrown) {
      contestError = thrown instanceof Error ? thrown.message : 'That did not work; try again';
      contestBusy = false;
      return;
    }
    try {
      home = await fetchGuildHome(home.guild.id);
    } catch {
      contestError = guildHomeCopy.contestRecordedRefreshFailed;
    } finally {
      contestBusy = false;
    }
  }

  function rowPath(row: GuildRosterRow): CharacterPath {
    return { region: row.region as Region, ruleset: row.ruleset as Ruleset, slug: characterSlug(row.name) };
  }

  /** Unchanged v1 rule: a contested claim freezes Remove for everyone except the row's own
   *  account or a site moderator -- Approve carries no such override. */
  function removeOverridesFrozen(row: GuildRosterRow): boolean {
    return myCharacterKeys.has(row.character_key) || isModerator;
  }

  // `home.summary` (control-centre contract) is the authoritative, guild-wide count --
  // preferred whenever it has loaded; the public `data.progression` fallback (today's own
  // public roster_best/progression read, unaffected by this round) covers a visitor who
  // never reaches a membership match, or a pre-contract home response.
  const killed = $derived(
    home?.summary?.named_encounters_down ?? (data?.progression ?? []).filter((row) => row.kills > 0).length,
  );
  const pulls = $derived(
    home?.summary?.pulls_this_tier ??
      (data?.progression ?? []).reduce((total, row) => total + row.pull_count, 0),
  );

  // ---------------------------------------------------------------------------------------
  // Tab routing (spec §4.0): one URL, hash-anchored, never a separate route. A public
  // visitor who opens a disallowed hash (a stale member/officer link) silently falls back
  // to the first public tab, never a 403 (tabFromHash's own contract).
  // ---------------------------------------------------------------------------------------
  function currentHashTab(): GuildTabId {
    const hash = typeof window === 'undefined' ? '' : window.location.hash;
    return tabFromHash(hash, tabsForRole(role));
  }

  /** A writable `$derived`: recomputes from `role` whenever the viewer's role resolves
   *  (so a tab disallowed for the role the page first guessed public falls back once
   *  home loads), and `selectTab`/the hashchange listener below both write directly to
   *  it, which locally overrides the derived value until `role` changes again. */
  let activeTab = $derived<GuildTabId>(currentHashTab());

  $effect(() => {
    if (typeof window === 'undefined') return;
    const onHashChange = (): void => {
      activeTab = currentHashTab();
    };
    window.addEventListener('hashchange', onHashChange);
    return () => window.removeEventListener('hashchange', onHashChange);
  });

  function selectTab(tab: GuildTabId): void {
    activeTab = tab;
    if (typeof window !== 'undefined') window.history.replaceState(null, '', `#${tab}`);
  }

  /** Starts the chunk fetch for whichever tab is now active -- `load()` is itself a no-op
   *  once a chunk has resolved or is already in flight, so this is safe to re-run on every
   *  `activeTab` change. */
  $effect(() => {
    if (activeTab === 'roster') rosterLazy.load();
    else if (activeTab === 'raids') raidsLazy.load();
    else if (activeTab === 'progression') progressionLazy.load();
    else if (activeTab === 'readiness') readinessLazy.load();
    else if (activeTab === 'loot') lootLazy.load();
    else if (activeTab === 'settings') settingsLazy.load();
  });

  // ---------------------------------------------------------------------------------------
  // Per-tab lazy fetches (spec §8): each of Raids/Progression/Readiness/Loot mounts only
  // once its own tab is opened, never all four on first paint. Any failure -- including a
  // 404 from an endpoint the api lane has not deployed yet -- lands the tab on 'missing',
  // its own honest empty state with a next-step line, never a crash (hard rule).
  // ---------------------------------------------------------------------------------------
  type TabFetchStatus = 'idle' | 'loading' | 'ready' | 'missing';

  let raids = $state<GuildRaidsPage | null>(null);
  let raidsStatus = $state<TabFetchStatus>('idle');
  $effect(() => {
    if (activeTab !== 'raids' || guildId === null || raidsStatus !== 'idle') return;
    raidsStatus = 'loading';
    void fetchGuildRaids(guildId)
      .then((page) => {
        raids = page;
        raidsStatus = 'ready';
      })
      .catch(() => {
        raidsStatus = 'missing';
      });
  });

  let progression = $state<GuildProgressionPage | null>(null);
  let progressionStatus = $state<TabFetchStatus>('idle');
  $effect(() => {
    if (activeTab !== 'progression' || guildId === null || progressionStatus !== 'idle') return;
    progressionStatus = 'loading';
    void fetchGuildProgression(guildId)
      .then((page) => {
        progression = page;
        progressionStatus = 'ready';
      })
      .catch(() => {
        progressionStatus = 'missing';
      });
  });

  let readiness = $state<GuildReadinessPage | null>(null);
  let readinessStatus = $state<TabFetchStatus>('idle');
  $effect(() => {
    if (activeTab !== 'readiness' || guildId === null || readinessStatus !== 'idle') return;
    readinessStatus = 'loading';
    void fetchGuildReadiness(guildId)
      .then((page) => {
        readiness = page;
        readinessStatus = 'ready';
      })
      .catch(() => {
        readinessStatus = 'missing';
      });
  });

  let loot = $state<GuildLootPage | null>(null);
  let lootStatus = $state<TabFetchStatus>('idle');
  let lootBusyItemId = $state<number | null>(null);
  let lootError = $state('');
  $effect(() => {
    if (activeTab !== 'loot' || guildId === null || lootStatus !== 'idle') return;
    lootStatus = 'loading';
    void fetchGuildLoot(guildId)
      .then((page) => {
        loot = page;
        lootStatus = 'ready';
      })
      .catch(() => {
        lootStatus = 'missing';
      });
  });

  async function onAward(item: GuildLootItem, charKey: string): Promise<void> {
    if (guildId === null || loot === null || loot.selected === null) return;
    lootBusyItemId = item.item_id;
    lootError = '';
    try {
      await awardLoot(guildId, {
        encounter_id: loot.selected,
        item_id: item.item_id,
        character_key: charKey,
      });
      loot = await fetchGuildLoot(guildId, loot.selected);
    } catch (thrown) {
      lootError = thrown instanceof Error ? thrown.message : 'That did not work; try again';
    } finally {
      lootBusyItemId = null;
    }
  }
</script>

{#if status === 'missing'}
  <h1 class="sr-only">Guild</h1>
  <p class="text-[14px]" data-testid="guild-missing">
    That is not a guild address. They look like <code class="font-mono">/guild/eu/normal/the-last-watch</code
    >.
  </p>
{:else if status === 'loading' || status === 'failed'}
  <h1 class="sr-only">Guild</h1>
  <GuildStatus
    status={status === 'loading' ? 'loading' : 'failed'}
    error={status === 'failed' ? error : ''}
    onRetry={() => (attempt += 1)}
    lines={GUILD_LOADING.home.lines}
    minHeight={GUILD_LOADING.home.minHeight}
    testid="guild"
  />
{:else if data !== null && resolved !== null}
  <div class="reveal flex flex-col gap-[22px] md:gap-8" data-testid="guild" id="guild">
    <!-- Fix round 1 (owner rule): a full-bleed band, inner content capped and centred at
         1344px, like every other page. Built with the same breakout `LogsHeroBand.svelte`
         and `PlannerHeaderBand.svelte` already use (`w-screen` + `ml-[calc(50%-50vw)]`),
         not a second way of doing the same thing (tenet 10) -- this page's own hero is
         only ever known client-side (guild name, standing, tab state), so there is no
         build-time-known markup to hoist into a separate Astro `beforeMain` region the way
         `ArtPanel`/`ClassHeader` do on the guides page; the CSS breakout keeps the one
         GuildShell island's state in one place instead of splitting it across two mount
         points. The band's own inner uses the identical max-width/gutter column
         `[...path].astro`'s `<main>` already gives the tab content below, so both share one
         left edge at every width -- verified at 1440 and 2000. -->
    <div class="bg-raised w-screen" style="margin-left:calc(50% - 50vw)">
      <div
        class="mx-auto flex w-full max-w-[1344px] flex-col gap-3 px-[18px] pt-4 pb-[22px] md:px-12 md:pt-7 md:pb-8"
      >
        <header class="flex flex-col gap-3">
          <span class="label text-gold flex items-center gap-2.5">
            <i class="bg-gold inline-block h-px w-7" aria-hidden="true"></i>
            Guild
          </span>
          <h1 class="font-display text-strong text-[22px] font-bold">{data.guild.name}</h1>
          {#if home?.summary?.updated_at !== undefined}
            <span class="text-muted text-[12px]">Updated {home.summary.updated_at.slice(0, 10)}</span>
          {/if}
          <p class="text-muted text-[13px]">
            {rulesetLabel(resolved.ruleset)}
            {resolved.region.toUpperCase()} ·
            <span class="tabular font-mono">{killed}</span> bosses down ·
            <span class="tabular font-mono">{pulls}</span> pulls
          </p>

          <!-- Spec §4.A.1 fourth branch (public visitor): one sign-in prompt in the hero,
               the Battle.net button style, returning to this page. Live defect 2026-10-05:
               the public view shipped without it. -->
          {#if meResolved && role === 'public' && myBattletag === null}
            <SignInPrompt
              line="See where you stand in this guild."
              next={typeof location === 'undefined' ? '/account?signed_in=1' : location.pathname}
              testid="guild-sign-in-prompt"
            />
          {/if}

          <!-- The role-based line (spec §4.0/§4.A.1): unchanged v1 mechanism for claim/
               settings, still above the tab strip. "Contest this claim" moved out of here
               entirely (fix round 1, item 2) -- it now lives in the Settings tab, under
               the claim block, and never reaches the claimant's own account. -->
          {#if homeStatus === 'ready' && home !== null}
            <div class="flex flex-col gap-2" data-testid="guild-role-line">
              <!-- Spec §4.A.1: the standing line's four branches (ranked / alone-in-spec /
                   unverified / signed-out) never branch on officer vs member -- every
                   verified viewer sees it, officer included; only the member-only "before
                   Thursday" sentence (§4.A.2) is role-gated. Kept unconditional here,
                   separate from the claim/settings row below it (unchanged v1 mechanism),
                   rather than the two fighting over the same slot. -->
              {#if home.standing !== null && home.standing !== undefined}
                <div class="flex flex-col gap-1" data-testid="guild-overview-standing">
                  <p class="flex flex-wrap items-center gap-3 text-[14px] font-semibold">
                    <span style={`color:${classColorVar(home.standing.class)}`}>
                      {myRosterRow !== null ? `${myRosterRow.name} · ` : ''}{standingSentence(home.standing)}
                    </span>
                    {#if viewerPlannerHref !== null}
                      <a
                        href={viewerPlannerHref}
                        class="text-[13px] font-bold"
                        data-testid="guild-standing-planner-link"
                      >
                        Gear in the planner
                      </a>
                    {/if}
                  </p>
                  {#if role === 'member'}
                    <p
                      class="text-[13px]"
                      style={home.standing.needs_before_next_raid.length === 0
                        ? 'color:#7bff5c'
                        : 'color:#ff6b5c'}
                      data-testid="guild-overview-needs-thursday"
                    >
                      {needsBeforeThursdaySentence(home.standing.needs_before_next_raid)}
                    </p>
                  {/if}
                </div>
              {/if}
              <div class="flex flex-wrap items-center justify-between gap-3">
                {#if isOfficerOrLeader && home.claim.state === 'unclaimed'}
                  <a
                    class="text-[13px] font-semibold"
                    href={guildClaimHref(resolved.region, resolved.ruleset, home.guild.name)}
                    data-testid="guild-claim-link"
                  >
                    {guildHomeCopy.claimLink}
                  </a>
                {:else if isOfficerOrLeader && home.claim.state === 'pending'}
                  <a
                    class="text-[13px] font-semibold"
                    href={guildClaimHref(resolved.region, resolved.ruleset, home.guild.name)}
                    data-testid="guild-confirm-claim-link"
                  >
                    {guildHomeCopy.confirmClaimLink}
                  </a>
                {:else if canManage}
                  <a
                    class="text-[13px] font-semibold"
                    href={guildSettingsHref(resolved.region, resolved.ruleset, home.guild.name)}
                    data-testid="guild-settings-link"
                  >
                    {guildHomeCopy.settingsLink}
                  </a>
                {:else if myGuildMembership !== null && !myGuildMembership.verified}
                  <p class="text-muted text-[13px]" data-testid="guild-home-not-verified-note">
                    {guildHomeCopy.notVerifiedNote}
                  </p>
                {/if}
              </div>

              <!-- The claimant's own officers (fix round 1, item 2): the glance strip
                   (claim badge, waiting-for-approval, invite rotate), now part of the band
                   so it is visible on every tab, not Overview-only. -->
              {#if canManage && home.claim.state === 'claimed' && guildId !== null}
                <GuildOfficerStrip
                  {guildId}
                  claimedByName={home.claim.claimed_by_name ?? null}
                  waitingCount={home.summary?.waiting_for_approval ?? 0}
                  onSelectRoster={() => selectTab('roster')}
                />
              {/if}

              {#if home.claim.frozen}
                <p class="text-[13px]" role="alert" data-testid="guild-home-frozen">
                  {guildHomeCopy.frozenNotice}
                </p>
              {/if}
            </div>
          {/if}
        </header>

        <GuildTabs active={activeTab} {role} onSelect={selectTab} />
      </div>
    </div>

    {#if activeTab === 'overview'}
      <GuildOverview
        {guildId}
        {role}
        reports={home?.reports ?? data.reports}
        summary={home?.summary}
        claimed={home?.claim.state === 'claimed'}
        claimedByName={home?.claim.claimed_by_name ?? null}
        onSelectTab={selectTab}
      />
    {:else if activeTab === 'roster' && home !== null}
      {#if rosterLazy.current}
        <rosterLazy.current
          roster={home.roster}
          pending={home.pending ?? []}
          officer={role === 'officer' || role === 'moderator'}
          {myCharacterKey}
          {rosterBusy}
          {rosterActionError}
          frozen={home.claim.frozen}
          {removeOverridesFrozen}
          onApprove={(row) => void onApprove(row)}
          onRemove={(row) => void onRemove(row)}
          onApproveAll={() => void onApproveAll()}
          {rowPath}
        />
      {:else}
        <Skeleton lines={6} minHeight="min-h-[320px]" testid="guild-roster-skeleton" />
        {@render lazyFallback(rosterLazy)}
      {/if}
    {:else if activeTab === 'raids'}
      {#if raidsStatus === 'loading' || (raidsStatus === 'ready' && !raidsLazy.current)}
        <Skeleton lines={5} minHeight="min-h-[240px]" testid="guild-raids-skeleton" />
        {@render lazyFallback(raidsLazy)}
      {:else if raidsStatus === 'missing'}
        <EmptyState
          message="Raid night detail isn't live yet -- the api lane is still deploying this endpoint. Check back after the next sync."
          testid="guild-raids-missing"
        />
      {:else if raidsLazy.current}
        <raidsLazy.current rows={raids?.rows ?? []} officer={role === 'officer' || role === 'moderator'} />
      {/if}
    {:else if activeTab === 'progression'}
      {#if progressionStatus === 'loading' || (progressionStatus === 'ready' && !progressionLazy.current)}
        <Skeleton lines={5} minHeight="min-h-[240px]" testid="guild-progression-skeleton" />
        {@render lazyFallback(progressionLazy)}
      {:else if progressionStatus === 'missing' && (data?.roster_best ?? []).length === 0}
        <EmptyState
          message="Progression detail isn't live yet -- the api lane is still deploying this endpoint. Check back after the next sync."
          testid="guild-progression-missing"
        />
      {:else if progressionLazy.current}
        <progressionLazy.current
          progression={progressionStatus === 'ready' ? progression : null}
          rosterBest={data?.roster_best ?? []}
          region={resolved.region}
          ruleset={resolved.ruleset}
        />
      {/if}
    {:else if activeTab === 'readiness' && home !== null}
      {#if readinessStatus === 'loading' || (readinessStatus === 'ready' && !readinessLazy.current)}
        <Skeleton lines={6} minHeight="min-h-[280px]" testid="guild-readiness-skeleton" />
        {@render lazyFallback(readinessLazy)}
      {:else if readinessStatus === 'missing'}
        <EmptyState
          message="Readiness checks aren't live yet -- the api lane is still deploying this endpoint. Check back after the next sync."
          testid="guild-readiness-missing"
        />
      {:else if readinessLazy.current}
        <readinessLazy.current
          rows={readiness?.rows ?? []}
          officer={role === 'officer' || role === 'moderator'}
          {myCharacterKey}
        />
      {/if}
    {:else if activeTab === 'loot'}
      {#if lootStatus === 'loading' || (lootStatus === 'ready' && !lootLazy.current)}
        <Skeleton lines={5} minHeight="min-h-[240px]" testid="guild-loot-skeleton" />
        {@render lazyFallback(lootLazy)}
      {:else if lootStatus === 'missing'}
        <EmptyState
          message="Loot ranking isn't live yet -- the api lane is still deploying this endpoint. Check back after the next sync."
          testid="guild-loot-missing"
        />
      {:else if lootLazy.current}
        <lootLazy.current
          loot={lootStatus === 'ready' ? loot : null}
          officer={role === 'officer' || role === 'moderator'}
          memberReadOnly={role === 'member'}
          busyItemId={lootBusyItemId}
          onAward={(item, charKey) => void onAward(item, charKey)}
        />
        {#if lootError !== ''}<p class="text-[13px]" role="alert" data-testid="guild-loot-error">
            {lootError}
          </p>{/if}
      {/if}
    {:else if activeTab === 'settings' && (role === 'officer' || role === 'moderator') && home !== null}
      {#if settingsLazy.current}
        <settingsLazy.current
          path={resolved}
          {home}
          {isClaimantAccount}
          {contestBusy}
          {contestError}
          {showContestConfirm}
          onContestStart={() => {
            showContestConfirm = true;
            contestError = '';
          }}
          onContestCancel={() => {
            showContestConfirm = false;
            contestError = '';
          }}
          onContestConfirm={() => void onContest()}
        />
      {:else}
        <Skeleton lines={6} minHeight="min-h-[400px]" testid="guild-settings-skeleton" />
        {@render lazyFallback(settingsLazy)}
      {/if}
    {/if}
  </div>
{/if}

{#snippet lazyFallback(lazy: LazyLoadState)}
  {#if lazy.error !== ''}
    <p class="text-muted text-[13px]" role="alert">
      {lazy.error}
      <button
        type="button"
        class="text-strong ml-1 inline-flex items-center underline"
        onclick={() => lazy.load()}
      >
        Try again
      </button>
    </p>
  {/if}
{/snippet}
