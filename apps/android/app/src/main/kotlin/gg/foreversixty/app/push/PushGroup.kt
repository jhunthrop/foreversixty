package gg.foreversixty.app.push

import android.app.NotificationChannel
import android.app.NotificationManager
import android.content.Context
import androidx.annotation.StringRes
import gg.foreversixty.app.R

/**
 * The notification groups (spec 2026-11-08 F4 and 4.4: the Activity filters and the notification
 * settings), one Android channel each. [wire] is the API's `events.Group`; the raids channel keeps
 * the id lane F4 created, so a player's existing channel setting survives.
 */
enum class PushGroup(
    val wire: String,
    val channelId: String,
    @param:StringRes val title: Int,
    @param:StringRes val line: Int,
    val importance: Int,
) {
    Raids("raids", "raids", R.string.push_group_raids, R.string.push_group_raids_line, NotificationManager.IMPORTANCE_HIGH),
    Gear("gear", "gear", R.string.push_group_gear, R.string.push_group_gear_line, NotificationManager.IMPORTANCE_DEFAULT),
    Logs("logs", "logs", R.string.push_group_logs, R.string.push_group_logs_line, NotificationManager.IMPORTANCE_DEFAULT),
    Ladder("ladder", "ladder", R.string.push_group_ladder, R.string.push_group_ladder_line, NotificationManager.IMPORTANCE_DEFAULT),
    Tiers("tiers", "tiers", R.string.push_group_tiers, R.string.push_group_tiers_line, NotificationManager.IMPORTANCE_LOW),
    Officer("officer", "officer", R.string.push_group_officer, R.string.push_group_officer_line, NotificationManager.IMPORTANCE_HIGH),
    Account("account", "account", R.string.push_group_account, R.string.push_group_account_line, NotificationManager.IMPORTANCE_DEFAULT),
    ;

    companion object {
        /**
         * The message's group: its `group` field, else the kind's group. The kind table mirrors
         * `api/internal/events/groups.go` only until every send carries `group`; an unknown kind reads
         * as account there too.
         */
        fun of(group: String?, kind: String?): PushGroup =
            entries.firstOrNull { it.wire == group } ?: kind?.let(::ofKind) ?: Account

        private fun ofKind(kind: String): PushGroup? = when {
            kind in OFFICER_KINDS -> Officer
            kind in LOG_KINDS -> Logs
            kind in GEAR_KINDS -> Gear
            kind.startsWith("raid.") -> Raids
            kind.startsWith("guildsim.") -> Ladder
            kind.startsWith("tier.") -> Tiers
            else -> null
        }

        private val OFFICER_KINDS = setOf("raid.short", "raid.late_change", "roster.pending", "guild.applicant")
        private val LOG_KINDS = setOf("report.ready", "coach.ready", "ranking.best", "raid.summary")
        private val GEAR_KINDS = setOf("loot.awarded", "bis.changed", "band.entered", "game.milestone", "sim.finished")

        /** Creates every channel; idempotent, so it runs before each post and at launch. */
        fun createChannels(context: Context) {
            val manager = context.getSystemService(NotificationManager::class.java) ?: return
            manager.createNotificationChannels(
                entries.map { group ->
                    NotificationChannel(group.channelId, context.getString(group.title), group.importance).apply {
                        description = context.getString(group.line)
                    }
                },
            )
        }
    }
}
