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
  import { fetchMeFresh, type Me } from '../lib/account/api';
  import { lookupAddonExport } from '../lib/addon-export';
  import {
    characterSlug,
    guildClaimHref,
    guildHref,
    guildRaidHref,
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
    unawardLoot,
    contestClaim,
    fetchGuildHome,
    fetchGuildLoot,
    fetchGuildProgression,
    fetchGuildRaids,
    removeCharacter,
    GuildApiError,
    type GuildHome,
    type GuildLootItem,
    type GuildLootPage,
    type GuildProgressionPage,
    type GuildRaidsPage,
    type GuildRosterRow,
    type GuildViewerRole,
  } from '../lib/guild/api';
  import { guildHomeCopy } from '../lib/guild/copy';
  import {
    officerToolsOn,
    raidGuildToolsOn,
    refusalGate,
    type GuildPlanRefusal,
  } from '../lib/guild/officer-tools';
  import { plannedReserveRows, rememberPlannedRows } from '../lib/guild/planned-reserve';
  import { browserStorage } from '../lib/guild/reserve-keys';
  import {
    bandHoldKey,
    forgetBandHold,
    readBandHold,
    rememberBandHold,
    type BandLayout,
  } from '../lib/guild/band-hold';
  import { GUILD_BANDLESS_TOP, PLANNED_RESERVE } from '../lib/guild/raid-layout';
  import { guildMarkSrc } from '../lib/guild/mark';
  import { bandFacts } from '../lib/guild/fact-line';
  import { showBandCounts, updatedLine } from '../lib/guild/overview-view';
  import { claimEntryCopy, claimEntryFor, claimEntryNotes, rosterGateState } from '../lib/guild/roster-gate';
  import { rosterPlannerHref } from '../lib/guild/roster-links';
  import { afterApprove, afterRemove } from '../lib/guild/roster-model';
  import { needsBeforeSentence, standingSentence } from '../lib/guild/standing';
  import { createRefreshingRead } from '../lib/guild/refreshing-read';
  import {
    attendanceSourceOf,
    raidsChangeAction,
    readinessTabView,
    type TabFetchStatus,
  } from '../lib/guild/guild-page-decisions';
  import { loadReadinessBoard, type ReadinessBoard } from '../lib/guild/readiness-raid';
  import { nextRaidDeadline, raidCopy, showsBandLine } from '../lib/guild/raid-copy';
  import { createLazyComponent, type LazyLoadState } from '../lib/report/lazy-component.svelte';
  import { classColorVar } from '../lib/report/format';
  import { fetchGuild, type GuildPage } from '../lib/rankings/api';
  import FactLine from './guild/FactLine.svelte';
  import GuildUnverifiedNote from './guild/GuildUnverifiedNote.svelte';
  import GuildRosterGate from './guild/GuildRosterGate.svelte';
  import FactionCrest from './character/FactionCrest.svelte';
  import SignInPrompt from './SignInPrompt.svelte';
  import GuildOverview from './guild/GuildOverview.svelte';
  import GuildTabs, {
    tabAfterHashChange,
    tabFromHash,
    tabsForRole,
    type GuildTabId,
  } from './guild/GuildTabs.svelte';
  import GuildStatus from './GuildStatus.svelte';
  import EmptyState from './ui/EmptyState.svelte';
  import PageHead from './ui/PageHead.svelte';
  import Skeleton from './ui/Skeleton.svelte';
  import { findMyGuild } from '../lib/guild/my-guild';
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
  /** The raid line in the band (raids spec 3.4): its own chunk, started the moment the viewer's
   *  membership is known, so it lands alongside /home and keeps the entry chunk under budget. */
  const bandLineLazy = createLazyComponent(() => import('./guild/RaidBandLine.svelte'));

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

  /** Header art round (spec §12.2, ruling 12.2.B): `guild.faction` is a stored column, read
   *  from the public `data` fetch only -- the same fetch that already gates the eyebrow/h1
   *  first paint (`status === 'ready'`), so the crest/vignette/watermark render in the same
   *  paint as the name, never as a later pop-in once the member-only `home` fetch resolves
   *  after it. Missing or `null` both read as "no faction" (§12.2's own states table). */
  const faction = $derived(data?.guild.faction ?? null);

  let home = $state<GuildHome | null>(null);
  let homeStatus = $state<'idle' | 'loading' | 'ready'>('idle');
  /** True once /v1/me has answered (signed in or not), so the public sign-in prompt never
   *  flashes for a member during the first render. */
  let meResolved = $state(false);
  /** `/v1/me` answered with a session (not a 401 or a failure). */
  let signedIn = $state(false);
  let rosterBusy = $state<string | null>(null);
  let rosterActionError = $state('');
  let contestBusy = $state(false);
  let contestError = $state('');
  let showContestConfirm = $state(false);

  let myGuildMembership = $state<{ rank: string; verified: boolean; can: string[] } | null>(null);
  let myCharacterKeys = $state<Set<string>>(new Set());
  /** The site's own record of each of the account's characters' spec (`/v1/me`), by character key: the
   *  header chip reads it, and the roster's spec line says where the addon's snapshot disagrees. */
  let mySiteSpecs = $state<Map<string, string>>(new Map());
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
    // Painted from the cache at once, then asked again: a join, a grant or an approval since
    // the last visit replaces the answer through `apply` (`onFresh`) without a reload.
    const apply = (result: Me | null): Promise<void> | null => {
      if (resolved !== requested) return null;
      const membership = findMyGuild(result, requested) ?? undefined;
      myCharacterKeys = new Set((result?.characters ?? []).map((c) => c.key));
      mySiteSpecs = new Map(
        (result?.characters ?? []).flatMap((c) => (c.spec === undefined ? [] : [[c.key, c.spec] as const])),
      );
      isModerator = result?.user.role === 'moderator' || result?.user.role === 'admin';
      myBattletag = result?.user.battletag ?? null;
      meResolved = true;
      signedIn = result !== null;
      if (membership === undefined) {
        myGuildMembership = null;
        home = null;
        homeStatus = 'idle';
        return null;
      }
      myGuildMembership = {
        rank: membership.rank ?? 'member',
        verified: membership.verified,
        can: membership.can ?? [],
      };
      const show = (page: GuildHome): void => {
        if (resolved !== requested) return;
        home = page;
        homeStatus = 'ready';
      };
      if (membership.verified) bandLineLazy.load();
      return fetchGuildHome(membership.id, undefined, undefined, show).then(show);
    };
    void fetchMeFresh((fresh) => void apply(fresh)?.catch(() => {}))
      .then(apply)
      .catch(() => {
        if (resolved !== requested) return;
        meResolved = true;
        homeStatus = 'idle';
      });
  });

  /** The roster is closed to everyone but verified members (`/home` answers them alone), so
   *  anyone else -- a visitor, another guild's member, a member waiting on an officer -- gets
   *  the `#roster` section's one sentence instead. A verified member whose `/home` merely
   *  failed gets nothing: they are not who that sentence is for. */
  const rosterGate = $derived(
    meResolved && home === null && homeStatus === 'idle' && myGuildMembership?.verified !== true
      ? rosterGateState({ signedIn, member: myGuildMembership !== null })
      : null,
  );

  /** The guild's own `#roster` link is followed before this section exists (the page renders
   *  after its fetches), so the browser's anchor jump has nothing to land on; do it here once
   *  the section is on the page. */
  $effect(() => {
    if (rosterGate === null || typeof window === 'undefined' || window.location.hash !== '#roster') return;
    document.getElementById('roster')?.scrollIntoView({ block: 'start' });
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

  /** What the band says about claiming: the "Run this guild? Claim it." prompt only for an officer
   *  rank in an unclaimed guild, who holds a claimed guild for its members, nothing for a stranger. */
  const claimEntry = $derived(
    claimEntryFor({
      signedIn,
      canManage,
      membership: myGuildMembership,
      claimState: home?.claim.state ?? null,
      claimedBy: home?.claim.claimed_by_name ?? null,
    }),
  );
  /** Control-centre contract addition: `home.viewer.role` is the server's own answer,
   *  preferred outright when present; a home response from before this contract landed
   *  (or a fixture using the old shape) falls back to the pre-existing `/v1/me`-membership
   *  derivation, so this page behaves identically either way. */
  const role = $derived<GuildViewerRole>(
    home?.viewer?.role ?? (home === null ? 'public' : canManage ? 'officer' : 'member'),
  );
  /** What this account may do here, the API's capability list from `/v1/me` (api
   *  `guilds.Capabilities`): the Roster tab draws Approve, Remove and Decline from it and from
   *  nothing else, so an officer who is not the guild master is never offered a guild master's
   *  controls by inference from her rank. */
  const can = $derived(myGuildMembership?.can ?? []);
  /** Officer tools gate §3.1/§4: true until the API says the guild's officer tools are off, or
   *  refuses an Award/Undo with `guild_plan_required` (the grant ended after the page loaded).
   *  `?? true` is deliberate: the server's 403 is the real gate. */
  let officerToolsEnded = $state(false);
  /** `raid_guild_required` on an Award/Undo: the guild is Guild tier, so Loot and Readiness gate to Raid guild. */
  let raidGuildEnded = $state(false);
  let planRefusal = $state<GuildPlanRefusal | null>(null);
  const officerTools = $derived(officerToolsOn(role, home?.viewer?.officer_tools, officerToolsEnded));
  const raidGuildTools = $derived(
    raidGuildToolsOn(role, home?.viewer?.raid_guild_tools ?? true, raidGuildEnded),
  );
  const officerView = $derived(role === 'officer' || role === 'moderator');
  /** The plan sentence's link (gate spec §3.2): the guild's settings tab. */
  const settingsHref = $derived(
    resolved !== null && home !== null
      ? guildSettingsHref(resolved.region, resolved.ruleset, home.guild.name)
      : '',
  );
  const guildId = $derived(home?.guild.id ?? data?.guild.id ?? null);
  /** Guild crest round (docs/contracts/2026-10-05-guild-crest-api.md): `home.guild` (the
   *  member-gated fetch) wins once it resolves -- it is the one place a just-saved/removed
   *  crest shows up without a reload (`onCrestChanged` below re-fetches `home`, nothing
   *  else) -- `data.guild` (the public fetch) is the first-paint fallback, same precedence
   *  `guildId` above uses. `guildMarkSrc` is the one helper that picks crest vs. faction
   *  logo vs. null; this component never re-derives that rule itself. */
  const guildMark = $derived(guildMarkSrc(home?.guild ?? data?.guild ?? null));
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

  /** After a roster action: the mutation dropped the cached home, so this reads the guild's
   *  own answer (roster, pending list, waiting count) rather than trusting the local edit.
   *  Best-effort: the local edit already shows the change. */
  async function refreshHome(): Promise<void> {
    if (home === null) return;
    try {
      home = await fetchGuildHome(home.guild.id);
    } catch {
      // The local edit stays.
    }
  }

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
        ...afterApprove(home, new Set([row.character_key])),
      };
      await refreshHome();
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
        ...afterApprove(home, approved),
      };
      await refreshHome();
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
      home = {
        ...home,
        roster: home.roster.filter((r) => r.character_key !== row.character_key),
        pending: (home.pending ?? []).filter((r) => r.character_key !== row.character_key),
        ...afterRemove(home, row.character_key),
      };
      await refreshHome();
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

  /** Guild crest round: `putGuildCrest`/`deleteGuildCrest` already invalidate the cached
   *  home/settings queries themselves (the same pattern every other guild mutation in
   *  `lib/guild/api.ts` follows) -- this just re-fetches so the header ring picks up the
   *  change without a page reload, mirroring `onContest`'s own refresh above. Best-effort:
   *  the crest block already showed its own success before calling this, so a failed
   *  refetch here is silent rather than a second error message fighting the first. */
  async function onCrestChanged(): Promise<void> {
    if (home === null) return;
    try {
      home = await fetchGuildHome(home.guild.id);
    } catch {
      // See comment above: silent.
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
  /** The band counts what the viewer can open: a visitor whose only report is guild-only is told
   *  nothing about bosses and pulls beside a tab reading "No public reports this week yet". */
  const showCounts = $derived(showBandCounts(role, data?.reports.length ?? 0));
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
      activeTab = tabAfterHashChange(window.location.hash, tabsForRole(role), activeTab);
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

  let raids = $state<GuildRaidsPage | null>(null);
  let raidsStatus = $state<TabFetchStatus>('idle');
  async function loadRaids(id: number): Promise<void> {
    const page = await fetchGuildRaids(id);
    raids = page;
    rememberPlannedRows(id, page, browserStorage());
  }
  /** The one raids read: a change that lands while it is in flight marks it stale, and it reads again
   *  when it settles, so a sign-up or officer action during a poll never leaves a stale list. */
  const raidsRead = createRefreshingRead(async () => {
    if (guildId !== null) await loadRaids(guildId);
  });
  /** A band press moved the Raids tab's counts and chips: an open tab (or one still loading) reads
   *  again in place (no skeleton flash); a closed one reads again when next shown. */
  function raidsChanged(): void {
    if (raidsChangeAction(guildId !== null, raidsStatus, activeTab === 'raids') === 'reread') {
      raidsRead.invalidate();
    } else raidsStatus = 'idle';
  }
  $effect(() => {
    if (activeTab !== 'raids' || guildId === null || raidsStatus !== 'idle') return;
    raidsStatus = 'loading';
    void raidsRead
      .run()
      .then(() => {
        raidsStatus = 'ready';
      })
      .catch(() => {
        raidsStatus = 'missing';
      });
  });
  /** Raids spec 4: a verified viewer's tab reserves the planned rows it last drew while it loads, from
   *  the band hold's first paint (a viewer last seen verified) until /home names the role. */
  const plannedReserve = $derived(
    guildId !== null && (role !== 'public' || holdBand)
      ? plannedReserveRows(guildId, browserStorage())
      : null,
  );
  const guildPath = $derived(
    resolved !== null && home !== null ? guildHref(resolved.region, resolved.ruleset, home.guild.name) : '',
  );

  /** The band's member lines (and the claim line /v1/me adds) arrive after the band paints; a viewer
   *  last seen verified here gets the band header's last height while /home is in flight, so the
   *  tabs under it never move (band-hold.ts). */
  const bandLayout: BandLayout =
    typeof window !== 'undefined' && window.matchMedia('(min-width: 768px)').matches ? 'desktop' : 'phone';
  const bandKey = $derived(
    resolved === null ? null : bandHoldKey(resolved.region, resolved.ruleset, resolved.slug),
  );
  const heldBand = $derived(bandKey === null ? null : readBandHold(browserStorage(), bandKey, bandLayout));
  const holdBand = $derived(
    heldBand !== null && homeStatus !== 'ready' && (!meResolved || homeStatus === 'loading'),
  );
  let bandHeaderHeight = $state(0);
  /** A refused band press's line is transient: the hold is saved only with the band at rest. */
  let bandLineStatus = $state(false);
  $effect(() => {
    if (bandKey === null || homeStatus !== 'ready' || !meResolved || bandHeaderHeight <= 0) return;
    if (bandLineStatus) return;
    if (myGuildMembership?.verified === true) {
      rememberBandHold(browserStorage(), bandKey, bandLayout, bandHeaderHeight);
    } else forgetBandHold(browserStorage(), bandKey);
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

  let readiness = $state<ReadinessBoard | null>(null);
  let readinessStatus = $state<TabFetchStatus>('idle');
  /** The next-raid toggle, on by default; held here so a re-read never resets it. */
  let readinessNextRaidOnly = $state(true);
  const readinessView = $derived(readinessTabView(readiness, readinessNextRaidOnly, home?.next_raid ?? null));
  $effect(() => {
    if (activeTab !== 'readiness' || guildId === null || readinessStatus !== 'idle') return;
    readinessStatus = 'loading';
    void loadReadinessBoard(guildId)
      .then((page) => {
        readiness = page;
        readinessStatus = 'ready';
      })
      .catch(() => {
        readinessStatus = 'missing';
      });
  });

  let loot = $state<GuildLootPage | null>(null);
  // Live-fix round, defect 5: unlike raids/progression/readiness above, Loot distinguishes
  // a real 404 (the api lane has not deployed this endpoint yet -- 'missing', the honest
  // "not live yet" copy) from every other failure ('failed' -- a network blip, a 5xx, a
  // parse error): those must surface their own message (lootError below), never the 404
  // copy, so an officer debugging a real outage is not told to wait for a deploy that
  // already happened. If the query cache already holds the page GuildOverview's own
  // summary-card prefetch resolved (fetchGuildLoot's shared cache key, data/query.ts), this
  // read returns it instantly rather than refetching.
  let lootStatus = $state<'idle' | 'loading' | 'ready' | 'missing' | 'failed'>('idle');
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
      .catch((thrown: unknown) => {
        if (thrown instanceof GuildApiError && thrown.status === 404) {
          lootStatus = 'missing';
          return;
        }
        lootStatus = 'failed';
        lootError = thrown instanceof Error ? thrown.message : 'That did not work; try again';
      });
  });

  /** §4: a `guild_plan_required` refusal flips both tabs to officer-read-only and says so once,
   *  in the clicked item's card, never in the tab-bottom error line. `raid_guild_required` flips
   *  them to the Raid guild gate sentence instead. True when it handled it. */
  function planRefused(thrown: unknown, item: GuildLootItem, kind: 'award' | 'undo'): boolean {
    if (!(thrown instanceof GuildApiError)) return false;
    const gate = refusalGate(thrown.code);
    if (gate === null) return false;
    if (gate === 'raid-guild') {
      raidGuildEnded = true;
      return true;
    }
    officerToolsEnded = true;
    if (loot !== null && loot.selected !== null) {
      planRefusal = { encounterId: loot.selected, itemId: item.item_id, kind };
    }
    return true;
  }

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
      if (planRefused(thrown, item, 'award')) return;
      lootError = thrown instanceof Error ? thrown.message : 'That did not work; try again';
    } finally {
      lootBusyItemId = null;
    }
  }

  /** The boss chooser: reads that encounter's loot. */
  async function onBoss(encounterId: number): Promise<void> {
    if (guildId === null) return;
    lootError = '';
    try {
      loot = await fetchGuildLoot(guildId, encounterId);
    } catch (thrown) {
      lootError = thrown instanceof Error ? thrown.message : 'That did not work; try again';
    }
  }

  /** Undo: removes the award the item shows, then reads the page again (the invalidation in
   *  `unawardLoot` clears the cached loot, so the read is fresh). */
  async function onUndo(item: GuildLootItem): Promise<void> {
    if (guildId === null || loot === null || loot.selected === null || item.awarded_to === null) return;
    lootBusyItemId = item.item_id;
    lootError = '';
    try {
      await unawardLoot(guildId, item.awarded_to.award_id);
      loot = await fetchGuildLoot(guildId, loot.selected);
    } catch (thrown) {
      if (planRefused(thrown, item, 'undo')) return;
      lootError = thrown instanceof Error ? thrown.message : 'That did not work; try again';
    } finally {
      lootBusyItemId = null;
    }
  }
</script>

{#if status === 'missing'}
  <h1 class="sr-only">Guild</h1>
  <!-- Keeps the loading state's box and its lede: this branch replaces the loading state on
       the client (an address that does not parse never fetches), so the server-rendered
       lede stays the first paint and a footer that jumps up the page by the skeleton's
       height (the layout shift the guild Lighthouse row measured) stays put. -->
  <div class="flex flex-col gap-3 {GUILD_LOADING.home.minHeight} {GUILD_BANDLESS_TOP}">
    <p class="text-muted text-[14px]" data-testid="guild-lede">{guildHomeCopy.lede}</p>
    <p class="text-[14px]" data-testid="guild-missing">
      That is not a guild address. They look like <code class="font-mono"
        >/guild/eu/normal/the-last-watch</code
      >.
    </p>
  </div>
{:else if status === 'loading' || status === 'failed'}
  <h1 class="sr-only">Guild</h1>
  <!-- The lede is server-rendered text: it is the page's first paint, so the largest
       contentful paint does not wait for the island's scripts (W.16). The box keeps the
       loading skeleton's height, so the lede and the skeleton rows share it and nothing
       moves when the guild arrives or the address turns out wrong. -->
  <div class="flex flex-col gap-3 {GUILD_LOADING.home.minHeight} {GUILD_BANDLESS_TOP}">
    <p class="text-muted text-[14px]" data-testid="guild-lede">{guildHomeCopy.lede}</p>
    <GuildStatus
      status={status === 'loading' ? 'loading' : 'failed'}
      error={status === 'failed' ? error : ''}
      onRetry={() => (attempt += 1)}
      lines={GUILD_LOADING.home.lines}
      minHeight=""
      testid="guild"
    />
  </div>
{:else if data !== null && resolved !== null}
  <div class="reveal flex flex-col gap-[22px] md:gap-8 {GUILD_BANDLESS_TOP}" data-testid="guild" id="guild">
    <!-- Site alignment (2026-10-29, D1): the head sits on the page ground like every other page's:
         the shared PageHead (eyebrow, the guild's mark at the left, the name, the facts line), then
         the viewer's own lines under it, then the tab strip. No band, no faction wash or watermark:
         the TabStrip's edge fade runs to the page ground. -->
    <div class="flex flex-col gap-4" data-testid="guild-header">
      <header
        class="flex flex-col gap-3"
        style:min-height={holdBand ? `${heldBand}px` : undefined}
        bind:offsetHeight={bandHeaderHeight}
        data-testid="guild-band-header"
      >
        <PageHead eyebrow={guildHomeCopy.eyebrow} title={data.guild.name}>
          {#snippet lead()}
            <FactionCrest {faction} src={guildMark} testid="guild-faction-crest" />
          {/snippet}
          {#snippet meta()}
            <span data-testid="guild-header-facts">
              <FactLine
                facts={bandFacts(
                  rulesetLabel(resolved.ruleset),
                  resolved.region,
                  showCounts ? { killed, pulls } : null,
                  home?.summary?.updated_at === undefined ? '' : updatedLine(home.summary.updated_at),
                )}
              />
            </span>
          {/snippet}
        </PageHead>

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

        <!-- Any signed-in viewer who is not a verified officer can start a claim from here;
               the claim page states the rules and who may claim. A verified officer gets
               their Guild settings link in the role line below instead. -->
        {#if meResolved && claimEntry.kind === 'claim'}
          <p class="text-muted text-[13px]" data-testid="guild-claim-entry">
            {claimEntryCopy.question}
            <a
              href={guildClaimHref(resolved.region, resolved.ruleset, data.guild.name)}
              data-testid="guild-claim-entry-link">{claimEntryCopy.action}</a
            >
          </p>
        {:else if meResolved && claimEntry.kind === 'getClaimed'}
          <p class="text-[13px]" data-testid="guild-claim-entry">
            <a
              href={guildClaimHref(resolved.region, resolved.ruleset, data.guild.name)}
              data-testid="guild-claim-entry-link">{claimEntryNotes.getClaimed}</a
            >
          </p>
        {:else if meResolved && claimEntry.kind === 'note'}
          <p class="text-muted text-[13px]" data-testid="guild-claim-note">{claimEntry.text}</p>
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
                   {Weekday}" sentence (§4.A.2) is role-gated. Kept unconditional here,
                   separate from the claim/settings row below it (unchanged v1 mechanism),
                   rather than the two fighting over the same slot. -->
            <div class="flex flex-col gap-1" data-testid="guild-overview-standing">
              {#if home.standing !== null && home.standing !== undefined}
                {@const sentence = standingSentence(home.standing)}
                <p class="flex flex-wrap items-center gap-3 text-[14px] font-semibold">
                  <!-- No line when the spec has no peer to rank against: the old "the only ... Hunter"
                         gloss read as a second spec beside the header chip's. -->
                  {#if sentence !== ''}
                    <span style={`color:${classColorVar(home.standing.class)}`}>
                      {myRosterRow !== null ? `${myRosterRow.name} · ` : ''}{sentence}
                    </span>
                  {/if}
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
              {/if}
              <!-- Raids spec 3.4: the next raid's line, for verified viewers only (R1), between
                     the standing line and the needs line. Its chunk loads alongside /home; the box
                     it fills (the resting caps, 44 desktop, 112 phone) is held until it lands, and
                     a chunk that fails leaves the plain link in its place. -->
              {#if showsBandLine(myGuildMembership?.verified === true, role)}
                {#if bandLineLazy.current !== null}
                  <bandLineLazy.current
                    guildId={home.guild.id}
                    nextRaid={home.next_raid ?? null}
                    officer={officerView}
                    {guildPath}
                    currentKey={myCharacterKey}
                    ownCharacters={home.roster.filter((r) => myCharacterKeys.has(r.character_key))}
                    defaultRole={myRosterRow?.role ?? 'dps'}
                    onChange={(next) => {
                      if (home !== null) home = { ...home, next_raid: next };
                      raidsChanged();
                    }}
                    onStatusLine={(shown) => (bandLineStatus = shown)}
                  />
                {:else if bandLineLazy.error !== '' && home.next_raid != null}
                  <a
                    href={guildRaidHref(guildPath, home.next_raid.id)}
                    class="text-gold text-[13px] font-bold"
                    data-testid="guild-raid-line-link">{home.next_raid.zone_name} →</a
                  >
                {:else if bandLineLazy.error !== '' && officerView}
                  <a href="#raids" class="text-[13px] font-bold" data-testid="guild-raid-line-plan"
                    >{raidCopy.bandNone.planLink}</a
                  >
                {:else if bandLineLazy.error === ''}
                  <div
                    class={home.next_raid == null ? 'min-h-5' : 'min-h-[112px] md:min-h-11'}
                    aria-hidden="true"
                    data-testid="guild-raid-line-reserve"
                  ></div>
                {/if}
              {/if}
              {#if role === 'member' && home.standing !== null && home.standing !== undefined}
                <!-- Guild spec 4.A.2: the fails read --text (a to-do list, not a failure); every check
                       passing reads green. -->
                <p
                  class="text-[13px] {home.standing.needs_before_next_raid.length === 0 ? '' : 'text-text'}"
                  style={home.standing.needs_before_next_raid.length === 0 ? 'color:#7bff5c' : undefined}
                  data-testid="guild-overview-needs"
                >
                  {needsBeforeSentence(
                    home.standing.needs_before_next_raid,
                    nextRaidDeadline(home.next_raid?.starts_at),
                  )}
                </p>
              {/if}
            </div>
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
              {:else if myGuildMembership !== null && !myGuildMembership.verified}
                <GuildUnverifiedNote
                  guildName={home.guild.name}
                  hasApprover={home.viewer?.verify?.has_approver}
                  claimHref={guildClaimHref(resolved.region, resolved.ruleset, home.guild.name)}
                />
              {/if}
            </div>

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

    {#if activeTab === 'overview'}
      <GuildOverview
        {guildId}
        {role}
        reports={home?.reports ?? data.reports}
        summary={home?.summary}
        {showCounts}
        onSelectTab={selectTab}
      />
    {:else if activeTab === 'roster' && home !== null}
      {#if rosterLazy.current}
        <rosterLazy.current
          roster={home.roster}
          pending={home.pending ?? []}
          members={home.members}
          counts={home.counts}
          {can}
          {myCharacterKey}
          {rosterBusy}
          {rosterActionError}
          frozen={home.claim.frozen}
          {removeOverridesFrozen}
          onApprove={(row) => void onApprove(row)}
          onRemove={(row) => void onRemove(row)}
          onDecline={(row) => void onRemove(row)}
          onApproveAll={() => void onApproveAll()}
          {rowPath}
          addonSnapshot={home.addon_snapshot ?? null}
          mySiteSpec={myCharacterKey === null ? null : (mySiteSpecs.get(myCharacterKey) ?? null)}
          attendanceSource={attendanceSourceOf(home)}
        />
      {:else}
        <Skeleton lines={6} minHeight="min-h-[320px]" testid="guild-roster-skeleton" />
        {@render lazyFallback(rosterLazy)}
      {/if}
    {:else if activeTab === 'raids'}
      {#if raidsStatus === 'loading' || (raidsStatus === 'ready' && !raidsLazy.current)}
        {#if plannedReserve !== null}
          <div
            class={PLANNED_RESERVE}
            style:--planned-rows={plannedReserve}
            aria-hidden="true"
            data-testid="guild-planned-reserve"
          ></div>
        {/if}
        <Skeleton lines={5} minHeight="min-h-[240px]" testid="guild-raids-skeleton" />
        {@render lazyFallback(raidsLazy)}
      {:else if raidsStatus === 'missing'}
        <EmptyState
          message="Raid night detail is not available yet. Check back after the next sync."
          testid="guild-raids-missing"
        />
      {:else if raidsLazy.current}
        <raidsLazy.current
          rows={raids?.rows ?? []}
          {guildId}
          publicView={role === 'public'}
          planned={raids?.planned}
          zones={raids?.zones ?? []}
          guildTimezone={raids?.guild_timezone ?? null}
          discord={raids?.discord}
          officer={role === 'officer'}
          {officerTools}
          memberView={myGuildMembership?.verified === true}
          {settingsHref}
          {guildPath}
          onPlanned={() => (guildId === null ? Promise.resolve() : loadRaids(guildId))}
          onPlanRefused={() => (officerToolsEnded = true)}
        />
      {/if}
    {:else if activeTab === 'progression'}
      {#if progressionStatus === 'loading' || (progressionStatus === 'ready' && !progressionLazy.current)}
        <Skeleton lines={5} minHeight="min-h-[240px]" testid="guild-progression-skeleton" />
        {@render lazyFallback(progressionLazy)}
      {:else if progressionStatus === 'missing' && (data?.roster_best ?? []).length === 0}
        <EmptyState
          message="Progression detail is not available yet. Check back after the next sync."
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
          message="Readiness checks are not available yet. Check back after the next sync."
          testid="guild-readiness-missing"
        />
      {:else if readinessLazy.current}
        <readinessLazy.current
          rows={readinessView.rows}
          raid={readinessView.raid}
          nextRaidOnly={readinessNextRaidOnly}
          onToggleNextRaid={() => (readinessNextRaidOnly = !readinessNextRaidOnly)}
          raidHref={readinessView.raidId === null ? null : guildRaidHref(guildPath, readinessView.raidId)}
          guildTimezone={readinessView.guildTimezone}
          officer={officerView}
          {officerTools}
          {raidGuildTools}
          {myCharacterKey}
          {settingsHref}
        />
      {/if}
    {:else if activeTab === 'loot'}
      {#if lootStatus === 'loading' || (lootStatus === 'ready' && !lootLazy.current)}
        <Skeleton lines={5} minHeight="min-h-[240px]" testid="guild-loot-skeleton" />
        {@render lazyFallback(lootLazy)}
      {:else if lootStatus === 'missing'}
        <EmptyState
          message="Loot ranking is not available yet. Check back after the next sync."
          testid="guild-loot-missing"
        />
      {:else if lootStatus === 'failed'}
        <!-- Live-fix round, defect 5: a network/parse/5xx failure is never the 404 "not
             live yet" copy above -- its own message, rendered without waiting on the lazy
             GuildLoot.svelte chunk (which a slow connection may not have finished loading
             yet when the fetch itself has already failed). -->
        <p class="text-[13px]" role="alert" data-testid="guild-loot-error">{lootError}</p>
      {:else if lootLazy.current}
        <lootLazy.current
          loot={lootStatus === 'ready' ? loot : null}
          officer={officerView}
          {officerTools}
          {raidGuildTools}
          memberReadOnly={role === 'member'}
          attendanceSource={attendanceSourceOf(home)}
          busyItemId={lootBusyItemId}
          refusal={planRefusal}
          {settingsHref}
          onAward={(item, charKey) => void onAward(item, charKey)}
          onUndo={(item) => void onUndo(item)}
          onBoss={(encounterId) => void onBoss(encounterId)}
        />
        {#if lootError !== ''}<p class="text-[13px]" role="alert" data-testid="guild-loot-error">
            {lootError}
          </p>{/if}
      {/if}
    {:else if activeTab === 'settings' && officerView && home !== null}
      {#if settingsLazy.current}
        <settingsLazy.current
          path={resolved}
          {home}
          {role}
          {isClaimantAccount}
          {contestBusy}
          {contestError}
          {showContestConfirm}
          {officerTools}
          onPlanRefused={() => (officerToolsEnded = true)}
          onCrestChanged={() => void onCrestChanged()}
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

    {#if rosterGate !== null}
      <GuildRosterGate state={rosterGate} />
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
