plugins {
    alias(libs.plugins.forever.android.library)
    alias(libs.plugins.kotlin.serialization)
    alias(libs.plugins.forever.contract.fixtures)
}

dependencies {
    api(libs.kotlinx.serialization.json)

    testImplementation(libs.junit)
}
