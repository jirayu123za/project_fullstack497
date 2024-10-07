package repositories

import "mime/multipart"

type MinIORepository interface {
	SaveFileToMinIO(file multipart.File, userGroupName, userName, fileName string) error
	//SaveFileToMinIO(file multipart.File, userGroupName, userName, fileExtension string) error
}
