package gg.foreversixty.app.widgets.glance

import android.content.Context
import android.graphics.Bitmap
import androidx.compose.runtime.collectAsState
import androidx.glance.GlanceId
import androidx.glance.LocalSize
import androidx.glance.action.ActionParameters
import androidx.glance.appwidget.GlanceAppWidget
import androidx.glance.appwidget.GlanceAppWidgetManager
import androidx.glance.appwidget.GlanceAppWidgetReceiver
import androidx.glance.appwidget.SizeMode
import androidx.glance.appwidget.action.ActionCallback
import androidx.glance.appwidget.provideContent
import gg.foreversixty.app.ForeverApplication
import gg.foreversixty.app.raids.logic.SignupState
import gg.foreversixty.app.widgets.WidgetRefreshWork
import gg.foreversixty.app.widgets.WidgetsFeature
import gg.foreversixty.app.widgets.data.WidgetStore
import gg.foreversixty.app.widgets.logic.WidgetFigures
import gg.foreversixty.app.widgets.logic.WidgetSnapshot
import java.time.Instant

private fun Context.widgets(): WidgetsFeature? = (applicationContext as ForeverApplication).graph.widgets

/** Which character a placed widget shows, and its stored snapshot. */
private suspend fun WidgetStore.forWidget(context: Context, id: GlanceId): Pair<String, WidgetSnapshot?> {
    val appWidgetId = GlanceAppWidgetManager(context).getAppWidgetId(id)
    val character = characterFor(appWidgetId)
    return character to snapshot(character)
}

/** Next raid (spec §5.6): 2 x 2 and 4 x 2, home screen and, where the OS allows, the lock screen. */
class NextRaidWidget : GlanceAppWidget() {
    override val sizeMode: SizeMode = SizeMode.Responsive(setOf(WidgetDesign.small, WidgetDesign.medium))

    override suspend fun provideGlance(context: Context, id: GlanceId) {
        val store = context.widgets()?.store ?: WidgetStore(context)
        val (character, first) = store.forWidget(context, id)
        val initial = first to store.overlay()
        val figures = WidgetFigures(context.resources)
        provideContent {
            // A running session redraws as the store changes; each composition reads the clock afresh.
            val (snapshot, overlay) = store.watch(character).collectAsState(initial).value
            val medium = LocalSize.current.width >= WidgetDesign.medium.width
            RaidWidgetContent(figures.raid(snapshot, Instant.now(), overlay), medium, snapshot?.widget?.character?.key)
        }
    }
}

/** Next upgrade (spec §5.6): 2 x 2 and 4 x 2. */
class NextUpgradeWidget : GlanceAppWidget() {
    override val sizeMode: SizeMode = SizeMode.Responsive(setOf(WidgetDesign.small, WidgetDesign.medium))

    override suspend fun provideGlance(context: Context, id: GlanceId) {
        val store = context.widgets()?.store ?: WidgetStore(context)
        val (_, snapshot) = store.forWidget(context, id)
        val figures = WidgetFigures(context.resources).upgrades(snapshot)
        val build = (context.applicationContext as ForeverApplication).graph.raids.iconBuild()
        val icons: Map<Int, Bitmap> = figures.rows.associate { row ->
            row.itemId to WidgetIcons.framed(WidgetIcons.load(context, row.icon, build), row.quality)
        }
        provideContent {
            UpgradeWidgetContent(figures, LocalSize.current.width >= WidgetDesign.medium.width, icons)
        }
    }
}

/** Placing the first widget starts the 30-minute refresh; removing the last leaves it to sign-out. */
abstract class ForeverWidgetReceiver : GlanceAppWidgetReceiver() {
    override fun onEnabled(context: Context) {
        super.onEnabled(context)
        WidgetRefreshWork.periodic(context)
        WidgetRefreshWork.now(context)
    }
}

class NextRaidWidgetReceiver : ForeverWidgetReceiver() {
    override val glanceAppWidget: GlanceAppWidget = NextRaidWidget()
}

class NextUpgradeWidgetReceiver : ForeverWidgetReceiver() {
    override val glanceAppWidget: GlanceAppWidget = NextUpgradeWidget()
}

/** In, Late or Out on the 4 x 2 widget: answers without opening the app (spec §5.6). */
class SignupCallback : ActionCallback {
    override suspend fun onAction(context: Context, glanceId: GlanceId, parameters: ActionParameters) {
        val feature = context.widgets() ?: return
        val state = SignupState.of(parameters[STATE]) ?: return
        val guild = parameters[GUILD] ?: return
        val raid = parameters[RAID] ?: return
        val character = parameters[CHARACTER] ?: return
        feature.signup.answer(guild, raid, character, state)
    }

    companion object {
        val GUILD = ActionParameters.Key<Long>("guild")
        val RAID = ActionParameters.Key<Long>("raid")
        val CHARACTER = ActionParameters.Key<String>("character")
        val STATE = ActionParameters.Key<String>("state")
    }
}
