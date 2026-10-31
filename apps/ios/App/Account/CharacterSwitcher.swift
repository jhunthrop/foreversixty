import Auth
import DesignSystem
import Models
import SwiftUI

/// The selector (spec §3.4): crest 34, the full name in the class colour, the short line and a
/// chevron. The one way to switch; Home, Raids and BiS put it in their toolbar with
/// `.characterSelector()`. The name never truncates: from AX1 it stops growing and the short
/// line drops, and a long press shows the Large Content Viewer.
struct CharacterSelector: View {
    let character: AccountCharacter
    let isOpen: Bool
    let action: () -> Void
    @Environment(\.dynamicTypeSize) private var typeSize

    var body: some View {
        Button(action: action) {
            HStack(spacing: FSSpacing.space2) {
                ClassCrest(className: character.className, size: 34)
                VStack(alignment: .leading, spacing: 0) {
                    Text(character.name)
                        .font(ScaledType.name)
                        .foregroundStyle(ClassColor.of(character.className))
                        .lineLimit(1)
                        .fixedSize()
                    if !typeSize.isAccessibilitySize, let detail = AccountCopy.selectorDetail(character) {
                        Text(detail).font(ScaledType.meta).foregroundStyle(FSColor.muted).lineLimit(1)
                    }
                }
                Image(systemName: "chevron.down").font(.footnote.weight(.semibold)).foregroundStyle(FSColor.muted)
            }
            .padding(.vertical, FSSpacing.space1)
            .padding(.leading, FSSpacing.space1)
            .padding(.trailing, FSSpacing.space3)
            .background(FSColor.raised, in: .capsule)
            .overlay(Capsule().strokeBorder(isOpen ? FSColor.gold : FSColor.gold.opacity(0.17)))
            .shadow(color: isOpen ? FSColor.gold.opacity(0.4) : .clear, radius: 8)
        }
        .buttonStyle(.plain)
        .dynamicTypeSize(...DynamicTypeSize.accessibility1)
        .accessibilityShowsLargeContentViewer()
        .accessibilityLabel(AccountCopy.selectorLabel(character))
        .accessibilityHint(Copy.accountSwitcherOpen)
        .accessibilityAction(named: Copy.accountSwitcherAction, action)
    }
}

/// F1.5: one 68 pt row per character, the current one washed and checked.
struct SwitcherSheet: View {
    let me: Me
    let currentKey: String?
    let onPick: (AccountCharacter) -> Void

    var body: some View {
        ScrollView {
            VStack(alignment: .leading, spacing: 0) {
                VStack(alignment: .leading, spacing: FSSpacing.space1) {
                    Text(Copy.switcherTitle)
                        .font(ScaledType.sec)
                        .dynamicTypeSize(...DynamicTypeSize.accessibility2)
                        .textCase(.uppercase)
                        .foregroundStyle(FSColor.strong)
                        .accessibilityAddTraits(.isHeader)
                    if let battletag = me.user.battletag {
                        Text(Copy.switcherWho(battletag: battletag)).font(ScaledType.meta).foregroundStyle(FSColor.muted)
                    }
                }
                .padding(FSSpacing.space4)
                ForEach(me.characters) { character in
                    row(character)
                    Divider().overlay(FSColor.lineSoft)
                }
                Text(Copy.switcherFooter)
                    .font(ScaledType.meta)
                    .foregroundStyle(FSColor.muted)
                    .padding(FSSpacing.space4)
            }
        }
        .background(FSColor.raised)
        .presentationDetents([.medium, .large])
        .presentationDragIndicator(.visible)
    }

    private func row(_ character: AccountCharacter) -> some View {
        let isCurrent = character.key == currentKey
        return Button { onPick(character) } label: {
            HStack(spacing: FSSpacing.space3) {
                ClassCrest(className: character.className, size: 44)
                VStack(alignment: .leading, spacing: 2) {
                    // `Thrall · US · PvP`: same-named characters differ by region and ruleset.
                    let name = Text(character.name).font(ScaledType.name).foregroundStyle(ClassColor.of(character.className))
                    let place = Text(" · \(AccountCopy.place(character))").font(ScaledType.meta).foregroundStyle(FSColor.muted)
                    Text("\(name)\(place)").fixedSize(horizontal: false, vertical: true)
                    Text(AccountCopy.switcherDetail(character)).font(ScaledType.meta).foregroundStyle(FSColor.text)
                    Text(character.guild?.name ?? Copy.switcherNoGuild).font(ScaledType.meta).foregroundStyle(FSColor.muted)
                }
                Spacer(minLength: 0)
                if isCurrent {
                    Image(systemName: "checkmark").foregroundStyle(FSColor.gold)
                }
            }
            .padding(.horizontal, FSSpacing.space4)
            .frame(minHeight: 68)
            .background(isCurrent ? Color.white.opacity(0.04) : .clear)
            .contentShape(.rect)
        }
        .buttonStyle(.plain)
        .accessibilityAddTraits(isCurrent ? .isSelected : [])
    }
}

/// Puts the selector in a root screen's toolbar and opens F1.5 from it. A no-op when signed out
/// or with no characters (the sheet is never empty).
struct CharacterSelectorToolbar: ViewModifier {
    @Environment(SessionStore.self) private var session
    @State private var isOpen = false

    func body(content: Content) -> some View {
        content
            .toolbar {
                if let character = session.currentCharacter {
                    ToolbarItem(placement: .topBarLeading) {
                        CharacterSelector(character: character, isOpen: isOpen) { isOpen = true }
                    }
                    .sharedBackgroundVisibility(.hidden)
                }
            }
            .sheet(isPresented: $isOpen) {
                if let me = session.phase.me {
                    SwitcherSheet(me: me, currentKey: session.currentCharacter?.key) { picked in
                        session.pick(picked.key)
                        isOpen = false
                    }
                }
            }
            .sensoryFeedback(.selection, trigger: session.pickedKey)
    }
}

extension View {
    /// The one character switcher, for the Home, Raids and BiS root screens (spec §3.4).
    func characterSelector() -> some View { modifier(CharacterSelectorToolbar()) }

    /// The selector only where the session is in play (sign-in on); a no-op otherwise.
    @ViewBuilder
    func characterSelector(isOn: Bool) -> some View {
        if isOn { characterSelector() } else { self }
    }
}
