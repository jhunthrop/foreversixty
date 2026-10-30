plugins {
    alias(libs.plugins.forever.android.library)
    alias(libs.plugins.forever.android.compose)
    alias(libs.plugins.forever.shared.fonts)
}

dependencies {
    api(libs.compose.material3)
    api(libs.compose.ui)
}
