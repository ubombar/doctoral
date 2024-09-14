package zmle

// This is the ZML, Zettelkasten Manipilation Language. It is designated to
// operate on set of text files. This is the executor instance. ZMLE.
//
// It stands for Zettelkasten Manipulation Language Executor. In the future,
// we might want to add an actual language to perform these. But for now
// it is done by the ZMLE interface.
type ZMLE struct {
	// This is the internal state representation for the groups.
	internalState zmleState
}
