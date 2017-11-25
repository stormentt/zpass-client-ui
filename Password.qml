import QtQuick 2.0

Item {
    id: password
    Item {
        height: 50
        Text {
            id: text1
            y: 18
            text: pwName
            anchors.verticalCenter: parent.verticalCenter
            anchors.left: parent.left
            anchors.leftMargin: 50
            font.pixelSize: 14
        }
    }

}
