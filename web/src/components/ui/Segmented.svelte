<!-- web/src/components/ui/Segmented.svelte -->
<!-- An inline choice of one among few (web layout spec 2026-11-05 §7.3): the sign-up's In / Late /
     Bench, later any in-row choice. The ControlBar's segmented look (segmented.css), as buttons that
     report a choice instead of links that navigate. A radio group with one roving tab stop: arrows
     move focus and never select, Space and Enter select (lib/ui/segmented.ts). `pending` is pressed
     but not saved yet, drawn dashed, never as checked. -->
<script lang="ts" generics="T extends string">
  import { segmentKeyAction } from '../../lib/ui/segmented';
  import './segmented.css';

  let {
    options,
    value,
    pending = null,
    label,
    size = 'md',
    fill = false,
    disabled = false,
    describedBy,
    onChange,
    testid,
  }: {
    options: readonly { value: T; label: string }[];
    value: T | null;
    pending?: T | null;
    /** The group's accessible name ("Hollowmend's sign-up"). */
    label: string;
    /** `md` is 44 px (the bar's phone size); `sm` 32 px from 768, for a choice inside a row (44 on phone). */
    size?: 'md' | 'sm';
    /** Stretch to the container's width; `phone` below 768 only; else as wide as the labels. */
    fill?: boolean | 'phone';
    disabled?: boolean;
    /** The line that says why the group is disabled. */
    describedBy?: string;
    onChange: (value: T) => void;
    testid?: string;
  } = $props();

  const buttons: HTMLButtonElement[] = [];
  const checkedIndex = $derived(options.findIndex((option) => option.value === value));
  const tabStop = $derived(checkedIndex === -1 ? 0 : checkedIndex);

  function select(index: number): void {
    const option = options[index];
    if (option !== undefined && option.value !== value) onChange(option.value);
  }

  function onKeydown(event: KeyboardEvent, index: number): void {
    const action = segmentKeyAction(event.key, index, options.length);
    if (action === null) return;
    event.preventDefault();
    if (action.kind === 'move') buttons[action.index]?.focus();
    else select(action.index);
  }
</script>

<div
  role="radiogroup"
  aria-label={label}
  aria-describedby={disabled ? describedBy : undefined}
  class={[
    'ui-seg',
    size === 'sm' && 'ui-seg--sm',
    fill === true ? 'w-full' : fill === 'phone' ? 'w-full md:w-fit' : 'w-fit',
  ]}
  data-testid={testid}
>
  {#each options as option, index (option.value)}
    <button
      bind:this={buttons[index]}
      type="button"
      role="radio"
      class="ui-seg-item"
      aria-checked={option.value === value ? 'true' : 'false'}
      aria-busy={pending === option.value ? 'true' : undefined}
      tabindex={index === tabStop ? 0 : -1}
      {disabled}
      data-testid={testid === undefined ? undefined : `${testid}-${option.value}`}
      onclick={() => select(index)}
      onkeydown={(event) => onKeydown(event, index)}>{option.label}</button
    >
  {/each}
</div>
