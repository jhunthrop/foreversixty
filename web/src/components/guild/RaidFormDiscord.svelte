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

  /** The inline switch: a 36 x 20 track with a 14 px thumb, gold when on. The input is the track, so
   *  the row's label stays the 44 px hit area. The polish lane's shared Switch replaces this. */
  const SWITCH =
    'relative h-5 w-9 shrink-0 cursor-pointer appearance-none rounded-full border border-line bg-bg transition-colors ' +
    "before:absolute before:top-1/2 before:left-0.5 before:h-3.5 before:w-3.5 before:-translate-y-1/2 before:rounded-full before:bg-muted before:transition-transform before:content-[''] " +
    'checked:border-gold checked:bg-gold checked:before:translate-x-4 checked:before:bg-bg ' +
    'focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-gold disabled:cursor-not-allowed';

  const uid = $props.id();
  const postHelperId = `${uid}-post-helper`;
  const remindHelperId = `${uid}-remind-helper`;
  const locked = $derived(remindLocked(values, helper.postRow));
</script>

<div class="-mt-2 flex min-h-[164px] flex-col sm:min-h-[144px]">
  {#if helper.postRow === 'switch'}
    <label class={CHECK_ROW}>
      <input
        type="checkbox"
        class={SWITCH}
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
    <p class="text-muted mt-1 text-[13px]" data-testid="raid-form-discord-posted">
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
          class={SWITCH}
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
      <p class="text-muted mt-1 text-[13px]" data-testid="raid-form-remind-helper">
        {helper.remind}
      </p>
    {/if}
  </div>
</div>
