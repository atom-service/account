package dao

// MustInitTables MustInitTables
func MustInitTables() {
	var err error
	err = InitTables()
	if err != nil {
		panic(err)
	}
}

// InitTables InitTables
func InitTables() error {
	var err error
	err = createUserTable()
	err = createLabelTable()
	err = createLabelMappingTable()
	err = createGroupTable()
	err = createGroupMappingTable()
	if err != nil {
		return err
	}
	return nil
}

// MustTruncateTables MustTruncateTables
func MustTruncateTables() {
	var err error
	err = TruncateTables()
	if err != nil {
		panic(err)
	}
}

// TruncateTables TruncateTables
func TruncateTables() error {
	var err error
	err = truncateUserTable()
	err = truncateLabelTable()
	err = truncateLabelMappingTable()
	err = truncateGroupTable()
	err = truncateGroupMappingTable()
	if err != nil {
		return err
	}
	return nil
}
