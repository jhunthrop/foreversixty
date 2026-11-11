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
import androidx.glance.layout.fillMaxHeight
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
import gg.foreversixty.app.designsystem.FsColor
import gg.foreversixty.app.widgets.logic.RaidCard
import gg.foreversixty.app.widgets.logic.RaidSurfaceState

/** A widget tap opens the app on the right screen (spec §5.6); null opens the app where it was. */
internal fun openApp(context: Context, url: String? = null): Action {
    val intent = Intent(context, MainActivity::class.java).addFlags(Intent.FLAG_ACTIVITY_NEW_TASK or Intent.FLAG_ACTIVITY_CLEAR_TOP)
    url?.let { intent.setAction(Intent.ACTION_VIEW).setData(it.toUri()) }
    return actionStartActivity(intent)
}

/**
 * Next raid at 2 x 2 ([medium] false) and 4 x 2 (mocks `widgets-guild`, `-noguild`, `-open`): the
 * card's label, big figure, name and line. The 4 x 2 keeps In, Late and Out beside the card, so the
 * answer stays one tap from the home screen (E1). Signed out and with no character it says what to
 * do. The card is one TalkBack element; each button its own.
 */
@Composable
fun RaidWidgetContent(figures: RaidFigures, medium: Boolean, characterKey: String?) {
    val context = LocalContext.current
    val raidLink = figures.guildId?.let { g -> figures.raidId?.let { r -> DeepLink.Raid(g, r).url() } }
    val card = figures.card
    Box(
        GlanceModifier.fillMaxSize().background(ImageProvider(R.drawable.widget_card)).padding(WidgetDesign.padding)
            .clickable(openApp(context, raidLink)),
    ) {
        when {
            card == null -> WordyRaid(figures)
            medium -> Row(GlanceModifier.fillMaxSize()) {
                Box(GlanceModifier.defaultWeight().fillMaxHeight()) { CardLines(card, figures, medium = true) }
                Spacer(GlanceModifier.width(12.dp))
                CardSide(figures, card, characterKey)
            }
            else -> CardLines(card, figures, medium = false)
        }
    }
}

/**
 * Label, big figure, name and line, with `Updated 14 h ago` at the foot on old data. Beside the
 * 4 x 2's buttons the name keeps one line and the line two, so large text never pushes them out.
 */
@Composable
private fun CardLines(card: RaidCard, figures: RaidFigures, medium: Boolean) {
    val styles = widgetStyles()
    Column(GlanceModifier.fillMaxSize().semantics { contentDescription = figures.description }) {
        Text(card.key, style = styles.cardKey, maxLines = 1)
        Spacer(GlanceModifier.height(WidgetDesign.cardGap))
        Text(card.big, style = styles.cardBig(if (card.cancelled) FsColor.wipe else FsColor.strong), maxLines = 1)
        Spacer(GlanceModifier.height(WidgetDesign.cardGap))
        Text(card.name, style = styles.cardName, maxLines = if (medium) 1 else 2)
        card.line?.let {
            Spacer(GlanceModifier.height(WidgetDesign.cardGap))
            Text(it, style = styles.cardLine, maxLines = if (medium) 2 else 3)
        }
        figures.stale?.let {
            Spacer(GlanceModifier.defaultWeight())
            Text(it, style = styles.muted, maxLines = 1)
        }
    }
}

/** The 4 x 2's side: who is coming at the top, In, Late and Out at the foot while answers are open. */
@Composable
private fun CardSide(figures: RaidFigures, card: RaidCard, characterKey: String?) {
    Column(GlanceModifier.fillMaxHeight(), horizontalAlignment = Alignment.End) {
        card.counts?.let { Text(it, style = widgetStyles().muted, maxLines = 1) }
        Spacer(GlanceModifier.defaultWeight())
        if (figures.buttons) Answers(figures, characterKey)
    }
}

@Composable
private fun Label(text: String, modifier: GlanceModifier = GlanceModifier) {
    Row(modifier, verticalAlignment = Alignment.CenterVertically) {
        Image(ImageProvider(R.drawable.brand_seal), contentDescription = null, modifier = GlanceModifier.size(WidgetDesign.sealSize))
        Spacer(GlanceModifier.width(6.dp))
        Text(text, style = widgetStyles().micro, maxLines = 1)
    }
}

/** In, Late, Bench, Out: 48 dp, the answer checked; a press answers in place, or asks the spec first. */
@Composable
private fun Answers(figures: RaidFigures, characterKey: String?) {
    val context = LocalContext.current
    Row {
        SignupState.entries.forEachIndexed { index, state ->
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
            Spacer(GlanceModifier.height(WidgetDesign.cardGap))
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
