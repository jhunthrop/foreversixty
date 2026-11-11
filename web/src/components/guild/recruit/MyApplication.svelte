<!-- web/src/components/guild/recruit/MyApplication.svelte -->
<!-- The viewer's own application as one open Row: the character, where it stands in a sentence,
     when it was sent or decided, and Withdraw while it waits. -->
<script lang="ts">
  import type { GuildApplication } from '../../../lib/guild/recruit-api';
  import { recruitCopy as copy } from '../../../lib/guild/recruit-copy';
  import { applicationLine, applicationStamp } from '../../../lib/guild/recruit-view';
  import { classTone } from '../../../lib/ui/primitives';
  import Crest from '../../ui/Crest.svelte';
  import Button from '../../ui/Button.svelte';
  import Row from '../../ui/Row.svelte';
  import Section from '../../ui/Section.svelte';
  import Tag from '../../ui/Tag.svelte';

  let {
    application,
    guildName,
    viewerTimezone,
    busy,
    error,
    onWithdraw,
  }: {
    application: GuildApplication;
    guildName: string;
    viewerTimezone: string;
    busy: boolean;
    error: string;
    onWithdraw: () => void;
  } = $props();

  const pending = $derived(application.status === 'pending');
  const TONE = { pending: 'gold', accepted: 'success', declined: 'muted', withdrawn: 'muted' } as const;
  const WORD = {
    pending: 'Waiting',
    accepted: 'Accepted',
    declined: 'Declined',
    withdrawn: 'Withdrawn',
  } as const;
</script>

<Section title={copy.mine.heading} panel testid="guild-recruit-mine">
  <Row
    title={application.character.name}
    titleTone={classTone(application.character.class ?? '')}
    sub={`${applicationLine(application, guildName)} ${applicationStamp(application, viewerTimezone)}.`}
    open
    wrapTrail
    testid="guild-recruit-mine-row"
  >
    {#snippet lead()}<Crest cls={application.character.class ?? ''} size={32} />{/snippet}
    {#snippet trail()}
      <Tag tone={TONE[application.status]} testid="guild-recruit-mine-status">{WORD[application.status]}</Tag>
      {#if pending}
        <Button size="sm" disabled={busy} onclick={onWithdraw} testid="guild-recruit-withdraw"
          >{busy ? copy.mine.withdrawing : copy.mine.withdraw}</Button
        >
      {/if}
    {/snippet}
  </Row>
  {#if error !== ''}<p class="text-wipe text-[13px]" role="alert">{error}</p>{/if}
</Section>
