<!-- web/src/components/guild/recruit/ApplyForm.svelte -->
<!-- The application, one FormRow (web layout spec 2026-11-05 §7.2 Recruitment): one of the viewer's
     own characters on the guild's realm, a message (1000 at most), when they can raid, and the
     Discord opt-in, off until they turn it on. The API's refusal reads in words under the form. -->
<script lang="ts">
  import type { ApplicationInput } from '../../../lib/guild/recruit-api';
  import { recruitCopy as copy } from '../../../lib/guild/recruit-copy';
  import { AVAILABILITY_MAX, MESSAGE_MAX, type CharacterChoice } from '../../../lib/guild/recruit-view';
  import Button from '../../ui/Button.svelte';
  import Field from '../../ui/Field.svelte';
  import FormRow from '../../ui/FormRow.svelte';
  import Select from '../../ui/Select.svelte';

  let {
    characters,
    busy,
    error,
    onApply,
  }: {
    characters: readonly CharacterChoice[];
    busy: boolean;
    /** The last refusal, in words; '' for none. */
    error: string;
    onApply: (input: ApplicationInput) => void;
  } = $props();

  // svelte-ignore state_referenced_locally
  let characterKey = $state(characters.length === 1 ? characters[0].value : '');
  let message = $state('');
  let availability = $state('');
  let discordDm = $state(false);
  let pickError = $state('');

  const options = $derived(
    characters.map((c) => ({
      value: c.value,
      label: c.label,
      meta: c.sub,
      colour: c.classSlug === '' ? undefined : `--color-class-${c.classSlug}`,
    })),
  );

  function send(event: SubmitEvent): void {
    event.preventDefault();
    if (busy) return;
    if (characterKey === '') {
      pickError = copy.apply.pickFirst;
      return;
    }
    pickError = '';
    onApply({
      character_key: characterKey,
      message: message.trim(),
      availability: availability.trim(),
      discord_dm: discordDm,
    });
  }
</script>

<FormRow label={copy.apply.heading} onsubmit={send} testid="guild-recruit-form">
  <div class="flex min-w-[220px] flex-1 flex-col gap-[6px]">
    <label class="text-label text-muted font-bold" for="guild-recruit-character">{copy.apply.character}</label
    >
    <Select
      id="guild-recruit-character"
      size="compact"
      bind:value={characterKey}
      {options}
      placeholder={copy.apply.pickCharacter}
      error={pickError === '' ? undefined : pickError}
      testid="guild-recruit-character"
    />
  </div>
  <Field
    id="guild-recruit-availability"
    label={copy.apply.availability}
    placeholder={copy.apply.availabilityHint}
    maxlength={AVAILABILITY_MAX}
    bind:value={availability}
    grow
    testid="guild-recruit-availability"
  />
  <div class="w-full">
    <Field
      id="guild-recruit-message"
      label={copy.apply.message}
      placeholder={copy.apply.messageHint}
      maxlength={MESSAGE_MAX}
      rows={4}
      hint={`${message.length} / ${MESSAGE_MAX}`}
      bind:value={message}
      testid="guild-recruit-message"
    />
  </div>
  <label class="text-text flex min-h-11 w-full items-center gap-3 text-[14px]">
    <input
      type="checkbox"
      class="accent-gold h-[18px] w-[18px]"
      bind:checked={discordDm}
      data-testid="guild-recruit-discord-dm"
    />
    {copy.apply.discordDm}
  </label>
  {#if error !== ''}
    <p class="text-wipe w-full text-[13px]" role="alert" data-testid="guild-recruit-error">{error}</p>
  {/if}
  {#snippet submit()}
    <Button tone="primary" type="submit" disabled={busy} testid="guild-recruit-submit"
      >{busy ? copy.apply.sending : copy.apply.submit}</Button
    >
  {/snippet}
</FormRow>
