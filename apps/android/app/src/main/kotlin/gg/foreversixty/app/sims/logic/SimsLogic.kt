package gg.foreversixty.app.sims.logic

import gg.foreversixty.app.model.SimProgress
import gg.foreversixty.app.model.SimRow
import gg.foreversixty.app.network.ApiError
import java.time.Instant
import kotlin.math.roundToInt
import kotlin.time.Duration.Companion.seconds

/** The seven states of a Server runs row, as the site's `rowState` (web/src/lib/sim/server-runs.ts). */
enum class ServerRunState {
    Queued, Running, RunningOverdue, RunningNoEstimate, Done, DoneUnseen, Failed;

    val isInFlight: Boolean get() = this == Queued || this == Running || this == RunningOverdue || this == RunningNoEstimate
    val isDone: Boolean get() = this == Done || this == DoneUnseen
}

/** A run's usual length in the one unit the state line names (`about 4 minutes`). */
sealed interface RunEstimate {
    data object AboutAMinute : RunEstimate
    data class Minutes(val n: Int) : RunEstimate
    data class Hours(val n: Int) : RunEstimate

    companion object {
        private const val SECONDS_PER_MINUTE = 60.0
        private const val MINUTES_PER_HOUR = 60.0

        /** From 90 minutes the estimate reads in hours, as on the site. */
        private const val HOURS_FROM_MINUTES = 90

        fun of(seconds: Int): RunEstimate {
            val minutes = (seconds / SECONDS_PER_MINUTE).roundToInt()
            return when {
                minutes <= 1 -> AboutAMinute
                minutes < HOURS_FROM_MINUTES -> Minutes(minutes)
                else -> Hours((minutes / MINUTES_PER_HOUR).roundToInt())
            }
        }
    }
}

/** The rules under Sims (spec F5): pure data in, data out. Every sentence is in `SimsCopy`. */
object ServerRuns {
    /** Server runs shows the newest five; older ones are in Your sims (the site's `MAX_VISIBLE_RUNS`). */
    const val MAX_VISIBLE = 5

    /** A run is overdue once it has taken more than this many times its estimate. */
    private const val OVERDUE_FACTOR = 1.5

    /** How often an in-flight row asks `GET /v1/sims/{id}/progress` while the screen shows. */
    val POLL_INTERVAL = 5.seconds

    const val SERVER_LANE = "server"

    fun state(row: SimRow, now: Instant): ServerRunState = when (row.state) {
        "queued" -> ServerRunState.Queued
        "running" -> {
            val estimate = row.estimateSeconds
            when {
                estimate == null -> ServerRunState.RunningNoEstimate
                isOverdue(estimate, elapsedSeconds(row, now)) -> ServerRunState.RunningOverdue
                else -> ServerRunState.Running
            }
        }
        "error" -> ServerRunState.Failed
        else -> if (row.seen == false) ServerRunState.DoneUnseen else ServerRunState.Done
    }

    fun isOverdue(estimate: Int, elapsed: Long): Boolean = elapsed > OVERDUE_FACTOR * estimate

    /** Whole seconds since the run was created, never negative. */
    fun elapsedSeconds(row: SimRow, now: Instant): Long {
        val created = row.createdAt?.let(::instantOrNull) ?: return 0
        return maxOf(0, now.epochSecond - created.epochSecond)
    }

    /** The newest five by creation time, and whether older rows were left out. */
    fun visible(rows: List<SimRow>): Pair<List<SimRow>, Boolean> {
        val newestFirst = rows.sortedByDescending { it.createdAt?.let(::instantOrNull) ?: Instant.MIN }
        return newestFirst.take(MAX_VISIBLE) to (newestFirst.size > MAX_VISIBLE)
    }

    /** The section shows for a Premium account, or one that still has server rows (a lapsed plan). */
    fun showsSection(entitled: Boolean, serverRows: List<SimRow>): Boolean = entitled || serverRows.isNotEmpty()

    /** The rows the poll asks about: queued or running. */
    fun inFlightIds(rows: List<SimRow>, now: Instant): List<String> =
        rows.filter { state(it, now).isInFlight }.mapNotNull { it.simId }

    /** The bar's fill, only when the run counts combinations (a bulk run); a plain run has no total. */
    fun fraction(progress: SimProgress?): Float? {
        if (progress == null || progress.combosTotal <= 0) return null
        return minOf(1f, progress.combosDone.toFloat() / progress.combosTotal)
    }

    /** A finished poll answer: the list is read again for its headline and DPS. */
    fun hasEnded(progress: SimProgress): Boolean = progress.state == "done" || progress.state == "error"

    fun instantOrNull(text: String): Instant? = runCatching { Instant.parse(text) }.getOrNull()
}

/** What a failed row's `Run again` shows (the site's `RerunStatus`). */
enum class RerunStatus {
    Idle, Pending, Started, Failed, Lapsed;

    /** The button cannot be pressed again while it starts or once it started. */
    val isLocked: Boolean get() = this == Pending || this == Started

    companion object {
        private const val PAYMENT_REQUIRED = 402

        /** A refused rerun: 402 is a lapsed plan, anything else could not start. */
        fun refused(error: ApiError): RerunStatus =
            if (error is ApiError.Server && error.status == PAYMENT_REQUIRED) Lapsed else Failed
    }
}
