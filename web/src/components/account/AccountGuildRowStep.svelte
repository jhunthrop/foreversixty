<!-- web/src/components/account/AccountGuildRowStep.svelte -->
<!-- What a guild on /account adds to its rows (premium spec 2026-11-08 section 7.4): `plan` is the
     tier Tag beside the guild name (`No plan` when it has none), `mark` the muted `Not verified` Tag,
     `line` the no-plan slim row: its sentence with `See the next step ›` to the checkout guild
     step. The words and the link come from lib/account/guild-row-line. -->
<script lang="ts">
  import {
    guildPlanTag,
    guildRowNextStep,
    notVerifiedMark,
    type GuildRowFacts,
  } from '../../lib/account/guild-row-line';
  import { LINK } from '../../lib/ui/primitives';
  import Row from '../ui/Row.svelte';
  import Tag from '../ui/Tag.svelte';

  let { guild, part }: { guild: GuildRowFacts; part: 'plan' | 'mark' | 'line' } = $props();

  const plan = $derived(guildPlanTag(guild));
  const mark = $derived(notVerifiedMark(guild));
  const next = $derived(guildRowNextStep(guild));
</script>

{#if part === 'plan' && plan !== null}
  <Tag tone={guild.plan ? 'gold' : 'muted'} testid="account-guild-plan-tag">{plan}</Tag>
{:else if part === 'mark' && mark !== null}
  <Tag tone="muted" testid="account-guild-not-verified">{mark}</Tag>
{:else if part === 'line' && next !== null}
  <Row title={next.text} density="slim" open wrapTrail testid="account-guild-no-plan">
    {#snippet trail()}
      <a class="{LINK} text-[15px] whitespace-nowrap" href={next.href} aria-label={next.label}
        >{next.action} ›</a
      >
    {/snippet}
  </Row>
{/if}
