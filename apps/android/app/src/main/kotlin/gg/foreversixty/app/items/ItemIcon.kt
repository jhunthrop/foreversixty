package gg.foreversixty.app.items

import android.content.Context
import androidx.compose.foundation.Image
import androidx.compose.foundation.background
import androidx.compose.foundation.border
import androidx.compose.foundation.layout.Box
import androidx.compose.foundation.layout.padding
import androidx.compose.foundation.layout.size
import androidx.compose.foundation.shape.RoundedCornerShape
import androidx.compose.runtime.Composable
import androidx.compose.runtime.staticCompositionLocalOf
import androidx.compose.ui.Modifier
import androidx.compose.ui.draw.clip
import androidx.compose.ui.draw.drawBehind
import androidx.compose.ui.draw.shadow
import androidx.compose.ui.geometry.CornerRadius
import androidx.compose.ui.graphics.Color
import androidx.compose.ui.graphics.PathEffect
import androidx.compose.ui.graphics.drawscope.Stroke
import androidx.compose.ui.graphics.painter.Painter
import androidx.compose.ui.layout.ContentScale
import androidx.compose.ui.platform.LocalContext
import androidx.compose.ui.unit.Dp
import androidx.compose.ui.unit.dp
import coil3.ImageLoader
import coil3.compose.AsyncImage
import coil3.disk.DiskCache
import coil3.disk.directory
import coil3.request.ImageRequest
import coil3.request.crossfade
import gg.foreversixty.app.designsystem.FsColor
import gg.foreversixty.app.designsystem.FsRadius
import gg.foreversixty.app.network.StaticDataClient

/**
 * Where item icons come from. Live: `/data/{build}/icons/{icon}.webp` on the site through Coil
 * with a 200 MB disk cache (spec §11). Screenshot tests provide a painter from the repo's own
 * icon files instead.
 */
fun interface ItemIconPainterSource {
    @Composable
    fun painter(icon: String, build: String): Painter?
}

/** Null in the app (Coil loads the CDN file); set by screenshot tests. */
val LocalItemIconPainter = staticCompositionLocalOf<ItemIconPainterSource?> { null }

/** The one icon loader, its disk cache under the app's cache directory. */
object ItemIconLoader {
    private const val DISK_BYTES = 200L * 1024 * 1024
    private const val FADE_MS = 120

    @Volatile
    private var loader: ImageLoader? = null

    fun get(context: Context): ImageLoader = loader ?: synchronized(this) {
        loader ?: ImageLoader.Builder(context.applicationContext)
            .diskCache {
                DiskCache.Builder()
                    .directory(context.applicationContext.cacheDir.resolve("item-icons"))
                    .maxSizeBytes(DISK_BYTES)
                    .build()
            }
            .crossfade(FADE_MS)
            .build()
            .also { loader = it }
    }

    fun url(icon: String, build: String): String = "${StaticDataClient.SITE_URL}/data/$build/icons/$icon.webp"
}

/**
 * A square game icon in its quality border (spec §4.3): 1.5 dp quality border, a 1 dp inner black
 * line and the quality glow. The tile draws at once on `#131824`; the image fades in, and a
 * missing icon leaves the tile, never a broken-image glyph (§8). Hidden from TalkBack: the name
 * beside it carries the item.
 */
@Composable
fun ItemIcon(icon: String?, quality: Int?, build: String, size: Dp, modifier: Modifier = Modifier, dashed: Boolean = false) {
    val color = ItemQuality(quality).border
    val shape = RoundedCornerShape(FsRadius.pill)
    val frame = if (dashed) {
        Modifier.drawBehind {
            drawRoundRect(
                color = FsColor.line,
                cornerRadius = CornerRadius(FsRadius.pill.toPx()),
                style = Stroke(width = 1.5.dp.toPx(), pathEffect = PathEffect.dashPathEffect(floatArrayOf(8f, 6f))),
            )
        }
    } else {
        Modifier
            .shadow(6.dp, shape, ambientColor = color.copy(alpha = GLOW), spotColor = color.copy(alpha = GLOW))
            .border(1.5.dp, color, shape)
    }
    Box(modifier.size(size).then(frame).clip(shape).background(FsColor.cardTop)) {
        if (icon != null && !dashed) {
            val local = LocalItemIconPainter.current
            if (local != null) {
                local.painter(icon, build)?.let {
                    Image(it, null, Modifier.size(size), contentScale = ContentScale.Crop)
                }
            } else {
                val context = LocalContext.current
                AsyncImage(
                    model = ImageRequest.Builder(context).data(ItemIconLoader.url(icon, build)).build(),
                    imageLoader = ItemIconLoader.get(context),
                    contentDescription = null,
                    contentScale = ContentScale.Crop,
                    modifier = Modifier.size(size),
                )
            }
        }
        if (!dashed) {
            Box(Modifier.size(size).padding(1.5.dp).border(1.dp, Color.Black, shape))
        }
    }
}

private const val GLOW = 0.22f
