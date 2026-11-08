package gg.foreversixty.app.push

import android.Manifest
import android.os.Build
import androidx.activity.compose.rememberLauncherForActivityResult
import androidx.activity.result.contract.ActivityResultContracts
import androidx.compose.runtime.Composable
import androidx.compose.runtime.remember

/**
 * The system's Allow for POST_NOTIFICATIONS (Android 13+; 9 to 12 need no ask, spec F4.4). Every ask in
 * the app goes through this one: the raid reminder and the first run's Turn on notifications. The FCM
 * token is registered by [PushRegistrar] whatever the answer.
 */
@Composable
fun rememberNotificationRequest(): () -> Unit {
    val launcher = rememberLauncherForActivityResult(ActivityResultContracts.RequestPermission()) { }
    return remember(launcher) {
        {
            if (Build.VERSION.SDK_INT >= Build.VERSION_CODES.TIRAMISU) launcher.launch(Manifest.permission.POST_NOTIFICATIONS)
        }
    }
}
