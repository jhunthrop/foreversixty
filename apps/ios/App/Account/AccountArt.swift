import DesignSystem
import SwiftUI

/// The hero band ground (`--surface-band`): the ember radial over the warm dark.
struct BandBackground: View {
    var body: some View {
        ZStack {
            FSColor.bg
            RadialGradient(
                colors: [FSColor.ember.opacity(0.28), .clear],
                center: .top,
                startRadius: 0,
                endRadius: 520
            )
        }
        .ignoresSafeArea()
    }
}

/// The toast at the top of F1.1: F1.4 and the 401 line in wipe red, a finished deletion neutral.
struct NoticeToast: View {
    let text: String
    /// A failure is wipe-red with the alert glyph; anything else is neutral with the check.
    var isFailure = true

    private var tint: Color { isFailure ? FSColor.wipe : FSColor.line }

    var body: some View {
        HStack(alignment: .top, spacing: FSSpacing.space3) {
            Image(systemName: isFailure ? "exclamationmark.circle" : "checkmark.circle")
                .foregroundStyle(isFailure ? FSColor.wipe : FSColor.gold)
            Text(text)
                .font(ScaledType.body)
                .foregroundStyle(FSColor.strong)
                .frame(maxWidth: .infinity, alignment: .leading)
                .fixedSize(horizontal: false, vertical: true)
        }
        .padding(FSSpacing.space4)
        .background(FSColor.raised, in: .rect(cornerRadius: FSRadius.panel))
        .overlay(RoundedRectangle(cornerRadius: FSRadius.panel).stroke(tint.opacity(0.45)))
        .accessibilityElement(children: .combine)
        .accessibilityAddTraits(.isStaticText)
    }
}
