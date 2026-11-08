package gg.foreversixty.app.firstrun

import android.content.Context
import androidx.activity.compose.BackHandler
import androidx.compose.foundation.background
import androidx.compose.foundation.layout.Box
import androidx.compose.foundation.layout.fillMaxSize
import androidx.compose.foundation.layout.systemBarsPadding
import androidx.compose.runtime.Composable
import androidx.compose.runtime.LaunchedEffect
import androidx.compose.runtime.getValue
import androidx.compose.runtime.remember
import androidx.compose.runtime.rememberCoroutineScope
import androidx.compose.ui.Modifier
import androidx.compose.ui.platform.LocalContext
import androidx.lifecycle.compose.collectAsStateWithLifecycle
import gg.foreversixty.app.auth.session.SessionManager
import gg.foreversixty.app.auth.session.me
import gg.foreversixty.app.designsystem.FsColor
import gg.foreversixty.app.push.rememberNotificationRequest
import gg.foreversixty.app.raids.logic.SystemNotificationPermission
import kotlinx.coroutines.launch

/**
 * The first run over the app (flows spec F8, iOS `FirstRunHost`): one full-screen step at a time, driven
 * by [FirstRunModel] from the session's account; with no step, [content] (the tabs).
 */
@Composable
fun FirstRunHost(session: SessionManager, content: @Composable () -> Unit) {
    val context = LocalContext.current
    val model = remember(session) { FirstRunModel(dependencies(context, session)) }
    val phase by session.phase.collectAsStateWithLifecycle()
    val me = phase.me
    LaunchedEffect(me) { model.update(me) }
    val scope = rememberCoroutineScope()
    val request = rememberNotificationRequest()
    when (val step = model.step) {
        null -> content()
        FirstRunModel.Step.PickMain -> Step {
            // Pick your main cannot be skipped: Back stays on it.
            BackHandler {}
            PickMainScreen(
                characters = me?.let(FirstRunModel::candidates).orEmpty(),
                selection = model.selection,
                saving = model.saving,
                failed = model.failed,
                onSelect = { model.selection = it },
                onContinue = { scope.launch { model.continueWithMain(session.phase.value.me) } },
            )
        }
        is FirstRunModel.Step.Permission -> Step {
            BackHandler { model.notNow() }
            FirstPermissionScreen(step.guild, onTurnOn = { model.turnOnNotifications(request) }, onNotNow = model::notNow)
        }
    }
}

@Composable
private fun Step(content: @Composable () -> Unit) {
    Box(Modifier.fillMaxSize().background(FsColor.bg).systemBarsPadding()) { content() }
}

private fun dependencies(context: Context, session: SessionManager) = FirstRunModel.Dependencies(
    setMain = session::setMain,
    permission = SystemNotificationPermission(context),
    asked = PrefsFirstRunAskStore(context.getSharedPreferences(PrefsFirstRunAskStore.FILE, Context.MODE_PRIVATE)),
)
