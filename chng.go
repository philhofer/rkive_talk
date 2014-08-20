// fetch myBlob

chng := func(o rkive.Object) error {
	v := o.(*rkive.Blob)
	
	// the change already happened
	if bytes.Equal(v.Data, []byte("New Data!") {
		return rkive.ErrDone
	}

	// the change hasn't happened yet
	if bytes.Equal(v.Data, old) {
		v.Data = []byte("New Data!")
	}
	
	// the value was changed in the database;
	// re-evaluate
	return rkive.ErrModified
}

err = rkive.PushChangeset(myBlob, chng, nil)
