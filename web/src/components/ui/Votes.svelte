<!-- web/src/components/ui/Votes.svelte -->
<!-- A vote tally (web layout spec 2026-11-05 §7.3): a yes Tag and a no Tag, then who can vote
     ("2 yes · 1 no of 4 officers"). The viewer's own vote carries the Tag's dot. With `onVote` the
     two Tags (toned once they hold a vote) are the vote buttons (aria-pressed on the viewer's own): pressing it again takes the
     vote back, pressing the other changes it. Loot council candidates, later applications. -->
<script lang="ts">
  import { nextVote, votersLine, votesLine, type VoteChoice } from '../../lib/ui/votes';
  import Tag from './Tag.svelte';

  let {
    yes,
    no,
    voters,
    mine = null,
    noun = 'officer',
    subject,
    onVote,
    disabled = false,
    describedBy,
    testid,
  }: {
    yes: number;
    no: number;
    voters: number;
    mine?: VoteChoice | null;
    /** Who votes, singular ("officer"). */
    noun?: string;
    /** What the vote is on, for the buttons' names ("Seraphine"). */
    subject: string;
    /** The vote to send (null takes the viewer's back); absent, the tally is read-only. */
    onVote?: (vote: VoteChoice | null) => void;
    disabled?: boolean;
    describedBy?: string;
    testid?: string;
  } = $props();

  const choices = $derived([
    { choice: 'yes' as const, count: yes, tone: 'success' as const },
    { choice: 'no' as const, count: no, tone: 'danger' as const },
  ]);
</script>

<span
  class="inline-flex flex-wrap items-center gap-[6px]"
  title={votesLine(yes, no, voters, noun)}
  data-testid={testid}
>
  {#each choices as { choice, count, tone } (choice)}
    {#if onVote}
      <button
        type="button"
        class="ui-vote rounded-field focus-visible:outline-gold inline-flex min-h-11 items-center focus-visible:outline-2 focus-visible:outline-offset-2 disabled:cursor-not-allowed disabled:opacity-45 md:min-h-0"
        aria-pressed={mine === choice ? 'true' : 'false'}
        aria-label={`Vote ${choice} for ${subject}`}
        aria-describedby={disabled ? describedBy : undefined}
        {disabled}
        onclick={() => onVote(nextVote(mine, choice))}
        data-testid={testid === undefined ? undefined : `${testid}-${choice}`}
        ><Tag tone={count > 0 ? tone : 'muted'} dot={mine === choice}>{count} {choice}</Tag></button
      >
    {:else}
      <Tag tone={count > 0 ? tone : 'muted'} dot={mine === choice}>{count} {choice}</Tag>
    {/if}
  {/each}
  <span class="text-muted text-[13px] leading-[18px]">{votersLine(voters, noun)}</span>
</span>
