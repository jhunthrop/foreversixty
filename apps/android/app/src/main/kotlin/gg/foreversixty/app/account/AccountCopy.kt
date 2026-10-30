package gg.foreversixty.app.account

import androidx.compose.runtime.Composable
import androidx.compose.ui.res.stringResource
import gg.foreversixty.app.R
import gg.foreversixty.app.auth.session.DeviceRow
import gg.foreversixty.app.auth.session.RelativeAge
import gg.foreversixty.app.auth.session.SignInFailure
import gg.foreversixty.app.auth.session.SignedOutNotice
import gg.foreversixty.app.model.AccountCharacter
import gg.foreversixty.app.model.Billing
import java.time.Instant
import java.time.ZoneId
import java.time.format.DateTimeFormatter
import java.util.Locale

/** The words for F1's states, all from the copy catalog. Dates: `30 Oct`, `27 Oct 2027`. */
object AccountCopy {
    @Composable
    fun notice(notice: SignedOutNotice): String = when (notice) {
        is SignedOutNotice.SignInFailed -> signInFailure(notice.failure)
        SignedOutNotice.SessionEnded -> stringResource(R.string.signin_signedOut)
        is SignedOutNotice.AccountDeleted -> stringResource(R.string.account_delete_done, longDate(notice.completeBy))
    }

    @Composable
    fun signInFailure(failure: SignInFailure): String = stringResource(
        when (failure) {
            SignInFailure.BattleNet -> R.string.signin_error_bnet
            SignInFailure.Expired -> R.string.signin_error_expired
            SignInFailure.Offline -> R.string.signin_error_offline
        },
    )

    @Composable
    fun age(age: RelativeAge): String = when (age) {
        RelativeAge.Now -> stringResource(R.string.account_age_now)
        is RelativeAge.Minutes -> stringResource(R.string.account_age_minutes, age.n)
        is RelativeAge.Hours -> stringResource(R.string.account_age_hours, age.n)
        is RelativeAge.Days -> stringResource(R.string.account_age_days, age.n)
    }

    @Composable
    fun deviceName(row: DeviceRow): String = if (row.isThisDevice) stringResource(R.string.devices_this) else row.name

    @Composable
    fun deviceMeta(row: DeviceRow): String? = when (val meta = row.meta) {
        is DeviceRow.Meta.SignedIn -> stringResource(R.string.devices_signedIn, shortDate(meta.at))
        is DeviceRow.Meta.LastUsed -> stringResource(R.string.devices_lastUsed, age(meta.age))
        DeviceRow.Meta.None -> null
    }

    /** `Orc Fury Warrior · Level 58` (switcher rows). */
    @Composable
    fun switcherDetail(character: AccountCharacter): String {
        val kind = listOfNotNull(character.race, character.spec, character.className).joinToString(" ") { words(it) }
        val level = character.level ?: return kind
        return stringResource(R.string.account_switcher_detail, kind, level)
    }

    /** `Fury · 58` (the selector's short line); the class when no spec is saved. */
    @Composable
    fun selectorDetail(character: AccountCharacter): String? {
        val level = character.level ?: return null
        val spec = (character.spec ?: character.className)?.let(::words).orEmpty()
        return stringResource(R.string.account_selector_detail, spec, level)
    }

    @Composable
    fun selectorLabel(character: AccountCharacter): String = stringResource(
        R.string.selector_a11y,
        character.name,
        character.spec?.let(::words).orEmpty(),
        character.className?.let(::words).orEmpty(),
        character.level ?: 0,
    )

    /** The site's `accountPlanCopy` lines: `Premium · renews 27 Oct 2027`, or `ends` once cancelled. */
    @Composable
    fun planLine(billing: Billing?): String {
        val end = billing?.currentPeriodEnd?.let { runCatching { Instant.parse(it) }.getOrNull() }
        return when {
            billing == null || end == null -> stringResource(R.string.you_plan_free)
            billing.cancelAtPeriodEnd -> stringResource(R.string.plan_ends, longDate(end))
            else -> stringResource(R.string.plan_renews, longDate(end))
        }
    }

    /** The API's lower-case words (`night elf`, `warrior`) as display words. */
    fun words(raw: String): String = raw.split(' ').joinToString(" ") { it.replaceFirstChar(Char::uppercaseChar) }

    fun shortDate(at: Instant): String = SHORT.format(at)

    fun longDate(at: Instant): String = LONG.format(at)

    /** The catalog is English (en.json); dates match it. */
    private val SHORT = DateTimeFormatter.ofPattern("d MMM", Locale.UK).withZone(ZoneId.systemDefault())
    private val LONG = DateTimeFormatter.ofPattern("d MMM yyyy", Locale.UK).withZone(ZoneId.systemDefault())
}
