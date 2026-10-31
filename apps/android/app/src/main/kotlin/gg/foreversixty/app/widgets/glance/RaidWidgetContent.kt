package gg.foreversixty.app.widgets.glance

import android.content.Context
import android.content.Intent
import androidx.compose.runtime.Composable
import androidx.compose.ui.unit.dp
import androidx.core.net.toUri
import androidx.glance.GlanceModifier
import androidx.glance.Image
import androidx.glance.ImageProvider
import androidx.glance.LocalContext
import androidx.glance.action.Action
import androidx.glance.action.actionParametersOf
import androidx.glance.appwidget.action.actionStartActivity
import androidx.glance.action.clickable
import androidx.glance.appwidget.action.actionRunCallback
import androidx.glance.background
import androidx.glance.layout.Alignment
import androidx.glance.layout.Box
import androidx.glance.layout.Column
import androidx.glance.layout.Row
import androidx.glance.layout.Spacer
import androidx.glance.layout.fillMaxWidth
import androidx.glance.layout.fillMaxSize
import androidx.glance.layout.height
import androidx.glance.layout.padding
import androidx.glance.layout.size
import androidx.glance.layout.width
import androidx.glance.semantics.contentDescription
import androidx.glance.semantics.semantics
import androidx.glance.text.Text
import gg.foreversixty.app.MainActivity
import gg.foreversixty.app.R
import gg.foreversixty.app.navigation.DeepLink
import gg.foreversixty.app.raids.logic.SignupState
import gg.foreversixty.app.widgets.logic.RaidFigures
import gg.foreversixty.app.widgets.logic.RaidSurfaceState
import gg.foreversixty.app.widgets.logic.TagKind

/** A widget tap opens the app on the right screen (spec §5.6); null opens the app where it was. */
internal fun openApp(context: Context, url: String? = null): Action {
    val intent = Intent(context, MainActivity::class.java).addFlags(Intent.FLAG_ACTIVITY_NEW_TASK or Intent.FLAG_ACTIVITY_CLEAR_TOP)
    url?.let { intent.setAction(Intent.ACTION_VIEW).setData(it.toUri()) }
    return actionStartActivity(intent)
}

/**
 * Next raid at 2 x 2 ([medium] false) and 4 x 2 (spec §5.6, boards F14.1, F14.2). Every state has
 * words: no spinner, no empty box. The whole widget is one TalkBack element; each button its own.
 */
@Composable
fun RaidWidgetContent(figures: RaidFigures, medium: Boolean, characterKey: String?) {
    val context = LocalContext.current
    val raidLink = figures.guildId?.let { g -> figures.raidId?.let { r -> DeepLink.Raid(g, r).url() } }
    Box(
        GlanceModifier.fillMaxSize().background(ImageProvider(R.drawable.widget_card)).padding(WidgetDesign.padding)
            .clickable(openApp(context, raidLink)),
    ) {
        if (figures.state != RaidSurfaceState.Raid) {
            WordyRaid(figures)
        } else if (medium) {
            MediumRaid(figures, characterKey)
        } else {
            Column(GlanceModifier.fillMaxSize().semantics { contentDescription = figures.description }) {
                RaidHead(figures)
                Spacer(GlanceModifier.defaultWeight())
                Foot(figures)
            }
        }
    }
}

/**
 * The 4 x 2: the count beside the label, the zone, the time and countdown on one line, then the
 * foot (tag or `Updated`) and In, Late, Out. Every line has the full width, so the zone and the
 * countdown stay whole up to the widget's type cap.
 */
@Composable
private fun MediumRaid(figures: RaidFigures, characterKey: String?) {
    val context = LocalContext.current
    Column(GlanceModifier.fillMaxSize()) {
        Column(GlanceModifier.fillMaxWidth().semantics { contentDescription = figures.description }) {
            Row(GlanceModifier.fillMaxWidth(), verticalAlignment = Alignment.CenterVertically) {
                Label(context.getString(R.string.widget_nextRaid), GlanceModifier.defaultWeight())
                figures.signedUp?.let { Text(it, style = widgetStyles().count, maxLines = 1) }
                Spacer(GlanceModifier.width(4.dp))
                Text(context.getString(R.string.widget_signedUpLabel), style = widgetStyles().muted, maxLines = 1)
            }
            Spacer(GlanceModifier.height(6.dp))
            Text(figures.title.uppercase(), style = widgetStyles().zone, maxLines = 1)
            Row(verticalAlignment = Alignment.CenterVertically) {
                figures.time?.let { Text(it, style = widgetStyles().muted, maxLines = 1) }
                figures.countdown?.let {
                    Spacer(GlanceModifier.width(8.dp))
                    Text(it, style = widgetStyles().countdown, maxLines = 1)
                }
            }
        }
        Spacer(GlanceModifier.defaultWeight())
        Row(GlanceModifier.fillMaxWidth(), verticalAlignment = Alignment.Bottom) {
            // The answer in words whenever no button says it: locked, cancelled, not signed up, choose a spec.
            val unanswered = figures.tagKind == TagKind.NotSignedUp || figures.tagKind == TagKind.ChooseSpec
            Box(GlanceModifier.defaultWeight()) { Foot(figures, showsTag = !figures.buttons || unanswered) }
            if (figures.buttons) Answers(figures, characterKey)
        }
    }
}

@Composable
private fun Label(text: String, modifier: GlanceModifier = GlanceModifier) {
    Row(modifier, verticalAlignment = Alignment.CenterVertically) {
        Image(ImageProvider(R.drawable.brand_seal), contentDescription = null, modifier = GlanceModifier.size(WidgetDesign.sealSize))
        Spacer(GlanceModifier.width(6.dp))
        Text(WidgetDesign.tracked(text), style = widgetStyles().micro, maxLines = 1)
    }
}

@Composable
private fun RaidHead(figures: RaidFigures) {
    val context = LocalContext.current
    Label(context.getString(R.string.widget_nextRaid))
    Spacer(GlanceModifier.height(6.dp))
    // One line: the time and the countdown below it never give way to a wrapped zone.
    Text(figures.title.uppercase(), style = widgetStyles().zone, maxLines = 1)
    figures.time?.let { Text(it, style = widgetStyles().muted, maxLines = 1) }
    figures.countdown?.let { Text(it, style = widgetStyles().countdown, maxLines = 1) }
}

/** The foot: the answer tag, and on old data `Updated 14 h ago` under it (the widget stays whole). */
@Composable
private fun Foot(figures: RaidFigures, showsTag: Boolean = true) {
    Column {
        val tag = figures.tag
        val kind = figures.tagKind
        if (showsTag && tag != null && kind != null) {
            Box(GlanceModifier.background(ImageProvider(WidgetDesign.tagBackground(kind))).padding(horizontal = 8.dp, vertical = 3.dp)) {
                Text(WidgetDesign.tracked(tag), style = widgetStyles().tag(kind), maxLines = 1)
            }
        }
        figures.stale?.let {
            if (showsTag && tag != null) Spacer(GlanceModifier.height(4.dp))
            Text(it, style = widgetStyles().muted, maxLines = 1)
        }
    }
}

/** In, Late, Out: 48 dp, the answer checked; a press answers in place, or asks the spec first. */
@Composable
private fun Answers(figures: RaidFigures, characterKey: String?) {
    val context = LocalContext.current
    Row {
        listOf(SignupState.In, SignupState.Late, SignupState.Out).forEachIndexed { index, state ->
            if (index > 0) Spacer(GlanceModifier.width(6.dp))
            val checked = figures.checked == state
            val failed = figures.notSaved == state
            val label = if (failed) context.getString(R.string.widget_notSaved) else context.getString(state.label)
            Box(
                GlanceModifier.height(WidgetDesign.buttonHeight)
                    // `Not saved` takes the room its words need; the answers keep a 52 dp target.
                    .let { if (failed) it.padding(horizontal = 14.dp) else it.width(52.dp) }
                    .background(ImageProvider(if (checked) R.drawable.widget_button_checked else R.drawable.widget_button))
                    .clickable(answerAction(context, figures, characterKey, state))
                    .semantics { contentDescription = label },
                contentAlignment = Alignment.Center,
            ) { Text(label, style = widgetStyles().button(checked, failed), maxLines = 1) }
        }
    }
}

private fun answerAction(context: Context, figures: RaidFigures, characterKey: String?, state: SignupState): Action {
    val guild = figures.guildId
    val raid = figures.raidId
    if (guild == null || raid == null || characterKey == null) return openApp(context)
    if (figures.asksSpec) return openApp(context, DeepLink.Raid(guild, raid, spec = "ask").url())
    return actionRunCallback<SignupCallback>(
        actionParametersOf(
            SignupCallback.GUILD to guild, SignupCallback.RAID to raid,
            SignupCallback.CHARACTER to characterKey, SignupCallback.STATE to state.wire,
        ),
    )
}

/** Signed out, no character, no guild, no raid: the words, and `Open the app` when it helps. */
@Composable
private fun WordyRaid(figures: RaidFigures) {
    val context = LocalContext.current
    val label = when (figures.state) {
        RaidSurfaceState.SignedOut -> context.getString(R.string.app_name)
        RaidSurfaceState.NoGuild -> context.getString(R.string.tab_raids)
        else -> context.getString(R.string.widget_nextRaid)
    }
    Column(GlanceModifier.fillMaxSize().semantics { contentDescription = figures.description }) {
        Label(label)
        Spacer(GlanceModifier.height(10.dp))
        Text(figures.title, style = widgetStyles().title, maxLines = 3)
        figures.subtitle?.let {
            Spacer(GlanceModifier.height(4.dp))
            Text(it, style = widgetStyles().muted, maxLines = 3)
        }
        Spacer(GlanceModifier.defaultWeight())
        when {
            figures.state == RaidSurfaceState.SignedOut || figures.state == RaidSurfaceState.NoCharacter ->
                Text(context.getString(R.string.widget_openApp), style = widgetStyles().link, maxLines = 1)
            figures.stale != null -> Text(figures.stale, style = widgetStyles().muted, maxLines = 1)
        }
    }
}
