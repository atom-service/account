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
	if err != nil {
		return err
	}
	err = createLabelTable()
	if err != nil {
		return err
	}
	err = createLabelMappingTable()
	if err != nil {
		return err
	}
	err = createGroupTable()
	if err != nil {
		return err
	}
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
	if err != nil {
		return err
	}
	err = truncateLabelTable()
	if err != nil {
		return err
	}
	err = truncateLabelMappingTable()
	if err != nil {
		return err
	}
	err = truncateGroupTable()
	if err != nil {
		return err
	}
	err = truncateGroupMappingTable()
	if err != nil {
		return err
	}
	return nil
}
