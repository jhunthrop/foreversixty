<!-- web/src/components/guild/recruit/ApplicantRow.svelte -->
<!-- One open application (guild dense spec, h-guild-tabs): crest, the name in class colour, the spec
     over the message and availability; from 1280 the figures in fixed columns under the list's head
     line (item level, best boss kill as a percentile and its figure, the seat the spec fills for the
     next raid, days waiting), then the officers' Votes with the viewer's own, Accept and Decline. -->
<script lang="ts">
  import type { CompGapRow } from '../../../lib/guild/dense/week-board';
  import type { ApplicationVote, OfficerApplication } from '../../../lib/guild/recruit-api';
  import { applicationsCopy as copy } from '../../../lib/guild/recruit-copy';
  import { applicantFacts } from '../../../lib/guild/applications-view';
  import { seatFill, waitedWords } from '../../../lib/guild/applicant-figures';
  import { outputFigure } from '../../../lib/guild/output-figure';

  import { classTone } from '../../../lib/ui/primitives';
  import Crest from '../../ui/Crest.svelte';
  import Button from '../../ui/Button.svelte';
  import Row from '../../ui/Row.svelte';
  import Votes from '../../ui/Votes.svelte';
  import { APPLICANT_ACTIONS, APPLICANT_COLUMNS } from './applicant-columns';
  import { SEP } from '../../../lib/guild/raid-copy';
  import { signed } from '../../../lib/guild/planner-facts';
  import { parseFigure } from '../../../lib/report/format';

  let {
    application,
    busy,
    error,
    gaps = null,
    median = null,
    now = new Date(),
    onVote,
    onDecide,
  }: {
    application: OfficerApplication;
    busy: boolean;
    error: string;
    /** The next raid's seats against its comp; null leaves the Fills column dashed. */
    gaps?: readonly CompGapRow[] | null;
    /** The guild's median item level; null leaves the delta off. */
    median?: number | null;
    now?: Date;
    onVote: (vote: ApplicationVote | null) => void;
    onDecide: (decision: 'accept' | 'decline') => void;
  } = $props();

  const character = $derived(application.character);
  const parse = $derived(character.best_parse);
  const fill = $derived(seatFill(character, gaps));
</script>

<Row
  title={character.name}
  titleTone={classTone(character.class ?? '')}
  wrapTrail="xl"
  density="slim"
  testid={`guild-application-${application.id}`}
>
  {#snippet lead()}<Crest cls={character.class ?? ''} size={28} />{/snippet}
  {#snippet subContent()}
    <span class="xl:hidden">{applicantFacts(character)}</span><span class="hidden xl:inline"
      >{character.spec_name ?? ''}</span
    >{#if fill !== null}&#32;<span
        class="block whitespace-nowrap xl:ml-1 xl:inline {fill.short ? 'text-gold' : 'text-muted'}"
        data-testid="guild-application-fills">{fill.text}</span
      >{/if}<span class="block"
      >{#if application.availability !== ''}<span data-testid="guild-application-availability"
          >{copy.availability(application.availability)}{SEP}</span
        >{/if}<span data-testid="guild-application-message"
        >{application.message === '' ? copy.noMessage : application.message}</span
      ></span
    >{#if error !== ''}<span class="text-wipe block" role="alert">{error}</span>{/if}
  {/snippet}
  {#snippet trail()}
    <span class="hidden {APPLICANT_COLUMNS} text-[14px] xl:grid" data-testid="guild-application-figures">
      <span class="text-strong"
        >{character.item_level ?? '–'}{#if median !== null && character.item_level !== null}<span
            class="text-muted ml-1 font-normal">{signed(character.item_level - median)}</span
          >{/if}</span
      >
      {#if parse === null}<span class="text-muted">–</span>
      {:else}
        <span
          >{#if parse.percentile !== null}{@const figure = parseFigure(parse.percentile)}<b
              style:color={figure.color}
              aria-label={figure.label}>{figure.text}</b
            >{/if}
          <span class="text-muted">{outputFigure(parse.value)}</span></span
        >
      {/if}
      <span>{waitedWords(application.created_at, now)}</span>
    </span>
    <span class="flex flex-wrap items-center justify-end gap-2 {APPLICANT_ACTIONS}">
      <Votes
        yes={application.votes.yes}
        no={application.votes.no}
        voters={application.votes.voters}
        mine={application.votes.mine}
        subject={character.name}
        {onVote}
        disabled={busy}
        testid="guild-application-votes"
      />
      <Button
        size="sm"
        tone="secondary"
        disabled={busy}
        onclick={() => onDecide('accept')}
        testid="guild-application-accept">{copy.accept}</Button
      >
      <Button
        size="sm"
        tone="ghost"
        disabled={busy}
        onclick={() => onDecide('decline')}
        testid="guild-application-decline">{copy.decline}</Button
      >
    </span>
  {/snippet}
</Row>
