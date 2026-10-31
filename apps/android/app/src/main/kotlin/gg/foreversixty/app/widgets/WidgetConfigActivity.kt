package gg.foreversixty.app.widgets

import android.appwidget.AppWidgetManager
import android.content.Intent
import android.os.Bundle
import androidx.activity.ComponentActivity
import androidx.activity.compose.setContent
import androidx.compose.foundation.background
import androidx.compose.foundation.clickable
import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.fillMaxSize
import androidx.compose.foundation.layout.fillMaxWidth
import androidx.compose.foundation.layout.heightIn
import androidx.compose.foundation.layout.padding
import androidx.compose.foundation.layout.statusBarsPadding
import androidx.compose.foundation.rememberScrollState
import androidx.compose.foundation.verticalScroll
import androidx.compose.material3.HorizontalDivider
import androidx.compose.material3.Text
import androidx.compose.runtime.Composable
import androidx.compose.runtime.LaunchedEffect
import androidx.compose.runtime.getValue
import androidx.compose.ui.Modifier
import androidx.compose.ui.res.stringResource
import androidx.compose.ui.unit.dp
import androidx.lifecycle.compose.collectAsStateWithLifecycle
import androidx.lifecycle.lifecycleScope
import gg.foreversixty.app.ForeverApplication
import gg.foreversixty.app.R
import gg.foreversixty.app.auth.session.me
import gg.foreversixty.app.designsystem.ForeverTheme
import gg.foreversixty.app.designsystem.FsColor
import gg.foreversixty.app.designsystem.FsSpacing
import gg.foreversixty.app.designsystem.FsType
import gg.foreversixty.app.widgets.data.WidgetStore
import kotlinx.coroutines.launch

/** One choice in the widget's one setting (spec §5.6 Configuration); [key] "" follows the app. */
data class CharacterChoice(val key: String, val name: String)

/**
 * The widgets' configuration (spec §5.6): the character, `Current character` by default. Picking
 * one saves it, refreshes that character's snapshot and places the widget.
 */
class WidgetConfigActivity : ComponentActivity() {
    override fun onCreate(savedInstanceState: Bundle?) {
        super.onCreate(savedInstanceState)
        val appWidgetId = intent?.getIntExtra(AppWidgetManager.EXTRA_APPWIDGET_ID, AppWidgetManager.INVALID_APPWIDGET_ID)
            ?: AppWidgetManager.INVALID_APPWIDGET_ID
        // Backing out leaves no widget behind (the launcher reads CANCELED).
        setResult(RESULT_CANCELED, result(appWidgetId))
        if (appWidgetId == AppWidgetManager.INVALID_APPWIDGET_ID) {
            finish()
            return
        }
        val graph = (application as ForeverApplication).graph
        setContent {
            ForeverTheme {
                val phase by graph.session.phase.collectAsStateWithLifecycle()
                LaunchedEffect(Unit) {
                    graph.session.start()
                    graph.session.refresh()
                }
                val current = CharacterChoice(WidgetStore.CURRENT, stringResource(R.string.widget_currentCharacter))
                val characters = phase.me?.characters.orEmpty().map { CharacterChoice(it.key, it.name) }
                CharacterPicker(listOf(current) + characters) { choice -> pick(appWidgetId, choice) }
            }
        }
    }

    private fun pick(appWidgetId: Int, choice: CharacterChoice) {
        val widgets = (application as ForeverApplication).graph.widgets
        lifecycleScope.launch {
            (widgets?.store ?: WidgetStore(this@WidgetConfigActivity)).setCharacter(appWidgetId, choice.key)
            widgets?.refresher?.refresh()
            setResult(RESULT_OK, result(appWidgetId))
            finish()
        }
    }

    private fun result(appWidgetId: Int) = Intent().putExtra(AppWidgetManager.EXTRA_APPWIDGET_ID, appWidgetId)
}

@Composable
fun CharacterPicker(choices: List<CharacterChoice>, onPick: (CharacterChoice) -> Unit) {
    Column(Modifier.fillMaxSize().background(FsColor.bg).statusBarsPadding().verticalScroll(rememberScrollState()).padding(FsSpacing.space4)) {
        Text(stringResource(R.string.widget_characterSetting), style = FsType.title, color = FsColor.strong)
        choices.forEachIndexed { index, choice ->
            if (index > 0) HorizontalDivider(color = FsColor.lineSoft)
            Text(
                choice.name, style = FsType.bodyStrong, color = FsColor.strong,
                modifier = Modifier.fillMaxWidth().heightIn(min = 56.dp).clickable { onPick(choice) }.padding(vertical = 16.dp),
            )
        }
    }
}
