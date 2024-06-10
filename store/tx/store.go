package tx

type TxStore struct {
}

func NewTxStore() *TxStore {

	return &TxStore{}
}
func (tx *TxStore) Upload(bucketName, objectKey, fileName string) error {

	return nil
}
