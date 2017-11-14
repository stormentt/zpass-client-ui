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
        }
    }

    header: Label {
        id: label
        text: "Unlock Vault"
        anchors.horizontalCenter: parent.horizontalCenter
        padding: 10
        horizontalAlignment: Text.AlignHCenter
    }

    TextField {
        id: vaultPw
        x: -67
        width: 231
        height: 46
        text: qsTr("")
        anchors.horizontalCenter: parent.horizontalCenter
        anchors.top: parent.top
        anchors.topMargin: 131
        leftPadding: 5
        padding: 0
        echoMode: TextInput.Password;
        placeholderText: qsTr("Password")
    }

    Button {
        id: button
        x: 145
        y: 432
        text: qsTr("Unlock")
        anchors.horizontalCenter: parent.horizontalCenter
        anchors.bottom: parent.bottom
        anchors.bottomMargin: 0
        onClicked: {
            console.log("Unlocking");
            console.log(myvault.unlock("test"));
            console.log("Called Unlocked?");
        }
    }
}
