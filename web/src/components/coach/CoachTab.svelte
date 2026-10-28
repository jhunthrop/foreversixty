<!-- web/src/components/coach/CoachTab.svelte -->
<!-- The Coach tab's container: reads who is looking, picks the one player the review is for,
     asks the coach API (and keeps asking while the coach writes), and hands the settled state to
     CoachPanel. Several owned characters get a row of crest buttons that write `source`. An
     officer's pick on another player's row resolves its character key through the pull's rating
     rows. Focus moves to the panel heading once, so a strip or Discord link lands on it. -->
<script lang="ts">
  import activeBuild from '../../data/active-build.json';
  import { coachCopy } from '../../lib/coach/copy';
  import { EMPTY_NAMES, loadCoachNames, type CoachNames } from '../../lib/coach/names';
  import { createCoachPanelState } from '../../lib/coach/panel-state.svelte';
  import { pickPlayer, showsOfficerPick, type RosterUnit } from '../../lib/coach/player-pick';
  import { playerCardLine } from '../../lib/coach/player-card';
  import { viewOf } from '../../lib/coach/states';
  import { createCoachViewerState } from '../../lib/coach/viewer-state.svelte';
  import { splitUnitName } from '../../lib/characters';
  import { createReportRatingsFetch } from '../../lib/rating/report-ratings.svelte';
  import { REPORT_LAZY_MIN_H } from '../../lib/report/layout';
  import { classSlugOf } from '../../lib/report/planner-link';
  import type { FightEntry } from '../../lib/report/types';
  import ClassIcon from '../report/ClassIcon.svelte';
  import CoachPanel from './CoachPanel.svelte';

  let {
    reportId,
    fight,
    firstFight,
    roster,
    source,
    onSelectPlayer,
    apiBase = undefined,
  }: {
    reportId: string;
    fight: FightEntry | null;
    firstFight: number;
    roster: RosterUnit[];
    source: string;
    onSelectPlayer: (guid: string) => void;
    apiBase?: string;
  } = $props();

  // svelte-ignore state_referenced_locally
  const viewerState = createCoachViewerState(apiBase);
  // svelte-ignore state_referenced_locally
  const panel = createCoachPanelState(apiBase);
  // svelte-ignore state_referenced_locally
  const ratings = createReportRatingsFetch(apiBase);

  const viewer = $derived(viewerState.viewer);
  const pick = $derived(pickPlayer(viewer, roster, source));
  const unit = $derived(pick.unit);
  const fightIndex = $derived(fight?.index ?? null);
  const classSlug = $derived(classSlugOf(unit?.class) ?? '');

  const ratedKey = $derived(
    unit === null
      ? null
      : (ratings.data?.players.find(
          (player) => splitUnitName(player.player_name).name === splitUnitName(unit.name).name,
        )?.player_key ?? null),
  );
  const playerKey = $derived(pick.ownKey ?? ratedKey);
  const unresolved = $derived(
    unit !== null && pick.ownKey === null && ['failed', 'none', 'flagged'].includes(ratings.status),
  );

  // An officer's pick needs the pull's rating rows to name the character key.
  $effect(() => {
    if (unit !== null && pick.ownKey === null && fightIndex !== null) ratings.load(reportId, fightIndex);
  });
  $effect(() => {
    if (playerKey !== null && fightIndex !== null) panel.load(reportId, fightIndex, playerKey);
  });
  $effect(() => panel.dispose);

  let names = $state<CoachNames>(EMPTY_NAMES);
  $effect(() => {
    const slug = classSlug;
    let current = true;
    if (slug === '') return;
    void loadCoachNames(activeBuild.build, slug).then((loaded) => {
      if (current) names = loaded;
    });
    return () => {
      current = false;
    };
  });

  let focused = false;
  $effect(() => {
    if (focused || panel.view.state === 'loading') return;
    // The panel's load can resolve after a test's DOM is torn down; there is nothing to focus then.
    if (typeof document === 'undefined') return;
    focused = true;
    const heading = document.getElementById('coach-panel-heading');
    heading?.focus({ preventScroll: true });
    // A phone scrolls the tab under its sticky bar (spec 1); a desktop already has it in view,
    // and scrolling it to the top would carry the sticky fight list down with it.
    const desktop = window.matchMedia('(min-width: 768px)').matches;
    heading?.scrollIntoView({ block: desktop ? 'nearest' : 'start' });
  });

  const view = $derived.by(() => {
    if (viewerState.resolved && !viewer.signedIn) return viewOf('signed_out');
    if (unresolved || (unit === null && viewerState.resolved)) return viewOf('not_available');
    return panel.view;
  });
  const signInHref = $derived(
    typeof window === 'undefined'
      ? '/login'
      : `/login?next=${encodeURIComponent(`${window.location.pathname}${window.location.search}`)}`,
  );
  const historyHref = $derived(
    viewer.premium && pick.ownKey !== null ? `/character/${pick.ownKey}#coach` : null,
  );
</script>

<!-- The tab's reserve (640 px desktop, 420 px phone) lives here, not on the card: a one-line
     state is a natural-height card over bare background. -->
<div class="flex flex-col gap-4 {REPORT_LAZY_MIN_H.coach}" data-testid="coach-tab">
  {#if pick.owned.length > 1}
    <ul class="flex gap-1" aria-label="Characters in this pull" data-testid="coach-players">
      {#each pick.owned as owned (owned.guid)}
        <li>
          <button
            type="button"
            class="focus-visible:outline-gold flex h-11 w-11 items-center justify-center rounded-full focus-visible:outline-2"
            style={owned.guid === unit?.guid ? 'box-shadow: 0 0 0 2px var(--color-gold)' : ''}
            aria-label={splitUnitName(owned.name).name}
            aria-pressed={owned.guid === unit?.guid}
            title={splitUnitName(owned.name).name}
            onclick={() => onSelectPlayer(owned.guid)}
          >
            <ClassIcon className={owned.class} size={32} />
          </button>
        </li>
      {/each}
    </ul>
  {/if}

  {#if viewerState.resolved && showsOfficerPick(viewer, pick)}
    <p class="text-muted text-[14px]" data-testid="coach-pick">{coachCopy.officerPick}</p>
  {:else}
    <CoachPanel
      {view}
      answer={panel.data}
      late={panel.late}
      reserve={panel.reserve}
      {names}
      build={activeBuild.build}
      {classSlug}
      spec={unit?.spec ?? ''}
      playerName={unit?.name ?? ''}
      playerGuid={unit?.guid ?? ''}
      {firstFight}
      className={unit?.class}
      cardLine={playerCardLine(
        panel.data !== null && panel.data.access !== 'teaser' ? panel.data.spec : '',
        unit?.class,
        fight,
      )}
      {historyHref}
      {signInHref}
      onRetry={() => panel.retry()}
    />
  {/if}
</div>
