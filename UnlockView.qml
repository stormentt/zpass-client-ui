import QtQuick 2.0
import QtQuick.Controls 2.0

Item {
    id: unlockView
    width: 360
    height: 480

    Rectangle {
        id: rectangle
        color: "#ffffff"
        anchors.fill: parent
    }

    Button {
        id: button
        text: qsTr("Unlock")
        anchors.horizontalCenter: parent.horizontalCenter
        anchors.bottom: parent.bottom
        anchors.bottomMargin: 10
        onClicked: {
            console.log("Unlocking");
            console.log(myvault.unlock(vaultPw.text));
            console.log("Called Unlocked?");
        }
    }

    TextField {
        id: vaultPw
        width: 230
        height: 50
        text: qsTr("")
        anchors.verticalCenterOffset: 0
        anchors.horizontalCenterOffset: 0
        anchors.horizontalCenter: parent.horizontalCenter
        anchors.verticalCenter: parent.verticalCenter
        leftPadding: 5
        echoMode: TextInput.Password;
        placeholderText: qsTr("Password")
    }
}
