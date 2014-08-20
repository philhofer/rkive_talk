type ObjectM interface {
        Object

        // Empty should return an initialized
        // (zero-value) object of the same underlying
        // type as the parent.
        NewEmpty() ObjectM

        // Merge should merge the argument object into the method receiver. It
        // is safe to type-assert the argument of Merge to the same type
        // as the type of the object satisfying the inteface. (Under the hood,
        // the argument passed to Merge is simply the value of NewEmpty() after
        // data has been read into it.) Merge is used to iteratively merge many sibling objects.
        Merge(o ObjectM)
}
