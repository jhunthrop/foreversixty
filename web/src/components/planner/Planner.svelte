<!-- web/src/components/planner/Planner.svelte -->
<!-- The planner island (spec 2026-11-08, mocks design/mocks/web/web-planner*). Mounted by
     src/pages/planner.astro (client:after-paint) and by src/planner-island.ts on the API-rendered
     /b/:id page, where `record` is supplied and the build starts read-only; Top Gear's inline "add a
     build" mounts it with `standalone` false. The page is L3 Workbench of panels (directions mock
     b2-planner): the PageHeader (crest, "{Class} talents", the meta line, the answer: the points
     and Sim this build; the simulator owns the numbers, so the live estimate runs only in the
     embed), the ControlBar (the nine class crests, the race), one note line, then the framed
     Talents panel (TalentTrees, edit) and Point order in main, and in the rail the Level build, In the
     game, Share and Import. Below 1280 the regions flow in one column in phone order (Ordered). -->
<script lang="ts">
  import { onMount, untrack } from 'svelte';
  import activeBuild from '../../data/active-build.json';
  import { fetchMeOnce, type Me, type MeCharacter } from '../../lib/account/api';
  import { adoptHeldCharacter } from '../../lib/account/adopt-held';
  import { selectedCharacter } from '../../lib/account/hero-character';
  import { recordPastedExport } from '../../lib/character-selector/paste';
  import { accountCharacterCode } from '../../lib/planner/account-character';
  import { characterPointsSpent } from '../../lib/planner/send-state';
  import { healWouldDiscardPlannerWork } from '../../lib/planner/heal-guard';
  import { mainCharacter } from '../../lib/account/main-character';
  import { createQueryState } from '../../lib/data/query.svelte';
  import { clearCurrent, onCurrentCharacterChange, readCurrent } from '../../lib/current-character';
  import { API_BASE_URL, DEFAULT_CLASS_SLUG } from '../../lib/planner/config';
  import type { BandDiffView } from '../../lib/planner/band-compare';
  import { bandTitleForDraft, loadFromBand } from '../../lib/planner/band-compare';
  import {
    decidePlannerLoad,
    followPointerInPlanner,
    isBarePlannerUrl,
  } from '../../lib/planner/current-character-planner';
  import { bandRangeLabel } from '../../lib/planner/band-title';
  import { bandForLevel } from '../../lib/bis/hover';
  import { plannerCopy } from '../../lib/planner/copy';
  import { uiCopy } from '../../lib/ui/copy';
  import { levelReached, ranksByTalent } from '../../lib/planner/derive';
  import {
    decodeFS1,
    encodeFS1,
    encodeFS1V2,
    orderFromRanks,
    remapTreeRanksByTalentId,
  } from '../../lib/planner/fs1';
  import { createLiveDps } from '../../lib/planner/live-dps.svelte';
  import { isConstrainedDevice, liveGate } from '../../lib/planner/live-gate';
  import { DATA_LOAD_FAILED, DataLoadError, loadReference, loadTalents } from '../../lib/planner/load';
  import { indexTalents, type TalentIndex } from '../../lib/planner/rules';
  import { createPlannerStore, READ_ONLY_REASON } from '../../lib/planner/store.svelte';
  import { yieldToMain } from '../../lib/yield';
  import { treeSourceNotice } from '../../lib/planner/tree-source';
  import { plannerSearchFor } from '../../lib/planner/url';
  import type { BuildRecord, Gear, TalentFile } from '../../lib/planner/types';
  import { characterFromPlanner, specOf } from '../../lib/sim/character';
  import { classDisplayName, specDisplayName, specLabel, specRow } from '../../lib/sim/spec-label';
  import { defaultSimState, simSearch, withSimState } from '../../lib/sim/url';
  import { importFromAddon, type ImportOutcome } from '../../lib/addon/import';
  import type { BisBand } from '../../lib/bis/types';
  import { bandCompareCopy, plannerCharacterCardCopy, plannerHeaderCopy } from '../../lib/planner/copy';
  import { editorTrees, GAME_TIERS, treeIconSrc } from '../../lib/planner/editor-trees';
  import { bandCodeGear, earnedOrder, estimateCharacter } from '../../lib/planner/estimate';
  import type { LadderOrderFor, LadderOrderForRanks } from '../../lib/planner/ladder-load';
  import { dataUrl } from '../../lib/planner/load';
  import SpellIcon from '../ui/SpellIcon.svelte';
  import { classTone, LINK, type ControlGroup } from '../../lib/ui/primitives';
  import { REGION } from '../../lib/ui/regions';
  import { openTreeIndex, type TalentTreesCopy } from '../../lib/ui/talent-tree';
  import classRows from '../../data/classes.json';
  import Crest from '../ui/Crest.svelte';
  import Button from '../ui/Button.svelte';
  import ControlBar from '../ui/ControlBar.svelte';
  import Notice from '../ui/Notice.svelte';
  import Ordered from '../ui/Ordered.svelte';
  import Page from '../ui/Page.svelte';
  import PageHeader from '../ui/PageHeader.svelte';
  import Plate from '../ui/Plate.svelte';
  import Section from '../ui/Section.svelte';
  import TalentTrees from '../ui/TalentTrees.svelte';
  import BandCompare from './BandCompare.svelte';
  import ImportBox from './ImportBox.svelte';
  import InGame from './InGame.svelte';
  import PlannerAnswer from './PlannerAnswer.svelte';
  import GuideBuilds from './GuideBuilds.svelte';
  import type { GuideBuild } from '../../lib/planner/guide-builds';
  import PlannerOrder from './PlannerOrder.svelte';
  import SharePanel from './SharePanel.svelte';
  import SummaryBar from './SummaryBar.svelte';
  import './planner.css';

  let {
    treeVersion,
    classSlug = DEFAULT_CLASS_SLUG,
    raceSlug,
    record = null,
    gear,
    oncode,
    standalone = true,
    openingArt = [],
    openingTrees = [],
    title = null,
    views = null,
    hideBar = false,
    guideBuilds = [],
  }: {
    treeVersion: string;
    classSlug?: string;
    raceSlug?: string;
    record?: BuildRecord | null;
    /**
     * Seeds the store's gear when there is no `record` (dps D39/D40) -- Top Gear's inline
     * "add a build" (TalentCandidates.svelte) passes the loaded character's gear so its live
     * DPS card sims the same equipment the comparison table does. `record?.gear` always wins
     * when a record is present: `/b/:id` and `/planner` must stay byte-identical to today.
     */
    gear?: Gear;
    /**
     * Called with the build's own FS1 code whenever it changes. Top Gear's "add a build"
     * (Task 14's TalentCandidates) mounts this component inline and reads the code back
     * through it; /planner and /b/:id pass nothing and the callback never fires.
     */
    oncode?: (code: string) => void;
    /** False only for Top Gear's inline "add a build" (Task 10): no chip, no restore, no write. */
    standalone?: boolean;
    /** The opening class's tree art (`trees/<background>.webp`, tab order): the loading Plate
     *  paints it so the largest paint is in the server render (planner.astro reads it). */
    openingArt?: readonly string[];
    /** The opening class's tree names, left to right: the loading Plate's heads, each at 0 points,
     *  so the first paint reads the trees the island is about to draw. */
    openingTrees?: readonly string[];
    /** The shared build's own name (`/b/:id`, premium spec 2026-11-08 decision B1): the H1 in
     *  place of `{Class} talents`; a record with none reads `{Race} {Class}`. */
    title?: string | null;
    /** The shared build's view count, on its meta line. */
    views?: number | null;
    /** No class and race bar (decision B3): a shared build's class and race are the record's. */
    hideBar?: boolean;
    /** Every spec guide's published builds (planner.astro reads the guides): the rail's Guide builds. */
    guideBuilds?: readonly GuideBuild[];
  } = $props();

  /** The shared build page (`/b/:id`, spec section 8): its own head, the Edit a copy line, no bar,
   *  In the game first and the level build set against it as the `Guide build`. */
  const shared = record !== null;

  // The page is static, so ?class= and ?race= can only be read in the browser. A record
  // (from /b/:id) wins over the query string: that build already names its class and race.
  function fromQuery(name: string): string | undefined {
    if (typeof window === 'undefined') return undefined;
    return new URLSearchParams(window.location.search).get(name) ?? undefined;
  }

  // Rebuild spec §4.I: `?spec=` (new, from the BiS page/a guide/the home Talents card)
  // derives `?class=` when `?class=` itself is absent -- `specRow(spec).class_slug`, the
  // same spec catalogue `specKeyFor`/`specOf` already read. Read once, untracked, same as
  // every other URL param this component reads on mount.
  const specParam = fromQuery('spec') ?? null;
  const specClassSlug = specParam !== null ? (specRow(specParam)?.class_slug ?? undefined) : undefined;
  // `?talents=` (new): a band's own digit string, applied once on arrival -- see the
  // `talentsApplied` flag and its one call site inside `load()`, below.
  const talentsParam = fromQuery('talents') ?? null;
  let talentsApplied = false;
  // The ladder orders (/data/leveling-plans.json): a band link is spent in the sim's and the
  // guide's own order, not a reconstruction. Only read for a `?talents=` link, on demand (its
  // own chunk), and awaited by `load()` beside the talent data so the link never misses it.
  const noLadder: LadderOrderFor = () => undefined;
  const ladderReady: Promise<LadderOrderFor> =
    talentsParam === null || typeof window === 'undefined'
      ? Promise.resolve(noLadder)
      : import('../../lib/planner/ladder-load').then((m) => m.loadLadderOrderFor()).catch(() => noLadder);
  // A guide's `Load this build` names its spec (`?code=&spec=`): the build is spent in the guide's
  // own levelling order when the spec's ladder reaches exactly the code's ranks (spec 2026-11-08
  // ruling SHOULD), read on demand like the band ladder above.
  const noRanksLadder: LadderOrderForRanks = () => undefined;
  const guideLadderReady: Promise<LadderOrderForRanks> =
    specParam === null || fromQuery('code') === undefined || typeof window === 'undefined'
      ? Promise.resolve(noRanksLadder)
      : import('../../lib/planner/ladder-load')
          .then((m) => m.loadLadderOrderForRanks())
          .catch(() => noRanksLadder);

  // Current-character pointer (Task 10, spec section 1) -- decision in current-character-
  // planner.ts. `codeParam` is `?code=` itself, or a restored pointer's code, same path.
  const plannerLoad = untrack(() => {
    const urlCode = record ? null : (fromQuery('code') ?? null);
    const search = typeof window === 'undefined' ? '' : window.location.search;
    return decidePlannerLoad(
      urlCode,
      isBarePlannerUrl(search),
      record !== null,
      standalone,
      readCurrent(),
      fromQuery('character') ?? null,
    );
  });
  const { decoded } = plannerLoad;
  // A dead restored pointer is forgotten, not shown as an error. Runs once, on mount.
  $effect(() => {
    if (plannerLoad.deadPointer) clearCurrent();
  });

  // A parsed count inside one of these renders in a `numeric` span, the same as every
  // other number the planner shows (OrderStrip, SummaryBar, BandCompare, TalentCell).
  // This stays a small discriminated union rather than a plain string so the template can wrap
  // just the digits with a real element -- `decoded.message` (current-character-planner.ts's
  // own `decodeFS1` call, above) stays flat text (its exact wording is pinned by fs1.test.ts
  // and is the shape Task 14 reads), so the split happens once, here, rather than by
  // reformatting arbitrary text at render time.
  type CodeNote =
    | { kind: 'message'; text: string }
    | { kind: 'tree-count'; got: string; want: string }
    // `olderBuild` is set only when the code's own stamped build differs from the active
    // one AND that older build's talent file could not be loaded to remap by id (see the
    // `remapTreeRanksByTalentId` call in `load()`, below) -- the ranks were then read
    // positionally against the active build same as before that remap existed, which can
    // misplace a point if the two builds' trees do not match tab for tab.
    | { kind: 'reconstructed'; dropped: number; gearOnly: boolean; olderBuild: boolean }
    // Fix round 2 (ux-designer review finding 2): a band load (`?talents=`, or BandCompare's
    // own Load button) is never a character or an addon import -- the 'reconstructed' note's
    // own "Talents loaded from a character" wording is false for it, so it gets its own kind
    // and its own copy rather than a borrowed branch of that one.
    | { kind: 'band'; bandLabel: string }
    // A character whose recorded ranks reach further than any legal order can: the note names
    // the count the planner kept against the count the record carries.
    | { kind: 'kept'; kept: number; total: number; name: string }
    // Bug B1: a code that names no character (a guide's Load this build, a shared ?code=) is a
    // link, never "Talents loaded from a character"; a guide's link that names its spec arrives
    // in that guide's own levelling order.
    | { kind: 'linked' }
    | { kind: 'guide'; specName: string };

  /** Set when the build came in as an FS1 code: either why it could not be read, or that its
   *  order is a reconstruction (nothing in the game records the order points were spent in). */
  let codeNote = $state<CodeNote | null>(null);

  /** The account character whose build is on screen: the Sim link speaks for it. */
  let loadedCharacter = $state<MeCharacter | null>(null);

  /** Set when the build on screen is somebody else's (a linked or guild build, never the
   *  account character's own): `owner` is their name, '' for a link that names nobody. The
   *  send button reads it: another player's build asks before it replaces anything. */
  let buildOrigin = $state<{ owner: string } | null>(null);
  // A build with nothing spent is nobody's.
  $effect(() => {
    if (store.spent === 0 && buildOrigin !== null) untrack(() => (buildOrigin = null));
  });

  // Matches the one message `decodeFS1` returns that carries two counts, so the digits can be
  // pulled out and wrapped here without `decodeFS1` itself having to return anything but a
  // flat, tested string.
  const TREE_COUNT_MESSAGE = /^That code has (\d+) talent trees; a build has (\d+)\.$/;

  function noteForMessage(message: string): CodeNote {
    const match = TREE_COUNT_MESSAGE.exec(message);
    return match ? { kind: 'tree-count', got: match[1], want: match[2] } : { kind: 'message', text: message };
  }

  // A code has exactly one turn on the class it names, tracked by two things together rather
  // than one flag armed early:
  //
  //  - Every `load()` run only ever considers applying the code when *this run's own class*
  //    matches the class the code names. A run for any other class leaves it alone no matter
  //    how the timing between two `load()` calls falls out, which is what stops a class switch
  //    (immediate, or after a failed load) from ever applying a code meant for a different
  //    class -- `orderFromRanks` only knows tab positions, not which class they belong to, so
  //    it would apply them onto the wrong tree without complaint if it were reachable at all.
  //
  //  - `codeApplied` gates *reuse within that one matching class*. It is armed only once the
  //    code has genuinely had its turn: applied successfully, or permanently refused because
  //    the class it names has no talent data (retrying that can never succeed). It is left
  //    unarmed by any other failure on that same class -- a network blip on the reference
  //    files, a 5xx on talents, sets or items -- so a same-class Retry still gets to apply the
  //    code. Arming it unconditionally on the first attempt, whatever the failure, would drop
  //    a perfectly good code the moment an unrelated transient error hit first and leave no
  //    trace once the retry quietly succeeded onto an empty default build.
  let codeApplied = false;

  // The store is seeded once, from the props as they arrive. `untrack` says so: without it
  // the compiler warns that these reads only capture the initial value, which is the point --
  // after mount the store owns the class, the race and the order, not the props.
  const store = untrack(() =>
    createPlannerStore({
      treeVersion: record?.tree_version ?? treeVersion,
      classSlug: record
        ? classSlug
        : decoded?.ok
          ? decoded.build.classSlug
          : (fromQuery('class') ?? specClassSlug ?? plannerLoad.initialClassSlug ?? classSlug),
      raceSlug: record
        ? (raceSlug ?? '')
        : decoded?.ok
          ? decoded.build.raceSlug
          : (fromQuery('race') ?? raceSlug ?? ''),
      order: record?.point_order,
      gear: record?.gear ?? gear,
      title: record?.title,
      sourceId: record?.id ?? null,
      readOnly: record !== null,
    }),
  );

  // Spec 4.2: "and when there is none, on the main's class." Only reached when there was no
  // pointer, no ?code=, no ?class= and no record at all (plannerLoad.initialClassSlug is
  // null exactly then) -- a synchronous pointer-derived class (any source) already won in
  // the store's own construction above and this never overrides it.
  const session = standalone
    ? createQueryState<Me | null>(`${API_BASE_URL}/v1/me`, () => fetchMeOnce(), {
        scope: 'private',
        ttlMs: 10 * 60 * 1000,
      })
    : null;
  /** A link that names a build (`?code=` from the URL, `?talents=`, `?class=`/`?spec=`) owns
   *  the planner: the signed-in main-class default and the hero's own level never apply to it.
   *  Live defect 2026-10-05: a guide's "Load this build" (`?code=` for a Paladin) opened
   *  signed in as a Warrior main; this effect switched class while the talent file was still
   *  loading, so the code no longer matched the class and was silently dropped. */
  const linkNamesBuild =
    (plannerLoad.codeParam !== null && !plannerLoad.restored) ||
    talentsParam !== null ||
    fromQuery('class') !== undefined ||
    specParam !== null;
  $effect(() => {
    if (session === null || linkNamesBuild || plannerLoad.initialClassSlug !== null || record !== null)
      return;
    const main =
      session.data === null ? null : mainCharacter(session.data.characters, session.data.main_character_key);
    const mainClassSlug = main?.class?.toLowerCase();
    // Only correct a build the visitor has not touched yet: no points spent, no race chosen
    // beyond the class's own default, and still on the class this mount opened with.
    if (mainClassSlug === undefined || mainClassSlug === store.classSlug || store.order.length > 0) return;
    store.selectClass(mainClassSlug);
  });

  // Fix round 2 (wow-player review finding 1): Level must read the same real character the
  // character header beside this page shows, not a reconstructed figure that can openly
  // disagree with it. The same hero-character lookup `createCharacterCardState` (the
  // character header's own hook) makes, read from this mount's own session so the two
  // panels can never pick two different characters. `null` (an anonymous visit, or before
  // the session resolves) leaves `store.level` on its existing `levelReached(order)`
  // fallback, unchanged for every bare build.
  // Bumped by every selector change so the level follows the newly selected character.
  let pointerVersion = $state(0);
  $effect(() => {
    void pointerVersion;
    if (session === null || linkNamesBuild) return;
    const hero =
      loadedCharacter ?? (session.data === null ? null : selectedCharacter(readCurrent(), session.data));
    store.setCharacterLevel(hero?.level ?? null);
    // The race picker starts on the character's race, as the class already follows it.
    // A link or a saved build names its own race; a manual pick wins (store.followRace).
    if (record === null) untrack(() => store.followRace(hero?.race ?? null));
  });

  /** The header selector's character (spec 3.6: the character lives in the header): the send is
   *  aimed at it. Null signed out, with none synced, or while the session answers. */
  const headerCharacter = $derived.by((): MeCharacter | null => {
    void pointerVersion;
    const me = session?.data ?? null;
    return me === null ? null : selectedCharacter(readCurrent(), me);
  });
  /** The character In the game sends to: the header's, read-only builds included. */
  const sendCharacter = $derived(headerCharacter);
  /** The same fact with the account's other characters counted: a build loaded from one of the
   *  player's own characters is another character's build beside a different current one. */
  const effectiveOrigin = $derived(
    buildOrigin ??
      (loadedCharacter !== null && headerCharacter !== null && loadedCharacter.key !== headerCharacter.key
        ? { owner: loadedCharacter.name }
        : null),
  );

  // The chip's "Copy for the game" link; shared with SharePanel's own buttons (Task 10, L8.1).

  // An addon export for another class: switch the planner to it and import once that
  // class's talents have loaded, instead of telling the player to switch by hand. The
  // pending code is cleared before anything else so a failed deferred import cannot retry
  // on every later talent load.
  // `origin` says who asked: a 'paste' is the player choosing an export (it earns the pointer), a
  // 'follow' is the header selector moving the page (the pointer is already right), a 'character'
  // is an account character opening with its own build.
  let pendingImport = $state<{
    classSlug: string;
    code: string;
    origin: 'paste' | 'follow' | 'character';
    character?: MeCharacter;
  } | null>(null);
  function switchClassForImport(
    classSlug: string,
    code: string,
    origin: 'paste' | 'follow' | 'character' = 'paste',
    character?: MeCharacter,
  ): void {
    pendingImport = { classSlug, code, origin, character };
    buildOrigin = null;
    if (origin !== 'character') characterLoading = false;
    store.selectClass(classSlug);
  }

  /** Opens the account character `key` with its own talents, gear, race and level. */
  // Each load (and each other move of the pointer) takes a number; a load that finds a newer one
  // when its reads come back stops, so the character chosen last is the one that opens.
  let characterRequest = 0;
  /** True from the moment a character is asked for until its build is on screen. */
  let characterLoading = $state(false);
  async function loadAccountCharacter(key: string, abandon?: () => boolean): Promise<void> {
    const mine = ++characterRequest;
    characterLoading = true;
    const me = await fetchMeOnce().catch(() => null);
    const character = me?.characters.find((row) => row.key === key);
    if (character === undefined || mine !== characterRequest) {
      if (mine === characterRequest) characterLoading = false;
      return;
    }
    const classSlug = character.class?.toLowerCase() ?? '';
    const code = await accountCharacterCode(character, activeBuild.build, API_BASE_URL);
    // A heal's load stops when the visitor edited while the code was coming: their build stays.
    if (mine !== characterRequest || abandon?.() === true) {
      if (mine === characterRequest) characterLoading = false;
      return;
    }
    if (code !== null && decodeFS1(code).ok) {
      switchClassForImport(classSlug, code, 'character', character);
      return;
    }
    // The account holds no build for it: its class and race are still its own.
    if (classSlug !== '') store.selectClass(classSlug);
    loadedCharacter = character;
    store.followRace(character.race ?? null);
    characterLoading = false;
  }
  if (plannerLoad.characterKey !== null) {
    const key = plannerLoad.characterKey;
    $effect(() => {
      untrack(() => void loadAccountCharacter(key));
    });
  }

  async function applyCharacterImport(
    pending: { classSlug: string; code: string; character?: MeCharacter },
    index: TalentIndex,
  ): Promise<void> {
    const decoded = decodeFS1(pending.code);
    if (!decoded.ok) return;
    let treeRanks = decoded.build.treeRanks;
    let remapDropped = 0;
    // The export may be from another client build than the trees on screen; the same by-id remap
    // a `?code=` gets keeps each point on its own talent.
    if (decoded.build.dataBuild !== store.treeVersion) {
      try {
        const older = await loadTalents(decoded.build.dataBuild, pending.classSlug);
        if (store.classSlug !== pending.classSlug) return;
        const remap = remapTreeRanksByTalentId(indexTalents(older), index, treeRanks);
        treeRanks = remap.treeRanks;
        remapDropped = remap.dropped.length;
      } catch {
        // Positional, as before the remap existed.
      }
    }
    if (store.classSlug !== pending.classSlug) {
      characterLoading = false;
      return;
    }
    const rebuilt = orderFromRanks(index, treeRanks);
    store.loadImported({
      classSlug: pending.classSlug,
      raceSlug: decoded.build.raceSlug,
      order: rebuilt.order,
      gear: decoded.build.gear,
    });
    loadedCharacter = pending.character ?? null;
    const total = pending.character === undefined ? null : characterPointsSpent(pending.character);
    codeNote =
      pending.character !== undefined && total !== null && store.spent < total
        ? { kind: 'kept', kept: store.spent, total, name: pending.character.name }
        : {
            kind: 'reconstructed',
            dropped: rebuilt.dropped.length + remapDropped,
            gearOnly: false,
            olderBuild: false,
          };
    characterLoading = false;
  }
  // The header selector moves the open planner along, the way a fresh load of that character
  // would (followPointerInPlanner): a pasted export imports, any other character opens its
  // class. The address follows through the class/race effect below; the selector already
  // replaced it for a character chosen there.
  // The session's default must not replace a link's build or the visitor's own edits.
  const healDiscardsWork = (): boolean =>
    healWouldDiscardPlannerWork({
      linkNamesBuild,
      pointsSpent: store.order.length,
      gearCount: Object.keys(store.gear).length,
    });
  $effect(() => {
    if (!standalone || record !== null) return;
    return onCurrentCharacterChange(({ fromPageLoad, fromSessionHeal }) => {
      pointerVersion += 1;
      if (fromPageLoad) return;
      if (fromSessionHeal && healDiscardsWork()) return;
      const follow = followPointerInPlanner(readCurrent(), store.classSlug);
      if (follow.kind !== 'none') {
        loadedCharacter = null;
        buildOrigin = null;
        characterRequest += 1;
      }
      if (follow.kind === 'character')
        void loadAccountCharacter(follow.key, fromSessionHeal ? healDiscardsWork : undefined);
      else if (follow.kind === 'import') switchClassForImport(follow.classSlug, follow.code, 'follow');
      else if (follow.kind === 'class') store.selectClass(follow.classSlug);
    });
  });
  $effect(() => {
    const pending = pendingImport;
    const index = store.talentIndex;
    if (pending === null || index === null || index.file.class_slug !== pending.classSlug) return;
    pendingImport = null;
    if (pending.origin === 'character') {
      void applyCharacterImport(pending, index);
      return;
    }
    const result = importFromAddon(pending.code, index, activeBuild.build);
    if (!result.ok) return;
    store.loadImported({
      classSlug: result.classSlug,
      raceSlug: result.raceSlug,
      order: result.order,
      gear: result.gear,
    });
    if (pending.origin === 'paste') void recordPastedExport(pending.code);
  });

  // The live DPS estimate (Task 20). Created once -- createLiveDps holds no pool until the
  // first request, so this costs nothing on mount and does not touch engine.ts until a talent
  // or slot actually changes. `untrack` for the same reason the store above needs it: the
  // factory call itself reads nothing reactive, and without it the compiler would warn that a
  // value read once here looks like a dependency.
  const live = untrack(() => createLiveDps());

  // Every edit to the build re-requests an estimate; live.request debounces the burst into
  // one run and cancels whatever was already in flight. `characterFromPlanner` returns null
  // until the talent file has loaded, which `request` treats as "off" rather than an error.
  // ...but only for a build worth simming, on a device that wants it (live-gate.ts): until
  // every point is spent nothing is asked of the engine, so nothing is downloaded either, and
  // a phone is asked first. Stopping keeps the last figure on screen, dimmed.
  const constrained = untrack(() => isConstrainedDevice());
  let dpsOptedIn = $state(false);
  // Every role has a live figure: damage, a healer's effective HPS or a tank's score, each
  // computed in the browser by the same sim/score code the nightly ranker uses.
  const gate = $derived(
    liveGate({ spent: store.spent, needed: store.pointsAvailable, constrained, optedIn: dpsOptedIn }),
  );
  /** The band the Level build section compares against, once its file has answered: the live
   *  estimate sims that band's gear (decision P1, gap G1). Undefined until the section reports. */
  let estimateBand = $state<{ band: BisBand | null; label: string } | undefined>(undefined);

  // A figure from another class is not a stale answer to this build. Declared BEFORE the
  // request effect below on purpose: effects flush in declaration order, and a build pasted
  // in for another class changes the class and the talents in one go -- clearing second
  // would cancel the very run that paste had just scheduled.
  $effect(() => {
    void store.classSlug;
    untrack(() => live.clear());
  });

  $effect(() => {
    void store.order;
    void store.gear;
    const index = store.talentIndex;
    const base = gate === 'run' && index !== null ? characterFromPlanner(store) : null;
    // Gap G7: a character's estimate is simmed at its level with only the points it has earned.
    // Only where the sim tools embed the planner: the standalone page links to the simulator
    // instead of showing its number (one job per page), so it never wakes the engine.
    if (!standalone && base !== null && index !== null)
      live.request(estimateCharacter(base, index, estimateBand?.band ?? null, store.characterLevel), index);
    else live.request(null, null);
  });

  // No reactive reads of its own: this effect's body runs once, on mount, purely to register
  // the teardown that runs it returns -- which is the only thing that has to happen when the
  // planner unmounts, since `live.request` above already cancels and re-schedules on its own.
  $effect(() => () => live.dispose());

  /** The current build's talent ranks, one array per tree in tab order -- encodeFS1's shape. */
  function treeRanksFor(index: TalentIndex, order: number[]): number[][] {
    const ranks = ranksByTalent(order);
    return index.trees.map((tree) => tree.talents.map((talent) => ranks.get(talent.id) ?? 0));
  }

  /**
   * The build's own FS1 code -- the planner's export format. Derived once so `simHref`
   * (below) and an embedder's `oncode` (Top Gear's "add a build") always read the identical
   * encoding of the identical build, rather than each calling `encodeFS1` again and risking
   * the two drifting apart.
   */
  const liveCode = $derived(
    store.talentIndex === null
      ? ''
      : encodeFS1({
          dataBuild: store.treeVersion,
          classSlug: store.classSlug,
          raceSlug: store.raceSlug,
          treeRanks: treeRanksFor(store.talentIndex, store.order),
          gear: store.gear,
        }),
  );

  // A `$effect` rather than a call inside a derivation: a derivation must stay a pure read,
  // and calling `oncode` is a side effect that has to run again on every build change.
  $effect(() => {
    if (liveCode !== '') oncode?.(liveCode);
  });

  // The build of an account character, as the export the simulator reads: the planner's talents
  // and gear under the character's own name, race and level, so the simulator opens on that
  // character rather than on a nameless code.
  const characterBuildCode = $derived.by(() => {
    const character = loadedCharacter;
    if (character === null || store.talentIndex === null) return '';
    if (character.class?.toLowerCase() !== store.classSlug) return '';
    return encodeFS1V2({
      dataBuild: store.treeVersion,
      classSlug: store.classSlug,
      raceSlug: store.raceSlug,
      treeRanks: treeRanksFor(store.talentIndex, store.order),
      level: store.level,
      gear: store.gear,
      character: { name: character.name, realm: character.realm ?? '' },
    });
  });
  // Moving off the character's class is leaving its build behind.
  $effect(() => {
    if (loadedCharacter !== null && loadedCharacter.class?.toLowerCase() !== store.classSlug) {
      untrack(() => (loadedCharacter = null));
    }
  });
  /** The head's estimate as a code (decision P1): the build on the Level build's gear at the
   *  character's level, so `Sim this build ›` sims what the head shows. Only for a build with no
   *  gear of its own and no account character behind it. */
  const estimateCode = $derived.by(() => {
    const band = estimateBand?.band ?? null;
    if (band === null || store.talentIndex === null || Object.keys(store.gear).length > 0) return '';
    const worn = bandCodeGear(band);
    return encodeFS1V2({
      dataBuild: store.treeVersion,
      classSlug: store.classSlug,
      raceSlug: store.raceSlug,
      treeRanks: treeRanksFor(store.talentIndex, earnedOrder(store.order, store.characterLevel)),
      level: store.characterLevel ?? 60,
      gear: worn.gear,
      gearSlots: worn.gearSlots,
    });
  });
  /** Nothing to sim: no points spent and no character behind the build. */
  const simLinkShown = $derived(!standalone || store.spent > 0 || loadedCharacter !== null);

  // "Sim this build": a saved build's own link when it has one, otherwise `liveCode` above.
  // The link is always present so a player can reach the full results whether or not the
  // live estimate has run, or could run at all. Both branches go through
  // `simSearch`/`withSimState` rather than building the query string by hand, so /sim's own
  // URL state (lib/sim/url.ts) is the one place that encodes it.
  const simHref = $derived(
    `/sim${simSearch(
      withSimState(
        defaultSimState(),
        store.sourceId !== null
          ? { source: 'build', ref: store.sourceId }
          : characterBuildCode !== ''
            ? { code: characterBuildCode }
            : estimateCode !== ''
              ? { code: estimateCode }
              : liveCode !== ''
                ? { code: liveCode }
                : {},
      ),
    )}`,
  );

  let status = $state<'loading' | 'ready' | 'failed'>('loading');
  let attempt = $state(0);
  // Reset asks in the toolbar rather than through window.confirm: a browser dialog cannot be
  // styled, cannot say what it is about to clear, and reads badly on phone.
  let confirmingReset = $state(false);
  // Which tree tab the phone shows. Desktop ignores it and lays every tree out side by
  // side.
  let activeTree = $state(0);

  /** `BandCompare`'s own diff, once a band has loaded (spec §4.D/§4.E.4) -- shared with
   *  every tree through `TalentTrees`, so the per-cell marker is one computation, not a
   *  second one forked by the tree row. */
  let bandDiff = $state<BandDiffView | null>(null);
  /** The level card draws nothing (a build that matches it). */
  let bandQuiet = $state(false);
  /** Share and Import leave the rail for the full-width row when the level card shows beside a spent
   *  build: with all three, the rail ran 160 to 345 px past the trees (h-audit 2026-11-09). */
  const importBelow = $derived(standalone && !shared && !bandQuiet && store.spent > 0);
  /** Import's code and outcome live here: the import itself spends points, which moves the box
   *  (importBelow) and remounts it, and its notes must survive the move. */
  let importCode = $state('');
  let importOutcome = $state<ImportOutcome | null>(null);

  // Nothing cancels a request that is already in flight, so switching class twice in quick
  // succession leaves two runs of this racing to write to the same store. The class this run
  // was started for is its generation token: a run whose class has since moved on stops before
  // it writes anything, rather than leaving the store holding one class's trees under
  // another's slug -- which `toDraft` would then post as that class's talent ids under the
  // wrong class_id, for the API to refuse. SharePanel.svelte draws the same line for saves,
  // by comparing drafts.
  async function load(): Promise<void> {
    const slug = store.classSlug;
    const stale = (): boolean => store.classSlug !== slug;

    // A decode failure is not tied to any class -- the code will never decode differently no
    // matter which class loads or how many times -- so it is shown, and `codeApplied` armed,
    // the first chance `load()` gets, independent of `slug`. A well-formed code only ever gets
    // a turn on the run whose class matches the one it names; see the comment on `codeApplied`
    // above for why that alone (with no early, unconditional arming) is already enough to stop
    // a class switch from ever applying it to the wrong class.
    if (decoded !== null && !decoded.ok && !codeApplied) {
      codeApplied = true;
      codeNote = noteForMessage(decoded.message);
    }
    const codeForThisClass =
      decoded !== null && decoded.ok && !codeApplied && decoded.build.classSlug === slug ? decoded : null;

    status = 'loading';
    try {
      // Fired together, not one after the other: `loadTalents` depends on nothing
      // `loadReference` produces (both only need `store.treeVersion`/`slug`, already known),
      // so starting it here instead of after `reference` resolves turns two sequential round
      // trips into one. Planner rebuild, fix round 1: this is also what the header's own
      // background art (once the header art returns) waits on, so halving this wait is half of the
      // Lighthouse LCP budget back. `.catch(() => {})` on a throwaway reference only silences the
      // "unhandled rejection" warning a promise nobody has awaited yet can raise while
      // `reference` is still in flight below -- the real rejection still reaches the
      // `await talentsPromise` inside the nested `try`, unaffected by this one.
      const talentsPromise = loadTalents(store.treeVersion, slug);
      talentsPromise.catch(() => {});

      const reference = await loadReference(store.treeVersion);
      if (stale()) return;
      store.setReference(reference);
      // Each step of the boot below ends the task it ran in (lib/yield.ts): the reference
      // landing, the trees landing and the page rendering them are each their own task, so none
      // of them blocks the main thread past the page's blocking-time budget under throttling.
      await yieldToMain();
      if (stale()) return;

      const ladderOrderFor = await ladderReady;
      const guideLadderFor = await guideLadderReady;
      let talents: TalentFile;
      try {
        talents = await talentsPromise;
      } catch (error) {
        // A class this planner has no talent data for will never load no matter how many
        // times this is retried, so that -- and only that -- earns the code its one turn even
        // though nothing was applied. Any other failure here (a network blip, a 5xx) leaves
        // `codeApplied` unarmed, so a same-class Retry still gets a real chance to apply it.
        // A person who followed a build link deserves to be told the class is why: the
        // generic panel below says nothing about the code, and its Retry button cannot
        // succeed without also changing class.
        if (codeForThisClass !== null && error instanceof DataLoadError && error.status === 404) {
          codeApplied = true;
          if (!stale()) {
            codeNote = { kind: 'message', text: plannerCopy.noClassNote(codeForThisClass.build.classSlug) };
          }
        }
        throw error;
      }
      if (stale()) return;
      store.setTalents(talents);
      await yieldToMain();
      if (stale()) return;

      if (codeForThisClass !== null && store.talentIndex !== null) {
        codeApplied = true;

        // The code's own stamped build may not be the one whose trees `store.talentIndex`
        // just loaded (the active build). decodeFS1 reads a tree's digits purely by tab
        // position, which only lands on the right talent when the two builds' trees match
        // tab for tab -- paladin and shaman do not, between every build this site has ever
        // shipped. Remap by stable talent id first, using the code's OWN build's talent
        // file (loaded the same way `loadTalents` loads the active one, above). When that
        // older build's file is not available on this site (pruned, or never synced),
        // fall back to reading the digits positionally, same as before this remap
        // existed, and say so in the note below.
        let treeRanksForOrder = codeForThisClass.build.treeRanks;
        let remapDropped = 0;
        let olderBuild = false;
        if (codeForThisClass.build.dataBuild !== store.treeVersion) {
          try {
            const codeBuildTalents = await loadTalents(codeForThisClass.build.dataBuild, slug);
            if (stale()) return;
            const remap = remapTreeRanksByTalentId(
              indexTalents(codeBuildTalents),
              store.talentIndex,
              codeForThisClass.build.treeRanks,
            );
            treeRanksForOrder = remap.treeRanks;
            remapDropped = remap.dropped.length;
          } catch {
            olderBuild = true;
          }
        }

        // A guide's link (`?code=&spec=`) spends the build in the guide's own levelling order.
        const guideOrder =
          specParam === null ? undefined : guideLadderFor(store.talentIndex, treeRanksForOrder, specParam);
        const rebuilt =
          guideOrder === undefined
            ? orderFromRanks(store.talentIndex, treeRanksForOrder)
            : { order: guideOrder, dropped: [] };
        store.applyOrder(rebuilt.order, codeForThisClass.build.gear);
        // A build link from the deaths recap (src/lib/report/planner-link.ts) carries gear
        // alone when the log's talents field could not be read as ranks -- every tree comes
        // through as all zeros. The note says so rather than claiming talents loaded.
        const gearOnly = codeForThisClass.build.treeRanks.every((tree) => tree.every((rank) => rank === 0));
        // Bug B1: a code that names no character is a link (a guide's, a shared one), not a
        // character's export; a restored pointer is the visitor's own paste or character.
        const linked =
          codeForThisClass.build.character === undefined && !plannerLoad.restored && !gearOnly && !olderBuild;
        codeNote =
          linked && guideOrder !== undefined && specParam !== null
            ? { kind: 'guide', specName: specDisplayName(specParam) }
            : linked && rebuilt.dropped.length + remapDropped === 0
              ? { kind: 'linked' }
              : {
                  kind: 'reconstructed',
                  dropped: rebuilt.dropped.length + remapDropped,
                  gearOnly,
                  olderBuild,
                };
      } else if (
        talentsParam !== null &&
        !talentsApplied &&
        store.talentIndex !== null &&
        store.order.length === 0
      ) {
        // Rebuild spec §4.D/§4.I: the BiS page's "Talents in planner" link (and a guide's
        // "Open this build in the planner") carries a band's own talents string directly.
        // `?code=` always wins when both are present (unchanged precedence, checked above);
        // this only ever gets a turn on a build that is still empty, same "never silently
        // overwrite a build already in progress" rule `?code=` already follows.
        talentsApplied = true;
        const ordered =
          specParam === null ? undefined : ladderOrderFor(store.talentIndex, talentsParam, specParam);
        const rebuilt = loadFromBand(store.talentIndex, talentsParam, ordered);
        store.applyOrder(rebuilt.order, {});
        const bandLabel = bandRangeLabel(bandForLevel(levelReached(rebuilt.order)));
        store.nameBandBuild(bandTitleForDraft(bandLabel, store.classSlug, store.split));
        codeNote = { kind: 'band', bandLabel };
      }

      // Task 10: a `?code=`/restored code just applied above. A shared build (`record`) is looked
      // at, not chosen, like any link: it leaves the header's character where it is, so In the game
      // sends to that character (premium spec 2026-11-08 section 8, one way to switch).
      if (record === null && codeForThisClass !== null && !plannerLoad.restored) {
        const named = codeForThisClass.build.character;
        if (named === undefined) {
          // A link that names nobody is somebody's build, not the current character's.
          buildOrigin = { owner: '' };
        } else {
          // A link is looked at, not chosen: it moves the pointer only to an account character it
          // names, and never to a nameless code or somebody else's export.
          void adoptHeldCharacter(named.name, codeForThisClass.build.classSlug).then((held) => {
            if (store.classSlug !== codeForThisClass.build.classSlug) return;
            if (held !== null) loadedCharacter = held;
            else buildOrigin = { owner: named.name };
          });
        }
      }

      // Planner rebuild §11 (owner ruling): the gear panel left the planner entirely --
      // GearPanel/ItemPicker/BisSlotPopover and their own sets/items/weights fetch are gone
      // with it (the simulator keeps its own gear step). `status` flips the moment the
      // trees have something to show; nothing else is fetched after it.
      status = 'ready';
      activeTree = openTreeIndex(editorTreeList);
    } catch {
      // A stale run's failure is not this class's failure: the run that replaced it owns the
      // status, and reporting this one would put a working planner behind a Retry button.
      if (stale()) return;
      status = 'failed';
    }
  }

  // Re-runs whenever the class changes (selectClass drops the loaded trees) or Retry bumps
  // `attempt`. Reading both synchronously here is what registers them as dependencies; the
  // writes `load` performs happen after the tracking window, so this cannot loop.
  $effect(() => {
    void store.classSlug;
    void attempt;
    void load();
  });

  // Switching class empties the build on its own, and the class selector stays reachable while
  // the confirm is open. Without this the planner comes back asking whether to clear a build the
  // switch already cleared. The phone tab goes back to the first tree for the same reason: the
  // open tree belonged to the class that just left, and a class with fewer trees than the last
  // one would leave no tab selected and no panel shown at all. Its own effect rather than a line
  // in the one above: that effect documents a careful no-loop invariant, and this has nothing to
  // do with loading.
  // Compared against the class this effect last saw, rather than cleared unconditionally, so
  // that the very first run -- which fires once on mount, in the same synchronous flush as the
  // load effect's own first run -- cannot wipe out a decode failure that load() may already
  // have written into `codeNote` moments earlier in that same flush (decode failures are
  // reported synchronously, before load()'s first `await`). Every run after the first is a
  // genuine class change, and only those should ever clear it.
  // The address mirrors the class and race on screen (lib/planner/url.ts), so a refresh or a
  // copied link lands on the build the visitor chose rather than the one the page opened on.
  // Only the standalone /planner page owns its address: /b/:id names a record, and Top Gear's
  // inline planner lives on another page's URL. replaceState, never pushState: a class switch
  // is not a navigation the back button should retrace.
  const decodedCode = decoded !== null && decoded.ok ? decoded.build : null;
  // The class and race the page came up with, read once the reference data has landed:
  // `setReference` moves an unset or illegal race to the class's first legal one, and that
  // move is the page opening, not the visitor choosing, so it must not rewrite a bare address.
  let opened: { classSlug: string; raceSlug: string } | null = null;
  $effect(() => {
    if (!standalone || record || store.races.length === 0) return;
    // An account character owns the address: a refresh reopens it from the account.
    if (opened === null) {
      opened = { classSlug: store.classSlug, raceSlug: store.raceSlug };
      return;
    }
    if (loadedCharacter !== null || characterLoading) return;
    const next = plannerSearchFor(
      window.location.search,
      store.classSlug,
      store.raceSlug,
      decodedCode === null ? null : { classSlug: decodedCode.classSlug, raceSlug: decodedCode.raceSlug },
      opened,
    );
    if (next !== null) window.history.replaceState(null, '', `${window.location.pathname}${next}`);
  });

  /** The trees as the editor draws them (game colours, the cap, planned points, the band dots). */
  const editorTreeList = $derived(
    store.talentIndex === null
      ? []
      : editorTrees({
          index: store.talentIndex,
          order: store.order,
          classSlug: store.classSlug,
          characterLevel: store.characterLevel,
          readOnly: store.readOnly,
          bandDiff,
          build: store.treeVersion,
        }),
  );
  const EDITOR_COPY: TalentTreesCopy = {
    resetTree: plannerCopy.resetTree,
    resetTreeLabel: plannerCopy.resetTreeLabel,
  };
  const TREES_LABEL = 'Talent trees';
  /** The reserve's tier count (spec 5.2): the loaded trees', or the seven every class has. */
  const reserveTiers = GAME_TIERS;

  /** False in the server render and the hydration pass; the crest re-renders once true. */
  let mounted = $state(false);
  onMount(() => {
    mounted = true;
  });
  const className = $derived(store.classRow?.name ?? classDisplayName(store.classSlug));
  /** True once one tree holds strictly more points than each other (`specKeyFor`'s tie-break). */
  const leadingSpec = $derived.by(() => {
    const split = store.split;
    const max = split.length === 0 ? 0 : Math.max(...split);
    if (max === 0 || split.filter((points) => points === max).length !== 1 || store.talentIndex === null)
      return '';
    return specLabel(specOf(store.talentIndex, [...store.order]));
  });
  /** The meta line: race, level, and the spec once a tree leads (spec 4.1). */
  const metaLine = $derived(
    [store.raceRow?.name ?? '', `Level ${store.level}`, leadingSpec]
      .filter((part) => part !== '')
      .join(' · '),
  );
  /** The H1: `{Class} talents`; on a shared build its own title, else `{Race} {Class}` (B1). */
  const headTitle = $derived(
    !shared
      ? plannerHeaderCopy.h1(className)
      : title !== null && title.trim() !== ''
        ? title
        : [store.raceRow?.name ?? '', className].filter((part) => part !== '').join(' '),
  );
  /** Whose build this is when it is not the player's own; said once (ruling MUST 1): never beside a note. */
  const originLine = $derived(
    effectiveOrigin === null || codeNote !== null
      ? ''
      : plannerCharacterCardCopy.originLine(effectiveOrigin.owner, leadingSpec || className),
  );

  /** The one note line (spec 4.3), its phone sentence when it has one. A guide's link and a clean
   *  character load say nothing (h-audit 2026-11-09): the guide row's Loaded tag and the header's
   *  character already say where the build came from. */
  const noteText = $derived.by((): { text: string; short?: string } | null => {
    const note = codeNote;
    if (note === null) return null;
    switch (note.kind) {
      case 'linked':
        return { text: plannerCopy.linkedNote, short: plannerCopy.linkedNoteShort };
      case 'guide':
        return null;
      case 'tree-count':
        return { text: `That code has ${note.got} talent trees; a build has ${note.want}.` };
      case 'kept':
        return { text: plannerCopy.keptNote(note.kept, note.total, note.name) };
      case 'band':
        return { text: plannerCopy.bandLoadedNote(note.bandLabel) };
      case 'message':
        return { text: note.text };
      case 'reconstructed': {
        const head = note.gearOnly
          ? plannerCopy.noTalentRanksNote
          : note.dropped === 0
            ? ''
            : plannerCopy.reconstructedDroppedNote(note.dropped);
        const text = note.olderBuild ? `${head} ${plannerCopy.olderBuildNote}`.trim() : head;
        return text === '' ? null : { text };
      }
    }
  });

  /** The bar (spec 4.2): the nine class crests, then the class's legal races. */
  const barGroups = $derived<ControlGroup[]>([
    {
      label: 'Class',
      param: 'class',
      kind: 'segmented',
      compact: true,
      value: store.classSlug,
      options: classRows.map((row) => ({
        value: row.slug,
        label: row.name,
        crest: row.slug,
        disabled: store.readOnly,
        disabledReason: READ_ONLY_REASON,
      })),
    },
    {
      label: 'Race',
      param: 'race',
      kind: 'select',
      value: store.raceSlug,
      fieldWidth: 130,
      options: store.legalRaces.map((race) => ({
        value: race.slug,
        label: race.name,
        disabled: store.readOnly,
      })),
    },
  ]);
  function chooseFromBar(param: string, value: string): void {
    if (store.readOnly) return;
    if (param === 'class') store.selectClass(value);
    else if (param === 'race') store.selectRace(value);
  }

  let classSlugForNoteReset = store.classSlug;
  $effect(() => {
    void store.classSlug;
    confirmingReset = false;
    activeTree = 0;
    if (store.classSlug !== classSlugForNoteReset) {
      // A code's note describes why the build looked the way it did on the class it was shown
      // under -- switching class already discards that build (selectClass), so a note left
      // behind (most visibly the "does not have" message: switching class is its entire
      // remedy) would sit under an unrelated, working build claiming something no longer true.
      // Every kind of note is cleared the same way; none of them describes anything about a
      // class the planner has since moved on from.
      codeNote = null;
      buildOrigin = null;
      classSlugForNoteReset = store.classSlug;
    }
  });
</script>

{#snippet foot()}
  <!-- The status line and the source share one line: a refusal stands over the source while it
       shows, so no empty line is held under the trees. -->
  <div class="pl-foot-l">
    <p
      class="st2"
      class:wipe={store.refusal !== null}
      role="status"
      aria-live="polite"
      data-testid="planner-status"
    >
      {store.refusal ?? ''}
    </p>
    <span
      class="text-small text-muted"
      class:pl-veiled={store.refusal !== null}
      data-testid="planner-tree-source">{treeSourceNotice(store.treeVersion)}</span
    >
  </div>
  <span class="rs">
    {#if store.readOnly}
      <!-- A shared build's Edit a copy sits in the notes line, not under the trees. -->
    {:else if confirmingReset}
      <span class="text-small text-muted">{plannerCopy.resetAsk}</span>
      <Button
        size="sm"
        onclick={() => {
          store.reset();
          confirmingReset = false;
        }}
        testid="planner-reset-confirm">{plannerCopy.resetYes}</Button
      >
      <!-- Focus lands on the safe answer: a second Enter out of habit keeps the build. -->
      <span class="contents" {@attach (node) => node.querySelector('button')?.focus()}>
        <Button size="sm" onclick={() => (confirmingReset = false)} testid="planner-reset-keep"
          >{plannerCopy.resetNo}</Button
        >
      </span>
    {:else if store.spent > 0}
      <button
        type="button"
        class="{LINK} text-[14px] leading-5 max-md:-my-3 max-md:py-3 [@media(pointer:coarse)]:-my-3 [@media(pointer:coarse)]:py-3"
        onclick={() => (confirmingReset = true)}
        data-testid="planner-reset">{plannerCopy.reset}</button
      >
    {/if}
  </span>
{/snippet}

{#snippet treesReserve()}
  <div class="pl-reserve" style:--tiers={reserveTiers}>
    {#if status === 'failed'}
      <div class="pl-err" role="alert" data-testid="planner-load-error">
        <p class="text-strong">{DATA_LOAD_FAILED}</p>
        <p class="text-small text-muted">{plannerCopy.loadFailedDetail(store.treeVersion)}</p>
        <div class="pl-acts">
          <Button onclick={() => (attempt += 1)} label={plannerCopy.reloadTalents} testid="planner-retry"
            >{plannerCopy.retry}</Button
          >
        </div>
      </div>
    {:else}
      <div
        class="pl-skel"
        role="status"
        aria-busy="true"
        aria-label="Loading talent data"
        data-testid="planner-talent-skeleton"
      >
        {#each [0, 1, 2] as tree (tree)}
          <div class="pl-skt">
            {#if openingTrees[tree] !== undefined && store.classSlug === classSlug}
              <!-- The opening class's tree head as the editor draws it (TalentTree .tt-h). -->
              {@const iconSrc = treeIconSrc(classSlug, tree, store.treeVersion)}
              <div class="pl-skh">
                <SpellIcon url={iconSrc === '' ? null : iconSrc} size={28} />
                <h2>{openingTrees[tree]}</h2>
                <b>0</b>
              </div>
            {:else}
              <span class="skeleton-block block" style="width:180px;height:24px"></span>
            {/if}
            {#if openingArt[tree] !== undefined && store.classSlug === classSlug}
              <!-- The opening class's tree art in the server render: the page's largest paint
                       arrives with the first paint instead of after the island and the trees. -->
              <img
                class="pl-art"
                src={dataUrl(store.treeVersion, openingArt[tree])}
                alt=""
                width="216"
                height="384"
                loading="eager"
                fetchpriority={tree === 0 ? 'high' : 'auto'}
              />
            {:else}
              <span class="skeleton-block block max-w-full" style="width:216px;height:384px"></span>
            {/if}
          </div>
        {/each}
      </div>
    {/if}
  </div>
{/snippet}

{#snippet plate()}
  {#if status === 'ready' && store.talentIndex}
    <!-- Never the editor's readOnly: a shared build's press reaches the store, which refuses and
         says why in the status line (fork it to edit). -->
    <TalentTrees
      mode="edit"
      trees={editorTreeList}
      build={store.treeVersion}
      bind:active={activeTree}
      copy={EDITOR_COPY}
      onadd={(id) => store.addPoint(id)}
      onremove={(id) => store.removePoint(id)}
      onresettree={store.readOnly ? undefined : (i) => store.resetTree(i)}
      {foot}
      title={standalone ? plannerCopy.talentsTitle : undefined}
      frame="framed"
      testid="planner-trees"
    />
  {:else if standalone}
    <!-- The ready trees' panel, its heading included, so the trees arriving move nothing. -->
    <Section title={plannerCopy.talentsTitle} panel frame="framed" flush testid="planner-trees-loading">
      {@render treesReserve()}
    </Section>
  {:else}
    <Plate label={TREES_LABEL}>{@render treesReserve()}</Plate>
  {/if}
{/snippet}

{#snippet notes()}
  {#if shared && store.readOnly}
    <!-- Decision B2: Edit a copy sits in the first screen beside the shared-link line; pressing it
         makes the planner editable in place (store.fork), and the line goes. -->
    <div class="pl-notes">
      <div data-region={REGION.notes}>
        <!-- The Notice line's type and rhythm with a button as its second item (ui/Notice takes a
             link only). -->
        <p class="pl-shared-note" data-testid="planner-shared-note">
          <span>{plannerCopy.sharedLinkNote}</span>
          <Button onclick={() => store.fork()} testid="planner-fork-button">{plannerCopy.fork}</Button>
        </p>
      </div>
    </div>
  {:else if noteText !== null}
    <div class="pl-notes">
      <div data-region={REGION.notes}>
        <Notice text={noteText.text} short={noteText.short} tone="neutral" testid="planner-code-note" />
      </div>
    </div>
  {/if}
{/snippet}

{#snippet bandSection()}
  <BandCompare
    {store}
    panel={standalone}
    heading={shared ? bandCompareCopy.sharedHeading : undefined}
    quietMatch={standalone && !shared}
    characterLevel={store.characterLevel}
    specOverride={specParam ?? undefined}
    onbanddiff={(diff) => (bandDiff = diff)}
    onquiet={(quiet) => (bandQuiet = quiet)}
    onloaded={(bandLabel) => (codeNote = { kind: 'band', bandLabel })}
    onband={(band, label) => (estimateBand = { band, label })}
  />
{/snippet}

<!-- The take-it-with-you parts under In the game: Share (once a point is spent) and Import, unless
     they sit in the full-width row below. -->
{#snippet takeItParts()}
  {#if store.spent > 0 && !importBelow}<SharePanel {store} part />{/if}
  {#if !importBelow}{@render importSection(true)}{/if}
{/snippet}

{#snippet importSection(part: boolean)}
  {#if store.talentIndex}
    <ImportBox
      bind:code={importCode}
      bind:outcome={importOutcome}
      panel={standalone}
      {part}
      talents={store.talentIndex}
      activeBuild={activeBuild.build}
      onimport={(build, pastedCode) => {
        store.loadImported(build);
        loadedCharacter = null;
        buildOrigin = null;
        void recordPastedExport(pastedCode);
      }}
      onwrongclass={standalone ? switchClassForImport : undefined}
    />
  {/if}
{/snippet}

{#snippet railSkeleton(title: string)}
  <Section {title} panel>
    <div class="pl-sec" role="status" aria-busy="true">
      <span class="sr-only">{uiCopy.loading}</span>
      <span class="skeleton-block block max-w-full" style="width:260px;height:20px" aria-hidden="true"></span>
      <span class="skeleton-block block max-w-full" style="width:200px;height:44px" aria-hidden="true"></span>
    </div>
  </Section>
{/snippet}

<!-- Point order's place while the talent data loads: on a phone it sits between the trees and the
     rail, so without it the level build and everything under it jumped down on load. -->
{#snippet orderSkeleton()}
  <Section title={plannerCopy.pointOrderTitle} panel>
    <p class="pl-pad" role="status" aria-busy="true">
      <span class="sr-only">{uiCopy.loading}</span>
      <span class="skeleton-block block max-w-full" style="width:160px;height:20px" aria-hidden="true"></span>
    </p>
  </Section>
{/snippet}

{#if standalone}
  <Page layout="workbench" panels stickyRail testid="planner">
    {#snippet header()}
      <PageHeader
        title={headTitle}
        tone={classTone(store.classSlug)}
        lead={shared ? undefined : plannerCopy.lead}
        testid="planner-head"
      >
        {#snippet crest()}
          <!-- Keyed on mount: hydration keeps the server's src, which is the default class's. -->
          {#key mounted}<Crest cls={className} size={56} testid="planner-header-crest" />{/key}
        {/snippet}
        {#snippet meta()}
          {#if shared}<span data-testid="planner-shared-meta">{plannerCopy.sharedMeta}</span>{/if}
          <span data-testid="planner-meta">{metaLine}</span>
          {#if shared}
            <span class="numeric" data-testid="planner-split">{store.split.join('/')}</span>
            {#if views !== null}<span data-testid="planner-views">{plannerCopy.sharedViews(views)}</span>{/if}
          {/if}
          {#if originLine !== ''}<span class="text-gold" data-testid="planner-build-origin">{originLine}</span
            >{/if}
        {/snippet}
        {#snippet answer()}
          {#if status === 'loading'}
            <div class="pl-ans-skel" role="status" aria-busy="true">
              <span class="sr-only">{uiCopy.loading}</span>
              <span class="skeleton-block block" style="width:150px;height:68px" aria-hidden="true"></span>
            </div>
          {:else}
            <PlannerAnswer {store} simHref={simLinkShown ? simHref : undefined} />
          {/if}
        {/snippet}
      </PageHeader>
    {/snippet}
    {#snippet bar()}
      {#if !hideBar}
        <ControlBar
          groups={barGroups}
          label={plannerCopy.barLabel}
          onchange={chooseFromBar}
          testid="planner-bar"
          id="planner-bar"
        ></ControlBar>
      {/if}
      {@render notes()}
    {/snippet}
    {#snippet rail()}
      {#if status === 'ready'}
        <!-- A shared build leads with In the game (spec section 8); the planner with its level build. -->
        <Ordered order={shared ? 3 : 2}>{@render bandSection()}</Ordered>
        <!-- One take-it-with-you plate (h-audit 2026-11-09): In the game, then Share (once a point is
             spent) and Import under hairlines; last in the phone flow, after the guide builds. -->
        <!-- A shared build (read-only) has no Share or Import parts, so InGame keeps its own addon line. -->
        <Ordered order={shared ? 2 : 5}
          ><InGame
            {store}
            character={sendCharacter}
            origin={effectiveOrigin}
            plainOtherClass={shared}
            children={store.readOnly ? undefined : takeItParts}
          /></Ordered
        >
      {:else if status === 'loading'}
        {#if shared}
          <Ordered order={2}>{@render railSkeleton(plannerCopy.inGameTitle)}</Ordered>
          <Ordered order={3}>{@render railSkeleton(bandCompareCopy.sharedHeading)}</Ordered>
        {:else}
          <Ordered order={2}>{@render railSkeleton(bandCompareCopy.heading('60'))}</Ordered>
          <Ordered order={3}>{@render railSkeleton(plannerCopy.inGameTitle)}</Ordered>
        {/if}
      {/if}
    {/snippet}
    <!-- The trees beside the rail, then Point order and Build vs guide full width beneath both (h-audit
         2026-11-09: neither column leaves a void); on a phone Point order follows the trees. -->
    {#snippet below()}
      {#if status === 'ready'}
        <Ordered order={1} span
          ><PlannerOrder {store} characterName={sendCharacter?.name ?? null} panel /></Ordered
        >
        <!-- Only a class with a guide build takes a slot: an empty one would still hold a flow gap. -->
        {#if guideBuilds.some((build) => build.classSlug === store.classSlug)}
          <Ordered order={4} span><GuideBuilds builds={guideBuilds} {store} /></Ordered>
        {/if}
        {#if importBelow && !store.readOnly}
          <Ordered order={7}><SharePanel {store} /></Ordered>
          <Ordered order={8}>{@render importSection(false)}</Ordered>
        {/if}
      {:else if status === 'loading'}
        <Ordered order={1} span>{@render orderSkeleton()}</Ordered>
      {/if}
    {/snippet}
    <Ordered order={1} span>{@render plate()}</Ordered>
  </Page>
{:else}
  <!-- Top Gear's inline "add a build" (TalentCandidates.svelte): the facts bar, the trees, the
       order, the band and the paste box, in one column inside that page's own layout. -->
  <div class="flex flex-col gap-8" data-testid="planner">
    <SummaryBar {store} {live} {simHref} {gate} {standalone} onshowdps={() => (dpsOptedIn = true)} />
    {@render plate()}
    {#if status === 'ready'}
      <PlannerOrder {store} />
      {@render bandSection()}
      {@render importSection(false)}
    {/if}
  </div>
{/if}

<style>
  .pl-reserve {
    /* The ready trees' height (spec 5.2): seven tiers of cells (GAME_TIERS) plus the tree head,
       the grid's padding and the foot, measured per breakpoint, so loading, failed and ready
       stand the same height and nothing under the Plate moves when the trees arrive. */
    --cell: 44px;
    --cgap: 8px;
    --chrome: 157px;
    min-height: calc(var(--tiers) * (var(--cell) + var(--cgap)) + var(--chrome));
  }
  @media (min-width: 768px) {
    .pl-reserve {
      --cell: 48px;
      --chrome: 123px;
    }
  }
  @media (min-width: 1280px) {
    .pl-reserve {
      --chrome: 127px;
    }
  }
  @media (min-width: 1760px) {
    .pl-reserve {
      --cell: 56px;
      --cgap: 10px;
      --chrome: 125px;
    }
  }
  .pl-skel {
    display: grid;
    grid-template-columns: repeat(3, minmax(0, 1fr));
  }
  /* TalentTrees' .tt-col padding, so a head sits where the editor's head lands. */
  .pl-skt {
    display: flex;
    flex-direction: column;
    align-items: center;
    gap: 14px;
    padding: 16px 12px;
    min-width: 0;
  }
  @media (max-width: 767px) {
    .pl-skel {
      grid-template-columns: minmax(0, 1fr);
    }
    .pl-skt + .pl-skt {
      display: none;
    }
  }
  .pl-skh {
    align-self: stretch;
    display: flex;
    align-items: center;
    gap: 10px;
    min-width: 0;
    height: 28px;
  }
  .pl-skh h2 {
    font: 700 15px/20px var(--font-body);
    color: var(--color-strong);
    margin: 0;
    min-width: 0;
  }
  .pl-skh b {
    margin-left: auto;
    font: 700 18px/24px var(--font-body);
    font-variant-numeric: tabular-nums;
    color: var(--color-muted);
  }
  .pl-art {
    display: block;
    max-width: 100%;
    height: 384px;
    object-fit: cover;
    border-radius: 4px;
    opacity: 0.5;
  }
  .pl-err {
    display: flex;
    flex-direction: column;
    gap: 8px;
    padding: 20px;
  }
  /* The ready answer's height (spec 5.2): the trees arriving move nothing under the head. */
  .pl-ans-skel {
    display: flex;
    align-items: flex-start;
  }
  @media (min-width: 900px) {
    .pl-ans-skel {
      align-items: flex-end;
    }
  }
  .pl-notes {
    padding-inline: var(--page-inset);
    padding-block: 16px 0;
    max-width: calc(var(--measure-reading) + 2 * var(--page-inset));
  }
  .pl-foot-l {
    flex: 1 1 auto;
    display: grid;
    min-width: 0;
  }
  .pl-foot-l > :global(*) {
    grid-area: 1 / 1;
    align-self: center;
  }
  .pl-veiled {
    visibility: hidden;
  }
  .pl-shared-note {
    display: flex;
    flex-wrap: wrap;
    align-items: baseline;
    gap: 6px 14px;
    padding: 14px 0 4px;
    font-size: 15px;
    line-height: 24px;
    color: var(--color-text);
  }
  .pl-notes :global(.ui-notice) {
    padding-top: 4px;
  }
</style>
