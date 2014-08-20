package main

import (
	"github.com/philhofer/rkive"
)

func main() {
	// START OMIT
	// Dial Riak
	client, err := rkive.DialOne("localhost:8078", "test_client")
	if err != nil {
		// ...
	}

	// use Blob
	myblob := &rkive.Blob{
		RiakInfo: &rkive.Info{},
		Data: []byte("Hello, World!"),
	}

	// new @ bucket = "test_bucket", key = "my_first_blob"
	err = client.Buckets("test_bucket").New(myblob, "my_first_blob")
	// END OMIT
	if err != nil {
		// ...
	}
}
