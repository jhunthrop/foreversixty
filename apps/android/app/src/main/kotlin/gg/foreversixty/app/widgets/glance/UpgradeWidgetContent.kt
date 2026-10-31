package gg.foreversixty.app.widgets.glance

import android.graphics.Bitmap
import androidx.compose.runtime.Composable
import androidx.compose.ui.unit.dp
import androidx.glance.GlanceModifier
import androidx.glance.Image
import androidx.glance.ImageProvider
import androidx.glance.LocalContext
import androidx.glance.action.clickable
import androidx.glance.background
import androidx.glance.layout.Alignment
import androidx.glance.layout.Box
import androidx.glance.layout.Column
import androidx.glance.layout.Row
import androidx.glance.layout.Spacer
import androidx.glance.layout.fillMaxSize
import androidx.glance.layout.fillMaxWidth
import androidx.glance.layout.height
import androidx.glance.layout.padding
import androidx.glance.layout.size
import androidx.glance.layout.width
import androidx.glance.semantics.contentDescription
import androidx.glance.semantics.semantics
import androidx.glance.text.Text
import gg.foreversixty.app.R
import gg.foreversixty.app.items.ItemQuality
import gg.foreversixty.app.widgets.logic.UpgradeFigures
import gg.foreversixty.app.widgets.logic.UpgradeLine

/**
 * Next upgrade at 2 x 2 (the first row large) and 4 x 2 (`Upgrades · {name}`, three rows and the
 * footer), spec §5.6. Rows keep the server's order; names in quality colour; a tap opens the app.
 * [icons] are the framed item icons by item id; a missing one is the bare tile, never a hole.
 */
@Composable
fun UpgradeWidgetContent(figures: UpgradeFigures, medium: Boolean, icons: Map<Int, Bitmap>) {
    val context = LocalContext.current
    Box(
        GlanceModifier.fillMaxSize().background(ImageProvider(R.drawable.widget_card)).padding(WidgetDesign.padding)
            .clickable(openApp(context)),
    ) {
        Column(GlanceModifier.fillMaxSize().semantics { contentDescription = figures.description }) {
            Text(
                (if (medium && figures.sentence == null) figures.header else context.getString(R.string.widget_nextUpgrade)).uppercase(),
                style = widgetStyles().micro, maxLines = 1,
            )
            Spacer(GlanceModifier.height(8.dp))
            val sentence = figures.sentence
            when {
                sentence != null -> {
                    Text(sentence, style = if (figures.opensApp) widgetStyles().link else widgetStyles().body, maxLines = 3)
                }
                medium -> {
                    figures.rows.forEach { Row(it, icons[it.itemId]) }
                    Spacer(GlanceModifier.defaultWeight())
                    figures.footer?.let { Text(it, style = widgetStyles().muted, maxLines = 1) }
                }
                else -> figures.rows.firstOrNull()?.let { first ->
                    Image(ImageProvider(icons[first.itemId] ?: WidgetIcons.framed(null, first.quality)), contentDescription = null,
                        modifier = GlanceModifier.size(44.dp))
                    Spacer(GlanceModifier.height(6.dp))
                    Text(first.name, style = widgetStyles().quality(ItemQuality(first.quality).text), maxLines = 2)
                    Text(first.gain, style = widgetStyles().gain, maxLines = 1)
                }
            }
        }
    }
}

@Composable
private fun Row(line: UpgradeLine, icon: Bitmap?) {
    Row(GlanceModifier.fillMaxWidth().padding(vertical = 3.dp), verticalAlignment = Alignment.CenterVertically) {
        Image(ImageProvider(icon ?: WidgetIcons.framed(null, line.quality)), contentDescription = null, modifier = GlanceModifier.size(28.dp))
        Spacer(GlanceModifier.width(8.dp))
        // The one place a name may cut: a widget's item name on one line (spec §9).
        Text(line.name, style = widgetStyles().quality(ItemQuality(line.quality).text, 13f), maxLines = 1, modifier = GlanceModifier.defaultWeight())
        Spacer(GlanceModifier.width(8.dp))
        Text(line.gain, style = widgetStyles().gain, maxLines = 1)
    }
}
