import QtQuick 2.0

Item {
    id: passwordView
    width: 360
    height: 480

    Row {
        id: row
        x: 0
        y: 0
        width: 360
        height: 480


        Item {
            id: item1
            width: (parent.width * 0.4)
            height: (parent.height)

            Rectangle {
                id: rectangle
                color: "#ffffff"
                anchors.fill: parent
            }

            ListView {
                id: pwList
                anchors.fill: parent
                model: ListModel {
                    ListElement {
                        pwName: "Grey"
                    }
                }
                delegate: Password {
                    x: 5
                }
            }

        }

        Item {
            id: item2
            width: parent.width * 0.6
            height: parent.height
            anchors.right: parent.right

            Rectangle {
                id: rectangle1
                color: "#efefef"
                anchors.fill: parent
            }
        }
    }

}
