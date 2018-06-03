package domain

type (
	HostProvider int
	FlowType int
)

type CloudStorage struct {
	HostProvider HostProvider
	ProjectId string
	FlowType FlowType
}

type LocalRepository interface {
	FileExists(filename string) (fileExists bool, err error)
}

type BucketRepository interface {
	List(projectId string) (buckets []Bucket, err error)
	LocalRepository
	DownloadFile(filename string) (success bool, err error)
	UploadFile(flename string) (success bool, err error)
}

type FileRepository interface {
	Store(file File)
	FindById(id int) File
}

type Bucket struct {
	Name string `json:"name"`
}

type Buckets struct {
	Buckets []Bucket `json:"buckets"`
}
