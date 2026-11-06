<!-- web/src/components/guild/recruit/WantedEditor.svelte -->
<!-- The "Looking for" list in the card's edit: one line per class (any spec, or one of its specs)
     with a short note, Remove on each, Add a class up to the API's twelve. -->
<script lang="ts">
  import type { RecruitWantedInput } from '../../../lib/guild/recruit-api';
  import { applicationsCopy as copy } from '../../../lib/guild/recruit-copy';
  import {
    WANTED_MAX,
    WANTED_NOTE_MAX,
    classOptions,
    specOptions,
    wantedLabel,
  } from '../../../lib/guild/applications-view';
  import Button from '../../ui/Button.svelte';
  import Field from '../../ui/Field.svelte';
  import Select from '../../ui/Select.svelte';

  let { lines = $bindable() }: { lines: RecruitWantedInput[] } = $props();

  const classes = classOptions();

  function setLine(index: number, change: Partial<RecruitWantedInput>): void {
    lines = lines.map((line, i) => (i === index ? { ...line, ...change } : line));
  }
</script>

<fieldset class="flex flex-col gap-3" data-testid="guild-recruit-edit-wanted">
  <legend class="text-label tracking-label text-muted mb-2 font-bold uppercase">{copy.card.wanted}</legend>
  {#each lines as line, i (i)}
    <div class="flex flex-wrap items-end gap-2" data-testid="guild-recruit-edit-wanted-line">
      <Select
        id={`guild-recruit-wanted-class-${i}`}
        label={copy.card.wantedClass}
        size="compact"
        options={classes}
        value={line.class}
        onchange={(event) =>
          setLine(i, { class: (event.currentTarget as HTMLSelectElement).value, spec: null })}
      />
      <Select
        id={`guild-recruit-wanted-spec-${i}`}
        label={copy.card.wantedSpec}
        size="compact"
        options={specOptions(line.class)}
        value={line.spec ?? ''}
        onchange={(event) => {
          const value = (event.currentTarget as HTMLSelectElement).value;
          setLine(i, { spec: value === '' ? null : value });
        }}
      />
      <Field
        id={`guild-recruit-wanted-note-${i}`}
        label={copy.card.wantedNote}
        maxlength={WANTED_NOTE_MAX}
        bind:value={lines[i].note}
        grow
      />
      <Button
        size="sm"
        onclick={() => (lines = lines.filter((_, j) => j !== i))}
        testid="guild-recruit-edit-wanted-remove"
        >✕<span class="sr-only">{copy.card.removeWanted(wantedLabel(line))}</span></Button
      >
    </div>
  {/each}
  {#if lines.length < WANTED_MAX}
    <span>
      <Button
        size="sm"
        onclick={() => (lines = [...lines, { class: 'warrior', spec: null, note: '' }])}
        testid="guild-recruit-edit-wanted-add">{copy.card.addWanted}</Button
      >
    </span>
  {/if}
</fieldset>
