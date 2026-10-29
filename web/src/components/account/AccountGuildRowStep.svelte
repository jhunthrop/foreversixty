<!-- web/src/components/account/AccountGuildRowStep.svelte -->
<!-- Spec 6.5, in the polish component set: the three things a guild block adds. `plan` is the tier
     tag beside the guild name (`No plan` when it has none), `mark` the warning `Not verified` tag,
     `line` the no-plan row: its sentence under a `See the next step` link row to the checkout
     guild step. The words and the link come from lib/account/guild-row-line. -->
<script lang="ts">
  import {
    guildPlanTag,
    guildRowNextStep,
    notVerifiedMark,
    type GuildRowFacts,
  } from '../../lib/account/guild-row-line';
  import SettingsRow from '../ui/SettingsRow.svelte';
  import Tag from '../ui/Tag.svelte';

  let { guild, part }: { guild: GuildRowFacts; part: 'plan' | 'mark' | 'line' } = $props();

  const plan = $derived(guildPlanTag(guild));
  const mark = $derived(notVerifiedMark(guild));
  const next = $derived(guildRowNextStep(guild));
</script>

{#if part === 'plan' && plan !== null}
  <Tag tone={guild.plan ? 'gold' : 'muted'} testid="account-guild-plan-tag">{plan}</Tag>
{:else if part === 'mark' && mark !== null}
  <Tag tone="warn" testid="account-guild-not-verified">{mark}</Tag>
{:else if part === 'line' && next !== null}
  <SettingsRow
    glyph="card"
    label={next.action}
    sub={next.text}
    href={next.href}
    testid="account-guild-no-plan"
  />
{/if}
