package gg.foreversixty.app.raids.logic

import androidx.compose.ui.graphics.Color
import androidx.compose.ui.text.AnnotatedString
import androidx.compose.ui.text.LinkAnnotation
import androidx.compose.ui.text.SpanStyle
import androidx.compose.ui.text.TextLinkStyles
import androidx.compose.ui.text.buildAnnotatedString
import androidx.compose.ui.text.font.FontStyle
import androidx.compose.ui.text.font.FontWeight
import androidx.compose.ui.text.style.TextDecoration
import androidx.compose.ui.text.withLink
import androidx.compose.ui.text.withStyle

/**
 * The API's sanitised `notes_html` as annotated text (spec F4.2 item 7): paragraphs, line breaks,
 * list items, bold, italic and https links. The API has already sanitised it; any other tag drops
 * to its text.
 */
object NotesText {
    private val TOKEN = Regex("<([^>]*)>|([^<]+)")
    private val HREF = Regex("href=\"([^\"]*)\"")
    private val BREAKS = setOf("/p", "br", "br/", "/li", "/ul", "/ol", "/h1", "/h2", "/h3")

    fun annotated(html: String, link: Color): AnnotatedString = buildAnnotatedString {
        var bold = false
        var italic = false
        var href: String? = null
        var pendingBreak = false
        for (match in TOKEN.findAll(html)) {
            val tag = match.groups[1]?.value
            if (tag == null) {
                val text = decode(match.value)
                if (text.isBlank()) continue
                if (pendingBreak && length > 0) append('\n')
                pendingBreak = false
                appendRun(text.replace('\n', ' '), bold, italic, href, link)
                continue
            }
            when (val name = tag.substringBefore(' ').lowercase()) {
                "strong", "b" -> bold = true
                "/strong", "/b" -> bold = false
                "em", "i" -> italic = true
                "/em", "/i" -> italic = false
                "a" -> href = HREF.find(tag)?.groupValues?.get(1)?.let(::decode)?.takeIf { it.startsWith("https://") }
                "/a" -> href = null
                "li" -> {
                    if (length > 0) append('\n')
                    pendingBreak = false
                    append("• ")
                }
                else -> if (name in BREAKS) pendingBreak = true
            }
        }
    }

    private fun AnnotatedString.Builder.appendRun(text: String, bold: Boolean, italic: Boolean, href: String?, link: Color) {
        val style = SpanStyle(
            fontWeight = if (bold) FontWeight.SemiBold else null,
            fontStyle = if (italic) FontStyle.Italic else null,
        )
        if (href == null) {
            withStyle(style) { append(text) }
        } else {
            val linkStyle = TextLinkStyles(style.merge(SpanStyle(color = link, textDecoration = TextDecoration.Underline)))
            withLink(LinkAnnotation.Url(href, linkStyle)) { append(text) }
        }
    }

    private fun decode(text: String): String = listOf(
        "&amp;" to "&", "&lt;" to "<", "&gt;" to ">", "&quot;" to "\"", "&#39;" to "'", "&#x27;" to "'", "&nbsp;" to " ",
    ).fold(text) { acc, (from, to) -> acc.replace(from, to) }
}
