package artifacts

type Artifact struct {
	Name  string
	S3Key string
}

type Database interface {
	Find(id string) (Artifact, error)
}

func Query(id string, database Database) (Artifact, error) {
	return database.Find(id)
}
