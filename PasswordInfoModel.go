package main

import "github.com/therecipe/qt/core"

const (
	Name = int(core.Qt__UserRole) + 1<<iota
	Selector
)

type PasswordModel struct {
	core.QAbstractListModel

	_ func() `constructor:"init"`

	_ map[int]*core.QByteArray `property:"roles"`
	_ []*Password              `property:"passwords"`

	_ func(*Password)                      `slot:"addPassword"`
	_ func(row int, name, selector string) `slot:"editPassword"`
	_ func(row int)                        `slot:"removePassword"`
}

type Password struct {
	core.QObject

	_ string `property:"name"`
	_ string `property:"selector"`
}

func init() {
	Password_QRegisterMetaType()
}

func (m *Password) init() {
	m.SetRoles(map[int]*core.QByteArray{
		Name:     core.NewQByteArray2("name", len("name")),
		Selector: core.NewQByteArray2("selector", len("selector")),
	})

	m.ConnectData(m.data)
	m.ConnectRowCount(m.rowCount)
	m.ConnectColumnCount(m.columnCount)
	m.ConnectRoleNames(m.roleNames)

	m.ConnectAddPassword(m.addPassword)
	m.ConnectEditPassword(m.editPassword)
	m.ConnectRemovePassword(m.removePassword)
}

func (m *PasswordModel) data(index *core.QModelIndex, role int) *core.QVariant {
	if !index.IsValid() {
		return core.NewQVariant()
	}

	if index.Row() >= len(m.Passwords()) {
		return core.NewQVariant()
	}

	var p = m.Passwords()[index.Row()]

	switch role {
	case Name:
		{
			return core.NewQVariant14(p.Name())
		}

	case Selector:
		{
			return core.NewQVariant14(p.Selector())
		}

	default:
		{
			return core.NewQVariant()
		}
	}
}

func (m *PasswordModel) rowCount(parent *core.QModelIndex) int {
	return len(m.Passwords())
}

func (m *PasswordModel) columnCount(parent *core.QModelIndex) int {
	return 1
}

func (m *PasswordModel) roleNames() map[int]*core.QByteArray {
	return m.Roles()
}

func (m *PasswordModel) addPassword(p *Password) {
	m.BeginInsertRows(core.NewQModelIndex(), len(m.Passwords()), len(m.Passwords()))
	m.SetPasswords(append(m.Passwords(), p))
	m.EndInsertRows()
}

func (m *PasswordModel) editPassword(row int, name string, selector string) {
	var p = m.Passwords()[row]

	if name != "" {
		p.SetName(name)
	}

	if selector != "" {
		p.SetSelector(selector)
	}

	var pIndex = m.Index(row, 0, core.NewQModelIndex())
	m.DataChanged(pIndex, pIndex, []int{Name, Selector})
}

func (m *PasswordModel) removePassword(row int) {
	m.BeginRemoveRows(core.NewQModelIndex(), row, row)
	m.SetPasswords(append(m.Passwords()[:row], m.Passwords()[row+1:]...))
	m.EndRemoveRows()
}
