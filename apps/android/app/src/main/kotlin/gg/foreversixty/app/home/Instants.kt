package gg.foreversixty.app.home

import java.time.Instant

/** The API's RFC 3339 timestamps; null when absent or unreadable. */
fun parseInstant(text: String?): Instant? = text?.let { runCatching { Instant.parse(it) }.getOrNull() }
