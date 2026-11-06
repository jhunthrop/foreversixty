<!-- web/src/components/guild/planner/CompTargetForm.svelte -->
<!-- Officers: this raid's seats per role (web layout spec 2026-11-05 §7.2, owner decision D13).
     Four number Fields, Save, and Back to the default once the raid has its own target. A field
     the browser cannot read, or a 422 from the API, puts its line under the field; the API's
     `total` line (past the raid's size, or no seats at all) sits under the buttons. -->
<script lang="ts">
  import type { GuildCompRole, GuildCompTarget } from '../../../lib/guild/overview-api';
  import {
    COMP_ROLES,
    compDraft,
    parseCompDraft,
    type CompDraft,
  } from '../../../lib/guild/raid-planner-view';
  import Button from '../../ui/Button.svelte';
  import Field from '../../ui/Field.svelte';

  let {
    target,
    disabled = false,
    busy = null,
    fields = {},
    line = null,
    onSave,
    onReset,
    onRefuse,
    onDone,
  }: {
    target: GuildCompTarget;
    disabled?: boolean;
    busy?: 'save' | 'reset' | null;
    /** Lines by field (`tanks`, `total`), the browser's or the API's. */
    fields?: Readonly<Record<string, string>>;
    line?: string | null;
    onSave: (target: Readonly<Record<GuildCompRole, number>>) => Promise<boolean>;
    onReset: () => Promise<boolean>;
    onRefuse: (fields: Readonly<Record<string, string>>) => void;
    /** A save or reset went through: the caller closes the form. */
    onDone: () => void;
  } = $props();

  const LABELS: Readonly<Record<GuildCompRole, string>> = {
    tanks: 'Tanks',
    healers: 'Healers',
    melee: 'Melee',
    ranged: 'Ranged',
  };

  // The draft starts from the target the form opened on; a poll never overwrites typing.
  // svelte-ignore state_referenced_locally
  let draft = $state<Record<GuildCompRole, string>>({ ...compDraft(target) });

  async function save(event: SubmitEvent): Promise<void> {
    event.preventDefault();
    const parsed = parseCompDraft(draft as CompDraft);
    if (!parsed.ok) {
      onRefuse(parsed.fields);
      return;
    }
    if (await onSave(parsed.target)) onDone();
  }

  async function reset(): Promise<void> {
    if (await onReset()) onDone();
  }
</script>

<form
  class="flex flex-col gap-3"
  onsubmit={(event) => void save(event)}
  novalidate
  data-testid="raid-comp-form"
>
  <div class="grid grid-cols-2 gap-3">
    {#each COMP_ROLES as role (role)}
      <Field
        id={`raid-comp-${role}`}
        name={role}
        label={LABELS[role]}
        inputmode="numeric"
        bind:value={draft[role]}
        error={fields[role]}
        testid={`raid-comp-${role}`}
      />
    {/each}
  </div>
  {#if fields.total}
    <p class="text-wipe text-[13px]" role="alert" data-testid="raid-comp-total">{fields.total}</p>
  {/if}
  {#if line !== null}
    <p class="text-wipe text-[13px]" role="alert" data-testid="raid-comp-line">{line}</p>
  {/if}
  <div class="flex flex-wrap gap-2">
    <Button type="submit" size="sm" disabled={disabled || busy !== null} testid="raid-comp-save"
      >{busy === 'save' ? 'Saving' : 'Save'}</Button
    >
    {#if target.source === 'raid'}
      <Button
        size="sm"
        disabled={disabled || busy !== null}
        onclick={() => void reset()}
        testid="raid-comp-reset">Back to the default</Button
      >
    {/if}
  </div>
</form>
