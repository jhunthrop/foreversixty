<!-- web/src/components/guild/recruit/RecruitEdit.svelte -->
<!-- The rail of the applications page (mock web-guild-applications): the public card's fields as
     visitors read them (what we look for, the raid nights), each with Edit. Edit opens the form in
     place: taking applications, why this guild, raid nights, what we look for and the Discord role
     an accepted applicant gets, with one line of help. One Save sends the whole card. -->
<script lang="ts">
  import {
    draftOf,
    NIGHTS_MAX,
    PITCH_MAX,
    roleIdError,
    type CardDraft,
  } from '../../../lib/guild/applications-view';
  import { wantedSummary } from '../../../lib/guild/recruit-view';
  import type { RecruitCard } from '../../../lib/guild/recruit-api';
  import { applicationsCopy as copy } from '../../../lib/guild/recruit-copy';
  import Button from '../../ui/Button.svelte';
  import Row from '../../ui/Row.svelte';
  import Field from '../../ui/Field.svelte';
  import Section from '../../ui/Section.svelte';
  import Switch from '../../ui/Switch.svelte';
  import WantedEditor from './WantedEditor.svelte';

  let {
    card,
    roleId,
    recruitHref,
    saving,
    line,
    onSave,
    editing: startEditing = false,
  }: {
    card: RecruitCard;
    roleId: string | null;
    recruitHref: string;
    saving: boolean;
    line: { tone: 'ok' | 'error'; text: string } | null;
    onSave: (draft: CardDraft) => void;
    /** Start with the form open (a test's render); closed, the card's fields read with Edit. */
    editing?: boolean;
  } = $props();

  // The form starts from the card as it stands and keeps the officer's edits until they save.
  // svelte-ignore state_referenced_locally
  let draft = $state<CardDraft>(draftOf(card, roleId));
  const roleError = $derived(roleIdError(draft.roleId));
  // svelte-ignore state_referenced_locally
  let editing = $state(startEditing);
  const fields = $derived([
    { key: 'wanted', title: copy.card.wanted, sub: wantedSummary(card.wanted) },
    { key: 'nights', title: copy.card.nights, sub: card.raid_nights.text ?? copy.card.noNights },
  ]);

  function save(event: SubmitEvent): void {
    event.preventDefault();
    if (roleError !== '' || saving) return;
    onSave($state.snapshot(draft));
  }
</script>

<Section
  title={copy.card.heading}
  action={{ label: copy.card.viewPublic, href: recruitHref }}
  testid="guild-recruit-edit"
>
  {#if !editing}
    {#each fields as field (field.key)}
      <Row title={field.title} sub={field.sub} open testid={`guild-recruit-edit-${field.key}-row`}>
        {#snippet trail()}
          <Button
            size="sm"
            tone="ghost"
            label={`${copy.card.edit}: ${field.title}`}
            onclick={() => (editing = true)}
            testid={`guild-recruit-edit-${field.key}-open`}>{copy.card.edit}</Button
          >
        {/snippet}
      </Row>
    {/each}
  {:else}
    <form class="flex flex-col gap-4" onsubmit={save} aria-label={copy.card.heading}>
      <div class="flex min-h-11 items-center justify-between gap-3">
        <span id="guild-recruit-edit-open-label" class="text-strong text-[14px] font-bold"
          >{copy.card.open}</span
        >
        <Switch
          checked={draft.open}
          labelledby="guild-recruit-edit-open-label"
          onchange={(next) => (draft.open = next)}
          testid="guild-recruit-edit-open"
        />
      </div>
      <Field
        id="guild-recruit-edit-pitch"
        label={copy.card.pitch}
        placeholder={copy.card.pitchHint}
        maxlength={PITCH_MAX}
        rows={3}
        bind:value={draft.pitch}
        testid="guild-recruit-edit-pitch"
      />
      <Field
        id="guild-recruit-edit-nights"
        label={copy.card.nights}
        placeholder={copy.card.nightsHint}
        maxlength={NIGHTS_MAX}
        bind:value={draft.nights}
        testid="guild-recruit-edit-nights"
      />
      <WantedEditor bind:lines={draft.wanted} />
      <Field
        id="guild-recruit-edit-role"
        label={copy.card.role}
        inputmode="numeric"
        hint={copy.card.roleHint}
        error={roleError === '' ? undefined : roleError}
        bind:value={draft.roleId}
        testid="guild-recruit-edit-role"
      />
      <div class="flex flex-wrap items-center gap-3">
        <Button
          tone="primary"
          type="submit"
          disabled={saving || roleError !== ''}
          testid="guild-recruit-edit-save">{saving ? copy.card.saving : copy.card.save}</Button
        >
        {#if line !== null}
          <span
            class="text-[13px] {line.tone === 'ok' ? 'text-kill' : 'text-wipe'}"
            role={line.tone === 'ok' ? 'status' : 'alert'}
            data-testid="guild-recruit-edit-line">{line.text}</span
          >
        {/if}
      </div>
    </form>
  {/if}
</Section>
