package id

func GetUniqueId() (int64, error) {
	snowflake, err := NewSnowflake(1, 1)
	if err != nil {
		return 0, err
	}

	return snowflake.GenerateID(), nil
}
