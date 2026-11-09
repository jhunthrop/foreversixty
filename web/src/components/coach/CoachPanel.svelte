<!-- web/src/components/coach/CoachPanel.svelte -->
<!-- The Coach tab's panel (spec 3.1, 4.2 and 6; logs-reports 6.2 Coach): a Section titled for the
     review (meta: who it is for, spec, boss, result) whose Plate holds the review in type hierarchy,
     then the open rail Sections `What the coach could not check` and `Measured against`; the rail
     sits under the review, and beside it (320 px) only when the content column is 1120 px or wider
     (a container query, so nothing ever runs past the column). It is presentational: the container owns the
     fetch (panel-state) and hands over the view, the answer and the reserve. The status sits in the
     Section's head row and never moves a line; a writing review is held at its reserved
     heights. The review is always its natural height: the tab's reserve lives on CoachTab. -->
<script lang="ts">
  import { busyLine } from '../../lib/coach/busy-line';
  import { COACH_SEPARATOR, coachCopy, fill } from '../../lib/coach/copy';
  import { footerLine } from '../../lib/coach/footer';
  import type { CoachNames } from '../../lib/coach/names';
  import type { CoachAnswer } from '../../lib/coach/panel-state.svelte';
  import type { Reserve } from '../../lib/coach/reserves';
  import type { CoachPanelView } from '../../lib/coach/states';
  import { LINK } from '../../lib/ui/primitives';
  import { splitUnitName } from '../../lib/characters';
  import Button from '../ui/Button.svelte';
  import Plate from '../ui/Plate.svelte';
  import Section from '../ui/Section.svelte';
  import Tag from '../ui/Tag.svelte';
  import CoachReviewBody from './CoachReviewBody.svelte';
  import CoachTeaser from './CoachTeaser.svelte';
  import CoachText from './CoachText.svelte';

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
    /** The Section's meta after the name: spec, boss, result. */
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
  /** The status words last shown: once the review swaps in, they stay as an invisible placeholder,
   *  so on a narrow card (where the status wraps under the title) the head keeps its height and
   *  nothing below moves (the in-place swap's promise). */
  let lastStatus = $state('');
  let lastDot = $state(false);
  $effect(() => {
    if (statusText === '') return;
    lastStatus = statusText;
    lastDot = view.state === 'writing' && !late;
  });
  const showBody = $derived(
    full !== null &&
      ['writing', 'busy', 'ready_model', 'ready_rules', 'no_findings', 'raid_buffs_only'].includes(
        view.state,
      ),
  );
  const footer = $derived(full === null ? '' : footerLine(full.reviewed_at, full.basis_labels));
  /** The Section's meta: who the review is for, then the card line (spec, boss, result). */
  const meta = $derived(cardLine === '' ? displayName : `${displayName}${COACH_SEPARATOR}${cardLine}`);
</script>

<div class="@container" data-testid="coach-panel" data-state={view.state}>
  <div
    class="flex flex-col gap-[var(--section-gap)] @min-[1120px]:grid @min-[1120px]:grid-cols-[minmax(0,1fr)_320px] @min-[1120px]:items-start @min-[1120px]:gap-x-10"
  >
    <Section {title} {meta} focusId="coach-panel-heading" titleTestid="coach-title" testid="coach-card">
      {#snippet head()}
        <span class="flex h-5 items-center gap-2 text-[13px]" aria-live="polite" data-testid="coach-status">
          {#if view.officer}
            <Tag tone="gold" testid="coach-officer-pill">{coachCopy.officerPill}</Tag>
          {:else if statusText !== ''}
            {#if view.state === 'writing' && !late}
              <span class="bg-gold h-2 w-2 rounded-full motion-safe:animate-pulse" aria-hidden="true"></span>
            {/if}
            <span class="text-muted">{statusText}</span>
          {:else if swapped}
            <span class="sr-only">{coachCopy.ready}</span>
            {#if lastStatus !== ''}
              {#if lastDot}<span class="invisible h-2 w-2" aria-hidden="true"></span>{/if}
              <span class="invisible" aria-hidden="true">{lastStatus}</span>
            {/if}
          {/if}
        </span>
      {/snippet}
      <Plate flush={false}>
        {#if view.starter}
          <p class="text-small text-muted mb-3" data-testid="coach-starter-hint">
            {fill(coachCopy.starterHint, { Spec: specName })}
          </p>
        {/if}
        {#if view.officer}
          <p class="text-small text-muted mb-3" data-testid="coach-officer-note">
            {fill(coachCopy.officerNote, { Name: displayName })}
          </p>
        {/if}
        {#if view.state === 'busy'}
          <p class="text-small text-ember mb-3" data-testid="coach-busy">
            {full === null ? coachCopy.busy : busyLine(full)}
          </p>
        {/if}

        <div aria-busy={view.state === 'writing' || view.state === 'loading'}>
          {#if showBody && full !== null}
            <CoachReviewBody answer={full} {names} {build} {classSlug} {firstFight} {playerGuid} {reserve} />
          {:else if view.state === 'loading'}
            <div class="flex flex-col gap-3" data-testid="coach-loading">
              <span class="skeleton-block h-4 w-3/4"></span>
              <span class="skeleton-block h-4 w-1/2"></span>
              <span class="skeleton-block mt-4 h-3 w-24"></span>
              <span class="skeleton-block h-6 w-2/3"></span>
              <span class="skeleton-block h-12 w-full"></span>
            </div>
          {:else if (view.state === 'teaser' || view.state === 'teaser_empty') && answer?.access === 'teaser'}
            <CoachTeaser {answer} {names} />
          {:else if view.state === 'preparing'}
            <p class="text-text text-[15px]" data-testid="coach-preparing">
              {fill(coachCopy.preparing, { Spec: specName })}
            </p>
            <span class="mt-4 inline-flex" aria-busy={refetching}>
              <Button disabled={refetching} onclick={onRetry}>{coachCopy.checkAgain}</Button>
            </span>
          {:else if view.state === 'not_reviewable'}
            <div class="flex flex-col gap-1" data-testid="coach-not-reviewable">
              {#each notReviewableLines as line, index (index)}
                <p class="text-text text-[15px] leading-6">{line}</p>
              {/each}
            </div>
          {:else if view.state === 'error'}
            <p class="text-text text-[15px]" role="alert" data-testid="coach-error">{coachCopy.error}</p>
            <span class="mt-4 inline-flex" aria-busy={refetching}>
              <Button disabled={refetching} onclick={onRetry}>{coachCopy.rowRetry}</Button>
            </span>
          {:else if view.state === 'not_available'}
            <p class="text-text text-[15px]" data-testid="coach-not-available">{coachCopy.notAvailable}</p>
          {:else if view.state === 'signed_out'}
            <p class="text-text text-[15px]" data-testid="coach-signed-out">{coachCopy.signedOut}</p>
            <span class="mt-4 inline-flex">
              <Button tone="primary" href={signInHref} testid="coach-sign-in">{coachCopy.signIn}</Button>
            </span>
          {/if}
        </div>
      </Plate>
    </Section>

    {#if full !== null && showBody}
      <aside class="flex flex-col gap-[var(--section-gap)]" data-testid="coach-rail">
        {#if full.review.couldntCheck.length > 0}
          <Section title={coachCopy.labelCannot} level={3} testid="coach-cannot">
            <ul class="flex flex-col gap-1">
              {#each full.review.couldntCheck as line, index (index)}
                <li class="text-small text-muted">
                  <CoachText text={line} {names} {build} {classSlug} />
                </li>
              {/each}
            </ul>
          </Section>
        {/if}
        <Section title={coachCopy.labelBasis} level={3} testid="coach-footer-card">
          <p class="text-small text-muted" data-testid="coach-footer">{footer}</p>
          {#if historyHref !== null && !view.officer}
            <a
              href={historyHref}
              class="{LINK} mt-1 inline-flex min-h-11 items-center text-[14px]"
              data-testid="coach-history">{coachCopy.history} ›</a
            >
          {/if}
        </Section>
      </aside>
    {/if}
  </div>
</div>
