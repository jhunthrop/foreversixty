package gg.foreversixty.app.raids.push

import android.app.PendingIntent
import android.content.BroadcastReceiver
import android.content.Context
import android.content.Intent
import android.os.Bundle
import androidx.core.app.NotificationCompat
import gg.foreversixty.app.ForeverApplication
import gg.foreversixty.app.R
import gg.foreversixty.app.navigation.DeepLink
import gg.foreversixty.app.push.PushGroup
import gg.foreversixty.app.push.PushMessage
import gg.foreversixty.app.push.PushNotifier
import gg.foreversixty.app.raids.logic.LiveRaidService
import gg.foreversixty.app.raids.logic.SignupState
import gg.foreversixty.app.widgets.WidgetRefreshWork
import kotlinx.coroutines.CoroutineScope
import kotlinx.coroutines.Dispatchers
import kotlinx.coroutines.SupervisorJob
import kotlinx.coroutines.launch
import kotlinx.coroutines.withTimeoutOrNull

/**
 * In, Late, Bench and Out answered from the notification shade (spec 2026-11-08 F1): `goAsync`, the
 * PUT inside the receiver's window, then the answer replaces the notification by its tag. After Late the
 * answer offers How late? (15, 30, 60 min); a press there answers again with the minutes. A foreground
 * action (Choose a spec, Choose a character) is an activity PendingIntent on the raid page, never this
 * receiver (Android 12 blocks trampolines); a press the server cannot place replaces the notification
 * with one whose tap opens the raid page.
 */
class RaidActionReceiver : BroadcastReceiver() {
    override fun onReceive(context: Context, intent: Intent) {
        val action = intent.getStringExtra(EXTRA_ACTION) ?: return
        val raw = intent.getBundleExtra(EXTRA_DATA)?.toStringMap() ?: return
        val payload = RaidPushPayload.parse(raw) ?: return
        val tag = intent.getStringExtra(EXTRA_TAG)
        val id = intent.getIntExtra(EXTRA_ID, 0)
        val state = RaidPushCategory.stateFor(action) ?: return
        val minutes = RaidPushCategory.lateMinutes(action)
        val pending = goAsync()
        val app = context.applicationContext as ForeverApplication
        val answerer = RaidPushAnswerer(LiveRaidService(app.graph.apiClient), context.resources)
        CoroutineScope(SupervisorJob() + Dispatchers.IO).launch {
            try {
                val outcome = withTimeoutOrNull(BUDGET_MILLIS) { answerer.answer(payload, state, minutes) }
                    ?: RaidPushOutcome.Answered(payload.answer(RaidPushPayload.FAILED, context.resources))
                when (outcome) {
                    is RaidPushOutcome.Answered -> {
                        val askMinutes = state == SignupState.Late && minutes == null
                        post(context, outcome.answer, payload, raw, tag, id, askMinutes)
                        // Lane F14: the widgets, tile and countdown redraw from the server's answer.
                        WidgetRefreshWork.now(context)
                    }
                    is RaidPushOutcome.Open -> postOpen(context, outcome.link, raw, tag, id)
                }
            } finally {
                pending.finish()
            }
        }
    }

    /** The answer replaces the question: the same tag and id; after a plain Late, the minutes. */
    @Suppress("LongParameterList")
    private fun post(
        context: Context,
        answer: RaidPushPayload.Answer,
        payload: RaidPushPayload,
        raw: Map<String, String>,
        tag: String?,
        id: Int,
        askMinutes: Boolean,
    ) {
        val body = if (askMinutes) listOf(answer.body, context.getString(R.string.push_late_question)).filter { it.isNotEmpty() }.joinToString("\n") else answer.body
        val builder = PushNotifier.builder(context, PushGroup.Raids.channelId)
            .setContentTitle(answer.title)
            .setContentText(body)
            .setStyle(NotificationCompat.BigTextStyle().bigText(body))
            .setContentIntent(PushNotifier.openIntent(context, payload.deepLink().url(), tag))
            .setAutoCancel(true)
        if (askMinutes) {
            RaidNotificationActions.lateMinutes(context, raw, tag, id).forEach(builder::addAction)
        }
        PushNotifier.notify(context, tag, id, builder.build())
    }

    /**
     * A press the shade cannot finish (the spec to pick, the character to choose): the notification
     * asks for it and its tap opens the raid page there. A receiver may not start the activity itself.
     */
    private fun postOpen(context: Context, link: DeepLink, raw: Map<String, String>, tag: String?, id: Int) {
        val asking = (link as? DeepLink.Raid)?.spec == RaidPushPayload.SPEC_ASK
        val title = context.getString(if (asking) R.string.push_action_askSpec else R.string.push_action_choose)
        val body = raw[PushMessage.KEY_TITLE].orEmpty()
        val notification = PushNotifier.builder(context, PushGroup.Raids.channelId)
            .setContentTitle(title)
            .setContentText(body)
            .setContentIntent(PushNotifier.openIntent(context, link.url(), tag))
            .setAutoCancel(true)
            .build()
        PushNotifier.notify(context, tag, id, notification)
    }

    companion object {
        const val EXTRA_ACTION = "gg.foreversixty.raids.ACTION"
        const val EXTRA_DATA = "gg.foreversixty.raids.DATA"
        const val EXTRA_TAG = "gg.foreversixty.raids.TAG"
        const val EXTRA_ID = "gg.foreversixty.raids.ID"

        /** Inside the receiver's ~10 s window, with room to post the answer. */
        private const val BUDGET_MILLIS = 8_000L

        private fun Bundle.toStringMap(): Map<String, String> =
            keySet().mapNotNull { key -> getString(key)?.let { key to it } }.toMap()
    }
}

/** The buttons of a raid notification, in the order the spec draws them. */
object RaidNotificationActions {
    /** The Raids channel, created with the other groups ([PushGroup.createChannels]). */
    val CHANNEL_RAIDS: String get() = PushGroup.Raids.channelId

    fun actions(context: Context, category: String, data: Map<String, String>, tag: String?, id: Int): List<NotificationCompat.Action> {
        val ids = when (category) {
            // Raid night's countdown carries the same four answers (spec F2, mock m-android-lock).
            RaidPushCategory.SIGNUP, RaidPushCategory.LIVE -> listOf(
                RaidPushCategory.ACTION_IN to context.getString(R.string.push_action_in),
                RaidPushCategory.ACTION_LATE to context.getString(R.string.push_action_late),
                RaidPushCategory.ACTION_BENCH to context.getString(R.string.push_action_bench),
                RaidPushCategory.ACTION_OUT to context.getString(R.string.push_action_out),
            )
            RaidPushCategory.ASK_SPEC -> listOf(
                RaidPushCategory.ACTION_ASK_SPEC to context.getString(R.string.push_action_askSpec),
                RaidPushCategory.ACTION_OUT to context.getString(R.string.push_action_out),
            )
            RaidPushCategory.CHOOSE -> listOf(RaidPushCategory.ACTION_CHOOSE to context.getString(R.string.push_action_choose))
            else -> emptyList()
        }
        return build(context, ids, data, tag, id, slot = 0)
    }

    /** How late? 15, 30 and 60 min, under the Late answer. */
    fun lateMinutes(context: Context, data: Map<String, String>, tag: String?, id: Int): List<NotificationCompat.Action> {
        val ids = RaidPushCategory.LATE_MINUTES.map { minutes ->
            RaidPushCategory.ACTION_LATE_MINUTES + minutes to context.getString(R.string.push_late_minutes, minutes)
        }
        return build(context, ids, data, tag, id, slot = MINUTES_SLOT)
    }

    @Suppress("LongParameterList")
    private fun build(
        context: Context,
        ids: List<Pair<String, String>>,
        data: Map<String, String>,
        tag: String?,
        id: Int,
        slot: Int,
    ): List<NotificationCompat.Action> {
        val bundle = Bundle().apply { data.forEach { (key, value) -> putString(key, value) } }
        val payload = RaidPushPayload.parse(data)
        return ids.mapIndexed { index, (action, label) ->
            if (RaidPushCategory.stateFor(action) == null) {
                // Opens the raid page: an activity PendingIntent straight from the button (no trampoline).
                val link = payload?.deepLink(askingSpec = action == RaidPushCategory.ACTION_ASK_SPEC)?.url()
                val open = PushNotifier.openIntent(context, link, tag, slot = OPEN_SLOT + slot + index)
                return@mapIndexed NotificationCompat.Action.Builder(0, label, open).build()
            }
            val intent = Intent(context, RaidActionReceiver::class.java)
                .setAction(action)
                .putExtra(RaidActionReceiver.EXTRA_ACTION, action)
                .putExtra(RaidActionReceiver.EXTRA_DATA, bundle)
                .putExtra(RaidActionReceiver.EXTRA_TAG, tag)
                .putExtra(RaidActionReceiver.EXTRA_ID, id)
            val pending = PendingIntent.getBroadcast(
                context, (tag ?: "").hashCode() * ACTION_SLOTS + slot + index, intent,
                PendingIntent.FLAG_UPDATE_CURRENT or PendingIntent.FLAG_IMMUTABLE,
            )
            NotificationCompat.Action.Builder(0, label, pending).build()
        }
    }

    /** Room for the four answers, then the three minutes. */
    private const val ACTION_SLOTS = 8
    private const val MINUTES_SLOT = 4

    /** Keeps an opening button's request code apart from the notification's own tap. */
    private const val OPEN_SLOT = 1
}
