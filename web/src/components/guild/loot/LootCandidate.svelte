<!-- web/src/components/guild/loot/LootCandidate.svelte -->
<!-- One candidate for a drop (web layout spec 2026-11-05 §7.2): crest, the name in class colour, then
     the gain in the BiS page's unit and what it replaces, attendance and loot this tier; under it
     the MS or OS Tag (solid when they flagged it, muted when read off their BiS list), BiS for
     them, the officers' Votes with the viewer's yes and no, and Award, which opens the confirm in
     place. A closed drop keeps the facts and drops the controls. -->
<script lang="ts">
  import type { LootCouncilCandidate, MsOs, Vote } from '../../../lib/guild/loot-council-api';
  import { candidateSub, msOsTag } from '../../../lib/guild/loot-council-view';
  import { classSlugFromName } from '../../../lib/report/tree-sizes';
  import { isClassSlug } from '../../../lib/ui/primitives';
  import Crest from '../../ui/Crest.svelte';
  import Button from '../../ui/Button.svelte';
  import Row from '../../ui/Row.svelte';
  import Tag from '../../ui/Tag.svelte';
  import Votes from '../../ui/Votes.svelte';
  import AwardConfirm from './AwardConfirm.svelte';

  let {
    candidate,
    itemName,
    gear,
    closed,
    canAct,
    describedBy,
    confirming,
    busy,
    onVote,
    onAsk,
    onAward,
    onCancel,
  }: {
    candidate: LootCouncilCandidate;
    itemName: string;
    gear: boolean;
    /** The drop is awarded on the night: no votes, no Award. */
    closed: boolean;
    /** False while the plan gate holds the officer's tools: the controls show, disabled. */
    canAct: boolean;
    describedBy?: string;
    /** This candidate's Award confirm is open. */
    confirming: boolean;
    busy: boolean;
    onVote: (vote: Vote | null) => void;
    onAsk: () => void;
    onAward: (msOs: MsOs | null, reason: string | null) => void;
    onCancel: () => void;
  } = $props();

  const slug = $derived(classSlugFromName(candidate.class));
  const tag = $derived(msOsTag(candidate));
  /** More than half the council. */
  const majority = $derived(Math.floor(candidate.votes.voters / 2) + 1);
</script>

<Row
  title={candidate.name}
  titleTone={isClassSlug(slug) ? { classSlug: slug } : undefined}
  sub={candidateSub(candidate)}
  open
  testid={`guild-loot-candidate-${candidate.character_key}`}
>
  {#snippet lead()}<Crest cls={candidate.class} size={32} />{/snippet}
  {#snippet foot()}
    {#if tag !== null}<Tag tone={tag.declared ? 'gold' : 'muted'} title={tag.title} testid="guild-loot-msos"
        >{tag.label}</Tag
      >{/if}
    {#if candidate.bis_for_them}<Tag tone="gold" testid="guild-loot-bis">BiS for them</Tag>{/if}
    <Votes
      yes={candidate.votes.yes}
      no={candidate.votes.no}
      voters={candidate.votes.voters}
      mine={candidate.votes.mine}
      subject={candidate.name}
      onVote={closed ? undefined : onVote}
      disabled={!canAct || busy}
      {describedBy}
      testid="guild-loot-votes"
    />
    {#if !closed && !confirming}
      <span class="ml-auto flex items-center gap-2">
        <span class="text-small text-muted" data-testid="guild-loot-majority"
          >{candidate.votes.yes >= majority
            ? 'Majority reached'
            : `${majority - candidate.votes.yes} more yes for a majority`}</span
        >
        <span class={candidate.votes.yes >= majority ? '' : 'opacity-50'}
          ><Button
            size="sm"
            tone={candidate.votes.yes >= majority ? 'primary' : 'ghost'}
            disabled={!canAct || busy}
            onclick={onAsk}
            testid="guild-loot-award">Award</Button
          ></span
        >
      </span>
    {/if}
    {#if confirming}
      <AwardConfirm
        {itemName}
        name={candidate.name}
        {gear}
        initial={candidate.ms_os}
        {busy}
        onConfirm={onAward}
        {onCancel}
      />
    {/if}
  {/snippet}
</Row>
