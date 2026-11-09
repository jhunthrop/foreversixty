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
import gg.foreversixty.app.R
import gg.foreversixty.app.designsystem.FsColor
import gg.foreversixty.app.widgets.logic.UpgradeFigures

/**
 * Next upgrade at 2 x 2 and 4 x 2 (mocks `widgets-guild`, `-noguild`, `-open`): `Upgrades`, the gain
 * in gold and `next upgrade`; the 4 x 2 adds the item and where it comes from, with its icon in its
 * quality frame. The widget shows the next upgrade, never a total (upgrades overlap). Any state but
 * ok says its sentence; a tap opens the app. [icons] are the framed item icons by item id.
 */
@Composable
fun UpgradeWidgetContent(figures: UpgradeFigures, medium: Boolean, icons: Map<Int, Bitmap>) {
    val context = LocalContext.current
    val styles = widgetStyles()
    val first = figures.rows.firstOrNull()
    val gain = figures.gain
    val sentence = figures.sentence
    Box(
        GlanceModifier.fillMaxSize().background(ImageProvider(R.drawable.widget_card)).padding(WidgetDesign.padding)
            .clickable(openApp(context)),
    ) {
        Row(GlanceModifier.fillMaxSize().semantics { contentDescription = figures.description }) {
            Column(GlanceModifier.defaultWeight().fillMaxHeight()) {
                Text(context.getString(R.string.widget_card_upgrades), style = styles.cardKey, maxLines = 1)
                Spacer(GlanceModifier.height(WidgetDesign.cardGap))
                if (sentence != null || gain == null) {
                    Text(sentence.orEmpty(), style = if (figures.opensApp) styles.link else styles.body, maxLines = 4)
                } else {
                    Text(gain, style = styles.cardBig(FsColor.gold), maxLines = 1)
                    Spacer(GlanceModifier.height(WidgetDesign.cardGap))
                    Text(context.getString(R.string.widget_card_nextUpgrade), style = styles.cardLine, maxLines = 1)
                    figures.item?.takeIf { medium }?.let {
                        Spacer(GlanceModifier.height(WidgetDesign.cardGap))
                        Text(it, style = styles.cardLine, maxLines = 2)
                    }
                }
            }
            if (medium && first != null && sentence == null) {
                Spacer(GlanceModifier.width(12.dp))
                Image(
                    ImageProvider(icons[first.itemId] ?: WidgetIcons.framed(null, first.quality)), contentDescription = null,
                    modifier = GlanceModifier.size(ICON_SIZE.dp),
                )
            }
        }
    }
}

private const val ICON_SIZE = 56
