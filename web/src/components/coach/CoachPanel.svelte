<!-- web/src/components/coach/CoachPanel.svelte -->
<!-- The Coach tab's panel (spec 3.1, 4.2 and 6): one column inside the report's content column,
     the card at its full width and the rail under it as a 2-up grid (20 px gap); two columns, a
     760 px card and a 320 px rail, only when the content column is 1120 px or wider (a container
     query, so nothing ever runs past the column). It is presentational: the container owns the
     fetch (panel-state) and hands over the view, the answer and the reserve. The header is a fixed
     44 px row with a status slot that never moves a line; a writing review is held at its reserved
     heights. The card is always its natural height: the tab's reserve lives on CoachTab. -->
<script lang="ts">
  import { coachCopy, fill } from '../../lib/coach/copy';
  import { reviewedDate } from '../../lib/coach/dates';
  import type { CoachNames } from '../../lib/coach/names';
  import type { CoachAnswer } from '../../lib/coach/panel-state.svelte';
  import type { Reserve } from '../../lib/coach/reserves';
  import type { CoachPanelView } from '../../lib/coach/states';
  import { CARD } from '../../lib/ui/card';
  import { BUSY_CLASS } from '../../lib/ui/busy';
  import { classColorVar } from '../../lib/report/format';
  import { splitUnitName } from '../../lib/characters';
  import ClassIcon from '../report/ClassIcon.svelte';
  import CoachReviewBody from './CoachReviewBody.svelte';
  import CoachTeaser from './CoachTeaser.svelte';

  let {
    view,
    answer,
    late = false,
    reserve = null,
    names,
    build,
    classSlug,
    spec,
    playerName,
    playerGuid,
    firstFight,
    className = undefined,
    cardLine = '',
    historyHref = null,
    refetching = false,
    signInHref = '/login',
    onRetry,
  }: {
    view: CoachPanelView;
    answer: CoachAnswer | null;
    late?: boolean;
    reserve?: Reserve | null;
    names: CoachNames;
    build: string;
    classSlug: string;
    /** The roster's spec, for the title before an answer names it. */
    spec: string;
    playerName: string;
    playerGuid: string;
    firstFight: number;
    /** The roster's class, for the rail's player card. */
    className?: string;
    /** The card's muted line: spec, boss, result. */
    cardLine?: string;
    /** The character page's coach list, for Premium viewers of their own review. */
    historyHref?: string | null;
    refetching?: boolean;
    /** `/login?next=` this exact url, so sign-in returns to the panel. */
    signInHref?: string;
    onRetry: () => void;
  } = $props();

  const displayName = $derived(splitUnitName(playerName).name);
  const full = $derived(answer !== null && answer.access !== 'teaser' ? answer : null);
  const specName = $derived(answer?.spec ?? spec);
  const swapped = $derived(view.state === 'ready_model' && reserve !== null);

  /** The states whose body is the rule-only rendering when the answer says so. */
  const RULES_BODY: readonly CoachPanelView['state'][] = [
    'ready_rules',
    'writing',
    'busy',
    'no_findings',
    'raid_buffs_only',
  ];
  /** The states that name no spec: nothing is known about the reviewed player. */
  const GENERIC: readonly CoachPanelView['state'][] = ['not_available', 'signed_out', 'whole_night'];

  const title = $derived.by(() => {
    if (view.officer) return fill(coachCopy.titleOfficer, { Name: displayName, Spec: specName });
    if (specName === '' || GENERIC.includes(view.state)) return coachCopy.titleGeneric;
    if (view.starter) return fill(coachCopy.titleStarter, { Spec: specName });
    const rulesBody =
      full !== null &&
      RULES_BODY.includes(view.state) &&
      (full.source === 'rules' || full.status === 'writing');
    const rules = rulesBody || view.state === 'not_reviewable';
    return fill(rules ? coachCopy.titleRules : coachCopy.titleModel, { Spec: specName });
  });
  /** Why a pull is not reviewed: the answer's reason lines, else the generic sentence. */
  const notReviewableLines = $derived(
    full !== null && full.review.couldntCheck.length > 0
      ? full.review.couldntCheck
      : [coachCopy.notReviewable],
  );

  const statusText = $derived(
    view.state === 'loading'
      ? coachCopy.loading
      : view.state === 'writing'
        ? late
          ? coachCopy.writingLate
          : coachCopy.writing
        : '',
  );
  const showBody = $derived(
    full !== null &&
      ['writing', 'busy', 'ready_model', 'ready_rules', 'no_findings', 'raid_buffs_only'].includes(
        view.state,
      ),
  );
  const footer = $derived(
    full === null
      ? ''
      : fill(coachCopy.footer, {
          date: reviewedDate(full.reviewed_at),
          labels: basisList(full.basis_labels),
        }),
  );
  /** The basis labels are sentences: one list, each without its own full stop, `; ` between. */
  const basisList = (labels: readonly string[]): string =>
    labels.map((label) => label.replace(/\.$/, '')).join('; ');
  const BUTTON =
    'rounded-control border-line-warm hover:border-gold focus-visible:outline-gold active:bg-card-top inline-flex h-11 items-center border px-4 text-[13px] font-semibold focus-visible:outline-2 focus-visible:outline-offset-2';
</script>

<div class="@container" data-testid="coach-panel" data-state={view.state}>
  <div class="flex flex-col gap-5 @min-[1120px]:flex-row @min-[1120px]:items-start @min-[1120px]:gap-10">
    <section
      class="{CARD} w-full p-[18px] md:p-6 @min-[1120px]:w-[760px] @min-[1120px]:shrink-0"
      data-testid="coach-card"
    >
      <header class="flex flex-wrap items-center justify-between gap-x-4 md:min-h-11">
        <h2
          id="coach-panel-heading"
          tabindex="-1"
          class="font-display text-strong text-[17px] leading-6 font-bold outline-none"
          data-testid="coach-title"
        >
          {title}
        </h2>
        <div
          class="flex h-5 basis-full items-center gap-2 text-[13px] md:basis-auto"
          aria-live="polite"
          data-testid="coach-status"
        >
          {#if view.officer}
            <span class="pill-site" data-testid="coach-officer-pill">{coachCopy.officerPill}</span>
          {:else if statusText !== ''}
            {#if view.state === 'writing' && !late}
              <span class="bg-gold h-2 w-2 rounded-full motion-safe:animate-pulse" aria-hidden="true"></span>
            {/if}
            <span class="text-muted">{statusText}</span>
          {:else if swapped}
            <span class="sr-only">{coachCopy.ready}</span>
          {/if}
        </div>
      </header>

      {#if view.starter}
        <p class="text-muted mt-1 text-[13px] leading-5" data-testid="coach-starter-hint">
          {fill(coachCopy.starterHint, { Spec: specName })}
        </p>
      {/if}
      {#if view.officer}
        <p class="text-muted mt-1 text-[13px] leading-5" data-testid="coach-officer-note">
          {fill(coachCopy.officerNote, { Name: displayName })}
        </p>
      {/if}
      {#if view.state === 'busy'}
        <p
          class="border-line-warm text-muted mt-3 border-l-2 pl-3 text-[13px] leading-5"
          data-testid="coach-busy"
        >
          {coachCopy.busy}
        </p>
      {/if}

      <div class="mt-4" aria-busy={view.state === 'writing' || view.state === 'loading'}>
        {#if showBody && full !== null}
          <CoachReviewBody answer={full} {names} {build} {classSlug} {firstFight} {playerGuid} {reserve} />
        {:else if view.state === 'loading'}
          <div class="flex flex-col gap-3" data-testid="coach-loading">
            <span class="skeleton-block h-4 w-3/4"></span>
            <span class="skeleton-block h-4 w-1/2"></span>
            <span class="border-line mt-4 block h-24 border-l-2 pl-4"
              ><span class="skeleton-block block h-5 w-2/3"></span></span
            >
          </div>
        {:else if (view.state === 'teaser' || view.state === 'teaser_empty') && answer?.access === 'teaser'}
          <CoachTeaser {answer} {names} />
        {:else if view.state === 'preparing'}
          <p class="text-text text-[15px]" data-testid="coach-preparing">
            {fill(coachCopy.preparing, { Spec: specName })}
          </p>
          <button
            type="button"
            class="{BUTTON} mt-4 {refetching ? BUSY_CLASS : ''}"
            aria-busy={refetching}
            disabled={refetching}
            onclick={onRetry}
          >
            {coachCopy.checkAgain}
          </button>
        {:else if view.state === 'not_reviewable'}
          <div class="flex flex-col gap-1" data-testid="coach-not-reviewable">
            {#each notReviewableLines as line, index (index)}
              <p class="text-text text-[15px] leading-[22px]">{line}</p>
            {/each}
          </div>
        {:else if view.state === 'error'}
          <p class="text-text text-[15px]" role="alert" data-testid="coach-error">{coachCopy.error}</p>
          <button
            type="button"
            class="{BUTTON} mt-4 {refetching ? BUSY_CLASS : ''}"
            aria-busy={refetching}
            disabled={refetching}
            onclick={onRetry}
          >
            {coachCopy.rowRetry}
          </button>
        {:else if view.state === 'not_available'}
          <p class="text-text text-[15px]" data-testid="coach-not-available">{coachCopy.notAvailable}</p>
        {:else if view.state === 'signed_out'}
          <p class="text-text text-[15px]" data-testid="coach-signed-out">{coachCopy.signedOut}</p>
          <a
            href={signInHref}
            class="rounded-control border-gold text-gold hover:border-gold-hover mt-4 inline-flex h-11 items-center border px-4 text-[14px] font-semibold"
            data-testid="coach-sign-in">{coachCopy.signIn}</a
          >
        {/if}
      </div>
    </section>

    {#if full !== null && showBody}
      <aside
        class="grid w-full grid-cols-1 items-start gap-5 md:grid-cols-2 @min-[1120px]:flex @min-[1120px]:w-[320px] @min-[1120px]:shrink-0 @min-[1120px]:flex-col"
        data-testid="coach-rail"
      >
        <div class="flex flex-col gap-5 @min-[1120px]:contents">
          <div class="{CARD} hidden items-center gap-3 p-4 md:flex" data-testid="coach-player-card">
            <ClassIcon {className} size={40} />
            <div class="min-w-0">
              <p class="truncate text-[15px] font-semibold" style={`color: ${classColorVar(className)}`}>
                {displayName}
              </p>
              <p class="text-muted text-[13px] leading-5">{cardLine}</p>
            </div>
          </div>
          {#if full.review.couldntCheck.length > 0}
            <section class="{CARD} p-4" data-testid="coach-cannot">
              <h3 class="label text-muted">{coachCopy.labelCannot}</h3>
              <ul class="mt-1 flex flex-col gap-1">
                {#each full.review.couldntCheck as line, index (index)}
                  <li class="text-muted text-[13px] leading-5">{line}</li>
                {/each}
              </ul>
            </section>
          {/if}
        </div>
        <div class="flex flex-col gap-5 @min-[1120px]:contents">
          <div class="{CARD} p-4" data-testid="coach-footer-card">
            <p class="text-muted text-[12px] leading-[18px]" data-testid="coach-footer">{footer}</p>
            {#if historyHref !== null && !view.officer}
              <a
                href={historyHref}
                class="text-text hover:text-strong focus-visible:outline-gold inline-flex min-h-11 items-center text-[13px] underline focus-visible:outline-2"
                data-testid="coach-history">{coachCopy.history}</a
              >
            {/if}
          </div>
        </div>
      </aside>
    {/if}
  </div>
</div>
