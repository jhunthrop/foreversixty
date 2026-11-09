<!-- web/src/components/guild/recruit/ApplicantRow.svelte -->
<!-- One open application (web layout spec 2026-11-05 §7.2 Applications; mock web-guild-applications):
     crest, the name in class colour, spec, level, item level and best boss figure over their message
     and availability; in the trail the officers' Votes with the viewer's own, Accept and Decline. -->
<script lang="ts">
  import type { ApplicationVote, OfficerApplication } from '../../../lib/guild/recruit-api';
  import { applicationsCopy as copy } from '../../../lib/guild/recruit-copy';
  import { applicantFacts } from '../../../lib/guild/applications-view';
  import { classTone } from '../../../lib/ui/primitives';
  import ClassCrestRing from '../../character/ClassCrestRing.svelte';
  import Button from '../../ui/Button.svelte';
  import Row from '../../ui/Row.svelte';
  import Votes from '../../ui/Votes.svelte';

  let {
    application,
    busy,
    error,
    onVote,
    onDecide,
  }: {
    application: OfficerApplication;
    busy: boolean;
    error: string;
    onVote: (vote: ApplicationVote | null) => void;
    onDecide: (decision: 'accept' | 'decline') => void;
  } = $props();

  const character = $derived(application.character);
</script>

<Row
  title={character.name}
  titleTone={classTone(character.class ?? '')}
  wrapTrail
  testid={`guild-application-${application.id}`}
>
  {#snippet lead()}<ClassCrestRing characterClass={character.class ?? ''} size={36} />{/snippet}
  {#snippet subContent()}
    {applicantFacts(character)}<span class="block" data-testid="guild-application-message"
      >{application.message === ''
        ? copy.noMessage
        : application.message}{#if application.availability !== ''}<span
          data-testid="guild-application-availability"
        >
          · {copy.availability(application.availability)}</span
        >{/if}</span
    >{#if error !== ''}<span class="text-wipe block" role="alert">{error}</span>{/if}
  {/snippet}
  {#snippet trail()}
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
  {/snippet}
</Row>
