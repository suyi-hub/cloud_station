package aws

type AwsStore struct {
}

func NewAwsStore() *AwsStore {

	return &AwsStore{}
}

func (aws *AwsStore) Upload(bucketName, objectKey, fileName string) error {

	return nil
}
