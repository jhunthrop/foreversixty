package gg.foreversixty.app.model

import kotlinx.serialization.Serializable

/** `/data/active-build.json`: the client build whose static data the site serves now. */
@Serializable
data class ActiveBuild(val build: String)

/** One entry of `/data/{build}/classes.json`. [color] is `#RRGGBB`. */
@Serializable
data class ClassInfo(
    val id: Int,
    val name: String,
    val slug: String,
    val color: String,
)
