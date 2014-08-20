type Object interface {
        // Objects must maintain
        // a reference to an Info
        // struct, which contains
        // this object's riak
        // metadata. Info() must
        // never return nil, or it
        // will cause a panic.
        Info() *Info

        // Marshal should return the encoded
        // value of the object, and any
        // relevant errors.
        Marshal() ([]byte, error)

        // Unmarshal should unmarshal the object
        // from a []byte. It can safely use
        // zero-copy methods, as the byte slice
        // passed to it will "belong" to the
        // object.
        Unmarshal([]byte) error
}
