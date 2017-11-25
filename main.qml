import Vault 1.0
import QtQuick 2.7
import QtQuick.Controls 2.0
import QtQuick.Layouts 1.3

ApplicationWindow {
    id: applicationWindow
    visible: true
    width: 360
    height: 480
    title: qsTr("Zpass Client")

    Vault {
        id: myvault
        onUnlocked: {
            console.log("Unlocked");
            view.push(pwListView)
        }
    }

    header: Label {
        id: label
        text: "zpass"
        anchors.horizontalCenter: parent.horizontalCenter
        padding: 10
        horizontalAlignment: Text.AlignHCenter
    }

    StackView {
        id: view
        initialItem: unlockView
        PasswordView {
            id: pwListView
        }
        UnlockView {
            id: unlockView
        }
    }
}
