package gg.foreversixty.app.raids.push

import android.app.PendingIntent
import android.content.BroadcastReceiver
import android.content.Context
import android.content.Intent
import android.os.Bundle
import androidx.core.app.NotificationCompat
import gg.foreversixty.app.ForeverApplication
import gg.foreversixty.app.push.PushAction
import gg.foreversixty.app.push.PushActions
import gg.foreversixty.app.push.PushGroup
import gg.foreversixty.app.push.PushNotifier
import kotlinx.coroutines.CoroutineScope
import kotlinx.coroutines.Dispatchers
import kotlinx.coroutines.SupervisorJob
import kotlinx.coroutines.launch
import kotlinx.coroutines.withTimeoutOrNull

/**
 * An officer push's buttons answered from the shade: Accept or Decline (GUILD_APPLICANT), Mention
 * healers (RAID_SHORT). `goAsync`, the POST inside the receiver's window, then the answer replaces the
 * notification by its tag; a tap on the answer opens the application or the raid.
 */
class OfficerActionReceiver : BroadcastReceiver() {
    override fun onReceive(context: Context, intent: Intent) {
        val action = intent.getStringExtra(EXTRA_ACTION) ?: return
        val category = intent.getStringExtra(EXTRA_CATEGORY) ?: return
        val raw = intent.getBundleExtra(EXTRA_DATA)?.let { bundle -> bundle.keySet().mapNotNull { key -> bundle.getString(key)?.let { key to it } }.toMap() } ?: return
        val press = OfficerPress.of(context, category, raw) ?: return
        val tag = intent.getStringExtra(EXTRA_TAG)
        val id = intent.getIntExtra(EXTRA_ID, 0)
        val pending = goAsync()
        CoroutineScope(SupervisorJob() + Dispatchers.IO).launch {
            try {
                val answer = withTimeoutOrNull(BUDGET_MILLIS) { press.answer(action) } ?: press.failed(action)
                val notification = PushNotifier.builder(context, PushGroup.Officer.channelId)
                    .setContentTitle(answer.title)
                    .setContentText(answer.body)
                    .setStyle(NotificationCompat.BigTextStyle().bigText(answer.body))
                    .setContentIntent(PushNotifier.openIntent(context, press.link, tag))
                    .setAutoCancel(true)
                    .build()
                PushNotifier.notify(context, tag, id, notification)
            } finally {
                pending.finish()
            }
        }
    }

    companion object {
        const val EXTRA_ACTION = "gg.foreversixty.officer.ACTION"
        const val EXTRA_CATEGORY = "gg.foreversixty.officer.CATEGORY"
        const val EXTRA_DATA = "gg.foreversixty.officer.DATA"
        const val EXTRA_TAG = "gg.foreversixty.officer.TAG"
        const val EXTRA_ID = "gg.foreversixty.officer.ID"

        /** Inside the receiver's ~10 s window, with room to post the answer. */
        private const val BUDGET_MILLIS = 8_000L
    }
}

/** One officer push's background press: its answer, the failure words, and where the answer's tap goes. */
private class OfficerPress(
    val link: String?,
    private val run: suspend (String) -> RaidPushPayload.Answer,
    private val fail: (String) -> RaidPushPayload.Answer,
) {
    suspend fun answer(action: String) = run(action)

    fun failed(action: String) = fail(action)

    companion object {
        fun of(context: Context, category: String, data: Map<String, String>): OfficerPress? {
            val client = (context.applicationContext as ForeverApplication).graph.apiClient
            return when (category) {
                ApplicantPushCategory.CATEGORY -> ApplicantPushPayload.parse(data)?.let { payload ->
                    val answerer = ApplicantPushAnswerer(LiveApplicantDecider(client), context.resources)
                    OfficerPress(payload.link, { answerer.answer(payload, it) }, { answerer.failed(payload, it) })
                }
                RaidShortPushCategory.CATEGORY -> RaidShortPushPayload.parse(data)?.let { payload ->
                    val answerer = RaidShortPushAnswerer(LiveRaidMentioner(client), context.resources)
                    OfficerPress(payload.link, { answerer.answer(payload, it) }, { answerer.failed(payload) })
                }
                else -> null
            }
        }
    }
}

/** An officer notification's buttons, from the payload's `actions` (titles as the server words them). */
object OfficerNotificationActions {
    /** The categories whose buttons come from the payload and run in [OfficerActionReceiver]. */
    val CATEGORIES = setOf(ApplicantPushCategory.CATEGORY, RaidShortPushCategory.CATEGORY)

    fun actions(context: Context, category: String, data: Map<String, String>, tag: String?, id: Int): List<NotificationCompat.Action> {
        val bundle = Bundle().apply { data.forEach { (key, value) -> putString(key, value) } }
        return answerable(category, PushActions.of(data)).mapIndexed { index, action ->
            val intent = Intent(context, OfficerActionReceiver::class.java)
                .setAction(action.id)
                .putExtra(OfficerActionReceiver.EXTRA_ACTION, action.id)
                .putExtra(OfficerActionReceiver.EXTRA_CATEGORY, category)
                .putExtra(OfficerActionReceiver.EXTRA_DATA, bundle)
                .putExtra(OfficerActionReceiver.EXTRA_TAG, tag)
                .putExtra(OfficerActionReceiver.EXTRA_ID, id)
            val pending = PendingIntent.getBroadcast(
                context, (tag ?: "").hashCode() + index, intent,
                PendingIntent.FLAG_UPDATE_CURRENT or PendingIntent.FLAG_IMMUTABLE,
            )
            NotificationCompat.Action.Builder(0, action.title, pending).build()
        }
    }

    private fun answerable(category: String, actions: List<PushAction>): List<PushAction> = when (category) {
        ApplicantPushCategory.CATEGORY -> ApplicantPushCategory.answerable(actions)
        RaidShortPushCategory.CATEGORY -> RaidShortPushCategory.answerable(actions)
        else -> emptyList()
    }
}
