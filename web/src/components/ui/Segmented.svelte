<!-- web/src/components/ui/Segmented.svelte -->
<!-- The segmented control (web layout spec 2026-11-05 §3.4): a 1 px outlined group, items divided by
     hairlines, 14 px 600 text, the selected one in gold on the selected ground with a 2 px gold
     underline. A radio group of links (`?param=value`, or the option's own href for a path choice
     such as the spec), so it works before any script; with `onchange` a click is handled in place.
     Arrow keys move and select, Home and End jump, Tab leaves (lib/ui/control-bar.ts). 40 px high
     from 768, 44 px and full width below it when `fill` is set. -->
<script lang="ts">
  import { checkRadio, radioKeyTarget } from '../../lib/ui/control-bar';
  import type { ControlOption } from '../../lib/ui/primitives';

  let {
    param,
    options,
    value,
    label,
    labelledby,
    fill = false,
    onchange,
    testid,
  }: {
    param: string;
    options: readonly ControlOption[];
    value: string;
    /** The group's accessible name, when no visible label names it. */
    label?: string;
    labelledby?: string;
    /** Full width below 768 (the phone control stack). */
    fill?: boolean;
    onchange?: (param: string, value: string) => void;
    testid?: string;
  } = $props();

  function href(option: ControlOption): string {
    return option.href ?? `?${new URLSearchParams({ [param]: option.value }).toString()}`;
  }

  function choose(event: MouseEvent, option: ControlOption): void {
    if (onchange === undefined || option.href !== undefined) return;
    event.preventDefault();
    if (option.disabled) return;
    const change = checkRadio(event.currentTarget as HTMLElement);
    if (change) onchange(change.param, change.value);
  }

  function onkeydown(event: KeyboardEvent): void {
    if (onchange === undefined) return;
    const group = event.currentTarget as HTMLElement;
    const radios = [...group.querySelectorAll<HTMLElement>('[role="radio"]:not([aria-disabled="true"])')];
    const target = radioKeyTarget(event.key, radios.indexOf(event.target as HTMLElement), radios.length);
    if (target === null) return;
    event.preventDefault();
    radios[target]!.focus();
    radios[target]!.click();
  }
</script>

<!-- svelte-ignore a11y_interactive_supports_focus -->
<div
  class="ui-seg"
  class:ui-seg--fill={fill}
  role="radiogroup"
  aria-label={label}
  aria-labelledby={labelledby}
  data-param={param}
  data-testid={testid ?? `control-${param}`}
  onkeydown={onchange === undefined ? undefined : onkeydown}
>
  {#each options as option (option.value)}
    {@const checked = option.value === value}
    <a
      class="ui-seg-item"
      role="radio"
      href={href(option)}
      aria-checked={checked}
      aria-disabled={option.disabled ? 'true' : undefined}
      title={option.disabled ? (option.disabledReason ?? 'No data yet at this level') : undefined}
      tabindex={checked ? 0 : -1}
      data-value={option.value}
      data-navigate={option.href !== undefined ? 'true' : undefined}
      data-testid={`control-${param}-${option.value}`}
      onclick={(event) => choose(event, option)}
    >
      {#if option.icon}<img src={option.icon} alt="" width="22" height="22" class="rounded-[4px]" />{/if}
      <span>{option.label}</span>
      {#if option.count !== undefined}<span class="ui-seg-count tabular-nums">{option.count}</span>{/if}
    </a>
  {/each}
</div>

<style>
  .ui-seg {
    display: inline-flex;
    border: 1px solid var(--color-line);
    border-radius: var(--radius-field);
    overflow: hidden;
    max-width: 100%;
  }
  .ui-seg-item {
    flex: none;
    display: inline-flex;
    align-items: center;
    justify-content: center;
    gap: 8px;
    height: 40px;
    padding: 0 14px;
    font-size: 14px;
    line-height: 16px;
    font-weight: 600;
    color: var(--color-nav);
    white-space: nowrap;
    border-left: 1px solid var(--color-line-soft);
  }
  .ui-seg-item:first-child {
    border-left: 0;
  }
  .ui-seg-item:active {
    background: var(--color-pressed);
  }
  .ui-seg-item:focus-visible {
    outline: 2px solid var(--color-gold);
    outline-offset: -2px;
  }
  .ui-seg-item[aria-checked='true'] {
    background: var(--color-selected);
    color: var(--color-gold);
    box-shadow: inset 0 -2px 0 var(--color-gold);
  }
  .ui-seg-item[aria-disabled='true'] {
    opacity: 0.5;
    color: var(--color-muted);
    cursor: not-allowed;
  }
  .ui-seg-count {
    font-size: 12px;
    color: var(--color-muted);
  }
  .ui-seg-item[aria-checked='true'] .ui-seg-count {
    color: var(--color-gold);
  }
  @media (hover: hover) and (pointer: fine) {
    .ui-seg-item:not([aria-checked='true']):not([aria-disabled='true']):hover {
      background: var(--color-hover);
      color: var(--color-strong);
    }
  }
  @media (max-width: 767px) {
    .ui-seg--fill {
      display: flex;
      width: 100%;
    }
    .ui-seg--fill .ui-seg-item {
      flex: 1;
      height: 44px;
      padding: 0 8px;
    }
  }
</style>
