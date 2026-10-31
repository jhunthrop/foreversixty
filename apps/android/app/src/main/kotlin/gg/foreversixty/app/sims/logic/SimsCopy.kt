package gg.foreversixty.app.sims.logic

import android.content.res.Resources
import gg.foreversixty.app.R
import gg.foreversixty.app.auth.session.RelativeAge
import gg.foreversixty.app.model.SimRow
import gg.foreversixty.app.phrase
import java.time.Duration
import java.time.Instant
import java.time.ZoneId
import java.time.format.DateTimeFormatter
import java.util.Locale

/** Every F5 sentence built from data (spec §5.5); the words themselves live in the copy catalog. */
class SimsCopy(private val res: Resources, private val zone: ZoneId = ZoneId.systemDefault()) {
    fun estimate(estimate: RunEstimate): String = when (estimate) {
        RunEstimate.AboutAMinute -> res.getString(R.string.sims_estimate_minute)
        is RunEstimate.Minutes -> res.getString(R.string.sims_estimate_minutes, estimate.n)
        is RunEstimate.Hours -> res.getString(R.string.sims_estimate_hours, estimate.n)
    }

    /** The muted words after `Running`: the estimate, `taking longer than usual`, or nothing. */
    fun stateSuffix(row: SimRow, state: ServerRunState): String = when (state) {
        ServerRunState.Running -> row.estimateSeconds?.let { res.getString(R.string.sims_takes, estimate(RunEstimate.of(it))) }.orEmpty()
        ServerRunState.RunningOverdue -> res.getString(R.string.sims_takesLong)
        else -> ""
    }

    fun failedLine(row: SimRow): String {
        val reason = row.failureReason
        return if (reason.isNullOrEmpty()) res.getString(R.string.sims_failedPlain) else res.getString(R.string.sims_failed, reason)
    }

    fun finishedLine(row: SimRow, now: Instant): String =
        res.getString(R.string.sims_finished, age(instant(row.finishedAt ?: row.createdAt), now))

    /** Started, Finished or Ended, with its age (the right half of the state line). */
    fun ageLine(row: SimRow, state: ServerRunState, now: Instant): String = when {
        state.isDone -> finishedLine(row, now)
        state == ServerRunState.Failed -> res.getString(R.string.sims_ended, age(instant(row.finishedAt ?: row.createdAt), now))
        else -> res.getString(R.string.sims_started, age(instant(row.createdAt), now))
    }

    /** Relative inside 14 days (`2 h ago`), else a short date (`2 Nov`), as the site's `ageWords`. */
    fun age(at: Instant?, now: Instant): String {
        if (at == null) return RelativeAge.Now.phrase(res)
        return if (Duration.between(at, now) < RELATIVE_CEILING) RelativeAge.of(at, now).phrase(res) else shortDate(at)
    }

    /** `Saved from the browser · 26 Oct` or `Ran on the server · 26 Oct`. */
    fun origin(row: SimRow): String {
        val date = instant(row.createdAt)?.let(::shortDate).orEmpty()
        return if (row.lane == ServerRuns.SERVER_LANE) res.getString(R.string.sims_ranServer, date) else res.getString(R.string.sims_savedBrowser, date)
    }

    fun shortDate(at: Instant): String = SHORT_DATE.format(at.atZone(zone))

    fun kindLabel(kind: String?): String = res.getString(
        when (kind) {
            "gear" -> R.string.sims_kind_gear
            "talents" -> R.string.sims_kind_talents
            "drops" -> R.string.sims_kind_drops
            "weights" -> R.string.sims_kind_weights
            else -> R.string.sims_kind_run
        },
    )

    private fun instant(text: String?): Instant? = text?.let(ServerRuns::instantOrNull)

    private companion object {
        val RELATIVE_CEILING: Duration = Duration.ofDays(14)
        val SHORT_DATE: DateTimeFormatter = DateTimeFormatter.ofPattern("d MMM", Locale.US)
    }
}
