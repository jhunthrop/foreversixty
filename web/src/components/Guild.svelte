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
  import { needsBeforeThursdaySentence, standingSentence } from '../lib/guild/standing';
  import { SECONDARY_BUTTON_FIXED } from '../lib/planner/styles';
  import { fetchGuild, type GuildPage } from '../lib/rankings/api';
  import GuildOverview from './guild/GuildOverview.svelte';
  import GuildProgression from './guild/GuildProgression.svelte';
  import GuildRaids from './guild/GuildRaids.svelte';
  import GuildReadiness from './guild/GuildReadiness.svelte';
  import GuildLoot from './guild/GuildLoot.svelte';
  import GuildRosterTable from './guild/GuildRosterTable.svelte';
  import GuildSettingsTab from './guild/GuildSettingsTab.svelte';
  import GuildTabs, { tabFromHash, tabsForRole, type GuildTabId } from './guild/GuildTabs.svelte';
  import GuildStatus from './GuildStatus.svelte';
  import EmptyState from './ui/EmptyState.svelte';
  import Skeleton from './ui/Skeleton.svelte';
  import { GUILD_LOADING } from '../lib/guild/layout';

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
  let rosterBusy = $state<string | null>(null);
  let rosterActionError = $state('');
  let contestBusy = $state(false);
  let contestError = $state('');
  let showContestConfirm = $state(false);

  let myGuildMembership = $state<{ rank: string; verified: boolean } | null>(null);
  let myCharacterKeys = $state<Set<string>>(new Set());
  let isModerator = $state(false);

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
  const canContest = $derived(myGuildMembership !== null);

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

  const killed = $derived((data?.progression ?? []).filter((row) => row.kills > 0).length);
  const pulls = $derived((data?.progression ?? []).reduce((total, row) => total + row.pull_count, 0));

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
    <header class="flex flex-col gap-2">
      <h1 class="section-title text-[18px]">{data.guild.name}</h1>
      <p class="text-muted text-[13px]">
        {rulesetLabel(resolved.ruleset)}
        {resolved.region.toUpperCase()} ·
        <span class="tabular font-mono">{killed}</span> bosses down ·
        <span class="tabular font-mono">{pulls}</span> pulls
      </p>

      <!-- The role-based line (spec §4.0/§4.A.1): unchanged v1 mechanism for claim/
           contest/settings, now living here -- above the tab strip -- so it stays visible
           on every tab, not only Overview. -->
      {#if homeStatus === 'ready' && home !== null}
        <div class="flex flex-col gap-2" data-testid="guild-role-line">
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
            {:else if role === 'member' && home.standing !== null && home.standing !== undefined}
              <div class="flex flex-col gap-1" data-testid="guild-overview-standing">
                <p class="text-[14px]"><strong>{standingSentence(home.standing)}</strong></p>
                <p
                  class="text-[13px]"
                  style={home.standing.needs_before_next_raid.length === 0
                    ? 'color:#7bff5c'
                    : 'color:#ff6b5c'}
                  data-testid="guild-overview-needs-thursday"
                >
                  {needsBeforeThursdaySentence(home.standing.needs_before_next_raid)}
                </p>
              </div>
            {/if}

            {#if canContest && home.claim.state !== 'unclaimed' && home.claim.state !== 'contested'}
              <button
                class="{SECONDARY_BUTTON_FIXED} border-line-warm text-text px-3"
                onclick={() => {
                  showContestConfirm = true;
                  contestError = '';
                }}
                disabled={contestBusy}
                data-testid="guild-home-contest-button"
              >
                {guildHomeCopy.contestButton}
              </button>
            {/if}
          </div>

          {#if home.claim.frozen}
            <p class="text-[13px]" role="alert" data-testid="guild-home-frozen">
              {guildHomeCopy.frozenNotice}
            </p>
          {/if}

          {#if !showContestConfirm && contestError !== ''}
            <p class="text-[13px]" role="alert" data-testid="guild-home-contest-error">{contestError}</p>
          {/if}

          {#if showContestConfirm}
            <div
              class="border-line-soft flex flex-col gap-3 border p-4"
              data-testid="guild-home-contest-confirm"
            >
              <ul class="flex flex-col gap-1 text-[13px]">
                {#each guildHomeCopy.contestRules as rule (rule)}
                  <li>{rule}</li>
                {/each}
              </ul>
              <div class="flex gap-3">
                <button
                  class="{SECONDARY_BUTTON_FIXED} border-line-warm-strong text-strong px-3"
                  onclick={() => void onContest()}
                  disabled={contestBusy}
                  data-testid="guild-home-contest-confirm-button"
                >
                  {guildHomeCopy.contestConfirmButton}
                </button>
                <button
                  class="{SECONDARY_BUTTON_FIXED} border-line-warm text-text px-3"
                  onclick={() => {
                    showContestConfirm = false;
                    contestError = '';
                  }}
                  disabled={contestBusy}
                >
                  {guildHomeCopy.cancel}
                </button>
              </div>
              {#if contestError !== ''}<p class="text-[13px]" role="alert">{contestError}</p>{/if}
            </div>
          {/if}
        </div>
      {/if}
    </header>

    <GuildTabs active={activeTab} {role} onSelect={selectTab} />

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
      <GuildRosterTable
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
    {:else if activeTab === 'raids'}
      {#if raidsStatus === 'loading'}
        <Skeleton lines={5} minHeight="min-h-[240px]" testid="guild-raids-skeleton" />
      {:else if raidsStatus === 'missing'}
        <EmptyState
          message="Raid night detail isn't live yet -- the api lane is still deploying this endpoint. Check back after the next sync."
          testid="guild-raids-missing"
        />
      {:else}
        <GuildRaids rows={raids?.rows ?? []} officer={role === 'officer' || role === 'moderator'} />
      {/if}
    {:else if activeTab === 'progression'}
      {#if progressionStatus === 'loading'}
        <Skeleton lines={5} minHeight="min-h-[240px]" testid="guild-progression-skeleton" />
      {:else if progressionStatus === 'missing' && (data?.roster_best ?? []).length === 0}
        <EmptyState
          message="Progression detail isn't live yet -- the api lane is still deploying this endpoint. Check back after the next sync."
          testid="guild-progression-missing"
        />
      {:else}
        <GuildProgression
          progression={progressionStatus === 'ready' ? progression : null}
          rosterBest={data?.roster_best ?? []}
          region={resolved.region}
          ruleset={resolved.ruleset}
        />
      {/if}
    {:else if activeTab === 'readiness' && home !== null}
      {#if readinessStatus === 'loading'}
        <Skeleton lines={6} minHeight="min-h-[280px]" testid="guild-readiness-skeleton" />
      {:else if readinessStatus === 'missing'}
        <EmptyState
          message="Readiness checks aren't live yet -- the api lane is still deploying this endpoint. Check back after the next sync."
          testid="guild-readiness-missing"
        />
      {:else}
        <GuildReadiness
          rows={readiness?.rows ?? []}
          officer={role === 'officer' || role === 'moderator'}
          {myCharacterKey}
        />
      {/if}
    {:else if activeTab === 'loot'}
      {#if lootStatus === 'loading'}
        <Skeleton lines={5} minHeight="min-h-[240px]" testid="guild-loot-skeleton" />
      {:else if lootStatus === 'missing'}
        <EmptyState
          message="Loot ranking isn't live yet -- the api lane is still deploying this endpoint. Check back after the next sync."
          testid="guild-loot-missing"
        />
      {:else}
        <GuildLoot
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
    {:else if activeTab === 'settings' && (role === 'officer' || role === 'moderator')}
      <GuildSettingsTab path={resolved} />
    {/if}
  </div>
{/if}
