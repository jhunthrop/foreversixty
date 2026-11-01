package gg.foreversixty.app.widgets.logic

import android.content.res.Resources
import gg.foreversixty.app.R
import java.time.Duration
import java.time.Instant
import java.util.Locale

/**
 * The ongoing raid notification's words (spec §5.6, board F14.3), all from [RaidFigures]: never a
 * second source. The guild names the card's header; [cancelled] replaces the countdown with
 * `Cancelled` for 5 minutes.
 */
data class RaidLiveContent(
    val title: String,
    val text: String,
    val guild: String,
    val bigText: String,
    val countdownTo: Instant?,
    val guildId: Long,
    val raidId: Long,
    val description: String,
    val cancelled: Boolean = false,
) {
    companion object {
        /** Null outside the countdown's window: nothing to show, so the notification goes. */
        fun of(figures: RaidFigures, res: Resources): RaidLiveContent? {
            val guildId = figures.guildId ?: return null
            val raidId = figures.raidId ?: return null
            val guild = figures.guildName.orEmpty()
            if (figures.liveCancelled) {
                val cancelled = res.getString(R.string.widget_cancelled)
                return RaidLiveContent(
                    title = figures.title, text = cancelled, guild = guild, bigText = cancelled, countdownTo = null,
                    guildId = guildId, raidId = raidId, description = "${figures.title}. $cancelled", cancelled = true,
                )
            }
            if (!figures.liveActive) return null
            val title = figures.liveTitle ?: return null
            val you = figures.liveYou.orEmpty()
            val signedUp = figures.size?.let { res.getString(R.string.widget_signedUp, figures.signedUpCount, it) } ?: figures.signedUp
            return RaidLiveContent(
                title = title, text = you, guild = guild,
                bigText = listOfNotNull(figures.liveUntil, you, signedUp).joinToString("\n"),
                countdownTo = figures.liveTarget, guildId = guildId, raidId = raidId,
                description = listOf(title, you).joinToString(". "),
            )
        }
    }
}

/** The ongoing server-sim notification's words and bar (spec §5.6 Server sim). */
data class SimLiveContent(
    val title: String,
    val text: String,
    val subText: String?,
    /** done of max; max 0 with [indeterminate] for a plain run, null bar once it ended. */
    val progress: Pair<Int, Int>?,
    val indeterminate: Boolean,
    val ongoing: Boolean,
) {
    companion object {
        fun of(sim: SimLive, res: Resources): SimLiveContent = when (sim.state) {
            SimLive.State.Running -> running(sim, res)
            SimLive.State.Finished -> SimLiveContent(
                title = sim.title,
                text = res.getString(R.string.live_finishedSub, sim.dps?.let { String.format(Locale.US, "%,.0f", it) } ?: "—"),
                subText = res.getString(R.string.live_finished), progress = null, indeterminate = false, ongoing = false,
            )
            SimLive.State.Failed -> SimLiveContent(
                title = sim.title, text = res.getString(R.string.live_openRun), subText = res.getString(R.string.live_failed),
                progress = null, indeterminate = false, ongoing = false,
            )
        }

        private fun running(sim: SimLive, res: Resources): SimLiveContent {
            val total = sim.total
            val parts = buildList {
                add(total?.let { res.getString(R.string.live_sim, sim.done, it) } ?: res.getString(R.string.sims_running))
                sim.etaSeconds?.let { seconds ->
                    val minutes = WidgetFigures.ceilMinutes(Duration.ofSeconds(seconds.toLong())).toInt().coerceAtLeast(1)
                    add(res.getQuantityString(R.plurals.live_simLeft, minutes, minutes))
                }
            }
            return SimLiveContent(
                title = sim.title, text = parts.joinToString(" · "), subText = null,
                progress = if (total != null) sim.done to total else 0 to 0, indeterminate = total == null, ongoing = true,
            )
        }
    }
}

/** Updates a server run's notification at most every 15 s (spec §5.6); an ending always passes. */
class SimThrottle(private val every: Duration = Duration.ofSeconds(15)) {
    private val last = mutableMapOf<String, Instant>()

    fun allows(sim: SimLive, now: Instant): Boolean {
        if (sim.state != SimLive.State.Running) return true.also { last.remove(sim.simId) }
        val previous = last[sim.simId]
        if (previous != null && Duration.between(previous, now) < every) return false
        last[sim.simId] = now
        return true
    }
}
