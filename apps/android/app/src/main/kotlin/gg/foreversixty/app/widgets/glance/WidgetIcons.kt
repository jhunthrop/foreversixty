package gg.foreversixty.app.widgets.glance

import android.content.Context
import android.graphics.Bitmap
import android.graphics.Canvas
import android.graphics.Paint
import android.graphics.RectF
import androidx.compose.ui.graphics.toArgb
import androidx.core.graphics.createBitmap
import coil3.request.ImageRequest
import coil3.request.SuccessResult
import coil3.request.allowHardware
import coil3.toBitmap
import gg.foreversixty.app.designsystem.FsColor
import gg.foreversixty.app.items.ItemIconLoader
import gg.foreversixty.app.items.ItemQuality

/**
 * Item icons for the widgets (spec §4.3, §9: every item keeps its icon and quality colour). A widget
 * cannot draw a Compose border, so the quality frame is painted into the bitmap: the tile ground at
 * once, the art over it when the cached file loads, the 1.5 dp quality border on top.
 */
internal object WidgetIcons {
    private const val SIZE_PX = 96
    private const val BORDER_FRACTION = 0.05f
    private const val RADIUS_FRACTION = 0.12f

    /** The icon file through the app's one icon loader (disk cached), or null offline. */
    suspend fun load(context: Context, icon: String, build: String?): Bitmap? {
        if (build == null) return null
        val request = ImageRequest.Builder(context).data(ItemIconLoader.url(icon, build)).allowHardware(false).size(SIZE_PX).build()
        return (ItemIconLoader.get(context).execute(request) as? SuccessResult)?.image?.toBitmap()
    }

    fun framed(art: Bitmap?, quality: Int?): Bitmap {
        val out = createBitmap(SIZE_PX, SIZE_PX)
        val canvas = Canvas(out)
        val radius = SIZE_PX * RADIUS_FRACTION
        val border = SIZE_PX * BORDER_FRACTION
        val bounds = RectF(0f, 0f, SIZE_PX.toFloat(), SIZE_PX.toFloat())
        canvas.drawRoundRect(bounds, radius, radius, Paint(Paint.ANTI_ALIAS_FLAG).apply { color = FsColor.cardTop.toArgb() })
        art?.let { canvas.drawBitmap(it, null, RectF(border, border, SIZE_PX - border, SIZE_PX - border), Paint(Paint.FILTER_BITMAP_FLAG)) }
        val stroke = Paint(Paint.ANTI_ALIAS_FLAG).apply {
            style = Paint.Style.STROKE
            strokeWidth = border
            color = ItemQuality(quality).border.toArgb()
        }
        val inset = border / 2
        canvas.drawRoundRect(RectF(inset, inset, SIZE_PX - inset, SIZE_PX - inset), radius, radius, stroke)
        return out
    }
}
