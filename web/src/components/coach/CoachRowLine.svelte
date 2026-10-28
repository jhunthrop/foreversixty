<!-- web/src/components/coach/CoachRowLine.svelte -->
<!-- The coach line under a report row (spec 3.2): mounted by ActorRow as a sibling after the row
     button. It fetches only for the viewer's own characters on a reviewable pull, in the
     read-only `line=1` form that never queues a generation (Ruling 6), and shows the Premium or
     the free line. For an officer on another player's opened row it offers the on-demand
     `Coach review` button, which opens the Coach tab; nothing is fetched until the panel opens. -->
<script lang="ts">
  import activeBuild from '../../data/active-build.json';
  import { fetchFightCoach } from '../../lib/coach/api';
  import { coachCopy } from '../../lib/coach/copy';
  import { loadSpellIconUrls, type IconUrl } from '../../lib/coach/names';
  import { followReportLink, openReportLink, reportHref } from '../../lib/coach/navigate';
  import { useCoachRow } from '../../lib/coach/report-context';
  import { lineRequestFor, officerButtonFor, type LineRequest } from '../../lib/coach/row-fetch';
  import { rowLineFor, type RowLine, type StripPhase } from '../../lib/coach/row-line';
  import { createCoachViewerState } from '../../lib/coach/viewer-state.svelte';
  import { defaultState, withState } from '../../lib/report/url';
  import CoachRowStrip from './CoachRowStrip.svelte';

  let {
    actor,
    rowOpen = false,
    apiBase = undefined,
  }: {
    actor: { guid: string; name: string; class?: string };
    rowOpen?: boolean;
    apiBase?: string;
  } = $props();

  const readReport = useCoachRow();
  // svelte-ignore state_referenced_locally
  const viewerState = createCoachViewerState(apiBase);
  const report = $derived(readReport?.() ?? null);
  const request = $derived<LineRequest | null>(
    report === null ? null : lineRequestFor(viewerState.viewer, actor.name, report),
  );
  const officerButton = $derived(report !== null && officerButtonFor(viewerState.viewer, actor.name, report));

  let phase = $state<StripPhase>('loading');
  let line = $state<RowLine | null>(null);
  let iconUrl = $state<IconUrl>(() => null);
  let attempt = $state(0);

  // One request per (report, fight, player); a stale answer for an older key is dropped.
  $effect(() => {
    const wanted = request;
    void attempt;
    if (wanted === null) return;
    let current = true;
    phase = 'loading';
    line = null;
    void loadSpellIconUrls(activeBuild.build).then((urls) => {
      if (current) iconUrl = urls;
    });
    fetchFightCoach(wanted.reportId, wanted.fightIndex, wanted.playerKey, apiBase, { line: true })
      .then((answer) => {
        if (!current) return;
        line = rowLineFor(answer, actor.class ?? '');
        phase = 'line';
      })
      .catch(() => {
        if (current) phase = 'error';
      });
    return () => {
      current = false;
    };
  });

  const href = $derived(
    report?.fight == null
      ? '?'
      : reportHref(
          withState(defaultState(report.firstFight), {
            fight: report.fight.index,
            tab: 'coach',
            source: actor.guid,
          }),
          report.firstFight,
        ),
  );
  const label = $derived(line === null ? '' : `${line.lead}${line.title}`);
  const OFFICER_BUTTON =
    'rounded-control border-line-warm hover:border-gold focus-visible:outline-gold active:bg-card-top inline-flex h-11 items-center border px-3 text-[13px] font-semibold focus-visible:outline-2 focus-visible:outline-offset-2 md:h-9';
</script>

{#if request !== null && (phase !== 'line' || line !== null)}
  <CoachRowStrip
    {phase}
    {line}
    {href}
    {iconUrl}
    {label}
    onOpen={(event) => followReportLink(event, href)}
    onRetry={() => (attempt += 1)}
  />
{:else if officerButton && rowOpen}
  <div class="border-line-soft border-t bg-[#0b0f17] py-1 pl-12 md:pl-[100px]">
    <button
      type="button"
      class={OFFICER_BUTTON}
      onclick={() => openReportLink(href)}
      data-testid="coach-officer-open"
    >
      {coachCopy.rowOfficer}
    </button>
  </div>
{/if}
