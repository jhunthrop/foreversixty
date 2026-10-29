<!-- web/src/components/guild/RaidFormDiscord.svelte -->
<!-- The raid form's Discord block (guided-cards spec P.3): `Post to Discord` as a switch, its helper,
     and `Remind the day before` under it. Every choice of row and line is made in
     lib/guild/raid-form.ts (`discordHelper`); this file draws it. The site has no Switch control yet,
     so each row is the form's CHECK_ROW checkbox with role="switch" and aria-checked; swap for
     ui/Switch.svelte when it lands. The block reserves its height so loading the guild's
     defaults shifts nothing. -->
<script lang="ts">
  import { CHECK_ROW, DISABLED_LOOK } from '../../lib/guild/field-styles';
  import { raidCopy } from '../../lib/guild/raid-copy';
  import { remindLocked, type DiscordHelper, type RaidFormValues } from '../../lib/guild/raid-form';

  interface Props {
    values: RaidFormValues;
    helper: DiscordHelper;
    busy: boolean;
    set: <K extends keyof RaidFormValues>(key: K, value: RaidFormValues[K]) => void;
  }

  let { values, helper, busy, set }: Props = $props();

  const uid = $props.id();
  const postHelperId = `${uid}-post-helper`;
  const remindHelperId = `${uid}-remind-helper`;
  const locked = $derived(remindLocked(values, helper.postRow));
</script>

<div class="flex min-h-[120px] flex-col sm:min-h-[96px]">
  {#if helper.postRow === 'switch'}
    <label class={CHECK_ROW}>
      <input
        type="checkbox"
        role="switch"
        checked={values.postToDiscord}
        aria-checked={values.postToDiscord}
        aria-describedby={postHelperId}
        disabled={busy}
        onchange={(event) => set('postToDiscord', event.currentTarget.checked)}
        data-testid="raid-form-discord"
      />
      {raidCopy.form.postToDiscord}
    </label>
    <p id={postHelperId} class="text-muted mt-1 text-[13px]" data-testid="raid-form-discord-helper">
      {helper.post}
    </p>
  {:else if helper.postRow === 'posted'}
    <p class="text-muted flex min-h-11 items-center text-[13px]" data-testid="raid-form-discord-posted">
      {helper.post}
    </p>
  {:else}
    <label class={CHECK_ROW}>
      <input
        type="checkbox"
        checked={values.postToDiscord}
        disabled={busy}
        onchange={(event) => set('postToDiscord', event.currentTarget.checked)}
        data-testid="raid-form-discord"
      />
      {raidCopy.form.postToDiscordIfMoved}
    </label>
  {/if}

  <div class="mt-2 pl-6">
    {#if helper.remindRow === 'live'}
      <label class="{CHECK_ROW} {locked ? DISABLED_LOOK : ''}">
        <input
          type="checkbox"
          role="switch"
          checked={values.remindDayBefore}
          aria-checked={values.remindDayBefore}
          aria-describedby={remindHelperId}
          disabled={busy || locked}
          onchange={(event) => set('remindDayBefore', event.currentTarget.checked)}
          data-testid="raid-form-remind"
        />
        {raidCopy.form.remindDayBefore}
      </label>
      <p
        id={remindHelperId}
        class="text-muted mt-1 text-[13px] {locked ? DISABLED_LOOK : ''}"
        data-testid="raid-form-remind-helper"
      >
        {helper.remind}
      </p>
    {:else}
      <p class="text-muted flex min-h-11 items-center text-[13px]" data-testid="raid-form-remind-helper">
        {helper.remind}
      </p>
    {/if}
  </div>
</div>
