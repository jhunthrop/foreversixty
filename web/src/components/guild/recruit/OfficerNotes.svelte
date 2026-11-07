<!-- web/src/components/guild/recruit/OfficerNotes.svelte -->
<!-- Officer notes, the profile's rail for officers only (web layout spec 2026-11-05 §7.2; mock
     web-guild-profile): each note a Row, who wrote it and when, and under it (the rail is narrow) its
     "Officers only" Tag and Delete on the viewer's own (the guild master on any); a one-line FormRow
     adds one. Never mounted for a member. -->
<script lang="ts">
  import { NOTE_MAX, noteSub } from '../../../lib/guild/profile-view';
  import type { MemberNote } from '../../../lib/guild/recruit-api';
  import { profileCopy as copy } from '../../../lib/guild/recruit-copy';
  import Button from '../../ui/Button.svelte';
  import Field from '../../ui/Field.svelte';
  import FormRow from '../../ui/FormRow.svelte';
  import Row from '../../ui/Row.svelte';
  import Section from '../../ui/Section.svelte';
  import Tag from '../../ui/Tag.svelte';

  let {
    notes,
    viewerTimezone,
    locale,
    saving,
    deletingId,
    error,
    onAdd,
    onDelete,
  }: {
    notes: readonly MemberNote[];
    viewerTimezone: string;
    locale: string;
    saving: boolean;
    deletingId: number | null;
    error: string;
    /** Resolves true once the note is saved, so the field clears. */
    onAdd: (body: string) => Promise<boolean>;
    onDelete: (id: number) => void;
  } = $props();

  let body = $state('');

  async function add(event: SubmitEvent): Promise<void> {
    event.preventDefault();
    if (body.trim() === '' || saving) return;
    if (await onAdd(body)) body = '';
  }
</script>

<Section title={copy.notes} meta={copy.notesMeta} testid="guild-profile-notes">
  {#if notes.length === 0}
    <p class="text-muted text-[14px]" data-testid="guild-profile-notes-empty">{copy.noNotes}</p>
  {/if}
  {#each notes as note (note.id)}
    <Row title={note.body} sub={noteSub(note, viewerTimezone, locale)} open testid="guild-profile-note">
      {#snippet foot()}
        <Tag tone="muted">{copy.officersOnly}</Tag>
        {#if note.may_delete}
          <span class="ml-auto">
            <Button
              size="sm"
              disabled={deletingId === note.id}
              onclick={() => onDelete(note.id)}
              testid="guild-profile-note-delete">{copy.deleteNote}</Button
            >
          </span>
        {/if}
      {/snippet}
    </Row>
  {/each}
  <div class="pt-3">
    <FormRow label={copy.addNote} onsubmit={(event) => void add(event)} testid="guild-profile-note-form">
      <!-- The field takes what the rail leaves beside the Save button (a Field's own grow floor is 220). -->
      <div class="min-w-0 flex-1">
        <Field
          id="guild-profile-note-body"
          label={copy.addNote}
          placeholder={copy.addNoteHint}
          maxlength={NOTE_MAX}
          bind:value={body}
          testid="guild-profile-note-body"
        />
      </div>
      {#snippet submit()}
        <Button
          size="sm"
          type="submit"
          disabled={saving || body.trim() === ''}
          testid="guild-profile-note-save">{saving ? copy.saving : copy.saveNote}</Button
        >
      {/snippet}
    </FormRow>
  </div>
  {#if error !== ''}<p class="text-wipe text-[13px]" role="alert" data-testid="guild-profile-note-error">
      {error}
    </p>{/if}
</Section>
