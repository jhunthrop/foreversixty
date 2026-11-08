package gg.foreversixty.app.push

import android.annotation.SuppressLint
import android.content.Context
import androidx.core.content.edit
import com.google.firebase.FirebaseApp
import com.google.firebase.messaging.FirebaseMessaging
import com.google.firebase.messaging.FirebaseMessagingService
import com.google.firebase.messaging.RemoteMessage
import gg.foreversixty.app.ForeverApplication
import gg.foreversixty.app.widgets.logic.LivePush
import kotlinx.coroutines.CoroutineScope
import kotlinx.coroutines.Dispatchers
import kotlinx.coroutines.SupervisorJob
import kotlinx.coroutines.launch

/**
 * FCM's entry: a registration token is kept and registered, a message is drawn by [PushNotifier].
 * firebase-messaging 26 deprecates onNewToken for onRegistered, which lint does not know yet.
 */
@SuppressLint("MissingFirebaseInstanceTokenRefresh")
class ForeverMessagingService : FirebaseMessagingService() {
    private val scope = CoroutineScope(SupervisorJob() + Dispatchers.IO)

    override fun onRegistered(token: String) {
        FcmToken.save(this, token)
        val registrar = (application as ForeverApplication).graph.push
        scope.launch { registrar.register(token) }
    }

    override fun onUnregistered(token: String) {
        FcmToken.save(this, null)
    }

    override fun onMessageReceived(message: RemoteMessage) {
        // Live surfaces (live.raid.*, live.sim.*, widget.refresh) draw the ongoing notifications and widgets.
        if (LivePush.parse(message.data) != null) {
            val widgets = (application as ForeverApplication).graph.widgets ?: return
            scope.launch { widgets.handlePush(message.data) }
            return
        }
        val parsed = PushMessage.parse(
            data = message.data,
            notificationTitle = message.notification?.title,
            notificationBody = message.notification?.body,
            collapse = message.notification?.tag ?: message.collapseKey,
        ) ?: return
        PushNotifier.show(this, parsed)
    }
}

/**
 * The install's FCM registration token as [ForeverMessagingService.onRegistered] last gave it. A build
 * without Firebase's config (no google-services.json) never has one.
 */
object FcmToken {
    private const val PREFS = "push"
    private const val KEY = "fcm_token"

    fun current(context: Context): String? =
        context.getSharedPreferences(PREFS, Context.MODE_PRIVATE).getString(KEY, null)

    fun save(context: Context, token: String?) {
        context.getSharedPreferences(PREFS, Context.MODE_PRIVATE).edit { putString(KEY, token) }
    }

    /** Asks FCM to register; a new or refreshed token arrives in onRegistered. */
    fun request(context: Context) {
        if (FirebaseApp.getApps(context).isEmpty()) return
        FirebaseMessaging.getInstance().register()
    }
}
